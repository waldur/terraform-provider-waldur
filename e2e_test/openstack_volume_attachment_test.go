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

func TestOpenstackVolumeAttachment_CRUD(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Skipping acceptance test")
	}

	rec, cleanup := testhelpers.SetupVCR(t, "openstack_volume_attachment_crud")
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
				Config: testAccOpenstackVolumeAttachmentConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("waldur_openstack_volume_attachment.test", "volume", "waldur_openstack_volume.volume", "id"),
					resource.TestCheckResourceAttrPair("waldur_openstack_volume_attachment.test", "instance", "waldur_openstack_instance.test", "url"),
					resource.TestCheckResourceAttrSet("waldur_openstack_volume_attachment.test", "id"),
				),
			},
		},
	})
}

func testAccOpenstackVolumeAttachmentConfig_basic() string {
	return testhelpers.GetProviderConfig() + `

data "waldur_structure_project" "test" {
  filters = {
    name_exact = "Naked azanide as an aminating agent and a superbase"
  }
}

data "waldur_marketplace_offering" "instance" {
  filters = {
    name = "Virtual machine in test-vpc-1"
    project_uuid = data.waldur_structure_project.test.id
  }
}

data "waldur_marketplace_offering" "volume" {
  filters = {
    name = "Volume in test-vpc-1"
    project_uuid = data.waldur_structure_project.test.id
  }
}

// Ensure both instance and volume offerings are in the same tenant scope
data "waldur_openstack_flavor" "test" {
  filters = {
    name = "m1.small"
    tenant_uuid = data.waldur_marketplace_offering.instance.scope_uuid
  }
}

data "waldur_openstack_image" "test" {
  filters = {
    name = "cirros"
    tenant_uuid = data.waldur_marketplace_offering.instance.scope_uuid
  }
}

data "waldur_core_ssh_public_key" "test" {
  filters = {
    name = "my-ssh-key"
  }
}

data "waldur_openstack_subnet" "test" {
  filters = {
    tenant_uuid = data.waldur_marketplace_offering.instance.scope_uuid
    name = "test"
  }
}

resource "waldur_openstack_instance" "test" {
  name    = "tf-test-instance-attach-v1"
  flavor  = data.waldur_openstack_flavor.test.url
  image   = data.waldur_openstack_image.test.url
  project = data.waldur_structure_project.test.url
  offering = data.waldur_marketplace_offering.instance.url
  ssh_public_key = data.waldur_core_ssh_public_key.test.url
  system_volume_size = 1024
  data_volume_size = 1024
  ports = [
    {
       subnet = data.waldur_openstack_subnet.test.url
    }
  ]
}

resource "waldur_openstack_volume" "volume" {
  name     = "tf-test-volume-attach-v1"
  project  = data.waldur_structure_project.test.url
  offering = data.waldur_marketplace_offering.volume.url
  size     = 1024
}

resource "waldur_openstack_volume_attachment" "test" {
  volume   = waldur_openstack_volume.volume.id
  instance = waldur_openstack_instance.test.url
}
`
}
