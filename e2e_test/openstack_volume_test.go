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

func TestOpenstackVolume_Create(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Skipping acceptance test")
	}

	rec, cleanup := testhelpers.SetupVCR(t, "openstack_volume_create")
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
				Config: testAccOpenstackVolumeConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("waldur_openstack_volume.test", "name", "tf-test-volume-v11"),
					resource.TestCheckResourceAttr("waldur_openstack_volume.test", "size", "1024"),
					resource.TestCheckResourceAttrSet("waldur_openstack_volume.test", "id"),
					resource.TestCheckResourceAttrSet("waldur_openstack_volume.test", "project"),
					resource.TestCheckResourceAttrSet("waldur_openstack_volume.test", "offering"),
				),
			},
		},
	})
}

func testAccOpenstackVolumeConfig_basic() string {
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

resource "waldur_openstack_volume" "test" {
  name     = "tf-test-volume-v11"
  project  = data.waldur_structure_project.test.url
  offering = data.waldur_marketplace_offering.test.url
  size     = 1024
}
`
}
