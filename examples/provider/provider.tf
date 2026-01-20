terraform {
  required_providers {
    waldur = {
      source = "waldur/waldur"
    }
  }
}

# Option 1: Configure explicitly
provider "waldur" {
  endpoint = "https://waldur.example.com"
  token    = "your-api-token"
}

# Option 2: Use environment variables
# export WALDUR_API_URL="https://waldur.example.com"
# export WALDUR_ACCESS_TOKEN="your-api-token"
# 
# provider "waldur" {}
