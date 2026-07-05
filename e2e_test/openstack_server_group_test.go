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

func TestOpenstackServerGroup_CRUD(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Skipping acceptance test")
	}

	rec, cleanup := testhelpers.SetupVCR(t, "openstack_server_group_crud")
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
				Config: testAccOpenstackServerGroupConfig("test-server-group", "affinity"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("waldur_openstack_server_group.test", "name", "test-server-group"),
					resource.TestCheckResourceAttr("waldur_openstack_server_group.test", "policy", "affinity"),
					resource.TestCheckResourceAttrSet("waldur_openstack_server_group.test", "id"),
				),
			},
		},
	})
}

func testAccOpenstackServerGroupConfig(name, policy string) string {
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

resource "waldur_openstack_server_group" "test" {
  name   = %q
  policy = %q
  tenant = data.waldur_openstack_tenant.test.url
}
`, name, policy)
}
