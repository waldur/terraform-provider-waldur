package core

import (
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func GetResources() []func() resource.Resource {
	return []func() resource.Resource{}
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
