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

// TestOpenstackNetworkRbacPolicy_CRUD tests CRUD operations for OpenStack Network RBAC Policy.
// This test uses hardcoded URLs for existing network and tenant resources in the test environment.
// The test creates an RBAC policy to share the network with the target tenant,
// then updates the policy type and finally deletes it.
func TestOpenstackNetworkRbacPolicy_CRUD(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Skipping acceptance test")
	}

	rec, cleanup := testhelpers.SetupVCR(t, "openstack_network_rbac_policy_crud")
	defer cleanup()

	httpClient := &http.Client{Transport: rec}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"waldur": providerserver.NewProtocol6WithError(
				provider.NewWithHTTPClient("test", httpClient)(),
			),
		},
		Steps: []resource.TestStep{
			// Create and Read testing - create RBAC policy with access_as_shared
			{
				Config: testAccOpenstackNetworkRbacPolicyConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("waldur_openstack_network_rbac_policy.test", "id"),
					resource.TestCheckResourceAttrSet("waldur_openstack_network_rbac_policy.test", "url"),
					resource.TestCheckResourceAttrSet("waldur_openstack_network_rbac_policy.test", "backend_id"),
					resource.TestCheckResourceAttr("waldur_openstack_network_rbac_policy.test", "policy_type", "access_as_shared"),
				),
			},
			// Update testing - change policy type to access_as_external
			{
				Config: testAccOpenstackNetworkRbacPolicyConfig_updated(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("waldur_openstack_network_rbac_policy.test", "id"),
					resource.TestCheckResourceAttr("waldur_openstack_network_rbac_policy.test", "policy_type", "access_as_external"),
				),
			},
			// Delete testing is implicit when the test completes
		},
	})
}

func getProviderConfigRbacPolicy() string {
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

func testAccOpenstackNetworkRbacPolicyConfig_basic() string {
	return getProviderConfigRbacPolicy() + `
data "waldur_openstack_network" "net" {
  name = "91809-int-net"
}

data "waldur_structure_customer" "cust" {
  name = "UT-DevOpsCourse"
}

data "waldur_structure_project" "proj" {
  name          = "921091"
  customer      = data.waldur_structure_customer.cust.id
}

data "waldur_openstack_tenant" "tenant" {
  name         = "921091"
  project_uuid = data.waldur_structure_project.proj.id
}

resource "waldur_openstack_network_rbac_policy" "test" {
  network       = data.waldur_openstack_network.net.url
  target_tenant = data.waldur_openstack_tenant.tenant.url
  policy_type   = "access_as_shared"
}
`
}

func testAccOpenstackNetworkRbacPolicyConfig_updated() string {
	return getProviderConfigRbacPolicy() + `
data "waldur_openstack_network" "net" {
  name = "91809-int-net"
}

data "waldur_structure_customer" "cust" {
  name = "UT-DevOpsCourse"
}

data "waldur_structure_project" "proj" {
  name          = "921091"
  customer      = data.waldur_structure_customer.cust.id
}

data "waldur_openstack_tenant" "tenant" {
  name         = "921091"
  project_uuid = data.waldur_structure_project.proj.id
}

resource "waldur_openstack_network_rbac_policy" "test" {
  network       = data.waldur_openstack_network.net.url
  target_tenant = data.waldur_openstack_tenant.tenant.url
  policy_type   = "access_as_external"
}
`
}
