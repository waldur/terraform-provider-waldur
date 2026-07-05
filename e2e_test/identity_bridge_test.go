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

func TestIdentityBridgeLink_CRUD(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Skipping acceptance test")
	}

	rec, cleanup := testhelpers.SetupVCR(t, "identity_bridge_link_crud")
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
				Config: testAccIdentityBridgeLinkConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("waldur_identity_bridge_link.test", "username", "tf-test-cuid@example.com"),
					resource.TestCheckResourceAttr("waldur_identity_bridge_link.test", "source", "isd:puhuri"),
					resource.TestCheckResourceAttr("waldur_identity_bridge_link.test", "first_name", "TFTest"),
					resource.TestCheckResourceAttr("waldur_identity_bridge_link.test", "last_name", "User"),
					resource.TestCheckResourceAttr("waldur_identity_bridge_link.test", "email", "tftest@example.com"),
					resource.TestCheckResourceAttr("waldur_identity_bridge_link.test", "id", "isd:puhuri/tf-test-cuid@example.com"),
				),
			},
			{
				Config: testAccIdentityBridgeLinkConfig_updated(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("waldur_identity_bridge_link.test", "username", "tf-test-cuid@example.com"),
					resource.TestCheckResourceAttr("waldur_identity_bridge_link.test", "source", "isd:puhuri"),
					resource.TestCheckResourceAttr("waldur_identity_bridge_link.test", "first_name", "TFTestUpdated"),
					resource.TestCheckResourceAttr("waldur_identity_bridge_link.test", "last_name", "UserUpdated"),
					resource.TestCheckResourceAttr("waldur_identity_bridge_link.test", "email", "tftest_updated@example.com"),
				),
			},
		},
	})
}

func testAccIdentityBridgeLinkConfig_basic() string {
	return testhelpers.GetProviderConfig() + `
resource "waldur_identity_bridge_link" "test" {
  username   = "tf-test-cuid@example.com"
  source     = "isd:puhuri"
  first_name = "TFTest"
  last_name  = "User"
  email      = "tftest@example.com"
}
`
}

func testAccIdentityBridgeLinkConfig_updated() string {
	return testhelpers.GetProviderConfig() + `
resource "waldur_identity_bridge_link" "test" {
  username   = "tf-test-cuid@example.com"
  source     = "isd:puhuri"
  first_name = "TFTestUpdated"
  last_name  = "UserUpdated"
  email      = "tftest_updated@example.com"
}
`
}
