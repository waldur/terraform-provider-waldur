package security_group

import (
	"encoding/json"

	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackSecurityGroupCreateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`

	Name *string `json:"name" tfsdk:"name"`

	Rules *[]common.OpenStackSecurityGroupRuleCreateRequest `json:"rules,omitempty" tfsdk:"rules"`
}

type OpenstackSecurityGroupUpdateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`

	Name *string `json:"name,omitempty" tfsdk:"name"`
}

type OpenstackSecurityGroupSetRulesActionRequest struct {
	Rules []common.OpenStackSecurityGroupRuleCreateRequest `json:"-"`
}

func (r OpenstackSecurityGroupSetRulesActionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Rules)
}

type OpenstackSecurityGroupResponse struct {
	UUID *string `json:"uuid"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	Customer *string `json:"customer,omitempty" tfsdk:"customer"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	ErrorMessage *string `json:"error_message,omitempty" tfsdk:"error_message"`

	MarketplaceResourceUuid *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`

	Name *string `json:"name" tfsdk:"name"`

	Project *string `json:"project,omitempty" tfsdk:"project"`

	ResourceType *string `json:"resource_type,omitempty" tfsdk:"resource_type"`

	Rules *[]common.OpenStackSecurityGroupRuleCreateRequest `json:"rules,omitempty" tfsdk:"rules"`

	State *string `json:"state,omitempty" tfsdk:"state"`

	Tenant *string `json:"tenant" tfsdk:"tenant"`

	TenantName *string `json:"tenant_name,omitempty" tfsdk:"tenant_name"`

	TenantUuid *string `json:"tenant_uuid,omitempty" tfsdk:"tenant_uuid"`

	Url *string `json:"url,omitempty" tfsdk:"url"`
}

type OpenstackSecurityGroupRulesResponse struct {
	Cidr *string `json:"cidr,omitempty" tfsdk:"cidr"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	Direction *string `json:"direction,omitempty" tfsdk:"direction"`

	Ethertype *string `json:"ethertype,omitempty" tfsdk:"ethertype"`

	FromPort *int64 `json:"from_port,omitempty" tfsdk:"from_port"`

	Protocol *string `json:"protocol,omitempty" tfsdk:"protocol"`

	RemoteGroup *string `json:"remote_group,omitempty" tfsdk:"remote_group"`

	ToPort *int64 `json:"to_port,omitempty" tfsdk:"to_port"`

	Id *int64 `json:"id,omitempty" tfsdk:"id"`

	RemoteGroupName *string `json:"remote_group_name,omitempty" tfsdk:"remote_group_name"`

	RemoteGroupUuid *string `json:"remote_group_uuid,omitempty" tfsdk:"remote_group_uuid"`
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
