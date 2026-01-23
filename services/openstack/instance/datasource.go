package instance

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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

// OpenstackInstanceDataSource defines the data source implementation.
type OpenstackInstanceDataSource struct {
	client *Client
}

// OpenstackInstanceFiltersModel contains the filter parameters for querying.
type OpenstackInstanceFiltersModel struct {
	AttachVolumeUuid     types.String `tfsdk:"attach_volume_uuid"`
	AvailabilityZoneName types.String `tfsdk:"availability_zone_name"`
	BackendId            types.String `tfsdk:"backend_id"`
	CanManage            types.Bool   `tfsdk:"can_manage"`
	Customer             types.String `tfsdk:"customer"`
	CustomerAbbreviation types.String `tfsdk:"customer_abbreviation"`
	CustomerName         types.String `tfsdk:"customer_name"`
	CustomerNativeName   types.String `tfsdk:"customer_native_name"`
	CustomerUuid         types.String `tfsdk:"customer_uuid"`
	Description          types.String `tfsdk:"description"`
	ExternalIp           types.String `tfsdk:"external_ip"`
	Name                 types.String `tfsdk:"name"`
	NameExact            types.String `tfsdk:"name_exact"`
	Project              types.String `tfsdk:"project"`
	ProjectName          types.String `tfsdk:"project_name"`
	ProjectUuid          types.String `tfsdk:"project_uuid"`
	Query                types.String `tfsdk:"query"`
	RuntimeState         types.String `tfsdk:"runtime_state"`
	ServiceSettingsName  types.String `tfsdk:"service_settings_name"`
	ServiceSettingsUuid  types.String `tfsdk:"service_settings_uuid"`
	State                types.String `tfsdk:"state"`
	Tenant               types.String `tfsdk:"tenant"`
	TenantUuid           types.String `tfsdk:"tenant_uuid"`
	Uuid                 types.String `tfsdk:"uuid"`
}

