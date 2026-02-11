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

resource "waldur_openstack_port" "example" {
  name    = "example-port"
  network = data.waldur_openstack_network.example.url

  fixed_ips = [
    {
      ip_address = "10.0.1.50"
      subnet_id  = "subnet-backend-id"
    },
  ]

  security_groups = [
    {
      name = "default"
    },
  ]
}
