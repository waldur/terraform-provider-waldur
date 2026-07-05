package e2e_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/waldur/terraform-provider-waldur/internal/provider"
	"github.com/waldur/terraform-provider-waldur/internal/testhelpers"
)

func TestOpenstackSubnet_CRUD(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Skipping acceptance test")
	}

	rec, cleanup := testhelpers.SetupVCR(t, "openstack_subnet_crud")
	defer cleanup()

	httpClient := &http.Client{Transport: rec}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"waldur": providerserver.NewProtocol6WithError(
				provider.NewWithHTTPClient("test", httpClient)(),
			),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccOpenstackSubnetConfig_basic("test-subnet", "192.168.84.0/24"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("waldur_openstack_subnet.test", "name", "test-subnet"),
					resource.TestCheckResourceAttr("waldur_openstack_subnet.test", "cidr", "192.168.84.0/24"),
					resource.TestCheckResourceAttr("waldur_openstack_subnet.test", "allocation_pools.#", "1"),
					resource.TestCheckResourceAttr("waldur_openstack_subnet.test", "allocation_pools.0.start", "192.168.84.10"),
					resource.TestCheckResourceAttr("waldur_openstack_subnet.test", "allocation_pools.0.end", "192.168.84.200"),
				),
			},
			{
				Config: testAccOpenstackSubnetConfig_basic("test-subnet-updated", "192.168.84.0/24"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("waldur_openstack_subnet.test", "name", "test-subnet-updated"),
				),
			},
		},
	})
}

func testAccOpenstackSubnetConfig_basic(name, cidr string) string {
	return testhelpers.GetProviderConfig() + fmt.Sprintf(`

data "waldur_structure_project" "test" {
  filters = {
    name_exact = "Naked azanide as an aminating agent and a superbase"
  }
}

data "waldur_openstack_tenant" "test" {
  filters = {
    name = "test-vpc-1"
    project_uuid = data.waldur_structure_project.test.id
  }
}

resource "waldur_openstack_network" "test" {
  name   = "test-network-for-subnet"
  tenant = data.waldur_openstack_tenant.test.url
}

resource "waldur_openstack_subnet" "test" {
  name    = %q
  network = waldur_openstack_network.test.url
  cidr    = %q
  allocation_pools = [
    {
      start = "192.168.84.10"
      end   = "192.168.84.200"
    }
  ]
}
`, name, cidr)
}
