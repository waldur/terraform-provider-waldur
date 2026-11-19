# The access keys the site agent has issued for a resource
data "waldur_marketplace_resource_api_keys" "s3" {
  filters = {
    resource_uuid = waldur_marketplace_order.s3.marketplace_resource_uuid
    state         = "OK"
  }
}
