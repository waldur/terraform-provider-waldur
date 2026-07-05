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

func TestOpenstackSecurityGroup_CRUD(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Skipping acceptance test")
	}

	rec, cleanup := testhelpers.SetupVCR(t, "openstack_security_group_crud")
	defer cleanup()

	httpClient := &http.Client{Transport: rec}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"waldur": providerserver.NewProtocol6WithError(
				provider.NewWithHTTPClient("test", httpClient)(),
			),
		},
		Steps: []resource.TestStep{
			// Step 1: Create security group
			{
				Config: testAccOpenstackSecurityGroupConfig_basic("test-sg", "test security group"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("waldur_openstack_security_group.test", "name", "test-sg"),
					resource.TestCheckResourceAttr("waldur_openstack_security_group.test", "description", "test security group"),
					resource.TestCheckResourceAttrSet("waldur_openstack_security_group.test", "id"),
				),
			},
			// Step 2: Update rules
			{
				Config: testAccOpenstackSecurityGroupConfig_withRules("test-sg", "test security group"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("waldur_openstack_security_group.test", "rules.#", "2"),
					resource.TestCheckResourceAttr("waldur_openstack_security_group.test", "rules.0.from_port", "80"),
					resource.TestCheckResourceAttr("waldur_openstack_security_group.test", "rules.0.to_port", "80"),
					resource.TestCheckResourceAttr("waldur_openstack_security_group.test", "rules.1.from_port", "443"),
					resource.TestCheckResourceAttr("waldur_openstack_security_group.test", "rules.1.to_port", "443"),
				),
			},
		},
	})
}

func testAccOpenstackSecurityGroupConfig_setup() string {
	return testhelpers.GetProviderConfig() + `
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
`
}

func testAccOpenstackSecurityGroupConfig_basic(name, description string) string {
	return testAccOpenstackSecurityGroupConfig_setup() + fmt.Sprintf(`
resource "waldur_openstack_security_group" "test" {
  name        = %q
  description = %q
  tenant      = data.waldur_openstack_tenant.test.url
  rules       = []
}
`, name, description)
}

func testAccOpenstackSecurityGroupConfig_withRules(name, description string) string {
	return testAccOpenstackSecurityGroupConfig_setup() + fmt.Sprintf(`
resource "waldur_openstack_security_group" "test" {
  name        = %q
  description = %q
  tenant      = data.waldur_openstack_tenant.test.url

  rules = [
    {
      protocol  = "tcp"
      from_port = 80
      to_port   = 80
      ethertype = "IPv4"
      direction = "ingress"
    },
    {
      protocol  = "tcp"
      from_port = 443
      to_port   = 443
      ethertype = "IPv4"
      direction = "ingress"
    }
  ]
}
`, name, description)
}
