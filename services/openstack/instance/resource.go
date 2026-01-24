package instance

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OpenstackInstanceResource{}
var _ resource.ResourceWithImportState = &OpenstackInstanceResource{}

func NewOpenstackInstanceResource() resource.Resource {
	return &OpenstackInstanceResource{}
}

// OpenstackInstanceResource defines the resource implementation.
type OpenstackInstanceResource struct {
	client *Client
}

// OpenstackInstanceResourceModel describes the resource data model.
type OpenstackInstanceResourceModel struct {
	UUID                             types.String   `tfsdk:"id"`
	AccessUrl                        types.String   `tfsdk:"access_url"`
	Action                           types.String   `tfsdk:"action"`
	AvailabilityZone                 types.String   `tfsdk:"availability_zone"`
	AvailabilityZoneName             types.String   `tfsdk:"availability_zone_name"`
	BackendId                        types.String   `tfsdk:"backend_id"`
	ConnectDirectlyToExternalNetwork types.Bool     `tfsdk:"connect_directly_to_external_network"`
	Cores                            types.Int64    `tfsdk:"cores"`
	Created                          types.String   `tfsdk:"created"`
	Customer                         types.String   `tfsdk:"customer"`
	CustomerAbbreviation             types.String   `tfsdk:"customer_abbreviation"`
	CustomerName                     types.String   `tfsdk:"customer_name"`
	CustomerNativeName               types.String   `tfsdk:"customer_native_name"`
	CustomerUuid                     types.String   `tfsdk:"customer_uuid"`
	DataVolumeSize                   types.Int64    `tfsdk:"data_volume_size"`
	DataVolumeType                   types.String   `tfsdk:"data_volume_type"`
	DataVolumes                      types.List     `tfsdk:"data_volumes"`
	DeleteVolumes                    types.Bool     `tfsdk:"delete_volumes"`
	Description                      types.String   `tfsdk:"description"`
	Disk                             types.Int64    `tfsdk:"disk"`
	ErrorMessage                     types.String   `tfsdk:"error_message"`
	ErrorTraceback                   types.String   `tfsdk:"error_traceback"`
	ExternalAddress                  types.List     `tfsdk:"external_address"`
	ExternalIps                      types.List     `tfsdk:"external_ips"`
	Flavor                           types.String   `tfsdk:"flavor"`
	FlavorDisk                       types.Int64    `tfsdk:"flavor_disk"`
	FlavorName                       types.String   `tfsdk:"flavor_name"`
	FloatingIps                      types.List     `tfsdk:"floating_ips"`
	HypervisorHostname               types.String   `tfsdk:"hypervisor_hostname"`
	Image                            types.String   `tfsdk:"image"`
	ImageName                        types.String   `tfsdk:"image_name"`
	InternalIps                      types.List     `tfsdk:"internal_ips"`
	IsLimitBased                     types.Bool     `tfsdk:"is_limit_based"`
	IsUsageBased                     types.Bool     `tfsdk:"is_usage_based"`
	KeyFingerprint                   types.String   `tfsdk:"key_fingerprint"`
	KeyName                          types.String   `tfsdk:"key_name"`
	Latitude                         types.Float64  `tfsdk:"latitude"`
	Longitude                        types.Float64  `tfsdk:"longitude"`
	MarketplaceCategoryName          types.String   `tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid          types.String   `tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName          types.String   `tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid          types.String   `tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid              types.String   `tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState         types.String   `tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid          types.String   `tfsdk:"marketplace_resource_uuid"`
	MinDisk                          types.Int64    `tfsdk:"min_disk"`
	MinRam                           types.Int64    `tfsdk:"min_ram"`
	Modified                         types.String   `tfsdk:"modified"`
	Name                             types.String   `tfsdk:"name"`
	Offering                         types.String   `tfsdk:"offering"`
	Ports                            types.List     `tfsdk:"ports"`
	Project                          types.String   `tfsdk:"project"`
	ProjectName                      types.String   `tfsdk:"project_name"`
	ProjectUuid                      types.String   `tfsdk:"project_uuid"`
	Ram                              types.Int64    `tfsdk:"ram"`
	ReleaseFloatingIps               types.Bool     `tfsdk:"release_floating_ips"`
	ResourceType                     types.String   `tfsdk:"resource_type"`
	RuntimeState                     types.String   `tfsdk:"runtime_state"`
	SecurityGroups                   types.List     `tfsdk:"security_groups"`
	ServerGroup                      types.Object   `tfsdk:"server_group"`
	ServiceName                      types.String   `tfsdk:"service_name"`
	ServiceSettings                  types.String   `tfsdk:"service_settings"`
	ServiceSettingsErrorMessage      types.String   `tfsdk:"service_settings_error_message"`
	ServiceSettingsState             types.String   `tfsdk:"service_settings_state"`
	ServiceSettingsUuid              types.String   `tfsdk:"service_settings_uuid"`
	SshPublicKey                     types.String   `tfsdk:"ssh_public_key"`
	StartTime                        types.String   `tfsdk:"start_time"`
	State                            types.String   `tfsdk:"state"`
	SystemVolumeSize                 types.Int64    `tfsdk:"system_volume_size"`
	SystemVolumeType                 types.String   `tfsdk:"system_volume_type"`
	Tenant                           types.String   `tfsdk:"tenant"`
	TenantUuid                       types.String   `tfsdk:"tenant_uuid"`
	Url                              types.String   `tfsdk:"url"`
	UserData                         types.String   `tfsdk:"user_data"`
	Volumes                          types.List     `tfsdk:"volumes"`
	Timeouts                         timeouts.Value `tfsdk:"timeouts"`
}

