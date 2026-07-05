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

func TestOpenstackTenant_Create(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Skipping acceptance test")
	}

	rec, cleanup := testhelpers.SetupVCR(t, "openstack_tenant_create")
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
				Config: testAccOpenstackTenantConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("waldur_openstack_tenant.test", "name", "tf-test-tenant-v11"),
					resource.TestCheckResourceAttrSet("waldur_openstack_tenant.test", "id"),
					resource.TestCheckResourceAttrSet("waldur_openstack_tenant.test", "project"),
					resource.TestCheckResourceAttrSet("waldur_openstack_tenant.test", "offering"),
				),
			},
		},
	})
}

func testAccOpenstackTenantConfig_basic() string {
	return testhelpers.GetProviderConfig() + `
data "waldur_structure_project" "test" {
  filters = {
    name_exact = "Naked azanide as an aminating agent and a superbase"
  }
}

data "waldur_marketplace_offering" "test" {
  filters = {
    name = "MockOpenStack5"
    project_uuid = data.waldur_structure_project.test.id
  }
}

resource "waldur_openstack_tenant" "test" {
  name     = "tf-test-tenant-v11"
  project  = data.waldur_structure_project.test.url
  offering = data.waldur_marketplace_offering.test.url
  plan     = data.waldur_marketplace_offering.test.plans[0].url
  limits = {
    cores   = 10
    ram     = 10240
    storage = 1024000
  }
  skip_connection_extnet          = true
  skip_creation_of_default_router = true
}
`
}
