package actions

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// Ensure implementation satisfies interfaces.
var _ action.Action = &OpenstackPortEnablePortSecurityAction{}
var _ action.ActionWithConfigure = &OpenstackPortEnablePortSecurityAction{}

type OpenstackPortEnablePortSecurityAction struct {
	client *client.Client
}

func NewOpenstackPortEnablePortSecurityAction() action.Action {
	return &OpenstackPortEnablePortSecurityAction{}
}

func (a *OpenstackPortEnablePortSecurityAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_port_enable_port_security"
}
func (a *OpenstackPortEnablePortSecurityAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Perform enable_port_security action on openstack port",
		Attributes: map[string]schema.Attribute{
			"uuid": schema.StringAttribute{
				Description: "The UUID of the openstack port",
				Required:    true,
			},
		},
	}
}

func (a *OpenstackPortEnablePortSecurityAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Action Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T", req.ProviderData),
		)
		return
	}

	a.client = client
}

type OpenstackPortEnablePortSecurityModel struct {
	Uuid types.String `tfsdk:"uuid"`
}

func (a *OpenstackPortEnablePortSecurityAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data OpenstackPortEnablePortSecurityModel
	// Read validation data
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	uuid := data.Uuid.ValueString()
	// Perform the action
	path := "/api/openstack-ports/{uuid}/enable_port_security/"
	path = strings.Replace(path, "{uuid}", uuid, 1)

	err := a.client.Post(ctx, path, nil, nil)

	if err != nil {
		resp.Diagnostics.AddError(
			"Action Failed",
			fmt.Sprintf("Failed to perform enable_port_security on %s: %s", uuid, err),
		)
		return
	}
}
