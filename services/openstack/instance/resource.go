package instance

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

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
	client *OpenstackInstanceClient
}

// OpenstackInstanceResourceModel describes the resource data model.
type OpenstackInstanceResourceModel struct {
	OpenstackInstanceModel
	DataVolumeSize     types.Int64    `tfsdk:"data_volume_size"`
	DataVolumeType     types.String   `tfsdk:"data_volume_type"`
	DataVolumes        types.List     `tfsdk:"data_volumes"`
	DeleteVolumes      types.Bool     `tfsdk:"delete_volumes"`
	Flavor             types.String   `tfsdk:"flavor"`
	Image              types.String   `tfsdk:"image"`
	Limits             types.Map      `tfsdk:"limits"`
	Offering           types.String   `tfsdk:"offering"`
	Plan               types.String   `tfsdk:"plan"`
	ReleaseFloatingIps types.Bool     `tfsdk:"release_floating_ips"`
	SshPublicKey       types.String   `tfsdk:"ssh_public_key"`
	SystemVolumeSize   types.Int64    `tfsdk:"system_volume_size"`
	SystemVolumeType   types.String   `tfsdk:"system_volume_type"`
	Timeouts           timeouts.Value `tfsdk:"timeouts"`
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
				MarkdownDescription: "Openstack Instance UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
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
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
					boolplanmodifier.UseStateForUnknown(),
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
			"customer": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Customer",
			},
			"data_volume_size": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Size of the data volume in MiB. Minimum size is 1024 MiB (1 GiB)",
				Validators: []validator.Int64{
					int64validator.AtLeast(1024),
				},
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
							MarkdownDescription: "Volume Type",
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
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Termination attribute",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Description",
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
				MarkdownDescription: "Error Message",
			},
			"error_traceback": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Error Traceback",
			},
			"external_address": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "External Address",
			},
			"external_ips": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "External Ips",
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
			"floating_ips": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Existing floating IP address in selected OpenStack tenant to be assigned to new virtual machine",
						},
						"subnet": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Subnet",
						},
						"url": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Url",
						},
						"address": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "The public IPv4 address of the floating IP",
						},
						"port_fixed_ips": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ip_address": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "IP address to assign to the port",
									},
									"subnet_id": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "ID of the subnet in which to assign the IP address",
									},
								},
							},
							Computed: true,
							PlanModifiers: []planmodifier.List{
								listplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Port Fixed Ips",
						},
						"port_mac_address": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "MAC address of the port",
						},
						"subnet_cidr": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "IPv4 network address in CIDR format (e.g. 192.168.0.0/24)",
						},
						"subnet_description": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Subnet Description",
						},
						"subnet_name": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Subnet Name",
						},
						"subnet_uuid": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Subnet Uuid",
						},
						"uuid": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Uuid",
						},
					},
				},
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
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
				MarkdownDescription: "Image Name",
			},
			"internal_ips": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Internal Ips",
			},
			"key_fingerprint": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Key Fingerprint",
			},
			"key_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Key Name",
			},
			"latitude": schema.Float64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Latitude",
			},
			"limits": schema.MapAttribute{
				ElementType: types.Float64Type,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Map{
					mapplanmodifier.RequiresReplace(),
					mapplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Resource limits",
			},
			"longitude": schema.Float64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Longitude",
			},
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Marketplace Resource Uuid",
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
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name",
			},
			"offering": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Offering URL",
			},
			"plan": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Plan URL",
			},
			"ports": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"fixed_ips": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ip_address": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "IP address to assign to the port",
									},
									"subnet_id": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "ID of the subnet in which to assign the IP address",
									},
								},
							},
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.List{
								listplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Fixed Ips",
						},
						"port": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Port",
						},
						"subnet": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Subnet to which this port belongs",
						},
						"allowed_address_pairs": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"mac_address": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Mac Address",
									},
								},
							},
							Computed: true,
							PlanModifiers: []planmodifier.List{
								listplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Allowed Address Pairs",
						},
						"device_id": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "ID of device (instance, router etc) to which this port is connected",
						},
						"device_owner": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Entity that uses this port (e.g. network:router_interface)",
						},
						"mac_address": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "MAC address of the port",
						},
						"security_groups": schema.SetNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"backend_id": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Backend Id",
									},
									"customer": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Customer",
									},
									"description": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Description",
									},
									"error_message": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Error Message",
									},
									"error_traceback": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Error Traceback",
									},
									"marketplace_resource_uuid": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Marketplace Resource Uuid",
									},
									"name": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Name",
									},
									"project": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Project",
									},
									"resource_type": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Resource Type",
									},
									"rules": schema.ListNestedAttribute{
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"cidr": schema.StringAttribute{
													Optional: true,
													Computed: true,
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "CIDR notation for the source/destination network address range",
												},
												"description": schema.StringAttribute{
													Optional: true,
													Computed: true,
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "Description",
												},
												"direction": schema.StringAttribute{
													Optional: true,
													Computed: true,
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "Traffic direction - either 'ingress' (incoming) or 'egress' (outgoing)",
												},
												"ethertype": schema.StringAttribute{
													Optional: true,
													Computed: true,
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "IP protocol version - either 'IPv4' or 'IPv6'",
												},
												"from_port": schema.Int64Attribute{
													Optional: true,
													Computed: true,
													PlanModifiers: []planmodifier.Int64{
														int64planmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "Starting port number in the range (1-65535)",
													Validators: []validator.Int64{
														int64validator.AtLeast(-2147483648),
														int64validator.AtMost(65535),
													},
												},
												"id": schema.Int64Attribute{
													Computed: true,
													PlanModifiers: []planmodifier.Int64{
														int64planmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "Id",
												},
												"protocol": schema.StringAttribute{
													Optional: true,
													Computed: true,
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "The network protocol (TCP, UDP, ICMP, or empty for any protocol)",
												},
												"remote_group": schema.StringAttribute{
													Optional: true,
													Computed: true,
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "Remote security group that this rule references, if any",
												},
												"remote_group_name": schema.StringAttribute{
													Computed: true,
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "Remote Group Name",
												},
												"remote_group_uuid": schema.StringAttribute{
													Computed: true,
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "Remote Group Uuid",
												},
												"to_port": schema.Int64Attribute{
													Optional: true,
													Computed: true,
													PlanModifiers: []planmodifier.Int64{
														int64planmodifier.UseStateForUnknown(),
													},
													MarkdownDescription: "Ending port number in the range (1-65535)",
													Validators: []validator.Int64{
														int64validator.AtLeast(-2147483648),
														int64validator.AtMost(65535),
													},
												},
											},
										},
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.List{
											listplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Rules",
									},
									"state": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "State",
									},
									"tenant": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Tenant",
									},
									"tenant_name": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Tenant Name",
									},
									"tenant_uuid": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Tenant Uuid",
									},
									"url": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Url",
									},
									"uuid": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Uuid",
									},
								},
							},
							Computed: true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Security Groups",
						},
						"subnet_cidr": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "IPv4 network address in CIDR format (e.g. 192.168.0.0/24)",
						},
						"subnet_description": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Subnet Description",
						},
						"subnet_name": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Subnet Name",
						},
						"subnet_uuid": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Subnet Uuid",
						},
						"url": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Url",
						},
					},
				},
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Network ports to attach to the instance",
			},
			"project": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Project URL",
			},
			"ram": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Memory size in MiB",
			},
			"rancher_cluster": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"marketplace_uuid": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Marketplace Uuid",
					},
					"name": schema.StringAttribute{
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name",
					},
					"uuid": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Uuid",
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Rancher Cluster",
			},
			"release_floating_ips": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Termination attribute",
			},
			"resource_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Resource Type",
			},
			"runtime_state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Runtime State",
			},
			"security_groups": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"url": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Url",
						},
						"description": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Description",
						},
						"name": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Name",
						},
						"rules": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"cidr": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "CIDR notation for the source/destination network address range",
									},
									"description": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Description",
									},
									"direction": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Traffic direction - either 'ingress' (incoming) or 'egress' (outgoing)",
									},
									"ethertype": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "IP protocol version - either 'IPv4' or 'IPv6'",
									},
									"from_port": schema.Int64Attribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Starting port number in the range (1-65535)",
										Validators: []validator.Int64{
											int64validator.AtLeast(-2147483648),
											int64validator.AtMost(65535),
										},
									},
									"id": schema.Int64Attribute{
										Computed: true,
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Id",
									},
									"protocol": schema.StringAttribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "The network protocol (TCP, UDP, ICMP, or empty for any protocol)",
									},
									"remote_group_name": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Remote Group Name",
									},
									"remote_group_uuid": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Remote Group Uuid",
									},
									"to_port": schema.Int64Attribute{
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: "Ending port number in the range (1-65535)",
										Validators: []validator.Int64{
											int64validator.AtLeast(-2147483648),
											int64validator.AtMost(65535),
										},
									},
								},
							},
							Computed: true,
							PlanModifiers: []planmodifier.List{
								listplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Rules",
						},
						"state": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "State",
						},
					},
				},
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "List of security groups to apply to the instance",
			},
			"server_group": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"url": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Url",
					},
					"name": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Name",
					},
					"policy": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "Server group policy determining the rules for scheduling servers in this group",
					},
					"state": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						MarkdownDescription: "State",
					},
				},
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplace(),
					objectplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Server group for instance scheduling policy",
			},
			"ssh_public_key": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Ssh Public Key",
			},
			"start_time": schema.StringAttribute{
				CustomType: timetypes.RFC3339Type{},
				Computed:   true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Start Time",
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
				Validators: []validator.Int64{
					int64validator.AtLeast(1024),
				},
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
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Additional data that will be added to instance on provisioning",
			},
			"volumes": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"bootable": schema.BoolAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Indicates if this volume can be used to boot an instance",
						},
						"device": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Name of volume as instance device e.g. /dev/vdb.",
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^/dev/[a-zA-Z0-9]+$`), ""),
							},
						},
						"image_name": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Name of the image this volume was created from",
						},
						"marketplace_resource_uuid": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Marketplace Resource Uuid",
						},
						"name": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Name",
						},
						"resource_type": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Resource Type",
						},
						"size": schema.Int64Attribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Size in MiB",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							},
						},
						"state": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "State",
						},
						"type": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Type of the volume (e.g. SSD, HDD)",
						},
						"type_name": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Type Name",
						},
						"url": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Url",
						},
						"uuid": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: "Uuid",
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

	r.client = &OpenstackInstanceClient{}
	if err := r.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			err.Error(),
		)
		return
	}
}

// resolveUnknownAttributes ensures that fields not returned by the Waldur GET API
// are set to explicit null values instead of remaining "Unknown".
func (r *OpenstackInstanceResource) resolveUnknownAttributes(data *OpenstackInstanceResourceModel) {
	// Iterate over all model fields to handle Unknown values
	if data.Action.IsUnknown() {
		data.Action = types.StringNull()
	}
	if data.AvailabilityZone.IsUnknown() {
		data.AvailabilityZone = types.StringNull()
	}
	if data.AvailabilityZoneName.IsUnknown() {
		data.AvailabilityZoneName = types.StringNull()
	}
	if data.BackendId.IsUnknown() {
		data.BackendId = types.StringNull()
	}
	if data.ConnectDirectlyToExternalNetwork.IsUnknown() {
		data.ConnectDirectlyToExternalNetwork = types.BoolNull()
	}
	if data.Cores.IsUnknown() {
		data.Cores = types.Int64Null()
	}
	if data.Customer.IsUnknown() {
		data.Customer = types.StringNull()
	}
	if data.DataVolumeSize.IsUnknown() {
		data.DataVolumeSize = types.Int64Null()
	}
	if data.DataVolumeType.IsUnknown() {
		data.DataVolumeType = types.StringNull()
	}
	if data.DataVolumes.IsUnknown() {
		data.DataVolumes = types.ListNull(OpenStackDataVolumeRequestType())
	}
	if data.DeleteVolumes.IsUnknown() {
		data.DeleteVolumes = types.BoolNull()
	}
	if data.Description.IsUnknown() {
		data.Description = types.StringNull()
	}
	if data.Disk.IsUnknown() {
		data.Disk = types.Int64Null()
	}
	if data.ErrorMessage.IsUnknown() {
		data.ErrorMessage = types.StringNull()
	}
	if data.ErrorTraceback.IsUnknown() {
		data.ErrorTraceback = types.StringNull()
	}
	if data.ExternalAddress.IsUnknown() {
		data.ExternalAddress = types.ListNull(types.StringType)
	}
	if data.ExternalIps.IsUnknown() {
		data.ExternalIps = types.ListNull(types.StringType)
	}
	if data.Flavor.IsUnknown() {
		data.Flavor = types.StringNull()
	}
	if data.FlavorDisk.IsUnknown() {
		data.FlavorDisk = types.Int64Null()
	}
	if data.FlavorName.IsUnknown() {
		data.FlavorName = types.StringNull()
	}
	if data.FloatingIps.IsUnknown() {
		data.FloatingIps = types.SetNull(OpenStackCreateFloatingIPRequestType())
	}
	if data.HypervisorHostname.IsUnknown() {
		data.HypervisorHostname = types.StringNull()
	}
	if data.Image.IsUnknown() {
		data.Image = types.StringNull()
	}
	if data.ImageName.IsUnknown() {
		data.ImageName = types.StringNull()
	}
	if data.InternalIps.IsUnknown() {
		data.InternalIps = types.ListNull(types.StringType)
	}
	if data.KeyFingerprint.IsUnknown() {
		data.KeyFingerprint = types.StringNull()
	}
	if data.KeyName.IsUnknown() {
		data.KeyName = types.StringNull()
	}
	if data.Latitude.IsUnknown() {
		data.Latitude = types.Float64Null()
	}
	if data.Limits.IsUnknown() {
		data.Limits = types.MapNull(types.Float64Type)
	}
	if data.Longitude.IsUnknown() {
		data.Longitude = types.Float64Null()
	}
	if data.MarketplaceResourceUuid.IsUnknown() {
		data.MarketplaceResourceUuid = types.StringNull()
	}
	if data.MinDisk.IsUnknown() {
		data.MinDisk = types.Int64Null()
	}
	if data.MinRam.IsUnknown() {
		data.MinRam = types.Int64Null()
	}
	if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}
	if data.Offering.IsUnknown() {
		data.Offering = types.StringNull()
	}
	if data.Plan.IsUnknown() {
		data.Plan = types.StringNull()
	}
	if data.Ports.IsUnknown() {
		data.Ports = types.ListNull(OpenStackCreateInstancePortRequestType())
	}
	if data.Project.IsUnknown() {
		data.Project = types.StringNull()
	}
	if data.Ram.IsUnknown() {
		data.Ram = types.Int64Null()
	}
	if data.RancherCluster.IsUnknown() {
		data.RancherCluster = types.ObjectNull(RancherClusterType().AttrTypes)
	}
	if data.ReleaseFloatingIps.IsUnknown() {
		data.ReleaseFloatingIps = types.BoolNull()
	}
	if data.ResourceType.IsUnknown() {
		data.ResourceType = types.StringNull()
	}
	if data.RuntimeState.IsUnknown() {
		data.RuntimeState = types.StringNull()
	}
	if data.SecurityGroups.IsUnknown() {
		data.SecurityGroups = types.SetNull(OpenStackSecurityGroupHyperlinkRequestType())
	}
	if data.ServerGroup.IsUnknown() {
		data.ServerGroup = types.ObjectNull(ServerGroupType().AttrTypes)
	}
	if data.SshPublicKey.IsUnknown() {
		data.SshPublicKey = types.StringNull()
	}
	if data.StartTime.IsUnknown() {
		data.StartTime = timetypes.NewRFC3339Null()
	}
	if data.State.IsUnknown() {
		data.State = types.StringNull()
	}
	if data.SystemVolumeSize.IsUnknown() {
		data.SystemVolumeSize = types.Int64Null()
	}
	if data.SystemVolumeType.IsUnknown() {
		data.SystemVolumeType = types.StringNull()
	}
	if data.Tenant.IsUnknown() {
		data.Tenant = types.StringNull()
	}
	if data.TenantUuid.IsUnknown() {
		data.TenantUuid = types.StringNull()
	}
	if data.Url.IsUnknown() {
		data.Url = types.StringNull()
	}
	if data.UserData.IsUnknown() {
		data.UserData = types.StringNull()
	}
	if data.Volumes.IsUnknown() {
		data.Volumes = types.ListNull(OpenStackNestedVolumeType())
	}
}

func (r *OpenstackInstanceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var data OpenstackInstanceResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Phase 1: Payload Construction
	// We map the Terraform schema fields to the 'attributes' map required by the Marketplace Order API.
	attributes := OpenstackInstanceCreateAttributes{}
	if !data.AvailabilityZone.IsNull() && !data.AvailabilityZone.IsUnknown() {
		attributes.AvailabilityZone = data.AvailabilityZone.ValueStringPointer()
	}
	if !data.ConnectDirectlyToExternalNetwork.IsNull() && !data.ConnectDirectlyToExternalNetwork.IsUnknown() {
		attributes.ConnectDirectlyToExternalNetwork = data.ConnectDirectlyToExternalNetwork.ValueBoolPointer()
	}
	if !data.DataVolumeSize.IsNull() && !data.DataVolumeSize.IsUnknown() {
		attributes.DataVolumeSize = data.DataVolumeSize.ValueInt64Pointer()
	}
	if !data.DataVolumeType.IsNull() && !data.DataVolumeType.IsUnknown() {
		attributes.DataVolumeType = data.DataVolumeType.ValueStringPointer()
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		attributes.Description = data.Description.ValueStringPointer()
	}
	if !data.Flavor.IsNull() && !data.Flavor.IsUnknown() {
		attributes.Flavor = data.Flavor.ValueStringPointer()
	}
	if !data.Image.IsNull() && !data.Image.IsUnknown() {
		attributes.Image = data.Image.ValueStringPointer()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		attributes.Name = data.Name.ValueStringPointer()
	}
	if !data.SshPublicKey.IsNull() && !data.SshPublicKey.IsUnknown() {
		attributes.SshPublicKey = data.SshPublicKey.ValueStringPointer()
	}
	if !data.SystemVolumeSize.IsNull() && !data.SystemVolumeSize.IsUnknown() {
		attributes.SystemVolumeSize = data.SystemVolumeSize.ValueInt64Pointer()
	}
	if !data.SystemVolumeType.IsNull() && !data.SystemVolumeType.IsUnknown() {
		attributes.SystemVolumeType = data.SystemVolumeType.ValueStringPointer()
	}
	if !data.UserData.IsNull() && !data.UserData.IsUnknown() {
		attributes.UserData = data.UserData.ValueStringPointer()
	}
	resp.Diagnostics.Append(common.PopulateOptionalSliceField(ctx, data.DataVolumes, &attributes.DataVolumes)...)
	resp.Diagnostics.Append(common.PopulateOptionalSetField(ctx, data.FloatingIps, &attributes.FloatingIps)...)
	resp.Diagnostics.Append(common.PopulateOptionalSliceField(ctx, data.Ports, &attributes.Ports)...)
	resp.Diagnostics.Append(common.PopulateOptionalSetField(ctx, data.SecurityGroups, &attributes.SecurityGroups)...)
	resp.Diagnostics.Append(common.PopulateOptionalObjectField(ctx, data.ServerGroup, &attributes.ServerGroup)...)

	// Construct the Create Order Request
	payload := OpenstackInstanceCreateRequest{
		Project:    data.Project.ValueStringPointer(),
		Offering:   data.Offering.ValueStringPointer(),
		Attributes: attributes,
	}

	if !data.Plan.IsNull() && !data.Plan.IsUnknown() {
		payload.Plan = data.Plan.ValueStringPointer()
	}

	if !data.Limits.IsNull() && !data.Limits.IsUnknown() {
		limits := make(map[string]float64)
		diags := data.Limits.ElementsAs(ctx, &limits, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		payload.Limits = limits
	}

	// Phase 2: Submit Order
	orderRes, err := r.client.CreateOrder(ctx, &payload)
	if err != nil {
		resp.Diagnostics.AddError("Order Submission Failed", err.Error())
		return
	}

	// Phase 3: Poll for Completion
	// We use the 'time' package to handle the timeout specified in the TF config or default to global default.
	timeout, diags := data.Timeouts.Create(ctx, common.DefaultCreateTimeout)
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
	apiResp, err := r.client.Get(ctx, data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to Read Resource", err.Error())
		return
	}
	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	// Resolve unknown attributes to explicit null values
	r.resolveUnknownAttributes(&data)

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

	apiResp, err := r.client.Get(ctx, data.UUID.ValueString())
	if err != nil {
		if IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Unable to Read Openstack Instance",
			"An error occurred while reading the Openstack Instance: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Resolve unknown attributes to explicit null values
	r.resolveUnknownAttributes(&data)

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
	anyChanges := false
	var patchPayload OpenstackInstanceUpdateRequest
	if !data.Description.IsNull() && !data.Description.Equal(state.Description) {
		anyChanges = true
		patchPayload.Description = data.Description.ValueStringPointer()
	}
	if !data.Name.IsNull() && !data.Name.Equal(state.Name) {
		anyChanges = true
		patchPayload.Name = data.Name.ValueStringPointer()
	}

	if anyChanges {
		// Execute the PATCH request
		_, err := r.client.Update(ctx, data.UUID.ValueString(), &patchPayload)
		if err != nil {
			resp.Diagnostics.AddError("Update Failed", err.Error())
			return
		}
	}

	updateTimeout, diags := data.Timeouts.Update(ctx, common.DefaultUpdateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Phase 2: RPC Actions
	// These actions are triggered when their corresponding specific fields change.
	if !data.FloatingIps.Equal(state.FloatingIps) {
		// Convert Terraform value to API payload for the specific action
		var req OpenstackInstanceUpdateFloatingIpsActionRequest
		resp.Diagnostics.Append(common.PopulateSetField(ctx, data.FloatingIps, &req.FloatingIps)...)

		// Execute the Action
		if err := r.client.UpdateFloatingIps(ctx, data.UUID.ValueString(), &req); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: update_floating_ips", err.Error())
			return
		}

		// Wait for the resource to return to OK state
		apiResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackInstanceResponse, error) {
			return r.client.Get(ctx, data.UUID.ValueString())
		}, updateTimeout)
		if err != nil {
			resp.Diagnostics.AddError("Wait for RPC action failed", err.Error())
			return
		}
		resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
		state = data // Update local state to avoid repeated action calls if multiple fields changed (though actions are usually 1-to-1)
	}
	if !data.Ports.Equal(state.Ports) {
		// Convert Terraform value to API payload for the specific action
		var req OpenstackInstanceUpdatePortsActionRequest
		resp.Diagnostics.Append(common.PopulateSliceField(ctx, data.Ports, &req.Ports)...)

		// Execute the Action
		if err := r.client.UpdatePorts(ctx, data.UUID.ValueString(), &req); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: update_ports", err.Error())
			return
		}

		// Wait for the resource to return to OK state
		apiResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackInstanceResponse, error) {
			return r.client.Get(ctx, data.UUID.ValueString())
		}, updateTimeout)
		if err != nil {
			resp.Diagnostics.AddError("Wait for RPC action failed", err.Error())
			return
		}
		resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
		state = data // Update local state to avoid repeated action calls if multiple fields changed (though actions are usually 1-to-1)
	}
	if !data.SecurityGroups.Equal(state.SecurityGroups) {
		// Convert Terraform value to API payload for the specific action
		var req OpenstackInstanceUpdateSecurityGroupsActionRequest
		resp.Diagnostics.Append(common.PopulateSetField(ctx, data.SecurityGroups, &req.SecurityGroups)...)

		// Execute the Action
		if err := r.client.UpdateSecurityGroups(ctx, data.UUID.ValueString(), &req); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: update_security_groups", err.Error())
			return
		}

		// Wait for the resource to return to OK state
		apiResp, err := common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackInstanceResponse, error) {
			return r.client.Get(ctx, data.UUID.ValueString())
		}, updateTimeout)
		if err != nil {
			resp.Diagnostics.AddError("Wait for RPC action failed", err.Error())
			return
		}
		resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
		state = data // Update local state to avoid repeated action calls if multiple fields changed (though actions are usually 1-to-1)
	}

	// Fetch updated state after all changes
	apiResp, err := r.client.Get(ctx, data.UUID.ValueString())
	if err != nil {
		if IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Failed to Read Resource After Update", err.Error())
		return
	}
	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)

	// Resolve unknown attributes to explicit null values
	r.resolveUnknownAttributes(&data)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackInstanceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var data OpenstackInstanceResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Order-based Delete
	// OpenStack instances must be stopped before they can be terminated.
	// We check the runtime_state and stop it if it's not already SHUTOFF.
	currData, err := r.client.Get(ctx, data.UUID.ValueString())
	if err == nil && currData.RuntimeState != nil && *currData.RuntimeState == "ACTIVE" {
		tflog.Info(ctx, "Stopping Openstack instance before deletion", map[string]interface{}{
			"uuid": data.UUID.ValueString(),
		})
		// We ignore the initial stop error; if it fails, Terminate might still fail with 409
		// which is better than failing early if Stop is temporarily unavailable.
		_ = r.client.Stop(ctx, data.UUID.ValueString())

		// Wait for the instance to reach a stable OK state after stop.
		// Waldur will move it to OK with RuntimeState=SHUTOFF.
		timeout, _ := data.Timeouts.Delete(ctx, common.DefaultDeleteTimeout)
		_, _ = common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackInstanceResponse, error) {
			return r.client.Get(ctx, data.UUID.ValueString())
		}, timeout)
	}

	payload := map[string]interface{}{}
	if !data.DeleteVolumes.IsNull() {
		payload["delete_volumes"] = data.DeleteVolumes.ValueBool()
	}
	if !data.ReleaseFloatingIps.IsNull() {
		payload["release_floating_ips"] = data.ReleaseFloatingIps.ValueBool()
	}

	// Submit termination order
	resourceID := data.UUID.ValueString()
	if !data.MarketplaceResourceUuid.IsNull() && !data.MarketplaceResourceUuid.IsUnknown() {
		resourceID = data.MarketplaceResourceUuid.ValueString()
	}
	orderUUID, err := r.client.Terminate(ctx, resourceID, payload)
	if err != nil {
		resp.Diagnostics.AddError("Termination Failed", err.Error())
		return
	}

	// Wait for deletion if order UUID is returned
	if orderUUID != "" {
		timeout, diags := data.Timeouts.Delete(ctx, common.DefaultDeleteTimeout)
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

	uuid := req.ID
	if uuid == "" {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			"Import ID cannot be empty. Please provide the UUID of the Openstack Instance.",
		)
		return
	}

	tflog.Info(ctx, "Importing Openstack Instance", map[string]interface{}{
		"uuid": uuid,
	})

	apiResp, err := r.client.Get(ctx, uuid)
	if err != nil {
		if IsNotFoundError(err) {
			resp.Diagnostics.AddError(
				"Resource Not Found",
				fmt.Sprintf("Openstack Instance with UUID '%s' does not exist or is not accessible.", uuid),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to Import Openstack Instance",
			fmt.Sprintf("An error occurred while fetching the Openstack Instance: %s", err.Error()),
		)
		return
	}

	var data OpenstackInstanceResourceModel
	resp.Diagnostics.Append(data.CopyFrom(ctx, *apiResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