func (r *OpenstackInstanceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_instance"
}

func (r *OpenstackInstanceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Instance resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"access_url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Access url",
			},
			"action": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Action",
			},
			"availability_zone": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Availability zone where this instance is located",
			},
			"availability_zone_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the availability zone where instance is located",
			},
			"backend_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Instance ID in the OpenStack backend",
			},
			"connect_directly_to_external_network": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "If True, instance will be connected directly to external network",
			},
			"cores": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Number of cores in a VM",
			},
			"created": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Created",
			},
			"customer": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Customer",
			},
			"customer_abbreviation": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Customer abbreviation",
			},
			"customer_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the customer",
			},
			"customer_native_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the customer native",
			},
			"customer_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the customer",
			},
			"data_volume_size": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Size of the data volume in MiB. Minimum size is 1024 MiB (1 GiB)",
			},
			"data_volume_type": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Volume type for the data volume",
			},
			"data_volumes": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"size": schema.Int64Attribute{
							Required:            true,
							MarkdownDescription: "Size",
						},
						"volume_type": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Volume type",
						},
					},
				},
				Optional: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Additional data volumes to attach to the instance",
			},
			"delete_volumes": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Termination attribute",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Description of the resource",
			},
			"disk": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Disk size in MiB",
			},
			"error_message": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Error message",
			},
			"error_traceback": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Error traceback",
			},
			"external_address": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.StringType},
				Computed:   true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "External address",
			},
			"external_ips": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.StringType},
				Computed:   true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "External ips",
			},
			"flavor": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The flavor to use for the instance",
			},
			"flavor_disk": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Flavor disk size in MiB",
			},
			"flavor_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the flavor used by this instance",
			},
			"floating_ips": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Existing floating IP address in selected OpenStack tenant to be assigned to new virtual machine",
						},
						"subnet": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Subnet",
						},
						"url": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Url",
						},
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
					},
				},
				Optional:            true,
				MarkdownDescription: "Floating IPs to assign to the instance",
			},
			"hypervisor_hostname": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the hypervisor hosting this instance",
			},
			"image": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "The OS image to use for the instance",
			},
			"image_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the image",
			},
			"internal_ips": schema.ListAttribute{
				CustomType: types.ListType{ElemType: types.StringType},
				Computed:   true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Internal ips",
			},
			"is_limit_based": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Is limit based",
			},
			"is_usage_based": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Is usage based",
			},
			"key_fingerprint": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Key fingerprint",
			},
			"key_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the key",
			},
			"latitude": schema.Float64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Latitude",
			},
			"longitude": schema.Float64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Longitude",
			},
			"marketplace_category_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the marketplace category",
			},
			"marketplace_category_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the marketplace category",
			},
			"marketplace_offering_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the marketplace offering",
			},
			"marketplace_offering_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the marketplace offering",
			},
			"marketplace_plan_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the marketplace plan",
			},
			"marketplace_resource_state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Marketplace resource state",
			},
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the marketplace resource",
			},
			"min_disk": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Minimum disk size in MiB",
			},
			"min_ram": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Minimum memory size in MiB",
			},
			"modified": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Modified",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Name of the resource",
			},
			"offering": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Offering URL",
			},
			"ports": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"fixed_ips": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ip_address": schema.StringAttribute{
										Required:            true,
										MarkdownDescription: "IP address to assign to the port",
									},
									"subnet_id": schema.StringAttribute{
										Required:            true,
										MarkdownDescription: "ID of the subnet in which to assign the IP address",
									},
								},
							},
							Optional:            true,
							MarkdownDescription: "Fixed ips",
						},
						"port": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Port",
						},
						"subnet": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Subnet to which this port belongs",
						},
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
				MarkdownDescription: "Network ports to attach to the instance",
			},
			"project": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Project",
			},
			"project_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the project",
			},
			"project_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the project",
			},
			"ram": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Memory size in MiB",
			},
			"release_floating_ips": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Termination attribute",
			},
			"resource_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Resource type",
			},
			"runtime_state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Runtime state",
			},
			"security_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"url": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Url",
						},
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
					},
				},
				Optional:            true,
				MarkdownDescription: "List of security groups to apply to the instance",
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
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Server group",
			},
			"service_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the service",
			},
			"service_settings": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "OpenStack provider settings",
			},
			"service_settings_error_message": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Service settings error message",
			},
			"service_settings_state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Service settings state",
			},
			"service_settings_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the service settings",
			},
			"ssh_public_key": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Ssh public key",
			},
			"start_time": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Start time",
			},
			"state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "State",
			},
			"system_volume_size": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Size of the system volume in MiB. Minimum size is 1024 MiB (1 GiB)",
			},
			"system_volume_type": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Volume type for the system volume",
			},
			"tenant": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "The OpenStack tenant to create the instance in",
			},
			"tenant_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the OpenStack tenant",
			},
			"url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Url",
			},
			"user_data": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "List of volumes attached to the instance",
			},
		},

		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx, timeouts.Opts{
				Create: true,
				Update: true,
				Delete: true,
			}),
		},
	}
}

