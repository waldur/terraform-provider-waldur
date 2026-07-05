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

func TestCoreSshPublicKey_CRUD(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Skipping acceptance test")
	}

	rec, cleanup := testhelpers.SetupVCR(t, "ssh_public_key_crud")
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
				Config: testAccCoreSshPublicKeyConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("waldur_core_ssh_public_key.test", "name", "tf-test-key"),
					resource.TestCheckResourceAttrSet("waldur_core_ssh_public_key.test", "id"),
					resource.TestCheckResourceAttrSet("waldur_core_ssh_public_key.test", "public_key"),
				),
			},
		},
	})
}

func testAccCoreSshPublicKeyConfig_basic() string {
	return testhelpers.GetProviderConfig() + `
resource "waldur_core_ssh_public_key" "test" {
  name       = "tf-test-key"
  public_key = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAINznqmc6e9s1IC5MidoIc1bwAL/HHagG1Hc3yGNnpleT tf-test-key"
}
`
}
