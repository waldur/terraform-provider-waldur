# Changelog

## 2026-08-25

### ⚠️ Breaking changes

- `marketplace_offering`: removed attribute `scope_resource_name`
- `marketplace_offering`: removed attribute `scope_resource_uuid`
- `marketplace_offering`: removed attribute `scope_resource`
- `marketplace_offering`: removed attribute `user_has_offering_user`
- `structure_project`: removed attribute `display_credit_reports`

## 2026-08-24

### Added

- `marketplace_offering`: new attribute `user_has_offering_user`

## 2026-08-23

### Added

- `marketplace_offering`: new attribute `scope_resource_name`
- `marketplace_offering`: new attribute `scope_resource_uuid`
- `marketplace_offering`: new attribute `scope_resource`
- `structure_project`: new attribute `display_credit_reports`

## 2026-08-23

### ⚠️ Breaking changes

- `marketplace_offering`: removed attribute `scope_resource_name`
- `marketplace_offering`: removed attribute `scope_resource_uuid`
- `marketplace_offering`: removed attribute `scope_resource`
- `structure_project`: removed attribute `display_credit_reports`

## 2026-08-23

### Added

- `marketplace_offering`: new attribute `scope_resource_name`
- `marketplace_offering`: new attribute `scope_resource_uuid`
- `marketplace_offering`: new attribute `scope_resource`
- `structure_project`: new attribute `display_credit_reports`

## 2026-08-23

### ⚠️ Breaking changes

- `marketplace_offering`: removed attribute `scope_resource_name`
- `marketplace_offering`: removed attribute `scope_resource_uuid`
- `marketplace_offering`: removed attribute `scope_resource`
- `structure_project`: removed attribute `display_credit_reports`

## 2026-08-22

### Added

- `marketplace_offering`: new attribute `scope_resource_name`
- `marketplace_offering`: new attribute `scope_resource_uuid`
- `marketplace_offering`: new attribute `scope_resource`
- `structure_project`: new attribute `display_credit_reports`

## 2026-08-22

### ⚠️ Breaking changes

- `marketplace_offering`: removed attribute `scope_resource_name`
- `marketplace_offering`: removed attribute `scope_resource_uuid`
- `marketplace_offering`: removed attribute `scope_resource`
- `structure_project`: removed attribute `display_credit_reports`

## 2026-08-19

### Added

- `marketplace_offering`: new attribute `scope_resource_name`
- `marketplace_offering`: new attribute `scope_resource_uuid`
- `marketplace_offering`: new attribute `scope_resource`

## 2026-08-19

### Added

- `structure_project`: new attribute `display_credit_reports`

## 2026-08-12

### Added

- `marketplace_offering`: new attribute `plugin_options.show_ssh_key_loss_warning`

## 2026-08-11

### ⚠️ Breaking changes

- `marketplace_offering`: attribute `software_catalogs.enabled_cpu_family` type changed from `types.Map` to `types.List<string>`
- `marketplace_offering`: attribute `software_catalogs.enabled_cpu_microarchitectures` type changed from `types.Map` to `types.List<string>`

## 2026-08-10

### Added

- `marketplace_offering`: new attribute `plugin_options.enable_resource_end_date_change_requests`

## 2026-08-07

### Added

- `marketplace_order`: new attribute `consumer_message_updated_at`
- `marketplace_order`: new attribute `provider_message_updated_at`
- `marketplace_resource`: new attribute `order_in_progress.consumer_message_updated_at`
- `marketplace_resource`: new attribute `order_in_progress.provider_message_updated_at`

## 2026-08-06

### Added

- `marketplace_offering`: new attribute `open_for_proposals`

## 2026-08-03

### ⚠️ Breaking changes

- `marketplace_offering`: removed attribute `plans.components.discount_rate`
- `marketplace_offering`: removed attribute `plans.components.discount_threshold`
- `marketplace_offering`: removed attribute `plans.components.discounted_price`
- `marketplace_offering`: removed attribute `plugin_options.initial_primarygroup_number`
- `marketplace_offering`: removed attribute `plugin_options.initial_rolegroup_number`
- `marketplace_offering`: removed attribute `plugin_options.initial_uidnumber`
- `marketplace_offering`: removed attribute `plugin_options.initial_usergroup_number`
- `openstack_port`: attribute `allowed_address_pairs.ip_address` is now computed (read-only)
- `structure_customer`: attribute `customer_credit` type changed from `types.Float64` to `types.String`
- `structure_customer`: attribute `customer_unallocated_credit` type changed from `types.Float64` to `types.String`
- `structure_customer`: attribute `payment_profiles.attributes.contract_sum` type changed from `types.Int64` to `types.String`
- `structure_project`: removed attribute `project_metadata.answer`

### Added

