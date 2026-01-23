package marketplace

import (
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	pkg_offering "github.com/waldur/terraform-provider-waldur/services/marketplace/offering"
	pkg_order "github.com/waldur/terraform-provider-waldur/services/marketplace/order"
	pkg_resource "github.com/waldur/terraform-provider-waldur/services/marketplace/resource"
)

func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		pkg_order.NewMarketplaceOrderResource,
		pkg_offering.NewMarketplaceOfferingResource,
		pkg_resource.NewMarketplaceResourceResource,
	}
}

func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		pkg_order.NewMarketplaceOrderDataSource,
		pkg_offering.NewMarketplaceOfferingDataSource,
		pkg_resource.NewMarketplaceResourceDataSource,
	}
}

func GetActions() []func() action.Action {
	return []func() action.Action{
		pkg_resource.NewMarketplaceResourcePullAction,
		pkg_resource.NewMarketplaceResourceTerminateAction,
		pkg_resource.NewMarketplaceResourceUnlinkAction,
	}
}

func GetListResources() []func() list.ListResource {
	return []func() list.ListResource{
		pkg_order.NewMarketplaceOrderList,
		pkg_offering.NewMarketplaceOfferingList,
		pkg_resource.NewMarketplaceResourceList,
	}
}
