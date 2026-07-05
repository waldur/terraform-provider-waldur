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
			// Delete testing is implicit when the test completes
		},
	})
}

func testAccStructureCustomerConfig_basic() string {
	return testhelpers.GetProviderConfig() + `
resource "waldur_structure_customer" "test" {
  name = "tf-test-customer"
}
`
}
