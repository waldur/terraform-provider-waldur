resource "waldur_openstack_volume" "example" {
  name     = "example-volume"
  project  = data.waldur_structure_project.project.url
  offering = data.waldur_marketplace_offering.offering.url
  size     = 1024
}

data "waldur_structure_project" "project" {
  filters = {
    name_exact = "Project Name"
  }
}

data "waldur_marketplace_offering" "offering" {
  filters = {
    name = "Volume Offering"
  }
}
