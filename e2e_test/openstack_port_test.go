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

func TestOpenstackPort_CRUD(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Skipping acceptance test")
	}

	rec, cleanup := testhelpers.SetupVCR(t, "openstack_port_crud")
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
				Config: testAccOpenstackPortConfig("test-port", "test port description"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("waldur_openstack_port.test", "name", "test-port"),
					resource.TestCheckResourceAttr("waldur_openstack_port.test", "description", "test port description"),
					resource.TestCheckResourceAttrSet("waldur_openstack_port.test", "id"),
					resource.TestCheckResourceAttrSet("waldur_openstack_port.test", "network"),
				),
			},
			{
				Config: testAccOpenstackPortConfig("test-port-updated", "updated description"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("waldur_openstack_port.test", "name", "test-port-updated"),
					resource.TestCheckResourceAttr("waldur_openstack_port.test", "description", "updated description"),
				),
			},
		},
	})
}

func testAccOpenstackPortConfig(name, description string) string {
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
  name   = "test-network-for-port"
  tenant = data.waldur_openstack_tenant.test.url
}

resource "waldur_openstack_port" "test" {
  name        = %q
  description = %q
  network     = waldur_openstack_network.test.url
}
`, name, description)
}