- `marketplace_offering`: new attribute `default_access_subnets`
- `marketplace_offering`: new attribute `partitions.qos_options`
- `marketplace_offering`: new attribute `plans.components.discount_aggregation`
- `marketplace_offering`: new attribute `plans.components.discount_formula`
- `marketplace_offering`: new attribute `plugin_options.action_on_usage_limit`
- `marketplace_offering`: new attribute `plugin_options.auto_approve_for_roles`
- `marketplace_offering`: new attribute `plugin_options.billing_source`
- `marketplace_offering`: new attribute `plugin_options.conceal_subnet_restricted_resources`
- `marketplace_offering`: new attribute `plugin_options.disable_grace_period`
- `marketplace_offering`: new attribute `plugin_options.emit_display_name`
- `marketplace_offering`: new attribute `plugin_options.emit_waldur_username`
- `marketplace_offering`: new attribute `plugin_options.enable_membership_sync_status`
- `marketplace_offering`: new attribute `plugin_options.enable_posix_account`
- `marketplace_offering`: new attribute `plugin_options.enable_resource_access_subnets`
- `marketplace_offering`: new attribute `plugin_options.enforce_qos`
- `marketplace_offering`: new attribute `plugin_options.gid_source`
- `marketplace_offering`: new attribute `plugin_options.login_shell`
- `marketplace_offering`: new attribute `plugin_options.resource_projects_limit_policy`
- `marketplace_offering`: new attribute `plugin_options.restricted_to_roles`
- `marketplace_offering`: new attribute `plugin_options.uid_source`
- `marketplace_offering`: new attribute `qos_profiles`
- `marketplace_order`: new attribute `created_by_organization_address`
- `marketplace_order`: new attribute `created_by_organization_country`
- `marketplace_order`: new attribute `created_by_organization_vat_code`
- `marketplace_resource`: new attribute `has_api_keys`
- `marketplace_resource`: new attribute `order_in_progress.created_by_organization_address`
- `marketplace_resource`: new attribute `order_in_progress.created_by_organization_country`
- `marketplace_resource`: new attribute `order_in_progress.created_by_organization_vat_code`
- `marketplace_resource`: new attribute `resource_effective_end_date`
- `marketplace_resource`: new attribute `usage_limit_restriction`
- `openstack_instance`: new attribute `ports.allowed_address_pairs.ip_address`
- `structure_customer`: new attribute `has_active_helpdesk`
- `structure_customer`: new attribute `has_affiliate_links`

## 2026-08-03

### ⚠️ Breaking changes

- `marketplace_offering`: removed attribute `default_access_subnets`
- `marketplace_offering`: removed attribute `partitions.qos_options`
- `marketplace_offering`: removed attribute `plans.components.discount_aggregation`
- `marketplace_offering`: removed attribute `plans.components.discount_formula`
- `marketplace_offering`: removed attribute `plugin_options.action_on_usage_limit`
- `marketplace_offering`: removed attribute `plugin_options.auto_approve_for_roles`
- `marketplace_offering`: removed attribute `plugin_options.billing_source`
- `marketplace_offering`: removed attribute `plugin_options.conceal_subnet_restricted_resources`
- `marketplace_offering`: removed attribute `plugin_options.disable_grace_period`
- `marketplace_offering`: removed attribute `plugin_options.emit_display_name`
- `marketplace_offering`: removed attribute `plugin_options.emit_waldur_username`
- `marketplace_offering`: removed attribute `plugin_options.enable_membership_sync_status`
- `marketplace_offering`: removed attribute `plugin_options.enable_posix_account`
- `marketplace_offering`: removed attribute `plugin_options.enable_resource_access_subnets`
- `marketplace_offering`: removed attribute `plugin_options.enforce_qos`
- `marketplace_offering`: removed attribute `plugin_options.gid_source`
- `marketplace_offering`: removed attribute `plugin_options.login_shell`
- `marketplace_offering`: removed attribute `plugin_options.resource_projects_limit_policy`
- `marketplace_offering`: removed attribute `plugin_options.restricted_to_roles`
- `marketplace_offering`: removed attribute `plugin_options.uid_source`
- `marketplace_offering`: removed attribute `qos_profiles`
- `marketplace_order`: removed attribute `created_by_organization_address`
- `marketplace_order`: removed attribute `created_by_organization_country`
- `marketplace_order`: removed attribute `created_by_organization_vat_code`
- `marketplace_resource`: removed attribute `has_api_keys`
- `marketplace_resource`: removed attribute `order_in_progress.created_by_organization_address`
- `marketplace_resource`: removed attribute `order_in_progress.created_by_organization_country`
- `marketplace_resource`: removed attribute `order_in_progress.created_by_organization_vat_code`
- `marketplace_resource`: removed attribute `resource_effective_end_date`
- `marketplace_resource`: removed attribute `usage_limit_restriction`
- `openstack_instance`: removed attribute `ports.allowed_address_pairs.ip_address`
- `structure_customer`: attribute `customer_credit` type changed from `types.String` to `types.Float64`
- `structure_customer`: attribute `customer_unallocated_credit` type changed from `types.String` to `types.Float64`
- `structure_customer`: attribute `payment_profiles.attributes.contract_sum` type changed from `types.String` to `types.Int64`
- `structure_customer`: removed attribute `has_active_helpdesk`
- `structure_customer`: removed attribute `has_affiliate_links`

### Added

