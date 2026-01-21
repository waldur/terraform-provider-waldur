package resources

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OpenstackVolumeResource{}
var _ resource.ResourceWithImportState = &OpenstackVolumeResource{}

func NewOpenstackVolumeResource() resource.Resource {
	return &OpenstackVolumeResource{}
}

// OpenstackVolumeResource defines the resource implementation.
type OpenstackVolumeResource struct {
	client *client.Client
}

// OpenstackVolumeApiResponse is the API response model.
type OpenstackVolumeApiResponse struct {
	UUID *string `json:"uuid"`

	AccessUrl                   *string `json:"access_url" tfsdk:"access_url"`
	Action                      *string `json:"action" tfsdk:"action"`
	AvailabilityZone            *string `json:"availability_zone" tfsdk:"availability_zone"`
	AvailabilityZoneName        *string `json:"availability_zone_name" tfsdk:"availability_zone_name"`
	BackendId                   *string `json:"backend_id" tfsdk:"backend_id"`
	Bootable                    *bool   `json:"bootable" tfsdk:"bootable"`
	Created                     *string `json:"created" tfsdk:"created"`
	Customer                    *string `json:"customer" tfsdk:"customer"`
	CustomerAbbreviation        *string `json:"customer_abbreviation" tfsdk:"customer_abbreviation"`
	CustomerName                *string `json:"customer_name" tfsdk:"customer_name"`
	CustomerNativeName          *string `json:"customer_native_name" tfsdk:"customer_native_name"`
	CustomerUuid                *string `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description                 *string `json:"description" tfsdk:"description"`
	Device                      *string `json:"device" tfsdk:"device"`
	ErrorMessage                *string `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback              *string `json:"error_traceback" tfsdk:"error_traceback"`
	ExtendEnabled               *bool   `json:"extend_enabled" tfsdk:"extend_enabled"`
	Image                       *string `json:"image" tfsdk:"image"`
	ImageMetadata               *string `json:"image_metadata" tfsdk:"image_metadata"`
	ImageName                   *string `json:"image_name" tfsdk:"image_name"`
	Instance                    *string `json:"instance" tfsdk:"instance"`
	InstanceMarketplaceUuid     *string `json:"instance_marketplace_uuid" tfsdk:"instance_marketplace_uuid"`
	InstanceName                *string `json:"instance_name" tfsdk:"instance_name"`
	IsLimitBased                *bool   `json:"is_limit_based" tfsdk:"is_limit_based"`
	IsUsageBased                *bool   `json:"is_usage_based" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string `json:"marketplace_category_name" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string `json:"marketplace_category_uuid" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string `json:"marketplace_offering_name" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string `json:"marketplace_offering_uuid" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string `json:"marketplace_plan_uuid" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string `json:"marketplace_resource_state" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string `json:"modified" tfsdk:"modified"`
	Name                        *string `json:"name" tfsdk:"name"`
	Offering                    *string `json:"offering" tfsdk:"offering"`
	Project                     *string `json:"project" tfsdk:"project"`
	ProjectName                 *string `json:"project_name" tfsdk:"project_name"`
	ProjectUuid                 *string `json:"project_uuid" tfsdk:"project_uuid"`
	ResourceType                *string `json:"resource_type" tfsdk:"resource_type"`
	RuntimeState                *string `json:"runtime_state" tfsdk:"runtime_state"`
	ServiceName                 *string `json:"service_name" tfsdk:"service_name"`
	ServiceSettings             *string `json:"service_settings" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string `json:"service_settings_error_message" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string `json:"service_settings_state" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string `json:"service_settings_uuid" tfsdk:"service_settings_uuid"`
	Size                        *int64  `json:"size" tfsdk:"size"`
	SourceSnapshot              *string `json:"source_snapshot" tfsdk:"source_snapshot"`
	State                       *string `json:"state" tfsdk:"state"`
	Tenant                      *string `json:"tenant" tfsdk:"tenant"`
	TenantUuid                  *string `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Type                        *string `json:"type" tfsdk:"type"`
	TypeName                    *string `json:"type_name" tfsdk:"type_name"`
	Url                         *string `json:"url" tfsdk:"url"`
}

// OpenstackVolumeResourceModel describes the resource data model.
type OpenstackVolumeResourceModel struct {
	UUID                        types.String   `tfsdk:"id"`
	AccessUrl                   types.String   `tfsdk:"access_url"`
	Action                      types.String   `tfsdk:"action"`
	AvailabilityZone            types.String   `tfsdk:"availability_zone"`
	AvailabilityZoneName        types.String   `tfsdk:"availability_zone_name"`
	BackendId                   types.String   `tfsdk:"backend_id"`
	Bootable                    types.Bool     `tfsdk:"bootable"`
	Created                     types.String   `tfsdk:"created"`
	Customer                    types.String   `tfsdk:"customer"`
	CustomerAbbreviation        types.String   `tfsdk:"customer_abbreviation"`
	CustomerName                types.String   `tfsdk:"customer_name"`
	CustomerNativeName          types.String   `tfsdk:"customer_native_name"`
	CustomerUuid                types.String   `tfsdk:"customer_uuid"`
	Description                 types.String   `tfsdk:"description"`
	Device                      types.String   `tfsdk:"device"`
	ErrorMessage                types.String   `tfsdk:"error_message"`
	ErrorTraceback              types.String   `tfsdk:"error_traceback"`
	ExtendEnabled               types.Bool     `tfsdk:"extend_enabled"`
	Image                       types.String   `tfsdk:"image"`
	ImageMetadata               types.String   `tfsdk:"image_metadata"`
	ImageName                   types.String   `tfsdk:"image_name"`
	Instance                    types.String   `tfsdk:"instance"`
	InstanceMarketplaceUuid     types.String   `tfsdk:"instance_marketplace_uuid"`
	InstanceName                types.String   `tfsdk:"instance_name"`
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
	ResourceType                types.String   `tfsdk:"resource_type"`
	RuntimeState                types.String   `tfsdk:"runtime_state"`
	ServiceName                 types.String   `tfsdk:"service_name"`
	ServiceSettings             types.String   `tfsdk:"service_settings"`
	ServiceSettingsErrorMessage types.String   `tfsdk:"service_settings_error_message"`
	ServiceSettingsState        types.String   `tfsdk:"service_settings_state"`
	ServiceSettingsUuid         types.String   `tfsdk:"service_settings_uuid"`
	Size                        types.Int64    `tfsdk:"size"`
	SourceSnapshot              types.String   `tfsdk:"source_snapshot"`
	State                       types.String   `tfsdk:"state"`
	Tenant                      types.String   `tfsdk:"tenant"`
	TenantUuid                  types.String   `tfsdk:"tenant_uuid"`
	Type                        types.String   `tfsdk:"type"`
	TypeName                    types.String   `tfsdk:"type_name"`
	Url                         types.String   `tfsdk:"url"`
	Timeouts                    timeouts.Value `tfsdk:"timeouts"`
}

func (r *OpenstackVolumeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_volume"
}

func (r *OpenstackVolumeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Openstack Volume resource",

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
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Availability zone where this volume is located",
			},
			"availability_zone_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the availability zone",
			},
			"backend_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Volume ID in the OpenStack backend",
			},
			"bootable": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Indicates if this volume can be used to boot an instance",
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
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Description of the resource",
			},
			"device": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of volume as instance device e.g. /dev/vdb.",
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
			"extend_enabled": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Extend enabled",
			},
			"image": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Image that this volume was created from, if any",
			},
			"image_metadata": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Metadata of the image this volume was created from",
			},
			"image_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the image this volume was created from",
			},
			"instance": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Instance that this volume is attached to, if any",
			},
			"instance_marketplace_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the instance marketplace",
			},
			"instance_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the instance",
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
			"modified": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Modified",
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Name of the resource",
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
				MarkdownDescription: "Service settings",
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
			"size": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Size in MiB",
			},
			"source_snapshot": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Snapshot that this volume was created from, if any",
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
			"tenant_uuid": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the tenant",
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
				MarkdownDescription: "Name of the type",
			},
			"url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Url",
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

func (r *OpenstackVolumeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *OpenstackVolumeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackVolumeResourceModel
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
	if !data.Image.IsNull() {
		attributes["image"] = data.Image.ValueString()
	}
	if !data.Name.IsNull() {
		attributes["name"] = data.Name.ValueString()
	}
	if !data.Size.IsNull() {
		attributes["size"] = data.Size.ValueInt64()
	}
	if !data.Type.IsNull() {
		attributes["type"] = data.Type.ValueString()
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
						var apiResp OpenstackVolumeApiResponse
						retrievePath := strings.Replace("/api/openstack-volumes/{uuid}/", "{uuid}", pluginUUID, 1)
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
	var apiResp OpenstackVolumeApiResponse
	err = r.client.GetByUUID(ctx, "/api/openstack-volumes/{uuid}/", data.UUID.ValueString(), &apiResp)
	if err != nil {
		resp.Diagnostics.AddError("Failed to Read Resource", err.Error())
		return
	}
	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackVolumeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpenstackVolumeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource

	retrievePath := strings.Replace("/api/openstack-volumes/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	var apiResp OpenstackVolumeApiResponse
	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &apiResp)
	if err != nil {
		if client.IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Unable to Read Openstack Volume",
			"An error occurred while reading the Openstack Volume: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackVolumeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpenstackVolumeResourceModel
	var state OpenstackVolumeResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Phase 1: Standard PATCH (Simple fields)
	patchPayload := map[string]interface{}{}
	if !data.Bootable.IsNull() && !data.Bootable.Equal(state.Bootable) {
		patchPayload["bootable"] = data.Bootable.ValueBool()
	}
	if !data.Description.IsNull() && !data.Description.Equal(state.Description) {
		patchPayload["description"] = data.Description.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.Equal(state.Name) {
		patchPayload["name"] = data.Name.ValueString()
	}

	if len(patchPayload) > 0 {
		var result map[string]interface{}
		err := r.client.Update(ctx, "/api/openstack-volumes/{uuid}/", data.UUID.ValueString(), patchPayload, &result)
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
	if !data.Size.Equal(state.Size) {

		// Convert Terraform value to API payload
		actionPayloadExtend := map[string]interface{}{
			"size": data.Size.ValueInt64(),
		}
		actionUrlExtend := strings.Replace("/api/openstack-volumes/{uuid}/extend/", "{uuid}", data.UUID.ValueString(), 1)
		var actionResultExtend map[string]interface{}
		if err := r.client.Post(ctx, actionUrlExtend, actionPayloadExtend, &actionResultExtend); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: extend", err.Error())
			return
		}
	}
	if !data.Type.Equal(state.Type) {

		// Convert Terraform value to API payload
		actionPayloadRetype := map[string]interface{}{
			"type": data.Type.ValueString(),
		}
		actionUrlRetype := strings.Replace("/api/openstack-volumes/{uuid}/retype/", "{uuid}", data.UUID.ValueString(), 1)
		var actionResultRetype map[string]interface{}
		if err := r.client.Post(ctx, actionUrlRetype, actionPayloadRetype, &actionResultRetype); err != nil {
			resp.Diagnostics.AddError("RPC Action Failed: retype", err.Error())
			return
		}
	}

	// Fetch updated state
	// Call Waldur API to read resource
	var apiResp OpenstackVolumeApiResponse

	retrievePath := strings.Replace("/api/openstack-volumes/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &apiResp)
	if err != nil {
		if client.IsNotFoundError(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Failed to Read Resource After Update", err.Error())
		return
	}
	resp.Diagnostics.Append(r.mapResponseToModel(ctx, apiResp, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackVolumeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackVolumeResourceModel
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

func (r *OpenstackVolumeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *OpenstackVolumeResource) mapResponseToModel(ctx context.Context, apiResp OpenstackVolumeApiResponse, model *OpenstackVolumeResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	model.UUID = types.StringPointerValue(apiResp.UUID)
	model.AccessUrl = types.StringPointerValue(apiResp.AccessUrl)
	model.Action = types.StringPointerValue(apiResp.Action)
	model.AvailabilityZone = types.StringPointerValue(apiResp.AvailabilityZone)
	model.AvailabilityZoneName = types.StringPointerValue(apiResp.AvailabilityZoneName)
	model.BackendId = types.StringPointerValue(apiResp.BackendId)
	model.Bootable = types.BoolPointerValue(apiResp.Bootable)
	model.Created = types.StringPointerValue(apiResp.Created)
	model.Customer = types.StringPointerValue(apiResp.Customer)
	model.CustomerAbbreviation = types.StringPointerValue(apiResp.CustomerAbbreviation)
	model.CustomerName = types.StringPointerValue(apiResp.CustomerName)
	model.CustomerNativeName = types.StringPointerValue(apiResp.CustomerNativeName)
	model.CustomerUuid = types.StringPointerValue(apiResp.CustomerUuid)
	model.Description = types.StringPointerValue(apiResp.Description)
	model.Device = types.StringPointerValue(apiResp.Device)
	model.ErrorMessage = types.StringPointerValue(apiResp.ErrorMessage)
	model.ErrorTraceback = types.StringPointerValue(apiResp.ErrorTraceback)
	model.ExtendEnabled = types.BoolPointerValue(apiResp.ExtendEnabled)
	model.Image = types.StringPointerValue(apiResp.Image)
	model.ImageMetadata = types.StringPointerValue(apiResp.ImageMetadata)
	model.ImageName = types.StringPointerValue(apiResp.ImageName)
	model.Instance = types.StringPointerValue(apiResp.Instance)
	model.InstanceMarketplaceUuid = types.StringPointerValue(apiResp.InstanceMarketplaceUuid)
	model.InstanceName = types.StringPointerValue(apiResp.InstanceName)
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
	model.Name = types.StringPointerValue(apiResp.Name)
	model.Offering = types.StringPointerValue(apiResp.Offering)
	model.Project = types.StringPointerValue(apiResp.Project)
	model.ProjectName = types.StringPointerValue(apiResp.ProjectName)
	model.ProjectUuid = types.StringPointerValue(apiResp.ProjectUuid)
	model.ResourceType = types.StringPointerValue(apiResp.ResourceType)
	model.RuntimeState = types.StringPointerValue(apiResp.RuntimeState)
	model.ServiceName = types.StringPointerValue(apiResp.ServiceName)
	model.ServiceSettings = types.StringPointerValue(apiResp.ServiceSettings)
	model.ServiceSettingsErrorMessage = types.StringPointerValue(apiResp.ServiceSettingsErrorMessage)
	model.ServiceSettingsState = types.StringPointerValue(apiResp.ServiceSettingsState)
	model.ServiceSettingsUuid = types.StringPointerValue(apiResp.ServiceSettingsUuid)
	model.Size = types.Int64PointerValue(apiResp.Size)
	model.SourceSnapshot = types.StringPointerValue(apiResp.SourceSnapshot)
	model.State = types.StringPointerValue(apiResp.State)
	model.Tenant = types.StringPointerValue(apiResp.Tenant)
	model.TenantUuid = types.StringPointerValue(apiResp.TenantUuid)
	model.Type = types.StringPointerValue(apiResp.Type)
	model.TypeName = types.StringPointerValue(apiResp.TypeName)
	model.Url = types.StringPointerValue(apiResp.Url)

	return diags
}