func (r *OpenstackInstanceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			"Expected *client.Client, got something else. Please report this issue to the provider developers.",
		)
		return
	}

	r.client = NewClient(client)
}

func (r *OpenstackInstanceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackInstanceResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Phase 1: Payload Construction
	// We map the Terraform schema fields to the 'attributes' map required by the Marketplace Order API.
	attributes := OpenstackInstanceCreateAttributes{
		AvailabilityZone:                 data.AvailabilityZone.ValueStringPointer(),
		ConnectDirectlyToExternalNetwork: data.ConnectDirectlyToExternalNetwork.ValueBoolPointer(),
		DataVolumeSize:                   data.DataVolumeSize.ValueInt64Pointer(),
		DataVolumeType:                   data.DataVolumeType.ValueStringPointer(),
		Description:                      data.Description.ValueStringPointer(),
		Flavor:                           data.Flavor.ValueStringPointer(),
		Image:                            data.Image.ValueStringPointer(),
		Name:                             data.Name.ValueStringPointer(),
		SshPublicKey:                     data.SshPublicKey.ValueStringPointer(),
		SystemVolumeSize:                 data.SystemVolumeSize.ValueInt64Pointer(),
		SystemVolumeType:                 data.SystemVolumeType.ValueStringPointer(),
		UserData:                         data.UserData.ValueStringPointer(),
	}
	{
		var items []common.OpenStackDataVolumeRequest
		diags := data.DataVolumes.ElementsAs(ctx, &items, false)
		resp.Diagnostics.Append(diags...)
		if !diags.HasError() {
			if !data.DataVolumes.IsNull() && !data.DataVolumes.IsUnknown() {
				attributes.DataVolumes = &items
			}
		}
	}
	{
		var items []common.OpenStackCreateFloatingIPRequest
		diags := data.FloatingIps.ElementsAs(ctx, &items, false)
		resp.Diagnostics.Append(diags...)
		if !diags.HasError() {
			if !data.FloatingIps.IsNull() && !data.FloatingIps.IsUnknown() {
				attributes.FloatingIps = &items
			}
		}
	}
	{
		var items []common.OpenStackCreateInstancePortRequest
		diags := data.Ports.ElementsAs(ctx, &items, false)
		resp.Diagnostics.Append(diags...)
		if !diags.HasError() {
			if !data.Ports.IsNull() && !data.Ports.IsUnknown() {
				attributes.Ports = &items
			}
		}
	}
	{
		var items []common.OpenStackSecurityGroupHyperlinkRequest
		diags := data.SecurityGroups.ElementsAs(ctx, &items, false)
		resp.Diagnostics.Append(diags...)
		if !diags.HasError() {
			if !data.SecurityGroups.IsNull() && !data.SecurityGroups.IsUnknown() {
				attributes.SecurityGroups = &items
			}
		}
	}

	// Construct the Create Order Request
	payload := OpenstackInstanceCreateRequest{
		Project:    data.Project.ValueStringPointer(),
		Offering:   data.Offering.ValueStringPointer(),
		Attributes: attributes,
	}

	// Phase 2: Submit Order
	orderRes, err := r.client.CreateOpenstackInstanceOrder(ctx, &payload)
	if err != nil {
		resp.Diagnostics.AddError("Order Submission Failed", err.Error())
		return
	}

	// Phase 3: Poll for Completion
	// We use the 'time' package to handle the timeout specified in the TF config or default to 45m.
	timeout, diags := data.Timeouts.Create(ctx, 45*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Wait for the order to reach a terminal state (done/erred)
	finalOrder, err := common.WaitForOrder(ctx, r.client.Client, *orderRes.Uuid, timeout)
	if err != nil {
		resp.Diagnostics.AddError("Order Failed", err.Error())
		return
	}

	// Resolve the created Resource UUID from the completed order
	if uuid := common.ResolveResourceUUID(finalOrder); uuid != "" {
		data.UUID = types.StringValue(uuid)
	} else {
		resp.Diagnostics.AddError("Resource UUID Missing", "Order completed but resource UUID is missing")
		return
	}

	// Fetch final resource state to ensure Terraform state matches reality
	apiResp, err := r.client.GetOpenstackInstance(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to Read Resource", err.Error())
		return
	}
	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackInstanceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpenstackInstanceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource

	apiResp, err := r.client.GetOpenstackInstance(ctx, data.UUID.ValueString())
	if err != nil {
		if client.IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Unable to Read Openstack Instance",
			"An error occurred while reading the Openstack Instance: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackInstanceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpenstackInstanceResourceModel
	var state OpenstackInstanceResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Phase 1: Standard PATCH (Simple fields)
	// We compare the plan (data) with the state (state) to determine which fields changed.
	var patchPayload OpenstackInstanceUpdateRequest
	if !data.Description.IsNull() && !data.Description.Equal(state.Description) {
		patchPayload.Description = data.Description.ValueStringPointer()
	}
	if !data.Name.IsNull() && !data.Name.Equal(state.Name) {
		patchPayload.Name = data.Name.ValueStringPointer()
	}

	{
		// Execute the PATCH request
		_, err := r.client.UpdateOpenstackInstance(ctx, data.UUID.ValueString(), &patchPayload)
		if err != nil {
			resp.Diagnostics.AddError("Update Failed", err.Error())
			return
		}
	}

	// Phase 2: RPC Actions
	// These actions are triggered when their corresponding specific fields change.
	if !data.FloatingIps.Equal(state.FloatingIps) {
		// Convert Terraform value to API payload for the specific action
		var req OpenstackInstanceUpdateFloatingIpsActionRequest
		common.PopulateSliceField(ctx, data.FloatingIps, &req.FloatingIps)

		// Execute the Action
		if err := r.client.OpenstackInstanceUpdateFloatingIps(ctx, data.UUID.ValueString(), &req); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: update_floating_ips", err.Error())
			return
		}
	}
	if !data.Ports.Equal(state.Ports) {
		// Convert Terraform value to API payload for the specific action
		var req OpenstackInstanceUpdatePortsActionRequest
		common.PopulateSliceField(ctx, data.Ports, &req.Ports)

		// Execute the Action
		if err := r.client.OpenstackInstanceUpdatePorts(ctx, data.UUID.ValueString(), &req); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: update_ports", err.Error())
			return
		}
	}
	if !data.SecurityGroups.Equal(state.SecurityGroups) {
		// Convert Terraform value to API payload for the specific action
		var req OpenstackInstanceUpdateSecurityGroupsActionRequest
		common.PopulateSliceField(ctx, data.SecurityGroups, &req.SecurityGroups)

		// Execute the Action
		if err := r.client.OpenstackInstanceUpdateSecurityGroups(ctx, data.UUID.ValueString(), &req); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: update_security_groups", err.Error())
			return
		}
	}

	// Fetch updated state after all changes
	apiResp, err := r.client.GetOpenstackInstance(ctx, data.UUID.ValueString())
	if err != nil {
		if client.IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Failed to Read Resource After Update", err.Error())
		return
	}
	resp.Diagnostics.Append(r.mapResponseToModel(ctx, *apiResp, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackInstanceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackInstanceResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Order-based Delete
	payload := map[string]interface{}{}
	if !data.DeleteVolumes.IsNull() {
		payload["delete_volumes"] = data.DeleteVolumes.ValueBool()
	}
	if !data.ReleaseFloatingIps.IsNull() {
		payload["release_floating_ips"] = data.ReleaseFloatingIps.ValueBool()
	}

	// Submit termination order
	orderUUID, err := r.client.TerminateOpenstackInstance(ctx, data.UUID.ValueString(), payload)
	if err != nil {
		resp.Diagnostics.AddError("Termination Failed", err.Error())
		return
	}

	// Wait for deletion if order UUID is returned
	if orderUUID != "" {
		timeout, diags := data.Timeouts.Delete(ctx, 45*time.Minute)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		_, err := common.WaitForOrder(ctx, r.client.Client, orderUUID, timeout)
		if err != nil {
			resp.Diagnostics.AddError("Termination Order Failed", err.Error())
			return
		}
	}
}

func (r *OpenstackInstanceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *OpenstackInstanceResource) mapResponseToModel(ctx context.Context, apiResp OpenstackInstanceResponse, model *OpenstackInstanceResourceModel) diag.Diagnostics {
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
	listValExternalAddress, listDiagsExternalAddress := types.ListValueFrom(ctx, types.StringType, apiResp.ExternalAddress)
	model.ExternalAddress = listValExternalAddress
	diags.Append(listDiagsExternalAddress...)
	listValExternalIps, listDiagsExternalIps := types.ListValueFrom(ctx, types.StringType, apiResp.ExternalIps)
	model.ExternalIps = listValExternalIps
	diags.Append(listDiagsExternalIps...)
	model.FlavorDisk = types.Int64PointerValue(apiResp.FlavorDisk)
	model.FlavorName = types.StringPointerValue(apiResp.FlavorName)

	{
		listValFloatingIps, listDiagsFloatingIps := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: map[string]attr.Type{
			"ip_address": types.StringType,
			"subnet":     types.StringType,
			"url":        types.StringType,
			"address":    types.StringType,
			"port_fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
				"ip_address": types.StringType,
				"subnet_id":  types.StringType,
			}}},
			"port_mac_address":   types.StringType,
			"subnet_cidr":        types.StringType,
			"subnet_description": types.StringType,
			"subnet_name":        types.StringType,
			"subnet_uuid":        types.StringType,
		}}, apiResp.FloatingIps)
		diags.Append(listDiagsFloatingIps...)
		model.FloatingIps = listValFloatingIps
	}
	model.HypervisorHostname = types.StringPointerValue(apiResp.HypervisorHostname)
	model.ImageName = types.StringPointerValue(apiResp.ImageName)
	listValInternalIps, listDiagsInternalIps := types.ListValueFrom(ctx, types.StringType, apiResp.InternalIps)
	model.InternalIps = listValInternalIps
	diags.Append(listDiagsInternalIps...)
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
			"fixed_ips": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
				"ip_address": types.StringType,
				"subnet_id":  types.StringType,
			}}},
			"port":   types.StringType,
			"subnet": types.StringType,
			"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
				"mac_address": types.StringType,
			}}},
			"device_id":    types.StringType,
			"device_owner": types.StringType,
			"mac_address":  types.StringType,
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
			"url":         types.StringType,
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
