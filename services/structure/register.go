package structure

import (
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	pkg_customer "github.com/waldur/terraform-provider-waldur/services/structure/customer"
	pkg_project "github.com/waldur/terraform-provider-waldur/services/structure/project"
)

func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		pkg_project.NewStructureProjectResource,
		pkg_customer.NewStructureCustomerResource,
	}
}

func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		pkg_project.NewStructureProjectDataSource,
		pkg_customer.NewStructureCustomerDataSource,
	}
}

func GetActions() []func() action.Action {
	return []func() action.Action{}
}

func GetListResources() []func() list.ListResource {
	return []func() list.ListResource{
		pkg_project.NewStructureProjectList,
		pkg_customer.NewStructureCustomerList,
	}
}
