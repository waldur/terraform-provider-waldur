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
var _ action.Action = &OpenstackSubnetDisconnectAction{}
var _ action.ActionWithConfigure = &OpenstackSubnetDisconnectAction{}

type OpenstackSubnetDisconnectAction struct {
	client *client.Client
}

func NewOpenstackSubnetDisconnectAction() action.Action {
	return &OpenstackSubnetDisconnectAction{}
}

func (a *OpenstackSubnetDisconnectAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_subnet_disconnect"
}
func (a *OpenstackSubnetDisconnectAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Perform disconnect action on openstack_subnet",
		Attributes: map[string]schema.Attribute{
			"uuid": schema.StringAttribute{
				Description: "The UUID of the openstack_subnet",
				Required:    true,
			},
		},
	}
}

func (a *OpenstackSubnetDisconnectAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

type OpenstackSubnetDisconnectModel struct {
	Uuid types.String `tfsdk:"uuid"`
}

func (a *OpenstackSubnetDisconnectAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data OpenstackSubnetDisconnectModel
	// Read validation data
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	uuid := data.Uuid.ValueString()
	// Perform the action
	path := "/api/openstack-subnets/{uuid}/disconnect/"
	path = strings.Replace(path, "{uuid}", uuid, 1)

	err := a.client.Post(ctx, path, nil, nil)

	if err != nil {
		resp.Diagnostics.AddError(
			"Action Failed",
			fmt.Sprintf("Failed to perform disconnect on %s: %s", uuid, err),
		)
		return
	}
}
