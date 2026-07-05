package identity

import (
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	pkg_bridge_link "github.com/waldur/terraform-provider-waldur/services/identity/bridge_link"
)

func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		pkg_bridge_link.NewIdentityBridgeLinkResource,
	}
}

func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func GetActions() []func() action.Action {
	return []func() action.Action{}
}

func GetListResources() []func() list.ListResource {
	return []func() list.ListResource{}
}
