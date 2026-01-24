package volume

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

// Ensure implementation satisfies interfaces.
var _ action.Action = &OpenstackVolumePullAction{}
var _ action.ActionWithConfigure = &OpenstackVolumePullAction{}

type OpenstackVolumePullModel struct {
	Uuid    types.String `tfsdk:"uuid"`
	Timeout types.String `tfsdk:"timeout"`
}

type OpenstackVolumePullAction struct {
	client *Client
}

func NewOpenstackVolumePullAction() action.Action {
	return &OpenstackVolumePullAction{}
}

func (a *OpenstackVolumePullAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_volume_pull"
}

func (a *OpenstackVolumePullAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Perform pull action on openstack volume",
		Attributes: map[string]schema.Attribute{
			"uuid": schema.StringAttribute{
				Description: "The UUID of the openstack volume",
				Required:    true,
			},
			"timeout": schema.StringAttribute{
				Description: "Timeout for the action execution and state stabilization (e.g. '10m').",
				Optional:    true,
			},
		},
	}
}

func (a *OpenstackVolumePullAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	clientRaw, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Action Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T", req.ProviderData),
		)
		return
	}

	a.client = NewClient(clientRaw)
}

func (a *OpenstackVolumePullAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data OpenstackVolumePullModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	uuid := data.Uuid.ValueString()
	err := a.client.OpenstackVolumePull(ctx, uuid)

	if err != nil {
		resp.Diagnostics.AddError(
			"Action Failed",
			fmt.Sprintf("Failed to perform pull on %s: %s", uuid, err),
		)
		return
	}

	// Wait for resource to stabilize
	timeout := 10 * time.Minute
	if !data.Timeout.IsNull() {
		var err error
		timeout, err = time.ParseDuration(data.Timeout.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Invalid timeout", "Failed to parse timeout: "+err.Error())
			return
		}
	}

	_, err = common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackVolumeResponse, error) {
		return a.client.GetOpenstackVolume(ctx, uuid)
	}, timeout)
	if err != nil {
		resp.Diagnostics.AddWarning("Resource state check failed", err.Error())
	}
}
