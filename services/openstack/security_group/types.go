package security_group

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackSecurityGroupCreateRequest struct {
	Description *string                                          `json:"description,omitempty" tfsdk:"description"`
	Name        *string                                          `json:"name" tfsdk:"name"`
	Rules       []common.OpenStackSecurityGroupRuleCreateRequest `json:"rules" tfsdk:"rules"`
}

type OpenstackSecurityGroupUpdateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type OpenstackSecurityGroupSetRulesActionRequest struct {
	Rules []common.OpenStackSecurityGroupRuleCreateRequest `json:"rules"`
}

type OpenstackSecurityGroupResponse struct {
	UUID *string `json:"uuid"`

	BackendId               *string                                          `json:"backend_id" tfsdk:"backend_id"`
	Created                 *string                                          `json:"created" tfsdk:"created"`
	Customer                *string                                          `json:"customer" tfsdk:"customer"`
	Description             *string                                          `json:"description" tfsdk:"description"`
	ErrorMessage            *string                                          `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback          *string                                          `json:"error_traceback" tfsdk:"error_traceback"`
	MarketplaceResourceUuid *string                                          `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                *string                                          `json:"modified" tfsdk:"modified"`
	Name                    *string                                          `json:"name" tfsdk:"name"`
	Project                 *string                                          `json:"project" tfsdk:"project"`
	ResourceType            *string                                          `json:"resource_type" tfsdk:"resource_type"`
	Rules                   []common.OpenStackSecurityGroupRuleCreateRequest `json:"rules" tfsdk:"rules"`
	State                   *string                                          `json:"state" tfsdk:"state"`
	Tenant                  *string                                          `json:"tenant" tfsdk:"tenant"`
	TenantName              *string                                          `json:"tenant_name" tfsdk:"tenant_name"`
	TenantUuid              *string                                          `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Url                     *string                                          `json:"url" tfsdk:"url"`
}

type OpenstackSecurityGroupRulesResponse struct {
	Cidr            *string `json:"cidr" tfsdk:"cidr"`
	Description     *string `json:"description" tfsdk:"description"`
	Direction       *string `json:"direction" tfsdk:"direction"`
	Ethertype       *string `json:"ethertype" tfsdk:"ethertype"`
	FromPort        *int64  `json:"from_port" tfsdk:"from_port"`
	Protocol        *string `json:"protocol" tfsdk:"protocol"`
	RemoteGroup     *string `json:"remote_group" tfsdk:"remote_group"`
	ToPort          *int64  `json:"to_port" tfsdk:"to_port"`
	Id              *int64  `json:"id" tfsdk:"id"`
	RemoteGroupName *string `json:"remote_group_name" tfsdk:"remote_group_name"`
	RemoteGroupUuid *string `json:"remote_group_uuid" tfsdk:"remote_group_uuid"`
}

func (r *OpenstackSecurityGroupResponse) GetState() string {
	if r.State != nil {
		return *r.State
	}
	return "OK"
}

func (r *OpenstackSecurityGroupResponse) GetErrorMessage() string {
	if r.ErrorMessage != nil {
		return *r.ErrorMessage
	}
	return ""
}
