resource "waldur_structure_customer" "example" {
  name = "example-customer-for-project"
}

resource "waldur_structure_project" "example" {
  name     = "example-project"
  customer = waldur_structure_customer.example.url
}
