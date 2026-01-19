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

func TestStructureProject_CRUD(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Skipping acceptance test")
	}

	rec, cleanup := testhelpers.SetupVCR(t, "structure_project_crud")
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
				Config: testAccStructureProjectConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("waldur_structure_project.test", "name", "tf-test-project"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "id"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "url"),
					resource.TestCheckResourceAttrSet("waldur_structure_project.test", "customer"),
				),
			},
			// Update testing
			{
				Config: testAccStructureProjectConfig_updated(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("waldur_structure_project.test", "name", "tf-test-project-updated"),
					resource.TestCheckResourceAttr("waldur_structure_project.test", "description", "Updated project for Terraform testing"),
				),
			},
			// Delete testing is implicit when the test completes
		},
	})
}

func getProviderConfigProject() string {
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

func testAccStructureProjectConfig_basic() string {
	return getProviderConfigProject() + `
resource "waldur_structure_customer" "test" {
  name = "tf-test-customer-for-project"
}

resource "waldur_structure_project" "test" {
  name     = "tf-test-project"
  customer = waldur_structure_customer.test.url
}
`
}

func testAccStructureProjectConfig_updated() string {
	return getProviderConfigProject() + `
resource "waldur_structure_customer" "test" {
  name = "tf-test-customer-for-project"
}

resource "waldur_structure_project" "test" {
  name        = "tf-test-project-updated"
  customer    = waldur_structure_customer.test.url
  description = "Updated project for Terraform testing"
}
`
}
