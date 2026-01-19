package resources

import (
	"context"
	"fmt"
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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OpenstackTenantResource{}
var _ resource.ResourceWithImportState = &OpenstackTenantResource{}

func NewOpenstackTenantResource() resource.Resource {
	return &OpenstackTenantResource{}
}

// OpenstackTenantResource defines the resource implementation.
type OpenstackTenantResource struct {
	client *client.Client
}

// OpenstackTenantResourceModel describes the resource data model.
type OpenstackTenantResourceModel struct {
	UUID                        types.String   `tfsdk:"id"`
	AccessUrl                   types.String   `tfsdk:"access_url"`
	AvailabilityZone            types.String   `tfsdk:"availability_zone"`
	BackendId                   types.String   `tfsdk:"backend_id"`
	Created                     types.String   `tfsdk:"created"`
	Customer                    types.String   `tfsdk:"customer"`
	CustomerAbbreviation        types.String   `tfsdk:"customer_abbreviation"`
	CustomerName                types.String   `tfsdk:"customer_name"`
	CustomerNativeName          types.String   `tfsdk:"customer_native_name"`
	CustomerUuid                types.String   `tfsdk:"customer_uuid"`
	DefaultVolumeTypeName       types.String   `tfsdk:"default_volume_type_name"`
	Description                 types.String   `tfsdk:"description"`
	ErrorMessage                types.String   `tfsdk:"error_message"`
	ErrorTraceback              types.String   `tfsdk:"error_traceback"`
	ExternalNetworkId           types.String   `tfsdk:"external_network_id"`
	InternalNetworkId           types.String   `tfsdk:"internal_network_id"`
	IsLimitBased                types.Bool     `tfsdk:"is_limit_based"`
	IsUsageBased                types.Bool     `tfsdk:"is_usage_based"`
	MarketplaceCategoryName     types.String   `tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     types.String   `tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     types.String   `tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     types.String   `tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         types.String   `tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    types.String   `tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     types.String   `tfsdk:"marketplace_resource_uuid"`
	Modified                    types.String   `tfsdk:"modified"`
	Name                        types.String   `tfsdk:"name"`
	Offering                    types.String   `tfsdk:"offering"`
	Project                     types.String   `tfsdk:"project"`
	ProjectName                 types.String   `tfsdk:"project_name"`
	ProjectUuid                 types.String   `tfsdk:"project_uuid"`
	Quotas                      types.List     `tfsdk:"quotas"`
	ResourceType                types.String   `tfsdk:"resource_type"`
	SecurityGroups              types.List     `tfsdk:"security_groups"`
	ServiceName                 types.String   `tfsdk:"service_name"`
	ServiceSettings             types.String   `tfsdk:"service_settings"`
	ServiceSettingsErrorMessage types.String   `tfsdk:"service_settings_error_message"`
	ServiceSettingsState        types.String   `tfsdk:"service_settings_state"`
	ServiceSettingsUuid         types.String   `tfsdk:"service_settings_uuid"`
	SkipConnectionExtnet        types.Bool     `tfsdk:"skip_connection_extnet"`
	SkipCreationOfDefaultRouter types.Bool     `tfsdk:"skip_creation_of_default_router"`
	State                       types.String   `tfsdk:"state"`
	SubnetCidr                  types.String   `tfsdk:"subnet_cidr"`
	Url                         types.String   `tfsdk:"url"`
	UserPassword                types.String   `tfsdk:"user_password"`
	UserUsername                types.String   `tfsdk:"user_username"`
	Timeouts                    timeouts.Value `tfsdk:"timeouts"`
}

func (r *OpenstackTenantResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_tenant"
}

func (r *OpenstackTenantResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "OpenstackTenant resource",

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
				MarkdownDescription: " ",
			},
			"availability_zone": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Optional availability group. Will be used for all instances provisioned in this tenant",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of tenant in the OpenStack backend",
			},
			"created": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"customer": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"customer_abbreviation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"customer_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"customer_native_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"customer_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"default_volume_type_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Volume type name to use when creating volumes.",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"external_network_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of external network connected to OpenStack tenant",
			},
			"internal_network_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ID of internal network in OpenStack tenant",
			},
			"is_limit_based": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"is_usage_based": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_category_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_category_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_offering_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_offering_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_plan_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_resource_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"marketplace_resource_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"modified": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"offering": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Offering UUID or URL",
			},
			"project": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: " ",
			},
			"project_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"project_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"quotas": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"limit": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"name": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"usage": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: " ",
			},
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"security_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"description": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: " ",
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
										MarkdownDescription: " ",
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
									"to_port": schema.Int64Attribute{
										Optional:            true,
										Computed:            true,
										MarkdownDescription: "Ending port number in the range (1-65535)",
									},
								},
							},
							Optional:            true,
							Computed:            true,
							MarkdownDescription: " ",
						},
					},
				},
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"service_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"service_settings": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"service_settings_error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"service_settings_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"service_settings_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"skip_connection_extnet": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"skip_creation_of_default_router": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"subnet_cidr": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"user_password": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Password of the tenant user",
			},
			"user_username": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Username of the tenant user",
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

