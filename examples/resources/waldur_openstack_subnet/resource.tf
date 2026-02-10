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

resource "waldur_openstack_subnet" "example" {
  name            = "example-subnet"
  tenant          = data.waldur_openstack_tenant.example.url
  network         = data.waldur_openstack_network.example.url
  cidr            = "10.0.1.0/24"
  gateway_ip      = "10.0.1.1"
  allocation_pools = [
    {
      start = "10.0.1.10"
      end   = "10.0.1.100"
    }
  ]
}
