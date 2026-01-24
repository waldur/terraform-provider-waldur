package order

// MarketplaceOrder Structs

type MarketplaceOrderCreateRequest struct {
	AcceptingTermsOfService *bool                  `json:"accepting_terms_of_service,omitempty" tfsdk:"accepting_terms_of_service"`
	Attributes              map[string]interface{} `json:"attributes" tfsdk:"attributes"`
	CallbackUrl             *string                `json:"callback_url,omitempty" tfsdk:"callback_url"`
	Offering                *string                `json:"offering" tfsdk:"offering"`
	Plan                    *string                `json:"plan,omitempty" tfsdk:"plan"`
	Project                 *string                `json:"project" tfsdk:"project"`
	RequestComment          *string                `json:"request_comment,omitempty" tfsdk:"request_comment"`
	Slug                    *string                `json:"slug,omitempty" tfsdk:"slug"`
	StartDate               *string                `json:"start_date,omitempty" tfsdk:"start_date"`
	Type                    *string                `json:"type,omitempty" tfsdk:"type"`
}

type MarketplaceOrderCreateAttributesRequest struct {
}

type MarketplaceOrderResponse struct {
	UUID *string `json:"uuid"`

	ActivationPrice            *float64 `json:"activation_price" tfsdk:"activation_price"`
	Attachment                 *string  `json:"attachment" tfsdk:"attachment"`
	BackendId                  *string  `json:"backend_id" tfsdk:"backend_id"`
	CallbackUrl                *string  `json:"callback_url" tfsdk:"callback_url"`
	CanTerminate               *bool    `json:"can_terminate" tfsdk:"can_terminate"`
	CategoryIcon               *string  `json:"category_icon" tfsdk:"category_icon"`
	CategoryTitle              *string  `json:"category_title" tfsdk:"category_title"`
	CategoryUuid               *string  `json:"category_uuid" tfsdk:"category_uuid"`
	CompletedAt                *string  `json:"completed_at" tfsdk:"completed_at"`
	ConsumerReviewedAt         *string  `json:"consumer_reviewed_at" tfsdk:"consumer_reviewed_at"`
	ConsumerReviewedBy         *string  `json:"consumer_reviewed_by" tfsdk:"consumer_reviewed_by"`
	ConsumerReviewedByFullName *string  `json:"consumer_reviewed_by_full_name" tfsdk:"consumer_reviewed_by_full_name"`
	ConsumerReviewedByUsername *string  `json:"consumer_reviewed_by_username" tfsdk:"consumer_reviewed_by_username"`
	Cost                       *string  `json:"cost" tfsdk:"cost"`
	Created                    *string  `json:"created" tfsdk:"created"`
	CreatedByCivilNumber       *string  `json:"created_by_civil_number" tfsdk:"created_by_civil_number"`
	CreatedByFullName          *string  `json:"created_by_full_name" tfsdk:"created_by_full_name"`
	CreatedByUsername          *string  `json:"created_by_username" tfsdk:"created_by_username"`
	CustomerName               *string  `json:"customer_name" tfsdk:"customer_name"`
	CustomerSlug               *string  `json:"customer_slug" tfsdk:"customer_slug"`
	CustomerUuid               *string  `json:"customer_uuid" tfsdk:"customer_uuid"`
	ErrorMessage               *string  `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback             *string  `json:"error_traceback" tfsdk:"error_traceback"`
	FixedPrice                 *float64 `json:"fixed_price" tfsdk:"fixed_price"`
	MarketplaceResourceUuid    *string  `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                   *string  `json:"modified" tfsdk:"modified"`
	NewCostEstimate            *string  `json:"new_cost_estimate" tfsdk:"new_cost_estimate"`
	NewPlanName                *string  `json:"new_plan_name" tfsdk:"new_plan_name"`
	NewPlanUuid                *string  `json:"new_plan_uuid" tfsdk:"new_plan_uuid"`
	Offering                   *string  `json:"offering" tfsdk:"offering"`
	OfferingBillable           *bool    `json:"offering_billable" tfsdk:"offering_billable"`
	OfferingDescription        *string  `json:"offering_description" tfsdk:"offering_description"`
	OfferingImage              *string  `json:"offering_image" tfsdk:"offering_image"`
	OfferingName               *string  `json:"offering_name" tfsdk:"offering_name"`
	OfferingShared             *bool    `json:"offering_shared" tfsdk:"offering_shared"`
	OfferingThumbnail          *string  `json:"offering_thumbnail" tfsdk:"offering_thumbnail"`
	OfferingType               *string  `json:"offering_type" tfsdk:"offering_type"`
	OfferingUuid               *string  `json:"offering_uuid" tfsdk:"offering_uuid"`
	OldCostEstimate            *float64 `json:"old_cost_estimate" tfsdk:"old_cost_estimate"`
	OldPlanName                *string  `json:"old_plan_name" tfsdk:"old_plan_name"`
	OldPlanUuid                *string  `json:"old_plan_uuid" tfsdk:"old_plan_uuid"`
	OrderSubtype               *string  `json:"order_subtype" tfsdk:"order_subtype"`
	Output                     *string  `json:"output" tfsdk:"output"`
	Plan                       *string  `json:"plan" tfsdk:"plan"`
	PlanDescription            *string  `json:"plan_description" tfsdk:"plan_description"`
	PlanName                   *string  `json:"plan_name" tfsdk:"plan_name"`
	PlanUnit                   *string  `json:"plan_unit" tfsdk:"plan_unit"`
	PlanUuid                   *string  `json:"plan_uuid" tfsdk:"plan_uuid"`
	ProjectDescription         *string  `json:"project_description" tfsdk:"project_description"`
	ProjectName                *string  `json:"project_name" tfsdk:"project_name"`
	ProjectSlug                *string  `json:"project_slug" tfsdk:"project_slug"`
	ProjectUuid                *string  `json:"project_uuid" tfsdk:"project_uuid"`
	ProviderName               *string  `json:"provider_name" tfsdk:"provider_name"`
	ProviderReviewedAt         *string  `json:"provider_reviewed_at" tfsdk:"provider_reviewed_at"`
	ProviderReviewedBy         *string  `json:"provider_reviewed_by" tfsdk:"provider_reviewed_by"`
	ProviderReviewedByFullName *string  `json:"provider_reviewed_by_full_name" tfsdk:"provider_reviewed_by_full_name"`
	ProviderReviewedByUsername *string  `json:"provider_reviewed_by_username" tfsdk:"provider_reviewed_by_username"`
	ProviderSlug               *string  `json:"provider_slug" tfsdk:"provider_slug"`
	ProviderUuid               *string  `json:"provider_uuid" tfsdk:"provider_uuid"`
	RequestComment             *string  `json:"request_comment" tfsdk:"request_comment"`
	ResourceName               *string  `json:"resource_name" tfsdk:"resource_name"`
	ResourceType               *string  `json:"resource_type" tfsdk:"resource_type"`
	ResourceUuid               *string  `json:"resource_uuid" tfsdk:"resource_uuid"`
	Slug                       *string  `json:"slug" tfsdk:"slug"`
	StartDate                  *string  `json:"start_date" tfsdk:"start_date"`
	State                      *string  `json:"state" tfsdk:"state"`
	TerminationComment         *string  `json:"termination_comment" tfsdk:"termination_comment"`
	Type                       *string  `json:"type" tfsdk:"type"`
	Url                        *string  `json:"url" tfsdk:"url"`
}

func (r *MarketplaceOrderResponse) GetState() string {
	if r.State != nil {
		return *r.State
	}
	return "OK"
}

func (r *MarketplaceOrderResponse) GetErrorMessage() string {
	if r.ErrorMessage != nil {
		return *r.ErrorMessage
	}
	return ""
}
