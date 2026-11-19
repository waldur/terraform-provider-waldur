# Reveal one access key. The secret is fetched on every run and is never
# written to the plan or the state, so it can only be passed on to arguments
# that accept ephemeral values: provider configuration, provisioners, other
# ephemeral resources and write-only attributes.
ephemeral "waldur_marketplace_resource_api_key" "s3" {
  id = data.waldur_marketplace_resource_api_keys.s3.items[0].id
}

# Rewritten on every apply: a rotated key keeps its id, so only a fresh
# reveal picks up its new secret.
resource "terraform_data" "credentials_file" {
  triggers_replace = [plantimestamp()]

  provisioner "local-exec" {
    command = <<-EOT
      umask 077
      printf 'export AWS_ACCESS_KEY_ID=%s\nexport AWS_SECRET_ACCESS_KEY=%s\n' "$KEY_ID" "$SECRET" > s3.env
    EOT
    environment = {
      KEY_ID = ephemeral.waldur_marketplace_resource_api_key.s3.client_id
      SECRET = ephemeral.waldur_marketplace_resource_api_key.s3.api_key
    }
  }
}
