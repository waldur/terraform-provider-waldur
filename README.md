# Terraform Provider Waldur

A Terraform provider for managing resources via the [Waldur](https://waldur.com/) API.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.24 (to build the provider plugin)

## Installation

### From Terraform Registry

```hcl
terraform {
  required_providers {
    waldur = {
      source = "waldur/waldur"
    }
  }
}
```

## Provider Configuration

```hcl
provider "waldur" {
  endpoint = "https://your-waldur-instance.com/api/"
  token    = "your-api-token"
}
```

| Argument | Description | Required |
|----------|-------------|----------|
| `endpoint` | The Waldur API endpoint URL | Yes |
| `token` | API authentication token | Yes |

## Resources

| Resource | Description |
|----------|-------------|
| `waldur_structure_project` | Manages a Project |
| `waldur_structure_customer` | Manages a Customer |
| `waldur_marketplace_order` | Manages a Order |
| `waldur_marketplace_offering` | Manages a Offering |
| `waldur_marketplace_resource` | Manages a Resource |
| `waldur_openstack_tenant` | Manages a Tenant |
| `waldur_openstack_volume` | Manages a Volume |
| `waldur_openstack_volume_attachment` | Manages a VolumeAttachment |
| `waldur_openstack_instance` | Manages a Instance |
| `waldur_openstack_security_group` | Manages a SecurityGroup |
| `waldur_openstack_network` | Manages a Network |
| `waldur_openstack_server_group` | Manages a ServerGroup |
| `waldur_openstack_subnet` | Manages a Subnet |
| `waldur_openstack_network_rbac_policy` | Manages a NetworkRbacPolicy |
| `waldur_openstack_floating_ip` | Manages a FloatingIp |
| `waldur_openstack_port` | Manages a Port |

## Actions

| Action | Resource | Description |
|--------|----------|-------------|
| `waldur_marketplace_resource_pull` | `waldur_marketplace_resource` | Pull action |
| `waldur_marketplace_resource_terminate` | `waldur_marketplace_resource` | Terminate action |
| `waldur_marketplace_resource_unlink` | `waldur_marketplace_resource` | Unlink action |
| `waldur_openstack_tenant_pull` | `waldur_openstack_tenant` | Pull action |
| `waldur_openstack_tenant_unlink` | `waldur_openstack_tenant` | Unlink action |
| `waldur_openstack_volume_pull` | `waldur_openstack_volume` | Pull action |
| `waldur_openstack_volume_unlink` | `waldur_openstack_volume` | Unlink action |
| `waldur_openstack_instance_start` | `waldur_openstack_instance` | Start action |
| `waldur_openstack_instance_stop` | `waldur_openstack_instance` | Stop action |
| `waldur_openstack_instance_restart` | `waldur_openstack_instance` | Restart action |
| `waldur_openstack_instance_pull` | `waldur_openstack_instance` | Pull action |
| `waldur_openstack_instance_unlink` | `waldur_openstack_instance` | Unlink action |
| `waldur_openstack_network_pull` | `waldur_openstack_network` | Pull action |
| `waldur_openstack_network_unlink` | `waldur_openstack_network` | Unlink action |
| `waldur_openstack_subnet_connect` | `waldur_openstack_subnet` | Connect action |
| `waldur_openstack_subnet_disconnect` | `waldur_openstack_subnet` | Disconnect action |
| `waldur_openstack_subnet_pull` | `waldur_openstack_subnet` | Pull action |
| `waldur_openstack_subnet_unlink` | `waldur_openstack_subnet` | Unlink action |
| `waldur_openstack_port_enable_port` | `waldur_openstack_port` | EnablePort action |
| `waldur_openstack_port_disable_port` | `waldur_openstack_port` | DisablePort action |
| `waldur_openstack_port_enable_port_security` | `waldur_openstack_port` | EnablePortSecurity action |
| `waldur_openstack_port_disable_port_security` | `waldur_openstack_port` | DisablePortSecurity action |
| `waldur_openstack_port_pull` | `waldur_openstack_port` | Pull action |
| `waldur_openstack_port_unlink` | `waldur_openstack_port` | Unlink action |

## Data Sources

| Data Source | Description |
|-------------|-------------|
| `waldur_structure_project` | Retrieves Project data |
| `waldur_structure_customer` | Retrieves Customer data |
| `waldur_core_ssh_public_key` | Retrieves SshPublicKey data |
| `waldur_marketplace_offering` | Retrieves Offering data |
| `waldur_marketplace_order` | Retrieves Order data |
| `waldur_marketplace_resource` | Retrieves Resource data |
| `waldur_openstack_tenant` | Retrieves Tenant data |
| `waldur_openstack_volume` | Retrieves Volume data |
| `waldur_openstack_instance` | Retrieves Instance data |
| `waldur_openstack_security_group` | Retrieves SecurityGroup data |
| `waldur_openstack_subnet` | Retrieves Subnet data |
| `waldur_openstack_network` | Retrieves Network data |
| `waldur_openstack_flavor` | Retrieves Flavor data |
| `waldur_openstack_image` | Retrieves Image data |
| `waldur_openstack_volume_type` | Retrieves VolumeType data |
| `waldur_openstack_port` | Retrieves Port data |

## Documentation

For detailed documentation on each resource and data source, please refer to the
[Terraform Registry documentation](https://registry.terraform.io/providers/waldur/waldur/latest/docs).

## Links

- [Waldur](https://waldur.com/)
- [Terraform Registry](https://registry.terraform.io/)
