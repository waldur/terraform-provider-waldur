package resources

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-waldur-provider/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OpenstackInstanceResource{}
var _ resource.ResourceWithImportState = &OpenstackInstanceResource{}

func NewOpenstackInstanceResource() resource.Resource {
	return &OpenstackInstanceResource{}
}

// OpenstackInstanceResource defines the resource implementation.
type OpenstackInstanceResource struct {
	client *client.Client
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
		MarkdownDescription: "OpenstackInstance resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"access_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"action": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"availability_zone": schema.StringAttribute{
				Optional:            true,
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
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "If True, instance will be connected directly to external network",
			},
			"cores": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Number of cores in a VM",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"customer": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"customer_abbreviation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"customer_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"customer_native_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"customer_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"data_volume_size": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Size of the data volume in MiB. Minimum size is 1024 MiB (1 GiB)",
			},
			"data_volume_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Volume type for the data volume",
			},
			"data_volumes": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"size": schema.Int64Attribute{
							Required:            true,
							MarkdownDescription: "",
						},
						"volume_type": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "",
						},
					},
				},
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Additional data volumes to attach to the instance",
			},
			"delete_volumes": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Termination attribute",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"disk": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Disk size in MiB",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"external_address": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.StringType},
				Computed:            true,
				MarkdownDescription: "",
			},
			"external_ips": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.StringType},
				Computed:            true,
				MarkdownDescription: "",
			},
			"flavor": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The flavor to use for the instance",
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
						"ip_address": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Existing floating IP address in selected OpenStack tenant to be assigned to new virtual machine",
						},
						"subnet": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "",
						},
						"url": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "",
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
										Computed:            true,
										MarkdownDescription: "IP address to assign to the port",
									},
									"subnet_id": schema.StringAttribute{
										Optional:            true,
										Computed:            true,
										MarkdownDescription: "ID of the subnet in which to assign the IP address",
									},
								},
							},
							Computed:            true,
							MarkdownDescription: "",
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
							MarkdownDescription: "",
						},
						"subnet_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
						"subnet_uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
					},
				},
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Floating IPs to assign to the instance",
			},
			"hypervisor_hostname": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the hypervisor hosting this instance",
			},
			"image": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The OS image to use for the instance",
			},
			"image_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"internal_ips": schema.ListAttribute{
				CustomType:          types.ListType{ElemType: types.StringType},
				Computed:            true,
				MarkdownDescription: "",
			},
			"is_limit_based": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"is_usage_based": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"key_fingerprint": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"key_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"latitude": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"longitude": schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"marketplace_category_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"marketplace_category_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"marketplace_offering_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"marketplace_offering_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"marketplace_plan_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"marketplace_resource_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
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
				MarkdownDescription: "",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"offering": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Offering UUID or URL",
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
							Computed:            true,
							MarkdownDescription: "",
						},
						"port": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "",
						},
						"subnet": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Subnet to which this port belongs",
						},
						"allowed_address_pairs": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"mac_address": schema.StringAttribute{
										Optional:            true,
										Computed:            true,
										MarkdownDescription: "",
									},
								},
							},
							Computed:            true,
							MarkdownDescription: "",
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
										MarkdownDescription: "",
									},
									"backend_id": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"created": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"customer": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"customer_abbreviation": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"customer_name": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"customer_native_name": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"customer_uuid": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"description": schema.StringAttribute{
										Optional:            true,
										Computed:            true,
										MarkdownDescription: "",
									},
									"error_message": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"error_traceback": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"is_limit_based": schema.BoolAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"is_usage_based": schema.BoolAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"marketplace_category_name": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"marketplace_category_uuid": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"marketplace_offering_name": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"marketplace_offering_uuid": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"marketplace_plan_uuid": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"marketplace_resource_state": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"marketplace_resource_uuid": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"modified": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"name": schema.StringAttribute{
										Optional:            true,
										Computed:            true,
										MarkdownDescription: "",
									},
									"project": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										MarkdownDescription: "",
									},
									"project_name": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"project_uuid": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"resource_type": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"rules": schema.ListNestedAttribute{
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"cidr": schema.StringAttribute{
													Optional:            true,
													Computed:            true,
													MarkdownDescription: "CIDR notation for the source/destination network address range",
												},
												"description": schema.StringAttribute{
													Optional:            true,
													Computed:            true,
													MarkdownDescription: "",
												},
												"direction": schema.StringAttribute{
													Optional:            true,
													Computed:            true,
													MarkdownDescription: "Traffic direction - either 'ingress' (incoming) or 'egress' (outgoing)",
												},
												"ethertype": schema.StringAttribute{
													Optional:            true,
													Computed:            true,
													MarkdownDescription: "IP protocol version - either 'IPv4' or 'IPv6'",
												},
												"from_port": schema.Int64Attribute{
													Optional:            true,
													Computed:            true,
													MarkdownDescription: "Starting port number in the range (1-65535)",
												},
												"id": schema.Int64Attribute{
													Computed:            true,
													MarkdownDescription: "",
												},
												"protocol": schema.StringAttribute{
													Optional:            true,
													Computed:            true,
													MarkdownDescription: "The network protocol (TCP, UDP, ICMP, or empty for any protocol)",
												},
												"remote_group": schema.StringAttribute{
													Optional:            true,
													Computed:            true,
													MarkdownDescription: "Remote security group that this rule references, if any",
												},
												"remote_group_name": schema.StringAttribute{
													Computed:            true,
													MarkdownDescription: "",
												},
												"remote_group_uuid": schema.StringAttribute{
													Computed:            true,
													MarkdownDescription: "",
												},
												"to_port": schema.Int64Attribute{
													Optional:            true,
													Computed:            true,
													MarkdownDescription: "Ending port number in the range (1-65535)",
												},
											},
										},
										Optional:            true,
										Computed:            true,
										MarkdownDescription: "",
									},
									"service_name": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"service_settings": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"service_settings_error_message": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"service_settings_state": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"service_settings_uuid": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"state": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"tenant": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"tenant_name": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"tenant_uuid": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"url": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
								},
							},
							Computed:            true,
							MarkdownDescription: "",
						},
						"subnet_cidr": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "IPv4 network address in CIDR format (e.g. 192.168.0.0/24)",
						},
						"subnet_description": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
						"subnet_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
						"subnet_uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
					},
				},
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Network ports to attach to the instance",
			},
			"project": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "",
			},
			"project_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"project_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"ram": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Memory size in MiB",
			},
			"release_floating_ips": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Termination attribute",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"runtime_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"security_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"url": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "",
						},
						"description": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
						"rules": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"cidr": schema.StringAttribute{
										Optional:            true,
										Computed:            true,
										MarkdownDescription: "CIDR notation for the source/destination network address range",
									},
									"description": schema.StringAttribute{
										Optional:            true,
										Computed:            true,
										MarkdownDescription: "",
									},
									"direction": schema.StringAttribute{
										Optional:            true,
										Computed:            true,
										MarkdownDescription: "Traffic direction - either 'ingress' (incoming) or 'egress' (outgoing)",
									},
									"ethertype": schema.StringAttribute{
										Optional:            true,
										Computed:            true,
										MarkdownDescription: "IP protocol version - either 'IPv4' or 'IPv6'",
									},
									"from_port": schema.Int64Attribute{
										Optional:            true,
										Computed:            true,
										MarkdownDescription: "Starting port number in the range (1-65535)",
									},
									"id": schema.Int64Attribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"protocol": schema.StringAttribute{
										Optional:            true,
										Computed:            true,
										MarkdownDescription: "The network protocol (TCP, UDP, ICMP, or empty for any protocol)",
									},
									"remote_group_name": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"remote_group_uuid": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "",
									},
									"to_port": schema.Int64Attribute{
										Optional:            true,
										Computed:            true,
										MarkdownDescription: "Ending port number in the range (1-65535)",
									},
								},
							},
							Computed:            true,
							MarkdownDescription: "",
						},
						"state": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
					},
				},
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "List of security groups to apply to the instance",
			},
			"server_group": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "",
					},
					"policy": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "Server group policy determining the rules for scheduling servers in this group",
					},
					"state": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "",
					},
					"url": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: "",
					},
				},
				Computed:            true,
				MarkdownDescription: "",
			},
			"service_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"service_settings": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "OpenStack provider settings",
			},
			"service_settings_error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"service_settings_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"service_settings_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"ssh_public_key": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"start_time": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"system_volume_size": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Size of the system volume in MiB. Minimum size is 1024 MiB (1 GiB)",
			},
			"system_volume_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Volume type for the system volume",
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
				MarkdownDescription: "",
			},
			"user_data": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Additional data that will be added to instance on provisioning",
			},
			"volumes": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"bootable": schema.BoolAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Indicates if this volume can be used to boot an instance",
						},
						"device": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Name of volume as instance device e.g. /dev/vdb.",
						},
						"image_name": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Name of the image this volume was created from",
						},
						"marketplace_resource_uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
						"resource_type": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
						"size": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Size in MiB",
						},
						"state": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
						"type": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Type of the volume (e.g. SSD, HDD)",
						},
						"type_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
						"url": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "",
						},
					},
				},
				Computed:            true,
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

func (r *OpenstackInstanceResource) convertTFValue(v attr.Value) interface{} {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	switch val := v.(type) {
	case types.String:
		return val.ValueString()
	case types.Int64:
		return val.ValueInt64()
	case types.Bool:
		return val.ValueBool()
	case types.Float64:
		return val.ValueFloat64()
	case types.List:
		items := make([]interface{}, len(val.Elements()))
		for i, elem := range val.Elements() {
			items[i] = r.convertTFValue(elem)
		}
		return items
	case types.Object:
		obj := make(map[string]interface{})
		for k, attr := range val.Attributes() {
			if converted := r.convertTFValue(attr); converted != nil {
				obj[k] = converted
			}
		}
		return obj
	}
	return nil
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

	r.client = client
}

