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

func TestStructureCustomer_CRUD(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Skipping acceptance test")
	}

	rec, cleanup := testhelpers.SetupVCR(t, "structure_customer_crud")
	defer cleanup()

	httpClient := &http.Client{Transport: rec}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"waldur": providerserver.NewProtocol6WithError(
				provider.NewWithHTTPClient("test", httpClient)(),
			),
		},
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccStructureCustomerConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("waldur_structure_customer.test", "name", "tf-test-customer"),
					resource.TestCheckResourceAttrSet("waldur_structure_customer.test", "id"),
					resource.TestCheckResourceAttrSet("waldur_structure_customer.test", "url"),
				),
			},
			// Update testing
			{
				Config: testAccStructureCustomerConfig_updated(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("waldur_structure_customer.test", "name", "tf-test-customer-updated"),
					resource.TestCheckResourceAttr("waldur_structure_customer.test", "description", "Updated customer for Terraform testing"),
				),
			},
			// Delete testing is implicit when the test completes
		},
	})
}

func getProviderConfigCustomer() string {
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

func testAccStructureCustomerConfig_basic() string {
	return getProviderConfigCustomer() + `
resource "waldur_structure_customer" "test" {
  name = "tf-test-customer"
}
`
}

func testAccStructureCustomerConfig_updated() string {
	return getProviderConfigCustomer() + `
resource "waldur_structure_customer" "test" {
  name        = "tf-test-customer-updated"
  description = "Updated customer for Terraform testing"
}
`
}