- `marketplace_offering`: new attribute `plans.components.discount_rate`
- `marketplace_offering`: new attribute `plans.components.discount_threshold`
- `marketplace_offering`: new attribute `plans.components.discounted_price`
- `marketplace_offering`: new attribute `plugin_options.initial_primarygroup_number`
- `marketplace_offering`: new attribute `plugin_options.initial_rolegroup_number`
- `marketplace_offering`: new attribute `plugin_options.initial_uidnumber`
- `marketplace_offering`: new attribute `plugin_options.initial_usergroup_number`
- `structure_project`: new attribute `project_metadata.answer`

## 2026-07-28

### Added

- `marketplace_offering`: new attribute `partitions.qos_options`
- `marketplace_offering`: new attribute `plugin_options.enforce_qos`
- `marketplace_offering`: new attribute `qos_profiles`

## 2026-07-27

### Added

- `marketplace_offering`: new attribute `plugin_options.enable_membership_sync_status`

## 2026-07-26

### Added

- `marketplace_resource`: new attribute `has_api_keys`

## 2026-07-23

### Added

- `marketplace_offering`: new attribute `plugin_options.conceal_subnet_restricted_resources`
- `marketplace_offering`: new attribute `plugin_options.enable_resource_access_subnets`
- `marketplace_offering`: new attribute `plugin_options.resource_projects_limit_policy`

## 2026-07-23

### Added

- `structure_customer`: new attribute `has_active_helpdesk`

## 2026-07-21

### ⚠️ Breaking changes

- `structure_project`: removed attribute `project_metadata.answer`

## 2026-07-21

### ⚠️ Breaking changes

- `structure_customer`: attribute `customer_credit` type changed from `types.Float64` to `types.String`
- `structure_customer`: attribute `customer_unallocated_credit` type changed from `types.Float64` to `types.String`
- `structure_customer`: attribute `payment_profiles.attributes.contract_sum` type changed from `types.Int64` to `types.String`

## 2026-07-20

### ⚠️ Breaking changes

- `marketplace_offering`: removed attribute `plans.components.discount_rate`
- `marketplace_offering`: removed attribute `plans.components.discount_threshold`
- `marketplace_offering`: removed attribute `plans.components.discounted_price`
- `marketplace_offering`: removed attribute `plugin_options.initial_primarygroup_number`
- `marketplace_offering`: removed attribute `plugin_options.initial_rolegroup_number`
- `marketplace_offering`: removed attribute `plugin_options.initial_uidnumber`
- `marketplace_offering`: removed attribute `plugin_options.initial_usergroup_number`
- `openstack_port`: attribute `allowed_address_pairs.ip_address` is now computed (read-only)

### Added

- `marketplace_offering`: new attribute `default_access_subnets`
- `marketplace_offering`: new attribute `plans.components.discount_aggregation`
- `marketplace_offering`: new attribute `plans.components.discount_formula`
- `marketplace_offering`: new attribute `plugin_options.action_on_usage_limit`
- `marketplace_offering`: new attribute `plugin_options.auto_approve_for_roles`
- `marketplace_offering`: new attribute `plugin_options.billing_source`
- `marketplace_offering`: new attribute `plugin_options.disable_grace_period`
- `marketplace_offering`: new attribute `plugin_options.emit_display_name`
- `marketplace_offering`: new attribute `plugin_options.emit_waldur_username`
- `marketplace_offering`: new attribute `plugin_options.enable_posix_account`
- `marketplace_offering`: new attribute `plugin_options.gid_source`
- `marketplace_offering`: new attribute `plugin_options.login_shell`
- `marketplace_offering`: new attribute `plugin_options.restricted_to_roles`
- `marketplace_offering`: new attribute `plugin_options.uid_source`
- `marketplace_order`: new attribute `created_by_organization_address`
- `marketplace_order`: new attribute `created_by_organization_country`
- `marketplace_order`: new attribute `created_by_organization_vat_code`
- `marketplace_resource`: new attribute `order_in_progress.created_by_organization_address`
- `marketplace_resource`: new attribute `order_in_progress.created_by_organization_country`
- `marketplace_resource`: new attribute `order_in_progress.created_by_organization_vat_code`
- `marketplace_resource`: new attribute `resource_effective_end_date`
- `marketplace_resource`: new attribute `usage_limit_restriction`
- `openstack_instance`: new attribute `ports.allowed_address_pairs.ip_address`
- `structure_customer`: new attribute `has_affiliate_links`

## 2026-07-14

### Added

- New data source `marketplace_offering`
- New data source `openstack_flavor`
- New data source `openstack_image`
- New data source `openstack_volume_type`
- New resource `core_ssh_public_key`
- New resource `customer_permission`
- New resource `identity_bridge_link`
- New resource `marketplace_order`
- New resource `marketplace_resource`
- New resource `openstack_floating_ip`
- New resource `openstack_instance`
- New resource `openstack_network_rbac_policy`
- New resource `openstack_network`
- New resource `openstack_port`
- New resource `openstack_security_group`
- New resource `openstack_server_group`
- New resource `openstack_subnet`
- New resource `openstack_tenant`
- New resource `openstack_volume_attachment`
- New resource `openstack_volume`
- New resource `project_permission`
- New resource `structure_customer`
- New resource `structure_project`