type OpenstackInstanceDataSourceModel struct {
	UUID                             types.String                   `tfsdk:"id"`
	Filters                          *OpenstackInstanceFiltersModel `tfsdk:"filters"`
	AccessUrl                        types.String                   `tfsdk:"access_url"`
	Action                           types.String                   `tfsdk:"action"`
	AvailabilityZone                 types.String                   `tfsdk:"availability_zone"`
	AvailabilityZoneName             types.String                   `tfsdk:"availability_zone_name"`
	BackendId                        types.String                   `tfsdk:"backend_id"`
	ConnectDirectlyToExternalNetwork types.Bool                     `tfsdk:"connect_directly_to_external_network"`
	Cores                            types.Int64                    `tfsdk:"cores"`
	Created                          types.String                   `tfsdk:"created"`
	Customer                         types.String                   `tfsdk:"customer"`
	CustomerAbbreviation             types.String                   `tfsdk:"customer_abbreviation"`
	CustomerName                     types.String                   `tfsdk:"customer_name"`
	CustomerNativeName               types.String                   `tfsdk:"customer_native_name"`
	CustomerUuid                     types.String                   `tfsdk:"customer_uuid"`
	Description                      types.String                   `tfsdk:"description"`
	Disk                             types.Int64                    `tfsdk:"disk"`
	ErrorMessage                     types.String                   `tfsdk:"error_message"`
	ErrorTraceback                   types.String                   `tfsdk:"error_traceback"`
	ExternalAddress                  types.List                     `tfsdk:"external_address"`
	ExternalIps                      types.List                     `tfsdk:"external_ips"`
	FlavorDisk                       types.Int64                    `tfsdk:"flavor_disk"`
	FlavorName                       types.String                   `tfsdk:"flavor_name"`
	FloatingIps                      types.List                     `tfsdk:"floating_ips"`
	HypervisorHostname               types.String                   `tfsdk:"hypervisor_hostname"`
	ImageName                        types.String                   `tfsdk:"image_name"`
	InternalIps                      types.List                     `tfsdk:"internal_ips"`
	IsLimitBased                     types.Bool                     `tfsdk:"is_limit_based"`
	IsUsageBased                     types.Bool                     `tfsdk:"is_usage_based"`
	KeyFingerprint                   types.String                   `tfsdk:"key_fingerprint"`
	KeyName                          types.String                   `tfsdk:"key_name"`
	Latitude                         types.Float64                  `tfsdk:"latitude"`
	Longitude                        types.Float64                  `tfsdk:"longitude"`
	MarketplaceCategoryName          types.String                   `tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid          types.String                   `tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName          types.String                   `tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid          types.String                   `tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid              types.String                   `tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState         types.String                   `tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid          types.String                   `tfsdk:"marketplace_resource_uuid"`
	MinDisk                          types.Int64                    `tfsdk:"min_disk"`
	MinRam                           types.Int64                    `tfsdk:"min_ram"`
	Modified                         types.String                   `tfsdk:"modified"`
	Name                             types.String                   `tfsdk:"name"`
	Ports                            types.List                     `tfsdk:"ports"`
	Project                          types.String                   `tfsdk:"project"`
	ProjectName                      types.String                   `tfsdk:"project_name"`
	ProjectUuid                      types.String                   `tfsdk:"project_uuid"`
	Ram                              types.Int64                    `tfsdk:"ram"`
	ResourceType                     types.String                   `tfsdk:"resource_type"`
	RuntimeState                     types.String                   `tfsdk:"runtime_state"`
	SecurityGroups                   types.List                     `tfsdk:"security_groups"`
	ServerGroup                      types.Object                   `tfsdk:"server_group"`
	ServiceName                      types.String                   `tfsdk:"service_name"`
	ServiceSettings                  types.String                   `tfsdk:"service_settings"`
	ServiceSettingsErrorMessage      types.String                   `tfsdk:"service_settings_error_message"`
	ServiceSettingsState             types.String                   `tfsdk:"service_settings_state"`
	ServiceSettingsUuid              types.String                   `tfsdk:"service_settings_uuid"`
	StartTime                        types.String                   `tfsdk:"start_time"`
	State                            types.String                   `tfsdk:"state"`
	Tenant                           types.String                   `tfsdk:"tenant"`
	TenantUuid                       types.String                   `tfsdk:"tenant_uuid"`
	Url                              types.String                   `tfsdk:"url"`
	UserData                         types.String                   `tfsdk:"user_data"`
	Volumes                          types.List                     `tfsdk:"volumes"`
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
					"state": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "State Allowed values: `CREATING`, `CREATION_SCHEDULED`, `DELETING`, `DELETION_SCHEDULED`, `ERRED`, `OK`, `UPDATE_SCHEDULED`, `UPDATING`.",
						Validators: []validator.String{
							stringvalidator.OneOf("CREATING", "CREATION_SCHEDULED", "DELETING", "DELETION_SCHEDULED", "ERRED", "OK", "UPDATE_SCHEDULED", "UPDATING"),
						},
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
				Computed:            true,
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
				Computed:            true,
				MarkdownDescription: "If True, instance will be connected directly to external network",
			},
			"cores": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Number of cores in a VM",
			},
			"created": schema.StringAttribute{
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
				Computed:            true,
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
			"floating_ips": schema.ListNestedAttribute{
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
				Computed:            true,
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
				Computed:            true,
				MarkdownDescription: "Modified",
			},
			"name": schema.StringAttribute{
				Computed:            true,
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
						"security_groups": schema.ListNestedAttribute{
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
				Computed:            true,
				MarkdownDescription: "Ports",
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
			"security_groups": schema.ListNestedAttribute{
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
				Computed:            true,
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
				Computed:            true,
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
				Computed:            true,
				MarkdownDescription: "Start time",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "State",
			},
			"tenant": schema.StringAttribute{
				Computed:            true,
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
				Computed:            true,
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

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, *apiResp, &data)...)

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

		resp.Diagnostics.Append(d.mapResponseToModel(ctx, results[0], &data)...)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *OpenstackInstanceDataSource) mapResponseToModel(ctx context.Context, apiResp OpenstackInstanceResponse, model *OpenstackInstanceDataSourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
	model.Action = types.StringPointerValue(apiResp.Action)
	model.AvailabilityZone = types.StringPointerValue(apiResp.AvailabilityZone)
	model.AvailabilityZoneName = types.StringPointerValue(apiResp.AvailabilityZoneName)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.ConnectDirectlyToExternalNetwork = types.BoolPointerValue(apiResp.ConnectDirectlyToExternalNetwork)
	model.Cores = types.Int64PointerValue(apiResp.Cores)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.Customer = types.StringPointerValue(apiResp.Customer)
	model.CustomerAbbreviation = types.StringPointerValue(apiResp.CustomerAbbreviation)
	model.CustomerName = types.StringPointerValue(apiResp.CustomerName)
	model.CustomerNativeName = types.StringPointerValue(apiResp.CustomerNativeName)
	model.CustomerUuid = types.StringPointerValue(apiResp.CustomerUuid)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.Disk = types.Int64PointerValue(apiResp.Disk)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.ExternalAddress, _ = types.ListValueFrom(ctx, types.StringType, apiResp.ExternalAddress)
	model.ExternalIps, _ = types.ListValueFrom(ctx, types.StringType, apiResp.ExternalIps)
	model.FlavorDisk = types.Int64PointerValue(apiResp.FlavorDisk)
	model.FlavorName = types.StringPointerValue(apiResp.FlavorName)

	{
		listValFloatingIps, listDiagsFloatingIps := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"address": types.StringType,
			"port_fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
				"ip_address": types.StringType,
				"subnet_id":  types.StringType,
			}}},
			"port_mac_address":   types.StringType,
			"subnet":             types.StringType,
			"subnet_cidr":        types.StringType,
			"subnet_description": types.StringType,
			"subnet_name":        types.StringType,
			"subnet_uuid":        types.StringType,
			"url":                types.StringType,
		}}, apiResp.FloatingIps)
		diags.Append(listDiagsFloatingIps...)
		model.FloatingIps = listValFloatingIps
	}
	model.HypervisorHostname = types.StringPointerValue(apiResp.HypervisorHostname)
	model.ImageName = types.StringPointerValue(apiResp.ImageName)
	model.InternalIps, _ = types.ListValueFrom(ctx, types.StringType, apiResp.InternalIps)
	model.IsLimitBased = types.BoolPointerValue(apiResp.IsLimitBased)
	model.IsUsageBased = types.BoolPointerValue(apiResp.IsUsageBased)
	model.KeyFingerprint = types.StringPointerValue(apiResp.KeyFingerprint)
	model.KeyName = types.StringPointerValue(apiResp.KeyName)
	model.Latitude = types.Float64PointerValue(apiResp.Latitude)
	model.Longitude = types.Float64PointerValue(apiResp.Longitude)
	model.MarketplaceCategoryName = types.StringPointerValue(apiResp.MarketplaceCategoryName)
	model.MarketplaceCategoryUuid = types.StringPointerValue(apiResp.MarketplaceCategoryUuid)
	model.MarketplaceOfferingName = types.StringPointerValue(apiResp.MarketplaceOfferingName)
	model.MarketplaceOfferingUuid = types.StringPointerValue(apiResp.MarketplaceOfferingUuid)
	model.MarketplacePlanUuid = types.StringPointerValue(apiResp.MarketplacePlanUuid)
	model.MarketplaceResourceState = types.StringPointerValue(apiResp.MarketplaceResourceState)
	model.MarketplaceResourceUuid = types.StringPointerValue(apiResp.MarketplaceResourceUuid)
	model.MinDisk = types.Int64PointerValue(apiResp.MinDisk)
	model.MinRam = types.Int64PointerValue(apiResp.MinRam)
	model.Modified = types.StringPointerValue(apiResp.Modified)
	model.Name = types.StringPointerValue(apiResp.Name)

	{
		listValPorts, listDiagsPorts := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
				"mac_address": types.StringType,
			}}},
			"device_id":    types.StringType,
			"device_owner": types.StringType,
			"fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
				"ip_address": types.StringType,
				"subnet_id":  types.StringType,
			}}},
			"mac_address": types.StringType,
			"security_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
				"access_url":                 types.StringType,
				"backend_id":                 types.StringType,
				"created":                    types.StringType,
				"customer":                   types.StringType,
				"customer_abbreviation":      types.StringType,
				"customer_name":              types.StringType,
				"customer_native_name":       types.StringType,
				"customer_uuid":              types.StringType,
				"description":                types.StringType,
				"error_message":              types.StringType,
				"error_traceback":            types.StringType,
				"is_limit_based":             types.BoolType,
				"is_usage_based":             types.BoolType,
				"marketplace_category_name":  types.StringType,
				"marketplace_category_uuid":  types.StringType,
				"marketplace_offering_name":  types.StringType,
				"marketplace_offering_uuid":  types.StringType,
				"marketplace_plan_uuid":      types.StringType,
				"marketplace_resource_state": types.StringType,
				"marketplace_resource_uuid":  types.StringType,
				"modified":                   types.StringType,
				"name":                       types.StringType,
				"project":                    types.StringType,
				"project_name":               types.StringType,
				"project_uuid":               types.StringType,
				"resource_type":              types.StringType,
				"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"cidr":              types.StringType,
					"description":       types.StringType,
					"direction":         types.StringType,
					"ethertype":         types.StringType,
					"from_port":         types.Int64Type,
					"id":                types.Int64Type,
					"protocol":          types.StringType,
					"remote_group":      types.StringType,
					"remote_group_name": types.StringType,
					"remote_group_uuid": types.StringType,
					"to_port":           types.Int64Type,
				}}},
				"service_name":                   types.StringType,
				"service_settings":               types.StringType,
				"service_settings_error_message": types.StringType,
				"service_settings_state":         types.StringType,
				"service_settings_uuid":          types.StringType,
				"state":                          types.StringType,
				"tenant":                         types.StringType,
				"tenant_name":                    types.StringType,
				"tenant_uuid":                    types.StringType,
				"url":                            types.StringType,
			}}},
			"subnet":             types.StringType,
			"subnet_cidr":        types.StringType,
			"subnet_description": types.StringType,
			"subnet_name":        types.StringType,
			"subnet_uuid":        types.StringType,
			"url":                types.StringType,
		}}, apiResp.Ports)
		diags.Append(listDiagsPorts...)
		model.Ports = listValPorts
	}
	model.Project = types.StringPointerValue(apiResp.Project)
	model.ProjectName = types.StringPointerValue(apiResp.ProjectName)
	model.ProjectUuid = types.StringPointerValue(apiResp.ProjectUuid)
	model.Ram = types.Int64PointerValue(apiResp.Ram)
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	model.RuntimeState = types.StringPointerValue(apiResp.RuntimeState)

	{
		listValSecurityGroups, listDiagsSecurityGroups := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"name":        types.StringType,
			"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
				"cidr":              types.StringType,
				"description":       types.StringType,
				"direction":         types.StringType,
				"ethertype":         types.StringType,
				"from_port":         types.Int64Type,
				"id":                types.Int64Type,
				"protocol":          types.StringType,
				"remote_group_name": types.StringType,
				"remote_group_uuid": types.StringType,
				"to_port":           types.Int64Type,
			}}},
			"state": types.StringType,
			"url":   types.StringType,
		}}, apiResp.SecurityGroups)
		diags.Append(listDiagsSecurityGroups...)
		model.SecurityGroups = listValSecurityGroups
	}
	if apiResp.ServerGroup != nil {
		objValServerGroup, objDiagsServerGroup := types.ObjectValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"name":   types.StringType,
			"policy": types.StringType,
			"state":  types.StringType,
			"url":    types.StringType,
		}}.AttrTypes, *apiResp.ServerGroup)
		diags.Append(objDiagsServerGroup...)
		model.ServerGroup = objValServerGroup
	} else {
		model.ServerGroup = types.ObjectNull(types.ObjectType{AttrTypes: map[string]attr.Type{
			"name":   types.StringType,
			"policy": types.StringType,
			"state":  types.StringType,
			"url":    types.StringType,
		}}.AttrTypes)
	}
	model.ServiceName = types.StringPointerValue(apiResp.ServiceName)
	model.ServiceSettings = types.StringPointerValue(apiResp.ServiceSettings)
	model.ServiceSettingsErrorMessage = types.StringPointerValue(apiResp.ServiceSettingsErrorMessage)
	model.ServiceSettingsState = types.StringPointerValue(apiResp.ServiceSettingsState)
	model.ServiceSettingsUuid = types.StringPointerValue(apiResp.ServiceSettingsUuid)
	model.StartTime = types.StringPointerValue(apiResp.StartTime)
	model.State = types.StringPointerValue(apiResp.State)
	model.Tenant = types.StringPointerValue(apiResp.Tenant)
	model.TenantUuid = types.StringPointerValue(apiResp.TenantUuid)
	model.Url = types.StringPointerValue(apiResp.Url)
	model.UserData = types.StringPointerValue(apiResp.UserData)

	{
		listValVolumes, listDiagsVolumes := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"bootable":                  types.BoolType,
			"device":                    types.StringType,
			"image_name":                types.StringType,
			"marketplace_resource_uuid": types.StringType,
			"name":                      types.StringType,
			"resource_type":             types.StringType,
			"size":                      types.Int64Type,
			"state":                     types.StringType,
			"type":                      types.StringType,
			"type_name":                 types.StringType,
			"url":                       types.StringType,
		}}, apiResp.Volumes)
		diags.Append(listDiagsVolumes...)
		model.Volumes = listValVolumes
	}

	return diags
}
