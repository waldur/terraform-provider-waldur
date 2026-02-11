data "waldur_structure_project" "example" {
  filters = {
    name = "Example Project"
  }
}

data "waldur_openstack_tenant" "example" {
  filters = {
    name         = "example-tenant"
    project_uuid = data.waldur_structure_project.example.id
  }
}

resource "waldur_openstack_floating_ip" "example" {
  tenant = data.waldur_openstack_tenant.example.url
}
