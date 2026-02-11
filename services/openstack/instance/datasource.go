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

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OpenstackInstanceDataSource{}

func NewOpenstackInstanceDataSource() datasource.DataSource {
	return &OpenstackInstanceDataSource{}
}

type OpenstackInstanceDataSource struct {
	client *OpenstackInstanceClient
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
				MarkdownDescription: "Openstack Instance UUID",
			},
			"filters": (&OpenstackInstanceFiltersModel{}).GetSchema(),
			"action": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Action"},
			"availability_zone": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Availability zone where this instance is located"},
			"availability_zone_name": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Name of the availability zone where instance is located"},
			"backend_id": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Instance ID in the OpenStack backend"},
			"connect_directly_to_external_network": schema.BoolAttribute{
				Computed: true, MarkdownDescription: "If True, instance will be connected directly to external network"},
			"cores": schema.Int64Attribute{
				Computed: true, MarkdownDescription: "Number of cores in a VM"},
			"customer": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Customer"},
			"description": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Description"},
			"disk": schema.Int64Attribute{
				Computed: true, MarkdownDescription: "Disk size in MiB"},
			"error_message": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Error Message"},
			"external_address": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true, MarkdownDescription: "External Address"},
			"external_ips": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true, MarkdownDescription: "External Ips"},
			"flavor_disk": schema.Int64Attribute{
				Computed: true, MarkdownDescription: "Flavor disk size in MiB"},
			"flavor_name": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Name of the flavor used by this instance"},
			"floating_ips": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Existing floating IP address in selected OpenStack tenant to be assigned to new virtual machine"},
						"subnet": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Subnet"},
						"url": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Url"},
						"address": schema.StringAttribute{
							Computed: true, MarkdownDescription: "The public IPv4 address of the floating IP"},
						"port_fixed_ips": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ip_address": schema.StringAttribute{
										Computed: true, MarkdownDescription: "IP address to assign to the port"},
									"subnet_id": schema.StringAttribute{
										Computed: true, MarkdownDescription: "ID of the subnet in which to assign the IP address"},
								},
							},
							Computed: true, MarkdownDescription: "Port Fixed Ips",
						},
						"port_mac_address": schema.StringAttribute{
							Computed: true, MarkdownDescription: "MAC address of the port"},
						"subnet_cidr": schema.StringAttribute{
							Computed: true, MarkdownDescription: "IPv4 network address in CIDR format (e.g. 192.168.0.0/24)"},
						"subnet_description": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Subnet Description"},
						"subnet_name": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Subnet Name"},
						"subnet_uuid": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Subnet Uuid"},
						"uuid": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Uuid"},
					},
				},
				Computed: true, MarkdownDescription: "Floating IPs to assign to the instance",
			},
			"hypervisor_hostname": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Name of the hypervisor hosting this instance"},
			"image_name": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Image Name"},
			"internal_ips": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true, MarkdownDescription: "Internal Ips"},
			"key_fingerprint": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Key Fingerprint"},
			"key_name": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Key Name"},
			"latitude": schema.Float64Attribute{
				Computed: true, MarkdownDescription: "Latitude"},
			"longitude": schema.Float64Attribute{
				Computed: true, MarkdownDescription: "Longitude"},
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Marketplace Resource Uuid"},
			"min_disk": schema.Int64Attribute{
				Computed: true, MarkdownDescription: "Minimum disk size in MiB"},
			"min_ram": schema.Int64Attribute{
				Computed: true, MarkdownDescription: "Minimum memory size in MiB"},
			"name": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Name"},
			"ports": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"fixed_ips": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ip_address": schema.StringAttribute{
										Computed: true, MarkdownDescription: "IP address to assign to the port"},
									"subnet_id": schema.StringAttribute{
										Computed: true, MarkdownDescription: "ID of the subnet in which to assign the IP address"},
								},
							},
							Computed: true, MarkdownDescription: "Fixed Ips",
						},
						"port": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Port"},
						"subnet": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Subnet to which this port belongs"},
						"allowed_address_pairs": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"mac_address": schema.StringAttribute{
										Computed: true, MarkdownDescription: "Mac Address"},
								},
							},
							Computed: true, MarkdownDescription: "Allowed Address Pairs",
						},
						"device_id": schema.StringAttribute{
							Computed: true, MarkdownDescription: "ID of device (instance, router etc) to which this port is connected"},
						"device_owner": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Entity that uses this port (e.g. network:router_interface)"},
						"mac_address": schema.StringAttribute{
							Computed: true, MarkdownDescription: "MAC address of the port"},
						"security_groups": schema.SetNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"backend_id": schema.StringAttribute{
										Computed: true, MarkdownDescription: "Backend Id"},
									"customer": schema.StringAttribute{
										Computed: true, MarkdownDescription: "Customer"},
									"description": schema.StringAttribute{
										Computed: true, MarkdownDescription: "Description"},
									"error_message": schema.StringAttribute{
										Computed: true, MarkdownDescription: "Error Message"},
									"marketplace_resource_uuid": schema.StringAttribute{
										Computed: true, MarkdownDescription: "Marketplace Resource Uuid"},
									"name": schema.StringAttribute{
										Computed: true, MarkdownDescription: "Name"},
									"project": schema.StringAttribute{
										Computed: true, MarkdownDescription: "Project"},
									"resource_type": schema.StringAttribute{
										Computed: true, MarkdownDescription: "Resource Type"},
									"rules": schema.ListNestedAttribute{
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"cidr": schema.StringAttribute{
													Computed: true, MarkdownDescription: "CIDR notation for the source/destination network address range"},
												"description": schema.StringAttribute{
													Computed: true, MarkdownDescription: "Description"},
												"direction": schema.StringAttribute{
													Computed: true, MarkdownDescription: "Traffic direction - either 'ingress' (incoming) or 'egress' (outgoing)"},
												"ethertype": schema.StringAttribute{
													Computed: true, MarkdownDescription: "IP protocol version - either 'IPv4' or 'IPv6'"},
												"from_port": schema.Int64Attribute{
													Computed: true, MarkdownDescription: "Starting port number in the range (1-65535)",
													Validators: []validator.Int64{
														int64validator.AtLeast(-2147483648),
														int64validator.AtMost(65535),
													}},
												"id": schema.Int64Attribute{
													Computed: true, MarkdownDescription: "Id"},
												"protocol": schema.StringAttribute{
													Computed: true, MarkdownDescription: "The network protocol (TCP, UDP, ICMP, or empty for any protocol)"},
												"remote_group": schema.StringAttribute{
													Computed: true, MarkdownDescription: "Remote security group that this rule references, if any"},
												"remote_group_name": schema.StringAttribute{
													Computed: true, MarkdownDescription: "Remote Group Name"},
												"remote_group_uuid": schema.StringAttribute{
													Computed: true, MarkdownDescription: "Remote Group Uuid"},
												"to_port": schema.Int64Attribute{
													Computed: true, MarkdownDescription: "Ending port number in the range (1-65535)",
													Validators: []validator.Int64{
														int64validator.AtLeast(-2147483648),
														int64validator.AtMost(65535),
													}},
											},
										},
										Computed: true, MarkdownDescription: "Rules",
									},
									"state": schema.StringAttribute{
										Computed: true, MarkdownDescription: "State"},
									"tenant": schema.StringAttribute{
										Computed: true, MarkdownDescription: "Tenant"},
									"tenant_name": schema.StringAttribute{
										Computed: true, MarkdownDescription: "Tenant Name"},
									"tenant_uuid": schema.StringAttribute{
										Computed: true, MarkdownDescription: "Tenant Uuid"},
									"url": schema.StringAttribute{
										Computed: true, MarkdownDescription: "Url"},
									"uuid": schema.StringAttribute{
										Computed: true, MarkdownDescription: "Uuid"},
								},
							},
							Computed: true, MarkdownDescription: "Security Groups",
						},
						"subnet_cidr": schema.StringAttribute{
							Computed: true, MarkdownDescription: "IPv4 network address in CIDR format (e.g. 192.168.0.0/24)"},
						"subnet_description": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Subnet Description"},
						"subnet_name": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Subnet Name"},
						"subnet_uuid": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Subnet Uuid"},
						"url": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Url"},
					},
				},
				Computed: true, MarkdownDescription: "Network ports to attach to the instance",
			},
			"project": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Project URL"},
			"ram": schema.Int64Attribute{
				Computed: true, MarkdownDescription: "Memory size in MiB"},
			"rancher_cluster": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"marketplace_uuid": schema.StringAttribute{
						Computed: true, MarkdownDescription: "Marketplace Uuid"},
					"name": schema.StringAttribute{
						Computed: true, MarkdownDescription: "Name"},
					"uuid": schema.StringAttribute{
						Computed: true, MarkdownDescription: "Uuid"},
				},
				Computed: true, MarkdownDescription: "Rancher Cluster",
			},
			"resource_type": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Resource Type"},
			"runtime_state": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Runtime State"},
			"security_groups": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"url": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Url"},
						"description": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Description"},
						"name": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Name"},
						"rules": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"cidr": schema.StringAttribute{
										Computed: true, MarkdownDescription: "CIDR notation for the source/destination network address range"},
									"description": schema.StringAttribute{
										Computed: true, MarkdownDescription: "Description"},
									"direction": schema.StringAttribute{
										Computed: true, MarkdownDescription: "Traffic direction - either 'ingress' (incoming) or 'egress' (outgoing)"},
									"ethertype": schema.StringAttribute{
										Computed: true, MarkdownDescription: "IP protocol version - either 'IPv4' or 'IPv6'"},
									"from_port": schema.Int64Attribute{
										Computed: true, MarkdownDescription: "Starting port number in the range (1-65535)",
										Validators: []validator.Int64{
											int64validator.AtLeast(-2147483648),
											int64validator.AtMost(65535),
										}},
									"id": schema.Int64Attribute{
										Computed: true, MarkdownDescription: "Id"},
									"protocol": schema.StringAttribute{
										Computed: true, MarkdownDescription: "The network protocol (TCP, UDP, ICMP, or empty for any protocol)"},
									"remote_group_name": schema.StringAttribute{
										Computed: true, MarkdownDescription: "Remote Group Name"},
									"remote_group_uuid": schema.StringAttribute{
										Computed: true, MarkdownDescription: "Remote Group Uuid"},
									"to_port": schema.Int64Attribute{
										Computed: true, MarkdownDescription: "Ending port number in the range (1-65535)",
										Validators: []validator.Int64{
											int64validator.AtLeast(-2147483648),
											int64validator.AtMost(65535),
										}},
								},
							},
							Computed: true, MarkdownDescription: "Rules",
						},
						"state": schema.StringAttribute{
							Computed: true, MarkdownDescription: "State"},
					},
				},
				Computed: true, MarkdownDescription: "List of security groups to apply to the instance",
			},
			"server_group": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"url": schema.StringAttribute{
						Computed: true, MarkdownDescription: "Url"},
					"name": schema.StringAttribute{
						Computed: true, MarkdownDescription: "Name"},
					"policy": schema.StringAttribute{
						Computed: true, MarkdownDescription: "Server group policy determining the rules for scheduling servers in this group"},
					"state": schema.StringAttribute{
						Computed: true, MarkdownDescription: "State"},
				},
				Computed: true, MarkdownDescription: "Server group for instance scheduling policy",
			},
			"start_time": schema.StringAttribute{
				CustomType: timetypes.RFC3339Type{},
				Computed:   true, MarkdownDescription: "Start Time"},
			"state": schema.StringAttribute{
				Computed: true, MarkdownDescription: "State"},
			"tenant": schema.StringAttribute{
				Computed: true, MarkdownDescription: "The OpenStack tenant to create the instance in"},
			"tenant_uuid": schema.StringAttribute{
				Computed: true, MarkdownDescription: "UUID of the OpenStack tenant"},
			"url": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Url"},
			"user_data": schema.StringAttribute{
				Computed: true, MarkdownDescription: "Additional data that will be added to instance on provisioning"},
			"volumes": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"bootable": schema.BoolAttribute{
							Computed: true, MarkdownDescription: "Indicates if this volume can be used to boot an instance"},
						"device": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Name of volume as instance device e.g. /dev/vdb.",
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^/dev/[a-zA-Z0-9]+$`), ""),
							}},
						"image_name": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Name of the image this volume was created from"},
						"marketplace_resource_uuid": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Marketplace Resource Uuid"},
						"name": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Name"},
						"resource_type": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Resource Type"},
						"size": schema.Int64Attribute{
							Computed: true, MarkdownDescription: "Size in MiB",
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
								int64validator.AtMost(2147483647),
							}},
						"state": schema.StringAttribute{
							Computed: true, MarkdownDescription: "State"},
						"type": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Type of the volume (e.g. SSD, HDD)"},
						"type_name": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Type Name"},
						"url": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Url"},
						"uuid": schema.StringAttribute{
							Computed: true, MarkdownDescription: "Uuid"},
					},
				},
				Computed: true, MarkdownDescription: "List of volumes attached to the instance",
			},
		},
	}
}

func (d *OpenstackInstanceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	d.client = &OpenstackInstanceClient{}
	if err := d.client.Configure(ctx, req.ProviderData); err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			err.Error(),
		)
		return
	}
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
		apiResp, err := d.client.Get(ctx, data.UUID.ValueString())
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

		results, err := d.client.List(ctx, filters)
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
