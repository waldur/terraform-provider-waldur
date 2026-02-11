package server_group

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackServerGroupCreateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`

	Name *string `json:"name" tfsdk:"name"`

	Policy *string `json:"policy,omitempty" tfsdk:"policy"`
}

type OpenstackServerGroupUpdateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Policy *string `json:"policy,omitempty" tfsdk:"policy"`
}

type OpenstackServerGroupResponse struct {
	UUID *string `json:"uuid"`

	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	Customer *string `json:"customer,omitempty" tfsdk:"customer"`

	Description *string `json:"description,omitempty" tfsdk:"description"`

	DisplayName *string `json:"display_name,omitempty" tfsdk:"display_name"`

	ErrorMessage *string `json:"error_message,omitempty" tfsdk:"error_message"`

	Instances *[]common.OpenStackNestedInstance `json:"instances,omitempty" tfsdk:"instances"`

	MarketplaceResourceUuid *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`

	Name *string `json:"name" tfsdk:"name"`

	Policy *string `json:"policy,omitempty" tfsdk:"policy"`

	Project *string `json:"project,omitempty" tfsdk:"project"`

	ResourceType *string `json:"resource_type,omitempty" tfsdk:"resource_type"`

	State *string `json:"state,omitempty" tfsdk:"state"`

	Tenant *string `json:"tenant" tfsdk:"tenant"`

	TenantName *string `json:"tenant_name,omitempty" tfsdk:"tenant_name"`

	TenantUuid *string `json:"tenant_uuid,omitempty" tfsdk:"tenant_uuid"`

	Url *string `json:"url,omitempty" tfsdk:"url"`
}

type OpenstackServerGroupInstancesResponse struct {
	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`

	Name *string `json:"name,omitempty" tfsdk:"name"`

	Uuid *string `json:"uuid,omitempty" tfsdk:"uuid"`
}

func (r *OpenstackServerGroupResponse) GetState() string {
	if r.State != nil {
		return *r.State
	}
	return "OK"
}

func (r *OpenstackServerGroupResponse) GetErrorMessage() string {
	if r.ErrorMessage != nil {
		return *r.ErrorMessage
	}
	return ""
}
