package resources

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
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
var _ resource.Resource = &OpenstackVolumeResource{}
var _ resource.ResourceWithImportState = &OpenstackVolumeResource{}

func NewOpenstackVolumeResource() resource.Resource {
	return &OpenstackVolumeResource{}
}

// OpenstackVolumeResource defines the resource implementation.
type OpenstackVolumeResource struct {
	client *client.Client
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
		MarkdownDescription: "OpenstackVolume resource",

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
			"action": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"availability_zone": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Availability zone where this volume is located",
			},
			"availability_zone_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"backend_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Volume ID in the OpenStack backend",
			},
			"bootable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Indicates if this volume can be used to boot an instance",
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
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: " ",
			},
			"device": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of volume as instance device e.g. /dev/vdb.",
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"error_traceback": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"extend_enabled": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"image": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Image that this volume was created from, if any",
			},
			"image_metadata": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Metadata of the image this volume was created from",
			},
			"image_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the image this volume was created from",
			},
			"instance": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Instance that this volume is attached to, if any",
			},
			"instance_marketplace_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"instance_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
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
			"resource_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"runtime_state": schema.StringAttribute{
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
			"size": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Size in MiB",
			},
			"source_snapshot": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Snapshot that this volume was created from, if any",
			},
			"state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"tenant": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"tenant_uuid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Type of the volume (e.g. SSD, HDD)",
			},
			"type_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: " ",
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
						var pluginRes map[string]interface{}
						retrievePath := strings.Replace("/api/openstack-volumes/{uuid}/", "{uuid}", pluginUUID, 1)
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
							if val, ok := sourceMap["bootable"]; ok && val != nil {
								if b, ok := val.(bool); ok {
									data.Bootable = types.BoolValue(b)
								}
							} else {
								if data.Bootable.IsUnknown() {
									data.Bootable = types.BoolNull()
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
							if val, ok := sourceMap["description"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.Description = types.StringValue(str)
								}
							} else {
								if data.Description.IsUnknown() {
									data.Description = types.StringNull()
								}
							}
							if val, ok := sourceMap["device"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.Device = types.StringValue(str)
								}
							} else {
								if data.Device.IsUnknown() {
									data.Device = types.StringNull()
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
							if val, ok := sourceMap["extend_enabled"]; ok && val != nil {
								if b, ok := val.(bool); ok {
									data.ExtendEnabled = types.BoolValue(b)
								}
							} else {
								if data.ExtendEnabled.IsUnknown() {
									data.ExtendEnabled = types.BoolNull()
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
							if val, ok := sourceMap["image_metadata"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.ImageMetadata = types.StringValue(str)
								}
							} else {
								if data.ImageMetadata.IsUnknown() {
									data.ImageMetadata = types.StringNull()
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
							if val, ok := sourceMap["instance"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.Instance = types.StringValue(str)
								}
							} else {
								if data.Instance.IsUnknown() {
									data.Instance = types.StringNull()
								}
							}
							if val, ok := sourceMap["instance_marketplace_uuid"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.InstanceMarketplaceUuid = types.StringValue(str)
								}
							} else {
								if data.InstanceMarketplaceUuid.IsUnknown() {
									data.InstanceMarketplaceUuid = types.StringNull()
								}
							}
							if val, ok := sourceMap["instance_name"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.InstanceName = types.StringValue(str)
								}
							} else {
								if data.InstanceName.IsUnknown() {
									data.InstanceName = types.StringNull()
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
							if val, ok := sourceMap["size"]; ok && val != nil {
								if num, ok := val.(float64); ok {
									data.Size = types.Int64Value(int64(num))
								}
							} else {
								if data.Size.IsUnknown() {
									data.Size = types.Int64Null()
								}
							}
							if val, ok := sourceMap["source_snapshot"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.SourceSnapshot = types.StringValue(str)
								}
							} else {
								if data.SourceSnapshot.IsUnknown() {
									data.SourceSnapshot = types.StringNull()
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
							if val, ok := sourceMap["type"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.Type = types.StringValue(str)
								}
							} else {
								if data.Type.IsUnknown() {
									data.Type = types.StringNull()
								}
							}
							if val, ok := sourceMap["type_name"]; ok && val != nil {
								if str, ok := val.(string); ok {
									data.TypeName = types.StringValue(str)
								}
							} else {
								if data.TypeName.IsUnknown() {
									data.TypeName = types.StringNull()
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
	err = r.client.GetByUUID(ctx, "/api/openstack-volumes/{uuid}/", data.UUID.ValueString(), &finalState)
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
	if val, ok := sourceMap["bootable"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Bootable = types.BoolValue(b)
		}
	} else {
		if data.Bootable.IsUnknown() {
			data.Bootable = types.BoolNull()
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
	if val, ok := sourceMap["description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
		}
	} else {
		if data.Description.IsUnknown() {
			data.Description = types.StringNull()
		}
	}
	if val, ok := sourceMap["device"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Device = types.StringValue(str)
		}
	} else {
		if data.Device.IsUnknown() {
			data.Device = types.StringNull()
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
	if val, ok := sourceMap["extend_enabled"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.ExtendEnabled = types.BoolValue(b)
		}
	} else {
		if data.ExtendEnabled.IsUnknown() {
			data.ExtendEnabled = types.BoolNull()
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
	if val, ok := sourceMap["image_metadata"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ImageMetadata = types.StringValue(str)
		}
	} else {
		if data.ImageMetadata.IsUnknown() {
			data.ImageMetadata = types.StringNull()
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
	if val, ok := sourceMap["instance"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Instance = types.StringValue(str)
		}
	} else {
		if data.Instance.IsUnknown() {
			data.Instance = types.StringNull()
		}
	}
	if val, ok := sourceMap["instance_marketplace_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.InstanceMarketplaceUuid = types.StringValue(str)
		}
	} else {
		if data.InstanceMarketplaceUuid.IsUnknown() {
			data.InstanceMarketplaceUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["instance_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.InstanceName = types.StringValue(str)
		}
	} else {
		if data.InstanceName.IsUnknown() {
			data.InstanceName = types.StringNull()
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
	if val, ok := sourceMap["size"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Size = types.Int64Value(int64(num))
		}
	} else {
		if data.Size.IsUnknown() {
			data.Size = types.Int64Null()
		}
	}
	if val, ok := sourceMap["source_snapshot"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.SourceSnapshot = types.StringValue(str)
		}
	} else {
		if data.SourceSnapshot.IsUnknown() {
			data.SourceSnapshot = types.StringNull()
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
	if val, ok := sourceMap["type"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Type = types.StringValue(str)
		}
	} else {
		if data.Type.IsUnknown() {
			data.Type = types.StringNull()
		}
	}
	if val, ok := sourceMap["type_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.TypeName = types.StringValue(str)
		}
	} else {
		if data.TypeName.IsUnknown() {
			data.TypeName = types.StringNull()
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

	// Map filter parameters from response if available

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
	var result map[string]interface{}

	retrievePath := strings.Replace("/api/openstack-volumes/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read OpenstackVolume",
			"An error occurred while reading the openstack_volume: "+err.Error(),
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
	if val, ok := sourceMap["bootable"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Bootable = types.BoolValue(b)
		}
	} else {
		if data.Bootable.IsUnknown() {
			data.Bootable = types.BoolNull()
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
	if val, ok := sourceMap["description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
		}
	} else {
		if data.Description.IsUnknown() {
			data.Description = types.StringNull()
		}
	}
	if val, ok := sourceMap["device"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Device = types.StringValue(str)
		}
	} else {
		if data.Device.IsUnknown() {
			data.Device = types.StringNull()
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
	if val, ok := sourceMap["extend_enabled"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.ExtendEnabled = types.BoolValue(b)
		}
	} else {
		if data.ExtendEnabled.IsUnknown() {
			data.ExtendEnabled = types.BoolNull()
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
	if val, ok := sourceMap["image_metadata"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ImageMetadata = types.StringValue(str)
		}
	} else {
		if data.ImageMetadata.IsUnknown() {
			data.ImageMetadata = types.StringNull()
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
	if val, ok := sourceMap["instance"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Instance = types.StringValue(str)
		}
	} else {
		if data.Instance.IsUnknown() {
			data.Instance = types.StringNull()
		}
	}
	if val, ok := sourceMap["instance_marketplace_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.InstanceMarketplaceUuid = types.StringValue(str)
		}
	} else {
		if data.InstanceMarketplaceUuid.IsUnknown() {
			data.InstanceMarketplaceUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["instance_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.InstanceName = types.StringValue(str)
		}
	} else {
		if data.InstanceName.IsUnknown() {
			data.InstanceName = types.StringNull()
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
	if val, ok := sourceMap["size"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Size = types.Int64Value(int64(num))
		}
	} else {
		if data.Size.IsUnknown() {
			data.Size = types.Int64Null()
		}
	}
	if val, ok := sourceMap["source_snapshot"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.SourceSnapshot = types.StringValue(str)
		}
	} else {
		if data.SourceSnapshot.IsUnknown() {
			data.SourceSnapshot = types.StringNull()
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
	if val, ok := sourceMap["type"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Type = types.StringValue(str)
		}
	} else {
		if data.Type.IsUnknown() {
			data.Type = types.StringNull()
		}
	}
	if val, ok := sourceMap["type_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.TypeName = types.StringValue(str)
		}
	} else {
		if data.TypeName.IsUnknown() {
			data.TypeName = types.StringNull()
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

	// Map filter parameters from response if available

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

	data.UUID = state.UUID

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
	var result map[string]interface{}

	retrievePath := strings.Replace("/api/openstack-volumes/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)

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
	if val, ok := sourceMap["bootable"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Bootable = types.BoolValue(b)
		}
	} else {
		if data.Bootable.IsUnknown() {
			data.Bootable = types.BoolNull()
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
	if val, ok := sourceMap["description"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
		}
	} else {
		if data.Description.IsUnknown() {
			data.Description = types.StringNull()
		}
	}
	if val, ok := sourceMap["device"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Device = types.StringValue(str)
		}
	} else {
		if data.Device.IsUnknown() {
			data.Device = types.StringNull()
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
	if val, ok := sourceMap["extend_enabled"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.ExtendEnabled = types.BoolValue(b)
		}
	} else {
		if data.ExtendEnabled.IsUnknown() {
			data.ExtendEnabled = types.BoolNull()
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
	if val, ok := sourceMap["image_metadata"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.ImageMetadata = types.StringValue(str)
		}
	} else {
		if data.ImageMetadata.IsUnknown() {
			data.ImageMetadata = types.StringNull()
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
	if val, ok := sourceMap["instance"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Instance = types.StringValue(str)
		}
	} else {
		if data.Instance.IsUnknown() {
			data.Instance = types.StringNull()
		}
	}
	if val, ok := sourceMap["instance_marketplace_uuid"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.InstanceMarketplaceUuid = types.StringValue(str)
		}
	} else {
		if data.InstanceMarketplaceUuid.IsUnknown() {
			data.InstanceMarketplaceUuid = types.StringNull()
		}
	}
	if val, ok := sourceMap["instance_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.InstanceName = types.StringValue(str)
		}
	} else {
		if data.InstanceName.IsUnknown() {
			data.InstanceName = types.StringNull()
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
	if val, ok := sourceMap["size"]; ok && val != nil {
		if num, ok := val.(float64); ok {
			data.Size = types.Int64Value(int64(num))
		}
	} else {
		if data.Size.IsUnknown() {
			data.Size = types.Int64Null()
		}
	}
	if val, ok := sourceMap["source_snapshot"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.SourceSnapshot = types.StringValue(str)
		}
	} else {
		if data.SourceSnapshot.IsUnknown() {
			data.SourceSnapshot = types.StringNull()
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
	if val, ok := sourceMap["type"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.Type = types.StringValue(str)
		}
	} else {
		if data.Type.IsUnknown() {
			data.Type = types.StringNull()
		}
	}
	if val, ok := sourceMap["type_name"]; ok && val != nil {
		if str, ok := val.(string); ok {
			data.TypeName = types.StringValue(str)
		}
	} else {
		if data.TypeName.IsUnknown() {
			data.TypeName = types.StringNull()
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

	// Map filter parameters from response if available

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
