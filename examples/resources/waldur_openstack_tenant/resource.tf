resource "waldur_openstack_tenant" "example" {
  name     = "example-tenant"
  project  = data.waldur_structure_project.project.url
  offering = data.waldur_marketplace_offering.offering.url
  plan     = data.waldur_marketplace_offering.offering.plans[0].url
  
  limits = {
    cores   = 10
    ram     = 10240
    storage = 100
  }

  skip_connection_extnet          = true
  skip_creation_of_default_router = true
}

data "waldur_structure_project" "project" {
  filters = {
    name_exact = "Project Name"
  }
}

data "waldur_marketplace_offering" "offering" {
  filters = {
    name = "OpenStack Offering"
  }
}
