data "waldur_structure_project" "example" {
  filters = {
    name = "Example Project"
  }
}

data "waldur_marketplace_offering" "example" {
  filters = {
    name = "Example Offering"
  }
}

resource "waldur_marketplace_order" "example" {
  offering = data.waldur_marketplace_offering.example.url
  project  = data.waldur_structure_project.example.url
  plan     = data.waldur_marketplace_offering.example.plans[0].url
}
