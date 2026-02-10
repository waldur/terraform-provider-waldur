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

resource "waldur_openstack_network" "example" {
  name   = "example-network"
  tenant = data.waldur_openstack_tenant.example.url
}
