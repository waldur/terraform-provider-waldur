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

func TestMarketplaceOrder_ProvisionVolume(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Skipping acceptance test")
	}

	rec, cleanup := testhelpers.SetupVCR(t, "marketplace_order_provision_volume")
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
				Config: testAccMarketplaceOrderConfig_provisionVolume(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("waldur_marketplace_order.test", "resource_name", "tf-test-volume-provision"),
					resource.TestCheckResourceAttrSet("waldur_marketplace_order.test", "id"),
					resource.TestCheckResourceAttrSet("waldur_marketplace_order.test", "project"),
					resource.TestCheckResourceAttrSet("waldur_marketplace_order.test", "offering"),
				),
			},
		},
	})
}

func testAccMarketplaceOrderConfig_provisionVolume() string {
	return testhelpers.GetProviderConfig() + `
data "waldur_structure_project" "test" {
  filters = {
    name_exact = "Naked azanide as an aminating agent and a superbase"
  }
}

data "waldur_marketplace_offering" "test" {
  filters = {
    name = "Volume in tf-test-tenant-v5"
  }
}

resource "waldur_marketplace_order" "test" {
  project  = data.waldur_structure_project.test.url
  offering = data.waldur_marketplace_offering.test.url
  attributes = {
    name       = "tf-test-volume-provision"
    size       = "1024"
    description = "Provisioned via marketplace order E2E test"
  }
}
`
}
