package port

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
var _ action.Action = &OpenstackPortDisablePortSecurityAction{}
var _ action.ActionWithConfigure = &OpenstackPortDisablePortSecurityAction{}

type OpenstackPortDisablePortSecurityModel struct {
	Uuid    types.String `tfsdk:"uuid"`
	Timeout types.String `tfsdk:"timeout"`
}

type OpenstackPortDisablePortSecurityAction struct {
	client *Client
}

func NewOpenstackPortDisablePortSecurityAction() action.Action {
	return &OpenstackPortDisablePortSecurityAction{}
}

func (a *OpenstackPortDisablePortSecurityAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_port_disable_port_security"
}

func (a *OpenstackPortDisablePortSecurityAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Perform disable_port_security action on openstack port",
		Attributes: map[string]schema.Attribute{
			"uuid": schema.StringAttribute{
				Description: "The UUID of the openstack port",
				Required:    true,
			},
			"timeout": schema.StringAttribute{
				Description: "Timeout for the action execution and state stabilization (e.g. '10m').",
				Optional:    true,
			},
		},
	}
}

func (a *OpenstackPortDisablePortSecurityAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

func (a *OpenstackPortDisablePortSecurityAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data OpenstackPortDisablePortSecurityModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	uuid := data.Uuid.ValueString()
	err := a.client.OpenstackPortDisablePortSecurity(ctx, uuid)

	if err != nil {
		resp.Diagnostics.AddError(
			"Action Failed",
			fmt.Sprintf("Failed to perform disable_port_security on %s: %s", uuid, err),
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

	_, err = common.WaitForResource(ctx, func(ctx context.Context) (*OpenstackPortResponse, error) {
		return a.client.GetOpenstackPort(ctx, uuid)
	}, timeout)
	if err != nil {
		resp.Diagnostics.AddWarning("Resource state check failed", err.Error())
	}
}
