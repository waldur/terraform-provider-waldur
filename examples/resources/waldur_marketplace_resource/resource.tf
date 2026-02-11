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

resource "waldur_marketplace_resource" "example" {
  name     = "example-resource"
  offering = data.waldur_marketplace_offering.example.url
  plan     = data.waldur_marketplace_offering.example.plans[0].url
  project  = data.waldur_structure_project.example.url

  limits = {
    cores   = 4
    ram     = 8192
    storage = 102400
  }
}
