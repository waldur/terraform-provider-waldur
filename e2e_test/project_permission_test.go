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

func TestProjectPermission_CRUD(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Skipping acceptance test")
	}

	rec, cleanup := testhelpers.SetupVCR(t, "project_permission_crud")
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
				Config: testAccProjectPermissionConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("waldur_project_permission.test", "id"),
					resource.TestCheckResourceAttr("waldur_project_permission.test", "role", "PROJECT.ADMIN"),
					resource.TestCheckResourceAttr("waldur_project_permission.test", "user", "e2e00000000000000000000000000001"),
				),
			},
			{
				Config: testAccProjectPermissionConfig_updated(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("waldur_project_permission.test", "id"),
					resource.TestCheckResourceAttr("waldur_project_permission.test", "role", "PROJECT.ADMIN"),
					resource.TestCheckResourceAttr("waldur_project_permission.test", "user", "e2e00000000000000000000000000001"),
					resource.TestCheckResourceAttr("waldur_project_permission.test", "expiration_time", "2030-01-01T00:00:00Z"),
				),
			},
		},
	})
}

func testAccProjectPermissionConfig_basic() string {
	return testhelpers.GetProviderConfig() + `
resource "waldur_structure_customer" "test" {
  name = "tf-test-customer-for-perm"
}

resource "waldur_structure_project" "test" {
  name     = "tf-test-project-for-perm"
  customer = waldur_structure_customer.test.url
}

resource "waldur_project_permission" "test" {
  project = waldur_structure_project.test.id
  user    = "e2e00000000000000000000000000001"
  role    = "PROJECT.ADMIN"
}
`
}

func testAccProjectPermissionConfig_updated() string {
	return testhelpers.GetProviderConfig() + `
resource "waldur_structure_customer" "test" {
  name = "tf-test-customer-for-perm"
}

resource "waldur_structure_project" "test" {
  name     = "tf-test-project-for-perm"
  customer = waldur_structure_customer.test.url
}

resource "waldur_project_permission" "test" {
  project         = waldur_structure_project.test.id
  user            = "e2e00000000000000000000000000001"
  role            = "PROJECT.ADMIN"
  expiration_time = "2030-01-01T00:00:00Z"
}
`
}
