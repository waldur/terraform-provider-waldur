package server_group

import (
	"github.com/waldur/terraform-provider-waldur/internal/sdk/common"
)

type OpenstackServerGroupCreateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
	Policy      *string `json:"policy,omitempty" tfsdk:"policy"`
}

type OpenstackServerGroupUpdateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
	Policy      *string `json:"policy,omitempty" tfsdk:"policy"`
}

type OpenstackServerGroupResponse struct {
	UUID *string `json:"uuid"`

	BackendId               *string                           `json:"backend_id" tfsdk:"backend_id"`
	Customer                *string                           `json:"customer" tfsdk:"customer"`
	Description             *string                           `json:"description" tfsdk:"description"`
	DisplayName             *string                           `json:"display_name" tfsdk:"display_name"`
	ErrorMessage            *string                           `json:"error_message" tfsdk:"error_message"`
	Instances               *[]common.OpenStackNestedInstance `json:"instances" tfsdk:"instances"`
	MarketplaceResourceUuid *string                           `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Name                    *string                           `json:"name" tfsdk:"name"`
	Policy                  *string                           `json:"policy" tfsdk:"policy"`
	Project                 *string                           `json:"project" tfsdk:"project"`
	ResourceType            *string                           `json:"resource_type" tfsdk:"resource_type"`
	State                   *string                           `json:"state" tfsdk:"state"`
	Tenant                  *string                           `json:"tenant" tfsdk:"tenant"`
	TenantName              *string                           `json:"tenant_name" tfsdk:"tenant_name"`
	TenantUuid              *string                           `json:"tenant_uuid" tfsdk:"tenant_uuid"`
	Url                     *string                           `json:"url" tfsdk:"url"`
}

type OpenstackServerGroupInstancesResponse struct {
	BackendId *string `json:"backend_id" tfsdk:"backend_id"`
	Name      *string `json:"name" tfsdk:"name"`
	Uuid      *string `json:"uuid" tfsdk:"uuid"`
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
