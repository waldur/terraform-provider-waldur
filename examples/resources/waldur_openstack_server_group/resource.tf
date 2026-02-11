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

resource "waldur_openstack_server_group" "example" {
  name        = "example-server-group"
  description = "Anti-affinity server group for HA workloads"
  policy      = "anti-affinity"
  tenant      = data.waldur_openstack_tenant.example.url
}
