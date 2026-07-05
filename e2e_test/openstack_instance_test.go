package e2e_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/waldur/terraform-provider-waldur/internal/provider"
	"github.com/waldur/terraform-provider-waldur/internal/testhelpers"
)

func TestOpenstackInstance_CRUD(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Skipping acceptance test")
	}

	rec, cleanup := testhelpers.SetupVCR(t, "openstack_instance_crud")
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
				Config: testAccOpenstackInstanceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "name", "test-instance-v6"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "id"),
				),
			},
		},
	})
}

func testAccOpenstackInstanceConfig_basic() string {
	return testhelpers.GetProviderConfig() + `

data "waldur_structure_project" "test" {
  filters = {
    name_exact = "Naked azanide as an aminating agent and a superbase"
  }
}

data "waldur_marketplace_offering" "test" {
  filters = {
    name = "Virtual machine in test-vpc-1"
    project_uuid = data.waldur_structure_project.test.id
  }
}

data "waldur_openstack_flavor" "test" {
  filters = {
    name = "m1.small"
    tenant_uuid = data.waldur_marketplace_offering.test.scope_uuid
  }
}

data "waldur_openstack_image" "test" {
  filters = {
    name = "cirros"
    tenant_uuid = data.waldur_marketplace_offering.test.scope_uuid
  }
}

data "waldur_core_ssh_public_key" "test" {
  filters = {
    name = "my-ssh-key"
  }
}

data "waldur_openstack_subnet" "test" {
  filters = {
    tenant_uuid = data.waldur_marketplace_offering.test.scope_uuid
	name = "test"
  }
}

resource "waldur_openstack_instance" "test" {
  name    = "test-instance-v6"
  flavor  = data.waldur_openstack_flavor.test.url
  image   = data.waldur_openstack_image.test.url
  project = data.waldur_structure_project.test.url
  offering = data.waldur_marketplace_offering.test.url
  ssh_public_key = data.waldur_core_ssh_public_key.test.url
  system_volume_size = 1024
  data_volume_size = 1024
  ports = [
    {
       subnet = data.waldur_openstack_subnet.test.url
    }
  ]
}
`
}