func (r *OpenstackInstanceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackInstanceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Phase 1: Payload Construction
	attributes := map[string]interface{}{}
	if !data.AvailabilityZone.IsNull() {
		attributes["availability_zone"] = data.AvailabilityZone.ValueString()
	}
	if !data.ConnectDirectlyToExternalNetwork.IsNull() {
		attributes["connect_directly_to_external_network"] = data.ConnectDirectlyToExternalNetwork.ValueBool()
	}
	if !data.DataVolumeSize.IsNull() {
		attributes["data_volume_size"] = data.DataVolumeSize.ValueInt64()
	}
	if !data.DataVolumeType.IsNull() {
		attributes["data_volume_type"] = data.DataVolumeType.ValueString()
	}
	if !data.DataVolumes.IsNull() {
		items := make([]interface{}, 0)
		for _, elem := range data.DataVolumes.Elements() {
			items = append(items, r.convertTFValue(elem))
		}
		attributes["data_volumes"] = items
	}
	if !data.Description.IsNull() {
		attributes["description"] = data.Description.ValueString()
	}
	if !data.Flavor.IsNull() {
		attributes["flavor"] = data.Flavor.ValueString()
	}
	if !data.FloatingIps.IsNull() {
		items := make([]interface{}, 0)
		for _, elem := range data.FloatingIps.Elements() {
			items = append(items, r.convertTFValue(elem))
		}
		attributes["floating_ips"] = items
	}
	if !data.Image.IsNull() {
		attributes["image"] = data.Image.ValueString()
	}
	if !data.Name.IsNull() {
		attributes["name"] = data.Name.ValueString()
	}
	if !data.Ports.IsNull() {
		items := make([]interface{}, 0)
		for _, elem := range data.Ports.Elements() {
			items = append(items, r.convertTFValue(elem))
		}
		attributes["ports"] = items
	}
	if !data.SecurityGroups.IsNull() {
		items := make([]interface{}, 0)
		for _, elem := range data.SecurityGroups.Elements() {
			items = append(items, r.convertTFValue(elem))
		}
		attributes["security_groups"] = items
	}
	if !data.SshPublicKey.IsNull() {
		attributes["ssh_public_key"] = data.SshPublicKey.ValueString()
	}
	if !data.SystemVolumeSize.IsNull() {
		attributes["system_volume_size"] = data.SystemVolumeSize.ValueInt64()
	}
	if !data.SystemVolumeType.IsNull() {
		attributes["system_volume_type"] = data.SystemVolumeType.ValueString()
	}
	if !data.UserData.IsNull() {
		attributes["user_data"] = data.UserData.ValueString()
	}

	payload := map[string]interface{}{
		"project":    data.Project.ValueString(),
		"offering":   data.Offering.ValueString(), // Assuming offering is passed as URL or UUID handled by API
		"attributes": attributes,
	}

	// Phase 2: Submit Order
	var orderRes map[string]interface{}
	err := r.client.Post(ctx, "/api/marketplace-orders/", payload, &orderRes)
	if err != nil {
		resp.Diagnostics.AddError("Order Submission Failed", err.Error())
		return
	}

	orderUUID, ok := orderRes["uuid"].(string)
	if !ok {
		resp.Diagnostics.AddError("Invalid Response", "Order UUID not found")
		return
	}

	// Phase 3: Poll for Completion
	// Attempt to resolve UUID
	if uuid, ok := orderRes["resource_uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	} else {
		data.UUID = types.StringValue(orderUUID)
	}

	// Attempt to fetch the resource to populate state
	{
		var mpUUID string
		if uuid, ok := orderRes["resource_uuid"].(string); ok {
			mpUUID = uuid
		} else if uuid, ok := orderRes["marketplace_resource_uuid"].(string); ok {
			mpUUID = uuid
		}

		if mpUUID != "" {
			var mpRes map[string]interface{}
			err = r.client.GetByUUID(ctx, "/api/marketplace-resources/{uuid}/", mpUUID, &mpRes)
			if err == nil {
				// Debug logging
				tflog.Warn(ctx, fmt.Sprintf("Fetched MP Resource: %+v", mpRes))
				if val, exists := mpRes["resource_uuid"]; exists {
					tflog.Warn(ctx, fmt.Sprintf("resource_uuid type: %T, value: %v", val, val))
				} else {
					tflog.Warn(ctx, "resource_uuid key missing in MP response")
				}

				// Plugin Resource UUID is available directly in resource_uuid field
				if pluginUUID, ok := mpRes["resource_uuid"].(string); ok {
					if pluginUUID != "" {
						data.UUID = types.StringValue(pluginUUID)

						// Fetch Plugin Resource
						var pluginRes map[string]interface{}
						retrievePath := strings.Replace("/api/openstack-instances/{uuid}/", "{uuid}", pluginUUID, 1)
						tflog.Warn(ctx, "Attempting to fetch plugin resource at: "+retrievePath)
						err = r.client.GetByUUID(ctx, retrievePath, pluginUUID, &pluginRes)
						if err == nil {
							tflog.Warn(ctx, "Successfully fetched plugin resource")
							sourceMap := pluginRes
							// Map response fields to data model
							_ = sourceMap
							if val, ok := sourceMap["access_url"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.AccessUrl = types.StringValue(str)
								}
							} else {
								if data.AccessUrl.IsUnknown() {
									data.AccessUrl = types.StringNull()
								}
							}
							if val, ok := sourceMap["action"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.Action = types.StringValue(str)
								}
							} else {
								if data.Action.IsUnknown() {
									data.Action = types.StringNull()
								}
							}
							if val, ok := sourceMap["availability_zone"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.AvailabilityZone = types.StringValue(str)
								}
							} else {
								if data.AvailabilityZone.IsUnknown() {
									data.AvailabilityZone = types.StringNull()
								}
							}
							if val, ok := sourceMap["availability_zone_name"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.AvailabilityZoneName = types.StringValue(str)
								}
							} else {
								if data.AvailabilityZoneName.IsUnknown() {
									data.AvailabilityZoneName = types.StringNull()
								}
							}
							if val, ok := sourceMap["backend_id"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.BackendId = types.StringValue(str)
								}
							} else {
								if data.BackendId.IsUnknown() {
									data.BackendId = types.StringNull()
								}
							}
							if val, ok := sourceMap["connect_directly_to_external_network"]; ok && val != nil {
								if b, ok := val.(bool); ok {
									data.ConnectDirectlyToExternalNetwork = types.BoolValue(b)
								}
							} else {
								if data.ConnectDirectlyToExternalNetwork.IsUnknown() {
									data.ConnectDirectlyToExternalNetwork = types.BoolNull()
								}
							}
							if val, ok := sourceMap["cores"]; ok && val != nil {
								if num, ok := val.(float64); ok {
									data.Cores = types.Int64Value(int64(num))
								}
							} else {
								if data.Cores.IsUnknown() {
									data.Cores = types.Int64Null()
								}
							}
							if val, ok := sourceMap["created"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.Created = types.StringValue(str)
								}
							} else {
								if data.Created.IsUnknown() {
									data.Created = types.StringNull()
								}
							}
							if val, ok := sourceMap["customer"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.Customer = types.StringValue(str)
								}
							} else {
								if data.Customer.IsUnknown() {
									data.Customer = types.StringNull()
								}
							}
							if val, ok := sourceMap["customer_abbreviation"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.CustomerAbbreviation = types.StringValue(str)
								}
							} else {
								if data.CustomerAbbreviation.IsUnknown() {
									data.CustomerAbbreviation = types.StringNull()
								}
							}
							if val, ok := sourceMap["customer_name"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.CustomerName = types.StringValue(str)
								}
							} else {
								if data.CustomerName.IsUnknown() {
									data.CustomerName = types.StringNull()
								}
							}
							if val, ok := sourceMap["customer_native_name"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.CustomerNativeName = types.StringValue(str)
								}
							} else {
								if data.CustomerNativeName.IsUnknown() {
									data.CustomerNativeName = types.StringNull()
								}
							}
							if val, ok := sourceMap["customer_uuid"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.CustomerUuid = types.StringValue(str)
								}
							} else {
								if data.CustomerUuid.IsUnknown() {
									data.CustomerUuid = types.StringNull()
								}
							}
							if val, ok := sourceMap["data_volume_size"]; ok && val != nil {
								if num, ok := val.(float64); ok {
									data.DataVolumeSize = types.Int64Value(int64(num))
								}
							} else {
								if data.DataVolumeSize.IsUnknown() {
									data.DataVolumeSize = types.Int64Null()
								}
							}
							if val, ok := sourceMap["data_volume_type"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.DataVolumeType = types.StringValue(str)
								}
							} else {
								if data.DataVolumeType.IsUnknown() {
									data.DataVolumeType = types.StringNull()
								}
							}
							if val, ok := sourceMap["data_volumes"]; ok && val != nil {
								// List of objects
								if arr, ok := val.([]interface{}); ok {
									items := make([]attr.Value, 0, len(arr))
									for _, item := range arr {
										if objMap, ok := item.(map[string]interface{}); ok {
											attrTypes := map[string]attr.Type{
												"size":        types.Int64Type,
												"volume_type": types.StringType,
											}
											attrValues := map[string]attr.Value{
												"size": func() attr.Value {
													if v, ok := objMap["size"].(float64); ok {
														return types.Int64Value(int64(v))
													}
													return types.Int64Null()
												}(),
												"volume_type": func() attr.Value {
													if v, ok := objMap["volume_type"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
											}
											objVal, _ := types.ObjectValue(attrTypes, attrValues)
											items = append(items, objVal)
										}
									}
									listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
										"size":        types.Int64Type,
										"volume_type": types.StringType,
									}}, items)
									data.DataVolumes = listVal
								}
							} else {
								if data.DataVolumes.IsUnknown() {
									data.DataVolumes = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
										"size":        types.Int64Type,
										"volume_type": types.StringType,
									}})
								}
							}
							if val, ok := sourceMap["delete_volumes"]; ok && val != nil {
								if b, ok := val.(bool); ok {
									data.DeleteVolumes = types.BoolValue(b)
								}
							} else {
								if data.DeleteVolumes.IsUnknown() {
									data.DeleteVolumes = types.BoolNull()
								}
							}
							if val, ok := sourceMap["description"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.Description = types.StringValue(str)
								}
							} else {
								if data.Description.IsUnknown() {
									data.Description = types.StringNull()
								}
							}
							if val, ok := sourceMap["disk"]; ok && val != nil {
								if num, ok := val.(float64); ok {
									data.Disk = types.Int64Value(int64(num))
								}
							} else {
								if data.Disk.IsUnknown() {
									data.Disk = types.Int64Null()
								}
							}
							if val, ok := sourceMap["error_message"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.ErrorMessage = types.StringValue(str)
								}
							} else {
								if data.ErrorMessage.IsUnknown() {
									data.ErrorMessage = types.StringNull()
								}
							}
							if val, ok := sourceMap["error_traceback"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.ErrorTraceback = types.StringValue(str)
								}
							} else {
								if data.ErrorTraceback.IsUnknown() {
									data.ErrorTraceback = types.StringNull()
								}
							}
							if val, ok := sourceMap["external_address"]; ok && val != nil {
								// List of strings (or flattened objects)
								if arr, ok := val.([]interface{}); ok {
									items := make([]attr.Value, 0, len(arr))
									for _, item := range arr {
										if str, ok := item.(string); ok {
											items = append(items, types.StringValue(str))
										} else if obj, ok := item.(map[string]interface{}); ok {
											// Flattening logic: extract URL or UUID
											if url, ok := obj["url"].(string); ok {
												parts := strings.Split(strings.TrimRight(url, "/"), "/")
												uuid := parts[len(parts)-1]
												items = append(items, types.StringValue(uuid))
											} else if uuid, ok := obj["uuid"].(string); ok {
												items = append(items, types.StringValue(uuid))
											} else if name, ok := obj["name"].(string); ok {
												items = append(items, types.StringValue(name))
											}
										}
									}
									listVal, _ := types.ListValue(types.StringType, items)
									data.ExternalAddress = listVal
								}
							} else {
								if data.ExternalAddress.IsUnknown() {
									data.ExternalAddress = types.ListNull(types.StringType)
								}
							}
							if val, ok := sourceMap["external_ips"]; ok && val != nil {
								// List of strings (or flattened objects)
								if arr, ok := val.([]interface{}); ok {
									items := make([]attr.Value, 0, len(arr))
									for _, item := range arr {
										if str, ok := item.(string); ok {
											items = append(items, types.StringValue(str))
										} else if obj, ok := item.(map[string]interface{}); ok {
											// Flattening logic: extract URL or UUID
											if url, ok := obj["url"].(string); ok {
												parts := strings.Split(strings.TrimRight(url, "/"), "/")
												uuid := parts[len(parts)-1]
												items = append(items, types.StringValue(uuid))
											} else if uuid, ok := obj["uuid"].(string); ok {
												items = append(items, types.StringValue(uuid))
											} else if name, ok := obj["name"].(string); ok {
												items = append(items, types.StringValue(name))
											}
										}
									}
									listVal, _ := types.ListValue(types.StringType, items)
									data.ExternalIps = listVal
								}
							} else {
								if data.ExternalIps.IsUnknown() {
									data.ExternalIps = types.ListNull(types.StringType)
								}
							}
							if val, ok := sourceMap["flavor"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.Flavor = types.StringValue(str)
								}
							} else {
								if data.Flavor.IsUnknown() {
									data.Flavor = types.StringNull()
								}
							}
							if val, ok := sourceMap["flavor_disk"]; ok && val != nil {
								if num, ok := val.(float64); ok {
									data.FlavorDisk = types.Int64Value(int64(num))
								}
							} else {
								if data.FlavorDisk.IsUnknown() {
									data.FlavorDisk = types.Int64Null()
								}
							}
							if val, ok := sourceMap["flavor_name"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.FlavorName = types.StringValue(str)
								}
							} else {
								if data.FlavorName.IsUnknown() {
									data.FlavorName = types.StringNull()
								}
							}
							if val, ok := sourceMap["floating_ips"]; ok && val != nil {
								// List of objects
								if arr, ok := val.([]interface{}); ok {
									items := make([]attr.Value, 0, len(arr))
									for _, item := range arr {
										if objMap, ok := item.(map[string]interface{}); ok {
											attrTypes := map[string]attr.Type{
												"ip_address":         types.StringType,
												"subnet":             types.StringType,
												"url":                types.StringType,
												"address":            types.StringType,
												"port_fixed_ips":     types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
												"port_mac_address":   types.StringType,
												"subnet_cidr":        types.StringType,
												"subnet_description": types.StringType,
												"subnet_name":        types.StringType,
												"subnet_uuid":        types.StringType,
											}
											attrValues := map[string]attr.Value{
												"ip_address": func() attr.Value {
													if v, ok := objMap["ip_address"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"subnet": func() attr.Value {
													if v, ok := objMap["subnet"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"url": func() attr.Value {
													if v, ok := objMap["url"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"address": func() attr.Value {
													if v, ok := objMap["address"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"port_fixed_ips": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}}.ElemType),
												"port_mac_address": func() attr.Value {
													if v, ok := objMap["port_mac_address"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"subnet_cidr": func() attr.Value {
													if v, ok := objMap["subnet_cidr"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"subnet_description": func() attr.Value {
													if v, ok := objMap["subnet_description"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"subnet_name": func() attr.Value {
													if v, ok := objMap["subnet_name"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"subnet_uuid": func() attr.Value {
													if v, ok := objMap["subnet_uuid"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
											}
											objVal, _ := types.ObjectValue(attrTypes, attrValues)
											items = append(items, objVal)
										}
									}
									listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
										"ip_address":         types.StringType,
										"subnet":             types.StringType,
										"url":                types.StringType,
										"address":            types.StringType,
										"port_fixed_ips":     types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
										"port_mac_address":   types.StringType,
										"subnet_cidr":        types.StringType,
										"subnet_description": types.StringType,
										"subnet_name":        types.StringType,
										"subnet_uuid":        types.StringType,
									}}, items)
									data.FloatingIps = listVal
								}
							} else {
								if data.FloatingIps.IsUnknown() {
									data.FloatingIps = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
										"ip_address":         types.StringType,
										"subnet":             types.StringType,
										"url":                types.StringType,
										"address":            types.StringType,
										"port_fixed_ips":     types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
										"port_mac_address":   types.StringType,
										"subnet_cidr":        types.StringType,
										"subnet_description": types.StringType,
										"subnet_name":        types.StringType,
										"subnet_uuid":        types.StringType,
									}})
								}
							}
							if val, ok := sourceMap["hypervisor_hostname"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.HypervisorHostname = types.StringValue(str)
								}
							} else {
								if data.HypervisorHostname.IsUnknown() {
									data.HypervisorHostname = types.StringNull()
								}
							}
							if val, ok := sourceMap["image"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.Image = types.StringValue(str)
								}
							} else {
								if data.Image.IsUnknown() {
									data.Image = types.StringNull()
								}
							}
							if val, ok := sourceMap["image_name"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.ImageName = types.StringValue(str)
								}
							} else {
								if data.ImageName.IsUnknown() {
									data.ImageName = types.StringNull()
								}
							}
							if val, ok := sourceMap["internal_ips"]; ok && val != nil {
								// List of strings (or flattened objects)
								if arr, ok := val.([]interface{}); ok {
									items := make([]attr.Value, 0, len(arr))
									for _, item := range arr {
										if str, ok := item.(string); ok {
											items = append(items, types.StringValue(str))
										} else if obj, ok := item.(map[string]interface{}); ok {
											// Flattening logic: extract URL or UUID
											if url, ok := obj["url"].(string); ok {
												parts := strings.Split(strings.TrimRight(url, "/"), "/")
												uuid := parts[len(parts)-1]
												items = append(items, types.StringValue(uuid))
											} else if uuid, ok := obj["uuid"].(string); ok {
												items = append(items, types.StringValue(uuid))
											} else if name, ok := obj["name"].(string); ok {
												items = append(items, types.StringValue(name))
											}
										}
									}
									listVal, _ := types.ListValue(types.StringType, items)
									data.InternalIps = listVal
								}
							} else {
								if data.InternalIps.IsUnknown() {
									data.InternalIps = types.ListNull(types.StringType)
								}
							}
							if val, ok := sourceMap["is_limit_based"]; ok && val != nil {
								if b, ok := val.(bool); ok {
									data.IsLimitBased = types.BoolValue(b)
								}
							} else {
								if data.IsLimitBased.IsUnknown() {
									data.IsLimitBased = types.BoolNull()
								}
							}
							if val, ok := sourceMap["is_usage_based"]; ok && val != nil {
								if b, ok := val.(bool); ok {
									data.IsUsageBased = types.BoolValue(b)
								}
							} else {
								if data.IsUsageBased.IsUnknown() {
									data.IsUsageBased = types.BoolNull()
								}
							}
							if val, ok := sourceMap["key_fingerprint"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.KeyFingerprint = types.StringValue(str)
								}
							} else {
								if data.KeyFingerprint.IsUnknown() {
									data.KeyFingerprint = types.StringNull()
								}
							}
							if val, ok := sourceMap["key_name"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.KeyName = types.StringValue(str)
								}
							} else {
								if data.KeyName.IsUnknown() {
									data.KeyName = types.StringNull()
								}
							}
							if val, ok := sourceMap["latitude"]; ok && val != nil {
								if num, ok := val.(float64); ok {
									data.Latitude = types.Float64Value(num)
								}
							} else {
								if data.Latitude.IsUnknown() {
									data.Latitude = types.Float64Null()
								}
							}
							if val, ok := sourceMap["longitude"]; ok && val != nil {
								if num, ok := val.(float64); ok {
									data.Longitude = types.Float64Value(num)
								}
							} else {
								if data.Longitude.IsUnknown() {
									data.Longitude = types.Float64Null()
								}
							}
							if val, ok := sourceMap["marketplace_category_name"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.MarketplaceCategoryName = types.StringValue(str)
								}
							} else {
								if data.MarketplaceCategoryName.IsUnknown() {
									data.MarketplaceCategoryName = types.StringNull()
								}
							}
							if val, ok := sourceMap["marketplace_category_uuid"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.MarketplaceCategoryUuid = types.StringValue(str)
								}
							} else {
								if data.MarketplaceCategoryUuid.IsUnknown() {
									data.MarketplaceCategoryUuid = types.StringNull()
								}
							}
							if val, ok := sourceMap["marketplace_offering_name"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.MarketplaceOfferingName = types.StringValue(str)
								}
							} else {
								if data.MarketplaceOfferingName.IsUnknown() {
									data.MarketplaceOfferingName = types.StringNull()
								}
							}
							if val, ok := sourceMap["marketplace_offering_uuid"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.MarketplaceOfferingUuid = types.StringValue(str)
								}
							} else {
								if data.MarketplaceOfferingUuid.IsUnknown() {
									data.MarketplaceOfferingUuid = types.StringNull()
								}
							}
							if val, ok := sourceMap["marketplace_plan_uuid"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.MarketplacePlanUuid = types.StringValue(str)
								}
							} else {
								if data.MarketplacePlanUuid.IsUnknown() {
									data.MarketplacePlanUuid = types.StringNull()
								}
							}
							if val, ok := sourceMap["marketplace_resource_state"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.MarketplaceResourceState = types.StringValue(str)
								}
							} else {
								if data.MarketplaceResourceState.IsUnknown() {
									data.MarketplaceResourceState = types.StringNull()
								}
							}
							if val, ok := sourceMap["marketplace_resource_uuid"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.MarketplaceResourceUuid = types.StringValue(str)
								}
							} else {
								if data.MarketplaceResourceUuid.IsUnknown() {
									data.MarketplaceResourceUuid = types.StringNull()
								}
							}
							if val, ok := sourceMap["min_disk"]; ok && val != nil {
								if num, ok := val.(float64); ok {
									data.MinDisk = types.Int64Value(int64(num))
								}
							} else {
								if data.MinDisk.IsUnknown() {
									data.MinDisk = types.Int64Null()
								}
							}
							if val, ok := sourceMap["min_ram"]; ok && val != nil {
								if num, ok := val.(float64); ok {
									data.MinRam = types.Int64Value(int64(num))
								}
							} else {
								if data.MinRam.IsUnknown() {
									data.MinRam = types.Int64Null()
								}
							}
							if val, ok := sourceMap["modified"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.Modified = types.StringValue(str)
								}
							} else {
								if data.Modified.IsUnknown() {
									data.Modified = types.StringNull()
								}
							}
							if val, ok := sourceMap["offering"]; ok && val != nil {
								if str, ok := val.(string); ok {
									// Normalize URL to UUID
									parts := strings.Split(strings.TrimRight(str, "/"), "/")
									uuid := parts[len(parts)-1]
									data.Offering = types.StringValue(uuid)
								} else {
									data.Offering = types.StringNull()
								}
							} else {
								if data.Offering.IsUnknown() {
									data.Offering = types.StringNull()
								}
							}
							if val, ok := sourceMap["ports"]; ok && val != nil {
								// List of objects
								if arr, ok := val.([]interface{}); ok {
									items := make([]attr.Value, 0, len(arr))
									for _, item := range arr {
										if objMap, ok := item.(map[string]interface{}); ok {
											attrTypes := map[string]attr.Type{
												"fixed_ips":             types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
												"port":                  types.StringType,
												"subnet":                types.StringType,
												"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"mac_address": types.StringType}}},
												"device_id":             types.StringType,
												"device_owner":          types.StringType,
												"mac_address":           types.StringType,
												"security_groups":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"access_url": types.StringType, "backend_id": types.StringType, "created": types.StringType, "customer": types.StringType, "customer_abbreviation": types.StringType, "customer_name": types.StringType, "customer_native_name": types.StringType, "customer_uuid": types.StringType, "description": types.StringType, "error_message": types.StringType, "error_traceback": types.StringType, "is_limit_based": types.BoolType, "is_usage_based": types.BoolType, "marketplace_category_name": types.StringType, "marketplace_category_uuid": types.StringType, "marketplace_offering_name": types.StringType, "marketplace_offering_uuid": types.StringType, "marketplace_plan_uuid": types.StringType, "marketplace_resource_state": types.StringType, "marketplace_resource_uuid": types.StringType, "modified": types.StringType, "name": types.StringType, "project": types.StringType, "project_name": types.StringType, "project_uuid": types.StringType, "resource_type": types.StringType, "rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}}, "service_name": types.StringType, "service_settings": types.StringType, "service_settings_error_message": types.StringType, "service_settings_state": types.StringType, "service_settings_uuid": types.StringType, "state": types.StringType, "tenant": types.StringType, "tenant_name": types.StringType, "tenant_uuid": types.StringType, "url": types.StringType}}},
												"subnet_cidr":           types.StringType,
												"subnet_description":    types.StringType,
												"subnet_name":           types.StringType,
												"subnet_uuid":           types.StringType,
												"url":                   types.StringType,
											}
											attrValues := map[string]attr.Value{
												"fixed_ips": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}}.ElemType),
												"port": func() attr.Value {
													if v, ok := objMap["port"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"subnet": func() attr.Value {
													if v, ok := objMap["subnet"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"allowed_address_pairs": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"mac_address": types.StringType}}}.ElemType),
												"device_id": func() attr.Value {
													if v, ok := objMap["device_id"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"device_owner": func() attr.Value {
													if v, ok := objMap["device_owner"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"mac_address": func() attr.Value {
													if v, ok := objMap["mac_address"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"security_groups": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"access_url": types.StringType, "backend_id": types.StringType, "created": types.StringType, "customer": types.StringType, "customer_abbreviation": types.StringType, "customer_name": types.StringType, "customer_native_name": types.StringType, "customer_uuid": types.StringType, "description": types.StringType, "error_message": types.StringType, "error_traceback": types.StringType, "is_limit_based": types.BoolType, "is_usage_based": types.BoolType, "marketplace_category_name": types.StringType, "marketplace_category_uuid": types.StringType, "marketplace_offering_name": types.StringType, "marketplace_offering_uuid": types.StringType, "marketplace_plan_uuid": types.StringType, "marketplace_resource_state": types.StringType, "marketplace_resource_uuid": types.StringType, "modified": types.StringType, "name": types.StringType, "project": types.StringType, "project_name": types.StringType, "project_uuid": types.StringType, "resource_type": types.StringType, "rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}}, "service_name": types.StringType, "service_settings": types.StringType, "service_settings_error_message": types.StringType, "service_settings_state": types.StringType, "service_settings_uuid": types.StringType, "state": types.StringType, "tenant": types.StringType, "tenant_name": types.StringType, "tenant_uuid": types.StringType, "url": types.StringType}}}.ElemType),
												"subnet_cidr": func() attr.Value {
													if v, ok := objMap["subnet_cidr"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"subnet_description": func() attr.Value {
													if v, ok := objMap["subnet_description"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"subnet_name": func() attr.Value {
													if v, ok := objMap["subnet_name"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"subnet_uuid": func() attr.Value {
													if v, ok := objMap["subnet_uuid"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"url": func() attr.Value {
													if v, ok := objMap["url"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
											}
											objVal, _ := types.ObjectValue(attrTypes, attrValues)
											items = append(items, objVal)
										}
									}
									listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
										"fixed_ips":             types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
										"port":                  types.StringType,
										"subnet":                types.StringType,
										"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"mac_address": types.StringType}}},
										"device_id":             types.StringType,
										"device_owner":          types.StringType,
										"mac_address":           types.StringType,
										"security_groups":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"access_url": types.StringType, "backend_id": types.StringType, "created": types.StringType, "customer": types.StringType, "customer_abbreviation": types.StringType, "customer_name": types.StringType, "customer_native_name": types.StringType, "customer_uuid": types.StringType, "description": types.StringType, "error_message": types.StringType, "error_traceback": types.StringType, "is_limit_based": types.BoolType, "is_usage_based": types.BoolType, "marketplace_category_name": types.StringType, "marketplace_category_uuid": types.StringType, "marketplace_offering_name": types.StringType, "marketplace_offering_uuid": types.StringType, "marketplace_plan_uuid": types.StringType, "marketplace_resource_state": types.StringType, "marketplace_resource_uuid": types.StringType, "modified": types.StringType, "name": types.StringType, "project": types.StringType, "project_name": types.StringType, "project_uuid": types.StringType, "resource_type": types.StringType, "rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}}, "service_name": types.StringType, "service_settings": types.StringType, "service_settings_error_message": types.StringType, "service_settings_state": types.StringType, "service_settings_uuid": types.StringType, "state": types.StringType, "tenant": types.StringType, "tenant_name": types.StringType, "tenant_uuid": types.StringType, "url": types.StringType}}},
										"subnet_cidr":           types.StringType,
										"subnet_description":    types.StringType,
										"subnet_name":           types.StringType,
										"subnet_uuid":           types.StringType,
										"url":                   types.StringType,
									}}, items)
									data.Ports = listVal
								}
							} else {
								if data.Ports.IsUnknown() {
									data.Ports = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
										"fixed_ips":             types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
										"port":                  types.StringType,
										"subnet":                types.StringType,
										"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"mac_address": types.StringType}}},
										"device_id":             types.StringType,
										"device_owner":          types.StringType,
										"mac_address":           types.StringType,
										"security_groups":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"access_url": types.StringType, "backend_id": types.StringType, "created": types.StringType, "customer": types.StringType, "customer_abbreviation": types.StringType, "customer_name": types.StringType, "customer_native_name": types.StringType, "customer_uuid": types.StringType, "description": types.StringType, "error_message": types.StringType, "error_traceback": types.StringType, "is_limit_based": types.BoolType, "is_usage_based": types.BoolType, "marketplace_category_name": types.StringType, "marketplace_category_uuid": types.StringType, "marketplace_offering_name": types.StringType, "marketplace_offering_uuid": types.StringType, "marketplace_plan_uuid": types.StringType, "marketplace_resource_state": types.StringType, "marketplace_resource_uuid": types.StringType, "modified": types.StringType, "name": types.StringType, "project": types.StringType, "project_name": types.StringType, "project_uuid": types.StringType, "resource_type": types.StringType, "rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}}, "service_name": types.StringType, "service_settings": types.StringType, "service_settings_error_message": types.StringType, "service_settings_state": types.StringType, "service_settings_uuid": types.StringType, "state": types.StringType, "tenant": types.StringType, "tenant_name": types.StringType, "tenant_uuid": types.StringType, "url": types.StringType}}},
										"subnet_cidr":           types.StringType,
										"subnet_description":    types.StringType,
										"subnet_name":           types.StringType,
										"subnet_uuid":           types.StringType,
										"url":                   types.StringType,
									}})
								}
							}
							if val, ok := sourceMap["project"]; ok && val != nil {
								if str, ok := val.(string); ok {
									// Normalize URL to UUID
									parts := strings.Split(strings.TrimRight(str, "/"), "/")
									uuid := parts[len(parts)-1]
									data.Project = types.StringValue(uuid)
								} else {
									data.Project = types.StringNull()
								}
							} else {
								if data.Project.IsUnknown() {
									data.Project = types.StringNull()
								}
							}
							if val, ok := sourceMap["project_name"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.ProjectName = types.StringValue(str)
								}
							} else {
								if data.ProjectName.IsUnknown() {
									data.ProjectName = types.StringNull()
								}
							}
							if val, ok := sourceMap["project_uuid"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.ProjectUuid = types.StringValue(str)
								}
							} else {
								if data.ProjectUuid.IsUnknown() {
									data.ProjectUuid = types.StringNull()
								}
							}
							if val, ok := sourceMap["ram"]; ok && val != nil {
								if num, ok := val.(float64); ok {
									data.Ram = types.Int64Value(int64(num))
								}
							} else {
								if data.Ram.IsUnknown() {
									data.Ram = types.Int64Null()
								}
							}
							if val, ok := sourceMap["release_floating_ips"]; ok && val != nil {
								if b, ok := val.(bool); ok {
									data.ReleaseFloatingIps = types.BoolValue(b)
								}
							} else {
								if data.ReleaseFloatingIps.IsUnknown() {
									data.ReleaseFloatingIps = types.BoolNull()
								}
							}
							if val, ok := sourceMap["resource_type"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.ResourceType = types.StringValue(str)
								}
							} else {
								if data.ResourceType.IsUnknown() {
									data.ResourceType = types.StringNull()
								}
							}
							if val, ok := sourceMap["runtime_state"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.RuntimeState = types.StringValue(str)
								}
							} else {
								if data.RuntimeState.IsUnknown() {
									data.RuntimeState = types.StringNull()
								}
							}
							if val, ok := sourceMap["security_groups"]; ok && val != nil {
								// List of objects
								if arr, ok := val.([]interface{}); ok {
									items := make([]attr.Value, 0, len(arr))
									for _, item := range arr {
										if objMap, ok := item.(map[string]interface{}); ok {
											attrTypes := map[string]attr.Type{
												"url":         types.StringType,
												"description": types.StringType,
												"name":        types.StringType,
												"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}},
												"state":       types.StringType,
											}
											attrValues := map[string]attr.Value{
												"url": func() attr.Value {
													if v, ok := objMap["url"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"description": func() attr.Value {
													if v, ok := objMap["description"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"name": func() attr.Value {
													if v, ok := objMap["name"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"rules": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}}.ElemType),
												"state": func() attr.Value {
													if v, ok := objMap["state"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
											}
											objVal, _ := types.ObjectValue(attrTypes, attrValues)
											items = append(items, objVal)
										}
									}
									listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
										"url":         types.StringType,
										"description": types.StringType,
										"name":        types.StringType,
										"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}},
										"state":       types.StringType,
									}}, items)
									data.SecurityGroups = listVal
								}
							} else {
								if data.SecurityGroups.IsUnknown() {
									data.SecurityGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
										"url":         types.StringType,
										"description": types.StringType,
										"name":        types.StringType,
										"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}},
										"state":       types.StringType,
									}})
								}
							}
							if val, ok := sourceMap["server_group"]; ok && val != nil {
								// Nested object
								if objMap, ok := val.(map[string]interface{}); ok {
									_ = objMap // Avoid unused variable if properties are empty
									attrTypes := map[string]attr.Type{
										"name":   types.StringType,
										"policy": types.StringType,
										"state":  types.StringType,
										"url":    types.StringType,
									}
									attrValues := map[string]attr.Value{
										"name": func() attr.Value {
											if v, ok := objMap["name"].(string); ok {
												return types.StringValue(v)
											}
											return types.StringNull()
										}(),
										"policy": func() attr.Value {
											if v, ok := objMap["policy"].(string); ok {
												return types.StringValue(v)
											}
											return types.StringNull()
										}(),
										"state": func() attr.Value {
											if v, ok := objMap["state"].(string); ok {
												return types.StringValue(v)
											}
											return types.StringNull()
										}(),
										"url": func() attr.Value {
											if v, ok := objMap["url"].(string); ok {
												return types.StringValue(v)
											}
											return types.StringNull()
										}(),
									}
									objVal, _ := types.ObjectValue(attrTypes, attrValues)
									data.ServerGroup = objVal
								}
							} else {
								if data.ServerGroup.IsUnknown() {
									data.ServerGroup = types.ObjectNull(map[string]attr.Type{
										"name":   types.StringType,
										"policy": types.StringType,
										"state":  types.StringType,
										"url":    types.StringType,
									})
								}
							}
							if val, ok := sourceMap["service_name"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.ServiceName = types.StringValue(str)
								}
							} else {
								if data.ServiceName.IsUnknown() {
									data.ServiceName = types.StringNull()
								}
							}
							if val, ok := sourceMap["service_settings"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.ServiceSettings = types.StringValue(str)
								}
							} else {
								if data.ServiceSettings.IsUnknown() {
									data.ServiceSettings = types.StringNull()
								}
							}
							if val, ok := sourceMap["service_settings_error_message"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.ServiceSettingsErrorMessage = types.StringValue(str)
								}
							} else {
								if data.ServiceSettingsErrorMessage.IsUnknown() {
									data.ServiceSettingsErrorMessage = types.StringNull()
								}
							}
							if val, ok := sourceMap["service_settings_state"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.ServiceSettingsState = types.StringValue(str)
								}
							} else {
								if data.ServiceSettingsState.IsUnknown() {
									data.ServiceSettingsState = types.StringNull()
								}
							}
							if val, ok := sourceMap["service_settings_uuid"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.ServiceSettingsUuid = types.StringValue(str)
								}
							} else {
								if data.ServiceSettingsUuid.IsUnknown() {
									data.ServiceSettingsUuid = types.StringNull()
								}
							}
							if val, ok := sourceMap["ssh_public_key"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.SshPublicKey = types.StringValue(str)
								}
							} else {
								if data.SshPublicKey.IsUnknown() {
									data.SshPublicKey = types.StringNull()
								}
							}
							if val, ok := sourceMap["start_time"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.StartTime = types.StringValue(str)
								}
							} else {
								if data.StartTime.IsUnknown() {
									data.StartTime = types.StringNull()
								}
							}
							if val, ok := sourceMap["state"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.State = types.StringValue(str)
								}
							} else {
								if data.State.IsUnknown() {
									data.State = types.StringNull()
								}
							}
							if val, ok := sourceMap["system_volume_size"]; ok && val != nil {
								if num, ok := val.(float64); ok {
									data.SystemVolumeSize = types.Int64Value(int64(num))
								}
							} else {
								if data.SystemVolumeSize.IsUnknown() {
									data.SystemVolumeSize = types.Int64Null()
								}
							}
							if val, ok := sourceMap["system_volume_type"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.SystemVolumeType = types.StringValue(str)
								}
							} else {
								if data.SystemVolumeType.IsUnknown() {
									data.SystemVolumeType = types.StringNull()
								}
							}
							if val, ok := sourceMap["tenant"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.Tenant = types.StringValue(str)
								}
							} else {
								if data.Tenant.IsUnknown() {
									data.Tenant = types.StringNull()
								}
							}
							if val, ok := sourceMap["tenant_uuid"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.TenantUuid = types.StringValue(str)
								}
							} else {
								if data.TenantUuid.IsUnknown() {
									data.TenantUuid = types.StringNull()
								}
							}
							if val, ok := sourceMap["url"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.Url = types.StringValue(str)
								}
							} else {
								if data.Url.IsUnknown() {
									data.Url = types.StringNull()
								}
							}
							if val, ok := sourceMap["user_data"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.UserData = types.StringValue(str)
								}
							} else {
								if data.UserData.IsUnknown() {
									data.UserData = types.StringNull()
								}
							}
							if val, ok := sourceMap["volumes"]; ok && val != nil {
								// List of objects
								if arr, ok := val.([]interface{}); ok {
									items := make([]attr.Value, 0, len(arr))
									for _, item := range arr {
										if objMap, ok := item.(map[string]interface{}); ok {
											attrTypes := map[string]attr.Type{
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
											}
											attrValues := map[string]attr.Value{
												"bootable": func() attr.Value {
													if v, ok := objMap["bootable"].(bool); ok {
														return types.BoolValue(v)
													}
													return types.BoolNull()
												}(),
												"device": func() attr.Value {
													if v, ok := objMap["device"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"image_name": func() attr.Value {
													if v, ok := objMap["image_name"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"marketplace_resource_uuid": func() attr.Value {
													if v, ok := objMap["marketplace_resource_uuid"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"name": func() attr.Value {
													if v, ok := objMap["name"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"resource_type": func() attr.Value {
													if v, ok := objMap["resource_type"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"size": func() attr.Value {
													if v, ok := objMap["size"].(float64); ok {
														return types.Int64Value(int64(v))
													}
													return types.Int64Null()
												}(),
												"state": func() attr.Value {
													if v, ok := objMap["state"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"type": func() attr.Value {
													if v, ok := objMap["type"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"type_name": func() attr.Value {
													if v, ok := objMap["type_name"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"url": func() attr.Value {
													if v, ok := objMap["url"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
											}
											objVal, _ := types.ObjectValue(attrTypes, attrValues)
											items = append(items, objVal)
										}
									}
									listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
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
									}}, items)
									data.Volumes = listVal
								}
							} else {
								if data.Volumes.IsUnknown() {
									data.Volumes = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
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
									}})
								}
							}

							// Map filter parameters from response if available
						} else {
							tflog.Warn(ctx, "Failed to fetch plugin resource: "+err.Error())
						}
					} else {
						tflog.Warn(ctx, "resource_uuid is empty string")
					}
				} else {
					tflog.Warn(ctx, "Failed to cast resource_uuid to string")
				}
			} else {
				tflog.Warn(ctx, "Failed to fetch MP resource: "+err.Error())
			}
		}
	}

	if os.Getenv("WALDUR_E2E_SKIP_WAIT") != "" {
		tflog.Warn(ctx, "Skipping wait for order completion due to WALDUR_E2E_SKIP_WAIT")
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}

	stateConf := &retry.StateChangeConf{
		Pending: []string{"pending", "executing", "created"},
		Target:  []string{"done"},
		Refresh: func() (interface{}, string, error) {
			var res map[string]interface{}
			err := r.client.GetByUUID(ctx, "/api/marketplace-orders/{uuid}/", orderUUID, &res)
			if err != nil {
				return nil, "", err
			}

			state, _ := res["state"].(string)
			if state == "erred" || state == "rejected" {
				msg, _ := res["error_message"].(string)
				return res, "failed", fmt.Errorf("order failed: %s", msg)
			}
			return res, state, nil
		},
		Timeout: func() time.Duration {
			timeout, diags := data.Timeouts.Create(ctx, 45*time.Minute)
			resp.Diagnostics.Append(diags...)
			return timeout
		}(),
		Delay:      10 * time.Second,
		MinTimeout: 5 * time.Second,
	}

	rawResult, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Order Failed", err.Error())
		return
	}

	result := rawResult.(map[string]interface{})
	if resourceUUID, ok := result["marketplace_resource_uuid"].(string); ok {
		data.UUID = types.StringValue(resourceUUID)
	} else {
		resp.Diagnostics.AddError("Resource UUID Missing", "Order completed but marketplace_resource_uuid is missing")
		return
	}

	// Fetch final resource state
	var finalState map[string]interface{}
	err = r.client.GetByUUID(ctx, "/api/openstack-instances/{uuid}/", data.UUID.ValueString(), &finalState)
	if err != nil {
		resp.Diagnostics.AddError("Failed to Read Resource", err.Error())
		return
	}
	sourceMap := finalState
	// Map response fields to data model
	_ = sourceMap
	if val, ok := sourceMap["access_url"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AccessUrl = types.StringValue(str)
		}
	} else {
		if data.AccessUrl.IsUnknown() {
			data.AccessUrl = types.StringNull()
		}
	}
	if val, ok := sourceMap["action"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Action = types.StringValue(str)
		}
	} else {
		if data.Action.IsUnknown() {
			data.Action = types.StringNull()
		}
	}
	if val, ok := sourceMap["availability_zone"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AvailabilityZone = types.StringValue(str)
		}
	} else {
		if data.AvailabilityZone.IsUnknown() {
			data.AvailabilityZone = types.StringNull()
		}
	}
	if val, ok := sourceMap["availability_zone_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AvailabilityZoneName = types.StringValue(str)
		}
	} else {
		if data.AvailabilityZoneName.IsUnknown() {
			data.AvailabilityZoneName = types.StringNull()
		}
	}
	if val, ok := sourceMap["backend_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.BackendId = types.StringValue(str)
		}
	} else {
		if data.BackendId.IsUnknown() {
			data.BackendId = types.StringNull()
		}
	}
	if val, ok := sourceMap["connect_directly_to_external_network"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.ConnectDirectlyToExternalNetwork = types.BoolValue(b)
		}
	} else {
		if data.ConnectDirectlyToExternalNetwork.IsUnknown() {
			data.ConnectDirectlyToExternalNetwork = types.BoolNull()
		}
	}
	if val, ok := sourceMap["cores"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Cores = types.Int64Value(int64(num))
		}
	} else {
		if data.Cores.IsUnknown() {
			data.Cores = types.Int64Null()
		}
	}
	if val, ok := sourceMap["created"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Created = types.StringValue(str)
		}
	} else {
		if data.Created.IsUnknown() {
			data.Created = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Customer = types.StringValue(str)
		}
	} else {
		if data.Customer.IsUnknown() {
			data.Customer = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer_abbreviation"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CustomerAbbreviation = types.StringValue(str)
		}
	} else {
		if data.CustomerAbbreviation.IsUnknown() {
			data.CustomerAbbreviation = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CustomerName = types.StringValue(str)
		}
	} else {
		if data.CustomerName.IsUnknown() {
			data.CustomerName = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer_native_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CustomerNativeName = types.StringValue(str)
		}
	} else {
		if data.CustomerNativeName.IsUnknown() {
			data.CustomerNativeName = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CustomerUuid = types.StringValue(str)
		}
	} else {
		if data.CustomerUuid.IsUnknown() {
			data.CustomerUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["data_volume_size"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.DataVolumeSize = types.Int64Value(int64(num))
		}
	} else {
		if data.DataVolumeSize.IsUnknown() {
			data.DataVolumeSize = types.Int64Null()
		}
	}
	if val, ok := sourceMap["data_volume_type"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.DataVolumeType = types.StringValue(str)
		}
	} else {
		if data.DataVolumeType.IsUnknown() {
			data.DataVolumeType = types.StringNull()
		}
	}
	if val, ok := sourceMap["data_volumes"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"size":        types.Int64Type,
						"volume_type": types.StringType,
					}
					attrValues := map[string]attr.Value{
						"size": func() attr.Value {
							if v, ok := objMap["size"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"volume_type": func() attr.Value {
							if v, ok := objMap["volume_type"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"size":        types.Int64Type,
				"volume_type": types.StringType,
			}}, items)
			data.DataVolumes = listVal
		}
	} else {
		if data.DataVolumes.IsUnknown() {
			data.DataVolumes = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"size":        types.Int64Type,
				"volume_type": types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["delete_volumes"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.DeleteVolumes = types.BoolValue(b)
		}
	} else {
		if data.DeleteVolumes.IsUnknown() {
			data.DeleteVolumes = types.BoolNull()
		}
	}
	if val, ok := sourceMap["description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
		}
	} else {
		if data.Description.IsUnknown() {
			data.Description = types.StringNull()
		}
	}
	if val, ok := sourceMap["disk"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Disk = types.Int64Value(int64(num))
		}
	} else {
		if data.Disk.IsUnknown() {
			data.Disk = types.Int64Null()
		}
	}
	if val, ok := sourceMap["error_message"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ErrorMessage = types.StringValue(str)
		}
	} else {
		if data.ErrorMessage.IsUnknown() {
			data.ErrorMessage = types.StringNull()
		}
	}
	if val, ok := sourceMap["error_traceback"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ErrorTraceback = types.StringValue(str)
		}
	} else {
		if data.ErrorTraceback.IsUnknown() {
			data.ErrorTraceback = types.StringNull()
		}
	}
	if val, ok := sourceMap["external_address"]; ok && val != nil {
		// List of strings (or flattened objects)
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if str, ok := item.(string); ok {
					items = append(items, types.StringValue(str))
				} else if obj, ok := item.(map[string]interface{}); ok {
					// Flattening logic: extract URL or UUID
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					} else if name, ok := obj["name"].(string); ok {
						items = append(items, types.StringValue(name))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.ExternalAddress = listVal
		}
	} else {
		if data.ExternalAddress.IsUnknown() {
			data.ExternalAddress = types.ListNull(types.StringType)
		}
	}
	if val, ok := sourceMap["external_ips"]; ok && val != nil {
		// List of strings (or flattened objects)
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if str, ok := item.(string); ok {
					items = append(items, types.StringValue(str))
				} else if obj, ok := item.(map[string]interface{}); ok {
					// Flattening logic: extract URL or UUID
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					} else if name, ok := obj["name"].(string); ok {
						items = append(items, types.StringValue(name))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.ExternalIps = listVal
		}
	} else {
		if data.ExternalIps.IsUnknown() {
			data.ExternalIps = types.ListNull(types.StringType)
		}
	}
	if val, ok := sourceMap["flavor"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Flavor = types.StringValue(str)
		}
	} else {
		if data.Flavor.IsUnknown() {
			data.Flavor = types.StringNull()
		}
	}
	if val, ok := sourceMap["flavor_disk"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.FlavorDisk = types.Int64Value(int64(num))
		}
	} else {
		if data.FlavorDisk.IsUnknown() {
			data.FlavorDisk = types.Int64Null()
		}
	}
	if val, ok := sourceMap["flavor_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.FlavorName = types.StringValue(str)
		}
	} else {
		if data.FlavorName.IsUnknown() {
			data.FlavorName = types.StringNull()
		}
	}
	if val, ok := sourceMap["floating_ips"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"ip_address":         types.StringType,
						"subnet":             types.StringType,
						"url":                types.StringType,
						"address":            types.StringType,
						"port_fixed_ips":     types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
						"port_mac_address":   types.StringType,
						"subnet_cidr":        types.StringType,
						"subnet_description": types.StringType,
						"subnet_name":        types.StringType,
						"subnet_uuid":        types.StringType,
					}
					attrValues := map[string]attr.Value{
						"ip_address": func() attr.Value {
							if v, ok := objMap["ip_address"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet": func() attr.Value {
							if v, ok := objMap["subnet"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"address": func() attr.Value {
							if v, ok := objMap["address"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"port_fixed_ips": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}}.ElemType),
						"port_mac_address": func() attr.Value {
							if v, ok := objMap["port_mac_address"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet_cidr": func() attr.Value {
							if v, ok := objMap["subnet_cidr"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet_description": func() attr.Value {
							if v, ok := objMap["subnet_description"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet_name": func() attr.Value {
							if v, ok := objMap["subnet_name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet_uuid": func() attr.Value {
							if v, ok := objMap["subnet_uuid"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"ip_address":         types.StringType,
				"subnet":             types.StringType,
				"url":                types.StringType,
				"address":            types.StringType,
				"port_fixed_ips":     types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
				"port_mac_address":   types.StringType,
				"subnet_cidr":        types.StringType,
				"subnet_description": types.StringType,
				"subnet_name":        types.StringType,
				"subnet_uuid":        types.StringType,
			}}, items)
			data.FloatingIps = listVal
		}
	} else {
		if data.FloatingIps.IsUnknown() {
			data.FloatingIps = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"ip_address":         types.StringType,
				"subnet":             types.StringType,
				"url":                types.StringType,
				"address":            types.StringType,
				"port_fixed_ips":     types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
				"port_mac_address":   types.StringType,
				"subnet_cidr":        types.StringType,
				"subnet_description": types.StringType,
				"subnet_name":        types.StringType,
				"subnet_uuid":        types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["hypervisor_hostname"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.HypervisorHostname = types.StringValue(str)
		}
	} else {
		if data.HypervisorHostname.IsUnknown() {
			data.HypervisorHostname = types.StringNull()
		}
	}
	if val, ok := sourceMap["image"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Image = types.StringValue(str)
		}
	} else {
		if data.Image.IsUnknown() {
			data.Image = types.StringNull()
		}
	}
	if val, ok := sourceMap["image_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ImageName = types.StringValue(str)
		}
	} else {
		if data.ImageName.IsUnknown() {
			data.ImageName = types.StringNull()
		}
	}
	if val, ok := sourceMap["internal_ips"]; ok && val != nil {
		// List of strings (or flattened objects)
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if str, ok := item.(string); ok {
					items = append(items, types.StringValue(str))
				} else if obj, ok := item.(map[string]interface{}); ok {
					// Flattening logic: extract URL or UUID
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					} else if name, ok := obj["name"].(string); ok {
						items = append(items, types.StringValue(name))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.InternalIps = listVal
		}
	} else {
		if data.InternalIps.IsUnknown() {
			data.InternalIps = types.ListNull(types.StringType)
		}
	}
	if val, ok := sourceMap["is_limit_based"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.IsLimitBased = types.BoolValue(b)
		}
	} else {
		if data.IsLimitBased.IsUnknown() {
			data.IsLimitBased = types.BoolNull()
		}
	}
	if val, ok := sourceMap["is_usage_based"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.IsUsageBased = types.BoolValue(b)
		}
	} else {
		if data.IsUsageBased.IsUnknown() {
			data.IsUsageBased = types.BoolNull()
		}
	}
	if val, ok := sourceMap["key_fingerprint"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.KeyFingerprint = types.StringValue(str)
		}
	} else {
		if data.KeyFingerprint.IsUnknown() {
			data.KeyFingerprint = types.StringNull()
		}
	}
	if val, ok := sourceMap["key_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.KeyName = types.StringValue(str)
		}
	} else {
		if data.KeyName.IsUnknown() {
			data.KeyName = types.StringNull()
		}
	}
	if val, ok := sourceMap["latitude"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Latitude = types.Float64Value(num)
		}
	} else {
		if data.Latitude.IsUnknown() {
			data.Latitude = types.Float64Null()
		}
	}
	if val, ok := sourceMap["longitude"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Longitude = types.Float64Value(num)
		}
	} else {
		if data.Longitude.IsUnknown() {
			data.Longitude = types.Float64Null()
		}
	}
	if val, ok := sourceMap["marketplace_category_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.MarketplaceCategoryName = types.StringValue(str)
		}
	} else {
		if data.MarketplaceCategoryName.IsUnknown() {
			data.MarketplaceCategoryName = types.StringNull()
		}
	}
	if val, ok := sourceMap["marketplace_category_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.MarketplaceCategoryUuid = types.StringValue(str)
		}
	} else {
		if data.MarketplaceCategoryUuid.IsUnknown() {
			data.MarketplaceCategoryUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["marketplace_offering_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.MarketplaceOfferingName = types.StringValue(str)
		}
	} else {
		if data.MarketplaceOfferingName.IsUnknown() {
			data.MarketplaceOfferingName = types.StringNull()
		}
	}
	if val, ok := sourceMap["marketplace_offering_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.MarketplaceOfferingUuid = types.StringValue(str)
		}
	} else {
		if data.MarketplaceOfferingUuid.IsUnknown() {
			data.MarketplaceOfferingUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["marketplace_plan_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.MarketplacePlanUuid = types.StringValue(str)
		}
	} else {
		if data.MarketplacePlanUuid.IsUnknown() {
			data.MarketplacePlanUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["marketplace_resource_state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.MarketplaceResourceState = types.StringValue(str)
		}
	} else {
		if data.MarketplaceResourceState.IsUnknown() {
			data.MarketplaceResourceState = types.StringNull()
		}
	}
	if val, ok := sourceMap["marketplace_resource_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.MarketplaceResourceUuid = types.StringValue(str)
		}
	} else {
		if data.MarketplaceResourceUuid.IsUnknown() {
			data.MarketplaceResourceUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["min_disk"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.MinDisk = types.Int64Value(int64(num))
		}
	} else {
		if data.MinDisk.IsUnknown() {
			data.MinDisk = types.Int64Null()
		}
	}
	if val, ok := sourceMap["min_ram"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.MinRam = types.Int64Value(int64(num))
		}
	} else {
		if data.MinRam.IsUnknown() {
			data.MinRam = types.Int64Null()
		}
	}
	if val, ok := sourceMap["modified"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Modified = types.StringValue(str)
		}
	} else {
		if data.Modified.IsUnknown() {
			data.Modified = types.StringNull()
		}
	}
	if val, ok := sourceMap["offering"]; ok && val != nil {
		if str, ok := val.(string); ok {
			// Normalize URL to UUID
			parts := strings.Split(strings.TrimRight(str, "/"), "/")
			uuid := parts[len(parts)-1]
			data.Offering = types.StringValue(uuid)
		} else {
			data.Offering = types.StringNull()
		}
	} else {
		if data.Offering.IsUnknown() {
			data.Offering = types.StringNull()
		}
	}
	if val, ok := sourceMap["ports"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"fixed_ips":             types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
						"port":                  types.StringType,
						"subnet":                types.StringType,
						"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"mac_address": types.StringType}}},
						"device_id":             types.StringType,
						"device_owner":          types.StringType,
						"mac_address":           types.StringType,
						"security_groups":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"access_url": types.StringType, "backend_id": types.StringType, "created": types.StringType, "customer": types.StringType, "customer_abbreviation": types.StringType, "customer_name": types.StringType, "customer_native_name": types.StringType, "customer_uuid": types.StringType, "description": types.StringType, "error_message": types.StringType, "error_traceback": types.StringType, "is_limit_based": types.BoolType, "is_usage_based": types.BoolType, "marketplace_category_name": types.StringType, "marketplace_category_uuid": types.StringType, "marketplace_offering_name": types.StringType, "marketplace_offering_uuid": types.StringType, "marketplace_plan_uuid": types.StringType, "marketplace_resource_state": types.StringType, "marketplace_resource_uuid": types.StringType, "modified": types.StringType, "name": types.StringType, "project": types.StringType, "project_name": types.StringType, "project_uuid": types.StringType, "resource_type": types.StringType, "rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}}, "service_name": types.StringType, "service_settings": types.StringType, "service_settings_error_message": types.StringType, "service_settings_state": types.StringType, "service_settings_uuid": types.StringType, "state": types.StringType, "tenant": types.StringType, "tenant_name": types.StringType, "tenant_uuid": types.StringType, "url": types.StringType}}},
						"subnet_cidr":           types.StringType,
						"subnet_description":    types.StringType,
						"subnet_name":           types.StringType,
						"subnet_uuid":           types.StringType,
						"url":                   types.StringType,
					}
					attrValues := map[string]attr.Value{
						"fixed_ips": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}}.ElemType),
						"port": func() attr.Value {
							if v, ok := objMap["port"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet": func() attr.Value {
							if v, ok := objMap["subnet"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"allowed_address_pairs": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"mac_address": types.StringType}}}.ElemType),
						"device_id": func() attr.Value {
							if v, ok := objMap["device_id"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"device_owner": func() attr.Value {
							if v, ok := objMap["device_owner"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"mac_address": func() attr.Value {
							if v, ok := objMap["mac_address"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"security_groups": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"access_url": types.StringType, "backend_id": types.StringType, "created": types.StringType, "customer": types.StringType, "customer_abbreviation": types.StringType, "customer_name": types.StringType, "customer_native_name": types.StringType, "customer_uuid": types.StringType, "description": types.StringType, "error_message": types.StringType, "error_traceback": types.StringType, "is_limit_based": types.BoolType, "is_usage_based": types.BoolType, "marketplace_category_name": types.StringType, "marketplace_category_uuid": types.StringType, "marketplace_offering_name": types.StringType, "marketplace_offering_uuid": types.StringType, "marketplace_plan_uuid": types.StringType, "marketplace_resource_state": types.StringType, "marketplace_resource_uuid": types.StringType, "modified": types.StringType, "name": types.StringType, "project": types.StringType, "project_name": types.StringType, "project_uuid": types.StringType, "resource_type": types.StringType, "rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}}, "service_name": types.StringType, "service_settings": types.StringType, "service_settings_error_message": types.StringType, "service_settings_state": types.StringType, "service_settings_uuid": types.StringType, "state": types.StringType, "tenant": types.StringType, "tenant_name": types.StringType, "tenant_uuid": types.StringType, "url": types.StringType}}}.ElemType),
						"subnet_cidr": func() attr.Value {
							if v, ok := objMap["subnet_cidr"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet_description": func() attr.Value {
							if v, ok := objMap["subnet_description"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet_name": func() attr.Value {
							if v, ok := objMap["subnet_name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet_uuid": func() attr.Value {
							if v, ok := objMap["subnet_uuid"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"fixed_ips":             types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
				"port":                  types.StringType,
				"subnet":                types.StringType,
				"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"mac_address": types.StringType}}},
				"device_id":             types.StringType,
				"device_owner":          types.StringType,
				"mac_address":           types.StringType,
				"security_groups":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"access_url": types.StringType, "backend_id": types.StringType, "created": types.StringType, "customer": types.StringType, "customer_abbreviation": types.StringType, "customer_name": types.StringType, "customer_native_name": types.StringType, "customer_uuid": types.StringType, "description": types.StringType, "error_message": types.StringType, "error_traceback": types.StringType, "is_limit_based": types.BoolType, "is_usage_based": types.BoolType, "marketplace_category_name": types.StringType, "marketplace_category_uuid": types.StringType, "marketplace_offering_name": types.StringType, "marketplace_offering_uuid": types.StringType, "marketplace_plan_uuid": types.StringType, "marketplace_resource_state": types.StringType, "marketplace_resource_uuid": types.StringType, "modified": types.StringType, "name": types.StringType, "project": types.StringType, "project_name": types.StringType, "project_uuid": types.StringType, "resource_type": types.StringType, "rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}}, "service_name": types.StringType, "service_settings": types.StringType, "service_settings_error_message": types.StringType, "service_settings_state": types.StringType, "service_settings_uuid": types.StringType, "state": types.StringType, "tenant": types.StringType, "tenant_name": types.StringType, "tenant_uuid": types.StringType, "url": types.StringType}}},
				"subnet_cidr":           types.StringType,
				"subnet_description":    types.StringType,
				"subnet_name":           types.StringType,
				"subnet_uuid":           types.StringType,
				"url":                   types.StringType,
			}}, items)
			data.Ports = listVal
		}
	} else {
		if data.Ports.IsUnknown() {
			data.Ports = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"fixed_ips":             types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
				"port":                  types.StringType,
				"subnet":                types.StringType,
				"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"mac_address": types.StringType}}},
				"device_id":             types.StringType,
				"device_owner":          types.StringType,
				"mac_address":           types.StringType,
				"security_groups":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"access_url": types.StringType, "backend_id": types.StringType, "created": types.StringType, "customer": types.StringType, "customer_abbreviation": types.StringType, "customer_name": types.StringType, "customer_native_name": types.StringType, "customer_uuid": types.StringType, "description": types.StringType, "error_message": types.StringType, "error_traceback": types.StringType, "is_limit_based": types.BoolType, "is_usage_based": types.BoolType, "marketplace_category_name": types.StringType, "marketplace_category_uuid": types.StringType, "marketplace_offering_name": types.StringType, "marketplace_offering_uuid": types.StringType, "marketplace_plan_uuid": types.StringType, "marketplace_resource_state": types.StringType, "marketplace_resource_uuid": types.StringType, "modified": types.StringType, "name": types.StringType, "project": types.StringType, "project_name": types.StringType, "project_uuid": types.StringType, "resource_type": types.StringType, "rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}}, "service_name": types.StringType, "service_settings": types.StringType, "service_settings_error_message": types.StringType, "service_settings_state": types.StringType, "service_settings_uuid": types.StringType, "state": types.StringType, "tenant": types.StringType, "tenant_name": types.StringType, "tenant_uuid": types.StringType, "url": types.StringType}}},
				"subnet_cidr":           types.StringType,
				"subnet_description":    types.StringType,
				"subnet_name":           types.StringType,
				"subnet_uuid":           types.StringType,
				"url":                   types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["project"]; ok && val != nil {
		if str, ok := val.(string); ok {
			// Normalize URL to UUID
			parts := strings.Split(strings.TrimRight(str, "/"), "/")
			uuid := parts[len(parts)-1]
			data.Project = types.StringValue(uuid)
		} else {
			data.Project = types.StringNull()
		}
	} else {
		if data.Project.IsUnknown() {
			data.Project = types.StringNull()
		}
	}
	if val, ok := sourceMap["project_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ProjectName = types.StringValue(str)
		}
	} else {
		if data.ProjectName.IsUnknown() {
			data.ProjectName = types.StringNull()
		}
	}
	if val, ok := sourceMap["project_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ProjectUuid = types.StringValue(str)
		}
	} else {
		if data.ProjectUuid.IsUnknown() {
			data.ProjectUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["ram"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Ram = types.Int64Value(int64(num))
		}
	} else {
		if data.Ram.IsUnknown() {
			data.Ram = types.Int64Null()
		}
	}
	if val, ok := sourceMap["release_floating_ips"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.ReleaseFloatingIps = types.BoolValue(b)
		}
	} else {
		if data.ReleaseFloatingIps.IsUnknown() {
			data.ReleaseFloatingIps = types.BoolNull()
		}
	}
	if val, ok := sourceMap["resource_type"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ResourceType = types.StringValue(str)
		}
	} else {
		if data.ResourceType.IsUnknown() {
			data.ResourceType = types.StringNull()
		}
	}
	if val, ok := sourceMap["runtime_state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.RuntimeState = types.StringValue(str)
		}
	} else {
		if data.RuntimeState.IsUnknown() {
			data.RuntimeState = types.StringNull()
		}
	}
	if val, ok := sourceMap["security_groups"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"url":         types.StringType,
						"description": types.StringType,
						"name":        types.StringType,
						"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}},
						"state":       types.StringType,
					}
					attrValues := map[string]attr.Value{
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"description": func() attr.Value {
							if v, ok := objMap["description"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"rules": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}}.ElemType),
						"state": func() attr.Value {
							if v, ok := objMap["state"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"url":         types.StringType,
				"description": types.StringType,
				"name":        types.StringType,
				"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}},
				"state":       types.StringType,
			}}, items)
			data.SecurityGroups = listVal
		}
	} else {
		if data.SecurityGroups.IsUnknown() {
			data.SecurityGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"url":         types.StringType,
				"description": types.StringType,
				"name":        types.StringType,
				"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}},
				"state":       types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["server_group"]; ok && val != nil {
		// Nested object
		if objMap, ok := val.(map[string]interface{}); ok {
			_ = objMap // Avoid unused variable if properties are empty
			attrTypes := map[string]attr.Type{
				"name":   types.StringType,
				"policy": types.StringType,
				"state":  types.StringType,
				"url":    types.StringType,
			}
			attrValues := map[string]attr.Value{
				"name": func() attr.Value {
					if v, ok := objMap["name"].(string); ok {
						return types.StringValue(v)
					}
					return types.StringNull()
				}(),
				"policy": func() attr.Value {
					if v, ok := objMap["policy"].(string); ok {
						return types.StringValue(v)
					}
					return types.StringNull()
				}(),
				"state": func() attr.Value {
					if v, ok := objMap["state"].(string); ok {
						return types.StringValue(v)
					}
					return types.StringNull()
				}(),
				"url": func() attr.Value {
					if v, ok := objMap["url"].(string); ok {
						return types.StringValue(v)
					}
					return types.StringNull()
				}(),
			}
			objVal, _ := types.ObjectValue(attrTypes, attrValues)
			data.ServerGroup = objVal
		}
	} else {
		if data.ServerGroup.IsUnknown() {
			data.ServerGroup = types.ObjectNull(map[string]attr.Type{
				"name":   types.StringType,
				"policy": types.StringType,
				"state":  types.StringType,
				"url":    types.StringType,
			})
		}
	}
	if val, ok := sourceMap["service_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ServiceName = types.StringValue(str)
		}
	} else {
		if data.ServiceName.IsUnknown() {
			data.ServiceName = types.StringNull()
		}
	}
	if val, ok := sourceMap["service_settings"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ServiceSettings = types.StringValue(str)
		}
	} else {
		if data.ServiceSettings.IsUnknown() {
			data.ServiceSettings = types.StringNull()
		}
	}
	if val, ok := sourceMap["service_settings_error_message"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ServiceSettingsErrorMessage = types.StringValue(str)
		}
	} else {
		if data.ServiceSettingsErrorMessage.IsUnknown() {
			data.ServiceSettingsErrorMessage = types.StringNull()
		}
	}
	if val, ok := sourceMap["service_settings_state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ServiceSettingsState = types.StringValue(str)
		}
	} else {
		if data.ServiceSettingsState.IsUnknown() {
			data.ServiceSettingsState = types.StringNull()
		}
	}
	if val, ok := sourceMap["service_settings_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ServiceSettingsUuid = types.StringValue(str)
		}
	} else {
		if data.ServiceSettingsUuid.IsUnknown() {
			data.ServiceSettingsUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["ssh_public_key"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.SshPublicKey = types.StringValue(str)
		}
	} else {
		if data.SshPublicKey.IsUnknown() {
			data.SshPublicKey = types.StringNull()
		}
	}
	if val, ok := sourceMap["start_time"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.StartTime = types.StringValue(str)
		}
	} else {
		if data.StartTime.IsUnknown() {
			data.StartTime = types.StringNull()
		}
	}
	if val, ok := sourceMap["state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.State = types.StringValue(str)
		}
	} else {
		if data.State.IsUnknown() {
			data.State = types.StringNull()
		}
	}
	if val, ok := sourceMap["system_volume_size"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.SystemVolumeSize = types.Int64Value(int64(num))
		}
	} else {
		if data.SystemVolumeSize.IsUnknown() {
			data.SystemVolumeSize = types.Int64Null()
		}
	}
	if val, ok := sourceMap["system_volume_type"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.SystemVolumeType = types.StringValue(str)
		}
	} else {
		if data.SystemVolumeType.IsUnknown() {
			data.SystemVolumeType = types.StringNull()
		}
	}
	if val, ok := sourceMap["tenant"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Tenant = types.StringValue(str)
		}
	} else {
		if data.Tenant.IsUnknown() {
			data.Tenant = types.StringNull()
		}
	}
	if val, ok := sourceMap["tenant_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.TenantUuid = types.StringValue(str)
		}
	} else {
		if data.TenantUuid.IsUnknown() {
			data.TenantUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["url"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Url = types.StringValue(str)
		}
	} else {
		if data.Url.IsUnknown() {
			data.Url = types.StringNull()
		}
	}
	if val, ok := sourceMap["user_data"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.UserData = types.StringValue(str)
		}
	} else {
		if data.UserData.IsUnknown() {
			data.UserData = types.StringNull()
		}
	}
	if val, ok := sourceMap["volumes"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
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
					}
					attrValues := map[string]attr.Value{
						"bootable": func() attr.Value {
							if v, ok := objMap["bootable"].(bool); ok {
								return types.BoolValue(v)
							}
							return types.BoolNull()
						}(),
						"device": func() attr.Value {
							if v, ok := objMap["device"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"image_name": func() attr.Value {
							if v, ok := objMap["image_name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"marketplace_resource_uuid": func() attr.Value {
							if v, ok := objMap["marketplace_resource_uuid"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"resource_type": func() attr.Value {
							if v, ok := objMap["resource_type"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"size": func() attr.Value {
							if v, ok := objMap["size"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"state": func() attr.Value {
							if v, ok := objMap["state"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"type": func() attr.Value {
							if v, ok := objMap["type"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"type_name": func() attr.Value {
							if v, ok := objMap["type_name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}}, items)
			data.Volumes = listVal
		}
	} else {
		if data.Volumes.IsUnknown() {
			data.Volumes = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}})
		}
	}

	// Map filter parameters from response if available

	// Save data into Terraform state
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
	var result map[string]interface{}

	retrievePath := strings.Replace("/api/openstack-instances/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read OpenstackInstance",
			"An error occurred while reading the openstack_instance: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	sourceMap := result
	// Map response fields to data model
	_ = sourceMap
	if val, ok := sourceMap["access_url"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AccessUrl = types.StringValue(str)
		}
	} else {
		if data.AccessUrl.IsUnknown() {
			data.AccessUrl = types.StringNull()
		}
	}
	if val, ok := sourceMap["action"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Action = types.StringValue(str)
		}
	} else {
		if data.Action.IsUnknown() {
			data.Action = types.StringNull()
		}
	}
	if val, ok := sourceMap["availability_zone"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AvailabilityZone = types.StringValue(str)
		}
	} else {
		if data.AvailabilityZone.IsUnknown() {
			data.AvailabilityZone = types.StringNull()
		}
	}
	if val, ok := sourceMap["availability_zone_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AvailabilityZoneName = types.StringValue(str)
		}
	} else {
		if data.AvailabilityZoneName.IsUnknown() {
			data.AvailabilityZoneName = types.StringNull()
		}
	}
	if val, ok := sourceMap["backend_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.BackendId = types.StringValue(str)
		}
	} else {
		if data.BackendId.IsUnknown() {
			data.BackendId = types.StringNull()
		}
	}
	if val, ok := sourceMap["connect_directly_to_external_network"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.ConnectDirectlyToExternalNetwork = types.BoolValue(b)
		}
	} else {
		if data.ConnectDirectlyToExternalNetwork.IsUnknown() {
			data.ConnectDirectlyToExternalNetwork = types.BoolNull()
		}
	}
	if val, ok := sourceMap["cores"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Cores = types.Int64Value(int64(num))
		}
	} else {
		if data.Cores.IsUnknown() {
			data.Cores = types.Int64Null()
		}
	}
	if val, ok := sourceMap["created"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Created = types.StringValue(str)
		}
	} else {
		if data.Created.IsUnknown() {
			data.Created = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Customer = types.StringValue(str)
		}
	} else {
		if data.Customer.IsUnknown() {
			data.Customer = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer_abbreviation"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CustomerAbbreviation = types.StringValue(str)
		}
	} else {
		if data.CustomerAbbreviation.IsUnknown() {
			data.CustomerAbbreviation = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CustomerName = types.StringValue(str)
		}
	} else {
		if data.CustomerName.IsUnknown() {
			data.CustomerName = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer_native_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CustomerNativeName = types.StringValue(str)
		}
	} else {
		if data.CustomerNativeName.IsUnknown() {
			data.CustomerNativeName = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CustomerUuid = types.StringValue(str)
		}
	} else {
		if data.CustomerUuid.IsUnknown() {
			data.CustomerUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["data_volume_size"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.DataVolumeSize = types.Int64Value(int64(num))
		}
	} else {
		if data.DataVolumeSize.IsUnknown() {
			data.DataVolumeSize = types.Int64Null()
		}
	}
	if val, ok := sourceMap["data_volume_type"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.DataVolumeType = types.StringValue(str)
		}
	} else {
		if data.DataVolumeType.IsUnknown() {
			data.DataVolumeType = types.StringNull()
		}
	}
	if val, ok := sourceMap["data_volumes"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"size":        types.Int64Type,
						"volume_type": types.StringType,
					}
					attrValues := map[string]attr.Value{
						"size": func() attr.Value {
							if v, ok := objMap["size"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"volume_type": func() attr.Value {
							if v, ok := objMap["volume_type"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"size":        types.Int64Type,
				"volume_type": types.StringType,
			}}, items)
			data.DataVolumes = listVal
		}
	} else {
		if data.DataVolumes.IsUnknown() {
			data.DataVolumes = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"size":        types.Int64Type,
				"volume_type": types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["delete_volumes"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.DeleteVolumes = types.BoolValue(b)
		}
	} else {
		if data.DeleteVolumes.IsUnknown() {
			data.DeleteVolumes = types.BoolNull()
		}
	}
	if val, ok := sourceMap["description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
		}
	} else {
		if data.Description.IsUnknown() {
			data.Description = types.StringNull()
		}
	}
	if val, ok := sourceMap["disk"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Disk = types.Int64Value(int64(num))
		}
	} else {
		if data.Disk.IsUnknown() {
			data.Disk = types.Int64Null()
		}
	}
	if val, ok := sourceMap["error_message"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ErrorMessage = types.StringValue(str)
		}
	} else {
		if data.ErrorMessage.IsUnknown() {
			data.ErrorMessage = types.StringNull()
		}
	}
	if val, ok := sourceMap["error_traceback"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ErrorTraceback = types.StringValue(str)
		}
	} else {
		if data.ErrorTraceback.IsUnknown() {
			data.ErrorTraceback = types.StringNull()
		}
	}
	if val, ok := sourceMap["external_address"]; ok && val != nil {
		// List of strings (or flattened objects)
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if str, ok := item.(string); ok {
					items = append(items, types.StringValue(str))
				} else if obj, ok := item.(map[string]interface{}); ok {
					// Flattening logic: extract URL or UUID
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					} else if name, ok := obj["name"].(string); ok {
						items = append(items, types.StringValue(name))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.ExternalAddress = listVal
		}
	} else {
		if data.ExternalAddress.IsUnknown() {
			data.ExternalAddress = types.ListNull(types.StringType)
		}
	}
	if val, ok := sourceMap["external_ips"]; ok && val != nil {
		// List of strings (or flattened objects)
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if str, ok := item.(string); ok {
					items = append(items, types.StringValue(str))
				} else if obj, ok := item.(map[string]interface{}); ok {
					// Flattening logic: extract URL or UUID
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					} else if name, ok := obj["name"].(string); ok {
						items = append(items, types.StringValue(name))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.ExternalIps = listVal
		}
	} else {
		if data.ExternalIps.IsUnknown() {
			data.ExternalIps = types.ListNull(types.StringType)
		}
	}
	if val, ok := sourceMap["flavor"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Flavor = types.StringValue(str)
		}
	} else {
		if data.Flavor.IsUnknown() {
			data.Flavor = types.StringNull()
		}
	}
	if val, ok := sourceMap["flavor_disk"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.FlavorDisk = types.Int64Value(int64(num))
		}
	} else {
		if data.FlavorDisk.IsUnknown() {
			data.FlavorDisk = types.Int64Null()
		}
	}
	if val, ok := sourceMap["flavor_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.FlavorName = types.StringValue(str)
		}
	} else {
		if data.FlavorName.IsUnknown() {
			data.FlavorName = types.StringNull()
		}
	}
	if val, ok := sourceMap["floating_ips"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"ip_address":         types.StringType,
						"subnet":             types.StringType,
						"url":                types.StringType,
						"address":            types.StringType,
						"port_fixed_ips":     types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
						"port_mac_address":   types.StringType,
						"subnet_cidr":        types.StringType,
						"subnet_description": types.StringType,
						"subnet_name":        types.StringType,
						"subnet_uuid":        types.StringType,
					}
					attrValues := map[string]attr.Value{
						"ip_address": func() attr.Value {
							if v, ok := objMap["ip_address"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet": func() attr.Value {
							if v, ok := objMap["subnet"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"address": func() attr.Value {
							if v, ok := objMap["address"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"port_fixed_ips": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}}.ElemType),
						"port_mac_address": func() attr.Value {
							if v, ok := objMap["port_mac_address"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet_cidr": func() attr.Value {
							if v, ok := objMap["subnet_cidr"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet_description": func() attr.Value {
							if v, ok := objMap["subnet_description"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet_name": func() attr.Value {
							if v, ok := objMap["subnet_name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet_uuid": func() attr.Value {
							if v, ok := objMap["subnet_uuid"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"ip_address":         types.StringType,
				"subnet":             types.StringType,
				"url":                types.StringType,
				"address":            types.StringType,
				"port_fixed_ips":     types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
				"port_mac_address":   types.StringType,
				"subnet_cidr":        types.StringType,
				"subnet_description": types.StringType,
				"subnet_name":        types.StringType,
				"subnet_uuid":        types.StringType,
			}}, items)
			data.FloatingIps = listVal
		}
	} else {
		if data.FloatingIps.IsUnknown() {
			data.FloatingIps = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"ip_address":         types.StringType,
				"subnet":             types.StringType,
				"url":                types.StringType,
				"address":            types.StringType,
				"port_fixed_ips":     types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
				"port_mac_address":   types.StringType,
				"subnet_cidr":        types.StringType,
				"subnet_description": types.StringType,
				"subnet_name":        types.StringType,
				"subnet_uuid":        types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["hypervisor_hostname"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.HypervisorHostname = types.StringValue(str)
		}
	} else {
		if data.HypervisorHostname.IsUnknown() {
			data.HypervisorHostname = types.StringNull()
		}
	}
	if val, ok := sourceMap["image"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Image = types.StringValue(str)
		}
	} else {
		if data.Image.IsUnknown() {
			data.Image = types.StringNull()
		}
	}
	if val, ok := sourceMap["image_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ImageName = types.StringValue(str)
		}
	} else {
		if data.ImageName.IsUnknown() {
			data.ImageName = types.StringNull()
		}
	}
	if val, ok := sourceMap["internal_ips"]; ok && val != nil {
		// List of strings (or flattened objects)
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if str, ok := item.(string); ok {
					items = append(items, types.StringValue(str))
				} else if obj, ok := item.(map[string]interface{}); ok {
					// Flattening logic: extract URL or UUID
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					} else if name, ok := obj["name"].(string); ok {
						items = append(items, types.StringValue(name))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.InternalIps = listVal
		}
	} else {
		if data.InternalIps.IsUnknown() {
			data.InternalIps = types.ListNull(types.StringType)
		}
	}
	if val, ok := sourceMap["is_limit_based"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.IsLimitBased = types.BoolValue(b)
		}
	} else {
		if data.IsLimitBased.IsUnknown() {
			data.IsLimitBased = types.BoolNull()
		}
	}
	if val, ok := sourceMap["is_usage_based"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.IsUsageBased = types.BoolValue(b)
		}
	} else {
		if data.IsUsageBased.IsUnknown() {
			data.IsUsageBased = types.BoolNull()
		}
	}
	if val, ok := sourceMap["key_fingerprint"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.KeyFingerprint = types.StringValue(str)
		}
	} else {
		if data.KeyFingerprint.IsUnknown() {
			data.KeyFingerprint = types.StringNull()
		}
	}
	if val, ok := sourceMap["key_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.KeyName = types.StringValue(str)
		}
	} else {
		if data.KeyName.IsUnknown() {
			data.KeyName = types.StringNull()
		}
	}
	if val, ok := sourceMap["latitude"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Latitude = types.Float64Value(num)
		}
	} else {
		if data.Latitude.IsUnknown() {
			data.Latitude = types.Float64Null()
		}
	}
	if val, ok := sourceMap["longitude"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Longitude = types.Float64Value(num)
		}
	} else {
		if data.Longitude.IsUnknown() {
			data.Longitude = types.Float64Null()
		}
	}
	if val, ok := sourceMap["marketplace_category_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.MarketplaceCategoryName = types.StringValue(str)
		}
	} else {
		if data.MarketplaceCategoryName.IsUnknown() {
			data.MarketplaceCategoryName = types.StringNull()
		}
	}
	if val, ok := sourceMap["marketplace_category_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.MarketplaceCategoryUuid = types.StringValue(str)
		}
	} else {
		if data.MarketplaceCategoryUuid.IsUnknown() {
			data.MarketplaceCategoryUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["marketplace_offering_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.MarketplaceOfferingName = types.StringValue(str)
		}
	} else {
		if data.MarketplaceOfferingName.IsUnknown() {
			data.MarketplaceOfferingName = types.StringNull()
		}
	}
	if val, ok := sourceMap["marketplace_offering_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.MarketplaceOfferingUuid = types.StringValue(str)
		}
	} else {
		if data.MarketplaceOfferingUuid.IsUnknown() {
			data.MarketplaceOfferingUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["marketplace_plan_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.MarketplacePlanUuid = types.StringValue(str)
		}
	} else {
		if data.MarketplacePlanUuid.IsUnknown() {
			data.MarketplacePlanUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["marketplace_resource_state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.MarketplaceResourceState = types.StringValue(str)
		}
	} else {
		if data.MarketplaceResourceState.IsUnknown() {
			data.MarketplaceResourceState = types.StringNull()
		}
	}
	if val, ok := sourceMap["marketplace_resource_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.MarketplaceResourceUuid = types.StringValue(str)
		}
	} else {
		if data.MarketplaceResourceUuid.IsUnknown() {
			data.MarketplaceResourceUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["min_disk"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.MinDisk = types.Int64Value(int64(num))
		}
	} else {
		if data.MinDisk.IsUnknown() {
			data.MinDisk = types.Int64Null()
		}
	}
	if val, ok := sourceMap["min_ram"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.MinRam = types.Int64Value(int64(num))
		}
	} else {
		if data.MinRam.IsUnknown() {
			data.MinRam = types.Int64Null()
		}
	}
	if val, ok := sourceMap["modified"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Modified = types.StringValue(str)
		}
	} else {
		if data.Modified.IsUnknown() {
			data.Modified = types.StringNull()
		}
	}
	if val, ok := sourceMap["offering"]; ok && val != nil {
		if str, ok := val.(string); ok {
			// Normalize URL to UUID
			parts := strings.Split(strings.TrimRight(str, "/"), "/")
			uuid := parts[len(parts)-1]
			data.Offering = types.StringValue(uuid)
		} else {
			data.Offering = types.StringNull()
		}
	} else {
		if data.Offering.IsUnknown() {
			data.Offering = types.StringNull()
		}
	}
	if val, ok := sourceMap["ports"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"fixed_ips":             types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
						"port":                  types.StringType,
						"subnet":                types.StringType,
						"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"mac_address": types.StringType}}},
						"device_id":             types.StringType,
						"device_owner":          types.StringType,
						"mac_address":           types.StringType,
						"security_groups":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"access_url": types.StringType, "backend_id": types.StringType, "created": types.StringType, "customer": types.StringType, "customer_abbreviation": types.StringType, "customer_name": types.StringType, "customer_native_name": types.StringType, "customer_uuid": types.StringType, "description": types.StringType, "error_message": types.StringType, "error_traceback": types.StringType, "is_limit_based": types.BoolType, "is_usage_based": types.BoolType, "marketplace_category_name": types.StringType, "marketplace_category_uuid": types.StringType, "marketplace_offering_name": types.StringType, "marketplace_offering_uuid": types.StringType, "marketplace_plan_uuid": types.StringType, "marketplace_resource_state": types.StringType, "marketplace_resource_uuid": types.StringType, "modified": types.StringType, "name": types.StringType, "project": types.StringType, "project_name": types.StringType, "project_uuid": types.StringType, "resource_type": types.StringType, "rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}}, "service_name": types.StringType, "service_settings": types.StringType, "service_settings_error_message": types.StringType, "service_settings_state": types.StringType, "service_settings_uuid": types.StringType, "state": types.StringType, "tenant": types.StringType, "tenant_name": types.StringType, "tenant_uuid": types.StringType, "url": types.StringType}}},
						"subnet_cidr":           types.StringType,
						"subnet_description":    types.StringType,
						"subnet_name":           types.StringType,
						"subnet_uuid":           types.StringType,
						"url":                   types.StringType,
					}
					attrValues := map[string]attr.Value{
						"fixed_ips": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}}.ElemType),
						"port": func() attr.Value {
							if v, ok := objMap["port"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet": func() attr.Value {
							if v, ok := objMap["subnet"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"allowed_address_pairs": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"mac_address": types.StringType}}}.ElemType),
						"device_id": func() attr.Value {
							if v, ok := objMap["device_id"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"device_owner": func() attr.Value {
							if v, ok := objMap["device_owner"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"mac_address": func() attr.Value {
							if v, ok := objMap["mac_address"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"security_groups": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"access_url": types.StringType, "backend_id": types.StringType, "created": types.StringType, "customer": types.StringType, "customer_abbreviation": types.StringType, "customer_name": types.StringType, "customer_native_name": types.StringType, "customer_uuid": types.StringType, "description": types.StringType, "error_message": types.StringType, "error_traceback": types.StringType, "is_limit_based": types.BoolType, "is_usage_based": types.BoolType, "marketplace_category_name": types.StringType, "marketplace_category_uuid": types.StringType, "marketplace_offering_name": types.StringType, "marketplace_offering_uuid": types.StringType, "marketplace_plan_uuid": types.StringType, "marketplace_resource_state": types.StringType, "marketplace_resource_uuid": types.StringType, "modified": types.StringType, "name": types.StringType, "project": types.StringType, "project_name": types.StringType, "project_uuid": types.StringType, "resource_type": types.StringType, "rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}}, "service_name": types.StringType, "service_settings": types.StringType, "service_settings_error_message": types.StringType, "service_settings_state": types.StringType, "service_settings_uuid": types.StringType, "state": types.StringType, "tenant": types.StringType, "tenant_name": types.StringType, "tenant_uuid": types.StringType, "url": types.StringType}}}.ElemType),
						"subnet_cidr": func() attr.Value {
							if v, ok := objMap["subnet_cidr"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet_description": func() attr.Value {
							if v, ok := objMap["subnet_description"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet_name": func() attr.Value {
							if v, ok := objMap["subnet_name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet_uuid": func() attr.Value {
							if v, ok := objMap["subnet_uuid"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"fixed_ips":             types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
				"port":                  types.StringType,
				"subnet":                types.StringType,
				"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"mac_address": types.StringType}}},
				"device_id":             types.StringType,
				"device_owner":          types.StringType,
				"mac_address":           types.StringType,
				"security_groups":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"access_url": types.StringType, "backend_id": types.StringType, "created": types.StringType, "customer": types.StringType, "customer_abbreviation": types.StringType, "customer_name": types.StringType, "customer_native_name": types.StringType, "customer_uuid": types.StringType, "description": types.StringType, "error_message": types.StringType, "error_traceback": types.StringType, "is_limit_based": types.BoolType, "is_usage_based": types.BoolType, "marketplace_category_name": types.StringType, "marketplace_category_uuid": types.StringType, "marketplace_offering_name": types.StringType, "marketplace_offering_uuid": types.StringType, "marketplace_plan_uuid": types.StringType, "marketplace_resource_state": types.StringType, "marketplace_resource_uuid": types.StringType, "modified": types.StringType, "name": types.StringType, "project": types.StringType, "project_name": types.StringType, "project_uuid": types.StringType, "resource_type": types.StringType, "rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}}, "service_name": types.StringType, "service_settings": types.StringType, "service_settings_error_message": types.StringType, "service_settings_state": types.StringType, "service_settings_uuid": types.StringType, "state": types.StringType, "tenant": types.StringType, "tenant_name": types.StringType, "tenant_uuid": types.StringType, "url": types.StringType}}},
				"subnet_cidr":           types.StringType,
				"subnet_description":    types.StringType,
				"subnet_name":           types.StringType,
				"subnet_uuid":           types.StringType,
				"url":                   types.StringType,
			}}, items)
			data.Ports = listVal
		}
	} else {
		if data.Ports.IsUnknown() {
			data.Ports = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"fixed_ips":             types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
				"port":                  types.StringType,
				"subnet":                types.StringType,
				"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"mac_address": types.StringType}}},
				"device_id":             types.StringType,
				"device_owner":          types.StringType,
				"mac_address":           types.StringType,
				"security_groups":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"access_url": types.StringType, "backend_id": types.StringType, "created": types.StringType, "customer": types.StringType, "customer_abbreviation": types.StringType, "customer_name": types.StringType, "customer_native_name": types.StringType, "customer_uuid": types.StringType, "description": types.StringType, "error_message": types.StringType, "error_traceback": types.StringType, "is_limit_based": types.BoolType, "is_usage_based": types.BoolType, "marketplace_category_name": types.StringType, "marketplace_category_uuid": types.StringType, "marketplace_offering_name": types.StringType, "marketplace_offering_uuid": types.StringType, "marketplace_plan_uuid": types.StringType, "marketplace_resource_state": types.StringType, "marketplace_resource_uuid": types.StringType, "modified": types.StringType, "name": types.StringType, "project": types.StringType, "project_name": types.StringType, "project_uuid": types.StringType, "resource_type": types.StringType, "rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}}, "service_name": types.StringType, "service_settings": types.StringType, "service_settings_error_message": types.StringType, "service_settings_state": types.StringType, "service_settings_uuid": types.StringType, "state": types.StringType, "tenant": types.StringType, "tenant_name": types.StringType, "tenant_uuid": types.StringType, "url": types.StringType}}},
				"subnet_cidr":           types.StringType,
				"subnet_description":    types.StringType,
				"subnet_name":           types.StringType,
				"subnet_uuid":           types.StringType,
				"url":                   types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["project"]; ok && val != nil {
		if str, ok := val.(string); ok {
			// Normalize URL to UUID
			parts := strings.Split(strings.TrimRight(str, "/"), "/")
			uuid := parts[len(parts)-1]
			data.Project = types.StringValue(uuid)
		} else {
			data.Project = types.StringNull()
		}
	} else {
		if data.Project.IsUnknown() {
			data.Project = types.StringNull()
		}
	}
	if val, ok := sourceMap["project_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ProjectName = types.StringValue(str)
		}
	} else {
		if data.ProjectName.IsUnknown() {
			data.ProjectName = types.StringNull()
		}
	}
	if val, ok := sourceMap["project_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ProjectUuid = types.StringValue(str)
		}
	} else {
		if data.ProjectUuid.IsUnknown() {
			data.ProjectUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["ram"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Ram = types.Int64Value(int64(num))
		}
	} else {
		if data.Ram.IsUnknown() {
			data.Ram = types.Int64Null()
		}
	}
	if val, ok := sourceMap["release_floating_ips"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.ReleaseFloatingIps = types.BoolValue(b)
		}
	} else {
		if data.ReleaseFloatingIps.IsUnknown() {
			data.ReleaseFloatingIps = types.BoolNull()
		}
	}
	if val, ok := sourceMap["resource_type"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ResourceType = types.StringValue(str)
		}
	} else {
		if data.ResourceType.IsUnknown() {
			data.ResourceType = types.StringNull()
		}
	}
	if val, ok := sourceMap["runtime_state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.RuntimeState = types.StringValue(str)
		}
	} else {
		if data.RuntimeState.IsUnknown() {
			data.RuntimeState = types.StringNull()
		}
	}
	if val, ok := sourceMap["security_groups"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"url":         types.StringType,
						"description": types.StringType,
						"name":        types.StringType,
						"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}},
						"state":       types.StringType,
					}
					attrValues := map[string]attr.Value{
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"description": func() attr.Value {
							if v, ok := objMap["description"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"rules": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}}.ElemType),
						"state": func() attr.Value {
							if v, ok := objMap["state"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"url":         types.StringType,
				"description": types.StringType,
				"name":        types.StringType,
				"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}},
				"state":       types.StringType,
			}}, items)
			data.SecurityGroups = listVal
		}
	} else {
		if data.SecurityGroups.IsUnknown() {
			data.SecurityGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"url":         types.StringType,
				"description": types.StringType,
				"name":        types.StringType,
				"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}},
				"state":       types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["server_group"]; ok && val != nil {
		// Nested object
		if objMap, ok := val.(map[string]interface{}); ok {
			_ = objMap // Avoid unused variable if properties are empty
			attrTypes := map[string]attr.Type{
				"name":   types.StringType,
				"policy": types.StringType,
				"state":  types.StringType,
				"url":    types.StringType,
			}
			attrValues := map[string]attr.Value{
				"name": func() attr.Value {
					if v, ok := objMap["name"].(string); ok {
						return types.StringValue(v)
					}
					return types.StringNull()
				}(),
				"policy": func() attr.Value {
					if v, ok := objMap["policy"].(string); ok {
						return types.StringValue(v)
					}
					return types.StringNull()
				}(),
				"state": func() attr.Value {
					if v, ok := objMap["state"].(string); ok {
						return types.StringValue(v)
					}
					return types.StringNull()
				}(),
				"url": func() attr.Value {
					if v, ok := objMap["url"].(string); ok {
						return types.StringValue(v)
					}
					return types.StringNull()
				}(),
			}
			objVal, _ := types.ObjectValue(attrTypes, attrValues)
			data.ServerGroup = objVal
		}
	} else {
		if data.ServerGroup.IsUnknown() {
			data.ServerGroup = types.ObjectNull(map[string]attr.Type{
				"name":   types.StringType,
				"policy": types.StringType,
				"state":  types.StringType,
				"url":    types.StringType,
			})
		}
	}
	if val, ok := sourceMap["service_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ServiceName = types.StringValue(str)
		}
	} else {
		if data.ServiceName.IsUnknown() {
			data.ServiceName = types.StringNull()
		}
	}
	if val, ok := sourceMap["service_settings"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ServiceSettings = types.StringValue(str)
		}
	} else {
		if data.ServiceSettings.IsUnknown() {
			data.ServiceSettings = types.StringNull()
		}
	}
	if val, ok := sourceMap["service_settings_error_message"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ServiceSettingsErrorMessage = types.StringValue(str)
		}
	} else {
		if data.ServiceSettingsErrorMessage.IsUnknown() {
			data.ServiceSettingsErrorMessage = types.StringNull()
		}
	}
	if val, ok := sourceMap["service_settings_state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ServiceSettingsState = types.StringValue(str)
		}
	} else {
		if data.ServiceSettingsState.IsUnknown() {
			data.ServiceSettingsState = types.StringNull()
		}
	}
	if val, ok := sourceMap["service_settings_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ServiceSettingsUuid = types.StringValue(str)
		}
	} else {
		if data.ServiceSettingsUuid.IsUnknown() {
			data.ServiceSettingsUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["ssh_public_key"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.SshPublicKey = types.StringValue(str)
		}
	} else {
		if data.SshPublicKey.IsUnknown() {
			data.SshPublicKey = types.StringNull()
		}
	}
	if val, ok := sourceMap["start_time"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.StartTime = types.StringValue(str)
		}
	} else {
		if data.StartTime.IsUnknown() {
			data.StartTime = types.StringNull()
		}
	}
	if val, ok := sourceMap["state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.State = types.StringValue(str)
		}
	} else {
		if data.State.IsUnknown() {
			data.State = types.StringNull()
		}
	}
	if val, ok := sourceMap["system_volume_size"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.SystemVolumeSize = types.Int64Value(int64(num))
		}
	} else {
		if data.SystemVolumeSize.IsUnknown() {
			data.SystemVolumeSize = types.Int64Null()
		}
	}
	if val, ok := sourceMap["system_volume_type"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.SystemVolumeType = types.StringValue(str)
		}
	} else {
		if data.SystemVolumeType.IsUnknown() {
			data.SystemVolumeType = types.StringNull()
		}
	}
	if val, ok := sourceMap["tenant"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Tenant = types.StringValue(str)
		}
	} else {
		if data.Tenant.IsUnknown() {
			data.Tenant = types.StringNull()
		}
	}
	if val, ok := sourceMap["tenant_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.TenantUuid = types.StringValue(str)
		}
	} else {
		if data.TenantUuid.IsUnknown() {
			data.TenantUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["url"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Url = types.StringValue(str)
		}
	} else {
		if data.Url.IsUnknown() {
			data.Url = types.StringNull()
		}
	}
	if val, ok := sourceMap["user_data"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.UserData = types.StringValue(str)
		}
	} else {
		if data.UserData.IsUnknown() {
			data.UserData = types.StringNull()
		}
	}
	if val, ok := sourceMap["volumes"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
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
					}
					attrValues := map[string]attr.Value{
						"bootable": func() attr.Value {
							if v, ok := objMap["bootable"].(bool); ok {
								return types.BoolValue(v)
							}
							return types.BoolNull()
						}(),
						"device": func() attr.Value {
							if v, ok := objMap["device"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"image_name": func() attr.Value {
							if v, ok := objMap["image_name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"marketplace_resource_uuid": func() attr.Value {
							if v, ok := objMap["marketplace_resource_uuid"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"resource_type": func() attr.Value {
							if v, ok := objMap["resource_type"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"size": func() attr.Value {
							if v, ok := objMap["size"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"state": func() attr.Value {
							if v, ok := objMap["state"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"type": func() attr.Value {
							if v, ok := objMap["type"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"type_name": func() attr.Value {
							if v, ok := objMap["type_name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}}, items)
			data.Volumes = listVal
		}
	} else {
		if data.Volumes.IsUnknown() {
			data.Volumes = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}})
		}
	}

	// Map filter parameters from response if available

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackInstanceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpenstackInstanceResourceModel
	var state OpenstackInstanceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	// Read current state to get the UUID (which is computed and not in plan)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Use UUID from state
	data.UUID = state.UUID
	// Phase 1: Standard PATCH (Simple fields)
	patchPayload := map[string]interface{}{}
	if !data.Description.IsNull() && !data.Description.Equal(state.Description) {
		patchPayload["description"] = data.Description.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.Equal(state.Name) {
		patchPayload["name"] = data.Name.ValueString()
	}

	if len(patchPayload) > 0 {
		var result map[string]interface{}
		err := r.client.Update(ctx, "/api/openstack-instances/{uuid}/", data.UUID.ValueString(), patchPayload, &result)
		if err != nil {
			resp.Diagnostics.AddError("Update Failed", err.Error())
			return
		}
	}

	// Phase 2: RPC Actions
	if os.Getenv("WALDUR_E2E_SKIP_WAIT") != "" {
		tflog.Warn(ctx, "Skipping wait for update order completion due to WALDUR_E2E_SKIP_WAIT")
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}
	if !data.FloatingIps.Equal(state.FloatingIps) {

		// Convert Terraform value to API payload

		var itemsUpdateFloatingIps []interface{}
		for _, elem := range data.FloatingIps.Elements() {
			if objVal, ok := elem.(types.Object); ok {
				objMap := make(map[string]interface{})
				for key, attr := range objVal.Attributes() {
					switch v := attr.(type) {
					case types.String:
						objMap[key] = v.ValueString()
					case types.Int64:
						objMap[key] = v.ValueInt64()
					case types.Bool:
						objMap[key] = v.ValueBool()
					case types.Float64:
						objMap[key] = v.ValueFloat64()
					}
				}
				itemsUpdateFloatingIps = append(itemsUpdateFloatingIps, objMap)
			}
		}
		actionPayloadUpdateFloatingIps := map[string]interface{}{
			"floating_ips": itemsUpdateFloatingIps,
		}
		actionUrlUpdateFloatingIps := strings.Replace("/api/openstack-instances/{uuid}/update_floating_ips/", "{uuid}", data.UUID.ValueString(), 1)
		var actionResultUpdateFloatingIps map[string]interface{}
		if err := r.client.Post(ctx, actionUrlUpdateFloatingIps, actionPayloadUpdateFloatingIps, &actionResultUpdateFloatingIps); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: update_floating_ips", err.Error())
			return
		}
	}
	if !data.Ports.Equal(state.Ports) {

		// Convert Terraform value to API payload

		var itemsUpdatePorts []interface{}
		for _, elem := range data.Ports.Elements() {
			if objVal, ok := elem.(types.Object); ok {
				objMap := make(map[string]interface{})
				for key, attr := range objVal.Attributes() {
					switch v := attr.(type) {
					case types.String:
						objMap[key] = v.ValueString()
					case types.Int64:
						objMap[key] = v.ValueInt64()
					case types.Bool:
						objMap[key] = v.ValueBool()
					case types.Float64:
						objMap[key] = v.ValueFloat64()
					}
				}
				itemsUpdatePorts = append(itemsUpdatePorts, objMap)
			}
		}
		actionPayloadUpdatePorts := map[string]interface{}{
			"ports": itemsUpdatePorts,
		}
		actionUrlUpdatePorts := strings.Replace("/api/openstack-instances/{uuid}/update_ports/", "{uuid}", data.UUID.ValueString(), 1)
		var actionResultUpdatePorts map[string]interface{}
		if err := r.client.Post(ctx, actionUrlUpdatePorts, actionPayloadUpdatePorts, &actionResultUpdatePorts); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: update_ports", err.Error())
			return
		}
	}
	if !data.SecurityGroups.Equal(state.SecurityGroups) {

		// Convert Terraform value to API payload

		var itemsUpdateSecurityGroups []interface{}
		for _, elem := range data.SecurityGroups.Elements() {
			if objVal, ok := elem.(types.Object); ok {
				objMap := make(map[string]interface{})
				for key, attr := range objVal.Attributes() {
					switch v := attr.(type) {
					case types.String:
						objMap[key] = v.ValueString()
					case types.Int64:
						objMap[key] = v.ValueInt64()
					case types.Bool:
						objMap[key] = v.ValueBool()
					case types.Float64:
						objMap[key] = v.ValueFloat64()
					}
				}
				itemsUpdateSecurityGroups = append(itemsUpdateSecurityGroups, objMap)
			}
		}
		actionPayloadUpdateSecurityGroups := map[string]interface{}{
			"security_groups": itemsUpdateSecurityGroups,
		}
		actionUrlUpdateSecurityGroups := strings.Replace("/api/openstack-instances/{uuid}/update_security_groups/", "{uuid}", data.UUID.ValueString(), 1)
		var actionResultUpdateSecurityGroups map[string]interface{}
		if err := r.client.Post(ctx, actionUrlUpdateSecurityGroups, actionPayloadUpdateSecurityGroups, &actionResultUpdateSecurityGroups); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: update_security_groups", err.Error())
			return
		}
	}

	// Fetch updated state
	// Call Waldur API to read resource
	var result map[string]interface{}

	retrievePath := strings.Replace("/api/openstack-instances/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &result)
	if err != nil {
		resp.Diagnostics.AddError("Failed to Read Resource After Update", err.Error())
		return
	}
	sourceMap := result
	// Map response fields to data model
	_ = sourceMap
	if val, ok := sourceMap["access_url"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AccessUrl = types.StringValue(str)
		}
	} else {
		if data.AccessUrl.IsUnknown() {
			data.AccessUrl = types.StringNull()
		}
	}
	if val, ok := sourceMap["action"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Action = types.StringValue(str)
		}
	} else {
		if data.Action.IsUnknown() {
			data.Action = types.StringNull()
		}
	}
	if val, ok := sourceMap["availability_zone"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AvailabilityZone = types.StringValue(str)
		}
	} else {
		if data.AvailabilityZone.IsUnknown() {
			data.AvailabilityZone = types.StringNull()
		}
	}
	if val, ok := sourceMap["availability_zone_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AvailabilityZoneName = types.StringValue(str)
		}
	} else {
		if data.AvailabilityZoneName.IsUnknown() {
			data.AvailabilityZoneName = types.StringNull()
		}
	}
	if val, ok := sourceMap["backend_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.BackendId = types.StringValue(str)
		}
	} else {
		if data.BackendId.IsUnknown() {
			data.BackendId = types.StringNull()
		}
	}
	if val, ok := sourceMap["connect_directly_to_external_network"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.ConnectDirectlyToExternalNetwork = types.BoolValue(b)
		}
	} else {
		if data.ConnectDirectlyToExternalNetwork.IsUnknown() {
			data.ConnectDirectlyToExternalNetwork = types.BoolNull()
		}
	}
	if val, ok := sourceMap["cores"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Cores = types.Int64Value(int64(num))
		}
	} else {
		if data.Cores.IsUnknown() {
			data.Cores = types.Int64Null()
		}
	}
	if val, ok := sourceMap["created"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Created = types.StringValue(str)
		}
	} else {
		if data.Created.IsUnknown() {
			data.Created = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Customer = types.StringValue(str)
		}
	} else {
		if data.Customer.IsUnknown() {
			data.Customer = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer_abbreviation"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CustomerAbbreviation = types.StringValue(str)
		}
	} else {
		if data.CustomerAbbreviation.IsUnknown() {
			data.CustomerAbbreviation = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CustomerName = types.StringValue(str)
		}
	} else {
		if data.CustomerName.IsUnknown() {
			data.CustomerName = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer_native_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CustomerNativeName = types.StringValue(str)
		}
	} else {
		if data.CustomerNativeName.IsUnknown() {
			data.CustomerNativeName = types.StringNull()
		}
	}
	if val, ok := sourceMap["customer_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.CustomerUuid = types.StringValue(str)
		}
	} else {
		if data.CustomerUuid.IsUnknown() {
			data.CustomerUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["data_volume_size"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.DataVolumeSize = types.Int64Value(int64(num))
		}
	} else {
		if data.DataVolumeSize.IsUnknown() {
			data.DataVolumeSize = types.Int64Null()
		}
	}
	if val, ok := sourceMap["data_volume_type"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.DataVolumeType = types.StringValue(str)
		}
	} else {
		if data.DataVolumeType.IsUnknown() {
			data.DataVolumeType = types.StringNull()
		}
	}
	if val, ok := sourceMap["data_volumes"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"size":        types.Int64Type,
						"volume_type": types.StringType,
					}
					attrValues := map[string]attr.Value{
						"size": func() attr.Value {
							if v, ok := objMap["size"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"volume_type": func() attr.Value {
							if v, ok := objMap["volume_type"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"size":        types.Int64Type,
				"volume_type": types.StringType,
			}}, items)
			data.DataVolumes = listVal
		}
	} else {
		if data.DataVolumes.IsUnknown() {
			data.DataVolumes = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"size":        types.Int64Type,
				"volume_type": types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["delete_volumes"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.DeleteVolumes = types.BoolValue(b)
		}
	} else {
		if data.DeleteVolumes.IsUnknown() {
			data.DeleteVolumes = types.BoolNull()
		}
	}
	if val, ok := sourceMap["description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
		}
	} else {
		if data.Description.IsUnknown() {
			data.Description = types.StringNull()
		}
	}
	if val, ok := sourceMap["disk"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Disk = types.Int64Value(int64(num))
		}
	} else {
		if data.Disk.IsUnknown() {
			data.Disk = types.Int64Null()
		}
	}
	if val, ok := sourceMap["error_message"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ErrorMessage = types.StringValue(str)
		}
	} else {
		if data.ErrorMessage.IsUnknown() {
			data.ErrorMessage = types.StringNull()
		}
	}
	if val, ok := sourceMap["error_traceback"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ErrorTraceback = types.StringValue(str)
		}
	} else {
		if data.ErrorTraceback.IsUnknown() {
			data.ErrorTraceback = types.StringNull()
		}
	}
	if val, ok := sourceMap["external_address"]; ok && val != nil {
		// List of strings (or flattened objects)
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if str, ok := item.(string); ok {
					items = append(items, types.StringValue(str))
				} else if obj, ok := item.(map[string]interface{}); ok {
					// Flattening logic: extract URL or UUID
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					} else if name, ok := obj["name"].(string); ok {
						items = append(items, types.StringValue(name))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.ExternalAddress = listVal
		}
	} else {
		if data.ExternalAddress.IsUnknown() {
			data.ExternalAddress = types.ListNull(types.StringType)
		}
	}
	if val, ok := sourceMap["external_ips"]; ok && val != nil {
		// List of strings (or flattened objects)
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if str, ok := item.(string); ok {
					items = append(items, types.StringValue(str))
				} else if obj, ok := item.(map[string]interface{}); ok {
					// Flattening logic: extract URL or UUID
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					} else if name, ok := obj["name"].(string); ok {
						items = append(items, types.StringValue(name))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.ExternalIps = listVal
		}
	} else {
		if data.ExternalIps.IsUnknown() {
			data.ExternalIps = types.ListNull(types.StringType)
		}
	}
	if val, ok := sourceMap["flavor"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Flavor = types.StringValue(str)
		}
	} else {
		if data.Flavor.IsUnknown() {
			data.Flavor = types.StringNull()
		}
	}
	if val, ok := sourceMap["flavor_disk"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.FlavorDisk = types.Int64Value(int64(num))
		}
	} else {
		if data.FlavorDisk.IsUnknown() {
			data.FlavorDisk = types.Int64Null()
		}
	}
	if val, ok := sourceMap["flavor_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.FlavorName = types.StringValue(str)
		}
	} else {
		if data.FlavorName.IsUnknown() {
			data.FlavorName = types.StringNull()
		}
	}
	if val, ok := sourceMap["floating_ips"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"ip_address":         types.StringType,
						"subnet":             types.StringType,
						"url":                types.StringType,
						"address":            types.StringType,
						"port_fixed_ips":     types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
						"port_mac_address":   types.StringType,
						"subnet_cidr":        types.StringType,
						"subnet_description": types.StringType,
						"subnet_name":        types.StringType,
						"subnet_uuid":        types.StringType,
					}
					attrValues := map[string]attr.Value{
						"ip_address": func() attr.Value {
							if v, ok := objMap["ip_address"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet": func() attr.Value {
							if v, ok := objMap["subnet"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"address": func() attr.Value {
							if v, ok := objMap["address"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"port_fixed_ips": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}}.ElemType),
						"port_mac_address": func() attr.Value {
							if v, ok := objMap["port_mac_address"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet_cidr": func() attr.Value {
							if v, ok := objMap["subnet_cidr"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet_description": func() attr.Value {
							if v, ok := objMap["subnet_description"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet_name": func() attr.Value {
							if v, ok := objMap["subnet_name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet_uuid": func() attr.Value {
							if v, ok := objMap["subnet_uuid"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"ip_address":         types.StringType,
				"subnet":             types.StringType,
				"url":                types.StringType,
				"address":            types.StringType,
				"port_fixed_ips":     types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
				"port_mac_address":   types.StringType,
				"subnet_cidr":        types.StringType,
				"subnet_description": types.StringType,
				"subnet_name":        types.StringType,
				"subnet_uuid":        types.StringType,
			}}, items)
			data.FloatingIps = listVal
		}
	} else {
		if data.FloatingIps.IsUnknown() {
			data.FloatingIps = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"ip_address":         types.StringType,
				"subnet":             types.StringType,
				"url":                types.StringType,
				"address":            types.StringType,
				"port_fixed_ips":     types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
				"port_mac_address":   types.StringType,
				"subnet_cidr":        types.StringType,
				"subnet_description": types.StringType,
				"subnet_name":        types.StringType,
				"subnet_uuid":        types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["hypervisor_hostname"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.HypervisorHostname = types.StringValue(str)
		}
	} else {
		if data.HypervisorHostname.IsUnknown() {
			data.HypervisorHostname = types.StringNull()
		}
	}
	if val, ok := sourceMap["image"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Image = types.StringValue(str)
		}
	} else {
		if data.Image.IsUnknown() {
			data.Image = types.StringNull()
		}
	}
	if val, ok := sourceMap["image_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ImageName = types.StringValue(str)
		}
	} else {
		if data.ImageName.IsUnknown() {
			data.ImageName = types.StringNull()
		}
	}
	if val, ok := sourceMap["internal_ips"]; ok && val != nil {
		// List of strings (or flattened objects)
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if str, ok := item.(string); ok {
					items = append(items, types.StringValue(str))
				} else if obj, ok := item.(map[string]interface{}); ok {
					// Flattening logic: extract URL or UUID
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					} else if name, ok := obj["name"].(string); ok {
						items = append(items, types.StringValue(name))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.InternalIps = listVal
		}
	} else {
		if data.InternalIps.IsUnknown() {
			data.InternalIps = types.ListNull(types.StringType)
		}
	}
	if val, ok := sourceMap["is_limit_based"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.IsLimitBased = types.BoolValue(b)
		}
	} else {
		if data.IsLimitBased.IsUnknown() {
			data.IsLimitBased = types.BoolNull()
		}
	}
	if val, ok := sourceMap["is_usage_based"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.IsUsageBased = types.BoolValue(b)
		}
	} else {
		if data.IsUsageBased.IsUnknown() {
			data.IsUsageBased = types.BoolNull()
		}
	}
	if val, ok := sourceMap["key_fingerprint"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.KeyFingerprint = types.StringValue(str)
		}
	} else {
		if data.KeyFingerprint.IsUnknown() {
			data.KeyFingerprint = types.StringNull()
		}
	}
	if val, ok := sourceMap["key_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.KeyName = types.StringValue(str)
		}
	} else {
		if data.KeyName.IsUnknown() {
			data.KeyName = types.StringNull()
		}
	}
	if val, ok := sourceMap["latitude"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Latitude = types.Float64Value(num)
		}
	} else {
		if data.Latitude.IsUnknown() {
			data.Latitude = types.Float64Null()
		}
	}
	if val, ok := sourceMap["longitude"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Longitude = types.Float64Value(num)
		}
	} else {
		if data.Longitude.IsUnknown() {
			data.Longitude = types.Float64Null()
		}
	}
	if val, ok := sourceMap["marketplace_category_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.MarketplaceCategoryName = types.StringValue(str)
		}
	} else {
		if data.MarketplaceCategoryName.IsUnknown() {
			data.MarketplaceCategoryName = types.StringNull()
		}
	}
	if val, ok := sourceMap["marketplace_category_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.MarketplaceCategoryUuid = types.StringValue(str)
		}
	} else {
		if data.MarketplaceCategoryUuid.IsUnknown() {
			data.MarketplaceCategoryUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["marketplace_offering_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.MarketplaceOfferingName = types.StringValue(str)
		}
	} else {
		if data.MarketplaceOfferingName.IsUnknown() {
			data.MarketplaceOfferingName = types.StringNull()
		}
	}
	if val, ok := sourceMap["marketplace_offering_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.MarketplaceOfferingUuid = types.StringValue(str)
		}
	} else {
		if data.MarketplaceOfferingUuid.IsUnknown() {
			data.MarketplaceOfferingUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["marketplace_plan_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.MarketplacePlanUuid = types.StringValue(str)
		}
	} else {
		if data.MarketplacePlanUuid.IsUnknown() {
			data.MarketplacePlanUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["marketplace_resource_state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.MarketplaceResourceState = types.StringValue(str)
		}
	} else {
		if data.MarketplaceResourceState.IsUnknown() {
			data.MarketplaceResourceState = types.StringNull()
		}
	}
	if val, ok := sourceMap["marketplace_resource_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.MarketplaceResourceUuid = types.StringValue(str)
		}
	} else {
		if data.MarketplaceResourceUuid.IsUnknown() {
			data.MarketplaceResourceUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["min_disk"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.MinDisk = types.Int64Value(int64(num))
		}
	} else {
		if data.MinDisk.IsUnknown() {
			data.MinDisk = types.Int64Null()
		}
	}
	if val, ok := sourceMap["min_ram"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.MinRam = types.Int64Value(int64(num))
		}
	} else {
		if data.MinRam.IsUnknown() {
			data.MinRam = types.Int64Null()
		}
	}
	if val, ok := sourceMap["modified"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Modified = types.StringValue(str)
		}
	} else {
		if data.Modified.IsUnknown() {
			data.Modified = types.StringNull()
		}
	}
	if val, ok := sourceMap["offering"]; ok && val != nil {
		if str, ok := val.(string); ok {
			// Normalize URL to UUID
			parts := strings.Split(strings.TrimRight(str, "/"), "/")
			uuid := parts[len(parts)-1]
			data.Offering = types.StringValue(uuid)
		} else {
			data.Offering = types.StringNull()
		}
	} else {
		if data.Offering.IsUnknown() {
			data.Offering = types.StringNull()
		}
	}
	if val, ok := sourceMap["ports"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"fixed_ips":             types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
						"port":                  types.StringType,
						"subnet":                types.StringType,
						"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"mac_address": types.StringType}}},
						"device_id":             types.StringType,
						"device_owner":          types.StringType,
						"mac_address":           types.StringType,
						"security_groups":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"access_url": types.StringType, "backend_id": types.StringType, "created": types.StringType, "customer": types.StringType, "customer_abbreviation": types.StringType, "customer_name": types.StringType, "customer_native_name": types.StringType, "customer_uuid": types.StringType, "description": types.StringType, "error_message": types.StringType, "error_traceback": types.StringType, "is_limit_based": types.BoolType, "is_usage_based": types.BoolType, "marketplace_category_name": types.StringType, "marketplace_category_uuid": types.StringType, "marketplace_offering_name": types.StringType, "marketplace_offering_uuid": types.StringType, "marketplace_plan_uuid": types.StringType, "marketplace_resource_state": types.StringType, "marketplace_resource_uuid": types.StringType, "modified": types.StringType, "name": types.StringType, "project": types.StringType, "project_name": types.StringType, "project_uuid": types.StringType, "resource_type": types.StringType, "rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}}, "service_name": types.StringType, "service_settings": types.StringType, "service_settings_error_message": types.StringType, "service_settings_state": types.StringType, "service_settings_uuid": types.StringType, "state": types.StringType, "tenant": types.StringType, "tenant_name": types.StringType, "tenant_uuid": types.StringType, "url": types.StringType}}},
						"subnet_cidr":           types.StringType,
						"subnet_description":    types.StringType,
						"subnet_name":           types.StringType,
						"subnet_uuid":           types.StringType,
						"url":                   types.StringType,
					}
					attrValues := map[string]attr.Value{
						"fixed_ips": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}}.ElemType),
						"port": func() attr.Value {
							if v, ok := objMap["port"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet": func() attr.Value {
							if v, ok := objMap["subnet"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"allowed_address_pairs": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"mac_address": types.StringType}}}.ElemType),
						"device_id": func() attr.Value {
							if v, ok := objMap["device_id"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"device_owner": func() attr.Value {
							if v, ok := objMap["device_owner"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"mac_address": func() attr.Value {
							if v, ok := objMap["mac_address"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"security_groups": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"access_url": types.StringType, "backend_id": types.StringType, "created": types.StringType, "customer": types.StringType, "customer_abbreviation": types.StringType, "customer_name": types.StringType, "customer_native_name": types.StringType, "customer_uuid": types.StringType, "description": types.StringType, "error_message": types.StringType, "error_traceback": types.StringType, "is_limit_based": types.BoolType, "is_usage_based": types.BoolType, "marketplace_category_name": types.StringType, "marketplace_category_uuid": types.StringType, "marketplace_offering_name": types.StringType, "marketplace_offering_uuid": types.StringType, "marketplace_plan_uuid": types.StringType, "marketplace_resource_state": types.StringType, "marketplace_resource_uuid": types.StringType, "modified": types.StringType, "name": types.StringType, "project": types.StringType, "project_name": types.StringType, "project_uuid": types.StringType, "resource_type": types.StringType, "rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}}, "service_name": types.StringType, "service_settings": types.StringType, "service_settings_error_message": types.StringType, "service_settings_state": types.StringType, "service_settings_uuid": types.StringType, "state": types.StringType, "tenant": types.StringType, "tenant_name": types.StringType, "tenant_uuid": types.StringType, "url": types.StringType}}}.ElemType),
						"subnet_cidr": func() attr.Value {
							if v, ok := objMap["subnet_cidr"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet_description": func() attr.Value {
							if v, ok := objMap["subnet_description"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet_name": func() attr.Value {
							if v, ok := objMap["subnet_name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"subnet_uuid": func() attr.Value {
							if v, ok := objMap["subnet_uuid"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"fixed_ips":             types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
				"port":                  types.StringType,
				"subnet":                types.StringType,
				"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"mac_address": types.StringType}}},
				"device_id":             types.StringType,
				"device_owner":          types.StringType,
				"mac_address":           types.StringType,
				"security_groups":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"access_url": types.StringType, "backend_id": types.StringType, "created": types.StringType, "customer": types.StringType, "customer_abbreviation": types.StringType, "customer_name": types.StringType, "customer_native_name": types.StringType, "customer_uuid": types.StringType, "description": types.StringType, "error_message": types.StringType, "error_traceback": types.StringType, "is_limit_based": types.BoolType, "is_usage_based": types.BoolType, "marketplace_category_name": types.StringType, "marketplace_category_uuid": types.StringType, "marketplace_offering_name": types.StringType, "marketplace_offering_uuid": types.StringType, "marketplace_plan_uuid": types.StringType, "marketplace_resource_state": types.StringType, "marketplace_resource_uuid": types.StringType, "modified": types.StringType, "name": types.StringType, "project": types.StringType, "project_name": types.StringType, "project_uuid": types.StringType, "resource_type": types.StringType, "rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}}, "service_name": types.StringType, "service_settings": types.StringType, "service_settings_error_message": types.StringType, "service_settings_state": types.StringType, "service_settings_uuid": types.StringType, "state": types.StringType, "tenant": types.StringType, "tenant_name": types.StringType, "tenant_uuid": types.StringType, "url": types.StringType}}},
				"subnet_cidr":           types.StringType,
				"subnet_description":    types.StringType,
				"subnet_name":           types.StringType,
				"subnet_uuid":           types.StringType,
				"url":                   types.StringType,
			}}, items)
			data.Ports = listVal
		}
	} else {
		if data.Ports.IsUnknown() {
			data.Ports = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"fixed_ips":             types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"ip_address": types.StringType, "subnet_id": types.StringType}}},
				"port":                  types.StringType,
				"subnet":                types.StringType,
				"allowed_address_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"mac_address": types.StringType}}},
				"device_id":             types.StringType,
				"device_owner":          types.StringType,
				"mac_address":           types.StringType,
				"security_groups":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"access_url": types.StringType, "backend_id": types.StringType, "created": types.StringType, "customer": types.StringType, "customer_abbreviation": types.StringType, "customer_name": types.StringType, "customer_native_name": types.StringType, "customer_uuid": types.StringType, "description": types.StringType, "error_message": types.StringType, "error_traceback": types.StringType, "is_limit_based": types.BoolType, "is_usage_based": types.BoolType, "marketplace_category_name": types.StringType, "marketplace_category_uuid": types.StringType, "marketplace_offering_name": types.StringType, "marketplace_offering_uuid": types.StringType, "marketplace_plan_uuid": types.StringType, "marketplace_resource_state": types.StringType, "marketplace_resource_uuid": types.StringType, "modified": types.StringType, "name": types.StringType, "project": types.StringType, "project_name": types.StringType, "project_uuid": types.StringType, "resource_type": types.StringType, "rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}}, "service_name": types.StringType, "service_settings": types.StringType, "service_settings_error_message": types.StringType, "service_settings_state": types.StringType, "service_settings_uuid": types.StringType, "state": types.StringType, "tenant": types.StringType, "tenant_name": types.StringType, "tenant_uuid": types.StringType, "url": types.StringType}}},
				"subnet_cidr":           types.StringType,
				"subnet_description":    types.StringType,
				"subnet_name":           types.StringType,
				"subnet_uuid":           types.StringType,
				"url":                   types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["project"]; ok && val != nil {
		if str, ok := val.(string); ok {
			// Normalize URL to UUID
			parts := strings.Split(strings.TrimRight(str, "/"), "/")
			uuid := parts[len(parts)-1]
			data.Project = types.StringValue(uuid)
		} else {
			data.Project = types.StringNull()
		}
	} else {
		if data.Project.IsUnknown() {
			data.Project = types.StringNull()
		}
	}
	if val, ok := sourceMap["project_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ProjectName = types.StringValue(str)
		}
	} else {
		if data.ProjectName.IsUnknown() {
			data.ProjectName = types.StringNull()
		}
	}
	if val, ok := sourceMap["project_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ProjectUuid = types.StringValue(str)
		}
	} else {
		if data.ProjectUuid.IsUnknown() {
			data.ProjectUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["ram"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Ram = types.Int64Value(int64(num))
		}
	} else {
		if data.Ram.IsUnknown() {
			data.Ram = types.Int64Null()
		}
	}
	if val, ok := sourceMap["release_floating_ips"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.ReleaseFloatingIps = types.BoolValue(b)
		}
	} else {
		if data.ReleaseFloatingIps.IsUnknown() {
			data.ReleaseFloatingIps = types.BoolNull()
		}
	}
	if val, ok := sourceMap["resource_type"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ResourceType = types.StringValue(str)
		}
	} else {
		if data.ResourceType.IsUnknown() {
			data.ResourceType = types.StringNull()
		}
	}
	if val, ok := sourceMap["runtime_state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.RuntimeState = types.StringValue(str)
		}
	} else {
		if data.RuntimeState.IsUnknown() {
			data.RuntimeState = types.StringNull()
		}
	}
	if val, ok := sourceMap["security_groups"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"url":         types.StringType,
						"description": types.StringType,
						"name":        types.StringType,
						"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}},
						"state":       types.StringType,
					}
					attrValues := map[string]attr.Value{
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"description": func() attr.Value {
							if v, ok := objMap["description"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"rules": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}}.ElemType),
						"state": func() attr.Value {
							if v, ok := objMap["state"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"url":         types.StringType,
				"description": types.StringType,
				"name":        types.StringType,
				"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}},
				"state":       types.StringType,
			}}, items)
			data.SecurityGroups = listVal
		}
	} else {
		if data.SecurityGroups.IsUnknown() {
			data.SecurityGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"url":         types.StringType,
				"description": types.StringType,
				"name":        types.StringType,
				"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "id": types.Int64Type, "protocol": types.StringType, "remote_group_name": types.StringType, "remote_group_uuid": types.StringType, "to_port": types.Int64Type}}},
				"state":       types.StringType,
			}})
		}
	}
	if val, ok := sourceMap["server_group"]; ok && val != nil {
		// Nested object
		if objMap, ok := val.(map[string]interface{}); ok {
			_ = objMap // Avoid unused variable if properties are empty
			attrTypes := map[string]attr.Type{
				"name":   types.StringType,
				"policy": types.StringType,
				"state":  types.StringType,
				"url":    types.StringType,
			}
			attrValues := map[string]attr.Value{
				"name": func() attr.Value {
					if v, ok := objMap["name"].(string); ok {
						return types.StringValue(v)
					}
					return types.StringNull()
				}(),
				"policy": func() attr.Value {
					if v, ok := objMap["policy"].(string); ok {
						return types.StringValue(v)
					}
					return types.StringNull()
				}(),
				"state": func() attr.Value {
					if v, ok := objMap["state"].(string); ok {
						return types.StringValue(v)
					}
					return types.StringNull()
				}(),
				"url": func() attr.Value {
					if v, ok := objMap["url"].(string); ok {
						return types.StringValue(v)
					}
					return types.StringNull()
				}(),
			}
			objVal, _ := types.ObjectValue(attrTypes, attrValues)
			data.ServerGroup = objVal
		}
	} else {
		if data.ServerGroup.IsUnknown() {
			data.ServerGroup = types.ObjectNull(map[string]attr.Type{
				"name":   types.StringType,
				"policy": types.StringType,
				"state":  types.StringType,
				"url":    types.StringType,
			})
		}
	}
	if val, ok := sourceMap["service_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ServiceName = types.StringValue(str)
		}
	} else {
		if data.ServiceName.IsUnknown() {
			data.ServiceName = types.StringNull()
		}
	}
	if val, ok := sourceMap["service_settings"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ServiceSettings = types.StringValue(str)
		}
	} else {
		if data.ServiceSettings.IsUnknown() {
			data.ServiceSettings = types.StringNull()
		}
	}
	if val, ok := sourceMap["service_settings_error_message"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ServiceSettingsErrorMessage = types.StringValue(str)
		}
	} else {
		if data.ServiceSettingsErrorMessage.IsUnknown() {
			data.ServiceSettingsErrorMessage = types.StringNull()
		}
	}
	if val, ok := sourceMap["service_settings_state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ServiceSettingsState = types.StringValue(str)
		}
	} else {
		if data.ServiceSettingsState.IsUnknown() {
			data.ServiceSettingsState = types.StringNull()
		}
	}
	if val, ok := sourceMap["service_settings_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ServiceSettingsUuid = types.StringValue(str)
		}
	} else {
		if data.ServiceSettingsUuid.IsUnknown() {
			data.ServiceSettingsUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["ssh_public_key"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.SshPublicKey = types.StringValue(str)
		}
	} else {
		if data.SshPublicKey.IsUnknown() {
			data.SshPublicKey = types.StringNull()
		}
	}
	if val, ok := sourceMap["start_time"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.StartTime = types.StringValue(str)
		}
	} else {
		if data.StartTime.IsUnknown() {
			data.StartTime = types.StringNull()
		}
	}
	if val, ok := sourceMap["state"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.State = types.StringValue(str)
		}
	} else {
		if data.State.IsUnknown() {
			data.State = types.StringNull()
		}
	}
	if val, ok := sourceMap["system_volume_size"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.SystemVolumeSize = types.Int64Value(int64(num))
		}
	} else {
		if data.SystemVolumeSize.IsUnknown() {
			data.SystemVolumeSize = types.Int64Null()
		}
	}
	if val, ok := sourceMap["system_volume_type"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.SystemVolumeType = types.StringValue(str)
		}
	} else {
		if data.SystemVolumeType.IsUnknown() {
			data.SystemVolumeType = types.StringNull()
		}
	}
	if val, ok := sourceMap["tenant"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Tenant = types.StringValue(str)
		}
	} else {
		if data.Tenant.IsUnknown() {
			data.Tenant = types.StringNull()
		}
	}
	if val, ok := sourceMap["tenant_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.TenantUuid = types.StringValue(str)
		}
	} else {
		if data.TenantUuid.IsUnknown() {
			data.TenantUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["url"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Url = types.StringValue(str)
		}
	} else {
		if data.Url.IsUnknown() {
			data.Url = types.StringNull()
		}
	}
	if val, ok := sourceMap["user_data"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.UserData = types.StringValue(str)
		}
	} else {
		if data.UserData.IsUnknown() {
			data.UserData = types.StringNull()
		}
	}
	if val, ok := sourceMap["volumes"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
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
					}
					attrValues := map[string]attr.Value{
						"bootable": func() attr.Value {
							if v, ok := objMap["bootable"].(bool); ok {
								return types.BoolValue(v)
							}
							return types.BoolNull()
						}(),
						"device": func() attr.Value {
							if v, ok := objMap["device"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"image_name": func() attr.Value {
							if v, ok := objMap["image_name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"marketplace_resource_uuid": func() attr.Value {
							if v, ok := objMap["marketplace_resource_uuid"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"resource_type": func() attr.Value {
							if v, ok := objMap["resource_type"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"size": func() attr.Value {
							if v, ok := objMap["size"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"state": func() attr.Value {
							if v, ok := objMap["state"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"type": func() attr.Value {
							if v, ok := objMap["type"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"type_name": func() attr.Value {
							if v, ok := objMap["type_name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"url": func() attr.Value {
							if v, ok := objMap["url"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}}, items)
			data.Volumes = listVal
		}
	} else {
		if data.Volumes.IsUnknown() {
			data.Volumes = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
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
			}})
		}
	}

	// Map filter parameters from response if available

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackInstanceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackInstanceResourceModel

	// Read Terraform prior state data into the model
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

	url := fmt.Sprintf("/api/marketplace-resources/%s/terminate/", data.UUID.ValueString())
	var res map[string]interface{}
	err := r.client.Post(ctx, url, payload, &res)
	if err != nil {
		resp.Diagnostics.AddError("Termination Failed", err.Error())
		return
	}

	// Wait for deletion if order UUID is returned
	if orderUUID, ok := res["uuid"].(string); ok {
		stateConf := &retry.StateChangeConf{
			Pending: []string{"pending", "executing", "created"},
			Target:  []string{"done"},
			Refresh: func() (interface{}, string, error) {
				var res map[string]interface{}
				err := r.client.GetByUUID(ctx, "/api/marketplace-orders/{uuid}/", orderUUID, &res)
				if err != nil {
					return nil, "", err
				}
				state, _ := res["state"].(string)
				if state == "erred" || state == "rejected" {
					return res, "failed", fmt.Errorf("termination order failed")
				}
				return res, state, nil
			},
			Timeout: func() time.Duration {
				timeout, diags := data.Timeouts.Delete(ctx, 45*time.Minute)
				resp.Diagnostics.Append(diags...)
				return timeout
			}(),
			Delay:      10 * time.Second,
			MinTimeout: 5 * time.Second,
		}
		_, err := stateConf.WaitForStateContext(ctx)
		if err != nil {
			resp.Diagnostics.AddError("Termination Order Failed", err.Error())
			return
		}
	}
}

func (r *OpenstackInstanceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
