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

data "waldur_openstack_network" "example" {
  filters = {
    name        = "example-network"
    tenant_uuid = data.waldur_openstack_tenant.example.id
  }
}

data "waldur_openstack_tenant" "target" {
  filters = {
    name         = "target-tenant"
    project_uuid = data.waldur_structure_project.example.id
  }
}

resource "waldur_openstack_network_rbac_policy" "example" {
  network       = data.waldur_openstack_network.example.url
  target_tenant = data.waldur_openstack_tenant.target.url
  policy_type   = "access_as_shared"
}
