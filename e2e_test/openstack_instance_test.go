package e2e_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/waldur/terraform-waldur-provider/internal/provider"
	"github.com/waldur/terraform-waldur-provider/internal/testhelpers"
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
					resource.TestCheckResourceAttr("waldur_openstack_instance.test", "name", "test-instance"),
					resource.TestCheckResourceAttrSet("waldur_openstack_instance.test", "id"),
				),
			},
		},
	})
}

func getProviderConfig() string {
	endpoint := os.Getenv("WALDUR_API_URL")
	if endpoint == "" {
		endpoint = "https://api.waldur.example.com"
	}
	token := os.Getenv("WALDUR_ACCESS_TOKEN")
	if token == "" {
		token = "test-token-sanitized"
	}
	return fmt.Sprintf(`provider "waldur" {
  endpoint = %q
  token    = %q
}
`, endpoint, token)
}

func testAccOpenstackInstanceConfig_basic() string {
	return getProviderConfig() + `

data "waldur_structure_project" "test" {
  name = "waldur-test"
}

data "waldur_marketplace_offering" "test" {
  name = "Virtual machine in waldur-dev"
  project_uuid = data.waldur_structure_project.test.id
}

data "waldur_openstack_flavor" "test" {
  name = "m1.small"
  tenant_uuid = data.waldur_marketplace_offering.test.scope_uuid
}

data "waldur_openstack_image" "test" {
  name = "cirros"
  tenant_uuid = data.waldur_marketplace_offering.test.scope_uuid
}

data "waldur_core_ssh_public_key" "test" {
  name = "my-ssh-key"
}

data "waldur_openstack_port" "test" {
  backend_id = "31771cfe-dc77-417a-bacb-460bf8fa0791"
}

resource "waldur_openstack_instance" "test" {
  name    = "test-instance"
  flavor  = data.waldur_openstack_flavor.test.url
  image   = data.waldur_openstack_image.test.url
  project = data.waldur_structure_project.test.url
  offering = data.waldur_marketplace_offering.test.url
  ssh_public_key = data.waldur_core_ssh_public_key.test.url
  system_volume_size = 1024
  data_volume_size = 1024
  ports = [
    {
       port = data.waldur_openstack_port.test.url
    }
  ]
}
`
}
