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

func TestOpenstackNetworkRbacPolicy_CRUD(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Skipping acceptance test")
	}

	rec, cleanup := testhelpers.SetupVCR(t, "openstack_network_rbac_policy_crud")
	defer cleanup()

	httpClient := &http.Client{Transport: rec}

	networkName := "test-network-rbac"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"waldur": providerserver.NewProtocol6WithError(
				provider.NewWithHTTPClient("test", httpClient)(),
			),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccOpenstackNetworkRbacPolicyConfig_basic(networkName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("waldur_openstack_network_rbac_policy.test", "network_name", networkName),
					resource.TestCheckResourceAttr("waldur_openstack_network_rbac_policy.test", "policy_type", "access_as_shared"),
				),
			},
		},
	})
}

func testAccOpenstackNetworkRbacPolicyConfig_basic(networkName string) string {
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
  name   = %q
  tenant = data.waldur_openstack_tenant.test.url
}

resource "waldur_openstack_network_rbac_policy" "test" {
  network       = waldur_openstack_network.test.url
  target_tenant = data.waldur_openstack_tenant.test.url
  policy_type   = "access_as_shared"
}
`, networkName)
}
