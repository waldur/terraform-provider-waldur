package resources

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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

// OpenstackTenantApiResponse is the API response model.
type OpenstackTenantApiResponse struct {
	UUID *string `json:"uuid"`

	AccessUrl                   *string                                 `json:"access_url" tfsdk:"access_url"`
	AvailabilityZone            *string                                 `json:"availability_zone" tfsdk:"availability_zone"`
	BackendId                   *string                                 `json:"backend_id" tfsdk:"backend_id"`
	Created                     *string                                 `json:"created" tfsdk:"created"`
	Customer                    *string                                 `json:"customer" tfsdk:"customer"`
	CustomerAbbreviation        *string                                 `json:"customer_abbreviation" tfsdk:"customer_abbreviation"`
	CustomerName                *string                                 `json:"customer_name" tfsdk:"customer_name"`
	CustomerNativeName          *string                                 `json:"customer_native_name" tfsdk:"customer_native_name"`
	CustomerUuid                *string                                 `json:"customer_uuid" tfsdk:"customer_uuid"`
	DefaultVolumeTypeName       *string                                 `json:"default_volume_type_name" tfsdk:"default_volume_type_name"`
	Description                 *string                                 `json:"description" tfsdk:"description"`
	ErrorMessage                *string                                 `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback              *string                                 `json:"error_traceback" tfsdk:"error_traceback"`
	ExternalNetworkId           *string                                 `json:"external_network_id" tfsdk:"external_network_id"`
	InternalNetworkId           *string                                 `json:"internal_network_id" tfsdk:"internal_network_id"`
	IsLimitBased                *bool                                   `json:"is_limit_based" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                                   `json:"is_usage_based" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string                                 `json:"marketplace_category_name" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                                 `json:"marketplace_category_uuid" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                                 `json:"marketplace_offering_name" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                                 `json:"marketplace_offering_uuid" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                                 `json:"marketplace_plan_uuid" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                                 `json:"marketplace_resource_state" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                                 `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                                 `json:"modified" tfsdk:"modified"`
	Offering                    *string                                 `json:"offering" tfsdk:"offering"`
	Project                     *string                                 `json:"project" tfsdk:"project"`
	ProjectName                 *string                                 `json:"project_name" tfsdk:"project_name"`
	ProjectUuid                 *string                                 `json:"project_uuid" tfsdk:"project_uuid"`
	Quotas                      []OpenstackTenantQuotasResponse         `json:"quotas" tfsdk:"quotas"`
	ResourceType                *string                                 `json:"resource_type" tfsdk:"resource_type"`
	SecurityGroups              []OpenstackTenantSecurityGroupsResponse `json:"security_groups" tfsdk:"security_groups"`
	ServiceName                 *string                                 `json:"service_name" tfsdk:"service_name"`
	ServiceSettings             *string                                 `json:"service_settings" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                                 `json:"service_settings_error_message" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                                 `json:"service_settings_state" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                                 `json:"service_settings_uuid" tfsdk:"service_settings_uuid"`
	SkipConnectionExtnet        *bool                                   `json:"skip_connection_extnet" tfsdk:"skip_connection_extnet"`
	SkipCreationOfDefaultRouter *bool                                   `json:"skip_creation_of_default_router" tfsdk:"skip_creation_of_default_router"`
	SkipCreationOfDefaultSubnet *bool                                   `json:"skip_creation_of_default_subnet" tfsdk:"skip_creation_of_default_subnet"`
	State                       *string                                 `json:"state" tfsdk:"state"`
	SubnetCidr                  *string                                 `json:"subnet_cidr" tfsdk:"subnet_cidr"`
	Url                         *string                                 `json:"url" tfsdk:"url"`
	UserPassword                *string                                 `json:"user_password" tfsdk:"user_password"`
	UserUsername                *string                                 `json:"user_username" tfsdk:"user_username"`
}

type OpenstackTenantQuotasResponse struct {
	Limit *int64 `json:"limit" tfsdk:"limit"`
	Usage *int64 `json:"usage" tfsdk:"usage"`
}

type OpenstackTenantSecurityGroupsResponse struct {
	Description *string                                      `json:"description" tfsdk:"description"`
	Rules       []OpenstackTenantSecurityGroupsRulesResponse `json:"rules" tfsdk:"rules"`
}

type OpenstackTenantSecurityGroupsRulesResponse struct {
	Cidr        *string `json:"cidr" tfsdk:"cidr"`
	Description *string `json:"description" tfsdk:"description"`
	Direction   *string `json:"direction" tfsdk:"direction"`
	Ethertype   *string `json:"ethertype" tfsdk:"ethertype"`
	FromPort    *int64  `json:"from_port" tfsdk:"from_port"`
	Protocol    *string `json:"protocol" tfsdk:"protocol"`
	RemoteGroup *string `json:"remote_group" tfsdk:"remote_group"`
	ToPort      *int64  `json:"to_port" tfsdk:"to_port"`
}

var openstacktenant_quotasAttrTypes = map[string]attr.Type{
	"limit": types.Int64Type,
	"name":  types.StringType,
	"usage": types.Int64Type,
}
var openstacktenant_quotasObjectType = types.ObjectType{
	AttrTypes: openstacktenant_quotasAttrTypes,
}

var openstacktenant_security_groupsAttrTypes = map[string]attr.Type{
	"description": types.StringType,
	"name":        types.StringType,
	"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
		"cidr":         types.StringType,
		"description":  types.StringType,
		"direction":    types.StringType,
		"ethertype":    types.StringType,
		"from_port":    types.Int64Type,
		"protocol":     types.StringType,
		"remote_group": types.StringType,
		"to_port":      types.Int64Type,
	}}},
}
var openstacktenant_security_groupsObjectType = types.ObjectType{
	AttrTypes: openstacktenant_security_groupsAttrTypes,
}

var openstacktenantsecuritygroups_rulesAttrTypes = map[string]attr.Type{
	"cidr":         types.StringType,
	"description":  types.StringType,
	"direction":    types.StringType,
	"ethertype":    types.StringType,
	"from_port":    types.Int64Type,
	"protocol":     types.StringType,
	"remote_group": types.StringType,
	"to_port":      types.Int64Type,
}
var openstacktenantsecuritygroups_rulesObjectType = types.ObjectType{
	AttrTypes: openstacktenantsecuritygroups_rulesAttrTypes,
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
	SkipCreationOfDefaultSubnet types.Bool     `tfsdk:"skip_creation_of_default_subnet"`
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
		MarkdownDescription: "Openstack Tenant resource",

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
				MarkdownDescription: "Offering URL",
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
			"skip_creation_of_default_subnet": schema.BoolAttribute{
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
			items = append(items, ConvertTFValue(elem))
		}
		attributes["security_groups"] = items
	}
	if !data.SkipConnectionExtnet.IsNull() {
		attributes["skip_connection_extnet"] = data.SkipConnectionExtnet.ValueBool()
	}
	if !data.SkipCreationOfDefaultRouter.IsNull() {
		attributes["skip_creation_of_default_router"] = data.SkipCreationOfDefaultRouter.ValueBool()
	}
	if !data.SkipCreationOfDefaultSubnet.IsNull() {
		attributes["skip_creation_of_default_subnet"] = data.SkipCreationOfDefaultSubnet.ValueBool()
	}
	if !data.SubnetCidr.IsNull() {
		attributes["subnet_cidr"] = data.SubnetCidr.ValueString()
	}

	payload := map[string]interface{}{
		"project":    data.Project.ValueString(),
		"offering":   data.Offering.ValueString(),
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
						var apiResp OpenstackTenantApiResponse
						retrievePath := strings.Replace("/api/openstack-tenants/{uuid}/", "{uuid}", pluginUUID, 1)
						tflog.Warn(ctx, "Attempting to fetch plugin resource at: "+retrievePath)
						err = r.client.GetByUUID(ctx, retrievePath, pluginUUID, &apiResp)
						if err == nil {
							tflog.Warn(ctx, "Successfully fetched plugin resource")
							resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)
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
	var apiResp OpenstackTenantApiResponse
	err = r.client.GetByUUID(ctx, "/api/openstack-tenants/{uuid}/", data.UUID.ValueString(), &apiResp)
	if err != nil {
		resp.Diagnostics.AddError("Failed to Read Resource", err.Error())
		return
	}
	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

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

	retrievePath := strings.Replace("/api/openstack-tenants/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	var apiResp OpenstackTenantApiResponse
	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Openstack Tenant",
			"An error occurred while reading the Openstack Tenant: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackTenantResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpenstackTenantResourceModel
	var state OpenstackTenantResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

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
	if !data.SkipCreationOfDefaultRouter.IsNull() && !data.SkipCreationOfDefaultRouter.Equal(state.SkipCreationOfDefaultRouter) {
		patchPayload["skip_creation_of_default_router"] = data.SkipCreationOfDefaultRouter.ValueBool()
	}
	if !data.SkipCreationOfDefaultSubnet.IsNull() && !data.SkipCreationOfDefaultSubnet.Equal(state.SkipCreationOfDefaultSubnet) {
		patchPayload["skip_creation_of_default_subnet"] = data.SkipCreationOfDefaultSubnet.ValueBool()
	}

	if len(patchPayload) > 0 {
		var result map[string]interface{}
		err := r.client.Update(ctx, "/api/openstack-tenants/{uuid}/", data.UUID.ValueString(), patchPayload, &result)
		if err != nil {
			resp.Diagnostics.AddError("Update Failed", err.Error())
			return
		}
		_ = result
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
	var apiResp OpenstackTenantApiResponse

	retrievePath := strings.Replace("/api/openstack-tenants/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &apiResp)
	if err != nil {
		resp.Diagnostics.AddError("Failed to Read Resource After Update", err.Error())
		return
	}
	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackTenantResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackTenantResourceModel
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

func (r *OpenstackTenantResource) mapResponseToModel(ctx context.Context, apiResp OpenstackTenantApiResponse, model *OpenstackTenantResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
	model.AvailabilityZone = types.StringPointerValue(apiResp.AvailabilityZone)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.Customer = types.StringPointerValue(apiResp.Customer)
	model.CustomerAbbreviation = types.StringPointerValue(apiResp.CustomerAbbreviation)
	model.CustomerName = types.StringPointerValue(apiResp.CustomerName)
	model.CustomerNativeName = types.StringPointerValue(apiResp.CustomerNativeName)
	model.CustomerUuid = types.StringPointerValue(apiResp.CustomerUuid)
	model.DefaultVolumeTypeName = types.StringPointerValue(apiResp.DefaultVolumeTypeName)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.ExternalNetworkId = types.StringPointerValue(apiResp.ExternalNetworkId)
	model.InternalNetworkId = types.StringPointerValue(apiResp.InternalNetworkId)
	model.IsLimitBased = types.BoolPointerValue(apiResp.IsLimitBased)
	model.IsUsageBased = types.BoolPointerValue(apiResp.IsUsageBased)
	model.MarketplaceCategoryName = types.StringPointerValue(apiResp.MarketplaceCategoryName)
	model.MarketplaceCategoryUuid = types.StringPointerValue(apiResp.MarketplaceCategoryUuid)
	model.MarketplaceOfferingName = types.StringPointerValue(apiResp.MarketplaceOfferingName)
	model.MarketplaceOfferingUuid = types.StringPointerValue(apiResp.MarketplaceOfferingUuid)
	model.MarketplacePlanUuid = types.StringPointerValue(apiResp.MarketplacePlanUuid)
	model.MarketplaceResourceState = types.StringPointerValue(apiResp.MarketplaceResourceState)
	model.MarketplaceResourceUuid = types.StringPointerValue(apiResp.MarketplaceResourceUuid)
	model.Modified = types.StringPointerValue(apiResp.Modified)
	if apiResp.Offering != nil {
		parts := strings.Split(strings.TrimRight(*apiResp.Offering, "/"), "/")
		model.Offering = types.StringValue(parts[len(parts)-1])
	} else {
		model.Offering = types.StringNull()
	}
	if apiResp.Project != nil {
		parts := strings.Split(strings.TrimRight(*apiResp.Project, "/"), "/")
		model.Project = types.StringValue(parts[len(parts)-1])
	} else {
		model.Project = types.StringNull()
	}
	model.ProjectName = types.StringPointerValue(apiResp.ProjectName)
	model.ProjectUuid = types.StringPointerValue(apiResp.ProjectUuid)
	listValQuotas, listDiagsQuotas := types.ListValueFrom(ctx, openstacktenant_quotasObjectType, apiResp.Quotas)
	diags.Append(listDiagsQuotas...)
	model.Quotas = listValQuotas
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	listValSecurityGroups, listDiagsSecurityGroups := types.ListValueFrom(ctx, openstacktenant_security_groupsObjectType, apiResp.SecurityGroups)
	diags.Append(listDiagsSecurityGroups...)
	model.SecurityGroups = listValSecurityGroups
	model.ServiceName = types.StringPointerValue(apiResp.ServiceName)
	model.ServiceSettings = types.StringPointerValue(apiResp.ServiceSettings)
	model.ServiceSettingsErrorMessage = types.StringPointerValue(apiResp.ServiceSettingsErrorMessage)
	model.ServiceSettingsState = types.StringPointerValue(apiResp.ServiceSettingsState)
	model.ServiceSettingsUuid = types.StringPointerValue(apiResp.ServiceSettingsUuid)
	model.SkipConnectionExtnet = types.BoolPointerValue(apiResp.SkipConnectionExtnet)
	model.SkipCreationOfDefaultRouter = types.BoolPointerValue(apiResp.SkipCreationOfDefaultRouter)
	model.SkipCreationOfDefaultSubnet = types.BoolPointerValue(apiResp.SkipCreationOfDefaultSubnet)
	model.State = types.StringPointerValue(apiResp.State)
	model.SubnetCidr = types.StringPointerValue(apiResp.SubnetCidr)
	model.Url = types.StringPointerValue(apiResp.Url)
	model.UserPassword = types.StringPointerValue(apiResp.UserPassword)
	model.UserUsername = types.StringPointerValue(apiResp.UserUsername)

	return diags
}
