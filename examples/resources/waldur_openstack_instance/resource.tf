data "waldur_structure_project" "example" {
  filters = {
    name = "Example Project"
  }
}

data "waldur_marketplace_offering" "example" {
  filters = {
    name         = "Virtual machine in test-vpc-1"
    project_uuid = data.waldur_structure_project.example.id
  }
}

data "waldur_openstack_flavor" "example" {
  filters = {
    name        = "m1.small"
    tenant_uuid = data.waldur_marketplace_offering.example.scope_uuid
  }
}

data "waldur_openstack_image" "example" {
  filters = {
    name        = "cirros"
    tenant_uuid = data.waldur_marketplace_offering.example.scope_uuid
  }
}

data "waldur_core_ssh_public_key" "example" {
  filters = {
    name = "example-ssh-key"
  }
}

data "waldur_openstack_subnet" "example" {
  filters = {
    tenant_uuid = data.waldur_marketplace_offering.example.scope_uuid
    name        = "example-subnet"
  }
}

resource "waldur_openstack_instance" "example" {
  name               = "example-instance"
  flavor             = data.waldur_openstack_flavor.example.url
  image              = data.waldur_openstack_image.example.url
  project            = data.waldur_structure_project.example.url
  offering           = data.waldur_marketplace_offering.example.url
  ssh_public_key     = data.waldur_core_ssh_public_key.example.url
  system_volume_size = 1024
  data_volume_size   = 1024
  ports = [
    {
      subnet = data.waldur_openstack_subnet.example.url
    }
  ]
}