func (r *OpenstackTenantResource) convertTFValue(v attr.Value) interface{} {
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

func (r *OpenstackTenantResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *OpenstackTenantResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackTenantResourceModel

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
	if !data.Description.IsNull() {
		attributes["description"] = data.Description.ValueString()
	}
	if !data.Name.IsNull() {
		attributes["name"] = data.Name.ValueString()
	}
	if !data.SecurityGroups.IsNull() {
		items := make([]interface{}, 0)
		for _, elem := range data.SecurityGroups.Elements() {
			items = append(items, r.convertTFValue(elem))
		}
		attributes["security_groups"] = items
	}
	if !data.SkipConnectionExtnet.IsNull() {
		attributes["skip_connection_extnet"] = data.SkipConnectionExtnet.ValueBool()
	}
	if !data.SkipCreationOfDefaultRouter.IsNull() {
		attributes["skip_creation_of_default_router"] = data.SkipCreationOfDefaultRouter.ValueBool()
	}
	if !data.SubnetCidr.IsNull() {
		attributes["subnet_cidr"] = data.SubnetCidr.ValueString()
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
						retrievePath := strings.Replace("/api/openstack-tenants/{uuid}/", "{uuid}", pluginUUID, 1)
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
							if val, ok := sourceMap["availability_zone"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.AvailabilityZone = types.StringValue(str)
								}
							} else {
								if data.AvailabilityZone.IsUnknown() {
									data.AvailabilityZone = types.StringNull()
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
							if val, ok := sourceMap["default_volume_type_name"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.DefaultVolumeTypeName = types.StringValue(str)
								}
							} else {
								if data.DefaultVolumeTypeName.IsUnknown() {
									data.DefaultVolumeTypeName = types.StringNull()
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
							if val, ok := sourceMap["external_network_id"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.ExternalNetworkId = types.StringValue(str)
								}
							} else {
								if data.ExternalNetworkId.IsUnknown() {
									data.ExternalNetworkId = types.StringNull()
								}
							}
							if val, ok := sourceMap["internal_network_id"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.InternalNetworkId = types.StringValue(str)
								}
							} else {
								if data.InternalNetworkId.IsUnknown() {
									data.InternalNetworkId = types.StringNull()
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
							if val, ok := sourceMap["quotas"]; ok && val != nil {
								// List of objects
								if arr, ok := val.([]interface{}); ok {
									items := make([]attr.Value, 0, len(arr))
									for _, item := range arr {
										if objMap, ok := item.(map[string]interface{}); ok {
											attrTypes := map[string]attr.Type{
												"limit": types.Int64Type,
												"name":  types.StringType,
												"usage": types.Int64Type,
											}
											attrValues := map[string]attr.Value{
												"limit": func() attr.Value {
													if v, ok := objMap["limit"].(float64); ok {
														return types.Int64Value(int64(v))
													}
													return types.Int64Null()
												}(),
												"name": func() attr.Value {
													if v, ok := objMap["name"].(string); ok {
														return types.StringValue(v)
													}
													return types.StringNull()
												}(),
												"usage": func() attr.Value {
													if v, ok := objMap["usage"].(float64); ok {
														return types.Int64Value(int64(v))
													}
													return types.Int64Null()
												}(),
											}
											objVal, _ := types.ObjectValue(attrTypes, attrValues)
											items = append(items, objVal)
										}
									}
									listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
										"limit": types.Int64Type,
										"name":  types.StringType,
										"usage": types.Int64Type,
									}}, items)
									data.Quotas = listVal
								}
							} else {
								if data.Quotas.IsUnknown() {
									data.Quotas = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
										"limit": types.Int64Type,
										"name":  types.StringType,
										"usage": types.Int64Type,
									}})
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
							if val, ok := sourceMap["security_groups"]; ok && val != nil {
								// List of objects
								if arr, ok := val.([]interface{}); ok {
									items := make([]attr.Value, 0, len(arr))
									for _, item := range arr {
										if objMap, ok := item.(map[string]interface{}); ok {
											attrTypes := map[string]attr.Type{
												"description": types.StringType,
												"name":        types.StringType,
												"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "to_port": types.Int64Type}}},
											}
											attrValues := map[string]attr.Value{
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
												"rules": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "to_port": types.Int64Type}}}.ElemType),
											}
											objVal, _ := types.ObjectValue(attrTypes, attrValues)
											items = append(items, objVal)
										}
									}
									listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
										"description": types.StringType,
										"name":        types.StringType,
										"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "to_port": types.Int64Type}}},
									}}, items)
									data.SecurityGroups = listVal
								}
							} else {
								if data.SecurityGroups.IsUnknown() {
									data.SecurityGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
										"description": types.StringType,
										"name":        types.StringType,
										"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "to_port": types.Int64Type}}},
									}})
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
							if val, ok := sourceMap["skip_connection_extnet"]; ok && val != nil {
								if b, ok := val.(bool); ok {
									data.SkipConnectionExtnet = types.BoolValue(b)
								}
							} else {
								if data.SkipConnectionExtnet.IsUnknown() {
									data.SkipConnectionExtnet = types.BoolNull()
								}
							}
							if val, ok := sourceMap["skip_creation_of_default_router"]; ok && val != nil {
								if b, ok := val.(bool); ok {
									data.SkipCreationOfDefaultRouter = types.BoolValue(b)
								}
							} else {
								if data.SkipCreationOfDefaultRouter.IsUnknown() {
									data.SkipCreationOfDefaultRouter = types.BoolNull()
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
							if val, ok := sourceMap["subnet_cidr"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.SubnetCidr = types.StringValue(str)
								}
							} else {
								if data.SubnetCidr.IsUnknown() {
									data.SubnetCidr = types.StringNull()
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
							if val, ok := sourceMap["user_password"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.UserPassword = types.StringValue(str)
								}
							} else {
								if data.UserPassword.IsUnknown() {
									data.UserPassword = types.StringNull()
								}
							}
							if val, ok := sourceMap["user_username"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.UserUsername = types.StringValue(str)
								}
							} else {
								if data.UserUsername.IsUnknown() {
									data.UserUsername = types.StringNull()
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
	err = r.client.GetByUUID(ctx, "/api/openstack-tenants/{uuid}/", data.UUID.ValueString(), &finalState)
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
	if val, ok := sourceMap["availability_zone"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AvailabilityZone = types.StringValue(str)
		}
	} else {
		if data.AvailabilityZone.IsUnknown() {
			data.AvailabilityZone = types.StringNull()
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
	if val, ok := sourceMap["default_volume_type_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.DefaultVolumeTypeName = types.StringValue(str)
		}
	} else {
		if data.DefaultVolumeTypeName.IsUnknown() {
			data.DefaultVolumeTypeName = types.StringNull()
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
	if val, ok := sourceMap["external_network_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ExternalNetworkId = types.StringValue(str)
		}
	} else {
		if data.ExternalNetworkId.IsUnknown() {
			data.ExternalNetworkId = types.StringNull()
		}
	}
	if val, ok := sourceMap["internal_network_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.InternalNetworkId = types.StringValue(str)
		}
	} else {
		if data.InternalNetworkId.IsUnknown() {
			data.InternalNetworkId = types.StringNull()
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
	if val, ok := sourceMap["quotas"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"limit": types.Int64Type,
						"name":  types.StringType,
						"usage": types.Int64Type,
					}
					attrValues := map[string]attr.Value{
						"limit": func() attr.Value {
							if v, ok := objMap["limit"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"usage": func() attr.Value {
							if v, ok := objMap["usage"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"limit": types.Int64Type,
				"name":  types.StringType,
				"usage": types.Int64Type,
			}}, items)
			data.Quotas = listVal
		}
	} else {
		if data.Quotas.IsUnknown() {
			data.Quotas = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"limit": types.Int64Type,
				"name":  types.StringType,
				"usage": types.Int64Type,
			}})
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
	if val, ok := sourceMap["security_groups"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"description": types.StringType,
						"name":        types.StringType,
						"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "to_port": types.Int64Type}}},
					}
					attrValues := map[string]attr.Value{
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
						"rules": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "to_port": types.Int64Type}}}.ElemType),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"description": types.StringType,
				"name":        types.StringType,
				"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "to_port": types.Int64Type}}},
			}}, items)
			data.SecurityGroups = listVal
		}
	} else {
		if data.SecurityGroups.IsUnknown() {
			data.SecurityGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"description": types.StringType,
				"name":        types.StringType,
				"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "to_port": types.Int64Type}}},
			}})
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
	if val, ok := sourceMap["skip_connection_extnet"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.SkipConnectionExtnet = types.BoolValue(b)
		}
	} else {
		if data.SkipConnectionExtnet.IsUnknown() {
			data.SkipConnectionExtnet = types.BoolNull()
		}
	}
	if val, ok := sourceMap["skip_creation_of_default_router"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.SkipCreationOfDefaultRouter = types.BoolValue(b)
		}
	} else {
		if data.SkipCreationOfDefaultRouter.IsUnknown() {
			data.SkipCreationOfDefaultRouter = types.BoolNull()
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
	if val, ok := sourceMap["subnet_cidr"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.SubnetCidr = types.StringValue(str)
		}
	} else {
		if data.SubnetCidr.IsUnknown() {
			data.SubnetCidr = types.StringNull()
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
	if val, ok := sourceMap["user_password"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.UserPassword = types.StringValue(str)
		}
	} else {
		if data.UserPassword.IsUnknown() {
			data.UserPassword = types.StringNull()
		}
	}
	if val, ok := sourceMap["user_username"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.UserUsername = types.StringValue(str)
		}
	} else {
		if data.UserUsername.IsUnknown() {
			data.UserUsername = types.StringNull()
		}
	}

	// Map filter parameters from response if available

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackTenantResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpenstackTenantResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource
	var result map[string]interface{}

	retrievePath := strings.Replace("/api/openstack-tenants/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read OpenstackTenant",
			"An error occurred while reading the openstack_tenant: "+err.Error(),
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
	if val, ok := sourceMap["availability_zone"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AvailabilityZone = types.StringValue(str)
		}
	} else {
		if data.AvailabilityZone.IsUnknown() {
			data.AvailabilityZone = types.StringNull()
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
	if val, ok := sourceMap["default_volume_type_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.DefaultVolumeTypeName = types.StringValue(str)
		}
	} else {
		if data.DefaultVolumeTypeName.IsUnknown() {
			data.DefaultVolumeTypeName = types.StringNull()
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
	if val, ok := sourceMap["external_network_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ExternalNetworkId = types.StringValue(str)
		}
	} else {
		if data.ExternalNetworkId.IsUnknown() {
			data.ExternalNetworkId = types.StringNull()
		}
	}
	if val, ok := sourceMap["internal_network_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.InternalNetworkId = types.StringValue(str)
		}
	} else {
		if data.InternalNetworkId.IsUnknown() {
			data.InternalNetworkId = types.StringNull()
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
	if val, ok := sourceMap["quotas"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"limit": types.Int64Type,
						"name":  types.StringType,
						"usage": types.Int64Type,
					}
					attrValues := map[string]attr.Value{
						"limit": func() attr.Value {
							if v, ok := objMap["limit"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"usage": func() attr.Value {
							if v, ok := objMap["usage"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"limit": types.Int64Type,
				"name":  types.StringType,
				"usage": types.Int64Type,
			}}, items)
			data.Quotas = listVal
		}
	} else {
		if data.Quotas.IsUnknown() {
			data.Quotas = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"limit": types.Int64Type,
				"name":  types.StringType,
				"usage": types.Int64Type,
			}})
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
	if val, ok := sourceMap["security_groups"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"description": types.StringType,
						"name":        types.StringType,
						"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "to_port": types.Int64Type}}},
					}
					attrValues := map[string]attr.Value{
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
						"rules": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "to_port": types.Int64Type}}}.ElemType),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"description": types.StringType,
				"name":        types.StringType,
				"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "to_port": types.Int64Type}}},
			}}, items)
			data.SecurityGroups = listVal
		}
	} else {
		if data.SecurityGroups.IsUnknown() {
			data.SecurityGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"description": types.StringType,
				"name":        types.StringType,
				"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "to_port": types.Int64Type}}},
			}})
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
	if val, ok := sourceMap["skip_connection_extnet"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.SkipConnectionExtnet = types.BoolValue(b)
		}
	} else {
		if data.SkipConnectionExtnet.IsUnknown() {
			data.SkipConnectionExtnet = types.BoolNull()
		}
	}
	if val, ok := sourceMap["skip_creation_of_default_router"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.SkipCreationOfDefaultRouter = types.BoolValue(b)
		}
	} else {
		if data.SkipCreationOfDefaultRouter.IsUnknown() {
			data.SkipCreationOfDefaultRouter = types.BoolNull()
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
	if val, ok := sourceMap["subnet_cidr"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.SubnetCidr = types.StringValue(str)
		}
	} else {
		if data.SubnetCidr.IsUnknown() {
			data.SubnetCidr = types.StringNull()
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
	if val, ok := sourceMap["user_password"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.UserPassword = types.StringValue(str)
		}
	} else {
		if data.UserPassword.IsUnknown() {
			data.UserPassword = types.StringNull()
		}
	}
	if val, ok := sourceMap["user_username"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.UserUsername = types.StringValue(str)
		}
	} else {
		if data.UserUsername.IsUnknown() {
			data.UserUsername = types.StringNull()
		}
	}

	// Map filter parameters from response if available

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackTenantResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpenstackTenantResourceModel
	var state OpenstackTenantResourceModel

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
	if !data.AvailabilityZone.IsNull() && !data.AvailabilityZone.Equal(state.AvailabilityZone) {
		patchPayload["availability_zone"] = data.AvailabilityZone.ValueString()
	}
	if !data.DefaultVolumeTypeName.IsNull() && !data.DefaultVolumeTypeName.Equal(state.DefaultVolumeTypeName) {
		patchPayload["default_volume_type_name"] = data.DefaultVolumeTypeName.ValueString()
	}
	if !data.Description.IsNull() && !data.Description.Equal(state.Description) {
		patchPayload["description"] = data.Description.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.Equal(state.Name) {
		patchPayload["name"] = data.Name.ValueString()
	}

	if len(patchPayload) > 0 {
		var result map[string]interface{}
		err := r.client.Update(ctx, "/api/openstack-tenants/{uuid}/", data.UUID.ValueString(), patchPayload, &result)
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
	if !data.SecurityGroups.Equal(state.SecurityGroups) {

		// Convert Terraform value to API payload

		var itemsPushSecurityGroups []interface{}
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
				itemsPushSecurityGroups = append(itemsPushSecurityGroups, objMap)
			}
		}
		actionPayloadPushSecurityGroups := map[string]interface{}{
			"security_groups": itemsPushSecurityGroups,
		}
		actionUrlPushSecurityGroups := strings.Replace("/api/openstack-tenants/{uuid}/push_security_groups/", "{uuid}", data.UUID.ValueString(), 1)
		var actionResultPushSecurityGroups map[string]interface{}
		if err := r.client.Post(ctx, actionUrlPushSecurityGroups, actionPayloadPushSecurityGroups, &actionResultPushSecurityGroups); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: push_security_groups", err.Error())
			return
		}
	}

	// Fetch updated state
	// Call Waldur API to read resource
	var result map[string]interface{}

	retrievePath := strings.Replace("/api/openstack-tenants/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

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
	if val, ok := sourceMap["availability_zone"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.AvailabilityZone = types.StringValue(str)
		}
	} else {
		if data.AvailabilityZone.IsUnknown() {
			data.AvailabilityZone = types.StringNull()
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
	if val, ok := sourceMap["default_volume_type_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.DefaultVolumeTypeName = types.StringValue(str)
		}
	} else {
		if data.DefaultVolumeTypeName.IsUnknown() {
			data.DefaultVolumeTypeName = types.StringNull()
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
	if val, ok := sourceMap["external_network_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ExternalNetworkId = types.StringValue(str)
		}
	} else {
		if data.ExternalNetworkId.IsUnknown() {
			data.ExternalNetworkId = types.StringNull()
		}
	}
	if val, ok := sourceMap["internal_network_id"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.InternalNetworkId = types.StringValue(str)
		}
	} else {
		if data.InternalNetworkId.IsUnknown() {
			data.InternalNetworkId = types.StringNull()
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
	if val, ok := sourceMap["quotas"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"limit": types.Int64Type,
						"name":  types.StringType,
						"usage": types.Int64Type,
					}
					attrValues := map[string]attr.Value{
						"limit": func() attr.Value {
							if v, ok := objMap["limit"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
						"name": func() attr.Value {
							if v, ok := objMap["name"].(string); ok {
								return types.StringValue(v)
							}
							return types.StringNull()
						}(),
						"usage": func() attr.Value {
							if v, ok := objMap["usage"].(float64); ok {
								return types.Int64Value(int64(v))
							}
							return types.Int64Null()
						}(),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"limit": types.Int64Type,
				"name":  types.StringType,
				"usage": types.Int64Type,
			}}, items)
			data.Quotas = listVal
		}
	} else {
		if data.Quotas.IsUnknown() {
			data.Quotas = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"limit": types.Int64Type,
				"name":  types.StringType,
				"usage": types.Int64Type,
			}})
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
	if val, ok := sourceMap["security_groups"]; ok && val != nil {
		// List of objects
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if objMap, ok := item.(map[string]interface{}); ok {
					attrTypes := map[string]attr.Type{
						"description": types.StringType,
						"name":        types.StringType,
						"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "to_port": types.Int64Type}}},
					}
					attrValues := map[string]attr.Value{
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
						"rules": types.ListNull(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "to_port": types.Int64Type}}}.ElemType),
					}
					objVal, _ := types.ObjectValue(attrTypes, attrValues)
					items = append(items, objVal)
				}
			}
			listVal, _ := types.ListValue(types.ObjectType{AttrTypes: map[string]attr.Type{
				"description": types.StringType,
				"name":        types.StringType,
				"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "to_port": types.Int64Type}}},
			}}, items)
			data.SecurityGroups = listVal
		}
	} else {
		if data.SecurityGroups.IsUnknown() {
			data.SecurityGroups = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
				"description": types.StringType,
				"name":        types.StringType,
				"rules":       types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"cidr": types.StringType, "description": types.StringType, "direction": types.StringType, "ethertype": types.StringType, "from_port": types.Int64Type, "protocol": types.StringType, "remote_group": types.StringType, "to_port": types.Int64Type}}},
			}})
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
	if val, ok := sourceMap["skip_connection_extnet"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.SkipConnectionExtnet = types.BoolValue(b)
		}
	} else {
		if data.SkipConnectionExtnet.IsUnknown() {
			data.SkipConnectionExtnet = types.BoolNull()
		}
	}
	if val, ok := sourceMap["skip_creation_of_default_router"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.SkipCreationOfDefaultRouter = types.BoolValue(b)
		}
	} else {
		if data.SkipCreationOfDefaultRouter.IsUnknown() {
			data.SkipCreationOfDefaultRouter = types.BoolNull()
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
	if val, ok := sourceMap["subnet_cidr"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.SubnetCidr = types.StringValue(str)
		}
	} else {
		if data.SubnetCidr.IsUnknown() {
			data.SubnetCidr = types.StringNull()
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
	if val, ok := sourceMap["user_password"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.UserPassword = types.StringValue(str)
		}
	} else {
		if data.UserPassword.IsUnknown() {
			data.UserPassword = types.StringNull()
		}
	}
	if val, ok := sourceMap["user_username"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.UserUsername = types.StringValue(str)
		}
	} else {
		if data.UserUsername.IsUnknown() {
			data.UserUsername = types.StringNull()
		}
	}

	// Map filter parameters from response if available

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackTenantResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackTenantResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Order-based Delete
	payload := map[string]interface{}{}

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

func (r *OpenstackTenantResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
