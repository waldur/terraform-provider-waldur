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

resource "waldur_openstack_security_group" "example" {
  name        = "example-security-group"
  description = "Allow SSH and HTTP traffic"
  tenant      = data.waldur_openstack_tenant.example.url

  rules = [
    {
      direction = "ingress"
      protocol  = "tcp"
      from_port = 22
      to_port   = 22
      cidr      = "0.0.0.0/0"
      ethertype = "IPv4"
    },
    {
      direction = "ingress"
      protocol  = "tcp"
      from_port = 80
      to_port   = 80
      cidr      = "0.0.0.0/0"
      ethertype = "IPv4"
    },
  ]
}
