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
var _ action.Action = &OpenstackVolumeUnlinkAction{}
var _ action.ActionWithConfigure = &OpenstackVolumeUnlinkAction{}

type OpenstackVolumeUnlinkAction struct {
	client *client.Client
}

func NewOpenstackVolumeUnlinkAction() action.Action {
	return &OpenstackVolumeUnlinkAction{}
}

func (a *OpenstackVolumeUnlinkAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_volume_unlink"
}
func (a *OpenstackVolumeUnlinkAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Perform unlink action on openstack_volume",
		Attributes: map[string]schema.Attribute{
			"uuid": schema.StringAttribute{
				Description: "The UUID of the openstack_volume",
				Required:    true,
			},
		},
	}
}

func (a *OpenstackVolumeUnlinkAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

type OpenstackVolumeUnlinkModel struct {
	Uuid types.String `tfsdk:"uuid"`
}

func (a *OpenstackVolumeUnlinkAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data OpenstackVolumeUnlinkModel
	// Read validation data
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	uuid := data.Uuid.ValueString()
	// Perform the action
	path := "/api/openstack-volumes/{uuid}/unlink/"
	path = strings.Replace(path, "{uuid}", uuid, 1)

	err := a.client.Post(ctx, path, nil, nil)

	if err != nil {
		resp.Diagnostics.AddError(
			"Action Failed",
			fmt.Sprintf("Failed to perform unlink on %s: %s", uuid, err),
		)
		return
	}
}
