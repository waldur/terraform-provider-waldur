package openstack

import (
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	pkg_floating_ip "github.com/waldur/terraform-provider-waldur/services/openstack/floating_ip"
	pkg_instance "github.com/waldur/terraform-provider-waldur/services/openstack/instance"
	pkg_network "github.com/waldur/terraform-provider-waldur/services/openstack/network"
	pkg_network_rbac_policy "github.com/waldur/terraform-provider-waldur/services/openstack/network_rbac_policy"
	pkg_port "github.com/waldur/terraform-provider-waldur/services/openstack/port"
	pkg_security_group "github.com/waldur/terraform-provider-waldur/services/openstack/security_group"
	pkg_server_group "github.com/waldur/terraform-provider-waldur/services/openstack/server_group"
	pkg_subnet "github.com/waldur/terraform-provider-waldur/services/openstack/subnet"
	pkg_tenant "github.com/waldur/terraform-provider-waldur/services/openstack/tenant"
	pkg_volume "github.com/waldur/terraform-provider-waldur/services/openstack/volume"
	pkg_volume_attachment "github.com/waldur/terraform-provider-waldur/services/openstack/volume_attachment"
)

func GetResources() []func() resource.Resource {
	return []func() resource.Resource{
		pkg_tenant.NewOpenstackTenantResource,
		pkg_volume.NewOpenstackVolumeResource,
		pkg_volume_attachment.NewOpenstackVolumeAttachmentResource,
		pkg_instance.NewOpenstackInstanceResource,
		pkg_security_group.NewOpenstackSecurityGroupResource,
		pkg_network.NewOpenstackNetworkResource,
		pkg_server_group.NewOpenstackServerGroupResource,
		pkg_subnet.NewOpenstackSubnetResource,
		pkg_network_rbac_policy.NewOpenstackNetworkRbacPolicyResource,
		pkg_floating_ip.NewOpenstackFloatingIpResource,
		pkg_port.NewOpenstackPortResource,
	}
}

func GetDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{
		pkg_tenant.NewOpenstackTenantDataSource,
		pkg_volume.NewOpenstackVolumeDataSource,
		pkg_instance.NewOpenstackInstanceDataSource,
		pkg_security_group.NewOpenstackSecurityGroupDataSource,
		pkg_network.NewOpenstackNetworkDataSource,
		pkg_server_group.NewOpenstackServerGroupDataSource,
		pkg_subnet.NewOpenstackSubnetDataSource,
		pkg_network_rbac_policy.NewOpenstackNetworkRbacPolicyDataSource,
		pkg_floating_ip.NewOpenstackFloatingIpDataSource,
		pkg_port.NewOpenstackPortDataSource,
	}
}

func GetActions() []func() action.Action {
	return []func() action.Action{
		pkg_tenant.NewOpenstackTenantPullAction,
		pkg_tenant.NewOpenstackTenantUnlinkAction,
		pkg_volume.NewOpenstackVolumePullAction,
		pkg_volume.NewOpenstackVolumeUnlinkAction,
		pkg_instance.NewOpenstackInstanceStartAction,
		pkg_instance.NewOpenstackInstanceStopAction,
		pkg_instance.NewOpenstackInstanceRestartAction,
		pkg_instance.NewOpenstackInstancePullAction,
		pkg_instance.NewOpenstackInstanceUnlinkAction,
		pkg_network.NewOpenstackNetworkPullAction,
		pkg_network.NewOpenstackNetworkUnlinkAction,
		pkg_subnet.NewOpenstackSubnetConnectAction,
		pkg_subnet.NewOpenstackSubnetDisconnectAction,
		pkg_subnet.NewOpenstackSubnetPullAction,
		pkg_subnet.NewOpenstackSubnetUnlinkAction,
		pkg_port.NewOpenstackPortEnablePortAction,
		pkg_port.NewOpenstackPortDisablePortAction,
		pkg_port.NewOpenstackPortEnablePortSecurityAction,
		pkg_port.NewOpenstackPortDisablePortSecurityAction,
		pkg_port.NewOpenstackPortPullAction,
		pkg_port.NewOpenstackPortUnlinkAction,
	}
}

func GetListResources() []func() list.ListResource {
	return []func() list.ListResource{
		pkg_tenant.NewOpenstackTenantList,
		pkg_volume.NewOpenstackVolumeList,
		pkg_volume_attachment.NewOpenstackVolumeAttachmentList,
		pkg_instance.NewOpenstackInstanceList,
		pkg_security_group.NewOpenstackSecurityGroupList,
		pkg_network.NewOpenstackNetworkList,
		pkg_server_group.NewOpenstackServerGroupList,
		pkg_subnet.NewOpenstackSubnetList,
		pkg_network_rbac_policy.NewOpenstackNetworkRbacPolicyList,
		pkg_floating_ip.NewOpenstackFloatingIpList,
		pkg_port.NewOpenstackPortList,
	}
}
