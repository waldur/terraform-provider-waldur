package instance

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OpenstackInstanceDataSource{}

func NewOpenstackInstanceDataSource() datasource.DataSource {
	return &OpenstackInstanceDataSource{}
}

type OpenstackInstanceDataSource struct {
	client *Client
}

type OpenstackInstanceDataSourceModel struct {
	OpenstackInstanceModel
	Filters *OpenstackInstanceFiltersModel `tfsdk:"filters"`
}

func (d *OpenstackInstanceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_instance"
}

func (d *OpenstackInstanceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Instance data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"filters": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Filter parameters for querying Openstack Instance",
				Attributes: map[string]schema.Attribute{
					"attach_volume_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Filter for attachment to volume UUID",
					},
					"availability_zone_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Availability zone name",
					},
					"backend_id": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Backend ID",
					},
					"can_manage": schema.BoolAttribute{
						Optional:            true,
						MarkdownDescription: "Can manage",
					},
					"customer": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Customer UUID",
					},
					"customer_abbreviation": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Customer abbreviation",
					},
					"customer_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Customer name",
					},
					"customer_native_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Customer native name",
					},
					"customer_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Customer UUID",
					},
					"description": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Description",
					},
					"external_ip": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "External IP",
					},
					"name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Name",
					},
					"name_exact": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Name (exact)",
					},
					"project": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Project UUID",
					},
					"project_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Project name",
					},
					"project_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Project UUID",
					},
					"query": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Search by name, internal IP, or external IP",
					},
					"runtime_state": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Runtime state",
					},
					"service_settings_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Service settings name",
					},
					"service_settings_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Service settings UUID",
					},
					"tenant": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Tenant URL",
					},
					"tenant_uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Tenant UUID",
					},
					"uuid": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "UUID",
					},
				},
			},
			"access_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Access url",
			},
			"action": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Action",
			},
			"availability_zone": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Availability zone where this instance is located",
			},
			"availability_zone_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the availability zone where instance is located",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Instance ID in the OpenStack backend",
			},
			"connect_directly_to_external_network": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "If True, instance will be connected directly to external network",
			},
			"cores": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Number of cores in a VM",
			},
			"created": schema.StringAttribute{
				CustomType:          timetypes.RFC3339Type{},
				Computed:            true,
				MarkdownDescription: "Created",
			},
			"customer": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Customer",
			},
			"customer_abbreviation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Customer abbreviation",
			},
			"customer_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the customer",
			},
			"customer_native_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the customer native",
			},
			"customer_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the customer",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Description of the resource",
			},
			"disk": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Disk size in MiB",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error message",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error traceback",
			},
			"external_address": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "External address",
			},
			"external_ips": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "External ips",
			},
			"flavor_disk": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Flavor disk size in MiB",
			},
			"flavor_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the flavor used by this instance",
			},
			"floating_ips": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "The public IPv4 address of the floating IP",
						},
						"port_fixed_ips": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ip_address": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "IP address to assign to the port",
									},
									"subnet_id": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "ID of the subnet in which to assign the IP address",
									},
								},
							},
							Computed:            true,
							MarkdownDescription: "Port fixed ips",
						},
						"port_mac_address": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "MAC address of the port",
						},
						"subnet": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Subnet",
						},
						"subnet_cidr": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "IPv4 network address in CIDR format (e.g. 192.168.0.0/24)",
						},
						"subnet_description": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Subnet description",
						},
						"subnet_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name of the subnet",
						},
						"subnet_uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "UUID of the subnet",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Url",
						},
					},
				},
				Optional:            true,
				MarkdownDescription: "Floating ips",
			},
			"hypervisor_hostname": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the hypervisor hosting this instance",
			},
			"image_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the image",
			},
			"internal_ips": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "Internal ips",
			},
			"is_limit_based": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Is limit based",
			},
			"is_usage_based": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Is usage based",
			},
			"key_fingerprint": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Key fingerprint",
			},
			"key_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the key",
			},
			"latitude": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "Latitude",
			},
			"longitude": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "Longitude",
			},
			"marketplace_category_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the marketplace category",
			},
			"marketplace_category_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the marketplace category",
			},
			"marketplace_offering_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the marketplace offering",
			},
			"marketplace_offering_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the marketplace offering",
			},
			"marketplace_plan_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the marketplace plan",
			},
			"marketplace_resource_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Marketplace resource state",
			},
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the marketplace resource",
			},
			"min_disk": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Minimum disk size in MiB",
			},
			"min_ram": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Minimum memory size in MiB",
			},
			"modified": schema.StringAttribute{
				CustomType:          timetypes.RFC3339Type{},
				Computed:            true,
				MarkdownDescription: "Modified",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name of the resource",
			},
			"ports": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"allowed_address_pairs": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"mac_address": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Mac address",
									},
								},
							},
							Computed:            true,
							MarkdownDescription: "Allowed address pairs",
						},
						"device_id": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "ID of device (instance, router etc) to which this port is connected",
						},
						"device_owner": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Entity that uses this port (e.g. network:router_interface)",
						},
						"fixed_ips": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ip_address": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "IP address to assign to the port",
									},
									"subnet_id": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "ID of the subnet in which to assign the IP address",
									},
								},
							},
							Optional:            true,
							MarkdownDescription: "Fixed ips",
						},
						"mac_address": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "MAC address of the port",
						},
						"security_groups": schema.SetNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"access_url": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Access url",
									},
									"backend_id": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "ID of the backend",
									},
									"created": schema.StringAttribute{
										CustomType:          timetypes.RFC3339Type{},
										Computed:            true,
										MarkdownDescription: "Created",
									},
									"customer": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Customer",
									},
									"customer_abbreviation": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Customer abbreviation",
									},
									"customer_name": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Name of the customer",
									},
									"customer_native_name": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Name of the customer native",
									},
									"customer_uuid": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "UUID of the customer",
									},
									"description": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Description of the resource",
									},
									"error_message": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Error message",
									},
									"error_traceback": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Error traceback",
									},
									"is_limit_based": schema.BoolAttribute{
										Computed:            true,
										MarkdownDescription: "Is limit based",
									},
									"is_usage_based": schema.BoolAttribute{
										Computed:            true,
										MarkdownDescription: "Is usage based",
									},
									"marketplace_category_name": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Name of the marketplace category",
									},
									"marketplace_category_uuid": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "UUID of the marketplace category",
									},
									"marketplace_offering_name": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Name of the marketplace offering",
									},
									"marketplace_offering_uuid": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "UUID of the marketplace offering",
									},
									"marketplace_plan_uuid": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "UUID of the marketplace plan",
									},
									"marketplace_resource_state": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Marketplace resource state",
									},
									"marketplace_resource_uuid": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "UUID of the marketplace resource",
									},
									"modified": schema.StringAttribute{
										CustomType:          timetypes.RFC3339Type{},
										Computed:            true,
										MarkdownDescription: "Modified",
									},
									"name": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Name of the resource",
									},
									"project": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Project",
									},
									"project_name": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Name of the project",
									},
									"project_uuid": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "UUID of the project",
									},
									"resource_type": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Resource type",
									},
									"rules": schema.ListNestedAttribute{
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"cidr": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "CIDR notation for the source/destination network address range",
												},
												"description": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "Description of the resource",
												},
												"direction": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "Traffic direction - either 'ingress' (incoming) or 'egress' (outgoing)",
												},
												"ethertype": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "IP protocol version - either 'IPv4' or 'IPv6'",
												},
												"from_port": schema.Int64Attribute{
													Optional:            true,
													MarkdownDescription: "Starting port number in the range (1-65535)",
													Validators: []validator.Int64{
														int64validator.AtLeast(-2147483648),
														int64validator.AtMost(65535),
													},
												},
												"id": schema.Int64Attribute{
													Computed:            true,
													MarkdownDescription: "Id",
												},
												"protocol": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "The network protocol (TCP, UDP, ICMP, or empty for any protocol)",
												},
												"remote_group": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "Remote security group that this rule references, if any",
												},
												"remote_group_name": schema.StringAttribute{
													Computed:            true,
													MarkdownDescription: "Name of the remote group",
												},
												"remote_group_uuid": schema.StringAttribute{
													Computed:            true,
													MarkdownDescription: "UUID of the remote group",
												},
												"to_port": schema.Int64Attribute{
													Optional:            true,
													MarkdownDescription: "Ending port number in the range (1-65535)",
													Validators: []validator.Int64{
														int64validator.AtLeast(-2147483648),
														int64validator.AtMost(65535),
													},
												},
											},
										},
										Optional:            true,
										MarkdownDescription: "Rules",
									},
									"service_name": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Name of the service",
									},
									"service_settings": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Service settings",
									},
									"service_settings_error_message": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Service settings error message",
									},
									"service_settings_state": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Service settings state",
									},
									"service_settings_uuid": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "UUID of the service settings",
									},
									"state": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "State",
									},
									"tenant": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Tenant",
									},
									"tenant_name": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Name of the tenant",
									},
									"tenant_uuid": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "UUID of the tenant",
									},
									"url": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Url",
									},
								},
							},
							Computed:            true,
							MarkdownDescription: "Security groups",
						},
						"subnet": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Subnet to which this port belongs",
						},
						"subnet_cidr": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "IPv4 network address in CIDR format (e.g. 192.168.0.0/24)",
						},
						"subnet_description": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Subnet description",
						},
						"subnet_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name of the subnet",
						},
						"subnet_uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "UUID of the subnet",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Url",
						},
					},
				},
				Optional:            true,
				MarkdownDescription: "Ports",
			},
			"project": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Project",
			},
			"project_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the project",
			},
			"project_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the project",
			},
			"ram": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Memory size in MiB",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource type",
			},
			"runtime_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Runtime state",
			},
			"security_groups": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"description": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Description of the resource",
						},
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name of the resource",
						},
						"rules": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"cidr": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "CIDR notation for the source/destination network address range",
									},
									"description": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Description of the resource",
									},
									"direction": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Traffic direction - either 'ingress' (incoming) or 'egress' (outgoing)",
									},
									"ethertype": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "IP protocol version - either 'IPv4' or 'IPv6'",
									},
									"from_port": schema.Int64Attribute{
										Optional:            true,
										MarkdownDescription: "Starting port number in the range (1-65535)",
										Validators: []validator.Int64{
											int64validator.AtLeast(-2147483648),
											int64validator.AtMost(65535),
										},
									},
									"id": schema.Int64Attribute{
										Computed:            true,
										MarkdownDescription: "Id",
									},
									"protocol": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "The network protocol (TCP, UDP, ICMP, or empty for any protocol)",
									},
									"remote_group_name": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Name of the remote group",
									},
									"remote_group_uuid": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "UUID of the remote group",
									},
									"to_port": schema.Int64Attribute{
										Optional:            true,
										MarkdownDescription: "Ending port number in the range (1-65535)",
										Validators: []validator.Int64{
											int64validator.AtLeast(-2147483648),
											int64validator.AtMost(65535),
										},
									},
								},
							},
							Computed:            true,
							MarkdownDescription: "Rules",
						},
						"state": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "State",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Url",
						},
					},
				},
				Optional:            true,
				MarkdownDescription: "Security groups",
			},
			"server_group": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Name of the resource",
					},
					"policy": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Server group policy determining the rules for scheduling servers in this group",
					},
					"state": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "State",
					},
					"url": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Url",
					},
				},
				Optional:            true,
				MarkdownDescription: "Server group",
			},
			"service_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the service",
			},
			"service_settings": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "OpenStack provider settings",
			},
			"service_settings_error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Service settings error message",
			},
			"service_settings_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Service settings state",
			},
			"service_settings_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the service settings",
			},
			"start_time": schema.StringAttribute{
				CustomType:          timetypes.RFC3339Type{},
				Computed:            true,
				MarkdownDescription: "Start time",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "State",
			},
			"tenant": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The OpenStack tenant to create the instance in",
			},
			"tenant_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the OpenStack tenant",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Url",
			},
			"user_data": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Additional data that will be added to instance on provisioning",
			},
			"volumes": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"bootable": schema.BoolAttribute{
							Optional:            true,
							MarkdownDescription: "Indicates if this volume can be used to boot an instance",
						},
						"device": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of volume as instance device e.g. /dev/vdb.",
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^/dev/[a-zA-Z0-9]+$`), ""),
							},
						},
						"image_name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Name of the image this volume was created from",
						},
						"marketplace_resource_uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "UUID of the marketplace resource",
						},
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name of the resource",
						},
						"resource_type": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Resource type",
						},
						"size": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "Size in MiB",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"state": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "State",
						},
						"type": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Type of the volume (e.g. SSD, HDD)",
						},
						"type_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Name of the type",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Url",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "List of volumes attached to the instance",
			},
		},
	}
}

func (d *OpenstackInstanceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	rawClient, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Expected *client.Client, got something else. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = NewClient(rawClient)
}

func (d *OpenstackInstanceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackInstanceDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		apiResp, err := d.client.GetOpenstackInstance(ctx, data.UUID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Openstack Instance",
				"An error occurred while reading the Openstack Instance by UUID: "+err.Error(),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	} else {
		filters := common.BuildQueryFilters(data.Filters)

		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_instance.",
			)
			return
		}

		results, err := d.client.ListOpenstackInstance(ctx, filters)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Openstack Instance",
				"An error occurred while filtering Openstack Instance: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Openstack Instance Not Found",
				"No Openstack Instance found with provided filters.",
			)
			return
		}

		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Openstack Instances Found",
				fmt.Sprintf("Found %d Openstack Instances with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		resp.Diagnostics.Append(data.CopyFrom(ctx, results[0])...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
