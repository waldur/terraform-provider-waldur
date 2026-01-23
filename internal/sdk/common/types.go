package common

type AccessSubnet struct {
	Customer    *string `json:"customer" tfsdk:"customer"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Inet        *string `json:"inet" tfsdk:"inet"`
}

type AccessSubnetRequest struct {
	Customer    *string `json:"customer" tfsdk:"customer"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Inet        *string `json:"inet" tfsdk:"inet"`
}

type AccountNameGenerationPolicyEnum struct {
}

type AdminAnnouncement struct {
	ActiveFrom                      *string                                         `json:"active_from,omitempty" tfsdk:"active_from"`
	ActiveTo                        *string                                         `json:"active_to,omitempty" tfsdk:"active_to"`
	Created                         *string                                         `json:"created,omitempty" tfsdk:"created"`
	Description                     *string                                         `json:"description,omitempty" tfsdk:"description"`
	IsActive                        *bool                                           `json:"is_active,omitempty" tfsdk:"is_active"`
	MaintenanceAffectedOfferings    []AdminAnnouncementMaintenanceAffectedOfferings `json:"maintenance_affected_offerings,omitempty" tfsdk:"maintenance_affected_offerings"`
	MaintenanceExternalReferenceUrl *string                                         `json:"maintenance_external_reference_url,omitempty" tfsdk:"maintenance_external_reference_url"`
	MaintenanceName                 *string                                         `json:"maintenance_name,omitempty" tfsdk:"maintenance_name"`
	MaintenanceScheduledEnd         *string                                         `json:"maintenance_scheduled_end,omitempty" tfsdk:"maintenance_scheduled_end"`
	MaintenanceScheduledStart       *string                                         `json:"maintenance_scheduled_start,omitempty" tfsdk:"maintenance_scheduled_start"`
	MaintenanceServiceProvider      *string                                         `json:"maintenance_service_provider,omitempty" tfsdk:"maintenance_service_provider"`
	MaintenanceState                *string                                         `json:"maintenance_state,omitempty" tfsdk:"maintenance_state"`
	MaintenanceType                 *string                                         `json:"maintenance_type,omitempty" tfsdk:"maintenance_type"`
	MaintenanceUuid                 *string                                         `json:"maintenance_uuid,omitempty" tfsdk:"maintenance_uuid"`
	Type                            *string                                         `json:"type,omitempty" tfsdk:"type"`
}

type AdminAnnouncementMaintenanceAffectedOfferings struct {
	ImpactDescription  *string `json:"impact_description,omitempty" tfsdk:"impact_description"`
	ImpactLevel        *string `json:"impact_level,omitempty" tfsdk:"impact_level"`
	ImpactLevelDisplay *string `json:"impact_level_display,omitempty" tfsdk:"impact_level_display"`
	Name               *string `json:"name,omitempty" tfsdk:"name"`
}

type AdminAnnouncementRequest struct {
	ActiveFrom  *string `json:"active_from" tfsdk:"active_from"`
	ActiveTo    *string `json:"active_to" tfsdk:"active_to"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Type        *string `json:"type,omitempty" tfsdk:"type"`
}

type AdminAnnouncementTypeEnum struct {
}

type AgentEventSubscriptionCreateRequest struct {
	Description          *string `json:"description,omitempty" tfsdk:"description"`
	ObservableObjectType *string `json:"observable_object_type" tfsdk:"observable_object_type"`
}

type AgentIdentity struct {
	ConfigFileContent *string              `json:"config_file_content,omitempty" tfsdk:"config_file_content"`
	ConfigFilePath    *string              `json:"config_file_path,omitempty" tfsdk:"config_file_path"`
	Created           *string              `json:"created" tfsdk:"created"`
	LastRestarted     *string              `json:"last_restarted,omitempty" tfsdk:"last_restarted"`
	Modified          *string              `json:"modified" tfsdk:"modified"`
	Name              *string              `json:"name" tfsdk:"name"`
	Offering          *string              `json:"offering" tfsdk:"offering"`
	Services          []NestedAgentService `json:"services" tfsdk:"services"`
	Url               *string              `json:"url" tfsdk:"url"`
	Version           *string              `json:"version,omitempty" tfsdk:"version"`
}

type AgentIdentityRequest struct {
	ConfigFileContent *string `json:"config_file_content,omitempty" tfsdk:"config_file_content"`
	ConfigFilePath    *string `json:"config_file_path,omitempty" tfsdk:"config_file_path"`
	LastRestarted     *string `json:"last_restarted,omitempty" tfsdk:"last_restarted"`
	Name              *string `json:"name" tfsdk:"name"`
	Offering          *string `json:"offering" tfsdk:"offering"`
	Version           *string `json:"version,omitempty" tfsdk:"version"`
}

type AgentProcessor struct {
	BackendType    *string `json:"backend_type" tfsdk:"backend_type"`
	BackendVersion *string `json:"backend_version,omitempty" tfsdk:"backend_version"`
	Created        *string `json:"created" tfsdk:"created"`
	LastRun        *string `json:"last_run,omitempty" tfsdk:"last_run"`
	Modified       *string `json:"modified" tfsdk:"modified"`
	Name           *string `json:"name" tfsdk:"name"`
	Service        *string `json:"service" tfsdk:"service"`
	ServiceName    *string `json:"service_name" tfsdk:"service_name"`
	Url            *string `json:"url" tfsdk:"url"`
}

type AgentProcessorCreateRequest struct {
	BackendType    *string `json:"backend_type" tfsdk:"backend_type"`
	BackendVersion *string `json:"backend_version,omitempty" tfsdk:"backend_version"`
	Name           *string `json:"name" tfsdk:"name"`
}

type AgentService struct {
	Created      *string                `json:"created" tfsdk:"created"`
	Identity     *string                `json:"identity" tfsdk:"identity"`
	IdentityName *string                `json:"identity_name" tfsdk:"identity_name"`
	Mode         *string                `json:"mode,omitempty" tfsdk:"mode"`
	Modified     *string                `json:"modified" tfsdk:"modified"`
	Name         *string                `json:"name" tfsdk:"name"`
	Processors   []NestedAgentProcessor `json:"processors" tfsdk:"processors"`
	State        *string                `json:"state" tfsdk:"state"`
	Url          *string                `json:"url" tfsdk:"url"`
}

type AgentServiceCreateRequest struct {
	Mode *string `json:"mode,omitempty" tfsdk:"mode"`
	Name *string `json:"name" tfsdk:"name"`
}

type AgentServiceState struct {
}

type AgentServiceStatisticsRequest struct {
}

type AgentTypeEnum struct {
}

type AgreementTypeEnum struct {
}

type Allocation struct {
	AccessUrl                   *string `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created                     *string `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	Groupname                   *string `json:"groupname,omitempty" tfsdk:"groupname"`
	IsActive                    *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	IsLimitBased                *bool   `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool   `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string `json:"name,omitempty" tfsdk:"name"`
	NodeLimit                   *int64  `json:"node_limit,omitempty" tfsdk:"node_limit"`
	NodeUsage                   *string `json:"node_usage,omitempty" tfsdk:"node_usage"`
	Project                     *string `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ResourceType                *string `json:"resource_type,omitempty" tfsdk:"resource_type"`
	ServiceName                 *string `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string `json:"state,omitempty" tfsdk:"state"`
	Url                         *string `json:"url,omitempty" tfsdk:"url"`
}

type AllocationRequest struct {
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	Groupname       *string `json:"groupname,omitempty" tfsdk:"groupname"`
	Name            *string `json:"name" tfsdk:"name"`
	NodeLimit       *int64  `json:"node_limit,omitempty" tfsdk:"node_limit"`
	Project         *string `json:"project" tfsdk:"project"`
	ServiceSettings *string `json:"service_settings" tfsdk:"service_settings"`
}

type AllocationSetLimits struct {
	NodeLimit *int64 `json:"node_limit" tfsdk:"node_limit"`
}

type AllocationSetLimitsRequest struct {
	NodeLimit *int64 `json:"node_limit" tfsdk:"node_limit"`
}

type AllocationTimeEnum struct {
}

type AllocationUserUsage struct {
	Allocation *string `json:"allocation" tfsdk:"allocation"`
	FullName   *string `json:"full_name" tfsdk:"full_name"`
	Month      *int64  `json:"month" tfsdk:"month"`
	NodeUsage  *string `json:"node_usage,omitempty" tfsdk:"node_usage"`
	User       *string `json:"user,omitempty" tfsdk:"user"`
	Username   *string `json:"username" tfsdk:"username"`
	Year       *int64  `json:"year" tfsdk:"year"`
}

type Answer struct {
	Created             *string `json:"created" tfsdk:"created"`
	Modified            *string `json:"modified" tfsdk:"modified"`
	QuestionDescription *string `json:"question_description" tfsdk:"question_description"`
	QuestionRequired    *bool   `json:"question_required" tfsdk:"question_required"`
	QuestionType        *string `json:"question_type" tfsdk:"question_type"`
	RequiresReview      *bool   `json:"requires_review" tfsdk:"requires_review"`
	User                *int64  `json:"user" tfsdk:"user"`
	UserName            *string `json:"user_name" tfsdk:"user_name"`
}

type AnswerSubmitRequest struct {
	QuestionUuid *string `json:"question_uuid" tfsdk:"question_uuid"`
}

type AnswerSubmitResponse struct {
	Completion *ChecklistCompletion `json:"completion" tfsdk:"completion"`
	Detail     *string              `json:"detail" tfsdk:"detail"`
}

type Association struct {
	Allocation     *string `json:"allocation" tfsdk:"allocation"`
	Groupname      *string `json:"groupname,omitempty" tfsdk:"groupname"`
	Useridentifier *string `json:"useridentifier,omitempty" tfsdk:"useridentifier"`
	Username       *string `json:"username,omitempty" tfsdk:"username"`
}

type AtlassianCredentialsRequest struct {
	ApiUrl              *string `json:"api_url" tfsdk:"api_url"`
	AuthMethod          *string `json:"auth_method" tfsdk:"auth_method"`
	Email               *string `json:"email,omitempty" tfsdk:"email"`
	Password            *string `json:"password,omitempty" tfsdk:"password"`
	PersonalAccessToken *string `json:"personal_access_token,omitempty" tfsdk:"personal_access_token"`
	Token               *string `json:"token,omitempty" tfsdk:"token"`
	Username            *string `json:"username,omitempty" tfsdk:"username"`
	VerifySsl           *bool   `json:"verify_ssl,omitempty" tfsdk:"verify_ssl"`
}

type AtlassianCustomFieldResponse struct {
	ClauseNames []string `json:"clause_names,omitempty" tfsdk:"clause_names"`
	FieldType   *string  `json:"field_type,omitempty" tfsdk:"field_type"`
	Id          *string  `json:"id" tfsdk:"id"`
	Name        *string  `json:"name" tfsdk:"name"`
	Required    *bool    `json:"required,omitempty" tfsdk:"required"`
}

type AtlassianPriorityResponse struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	IconUrl     *string `json:"icon_url,omitempty" tfsdk:"icon_url"`
	Id          *string `json:"id" tfsdk:"id"`
	Name        *string `json:"name" tfsdk:"name"`
}

type AtlassianProjectResponse struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Id          *string `json:"id" tfsdk:"id"`
	Key         *string `json:"key" tfsdk:"key"`
	Name        *string `json:"name" tfsdk:"name"`
}

type AtlassianRequestTypeResponse struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Id          *string `json:"id" tfsdk:"id"`
	IssueTypeId *string `json:"issue_type_id,omitempty" tfsdk:"issue_type_id"`
	Name        *string `json:"name" tfsdk:"name"`
}

type AtlassianSettingsPreviewRequest struct {
	AffectedResourceField     *string  `json:"affected_resource_field,omitempty" tfsdk:"affected_resource_field"`
	ApiUrl                    *string  `json:"api_url" tfsdk:"api_url"`
	AuthMethod                *string  `json:"auth_method" tfsdk:"auth_method"`
	CallerField               *string  `json:"caller_field,omitempty" tfsdk:"caller_field"`
	CustomFieldMappingEnabled *bool    `json:"custom_field_mapping_enabled,omitempty" tfsdk:"custom_field_mapping_enabled"`
	DefaultOfferingIssueType  *string  `json:"default_offering_issue_type,omitempty" tfsdk:"default_offering_issue_type"`
	Email                     *string  `json:"email,omitempty" tfsdk:"email"`
	ImpactField               *string  `json:"impact_field,omitempty" tfsdk:"impact_field"`
	IssueTypes                []string `json:"issue_types,omitempty" tfsdk:"issue_types"`
	OrganisationField         *string  `json:"organisation_field,omitempty" tfsdk:"organisation_field"`
	Password                  *string  `json:"password,omitempty" tfsdk:"password"`
	PersonalAccessToken       *string  `json:"personal_access_token,omitempty" tfsdk:"personal_access_token"`
	ProjectField              *string  `json:"project_field,omitempty" tfsdk:"project_field"`
	ProjectId                 *string  `json:"project_id" tfsdk:"project_id"`
	ReporterField             *string  `json:"reporter_field,omitempty" tfsdk:"reporter_field"`
	RequestFeedbackField      *string  `json:"request_feedback_field,omitempty" tfsdk:"request_feedback_field"`
	ResolutionSlaField        *string  `json:"resolution_sla_field,omitempty" tfsdk:"resolution_sla_field"`
	SatisfactionField         *string  `json:"satisfaction_field,omitempty" tfsdk:"satisfaction_field"`
	SlaField                  *string  `json:"sla_field,omitempty" tfsdk:"sla_field"`
	TemplateField             *string  `json:"template_field,omitempty" tfsdk:"template_field"`
	Token                     *string  `json:"token,omitempty" tfsdk:"token"`
	UseOldApi                 *bool    `json:"use_old_api,omitempty" tfsdk:"use_old_api"`
	Username                  *string  `json:"username,omitempty" tfsdk:"username"`
	VerifySsl                 *bool    `json:"verify_ssl,omitempty" tfsdk:"verify_ssl"`
	WaldurBackendIdField      *string  `json:"waldur_backend_id_field,omitempty" tfsdk:"waldur_backend_id_field"`
}

type AtlassianSettingsSaveRequest struct {
	AffectedResourceField     *string  `json:"affected_resource_field,omitempty" tfsdk:"affected_resource_field"`
	ApiUrl                    *string  `json:"api_url" tfsdk:"api_url"`
	AuthMethod                *string  `json:"auth_method" tfsdk:"auth_method"`
	CallerField               *string  `json:"caller_field,omitempty" tfsdk:"caller_field"`
	ConfirmSave               *bool    `json:"confirm_save" tfsdk:"confirm_save"`
	CustomFieldMappingEnabled *bool    `json:"custom_field_mapping_enabled,omitempty" tfsdk:"custom_field_mapping_enabled"`
	DefaultOfferingIssueType  *string  `json:"default_offering_issue_type,omitempty" tfsdk:"default_offering_issue_type"`
	Email                     *string  `json:"email,omitempty" tfsdk:"email"`
	ImpactField               *string  `json:"impact_field,omitempty" tfsdk:"impact_field"`
	IssueTypes                []string `json:"issue_types,omitempty" tfsdk:"issue_types"`
	OrganisationField         *string  `json:"organisation_field,omitempty" tfsdk:"organisation_field"`
	Password                  *string  `json:"password,omitempty" tfsdk:"password"`
	PersonalAccessToken       *string  `json:"personal_access_token,omitempty" tfsdk:"personal_access_token"`
	ProjectField              *string  `json:"project_field,omitempty" tfsdk:"project_field"`
	ProjectId                 *string  `json:"project_id" tfsdk:"project_id"`
	ReporterField             *string  `json:"reporter_field,omitempty" tfsdk:"reporter_field"`
	RequestFeedbackField      *string  `json:"request_feedback_field,omitempty" tfsdk:"request_feedback_field"`
	ResolutionSlaField        *string  `json:"resolution_sla_field,omitempty" tfsdk:"resolution_sla_field"`
	SatisfactionField         *string  `json:"satisfaction_field,omitempty" tfsdk:"satisfaction_field"`
	SlaField                  *string  `json:"sla_field,omitempty" tfsdk:"sla_field"`
	TemplateField             *string  `json:"template_field,omitempty" tfsdk:"template_field"`
	Token                     *string  `json:"token,omitempty" tfsdk:"token"`
	UseOldApi                 *bool    `json:"use_old_api,omitempty" tfsdk:"use_old_api"`
	Username                  *string  `json:"username,omitempty" tfsdk:"username"`
	VerifySsl                 *bool    `json:"verify_ssl,omitempty" tfsdk:"verify_ssl"`
	WaldurBackendIdField      *string  `json:"waldur_backend_id_field,omitempty" tfsdk:"waldur_backend_id_field"`
}

type Attachment struct {
	BackendId          *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created            *string `json:"created,omitempty" tfsdk:"created"`
	DestroyIsAvailable *bool   `json:"destroy_is_available,omitempty" tfsdk:"destroy_is_available"`
	File               *string `json:"file,omitempty" tfsdk:"file"`
	FileName           *string `json:"file_name,omitempty" tfsdk:"file_name"`
	FileSize           *int64  `json:"file_size,omitempty" tfsdk:"file_size"`
	Issue              *string `json:"issue,omitempty" tfsdk:"issue"`
	IssueKey           *string `json:"issue_key,omitempty" tfsdk:"issue_key"`
	MimeType           *string `json:"mime_type,omitempty" tfsdk:"mime_type"`
	Url                *string `json:"url,omitempty" tfsdk:"url"`
}

type AttachmentRequest struct {
	File  *string `json:"file" tfsdk:"file"`
	Issue *string `json:"issue" tfsdk:"issue"`
}

type AttachmentRequestForm struct {
	File  *string `json:"file" tfsdk:"file"`
	Issue *string `json:"issue" tfsdk:"issue"`
}

type AttachmentRequestMultipart struct {
	File  *string `json:"file" tfsdk:"file"`
	Issue *string `json:"issue" tfsdk:"issue"`
}

type AuthMethodEnum struct {
}

type AuthResult struct {
	Details      *string `json:"details" tfsdk:"details"`
	ErrorMessage *string `json:"error_message" tfsdk:"error_message"`
	Message      *string `json:"message" tfsdk:"message"`
	Phone        *string `json:"phone" tfsdk:"phone"`
	State        *string `json:"state" tfsdk:"state"`
	Token        *string `json:"token" tfsdk:"token"`
}

type AuthResultRequest struct {
	Phone *string `json:"phone" tfsdk:"phone"`
}

type AuthResultStateEnum struct {
}

type AuthResultUUIDRequest struct {
}

type AuthToken struct {
	Created           *string `json:"created" tfsdk:"created"`
	Url               *string `json:"url" tfsdk:"url"`
	User              *string `json:"user" tfsdk:"user"`
	UserFirstName     *string `json:"user_first_name" tfsdk:"user_first_name"`
	UserIsActive      *bool   `json:"user_is_active" tfsdk:"user_is_active"`
	UserLastName      *string `json:"user_last_name" tfsdk:"user_last_name"`
	UserTokenLifetime *int64  `json:"user_token_lifetime" tfsdk:"user_token_lifetime"`
	UserUsername      *string `json:"user_username" tfsdk:"user_username"`
}

type AvailableChecklist struct {
	CategoryName   *string `json:"category_name" tfsdk:"category_name"`
	CategoryUuid   *string `json:"category_uuid" tfsdk:"category_uuid"`
	ChecklistType  *string `json:"checklist_type" tfsdk:"checklist_type"`
	Description    *string `json:"description" tfsdk:"description"`
	Name           *string `json:"name" tfsdk:"name"`
	QuestionsCount *int64  `json:"questions_count" tfsdk:"questions_count"`
}

type AwsImage struct {
	Name *string `json:"name" tfsdk:"name"`
	Url  *string `json:"url" tfsdk:"url"`
}

type AwsInstance struct {
	AccessUrl                   *string  `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string  `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Cores                       *int64   `json:"cores,omitempty" tfsdk:"cores"`
	Created                     *string  `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string  `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string  `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string  `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string  `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string  `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string  `json:"description,omitempty" tfsdk:"description"`
	Disk                        *int64   `json:"disk,omitempty" tfsdk:"disk"`
	ErrorMessage                *string  `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string  `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	ExternalIps                 []string `json:"external_ips,omitempty" tfsdk:"external_ips"`
	ImageName                   *string  `json:"image_name,omitempty" tfsdk:"image_name"`
	InternalIps                 []string `json:"internal_ips,omitempty" tfsdk:"internal_ips"`
	IsLimitBased                *bool    `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool    `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	KeyFingerprint              *string  `json:"key_fingerprint,omitempty" tfsdk:"key_fingerprint"`
	KeyName                     *string  `json:"key_name,omitempty" tfsdk:"key_name"`
	Latitude                    *float64 `json:"latitude,omitempty" tfsdk:"latitude"`
	Longitude                   *float64 `json:"longitude,omitempty" tfsdk:"longitude"`
	MarketplaceCategoryName     *string  `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string  `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string  `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string  `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string  `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string  `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string  `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	MinDisk                     *int64   `json:"min_disk,omitempty" tfsdk:"min_disk"`
	MinRam                      *int64   `json:"min_ram,omitempty" tfsdk:"min_ram"`
	Modified                    *string  `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string  `json:"name,omitempty" tfsdk:"name"`
	Project                     *string  `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string  `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string  `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	Ram                         *int64   `json:"ram,omitempty" tfsdk:"ram"`
	ResourceType                *string  `json:"resource_type,omitempty" tfsdk:"resource_type"`
	ServiceName                 *string  `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string  `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string  `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string  `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string  `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	StartTime                   *string  `json:"start_time,omitempty" tfsdk:"start_time"`
	State                       *string  `json:"state,omitempty" tfsdk:"state"`
	Url                         *string  `json:"url,omitempty" tfsdk:"url"`
	UserData                    *string  `json:"user_data,omitempty" tfsdk:"user_data"`
}

type AwsInstanceRequest struct {
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	Image           *string `json:"image" tfsdk:"image"`
	Name            *string `json:"name" tfsdk:"name"`
	Project         *string `json:"project" tfsdk:"project"`
	Region          *string `json:"region" tfsdk:"region"`
	ServiceSettings *string `json:"service_settings" tfsdk:"service_settings"`
	Size            *string `json:"size" tfsdk:"size"`
	SshPublicKey    *string `json:"ssh_public_key,omitempty" tfsdk:"ssh_public_key"`
	UserData        *string `json:"user_data,omitempty" tfsdk:"user_data"`
}

type AwsInstanceResize struct {
	Size *string `json:"size" tfsdk:"size"`
}

type AwsInstanceResizeRequest struct {
	Size *string `json:"size" tfsdk:"size"`
}

type AwsRegion struct {
	Name *string `json:"name" tfsdk:"name"`
	Url  *string `json:"url" tfsdk:"url"`
}

type AwsSize struct {
	Cores       *int64      `json:"cores" tfsdk:"cores"`
	Description *string     `json:"description" tfsdk:"description"`
	Disk        *int64      `json:"disk" tfsdk:"disk"`
	Name        *string     `json:"name" tfsdk:"name"`
	Ram         *int64      `json:"ram" tfsdk:"ram"`
	Regions     []AwsRegion `json:"regions" tfsdk:"regions"`
	Url         *string     `json:"url" tfsdk:"url"`
}

type AwsVolume struct {
	AccessUrl                   *string `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created                     *string `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string `json:"description,omitempty" tfsdk:"description"`
	Device                      *string `json:"device,omitempty" tfsdk:"device"`
	ErrorMessage                *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	Instance                    *string `json:"instance,omitempty" tfsdk:"instance"`
	IsLimitBased                *bool   `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool   `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string `json:"name,omitempty" tfsdk:"name"`
	Project                     *string `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ResourceType                *string `json:"resource_type,omitempty" tfsdk:"resource_type"`
	RuntimeState                *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceName                 *string `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	Size                        *int64  `json:"size,omitempty" tfsdk:"size"`
	State                       *string `json:"state,omitempty" tfsdk:"state"`
	Url                         *string `json:"url,omitempty" tfsdk:"url"`
	VolumeType                  *string `json:"volume_type,omitempty" tfsdk:"volume_type"`
}

type AwsVolumeAttach struct {
	Device   *string `json:"device" tfsdk:"device"`
	Instance *string `json:"instance" tfsdk:"instance"`
}

type AwsVolumeAttachRequest struct {
	Device   *string `json:"device" tfsdk:"device"`
	Instance *string `json:"instance" tfsdk:"instance"`
}

type AwsVolumeRequest struct {
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	Name            *string `json:"name" tfsdk:"name"`
	Project         *string `json:"project" tfsdk:"project"`
	Region          *string `json:"region" tfsdk:"region"`
	ServiceSettings *string `json:"service_settings" tfsdk:"service_settings"`
	Size            *int64  `json:"size" tfsdk:"size"`
	VolumeType      *string `json:"volume_type" tfsdk:"volume_type"`
}

type AzureImage struct {
	Name      *string `json:"name" tfsdk:"name"`
	Publisher *string `json:"publisher" tfsdk:"publisher"`
	Sku       *string `json:"sku" tfsdk:"sku"`
	Url       *string `json:"url" tfsdk:"url"`
	Version   *string `json:"version" tfsdk:"version"`
}

type AzureLocation struct {
	Latitude  *float64 `json:"latitude,omitempty" tfsdk:"latitude"`
	Longitude *float64 `json:"longitude,omitempty" tfsdk:"longitude"`
	Name      *string  `json:"name" tfsdk:"name"`
	Url       *string  `json:"url" tfsdk:"url"`
}

type AzurePublicIP struct {
	AccessUrl                   *string `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created                     *string `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	IsLimitBased                *bool   `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool   `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	Location                    *string `json:"location,omitempty" tfsdk:"location"`
	MarketplaceCategoryName     *string `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string `json:"name,omitempty" tfsdk:"name"`
	Project                     *string `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ResourceGroup               *string `json:"resource_group,omitempty" tfsdk:"resource_group"`
	ResourceType                *string `json:"resource_type,omitempty" tfsdk:"resource_type"`
	ServiceName                 *string `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string `json:"state,omitempty" tfsdk:"state"`
	Url                         *string `json:"url,omitempty" tfsdk:"url"`
}

type AzurePublicIPRequest struct {
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	Location        *string `json:"location" tfsdk:"location"`
	Name            *string `json:"name" tfsdk:"name"`
	Project         *string `json:"project" tfsdk:"project"`
	ResourceGroup   *string `json:"resource_group" tfsdk:"resource_group"`
	ServiceSettings *string `json:"service_settings" tfsdk:"service_settings"`
}

type AzureResourceGroup struct {
	AccessUrl                   *string `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created                     *string `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	IsLimitBased                *bool   `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool   `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	Location                    *string `json:"location,omitempty" tfsdk:"location"`
	MarketplaceCategoryName     *string `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string `json:"name,omitempty" tfsdk:"name"`
	Project                     *string `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ResourceType                *string `json:"resource_type,omitempty" tfsdk:"resource_type"`
	ServiceName                 *string `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string `json:"state,omitempty" tfsdk:"state"`
	Url                         *string `json:"url,omitempty" tfsdk:"url"`
}

type AzureSQLServerCreateOrderAttributes struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Location    *string `json:"location" tfsdk:"location"`
	Name        *string `json:"name" tfsdk:"name"`
}

type AzureSize struct {
	MaxDataDiskCount     *int64  `json:"max_data_disk_count" tfsdk:"max_data_disk_count"`
	MemoryInMb           *int64  `json:"memory_in_mb" tfsdk:"memory_in_mb"`
	Name                 *string `json:"name" tfsdk:"name"`
	NumberOfCores        *int64  `json:"number_of_cores" tfsdk:"number_of_cores"`
	OsDiskSizeInMb       *int64  `json:"os_disk_size_in_mb" tfsdk:"os_disk_size_in_mb"`
	ResourceDiskSizeInMb *int64  `json:"resource_disk_size_in_mb" tfsdk:"resource_disk_size_in_mb"`
	Url                  *string `json:"url" tfsdk:"url"`
}

type AzureSqlDatabase struct {
	AccessUrl                   *string `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Charset                     *string `json:"charset,omitempty" tfsdk:"charset"`
	Collation                   *string `json:"collation,omitempty" tfsdk:"collation"`
	Created                     *string `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	IsLimitBased                *bool   `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool   `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	LocationName                *string `json:"location_name,omitempty" tfsdk:"location_name"`
	MarketplaceCategoryName     *string `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string `json:"name,omitempty" tfsdk:"name"`
	Project                     *string `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ResourceGroupName           *string `json:"resource_group_name,omitempty" tfsdk:"resource_group_name"`
	ResourceType                *string `json:"resource_type,omitempty" tfsdk:"resource_type"`
	Server                      *string `json:"server,omitempty" tfsdk:"server"`
	ServerMarketplaceUuid       *string `json:"server_marketplace_uuid,omitempty" tfsdk:"server_marketplace_uuid"`
	ServerName                  *string `json:"server_name,omitempty" tfsdk:"server_name"`
	ServerUuid                  *string `json:"server_uuid,omitempty" tfsdk:"server_uuid"`
	ServiceName                 *string `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string `json:"state,omitempty" tfsdk:"state"`
	Url                         *string `json:"url,omitempty" tfsdk:"url"`
}

type AzureSqlDatabaseCreate struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
}

type AzureSqlDatabaseCreateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
}

type AzureSqlDatabaseRequest struct {
	Charset         *string `json:"charset,omitempty" tfsdk:"charset"`
	Collation       *string `json:"collation,omitempty" tfsdk:"collation"`
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	Name            *string `json:"name" tfsdk:"name"`
	Project         *string `json:"project" tfsdk:"project"`
	Server          *string `json:"server" tfsdk:"server"`
	ServiceSettings *string `json:"service_settings" tfsdk:"service_settings"`
}

type AzureSqlServer struct {
	AccessUrl                   *string `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created                     *string `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	Fqdn                        *string `json:"fqdn,omitempty" tfsdk:"fqdn"`
	IsLimitBased                *bool   `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool   `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	LocationName                *string `json:"location_name,omitempty" tfsdk:"location_name"`
	MarketplaceCategoryName     *string `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string `json:"name,omitempty" tfsdk:"name"`
	Password                    *string `json:"password,omitempty" tfsdk:"password"`
	Project                     *string `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ResourceGroup               *string `json:"resource_group,omitempty" tfsdk:"resource_group"`
	ResourceGroupName           *string `json:"resource_group_name,omitempty" tfsdk:"resource_group_name"`
	ResourceType                *string `json:"resource_type,omitempty" tfsdk:"resource_type"`
	ServiceName                 *string `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string `json:"state,omitempty" tfsdk:"state"`
	StorageMb                   *int64  `json:"storage_mb,omitempty" tfsdk:"storage_mb"`
	Url                         *string `json:"url,omitempty" tfsdk:"url"`
	Username                    *string `json:"username,omitempty" tfsdk:"username"`
}

type AzureSqlServerRequest struct {
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	Location        *string `json:"location" tfsdk:"location"`
	Name            *string `json:"name" tfsdk:"name"`
	Project         *string `json:"project" tfsdk:"project"`
	ServiceSettings *string `json:"service_settings" tfsdk:"service_settings"`
	StorageMb       *int64  `json:"storage_mb,omitempty" tfsdk:"storage_mb"`
}

type AzureVirtualMachine struct {
	AccessUrl                   *string  `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string  `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Cores                       *int64   `json:"cores,omitempty" tfsdk:"cores"`
	Created                     *string  `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string  `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string  `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string  `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string  `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string  `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string  `json:"description,omitempty" tfsdk:"description"`
	Disk                        *int64   `json:"disk,omitempty" tfsdk:"disk"`
	ErrorMessage                *string  `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string  `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	ExternalIps                 []string `json:"external_ips,omitempty" tfsdk:"external_ips"`
	Image                       *string  `json:"image,omitempty" tfsdk:"image"`
	ImageName                   *string  `json:"image_name,omitempty" tfsdk:"image_name"`
	InternalIps                 []string `json:"internal_ips,omitempty" tfsdk:"internal_ips"`
	IsLimitBased                *bool    `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool    `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	KeyFingerprint              *string  `json:"key_fingerprint,omitempty" tfsdk:"key_fingerprint"`
	KeyName                     *string  `json:"key_name,omitempty" tfsdk:"key_name"`
	Latitude                    *float64 `json:"latitude,omitempty" tfsdk:"latitude"`
	LocationName                *string  `json:"location_name,omitempty" tfsdk:"location_name"`
	Longitude                   *float64 `json:"longitude,omitempty" tfsdk:"longitude"`
	MarketplaceCategoryName     *string  `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string  `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string  `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string  `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string  `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string  `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string  `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	MinDisk                     *int64   `json:"min_disk,omitempty" tfsdk:"min_disk"`
	MinRam                      *int64   `json:"min_ram,omitempty" tfsdk:"min_ram"`
	Modified                    *string  `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string  `json:"name,omitempty" tfsdk:"name"`
	Password                    *string  `json:"password,omitempty" tfsdk:"password"`
	Project                     *string  `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string  `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string  `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	Ram                         *int64   `json:"ram,omitempty" tfsdk:"ram"`
	ResourceGroup               *string  `json:"resource_group,omitempty" tfsdk:"resource_group"`
	ResourceGroupName           *string  `json:"resource_group_name,omitempty" tfsdk:"resource_group_name"`
	ResourceType                *string  `json:"resource_type,omitempty" tfsdk:"resource_type"`
	RuntimeState                *string  `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceName                 *string  `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string  `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string  `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string  `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string  `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	Size                        *string  `json:"size,omitempty" tfsdk:"size"`
	SizeName                    *string  `json:"size_name,omitempty" tfsdk:"size_name"`
	StartTime                   *string  `json:"start_time,omitempty" tfsdk:"start_time"`
	State                       *string  `json:"state,omitempty" tfsdk:"state"`
	Url                         *string  `json:"url,omitempty" tfsdk:"url"`
	UserData                    *string  `json:"user_data,omitempty" tfsdk:"user_data"`
	Username                    *string  `json:"username,omitempty" tfsdk:"username"`
}

type AzureVirtualMachineCreateOrderAttributes struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Image       *string `json:"image" tfsdk:"image"`
	Location    *string `json:"location" tfsdk:"location"`
	Name        *string `json:"name" tfsdk:"name"`
	Size        *string `json:"size" tfsdk:"size"`
}

type AzureVirtualMachineRequest struct {
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	Image           *string `json:"image" tfsdk:"image"`
	Location        *string `json:"location" tfsdk:"location"`
	Name            *string `json:"name" tfsdk:"name"`
	Project         *string `json:"project" tfsdk:"project"`
	ServiceSettings *string `json:"service_settings" tfsdk:"service_settings"`
	Size            *string `json:"size" tfsdk:"size"`
	SshPublicKey    *string `json:"ssh_public_key,omitempty" tfsdk:"ssh_public_key"`
	UserData        *string `json:"user_data,omitempty" tfsdk:"user_data"`
}

type BackendIdRequest struct {
	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
}

type BackendMetadata struct {
	Action       *string `json:"action,omitempty" tfsdk:"action"`
	InstanceName *string `json:"instance_name,omitempty" tfsdk:"instance_name"`
	RuntimeState *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	State        *string `json:"state,omitempty" tfsdk:"state"`
}

type BackendResource struct {
	BackendId    *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created      *string `json:"created" tfsdk:"created"`
	Modified     *string `json:"modified" tfsdk:"modified"`
	Name         *string `json:"name" tfsdk:"name"`
	Offering     *string `json:"offering" tfsdk:"offering"`
	OfferingName *string `json:"offering_name" tfsdk:"offering_name"`
	OfferingUrl  *string `json:"offering_url" tfsdk:"offering_url"`
	Project      *string `json:"project" tfsdk:"project"`
	ProjectName  *string `json:"project_name" tfsdk:"project_name"`
	ProjectUrl   *string `json:"project_url" tfsdk:"project_url"`
	Url          *string `json:"url" tfsdk:"url"`
}

type BackendResourceImportRequest struct {
	Plan *string `json:"plan,omitempty" tfsdk:"plan"`
}

type BackendResourceReq struct {
	Created        *string `json:"created" tfsdk:"created"`
	ErrorMessage   *string `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback *string `json:"error_traceback" tfsdk:"error_traceback"`
	Finished       *string `json:"finished" tfsdk:"finished"`
	Modified       *string `json:"modified" tfsdk:"modified"`
	Offering       *string `json:"offering" tfsdk:"offering"`
	OfferingName   *string `json:"offering_name" tfsdk:"offering_name"`
	OfferingUrl    *string `json:"offering_url" tfsdk:"offering_url"`
	Started        *string `json:"started" tfsdk:"started"`
	State          *string `json:"state" tfsdk:"state"`
	Url            *string `json:"url" tfsdk:"url"`
}

type BackendResourceReqRequest struct {
	Offering *string `json:"offering" tfsdk:"offering"`
}

type BackendResourceReqStateEnum struct {
}

type BackendResourceRequest struct {
	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Name      *string `json:"name" tfsdk:"name"`
	Offering  *string `json:"offering" tfsdk:"offering"`
	Project   *string `json:"project" tfsdk:"project"`
}

type BackendResourceRequestSetErredRequest struct {
	ErrorMessage   *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
}

type BaseComponentUsage struct {
	Created      *string `json:"created" tfsdk:"created"`
	Date         *string `json:"date" tfsdk:"date"`
	Description  *string `json:"description,omitempty" tfsdk:"description"`
	MeasuredUnit *string `json:"measured_unit" tfsdk:"measured_unit"`
	Name         *string `json:"name" tfsdk:"name"`
	Recurring    *bool   `json:"recurring,omitempty" tfsdk:"recurring"`
	Type         *string `json:"type" tfsdk:"type"`
	Usage        *string `json:"usage,omitempty" tfsdk:"usage"`
}

type BaseProviderPlan struct {
	Archived           *bool                 `json:"archived,omitempty" tfsdk:"archived"`
	ArticleCode        *string               `json:"article_code,omitempty" tfsdk:"article_code"`
	BackendId          *string               `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Components         []NestedPlanComponent `json:"components,omitempty" tfsdk:"components"`
	Description        *string               `json:"description,omitempty" tfsdk:"description"`
	InitPrice          *float64              `json:"init_price,omitempty" tfsdk:"init_price"`
	IsActive           *bool                 `json:"is_active,omitempty" tfsdk:"is_active"`
	MaxAmount          *int64                `json:"max_amount,omitempty" tfsdk:"max_amount"`
	MinimalPrice       *float64              `json:"minimal_price,omitempty" tfsdk:"minimal_price"`
	Name               *string               `json:"name,omitempty" tfsdk:"name"`
	OrganizationGroups []OrganizationGroup   `json:"organization_groups,omitempty" tfsdk:"organization_groups"`
	PlanType           *string               `json:"plan_type,omitempty" tfsdk:"plan_type"`
	ResourcesCount     *int64                `json:"resources_count,omitempty" tfsdk:"resources_count"`
	SwitchPrice        *float64              `json:"switch_price,omitempty" tfsdk:"switch_price"`
	Unit               *string               `json:"unit,omitempty" tfsdk:"unit"`
	UnitPrice          *string               `json:"unit_price,omitempty" tfsdk:"unit_price"`
	Url                *string               `json:"url,omitempty" tfsdk:"url"`
}

type BaseProviderPlanRequest struct {
	Archived    *bool   `json:"archived,omitempty" tfsdk:"archived"`
	ArticleCode *string `json:"article_code,omitempty" tfsdk:"article_code"`
	BackendId   *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	MaxAmount   *int64  `json:"max_amount,omitempty" tfsdk:"max_amount"`
	Name        *string `json:"name" tfsdk:"name"`
	Unit        *string `json:"unit,omitempty" tfsdk:"unit"`
	UnitPrice   *string `json:"unit_price,omitempty" tfsdk:"unit_price"`
}

type BasePublicPlan struct {
	Archived           *bool                 `json:"archived,omitempty" tfsdk:"archived"`
	ArticleCode        *string               `json:"article_code,omitempty" tfsdk:"article_code"`
	BackendId          *string               `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Components         []NestedPlanComponent `json:"components,omitempty" tfsdk:"components"`
	Description        *string               `json:"description,omitempty" tfsdk:"description"`
	InitPrice          *float64              `json:"init_price,omitempty" tfsdk:"init_price"`
	IsActive           *bool                 `json:"is_active,omitempty" tfsdk:"is_active"`
	MaxAmount          *int64                `json:"max_amount,omitempty" tfsdk:"max_amount"`
	MinimalPrice       *float64              `json:"minimal_price,omitempty" tfsdk:"minimal_price"`
	Name               *string               `json:"name,omitempty" tfsdk:"name"`
	OrganizationGroups []OrganizationGroup   `json:"organization_groups,omitempty" tfsdk:"organization_groups"`
	PlanType           *string               `json:"plan_type,omitempty" tfsdk:"plan_type"`
	ResourcesCount     *int64                `json:"resources_count,omitempty" tfsdk:"resources_count"`
	SwitchPrice        *float64              `json:"switch_price,omitempty" tfsdk:"switch_price"`
	Unit               *string               `json:"unit,omitempty" tfsdk:"unit"`
	UnitPrice          *string               `json:"unit_price,omitempty" tfsdk:"unit_price"`
	Url                *string               `json:"url,omitempty" tfsdk:"url"`
}

type BasePublicPlanRequest struct {
	Archived    *bool   `json:"archived,omitempty" tfsdk:"archived"`
	ArticleCode *string `json:"article_code,omitempty" tfsdk:"article_code"`
	BackendId   *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	MaxAmount   *int64  `json:"max_amount,omitempty" tfsdk:"max_amount"`
	Name        *string `json:"name" tfsdk:"name"`
	Unit        *string `json:"unit,omitempty" tfsdk:"unit"`
	UnitPrice   *string `json:"unit_price,omitempty" tfsdk:"unit_price"`
	Url         *string `json:"url" tfsdk:"url"`
}

type BasicCustomer struct {
	Name *string `json:"name" tfsdk:"name"`
}

type BasicUser struct {
	Email      *string `json:"email,omitempty" tfsdk:"email"`
	FullName   *string `json:"full_name,omitempty" tfsdk:"full_name"`
	Image      *string `json:"image,omitempty" tfsdk:"image"`
	NativeName *string `json:"native_name,omitempty" tfsdk:"native_name"`
	Url        *string `json:"url,omitempty" tfsdk:"url"`
	Username   *string `json:"username,omitempty" tfsdk:"username"`
}

type BillingTypeEnum struct {
}

type BillingUnit struct {
}

type BlankEnum struct {
}

type Booking struct {
	CreatedByFullName *string `json:"created_by_full_name" tfsdk:"created_by_full_name"`
	End               *string `json:"end" tfsdk:"end"`
	Start             *string `json:"start" tfsdk:"start"`
}

type BookingResource struct {
	AvailableActions           []string            `json:"available_actions,omitempty" tfsdk:"available_actions"`
	BackendId                  *string             `json:"backend_id,omitempty" tfsdk:"backend_id"`
	CanTerminate               *bool               `json:"can_terminate,omitempty" tfsdk:"can_terminate"`
	CategoryIcon               *string             `json:"category_icon,omitempty" tfsdk:"category_icon"`
	CategoryTitle              *string             `json:"category_title,omitempty" tfsdk:"category_title"`
	CategoryUuid               *string             `json:"category_uuid,omitempty" tfsdk:"category_uuid"`
	ConsumerReviewedBy         *string             `json:"consumer_reviewed_by,omitempty" tfsdk:"consumer_reviewed_by"`
	ConsumerReviewedByFullName *string             `json:"consumer_reviewed_by_full_name,omitempty" tfsdk:"consumer_reviewed_by_full_name"`
	ConsumerReviewedByUsername *string             `json:"consumer_reviewed_by_username,omitempty" tfsdk:"consumer_reviewed_by_username"`
	Created                    *string             `json:"created,omitempty" tfsdk:"created"`
	CreatedBy                  *string             `json:"created_by,omitempty" tfsdk:"created_by"`
	CreatedByFullName          *string             `json:"created_by_full_name,omitempty" tfsdk:"created_by_full_name"`
	CreatedByUsername          *string             `json:"created_by_username,omitempty" tfsdk:"created_by_username"`
	CustomerName               *string             `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerSlug               *string             `json:"customer_slug,omitempty" tfsdk:"customer_slug"`
	CustomerUuid               *string             `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                *string             `json:"description,omitempty" tfsdk:"description"`
	Downscaled                 *bool               `json:"downscaled,omitempty" tfsdk:"downscaled"`
	EffectiveId                *string             `json:"effective_id,omitempty" tfsdk:"effective_id"`
	EndDate                    *string             `json:"end_date,omitempty" tfsdk:"end_date"`
	EndDateRequestedBy         *string             `json:"end_date_requested_by,omitempty" tfsdk:"end_date_requested_by"`
	Endpoints                  []NestedEndpoint    `json:"endpoints,omitempty" tfsdk:"endpoints"`
	ErrorMessage               *string             `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback             *string             `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	IsLimitBased               *bool               `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased               *bool               `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	LastSync                   *string             `json:"last_sync,omitempty" tfsdk:"last_sync"`
	Modified                   *string             `json:"modified,omitempty" tfsdk:"modified"`
	Name                       *string             `json:"name,omitempty" tfsdk:"name"`
	Offering                   *string             `json:"offering,omitempty" tfsdk:"offering"`
	OfferingBillable           *bool               `json:"offering_billable,omitempty" tfsdk:"offering_billable"`
	OfferingComponents         []OfferingComponent `json:"offering_components,omitempty" tfsdk:"offering_components"`
	OfferingDescription        *string             `json:"offering_description,omitempty" tfsdk:"offering_description"`
	OfferingImage              *string             `json:"offering_image,omitempty" tfsdk:"offering_image"`
	OfferingName               *string             `json:"offering_name,omitempty" tfsdk:"offering_name"`
	OfferingShared             *bool               `json:"offering_shared,omitempty" tfsdk:"offering_shared"`
	OfferingSlug               *string             `json:"offering_slug,omitempty" tfsdk:"offering_slug"`
	OfferingState              *string             `json:"offering_state,omitempty" tfsdk:"offering_state"`
	OfferingThumbnail          *string             `json:"offering_thumbnail,omitempty" tfsdk:"offering_thumbnail"`
	OfferingType               *string             `json:"offering_type,omitempty" tfsdk:"offering_type"`
	OfferingUuid               *string             `json:"offering_uuid,omitempty" tfsdk:"offering_uuid"`
	ParentName                 *string             `json:"parent_name,omitempty" tfsdk:"parent_name"`
	ParentOfferingName         *string             `json:"parent_offering_name,omitempty" tfsdk:"parent_offering_name"`
	ParentOfferingSlug         *string             `json:"parent_offering_slug,omitempty" tfsdk:"parent_offering_slug"`
	ParentOfferingUuid         *string             `json:"parent_offering_uuid,omitempty" tfsdk:"parent_offering_uuid"`
	ParentUuid                 *string             `json:"parent_uuid,omitempty" tfsdk:"parent_uuid"`
	Paused                     *bool               `json:"paused,omitempty" tfsdk:"paused"`
	Plan                       *string             `json:"plan,omitempty" tfsdk:"plan"`
	PlanDescription            *string             `json:"plan_description,omitempty" tfsdk:"plan_description"`
	PlanName                   *string             `json:"plan_name,omitempty" tfsdk:"plan_name"`
	PlanUnit                   *string             `json:"plan_unit,omitempty" tfsdk:"plan_unit"`
	PlanUuid                   *string             `json:"plan_uuid,omitempty" tfsdk:"plan_uuid"`
	Project                    *string             `json:"project,omitempty" tfsdk:"project"`
	ProjectDescription         *string             `json:"project_description,omitempty" tfsdk:"project_description"`
	ProjectEndDate             *string             `json:"project_end_date,omitempty" tfsdk:"project_end_date"`
	ProjectEndDateRequestedBy  *string             `json:"project_end_date_requested_by,omitempty" tfsdk:"project_end_date_requested_by"`
	ProjectName                *string             `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectSlug                *string             `json:"project_slug,omitempty" tfsdk:"project_slug"`
	ProjectUuid                *string             `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ProviderName               *string             `json:"provider_name,omitempty" tfsdk:"provider_name"`
	ProviderSlug               *string             `json:"provider_slug,omitempty" tfsdk:"provider_slug"`
	ProviderUuid               *string             `json:"provider_uuid,omitempty" tfsdk:"provider_uuid"`
	Report                     []ReportSection     `json:"report,omitempty" tfsdk:"report"`
	ResourceType               *string             `json:"resource_type,omitempty" tfsdk:"resource_type"`
	ResourceUuid               *string             `json:"resource_uuid,omitempty" tfsdk:"resource_uuid"`
	RestrictMemberAccess       *bool               `json:"restrict_member_access,omitempty" tfsdk:"restrict_member_access"`
	Scope                      *string             `json:"scope,omitempty" tfsdk:"scope"`
	ServiceSettingsUuid        *string             `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	Slots                      []BookingSlot       `json:"slots,omitempty" tfsdk:"slots"`
	Slug                       *string             `json:"slug,omitempty" tfsdk:"slug"`
	State                      *string             `json:"state,omitempty" tfsdk:"state"`
	Url                        *string             `json:"url,omitempty" tfsdk:"url"`
	UserRequiresReconsent      *bool               `json:"user_requires_reconsent,omitempty" tfsdk:"user_requires_reconsent"`
	Username                   *string             `json:"username,omitempty" tfsdk:"username"`
}

type BookingSlot struct {
	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	End       *string `json:"end,omitempty" tfsdk:"end"`
	Start     *string `json:"start,omitempty" tfsdk:"start"`
}

type BroadcastMessage struct {
	AuthorFullName *string `json:"author_full_name,omitempty" tfsdk:"author_full_name"`
	Body           *string `json:"body,omitempty" tfsdk:"body"`
	Created        *string `json:"created,omitempty" tfsdk:"created"`
	SendAt         *string `json:"send_at,omitempty" tfsdk:"send_at"`
	State          *string `json:"state,omitempty" tfsdk:"state"`
	Subject        *string `json:"subject,omitempty" tfsdk:"subject"`
}

type BroadcastMessageRequest struct {
	Body    *string `json:"body" tfsdk:"body"`
	SendAt  *string `json:"send_at,omitempty" tfsdk:"send_at"`
	Subject *string `json:"subject" tfsdk:"subject"`
}

type BroadcastMessageStateEnum struct {
}

type BulkSilenceResponse struct {
	Count        *int64  `json:"count" tfsdk:"count"`
	DurationDays *int64  `json:"duration_days,omitempty" tfsdk:"duration_days"`
	Status       *string `json:"status" tfsdk:"status"`
}

type CallAttachDocumentsRequest struct {
	Description *string  `json:"description,omitempty" tfsdk:"description"`
	Documents   []string `json:"documents" tfsdk:"documents"`
}

type CallComplianceOverview struct {
}

type CallComplianceReviewRequest struct {
	ProposalUuid *string `json:"proposal_uuid" tfsdk:"proposal_uuid"`
	ReviewNotes  *string `json:"review_notes,omitempty" tfsdk:"review_notes"`
}

type CallDetachDocumentsRequest struct {
	Documents []string `json:"documents" tfsdk:"documents"`
}

type CallDocument struct {
	Created     *string `json:"created,omitempty" tfsdk:"created"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	File        *string `json:"file,omitempty" tfsdk:"file"`
	FileName    *string `json:"file_name,omitempty" tfsdk:"file_name"`
	FileSize    *int64  `json:"file_size,omitempty" tfsdk:"file_size"`
}

type CallDocumentRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	File        *string `json:"file,omitempty" tfsdk:"file"`
}

type CallManagingOrganisation struct {
	Created              *string `json:"created" tfsdk:"created"`
	Customer             *string `json:"customer" tfsdk:"customer"`
	CustomerAbbreviation *string `json:"customer_abbreviation" tfsdk:"customer_abbreviation"`
	CustomerCountry      *string `json:"customer_country" tfsdk:"customer_country"`
	CustomerImage        *string `json:"customer_image" tfsdk:"customer_image"`
	CustomerName         *string `json:"customer_name" tfsdk:"customer_name"`
	CustomerNativeName   *string `json:"customer_native_name" tfsdk:"customer_native_name"`
	CustomerUuid         *string `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description          *string `json:"description,omitempty" tfsdk:"description"`
	Image                *string `json:"image,omitempty" tfsdk:"image"`
	Url                  *string `json:"url" tfsdk:"url"`
}

type CallManagingOrganisationRequest struct {
	Customer    *string `json:"customer" tfsdk:"customer"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Image       *string `json:"image,omitempty" tfsdk:"image"`
}

type CallManagingOrganisationRequestForm struct {
	Customer    *string `json:"customer" tfsdk:"customer"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Image       *string `json:"image,omitempty" tfsdk:"image"`
}

type CallManagingOrganisationRequestMultipart struct {
	Customer    *string `json:"customer" tfsdk:"customer"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Image       *string `json:"image,omitempty" tfsdk:"image"`
}

type CallManagingOrganisationStat struct {
	AcceptedProposals       *int64 `json:"accepted_proposals" tfsdk:"accepted_proposals"`
	ActiveRounds            *int64 `json:"active_rounds" tfsdk:"active_rounds"`
	CallsClosingInOneWeek   *int64 `json:"calls_closing_in_one_week" tfsdk:"calls_closing_in_one_week"`
	OfferingRequestsPending *int64 `json:"offering_requests_pending" tfsdk:"offering_requests_pending"`
	OpenCalls               *int64 `json:"open_calls" tfsdk:"open_calls"`
	PendingProposals        *int64 `json:"pending_proposals" tfsdk:"pending_proposals"`
	PendingReview           *int64 `json:"pending_review" tfsdk:"pending_review"`
	RoundsClosingInOneWeek  *int64 `json:"rounds_closing_in_one_week" tfsdk:"rounds_closing_in_one_week"`
}

type CallResourceTemplate struct {
	Created               *string `json:"created,omitempty" tfsdk:"created"`
	CreatedBy             *string `json:"created_by,omitempty" tfsdk:"created_by"`
	CreatedByName         *string `json:"created_by_name,omitempty" tfsdk:"created_by_name"`
	Description           *string `json:"description,omitempty" tfsdk:"description"`
	IsRequired            *bool   `json:"is_required,omitempty" tfsdk:"is_required"`
	Name                  *string `json:"name,omitempty" tfsdk:"name"`
	RequestedOffering     *string `json:"requested_offering,omitempty" tfsdk:"requested_offering"`
	RequestedOfferingName *string `json:"requested_offering_name,omitempty" tfsdk:"requested_offering_name"`
	RequestedOfferingUuid *string `json:"requested_offering_uuid,omitempty" tfsdk:"requested_offering_uuid"`
	Url                   *string `json:"url,omitempty" tfsdk:"url"`
}

type CallResourceTemplateRequest struct {
	Description       *string `json:"description,omitempty" tfsdk:"description"`
	IsRequired        *bool   `json:"is_required,omitempty" tfsdk:"is_required"`
	Name              *string `json:"name" tfsdk:"name"`
	RequestedOffering *string `json:"requested_offering" tfsdk:"requested_offering"`
}

type CallRound struct {
	CallName   *string `json:"call_name" tfsdk:"call_name"`
	CallUuid   *string `json:"call_uuid" tfsdk:"call_uuid"`
	CutoffTime *string `json:"cutoff_time" tfsdk:"cutoff_time"`
	Slug       *string `json:"slug,omitempty" tfsdk:"slug"`
	StartTime  *string `json:"start_time" tfsdk:"start_time"`
	Status     *string `json:"status" tfsdk:"status"`
	Url        *string `json:"url" tfsdk:"url"`
}

type CallStates struct {
}

type Campaign struct {
	AutoApply         *bool              `json:"auto_apply,omitempty" tfsdk:"auto_apply"`
	Coupon            *string            `json:"coupon,omitempty" tfsdk:"coupon"`
	Description       *string            `json:"description,omitempty" tfsdk:"description"`
	Discount          *int64             `json:"discount" tfsdk:"discount"`
	DiscountType      *string            `json:"discount_type" tfsdk:"discount_type"`
	EndDate           *string            `json:"end_date" tfsdk:"end_date"`
	Months            *int64             `json:"months,omitempty" tfsdk:"months"`
	Name              *string            `json:"name" tfsdk:"name"`
	Offerings         []CampaignOffering `json:"offerings" tfsdk:"offerings"`
	RequiredOfferings []string           `json:"required_offerings,omitempty" tfsdk:"required_offerings"`
	ServiceProvider   *string            `json:"service_provider" tfsdk:"service_provider"`
	StartDate         *string            `json:"start_date" tfsdk:"start_date"`
	State             *string            `json:"state" tfsdk:"state"`
	Stock             *int64             `json:"stock,omitempty" tfsdk:"stock"`
	Url               *string            `json:"url" tfsdk:"url"`
}

type CampaignOffering struct {
	Name *string `json:"name" tfsdk:"name"`
}

type CampaignRequest struct {
	AutoApply         *bool    `json:"auto_apply,omitempty" tfsdk:"auto_apply"`
	Coupon            *string  `json:"coupon,omitempty" tfsdk:"coupon"`
	Description       *string  `json:"description,omitempty" tfsdk:"description"`
	Discount          *int64   `json:"discount" tfsdk:"discount"`
	DiscountType      *string  `json:"discount_type" tfsdk:"discount_type"`
	EndDate           *string  `json:"end_date" tfsdk:"end_date"`
	Months            *int64   `json:"months,omitempty" tfsdk:"months"`
	Name              *string  `json:"name" tfsdk:"name"`
	Offerings         []string `json:"offerings" tfsdk:"offerings"`
	RequiredOfferings []string `json:"required_offerings,omitempty" tfsdk:"required_offerings"`
	ServiceProvider   *string  `json:"service_provider" tfsdk:"service_provider"`
	StartDate         *string  `json:"start_date" tfsdk:"start_date"`
	Stock             *int64   `json:"stock,omitempty" tfsdk:"stock"`
}

type CancelRequestResponse struct {
	ScopeName *string `json:"scope_name" tfsdk:"scope_name"`
	ScopeUuid *string `json:"scope_uuid" tfsdk:"scope_uuid"`
}

type CascadeConfig struct {
	Steps []CascadeStep `json:"steps" tfsdk:"steps"`
}

type CascadeConfigRequest struct {
	Steps []CascadeStepRequest `json:"steps" tfsdk:"steps"`
}

type CascadeStep struct {
	DependsOn *string `json:"depends_on,omitempty" tfsdk:"depends_on"`
	Label     *string `json:"label" tfsdk:"label"`
	Name      *string `json:"name" tfsdk:"name"`
	Type      *string `json:"type" tfsdk:"type"`
}

type CascadeStepRequest struct {
	DependsOn *string `json:"depends_on,omitempty" tfsdk:"depends_on"`
	Label     *string `json:"label" tfsdk:"label"`
	Name      *string `json:"name" tfsdk:"name"`
	Type      *string `json:"type" tfsdk:"type"`
}

type CascadeStepTypeEnum struct {
}

type CatalogTypeEnum struct {
}

type CategoryColumn struct {
	Attribute *string `json:"attribute,omitempty" tfsdk:"attribute"`
	Category  *string `json:"category" tfsdk:"category"`
	Index     *int64  `json:"index" tfsdk:"index"`
	Title     *string `json:"title" tfsdk:"title"`
	Widget    *string `json:"widget,omitempty" tfsdk:"widget"`
}

type CategoryColumnRequest struct {
	Attribute *string `json:"attribute,omitempty" tfsdk:"attribute"`
	Category  *string `json:"category" tfsdk:"category"`
	Index     *int64  `json:"index" tfsdk:"index"`
	Title     *string `json:"title" tfsdk:"title"`
	Widget    *string `json:"widget,omitempty" tfsdk:"widget"`
}

type CategoryComponent struct {
	Description  *string `json:"description,omitempty" tfsdk:"description"`
	MeasuredUnit *string `json:"measured_unit,omitempty" tfsdk:"measured_unit"`
	Name         *string `json:"name,omitempty" tfsdk:"name"`
	Type         *string `json:"type,omitempty" tfsdk:"type"`
}

type CategoryComponentRequest struct {
	Description  *string `json:"description,omitempty" tfsdk:"description"`
	MeasuredUnit *string `json:"measured_unit,omitempty" tfsdk:"measured_unit"`
	Name         *string `json:"name" tfsdk:"name"`
	Type         *string `json:"type" tfsdk:"type"`
}

type CategoryComponentUsage struct {
	CategoryTitle *string `json:"category_title,omitempty" tfsdk:"category_title"`
	CategoryUuid  *string `json:"category_uuid,omitempty" tfsdk:"category_uuid"`
	Date          *string `json:"date,omitempty" tfsdk:"date"`
	FixedUsage    *int64  `json:"fixed_usage,omitempty" tfsdk:"fixed_usage"`
	MeasuredUnit  *string `json:"measured_unit,omitempty" tfsdk:"measured_unit"`
	Name          *string `json:"name,omitempty" tfsdk:"name"`
	ReportedUsage *int64  `json:"reported_usage,omitempty" tfsdk:"reported_usage"`
	Scope         *string `json:"scope,omitempty" tfsdk:"scope"`
	Type          *string `json:"type,omitempty" tfsdk:"type"`
}

type CategoryComponents struct {
	Category     *CategorySerializerForForNestedFields `json:"category" tfsdk:"category"`
	Description  *string                               `json:"description,omitempty" tfsdk:"description"`
	MeasuredUnit *string                               `json:"measured_unit,omitempty" tfsdk:"measured_unit"`
	Name         *string                               `json:"name" tfsdk:"name"`
	Type         *string                               `json:"type" tfsdk:"type"`
}

type CategoryComponentsRequest struct {
	Category     *CategorySerializerForForNestedFieldsRequest `json:"category" tfsdk:"category"`
	Description  *string                                      `json:"description,omitempty" tfsdk:"description"`
	MeasuredUnit *string                                      `json:"measured_unit,omitempty" tfsdk:"measured_unit"`
	Name         *string                                      `json:"name" tfsdk:"name"`
	Type         *string                                      `json:"type" tfsdk:"type"`
}

type CategoryEnum struct {
}

type CategoryGroup struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Icon        *string `json:"icon,omitempty" tfsdk:"icon"`
	Title       *string `json:"title,omitempty" tfsdk:"title"`
	Url         *string `json:"url,omitempty" tfsdk:"url"`
}

type CategoryGroupRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Icon        *string `json:"icon,omitempty" tfsdk:"icon"`
	Title       *string `json:"title" tfsdk:"title"`
}

type CategoryGroupRequestForm struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Icon        *string `json:"icon,omitempty" tfsdk:"icon"`
	Title       *string `json:"title" tfsdk:"title"`
}

type CategoryGroupRequestMultipart struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Icon        *string `json:"icon,omitempty" tfsdk:"icon"`
	Title       *string `json:"title" tfsdk:"title"`
}

type CategoryHelpArticle struct {
	Title *string `json:"title,omitempty" tfsdk:"title"`
	Url   *string `json:"url,omitempty" tfsdk:"url"`
}

type CategoryHelpArticleRequest struct {
	Title *string `json:"title,omitempty" tfsdk:"title"`
	Url   *string `json:"url" tfsdk:"url"`
}

type CategoryHelpArticles struct {
	Categories []CategorySerializerForForNestedFields `json:"categories" tfsdk:"categories"`
	Title      *string                                `json:"title,omitempty" tfsdk:"title"`
	Url        *string                                `json:"url" tfsdk:"url"`
}

type CategoryHelpArticlesRequest struct {
	Categories []CategorySerializerForForNestedFieldsRequest `json:"categories" tfsdk:"categories"`
	Title      *string                                       `json:"title,omitempty" tfsdk:"title"`
	Url        *string                                       `json:"url" tfsdk:"url"`
}

type CategorySerializerForForNestedFields struct {
	Title *string `json:"title" tfsdk:"title"`
	Url   *string `json:"url" tfsdk:"url"`
}

type CategorySerializerForForNestedFieldsRequest struct {
	Title *string `json:"title" tfsdk:"title"`
}

type CeleryBroker struct {
	Alternates       []string `json:"alternates" tfsdk:"alternates"`
	ConnectTimeout   *int64   `json:"connect_timeout" tfsdk:"connect_timeout"`
	FailoverStrategy *string  `json:"failover_strategy" tfsdk:"failover_strategy"`
	Heartbeat        *float64 `json:"heartbeat" tfsdk:"heartbeat"`
	Hostname         *string  `json:"hostname" tfsdk:"hostname"`
	Insist           *bool    `json:"insist" tfsdk:"insist"`
	LoginMethod      *string  `json:"login_method" tfsdk:"login_method"`
	Port             *int64   `json:"port" tfsdk:"port"`
	Ssl              *bool    `json:"ssl" tfsdk:"ssl"`
	Transport        *string  `json:"transport" tfsdk:"transport"`
	UriPrefix        *string  `json:"uri_prefix" tfsdk:"uri_prefix"`
	Userid           *string  `json:"userid" tfsdk:"userid"`
	VirtualHost      *string  `json:"virtual_host" tfsdk:"virtual_host"`
}

type CeleryScheduledTask struct {
	Eta      *string `json:"eta" tfsdk:"eta"`
	Priority *int64  `json:"priority" tfsdk:"priority"`
}

type CeleryStatsResponse struct {
}

type CeleryTask struct {
	Acknowledged *bool    `json:"acknowledged" tfsdk:"acknowledged"`
	Hostname     *string  `json:"hostname" tfsdk:"hostname"`
	Id           *string  `json:"id" tfsdk:"id"`
	Name         *string  `json:"name" tfsdk:"name"`
	TimeStart    *float64 `json:"time_start" tfsdk:"time_start"`
	Type         *string  `json:"type" tfsdk:"type"`
	WorkerPid    *int64   `json:"worker_pid" tfsdk:"worker_pid"`
}

type CeleryWorkerPool struct {
	MaxConcurrency        *int64 `json:"max_concurrency" tfsdk:"max_concurrency"`
	MaxTasksPerChild      *int64 `json:"max_tasks_per_child" tfsdk:"max_tasks_per_child"`
	PutGuardedBySemaphore *bool  `json:"put_guarded_by_semaphore" tfsdk:"put_guarded_by_semaphore"`
}

type CeleryWorkerStats struct {
	Clock         *string  `json:"clock" tfsdk:"clock"`
	Pid           *int64   `json:"pid" tfsdk:"pid"`
	PrefetchCount *int64   `json:"prefetch_count" tfsdk:"prefetch_count"`
	Uptime        *float64 `json:"uptime" tfsdk:"uptime"`
}

type ChatRequestRequest struct {
	Input *string `json:"input" tfsdk:"input"`
}

type CheckUniqueBackendIDRequest struct {
	BackendId         *string `json:"backend_id" tfsdk:"backend_id"`
	CheckAllOfferings *bool   `json:"check_all_offerings,omitempty" tfsdk:"check_all_offerings"`
}

type CheckUniqueBackendIDResponse struct {
	IsUnique *bool `json:"is_unique" tfsdk:"is_unique"`
}

type Checklist struct {
	Category       *string `json:"category,omitempty" tfsdk:"category"`
	CategoryName   *string `json:"category_name" tfsdk:"category_name"`
	CategoryUuid   *string `json:"category_uuid" tfsdk:"category_uuid"`
	ChecklistType  *string `json:"checklist_type,omitempty" tfsdk:"checklist_type"`
	Description    *string `json:"description,omitempty" tfsdk:"description"`
	Name           *string `json:"name" tfsdk:"name"`
	QuestionsCount *int64  `json:"questions_count" tfsdk:"questions_count"`
	Url            *string `json:"url" tfsdk:"url"`
}

type ChecklistCategory struct {
	ChecklistsCount *int64  `json:"checklists_count" tfsdk:"checklists_count"`
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	Icon            *string `json:"icon,omitempty" tfsdk:"icon"`
	Name            *string `json:"name" tfsdk:"name"`
	Url             *string `json:"url" tfsdk:"url"`
}

type ChecklistCategoryRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Icon        *string `json:"icon,omitempty" tfsdk:"icon"`
	Name        *string `json:"name" tfsdk:"name"`
}

type ChecklistCategoryRequestForm struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Icon        *string `json:"icon,omitempty" tfsdk:"icon"`
	Name        *string `json:"name" tfsdk:"name"`
}

type ChecklistCategoryRequestMultipart struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Icon        *string `json:"icon,omitempty" tfsdk:"icon"`
	Name        *string `json:"name" tfsdk:"name"`
}

type ChecklistCompletion struct {
	ChecklistDescription *string  `json:"checklist_description" tfsdk:"checklist_description"`
	ChecklistName        *string  `json:"checklist_name" tfsdk:"checklist_name"`
	CompletionPercentage *float64 `json:"completion_percentage" tfsdk:"completion_percentage"`
	Created              *string  `json:"created" tfsdk:"created"`
	IsCompleted          *bool    `json:"is_completed" tfsdk:"is_completed"`
	Modified             *string  `json:"modified" tfsdk:"modified"`
}

type ChecklistCompletionReviewer struct {
	ChecklistDescription *string  `json:"checklist_description" tfsdk:"checklist_description"`
	ChecklistName        *string  `json:"checklist_name" tfsdk:"checklist_name"`
	CompletionPercentage *float64 `json:"completion_percentage" tfsdk:"completion_percentage"`
	Created              *string  `json:"created" tfsdk:"created"`
	IsCompleted          *bool    `json:"is_completed" tfsdk:"is_completed"`
	Modified             *string  `json:"modified" tfsdk:"modified"`
	RequiresReview       *bool    `json:"requires_review" tfsdk:"requires_review"`
	ReviewNotes          *string  `json:"review_notes,omitempty" tfsdk:"review_notes"`
	ReviewedAt           *string  `json:"reviewed_at,omitempty" tfsdk:"reviewed_at"`
	ReviewedBy           *int64   `json:"reviewed_by,omitempty" tfsdk:"reviewed_by"`
	ReviewedByName       *string  `json:"reviewed_by_name" tfsdk:"reviewed_by_name"`
}

type ChecklistInfo struct {
	ChecklistType *string `json:"checklist_type" tfsdk:"checklist_type"`
	Name          *string `json:"name" tfsdk:"name"`
}

type ChecklistOperators struct {
}

type ChecklistRequest struct {
	Category      *string `json:"category,omitempty" tfsdk:"category"`
	ChecklistType *string `json:"checklist_type,omitempty" tfsdk:"checklist_type"`
	Description   *string `json:"description,omitempty" tfsdk:"description"`
	Name          *string `json:"name" tfsdk:"name"`
}

type ChecklistResponse struct {
	Completion *ChecklistCompletion `json:"completion" tfsdk:"completion"`
	Questions  []QuestionWithAnswer `json:"questions" tfsdk:"questions"`
}

type ChecklistReviewerResponse struct {
	Completion *ChecklistCompletionReviewer `json:"completion" tfsdk:"completion"`
	Questions  []QuestionWithAnswerReviewer `json:"questions" tfsdk:"questions"`
}

type ChecklistTemplate struct {
	InitialVisibleQuestions []Question `json:"initial_visible_questions" tfsdk:"initial_visible_questions"`
	Questions               []Question `json:"questions" tfsdk:"questions"`
}

type ChecklistTypeEnum struct {
}

type ClusterSecurityGroup struct {
	Description *string                           `json:"description" tfsdk:"description"`
	Name        *string                           `json:"name" tfsdk:"name"`
	Rules       []RancherClusterSecurityGroupRule `json:"rules" tfsdk:"rules"`
}

type ClusterSecurityGroupRequest struct {
	Rules []RancherClusterSecurityGroupRuleRequest `json:"rules" tfsdk:"rules"`
}

type Comment struct {
	AuthorEmail        *string `json:"author_email" tfsdk:"author_email"`
	AuthorName         *string `json:"author_name" tfsdk:"author_name"`
	AuthorUser         *string `json:"author_user" tfsdk:"author_user"`
	AuthorUuid         *string `json:"author_uuid" tfsdk:"author_uuid"`
	BackendId          *string `json:"backend_id" tfsdk:"backend_id"`
	Created            *string `json:"created" tfsdk:"created"`
	Description        *string `json:"description" tfsdk:"description"`
	DestroyIsAvailable *bool   `json:"destroy_is_available" tfsdk:"destroy_is_available"`
	IsPublic           *bool   `json:"is_public,omitempty" tfsdk:"is_public"`
	Issue              *string `json:"issue" tfsdk:"issue"`
	IssueKey           *string `json:"issue_key" tfsdk:"issue_key"`
	RemoteId           *string `json:"remote_id,omitempty" tfsdk:"remote_id"`
	UpdateIsAvailable  *bool   `json:"update_is_available" tfsdk:"update_is_available"`
	Url                *string `json:"url" tfsdk:"url"`
}

type CommentRequest struct {
	Description *string `json:"description" tfsdk:"description"`
	IsPublic    *bool   `json:"is_public,omitempty" tfsdk:"is_public"`
}

type ComplianceOverview struct {
	AverageCompletionPercentage *float64 `json:"average_completion_percentage" tfsdk:"average_completion_percentage"`
	FullyCompletedProjects      *int64   `json:"fully_completed_projects" tfsdk:"fully_completed_projects"`
	ProjectsRequiringReview     *int64   `json:"projects_requiring_review" tfsdk:"projects_requiring_review"`
	ProjectsWithCompletions     *int64   `json:"projects_with_completions" tfsdk:"projects_with_completions"`
	TotalProjects               *int64   `json:"total_projects" tfsdk:"total_projects"`
}

type ComponentMultiplierConfig struct {
	ComponentType *string `json:"component_type" tfsdk:"component_type"`
	Factor        *int64  `json:"factor" tfsdk:"factor"`
	MaxLimit      *int64  `json:"max_limit,omitempty" tfsdk:"max_limit"`
	MinLimit      *int64  `json:"min_limit,omitempty" tfsdk:"min_limit"`
}

type ComponentMultiplierConfigRequest struct {
	ComponentType *string `json:"component_type" tfsdk:"component_type"`
	Factor        *int64  `json:"factor" tfsdk:"factor"`
	MaxLimit      *int64  `json:"max_limit,omitempty" tfsdk:"max_limit"`
	MinLimit      *int64  `json:"min_limit,omitempty" tfsdk:"min_limit"`
}

type ComponentStats struct {
	BillingType  *string `json:"billing_type" tfsdk:"billing_type"`
	Description  *string `json:"description" tfsdk:"description"`
	Limit        *int64  `json:"limit" tfsdk:"limit"`
	LimitUsage   *int64  `json:"limit_usage" tfsdk:"limit_usage"`
	MeasuredUnit *string `json:"measured_unit" tfsdk:"measured_unit"`
	Name         *string `json:"name" tfsdk:"name"`
	OfferingName *string `json:"offering_name" tfsdk:"offering_name"`
	OfferingUuid *string `json:"offering_uuid" tfsdk:"offering_uuid"`
	Type         *string `json:"type" tfsdk:"type"`
	Usage        *int64  `json:"usage" tfsdk:"usage"`
}

type ComponentUsage struct {
	BillingPeriod *string `json:"billing_period,omitempty" tfsdk:"billing_period"`
	Created       *string `json:"created,omitempty" tfsdk:"created"`
	CustomerName  *string `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerUuid  *string `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Date          *string `json:"date,omitempty" tfsdk:"date"`
	Description   *string `json:"description,omitempty" tfsdk:"description"`
	MeasuredUnit  *string `json:"measured_unit,omitempty" tfsdk:"measured_unit"`
	ModifiedBy    *int64  `json:"modified_by,omitempty" tfsdk:"modified_by"`
	Name          *string `json:"name,omitempty" tfsdk:"name"`
	OfferingName  *string `json:"offering_name,omitempty" tfsdk:"offering_name"`
	OfferingUuid  *string `json:"offering_uuid,omitempty" tfsdk:"offering_uuid"`
	ProjectName   *string `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid   *string `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	Recurring     *bool   `json:"recurring,omitempty" tfsdk:"recurring"`
	ResourceName  *string `json:"resource_name,omitempty" tfsdk:"resource_name"`
	ResourceUuid  *string `json:"resource_uuid,omitempty" tfsdk:"resource_uuid"`
	Type          *string `json:"type,omitempty" tfsdk:"type"`
	Usage         *int64  `json:"usage,omitempty" tfsdk:"usage"`
}

type ComponentUsageCreateRequest struct {
	Date       *string                     `json:"date,omitempty" tfsdk:"date"`
	PlanPeriod *string                     `json:"plan_period,omitempty" tfsdk:"plan_period"`
	Resource   *string                     `json:"resource,omitempty" tfsdk:"resource"`
	Usages     []ComponentUsageItemRequest `json:"usages" tfsdk:"usages"`
}

type ComponentUsageItemRequest struct {
	Amount      *string `json:"amount" tfsdk:"amount"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Recurring   *bool   `json:"recurring,omitempty" tfsdk:"recurring"`
	Type        *string `json:"type" tfsdk:"type"`
}

type ComponentUsagesPerMonthStats struct {
	ComponentType         *string `json:"component_type" tfsdk:"component_type"`
	Month                 *int64  `json:"month" tfsdk:"month"`
	OfferingCountry       *string `json:"offering_country" tfsdk:"offering_country"`
	OfferingUuid          *string `json:"offering_uuid" tfsdk:"offering_uuid"`
	OrganizationGroupName *string `json:"organization_group_name" tfsdk:"organization_group_name"`
	OrganizationGroupUuid *string `json:"organization_group_uuid" tfsdk:"organization_group_uuid"`
	Usage                 *string `json:"usage" tfsdk:"usage"`
	Year                  *int64  `json:"year" tfsdk:"year"`
}

type ComponentUsagesPerProject struct {
	ComponentType *string `json:"component_type" tfsdk:"component_type"`
	ProjectUuid   *string `json:"project_uuid" tfsdk:"project_uuid"`
	Usage         *int64  `json:"usage" tfsdk:"usage"`
}

type ComponentUsagesStats struct {
	ComponentType         *string `json:"component_type" tfsdk:"component_type"`
	OfferingCountry       *string `json:"offering_country" tfsdk:"offering_country"`
	OfferingUuid          *string `json:"offering_uuid" tfsdk:"offering_uuid"`
	OrganizationGroupName *string `json:"organization_group_name" tfsdk:"organization_group_name"`
	OrganizationGroupUuid *string `json:"organization_group_uuid" tfsdk:"organization_group_uuid"`
	Usage                 *string `json:"usage" tfsdk:"usage"`
}

type ComponentUserUsage struct {
	BackendId      *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	BillingPeriod  *string `json:"billing_period,omitempty" tfsdk:"billing_period"`
	ComponentType  *string `json:"component_type,omitempty" tfsdk:"component_type"`
	ComponentUsage *string `json:"component_usage,omitempty" tfsdk:"component_usage"`
	Created        *string `json:"created,omitempty" tfsdk:"created"`
	CustomerName   *string `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerUuid   *string `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Date           *string `json:"date,omitempty" tfsdk:"date"`
	Description    *string `json:"description,omitempty" tfsdk:"description"`
	MeasuredUnit   *string `json:"measured_unit,omitempty" tfsdk:"measured_unit"`
	Modified       *string `json:"modified,omitempty" tfsdk:"modified"`
	OfferingName   *string `json:"offering_name,omitempty" tfsdk:"offering_name"`
	OfferingUuid   *string `json:"offering_uuid,omitempty" tfsdk:"offering_uuid"`
	ProjectName    *string `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid    *string `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ResourceName   *string `json:"resource_name,omitempty" tfsdk:"resource_name"`
	ResourceUuid   *string `json:"resource_uuid,omitempty" tfsdk:"resource_uuid"`
	Usage          *int64  `json:"usage,omitempty" tfsdk:"usage"`
	User           *string `json:"user,omitempty" tfsdk:"user"`
	Username       *string `json:"username,omitempty" tfsdk:"username"`
}

type ComponentUserUsageCreateRequest struct {
	Date     *string `json:"date,omitempty" tfsdk:"date"`
	Usage    *string `json:"usage,omitempty" tfsdk:"usage"`
	User     *string `json:"user,omitempty" tfsdk:"user"`
	Username *string `json:"username" tfsdk:"username"`
}

type ComponentUserUsageLimit struct {
	Component     *string `json:"component" tfsdk:"component"`
	ComponentType *string `json:"component_type" tfsdk:"component_type"`
	Limit         *string `json:"limit,omitempty" tfsdk:"limit"`
	Resource      *string `json:"resource" tfsdk:"resource"`
	Url           *string `json:"url" tfsdk:"url"`
	User          *string `json:"user" tfsdk:"user"`
}

type ComponentUserUsageLimitRequest struct {
	Component *string `json:"component" tfsdk:"component"`
	Limit     *string `json:"limit,omitempty" tfsdk:"limit"`
	Resource  *string `json:"resource" tfsdk:"resource"`
	User      *string `json:"user" tfsdk:"user"`
}

type ComponentsUsageStats struct {
	Components []ComponentStats `json:"components" tfsdk:"components"`
}

type ConfirmEmailRequestRequest struct {
	Code *string `json:"code" tfsdk:"code"`
}

type ConsoleUrl struct {
	Url *string `json:"url" tfsdk:"url"`
}

type ConstanceSettings struct {
	ANONYMOUSUSERCANVIEWOFFERINGS                  *bool    `json:"ANONYMOUS_USER_CAN_VIEW_OFFERINGS,omitempty" tfsdk:"ANONYMOUS_USER_CAN_VIEW_OFFERINGS"`
	ANONYMOUSUSERCANVIEWPLANS                      *bool    `json:"ANONYMOUS_USER_CAN_VIEW_PLANS,omitempty" tfsdk:"ANONYMOUS_USER_CAN_VIEW_PLANS"`
	ATLASSIANAFFECTEDRESOURCEFIELD                 *string  `json:"ATLASSIAN_AFFECTED_RESOURCE_FIELD,omitempty" tfsdk:"ATLASSIAN_AFFECTED_RESOURCE_FIELD"`
	ATLASSIANAPIURL                                *string  `json:"ATLASSIAN_API_URL,omitempty" tfsdk:"ATLASSIAN_API_URL"`
	ATLASSIANCALLERFIELD                           *string  `json:"ATLASSIAN_CALLER_FIELD,omitempty" tfsdk:"ATLASSIAN_CALLER_FIELD"`
	ATLASSIANCUSTOMISSUEFIELDMAPPINGENABLED        *bool    `json:"ATLASSIAN_CUSTOM_ISSUE_FIELD_MAPPING_ENABLED,omitempty" tfsdk:"ATLASSIAN_CUSTOM_ISSUE_FIELD_MAPPING_ENABLED"`
	ATLASSIANDEFAULTOFFERINGISSUETYPE              *string  `json:"ATLASSIAN_DEFAULT_OFFERING_ISSUE_TYPE,omitempty" tfsdk:"ATLASSIAN_DEFAULT_OFFERING_ISSUE_TYPE"`
	ATLASSIANDESCRIPTIONTEMPLATE                   *string  `json:"ATLASSIAN_DESCRIPTION_TEMPLATE,omitempty" tfsdk:"ATLASSIAN_DESCRIPTION_TEMPLATE"`
	ATLASSIANEMAIL                                 *string  `json:"ATLASSIAN_EMAIL,omitempty" tfsdk:"ATLASSIAN_EMAIL"`
	ATLASSIANEXCLUDEDATTACHMENTTYPES               *string  `json:"ATLASSIAN_EXCLUDED_ATTACHMENT_TYPES,omitempty" tfsdk:"ATLASSIAN_EXCLUDED_ATTACHMENT_TYPES"`
	ATLASSIANIMPACTFIELD                           *string  `json:"ATLASSIAN_IMPACT_FIELD,omitempty" tfsdk:"ATLASSIAN_IMPACT_FIELD"`
	ATLASSIANLINKEDISSUETYPE                       *string  `json:"ATLASSIAN_LINKED_ISSUE_TYPE,omitempty" tfsdk:"ATLASSIAN_LINKED_ISSUE_TYPE"`
	ATLASSIANMAPWALDURUSERSTOSERVICEDESKAGENTS     *bool    `json:"ATLASSIAN_MAP_WALDUR_USERS_TO_SERVICEDESK_AGENTS,omitempty" tfsdk:"ATLASSIAN_MAP_WALDUR_USERS_TO_SERVICEDESK_AGENTS"`
	ATLASSIANOAUTH2ACCESSTOKEN                     *string  `json:"ATLASSIAN_OAUTH2_ACCESS_TOKEN,omitempty" tfsdk:"ATLASSIAN_OAUTH2_ACCESS_TOKEN"`
	ATLASSIANOAUTH2CLIENTID                        *string  `json:"ATLASSIAN_OAUTH2_CLIENT_ID,omitempty" tfsdk:"ATLASSIAN_OAUTH2_CLIENT_ID"`
	ATLASSIANOAUTH2TOKENTYPE                       *string  `json:"ATLASSIAN_OAUTH2_TOKEN_TYPE,omitempty" tfsdk:"ATLASSIAN_OAUTH2_TOKEN_TYPE"`
	ATLASSIANORGANISATIONFIELD                     *string  `json:"ATLASSIAN_ORGANISATION_FIELD,omitempty" tfsdk:"ATLASSIAN_ORGANISATION_FIELD"`
	ATLASSIANPASSWORD                              *string  `json:"ATLASSIAN_PASSWORD,omitempty" tfsdk:"ATLASSIAN_PASSWORD"`
	ATLASSIANPERSONALACCESSTOKEN                   *string  `json:"ATLASSIAN_PERSONAL_ACCESS_TOKEN,omitempty" tfsdk:"ATLASSIAN_PERSONAL_ACCESS_TOKEN"`
	ATLASSIANPROJECTFIELD                          *string  `json:"ATLASSIAN_PROJECT_FIELD,omitempty" tfsdk:"ATLASSIAN_PROJECT_FIELD"`
	ATLASSIANPROJECTID                             *string  `json:"ATLASSIAN_PROJECT_ID,omitempty" tfsdk:"ATLASSIAN_PROJECT_ID"`
	ATLASSIANREPORTERFIELD                         *string  `json:"ATLASSIAN_REPORTER_FIELD,omitempty" tfsdk:"ATLASSIAN_REPORTER_FIELD"`
	ATLASSIANREQUESTFEEDBACKFIELD                  *string  `json:"ATLASSIAN_REQUEST_FEEDBACK_FIELD,omitempty" tfsdk:"ATLASSIAN_REQUEST_FEEDBACK_FIELD"`
	ATLASSIANRESOLUTIONSLAFIELD                    *string  `json:"ATLASSIAN_RESOLUTION_SLA_FIELD,omitempty" tfsdk:"ATLASSIAN_RESOLUTION_SLA_FIELD"`
	ATLASSIANSATISFACTIONFIELD                     *string  `json:"ATLASSIAN_SATISFACTION_FIELD,omitempty" tfsdk:"ATLASSIAN_SATISFACTION_FIELD"`
	ATLASSIANSHAREDUSERNAME                        *bool    `json:"ATLASSIAN_SHARED_USERNAME,omitempty" tfsdk:"ATLASSIAN_SHARED_USERNAME"`
	ATLASSIANSLAFIELD                              *string  `json:"ATLASSIAN_SLA_FIELD,omitempty" tfsdk:"ATLASSIAN_SLA_FIELD"`
	ATLASSIANSUMMARYTEMPLATE                       *string  `json:"ATLASSIAN_SUMMARY_TEMPLATE,omitempty" tfsdk:"ATLASSIAN_SUMMARY_TEMPLATE"`
	ATLASSIANTEMPLATEFIELD                         *string  `json:"ATLASSIAN_TEMPLATE_FIELD,omitempty" tfsdk:"ATLASSIAN_TEMPLATE_FIELD"`
	ATLASSIANTOKEN                                 *string  `json:"ATLASSIAN_TOKEN,omitempty" tfsdk:"ATLASSIAN_TOKEN"`
	ATLASSIANUSERNAME                              *string  `json:"ATLASSIAN_USERNAME,omitempty" tfsdk:"ATLASSIAN_USERNAME"`
	ATLASSIANUSEOLDAPI                             *bool    `json:"ATLASSIAN_USE_OLD_API,omitempty" tfsdk:"ATLASSIAN_USE_OLD_API"`
	ATLASSIANVERIFYSSL                             *bool    `json:"ATLASSIAN_VERIFY_SSL,omitempty" tfsdk:"ATLASSIAN_VERIFY_SSL"`
	ATLASSIANWALDURBACKENDIDFIELD                  *string  `json:"ATLASSIAN_WALDUR_BACKEND_ID_FIELD,omitempty" tfsdk:"ATLASSIAN_WALDUR_BACKEND_ID_FIELD"`
	AUTOAPPROVEUSERTOS                             *bool    `json:"AUTO_APPROVE_USER_TOS,omitempty" tfsdk:"AUTO_APPROVE_USER_TOS"`
	BRANDCOLOR                                     *string  `json:"BRAND_COLOR,omitempty" tfsdk:"BRAND_COLOR"`
	CALLMANAGEMENTHEROIMAGE                        *string  `json:"CALL_MANAGEMENT_HERO_IMAGE,omitempty" tfsdk:"CALL_MANAGEMENT_HERO_IMAGE"`
	COMMONFOOTERHTML                               *string  `json:"COMMON_FOOTER_HTML,omitempty" tfsdk:"COMMON_FOOTER_HTML"`
	COMMONFOOTERTEXT                               *string  `json:"COMMON_FOOTER_TEXT,omitempty" tfsdk:"COMMON_FOOTER_TEXT"`
	COUNTRIES                                      []string `json:"COUNTRIES,omitempty" tfsdk:"COUNTRIES"`
	CURRENCYNAME                                   *string  `json:"CURRENCY_NAME,omitempty" tfsdk:"CURRENCY_NAME"`
	DEACTIVATEUSERIFNOROLES                        *bool    `json:"DEACTIVATE_USER_IF_NO_ROLES,omitempty" tfsdk:"DEACTIVATE_USER_IF_NO_ROLES"`
	DEFAULTIDP                                     *string  `json:"DEFAULT_IDP,omitempty" tfsdk:"DEFAULT_IDP"`
	DISABLEDOFFERINGTYPES                          []string `json:"DISABLED_OFFERING_TYPES,omitempty" tfsdk:"DISABLED_OFFERING_TYPES"`
	DISABLEDARKTHEME                               *bool    `json:"DISABLE_DARK_THEME,omitempty" tfsdk:"DISABLE_DARK_THEME"`
	DISABLESENDINGNOTIFICATIONSABOUTRESOURCEUPDATE *bool    `json:"DISABLE_SENDING_NOTIFICATIONS_ABOUT_RESOURCE_UPDATE,omitempty" tfsdk:"DISABLE_SENDING_NOTIFICATIONS_ABOUT_RESOURCE_UPDATE"`
	DOCKERCLIENT                                   *string  `json:"DOCKER_CLIENT,omitempty" tfsdk:"DOCKER_CLIENT"`
	DOCKERIMAGES                                   *string  `json:"DOCKER_IMAGES,omitempty" tfsdk:"DOCKER_IMAGES"`
	DOCKERREMOVECONTAINER                          *bool    `json:"DOCKER_REMOVE_CONTAINER,omitempty" tfsdk:"DOCKER_REMOVE_CONTAINER"`
	DOCKERRUNOPTIONS                               *string  `json:"DOCKER_RUN_OPTIONS,omitempty" tfsdk:"DOCKER_RUN_OPTIONS"`
	DOCKERSCRIPTDIR                                *string  `json:"DOCKER_SCRIPT_DIR,omitempty" tfsdk:"DOCKER_SCRIPT_DIR"`
	DOCKERVOLUMENAME                               *string  `json:"DOCKER_VOLUME_NAME,omitempty" tfsdk:"DOCKER_VOLUME_NAME"`
	DOCSURL                                        *string  `json:"DOCS_URL,omitempty" tfsdk:"DOCS_URL"`
	ENABLEMOCKCOURSEACCOUNTBACKEND                 *bool    `json:"ENABLE_MOCK_COURSE_ACCOUNT_BACKEND,omitempty" tfsdk:"ENABLE_MOCK_COURSE_ACCOUNT_BACKEND"`
	ENABLEMOCKSERVICEACCOUNTBACKEND                *bool    `json:"ENABLE_MOCK_SERVICE_ACCOUNT_BACKEND,omitempty" tfsdk:"ENABLE_MOCK_SERVICE_ACCOUNT_BACKEND"`
	ENABLEORDERSTARTDATE                           *bool    `json:"ENABLE_ORDER_START_DATE,omitempty" tfsdk:"ENABLE_ORDER_START_DATE"`
	ENABLESTALERESOURCENOTIFICATIONS               *bool    `json:"ENABLE_STALE_RESOURCE_NOTIFICATIONS,omitempty" tfsdk:"ENABLE_STALE_RESOURCE_NOTIFICATIONS"`
	ENABLESTRICTCHECKACCEPTINGINVITATION           *bool    `json:"ENABLE_STRICT_CHECK_ACCEPTING_INVITATION,omitempty" tfsdk:"ENABLE_STRICT_CHECK_ACCEPTING_INVITATION"`
	ENFORCEUSERCONSENTFOROFFERINGS                 *bool    `json:"ENFORCE_USER_CONSENT_FOR_OFFERINGS,omitempty" tfsdk:"ENFORCE_USER_CONSENT_FOR_OFFERINGS"`
	FAVICON                                        *string  `json:"FAVICON,omitempty" tfsdk:"FAVICON"`
	FREEIPABLACKLISTEDUSERNAMES                    []string `json:"FREEIPA_BLACKLISTED_USERNAMES,omitempty" tfsdk:"FREEIPA_BLACKLISTED_USERNAMES"`
	FREEIPAENABLED                                 *bool    `json:"FREEIPA_ENABLED,omitempty" tfsdk:"FREEIPA_ENABLED"`
	FREEIPAGROUPNAMEPREFIX                         *string  `json:"FREEIPA_GROUPNAME_PREFIX,omitempty" tfsdk:"FREEIPA_GROUPNAME_PREFIX"`
	FREEIPAGROUPSYNCHRONIZATIONENABLED             *bool    `json:"FREEIPA_GROUP_SYNCHRONIZATION_ENABLED,omitempty" tfsdk:"FREEIPA_GROUP_SYNCHRONIZATION_ENABLED"`
	FREEIPAHOSTNAME                                *string  `json:"FREEIPA_HOSTNAME,omitempty" tfsdk:"FREEIPA_HOSTNAME"`
	FREEIPAPASSWORD                                *string  `json:"FREEIPA_PASSWORD,omitempty" tfsdk:"FREEIPA_PASSWORD"`
	FREEIPAUSERNAME                                *string  `json:"FREEIPA_USERNAME,omitempty" tfsdk:"FREEIPA_USERNAME"`
	FREEIPAUSERNAMEPREFIX                          *string  `json:"FREEIPA_USERNAME_PREFIX,omitempty" tfsdk:"FREEIPA_USERNAME_PREFIX"`
	FREEIPAVERIFYSSL                               *bool    `json:"FREEIPA_VERIFY_SSL,omitempty" tfsdk:"FREEIPA_VERIFY_SSL"`
	FULLPAGETITLE                                  *string  `json:"FULL_PAGE_TITLE,omitempty" tfsdk:"FULL_PAGE_TITLE"`
	HEROIMAGE                                      *string  `json:"HERO_IMAGE,omitempty" tfsdk:"HERO_IMAGE"`
	HEROLINKLABEL                                  *string  `json:"HERO_LINK_LABEL,omitempty" tfsdk:"HERO_LINK_LABEL"`
	HEROLINKURL                                    *string  `json:"HERO_LINK_URL,omitempty" tfsdk:"HERO_LINK_URL"`
	HOMEPORTURL                                    *string  `json:"HOMEPORT_URL,omitempty" tfsdk:"HOMEPORT_URL"`
	INVITATIONDISABLEMULTIPLEROLES                 *bool    `json:"INVITATION_DISABLE_MULTIPLE_ROLES,omitempty" tfsdk:"INVITATION_DISABLE_MULTIPLE_ROLES"`
	K8SCONFIGPATH                                  *string  `json:"K8S_CONFIG_PATH,omitempty" tfsdk:"K8S_CONFIG_PATH"`
	K8SJOBTIMEOUT                                  *int64   `json:"K8S_JOB_TIMEOUT,omitempty" tfsdk:"K8S_JOB_TIMEOUT"`
	K8SNAMESPACE                                   *string  `json:"K8S_NAMESPACE,omitempty" tfsdk:"K8S_NAMESPACE"`
	KEYCLOAKICON                                   *string  `json:"KEYCLOAK_ICON,omitempty" tfsdk:"KEYCLOAK_ICON"`
	LANGUAGECHOICES                                *string  `json:"LANGUAGE_CHOICES,omitempty" tfsdk:"LANGUAGE_CHOICES"`
	LLMCHATENABLED                                 *bool    `json:"LLM_CHAT_ENABLED,omitempty" tfsdk:"LLM_CHAT_ENABLED"`
	LLMINFERENCESAPITOKEN                          *string  `json:"LLM_INFERENCES_API_TOKEN,omitempty" tfsdk:"LLM_INFERENCES_API_TOKEN"`
	LLMINFERENCESAPIURL                            *string  `json:"LLM_INFERENCES_API_URL,omitempty" tfsdk:"LLM_INFERENCES_API_URL"`
	LLMINFERENCESBACKENDTYPE                       *string  `json:"LLM_INFERENCES_BACKEND_TYPE,omitempty" tfsdk:"LLM_INFERENCES_BACKEND_TYPE"`
	LLMINFERENCESMODEL                             *string  `json:"LLM_INFERENCES_MODEL,omitempty" tfsdk:"LLM_INFERENCES_MODEL"`
	LOGINLOGO                                      *string  `json:"LOGIN_LOGO,omitempty" tfsdk:"LOGIN_LOGO"`
	LOGINPAGELAYOUT                                *string  `json:"LOGIN_PAGE_LAYOUT,omitempty" tfsdk:"LOGIN_PAGE_LAYOUT"`
	LOGINPAGEVIDEOURL                              *string  `json:"LOGIN_PAGE_VIDEO_URL,omitempty" tfsdk:"LOGIN_PAGE_VIDEO_URL"`
	MAINTENANCEANNOUNCEMENTNOTIFYBEFOREMINUTES     *int64   `json:"MAINTENANCE_ANNOUNCEMENT_NOTIFY_BEFORE_MINUTES,omitempty" tfsdk:"MAINTENANCE_ANNOUNCEMENT_NOTIFY_BEFORE_MINUTES"`
	MAINTENANCEANNOUNCEMENTNOTIFYSYSTEM            []string `json:"MAINTENANCE_ANNOUNCEMENT_NOTIFY_SYSTEM,omitempty" tfsdk:"MAINTENANCE_ANNOUNCEMENT_NOTIFY_SYSTEM"`
	MARKETPLACEHEROIMAGE                           *string  `json:"MARKETPLACE_HERO_IMAGE,omitempty" tfsdk:"MARKETPLACE_HERO_IMAGE"`
	MARKETPLACELANDINGPAGE                         *string  `json:"MARKETPLACE_LANDING_PAGE,omitempty" tfsdk:"MARKETPLACE_LANDING_PAGE"`
	NOTIFYABOUTRESOURCECHANGE                      *bool    `json:"NOTIFY_ABOUT_RESOURCE_CHANGE,omitempty" tfsdk:"NOTIFY_ABOUT_RESOURCE_CHANGE"`
	NOTIFYSTAFFABOUTAPPROVALS                      *bool    `json:"NOTIFY_STAFF_ABOUT_APPROVALS,omitempty" tfsdk:"NOTIFY_STAFF_ABOUT_APPROVALS"`
	OFFERINGLOGOPLACEHOLDER                        *string  `json:"OFFERING_LOGO_PLACEHOLDER,omitempty" tfsdk:"OFFERING_LOGO_PLACEHOLDER"`
	OIDCACCESSTOKENENABLED                         *bool    `json:"OIDC_ACCESS_TOKEN_ENABLED,omitempty" tfsdk:"OIDC_ACCESS_TOKEN_ENABLED"`
	OIDCAUTHURL                                    *string  `json:"OIDC_AUTH_URL,omitempty" tfsdk:"OIDC_AUTH_URL"`
	OIDCBLOCKCREATIONOFUNINVITEDUSERS              *bool    `json:"OIDC_BLOCK_CREATION_OF_UNINVITED_USERS,omitempty" tfsdk:"OIDC_BLOCK_CREATION_OF_UNINVITED_USERS"`
	OIDCCACHETIMEOUT                               *int64   `json:"OIDC_CACHE_TIMEOUT,omitempty" tfsdk:"OIDC_CACHE_TIMEOUT"`
	OIDCCLIENTID                                   *string  `json:"OIDC_CLIENT_ID,omitempty" tfsdk:"OIDC_CLIENT_ID"`
	OIDCCLIENTSECRET                               *string  `json:"OIDC_CLIENT_SECRET,omitempty" tfsdk:"OIDC_CLIENT_SECRET"`
	OIDCINTROSPECTIONURL                           *string  `json:"OIDC_INTROSPECTION_URL,omitempty" tfsdk:"OIDC_INTROSPECTION_URL"`
	OIDCUSERFIELD                                  *string  `json:"OIDC_USER_FIELD,omitempty" tfsdk:"OIDC_USER_FIELD"`
	ONBOARDINGARIREGISTERBASEURL                   *string  `json:"ONBOARDING_ARIREGISTER_BASE_URL,omitempty" tfsdk:"ONBOARDING_ARIREGISTER_BASE_URL"`
	ONBOARDINGARIREGISTERPASSWORD                  *string  `json:"ONBOARDING_ARIREGISTER_PASSWORD,omitempty" tfsdk:"ONBOARDING_ARIREGISTER_PASSWORD"`
	ONBOARDINGARIREGISTERTIMEOUT                   *int64   `json:"ONBOARDING_ARIREGISTER_TIMEOUT,omitempty" tfsdk:"ONBOARDING_ARIREGISTER_TIMEOUT"`
	ONBOARDINGARIREGISTERUSERNAME                  *string  `json:"ONBOARDING_ARIREGISTER_USERNAME,omitempty" tfsdk:"ONBOARDING_ARIREGISTER_USERNAME"`
	ONBOARDINGBOLAGSVERKETAPIURL                   *string  `json:"ONBOARDING_BOLAGSVERKET_API_URL,omitempty" tfsdk:"ONBOARDING_BOLAGSVERKET_API_URL"`
	ONBOARDINGBOLAGSVERKETCLIENTID                 *string  `json:"ONBOARDING_BOLAGSVERKET_CLIENT_ID,omitempty" tfsdk:"ONBOARDING_BOLAGSVERKET_CLIENT_ID"`
	ONBOARDINGBOLAGSVERKETCLIENTSECRET             *string  `json:"ONBOARDING_BOLAGSVERKET_CLIENT_SECRET,omitempty" tfsdk:"ONBOARDING_BOLAGSVERKET_CLIENT_SECRET"`
	ONBOARDINGBOLAGSVERKETTOKENAPIURL              *string  `json:"ONBOARDING_BOLAGSVERKET_TOKEN_API_URL,omitempty" tfsdk:"ONBOARDING_BOLAGSVERKET_TOKEN_API_URL"`
	ONBOARDINGBREGAPIURL                           *string  `json:"ONBOARDING_BREG_API_URL,omitempty" tfsdk:"ONBOARDING_BREG_API_URL"`
	ONBOARDINGCOUNTRY                              *string  `json:"ONBOARDING_COUNTRY,omitempty" tfsdk:"ONBOARDING_COUNTRY"`
	ONBOARDINGVERIFICATIONEXPIRYHOURS              *int64   `json:"ONBOARDING_VERIFICATION_EXPIRY_HOURS,omitempty" tfsdk:"ONBOARDING_VERIFICATION_EXPIRY_HOURS"`
	ONBOARDINGWICOAPIURL                           *string  `json:"ONBOARDING_WICO_API_URL,omitempty" tfsdk:"ONBOARDING_WICO_API_URL"`
	ONBOARDINGWICOTOKEN                            *string  `json:"ONBOARDING_WICO_TOKEN,omitempty" tfsdk:"ONBOARDING_WICO_TOKEN"`
	POWEREDBYLOGO                                  *string  `json:"POWERED_BY_LOGO,omitempty" tfsdk:"POWERED_BY_LOGO"`
	PROJECTENDDATEMANDATORY                        *bool    `json:"PROJECT_END_DATE_MANDATORY,omitempty" tfsdk:"PROJECT_END_DATE_MANDATORY"`
	PROPOSALREVIEWDURATION                         *int64   `json:"PROPOSAL_REVIEW_DURATION,omitempty" tfsdk:"PROPOSAL_REVIEW_DURATION"`
	RANCHERUSERNAMEINPUTLABEL                      *string  `json:"RANCHER_USERNAME_INPUT_LABEL,omitempty" tfsdk:"RANCHER_USERNAME_INPUT_LABEL"`
	SCRIPTRUNMODE                                  *string  `json:"SCRIPT_RUN_MODE,omitempty" tfsdk:"SCRIPT_RUN_MODE"`
	SHORTPAGETITLE                                 *string  `json:"SHORT_PAGE_TITLE,omitempty" tfsdk:"SHORT_PAGE_TITLE"`
	SIDEBARLOGO                                    *string  `json:"SIDEBAR_LOGO,omitempty" tfsdk:"SIDEBAR_LOGO"`
	SIDEBARLOGODARK                                *string  `json:"SIDEBAR_LOGO_DARK,omitempty" tfsdk:"SIDEBAR_LOGO_DARK"`
	SIDEBARLOGOMOBILE                              *string  `json:"SIDEBAR_LOGO_MOBILE,omitempty" tfsdk:"SIDEBAR_LOGO_MOBILE"`
	SIDEBARSTYLE                                   *string  `json:"SIDEBAR_STYLE,omitempty" tfsdk:"SIDEBAR_STYLE"`
	SITEADDRESS                                    *string  `json:"SITE_ADDRESS,omitempty" tfsdk:"SITE_ADDRESS"`
	SITEDESCRIPTION                                *string  `json:"SITE_DESCRIPTION,omitempty" tfsdk:"SITE_DESCRIPTION"`
	SITEEMAIL                                      *string  `json:"SITE_EMAIL,omitempty" tfsdk:"SITE_EMAIL"`
	SITELOGO                                       *string  `json:"SITE_LOGO,omitempty" tfsdk:"SITE_LOGO"`
	SITENAME                                       *string  `json:"SITE_NAME,omitempty" tfsdk:"SITE_NAME"`
	SITEPHONE                                      *string  `json:"SITE_PHONE,omitempty" tfsdk:"SITE_PHONE"`
	SMAXAFFECTEDRESOURCEFIELD                      *string  `json:"SMAX_AFFECTED_RESOURCE_FIELD,omitempty" tfsdk:"SMAX_AFFECTED_RESOURCE_FIELD"`
	SMAXAPIURL                                     *string  `json:"SMAX_API_URL,omitempty" tfsdk:"SMAX_API_URL"`
	SMAXCREATIONSOURCENAME                         *string  `json:"SMAX_CREATION_SOURCE_NAME,omitempty" tfsdk:"SMAX_CREATION_SOURCE_NAME"`
	SMAXLOGIN                                      *string  `json:"SMAX_LOGIN,omitempty" tfsdk:"SMAX_LOGIN"`
	SMAXORGANISATIONFIELD                          *string  `json:"SMAX_ORGANISATION_FIELD,omitempty" tfsdk:"SMAX_ORGANISATION_FIELD"`
	SMAXPASSWORD                                   *string  `json:"SMAX_PASSWORD,omitempty" tfsdk:"SMAX_PASSWORD"`
	SMAXPROJECTFIELD                               *string  `json:"SMAX_PROJECT_FIELD,omitempty" tfsdk:"SMAX_PROJECT_FIELD"`
	SMAXREQUESTSOFFERING                           *string  `json:"SMAX_REQUESTS_OFFERING,omitempty" tfsdk:"SMAX_REQUESTS_OFFERING"`
	SMAXSECONDSTOWAIT                              *int64   `json:"SMAX_SECONDS_TO_WAIT,omitempty" tfsdk:"SMAX_SECONDS_TO_WAIT"`
	SMAXTENANTID                                   *string  `json:"SMAX_TENANT_ID,omitempty" tfsdk:"SMAX_TENANT_ID"`
	SMAXTIMESTOPULL                                *int64   `json:"SMAX_TIMES_TO_PULL,omitempty" tfsdk:"SMAX_TIMES_TO_PULL"`
	SMAXVERIFYSSL                                  *bool    `json:"SMAX_VERIFY_SSL,omitempty" tfsdk:"SMAX_VERIFY_SSL"`
	SOFTWARECATALOGCLEANUPENABLED                  *bool    `json:"SOFTWARE_CATALOG_CLEANUP_ENABLED,omitempty" tfsdk:"SOFTWARE_CATALOG_CLEANUP_ENABLED"`
	SOFTWARECATALOGEESSIAPIURL                     *string  `json:"SOFTWARE_CATALOG_EESSI_API_URL,omitempty" tfsdk:"SOFTWARE_CATALOG_EESSI_API_URL"`
	SOFTWARECATALOGEESSIINCLUDEEXTENSIONS          *bool    `json:"SOFTWARE_CATALOG_EESSI_INCLUDE_EXTENSIONS,omitempty" tfsdk:"SOFTWARE_CATALOG_EESSI_INCLUDE_EXTENSIONS"`
	SOFTWARECATALOGEESSIUPDATEENABLED              *bool    `json:"SOFTWARE_CATALOG_EESSI_UPDATE_ENABLED,omitempty" tfsdk:"SOFTWARE_CATALOG_EESSI_UPDATE_ENABLED"`
	SOFTWARECATALOGEESSIVERSION                    *string  `json:"SOFTWARE_CATALOG_EESSI_VERSION,omitempty" tfsdk:"SOFTWARE_CATALOG_EESSI_VERSION"`
	SOFTWARECATALOGRETENTIONDAYS                   *int64   `json:"SOFTWARE_CATALOG_RETENTION_DAYS,omitempty" tfsdk:"SOFTWARE_CATALOG_RETENTION_DAYS"`
	SOFTWARECATALOGSPACKDATAURL                    *string  `json:"SOFTWARE_CATALOG_SPACK_DATA_URL,omitempty" tfsdk:"SOFTWARE_CATALOG_SPACK_DATA_URL"`
	SOFTWARECATALOGSPACKUPDATEENABLED              *bool    `json:"SOFTWARE_CATALOG_SPACK_UPDATE_ENABLED,omitempty" tfsdk:"SOFTWARE_CATALOG_SPACK_UPDATE_ENABLED"`
	SOFTWARECATALOGSPACKVERSION                    *string  `json:"SOFTWARE_CATALOG_SPACK_VERSION,omitempty" tfsdk:"SOFTWARE_CATALOG_SPACK_VERSION"`
	SOFTWARECATALOGUPDATEEXISTINGPACKAGES          *bool    `json:"SOFTWARE_CATALOG_UPDATE_EXISTING_PACKAGES,omitempty" tfsdk:"SOFTWARE_CATALOG_UPDATE_EXISTING_PACKAGES"`
	SUPPORTPORTALURL                               *string  `json:"SUPPORT_PORTAL_URL,omitempty" tfsdk:"SUPPORT_PORTAL_URL"`
	TELEMETRYURL                                   *string  `json:"TELEMETRY_URL,omitempty" tfsdk:"TELEMETRY_URL"`
	TELEMETRYVERSION                               *int64   `json:"TELEMETRY_VERSION,omitempty" tfsdk:"TELEMETRY_VERSION"`
	THUMBNAILSIZE                                  *string  `json:"THUMBNAIL_SIZE,omitempty" tfsdk:"THUMBNAIL_SIZE"`
	USERTABLECOLUMNS                               *string  `json:"USER_TABLE_COLUMNS,omitempty" tfsdk:"USER_TABLE_COLUMNS"`
	WALDURAUTHSOCIALROLECLAIM                      *string  `json:"WALDUR_AUTH_SOCIAL_ROLE_CLAIM,omitempty" tfsdk:"WALDUR_AUTH_SOCIAL_ROLE_CLAIM"`
	WALDURSUPPORTACTIVEBACKENDTYPE                 *string  `json:"WALDUR_SUPPORT_ACTIVE_BACKEND_TYPE,omitempty" tfsdk:"WALDUR_SUPPORT_ACTIVE_BACKEND_TYPE"`
	WALDURSUPPORTDISPLAYREQUESTTYPE                *bool    `json:"WALDUR_SUPPORT_DISPLAY_REQUEST_TYPE,omitempty" tfsdk:"WALDUR_SUPPORT_DISPLAY_REQUEST_TYPE"`
	WALDURSUPPORTENABLED                           *bool    `json:"WALDUR_SUPPORT_ENABLED,omitempty" tfsdk:"WALDUR_SUPPORT_ENABLED"`
	ZAMMADAPIURL                                   *string  `json:"ZAMMAD_API_URL,omitempty" tfsdk:"ZAMMAD_API_URL"`
	ZAMMADARTICLETYPE                              *string  `json:"ZAMMAD_ARTICLE_TYPE,omitempty" tfsdk:"ZAMMAD_ARTICLE_TYPE"`
	ZAMMADCOMMENTCOOLDOWNDURATION                  *int64   `json:"ZAMMAD_COMMENT_COOLDOWN_DURATION,omitempty" tfsdk:"ZAMMAD_COMMENT_COOLDOWN_DURATION"`
	ZAMMADCOMMENTMARKER                            *string  `json:"ZAMMAD_COMMENT_MARKER,omitempty" tfsdk:"ZAMMAD_COMMENT_MARKER"`
	ZAMMADCOMMENTPREFIX                            *string  `json:"ZAMMAD_COMMENT_PREFIX,omitempty" tfsdk:"ZAMMAD_COMMENT_PREFIX"`
	ZAMMADGROUP                                    *string  `json:"ZAMMAD_GROUP,omitempty" tfsdk:"ZAMMAD_GROUP"`
	ZAMMADTOKEN                                    *string  `json:"ZAMMAD_TOKEN,omitempty" tfsdk:"ZAMMAD_TOKEN"`
}

type ConstanceSettingsRequest struct {
	ANONYMOUSUSERCANVIEWOFFERINGS                  *bool    `json:"ANONYMOUS_USER_CAN_VIEW_OFFERINGS,omitempty" tfsdk:"ANONYMOUS_USER_CAN_VIEW_OFFERINGS"`
	ANONYMOUSUSERCANVIEWPLANS                      *bool    `json:"ANONYMOUS_USER_CAN_VIEW_PLANS,omitempty" tfsdk:"ANONYMOUS_USER_CAN_VIEW_PLANS"`
	ATLASSIANAFFECTEDRESOURCEFIELD                 *string  `json:"ATLASSIAN_AFFECTED_RESOURCE_FIELD,omitempty" tfsdk:"ATLASSIAN_AFFECTED_RESOURCE_FIELD"`
	ATLASSIANAPIURL                                *string  `json:"ATLASSIAN_API_URL,omitempty" tfsdk:"ATLASSIAN_API_URL"`
	ATLASSIANCALLERFIELD                           *string  `json:"ATLASSIAN_CALLER_FIELD,omitempty" tfsdk:"ATLASSIAN_CALLER_FIELD"`
	ATLASSIANCUSTOMISSUEFIELDMAPPINGENABLED        *bool    `json:"ATLASSIAN_CUSTOM_ISSUE_FIELD_MAPPING_ENABLED,omitempty" tfsdk:"ATLASSIAN_CUSTOM_ISSUE_FIELD_MAPPING_ENABLED"`
	ATLASSIANDEFAULTOFFERINGISSUETYPE              *string  `json:"ATLASSIAN_DEFAULT_OFFERING_ISSUE_TYPE,omitempty" tfsdk:"ATLASSIAN_DEFAULT_OFFERING_ISSUE_TYPE"`
	ATLASSIANDESCRIPTIONTEMPLATE                   *string  `json:"ATLASSIAN_DESCRIPTION_TEMPLATE,omitempty" tfsdk:"ATLASSIAN_DESCRIPTION_TEMPLATE"`
	ATLASSIANEMAIL                                 *string  `json:"ATLASSIAN_EMAIL,omitempty" tfsdk:"ATLASSIAN_EMAIL"`
	ATLASSIANEXCLUDEDATTACHMENTTYPES               *string  `json:"ATLASSIAN_EXCLUDED_ATTACHMENT_TYPES,omitempty" tfsdk:"ATLASSIAN_EXCLUDED_ATTACHMENT_TYPES"`
	ATLASSIANIMPACTFIELD                           *string  `json:"ATLASSIAN_IMPACT_FIELD,omitempty" tfsdk:"ATLASSIAN_IMPACT_FIELD"`
	ATLASSIANLINKEDISSUETYPE                       *string  `json:"ATLASSIAN_LINKED_ISSUE_TYPE,omitempty" tfsdk:"ATLASSIAN_LINKED_ISSUE_TYPE"`
	ATLASSIANMAPWALDURUSERSTOSERVICEDESKAGENTS     *bool    `json:"ATLASSIAN_MAP_WALDUR_USERS_TO_SERVICEDESK_AGENTS,omitempty" tfsdk:"ATLASSIAN_MAP_WALDUR_USERS_TO_SERVICEDESK_AGENTS"`
	ATLASSIANOAUTH2ACCESSTOKEN                     *string  `json:"ATLASSIAN_OAUTH2_ACCESS_TOKEN,omitempty" tfsdk:"ATLASSIAN_OAUTH2_ACCESS_TOKEN"`
	ATLASSIANOAUTH2CLIENTID                        *string  `json:"ATLASSIAN_OAUTH2_CLIENT_ID,omitempty" tfsdk:"ATLASSIAN_OAUTH2_CLIENT_ID"`
	ATLASSIANOAUTH2TOKENTYPE                       *string  `json:"ATLASSIAN_OAUTH2_TOKEN_TYPE,omitempty" tfsdk:"ATLASSIAN_OAUTH2_TOKEN_TYPE"`
	ATLASSIANORGANISATIONFIELD                     *string  `json:"ATLASSIAN_ORGANISATION_FIELD,omitempty" tfsdk:"ATLASSIAN_ORGANISATION_FIELD"`
	ATLASSIANPASSWORD                              *string  `json:"ATLASSIAN_PASSWORD,omitempty" tfsdk:"ATLASSIAN_PASSWORD"`
	ATLASSIANPERSONALACCESSTOKEN                   *string  `json:"ATLASSIAN_PERSONAL_ACCESS_TOKEN,omitempty" tfsdk:"ATLASSIAN_PERSONAL_ACCESS_TOKEN"`
	ATLASSIANPROJECTFIELD                          *string  `json:"ATLASSIAN_PROJECT_FIELD,omitempty" tfsdk:"ATLASSIAN_PROJECT_FIELD"`
	ATLASSIANPROJECTID                             *string  `json:"ATLASSIAN_PROJECT_ID,omitempty" tfsdk:"ATLASSIAN_PROJECT_ID"`
	ATLASSIANREPORTERFIELD                         *string  `json:"ATLASSIAN_REPORTER_FIELD,omitempty" tfsdk:"ATLASSIAN_REPORTER_FIELD"`
	ATLASSIANREQUESTFEEDBACKFIELD                  *string  `json:"ATLASSIAN_REQUEST_FEEDBACK_FIELD,omitempty" tfsdk:"ATLASSIAN_REQUEST_FEEDBACK_FIELD"`
	ATLASSIANRESOLUTIONSLAFIELD                    *string  `json:"ATLASSIAN_RESOLUTION_SLA_FIELD,omitempty" tfsdk:"ATLASSIAN_RESOLUTION_SLA_FIELD"`
	ATLASSIANSATISFACTIONFIELD                     *string  `json:"ATLASSIAN_SATISFACTION_FIELD,omitempty" tfsdk:"ATLASSIAN_SATISFACTION_FIELD"`
	ATLASSIANSHAREDUSERNAME                        *bool    `json:"ATLASSIAN_SHARED_USERNAME,omitempty" tfsdk:"ATLASSIAN_SHARED_USERNAME"`
	ATLASSIANSLAFIELD                              *string  `json:"ATLASSIAN_SLA_FIELD,omitempty" tfsdk:"ATLASSIAN_SLA_FIELD"`
	ATLASSIANSUMMARYTEMPLATE                       *string  `json:"ATLASSIAN_SUMMARY_TEMPLATE,omitempty" tfsdk:"ATLASSIAN_SUMMARY_TEMPLATE"`
	ATLASSIANTEMPLATEFIELD                         *string  `json:"ATLASSIAN_TEMPLATE_FIELD,omitempty" tfsdk:"ATLASSIAN_TEMPLATE_FIELD"`
	ATLASSIANTOKEN                                 *string  `json:"ATLASSIAN_TOKEN,omitempty" tfsdk:"ATLASSIAN_TOKEN"`
	ATLASSIANUSERNAME                              *string  `json:"ATLASSIAN_USERNAME,omitempty" tfsdk:"ATLASSIAN_USERNAME"`
	ATLASSIANUSEOLDAPI                             *bool    `json:"ATLASSIAN_USE_OLD_API,omitempty" tfsdk:"ATLASSIAN_USE_OLD_API"`
	ATLASSIANVERIFYSSL                             *bool    `json:"ATLASSIAN_VERIFY_SSL,omitempty" tfsdk:"ATLASSIAN_VERIFY_SSL"`
	ATLASSIANWALDURBACKENDIDFIELD                  *string  `json:"ATLASSIAN_WALDUR_BACKEND_ID_FIELD,omitempty" tfsdk:"ATLASSIAN_WALDUR_BACKEND_ID_FIELD"`
	AUTOAPPROVEUSERTOS                             *bool    `json:"AUTO_APPROVE_USER_TOS,omitempty" tfsdk:"AUTO_APPROVE_USER_TOS"`
	BRANDCOLOR                                     *string  `json:"BRAND_COLOR,omitempty" tfsdk:"BRAND_COLOR"`
	CALLMANAGEMENTHEROIMAGE                        *string  `json:"CALL_MANAGEMENT_HERO_IMAGE,omitempty" tfsdk:"CALL_MANAGEMENT_HERO_IMAGE"`
	COMMONFOOTERHTML                               *string  `json:"COMMON_FOOTER_HTML,omitempty" tfsdk:"COMMON_FOOTER_HTML"`
	COMMONFOOTERTEXT                               *string  `json:"COMMON_FOOTER_TEXT,omitempty" tfsdk:"COMMON_FOOTER_TEXT"`
	COUNTRIES                                      []string `json:"COUNTRIES,omitempty" tfsdk:"COUNTRIES"`
	CURRENCYNAME                                   *string  `json:"CURRENCY_NAME,omitempty" tfsdk:"CURRENCY_NAME"`
	DEACTIVATEUSERIFNOROLES                        *bool    `json:"DEACTIVATE_USER_IF_NO_ROLES,omitempty" tfsdk:"DEACTIVATE_USER_IF_NO_ROLES"`
	DEFAULTIDP                                     *string  `json:"DEFAULT_IDP,omitempty" tfsdk:"DEFAULT_IDP"`
	DISABLEDOFFERINGTYPES                          []string `json:"DISABLED_OFFERING_TYPES,omitempty" tfsdk:"DISABLED_OFFERING_TYPES"`
	DISABLEDARKTHEME                               *bool    `json:"DISABLE_DARK_THEME,omitempty" tfsdk:"DISABLE_DARK_THEME"`
	DISABLESENDINGNOTIFICATIONSABOUTRESOURCEUPDATE *bool    `json:"DISABLE_SENDING_NOTIFICATIONS_ABOUT_RESOURCE_UPDATE,omitempty" tfsdk:"DISABLE_SENDING_NOTIFICATIONS_ABOUT_RESOURCE_UPDATE"`
	DOCKERCLIENT                                   *string  `json:"DOCKER_CLIENT,omitempty" tfsdk:"DOCKER_CLIENT"`
	DOCKERIMAGES                                   *string  `json:"DOCKER_IMAGES,omitempty" tfsdk:"DOCKER_IMAGES"`
	DOCKERREMOVECONTAINER                          *bool    `json:"DOCKER_REMOVE_CONTAINER,omitempty" tfsdk:"DOCKER_REMOVE_CONTAINER"`
	DOCKERRUNOPTIONS                               *string  `json:"DOCKER_RUN_OPTIONS,omitempty" tfsdk:"DOCKER_RUN_OPTIONS"`
	DOCKERSCRIPTDIR                                *string  `json:"DOCKER_SCRIPT_DIR,omitempty" tfsdk:"DOCKER_SCRIPT_DIR"`
	DOCKERVOLUMENAME                               *string  `json:"DOCKER_VOLUME_NAME,omitempty" tfsdk:"DOCKER_VOLUME_NAME"`
	DOCSURL                                        *string  `json:"DOCS_URL,omitempty" tfsdk:"DOCS_URL"`
	ENABLEMOCKCOURSEACCOUNTBACKEND                 *bool    `json:"ENABLE_MOCK_COURSE_ACCOUNT_BACKEND,omitempty" tfsdk:"ENABLE_MOCK_COURSE_ACCOUNT_BACKEND"`
	ENABLEMOCKSERVICEACCOUNTBACKEND                *bool    `json:"ENABLE_MOCK_SERVICE_ACCOUNT_BACKEND,omitempty" tfsdk:"ENABLE_MOCK_SERVICE_ACCOUNT_BACKEND"`
	ENABLEORDERSTARTDATE                           *bool    `json:"ENABLE_ORDER_START_DATE,omitempty" tfsdk:"ENABLE_ORDER_START_DATE"`
	ENABLESTALERESOURCENOTIFICATIONS               *bool    `json:"ENABLE_STALE_RESOURCE_NOTIFICATIONS,omitempty" tfsdk:"ENABLE_STALE_RESOURCE_NOTIFICATIONS"`
	ENABLESTRICTCHECKACCEPTINGINVITATION           *bool    `json:"ENABLE_STRICT_CHECK_ACCEPTING_INVITATION,omitempty" tfsdk:"ENABLE_STRICT_CHECK_ACCEPTING_INVITATION"`
	ENFORCEUSERCONSENTFOROFFERINGS                 *bool    `json:"ENFORCE_USER_CONSENT_FOR_OFFERINGS,omitempty" tfsdk:"ENFORCE_USER_CONSENT_FOR_OFFERINGS"`
	FAVICON                                        *string  `json:"FAVICON,omitempty" tfsdk:"FAVICON"`
	FREEIPABLACKLISTEDUSERNAMES                    []string `json:"FREEIPA_BLACKLISTED_USERNAMES,omitempty" tfsdk:"FREEIPA_BLACKLISTED_USERNAMES"`
	FREEIPAENABLED                                 *bool    `json:"FREEIPA_ENABLED,omitempty" tfsdk:"FREEIPA_ENABLED"`
	FREEIPAGROUPNAMEPREFIX                         *string  `json:"FREEIPA_GROUPNAME_PREFIX,omitempty" tfsdk:"FREEIPA_GROUPNAME_PREFIX"`
	FREEIPAGROUPSYNCHRONIZATIONENABLED             *bool    `json:"FREEIPA_GROUP_SYNCHRONIZATION_ENABLED,omitempty" tfsdk:"FREEIPA_GROUP_SYNCHRONIZATION_ENABLED"`
	FREEIPAHOSTNAME                                *string  `json:"FREEIPA_HOSTNAME,omitempty" tfsdk:"FREEIPA_HOSTNAME"`
	FREEIPAPASSWORD                                *string  `json:"FREEIPA_PASSWORD,omitempty" tfsdk:"FREEIPA_PASSWORD"`
	FREEIPAUSERNAME                                *string  `json:"FREEIPA_USERNAME,omitempty" tfsdk:"FREEIPA_USERNAME"`
	FREEIPAUSERNAMEPREFIX                          *string  `json:"FREEIPA_USERNAME_PREFIX,omitempty" tfsdk:"FREEIPA_USERNAME_PREFIX"`
	FREEIPAVERIFYSSL                               *bool    `json:"FREEIPA_VERIFY_SSL,omitempty" tfsdk:"FREEIPA_VERIFY_SSL"`
	FULLPAGETITLE                                  *string  `json:"FULL_PAGE_TITLE,omitempty" tfsdk:"FULL_PAGE_TITLE"`
	HEROIMAGE                                      *string  `json:"HERO_IMAGE,omitempty" tfsdk:"HERO_IMAGE"`
	HEROLINKLABEL                                  *string  `json:"HERO_LINK_LABEL,omitempty" tfsdk:"HERO_LINK_LABEL"`
	HEROLINKURL                                    *string  `json:"HERO_LINK_URL,omitempty" tfsdk:"HERO_LINK_URL"`
	HOMEPORTURL                                    *string  `json:"HOMEPORT_URL,omitempty" tfsdk:"HOMEPORT_URL"`
	INVITATIONDISABLEMULTIPLEROLES                 *bool    `json:"INVITATION_DISABLE_MULTIPLE_ROLES,omitempty" tfsdk:"INVITATION_DISABLE_MULTIPLE_ROLES"`
	K8SCONFIGPATH                                  *string  `json:"K8S_CONFIG_PATH,omitempty" tfsdk:"K8S_CONFIG_PATH"`
	K8SJOBTIMEOUT                                  *int64   `json:"K8S_JOB_TIMEOUT,omitempty" tfsdk:"K8S_JOB_TIMEOUT"`
	K8SNAMESPACE                                   *string  `json:"K8S_NAMESPACE,omitempty" tfsdk:"K8S_NAMESPACE"`
	KEYCLOAKICON                                   *string  `json:"KEYCLOAK_ICON,omitempty" tfsdk:"KEYCLOAK_ICON"`
	LANGUAGECHOICES                                *string  `json:"LANGUAGE_CHOICES,omitempty" tfsdk:"LANGUAGE_CHOICES"`
	LLMCHATENABLED                                 *bool    `json:"LLM_CHAT_ENABLED,omitempty" tfsdk:"LLM_CHAT_ENABLED"`
	LLMINFERENCESAPITOKEN                          *string  `json:"LLM_INFERENCES_API_TOKEN,omitempty" tfsdk:"LLM_INFERENCES_API_TOKEN"`
	LLMINFERENCESAPIURL                            *string  `json:"LLM_INFERENCES_API_URL,omitempty" tfsdk:"LLM_INFERENCES_API_URL"`
	LLMINFERENCESBACKENDTYPE                       *string  `json:"LLM_INFERENCES_BACKEND_TYPE,omitempty" tfsdk:"LLM_INFERENCES_BACKEND_TYPE"`
	LLMINFERENCESMODEL                             *string  `json:"LLM_INFERENCES_MODEL,omitempty" tfsdk:"LLM_INFERENCES_MODEL"`
	LOGINLOGO                                      *string  `json:"LOGIN_LOGO,omitempty" tfsdk:"LOGIN_LOGO"`
	LOGINPAGELAYOUT                                *string  `json:"LOGIN_PAGE_LAYOUT,omitempty" tfsdk:"LOGIN_PAGE_LAYOUT"`
	LOGINPAGEVIDEOURL                              *string  `json:"LOGIN_PAGE_VIDEO_URL,omitempty" tfsdk:"LOGIN_PAGE_VIDEO_URL"`
	MAINTENANCEANNOUNCEMENTNOTIFYBEFOREMINUTES     *int64   `json:"MAINTENANCE_ANNOUNCEMENT_NOTIFY_BEFORE_MINUTES,omitempty" tfsdk:"MAINTENANCE_ANNOUNCEMENT_NOTIFY_BEFORE_MINUTES"`
	MAINTENANCEANNOUNCEMENTNOTIFYSYSTEM            []string `json:"MAINTENANCE_ANNOUNCEMENT_NOTIFY_SYSTEM,omitempty" tfsdk:"MAINTENANCE_ANNOUNCEMENT_NOTIFY_SYSTEM"`
	MARKETPLACEHEROIMAGE                           *string  `json:"MARKETPLACE_HERO_IMAGE,omitempty" tfsdk:"MARKETPLACE_HERO_IMAGE"`
	MARKETPLACELANDINGPAGE                         *string  `json:"MARKETPLACE_LANDING_PAGE,omitempty" tfsdk:"MARKETPLACE_LANDING_PAGE"`
	NOTIFYABOUTRESOURCECHANGE                      *bool    `json:"NOTIFY_ABOUT_RESOURCE_CHANGE,omitempty" tfsdk:"NOTIFY_ABOUT_RESOURCE_CHANGE"`
	NOTIFYSTAFFABOUTAPPROVALS                      *bool    `json:"NOTIFY_STAFF_ABOUT_APPROVALS,omitempty" tfsdk:"NOTIFY_STAFF_ABOUT_APPROVALS"`
	OFFERINGLOGOPLACEHOLDER                        *string  `json:"OFFERING_LOGO_PLACEHOLDER,omitempty" tfsdk:"OFFERING_LOGO_PLACEHOLDER"`
	OIDCACCESSTOKENENABLED                         *bool    `json:"OIDC_ACCESS_TOKEN_ENABLED,omitempty" tfsdk:"OIDC_ACCESS_TOKEN_ENABLED"`
	OIDCAUTHURL                                    *string  `json:"OIDC_AUTH_URL,omitempty" tfsdk:"OIDC_AUTH_URL"`
	OIDCBLOCKCREATIONOFUNINVITEDUSERS              *bool    `json:"OIDC_BLOCK_CREATION_OF_UNINVITED_USERS,omitempty" tfsdk:"OIDC_BLOCK_CREATION_OF_UNINVITED_USERS"`
	OIDCCACHETIMEOUT                               *int64   `json:"OIDC_CACHE_TIMEOUT,omitempty" tfsdk:"OIDC_CACHE_TIMEOUT"`
	OIDCCLIENTID                                   *string  `json:"OIDC_CLIENT_ID,omitempty" tfsdk:"OIDC_CLIENT_ID"`
	OIDCCLIENTSECRET                               *string  `json:"OIDC_CLIENT_SECRET,omitempty" tfsdk:"OIDC_CLIENT_SECRET"`
	OIDCINTROSPECTIONURL                           *string  `json:"OIDC_INTROSPECTION_URL,omitempty" tfsdk:"OIDC_INTROSPECTION_URL"`
	OIDCUSERFIELD                                  *string  `json:"OIDC_USER_FIELD,omitempty" tfsdk:"OIDC_USER_FIELD"`
	ONBOARDINGARIREGISTERBASEURL                   *string  `json:"ONBOARDING_ARIREGISTER_BASE_URL,omitempty" tfsdk:"ONBOARDING_ARIREGISTER_BASE_URL"`
	ONBOARDINGARIREGISTERPASSWORD                  *string  `json:"ONBOARDING_ARIREGISTER_PASSWORD,omitempty" tfsdk:"ONBOARDING_ARIREGISTER_PASSWORD"`
	ONBOARDINGARIREGISTERTIMEOUT                   *int64   `json:"ONBOARDING_ARIREGISTER_TIMEOUT,omitempty" tfsdk:"ONBOARDING_ARIREGISTER_TIMEOUT"`
	ONBOARDINGARIREGISTERUSERNAME                  *string  `json:"ONBOARDING_ARIREGISTER_USERNAME,omitempty" tfsdk:"ONBOARDING_ARIREGISTER_USERNAME"`
	ONBOARDINGBOLAGSVERKETAPIURL                   *string  `json:"ONBOARDING_BOLAGSVERKET_API_URL,omitempty" tfsdk:"ONBOARDING_BOLAGSVERKET_API_URL"`
	ONBOARDINGBOLAGSVERKETCLIENTID                 *string  `json:"ONBOARDING_BOLAGSVERKET_CLIENT_ID,omitempty" tfsdk:"ONBOARDING_BOLAGSVERKET_CLIENT_ID"`
	ONBOARDINGBOLAGSVERKETCLIENTSECRET             *string  `json:"ONBOARDING_BOLAGSVERKET_CLIENT_SECRET,omitempty" tfsdk:"ONBOARDING_BOLAGSVERKET_CLIENT_SECRET"`
	ONBOARDINGBOLAGSVERKETTOKENAPIURL              *string  `json:"ONBOARDING_BOLAGSVERKET_TOKEN_API_URL,omitempty" tfsdk:"ONBOARDING_BOLAGSVERKET_TOKEN_API_URL"`
	ONBOARDINGBREGAPIURL                           *string  `json:"ONBOARDING_BREG_API_URL,omitempty" tfsdk:"ONBOARDING_BREG_API_URL"`
	ONBOARDINGCOUNTRY                              *string  `json:"ONBOARDING_COUNTRY,omitempty" tfsdk:"ONBOARDING_COUNTRY"`
	ONBOARDINGVERIFICATIONEXPIRYHOURS              *int64   `json:"ONBOARDING_VERIFICATION_EXPIRY_HOURS,omitempty" tfsdk:"ONBOARDING_VERIFICATION_EXPIRY_HOURS"`
	ONBOARDINGWICOAPIURL                           *string  `json:"ONBOARDING_WICO_API_URL,omitempty" tfsdk:"ONBOARDING_WICO_API_URL"`
	ONBOARDINGWICOTOKEN                            *string  `json:"ONBOARDING_WICO_TOKEN,omitempty" tfsdk:"ONBOARDING_WICO_TOKEN"`
	POWEREDBYLOGO                                  *string  `json:"POWERED_BY_LOGO,omitempty" tfsdk:"POWERED_BY_LOGO"`
	PROJECTENDDATEMANDATORY                        *bool    `json:"PROJECT_END_DATE_MANDATORY,omitempty" tfsdk:"PROJECT_END_DATE_MANDATORY"`
	PROPOSALREVIEWDURATION                         *int64   `json:"PROPOSAL_REVIEW_DURATION,omitempty" tfsdk:"PROPOSAL_REVIEW_DURATION"`
	RANCHERUSERNAMEINPUTLABEL                      *string  `json:"RANCHER_USERNAME_INPUT_LABEL,omitempty" tfsdk:"RANCHER_USERNAME_INPUT_LABEL"`
	SCRIPTRUNMODE                                  *string  `json:"SCRIPT_RUN_MODE,omitempty" tfsdk:"SCRIPT_RUN_MODE"`
	SHORTPAGETITLE                                 *string  `json:"SHORT_PAGE_TITLE,omitempty" tfsdk:"SHORT_PAGE_TITLE"`
	SIDEBARLOGO                                    *string  `json:"SIDEBAR_LOGO,omitempty" tfsdk:"SIDEBAR_LOGO"`
	SIDEBARLOGODARK                                *string  `json:"SIDEBAR_LOGO_DARK,omitempty" tfsdk:"SIDEBAR_LOGO_DARK"`
	SIDEBARLOGOMOBILE                              *string  `json:"SIDEBAR_LOGO_MOBILE,omitempty" tfsdk:"SIDEBAR_LOGO_MOBILE"`
	SIDEBARSTYLE                                   *string  `json:"SIDEBAR_STYLE,omitempty" tfsdk:"SIDEBAR_STYLE"`
	SITEADDRESS                                    *string  `json:"SITE_ADDRESS,omitempty" tfsdk:"SITE_ADDRESS"`
	SITEDESCRIPTION                                *string  `json:"SITE_DESCRIPTION,omitempty" tfsdk:"SITE_DESCRIPTION"`
	SITEEMAIL                                      *string  `json:"SITE_EMAIL,omitempty" tfsdk:"SITE_EMAIL"`
	SITELOGO                                       *string  `json:"SITE_LOGO,omitempty" tfsdk:"SITE_LOGO"`
	SITENAME                                       *string  `json:"SITE_NAME,omitempty" tfsdk:"SITE_NAME"`
	SITEPHONE                                      *string  `json:"SITE_PHONE,omitempty" tfsdk:"SITE_PHONE"`
	SMAXAFFECTEDRESOURCEFIELD                      *string  `json:"SMAX_AFFECTED_RESOURCE_FIELD,omitempty" tfsdk:"SMAX_AFFECTED_RESOURCE_FIELD"`
	SMAXAPIURL                                     *string  `json:"SMAX_API_URL,omitempty" tfsdk:"SMAX_API_URL"`
	SMAXCREATIONSOURCENAME                         *string  `json:"SMAX_CREATION_SOURCE_NAME,omitempty" tfsdk:"SMAX_CREATION_SOURCE_NAME"`
	SMAXLOGIN                                      *string  `json:"SMAX_LOGIN,omitempty" tfsdk:"SMAX_LOGIN"`
	SMAXORGANISATIONFIELD                          *string  `json:"SMAX_ORGANISATION_FIELD,omitempty" tfsdk:"SMAX_ORGANISATION_FIELD"`
	SMAXPASSWORD                                   *string  `json:"SMAX_PASSWORD,omitempty" tfsdk:"SMAX_PASSWORD"`
	SMAXPROJECTFIELD                               *string  `json:"SMAX_PROJECT_FIELD,omitempty" tfsdk:"SMAX_PROJECT_FIELD"`
	SMAXREQUESTSOFFERING                           *string  `json:"SMAX_REQUESTS_OFFERING,omitempty" tfsdk:"SMAX_REQUESTS_OFFERING"`
	SMAXSECONDSTOWAIT                              *int64   `json:"SMAX_SECONDS_TO_WAIT,omitempty" tfsdk:"SMAX_SECONDS_TO_WAIT"`
	SMAXTENANTID                                   *string  `json:"SMAX_TENANT_ID,omitempty" tfsdk:"SMAX_TENANT_ID"`
	SMAXTIMESTOPULL                                *int64   `json:"SMAX_TIMES_TO_PULL,omitempty" tfsdk:"SMAX_TIMES_TO_PULL"`
	SMAXVERIFYSSL                                  *bool    `json:"SMAX_VERIFY_SSL,omitempty" tfsdk:"SMAX_VERIFY_SSL"`
	SOFTWARECATALOGCLEANUPENABLED                  *bool    `json:"SOFTWARE_CATALOG_CLEANUP_ENABLED,omitempty" tfsdk:"SOFTWARE_CATALOG_CLEANUP_ENABLED"`
	SOFTWARECATALOGEESSIAPIURL                     *string  `json:"SOFTWARE_CATALOG_EESSI_API_URL,omitempty" tfsdk:"SOFTWARE_CATALOG_EESSI_API_URL"`
	SOFTWARECATALOGEESSIINCLUDEEXTENSIONS          *bool    `json:"SOFTWARE_CATALOG_EESSI_INCLUDE_EXTENSIONS,omitempty" tfsdk:"SOFTWARE_CATALOG_EESSI_INCLUDE_EXTENSIONS"`
	SOFTWARECATALOGEESSIUPDATEENABLED              *bool    `json:"SOFTWARE_CATALOG_EESSI_UPDATE_ENABLED,omitempty" tfsdk:"SOFTWARE_CATALOG_EESSI_UPDATE_ENABLED"`
	SOFTWARECATALOGEESSIVERSION                    *string  `json:"SOFTWARE_CATALOG_EESSI_VERSION,omitempty" tfsdk:"SOFTWARE_CATALOG_EESSI_VERSION"`
	SOFTWARECATALOGRETENTIONDAYS                   *int64   `json:"SOFTWARE_CATALOG_RETENTION_DAYS,omitempty" tfsdk:"SOFTWARE_CATALOG_RETENTION_DAYS"`
	SOFTWARECATALOGSPACKDATAURL                    *string  `json:"SOFTWARE_CATALOG_SPACK_DATA_URL,omitempty" tfsdk:"SOFTWARE_CATALOG_SPACK_DATA_URL"`
	SOFTWARECATALOGSPACKUPDATEENABLED              *bool    `json:"SOFTWARE_CATALOG_SPACK_UPDATE_ENABLED,omitempty" tfsdk:"SOFTWARE_CATALOG_SPACK_UPDATE_ENABLED"`
	SOFTWARECATALOGSPACKVERSION                    *string  `json:"SOFTWARE_CATALOG_SPACK_VERSION,omitempty" tfsdk:"SOFTWARE_CATALOG_SPACK_VERSION"`
	SOFTWARECATALOGUPDATEEXISTINGPACKAGES          *bool    `json:"SOFTWARE_CATALOG_UPDATE_EXISTING_PACKAGES,omitempty" tfsdk:"SOFTWARE_CATALOG_UPDATE_EXISTING_PACKAGES"`
	SUPPORTPORTALURL                               *string  `json:"SUPPORT_PORTAL_URL,omitempty" tfsdk:"SUPPORT_PORTAL_URL"`
	TELEMETRYURL                                   *string  `json:"TELEMETRY_URL,omitempty" tfsdk:"TELEMETRY_URL"`
	TELEMETRYVERSION                               *int64   `json:"TELEMETRY_VERSION,omitempty" tfsdk:"TELEMETRY_VERSION"`
	THUMBNAILSIZE                                  *string  `json:"THUMBNAIL_SIZE,omitempty" tfsdk:"THUMBNAIL_SIZE"`
	USERTABLECOLUMNS                               *string  `json:"USER_TABLE_COLUMNS,omitempty" tfsdk:"USER_TABLE_COLUMNS"`
	WALDURAUTHSOCIALROLECLAIM                      *string  `json:"WALDUR_AUTH_SOCIAL_ROLE_CLAIM,omitempty" tfsdk:"WALDUR_AUTH_SOCIAL_ROLE_CLAIM"`
	WALDURSUPPORTACTIVEBACKENDTYPE                 *string  `json:"WALDUR_SUPPORT_ACTIVE_BACKEND_TYPE,omitempty" tfsdk:"WALDUR_SUPPORT_ACTIVE_BACKEND_TYPE"`
	WALDURSUPPORTDISPLAYREQUESTTYPE                *bool    `json:"WALDUR_SUPPORT_DISPLAY_REQUEST_TYPE,omitempty" tfsdk:"WALDUR_SUPPORT_DISPLAY_REQUEST_TYPE"`
	WALDURSUPPORTENABLED                           *bool    `json:"WALDUR_SUPPORT_ENABLED,omitempty" tfsdk:"WALDUR_SUPPORT_ENABLED"`
	ZAMMADAPIURL                                   *string  `json:"ZAMMAD_API_URL,omitempty" tfsdk:"ZAMMAD_API_URL"`
	ZAMMADARTICLETYPE                              *string  `json:"ZAMMAD_ARTICLE_TYPE,omitempty" tfsdk:"ZAMMAD_ARTICLE_TYPE"`
	ZAMMADCOMMENTCOOLDOWNDURATION                  *int64   `json:"ZAMMAD_COMMENT_COOLDOWN_DURATION,omitempty" tfsdk:"ZAMMAD_COMMENT_COOLDOWN_DURATION"`
	ZAMMADCOMMENTMARKER                            *string  `json:"ZAMMAD_COMMENT_MARKER,omitempty" tfsdk:"ZAMMAD_COMMENT_MARKER"`
	ZAMMADCOMMENTPREFIX                            *string  `json:"ZAMMAD_COMMENT_PREFIX,omitempty" tfsdk:"ZAMMAD_COMMENT_PREFIX"`
	ZAMMADGROUP                                    *string  `json:"ZAMMAD_GROUP,omitempty" tfsdk:"ZAMMAD_GROUP"`
	ZAMMADTOKEN                                    *string  `json:"ZAMMAD_TOKEN,omitempty" tfsdk:"ZAMMAD_TOKEN"`
}

type ConstanceSettingsRequestForm struct {
	ANONYMOUSUSERCANVIEWOFFERINGS                  *bool    `json:"ANONYMOUS_USER_CAN_VIEW_OFFERINGS,omitempty" tfsdk:"ANONYMOUS_USER_CAN_VIEW_OFFERINGS"`
	ANONYMOUSUSERCANVIEWPLANS                      *bool    `json:"ANONYMOUS_USER_CAN_VIEW_PLANS,omitempty" tfsdk:"ANONYMOUS_USER_CAN_VIEW_PLANS"`
	ATLASSIANAFFECTEDRESOURCEFIELD                 *string  `json:"ATLASSIAN_AFFECTED_RESOURCE_FIELD,omitempty" tfsdk:"ATLASSIAN_AFFECTED_RESOURCE_FIELD"`
	ATLASSIANAPIURL                                *string  `json:"ATLASSIAN_API_URL,omitempty" tfsdk:"ATLASSIAN_API_URL"`
	ATLASSIANCALLERFIELD                           *string  `json:"ATLASSIAN_CALLER_FIELD,omitempty" tfsdk:"ATLASSIAN_CALLER_FIELD"`
	ATLASSIANCUSTOMISSUEFIELDMAPPINGENABLED        *bool    `json:"ATLASSIAN_CUSTOM_ISSUE_FIELD_MAPPING_ENABLED,omitempty" tfsdk:"ATLASSIAN_CUSTOM_ISSUE_FIELD_MAPPING_ENABLED"`
	ATLASSIANDEFAULTOFFERINGISSUETYPE              *string  `json:"ATLASSIAN_DEFAULT_OFFERING_ISSUE_TYPE,omitempty" tfsdk:"ATLASSIAN_DEFAULT_OFFERING_ISSUE_TYPE"`
	ATLASSIANDESCRIPTIONTEMPLATE                   *string  `json:"ATLASSIAN_DESCRIPTION_TEMPLATE,omitempty" tfsdk:"ATLASSIAN_DESCRIPTION_TEMPLATE"`
	ATLASSIANEMAIL                                 *string  `json:"ATLASSIAN_EMAIL,omitempty" tfsdk:"ATLASSIAN_EMAIL"`
	ATLASSIANEXCLUDEDATTACHMENTTYPES               *string  `json:"ATLASSIAN_EXCLUDED_ATTACHMENT_TYPES,omitempty" tfsdk:"ATLASSIAN_EXCLUDED_ATTACHMENT_TYPES"`
	ATLASSIANIMPACTFIELD                           *string  `json:"ATLASSIAN_IMPACT_FIELD,omitempty" tfsdk:"ATLASSIAN_IMPACT_FIELD"`
	ATLASSIANLINKEDISSUETYPE                       *string  `json:"ATLASSIAN_LINKED_ISSUE_TYPE,omitempty" tfsdk:"ATLASSIAN_LINKED_ISSUE_TYPE"`
	ATLASSIANMAPWALDURUSERSTOSERVICEDESKAGENTS     *bool    `json:"ATLASSIAN_MAP_WALDUR_USERS_TO_SERVICEDESK_AGENTS,omitempty" tfsdk:"ATLASSIAN_MAP_WALDUR_USERS_TO_SERVICEDESK_AGENTS"`
	ATLASSIANOAUTH2ACCESSTOKEN                     *string  `json:"ATLASSIAN_OAUTH2_ACCESS_TOKEN,omitempty" tfsdk:"ATLASSIAN_OAUTH2_ACCESS_TOKEN"`
	ATLASSIANOAUTH2CLIENTID                        *string  `json:"ATLASSIAN_OAUTH2_CLIENT_ID,omitempty" tfsdk:"ATLASSIAN_OAUTH2_CLIENT_ID"`
	ATLASSIANOAUTH2TOKENTYPE                       *string  `json:"ATLASSIAN_OAUTH2_TOKEN_TYPE,omitempty" tfsdk:"ATLASSIAN_OAUTH2_TOKEN_TYPE"`
	ATLASSIANORGANISATIONFIELD                     *string  `json:"ATLASSIAN_ORGANISATION_FIELD,omitempty" tfsdk:"ATLASSIAN_ORGANISATION_FIELD"`
	ATLASSIANPASSWORD                              *string  `json:"ATLASSIAN_PASSWORD,omitempty" tfsdk:"ATLASSIAN_PASSWORD"`
	ATLASSIANPERSONALACCESSTOKEN                   *string  `json:"ATLASSIAN_PERSONAL_ACCESS_TOKEN,omitempty" tfsdk:"ATLASSIAN_PERSONAL_ACCESS_TOKEN"`
	ATLASSIANPROJECTFIELD                          *string  `json:"ATLASSIAN_PROJECT_FIELD,omitempty" tfsdk:"ATLASSIAN_PROJECT_FIELD"`
	ATLASSIANPROJECTID                             *string  `json:"ATLASSIAN_PROJECT_ID,omitempty" tfsdk:"ATLASSIAN_PROJECT_ID"`
	ATLASSIANREPORTERFIELD                         *string  `json:"ATLASSIAN_REPORTER_FIELD,omitempty" tfsdk:"ATLASSIAN_REPORTER_FIELD"`
	ATLASSIANREQUESTFEEDBACKFIELD                  *string  `json:"ATLASSIAN_REQUEST_FEEDBACK_FIELD,omitempty" tfsdk:"ATLASSIAN_REQUEST_FEEDBACK_FIELD"`
	ATLASSIANRESOLUTIONSLAFIELD                    *string  `json:"ATLASSIAN_RESOLUTION_SLA_FIELD,omitempty" tfsdk:"ATLASSIAN_RESOLUTION_SLA_FIELD"`
	ATLASSIANSATISFACTIONFIELD                     *string  `json:"ATLASSIAN_SATISFACTION_FIELD,omitempty" tfsdk:"ATLASSIAN_SATISFACTION_FIELD"`
	ATLASSIANSHAREDUSERNAME                        *bool    `json:"ATLASSIAN_SHARED_USERNAME,omitempty" tfsdk:"ATLASSIAN_SHARED_USERNAME"`
	ATLASSIANSLAFIELD                              *string  `json:"ATLASSIAN_SLA_FIELD,omitempty" tfsdk:"ATLASSIAN_SLA_FIELD"`
	ATLASSIANSUMMARYTEMPLATE                       *string  `json:"ATLASSIAN_SUMMARY_TEMPLATE,omitempty" tfsdk:"ATLASSIAN_SUMMARY_TEMPLATE"`
	ATLASSIANTEMPLATEFIELD                         *string  `json:"ATLASSIAN_TEMPLATE_FIELD,omitempty" tfsdk:"ATLASSIAN_TEMPLATE_FIELD"`
	ATLASSIANTOKEN                                 *string  `json:"ATLASSIAN_TOKEN,omitempty" tfsdk:"ATLASSIAN_TOKEN"`
	ATLASSIANUSERNAME                              *string  `json:"ATLASSIAN_USERNAME,omitempty" tfsdk:"ATLASSIAN_USERNAME"`
	ATLASSIANUSEOLDAPI                             *bool    `json:"ATLASSIAN_USE_OLD_API,omitempty" tfsdk:"ATLASSIAN_USE_OLD_API"`
	ATLASSIANVERIFYSSL                             *bool    `json:"ATLASSIAN_VERIFY_SSL,omitempty" tfsdk:"ATLASSIAN_VERIFY_SSL"`
	ATLASSIANWALDURBACKENDIDFIELD                  *string  `json:"ATLASSIAN_WALDUR_BACKEND_ID_FIELD,omitempty" tfsdk:"ATLASSIAN_WALDUR_BACKEND_ID_FIELD"`
	AUTOAPPROVEUSERTOS                             *bool    `json:"AUTO_APPROVE_USER_TOS,omitempty" tfsdk:"AUTO_APPROVE_USER_TOS"`
	BRANDCOLOR                                     *string  `json:"BRAND_COLOR,omitempty" tfsdk:"BRAND_COLOR"`
	CALLMANAGEMENTHEROIMAGE                        *string  `json:"CALL_MANAGEMENT_HERO_IMAGE,omitempty" tfsdk:"CALL_MANAGEMENT_HERO_IMAGE"`
	COMMONFOOTERHTML                               *string  `json:"COMMON_FOOTER_HTML,omitempty" tfsdk:"COMMON_FOOTER_HTML"`
	COMMONFOOTERTEXT                               *string  `json:"COMMON_FOOTER_TEXT,omitempty" tfsdk:"COMMON_FOOTER_TEXT"`
	COUNTRIES                                      []string `json:"COUNTRIES,omitempty" tfsdk:"COUNTRIES"`
	CURRENCYNAME                                   *string  `json:"CURRENCY_NAME,omitempty" tfsdk:"CURRENCY_NAME"`
	DEACTIVATEUSERIFNOROLES                        *bool    `json:"DEACTIVATE_USER_IF_NO_ROLES,omitempty" tfsdk:"DEACTIVATE_USER_IF_NO_ROLES"`
	DEFAULTIDP                                     *string  `json:"DEFAULT_IDP,omitempty" tfsdk:"DEFAULT_IDP"`
	DISABLEDOFFERINGTYPES                          []string `json:"DISABLED_OFFERING_TYPES,omitempty" tfsdk:"DISABLED_OFFERING_TYPES"`
	DISABLEDARKTHEME                               *bool    `json:"DISABLE_DARK_THEME,omitempty" tfsdk:"DISABLE_DARK_THEME"`
	DISABLESENDINGNOTIFICATIONSABOUTRESOURCEUPDATE *bool    `json:"DISABLE_SENDING_NOTIFICATIONS_ABOUT_RESOURCE_UPDATE,omitempty" tfsdk:"DISABLE_SENDING_NOTIFICATIONS_ABOUT_RESOURCE_UPDATE"`
	DOCKERCLIENT                                   *string  `json:"DOCKER_CLIENT,omitempty" tfsdk:"DOCKER_CLIENT"`
	DOCKERIMAGES                                   *string  `json:"DOCKER_IMAGES,omitempty" tfsdk:"DOCKER_IMAGES"`
	DOCKERREMOVECONTAINER                          *bool    `json:"DOCKER_REMOVE_CONTAINER,omitempty" tfsdk:"DOCKER_REMOVE_CONTAINER"`
	DOCKERRUNOPTIONS                               *string  `json:"DOCKER_RUN_OPTIONS,omitempty" tfsdk:"DOCKER_RUN_OPTIONS"`
	DOCKERSCRIPTDIR                                *string  `json:"DOCKER_SCRIPT_DIR,omitempty" tfsdk:"DOCKER_SCRIPT_DIR"`
	DOCKERVOLUMENAME                               *string  `json:"DOCKER_VOLUME_NAME,omitempty" tfsdk:"DOCKER_VOLUME_NAME"`
	DOCSURL                                        *string  `json:"DOCS_URL,omitempty" tfsdk:"DOCS_URL"`
	ENABLEMOCKCOURSEACCOUNTBACKEND                 *bool    `json:"ENABLE_MOCK_COURSE_ACCOUNT_BACKEND,omitempty" tfsdk:"ENABLE_MOCK_COURSE_ACCOUNT_BACKEND"`
	ENABLEMOCKSERVICEACCOUNTBACKEND                *bool    `json:"ENABLE_MOCK_SERVICE_ACCOUNT_BACKEND,omitempty" tfsdk:"ENABLE_MOCK_SERVICE_ACCOUNT_BACKEND"`
	ENABLEORDERSTARTDATE                           *bool    `json:"ENABLE_ORDER_START_DATE,omitempty" tfsdk:"ENABLE_ORDER_START_DATE"`
	ENABLESTALERESOURCENOTIFICATIONS               *bool    `json:"ENABLE_STALE_RESOURCE_NOTIFICATIONS,omitempty" tfsdk:"ENABLE_STALE_RESOURCE_NOTIFICATIONS"`
	ENABLESTRICTCHECKACCEPTINGINVITATION           *bool    `json:"ENABLE_STRICT_CHECK_ACCEPTING_INVITATION,omitempty" tfsdk:"ENABLE_STRICT_CHECK_ACCEPTING_INVITATION"`
	ENFORCEUSERCONSENTFOROFFERINGS                 *bool    `json:"ENFORCE_USER_CONSENT_FOR_OFFERINGS,omitempty" tfsdk:"ENFORCE_USER_CONSENT_FOR_OFFERINGS"`
	FAVICON                                        *string  `json:"FAVICON,omitempty" tfsdk:"FAVICON"`
	FREEIPABLACKLISTEDUSERNAMES                    []string `json:"FREEIPA_BLACKLISTED_USERNAMES,omitempty" tfsdk:"FREEIPA_BLACKLISTED_USERNAMES"`
	FREEIPAENABLED                                 *bool    `json:"FREEIPA_ENABLED,omitempty" tfsdk:"FREEIPA_ENABLED"`
	FREEIPAGROUPNAMEPREFIX                         *string  `json:"FREEIPA_GROUPNAME_PREFIX,omitempty" tfsdk:"FREEIPA_GROUPNAME_PREFIX"`
	FREEIPAGROUPSYNCHRONIZATIONENABLED             *bool    `json:"FREEIPA_GROUP_SYNCHRONIZATION_ENABLED,omitempty" tfsdk:"FREEIPA_GROUP_SYNCHRONIZATION_ENABLED"`
	FREEIPAHOSTNAME                                *string  `json:"FREEIPA_HOSTNAME,omitempty" tfsdk:"FREEIPA_HOSTNAME"`
	FREEIPAPASSWORD                                *string  `json:"FREEIPA_PASSWORD,omitempty" tfsdk:"FREEIPA_PASSWORD"`
	FREEIPAUSERNAME                                *string  `json:"FREEIPA_USERNAME,omitempty" tfsdk:"FREEIPA_USERNAME"`
	FREEIPAUSERNAMEPREFIX                          *string  `json:"FREEIPA_USERNAME_PREFIX,omitempty" tfsdk:"FREEIPA_USERNAME_PREFIX"`
	FREEIPAVERIFYSSL                               *bool    `json:"FREEIPA_VERIFY_SSL,omitempty" tfsdk:"FREEIPA_VERIFY_SSL"`
	FULLPAGETITLE                                  *string  `json:"FULL_PAGE_TITLE,omitempty" tfsdk:"FULL_PAGE_TITLE"`
	HEROIMAGE                                      *string  `json:"HERO_IMAGE,omitempty" tfsdk:"HERO_IMAGE"`
	HEROLINKLABEL                                  *string  `json:"HERO_LINK_LABEL,omitempty" tfsdk:"HERO_LINK_LABEL"`
	HEROLINKURL                                    *string  `json:"HERO_LINK_URL,omitempty" tfsdk:"HERO_LINK_URL"`
	HOMEPORTURL                                    *string  `json:"HOMEPORT_URL,omitempty" tfsdk:"HOMEPORT_URL"`
	INVITATIONDISABLEMULTIPLEROLES                 *bool    `json:"INVITATION_DISABLE_MULTIPLE_ROLES,omitempty" tfsdk:"INVITATION_DISABLE_MULTIPLE_ROLES"`
	K8SCONFIGPATH                                  *string  `json:"K8S_CONFIG_PATH,omitempty" tfsdk:"K8S_CONFIG_PATH"`
	K8SJOBTIMEOUT                                  *int64   `json:"K8S_JOB_TIMEOUT,omitempty" tfsdk:"K8S_JOB_TIMEOUT"`
	K8SNAMESPACE                                   *string  `json:"K8S_NAMESPACE,omitempty" tfsdk:"K8S_NAMESPACE"`
	KEYCLOAKICON                                   *string  `json:"KEYCLOAK_ICON,omitempty" tfsdk:"KEYCLOAK_ICON"`
	LANGUAGECHOICES                                *string  `json:"LANGUAGE_CHOICES,omitempty" tfsdk:"LANGUAGE_CHOICES"`
	LLMCHATENABLED                                 *bool    `json:"LLM_CHAT_ENABLED,omitempty" tfsdk:"LLM_CHAT_ENABLED"`
	LLMINFERENCESAPITOKEN                          *string  `json:"LLM_INFERENCES_API_TOKEN,omitempty" tfsdk:"LLM_INFERENCES_API_TOKEN"`
	LLMINFERENCESAPIURL                            *string  `json:"LLM_INFERENCES_API_URL,omitempty" tfsdk:"LLM_INFERENCES_API_URL"`
	LLMINFERENCESBACKENDTYPE                       *string  `json:"LLM_INFERENCES_BACKEND_TYPE,omitempty" tfsdk:"LLM_INFERENCES_BACKEND_TYPE"`
	LLMINFERENCESMODEL                             *string  `json:"LLM_INFERENCES_MODEL,omitempty" tfsdk:"LLM_INFERENCES_MODEL"`
	LOGINLOGO                                      *string  `json:"LOGIN_LOGO,omitempty" tfsdk:"LOGIN_LOGO"`
	LOGINPAGELAYOUT                                *string  `json:"LOGIN_PAGE_LAYOUT,omitempty" tfsdk:"LOGIN_PAGE_LAYOUT"`
	LOGINPAGEVIDEOURL                              *string  `json:"LOGIN_PAGE_VIDEO_URL,omitempty" tfsdk:"LOGIN_PAGE_VIDEO_URL"`
	MAINTENANCEANNOUNCEMENTNOTIFYBEFOREMINUTES     *int64   `json:"MAINTENANCE_ANNOUNCEMENT_NOTIFY_BEFORE_MINUTES,omitempty" tfsdk:"MAINTENANCE_ANNOUNCEMENT_NOTIFY_BEFORE_MINUTES"`
	MAINTENANCEANNOUNCEMENTNOTIFYSYSTEM            []string `json:"MAINTENANCE_ANNOUNCEMENT_NOTIFY_SYSTEM,omitempty" tfsdk:"MAINTENANCE_ANNOUNCEMENT_NOTIFY_SYSTEM"`
	MARKETPLACEHEROIMAGE                           *string  `json:"MARKETPLACE_HERO_IMAGE,omitempty" tfsdk:"MARKETPLACE_HERO_IMAGE"`
	MARKETPLACELANDINGPAGE                         *string  `json:"MARKETPLACE_LANDING_PAGE,omitempty" tfsdk:"MARKETPLACE_LANDING_PAGE"`
	NOTIFYABOUTRESOURCECHANGE                      *bool    `json:"NOTIFY_ABOUT_RESOURCE_CHANGE,omitempty" tfsdk:"NOTIFY_ABOUT_RESOURCE_CHANGE"`
	NOTIFYSTAFFABOUTAPPROVALS                      *bool    `json:"NOTIFY_STAFF_ABOUT_APPROVALS,omitempty" tfsdk:"NOTIFY_STAFF_ABOUT_APPROVALS"`
	OFFERINGLOGOPLACEHOLDER                        *string  `json:"OFFERING_LOGO_PLACEHOLDER,omitempty" tfsdk:"OFFERING_LOGO_PLACEHOLDER"`
	OIDCACCESSTOKENENABLED                         *bool    `json:"OIDC_ACCESS_TOKEN_ENABLED,omitempty" tfsdk:"OIDC_ACCESS_TOKEN_ENABLED"`
	OIDCAUTHURL                                    *string  `json:"OIDC_AUTH_URL,omitempty" tfsdk:"OIDC_AUTH_URL"`
	OIDCBLOCKCREATIONOFUNINVITEDUSERS              *bool    `json:"OIDC_BLOCK_CREATION_OF_UNINVITED_USERS,omitempty" tfsdk:"OIDC_BLOCK_CREATION_OF_UNINVITED_USERS"`
	OIDCCACHETIMEOUT                               *int64   `json:"OIDC_CACHE_TIMEOUT,omitempty" tfsdk:"OIDC_CACHE_TIMEOUT"`
	OIDCCLIENTID                                   *string  `json:"OIDC_CLIENT_ID,omitempty" tfsdk:"OIDC_CLIENT_ID"`
	OIDCCLIENTSECRET                               *string  `json:"OIDC_CLIENT_SECRET,omitempty" tfsdk:"OIDC_CLIENT_SECRET"`
	OIDCINTROSPECTIONURL                           *string  `json:"OIDC_INTROSPECTION_URL,omitempty" tfsdk:"OIDC_INTROSPECTION_URL"`
	OIDCUSERFIELD                                  *string  `json:"OIDC_USER_FIELD,omitempty" tfsdk:"OIDC_USER_FIELD"`
	ONBOARDINGARIREGISTERBASEURL                   *string  `json:"ONBOARDING_ARIREGISTER_BASE_URL,omitempty" tfsdk:"ONBOARDING_ARIREGISTER_BASE_URL"`
	ONBOARDINGARIREGISTERPASSWORD                  *string  `json:"ONBOARDING_ARIREGISTER_PASSWORD,omitempty" tfsdk:"ONBOARDING_ARIREGISTER_PASSWORD"`
	ONBOARDINGARIREGISTERTIMEOUT                   *int64   `json:"ONBOARDING_ARIREGISTER_TIMEOUT,omitempty" tfsdk:"ONBOARDING_ARIREGISTER_TIMEOUT"`
	ONBOARDINGARIREGISTERUSERNAME                  *string  `json:"ONBOARDING_ARIREGISTER_USERNAME,omitempty" tfsdk:"ONBOARDING_ARIREGISTER_USERNAME"`
	ONBOARDINGBOLAGSVERKETAPIURL                   *string  `json:"ONBOARDING_BOLAGSVERKET_API_URL,omitempty" tfsdk:"ONBOARDING_BOLAGSVERKET_API_URL"`
	ONBOARDINGBOLAGSVERKETCLIENTID                 *string  `json:"ONBOARDING_BOLAGSVERKET_CLIENT_ID,omitempty" tfsdk:"ONBOARDING_BOLAGSVERKET_CLIENT_ID"`
	ONBOARDINGBOLAGSVERKETCLIENTSECRET             *string  `json:"ONBOARDING_BOLAGSVERKET_CLIENT_SECRET,omitempty" tfsdk:"ONBOARDING_BOLAGSVERKET_CLIENT_SECRET"`
	ONBOARDINGBOLAGSVERKETTOKENAPIURL              *string  `json:"ONBOARDING_BOLAGSVERKET_TOKEN_API_URL,omitempty" tfsdk:"ONBOARDING_BOLAGSVERKET_TOKEN_API_URL"`
	ONBOARDINGBREGAPIURL                           *string  `json:"ONBOARDING_BREG_API_URL,omitempty" tfsdk:"ONBOARDING_BREG_API_URL"`
	ONBOARDINGCOUNTRY                              *string  `json:"ONBOARDING_COUNTRY,omitempty" tfsdk:"ONBOARDING_COUNTRY"`
	ONBOARDINGVERIFICATIONEXPIRYHOURS              *int64   `json:"ONBOARDING_VERIFICATION_EXPIRY_HOURS,omitempty" tfsdk:"ONBOARDING_VERIFICATION_EXPIRY_HOURS"`
	ONBOARDINGWICOAPIURL                           *string  `json:"ONBOARDING_WICO_API_URL,omitempty" tfsdk:"ONBOARDING_WICO_API_URL"`
	ONBOARDINGWICOTOKEN                            *string  `json:"ONBOARDING_WICO_TOKEN,omitempty" tfsdk:"ONBOARDING_WICO_TOKEN"`
	POWEREDBYLOGO                                  *string  `json:"POWERED_BY_LOGO,omitempty" tfsdk:"POWERED_BY_LOGO"`
	PROJECTENDDATEMANDATORY                        *bool    `json:"PROJECT_END_DATE_MANDATORY,omitempty" tfsdk:"PROJECT_END_DATE_MANDATORY"`
	PROPOSALREVIEWDURATION                         *int64   `json:"PROPOSAL_REVIEW_DURATION,omitempty" tfsdk:"PROPOSAL_REVIEW_DURATION"`
	RANCHERUSERNAMEINPUTLABEL                      *string  `json:"RANCHER_USERNAME_INPUT_LABEL,omitempty" tfsdk:"RANCHER_USERNAME_INPUT_LABEL"`
	SCRIPTRUNMODE                                  *string  `json:"SCRIPT_RUN_MODE,omitempty" tfsdk:"SCRIPT_RUN_MODE"`
	SHORTPAGETITLE                                 *string  `json:"SHORT_PAGE_TITLE,omitempty" tfsdk:"SHORT_PAGE_TITLE"`
	SIDEBARLOGO                                    *string  `json:"SIDEBAR_LOGO,omitempty" tfsdk:"SIDEBAR_LOGO"`
	SIDEBARLOGODARK                                *string  `json:"SIDEBAR_LOGO_DARK,omitempty" tfsdk:"SIDEBAR_LOGO_DARK"`
	SIDEBARLOGOMOBILE                              *string  `json:"SIDEBAR_LOGO_MOBILE,omitempty" tfsdk:"SIDEBAR_LOGO_MOBILE"`
	SIDEBARSTYLE                                   *string  `json:"SIDEBAR_STYLE,omitempty" tfsdk:"SIDEBAR_STYLE"`
	SITEADDRESS                                    *string  `json:"SITE_ADDRESS,omitempty" tfsdk:"SITE_ADDRESS"`
	SITEDESCRIPTION                                *string  `json:"SITE_DESCRIPTION,omitempty" tfsdk:"SITE_DESCRIPTION"`
	SITEEMAIL                                      *string  `json:"SITE_EMAIL,omitempty" tfsdk:"SITE_EMAIL"`
	SITELOGO                                       *string  `json:"SITE_LOGO,omitempty" tfsdk:"SITE_LOGO"`
	SITENAME                                       *string  `json:"SITE_NAME,omitempty" tfsdk:"SITE_NAME"`
	SITEPHONE                                      *string  `json:"SITE_PHONE,omitempty" tfsdk:"SITE_PHONE"`
	SMAXAFFECTEDRESOURCEFIELD                      *string  `json:"SMAX_AFFECTED_RESOURCE_FIELD,omitempty" tfsdk:"SMAX_AFFECTED_RESOURCE_FIELD"`
	SMAXAPIURL                                     *string  `json:"SMAX_API_URL,omitempty" tfsdk:"SMAX_API_URL"`
	SMAXCREATIONSOURCENAME                         *string  `json:"SMAX_CREATION_SOURCE_NAME,omitempty" tfsdk:"SMAX_CREATION_SOURCE_NAME"`
	SMAXLOGIN                                      *string  `json:"SMAX_LOGIN,omitempty" tfsdk:"SMAX_LOGIN"`
	SMAXORGANISATIONFIELD                          *string  `json:"SMAX_ORGANISATION_FIELD,omitempty" tfsdk:"SMAX_ORGANISATION_FIELD"`
	SMAXPASSWORD                                   *string  `json:"SMAX_PASSWORD,omitempty" tfsdk:"SMAX_PASSWORD"`
	SMAXPROJECTFIELD                               *string  `json:"SMAX_PROJECT_FIELD,omitempty" tfsdk:"SMAX_PROJECT_FIELD"`
	SMAXREQUESTSOFFERING                           *string  `json:"SMAX_REQUESTS_OFFERING,omitempty" tfsdk:"SMAX_REQUESTS_OFFERING"`
	SMAXSECONDSTOWAIT                              *int64   `json:"SMAX_SECONDS_TO_WAIT,omitempty" tfsdk:"SMAX_SECONDS_TO_WAIT"`
	SMAXTENANTID                                   *string  `json:"SMAX_TENANT_ID,omitempty" tfsdk:"SMAX_TENANT_ID"`
	SMAXTIMESTOPULL                                *int64   `json:"SMAX_TIMES_TO_PULL,omitempty" tfsdk:"SMAX_TIMES_TO_PULL"`
	SMAXVERIFYSSL                                  *bool    `json:"SMAX_VERIFY_SSL,omitempty" tfsdk:"SMAX_VERIFY_SSL"`
	SOFTWARECATALOGCLEANUPENABLED                  *bool    `json:"SOFTWARE_CATALOG_CLEANUP_ENABLED,omitempty" tfsdk:"SOFTWARE_CATALOG_CLEANUP_ENABLED"`
	SOFTWARECATALOGEESSIAPIURL                     *string  `json:"SOFTWARE_CATALOG_EESSI_API_URL,omitempty" tfsdk:"SOFTWARE_CATALOG_EESSI_API_URL"`
	SOFTWARECATALOGEESSIINCLUDEEXTENSIONS          *bool    `json:"SOFTWARE_CATALOG_EESSI_INCLUDE_EXTENSIONS,omitempty" tfsdk:"SOFTWARE_CATALOG_EESSI_INCLUDE_EXTENSIONS"`
	SOFTWARECATALOGEESSIUPDATEENABLED              *bool    `json:"SOFTWARE_CATALOG_EESSI_UPDATE_ENABLED,omitempty" tfsdk:"SOFTWARE_CATALOG_EESSI_UPDATE_ENABLED"`
	SOFTWARECATALOGEESSIVERSION                    *string  `json:"SOFTWARE_CATALOG_EESSI_VERSION,omitempty" tfsdk:"SOFTWARE_CATALOG_EESSI_VERSION"`
	SOFTWARECATALOGRETENTIONDAYS                   *int64   `json:"SOFTWARE_CATALOG_RETENTION_DAYS,omitempty" tfsdk:"SOFTWARE_CATALOG_RETENTION_DAYS"`
	SOFTWARECATALOGSPACKDATAURL                    *string  `json:"SOFTWARE_CATALOG_SPACK_DATA_URL,omitempty" tfsdk:"SOFTWARE_CATALOG_SPACK_DATA_URL"`
	SOFTWARECATALOGSPACKUPDATEENABLED              *bool    `json:"SOFTWARE_CATALOG_SPACK_UPDATE_ENABLED,omitempty" tfsdk:"SOFTWARE_CATALOG_SPACK_UPDATE_ENABLED"`
	SOFTWARECATALOGSPACKVERSION                    *string  `json:"SOFTWARE_CATALOG_SPACK_VERSION,omitempty" tfsdk:"SOFTWARE_CATALOG_SPACK_VERSION"`
	SOFTWARECATALOGUPDATEEXISTINGPACKAGES          *bool    `json:"SOFTWARE_CATALOG_UPDATE_EXISTING_PACKAGES,omitempty" tfsdk:"SOFTWARE_CATALOG_UPDATE_EXISTING_PACKAGES"`
	SUPPORTPORTALURL                               *string  `json:"SUPPORT_PORTAL_URL,omitempty" tfsdk:"SUPPORT_PORTAL_URL"`
	TELEMETRYURL                                   *string  `json:"TELEMETRY_URL,omitempty" tfsdk:"TELEMETRY_URL"`
	TELEMETRYVERSION                               *int64   `json:"TELEMETRY_VERSION,omitempty" tfsdk:"TELEMETRY_VERSION"`
	THUMBNAILSIZE                                  *string  `json:"THUMBNAIL_SIZE,omitempty" tfsdk:"THUMBNAIL_SIZE"`
	USERTABLECOLUMNS                               *string  `json:"USER_TABLE_COLUMNS,omitempty" tfsdk:"USER_TABLE_COLUMNS"`
	WALDURAUTHSOCIALROLECLAIM                      *string  `json:"WALDUR_AUTH_SOCIAL_ROLE_CLAIM,omitempty" tfsdk:"WALDUR_AUTH_SOCIAL_ROLE_CLAIM"`
	WALDURSUPPORTACTIVEBACKENDTYPE                 *string  `json:"WALDUR_SUPPORT_ACTIVE_BACKEND_TYPE,omitempty" tfsdk:"WALDUR_SUPPORT_ACTIVE_BACKEND_TYPE"`
	WALDURSUPPORTDISPLAYREQUESTTYPE                *bool    `json:"WALDUR_SUPPORT_DISPLAY_REQUEST_TYPE,omitempty" tfsdk:"WALDUR_SUPPORT_DISPLAY_REQUEST_TYPE"`
	WALDURSUPPORTENABLED                           *bool    `json:"WALDUR_SUPPORT_ENABLED,omitempty" tfsdk:"WALDUR_SUPPORT_ENABLED"`
	ZAMMADAPIURL                                   *string  `json:"ZAMMAD_API_URL,omitempty" tfsdk:"ZAMMAD_API_URL"`
	ZAMMADARTICLETYPE                              *string  `json:"ZAMMAD_ARTICLE_TYPE,omitempty" tfsdk:"ZAMMAD_ARTICLE_TYPE"`
	ZAMMADCOMMENTCOOLDOWNDURATION                  *int64   `json:"ZAMMAD_COMMENT_COOLDOWN_DURATION,omitempty" tfsdk:"ZAMMAD_COMMENT_COOLDOWN_DURATION"`
	ZAMMADCOMMENTMARKER                            *string  `json:"ZAMMAD_COMMENT_MARKER,omitempty" tfsdk:"ZAMMAD_COMMENT_MARKER"`
	ZAMMADCOMMENTPREFIX                            *string  `json:"ZAMMAD_COMMENT_PREFIX,omitempty" tfsdk:"ZAMMAD_COMMENT_PREFIX"`
	ZAMMADGROUP                                    *string  `json:"ZAMMAD_GROUP,omitempty" tfsdk:"ZAMMAD_GROUP"`
	ZAMMADTOKEN                                    *string  `json:"ZAMMAD_TOKEN,omitempty" tfsdk:"ZAMMAD_TOKEN"`
}

type ConstanceSettingsRequestMultipart struct {
	ANONYMOUSUSERCANVIEWOFFERINGS                  *bool    `json:"ANONYMOUS_USER_CAN_VIEW_OFFERINGS,omitempty" tfsdk:"ANONYMOUS_USER_CAN_VIEW_OFFERINGS"`
	ANONYMOUSUSERCANVIEWPLANS                      *bool    `json:"ANONYMOUS_USER_CAN_VIEW_PLANS,omitempty" tfsdk:"ANONYMOUS_USER_CAN_VIEW_PLANS"`
	ATLASSIANAFFECTEDRESOURCEFIELD                 *string  `json:"ATLASSIAN_AFFECTED_RESOURCE_FIELD,omitempty" tfsdk:"ATLASSIAN_AFFECTED_RESOURCE_FIELD"`
	ATLASSIANAPIURL                                *string  `json:"ATLASSIAN_API_URL,omitempty" tfsdk:"ATLASSIAN_API_URL"`
	ATLASSIANCALLERFIELD                           *string  `json:"ATLASSIAN_CALLER_FIELD,omitempty" tfsdk:"ATLASSIAN_CALLER_FIELD"`
	ATLASSIANCUSTOMISSUEFIELDMAPPINGENABLED        *bool    `json:"ATLASSIAN_CUSTOM_ISSUE_FIELD_MAPPING_ENABLED,omitempty" tfsdk:"ATLASSIAN_CUSTOM_ISSUE_FIELD_MAPPING_ENABLED"`
	ATLASSIANDEFAULTOFFERINGISSUETYPE              *string  `json:"ATLASSIAN_DEFAULT_OFFERING_ISSUE_TYPE,omitempty" tfsdk:"ATLASSIAN_DEFAULT_OFFERING_ISSUE_TYPE"`
	ATLASSIANDESCRIPTIONTEMPLATE                   *string  `json:"ATLASSIAN_DESCRIPTION_TEMPLATE,omitempty" tfsdk:"ATLASSIAN_DESCRIPTION_TEMPLATE"`
	ATLASSIANEMAIL                                 *string  `json:"ATLASSIAN_EMAIL,omitempty" tfsdk:"ATLASSIAN_EMAIL"`
	ATLASSIANEXCLUDEDATTACHMENTTYPES               *string  `json:"ATLASSIAN_EXCLUDED_ATTACHMENT_TYPES,omitempty" tfsdk:"ATLASSIAN_EXCLUDED_ATTACHMENT_TYPES"`
	ATLASSIANIMPACTFIELD                           *string  `json:"ATLASSIAN_IMPACT_FIELD,omitempty" tfsdk:"ATLASSIAN_IMPACT_FIELD"`
	ATLASSIANLINKEDISSUETYPE                       *string  `json:"ATLASSIAN_LINKED_ISSUE_TYPE,omitempty" tfsdk:"ATLASSIAN_LINKED_ISSUE_TYPE"`
	ATLASSIANMAPWALDURUSERSTOSERVICEDESKAGENTS     *bool    `json:"ATLASSIAN_MAP_WALDUR_USERS_TO_SERVICEDESK_AGENTS,omitempty" tfsdk:"ATLASSIAN_MAP_WALDUR_USERS_TO_SERVICEDESK_AGENTS"`
	ATLASSIANOAUTH2ACCESSTOKEN                     *string  `json:"ATLASSIAN_OAUTH2_ACCESS_TOKEN,omitempty" tfsdk:"ATLASSIAN_OAUTH2_ACCESS_TOKEN"`
	ATLASSIANOAUTH2CLIENTID                        *string  `json:"ATLASSIAN_OAUTH2_CLIENT_ID,omitempty" tfsdk:"ATLASSIAN_OAUTH2_CLIENT_ID"`
	ATLASSIANOAUTH2TOKENTYPE                       *string  `json:"ATLASSIAN_OAUTH2_TOKEN_TYPE,omitempty" tfsdk:"ATLASSIAN_OAUTH2_TOKEN_TYPE"`
	ATLASSIANORGANISATIONFIELD                     *string  `json:"ATLASSIAN_ORGANISATION_FIELD,omitempty" tfsdk:"ATLASSIAN_ORGANISATION_FIELD"`
	ATLASSIANPASSWORD                              *string  `json:"ATLASSIAN_PASSWORD,omitempty" tfsdk:"ATLASSIAN_PASSWORD"`
	ATLASSIANPERSONALACCESSTOKEN                   *string  `json:"ATLASSIAN_PERSONAL_ACCESS_TOKEN,omitempty" tfsdk:"ATLASSIAN_PERSONAL_ACCESS_TOKEN"`
	ATLASSIANPROJECTFIELD                          *string  `json:"ATLASSIAN_PROJECT_FIELD,omitempty" tfsdk:"ATLASSIAN_PROJECT_FIELD"`
	ATLASSIANPROJECTID                             *string  `json:"ATLASSIAN_PROJECT_ID,omitempty" tfsdk:"ATLASSIAN_PROJECT_ID"`
	ATLASSIANREPORTERFIELD                         *string  `json:"ATLASSIAN_REPORTER_FIELD,omitempty" tfsdk:"ATLASSIAN_REPORTER_FIELD"`
	ATLASSIANREQUESTFEEDBACKFIELD                  *string  `json:"ATLASSIAN_REQUEST_FEEDBACK_FIELD,omitempty" tfsdk:"ATLASSIAN_REQUEST_FEEDBACK_FIELD"`
	ATLASSIANRESOLUTIONSLAFIELD                    *string  `json:"ATLASSIAN_RESOLUTION_SLA_FIELD,omitempty" tfsdk:"ATLASSIAN_RESOLUTION_SLA_FIELD"`
	ATLASSIANSATISFACTIONFIELD                     *string  `json:"ATLASSIAN_SATISFACTION_FIELD,omitempty" tfsdk:"ATLASSIAN_SATISFACTION_FIELD"`
	ATLASSIANSHAREDUSERNAME                        *bool    `json:"ATLASSIAN_SHARED_USERNAME,omitempty" tfsdk:"ATLASSIAN_SHARED_USERNAME"`
	ATLASSIANSLAFIELD                              *string  `json:"ATLASSIAN_SLA_FIELD,omitempty" tfsdk:"ATLASSIAN_SLA_FIELD"`
	ATLASSIANSUMMARYTEMPLATE                       *string  `json:"ATLASSIAN_SUMMARY_TEMPLATE,omitempty" tfsdk:"ATLASSIAN_SUMMARY_TEMPLATE"`
	ATLASSIANTEMPLATEFIELD                         *string  `json:"ATLASSIAN_TEMPLATE_FIELD,omitempty" tfsdk:"ATLASSIAN_TEMPLATE_FIELD"`
	ATLASSIANTOKEN                                 *string  `json:"ATLASSIAN_TOKEN,omitempty" tfsdk:"ATLASSIAN_TOKEN"`
	ATLASSIANUSERNAME                              *string  `json:"ATLASSIAN_USERNAME,omitempty" tfsdk:"ATLASSIAN_USERNAME"`
	ATLASSIANUSEOLDAPI                             *bool    `json:"ATLASSIAN_USE_OLD_API,omitempty" tfsdk:"ATLASSIAN_USE_OLD_API"`
	ATLASSIANVERIFYSSL                             *bool    `json:"ATLASSIAN_VERIFY_SSL,omitempty" tfsdk:"ATLASSIAN_VERIFY_SSL"`
	ATLASSIANWALDURBACKENDIDFIELD                  *string  `json:"ATLASSIAN_WALDUR_BACKEND_ID_FIELD,omitempty" tfsdk:"ATLASSIAN_WALDUR_BACKEND_ID_FIELD"`
	AUTOAPPROVEUSERTOS                             *bool    `json:"AUTO_APPROVE_USER_TOS,omitempty" tfsdk:"AUTO_APPROVE_USER_TOS"`
	BRANDCOLOR                                     *string  `json:"BRAND_COLOR,omitempty" tfsdk:"BRAND_COLOR"`
	CALLMANAGEMENTHEROIMAGE                        *string  `json:"CALL_MANAGEMENT_HERO_IMAGE,omitempty" tfsdk:"CALL_MANAGEMENT_HERO_IMAGE"`
	COMMONFOOTERHTML                               *string  `json:"COMMON_FOOTER_HTML,omitempty" tfsdk:"COMMON_FOOTER_HTML"`
	COMMONFOOTERTEXT                               *string  `json:"COMMON_FOOTER_TEXT,omitempty" tfsdk:"COMMON_FOOTER_TEXT"`
	COUNTRIES                                      []string `json:"COUNTRIES,omitempty" tfsdk:"COUNTRIES"`
	CURRENCYNAME                                   *string  `json:"CURRENCY_NAME,omitempty" tfsdk:"CURRENCY_NAME"`
	DEACTIVATEUSERIFNOROLES                        *bool    `json:"DEACTIVATE_USER_IF_NO_ROLES,omitempty" tfsdk:"DEACTIVATE_USER_IF_NO_ROLES"`
	DEFAULTIDP                                     *string  `json:"DEFAULT_IDP,omitempty" tfsdk:"DEFAULT_IDP"`
	DISABLEDOFFERINGTYPES                          []string `json:"DISABLED_OFFERING_TYPES,omitempty" tfsdk:"DISABLED_OFFERING_TYPES"`
	DISABLEDARKTHEME                               *bool    `json:"DISABLE_DARK_THEME,omitempty" tfsdk:"DISABLE_DARK_THEME"`
	DISABLESENDINGNOTIFICATIONSABOUTRESOURCEUPDATE *bool    `json:"DISABLE_SENDING_NOTIFICATIONS_ABOUT_RESOURCE_UPDATE,omitempty" tfsdk:"DISABLE_SENDING_NOTIFICATIONS_ABOUT_RESOURCE_UPDATE"`
	DOCKERCLIENT                                   *string  `json:"DOCKER_CLIENT,omitempty" tfsdk:"DOCKER_CLIENT"`
	DOCKERIMAGES                                   *string  `json:"DOCKER_IMAGES,omitempty" tfsdk:"DOCKER_IMAGES"`
	DOCKERREMOVECONTAINER                          *bool    `json:"DOCKER_REMOVE_CONTAINER,omitempty" tfsdk:"DOCKER_REMOVE_CONTAINER"`
	DOCKERRUNOPTIONS                               *string  `json:"DOCKER_RUN_OPTIONS,omitempty" tfsdk:"DOCKER_RUN_OPTIONS"`
	DOCKERSCRIPTDIR                                *string  `json:"DOCKER_SCRIPT_DIR,omitempty" tfsdk:"DOCKER_SCRIPT_DIR"`
	DOCKERVOLUMENAME                               *string  `json:"DOCKER_VOLUME_NAME,omitempty" tfsdk:"DOCKER_VOLUME_NAME"`
	DOCSURL                                        *string  `json:"DOCS_URL,omitempty" tfsdk:"DOCS_URL"`
	ENABLEMOCKCOURSEACCOUNTBACKEND                 *bool    `json:"ENABLE_MOCK_COURSE_ACCOUNT_BACKEND,omitempty" tfsdk:"ENABLE_MOCK_COURSE_ACCOUNT_BACKEND"`
	ENABLEMOCKSERVICEACCOUNTBACKEND                *bool    `json:"ENABLE_MOCK_SERVICE_ACCOUNT_BACKEND,omitempty" tfsdk:"ENABLE_MOCK_SERVICE_ACCOUNT_BACKEND"`
	ENABLEORDERSTARTDATE                           *bool    `json:"ENABLE_ORDER_START_DATE,omitempty" tfsdk:"ENABLE_ORDER_START_DATE"`
	ENABLESTALERESOURCENOTIFICATIONS               *bool    `json:"ENABLE_STALE_RESOURCE_NOTIFICATIONS,omitempty" tfsdk:"ENABLE_STALE_RESOURCE_NOTIFICATIONS"`
	ENABLESTRICTCHECKACCEPTINGINVITATION           *bool    `json:"ENABLE_STRICT_CHECK_ACCEPTING_INVITATION,omitempty" tfsdk:"ENABLE_STRICT_CHECK_ACCEPTING_INVITATION"`
	ENFORCEUSERCONSENTFOROFFERINGS                 *bool    `json:"ENFORCE_USER_CONSENT_FOR_OFFERINGS,omitempty" tfsdk:"ENFORCE_USER_CONSENT_FOR_OFFERINGS"`
	FAVICON                                        *string  `json:"FAVICON,omitempty" tfsdk:"FAVICON"`
	FREEIPABLACKLISTEDUSERNAMES                    []string `json:"FREEIPA_BLACKLISTED_USERNAMES,omitempty" tfsdk:"FREEIPA_BLACKLISTED_USERNAMES"`
	FREEIPAENABLED                                 *bool    `json:"FREEIPA_ENABLED,omitempty" tfsdk:"FREEIPA_ENABLED"`
	FREEIPAGROUPNAMEPREFIX                         *string  `json:"FREEIPA_GROUPNAME_PREFIX,omitempty" tfsdk:"FREEIPA_GROUPNAME_PREFIX"`
	FREEIPAGROUPSYNCHRONIZATIONENABLED             *bool    `json:"FREEIPA_GROUP_SYNCHRONIZATION_ENABLED,omitempty" tfsdk:"FREEIPA_GROUP_SYNCHRONIZATION_ENABLED"`
	FREEIPAHOSTNAME                                *string  `json:"FREEIPA_HOSTNAME,omitempty" tfsdk:"FREEIPA_HOSTNAME"`
	FREEIPAPASSWORD                                *string  `json:"FREEIPA_PASSWORD,omitempty" tfsdk:"FREEIPA_PASSWORD"`
	FREEIPAUSERNAME                                *string  `json:"FREEIPA_USERNAME,omitempty" tfsdk:"FREEIPA_USERNAME"`
	FREEIPAUSERNAMEPREFIX                          *string  `json:"FREEIPA_USERNAME_PREFIX,omitempty" tfsdk:"FREEIPA_USERNAME_PREFIX"`
	FREEIPAVERIFYSSL                               *bool    `json:"FREEIPA_VERIFY_SSL,omitempty" tfsdk:"FREEIPA_VERIFY_SSL"`
	FULLPAGETITLE                                  *string  `json:"FULL_PAGE_TITLE,omitempty" tfsdk:"FULL_PAGE_TITLE"`
	HEROIMAGE                                      *string  `json:"HERO_IMAGE,omitempty" tfsdk:"HERO_IMAGE"`
	HEROLINKLABEL                                  *string  `json:"HERO_LINK_LABEL,omitempty" tfsdk:"HERO_LINK_LABEL"`
	HEROLINKURL                                    *string  `json:"HERO_LINK_URL,omitempty" tfsdk:"HERO_LINK_URL"`
	HOMEPORTURL                                    *string  `json:"HOMEPORT_URL,omitempty" tfsdk:"HOMEPORT_URL"`
	INVITATIONDISABLEMULTIPLEROLES                 *bool    `json:"INVITATION_DISABLE_MULTIPLE_ROLES,omitempty" tfsdk:"INVITATION_DISABLE_MULTIPLE_ROLES"`
	K8SCONFIGPATH                                  *string  `json:"K8S_CONFIG_PATH,omitempty" tfsdk:"K8S_CONFIG_PATH"`
	K8SJOBTIMEOUT                                  *int64   `json:"K8S_JOB_TIMEOUT,omitempty" tfsdk:"K8S_JOB_TIMEOUT"`
	K8SNAMESPACE                                   *string  `json:"K8S_NAMESPACE,omitempty" tfsdk:"K8S_NAMESPACE"`
	KEYCLOAKICON                                   *string  `json:"KEYCLOAK_ICON,omitempty" tfsdk:"KEYCLOAK_ICON"`
	LANGUAGECHOICES                                *string  `json:"LANGUAGE_CHOICES,omitempty" tfsdk:"LANGUAGE_CHOICES"`
	LLMCHATENABLED                                 *bool    `json:"LLM_CHAT_ENABLED,omitempty" tfsdk:"LLM_CHAT_ENABLED"`
	LLMINFERENCESAPITOKEN                          *string  `json:"LLM_INFERENCES_API_TOKEN,omitempty" tfsdk:"LLM_INFERENCES_API_TOKEN"`
	LLMINFERENCESAPIURL                            *string  `json:"LLM_INFERENCES_API_URL,omitempty" tfsdk:"LLM_INFERENCES_API_URL"`
	LLMINFERENCESBACKENDTYPE                       *string  `json:"LLM_INFERENCES_BACKEND_TYPE,omitempty" tfsdk:"LLM_INFERENCES_BACKEND_TYPE"`
	LLMINFERENCESMODEL                             *string  `json:"LLM_INFERENCES_MODEL,omitempty" tfsdk:"LLM_INFERENCES_MODEL"`
	LOGINLOGO                                      *string  `json:"LOGIN_LOGO,omitempty" tfsdk:"LOGIN_LOGO"`
	LOGINPAGELAYOUT                                *string  `json:"LOGIN_PAGE_LAYOUT,omitempty" tfsdk:"LOGIN_PAGE_LAYOUT"`
	LOGINPAGEVIDEOURL                              *string  `json:"LOGIN_PAGE_VIDEO_URL,omitempty" tfsdk:"LOGIN_PAGE_VIDEO_URL"`
	MAINTENANCEANNOUNCEMENTNOTIFYBEFOREMINUTES     *int64   `json:"MAINTENANCE_ANNOUNCEMENT_NOTIFY_BEFORE_MINUTES,omitempty" tfsdk:"MAINTENANCE_ANNOUNCEMENT_NOTIFY_BEFORE_MINUTES"`
	MAINTENANCEANNOUNCEMENTNOTIFYSYSTEM            []string `json:"MAINTENANCE_ANNOUNCEMENT_NOTIFY_SYSTEM,omitempty" tfsdk:"MAINTENANCE_ANNOUNCEMENT_NOTIFY_SYSTEM"`
	MARKETPLACEHEROIMAGE                           *string  `json:"MARKETPLACE_HERO_IMAGE,omitempty" tfsdk:"MARKETPLACE_HERO_IMAGE"`
	MARKETPLACELANDINGPAGE                         *string  `json:"MARKETPLACE_LANDING_PAGE,omitempty" tfsdk:"MARKETPLACE_LANDING_PAGE"`
	NOTIFYABOUTRESOURCECHANGE                      *bool    `json:"NOTIFY_ABOUT_RESOURCE_CHANGE,omitempty" tfsdk:"NOTIFY_ABOUT_RESOURCE_CHANGE"`
	NOTIFYSTAFFABOUTAPPROVALS                      *bool    `json:"NOTIFY_STAFF_ABOUT_APPROVALS,omitempty" tfsdk:"NOTIFY_STAFF_ABOUT_APPROVALS"`
	OFFERINGLOGOPLACEHOLDER                        *string  `json:"OFFERING_LOGO_PLACEHOLDER,omitempty" tfsdk:"OFFERING_LOGO_PLACEHOLDER"`
	OIDCACCESSTOKENENABLED                         *bool    `json:"OIDC_ACCESS_TOKEN_ENABLED,omitempty" tfsdk:"OIDC_ACCESS_TOKEN_ENABLED"`
	OIDCAUTHURL                                    *string  `json:"OIDC_AUTH_URL,omitempty" tfsdk:"OIDC_AUTH_URL"`
	OIDCBLOCKCREATIONOFUNINVITEDUSERS              *bool    `json:"OIDC_BLOCK_CREATION_OF_UNINVITED_USERS,omitempty" tfsdk:"OIDC_BLOCK_CREATION_OF_UNINVITED_USERS"`
	OIDCCACHETIMEOUT                               *int64   `json:"OIDC_CACHE_TIMEOUT,omitempty" tfsdk:"OIDC_CACHE_TIMEOUT"`
	OIDCCLIENTID                                   *string  `json:"OIDC_CLIENT_ID,omitempty" tfsdk:"OIDC_CLIENT_ID"`
	OIDCCLIENTSECRET                               *string  `json:"OIDC_CLIENT_SECRET,omitempty" tfsdk:"OIDC_CLIENT_SECRET"`
	OIDCINTROSPECTIONURL                           *string  `json:"OIDC_INTROSPECTION_URL,omitempty" tfsdk:"OIDC_INTROSPECTION_URL"`
	OIDCUSERFIELD                                  *string  `json:"OIDC_USER_FIELD,omitempty" tfsdk:"OIDC_USER_FIELD"`
	ONBOARDINGARIREGISTERBASEURL                   *string  `json:"ONBOARDING_ARIREGISTER_BASE_URL,omitempty" tfsdk:"ONBOARDING_ARIREGISTER_BASE_URL"`
	ONBOARDINGARIREGISTERPASSWORD                  *string  `json:"ONBOARDING_ARIREGISTER_PASSWORD,omitempty" tfsdk:"ONBOARDING_ARIREGISTER_PASSWORD"`
	ONBOARDINGARIREGISTERTIMEOUT                   *int64   `json:"ONBOARDING_ARIREGISTER_TIMEOUT,omitempty" tfsdk:"ONBOARDING_ARIREGISTER_TIMEOUT"`
	ONBOARDINGARIREGISTERUSERNAME                  *string  `json:"ONBOARDING_ARIREGISTER_USERNAME,omitempty" tfsdk:"ONBOARDING_ARIREGISTER_USERNAME"`
	ONBOARDINGBOLAGSVERKETAPIURL                   *string  `json:"ONBOARDING_BOLAGSVERKET_API_URL,omitempty" tfsdk:"ONBOARDING_BOLAGSVERKET_API_URL"`
	ONBOARDINGBOLAGSVERKETCLIENTID                 *string  `json:"ONBOARDING_BOLAGSVERKET_CLIENT_ID,omitempty" tfsdk:"ONBOARDING_BOLAGSVERKET_CLIENT_ID"`
	ONBOARDINGBOLAGSVERKETCLIENTSECRET             *string  `json:"ONBOARDING_BOLAGSVERKET_CLIENT_SECRET,omitempty" tfsdk:"ONBOARDING_BOLAGSVERKET_CLIENT_SECRET"`
	ONBOARDINGBOLAGSVERKETTOKENAPIURL              *string  `json:"ONBOARDING_BOLAGSVERKET_TOKEN_API_URL,omitempty" tfsdk:"ONBOARDING_BOLAGSVERKET_TOKEN_API_URL"`
	ONBOARDINGBREGAPIURL                           *string  `json:"ONBOARDING_BREG_API_URL,omitempty" tfsdk:"ONBOARDING_BREG_API_URL"`
	ONBOARDINGCOUNTRY                              *string  `json:"ONBOARDING_COUNTRY,omitempty" tfsdk:"ONBOARDING_COUNTRY"`
	ONBOARDINGVERIFICATIONEXPIRYHOURS              *int64   `json:"ONBOARDING_VERIFICATION_EXPIRY_HOURS,omitempty" tfsdk:"ONBOARDING_VERIFICATION_EXPIRY_HOURS"`
	ONBOARDINGWICOAPIURL                           *string  `json:"ONBOARDING_WICO_API_URL,omitempty" tfsdk:"ONBOARDING_WICO_API_URL"`
	ONBOARDINGWICOTOKEN                            *string  `json:"ONBOARDING_WICO_TOKEN,omitempty" tfsdk:"ONBOARDING_WICO_TOKEN"`
	POWEREDBYLOGO                                  *string  `json:"POWERED_BY_LOGO,omitempty" tfsdk:"POWERED_BY_LOGO"`
	PROJECTENDDATEMANDATORY                        *bool    `json:"PROJECT_END_DATE_MANDATORY,omitempty" tfsdk:"PROJECT_END_DATE_MANDATORY"`
	PROPOSALREVIEWDURATION                         *int64   `json:"PROPOSAL_REVIEW_DURATION,omitempty" tfsdk:"PROPOSAL_REVIEW_DURATION"`
	RANCHERUSERNAMEINPUTLABEL                      *string  `json:"RANCHER_USERNAME_INPUT_LABEL,omitempty" tfsdk:"RANCHER_USERNAME_INPUT_LABEL"`
	SCRIPTRUNMODE                                  *string  `json:"SCRIPT_RUN_MODE,omitempty" tfsdk:"SCRIPT_RUN_MODE"`
	SHORTPAGETITLE                                 *string  `json:"SHORT_PAGE_TITLE,omitempty" tfsdk:"SHORT_PAGE_TITLE"`
	SIDEBARLOGO                                    *string  `json:"SIDEBAR_LOGO,omitempty" tfsdk:"SIDEBAR_LOGO"`
	SIDEBARLOGODARK                                *string  `json:"SIDEBAR_LOGO_DARK,omitempty" tfsdk:"SIDEBAR_LOGO_DARK"`
	SIDEBARLOGOMOBILE                              *string  `json:"SIDEBAR_LOGO_MOBILE,omitempty" tfsdk:"SIDEBAR_LOGO_MOBILE"`
	SIDEBARSTYLE                                   *string  `json:"SIDEBAR_STYLE,omitempty" tfsdk:"SIDEBAR_STYLE"`
	SITEADDRESS                                    *string  `json:"SITE_ADDRESS,omitempty" tfsdk:"SITE_ADDRESS"`
	SITEDESCRIPTION                                *string  `json:"SITE_DESCRIPTION,omitempty" tfsdk:"SITE_DESCRIPTION"`
	SITEEMAIL                                      *string  `json:"SITE_EMAIL,omitempty" tfsdk:"SITE_EMAIL"`
	SITELOGO                                       *string  `json:"SITE_LOGO,omitempty" tfsdk:"SITE_LOGO"`
	SITENAME                                       *string  `json:"SITE_NAME,omitempty" tfsdk:"SITE_NAME"`
	SITEPHONE                                      *string  `json:"SITE_PHONE,omitempty" tfsdk:"SITE_PHONE"`
	SMAXAFFECTEDRESOURCEFIELD                      *string  `json:"SMAX_AFFECTED_RESOURCE_FIELD,omitempty" tfsdk:"SMAX_AFFECTED_RESOURCE_FIELD"`
	SMAXAPIURL                                     *string  `json:"SMAX_API_URL,omitempty" tfsdk:"SMAX_API_URL"`
	SMAXCREATIONSOURCENAME                         *string  `json:"SMAX_CREATION_SOURCE_NAME,omitempty" tfsdk:"SMAX_CREATION_SOURCE_NAME"`
	SMAXLOGIN                                      *string  `json:"SMAX_LOGIN,omitempty" tfsdk:"SMAX_LOGIN"`
	SMAXORGANISATIONFIELD                          *string  `json:"SMAX_ORGANISATION_FIELD,omitempty" tfsdk:"SMAX_ORGANISATION_FIELD"`
	SMAXPASSWORD                                   *string  `json:"SMAX_PASSWORD,omitempty" tfsdk:"SMAX_PASSWORD"`
	SMAXPROJECTFIELD                               *string  `json:"SMAX_PROJECT_FIELD,omitempty" tfsdk:"SMAX_PROJECT_FIELD"`
	SMAXREQUESTSOFFERING                           *string  `json:"SMAX_REQUESTS_OFFERING,omitempty" tfsdk:"SMAX_REQUESTS_OFFERING"`
	SMAXSECONDSTOWAIT                              *int64   `json:"SMAX_SECONDS_TO_WAIT,omitempty" tfsdk:"SMAX_SECONDS_TO_WAIT"`
	SMAXTENANTID                                   *string  `json:"SMAX_TENANT_ID,omitempty" tfsdk:"SMAX_TENANT_ID"`
	SMAXTIMESTOPULL                                *int64   `json:"SMAX_TIMES_TO_PULL,omitempty" tfsdk:"SMAX_TIMES_TO_PULL"`
	SMAXVERIFYSSL                                  *bool    `json:"SMAX_VERIFY_SSL,omitempty" tfsdk:"SMAX_VERIFY_SSL"`
	SOFTWARECATALOGCLEANUPENABLED                  *bool    `json:"SOFTWARE_CATALOG_CLEANUP_ENABLED,omitempty" tfsdk:"SOFTWARE_CATALOG_CLEANUP_ENABLED"`
	SOFTWARECATALOGEESSIAPIURL                     *string  `json:"SOFTWARE_CATALOG_EESSI_API_URL,omitempty" tfsdk:"SOFTWARE_CATALOG_EESSI_API_URL"`
	SOFTWARECATALOGEESSIINCLUDEEXTENSIONS          *bool    `json:"SOFTWARE_CATALOG_EESSI_INCLUDE_EXTENSIONS,omitempty" tfsdk:"SOFTWARE_CATALOG_EESSI_INCLUDE_EXTENSIONS"`
	SOFTWARECATALOGEESSIUPDATEENABLED              *bool    `json:"SOFTWARE_CATALOG_EESSI_UPDATE_ENABLED,omitempty" tfsdk:"SOFTWARE_CATALOG_EESSI_UPDATE_ENABLED"`
	SOFTWARECATALOGEESSIVERSION                    *string  `json:"SOFTWARE_CATALOG_EESSI_VERSION,omitempty" tfsdk:"SOFTWARE_CATALOG_EESSI_VERSION"`
	SOFTWARECATALOGRETENTIONDAYS                   *int64   `json:"SOFTWARE_CATALOG_RETENTION_DAYS,omitempty" tfsdk:"SOFTWARE_CATALOG_RETENTION_DAYS"`
	SOFTWARECATALOGSPACKDATAURL                    *string  `json:"SOFTWARE_CATALOG_SPACK_DATA_URL,omitempty" tfsdk:"SOFTWARE_CATALOG_SPACK_DATA_URL"`
	SOFTWARECATALOGSPACKUPDATEENABLED              *bool    `json:"SOFTWARE_CATALOG_SPACK_UPDATE_ENABLED,omitempty" tfsdk:"SOFTWARE_CATALOG_SPACK_UPDATE_ENABLED"`
	SOFTWARECATALOGSPACKVERSION                    *string  `json:"SOFTWARE_CATALOG_SPACK_VERSION,omitempty" tfsdk:"SOFTWARE_CATALOG_SPACK_VERSION"`
	SOFTWARECATALOGUPDATEEXISTINGPACKAGES          *bool    `json:"SOFTWARE_CATALOG_UPDATE_EXISTING_PACKAGES,omitempty" tfsdk:"SOFTWARE_CATALOG_UPDATE_EXISTING_PACKAGES"`
	SUPPORTPORTALURL                               *string  `json:"SUPPORT_PORTAL_URL,omitempty" tfsdk:"SUPPORT_PORTAL_URL"`
	TELEMETRYURL                                   *string  `json:"TELEMETRY_URL,omitempty" tfsdk:"TELEMETRY_URL"`
	TELEMETRYVERSION                               *int64   `json:"TELEMETRY_VERSION,omitempty" tfsdk:"TELEMETRY_VERSION"`
	THUMBNAILSIZE                                  *string  `json:"THUMBNAIL_SIZE,omitempty" tfsdk:"THUMBNAIL_SIZE"`
	USERTABLECOLUMNS                               *string  `json:"USER_TABLE_COLUMNS,omitempty" tfsdk:"USER_TABLE_COLUMNS"`
	WALDURAUTHSOCIALROLECLAIM                      *string  `json:"WALDUR_AUTH_SOCIAL_ROLE_CLAIM,omitempty" tfsdk:"WALDUR_AUTH_SOCIAL_ROLE_CLAIM"`
	WALDURSUPPORTACTIVEBACKENDTYPE                 *string  `json:"WALDUR_SUPPORT_ACTIVE_BACKEND_TYPE,omitempty" tfsdk:"WALDUR_SUPPORT_ACTIVE_BACKEND_TYPE"`
	WALDURSUPPORTDISPLAYREQUESTTYPE                *bool    `json:"WALDUR_SUPPORT_DISPLAY_REQUEST_TYPE,omitempty" tfsdk:"WALDUR_SUPPORT_DISPLAY_REQUEST_TYPE"`
	WALDURSUPPORTENABLED                           *bool    `json:"WALDUR_SUPPORT_ENABLED,omitempty" tfsdk:"WALDUR_SUPPORT_ENABLED"`
	ZAMMADAPIURL                                   *string  `json:"ZAMMAD_API_URL,omitempty" tfsdk:"ZAMMAD_API_URL"`
	ZAMMADARTICLETYPE                              *string  `json:"ZAMMAD_ARTICLE_TYPE,omitempty" tfsdk:"ZAMMAD_ARTICLE_TYPE"`
	ZAMMADCOMMENTCOOLDOWNDURATION                  *int64   `json:"ZAMMAD_COMMENT_COOLDOWN_DURATION,omitempty" tfsdk:"ZAMMAD_COMMENT_COOLDOWN_DURATION"`
	ZAMMADCOMMENTMARKER                            *string  `json:"ZAMMAD_COMMENT_MARKER,omitempty" tfsdk:"ZAMMAD_COMMENT_MARKER"`
	ZAMMADCOMMENTPREFIX                            *string  `json:"ZAMMAD_COMMENT_PREFIX,omitempty" tfsdk:"ZAMMAD_COMMENT_PREFIX"`
	ZAMMADGROUP                                    *string  `json:"ZAMMAD_GROUP,omitempty" tfsdk:"ZAMMAD_GROUP"`
	ZAMMADTOKEN                                    *string  `json:"ZAMMAD_TOKEN,omitempty" tfsdk:"ZAMMAD_TOKEN"`
}

type ContainerFormatEnum struct {
}

type CoreAuthToken struct {
	Token *string `json:"token" tfsdk:"token"`
}

type CoreStates struct {
}

type CorrectiveAction struct {
	ApiEndpoint          *bool    `json:"api_endpoint,omitempty" tfsdk:"api_endpoint"`
	Category             *string  `json:"category" tfsdk:"category"`
	ConfirmationRequired *bool    `json:"confirmation_required,omitempty" tfsdk:"confirmation_required"`
	Label                *string  `json:"label" tfsdk:"label"`
	Method               *string  `json:"method,omitempty" tfsdk:"method"`
	PermissionsRequired  []string `json:"permissions_required,omitempty" tfsdk:"permissions_required"`
	RouteName            *string  `json:"route_name,omitempty" tfsdk:"route_name"`
	Severity             *string  `json:"severity" tfsdk:"severity"`
}

type CostsForPeriod struct {
	EndDate    *string `json:"end_date" tfsdk:"end_date"`
	StartDate  *string `json:"start_date" tfsdk:"start_date"`
	TotalPrice *string `json:"total_price" tfsdk:"total_price"`
}

type CountProjectsOfServiceProviders struct {
	Count                         *int64  `json:"count" tfsdk:"count"`
	CustomerName                  *string `json:"customer_name" tfsdk:"customer_name"`
	CustomerOrganizationGroupName *string `json:"customer_organization_group_name" tfsdk:"customer_organization_group_name"`
	CustomerOrganizationGroupUuid *string `json:"customer_organization_group_uuid" tfsdk:"customer_organization_group_uuid"`
	CustomerUuid                  *string `json:"customer_uuid" tfsdk:"customer_uuid"`
	ServiceProviderUuid           *string `json:"service_provider_uuid" tfsdk:"service_provider_uuid"`
}

type CountProjectsOfServiceProvidersGroupedByOecd struct {
	Count                         *int64  `json:"count" tfsdk:"count"`
	CustomerName                  *string `json:"customer_name" tfsdk:"customer_name"`
	CustomerOrganizationGroupName *string `json:"customer_organization_group_name" tfsdk:"customer_organization_group_name"`
	CustomerOrganizationGroupUuid *string `json:"customer_organization_group_uuid" tfsdk:"customer_organization_group_uuid"`
	CustomerUuid                  *string `json:"customer_uuid" tfsdk:"customer_uuid"`
	OecdFos2007Name               *string `json:"oecd_fos_2007_name" tfsdk:"oecd_fos_2007_name"`
	ServiceProviderUuid           *string `json:"service_provider_uuid" tfsdk:"service_provider_uuid"`
}

type CountStats struct {
	Count *int64  `json:"count" tfsdk:"count"`
	Name  *string `json:"name" tfsdk:"name"`
}

type CountUniqueUsersConnectedWithActiveResourcesOfServiceProvider struct {
	CountUsers   *int64  `json:"count_users" tfsdk:"count_users"`
	CustomerName *string `json:"customer_name" tfsdk:"customer_name"`
	CustomerUuid *string `json:"customer_uuid" tfsdk:"customer_uuid"`
}

type CountUsersOfServiceProviders struct {
	Count                         *int64  `json:"count" tfsdk:"count"`
	CustomerName                  *string `json:"customer_name" tfsdk:"customer_name"`
	CustomerOrganizationGroupName *string `json:"customer_organization_group_name" tfsdk:"customer_organization_group_name"`
	CustomerOrganizationGroupUuid *string `json:"customer_organization_group_uuid" tfsdk:"customer_organization_group_uuid"`
	CustomerUuid                  *string `json:"customer_uuid" tfsdk:"customer_uuid"`
	ServiceProviderUuid           *string `json:"service_provider_uuid" tfsdk:"service_provider_uuid"`
}

type Country struct {
	Label *string `json:"label" tfsdk:"label"`
	Value *string `json:"value" tfsdk:"value"`
}

type CountryEnum struct {
}

type CourseAccount struct {
	Created          *string `json:"created" tfsdk:"created"`
	CustomerName     *string `json:"customer_name" tfsdk:"customer_name"`
	CustomerUuid     *string `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description      *string `json:"description,omitempty" tfsdk:"description"`
	Email            *string `json:"email,omitempty" tfsdk:"email"`
	ErrorMessage     *string `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback   *string `json:"error_traceback" tfsdk:"error_traceback"`
	Modified         *string `json:"modified" tfsdk:"modified"`
	Project          *string `json:"project" tfsdk:"project"`
	ProjectEndDate   *string `json:"project_end_date" tfsdk:"project_end_date"`
	ProjectName      *string `json:"project_name" tfsdk:"project_name"`
	ProjectSlug      *string `json:"project_slug" tfsdk:"project_slug"`
	ProjectStartDate *string `json:"project_start_date" tfsdk:"project_start_date"`
	ProjectUuid      *string `json:"project_uuid" tfsdk:"project_uuid"`
	Url              *string `json:"url" tfsdk:"url"`
	UserUuid         *string `json:"user_uuid" tfsdk:"user_uuid"`
	Username         *string `json:"username" tfsdk:"username"`
}

type CourseAccountCreateNestedRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Email       *string `json:"email,omitempty" tfsdk:"email"`
}

type CourseAccountRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Email       *string `json:"email,omitempty" tfsdk:"email"`
	Project     *string `json:"project" tfsdk:"project"`
}

type CourseAccountsBulkCreateRequest struct {
	CourseAccounts []CourseAccountCreateNestedRequest `json:"course_accounts" tfsdk:"course_accounts"`
	Project        *string                            `json:"project" tfsdk:"project"`
}

type CreateAttachmentsRequest struct {
	Attachments []string `json:"attachments" tfsdk:"attachments"`
}

type CreateCustomerCredit struct {
	AllocatedToProjects       *float64 `json:"allocated_to_projects" tfsdk:"allocated_to_projects"`
	ApplyAsMinimalConsumption *bool    `json:"apply_as_minimal_consumption,omitempty" tfsdk:"apply_as_minimal_consumption"`
	ConsumptionLastMonth      *float64 `json:"consumption_last_month" tfsdk:"consumption_last_month"`
	Customer                  *string  `json:"customer" tfsdk:"customer"`
	CustomerName              *string  `json:"customer_name" tfsdk:"customer_name"`
	CustomerSlug              *string  `json:"customer_slug" tfsdk:"customer_slug"`
	CustomerUuid              *string  `json:"customer_uuid" tfsdk:"customer_uuid"`
	EndDate                   *string  `json:"end_date,omitempty" tfsdk:"end_date"`
	ExpectedConsumption       *string  `json:"expected_consumption,omitempty" tfsdk:"expected_consumption"`
	GraceCoefficient          *string  `json:"grace_coefficient,omitempty" tfsdk:"grace_coefficient"`
	MinimalConsumption        *float64 `json:"minimal_consumption" tfsdk:"minimal_consumption"`
	MinimalConsumptionLogic   *string  `json:"minimal_consumption_logic,omitempty" tfsdk:"minimal_consumption_logic"`
	Offerings                 []string `json:"offerings,omitempty" tfsdk:"offerings"`
	Url                       *string  `json:"url" tfsdk:"url"`
	Value                     *string  `json:"value,omitempty" tfsdk:"value"`
}

type CreateCustomerCreditRequest struct {
	ApplyAsMinimalConsumption *bool    `json:"apply_as_minimal_consumption,omitempty" tfsdk:"apply_as_minimal_consumption"`
	Customer                  *string  `json:"customer" tfsdk:"customer"`
	EndDate                   *string  `json:"end_date,omitempty" tfsdk:"end_date"`
	ExpectedConsumption       *string  `json:"expected_consumption,omitempty" tfsdk:"expected_consumption"`
	GraceCoefficient          *string  `json:"grace_coefficient,omitempty" tfsdk:"grace_coefficient"`
	MinimalConsumptionLogic   *string  `json:"minimal_consumption_logic,omitempty" tfsdk:"minimal_consumption_logic"`
	Offerings                 []string `json:"offerings,omitempty" tfsdk:"offerings"`
	Value                     *string  `json:"value,omitempty" tfsdk:"value"`
}

type CreateFeedback struct {
	Comment    *string `json:"comment,omitempty" tfsdk:"comment"`
	Evaluation *int64  `json:"evaluation" tfsdk:"evaluation"`
	Issue      *string `json:"issue" tfsdk:"issue"`
}

type CreateFeedbackRequest struct {
	Comment    *string `json:"comment,omitempty" tfsdk:"comment"`
	Evaluation *int64  `json:"evaluation" tfsdk:"evaluation"`
	Token      *string `json:"token" tfsdk:"token"`
}

type CreateRouter struct {
	Name            *string `json:"name" tfsdk:"name"`
	Project         *string `json:"project" tfsdk:"project"`
	ServiceSettings *string `json:"service_settings" tfsdk:"service_settings"`
	Tenant          *string `json:"tenant" tfsdk:"tenant"`
	Url             *string `json:"url" tfsdk:"url"`
}

type CreateRouterRequest struct {
	Name   *string `json:"name" tfsdk:"name"`
	Tenant *string `json:"tenant" tfsdk:"tenant"`
}

type Customer struct {
	Abbreviation                 *string             `json:"abbreviation,omitempty" tfsdk:"abbreviation"`
	AccessSubnets                *string             `json:"access_subnets,omitempty" tfsdk:"access_subnets"`
	AccountingStartDate          *string             `json:"accounting_start_date,omitempty" tfsdk:"accounting_start_date"`
	Address                      *string             `json:"address,omitempty" tfsdk:"address"`
	AgreementNumber              *string             `json:"agreement_number,omitempty" tfsdk:"agreement_number"`
	Archived                     *bool               `json:"archived,omitempty" tfsdk:"archived"`
	BackendId                    *string             `json:"backend_id,omitempty" tfsdk:"backend_id"`
	BankAccount                  *string             `json:"bank_account,omitempty" tfsdk:"bank_account"`
	BankName                     *string             `json:"bank_name,omitempty" tfsdk:"bank_name"`
	Blocked                      *bool               `json:"blocked,omitempty" tfsdk:"blocked"`
	CallManagingOrganizationUuid *string             `json:"call_managing_organization_uuid,omitempty" tfsdk:"call_managing_organization_uuid"`
	ContactDetails               *string             `json:"contact_details,omitempty" tfsdk:"contact_details"`
	Country                      *string             `json:"country,omitempty" tfsdk:"country"`
	CountryName                  *string             `json:"country_name,omitempty" tfsdk:"country_name"`
	Created                      *string             `json:"created,omitempty" tfsdk:"created"`
	CustomerCredit               *float64            `json:"customer_credit,omitempty" tfsdk:"customer_credit"`
	CustomerUnallocatedCredit    *float64            `json:"customer_unallocated_credit,omitempty" tfsdk:"customer_unallocated_credit"`
	DefaultTaxPercent            *string             `json:"default_tax_percent,omitempty" tfsdk:"default_tax_percent"`
	Description                  *string             `json:"description,omitempty" tfsdk:"description"`
	DisplayBillingInfoInProjects *bool               `json:"display_billing_info_in_projects,omitempty" tfsdk:"display_billing_info_in_projects"`
	DisplayName                  *string             `json:"display_name,omitempty" tfsdk:"display_name"`
	Domain                       *string             `json:"domain,omitempty" tfsdk:"domain"`
	Email                        *string             `json:"email,omitempty" tfsdk:"email"`
	GracePeriodDays              *int64              `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`
	Homepage                     *string             `json:"homepage,omitempty" tfsdk:"homepage"`
	Image                        *string             `json:"image,omitempty" tfsdk:"image"`
	IsServiceProvider            *bool               `json:"is_service_provider,omitempty" tfsdk:"is_service_provider"`
	Latitude                     *float64            `json:"latitude,omitempty" tfsdk:"latitude"`
	Longitude                    *float64            `json:"longitude,omitempty" tfsdk:"longitude"`
	MaxServiceAccounts           *int64              `json:"max_service_accounts,omitempty" tfsdk:"max_service_accounts"`
	Name                         *string             `json:"name,omitempty" tfsdk:"name"`
	NativeName                   *string             `json:"native_name,omitempty" tfsdk:"native_name"`
	NotificationEmails           *string             `json:"notification_emails,omitempty" tfsdk:"notification_emails"`
	OrganizationGroups           []OrganizationGroup `json:"organization_groups,omitempty" tfsdk:"organization_groups"`
	PaymentProfiles              []PaymentProfile    `json:"payment_profiles,omitempty" tfsdk:"payment_profiles"`
	PhoneNumber                  *string             `json:"phone_number,omitempty" tfsdk:"phone_number"`
	Postal                       *string             `json:"postal,omitempty" tfsdk:"postal"`
	ProjectMetadataChecklist     *string             `json:"project_metadata_checklist,omitempty" tfsdk:"project_metadata_checklist"`
	ProjectsCount                *int64              `json:"projects_count,omitempty" tfsdk:"projects_count"`
	RegistrationCode             *string             `json:"registration_code,omitempty" tfsdk:"registration_code"`
	ServiceProvider              *string             `json:"service_provider,omitempty" tfsdk:"service_provider"`
	ServiceProviderUuid          *string             `json:"service_provider_uuid,omitempty" tfsdk:"service_provider_uuid"`
	Slug                         *string             `json:"slug,omitempty" tfsdk:"slug"`
	SponsorNumber                *int64              `json:"sponsor_number,omitempty" tfsdk:"sponsor_number"`
	Url                          *string             `json:"url,omitempty" tfsdk:"url"`
	UsersCount                   *int64              `json:"users_count,omitempty" tfsdk:"users_count"`
	VatCode                      *string             `json:"vat_code,omitempty" tfsdk:"vat_code"`
}

type CustomerComponentUsagePolicy struct {
	Actions            *string                              `json:"actions" tfsdk:"actions"`
	ComponentLimitsSet []NestedCustomerUsagePolicyComponent `json:"component_limits_set" tfsdk:"component_limits_set"`
	Created            *string                              `json:"created" tfsdk:"created"`
	CreatedByFullName  *string                              `json:"created_by_full_name" tfsdk:"created_by_full_name"`
	CreatedByUsername  *string                              `json:"created_by_username" tfsdk:"created_by_username"`
	FiredDatetime      *string                              `json:"fired_datetime" tfsdk:"fired_datetime"`
	HasFired           *bool                                `json:"has_fired" tfsdk:"has_fired"`
	Scope              *string                              `json:"scope" tfsdk:"scope"`
	ScopeName          *string                              `json:"scope_name" tfsdk:"scope_name"`
	ScopeUuid          *string                              `json:"scope_uuid" tfsdk:"scope_uuid"`
	Url                *string                              `json:"url" tfsdk:"url"`
}

type CustomerComponentUsagePolicyRequest struct {
	Actions            *string                                     `json:"actions" tfsdk:"actions"`
	ComponentLimitsSet []NestedCustomerUsagePolicyComponentRequest `json:"component_limits_set" tfsdk:"component_limits_set"`
	Scope              *string                                     `json:"scope" tfsdk:"scope"`
}

type CustomerCredit struct {
	AllocatedToProjects       *float64                 `json:"allocated_to_projects" tfsdk:"allocated_to_projects"`
	ApplyAsMinimalConsumption *bool                    `json:"apply_as_minimal_consumption,omitempty" tfsdk:"apply_as_minimal_consumption"`
	ConsumptionLastMonth      *float64                 `json:"consumption_last_month" tfsdk:"consumption_last_month"`
	Customer                  *string                  `json:"customer" tfsdk:"customer"`
	CustomerName              *string                  `json:"customer_name" tfsdk:"customer_name"`
	CustomerSlug              *string                  `json:"customer_slug" tfsdk:"customer_slug"`
	CustomerUuid              *string                  `json:"customer_uuid" tfsdk:"customer_uuid"`
	EndDate                   *string                  `json:"end_date,omitempty" tfsdk:"end_date"`
	ExpectedConsumption       *string                  `json:"expected_consumption,omitempty" tfsdk:"expected_consumption"`
	GraceCoefficient          *string                  `json:"grace_coefficient,omitempty" tfsdk:"grace_coefficient"`
	MinimalConsumption        *float64                 `json:"minimal_consumption" tfsdk:"minimal_consumption"`
	MinimalConsumptionLogic   *string                  `json:"minimal_consumption_logic,omitempty" tfsdk:"minimal_consumption_logic"`
	Offerings                 []NestedProviderOffering `json:"offerings" tfsdk:"offerings"`
	Url                       *string                  `json:"url" tfsdk:"url"`
	Value                     *string                  `json:"value,omitempty" tfsdk:"value"`
}

type CustomerCreditConsumption struct {
	Date  *string `json:"date" tfsdk:"date"`
	Price *string `json:"price" tfsdk:"price"`
}

type CustomerCreditRequest struct {
	ApplyAsMinimalConsumption *bool   `json:"apply_as_minimal_consumption,omitempty" tfsdk:"apply_as_minimal_consumption"`
	Customer                  *string `json:"customer" tfsdk:"customer"`
	EndDate                   *string `json:"end_date,omitempty" tfsdk:"end_date"`
	ExpectedConsumption       *string `json:"expected_consumption,omitempty" tfsdk:"expected_consumption"`
	GraceCoefficient          *string `json:"grace_coefficient,omitempty" tfsdk:"grace_coefficient"`
	MinimalConsumptionLogic   *string `json:"minimal_consumption_logic,omitempty" tfsdk:"minimal_consumption_logic"`
	Value                     *string `json:"value,omitempty" tfsdk:"value"`
}

type CustomerDetails struct {
	Address     *string `json:"address,omitempty" tfsdk:"address"`
	BankAccount *string `json:"bank_account,omitempty" tfsdk:"bank_account"`
	BankName    *string `json:"bank_name,omitempty" tfsdk:"bank_name"`
	Country     *string `json:"country,omitempty" tfsdk:"country"`
	CountryName *string `json:"country_name,omitempty" tfsdk:"country_name"`
	Email       *string `json:"email,omitempty" tfsdk:"email"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
	PhoneNumber *string `json:"phone_number,omitempty" tfsdk:"phone_number"`
	Postal      *string `json:"postal,omitempty" tfsdk:"postal"`
	VatCode     *string `json:"vat_code,omitempty" tfsdk:"vat_code"`
}

type CustomerEstimatedCostPolicy struct {
	Actions           *string `json:"actions" tfsdk:"actions"`
	Created           *string `json:"created" tfsdk:"created"`
	CreatedByFullName *string `json:"created_by_full_name" tfsdk:"created_by_full_name"`
	CreatedByUsername *string `json:"created_by_username" tfsdk:"created_by_username"`
	CustomerCredit    *int64  `json:"customer_credit" tfsdk:"customer_credit"`
	FiredDatetime     *string `json:"fired_datetime" tfsdk:"fired_datetime"`
	HasFired          *bool   `json:"has_fired" tfsdk:"has_fired"`
	LimitCost         *int64  `json:"limit_cost" tfsdk:"limit_cost"`
	Period            *int64  `json:"period,omitempty" tfsdk:"period"`
	PeriodName        *string `json:"period_name" tfsdk:"period_name"`
	Scope             *string `json:"scope" tfsdk:"scope"`
	ScopeName         *string `json:"scope_name" tfsdk:"scope_name"`
	ScopeUuid         *string `json:"scope_uuid" tfsdk:"scope_uuid"`
	Url               *string `json:"url" tfsdk:"url"`
}

type CustomerEstimatedCostPolicyRequest struct {
	Actions   *string `json:"actions" tfsdk:"actions"`
	LimitCost *int64  `json:"limit_cost" tfsdk:"limit_cost"`
	Period    *int64  `json:"period,omitempty" tfsdk:"period"`
	Scope     *string `json:"scope" tfsdk:"scope"`
}

type CustomerIndustryFlagStats struct {
	Abbreviation *string `json:"abbreviation" tfsdk:"abbreviation"`
	Count        *int64  `json:"count" tfsdk:"count"`
	IsIndustry   *string `json:"is_industry" tfsdk:"is_industry"`
	Name         *string `json:"name" tfsdk:"name"`
}

type CustomerMemberCount struct {
	Abbreviation *string `json:"abbreviation" tfsdk:"abbreviation"`
	Count        *int64  `json:"count" tfsdk:"count"`
	HasResources *bool   `json:"has_resources" tfsdk:"has_resources"`
	Name         *string `json:"name" tfsdk:"name"`
}

type CustomerOecdCodeStats struct {
	Abbreviation *string `json:"abbreviation" tfsdk:"abbreviation"`
	Count        *int64  `json:"count" tfsdk:"count"`
	Name         *string `json:"name" tfsdk:"name"`
	Oecd         *string `json:"oecd" tfsdk:"oecd"`
}

type CustomerPermissionReview struct {
	Closed           *string `json:"closed" tfsdk:"closed"`
	Created          *string `json:"created" tfsdk:"created"`
	CustomerName     *string `json:"customer_name" tfsdk:"customer_name"`
	CustomerUuid     *string `json:"customer_uuid" tfsdk:"customer_uuid"`
	IsPending        *bool   `json:"is_pending" tfsdk:"is_pending"`
	ReviewerFullName *string `json:"reviewer_full_name" tfsdk:"reviewer_full_name"`
	ReviewerUuid     *string `json:"reviewer_uuid" tfsdk:"reviewer_uuid"`
	Url              *string `json:"url" tfsdk:"url"`
}

type CustomerQuotas struct {
	CustomerAbbreviation *string `json:"customer_abbreviation" tfsdk:"customer_abbreviation"`
	CustomerName         *string `json:"customer_name" tfsdk:"customer_name"`
	Value                *int64  `json:"value" tfsdk:"value"`
}

type CustomerRequest struct {
	Abbreviation                 *string  `json:"abbreviation,omitempty" tfsdk:"abbreviation"`
	AccessSubnets                *string  `json:"access_subnets,omitempty" tfsdk:"access_subnets"`
	AccountingStartDate          *string  `json:"accounting_start_date,omitempty" tfsdk:"accounting_start_date"`
	Address                      *string  `json:"address,omitempty" tfsdk:"address"`
	AgreementNumber              *string  `json:"agreement_number,omitempty" tfsdk:"agreement_number"`
	Archived                     *bool    `json:"archived,omitempty" tfsdk:"archived"`
	BackendId                    *string  `json:"backend_id,omitempty" tfsdk:"backend_id"`
	BankAccount                  *string  `json:"bank_account,omitempty" tfsdk:"bank_account"`
	BankName                     *string  `json:"bank_name,omitempty" tfsdk:"bank_name"`
	Blocked                      *bool    `json:"blocked,omitempty" tfsdk:"blocked"`
	ContactDetails               *string  `json:"contact_details,omitempty" tfsdk:"contact_details"`
	Country                      *string  `json:"country,omitempty" tfsdk:"country"`
	DefaultTaxPercent            *string  `json:"default_tax_percent,omitempty" tfsdk:"default_tax_percent"`
	Description                  *string  `json:"description,omitempty" tfsdk:"description"`
	DisplayBillingInfoInProjects *bool    `json:"display_billing_info_in_projects,omitempty" tfsdk:"display_billing_info_in_projects"`
	Domain                       *string  `json:"domain,omitempty" tfsdk:"domain"`
	Email                        *string  `json:"email,omitempty" tfsdk:"email"`
	GracePeriodDays              *int64   `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`
	Homepage                     *string  `json:"homepage,omitempty" tfsdk:"homepage"`
	Image                        *string  `json:"image,omitempty" tfsdk:"image"`
	Latitude                     *float64 `json:"latitude,omitempty" tfsdk:"latitude"`
	Longitude                    *float64 `json:"longitude,omitempty" tfsdk:"longitude"`
	MaxServiceAccounts           *int64   `json:"max_service_accounts,omitempty" tfsdk:"max_service_accounts"`
	Name                         *string  `json:"name" tfsdk:"name"`
	NativeName                   *string  `json:"native_name,omitempty" tfsdk:"native_name"`
	NotificationEmails           *string  `json:"notification_emails,omitempty" tfsdk:"notification_emails"`
	PhoneNumber                  *string  `json:"phone_number,omitempty" tfsdk:"phone_number"`
	Postal                       *string  `json:"postal,omitempty" tfsdk:"postal"`
	ProjectMetadataChecklist     *string  `json:"project_metadata_checklist,omitempty" tfsdk:"project_metadata_checklist"`
	RegistrationCode             *string  `json:"registration_code,omitempty" tfsdk:"registration_code"`
	Slug                         *string  `json:"slug,omitempty" tfsdk:"slug"`
	SponsorNumber                *int64   `json:"sponsor_number,omitempty" tfsdk:"sponsor_number"`
	VatCode                      *string  `json:"vat_code,omitempty" tfsdk:"vat_code"`
}

type CustomerRequestForm struct {
	Abbreviation                 *string  `json:"abbreviation,omitempty" tfsdk:"abbreviation"`
	AccessSubnets                *string  `json:"access_subnets,omitempty" tfsdk:"access_subnets"`
	AccountingStartDate          *string  `json:"accounting_start_date,omitempty" tfsdk:"accounting_start_date"`
	Address                      *string  `json:"address,omitempty" tfsdk:"address"`
	AgreementNumber              *string  `json:"agreement_number,omitempty" tfsdk:"agreement_number"`
	Archived                     *bool    `json:"archived,omitempty" tfsdk:"archived"`
	BackendId                    *string  `json:"backend_id,omitempty" tfsdk:"backend_id"`
	BankAccount                  *string  `json:"bank_account,omitempty" tfsdk:"bank_account"`
	BankName                     *string  `json:"bank_name,omitempty" tfsdk:"bank_name"`
	Blocked                      *bool    `json:"blocked,omitempty" tfsdk:"blocked"`
	ContactDetails               *string  `json:"contact_details,omitempty" tfsdk:"contact_details"`
	Country                      *string  `json:"country,omitempty" tfsdk:"country"`
	DefaultTaxPercent            *string  `json:"default_tax_percent,omitempty" tfsdk:"default_tax_percent"`
	Description                  *string  `json:"description,omitempty" tfsdk:"description"`
	DisplayBillingInfoInProjects *bool    `json:"display_billing_info_in_projects,omitempty" tfsdk:"display_billing_info_in_projects"`
	Domain                       *string  `json:"domain,omitempty" tfsdk:"domain"`
	Email                        *string  `json:"email,omitempty" tfsdk:"email"`
	GracePeriodDays              *int64   `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`
	Homepage                     *string  `json:"homepage,omitempty" tfsdk:"homepage"`
	Image                        *string  `json:"image,omitempty" tfsdk:"image"`
	Latitude                     *float64 `json:"latitude,omitempty" tfsdk:"latitude"`
	Longitude                    *float64 `json:"longitude,omitempty" tfsdk:"longitude"`
	MaxServiceAccounts           *int64   `json:"max_service_accounts,omitempty" tfsdk:"max_service_accounts"`
	Name                         *string  `json:"name" tfsdk:"name"`
	NativeName                   *string  `json:"native_name,omitempty" tfsdk:"native_name"`
	NotificationEmails           *string  `json:"notification_emails,omitempty" tfsdk:"notification_emails"`
	PhoneNumber                  *string  `json:"phone_number,omitempty" tfsdk:"phone_number"`
	Postal                       *string  `json:"postal,omitempty" tfsdk:"postal"`
	ProjectMetadataChecklist     *string  `json:"project_metadata_checklist,omitempty" tfsdk:"project_metadata_checklist"`
	RegistrationCode             *string  `json:"registration_code,omitempty" tfsdk:"registration_code"`
	Slug                         *string  `json:"slug,omitempty" tfsdk:"slug"`
	SponsorNumber                *int64   `json:"sponsor_number,omitempty" tfsdk:"sponsor_number"`
	VatCode                      *string  `json:"vat_code,omitempty" tfsdk:"vat_code"`
}

type CustomerRequestMultipart struct {
	Abbreviation                 *string  `json:"abbreviation,omitempty" tfsdk:"abbreviation"`
	AccessSubnets                *string  `json:"access_subnets,omitempty" tfsdk:"access_subnets"`
	AccountingStartDate          *string  `json:"accounting_start_date,omitempty" tfsdk:"accounting_start_date"`
	Address                      *string  `json:"address,omitempty" tfsdk:"address"`
	AgreementNumber              *string  `json:"agreement_number,omitempty" tfsdk:"agreement_number"`
	Archived                     *bool    `json:"archived,omitempty" tfsdk:"archived"`
	BackendId                    *string  `json:"backend_id,omitempty" tfsdk:"backend_id"`
	BankAccount                  *string  `json:"bank_account,omitempty" tfsdk:"bank_account"`
	BankName                     *string  `json:"bank_name,omitempty" tfsdk:"bank_name"`
	Blocked                      *bool    `json:"blocked,omitempty" tfsdk:"blocked"`
	ContactDetails               *string  `json:"contact_details,omitempty" tfsdk:"contact_details"`
	Country                      *string  `json:"country,omitempty" tfsdk:"country"`
	DefaultTaxPercent            *string  `json:"default_tax_percent,omitempty" tfsdk:"default_tax_percent"`
	Description                  *string  `json:"description,omitempty" tfsdk:"description"`
	DisplayBillingInfoInProjects *bool    `json:"display_billing_info_in_projects,omitempty" tfsdk:"display_billing_info_in_projects"`
	Domain                       *string  `json:"domain,omitempty" tfsdk:"domain"`
	Email                        *string  `json:"email,omitempty" tfsdk:"email"`
	GracePeriodDays              *int64   `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`
	Homepage                     *string  `json:"homepage,omitempty" tfsdk:"homepage"`
	Image                        *string  `json:"image,omitempty" tfsdk:"image"`
	Latitude                     *float64 `json:"latitude,omitempty" tfsdk:"latitude"`
	Longitude                    *float64 `json:"longitude,omitempty" tfsdk:"longitude"`
	MaxServiceAccounts           *int64   `json:"max_service_accounts,omitempty" tfsdk:"max_service_accounts"`
	Name                         *string  `json:"name" tfsdk:"name"`
	NativeName                   *string  `json:"native_name,omitempty" tfsdk:"native_name"`
	NotificationEmails           *string  `json:"notification_emails,omitempty" tfsdk:"notification_emails"`
	PhoneNumber                  *string  `json:"phone_number,omitempty" tfsdk:"phone_number"`
	Postal                       *string  `json:"postal,omitempty" tfsdk:"postal"`
	ProjectMetadataChecklist     *string  `json:"project_metadata_checklist,omitempty" tfsdk:"project_metadata_checklist"`
	RegistrationCode             *string  `json:"registration_code,omitempty" tfsdk:"registration_code"`
	Slug                         *string  `json:"slug,omitempty" tfsdk:"slug"`
	SponsorNumber                *int64   `json:"sponsor_number,omitempty" tfsdk:"sponsor_number"`
	VatCode                      *string  `json:"vat_code,omitempty" tfsdk:"vat_code"`
}

type CustomerServiceAccount struct {
	Created             *string `json:"created" tfsdk:"created"`
	Customer            *string `json:"customer" tfsdk:"customer"`
	CustomerName        *string `json:"customer_name" tfsdk:"customer_name"`
	CustomerUuid        *string `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description         *string `json:"description,omitempty" tfsdk:"description"`
	Email               *string `json:"email,omitempty" tfsdk:"email"`
	ErrorMessage        *string `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback      *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	ExpiresAt           *string `json:"expires_at" tfsdk:"expires_at"`
	Modified            *string `json:"modified" tfsdk:"modified"`
	PreferredIdentifier *string `json:"preferred_identifier,omitempty" tfsdk:"preferred_identifier"`
	Token               *string `json:"token" tfsdk:"token"`
	Url                 *string `json:"url" tfsdk:"url"`
	Username            *string `json:"username,omitempty" tfsdk:"username"`
}

type CustomerServiceAccountRequest struct {
	Customer            *string `json:"customer" tfsdk:"customer"`
	Description         *string `json:"description,omitempty" tfsdk:"description"`
	Email               *string `json:"email,omitempty" tfsdk:"email"`
	ErrorTraceback      *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	PreferredIdentifier *string `json:"preferred_identifier,omitempty" tfsdk:"preferred_identifier"`
	Username            *string `json:"username,omitempty" tfsdk:"username"`
}

type CustomerUser struct {
	Email          *string                   `json:"email,omitempty" tfsdk:"email"`
	ExpirationTime *string                   `json:"expiration_time,omitempty" tfsdk:"expiration_time"`
	FullName       *string                   `json:"full_name,omitempty" tfsdk:"full_name"`
	Image          *string                   `json:"image,omitempty" tfsdk:"image"`
	Projects       []NestedProjectPermission `json:"projects,omitempty" tfsdk:"projects"`
	RoleName       *string                   `json:"role_name,omitempty" tfsdk:"role_name"`
	Url            *string                   `json:"url,omitempty" tfsdk:"url"`
	Username       *string                   `json:"username,omitempty" tfsdk:"username"`
}

type DataVolume struct {
	Filesystem *string `json:"filesystem,omitempty" tfsdk:"filesystem"`
	MountPoint *string `json:"mount_point" tfsdk:"mount_point"`
	Size       *int64  `json:"size" tfsdk:"size"`
	VolumeType *string `json:"volume_type,omitempty" tfsdk:"volume_type"`
}

type DataVolumeRequest struct {
	Filesystem *string `json:"filesystem,omitempty" tfsdk:"filesystem"`
	MountPoint *string `json:"mount_point" tfsdk:"mount_point"`
	Size       *int64  `json:"size" tfsdk:"size"`
	VolumeType *string `json:"volume_type,omitempty" tfsdk:"volume_type"`
}

type DecidingEntityEnum struct {
}

type DeleteAttachmentsRequest struct {
	AttachmentIds []string `json:"attachment_ids" tfsdk:"attachment_ids"`
}

type DemoPreset struct {
	Description *string  `json:"description" tfsdk:"description"`
	Name        *string  `json:"name" tfsdk:"name"`
	Scenarios   []string `json:"scenarios" tfsdk:"scenarios"`
	Title       *string  `json:"title" tfsdk:"title"`
	Version     *string  `json:"version" tfsdk:"version"`
}

type DemoPresetLoadRequestRequest struct {
	CleanupFirst *bool `json:"cleanup_first,omitempty" tfsdk:"cleanup_first"`
	DryRun       *bool `json:"dry_run,omitempty" tfsdk:"dry_run"`
	SkipRoles    *bool `json:"skip_roles,omitempty" tfsdk:"skip_roles"`
	SkipUsers    *bool `json:"skip_users,omitempty" tfsdk:"skip_users"`
}

type DemoPresetLoadResponse struct {
	Message *string          `json:"message" tfsdk:"message"`
	Output  *string          `json:"output,omitempty" tfsdk:"output"`
	Success *bool            `json:"success" tfsdk:"success"`
	Users   []DemoPresetUser `json:"users,omitempty" tfsdk:"users"`
}

type DemoPresetUser struct {
	Email     *string `json:"email,omitempty" tfsdk:"email"`
	IsStaff   *bool   `json:"is_staff,omitempty" tfsdk:"is_staff"`
	IsSupport *bool   `json:"is_support,omitempty" tfsdk:"is_support"`
	Password  *string `json:"password" tfsdk:"password"`
	Username  *string `json:"username" tfsdk:"username"`
}

type DependencyLogicOperatorEnum struct {
}

type DeploymentModeEnum struct {
}

type DeprecatedNetworkRBACPolicy struct {
	BackendId        *string `json:"backend_id" tfsdk:"backend_id"`
	Created          *string `json:"created" tfsdk:"created"`
	Network          *string `json:"network" tfsdk:"network"`
	NetworkName      *string `json:"network_name" tfsdk:"network_name"`
	PolicyType       *string `json:"policy_type,omitempty" tfsdk:"policy_type"`
	TargetTenant     *string `json:"target_tenant" tfsdk:"target_tenant"`
	TargetTenantName *string `json:"target_tenant_name" tfsdk:"target_tenant_name"`
	Url              *string `json:"url" tfsdk:"url"`
}

type DeprecatedNetworkRBACPolicyRequest struct {
	PolicyType   *string `json:"policy_type,omitempty" tfsdk:"policy_type"`
	TargetTenant *string `json:"target_tenant" tfsdk:"target_tenant"`
}

type DetailState struct {
	Detail *string `json:"detail" tfsdk:"detail"`
	State  *string `json:"state" tfsdk:"state"`
}

type DigitalOceanDroplet struct {
	AccessUrl                   *string  `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string  `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Cores                       *int64   `json:"cores,omitempty" tfsdk:"cores"`
	Created                     *string  `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string  `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string  `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string  `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string  `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string  `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string  `json:"description,omitempty" tfsdk:"description"`
	Disk                        *int64   `json:"disk,omitempty" tfsdk:"disk"`
	ErrorMessage                *string  `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string  `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	ExternalIps                 []string `json:"external_ips,omitempty" tfsdk:"external_ips"`
	ImageName                   *string  `json:"image_name,omitempty" tfsdk:"image_name"`
	InternalIps                 []string `json:"internal_ips,omitempty" tfsdk:"internal_ips"`
	IsLimitBased                *bool    `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool    `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	KeyFingerprint              *string  `json:"key_fingerprint,omitempty" tfsdk:"key_fingerprint"`
	KeyName                     *string  `json:"key_name,omitempty" tfsdk:"key_name"`
	Latitude                    *float64 `json:"latitude,omitempty" tfsdk:"latitude"`
	Longitude                   *float64 `json:"longitude,omitempty" tfsdk:"longitude"`
	MarketplaceCategoryName     *string  `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string  `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string  `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string  `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string  `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string  `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string  `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	MinDisk                     *int64   `json:"min_disk,omitempty" tfsdk:"min_disk"`
	MinRam                      *int64   `json:"min_ram,omitempty" tfsdk:"min_ram"`
	Modified                    *string  `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string  `json:"name,omitempty" tfsdk:"name"`
	Project                     *string  `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string  `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string  `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	Ram                         *int64   `json:"ram,omitempty" tfsdk:"ram"`
	RegionName                  *string  `json:"region_name,omitempty" tfsdk:"region_name"`
	ResourceType                *string  `json:"resource_type,omitempty" tfsdk:"resource_type"`
	RuntimeState                *string  `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceName                 *string  `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string  `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string  `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string  `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string  `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	StartTime                   *string  `json:"start_time,omitempty" tfsdk:"start_time"`
	State                       *string  `json:"state,omitempty" tfsdk:"state"`
	Url                         *string  `json:"url,omitempty" tfsdk:"url"`
	UserData                    *string  `json:"user_data,omitempty" tfsdk:"user_data"`
}

type DigitalOceanDropletRequest struct {
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	Image           *string `json:"image" tfsdk:"image"`
	Name            *string `json:"name" tfsdk:"name"`
	Project         *string `json:"project" tfsdk:"project"`
	Region          *string `json:"region" tfsdk:"region"`
	ServiceSettings *string `json:"service_settings" tfsdk:"service_settings"`
	Size            *string `json:"size" tfsdk:"size"`
	SshPublicKey    *string `json:"ssh_public_key,omitempty" tfsdk:"ssh_public_key"`
	UserData        *string `json:"user_data,omitempty" tfsdk:"user_data"`
}

type DigitalOceanDropletResize struct {
	Disk *bool `json:"disk" tfsdk:"disk"`
}

type DigitalOceanDropletResizeRequest struct {
	Disk *bool   `json:"disk" tfsdk:"disk"`
	Size *string `json:"size" tfsdk:"size"`
}

type DigitalOceanImage struct {
	CreatedAt    *string              `json:"created_at,omitempty" tfsdk:"created_at"`
	Distribution *string              `json:"distribution" tfsdk:"distribution"`
	IsOfficial   *bool                `json:"is_official,omitempty" tfsdk:"is_official"`
	MinDiskSize  *int64               `json:"min_disk_size,omitempty" tfsdk:"min_disk_size"`
	Name         *string              `json:"name" tfsdk:"name"`
	Regions      []DigitalOceanRegion `json:"regions" tfsdk:"regions"`
	Type         *string              `json:"type" tfsdk:"type"`
	Url          *string              `json:"url" tfsdk:"url"`
}

type DigitalOceanRegion struct {
	Name *string `json:"name" tfsdk:"name"`
	Url  *string `json:"url" tfsdk:"url"`
}

type DigitalOceanSize struct {
	Cores    *int64               `json:"cores" tfsdk:"cores"`
	Disk     *int64               `json:"disk" tfsdk:"disk"`
	Name     *string              `json:"name" tfsdk:"name"`
	Ram      *int64               `json:"ram" tfsdk:"ram"`
	Regions  []DigitalOceanRegion `json:"regions" tfsdk:"regions"`
	Transfer *int64               `json:"transfer" tfsdk:"transfer"`
	Url      *string              `json:"url" tfsdk:"url"`
}

type DirectionEnum struct {
}

type DiscountConfigRequest struct {
	DiscountRate      *int64 `json:"discount_rate,omitempty" tfsdk:"discount_rate"`
	DiscountThreshold *int64 `json:"discount_threshold,omitempty" tfsdk:"discount_threshold"`
}

type DiscountTypeEnum struct {
}

type DiscountsUpdateRequest struct {
}

type DiscoverCustomFieldsRequestRequest struct {
	ApiUrl              *string `json:"api_url" tfsdk:"api_url"`
	AuthMethod          *string `json:"auth_method" tfsdk:"auth_method"`
	Email               *string `json:"email,omitempty" tfsdk:"email"`
	Password            *string `json:"password,omitempty" tfsdk:"password"`
	PersonalAccessToken *string `json:"personal_access_token,omitempty" tfsdk:"personal_access_token"`
	ProjectId           *string `json:"project_id,omitempty" tfsdk:"project_id"`
	RequestTypeId       *string `json:"request_type_id,omitempty" tfsdk:"request_type_id"`
	Token               *string `json:"token,omitempty" tfsdk:"token"`
	Username            *string `json:"username,omitempty" tfsdk:"username"`
	VerifySsl           *bool   `json:"verify_ssl,omitempty" tfsdk:"verify_ssl"`
}

type DiscoverPrioritiesRequestRequest struct {
	ApiUrl              *string `json:"api_url" tfsdk:"api_url"`
	AuthMethod          *string `json:"auth_method" tfsdk:"auth_method"`
	Email               *string `json:"email,omitempty" tfsdk:"email"`
	Password            *string `json:"password,omitempty" tfsdk:"password"`
	PersonalAccessToken *string `json:"personal_access_token,omitempty" tfsdk:"personal_access_token"`
	Token               *string `json:"token,omitempty" tfsdk:"token"`
	Username            *string `json:"username,omitempty" tfsdk:"username"`
	VerifySsl           *bool   `json:"verify_ssl,omitempty" tfsdk:"verify_ssl"`
}

type DiscoverProjectsRequestRequest struct {
	ApiUrl              *string `json:"api_url" tfsdk:"api_url"`
	AuthMethod          *string `json:"auth_method" tfsdk:"auth_method"`
	Email               *string `json:"email,omitempty" tfsdk:"email"`
	Password            *string `json:"password,omitempty" tfsdk:"password"`
	PersonalAccessToken *string `json:"personal_access_token,omitempty" tfsdk:"personal_access_token"`
	Token               *string `json:"token,omitempty" tfsdk:"token"`
	Username            *string `json:"username,omitempty" tfsdk:"username"`
	VerifySsl           *bool   `json:"verify_ssl,omitempty" tfsdk:"verify_ssl"`
}

type DiscoverRequestTypesRequestRequest struct {
	ApiUrl              *string `json:"api_url" tfsdk:"api_url"`
	AuthMethod          *string `json:"auth_method" tfsdk:"auth_method"`
	Email               *string `json:"email,omitempty" tfsdk:"email"`
	Password            *string `json:"password,omitempty" tfsdk:"password"`
	PersonalAccessToken *string `json:"personal_access_token,omitempty" tfsdk:"personal_access_token"`
	ProjectId           *string `json:"project_id" tfsdk:"project_id"`
	Token               *string `json:"token,omitempty" tfsdk:"token"`
	Username            *string `json:"username,omitempty" tfsdk:"username"`
	VerifySsl           *bool   `json:"verify_ssl,omitempty" tfsdk:"verify_ssl"`
}

type DiskFormatEnum struct {
}

type DryRun struct {
	Created         *string `json:"created" tfsdk:"created"`
	GetStateDisplay *string `json:"get_state_display" tfsdk:"get_state_display"`
	OrderOffering   *string `json:"order_offering,omitempty" tfsdk:"order_offering"`
	OrderType       *string `json:"order_type" tfsdk:"order_type"`
	Output          *string `json:"output" tfsdk:"output"`
	State           *int64  `json:"state" tfsdk:"state"`
	Url             *string `json:"url" tfsdk:"url"`
}

type DryRunRequest struct {
	OrderOffering *string `json:"order_offering,omitempty" tfsdk:"order_offering"`
	Plan          *string `json:"plan,omitempty" tfsdk:"plan"`
	Type          *string `json:"type,omitempty" tfsdk:"type"`
}

type DryRunStateEnum struct {
}

type DryRunTypeEnum struct {
}

type EmailHook struct {
	AuthorEmail    *string  `json:"author_email" tfsdk:"author_email"`
	AuthorFullname *string  `json:"author_fullname" tfsdk:"author_fullname"`
	AuthorUsername *string  `json:"author_username" tfsdk:"author_username"`
	AuthorUuid     *string  `json:"author_uuid" tfsdk:"author_uuid"`
	Created        *string  `json:"created" tfsdk:"created"`
	Email          *string  `json:"email" tfsdk:"email"`
	EventGroups    []string `json:"event_groups,omitempty" tfsdk:"event_groups"`
	EventTypes     []string `json:"event_types,omitempty" tfsdk:"event_types"`
	HookType       *string  `json:"hook_type" tfsdk:"hook_type"`
	IsActive       *bool    `json:"is_active,omitempty" tfsdk:"is_active"`
	Modified       *string  `json:"modified" tfsdk:"modified"`
	Url            *string  `json:"url" tfsdk:"url"`
}

type EmailHookRequest struct {
	Email       *string  `json:"email" tfsdk:"email"`
	EventGroups []string `json:"event_groups,omitempty" tfsdk:"event_groups"`
	EventTypes  []string `json:"event_types,omitempty" tfsdk:"event_types"`
	IsActive    *bool    `json:"is_active,omitempty" tfsdk:"is_active"`
}

type EmailLog struct {
	Body    *string  `json:"body" tfsdk:"body"`
	Emails  []string `json:"emails" tfsdk:"emails"`
	SentAt  *string  `json:"sent_at" tfsdk:"sent_at"`
	Subject *string  `json:"subject" tfsdk:"subject"`
	Url     *string  `json:"url" tfsdk:"url"`
}

type EndpointUUID struct {
}

type EndpointUUIDRequest struct {
}

type EthertypeEnum struct {
}

type Event struct {
	Created   *string `json:"created,omitempty" tfsdk:"created"`
	EventType *string `json:"event_type,omitempty" tfsdk:"event_type"`
	Message   *string `json:"message,omitempty" tfsdk:"message"`
}

type EventGroupsEnum struct {
}

type EventMetadataResponse struct {
}

type EventStats struct {
	Count *int64 `json:"count" tfsdk:"count"`
	Month *int64 `json:"month" tfsdk:"month"`
	Year  *int64 `json:"year" tfsdk:"year"`
}

type EventSubscription struct {
	Created      *string `json:"created" tfsdk:"created"`
	Description  *string `json:"description,omitempty" tfsdk:"description"`
	Modified     *string `json:"modified" tfsdk:"modified"`
	SourceIp     *string `json:"source_ip" tfsdk:"source_ip"`
	Url          *string `json:"url" tfsdk:"url"`
	User         *string `json:"user" tfsdk:"user"`
	UserFullName *string `json:"user_full_name" tfsdk:"user_full_name"`
	UserUsername *string `json:"user_username" tfsdk:"user_username"`
	UserUuid     *string `json:"user_uuid" tfsdk:"user_uuid"`
}

type EventSubscriptionRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
}

type EventTypesEnum struct {
}

type ExecuteActionErrorResponse struct {
	Error *string `json:"error" tfsdk:"error"`
}

type ExecuteActionRequest struct {
	ActionLabel *string `json:"action_label" tfsdk:"action_label"`
}

type ExecuteActionResponse struct {
	Action      *string `json:"action" tfsdk:"action"`
	Message     *string `json:"message,omitempty" tfsdk:"message"`
	RedirectUrl *string `json:"redirect_url,omitempty" tfsdk:"redirect_url"`
}

type ExecutionStateEnum struct {
}

type ExportComponentData struct {
	ArticleCode  *string  `json:"article_code" tfsdk:"article_code"`
	BackendId    *string  `json:"backend_id" tfsdk:"backend_id"`
	BillingType  *string  `json:"billing_type" tfsdk:"billing_type"`
	Description  *string  `json:"description" tfsdk:"description"`
	LimitAmount  *int64   `json:"limit_amount" tfsdk:"limit_amount"`
	LimitPeriod  *string  `json:"limit_period" tfsdk:"limit_period"`
	MeasuredUnit *string  `json:"measured_unit" tfsdk:"measured_unit"`
	Name         *string  `json:"name" tfsdk:"name"`
	Type         *string  `json:"type" tfsdk:"type"`
	UnitFactor   *float64 `json:"unit_factor" tfsdk:"unit_factor"`
}

type ExportComponentDataRequest struct {
	ArticleCode  *string  `json:"article_code" tfsdk:"article_code"`
	BackendId    *string  `json:"backend_id" tfsdk:"backend_id"`
	BillingType  *string  `json:"billing_type" tfsdk:"billing_type"`
	Description  *string  `json:"description" tfsdk:"description"`
	LimitAmount  *int64   `json:"limit_amount" tfsdk:"limit_amount"`
	LimitPeriod  *string  `json:"limit_period" tfsdk:"limit_period"`
	MeasuredUnit *string  `json:"measured_unit" tfsdk:"measured_unit"`
	Name         *string  `json:"name" tfsdk:"name"`
	Type         *string  `json:"type" tfsdk:"type"`
	UnitFactor   *float64 `json:"unit_factor" tfsdk:"unit_factor"`
}

type ExportEndpointData struct {
	Name *string `json:"name" tfsdk:"name"`
	Url  *string `json:"url" tfsdk:"url"`
}

type ExportEndpointDataRequest struct {
	Name *string `json:"name" tfsdk:"name"`
	Url  *string `json:"url" tfsdk:"url"`
}

type ExportFileData struct {
	ContentType *string `json:"content_type" tfsdk:"content_type"`
	FileContent *string `json:"file_content" tfsdk:"file_content"`
	Filename    *string `json:"filename" tfsdk:"filename"`
	Name        *string `json:"name" tfsdk:"name"`
}

type ExportFileDataRequest struct {
	ContentType *string `json:"content_type" tfsdk:"content_type"`
	FileContent *string `json:"file_content" tfsdk:"file_content"`
	Filename    *string `json:"filename" tfsdk:"filename"`
	Name        *string `json:"name" tfsdk:"name"`
}

type ExportOfferingData struct {
	AccessUrl        *string  `json:"access_url" tfsdk:"access_url"`
	Billable         *bool    `json:"billable" tfsdk:"billable"`
	CategoryName     *string  `json:"category_name" tfsdk:"category_name"`
	Country          *string  `json:"country" tfsdk:"country"`
	Description      *string  `json:"description" tfsdk:"description"`
	FullDescription  *string  `json:"full_description" tfsdk:"full_description"`
	GettingStarted   *string  `json:"getting_started" tfsdk:"getting_started"`
	IntegrationGuide *string  `json:"integration_guide" tfsdk:"integration_guide"`
	Latitude         *float64 `json:"latitude" tfsdk:"latitude"`
	Longitude        *float64 `json:"longitude" tfsdk:"longitude"`
	Name             *string  `json:"name" tfsdk:"name"`
	PausedReason     *string  `json:"paused_reason" tfsdk:"paused_reason"`
	Shared           *bool    `json:"shared" tfsdk:"shared"`
	State            *string  `json:"state" tfsdk:"state"`
	Type             *string  `json:"type" tfsdk:"type"`
	VendorDetails    *string  `json:"vendor_details" tfsdk:"vendor_details"`
}

type ExportOfferingDataRequest struct {
	AccessUrl        *string  `json:"access_url" tfsdk:"access_url"`
	Billable         *bool    `json:"billable" tfsdk:"billable"`
	CategoryName     *string  `json:"category_name" tfsdk:"category_name"`
	Country          *string  `json:"country" tfsdk:"country"`
	Description      *string  `json:"description" tfsdk:"description"`
	FullDescription  *string  `json:"full_description" tfsdk:"full_description"`
	GettingStarted   *string  `json:"getting_started" tfsdk:"getting_started"`
	IntegrationGuide *string  `json:"integration_guide" tfsdk:"integration_guide"`
	Latitude         *float64 `json:"latitude" tfsdk:"latitude"`
	Longitude        *float64 `json:"longitude" tfsdk:"longitude"`
	Name             *string  `json:"name" tfsdk:"name"`
	PausedReason     *string  `json:"paused_reason" tfsdk:"paused_reason"`
	Shared           *bool    `json:"shared" tfsdk:"shared"`
	State            *string  `json:"state" tfsdk:"state"`
	Type             *string  `json:"type" tfsdk:"type"`
	VendorDetails    *string  `json:"vendor_details" tfsdk:"vendor_details"`
}

type ExportOrganizationGroupData struct {
	Name       *string `json:"name" tfsdk:"name"`
	ParentName *string `json:"parent_name" tfsdk:"parent_name"`
}

type ExportOrganizationGroupDataRequest struct {
	Name       *string `json:"name" tfsdk:"name"`
	ParentName *string `json:"parent_name" tfsdk:"parent_name"`
}

type ExportPlanComponentData struct {
	Amount        *int64   `json:"amount" tfsdk:"amount"`
	ComponentType *string  `json:"component_type" tfsdk:"component_type"`
	FuturePrice   *float64 `json:"future_price" tfsdk:"future_price"`
	Price         *float64 `json:"price" tfsdk:"price"`
}

type ExportPlanComponentDataRequest struct {
	Amount        *int64   `json:"amount" tfsdk:"amount"`
	ComponentType *string  `json:"component_type" tfsdk:"component_type"`
	FuturePrice   *float64 `json:"future_price" tfsdk:"future_price"`
	Price         *float64 `json:"price" tfsdk:"price"`
}

type ExportPlanData struct {
	Archived    *bool                     `json:"archived" tfsdk:"archived"`
	ArticleCode *string                   `json:"article_code" tfsdk:"article_code"`
	BackendId   *string                   `json:"backend_id" tfsdk:"backend_id"`
	Components  []ExportPlanComponentData `json:"components" tfsdk:"components"`
	Description *string                   `json:"description" tfsdk:"description"`
	MaxAmount   *int64                    `json:"max_amount" tfsdk:"max_amount"`
	Name        *string                   `json:"name" tfsdk:"name"`
	Unit        *string                   `json:"unit" tfsdk:"unit"`
	UnitPrice   *float64                  `json:"unit_price" tfsdk:"unit_price"`
}

type ExportPlanDataRequest struct {
	Archived    *bool                            `json:"archived" tfsdk:"archived"`
	ArticleCode *string                          `json:"article_code" tfsdk:"article_code"`
	BackendId   *string                          `json:"backend_id" tfsdk:"backend_id"`
	Components  []ExportPlanComponentDataRequest `json:"components" tfsdk:"components"`
	Description *string                          `json:"description" tfsdk:"description"`
	MaxAmount   *int64                           `json:"max_amount" tfsdk:"max_amount"`
	Name        *string                          `json:"name" tfsdk:"name"`
	Unit        *string                          `json:"unit" tfsdk:"unit"`
	UnitPrice   *float64                         `json:"unit_price" tfsdk:"unit_price"`
}

type ExportScreenshotData struct {
	ContentType   *string `json:"content_type" tfsdk:"content_type"`
	Description   *string `json:"description" tfsdk:"description"`
	ImageContent  *string `json:"image_content" tfsdk:"image_content"`
	ImageFilename *string `json:"image_filename" tfsdk:"image_filename"`
	Name          *string `json:"name" tfsdk:"name"`
}

type ExportScreenshotDataRequest struct {
	ContentType   *string `json:"content_type" tfsdk:"content_type"`
	Description   *string `json:"description" tfsdk:"description"`
	ImageContent  *string `json:"image_content" tfsdk:"image_content"`
	ImageFilename *string `json:"image_filename" tfsdk:"image_filename"`
	Name          *string `json:"name" tfsdk:"name"`
}

type ExportTermsOfServiceData struct {
	GracePeriodDays    *int64  `json:"grace_period_days" tfsdk:"grace_period_days"`
	IsActive           *bool   `json:"is_active" tfsdk:"is_active"`
	RequiresReconsent  *bool   `json:"requires_reconsent" tfsdk:"requires_reconsent"`
	TermsOfService     *string `json:"terms_of_service" tfsdk:"terms_of_service"`
	TermsOfServiceLink *string `json:"terms_of_service_link" tfsdk:"terms_of_service_link"`
	Version            *string `json:"version" tfsdk:"version"`
}

type ExportTermsOfServiceDataRequest struct {
	GracePeriodDays    *int64  `json:"grace_period_days" tfsdk:"grace_period_days"`
	IsActive           *bool   `json:"is_active" tfsdk:"is_active"`
	RequiresReconsent  *bool   `json:"requires_reconsent" tfsdk:"requires_reconsent"`
	TermsOfService     *string `json:"terms_of_service" tfsdk:"terms_of_service"`
	TermsOfServiceLink *string `json:"terms_of_service_link" tfsdk:"terms_of_service_link"`
	Version            *string `json:"version" tfsdk:"version"`
}

type ExternalLink struct {
	Created     *string `json:"created" tfsdk:"created"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Image       *string `json:"image,omitempty" tfsdk:"image"`
	Link        *string `json:"link" tfsdk:"link"`
	Modified    *string `json:"modified" tfsdk:"modified"`
	Name        *string `json:"name" tfsdk:"name"`
	Url         *string `json:"url" tfsdk:"url"`
}

type ExternalLinkRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Image       *string `json:"image,omitempty" tfsdk:"image"`
	Link        *string `json:"link" tfsdk:"link"`
	Name        *string `json:"name" tfsdk:"name"`
}

type ExternalLinkRequestForm struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Image       *string `json:"image,omitempty" tfsdk:"image"`
	Link        *string `json:"link" tfsdk:"link"`
	Name        *string `json:"name" tfsdk:"name"`
}

type ExternalLinkRequestMultipart struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Image       *string `json:"image,omitempty" tfsdk:"image"`
	Link        *string `json:"link" tfsdk:"link"`
	Name        *string `json:"name" tfsdk:"name"`
}

type FeatureMetadataResponse struct {
}

type Feedback struct {
	Comment      *string `json:"comment,omitempty" tfsdk:"comment"`
	Created      *string `json:"created" tfsdk:"created"`
	Evaluation   *int64  `json:"evaluation" tfsdk:"evaluation"`
	IssueKey     *string `json:"issue_key" tfsdk:"issue_key"`
	IssueSummary *string `json:"issue_summary" tfsdk:"issue_summary"`
	IssueUuid    *string `json:"issue_uuid" tfsdk:"issue_uuid"`
	Modified     *string `json:"modified" tfsdk:"modified"`
	State        *string `json:"state" tfsdk:"state"`
	UserFullName *string `json:"user_full_name" tfsdk:"user_full_name"`
}

type FinancialReport struct {
	Abbreviation        *string          `json:"abbreviation,omitempty" tfsdk:"abbreviation"`
	AccountingStartDate *string          `json:"accounting_start_date,omitempty" tfsdk:"accounting_start_date"`
	AgreementNumber     *string          `json:"agreement_number,omitempty" tfsdk:"agreement_number"`
	Created             *string          `json:"created" tfsdk:"created"`
	Name                *string          `json:"name" tfsdk:"name"`
	PaymentProfiles     []PaymentProfile `json:"payment_profiles" tfsdk:"payment_profiles"`
	RegistrationCode    *string          `json:"registration_code,omitempty" tfsdk:"registration_code"`
}

type FinancialReportEmailRequest struct {
	Emails []string `json:"emails" tfsdk:"emails"`
	Month  *int64   `json:"month" tfsdk:"month"`
	Year   *int64   `json:"year" tfsdk:"year"`
}

type Fingerprint struct {
	Md5    *string `json:"md5,omitempty" tfsdk:"md5"`
	Sha256 *string `json:"sha256,omitempty" tfsdk:"sha256"`
	Sha512 *string `json:"sha512,omitempty" tfsdk:"sha512"`
}

type FirecrestJob struct {
	AccessUrl                   *string `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created                     *string `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	File                        *string `json:"file,omitempty" tfsdk:"file"`
	Modified                    *string `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string `json:"name,omitempty" tfsdk:"name"`
	Project                     *string `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ResourceType                *string `json:"resource_type,omitempty" tfsdk:"resource_type"`
	RuntimeState                *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceName                 *string `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string `json:"state,omitempty" tfsdk:"state"`
	Url                         *string `json:"url,omitempty" tfsdk:"url"`
	User                        *string `json:"user,omitempty" tfsdk:"user"`
	UserUsername                *string `json:"user_username,omitempty" tfsdk:"user_username"`
	UserUuid                    *string `json:"user_uuid,omitempty" tfsdk:"user_uuid"`
}

type FirecrestJobRequest struct {
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	File            *string `json:"file" tfsdk:"file"`
	Name            *string `json:"name" tfsdk:"name"`
	Project         *string `json:"project" tfsdk:"project"`
	RuntimeState    *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceSettings *string `json:"service_settings" tfsdk:"service_settings"`
}

type FirecrestJobRequestForm struct {
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	File            *string `json:"file" tfsdk:"file"`
	Name            *string `json:"name" tfsdk:"name"`
	Project         *string `json:"project" tfsdk:"project"`
	RuntimeState    *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceSettings *string `json:"service_settings" tfsdk:"service_settings"`
}

type FirecrestJobRequestMultipart struct {
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	File            *string `json:"file" tfsdk:"file"`
	Name            *string `json:"name" tfsdk:"name"`
	Project         *string `json:"project" tfsdk:"project"`
	RuntimeState    *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceSettings *string `json:"service_settings" tfsdk:"service_settings"`
}

type FreeipaProfile struct {
	AgreementDate *string `json:"agreement_date,omitempty" tfsdk:"agreement_date"`
	IsActive      *bool   `json:"is_active" tfsdk:"is_active"`
	User          *string `json:"user" tfsdk:"user"`
	UserFullName  *string `json:"user_full_name" tfsdk:"user_full_name"`
	UserUsername  *string `json:"user_username" tfsdk:"user_username"`
	UserUuid      *string `json:"user_uuid" tfsdk:"user_uuid"`
	Username      *string `json:"username" tfsdk:"username"`
}

type FreeipaProfileRequest struct {
	AgreementDate *string `json:"agreement_date,omitempty" tfsdk:"agreement_date"`
	Username      *string `json:"username" tfsdk:"username"`
}

type GenericOrderAttributes struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type GoogleCalendar struct {
	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	HttpLink  *string `json:"http_link,omitempty" tfsdk:"http_link"`
	Public    *bool   `json:"public,omitempty" tfsdk:"public"`
}

type GoogleCredentials struct {
	CalendarRefreshToken *string             `json:"calendar_refresh_token,omitempty" tfsdk:"calendar_refresh_token"`
	CalendarToken        *string             `json:"calendar_token,omitempty" tfsdk:"calendar_token"`
	Created              *string             `json:"created,omitempty" tfsdk:"created"`
	Customer             *string             `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation *string             `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerCountry      *string             `json:"customer_country,omitempty" tfsdk:"customer_country"`
	CustomerImage        *string             `json:"customer_image,omitempty" tfsdk:"customer_image"`
	CustomerName         *string             `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName   *string             `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerSlug         *string             `json:"customer_slug,omitempty" tfsdk:"customer_slug"`
	CustomerUuid         *string             `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description          *string             `json:"description,omitempty" tfsdk:"description"`
	EnableNotifications  *bool               `json:"enable_notifications,omitempty" tfsdk:"enable_notifications"`
	GoogleAuthUrl        *string             `json:"google_auth_url,omitempty" tfsdk:"google_auth_url"`
	Image                *string             `json:"image,omitempty" tfsdk:"image"`
	OfferingCount        *int64              `json:"offering_count,omitempty" tfsdk:"offering_count"`
	OrganizationGroups   []OrganizationGroup `json:"organization_groups,omitempty" tfsdk:"organization_groups"`
	Url                  *string             `json:"url,omitempty" tfsdk:"url"`
}

type GroupInvitation struct {
	AutoApprove         *bool   `json:"auto_approve,omitempty" tfsdk:"auto_approve"`
	AutoCreateProject   *bool   `json:"auto_create_project,omitempty" tfsdk:"auto_create_project"`
	Created             *string `json:"created" tfsdk:"created"`
	CreatedByFullName   *string `json:"created_by_full_name" tfsdk:"created_by_full_name"`
	CreatedByImage      *string `json:"created_by_image" tfsdk:"created_by_image"`
	CreatedByUsername   *string `json:"created_by_username" tfsdk:"created_by_username"`
	CustomerName        *string `json:"customer_name" tfsdk:"customer_name"`
	CustomerUuid        *string `json:"customer_uuid" tfsdk:"customer_uuid"`
	IsActive            *bool   `json:"is_active" tfsdk:"is_active"`
	IsPublic            *bool   `json:"is_public,omitempty" tfsdk:"is_public"`
	ProjectNameTemplate *string `json:"project_name_template,omitempty" tfsdk:"project_name_template"`
	ProjectRole         *string `json:"project_role,omitempty" tfsdk:"project_role"`
	Role                *string `json:"role" tfsdk:"role"`
	RoleDescription     *string `json:"role_description" tfsdk:"role_description"`
	RoleName            *string `json:"role_name" tfsdk:"role_name"`
	ScopeDescription    *string `json:"scope_description" tfsdk:"scope_description"`
	ScopeImage          *string `json:"scope_image" tfsdk:"scope_image"`
	ScopeName           *string `json:"scope_name" tfsdk:"scope_name"`
	ScopeType           *string `json:"scope_type" tfsdk:"scope_type"`
	ScopeUuid           *string `json:"scope_uuid" tfsdk:"scope_uuid"`
	Url                 *string `json:"url" tfsdk:"url"`
}

type GroupInvitationRequest struct {
	AutoApprove         *bool   `json:"auto_approve,omitempty" tfsdk:"auto_approve"`
	AutoCreateProject   *bool   `json:"auto_create_project,omitempty" tfsdk:"auto_create_project"`
	IsPublic            *bool   `json:"is_public,omitempty" tfsdk:"is_public"`
	ProjectNameTemplate *string `json:"project_name_template,omitempty" tfsdk:"project_name_template"`
	ProjectRole         *string `json:"project_role,omitempty" tfsdk:"project_role"`
	Role                *string `json:"role" tfsdk:"role"`
	Scope               *string `json:"scope" tfsdk:"scope"`
}

type GuestOsEnum struct {
}

type GuestPowerStateEnum struct {
}

type IPMapping struct {
	ExternalIp *string `json:"external_ip,omitempty" tfsdk:"external_ip"`
	FloatingIp *string `json:"floating_ip,omitempty" tfsdk:"floating_ip"`
}

type IPMappingRequest struct {
	ExternalIp *string `json:"external_ip" tfsdk:"external_ip"`
	FloatingIp *string `json:"floating_ip" tfsdk:"floating_ip"`
}

type IdentityProvider struct {
	AuthUrl                  *string `json:"auth_url" tfsdk:"auth_url"`
	ClientId                 *string `json:"client_id" tfsdk:"client_id"`
	ClientSecret             *string `json:"client_secret" tfsdk:"client_secret"`
	DiscoveryUrl             *string `json:"discovery_url" tfsdk:"discovery_url"`
	EnablePkce               *bool   `json:"enable_pkce,omitempty" tfsdk:"enable_pkce"`
	EnablePostLogoutRedirect *bool   `json:"enable_post_logout_redirect,omitempty" tfsdk:"enable_post_logout_redirect"`
	ExtraFields              *string `json:"extra_fields,omitempty" tfsdk:"extra_fields"`
	ExtraScope               *string `json:"extra_scope,omitempty" tfsdk:"extra_scope"`
	IsActive                 *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	Label                    *string `json:"label" tfsdk:"label"`
	LogoutUrl                *string `json:"logout_url" tfsdk:"logout_url"`
	ManagementUrl            *string `json:"management_url,omitempty" tfsdk:"management_url"`
	Provider                 *string `json:"provider" tfsdk:"provider"`
	TokenUrl                 *string `json:"token_url" tfsdk:"token_url"`
	UserClaim                *string `json:"user_claim,omitempty" tfsdk:"user_claim"`
	UserField                *string `json:"user_field,omitempty" tfsdk:"user_field"`
	UserinfoUrl              *string `json:"userinfo_url" tfsdk:"userinfo_url"`
	VerifySsl                *bool   `json:"verify_ssl,omitempty" tfsdk:"verify_ssl"`
}

type IdentityProviderRequest struct {
	ClientId                 *string `json:"client_id" tfsdk:"client_id"`
	ClientSecret             *string `json:"client_secret" tfsdk:"client_secret"`
	DiscoveryUrl             *string `json:"discovery_url" tfsdk:"discovery_url"`
	EnablePkce               *bool   `json:"enable_pkce,omitempty" tfsdk:"enable_pkce"`
	EnablePostLogoutRedirect *bool   `json:"enable_post_logout_redirect,omitempty" tfsdk:"enable_post_logout_redirect"`
	ExtraFields              *string `json:"extra_fields,omitempty" tfsdk:"extra_fields"`
	ExtraScope               *string `json:"extra_scope,omitempty" tfsdk:"extra_scope"`
	IsActive                 *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	Label                    *string `json:"label" tfsdk:"label"`
	ManagementUrl            *string `json:"management_url,omitempty" tfsdk:"management_url"`
	Provider                 *string `json:"provider" tfsdk:"provider"`
	UserClaim                *string `json:"user_claim,omitempty" tfsdk:"user_claim"`
	UserField                *string `json:"user_field,omitempty" tfsdk:"user_field"`
	VerifySsl                *bool   `json:"verify_ssl,omitempty" tfsdk:"verify_ssl"`
}

type ImageCreateRequest struct {
	ContainerFormat *string `json:"container_format,omitempty" tfsdk:"container_format"`
	DiskFormat      *string `json:"disk_format,omitempty" tfsdk:"disk_format"`
	MinDisk         *int64  `json:"min_disk,omitempty" tfsdk:"min_disk"`
	MinRam          *int64  `json:"min_ram,omitempty" tfsdk:"min_ram"`
	Name            *string `json:"name" tfsdk:"name"`
	Visibility      *string `json:"visibility,omitempty" tfsdk:"visibility"`
}

type ImageCreateResponse struct {
	ImageId   *string `json:"image_id" tfsdk:"image_id"`
	Name      *string `json:"name" tfsdk:"name"`
	Status    *string `json:"status" tfsdk:"status"`
	UploadUrl *string `json:"upload_url" tfsdk:"upload_url"`
}

type ImageUploadResponse struct {
	Message *string `json:"message" tfsdk:"message"`
	Status  *string `json:"status" tfsdk:"status"`
}

type ImpactLevelDisplayEnum struct {
}

type ImpactLevelEnum struct {
}

type ImportResourceRequest struct {
	BackendId *string `json:"backend_id" tfsdk:"backend_id"`
	Plan      *string `json:"plan,omitempty" tfsdk:"plan"`
	Project   *string `json:"project" tfsdk:"project"`
}

type ImportableResource struct {
	BackendId   *string `json:"backend_id" tfsdk:"backend_id"`
	Description *string `json:"description" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
	Type        *string `json:"type" tfsdk:"type"`
}

type InstanceFlavorChangeRequest struct {
	Flavor *string `json:"flavor" tfsdk:"flavor"`
}

type IntegrationStatus struct {
	AgentType            *string `json:"agent_type,omitempty" tfsdk:"agent_type"`
	LastRequestTimestamp *string `json:"last_request_timestamp,omitempty" tfsdk:"last_request_timestamp"`
	ServiceName          *string `json:"service_name,omitempty" tfsdk:"service_name"`
	Status               *string `json:"status,omitempty" tfsdk:"status"`
}

type IntegrationStatusDetails struct {
	LastRequestTimestamp *string `json:"last_request_timestamp" tfsdk:"last_request_timestamp"`
	Offering             *string `json:"offering" tfsdk:"offering"`
	Status               *string `json:"status" tfsdk:"status"`
	Url                  *string `json:"url" tfsdk:"url"`
}

type Invitation struct {
	CivilNumber         *string `json:"civil_number,omitempty" tfsdk:"civil_number"`
	Created             *string `json:"created" tfsdk:"created"`
	CreatedByFullName   *string `json:"created_by_full_name" tfsdk:"created_by_full_name"`
	CreatedByImage      *string `json:"created_by_image" tfsdk:"created_by_image"`
	CreatedByUsername   *string `json:"created_by_username" tfsdk:"created_by_username"`
	CustomerName        *string `json:"customer_name" tfsdk:"customer_name"`
	CustomerUuid        *string `json:"customer_uuid" tfsdk:"customer_uuid"`
	Email               *string `json:"email" tfsdk:"email"`
	ErrorMessage        *string `json:"error_message" tfsdk:"error_message"`
	ExecutionState      *string `json:"execution_state" tfsdk:"execution_state"`
	Expires             *string `json:"expires" tfsdk:"expires"`
	ExtraInvitationText *string `json:"extra_invitation_text,omitempty" tfsdk:"extra_invitation_text"`
	FullName            *string `json:"full_name,omitempty" tfsdk:"full_name"`
	JobTitle            *string `json:"job_title,omitempty" tfsdk:"job_title"`
	NativeName          *string `json:"native_name,omitempty" tfsdk:"native_name"`
	Organization        *string `json:"organization,omitempty" tfsdk:"organization"`
	PhoneNumber         *string `json:"phone_number,omitempty" tfsdk:"phone_number"`
	Role                *string `json:"role" tfsdk:"role"`
	RoleDescription     *string `json:"role_description" tfsdk:"role_description"`
	RoleName            *string `json:"role_name" tfsdk:"role_name"`
	ScopeDescription    *string `json:"scope_description" tfsdk:"scope_description"`
	ScopeName           *string `json:"scope_name" tfsdk:"scope_name"`
	ScopeType           *string `json:"scope_type" tfsdk:"scope_type"`
	ScopeUuid           *string `json:"scope_uuid" tfsdk:"scope_uuid"`
	State               *string `json:"state" tfsdk:"state"`
	Url                 *string `json:"url" tfsdk:"url"`
}

type InvitationCheck struct {
	CivilNumberRequired *bool   `json:"civil_number_required,omitempty" tfsdk:"civil_number_required"`
	Email               *string `json:"email" tfsdk:"email"`
}

type InvitationRequest struct {
	CivilNumber         *string `json:"civil_number,omitempty" tfsdk:"civil_number"`
	Email               *string `json:"email" tfsdk:"email"`
	ExtraInvitationText *string `json:"extra_invitation_text,omitempty" tfsdk:"extra_invitation_text"`
	FullName            *string `json:"full_name,omitempty" tfsdk:"full_name"`
	JobTitle            *string `json:"job_title,omitempty" tfsdk:"job_title"`
	NativeName          *string `json:"native_name,omitempty" tfsdk:"native_name"`
	Organization        *string `json:"organization,omitempty" tfsdk:"organization"`
	PhoneNumber         *string `json:"phone_number,omitempty" tfsdk:"phone_number"`
	Role                *string `json:"role" tfsdk:"role"`
	Scope               *string `json:"scope" tfsdk:"scope"`
}

type InvitationState struct {
}

type InvitationStateEnum struct {
}

type InvitationUpdate struct {
	Email *string `json:"email" tfsdk:"email"`
	Role  *string `json:"role,omitempty" tfsdk:"role"`
}

type InvitationUpdateRequest struct {
	Email *string `json:"email" tfsdk:"email"`
	Role  *string `json:"role,omitempty" tfsdk:"role"`
}

type Invoice struct {
	BackendId       *string          `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Compensations   *float64         `json:"compensations,omitempty" tfsdk:"compensations"`
	Customer        *string          `json:"customer,omitempty" tfsdk:"customer"`
	CustomerDetails *CustomerDetails `json:"customer_details,omitempty" tfsdk:"customer_details"`
	DueDate         *string          `json:"due_date,omitempty" tfsdk:"due_date"`
	IncurredCosts   *float64         `json:"incurred_costs,omitempty" tfsdk:"incurred_costs"`
	InvoiceDate     *string          `json:"invoice_date,omitempty" tfsdk:"invoice_date"`
	Items           []InvoiceItem    `json:"items,omitempty" tfsdk:"items"`
	Month           *int64           `json:"month,omitempty" tfsdk:"month"`
	Number          *int64           `json:"number,omitempty" tfsdk:"number"`
	PaymentUrl      *string          `json:"payment_url,omitempty" tfsdk:"payment_url"`
	Price           *string          `json:"price,omitempty" tfsdk:"price"`
	ReferenceNumber *string          `json:"reference_number,omitempty" tfsdk:"reference_number"`
	State           *string          `json:"state,omitempty" tfsdk:"state"`
	Tax             *string          `json:"tax,omitempty" tfsdk:"tax"`
	Total           *string          `json:"total,omitempty" tfsdk:"total"`
	Url             *string          `json:"url,omitempty" tfsdk:"url"`
	Year            *int64           `json:"year,omitempty" tfsdk:"year"`
}

type InvoiceCost struct {
	Month *int64   `json:"month" tfsdk:"month"`
	Price *float64 `json:"price" tfsdk:"price"`
	Year  *int64   `json:"year" tfsdk:"year"`
}

type InvoiceGrowth struct {
	CustomerPeriods []InvoiceGrowthCustomerPeriod `json:"customer_periods" tfsdk:"customer_periods"`
	Periods         []string                      `json:"periods" tfsdk:"periods"`
}

type InvoiceGrowthCustomerPeriod struct {
	Name *string `json:"name" tfsdk:"name"`
}

type InvoiceItem struct {
	ArticleCode  *string             `json:"article_code,omitempty" tfsdk:"article_code"`
	BackendUuid  *string             `json:"backend_uuid,omitempty" tfsdk:"backend_uuid"`
	BillingType  *string             `json:"billing_type,omitempty" tfsdk:"billing_type"`
	Credit       *bool               `json:"credit,omitempty" tfsdk:"credit"`
	Details      *InvoiceItemDetails `json:"details,omitempty" tfsdk:"details"`
	End          *string             `json:"end,omitempty" tfsdk:"end"`
	Factor       *int64              `json:"factor,omitempty" tfsdk:"factor"`
	MeasuredUnit *string             `json:"measured_unit,omitempty" tfsdk:"measured_unit"`
	Name         *string             `json:"name,omitempty" tfsdk:"name"`
	Price        *float64            `json:"price,omitempty" tfsdk:"price"`
	ProjectName  *string             `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid  *string             `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	Quantity     *string             `json:"quantity,omitempty" tfsdk:"quantity"`
	Resource     *string             `json:"resource,omitempty" tfsdk:"resource"`
	ResourceName *string             `json:"resource_name,omitempty" tfsdk:"resource_name"`
	ResourceUuid *string             `json:"resource_uuid,omitempty" tfsdk:"resource_uuid"`
	Start        *string             `json:"start,omitempty" tfsdk:"start"`
	Tax          *string             `json:"tax,omitempty" tfsdk:"tax"`
	Total        *string             `json:"total,omitempty" tfsdk:"total"`
	Unit         *string             `json:"unit,omitempty" tfsdk:"unit"`
	UnitPrice    *string             `json:"unit_price,omitempty" tfsdk:"unit_price"`
	Url          *string             `json:"url,omitempty" tfsdk:"url"`
}

type InvoiceItemCompensation struct {
	OfferingComponentName *string `json:"offering_component_name" tfsdk:"offering_component_name"`
}

type InvoiceItemCompensationRequest struct {
	OfferingComponentName *string `json:"offering_component_name" tfsdk:"offering_component_name"`
}

type InvoiceItemDetail struct {
	ArticleCode           *string `json:"article_code,omitempty" tfsdk:"article_code"`
	End                   *string `json:"end,omitempty" tfsdk:"end"`
	Invoice               *string `json:"invoice" tfsdk:"invoice"`
	MeasuredUnit          *string `json:"measured_unit,omitempty" tfsdk:"measured_unit"`
	Name                  *string `json:"name,omitempty" tfsdk:"name"`
	OfferingComponentType *string `json:"offering_component_type" tfsdk:"offering_component_type"`
	OfferingUuid          *string `json:"offering_uuid" tfsdk:"offering_uuid"`
	Quantity              *string `json:"quantity,omitempty" tfsdk:"quantity"`
	Resource              *string `json:"resource,omitempty" tfsdk:"resource"`
	Start                 *string `json:"start,omitempty" tfsdk:"start"`
	Unit                  *string `json:"unit,omitempty" tfsdk:"unit"`
	UnitPrice             *string `json:"unit_price,omitempty" tfsdk:"unit_price"`
}

type InvoiceItemDetails struct {
	OfferingComponentName *string               `json:"offering_component_name,omitempty" tfsdk:"offering_component_name"`
	OfferingComponentType *string               `json:"offering_component_type,omitempty" tfsdk:"offering_component_type"`
	OfferingName          *string               `json:"offering_name,omitempty" tfsdk:"offering_name"`
	OfferingType          *string               `json:"offering_type,omitempty" tfsdk:"offering_type"`
	OfferingUuid          *string               `json:"offering_uuid,omitempty" tfsdk:"offering_uuid"`
	PlanComponentId       *int64                `json:"plan_component_id,omitempty" tfsdk:"plan_component_id"`
	PlanName              *string               `json:"plan_name,omitempty" tfsdk:"plan_name"`
	PlanUuid              *string               `json:"plan_uuid,omitempty" tfsdk:"plan_uuid"`
	ResourceLimitPeriods  []ResourceLimitPeriod `json:"resource_limit_periods,omitempty" tfsdk:"resource_limit_periods"`
	ResourceName          *string               `json:"resource_name,omitempty" tfsdk:"resource_name"`
	ResourceUuid          *string               `json:"resource_uuid,omitempty" tfsdk:"resource_uuid"`
	ServiceProviderName   *string               `json:"service_provider_name,omitempty" tfsdk:"service_provider_name"`
	ServiceProviderUuid   *string               `json:"service_provider_uuid,omitempty" tfsdk:"service_provider_uuid"`
}

type InvoiceItemMigrateTo struct {
	Invoice *string `json:"invoice" tfsdk:"invoice"`
}

type InvoiceItemMigrateToRequest struct {
	Invoice *string `json:"invoice" tfsdk:"invoice"`
}

type InvoiceItemTotalPrice struct {
	TotalPrice *string `json:"total_price" tfsdk:"total_price"`
}

type InvoiceItemUpdate struct {
	ArticleCode *string `json:"article_code,omitempty" tfsdk:"article_code"`
	End         *string `json:"end,omitempty" tfsdk:"end"`
	Quantity    *string `json:"quantity,omitempty" tfsdk:"quantity"`
	Start       *string `json:"start,omitempty" tfsdk:"start"`
	UnitPrice   *string `json:"unit_price,omitempty" tfsdk:"unit_price"`
}

type InvoiceItemUpdateRequest struct {
	ArticleCode *string `json:"article_code,omitempty" tfsdk:"article_code"`
	End         *string `json:"end,omitempty" tfsdk:"end"`
	Quantity    *string `json:"quantity,omitempty" tfsdk:"quantity"`
	Start       *string `json:"start,omitempty" tfsdk:"start"`
	UnitPrice   *string `json:"unit_price,omitempty" tfsdk:"unit_price"`
}

type InvoiceStateEnum struct {
}

type InvoiceStatsOffering struct {
	AggregatedPrice      *float64 `json:"aggregated_price" tfsdk:"aggregated_price"`
	AggregatedTax        *float64 `json:"aggregated_tax" tfsdk:"aggregated_tax"`
	AggregatedTotal      *float64 `json:"aggregated_total" tfsdk:"aggregated_total"`
	OfferingName         *string  `json:"offering_name" tfsdk:"offering_name"`
	ServiceCategoryTitle *string  `json:"service_category_title" tfsdk:"service_category_title"`
	ServiceProviderName  *string  `json:"service_provider_name" tfsdk:"service_provider_name"`
	ServiceProviderUuid  *string  `json:"service_provider_uuid" tfsdk:"service_provider_uuid"`
}

type Issue struct {
	AddAttachmentIsAvailable *bool   `json:"add_attachment_is_available" tfsdk:"add_attachment_is_available"`
	AddCommentIsAvailable    *bool   `json:"add_comment_is_available" tfsdk:"add_comment_is_available"`
	Assignee                 *string `json:"assignee,omitempty" tfsdk:"assignee"`
	AssigneeName             *string `json:"assignee_name" tfsdk:"assignee_name"`
	AssigneeUuid             *string `json:"assignee_uuid" tfsdk:"assignee_uuid"`
	BackendId                *string `json:"backend_id" tfsdk:"backend_id"`
	BackendName              *string `json:"backend_name" tfsdk:"backend_name"`
	Caller                   *string `json:"caller,omitempty" tfsdk:"caller"`
	CallerFullName           *string `json:"caller_full_name" tfsdk:"caller_full_name"`
	CallerUuid               *string `json:"caller_uuid" tfsdk:"caller_uuid"`
	Created                  *string `json:"created" tfsdk:"created"`
	Customer                 *string `json:"customer,omitempty" tfsdk:"customer"`
	CustomerName             *string `json:"customer_name" tfsdk:"customer_name"`
	CustomerUuid             *string `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description              *string `json:"description,omitempty" tfsdk:"description"`
	DestroyIsAvailable       *bool   `json:"destroy_is_available" tfsdk:"destroy_is_available"`
	Key                      *string `json:"key" tfsdk:"key"`
	Link                     *string `json:"link" tfsdk:"link"`
	Modified                 *string `json:"modified" tfsdk:"modified"`
	OrderCustomerUuid        *string `json:"order_customer_uuid" tfsdk:"order_customer_uuid"`
	OrderProjectUuid         *string `json:"order_project_uuid" tfsdk:"order_project_uuid"`
	OrderResourceName        *string `json:"order_resource_name" tfsdk:"order_resource_name"`
	OrderUuid                *string `json:"order_uuid" tfsdk:"order_uuid"`
	Priority                 *string `json:"priority,omitempty" tfsdk:"priority"`
	Project                  *string `json:"project,omitempty" tfsdk:"project"`
	ProjectName              *string `json:"project_name" tfsdk:"project_name"`
	ProjectUuid              *string `json:"project_uuid" tfsdk:"project_uuid"`
	RemoteId                 *string `json:"remote_id,omitempty" tfsdk:"remote_id"`
	Reporter                 *string `json:"reporter" tfsdk:"reporter"`
	ReporterName             *string `json:"reporter_name" tfsdk:"reporter_name"`
	ReporterUuid             *string `json:"reporter_uuid" tfsdk:"reporter_uuid"`
	Resolution               *string `json:"resolution" tfsdk:"resolution"`
	Resolved                 *bool   `json:"resolved" tfsdk:"resolved"`
	Resource                 *string `json:"resource,omitempty" tfsdk:"resource"`
	ResourceName             *string `json:"resource_name" tfsdk:"resource_name"`
	ResourceType             *string `json:"resource_type" tfsdk:"resource_type"`
	Status                   *string `json:"status" tfsdk:"status"`
	Summary                  *string `json:"summary" tfsdk:"summary"`
	Template                 *string `json:"template,omitempty" tfsdk:"template"`
	Type                     *string `json:"type" tfsdk:"type"`
	UpdateIsAvailable        *bool   `json:"update_is_available" tfsdk:"update_is_available"`
	Url                      *string `json:"url" tfsdk:"url"`
}

type IssueReference struct {
	Key *string `json:"key,omitempty" tfsdk:"key"`
}

type IssueRequest struct {
	Assignee           *string `json:"assignee,omitempty" tfsdk:"assignee"`
	Caller             *string `json:"caller,omitempty" tfsdk:"caller"`
	Customer           *string `json:"customer,omitempty" tfsdk:"customer"`
	Description        *string `json:"description,omitempty" tfsdk:"description"`
	IsReportedManually *bool   `json:"is_reported_manually,omitempty" tfsdk:"is_reported_manually"`
	Priority           *string `json:"priority,omitempty" tfsdk:"priority"`
	Project            *string `json:"project,omitempty" tfsdk:"project"`
	RemoteId           *string `json:"remote_id,omitempty" tfsdk:"remote_id"`
	Resource           *string `json:"resource,omitempty" tfsdk:"resource"`
	Summary            *string `json:"summary" tfsdk:"summary"`
	Template           *string `json:"template,omitempty" tfsdk:"template"`
	Type               *string `json:"type" tfsdk:"type"`
}

type IssueStatus struct {
	Name        *string `json:"name" tfsdk:"name"`
	Type        *int64  `json:"type,omitempty" tfsdk:"type"`
	TypeDisplay *string `json:"type_display" tfsdk:"type_display"`
	Url         *string `json:"url" tfsdk:"url"`
}

type IssueStatusCreate struct {
	Name *string `json:"name" tfsdk:"name"`
	Type *int64  `json:"type,omitempty" tfsdk:"type"`
}

type IssueStatusCreateRequest struct {
	Name *string `json:"name" tfsdk:"name"`
	Type *int64  `json:"type,omitempty" tfsdk:"type"`
}

type IssueStatusType struct {
}

type IssueTypeEnum struct {
}

type JiraChangelog struct {
}

type JiraChangelogRequest struct {
}

type JiraComment struct {
	Id *string `json:"id" tfsdk:"id"`
}

type JiraCommentRequest struct {
	Id *string `json:"id" tfsdk:"id"`
}

type JiraIssue struct {
	Fields *JiraIssueFields `json:"fields" tfsdk:"fields"`
	Key    *string          `json:"key" tfsdk:"key"`
}

type JiraIssueFields struct {
	Project *JiraIssueProject `json:"project" tfsdk:"project"`
}

type JiraIssueFieldsRequest struct {
	Project *JiraIssueProjectRequest `json:"project" tfsdk:"project"`
}

type JiraIssueProject struct {
	Id   *string `json:"id" tfsdk:"id"`
	Key  *string `json:"key" tfsdk:"key"`
	Name *string `json:"name" tfsdk:"name"`
}

type JiraIssueProjectRequest struct {
	Id   *string `json:"id" tfsdk:"id"`
	Key  *string `json:"key" tfsdk:"key"`
	Name *string `json:"name" tfsdk:"name"`
}

type JiraIssueRequest struct {
	Fields *JiraIssueFieldsRequest `json:"fields" tfsdk:"fields"`
	Key    *string                 `json:"key" tfsdk:"key"`
}

type K8sDefaultConfiguration struct {
	AvailableKubernetesVersions   *string `json:"available_kubernetes_versions,omitempty" tfsdk:"available_kubernetes_versions"`
	DefaultControllerEtcdDiskGb   *int64  `json:"default_controller_etcd_disk_gb,omitempty" tfsdk:"default_controller_etcd_disk_gb"`
	DefaultControllerRamGb        *int64  `json:"default_controller_ram_gb,omitempty" tfsdk:"default_controller_ram_gb"`
	DefaultControllerSystemDiskGb *int64  `json:"default_controller_system_disk_gb,omitempty" tfsdk:"default_controller_system_disk_gb"`
	DefaultControllerVcpus        *int64  `json:"default_controller_vcpus,omitempty" tfsdk:"default_controller_vcpus"`
	DefaultLbLogsDiskGb           *int64  `json:"default_lb_logs_disk_gb,omitempty" tfsdk:"default_lb_logs_disk_gb"`
	DefaultLbRamGb                *int64  `json:"default_lb_ram_gb,omitempty" tfsdk:"default_lb_ram_gb"`
	DefaultLbSystemDiskGb         *int64  `json:"default_lb_system_disk_gb,omitempty" tfsdk:"default_lb_system_disk_gb"`
	DefaultLbVcpus                *int64  `json:"default_lb_vcpus,omitempty" tfsdk:"default_lb_vcpus"`
	DefaultStorageDataDiskGb      *int64  `json:"default_storage_data_disk_gb,omitempty" tfsdk:"default_storage_data_disk_gb"`
	DefaultStorageSanDiskGb       *int64  `json:"default_storage_san_disk_gb,omitempty" tfsdk:"default_storage_san_disk_gb"`
	DefaultWorkerDataDiskGb       *int64  `json:"default_worker_data_disk_gb,omitempty" tfsdk:"default_worker_data_disk_gb"`
	MinimalWorkerRamGb            *int64  `json:"minimal_worker_ram_gb,omitempty" tfsdk:"minimal_worker_ram_gb"`
	MinimalWorkerVcpus            *int64  `json:"minimal_worker_vcpus,omitempty" tfsdk:"minimal_worker_vcpus"`
}

type K8sDefaultConfigurationRequest struct {
	AvailableKubernetesVersions   *string `json:"available_kubernetes_versions,omitempty" tfsdk:"available_kubernetes_versions"`
	DefaultControllerEtcdDiskGb   *int64  `json:"default_controller_etcd_disk_gb,omitempty" tfsdk:"default_controller_etcd_disk_gb"`
	DefaultControllerRamGb        *int64  `json:"default_controller_ram_gb,omitempty" tfsdk:"default_controller_ram_gb"`
	DefaultControllerSystemDiskGb *int64  `json:"default_controller_system_disk_gb,omitempty" tfsdk:"default_controller_system_disk_gb"`
	DefaultControllerVcpus        *int64  `json:"default_controller_vcpus,omitempty" tfsdk:"default_controller_vcpus"`
	DefaultLbLogsDiskGb           *int64  `json:"default_lb_logs_disk_gb,omitempty" tfsdk:"default_lb_logs_disk_gb"`
	DefaultLbRamGb                *int64  `json:"default_lb_ram_gb,omitempty" tfsdk:"default_lb_ram_gb"`
	DefaultLbSystemDiskGb         *int64  `json:"default_lb_system_disk_gb,omitempty" tfsdk:"default_lb_system_disk_gb"`
	DefaultLbVcpus                *int64  `json:"default_lb_vcpus,omitempty" tfsdk:"default_lb_vcpus"`
	DefaultStorageDataDiskGb      *int64  `json:"default_storage_data_disk_gb,omitempty" tfsdk:"default_storage_data_disk_gb"`
	DefaultStorageSanDiskGb       *int64  `json:"default_storage_san_disk_gb,omitempty" tfsdk:"default_storage_san_disk_gb"`
	DefaultWorkerDataDiskGb       *int64  `json:"default_worker_data_disk_gb,omitempty" tfsdk:"default_worker_data_disk_gb"`
	MinimalWorkerRamGb            *int64  `json:"minimal_worker_ram_gb,omitempty" tfsdk:"minimal_worker_ram_gb"`
	MinimalWorkerVcpus            *int64  `json:"minimal_worker_vcpus,omitempty" tfsdk:"minimal_worker_vcpus"`
}

type KeycloakGroup struct {
	BackendId *string `json:"backend_id" tfsdk:"backend_id"`
	Created   *string `json:"created" tfsdk:"created"`
	Modified  *string `json:"modified" tfsdk:"modified"`
	Name      *string `json:"name" tfsdk:"name"`
	Role      *string `json:"role" tfsdk:"role"`
	ScopeName *string `json:"scope_name" tfsdk:"scope_name"`
	ScopeType *string `json:"scope_type" tfsdk:"scope_type"`
	ScopeUuid *string `json:"scope_uuid" tfsdk:"scope_uuid"`
	Url       *string `json:"url" tfsdk:"url"`
}

type KeycloakUserGroupMembership struct {
	Created        *string `json:"created" tfsdk:"created"`
	Email          *string `json:"email" tfsdk:"email"`
	ErrorMessage   *string `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback *string `json:"error_traceback" tfsdk:"error_traceback"`
	FirstName      *string `json:"first_name" tfsdk:"first_name"`
	Group          *string `json:"group" tfsdk:"group"`
	GroupName      *string `json:"group_name" tfsdk:"group_name"`
	GroupRole      *string `json:"group_role" tfsdk:"group_role"`
	GroupScopeName *string `json:"group_scope_name" tfsdk:"group_scope_name"`
	GroupScopeType *string `json:"group_scope_type" tfsdk:"group_scope_type"`
	LastChecked    *string `json:"last_checked" tfsdk:"last_checked"`
	LastName       *string `json:"last_name" tfsdk:"last_name"`
	Modified       *string `json:"modified" tfsdk:"modified"`
	State          *string `json:"state" tfsdk:"state"`
	Url            *string `json:"url" tfsdk:"url"`
	Username       *string `json:"username" tfsdk:"username"`
}

type KeycloakUserGroupMembershipRequest struct {
	Email     *string `json:"email" tfsdk:"email"`
	Role      *string `json:"role" tfsdk:"role"`
	ScopeUuid *string `json:"scope_uuid" tfsdk:"scope_uuid"`
	Username  *string `json:"username" tfsdk:"username"`
}

type KeycloakUserGroupMembershipState struct {
}

type KindEnum struct {
}

type LexisLink struct {
	Created              *string `json:"created" tfsdk:"created"`
	CustomerName         *string `json:"customer_name" tfsdk:"customer_name"`
	CustomerUuid         *string `json:"customer_uuid" tfsdk:"customer_uuid"`
	HeappeProjectId      *int64  `json:"heappe_project_id,omitempty" tfsdk:"heappe_project_id"`
	Modified             *string `json:"modified" tfsdk:"modified"`
	ProjectName          *string `json:"project_name" tfsdk:"project_name"`
	ProjectUuid          *string `json:"project_uuid" tfsdk:"project_uuid"`
	ResourceBackendId    *string `json:"resource_backend_id" tfsdk:"resource_backend_id"`
	ResourceEndDate      *string `json:"resource_end_date" tfsdk:"resource_end_date"`
	ResourceName         *string `json:"resource_name" tfsdk:"resource_name"`
	ResourceType         *string `json:"resource_type" tfsdk:"resource_type"`
	ResourceUuid         *string `json:"resource_uuid" tfsdk:"resource_uuid"`
	RobotAccount         *string `json:"robot_account" tfsdk:"robot_account"`
	RobotAccountType     *string `json:"robot_account_type" tfsdk:"robot_account_type"`
	RobotAccountUsername *string `json:"robot_account_username" tfsdk:"robot_account_username"`
	State                *string `json:"state" tfsdk:"state"`
	Url                  *string `json:"url" tfsdk:"url"`
}

type LexisLinkCreateRequest struct {
	Resource *string `json:"resource" tfsdk:"resource"`
}

type LexisLinkRequest struct {
	HeappeProjectId *int64 `json:"heappe_project_id,omitempty" tfsdk:"heappe_project_id"`
}

type LimitPeriodEnum struct {
}

type LimitTypeEnum struct {
}

type LinkOpenstackRequest struct {
	Instance *string `json:"instance" tfsdk:"instance"`
}

type LinkToInvoice struct {
	Invoice *string `json:"invoice" tfsdk:"invoice"`
}

type LinkToInvoiceRequest struct {
	Invoice *string `json:"invoice" tfsdk:"invoice"`
}

type Logout struct {
	LogoutUrl *string `json:"logout_url" tfsdk:"logout_url"`
}

type MaintenanceActionResponse struct {
	Detail *string `json:"detail" tfsdk:"detail"`
}

type MaintenanceAnnouncement struct {
	ActualEnd            *string                           `json:"actual_end" tfsdk:"actual_end"`
	ActualStart          *string                           `json:"actual_start" tfsdk:"actual_start"`
	AffectedOfferings    []MaintenanceAnnouncementOffering `json:"affected_offerings" tfsdk:"affected_offerings"`
	BackendId            *string                           `json:"backend_id" tfsdk:"backend_id"`
	CreatedBy            *string                           `json:"created_by" tfsdk:"created_by"`
	ExternalReferenceUrl *string                           `json:"external_reference_url,omitempty" tfsdk:"external_reference_url"`
	InternalNotes        *string                           `json:"internal_notes,omitempty" tfsdk:"internal_notes"`
	MaintenanceType      *int64                            `json:"maintenance_type,omitempty" tfsdk:"maintenance_type"`
	Message              *string                           `json:"message,omitempty" tfsdk:"message"`
	Name                 *string                           `json:"name" tfsdk:"name"`
	ScheduledEnd         *string                           `json:"scheduled_end" tfsdk:"scheduled_end"`
	ScheduledStart       *string                           `json:"scheduled_start" tfsdk:"scheduled_start"`
	ServiceProvider      *string                           `json:"service_provider" tfsdk:"service_provider"`
	ServiceProviderName  *string                           `json:"service_provider_name" tfsdk:"service_provider_name"`
	State                *string                           `json:"state" tfsdk:"state"`
	Url                  *string                           `json:"url" tfsdk:"url"`
}

type MaintenanceAnnouncementOffering struct {
	ImpactDescription  *string `json:"impact_description,omitempty" tfsdk:"impact_description"`
	ImpactLevel        *int64  `json:"impact_level,omitempty" tfsdk:"impact_level"`
	ImpactLevelDisplay *string `json:"impact_level_display" tfsdk:"impact_level_display"`
	Maintenance        *string `json:"maintenance" tfsdk:"maintenance"`
	Offering           *string `json:"offering" tfsdk:"offering"`
	OfferingName       *string `json:"offering_name" tfsdk:"offering_name"`
	Url                *string `json:"url" tfsdk:"url"`
}

type MaintenanceAnnouncementOfferingRequest struct {
	ImpactDescription *string `json:"impact_description,omitempty" tfsdk:"impact_description"`
	ImpactLevel       *int64  `json:"impact_level,omitempty" tfsdk:"impact_level"`
	Maintenance       *string `json:"maintenance" tfsdk:"maintenance"`
	Offering          *string `json:"offering" tfsdk:"offering"`
}

type MaintenanceAnnouncementOfferingTemplate struct {
	ImpactDescription   *string `json:"impact_description,omitempty" tfsdk:"impact_description"`
	ImpactLevel         *int64  `json:"impact_level,omitempty" tfsdk:"impact_level"`
	MaintenanceTemplate *string `json:"maintenance_template" tfsdk:"maintenance_template"`
	Offering            *string `json:"offering" tfsdk:"offering"`
	OfferingName        *string `json:"offering_name" tfsdk:"offering_name"`
	OfferingUuid        *string `json:"offering_uuid" tfsdk:"offering_uuid"`
	Url                 *string `json:"url" tfsdk:"url"`
}

type MaintenanceAnnouncementOfferingTemplateRequest struct {
	ImpactDescription   *string `json:"impact_description,omitempty" tfsdk:"impact_description"`
	ImpactLevel         *int64  `json:"impact_level,omitempty" tfsdk:"impact_level"`
	MaintenanceTemplate *string `json:"maintenance_template" tfsdk:"maintenance_template"`
	Offering            *string `json:"offering" tfsdk:"offering"`
}

type MaintenanceAnnouncementRequest struct {
	ExternalReferenceUrl *string `json:"external_reference_url,omitempty" tfsdk:"external_reference_url"`
	InternalNotes        *string `json:"internal_notes,omitempty" tfsdk:"internal_notes"`
	MaintenanceType      *int64  `json:"maintenance_type,omitempty" tfsdk:"maintenance_type"`
	Message              *string `json:"message,omitempty" tfsdk:"message"`
	Name                 *string `json:"name" tfsdk:"name"`
	ScheduledEnd         *string `json:"scheduled_end" tfsdk:"scheduled_end"`
	ScheduledStart       *string `json:"scheduled_start" tfsdk:"scheduled_start"`
	ServiceProvider      *string `json:"service_provider" tfsdk:"service_provider"`
}

type MaintenanceAnnouncementStateEnum struct {
}

type MaintenanceAnnouncementTemplate struct {
	AffectedOfferings []MaintenanceAnnouncementOffering `json:"affected_offerings" tfsdk:"affected_offerings"`
	MaintenanceType   *int64                            `json:"maintenance_type,omitempty" tfsdk:"maintenance_type"`
	Message           *string                           `json:"message,omitempty" tfsdk:"message"`
	Name              *string                           `json:"name" tfsdk:"name"`
	ServiceProvider   *string                           `json:"service_provider" tfsdk:"service_provider"`
	Url               *string                           `json:"url" tfsdk:"url"`
}

type MaintenanceAnnouncementTemplateRequest struct {
	MaintenanceType *int64  `json:"maintenance_type,omitempty" tfsdk:"maintenance_type"`
	Message         *string `json:"message,omitempty" tfsdk:"message"`
	Name            *string `json:"name" tfsdk:"name"`
	ServiceProvider *string `json:"service_provider" tfsdk:"service_provider"`
}

type MaintenanceTypeEnum struct {
}

type ManagedProject struct {
	Created            *string `json:"created" tfsdk:"created"`
	Destination        *string `json:"destination" tfsdk:"destination"`
	Identifier         *string `json:"identifier" tfsdk:"identifier"`
	LocalIdentifier    *string `json:"local_identifier,omitempty" tfsdk:"local_identifier"`
	Project            *string `json:"project" tfsdk:"project"`
	ProjectTemplate    *string `json:"project_template" tfsdk:"project_template"`
	ReviewComment      *string `json:"review_comment,omitempty" tfsdk:"review_comment"`
	ReviewedAt         *string `json:"reviewed_at" tfsdk:"reviewed_at"`
	ReviewedByFullName *string `json:"reviewed_by_full_name" tfsdk:"reviewed_by_full_name"`
	ReviewedByUuid     *string `json:"reviewed_by_uuid" tfsdk:"reviewed_by_uuid"`
	State              *string `json:"state" tfsdk:"state"`
}

type ManagedRancherCreateNodeRequest struct {
	Cpu              *int64              `json:"cpu,omitempty" tfsdk:"cpu"`
	DataVolumes      []DataVolumeRequest `json:"data_volumes,omitempty" tfsdk:"data_volumes"`
	Flavor           *string             `json:"flavor,omitempty" tfsdk:"flavor"`
	Memory           *int64              `json:"memory,omitempty" tfsdk:"memory"`
	Role             *string             `json:"role" tfsdk:"role"`
	SshPublicKey     *string             `json:"ssh_public_key,omitempty" tfsdk:"ssh_public_key"`
	Subnet           *string             `json:"subnet" tfsdk:"subnet"`
	SystemVolumeSize *int64              `json:"system_volume_size,omitempty" tfsdk:"system_volume_size"`
	SystemVolumeType *string             `json:"system_volume_type,omitempty" tfsdk:"system_volume_type"`
	Tenant           *string             `json:"tenant,omitempty" tfsdk:"tenant"`
}

type Mapping struct {
	Networks             []string            `json:"networks,omitempty" tfsdk:"networks"`
	SkipConnectionExtnet *bool               `json:"skip_connection_extnet,omitempty" tfsdk:"skip_connection_extnet"`
	Subnets              []SubNetMapping     `json:"subnets,omitempty" tfsdk:"subnets"`
	SyncInstancePorts    *bool               `json:"sync_instance_ports,omitempty" tfsdk:"sync_instance_ports"`
	VolumeTypes          []VolumeTypeMapping `json:"volume_types,omitempty" tfsdk:"volume_types"`
}

type MappingRequest struct {
	Networks             []string                   `json:"networks,omitempty" tfsdk:"networks"`
	SkipConnectionExtnet *bool                      `json:"skip_connection_extnet,omitempty" tfsdk:"skip_connection_extnet"`
	Subnets              []SubNetMappingRequest     `json:"subnets,omitempty" tfsdk:"subnets"`
	SyncInstancePorts    *bool                      `json:"sync_instance_ports,omitempty" tfsdk:"sync_instance_ports"`
	VolumeTypes          []VolumeTypeMappingRequest `json:"volume_types,omitempty" tfsdk:"volume_types"`
}

type MarketplaceCategory struct {
	Articles                []CategoryHelpArticle `json:"articles,omitempty" tfsdk:"articles"`
	AvailableOfferingsCount *int64                `json:"available_offerings_count,omitempty" tfsdk:"available_offerings_count"`
	Columns                 []NestedColumn        `json:"columns,omitempty" tfsdk:"columns"`
	Components              []CategoryComponent   `json:"components,omitempty" tfsdk:"components"`
	DefaultTenantCategory   *bool                 `json:"default_tenant_category,omitempty" tfsdk:"default_tenant_category"`
	DefaultVmCategory       *bool                 `json:"default_vm_category,omitempty" tfsdk:"default_vm_category"`
	DefaultVolumeCategory   *bool                 `json:"default_volume_category,omitempty" tfsdk:"default_volume_category"`
	Description             *string               `json:"description,omitempty" tfsdk:"description"`
	Group                   *string               `json:"group,omitempty" tfsdk:"group"`
	Icon                    *string               `json:"icon,omitempty" tfsdk:"icon"`
	OfferingCount           *int64                `json:"offering_count,omitempty" tfsdk:"offering_count"`
	Sections                []NestedSection       `json:"sections,omitempty" tfsdk:"sections"`
	Title                   *string               `json:"title,omitempty" tfsdk:"title"`
	Url                     *string               `json:"url,omitempty" tfsdk:"url"`
}

type MarketplaceCategoryRequest struct {
	DefaultTenantCategory *bool   `json:"default_tenant_category,omitempty" tfsdk:"default_tenant_category"`
	DefaultVmCategory     *bool   `json:"default_vm_category,omitempty" tfsdk:"default_vm_category"`
	DefaultVolumeCategory *bool   `json:"default_volume_category,omitempty" tfsdk:"default_volume_category"`
	Description           *string `json:"description,omitempty" tfsdk:"description"`
	Group                 *string `json:"group,omitempty" tfsdk:"group"`
	Icon                  *string `json:"icon,omitempty" tfsdk:"icon"`
	Title                 *string `json:"title" tfsdk:"title"`
}

type MarketplaceCategoryRequestForm struct {
	DefaultTenantCategory *bool   `json:"default_tenant_category,omitempty" tfsdk:"default_tenant_category"`
	DefaultVmCategory     *bool   `json:"default_vm_category,omitempty" tfsdk:"default_vm_category"`
	DefaultVolumeCategory *bool   `json:"default_volume_category,omitempty" tfsdk:"default_volume_category"`
	Description           *string `json:"description,omitempty" tfsdk:"description"`
	Group                 *string `json:"group,omitempty" tfsdk:"group"`
	Icon                  *string `json:"icon,omitempty" tfsdk:"icon"`
	Title                 *string `json:"title" tfsdk:"title"`
}

type MarketplaceCategoryRequestMultipart struct {
	DefaultTenantCategory *bool   `json:"default_tenant_category,omitempty" tfsdk:"default_tenant_category"`
	DefaultVmCategory     *bool   `json:"default_vm_category,omitempty" tfsdk:"default_vm_category"`
	DefaultVolumeCategory *bool   `json:"default_volume_category,omitempty" tfsdk:"default_volume_category"`
	Description           *string `json:"description,omitempty" tfsdk:"description"`
	Group                 *string `json:"group,omitempty" tfsdk:"group"`
	Icon                  *string `json:"icon,omitempty" tfsdk:"icon"`
	Title                 *string `json:"title" tfsdk:"title"`
}

type MarketplaceCustomerStats struct {
	Abbreviation *string `json:"abbreviation" tfsdk:"abbreviation"`
	Count        *int64  `json:"count" tfsdk:"count"`
	Name         *string `json:"name" tfsdk:"name"`
}

type MarketplaceOpenPortalCreateOrderAttributes struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
}

type MarketplaceOpenPortalRemoteCreateOrderAttributes struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
}

type MarketplaceProviderCustomer struct {
	Abbreviation    *string           `json:"abbreviation,omitempty" tfsdk:"abbreviation"`
	Email           *string           `json:"email,omitempty" tfsdk:"email"`
	Name            *string           `json:"name,omitempty" tfsdk:"name"`
	PaymentProfiles []PaymentProfile  `json:"payment_profiles,omitempty" tfsdk:"payment_profiles"`
	PhoneNumber     *string           `json:"phone_number,omitempty" tfsdk:"phone_number"`
	Projects        []ProviderProject `json:"projects,omitempty" tfsdk:"projects"`
	ProjectsCount   *int64            `json:"projects_count,omitempty" tfsdk:"projects_count"`
	Slug            *string           `json:"slug,omitempty" tfsdk:"slug"`
	Users           []ProviderUser    `json:"users,omitempty" tfsdk:"users"`
	UsersCount      *int64            `json:"users_count,omitempty" tfsdk:"users_count"`
}

type MarketplaceProviderCustomerProject struct {
	Description    *string `json:"description,omitempty" tfsdk:"description"`
	EndDate        *string `json:"end_date,omitempty" tfsdk:"end_date"`
	Name           *string `json:"name,omitempty" tfsdk:"name"`
	ResourcesCount *int64  `json:"resources_count,omitempty" tfsdk:"resources_count"`
	UsersCount     *int64  `json:"users_count,omitempty" tfsdk:"users_count"`
}

type MarketplaceServiceProviderUser struct {
	Email              *string `json:"email,omitempty" tfsdk:"email"`
	FirstName          *string `json:"first_name,omitempty" tfsdk:"first_name"`
	FullName           *string `json:"full_name,omitempty" tfsdk:"full_name"`
	IsActive           *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	LastName           *string `json:"last_name,omitempty" tfsdk:"last_name"`
	Organization       *string `json:"organization,omitempty" tfsdk:"organization"`
	PhoneNumber        *string `json:"phone_number,omitempty" tfsdk:"phone_number"`
	ProjectsCount      *int64  `json:"projects_count,omitempty" tfsdk:"projects_count"`
	RegistrationMethod *string `json:"registration_method,omitempty" tfsdk:"registration_method"`
	Username           *string `json:"username,omitempty" tfsdk:"username"`
}

type MergedPluginOptions struct {
	AutoApproveInServiceProviderProjects           *bool    `json:"auto_approve_in_service_provider_projects,omitempty" tfsdk:"auto_approve_in_service_provider_projects"`
	AutoApproveMarketplaceScript                   *bool    `json:"auto_approve_marketplace_script,omitempty" tfsdk:"auto_approve_marketplace_script"`
	AutoApproveRemoteOrders                        *bool    `json:"auto_approve_remote_orders,omitempty" tfsdk:"auto_approve_remote_orders"`
	BackendIdDisplayLabel                          *string  `json:"backend_id_display_label,omitempty" tfsdk:"backend_id_display_label"`
	CanRestoreResource                             *bool    `json:"can_restore_resource,omitempty" tfsdk:"can_restore_resource"`
	ConcealBillingData                             *bool    `json:"conceal_billing_data,omitempty" tfsdk:"conceal_billing_data"`
	CreateOrdersOnResourceOptionChange             *bool    `json:"create_orders_on_resource_option_change,omitempty" tfsdk:"create_orders_on_resource_option_change"`
	DefaultInternalNetworkMtu                      *int64   `json:"default_internal_network_mtu,omitempty" tfsdk:"default_internal_network_mtu"`
	DefaultResourceTerminationOffsetInDays         *int64   `json:"default_resource_termination_offset_in_days,omitempty" tfsdk:"default_resource_termination_offset_in_days"`
	DeploymentMode                                 *string  `json:"deployment_mode,omitempty" tfsdk:"deployment_mode"`
	DisableAutoapprove                             *bool    `json:"disable_autoapprove,omitempty" tfsdk:"disable_autoapprove"`
	EnableDisplayOfOrderActionsForServiceProvider  *bool    `json:"enable_display_of_order_actions_for_service_provider,omitempty" tfsdk:"enable_display_of_order_actions_for_service_provider"`
	EnableIssuesForMembershipChanges               *bool    `json:"enable_issues_for_membership_changes,omitempty" tfsdk:"enable_issues_for_membership_changes"`
	EnablePurchaseOrderUpload                      *bool    `json:"enable_purchase_order_upload,omitempty" tfsdk:"enable_purchase_order_upload"`
	FlavorsRegex                                   *string  `json:"flavors_regex,omitempty" tfsdk:"flavors_regex"`
	HeappeClusterId                                *string  `json:"heappe_cluster_id,omitempty" tfsdk:"heappe_cluster_id"`
	HeappeLocalBasePath                            *string  `json:"heappe_local_base_path,omitempty" tfsdk:"heappe_local_base_path"`
	HeappeUrl                                      *string  `json:"heappe_url,omitempty" tfsdk:"heappe_url"`
	HeappeUsername                                 *string  `json:"heappe_username,omitempty" tfsdk:"heappe_username"`
	HighlightBackendIdDisplay                      *bool    `json:"highlight_backend_id_display,omitempty" tfsdk:"highlight_backend_id_display"`
	HomedirPrefix                                  *string  `json:"homedir_prefix,omitempty" tfsdk:"homedir_prefix"`
	InitialPrimarygroupNumber                      *int64   `json:"initial_primarygroup_number,omitempty" tfsdk:"initial_primarygroup_number"`
	InitialUidnumber                               *int64   `json:"initial_uidnumber,omitempty" tfsdk:"initial_uidnumber"`
	InitialUsergroupNumber                         *int64   `json:"initial_usergroup_number,omitempty" tfsdk:"initial_usergroup_number"`
	IsResourceTerminationDateRequired              *bool    `json:"is_resource_termination_date_required,omitempty" tfsdk:"is_resource_termination_date_required"`
	LatestDateForResourceTermination               *string  `json:"latest_date_for_resource_termination,omitempty" tfsdk:"latest_date_for_resource_termination"`
	ManagedRancherLoadBalancerDataVolumeSizeGb     *int64   `json:"managed_rancher_load_balancer_data_volume_size_gb,omitempty" tfsdk:"managed_rancher_load_balancer_data_volume_size_gb"`
	ManagedRancherLoadBalancerDataVolumeTypeName   *string  `json:"managed_rancher_load_balancer_data_volume_type_name,omitempty" tfsdk:"managed_rancher_load_balancer_data_volume_type_name"`
	ManagedRancherLoadBalancerFlavorName           *string  `json:"managed_rancher_load_balancer_flavor_name,omitempty" tfsdk:"managed_rancher_load_balancer_flavor_name"`
	ManagedRancherLoadBalancerSystemVolumeSizeGb   *int64   `json:"managed_rancher_load_balancer_system_volume_size_gb,omitempty" tfsdk:"managed_rancher_load_balancer_system_volume_size_gb"`
	ManagedRancherLoadBalancerSystemVolumeTypeName *string  `json:"managed_rancher_load_balancer_system_volume_type_name,omitempty" tfsdk:"managed_rancher_load_balancer_system_volume_type_name"`
	ManagedRancherServerDataVolumeSizeGb           *int64   `json:"managed_rancher_server_data_volume_size_gb,omitempty" tfsdk:"managed_rancher_server_data_volume_size_gb"`
	ManagedRancherServerDataVolumeTypeName         *string  `json:"managed_rancher_server_data_volume_type_name,omitempty" tfsdk:"managed_rancher_server_data_volume_type_name"`
	ManagedRancherServerFlavorName                 *string  `json:"managed_rancher_server_flavor_name,omitempty" tfsdk:"managed_rancher_server_flavor_name"`
	ManagedRancherServerSystemVolumeSizeGb         *int64   `json:"managed_rancher_server_system_volume_size_gb,omitempty" tfsdk:"managed_rancher_server_system_volume_size_gb"`
	ManagedRancherServerSystemVolumeTypeName       *string  `json:"managed_rancher_server_system_volume_type_name,omitempty" tfsdk:"managed_rancher_server_system_volume_type_name"`
	ManagedRancherTenantMaxCpu                     *int64   `json:"managed_rancher_tenant_max_cpu,omitempty" tfsdk:"managed_rancher_tenant_max_cpu"`
	ManagedRancherTenantMaxDisk                    *int64   `json:"managed_rancher_tenant_max_disk,omitempty" tfsdk:"managed_rancher_tenant_max_disk"`
	ManagedRancherTenantMaxRam                     *int64   `json:"managed_rancher_tenant_max_ram,omitempty" tfsdk:"managed_rancher_tenant_max_ram"`
	ManagedRancherWorkerSystemVolumeSizeGb         *int64   `json:"managed_rancher_worker_system_volume_size_gb,omitempty" tfsdk:"managed_rancher_worker_system_volume_size_gb"`
	ManagedRancherWorkerSystemVolumeTypeName       *string  `json:"managed_rancher_worker_system_volume_type_name,omitempty" tfsdk:"managed_rancher_worker_system_volume_type_name"`
	MaxInstances                                   *int64   `json:"max_instances,omitempty" tfsdk:"max_instances"`
	MaxResourceTerminationOffsetInDays             *int64   `json:"max_resource_termination_offset_in_days,omitempty" tfsdk:"max_resource_termination_offset_in_days"`
	MaxSecurityGroups                              *int64   `json:"max_security_groups,omitempty" tfsdk:"max_security_groups"`
	MaxVolumes                                     *int64   `json:"max_volumes,omitempty" tfsdk:"max_volumes"`
	MaximalResourceCountPerProject                 *int64   `json:"maximal_resource_count_per_project,omitempty" tfsdk:"maximal_resource_count_per_project"`
	MinimalTeamCountForProvisioning                *int64   `json:"minimal_team_count_for_provisioning,omitempty" tfsdk:"minimal_team_count_for_provisioning"`
	OpenstackOfferingUuidList                      []string `json:"openstack_offering_uuid_list,omitempty" tfsdk:"openstack_offering_uuid_list"`
	ProjectPermanentDirectory                      *string  `json:"project_permanent_directory,omitempty" tfsdk:"project_permanent_directory"`
	RequirePurchaseOrderUpload                     *bool    `json:"require_purchase_order_upload,omitempty" tfsdk:"require_purchase_order_upload"`
	RequiredTeamRoleForProvisioning                *string  `json:"required_team_role_for_provisioning,omitempty" tfsdk:"required_team_role_for_provisioning"`
	ScratchProjectDirectory                        *string  `json:"scratch_project_directory,omitempty" tfsdk:"scratch_project_directory"`
	ServiceProviderCanCreateOfferingUser           *bool    `json:"service_provider_can_create_offering_user,omitempty" tfsdk:"service_provider_can_create_offering_user"`
	SnapshotSizeLimitGb                            *int64   `json:"snapshot_size_limit_gb,omitempty" tfsdk:"snapshot_size_limit_gb"`
	StorageMode                                    *string  `json:"storage_mode,omitempty" tfsdk:"storage_mode"`
	SupportsDownscaling                            *bool    `json:"supports_downscaling,omitempty" tfsdk:"supports_downscaling"`
	SupportsPausing                                *bool    `json:"supports_pausing,omitempty" tfsdk:"supports_pausing"`
	UsernameAnonymizedPrefix                       *string  `json:"username_anonymized_prefix,omitempty" tfsdk:"username_anonymized_prefix"`
	UsernameGenerationPolicy                       *string  `json:"username_generation_policy,omitempty" tfsdk:"username_generation_policy"`
}

type MergedPluginOptionsRequest struct {
	AutoApproveInServiceProviderProjects           *bool    `json:"auto_approve_in_service_provider_projects,omitempty" tfsdk:"auto_approve_in_service_provider_projects"`
	AutoApproveMarketplaceScript                   *bool    `json:"auto_approve_marketplace_script,omitempty" tfsdk:"auto_approve_marketplace_script"`
	AutoApproveRemoteOrders                        *bool    `json:"auto_approve_remote_orders,omitempty" tfsdk:"auto_approve_remote_orders"`
	BackendIdDisplayLabel                          *string  `json:"backend_id_display_label,omitempty" tfsdk:"backend_id_display_label"`
	CanRestoreResource                             *bool    `json:"can_restore_resource,omitempty" tfsdk:"can_restore_resource"`
	ConcealBillingData                             *bool    `json:"conceal_billing_data,omitempty" tfsdk:"conceal_billing_data"`
	CreateOrdersOnResourceOptionChange             *bool    `json:"create_orders_on_resource_option_change,omitempty" tfsdk:"create_orders_on_resource_option_change"`
	DefaultInternalNetworkMtu                      *int64   `json:"default_internal_network_mtu,omitempty" tfsdk:"default_internal_network_mtu"`
	DefaultResourceTerminationOffsetInDays         *int64   `json:"default_resource_termination_offset_in_days,omitempty" tfsdk:"default_resource_termination_offset_in_days"`
	DeploymentMode                                 *string  `json:"deployment_mode,omitempty" tfsdk:"deployment_mode"`
	DisableAutoapprove                             *bool    `json:"disable_autoapprove,omitempty" tfsdk:"disable_autoapprove"`
	EnableDisplayOfOrderActionsForServiceProvider  *bool    `json:"enable_display_of_order_actions_for_service_provider,omitempty" tfsdk:"enable_display_of_order_actions_for_service_provider"`
	EnableIssuesForMembershipChanges               *bool    `json:"enable_issues_for_membership_changes,omitempty" tfsdk:"enable_issues_for_membership_changes"`
	EnablePurchaseOrderUpload                      *bool    `json:"enable_purchase_order_upload,omitempty" tfsdk:"enable_purchase_order_upload"`
	FlavorsRegex                                   *string  `json:"flavors_regex,omitempty" tfsdk:"flavors_regex"`
	HeappeClusterId                                *string  `json:"heappe_cluster_id,omitempty" tfsdk:"heappe_cluster_id"`
	HeappeLocalBasePath                            *string  `json:"heappe_local_base_path,omitempty" tfsdk:"heappe_local_base_path"`
	HeappeUrl                                      *string  `json:"heappe_url,omitempty" tfsdk:"heappe_url"`
	HeappeUsername                                 *string  `json:"heappe_username,omitempty" tfsdk:"heappe_username"`
	HighlightBackendIdDisplay                      *bool    `json:"highlight_backend_id_display,omitempty" tfsdk:"highlight_backend_id_display"`
	HomedirPrefix                                  *string  `json:"homedir_prefix,omitempty" tfsdk:"homedir_prefix"`
	InitialPrimarygroupNumber                      *int64   `json:"initial_primarygroup_number,omitempty" tfsdk:"initial_primarygroup_number"`
	InitialUidnumber                               *int64   `json:"initial_uidnumber,omitempty" tfsdk:"initial_uidnumber"`
	InitialUsergroupNumber                         *int64   `json:"initial_usergroup_number,omitempty" tfsdk:"initial_usergroup_number"`
	IsResourceTerminationDateRequired              *bool    `json:"is_resource_termination_date_required,omitempty" tfsdk:"is_resource_termination_date_required"`
	LatestDateForResourceTermination               *string  `json:"latest_date_for_resource_termination,omitempty" tfsdk:"latest_date_for_resource_termination"`
	ManagedRancherLoadBalancerDataVolumeSizeGb     *int64   `json:"managed_rancher_load_balancer_data_volume_size_gb,omitempty" tfsdk:"managed_rancher_load_balancer_data_volume_size_gb"`
	ManagedRancherLoadBalancerDataVolumeTypeName   *string  `json:"managed_rancher_load_balancer_data_volume_type_name,omitempty" tfsdk:"managed_rancher_load_balancer_data_volume_type_name"`
	ManagedRancherLoadBalancerFlavorName           *string  `json:"managed_rancher_load_balancer_flavor_name,omitempty" tfsdk:"managed_rancher_load_balancer_flavor_name"`
	ManagedRancherLoadBalancerSystemVolumeSizeGb   *int64   `json:"managed_rancher_load_balancer_system_volume_size_gb,omitempty" tfsdk:"managed_rancher_load_balancer_system_volume_size_gb"`
	ManagedRancherLoadBalancerSystemVolumeTypeName *string  `json:"managed_rancher_load_balancer_system_volume_type_name,omitempty" tfsdk:"managed_rancher_load_balancer_system_volume_type_name"`
	ManagedRancherServerDataVolumeSizeGb           *int64   `json:"managed_rancher_server_data_volume_size_gb,omitempty" tfsdk:"managed_rancher_server_data_volume_size_gb"`
	ManagedRancherServerDataVolumeTypeName         *string  `json:"managed_rancher_server_data_volume_type_name,omitempty" tfsdk:"managed_rancher_server_data_volume_type_name"`
	ManagedRancherServerFlavorName                 *string  `json:"managed_rancher_server_flavor_name,omitempty" tfsdk:"managed_rancher_server_flavor_name"`
	ManagedRancherServerSystemVolumeSizeGb         *int64   `json:"managed_rancher_server_system_volume_size_gb,omitempty" tfsdk:"managed_rancher_server_system_volume_size_gb"`
	ManagedRancherServerSystemVolumeTypeName       *string  `json:"managed_rancher_server_system_volume_type_name,omitempty" tfsdk:"managed_rancher_server_system_volume_type_name"`
	ManagedRancherTenantMaxCpu                     *int64   `json:"managed_rancher_tenant_max_cpu,omitempty" tfsdk:"managed_rancher_tenant_max_cpu"`
	ManagedRancherTenantMaxDisk                    *int64   `json:"managed_rancher_tenant_max_disk,omitempty" tfsdk:"managed_rancher_tenant_max_disk"`
	ManagedRancherTenantMaxRam                     *int64   `json:"managed_rancher_tenant_max_ram,omitempty" tfsdk:"managed_rancher_tenant_max_ram"`
	ManagedRancherWorkerSystemVolumeSizeGb         *int64   `json:"managed_rancher_worker_system_volume_size_gb,omitempty" tfsdk:"managed_rancher_worker_system_volume_size_gb"`
	ManagedRancherWorkerSystemVolumeTypeName       *string  `json:"managed_rancher_worker_system_volume_type_name,omitempty" tfsdk:"managed_rancher_worker_system_volume_type_name"`
	MaxInstances                                   *int64   `json:"max_instances,omitempty" tfsdk:"max_instances"`
	MaxResourceTerminationOffsetInDays             *int64   `json:"max_resource_termination_offset_in_days,omitempty" tfsdk:"max_resource_termination_offset_in_days"`
	MaxSecurityGroups                              *int64   `json:"max_security_groups,omitempty" tfsdk:"max_security_groups"`
	MaxVolumes                                     *int64   `json:"max_volumes,omitempty" tfsdk:"max_volumes"`
	MaximalResourceCountPerProject                 *int64   `json:"maximal_resource_count_per_project,omitempty" tfsdk:"maximal_resource_count_per_project"`
	MinimalTeamCountForProvisioning                *int64   `json:"minimal_team_count_for_provisioning,omitempty" tfsdk:"minimal_team_count_for_provisioning"`
	OpenstackOfferingUuidList                      []string `json:"openstack_offering_uuid_list,omitempty" tfsdk:"openstack_offering_uuid_list"`
	ProjectPermanentDirectory                      *string  `json:"project_permanent_directory,omitempty" tfsdk:"project_permanent_directory"`
	RequirePurchaseOrderUpload                     *bool    `json:"require_purchase_order_upload,omitempty" tfsdk:"require_purchase_order_upload"`
	RequiredTeamRoleForProvisioning                *string  `json:"required_team_role_for_provisioning,omitempty" tfsdk:"required_team_role_for_provisioning"`
	ScratchProjectDirectory                        *string  `json:"scratch_project_directory,omitempty" tfsdk:"scratch_project_directory"`
	ServiceProviderCanCreateOfferingUser           *bool    `json:"service_provider_can_create_offering_user,omitempty" tfsdk:"service_provider_can_create_offering_user"`
	SnapshotSizeLimitGb                            *int64   `json:"snapshot_size_limit_gb,omitempty" tfsdk:"snapshot_size_limit_gb"`
	StorageMode                                    *string  `json:"storage_mode,omitempty" tfsdk:"storage_mode"`
	SupportsDownscaling                            *bool    `json:"supports_downscaling,omitempty" tfsdk:"supports_downscaling"`
	SupportsPausing                                *bool    `json:"supports_pausing,omitempty" tfsdk:"supports_pausing"`
	UsernameAnonymizedPrefix                       *string  `json:"username_anonymized_prefix,omitempty" tfsdk:"username_anonymized_prefix"`
	UsernameGenerationPolicy                       *string  `json:"username_generation_policy,omitempty" tfsdk:"username_generation_policy"`
}

type MergedSecretOptions struct {
	ApiUrl                                      *string     `json:"api_url,omitempty" tfsdk:"api_url"`
	ArgocdK8sKubeconfig                         *string     `json:"argocd_k8s_kubeconfig,omitempty" tfsdk:"argocd_k8s_kubeconfig"`
	ArgocdK8sNamespace                          *string     `json:"argocd_k8s_namespace,omitempty" tfsdk:"argocd_k8s_namespace"`
	BackendUrl                                  *string     `json:"backend_url,omitempty" tfsdk:"backend_url"`
	BaseImageName                               *string     `json:"base_image_name,omitempty" tfsdk:"base_image_name"`
	CloudInitTemplate                           *string     `json:"cloud_init_template,omitempty" tfsdk:"cloud_init_template"`
	Create                                      *string     `json:"create,omitempty" tfsdk:"create"`
	CustomerUuid                                *string     `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	DnsNameservers                              []string    `json:"dns_nameservers,omitempty" tfsdk:"dns_nameservers"`
	HeappeClusterPassword                       *string     `json:"heappe_cluster_password,omitempty" tfsdk:"heappe_cluster_password"`
	HeappePassword                              *string     `json:"heappe_password,omitempty" tfsdk:"heappe_password"`
	Ipv4ExternalIpMapping                       []IPMapping `json:"ipv4_external_ip_mapping,omitempty" tfsdk:"ipv4_external_ip_mapping"`
	K8sVersion                                  *string     `json:"k8s_version,omitempty" tfsdk:"k8s_version"`
	KeycloakPassword                            *string     `json:"keycloak_password,omitempty" tfsdk:"keycloak_password"`
	KeycloakRealm                               *string     `json:"keycloak_realm,omitempty" tfsdk:"keycloak_realm"`
	KeycloakSslVerify                           *bool       `json:"keycloak_ssl_verify,omitempty" tfsdk:"keycloak_ssl_verify"`
	KeycloakSyncFrequency                       *int64      `json:"keycloak_sync_frequency,omitempty" tfsdk:"keycloak_sync_frequency"`
	KeycloakUrl                                 *string     `json:"keycloak_url,omitempty" tfsdk:"keycloak_url"`
	KeycloakUserRealm                           *string     `json:"keycloak_user_realm,omitempty" tfsdk:"keycloak_user_realm"`
	KeycloakUsername                            *string     `json:"keycloak_username,omitempty" tfsdk:"keycloak_username"`
	Language                                    *string     `json:"language,omitempty" tfsdk:"language"`
	ManagedRancherLoadBalancerCloudInitTemplate *string     `json:"managed_rancher_load_balancer_cloud_init_template,omitempty" tfsdk:"managed_rancher_load_balancer_cloud_init_template"`
	NodeDiskDriver                              *string     `json:"node_disk_driver,omitempty" tfsdk:"node_disk_driver"`
	OpenstackApiTlsCertificate                  *string     `json:"openstack_api_tls_certificate,omitempty" tfsdk:"openstack_api_tls_certificate"`
	Password                                    *string     `json:"password,omitempty" tfsdk:"password"`
	PrivateRegistryPassword                     *string     `json:"private_registry_password,omitempty" tfsdk:"private_registry_password"`
	PrivateRegistryUrl                          *string     `json:"private_registry_url,omitempty" tfsdk:"private_registry_url"`
	PrivateRegistryUser                         *string     `json:"private_registry_user,omitempty" tfsdk:"private_registry_user"`
	Pull                                        *string     `json:"pull,omitempty" tfsdk:"pull"`
	SharedUserPassword                          *string     `json:"shared_user_password,omitempty" tfsdk:"shared_user_password"`
	TemplateConfirmationComment                 *string     `json:"template_confirmation_comment,omitempty" tfsdk:"template_confirmation_comment"`
	Terminate                                   *string     `json:"terminate,omitempty" tfsdk:"terminate"`
	Token                                       *string     `json:"token,omitempty" tfsdk:"token"`
	Update                                      *string     `json:"update,omitempty" tfsdk:"update"`
	Username                                    *string     `json:"username,omitempty" tfsdk:"username"`
	VaultHost                                   *string     `json:"vault_host,omitempty" tfsdk:"vault_host"`
	VaultPort                                   *int64      `json:"vault_port,omitempty" tfsdk:"vault_port"`
	VaultTlsVerify                              *bool       `json:"vault_tls_verify,omitempty" tfsdk:"vault_tls_verify"`
	VaultToken                                  *string     `json:"vault_token,omitempty" tfsdk:"vault_token"`
}

type MergedSecretOptionsRequest struct {
	ApiUrl                                      *string            `json:"api_url,omitempty" tfsdk:"api_url"`
	ArgocdK8sKubeconfig                         *string            `json:"argocd_k8s_kubeconfig,omitempty" tfsdk:"argocd_k8s_kubeconfig"`
	ArgocdK8sNamespace                          *string            `json:"argocd_k8s_namespace,omitempty" tfsdk:"argocd_k8s_namespace"`
	BackendUrl                                  *string            `json:"backend_url,omitempty" tfsdk:"backend_url"`
	BaseImageName                               *string            `json:"base_image_name,omitempty" tfsdk:"base_image_name"`
	CloudInitTemplate                           *string            `json:"cloud_init_template,omitempty" tfsdk:"cloud_init_template"`
	Create                                      *string            `json:"create,omitempty" tfsdk:"create"`
	CustomerUuid                                *string            `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	DnsNameservers                              []string           `json:"dns_nameservers,omitempty" tfsdk:"dns_nameservers"`
	HeappeClusterPassword                       *string            `json:"heappe_cluster_password,omitempty" tfsdk:"heappe_cluster_password"`
	HeappePassword                              *string            `json:"heappe_password,omitempty" tfsdk:"heappe_password"`
	Ipv4ExternalIpMapping                       []IPMappingRequest `json:"ipv4_external_ip_mapping,omitempty" tfsdk:"ipv4_external_ip_mapping"`
	K8sVersion                                  *string            `json:"k8s_version,omitempty" tfsdk:"k8s_version"`
	KeycloakPassword                            *string            `json:"keycloak_password,omitempty" tfsdk:"keycloak_password"`
	KeycloakRealm                               *string            `json:"keycloak_realm,omitempty" tfsdk:"keycloak_realm"`
	KeycloakSslVerify                           *bool              `json:"keycloak_ssl_verify,omitempty" tfsdk:"keycloak_ssl_verify"`
	KeycloakSyncFrequency                       *int64             `json:"keycloak_sync_frequency,omitempty" tfsdk:"keycloak_sync_frequency"`
	KeycloakUrl                                 *string            `json:"keycloak_url,omitempty" tfsdk:"keycloak_url"`
	KeycloakUserRealm                           *string            `json:"keycloak_user_realm,omitempty" tfsdk:"keycloak_user_realm"`
	KeycloakUsername                            *string            `json:"keycloak_username,omitempty" tfsdk:"keycloak_username"`
	Language                                    *string            `json:"language,omitempty" tfsdk:"language"`
	ManagedRancherLoadBalancerCloudInitTemplate *string            `json:"managed_rancher_load_balancer_cloud_init_template,omitempty" tfsdk:"managed_rancher_load_balancer_cloud_init_template"`
	NodeDiskDriver                              *string            `json:"node_disk_driver,omitempty" tfsdk:"node_disk_driver"`
	OpenstackApiTlsCertificate                  *string            `json:"openstack_api_tls_certificate,omitempty" tfsdk:"openstack_api_tls_certificate"`
	Password                                    *string            `json:"password,omitempty" tfsdk:"password"`
	PrivateRegistryPassword                     *string            `json:"private_registry_password,omitempty" tfsdk:"private_registry_password"`
	PrivateRegistryUrl                          *string            `json:"private_registry_url,omitempty" tfsdk:"private_registry_url"`
	PrivateRegistryUser                         *string            `json:"private_registry_user,omitempty" tfsdk:"private_registry_user"`
	Pull                                        *string            `json:"pull,omitempty" tfsdk:"pull"`
	SharedUserPassword                          *string            `json:"shared_user_password,omitempty" tfsdk:"shared_user_password"`
	TemplateConfirmationComment                 *string            `json:"template_confirmation_comment,omitempty" tfsdk:"template_confirmation_comment"`
	Terminate                                   *string            `json:"terminate,omitempty" tfsdk:"terminate"`
	Token                                       *string            `json:"token,omitempty" tfsdk:"token"`
	Update                                      *string            `json:"update,omitempty" tfsdk:"update"`
	Username                                    *string            `json:"username,omitempty" tfsdk:"username"`
	VaultHost                                   *string            `json:"vault_host,omitempty" tfsdk:"vault_host"`
	VaultPort                                   *int64             `json:"vault_port,omitempty" tfsdk:"vault_port"`
	VaultTlsVerify                              *bool              `json:"vault_tls_verify,omitempty" tfsdk:"vault_tls_verify"`
	VaultToken                                  *string            `json:"vault_token,omitempty" tfsdk:"vault_token"`
}

type MessageTemplate struct {
	Body    *string `json:"body" tfsdk:"body"`
	Name    *string `json:"name" tfsdk:"name"`
	Subject *string `json:"subject" tfsdk:"subject"`
	Url     *string `json:"url" tfsdk:"url"`
}

type MessageTemplateRequest struct {
	Body    *string `json:"body" tfsdk:"body"`
	Name    *string `json:"name" tfsdk:"name"`
	Subject *string `json:"subject" tfsdk:"subject"`
}

type MigrationCreate struct {
	Mappings    *Mapping `json:"mappings,omitempty" tfsdk:"mappings"`
	SrcResource *string  `json:"src_resource" tfsdk:"src_resource"`
}

type MigrationCreateRequest struct {
	Description *string         `json:"description,omitempty" tfsdk:"description"`
	DstOffering *string         `json:"dst_offering" tfsdk:"dst_offering"`
	DstPlan     *string         `json:"dst_plan" tfsdk:"dst_plan"`
	Mappings    *MappingRequest `json:"mappings,omitempty" tfsdk:"mappings"`
	Name        *string         `json:"name,omitempty" tfsdk:"name"`
	SrcResource *string         `json:"src_resource" tfsdk:"src_resource"`
}

type MigrationDetails struct {
	Created           *string  `json:"created" tfsdk:"created"`
	CreatedByFullName *string  `json:"created_by_full_name" tfsdk:"created_by_full_name"`
	CreatedByUuid     *string  `json:"created_by_uuid" tfsdk:"created_by_uuid"`
	DstOfferingName   *string  `json:"dst_offering_name" tfsdk:"dst_offering_name"`
	DstOfferingUuid   *string  `json:"dst_offering_uuid" tfsdk:"dst_offering_uuid"`
	DstResourceName   *string  `json:"dst_resource_name" tfsdk:"dst_resource_name"`
	DstResourceState  *string  `json:"dst_resource_state" tfsdk:"dst_resource_state"`
	DstResourceUuid   *string  `json:"dst_resource_uuid" tfsdk:"dst_resource_uuid"`
	ErrorMessage      *string  `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback    *string  `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	Mappings          *Mapping `json:"mappings" tfsdk:"mappings"`
	Modified          *string  `json:"modified" tfsdk:"modified"`
	SrcOfferingName   *string  `json:"src_offering_name" tfsdk:"src_offering_name"`
	SrcOfferingUuid   *string  `json:"src_offering_uuid" tfsdk:"src_offering_uuid"`
	SrcResourceName   *string  `json:"src_resource_name" tfsdk:"src_resource_name"`
	SrcResourceUuid   *string  `json:"src_resource_uuid" tfsdk:"src_resource_uuid"`
	State             *string  `json:"state" tfsdk:"state"`
}

type MigrationDetailsRequest struct {
	ErrorMessage   *string         `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback *string         `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	Mappings       *MappingRequest `json:"mappings" tfsdk:"mappings"`
}

type MinimalConsumptionLogicEnum struct {
}

type MoveOfferingRequest struct {
	Customer            *string `json:"customer" tfsdk:"customer"`
	PreservePermissions *bool   `json:"preserve_permissions" tfsdk:"preserve_permissions"`
}

type MoveProjectRequest struct {
	Customer            *string `json:"customer" tfsdk:"customer"`
	PreservePermissions *bool   `json:"preserve_permissions" tfsdk:"preserve_permissions"`
}

type MoveResourceRequest struct {
}

type NameUUID struct {
	Name *string `json:"name" tfsdk:"name"`
}

type NestedAgentProcessor struct {
	BackendType    *string `json:"backend_type" tfsdk:"backend_type"`
	BackendVersion *string `json:"backend_version,omitempty" tfsdk:"backend_version"`
	Created        *string `json:"created" tfsdk:"created"`
	LastRun        *string `json:"last_run,omitempty" tfsdk:"last_run"`
	Modified       *string `json:"modified" tfsdk:"modified"`
	Name           *string `json:"name" tfsdk:"name"`
	Url            *string `json:"url" tfsdk:"url"`
}

type NestedAgentService struct {
	Created  *string `json:"created" tfsdk:"created"`
	Mode     *string `json:"mode,omitempty" tfsdk:"mode"`
	Modified *string `json:"modified" tfsdk:"modified"`
	Name     *string `json:"name" tfsdk:"name"`
	State    *string `json:"state" tfsdk:"state"`
	Url      *string `json:"url" tfsdk:"url"`
}

type NestedAgentServiceRequest struct {
	Mode *string `json:"mode,omitempty" tfsdk:"mode"`
	Name *string `json:"name" tfsdk:"name"`
}

type NestedAttribute struct {
	Key      *string                 `json:"key,omitempty" tfsdk:"key"`
	Options  []NestedAttributeOption `json:"options,omitempty" tfsdk:"options"`
	Required *bool                   `json:"required,omitempty" tfsdk:"required"`
	Title    *string                 `json:"title,omitempty" tfsdk:"title"`
	Type     *string                 `json:"type,omitempty" tfsdk:"type"`
}

type NestedAttributeOption struct {
	Key   *string `json:"key,omitempty" tfsdk:"key"`
	Title *string `json:"title,omitempty" tfsdk:"title"`
}

type NestedAttributeOptionRequest struct {
	Key   *string `json:"key" tfsdk:"key"`
	Title *string `json:"title" tfsdk:"title"`
}

type NestedAttributeRequest struct {
	Key      *string                        `json:"key" tfsdk:"key"`
	Options  []NestedAttributeOptionRequest `json:"options" tfsdk:"options"`
	Required *bool                          `json:"required,omitempty" tfsdk:"required"`
	Title    *string                        `json:"title" tfsdk:"title"`
	Type     *string                        `json:"type" tfsdk:"type"`
}

type NestedAttributeTypeEnum struct {
}

type NestedCampaign struct {
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	Discount        *int64  `json:"discount,omitempty" tfsdk:"discount"`
	DiscountType    *string `json:"discount_type,omitempty" tfsdk:"discount_type"`
	EndDate         *string `json:"end_date,omitempty" tfsdk:"end_date"`
	Months          *int64  `json:"months,omitempty" tfsdk:"months"`
	Name            *string `json:"name,omitempty" tfsdk:"name"`
	ServiceProvider *string `json:"service_provider,omitempty" tfsdk:"service_provider"`
	StartDate       *string `json:"start_date,omitempty" tfsdk:"start_date"`
	Stock           *int64  `json:"stock,omitempty" tfsdk:"stock"`
}

type NestedColumn struct {
	Attribute *string `json:"attribute,omitempty" tfsdk:"attribute"`
	Index     *int64  `json:"index,omitempty" tfsdk:"index"`
	Title     *string `json:"title,omitempty" tfsdk:"title"`
	Widget    *string `json:"widget,omitempty" tfsdk:"widget"`
}

type NestedColumnRequest struct {
	Attribute *string `json:"attribute,omitempty" tfsdk:"attribute"`
	Index     *int64  `json:"index" tfsdk:"index"`
	Title     *string `json:"title" tfsdk:"title"`
	Widget    *string `json:"widget,omitempty" tfsdk:"widget"`
}

type NestedCustomerUsagePolicyComponent struct {
	Component  *string `json:"component" tfsdk:"component"`
	Limit      *int64  `json:"limit" tfsdk:"limit"`
	Period     *int64  `json:"period,omitempty" tfsdk:"period"`
	PeriodName *string `json:"period_name" tfsdk:"period_name"`
	Type       *string `json:"type" tfsdk:"type"`
}

type NestedCustomerUsagePolicyComponentRequest struct {
	Component *string `json:"component" tfsdk:"component"`
	Limit     *int64  `json:"limit" tfsdk:"limit"`
	Period    *int64  `json:"period,omitempty" tfsdk:"period"`
}

type NestedEndpoint struct {
	Name *string `json:"name,omitempty" tfsdk:"name"`
	Url  *string `json:"url,omitempty" tfsdk:"url"`
}

type NestedEndpointRequest struct {
	Name *string `json:"name" tfsdk:"name"`
	Url  *string `json:"url" tfsdk:"url"`
}

type NestedFeedback struct {
	Comment          *string `json:"comment,omitempty" tfsdk:"comment"`
	Evaluation       *int64  `json:"evaluation" tfsdk:"evaluation"`
	EvaluationNumber *int64  `json:"evaluation_number" tfsdk:"evaluation_number"`
	State            *string `json:"state" tfsdk:"state"`
}

type NestedFeedbackRequest struct {
	Comment *string `json:"comment,omitempty" tfsdk:"comment"`
}

type NestedOfferingComponentLimit struct {
	Limit *int64  `json:"limit" tfsdk:"limit"`
	Type  *string `json:"type" tfsdk:"type"`
}

type NestedOfferingComponentLimitRequest struct {
	Limit *int64  `json:"limit" tfsdk:"limit"`
	Type  *string `json:"type" tfsdk:"type"`
}

type NestedOfferingFile struct {
	Created *string `json:"created,omitempty" tfsdk:"created"`
	File    *string `json:"file,omitempty" tfsdk:"file"`
	Name    *string `json:"name,omitempty" tfsdk:"name"`
}

type NestedOfferingFileRequest struct {
	File *string `json:"file" tfsdk:"file"`
	Name *string `json:"name" tfsdk:"name"`
}

type NestedPartition struct {
	CpuBind          *int64  `json:"cpu_bind,omitempty" tfsdk:"cpu_bind"`
	DefCpuPerGpu     *int64  `json:"def_cpu_per_gpu,omitempty" tfsdk:"def_cpu_per_gpu"`
	DefMemPerCpu     *int64  `json:"def_mem_per_cpu,omitempty" tfsdk:"def_mem_per_cpu"`
	DefMemPerGpu     *int64  `json:"def_mem_per_gpu,omitempty" tfsdk:"def_mem_per_gpu"`
	DefMemPerNode    *int64  `json:"def_mem_per_node,omitempty" tfsdk:"def_mem_per_node"`
	DefaultTime      *int64  `json:"default_time,omitempty" tfsdk:"default_time"`
	ExclusiveTopo    *bool   `json:"exclusive_topo,omitempty" tfsdk:"exclusive_topo"`
	ExclusiveUser    *bool   `json:"exclusive_user,omitempty" tfsdk:"exclusive_user"`
	GraceTime        *int64  `json:"grace_time,omitempty" tfsdk:"grace_time"`
	MaxCpusPerNode   *int64  `json:"max_cpus_per_node,omitempty" tfsdk:"max_cpus_per_node"`
	MaxCpusPerSocket *int64  `json:"max_cpus_per_socket,omitempty" tfsdk:"max_cpus_per_socket"`
	MaxMemPerCpu     *int64  `json:"max_mem_per_cpu,omitempty" tfsdk:"max_mem_per_cpu"`
	MaxMemPerNode    *int64  `json:"max_mem_per_node,omitempty" tfsdk:"max_mem_per_node"`
	MaxNodes         *int64  `json:"max_nodes,omitempty" tfsdk:"max_nodes"`
	MaxTime          *int64  `json:"max_time,omitempty" tfsdk:"max_time"`
	MinNodes         *int64  `json:"min_nodes,omitempty" tfsdk:"min_nodes"`
	PartitionName    *string `json:"partition_name,omitempty" tfsdk:"partition_name"`
	PriorityTier     *int64  `json:"priority_tier,omitempty" tfsdk:"priority_tier"`
	Qos              *string `json:"qos,omitempty" tfsdk:"qos"`
	ReqResv          *bool   `json:"req_resv,omitempty" tfsdk:"req_resv"`
}

type NestedPartitionRequest struct {
	CpuBind          *int64  `json:"cpu_bind,omitempty" tfsdk:"cpu_bind"`
	DefCpuPerGpu     *int64  `json:"def_cpu_per_gpu,omitempty" tfsdk:"def_cpu_per_gpu"`
	DefMemPerCpu     *int64  `json:"def_mem_per_cpu,omitempty" tfsdk:"def_mem_per_cpu"`
	DefMemPerGpu     *int64  `json:"def_mem_per_gpu,omitempty" tfsdk:"def_mem_per_gpu"`
	DefMemPerNode    *int64  `json:"def_mem_per_node,omitempty" tfsdk:"def_mem_per_node"`
	DefaultTime      *int64  `json:"default_time,omitempty" tfsdk:"default_time"`
	ExclusiveTopo    *bool   `json:"exclusive_topo,omitempty" tfsdk:"exclusive_topo"`
	ExclusiveUser    *bool   `json:"exclusive_user,omitempty" tfsdk:"exclusive_user"`
	GraceTime        *int64  `json:"grace_time,omitempty" tfsdk:"grace_time"`
	MaxCpusPerNode   *int64  `json:"max_cpus_per_node,omitempty" tfsdk:"max_cpus_per_node"`
	MaxCpusPerSocket *int64  `json:"max_cpus_per_socket,omitempty" tfsdk:"max_cpus_per_socket"`
	MaxMemPerCpu     *int64  `json:"max_mem_per_cpu,omitempty" tfsdk:"max_mem_per_cpu"`
	MaxMemPerNode    *int64  `json:"max_mem_per_node,omitempty" tfsdk:"max_mem_per_node"`
	MaxNodes         *int64  `json:"max_nodes,omitempty" tfsdk:"max_nodes"`
	MaxTime          *int64  `json:"max_time,omitempty" tfsdk:"max_time"`
	MinNodes         *int64  `json:"min_nodes,omitempty" tfsdk:"min_nodes"`
	PartitionName    *string `json:"partition_name" tfsdk:"partition_name"`
	PriorityTier     *int64  `json:"priority_tier,omitempty" tfsdk:"priority_tier"`
	Qos              *string `json:"qos,omitempty" tfsdk:"qos"`
	ReqResv          *bool   `json:"req_resv,omitempty" tfsdk:"req_resv"`
}

type NestedPlanComponent struct {
	Amount            *int64  `json:"amount,omitempty" tfsdk:"amount"`
	DiscountRate      *int64  `json:"discount_rate,omitempty" tfsdk:"discount_rate"`
	DiscountThreshold *int64  `json:"discount_threshold,omitempty" tfsdk:"discount_threshold"`
	FuturePrice       *string `json:"future_price,omitempty" tfsdk:"future_price"`
	MeasuredUnit      *string `json:"measured_unit,omitempty" tfsdk:"measured_unit"`
	Name              *string `json:"name,omitempty" tfsdk:"name"`
	Price             *string `json:"price,omitempty" tfsdk:"price"`
	Type              *string `json:"type,omitempty" tfsdk:"type"`
}

type NestedPlanComponentRequest struct {
	Amount            *int64  `json:"amount,omitempty" tfsdk:"amount"`
	DiscountRate      *int64  `json:"discount_rate,omitempty" tfsdk:"discount_rate"`
	DiscountThreshold *int64  `json:"discount_threshold,omitempty" tfsdk:"discount_threshold"`
	FuturePrice       *string `json:"future_price,omitempty" tfsdk:"future_price"`
	Price             *string `json:"price,omitempty" tfsdk:"price"`
}

type NestedPriceEstimate struct {
	Current    *string `json:"current,omitempty" tfsdk:"current"`
	Tax        *string `json:"tax,omitempty" tfsdk:"tax"`
	TaxCurrent *string `json:"tax_current,omitempty" tfsdk:"tax_current"`
	Total      *string `json:"total,omitempty" tfsdk:"total"`
}

type NestedProject struct {
	Url *string `json:"url" tfsdk:"url"`
}

type NestedProjectPermission struct {
	ExpirationTime *string `json:"expiration_time,omitempty" tfsdk:"expiration_time"`
	Name           *string `json:"name,omitempty" tfsdk:"name"`
	RoleName       *string `json:"role_name,omitempty" tfsdk:"role_name"`
	Url            *string `json:"url,omitempty" tfsdk:"url"`
}

type NestedProviderOffering struct {
	Name *string `json:"name" tfsdk:"name"`
	Type *string `json:"type" tfsdk:"type"`
	Url  *string `json:"url" tfsdk:"url"`
}

type NestedProviderOfferingRequest struct {
	Name *string `json:"name" tfsdk:"name"`
	Type *string `json:"type" tfsdk:"type"`
}

type NestedPublicOffering struct {
	Name *string `json:"name" tfsdk:"name"`
	Type *string `json:"type" tfsdk:"type"`
	Url  *string `json:"url" tfsdk:"url"`
}

type NestedPublicOfferingRequest struct {
	Name *string `json:"name" tfsdk:"name"`
	Type *string `json:"type" tfsdk:"type"`
}

type NestedRemoteLocalCategory struct {
	LocalCategory      *string `json:"local_category" tfsdk:"local_category"`
	LocalCategoryName  *string `json:"local_category_name" tfsdk:"local_category_name"`
	LocalCategoryUuid  *string `json:"local_category_uuid" tfsdk:"local_category_uuid"`
	RemoteCategory     *string `json:"remote_category" tfsdk:"remote_category"`
	RemoteCategoryName *string `json:"remote_category_name,omitempty" tfsdk:"remote_category_name"`
}

type NestedRemoteLocalCategoryRequest struct {
	LocalCategory      *string `json:"local_category" tfsdk:"local_category"`
	RemoteCategory     *string `json:"remote_category" tfsdk:"remote_category"`
	RemoteCategoryName *string `json:"remote_category_name,omitempty" tfsdk:"remote_category_name"`
}

type NestedRequestedOffering struct {
	CallManagingOrganisation *string             `json:"call_managing_organisation,omitempty" tfsdk:"call_managing_organisation"`
	CategoryName             *string             `json:"category_name,omitempty" tfsdk:"category_name"`
	CategoryUuid             *string             `json:"category_uuid,omitempty" tfsdk:"category_uuid"`
	Components               []OfferingComponent `json:"components,omitempty" tfsdk:"components"`
	Created                  *string             `json:"created,omitempty" tfsdk:"created"`
	Offering                 *string             `json:"offering,omitempty" tfsdk:"offering"`
	OfferingName             *string             `json:"offering_name,omitempty" tfsdk:"offering_name"`
	OfferingUuid             *string             `json:"offering_uuid,omitempty" tfsdk:"offering_uuid"`
	Plan                     *string             `json:"plan,omitempty" tfsdk:"plan"`
	ProviderName             *string             `json:"provider_name,omitempty" tfsdk:"provider_name"`
	State                    *string             `json:"state,omitempty" tfsdk:"state"`
}

type NestedRequestedOfferingRequest struct {
	Offering *string `json:"offering" tfsdk:"offering"`
	Plan     *string `json:"plan,omitempty" tfsdk:"plan"`
}

type NestedRole struct {
	Name *string `json:"name,omitempty" tfsdk:"name"`
	Url  *string `json:"url,omitempty" tfsdk:"url"`
}

type NestedRoleRequest struct {
	Name *string `json:"name" tfsdk:"name"`
}

type NestedRound struct {
	AllocationDate           *string `json:"allocation_date,omitempty" tfsdk:"allocation_date"`
	AllocationTime           *string `json:"allocation_time,omitempty" tfsdk:"allocation_time"`
	CutoffTime               *string `json:"cutoff_time,omitempty" tfsdk:"cutoff_time"`
	DecidingEntity           *string `json:"deciding_entity,omitempty" tfsdk:"deciding_entity"`
	MinimalAverageScoring    *string `json:"minimal_average_scoring,omitempty" tfsdk:"minimal_average_scoring"`
	MinimumNumberOfReviewers *int64  `json:"minimum_number_of_reviewers,omitempty" tfsdk:"minimum_number_of_reviewers"`
	Name                     *string `json:"name,omitempty" tfsdk:"name"`
	ReviewDurationInDays     *int64  `json:"review_duration_in_days,omitempty" tfsdk:"review_duration_in_days"`
	ReviewStrategy           *string `json:"review_strategy,omitempty" tfsdk:"review_strategy"`
	Slug                     *string `json:"slug,omitempty" tfsdk:"slug"`
	StartTime                *string `json:"start_time,omitempty" tfsdk:"start_time"`
	Status                   *string `json:"status,omitempty" tfsdk:"status"`
}

type NestedRoundRequest struct {
	AllocationDate           *string `json:"allocation_date,omitempty" tfsdk:"allocation_date"`
	AllocationTime           *string `json:"allocation_time,omitempty" tfsdk:"allocation_time"`
	CutoffTime               *string `json:"cutoff_time" tfsdk:"cutoff_time"`
	DecidingEntity           *string `json:"deciding_entity,omitempty" tfsdk:"deciding_entity"`
	MinimalAverageScoring    *string `json:"minimal_average_scoring,omitempty" tfsdk:"minimal_average_scoring"`
	MinimumNumberOfReviewers *int64  `json:"minimum_number_of_reviewers,omitempty" tfsdk:"minimum_number_of_reviewers"`
	ReviewDurationInDays     *int64  `json:"review_duration_in_days,omitempty" tfsdk:"review_duration_in_days"`
	ReviewStrategy           *string `json:"review_strategy,omitempty" tfsdk:"review_strategy"`
	Slug                     *string `json:"slug,omitempty" tfsdk:"slug"`
	StartTime                *string `json:"start_time" tfsdk:"start_time"`
}

type NestedScreenshot struct {
	Created     *string `json:"created,omitempty" tfsdk:"created"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Image       *string `json:"image,omitempty" tfsdk:"image"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
	Thumbnail   *string `json:"thumbnail,omitempty" tfsdk:"thumbnail"`
}

type NestedScreenshotRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Image       *string `json:"image" tfsdk:"image"`
	Name        *string `json:"name" tfsdk:"name"`
}

type NestedSection struct {
	Attributes   []NestedAttribute `json:"attributes,omitempty" tfsdk:"attributes"`
	IsStandalone *bool             `json:"is_standalone,omitempty" tfsdk:"is_standalone"`
	Key          *string           `json:"key,omitempty" tfsdk:"key"`
	Title        *string           `json:"title,omitempty" tfsdk:"title"`
}

type NestedSectionRequest struct {
	IsStandalone *bool   `json:"is_standalone,omitempty" tfsdk:"is_standalone"`
	Key          *string `json:"key" tfsdk:"key"`
	Title        *string `json:"title" tfsdk:"title"`
}

type NestedSecurityGroupRule struct {
	Cidr            *string `json:"cidr,omitempty" tfsdk:"cidr"`
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	Direction       *string `json:"direction,omitempty" tfsdk:"direction"`
	Ethertype       *string `json:"ethertype,omitempty" tfsdk:"ethertype"`
	FromPort        *int64  `json:"from_port,omitempty" tfsdk:"from_port"`
	Id              *int64  `json:"id,omitempty" tfsdk:"id"`
	Protocol        *string `json:"protocol,omitempty" tfsdk:"protocol"`
	RemoteGroupName *string `json:"remote_group_name,omitempty" tfsdk:"remote_group_name"`
	RemoteGroupUuid *string `json:"remote_group_uuid,omitempty" tfsdk:"remote_group_uuid"`
	ToPort          *int64  `json:"to_port,omitempty" tfsdk:"to_port"`
}

type NestedSecurityGroupRuleRequest struct {
	Cidr        *string `json:"cidr,omitempty" tfsdk:"cidr"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Direction   *string `json:"direction,omitempty" tfsdk:"direction"`
	Ethertype   *string `json:"ethertype,omitempty" tfsdk:"ethertype"`
	FromPort    *int64  `json:"from_port,omitempty" tfsdk:"from_port"`
	Protocol    *string `json:"protocol,omitempty" tfsdk:"protocol"`
	ToPort      *int64  `json:"to_port,omitempty" tfsdk:"to_port"`
}

type NestedSoftwareCatalog struct {
	Catalog      *NestedSoftwareCatalogCatalog   `json:"catalog,omitempty" tfsdk:"catalog"`
	PackageCount *int64                          `json:"package_count,omitempty" tfsdk:"package_count"`
	Partition    *NestedSoftwareCatalogPartition `json:"partition,omitempty" tfsdk:"partition"`
}

type NestedSoftwareCatalogCatalog struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
	Version     *string `json:"version,omitempty" tfsdk:"version"`
}

type NestedSoftwareCatalogPartition struct {
	PartitionName *string `json:"partition_name,omitempty" tfsdk:"partition_name"`
	PriorityTier  *int64  `json:"priority_tier,omitempty" tfsdk:"priority_tier"`
	Qos           *string `json:"qos,omitempty" tfsdk:"qos"`
}

type NestedSoftwareCatalogRequest struct {
}

type NestedSoftwareTarget struct {
	Location      *string `json:"location,omitempty" tfsdk:"location"`
	TargetName    *string `json:"target_name,omitempty" tfsdk:"target_name"`
	TargetSubtype *string `json:"target_subtype,omitempty" tfsdk:"target_subtype"`
	TargetType    *string `json:"target_type,omitempty" tfsdk:"target_type"`
}

type NestedSoftwareTargetRequest struct {
	Location      *string `json:"location,omitempty" tfsdk:"location"`
	TargetName    *string `json:"target_name,omitempty" tfsdk:"target_name"`
	TargetSubtype *string `json:"target_subtype,omitempty" tfsdk:"target_subtype"`
	TargetType    *string `json:"target_type,omitempty" tfsdk:"target_type"`
}

type NestedSoftwareVersion struct {
	ReleaseDate *string                `json:"release_date,omitempty" tfsdk:"release_date"`
	Targets     []NestedSoftwareTarget `json:"targets" tfsdk:"targets"`
	Version     *string                `json:"version" tfsdk:"version"`
}

type NestedSoftwareVersionRequest struct {
	ReleaseDate *string `json:"release_date,omitempty" tfsdk:"release_date"`
	Version     *string `json:"version" tfsdk:"version"`
}

type NetworkRBACPolicy struct {
	BackendId        *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created          *string `json:"created,omitempty" tfsdk:"created"`
	Network          *string `json:"network,omitempty" tfsdk:"network"`
	NetworkName      *string `json:"network_name,omitempty" tfsdk:"network_name"`
	PolicyType       *string `json:"policy_type,omitempty" tfsdk:"policy_type"`
	TargetTenant     *string `json:"target_tenant,omitempty" tfsdk:"target_tenant"`
	TargetTenantName *string `json:"target_tenant_name,omitempty" tfsdk:"target_tenant_name"`
	Url              *string `json:"url,omitempty" tfsdk:"url"`
}

type NetworkRBACPolicyRequest struct {
	Network      *string `json:"network" tfsdk:"network"`
	PolicyType   *string `json:"policy_type,omitempty" tfsdk:"policy_type"`
	TargetTenant *string `json:"target_tenant" tfsdk:"target_tenant"`
}

type NodeDiskDriverEnum struct {
}

type Notification struct {
	Created     *string                                 `json:"created" tfsdk:"created"`
	Description *string                                 `json:"description,omitempty" tfsdk:"description"`
	Enabled     *bool                                   `json:"enabled" tfsdk:"enabled"`
	Key         *string                                 `json:"key" tfsdk:"key"`
	Templates   []NotificationTemplateDetailSerializers `json:"templates" tfsdk:"templates"`
	Url         *string                                 `json:"url" tfsdk:"url"`
}

type NotificationRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Key         *string `json:"key" tfsdk:"key"`
}

type NotificationTemplateDetailSerializers struct {
	Content             *string `json:"content" tfsdk:"content"`
	IsContentOverridden *bool   `json:"is_content_overridden" tfsdk:"is_content_overridden"`
	Name                *string `json:"name" tfsdk:"name"`
	OriginalContent     *string `json:"original_content" tfsdk:"original_content"`
	Path                *string `json:"path" tfsdk:"path"`
	Url                 *string `json:"url" tfsdk:"url"`
}

type NotificationTemplateDetailSerializersRequest struct {
	Name *string `json:"name" tfsdk:"name"`
	Path *string `json:"path" tfsdk:"path"`
}

type NotificationTemplateUpdateSerializersRequest struct {
	Content *string `json:"content" tfsdk:"content"`
}

type NullEnum struct {
}

type ObservableObjectTypeEnum struct {
}

type ObtainAuthTokenRequest struct {
	Password *string `json:"password" tfsdk:"password"`
	Username *string `json:"username" tfsdk:"username"`
}

type OecdFos2007CodeEnum struct {
}

type Offering struct {
	AccessUrl                 *string                 `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                 *string                 `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Billable                  *bool                   `json:"billable,omitempty" tfsdk:"billable"`
	BillingTypeClassification *string                 `json:"billing_type_classification,omitempty" tfsdk:"billing_type_classification"`
	Category                  *string                 `json:"category,omitempty" tfsdk:"category"`
	CategoryTitle             *string                 `json:"category_title,omitempty" tfsdk:"category_title"`
	CategoryUuid              *string                 `json:"category_uuid,omitempty" tfsdk:"category_uuid"`
	CitationCount             *int64                  `json:"citation_count,omitempty" tfsdk:"citation_count"`
	ComplianceChecklist       *string                 `json:"compliance_checklist,omitempty" tfsdk:"compliance_checklist"`
	Components                []OfferingComponent     `json:"components,omitempty" tfsdk:"components"`
	Country                   *string                 `json:"country,omitempty" tfsdk:"country"`
	Created                   *string                 `json:"created,omitempty" tfsdk:"created"`
	Customer                  *string                 `json:"customer,omitempty" tfsdk:"customer"`
	CustomerName              *string                 `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerUuid              *string                 `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	DataciteDoi               *string                 `json:"datacite_doi,omitempty" tfsdk:"datacite_doi"`
	Description               *string                 `json:"description,omitempty" tfsdk:"description"`
	Endpoints                 []NestedEndpoint        `json:"endpoints,omitempty" tfsdk:"endpoints"`
	Files                     []NestedOfferingFile    `json:"files,omitempty" tfsdk:"files"`
	FullDescription           *string                 `json:"full_description,omitempty" tfsdk:"full_description"`
	GettingStarted            *string                 `json:"getting_started,omitempty" tfsdk:"getting_started"`
	Googlecalendar            *GoogleCalendar         `json:"googlecalendar,omitempty" tfsdk:"googlecalendar"`
	HasComplianceRequirements *bool                   `json:"has_compliance_requirements,omitempty" tfsdk:"has_compliance_requirements"`
	Image                     *string                 `json:"image,omitempty" tfsdk:"image"`
	IntegrationGuide          *string                 `json:"integration_guide,omitempty" tfsdk:"integration_guide"`
	Latitude                  *float64                `json:"latitude,omitempty" tfsdk:"latitude"`
	Longitude                 *float64                `json:"longitude,omitempty" tfsdk:"longitude"`
	Name                      *string                 `json:"name,omitempty" tfsdk:"name"`
	OrderCount                *int64                  `json:"order_count,omitempty" tfsdk:"order_count"`
	OrganizationGroups        []OrganizationGroup     `json:"organization_groups,omitempty" tfsdk:"organization_groups"`
	ParentDescription         *string                 `json:"parent_description,omitempty" tfsdk:"parent_description"`
	ParentName                *string                 `json:"parent_name,omitempty" tfsdk:"parent_name"`
	ParentUuid                *string                 `json:"parent_uuid,omitempty" tfsdk:"parent_uuid"`
	Partitions                []NestedPartition       `json:"partitions,omitempty" tfsdk:"partitions"`
	PausedReason              *string                 `json:"paused_reason,omitempty" tfsdk:"paused_reason"`
	Plans                     []BasePublicPlan        `json:"plans,omitempty" tfsdk:"plans"`
	PrivacyPolicyLink         *string                 `json:"privacy_policy_link,omitempty" tfsdk:"privacy_policy_link"`
	Project                   *string                 `json:"project,omitempty" tfsdk:"project"`
	ProjectName               *string                 `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid               *string                 `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	Quotas                    []Quota                 `json:"quotas,omitempty" tfsdk:"quotas"`
	Roles                     []NestedRole            `json:"roles,omitempty" tfsdk:"roles"`
	Scope                     *string                 `json:"scope,omitempty" tfsdk:"scope"`
	ScopeErrorMessage         *string                 `json:"scope_error_message,omitempty" tfsdk:"scope_error_message"`
	ScopeName                 *string                 `json:"scope_name,omitempty" tfsdk:"scope_name"`
	ScopeState                *string                 `json:"scope_state,omitempty" tfsdk:"scope_state"`
	ScopeUuid                 *string                 `json:"scope_uuid,omitempty" tfsdk:"scope_uuid"`
	Screenshots               []NestedScreenshot      `json:"screenshots,omitempty" tfsdk:"screenshots"`
	Shared                    *bool                   `json:"shared,omitempty" tfsdk:"shared"`
	Slug                      *string                 `json:"slug,omitempty" tfsdk:"slug"`
	SoftwareCatalogs          []NestedSoftwareCatalog `json:"software_catalogs,omitempty" tfsdk:"software_catalogs"`
	State                     *string                 `json:"state,omitempty" tfsdk:"state"`
	Thumbnail                 *string                 `json:"thumbnail,omitempty" tfsdk:"thumbnail"`
	TotalCost                 *int64                  `json:"total_cost,omitempty" tfsdk:"total_cost"`
	TotalCostEstimated        *int64                  `json:"total_cost_estimated,omitempty" tfsdk:"total_cost_estimated"`
	TotalCustomers            *int64                  `json:"total_customers,omitempty" tfsdk:"total_customers"`
	Type                      *string                 `json:"type,omitempty" tfsdk:"type"`
	Url                       *string                 `json:"url,omitempty" tfsdk:"url"`
	UserHasConsent            *bool                   `json:"user_has_consent,omitempty" tfsdk:"user_has_consent"`
	VendorDetails             *string                 `json:"vendor_details,omitempty" tfsdk:"vendor_details"`
}

type OfferingBackendMetadataRequest struct {
}

type OfferingComplianceChecklistUpdateRequest struct {
	ComplianceChecklist *string `json:"compliance_checklist,omitempty" tfsdk:"compliance_checklist"`
}

type OfferingComponent struct {
	ArticleCode        *string `json:"article_code,omitempty" tfsdk:"article_code"`
	BillingType        *string `json:"billing_type,omitempty" tfsdk:"billing_type"`
	DefaultLimit       *int64  `json:"default_limit,omitempty" tfsdk:"default_limit"`
	Description        *string `json:"description,omitempty" tfsdk:"description"`
	Factor             *int64  `json:"factor,omitempty" tfsdk:"factor"`
	IsBoolean          *bool   `json:"is_boolean,omitempty" tfsdk:"is_boolean"`
	IsBuiltin          *bool   `json:"is_builtin,omitempty" tfsdk:"is_builtin"`
	IsPrepaid          *bool   `json:"is_prepaid,omitempty" tfsdk:"is_prepaid"`
	LimitAmount        *int64  `json:"limit_amount,omitempty" tfsdk:"limit_amount"`
	LimitPeriod        *string `json:"limit_period,omitempty" tfsdk:"limit_period"`
	MaxAvailableLimit  *int64  `json:"max_available_limit,omitempty" tfsdk:"max_available_limit"`
	MaxPrepaidDuration *int64  `json:"max_prepaid_duration,omitempty" tfsdk:"max_prepaid_duration"`
	MaxValue           *int64  `json:"max_value,omitempty" tfsdk:"max_value"`
	MeasuredUnit       *string `json:"measured_unit,omitempty" tfsdk:"measured_unit"`
	MinPrepaidDuration *int64  `json:"min_prepaid_duration,omitempty" tfsdk:"min_prepaid_duration"`
	MinValue           *int64  `json:"min_value,omitempty" tfsdk:"min_value"`
	Name               *string `json:"name,omitempty" tfsdk:"name"`
	OverageComponent   *string `json:"overage_component,omitempty" tfsdk:"overage_component"`
	Type               *string `json:"type,omitempty" tfsdk:"type"`
	UnitFactor         *int64  `json:"unit_factor,omitempty" tfsdk:"unit_factor"`
}

type OfferingComponentLimitRequest struct {
	Max               *int64 `json:"max" tfsdk:"max"`
	MaxAvailableLimit *int64 `json:"max_available_limit" tfsdk:"max_available_limit"`
	Min               *int64 `json:"min" tfsdk:"min"`
}

type OfferingComponentRequest struct {
	ArticleCode        *string `json:"article_code,omitempty" tfsdk:"article_code"`
	BillingType        *string `json:"billing_type" tfsdk:"billing_type"`
	DefaultLimit       *int64  `json:"default_limit,omitempty" tfsdk:"default_limit"`
	Description        *string `json:"description,omitempty" tfsdk:"description"`
	IsBoolean          *bool   `json:"is_boolean,omitempty" tfsdk:"is_boolean"`
	IsPrepaid          *bool   `json:"is_prepaid,omitempty" tfsdk:"is_prepaid"`
	LimitAmount        *int64  `json:"limit_amount,omitempty" tfsdk:"limit_amount"`
	LimitPeriod        *string `json:"limit_period,omitempty" tfsdk:"limit_period"`
	MaxAvailableLimit  *int64  `json:"max_available_limit,omitempty" tfsdk:"max_available_limit"`
	MaxPrepaidDuration *int64  `json:"max_prepaid_duration,omitempty" tfsdk:"max_prepaid_duration"`
	MaxValue           *int64  `json:"max_value,omitempty" tfsdk:"max_value"`
	MeasuredUnit       *string `json:"measured_unit,omitempty" tfsdk:"measured_unit"`
	MinPrepaidDuration *int64  `json:"min_prepaid_duration,omitempty" tfsdk:"min_prepaid_duration"`
	MinValue           *int64  `json:"min_value,omitempty" tfsdk:"min_value"`
	Name               *string `json:"name" tfsdk:"name"`
	OverageComponent   *string `json:"overage_component,omitempty" tfsdk:"overage_component"`
	Type               *string `json:"type" tfsdk:"type"`
	UnitFactor         *int64  `json:"unit_factor,omitempty" tfsdk:"unit_factor"`
}

type OfferingComponentStat struct {
	BillingPeriod *string `json:"billing_period" tfsdk:"billing_period"`
	Date          *string `json:"date" tfsdk:"date"`
	Description   *string `json:"description" tfsdk:"description"`
	MeasuredUnit  *string `json:"measured_unit" tfsdk:"measured_unit"`
	Name          *string `json:"name" tfsdk:"name"`
	Period        *string `json:"period" tfsdk:"period"`
	Type          *string `json:"type" tfsdk:"type"`
	Usage         *int64  `json:"usage" tfsdk:"usage"`
}

type OfferingCost struct {
	Cost         *float64 `json:"cost" tfsdk:"cost"`
	OfferingName *string  `json:"offering_name" tfsdk:"offering_name"`
	OfferingUuid *string  `json:"offering_uuid" tfsdk:"offering_uuid"`
}

type OfferingCountryStats struct {
	Count   *int64  `json:"count" tfsdk:"count"`
	Country *string `json:"country" tfsdk:"country"`
}

type OfferingCreateRequest struct {
	AccessUrl           *string                    `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId           *string                    `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Billable            *bool                      `json:"billable,omitempty" tfsdk:"billable"`
	Category            *string                    `json:"category" tfsdk:"category"`
	ComplianceChecklist *string                    `json:"compliance_checklist,omitempty" tfsdk:"compliance_checklist"`
	Components          []OfferingComponentRequest `json:"components,omitempty" tfsdk:"components"`
	Country             *string                    `json:"country,omitempty" tfsdk:"country"`
	Customer            *string                    `json:"customer,omitempty" tfsdk:"customer"`
	DataciteDoi         *string                    `json:"datacite_doi,omitempty" tfsdk:"datacite_doi"`
	Description         *string                    `json:"description,omitempty" tfsdk:"description"`
	FullDescription     *string                    `json:"full_description,omitempty" tfsdk:"full_description"`
	GettingStarted      *string                    `json:"getting_started,omitempty" tfsdk:"getting_started"`
	Image               *string                    `json:"image,omitempty" tfsdk:"image"`
	IntegrationGuide    *string                    `json:"integration_guide,omitempty" tfsdk:"integration_guide"`
	Latitude            *float64                   `json:"latitude,omitempty" tfsdk:"latitude"`
	Longitude           *float64                   `json:"longitude,omitempty" tfsdk:"longitude"`
	Name                *string                    `json:"name" tfsdk:"name"`
	Options             *OfferingOptionsRequest    `json:"options,omitempty" tfsdk:"options"`
	Plans               []BaseProviderPlanRequest  `json:"plans,omitempty" tfsdk:"plans"`
	PrivacyPolicyLink   *string                    `json:"privacy_policy_link,omitempty" tfsdk:"privacy_policy_link"`
	ResourceOptions     *OfferingOptionsRequest    `json:"resource_options,omitempty" tfsdk:"resource_options"`
	Shared              *bool                      `json:"shared,omitempty" tfsdk:"shared"`
	Slug                *string                    `json:"slug,omitempty" tfsdk:"slug"`
	Thumbnail           *string                    `json:"thumbnail,omitempty" tfsdk:"thumbnail"`
	Type                *string                    `json:"type" tfsdk:"type"`
	VendorDetails       *string                    `json:"vendor_details,omitempty" tfsdk:"vendor_details"`
}

type OfferingCreateRequestForm struct {
	AccessUrl           *string                    `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId           *string                    `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Billable            *bool                      `json:"billable,omitempty" tfsdk:"billable"`
	Category            *string                    `json:"category" tfsdk:"category"`
	ComplianceChecklist *string                    `json:"compliance_checklist,omitempty" tfsdk:"compliance_checklist"`
	Components          []OfferingComponentRequest `json:"components,omitempty" tfsdk:"components"`
	Country             *string                    `json:"country,omitempty" tfsdk:"country"`
	Customer            *string                    `json:"customer,omitempty" tfsdk:"customer"`
	DataciteDoi         *string                    `json:"datacite_doi,omitempty" tfsdk:"datacite_doi"`
	Description         *string                    `json:"description,omitempty" tfsdk:"description"`
	FullDescription     *string                    `json:"full_description,omitempty" tfsdk:"full_description"`
	GettingStarted      *string                    `json:"getting_started,omitempty" tfsdk:"getting_started"`
	Image               *string                    `json:"image,omitempty" tfsdk:"image"`
	IntegrationGuide    *string                    `json:"integration_guide,omitempty" tfsdk:"integration_guide"`
	Latitude            *float64                   `json:"latitude,omitempty" tfsdk:"latitude"`
	Longitude           *float64                   `json:"longitude,omitempty" tfsdk:"longitude"`
	Name                *string                    `json:"name" tfsdk:"name"`
	Options             *OfferingOptionsRequest    `json:"options,omitempty" tfsdk:"options"`
	Plans               []BaseProviderPlanRequest  `json:"plans,omitempty" tfsdk:"plans"`
	PrivacyPolicyLink   *string                    `json:"privacy_policy_link,omitempty" tfsdk:"privacy_policy_link"`
	ResourceOptions     *OfferingOptionsRequest    `json:"resource_options,omitempty" tfsdk:"resource_options"`
	Shared              *bool                      `json:"shared,omitempty" tfsdk:"shared"`
	Slug                *string                    `json:"slug,omitempty" tfsdk:"slug"`
	Thumbnail           *string                    `json:"thumbnail,omitempty" tfsdk:"thumbnail"`
	Type                *string                    `json:"type" tfsdk:"type"`
	VendorDetails       *string                    `json:"vendor_details,omitempty" tfsdk:"vendor_details"`
}

type OfferingCreateRequestMultipart struct {
	AccessUrl           *string                    `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId           *string                    `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Billable            *bool                      `json:"billable,omitempty" tfsdk:"billable"`
	Category            *string                    `json:"category" tfsdk:"category"`
	ComplianceChecklist *string                    `json:"compliance_checklist,omitempty" tfsdk:"compliance_checklist"`
	Components          []OfferingComponentRequest `json:"components,omitempty" tfsdk:"components"`
	Country             *string                    `json:"country,omitempty" tfsdk:"country"`
	Customer            *string                    `json:"customer,omitempty" tfsdk:"customer"`
	DataciteDoi         *string                    `json:"datacite_doi,omitempty" tfsdk:"datacite_doi"`
	Description         *string                    `json:"description,omitempty" tfsdk:"description"`
	FullDescription     *string                    `json:"full_description,omitempty" tfsdk:"full_description"`
	GettingStarted      *string                    `json:"getting_started,omitempty" tfsdk:"getting_started"`
	Image               *string                    `json:"image,omitempty" tfsdk:"image"`
	IntegrationGuide    *string                    `json:"integration_guide,omitempty" tfsdk:"integration_guide"`
	Latitude            *float64                   `json:"latitude,omitempty" tfsdk:"latitude"`
	Longitude           *float64                   `json:"longitude,omitempty" tfsdk:"longitude"`
	Name                *string                    `json:"name" tfsdk:"name"`
	Options             *OfferingOptionsRequest    `json:"options,omitempty" tfsdk:"options"`
	Plans               []BaseProviderPlanRequest  `json:"plans,omitempty" tfsdk:"plans"`
	PrivacyPolicyLink   *string                    `json:"privacy_policy_link,omitempty" tfsdk:"privacy_policy_link"`
	ResourceOptions     *OfferingOptionsRequest    `json:"resource_options,omitempty" tfsdk:"resource_options"`
	Shared              *bool                      `json:"shared,omitempty" tfsdk:"shared"`
	Slug                *string                    `json:"slug,omitempty" tfsdk:"slug"`
	Thumbnail           *string                    `json:"thumbnail,omitempty" tfsdk:"thumbnail"`
	Type                *string                    `json:"type" tfsdk:"type"`
	VendorDetails       *string                    `json:"vendor_details,omitempty" tfsdk:"vendor_details"`
}

type OfferingDescriptionUpdateRequest struct {
	Category *string `json:"category" tfsdk:"category"`
}

type OfferingEstimatedCostPolicy struct {
	Actions            *string  `json:"actions" tfsdk:"actions"`
	ApplyToAll         *bool    `json:"apply_to_all,omitempty" tfsdk:"apply_to_all"`
	Created            *string  `json:"created" tfsdk:"created"`
	CreatedByFullName  *string  `json:"created_by_full_name" tfsdk:"created_by_full_name"`
	CreatedByUsername  *string  `json:"created_by_username" tfsdk:"created_by_username"`
	FiredDatetime      *string  `json:"fired_datetime" tfsdk:"fired_datetime"`
	HasFired           *bool    `json:"has_fired" tfsdk:"has_fired"`
	LimitCost          *int64   `json:"limit_cost" tfsdk:"limit_cost"`
	OrganizationGroups []string `json:"organization_groups,omitempty" tfsdk:"organization_groups"`
	Period             *int64   `json:"period,omitempty" tfsdk:"period"`
	PeriodName         *string  `json:"period_name" tfsdk:"period_name"`
	Scope              *string  `json:"scope" tfsdk:"scope"`
	ScopeName          *string  `json:"scope_name" tfsdk:"scope_name"`
	ScopeUuid          *string  `json:"scope_uuid" tfsdk:"scope_uuid"`
	Url                *string  `json:"url" tfsdk:"url"`
}

type OfferingEstimatedCostPolicyRequest struct {
	Actions            *string  `json:"actions" tfsdk:"actions"`
	ApplyToAll         *bool    `json:"apply_to_all,omitempty" tfsdk:"apply_to_all"`
	LimitCost          *int64   `json:"limit_cost" tfsdk:"limit_cost"`
	OrganizationGroups []string `json:"organization_groups,omitempty" tfsdk:"organization_groups"`
	Period             *int64   `json:"period,omitempty" tfsdk:"period"`
	Scope              *string  `json:"scope" tfsdk:"scope"`
}

type OfferingExportData struct {
	Components         []ExportComponentData         `json:"components,omitempty" tfsdk:"components"`
	Endpoints          []ExportEndpointData          `json:"endpoints,omitempty" tfsdk:"endpoints"`
	Files              []ExportFileData              `json:"files,omitempty" tfsdk:"files"`
	Offering           *ExportOfferingData           `json:"offering" tfsdk:"offering"`
	OrganizationGroups []ExportOrganizationGroupData `json:"organization_groups,omitempty" tfsdk:"organization_groups"`
	Plans              []ExportPlanData              `json:"plans,omitempty" tfsdk:"plans"`
	Screenshots        []ExportScreenshotData        `json:"screenshots,omitempty" tfsdk:"screenshots"`
	TermsOfService     []ExportTermsOfServiceData    `json:"terms_of_service,omitempty" tfsdk:"terms_of_service"`
}

type OfferingExportDataRequest struct {
	Components         []ExportComponentDataRequest         `json:"components,omitempty" tfsdk:"components"`
	Endpoints          []ExportEndpointDataRequest          `json:"endpoints,omitempty" tfsdk:"endpoints"`
	Files              []ExportFileDataRequest              `json:"files,omitempty" tfsdk:"files"`
	Offering           *ExportOfferingDataRequest           `json:"offering" tfsdk:"offering"`
	OrganizationGroups []ExportOrganizationGroupDataRequest `json:"organization_groups,omitempty" tfsdk:"organization_groups"`
	Plans              []ExportPlanDataRequest              `json:"plans,omitempty" tfsdk:"plans"`
	Screenshots        []ExportScreenshotDataRequest        `json:"screenshots,omitempty" tfsdk:"screenshots"`
	TermsOfService     []ExportTermsOfServiceDataRequest    `json:"terms_of_service,omitempty" tfsdk:"terms_of_service"`
}

type OfferingExportParametersRequest struct {
	IncludeAttributes         *bool `json:"include_attributes,omitempty" tfsdk:"include_attributes"`
	IncludeComponents         *bool `json:"include_components,omitempty" tfsdk:"include_components"`
	IncludeEndpoints          *bool `json:"include_endpoints,omitempty" tfsdk:"include_endpoints"`
	IncludeFiles              *bool `json:"include_files,omitempty" tfsdk:"include_files"`
	IncludeOptions            *bool `json:"include_options,omitempty" tfsdk:"include_options"`
	IncludeOrganizationGroups *bool `json:"include_organization_groups,omitempty" tfsdk:"include_organization_groups"`
	IncludePlans              *bool `json:"include_plans,omitempty" tfsdk:"include_plans"`
	IncludePluginOptions      *bool `json:"include_plugin_options,omitempty" tfsdk:"include_plugin_options"`
	IncludeResourceOptions    *bool `json:"include_resource_options,omitempty" tfsdk:"include_resource_options"`
	IncludeScreenshots        *bool `json:"include_screenshots,omitempty" tfsdk:"include_screenshots"`
	IncludeSecretOptions      *bool `json:"include_secret_options,omitempty" tfsdk:"include_secret_options"`
	IncludeTermsOfService     *bool `json:"include_terms_of_service,omitempty" tfsdk:"include_terms_of_service"`
}

type OfferingExportResponse struct {
	ExportTimestamp    *string  `json:"export_timestamp" tfsdk:"export_timestamp"`
	ExportedComponents []string `json:"exported_components" tfsdk:"exported_components"`
	OfferingName       *string  `json:"offering_name" tfsdk:"offering_name"`
	OfferingUuid       *string  `json:"offering_uuid" tfsdk:"offering_uuid"`
}

type OfferingFile struct {
	Created  *string `json:"created,omitempty" tfsdk:"created"`
	File     *string `json:"file,omitempty" tfsdk:"file"`
	Name     *string `json:"name,omitempty" tfsdk:"name"`
	Offering *string `json:"offering,omitempty" tfsdk:"offering"`
	Url      *string `json:"url,omitempty" tfsdk:"url"`
}

type OfferingFileRequest struct {
	File     *string `json:"file" tfsdk:"file"`
	Name     *string `json:"name" tfsdk:"name"`
	Offering *string `json:"offering" tfsdk:"offering"`
}

type OfferingFileRequestForm struct {
	File     *string `json:"file" tfsdk:"file"`
	Name     *string `json:"name" tfsdk:"name"`
	Offering *string `json:"offering" tfsdk:"offering"`
}

type OfferingFileRequestMultipart struct {
	File     *string `json:"file" tfsdk:"file"`
	Name     *string `json:"name" tfsdk:"name"`
	Offering *string `json:"offering" tfsdk:"offering"`
}

type OfferingGroups struct {
	CustomerName *string             `json:"customer_name" tfsdk:"customer_name"`
	CustomerUuid *string             `json:"customer_uuid" tfsdk:"customer_uuid"`
	Offerings    []OfferingReference `json:"offerings" tfsdk:"offerings"`
}

type OfferingImageRequest struct {
	Image *string `json:"image" tfsdk:"image"`
}

type OfferingImageRequestForm struct {
	Image *string `json:"image" tfsdk:"image"`
}

type OfferingImageRequestMultipart struct {
	Image *string `json:"image" tfsdk:"image"`
}

type OfferingImportParametersRequest struct {
	Category                 *string `json:"category,omitempty" tfsdk:"category"`
	Customer                 *string `json:"customer,omitempty" tfsdk:"customer"`
	ImportComponents         *bool   `json:"import_components,omitempty" tfsdk:"import_components"`
	ImportEndpoints          *bool   `json:"import_endpoints,omitempty" tfsdk:"import_endpoints"`
	ImportFiles              *bool   `json:"import_files,omitempty" tfsdk:"import_files"`
	ImportOrganizationGroups *bool   `json:"import_organization_groups,omitempty" tfsdk:"import_organization_groups"`
	ImportPlans              *bool   `json:"import_plans,omitempty" tfsdk:"import_plans"`
	ImportPluginOptions      *bool   `json:"import_plugin_options,omitempty" tfsdk:"import_plugin_options"`
	ImportScreenshots        *bool   `json:"import_screenshots,omitempty" tfsdk:"import_screenshots"`
	ImportSecretOptions      *bool   `json:"import_secret_options,omitempty" tfsdk:"import_secret_options"`
	ImportTermsOfService     *bool   `json:"import_terms_of_service,omitempty" tfsdk:"import_terms_of_service"`
	OverwriteExisting        *bool   `json:"overwrite_existing,omitempty" tfsdk:"overwrite_existing"`
	Project                  *string `json:"project,omitempty" tfsdk:"project"`
}

type OfferingImportResponse struct {
	ImportTimestamp      *string  `json:"import_timestamp" tfsdk:"import_timestamp"`
	ImportedComponents   []string `json:"imported_components" tfsdk:"imported_components"`
	ImportedOfferingName *string  `json:"imported_offering_name" tfsdk:"imported_offering_name"`
	ImportedOfferingUuid *string  `json:"imported_offering_uuid" tfsdk:"imported_offering_uuid"`
	Warnings             []string `json:"warnings,omitempty" tfsdk:"warnings"`
}

type OfferingIntegrationUpdateRequest struct {
	BackendId     *string                     `json:"backend_id,omitempty" tfsdk:"backend_id"`
	PluginOptions *MergedPluginOptionsRequest `json:"plugin_options,omitempty" tfsdk:"plugin_options"`
	SecretOptions *MergedSecretOptionsRequest `json:"secret_options,omitempty" tfsdk:"secret_options"`
}

type OfferingLocationUpdateRequest struct {
	Latitude  *float64 `json:"latitude" tfsdk:"latitude"`
	Longitude *float64 `json:"longitude" tfsdk:"longitude"`
}

type OfferingOptions struct {
	Order []string `json:"order,omitempty" tfsdk:"order"`
}

type OfferingOptionsRequest struct {
	Order []string `json:"order" tfsdk:"order"`
}

type OfferingOptionsUpdateRequest struct {
	Options *OfferingOptionsRequest `json:"options" tfsdk:"options"`
}

type OfferingOverviewUpdateRequest struct {
	AccessUrl         *string `json:"access_url,omitempty" tfsdk:"access_url"`
	Description       *string `json:"description,omitempty" tfsdk:"description"`
	FullDescription   *string `json:"full_description,omitempty" tfsdk:"full_description"`
	GettingStarted    *string `json:"getting_started,omitempty" tfsdk:"getting_started"`
	IntegrationGuide  *string `json:"integration_guide,omitempty" tfsdk:"integration_guide"`
	Name              *string `json:"name" tfsdk:"name"`
	PrivacyPolicyLink *string `json:"privacy_policy_link,omitempty" tfsdk:"privacy_policy_link"`
	Slug              *string `json:"slug,omitempty" tfsdk:"slug"`
}

type OfferingPartition struct {
	CpuBind          *int64  `json:"cpu_bind,omitempty" tfsdk:"cpu_bind"`
	Created          *string `json:"created" tfsdk:"created"`
	DefCpuPerGpu     *int64  `json:"def_cpu_per_gpu,omitempty" tfsdk:"def_cpu_per_gpu"`
	DefMemPerCpu     *int64  `json:"def_mem_per_cpu,omitempty" tfsdk:"def_mem_per_cpu"`
	DefMemPerGpu     *int64  `json:"def_mem_per_gpu,omitempty" tfsdk:"def_mem_per_gpu"`
	DefMemPerNode    *int64  `json:"def_mem_per_node,omitempty" tfsdk:"def_mem_per_node"`
	DefaultTime      *int64  `json:"default_time,omitempty" tfsdk:"default_time"`
	ExclusiveTopo    *bool   `json:"exclusive_topo,omitempty" tfsdk:"exclusive_topo"`
	ExclusiveUser    *bool   `json:"exclusive_user,omitempty" tfsdk:"exclusive_user"`
	GraceTime        *int64  `json:"grace_time,omitempty" tfsdk:"grace_time"`
	MaxCpusPerNode   *int64  `json:"max_cpus_per_node,omitempty" tfsdk:"max_cpus_per_node"`
	MaxCpusPerSocket *int64  `json:"max_cpus_per_socket,omitempty" tfsdk:"max_cpus_per_socket"`
	MaxMemPerCpu     *int64  `json:"max_mem_per_cpu,omitempty" tfsdk:"max_mem_per_cpu"`
	MaxMemPerNode    *int64  `json:"max_mem_per_node,omitempty" tfsdk:"max_mem_per_node"`
	MaxNodes         *int64  `json:"max_nodes,omitempty" tfsdk:"max_nodes"`
	MaxTime          *int64  `json:"max_time,omitempty" tfsdk:"max_time"`
	MinNodes         *int64  `json:"min_nodes,omitempty" tfsdk:"min_nodes"`
	Modified         *string `json:"modified" tfsdk:"modified"`
	Offering         *string `json:"offering" tfsdk:"offering"`
	OfferingName     *string `json:"offering_name" tfsdk:"offering_name"`
	PartitionName    *string `json:"partition_name" tfsdk:"partition_name"`
	PriorityTier     *int64  `json:"priority_tier,omitempty" tfsdk:"priority_tier"`
	Qos              *string `json:"qos,omitempty" tfsdk:"qos"`
	ReqResv          *bool   `json:"req_resv,omitempty" tfsdk:"req_resv"`
}

type OfferingPartitionRequest struct {
	CpuBind          *int64  `json:"cpu_bind,omitempty" tfsdk:"cpu_bind"`
	DefCpuPerGpu     *int64  `json:"def_cpu_per_gpu,omitempty" tfsdk:"def_cpu_per_gpu"`
	DefMemPerCpu     *int64  `json:"def_mem_per_cpu,omitempty" tfsdk:"def_mem_per_cpu"`
	DefMemPerGpu     *int64  `json:"def_mem_per_gpu,omitempty" tfsdk:"def_mem_per_gpu"`
	DefMemPerNode    *int64  `json:"def_mem_per_node,omitempty" tfsdk:"def_mem_per_node"`
	DefaultTime      *int64  `json:"default_time,omitempty" tfsdk:"default_time"`
	ExclusiveTopo    *bool   `json:"exclusive_topo,omitempty" tfsdk:"exclusive_topo"`
	ExclusiveUser    *bool   `json:"exclusive_user,omitempty" tfsdk:"exclusive_user"`
	GraceTime        *int64  `json:"grace_time,omitempty" tfsdk:"grace_time"`
	MaxCpusPerNode   *int64  `json:"max_cpus_per_node,omitempty" tfsdk:"max_cpus_per_node"`
	MaxCpusPerSocket *int64  `json:"max_cpus_per_socket,omitempty" tfsdk:"max_cpus_per_socket"`
	MaxMemPerCpu     *int64  `json:"max_mem_per_cpu,omitempty" tfsdk:"max_mem_per_cpu"`
	MaxMemPerNode    *int64  `json:"max_mem_per_node,omitempty" tfsdk:"max_mem_per_node"`
	MaxNodes         *int64  `json:"max_nodes,omitempty" tfsdk:"max_nodes"`
	MaxTime          *int64  `json:"max_time,omitempty" tfsdk:"max_time"`
	MinNodes         *int64  `json:"min_nodes,omitempty" tfsdk:"min_nodes"`
	Offering         *string `json:"offering" tfsdk:"offering"`
	PartitionName    *string `json:"partition_name" tfsdk:"partition_name"`
	PriorityTier     *int64  `json:"priority_tier,omitempty" tfsdk:"priority_tier"`
	Qos              *string `json:"qos,omitempty" tfsdk:"qos"`
	ReqResv          *bool   `json:"req_resv,omitempty" tfsdk:"req_resv"`
}

type OfferingPauseRequest struct {
	PausedReason *string `json:"paused_reason,omitempty" tfsdk:"paused_reason"`
}

type OfferingPermission struct {
	Created        *string `json:"created" tfsdk:"created"`
	CreatedBy      *string `json:"created_by" tfsdk:"created_by"`
	ExpirationTime *string `json:"expiration_time,omitempty" tfsdk:"expiration_time"`
	Offering       *string `json:"offering" tfsdk:"offering"`
	OfferingName   *string `json:"offering_name" tfsdk:"offering_name"`
	OfferingSlug   *string `json:"offering_slug" tfsdk:"offering_slug"`
	OfferingUuid   *string `json:"offering_uuid" tfsdk:"offering_uuid"`
	Pk             *int64  `json:"pk" tfsdk:"pk"`
	Role           *string `json:"role" tfsdk:"role"`
	RoleName       *string `json:"role_name" tfsdk:"role_name"`
	Url            *string `json:"url" tfsdk:"url"`
	User           *string `json:"user" tfsdk:"user"`
	UserEmail      *string `json:"user_email" tfsdk:"user_email"`
	UserFullName   *string `json:"user_full_name" tfsdk:"user_full_name"`
	UserNativeName *string `json:"user_native_name" tfsdk:"user_native_name"`
	UserUsername   *string `json:"user_username" tfsdk:"user_username"`
	UserUuid       *string `json:"user_uuid" tfsdk:"user_uuid"`
}

type OfferingReference struct {
	OfferingName *string `json:"offering_name" tfsdk:"offering_name"`
	OfferingUuid *string `json:"offering_uuid" tfsdk:"offering_uuid"`
}

type OfferingReferral struct {
	Creator      *string `json:"creator,omitempty" tfsdk:"creator"`
	Pid          *string `json:"pid,omitempty" tfsdk:"pid"`
	Published    *string `json:"published,omitempty" tfsdk:"published"`
	Publisher    *string `json:"publisher,omitempty" tfsdk:"publisher"`
	ReferralUrl  *string `json:"referral_url,omitempty" tfsdk:"referral_url"`
	RelationType *string `json:"relation_type,omitempty" tfsdk:"relation_type"`
	ResourceType *string `json:"resource_type,omitempty" tfsdk:"resource_type"`
	Scope        *string `json:"scope" tfsdk:"scope"`
	ScopeUuid    *string `json:"scope_uuid" tfsdk:"scope_uuid"`
	Title        *string `json:"title,omitempty" tfsdk:"title"`
	Url          *string `json:"url" tfsdk:"url"`
}

type OfferingResourceOptionsUpdateRequest struct {
	ResourceOptions *OfferingOptionsRequest `json:"resource_options" tfsdk:"resource_options"`
}

type OfferingSoftwareCatalog struct {
	Catalog        *string `json:"catalog" tfsdk:"catalog"`
	CatalogName    *string `json:"catalog_name" tfsdk:"catalog_name"`
	CatalogVersion *string `json:"catalog_version" tfsdk:"catalog_version"`
	Created        *string `json:"created" tfsdk:"created"`
	Modified       *string `json:"modified" tfsdk:"modified"`
	Offering       *string `json:"offering" tfsdk:"offering"`
	OfferingName   *string `json:"offering_name" tfsdk:"offering_name"`
	Partition      *string `json:"partition,omitempty" tfsdk:"partition"`
	PartitionName  *string `json:"partition_name" tfsdk:"partition_name"`
}

type OfferingSoftwareCatalogRequest struct {
	Catalog   *string `json:"catalog" tfsdk:"catalog"`
	Offering  *string `json:"offering" tfsdk:"offering"`
	Partition *string `json:"partition,omitempty" tfsdk:"partition"`
}

type OfferingState struct {
}

type OfferingStats struct {
	Count   *int64  `json:"count" tfsdk:"count"`
	Country *string `json:"country" tfsdk:"country"`
	Name    *string `json:"name" tfsdk:"name"`
}

type OfferingStatsCounter struct {
	CategoryTitle       *string `json:"category_title" tfsdk:"category_title"`
	CategoryUuid        *string `json:"category_uuid" tfsdk:"category_uuid"`
	Count               *int64  `json:"count" tfsdk:"count"`
	ServiceProviderName *string `json:"service_provider_name" tfsdk:"service_provider_name"`
	ServiceProviderUuid *string `json:"service_provider_uuid" tfsdk:"service_provider_uuid"`
}

type OfferingTermsOfService struct {
	Created            *string `json:"created" tfsdk:"created"`
	GracePeriodDays    *int64  `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`
	HasUserConsent     *bool   `json:"has_user_consent" tfsdk:"has_user_consent"`
	IsActive           *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	Modified           *string `json:"modified" tfsdk:"modified"`
	OfferingName       *string `json:"offering_name" tfsdk:"offering_name"`
	OfferingUuid       *string `json:"offering_uuid" tfsdk:"offering_uuid"`
	RequiresReconsent  *bool   `json:"requires_reconsent,omitempty" tfsdk:"requires_reconsent"`
	TermsOfService     *string `json:"terms_of_service,omitempty" tfsdk:"terms_of_service"`
	TermsOfServiceLink *string `json:"terms_of_service_link,omitempty" tfsdk:"terms_of_service_link"`
	Version            *string `json:"version,omitempty" tfsdk:"version"`
}

type OfferingTermsOfServiceCreate struct {
	GracePeriodDays    *int64  `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`
	IsActive           *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	Offering           *string `json:"offering" tfsdk:"offering"`
	RequiresReconsent  *bool   `json:"requires_reconsent,omitempty" tfsdk:"requires_reconsent"`
	TermsOfService     *string `json:"terms_of_service,omitempty" tfsdk:"terms_of_service"`
	TermsOfServiceLink *string `json:"terms_of_service_link,omitempty" tfsdk:"terms_of_service_link"`
	Version            *string `json:"version,omitempty" tfsdk:"version"`
}

type OfferingTermsOfServiceCreateRequest struct {
	GracePeriodDays    *int64  `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`
	IsActive           *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	Offering           *string `json:"offering" tfsdk:"offering"`
	RequiresReconsent  *bool   `json:"requires_reconsent,omitempty" tfsdk:"requires_reconsent"`
	TermsOfService     *string `json:"terms_of_service,omitempty" tfsdk:"terms_of_service"`
	TermsOfServiceLink *string `json:"terms_of_service_link,omitempty" tfsdk:"terms_of_service_link"`
	Version            *string `json:"version,omitempty" tfsdk:"version"`
}

type OfferingTermsOfServiceRequest struct {
	GracePeriodDays    *int64  `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`
	IsActive           *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	TermsOfService     *string `json:"terms_of_service,omitempty" tfsdk:"terms_of_service"`
	TermsOfServiceLink *string `json:"terms_of_service_link,omitempty" tfsdk:"terms_of_service_link"`
}

type OfferingThumbnailRequest struct {
	Thumbnail *string `json:"thumbnail" tfsdk:"thumbnail"`
}

type OfferingThumbnailRequestForm struct {
	Thumbnail *string `json:"thumbnail" tfsdk:"thumbnail"`
}

type OfferingThumbnailRequestMultipart struct {
	Thumbnail *string `json:"thumbnail" tfsdk:"thumbnail"`
}

type OfferingUsagePolicy struct {
	Actions            *string                        `json:"actions" tfsdk:"actions"`
	ApplyToAll         *bool                          `json:"apply_to_all,omitempty" tfsdk:"apply_to_all"`
	ComponentLimitsSet []NestedOfferingComponentLimit `json:"component_limits_set" tfsdk:"component_limits_set"`
	Created            *string                        `json:"created" tfsdk:"created"`
	CreatedByFullName  *string                        `json:"created_by_full_name" tfsdk:"created_by_full_name"`
	CreatedByUsername  *string                        `json:"created_by_username" tfsdk:"created_by_username"`
	FiredDatetime      *string                        `json:"fired_datetime" tfsdk:"fired_datetime"`
	HasFired           *bool                          `json:"has_fired" tfsdk:"has_fired"`
	OrganizationGroups []string                       `json:"organization_groups,omitempty" tfsdk:"organization_groups"`
	Period             *int64                         `json:"period,omitempty" tfsdk:"period"`
	PeriodName         *string                        `json:"period_name" tfsdk:"period_name"`
	Scope              *string                        `json:"scope" tfsdk:"scope"`
	ScopeName          *string                        `json:"scope_name" tfsdk:"scope_name"`
	ScopeUuid          *string                        `json:"scope_uuid" tfsdk:"scope_uuid"`
	Url                *string                        `json:"url" tfsdk:"url"`
}

type OfferingUsagePolicyRequest struct {
	Actions            *string                               `json:"actions" tfsdk:"actions"`
	ApplyToAll         *bool                                 `json:"apply_to_all,omitempty" tfsdk:"apply_to_all"`
	ComponentLimitsSet []NestedOfferingComponentLimitRequest `json:"component_limits_set" tfsdk:"component_limits_set"`
	OrganizationGroups []string                              `json:"organization_groups,omitempty" tfsdk:"organization_groups"`
	Period             *int64                                `json:"period,omitempty" tfsdk:"period"`
	Scope              *string                               `json:"scope" tfsdk:"scope"`
}

type OfferingUser struct {
	Created                   *string `json:"created,omitempty" tfsdk:"created"`
	CustomerName              *string `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerUuid              *string `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	HasComplianceChecklist    *bool   `json:"has_compliance_checklist,omitempty" tfsdk:"has_compliance_checklist"`
	HasConsent                *bool   `json:"has_consent,omitempty" tfsdk:"has_consent"`
	IsRestricted              *bool   `json:"is_restricted,omitempty" tfsdk:"is_restricted"`
	Modified                  *string `json:"modified,omitempty" tfsdk:"modified"`
	Offering                  *string `json:"offering,omitempty" tfsdk:"offering"`
	OfferingName              *string `json:"offering_name,omitempty" tfsdk:"offering_name"`
	OfferingUuid              *string `json:"offering_uuid,omitempty" tfsdk:"offering_uuid"`
	RequiresReconsent         *bool   `json:"requires_reconsent,omitempty" tfsdk:"requires_reconsent"`
	ServiceProviderComment    *string `json:"service_provider_comment,omitempty" tfsdk:"service_provider_comment"`
	ServiceProviderCommentUrl *string `json:"service_provider_comment_url,omitempty" tfsdk:"service_provider_comment_url"`
	State                     *string `json:"state,omitempty" tfsdk:"state"`
	Url                       *string `json:"url,omitempty" tfsdk:"url"`
	User                      *string `json:"user,omitempty" tfsdk:"user"`
	UserEmail                 *string `json:"user_email,omitempty" tfsdk:"user_email"`
	UserFullName              *string `json:"user_full_name,omitempty" tfsdk:"user_full_name"`
	UserUsername              *string `json:"user_username,omitempty" tfsdk:"user_username"`
	UserUuid                  *string `json:"user_uuid,omitempty" tfsdk:"user_uuid"`
	Username                  *string `json:"username,omitempty" tfsdk:"username"`
}

type OfferingUserRequest struct {
	Offering     *string `json:"offering,omitempty" tfsdk:"offering"`
	OfferingUuid *string `json:"offering_uuid,omitempty" tfsdk:"offering_uuid"`
	User         *string `json:"user,omitempty" tfsdk:"user"`
	UserUuid     *string `json:"user_uuid,omitempty" tfsdk:"user_uuid"`
	Username     *string `json:"username,omitempty" tfsdk:"username"`
}

type OfferingUserRole struct {
	Name         *string `json:"name" tfsdk:"name"`
	Offering     *string `json:"offering" tfsdk:"offering"`
	OfferingName *string `json:"offering_name" tfsdk:"offering_name"`
	OfferingUuid *string `json:"offering_uuid" tfsdk:"offering_uuid"`
}

type OfferingUserRoleRequest struct {
	Name     *string `json:"name" tfsdk:"name"`
	Offering *string `json:"offering" tfsdk:"offering"`
}

type OfferingUserServiceProviderComment struct {
	ServiceProviderComment    *string `json:"service_provider_comment,omitempty" tfsdk:"service_provider_comment"`
	ServiceProviderCommentUrl *string `json:"service_provider_comment_url,omitempty" tfsdk:"service_provider_comment_url"`
}

type OfferingUserState struct {
}

type OfferingUserStateTransitionRequest struct {
	Comment    *string `json:"comment,omitempty" tfsdk:"comment"`
	CommentUrl *string `json:"comment_url,omitempty" tfsdk:"comment_url"`
}

type OfferingUserUpdateRestrictionRequest struct {
	IsRestricted *bool `json:"is_restricted" tfsdk:"is_restricted"`
}

type OnboardingCompanyValidationRequestRequest struct {
	Country               *string `json:"country" tfsdk:"country"`
	IsManualValidation    *bool   `json:"is_manual_validation,omitempty" tfsdk:"is_manual_validation"`
	LegalName             *string `json:"legal_name,omitempty" tfsdk:"legal_name"`
	LegalPersonIdentifier *string `json:"legal_person_identifier,omitempty" tfsdk:"legal_person_identifier"`
}

type OnboardingCountryChecklistConfiguration struct {
	Checklist     *string         `json:"checklist" tfsdk:"checklist"`
	ChecklistName *string         `json:"checklist_name" tfsdk:"checklist_name"`
	ChecklistUuid *string         `json:"checklist_uuid" tfsdk:"checklist_uuid"`
	Country       *string         `json:"country" tfsdk:"country"`
	Created       *string         `json:"created" tfsdk:"created"`
	IsActive      *bool           `json:"is_active,omitempty" tfsdk:"is_active"`
	Modified      *string         `json:"modified" tfsdk:"modified"`
	Questions     []QuestionAdmin `json:"questions" tfsdk:"questions"`
	Url           *string         `json:"url" tfsdk:"url"`
}

type OnboardingCountryChecklistConfigurationRequest struct {
	Checklist *string `json:"checklist" tfsdk:"checklist"`
	Country   *string `json:"country" tfsdk:"country"`
	IsActive  *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
}

type OnboardingJustification struct {
	Country                 *string                                `json:"country" tfsdk:"country"`
	Created                 *string                                `json:"created" tfsdk:"created"`
	ErrorMessage            *string                                `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback          *string                                `json:"error_traceback" tfsdk:"error_traceback"`
	LegalName               *string                                `json:"legal_name" tfsdk:"legal_name"`
	LegalPersonIdentifier   *string                                `json:"legal_person_identifier" tfsdk:"legal_person_identifier"`
	Modified                *string                                `json:"modified" tfsdk:"modified"`
	StaffNotes              *string                                `json:"staff_notes" tfsdk:"staff_notes"`
	SupportingDocumentation []OnboardingJustificationDocumentation `json:"supporting_documentation" tfsdk:"supporting_documentation"`
	User                    *string                                `json:"user" tfsdk:"user"`
	UserFullName            *string                                `json:"user_full_name" tfsdk:"user_full_name"`
	UserJustification       *string                                `json:"user_justification,omitempty" tfsdk:"user_justification"`
	ValidatedAt             *string                                `json:"validated_at" tfsdk:"validated_at"`
	ValidatedBy             *string                                `json:"validated_by" tfsdk:"validated_by"`
	ValidationDecision      *string                                `json:"validation_decision" tfsdk:"validation_decision"`
	Verification            *string                                `json:"verification" tfsdk:"verification"`
	VerificationUuid        *string                                `json:"verification_uuid" tfsdk:"verification_uuid"`
}

type OnboardingJustificationCreateRequest struct {
	UserJustification *string `json:"user_justification,omitempty" tfsdk:"user_justification"`
	VerificationUuid  *string `json:"verification_uuid" tfsdk:"verification_uuid"`
}

type OnboardingJustificationDocumentation struct {
	Created  *string `json:"created" tfsdk:"created"`
	File     *string `json:"file,omitempty" tfsdk:"file"`
	FileName *string `json:"file_name" tfsdk:"file_name"`
	FileSize *int64  `json:"file_size" tfsdk:"file_size"`
}

type OnboardingJustificationDocumentationRequest struct {
	File *string `json:"file,omitempty" tfsdk:"file"`
}

type OnboardingJustificationDocumentationRequestForm struct {
	File *string `json:"file,omitempty" tfsdk:"file"`
}

type OnboardingJustificationDocumentationRequestMultipart struct {
	File *string `json:"file,omitempty" tfsdk:"file"`
}

type OnboardingJustificationRequest struct {
	UserJustification *string `json:"user_justification,omitempty" tfsdk:"user_justification"`
	Verification      *string `json:"verification" tfsdk:"verification"`
}

type OnboardingJustificationReviewRequest struct {
	StaffNotes *string `json:"staff_notes,omitempty" tfsdk:"staff_notes"`
}

type OnboardingQuestionMetadata struct {
	ChecklistName       *string `json:"checklist_name" tfsdk:"checklist_name"`
	Created             *string `json:"created" tfsdk:"created"`
	IntentField         *string `json:"intent_field,omitempty" tfsdk:"intent_field"`
	MapsToCustomerField *string `json:"maps_to_customer_field,omitempty" tfsdk:"maps_to_customer_field"`
	Modified            *string `json:"modified" tfsdk:"modified"`
	Question            *string `json:"question" tfsdk:"question"`
	QuestionDescription *string `json:"question_description" tfsdk:"question_description"`
	QuestionUuid        *string `json:"question_uuid" tfsdk:"question_uuid"`
	Url                 *string `json:"url" tfsdk:"url"`
}

type OnboardingQuestionMetadataRequest struct {
	IntentField         *string `json:"intent_field,omitempty" tfsdk:"intent_field"`
	MapsToCustomerField *string `json:"maps_to_customer_field,omitempty" tfsdk:"maps_to_customer_field"`
	Question            *string `json:"question" tfsdk:"question"`
}

type OnboardingRunValidationRequestRequest struct {
	BirthDate        *string `json:"birth_date,omitempty" tfsdk:"birth_date"`
	FirstName        *string `json:"first_name,omitempty" tfsdk:"first_name"`
	LastName         *string `json:"last_name,omitempty" tfsdk:"last_name"`
	PersonIdentifier *string `json:"person_identifier,omitempty" tfsdk:"person_identifier"`
}

type OnboardingVerification struct {
	CanCustomerBeCreated         *bool                     `json:"can_customer_be_created" tfsdk:"can_customer_be_created"`
	Country                      *string                   `json:"country" tfsdk:"country"`
	Created                      *string                   `json:"created" tfsdk:"created"`
	Customer                     *string                   `json:"customer" tfsdk:"customer"`
	CustomerCreationErrorMessage *string                   `json:"customer_creation_error_message" tfsdk:"customer_creation_error_message"`
	ErrorMessage                 *string                   `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback               *string                   `json:"error_traceback" tfsdk:"error_traceback"`
	ExpiresAt                    *string                   `json:"expires_at,omitempty" tfsdk:"expires_at"`
	Justifications               []OnboardingJustification `json:"justifications" tfsdk:"justifications"`
	LegalName                    *string                   `json:"legal_name,omitempty" tfsdk:"legal_name"`
	LegalPersonIdentifier        *string                   `json:"legal_person_identifier,omitempty" tfsdk:"legal_person_identifier"`
	Modified                     *string                   `json:"modified" tfsdk:"modified"`
	Status                       *string                   `json:"status" tfsdk:"status"`
	User                         *string                   `json:"user" tfsdk:"user"`
	UserFullName                 *string                   `json:"user_full_name" tfsdk:"user_full_name"`
	ValidatedAt                  *string                   `json:"validated_at" tfsdk:"validated_at"`
	ValidationMethod             *string                   `json:"validation_method" tfsdk:"validation_method"`
}

type OnboardingVerificationRequest struct {
	Country               *string `json:"country" tfsdk:"country"`
	ExpiresAt             *string `json:"expires_at,omitempty" tfsdk:"expires_at"`
	LegalName             *string `json:"legal_name,omitempty" tfsdk:"legal_name"`
	LegalPersonIdentifier *string `json:"legal_person_identifier,omitempty" tfsdk:"legal_person_identifier"`
}

type OnboardingVerificationStatusEnum struct {
}

type OpenStackAllowedAddressPair struct {
	MacAddress *string `json:"mac_address,omitempty" tfsdk:"mac_address"`
}

type OpenStackAllowedAddressPairRequest struct {
	IpAddress  *string `json:"ip_address,omitempty" tfsdk:"ip_address"`
	MacAddress *string `json:"mac_address,omitempty" tfsdk:"mac_address"`
}

type OpenStackBackendInstance struct {
	AvailabilityZone   *string `json:"availability_zone" tfsdk:"availability_zone"`
	BackendId          *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created            *string `json:"created" tfsdk:"created"`
	HypervisorHostname *string `json:"hypervisor_hostname,omitempty" tfsdk:"hypervisor_hostname"`
	KeyName            *string `json:"key_name,omitempty" tfsdk:"key_name"`
	Name               *string `json:"name" tfsdk:"name"`
	RuntimeState       *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	StartTime          *string `json:"start_time,omitempty" tfsdk:"start_time"`
	State              *string `json:"state" tfsdk:"state"`
}

type OpenStackBackendVolumes struct {
	AvailabilityZone *string `json:"availability_zone" tfsdk:"availability_zone"`
	BackendId        *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Bootable         *bool   `json:"bootable,omitempty" tfsdk:"bootable"`
	Description      *string `json:"description,omitempty" tfsdk:"description"`
	Metadata         *string `json:"metadata,omitempty" tfsdk:"metadata"`
	Name             *string `json:"name" tfsdk:"name"`
	RuntimeState     *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	Size             *int64  `json:"size" tfsdk:"size"`
	State            *string `json:"state" tfsdk:"state"`
	Type             *string `json:"type" tfsdk:"type"`
}

type OpenStackBackup struct {
	AccessUrl                   *string                        `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string                        `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created                     *string                        `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string                        `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string                        `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string                        `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string                        `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string                        `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string                        `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string                        `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string                        `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	Instance                    *string                        `json:"instance,omitempty" tfsdk:"instance"`
	InstanceFloatingIps         []OpenStackNestedFloatingIP    `json:"instance_floating_ips,omitempty" tfsdk:"instance_floating_ips"`
	InstanceMarketplaceUuid     *string                        `json:"instance_marketplace_uuid,omitempty" tfsdk:"instance_marketplace_uuid"`
	InstanceName                *string                        `json:"instance_name,omitempty" tfsdk:"instance_name"`
	InstancePorts               []OpenStackNestedPort          `json:"instance_ports,omitempty" tfsdk:"instance_ports"`
	InstanceSecurityGroups      []OpenStackNestedSecurityGroup `json:"instance_security_groups,omitempty" tfsdk:"instance_security_groups"`
	IsLimitBased                *bool                          `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                          `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	KeptUntil                   *string                        `json:"kept_until,omitempty" tfsdk:"kept_until"`
	MarketplaceCategoryName     *string                        `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                        `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                        `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                        `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                        `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                        `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                        `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                        `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string                        `json:"name,omitempty" tfsdk:"name"`
	Project                     *string                        `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string                        `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string                        `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ResourceType                *string                        `json:"resource_type,omitempty" tfsdk:"resource_type"`
	Restorations                []OpenStackBackupRestoration   `json:"restorations,omitempty" tfsdk:"restorations"`
	ServiceName                 *string                        `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string                        `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                        `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                        `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                        `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string                        `json:"state,omitempty" tfsdk:"state"`
	TenantUuid                  *string                        `json:"tenant_uuid,omitempty" tfsdk:"tenant_uuid"`
	Url                         *string                        `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackBackupRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	KeptUntil   *string `json:"kept_until,omitempty" tfsdk:"kept_until"`
	Name        *string `json:"name" tfsdk:"name"`
}

type OpenStackBackupRestoration struct {
	Created        *string                        `json:"created,omitempty" tfsdk:"created"`
	Flavor         *string                        `json:"flavor,omitempty" tfsdk:"flavor"`
	FloatingIps    []OpenStackNestedFloatingIP    `json:"floating_ips,omitempty" tfsdk:"floating_ips"`
	Instance       *string                        `json:"instance,omitempty" tfsdk:"instance"`
	Name           *string                        `json:"name,omitempty" tfsdk:"name"`
	Ports          []OpenStackNestedPort          `json:"ports,omitempty" tfsdk:"ports"`
	SecurityGroups []OpenStackNestedSecurityGroup `json:"security_groups,omitempty" tfsdk:"security_groups"`
}

type OpenStackBackupRestorationCreateRequest struct {
	Flavor         *string                                  `json:"flavor" tfsdk:"flavor"`
	FloatingIps    []OpenStackCreateFloatingIPRequest       `json:"floating_ips,omitempty" tfsdk:"floating_ips"`
	Name           *string                                  `json:"name,omitempty" tfsdk:"name"`
	Ports          []OpenStackCreatePortRequest             `json:"ports,omitempty" tfsdk:"ports"`
	SecurityGroups []OpenStackSecurityGroupHyperlinkRequest `json:"security_groups,omitempty" tfsdk:"security_groups"`
}

type OpenStackBackupRestorationRequest struct {
	Flavor      *string                            `json:"flavor" tfsdk:"flavor"`
	FloatingIps []OpenStackNestedFloatingIPRequest `json:"floating_ips" tfsdk:"floating_ips"`
	Name        *string                            `json:"name,omitempty" tfsdk:"name"`
	Ports       []OpenStackNestedPortRequest       `json:"ports" tfsdk:"ports"`
}

type OpenStackCreateFloatingIPRequest struct {
	IpAddress *string `json:"ip_address,omitempty" tfsdk:"ip_address"`
	Subnet    *string `json:"subnet" tfsdk:"subnet"`
	Url       *string `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackCreateInstancePortRequest struct {
	FixedIps []OpenStackFixedIpRequest `json:"fixed_ips,omitempty" tfsdk:"fixed_ips"`
	Port     *string                   `json:"port,omitempty" tfsdk:"port"`
	Subnet   *string                   `json:"subnet,omitempty" tfsdk:"subnet"`
}

type OpenStackCreatePortRequest struct {
	FixedIps []OpenStackFixedIpRequest `json:"fixed_ips,omitempty" tfsdk:"fixed_ips"`
	Port     *string                   `json:"port,omitempty" tfsdk:"port"`
	Subnet   *string                   `json:"subnet,omitempty" tfsdk:"subnet"`
	Tenant   *string                   `json:"tenant,omitempty" tfsdk:"tenant"`
}

type OpenStackDataVolumeRequest struct {
	Size       *int64  `json:"size" tfsdk:"size"`
	VolumeType *string `json:"volume_type,omitempty" tfsdk:"volume_type"`
}

type OpenStackFixedIp struct {
	IpAddress *string `json:"ip_address,omitempty" tfsdk:"ip_address"`
	SubnetId  *string `json:"subnet_id,omitempty" tfsdk:"subnet_id"`
}

type OpenStackFixedIpRequest struct {
	IpAddress *string `json:"ip_address" tfsdk:"ip_address"`
	SubnetId  *string `json:"subnet_id" tfsdk:"subnet_id"`
}

type OpenStackFlavor struct {
	BackendId   *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Cores       *int64  `json:"cores,omitempty" tfsdk:"cores"`
	Disk        *int64  `json:"disk,omitempty" tfsdk:"disk"`
	DisplayName *string `json:"display_name,omitempty" tfsdk:"display_name"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
	Ram         *int64  `json:"ram,omitempty" tfsdk:"ram"`
	Settings    *string `json:"settings,omitempty" tfsdk:"settings"`
	Url         *string `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackFloatingIP struct {
	AccessUrl                   *string            `json:"access_url,omitempty" tfsdk:"access_url"`
	Address                     *string            `json:"address,omitempty" tfsdk:"address"`
	BackendId                   *string            `json:"backend_id,omitempty" tfsdk:"backend_id"`
	BackendNetworkId            *string            `json:"backend_network_id,omitempty" tfsdk:"backend_network_id"`
	Created                     *string            `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string            `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string            `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string            `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string            `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string            `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string            `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string            `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string            `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	ExternalAddress             *string            `json:"external_address,omitempty" tfsdk:"external_address"`
	InstanceName                *string            `json:"instance_name,omitempty" tfsdk:"instance_name"`
	InstanceUrl                 *string            `json:"instance_url,omitempty" tfsdk:"instance_url"`
	InstanceUuid                *string            `json:"instance_uuid,omitempty" tfsdk:"instance_uuid"`
	IsLimitBased                *bool              `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool              `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string            `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string            `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string            `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string            `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string            `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string            `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string            `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string            `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string            `json:"name,omitempty" tfsdk:"name"`
	Port                        *string            `json:"port,omitempty" tfsdk:"port"`
	PortFixedIps                []OpenStackFixedIp `json:"port_fixed_ips,omitempty" tfsdk:"port_fixed_ips"`
	Project                     *string            `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string            `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string            `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ResourceType                *string            `json:"resource_type,omitempty" tfsdk:"resource_type"`
	RuntimeState                *string            `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceName                 *string            `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string            `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string            `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string            `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string            `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string            `json:"state,omitempty" tfsdk:"state"`
	Tenant                      *string            `json:"tenant,omitempty" tfsdk:"tenant"`
	TenantName                  *string            `json:"tenant_name,omitempty" tfsdk:"tenant_name"`
	TenantUuid                  *string            `json:"tenant_uuid,omitempty" tfsdk:"tenant_uuid"`
	Url                         *string            `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackFloatingIPAttachRequest struct {
	Port *string `json:"port" tfsdk:"port"`
}

type OpenStackFloatingIPDescriptionUpdateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
}

type OpenStackImage struct {
	BackendId *string `json:"backend_id" tfsdk:"backend_id"`
	MinDisk   *int64  `json:"min_disk,omitempty" tfsdk:"min_disk"`
	MinRam    *int64  `json:"min_ram,omitempty" tfsdk:"min_ram"`
	Name      *string `json:"name" tfsdk:"name"`
	Settings  *string `json:"settings" tfsdk:"settings"`
	Url       *string `json:"url" tfsdk:"url"`
}

type OpenStackInstance struct {
	AccessUrl                        *string                        `json:"access_url,omitempty" tfsdk:"access_url"`
	Action                           *string                        `json:"action,omitempty" tfsdk:"action"`
	AvailabilityZone                 *string                        `json:"availability_zone,omitempty" tfsdk:"availability_zone"`
	AvailabilityZoneName             *string                        `json:"availability_zone_name,omitempty" tfsdk:"availability_zone_name"`
	BackendId                        *string                        `json:"backend_id,omitempty" tfsdk:"backend_id"`
	ConnectDirectlyToExternalNetwork *bool                          `json:"connect_directly_to_external_network,omitempty" tfsdk:"connect_directly_to_external_network"`
	Cores                            *int64                         `json:"cores,omitempty" tfsdk:"cores"`
	Created                          *string                        `json:"created,omitempty" tfsdk:"created"`
	Customer                         *string                        `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation             *string                        `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                     *string                        `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName               *string                        `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                     *string                        `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                      *string                        `json:"description,omitempty" tfsdk:"description"`
	Disk                             *int64                         `json:"disk,omitempty" tfsdk:"disk"`
	ErrorMessage                     *string                        `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback                   *string                        `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	ExternalAddress                  []string                       `json:"external_address,omitempty" tfsdk:"external_address"`
	ExternalIps                      []string                       `json:"external_ips,omitempty" tfsdk:"external_ips"`
	FlavorDisk                       *int64                         `json:"flavor_disk,omitempty" tfsdk:"flavor_disk"`
	FlavorName                       *string                        `json:"flavor_name,omitempty" tfsdk:"flavor_name"`
	FloatingIps                      []OpenStackNestedFloatingIP    `json:"floating_ips,omitempty" tfsdk:"floating_ips"`
	HypervisorHostname               *string                        `json:"hypervisor_hostname,omitempty" tfsdk:"hypervisor_hostname"`
	ImageName                        *string                        `json:"image_name,omitempty" tfsdk:"image_name"`
	InternalIps                      []string                       `json:"internal_ips,omitempty" tfsdk:"internal_ips"`
	IsLimitBased                     *bool                          `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                     *bool                          `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	KeyFingerprint                   *string                        `json:"key_fingerprint,omitempty" tfsdk:"key_fingerprint"`
	KeyName                          *string                        `json:"key_name,omitempty" tfsdk:"key_name"`
	Latitude                         *float64                       `json:"latitude,omitempty" tfsdk:"latitude"`
	Longitude                        *float64                       `json:"longitude,omitempty" tfsdk:"longitude"`
	MarketplaceCategoryName          *string                        `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid          *string                        `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName          *string                        `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid          *string                        `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid              *string                        `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState         *string                        `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid          *string                        `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	MinDisk                          *int64                         `json:"min_disk,omitempty" tfsdk:"min_disk"`
	MinRam                           *int64                         `json:"min_ram,omitempty" tfsdk:"min_ram"`
	Modified                         *string                        `json:"modified,omitempty" tfsdk:"modified"`
	Name                             *string                        `json:"name,omitempty" tfsdk:"name"`
	Ports                            []OpenStackNestedPort          `json:"ports,omitempty" tfsdk:"ports"`
	Project                          *string                        `json:"project,omitempty" tfsdk:"project"`
	ProjectName                      *string                        `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                      *string                        `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	Ram                              *int64                         `json:"ram,omitempty" tfsdk:"ram"`
	ResourceType                     *string                        `json:"resource_type,omitempty" tfsdk:"resource_type"`
	RuntimeState                     *string                        `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	SecurityGroups                   []OpenStackNestedSecurityGroup `json:"security_groups,omitempty" tfsdk:"security_groups"`
	ServerGroup                      *OpenStackNestedServerGroup    `json:"server_group,omitempty" tfsdk:"server_group"`
	ServiceName                      *string                        `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings                  *string                        `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage      *string                        `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState             *string                        `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid              *string                        `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	StartTime                        *string                        `json:"start_time,omitempty" tfsdk:"start_time"`
	State                            *string                        `json:"state,omitempty" tfsdk:"state"`
	Tenant                           *string                        `json:"tenant,omitempty" tfsdk:"tenant"`
	TenantUuid                       *string                        `json:"tenant_uuid,omitempty" tfsdk:"tenant_uuid"`
	Url                              *string                        `json:"url,omitempty" tfsdk:"url"`
	UserData                         *string                        `json:"user_data,omitempty" tfsdk:"user_data"`
	Volumes                          []OpenStackNestedVolume        `json:"volumes,omitempty" tfsdk:"volumes"`
}

type OpenStackInstanceAllowedAddressPairsUpdateRequest struct {
	AllowedAddressPairs []OpenStackAllowedAddressPairRequest `json:"allowed_address_pairs" tfsdk:"allowed_address_pairs"`
	Subnet              *string                              `json:"subnet" tfsdk:"subnet"`
}

type OpenStackInstanceAvailabilityZone struct {
	Available *bool   `json:"available,omitempty" tfsdk:"available"`
	Name      *string `json:"name" tfsdk:"name"`
	Settings  *string `json:"settings,omitempty" tfsdk:"settings"`
	Url       *string `json:"url" tfsdk:"url"`
}

type OpenStackInstanceCreateOrderAttributes struct {
	AvailabilityZone                 *string                                  `json:"availability_zone,omitempty" tfsdk:"availability_zone"`
	ConnectDirectlyToExternalNetwork *bool                                    `json:"connect_directly_to_external_network,omitempty" tfsdk:"connect_directly_to_external_network"`
	DataVolumeSize                   *int64                                   `json:"data_volume_size,omitempty" tfsdk:"data_volume_size"`
	DataVolumeType                   *string                                  `json:"data_volume_type,omitempty" tfsdk:"data_volume_type"`
	DataVolumes                      []OpenStackDataVolumeRequest             `json:"data_volumes,omitempty" tfsdk:"data_volumes"`
	Description                      *string                                  `json:"description,omitempty" tfsdk:"description"`
	Flavor                           *string                                  `json:"flavor" tfsdk:"flavor"`
	FloatingIps                      []OpenStackCreateFloatingIPRequest       `json:"floating_ips,omitempty" tfsdk:"floating_ips"`
	Image                            *string                                  `json:"image" tfsdk:"image"`
	Name                             *string                                  `json:"name" tfsdk:"name"`
	Ports                            []OpenStackCreateInstancePortRequest     `json:"ports" tfsdk:"ports"`
	SecurityGroups                   []OpenStackSecurityGroupHyperlinkRequest `json:"security_groups,omitempty" tfsdk:"security_groups"`
	SshPublicKey                     *string                                  `json:"ssh_public_key,omitempty" tfsdk:"ssh_public_key"`
	SystemVolumeSize                 *int64                                   `json:"system_volume_size" tfsdk:"system_volume_size"`
	SystemVolumeType                 *string                                  `json:"system_volume_type,omitempty" tfsdk:"system_volume_type"`
	UserData                         *string                                  `json:"user_data,omitempty" tfsdk:"user_data"`
}

type OpenStackInstanceFloatingIPsUpdateRequest struct {
	FloatingIps []OpenStackCreateFloatingIPRequest `json:"floating_ips,omitempty" tfsdk:"floating_ips"`
}

type OpenStackInstancePortsUpdateRequest struct {
	Ports []OpenStackCreatePortRequest `json:"ports" tfsdk:"ports"`
}

type OpenStackInstanceRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
}

type OpenStackInstanceSecurityGroupsUpdateRequest struct {
	SecurityGroups []string `json:"security_groups" tfsdk:"security_groups"`
}

type OpenStackNestedFloatingIP struct {
	Address           *string            `json:"address,omitempty" tfsdk:"address"`
	PortFixedIps      []OpenStackFixedIp `json:"port_fixed_ips,omitempty" tfsdk:"port_fixed_ips"`
	PortMacAddress    *string            `json:"port_mac_address,omitempty" tfsdk:"port_mac_address"`
	Subnet            *string            `json:"subnet,omitempty" tfsdk:"subnet"`
	SubnetCidr        *string            `json:"subnet_cidr,omitempty" tfsdk:"subnet_cidr"`
	SubnetDescription *string            `json:"subnet_description,omitempty" tfsdk:"subnet_description"`
	SubnetName        *string            `json:"subnet_name,omitempty" tfsdk:"subnet_name"`
	SubnetUuid        *string            `json:"subnet_uuid,omitempty" tfsdk:"subnet_uuid"`
	Url               *string            `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackNestedFloatingIPRequest struct {
	Subnet *string `json:"subnet" tfsdk:"subnet"`
}

type OpenStackNestedInstance struct {
	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Name      *string `json:"name,omitempty" tfsdk:"name"`
}

type OpenStackNestedPort struct {
	AllowedAddressPairs []OpenStackAllowedAddressPair `json:"allowed_address_pairs,omitempty" tfsdk:"allowed_address_pairs"`
	DeviceId            *string                       `json:"device_id,omitempty" tfsdk:"device_id"`
	DeviceOwner         *string                       `json:"device_owner,omitempty" tfsdk:"device_owner"`
	FixedIps            []OpenStackFixedIp            `json:"fixed_ips,omitempty" tfsdk:"fixed_ips"`
	MacAddress          *string                       `json:"mac_address,omitempty" tfsdk:"mac_address"`
	SecurityGroups      []OpenStackSecurityGroup      `json:"security_groups,omitempty" tfsdk:"security_groups"`
	Subnet              *string                       `json:"subnet,omitempty" tfsdk:"subnet"`
	SubnetCidr          *string                       `json:"subnet_cidr,omitempty" tfsdk:"subnet_cidr"`
	SubnetDescription   *string                       `json:"subnet_description,omitempty" tfsdk:"subnet_description"`
	SubnetName          *string                       `json:"subnet_name,omitempty" tfsdk:"subnet_name"`
	SubnetUuid          *string                       `json:"subnet_uuid,omitempty" tfsdk:"subnet_uuid"`
	Url                 *string                       `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackNestedPortRequest struct {
	FixedIps []OpenStackFixedIpRequest `json:"fixed_ips,omitempty" tfsdk:"fixed_ips"`
	Subnet   *string                   `json:"subnet,omitempty" tfsdk:"subnet"`
}

type OpenStackNestedSecurityGroup struct {
	Description *string                   `json:"description,omitempty" tfsdk:"description"`
	Name        *string                   `json:"name,omitempty" tfsdk:"name"`
	Rules       []NestedSecurityGroupRule `json:"rules,omitempty" tfsdk:"rules"`
	State       *string                   `json:"state,omitempty" tfsdk:"state"`
	Url         *string                   `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackNestedServerGroup struct {
	Name   *string `json:"name,omitempty" tfsdk:"name"`
	Policy *string `json:"policy,omitempty" tfsdk:"policy"`
	State  *string `json:"state,omitempty" tfsdk:"state"`
	Url    *string `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackNestedSubNet struct {
	AllocationPools []OpenStackSubNetAllocationPool `json:"allocation_pools,omitempty" tfsdk:"allocation_pools"`
	Cidr            *string                         `json:"cidr,omitempty" tfsdk:"cidr"`
	Description     *string                         `json:"description,omitempty" tfsdk:"description"`
	EnableDhcp      *bool                           `json:"enable_dhcp,omitempty" tfsdk:"enable_dhcp"`
	GatewayIp       *string                         `json:"gateway_ip,omitempty" tfsdk:"gateway_ip"`
	IpVersion       *int64                          `json:"ip_version,omitempty" tfsdk:"ip_version"`
	Name            *string                         `json:"name,omitempty" tfsdk:"name"`
}

type OpenStackNestedSubNetRequest struct {
	Cidr        *string `json:"cidr,omitempty" tfsdk:"cidr"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	EnableDhcp  *bool   `json:"enable_dhcp,omitempty" tfsdk:"enable_dhcp"`
	GatewayIp   *string `json:"gateway_ip,omitempty" tfsdk:"gateway_ip"`
	IpVersion   *int64  `json:"ip_version,omitempty" tfsdk:"ip_version"`
	Name        *string `json:"name" tfsdk:"name"`
}

type OpenStackNestedVolume struct {
	Bootable                *bool   `json:"bootable,omitempty" tfsdk:"bootable"`
	Device                  *string `json:"device,omitempty" tfsdk:"device"`
	ImageName               *string `json:"image_name,omitempty" tfsdk:"image_name"`
	MarketplaceResourceUuid *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Name                    *string `json:"name,omitempty" tfsdk:"name"`
	ResourceType            *string `json:"resource_type,omitempty" tfsdk:"resource_type"`
	Size                    *int64  `json:"size,omitempty" tfsdk:"size"`
	State                   *string `json:"state,omitempty" tfsdk:"state"`
	Type                    *string `json:"type,omitempty" tfsdk:"type"`
	TypeName                *string `json:"type_name,omitempty" tfsdk:"type_name"`
	Url                     *string `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackNestedVolumeRequest struct {
	Bootable  *bool   `json:"bootable,omitempty" tfsdk:"bootable"`
	Device    *string `json:"device,omitempty" tfsdk:"device"`
	ImageName *string `json:"image_name,omitempty" tfsdk:"image_name"`
	Size      *int64  `json:"size" tfsdk:"size"`
	Type      *string `json:"type,omitempty" tfsdk:"type"`
}

type OpenStackNetwork struct {
	AccessUrl                   *string                 `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string                 `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created                     *string                 `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string                 `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string                 `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string                 `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string                 `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string                 `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string                 `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string                 `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string                 `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	IsExternal                  *bool                   `json:"is_external,omitempty" tfsdk:"is_external"`
	IsLimitBased                *bool                   `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                   `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string                 `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                 `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                 `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                 `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                 `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                 `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                 `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                 `json:"modified,omitempty" tfsdk:"modified"`
	Mtu                         *int64                  `json:"mtu,omitempty" tfsdk:"mtu"`
	Name                        *string                 `json:"name,omitempty" tfsdk:"name"`
	Project                     *string                 `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string                 `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string                 `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	RbacPolicies                []NetworkRBACPolicy     `json:"rbac_policies,omitempty" tfsdk:"rbac_policies"`
	ResourceType                *string                 `json:"resource_type,omitempty" tfsdk:"resource_type"`
	SegmentationId              *int64                  `json:"segmentation_id,omitempty" tfsdk:"segmentation_id"`
	ServiceName                 *string                 `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string                 `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                 `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                 `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                 `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string                 `json:"state,omitempty" tfsdk:"state"`
	Subnets                     []OpenStackNestedSubNet `json:"subnets,omitempty" tfsdk:"subnets"`
	Tenant                      *string                 `json:"tenant,omitempty" tfsdk:"tenant"`
	TenantName                  *string                 `json:"tenant_name,omitempty" tfsdk:"tenant_name"`
	TenantUuid                  *string                 `json:"tenant_uuid,omitempty" tfsdk:"tenant_uuid"`
	Type                        *string                 `json:"type,omitempty" tfsdk:"type"`
	Url                         *string                 `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackNetworkRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
}

type OpenStackPort struct {
	AccessUrl                   *string                            `json:"access_url,omitempty" tfsdk:"access_url"`
	AdminStateUp                *bool                              `json:"admin_state_up,omitempty" tfsdk:"admin_state_up"`
	AllowedAddressPairs         []OpenStackAllowedAddressPair      `json:"allowed_address_pairs,omitempty" tfsdk:"allowed_address_pairs"`
	BackendId                   *string                            `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created                     *string                            `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string                            `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string                            `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string                            `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string                            `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string                            `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string                            `json:"description,omitempty" tfsdk:"description"`
	DeviceId                    *string                            `json:"device_id,omitempty" tfsdk:"device_id"`
	DeviceOwner                 *string                            `json:"device_owner,omitempty" tfsdk:"device_owner"`
	ErrorMessage                *string                            `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string                            `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	FixedIps                    []OpenStackFixedIp                 `json:"fixed_ips,omitempty" tfsdk:"fixed_ips"`
	FloatingIps                 []string                           `json:"floating_ips,omitempty" tfsdk:"floating_ips"`
	IsLimitBased                *bool                              `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                              `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	MacAddress                  *string                            `json:"mac_address,omitempty" tfsdk:"mac_address"`
	MarketplaceCategoryName     *string                            `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                            `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                            `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                            `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                            `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                            `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                            `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                            `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string                            `json:"name,omitempty" tfsdk:"name"`
	Network                     *string                            `json:"network,omitempty" tfsdk:"network"`
	NetworkName                 *string                            `json:"network_name,omitempty" tfsdk:"network_name"`
	NetworkUuid                 *string                            `json:"network_uuid,omitempty" tfsdk:"network_uuid"`
	PortSecurityEnabled         *bool                              `json:"port_security_enabled,omitempty" tfsdk:"port_security_enabled"`
	Project                     *string                            `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string                            `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string                            `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ResourceType                *string                            `json:"resource_type,omitempty" tfsdk:"resource_type"`
	SecurityGroups              []OpenStackPortNestedSecurityGroup `json:"security_groups,omitempty" tfsdk:"security_groups"`
	ServiceName                 *string                            `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string                            `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                            `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                            `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                            `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string                            `json:"state,omitempty" tfsdk:"state"`
	Status                      *string                            `json:"status,omitempty" tfsdk:"status"`
	Tenant                      *string                            `json:"tenant,omitempty" tfsdk:"tenant"`
	TenantName                  *string                            `json:"tenant_name,omitempty" tfsdk:"tenant_name"`
	TenantUuid                  *string                            `json:"tenant_uuid,omitempty" tfsdk:"tenant_uuid"`
	Url                         *string                            `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackPortIPUpdateRequest struct {
	IpAddress *string `json:"ip_address" tfsdk:"ip_address"`
	Subnet    *string `json:"subnet" tfsdk:"subnet"`
}

type OpenStackPortNestedSecurityGroup struct {
	Name *string `json:"name,omitempty" tfsdk:"name"`
	Url  *string `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackPortNestedSecurityGroupRequest struct {
	Name *string `json:"name" tfsdk:"name"`
}

type OpenStackPortRequest struct {
	AllowedAddressPairs []OpenStackAllowedAddressPairRequest      `json:"allowed_address_pairs,omitempty" tfsdk:"allowed_address_pairs"`
	Description         *string                                   `json:"description,omitempty" tfsdk:"description"`
	FixedIps            []OpenStackFixedIpRequest                 `json:"fixed_ips,omitempty" tfsdk:"fixed_ips"`
	MacAddress          *string                                   `json:"mac_address,omitempty" tfsdk:"mac_address"`
	Name                *string                                   `json:"name" tfsdk:"name"`
	Network             *string                                   `json:"network,omitempty" tfsdk:"network"`
	PortSecurityEnabled *bool                                     `json:"port_security_enabled,omitempty" tfsdk:"port_security_enabled"`
	SecurityGroups      []OpenStackPortNestedSecurityGroupRequest `json:"security_groups,omitempty" tfsdk:"security_groups"`
	TargetTenant        *string                                   `json:"target_tenant,omitempty" tfsdk:"target_tenant"`
}

type OpenStackRouter struct {
	AccessUrl                   *string                `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string                `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created                     *string                `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string                `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string                `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string                `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string                `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string                `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string                `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string                `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string                `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	FixedIps                    []OpenStackFixedIp     `json:"fixed_ips,omitempty" tfsdk:"fixed_ips"`
	IsLimitBased                *bool                  `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                  `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string                `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string                `json:"name,omitempty" tfsdk:"name"`
	OfferingExternalIps         []string               `json:"offering_external_ips,omitempty" tfsdk:"offering_external_ips"`
	Ports                       []OpenStackNestedPort  `json:"ports,omitempty" tfsdk:"ports"`
	Project                     *string                `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string                `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string                `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ResourceType                *string                `json:"resource_type,omitempty" tfsdk:"resource_type"`
	Routes                      []OpenStackStaticRoute `json:"routes,omitempty" tfsdk:"routes"`
	ServiceName                 *string                `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string                `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string                `json:"state,omitempty" tfsdk:"state"`
	Tenant                      *string                `json:"tenant,omitempty" tfsdk:"tenant"`
	TenantName                  *string                `json:"tenant_name,omitempty" tfsdk:"tenant_name"`
	TenantUuid                  *string                `json:"tenant_uuid,omitempty" tfsdk:"tenant_uuid"`
	Url                         *string                `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackRouterInterfaceRequest struct {
	Port   *string `json:"port,omitempty" tfsdk:"port"`
	Subnet *string `json:"subnet,omitempty" tfsdk:"subnet"`
}

type OpenStackRouterSetRoutes struct {
	Routes []OpenStackStaticRoute `json:"routes" tfsdk:"routes"`
}

type OpenStackRouterSetRoutesRequest struct {
	Routes []OpenStackStaticRouteRequest `json:"routes" tfsdk:"routes"`
}

type OpenStackSecurityGroup struct {
	AccessUrl                   *string                            `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string                            `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created                     *string                            `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string                            `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string                            `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string                            `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string                            `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string                            `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string                            `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string                            `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string                            `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	IsLimitBased                *bool                              `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                              `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string                            `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                            `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                            `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                            `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                            `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                            `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                            `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                            `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string                            `json:"name,omitempty" tfsdk:"name"`
	Project                     *string                            `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string                            `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string                            `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ResourceType                *string                            `json:"resource_type,omitempty" tfsdk:"resource_type"`
	Rules                       []OpenStackSecurityGroupRuleCreate `json:"rules,omitempty" tfsdk:"rules"`
	ServiceName                 *string                            `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string                            `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                            `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                            `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                            `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string                            `json:"state,omitempty" tfsdk:"state"`
	Tenant                      *string                            `json:"tenant,omitempty" tfsdk:"tenant"`
	TenantName                  *string                            `json:"tenant_name,omitempty" tfsdk:"tenant_name"`
	TenantUuid                  *string                            `json:"tenant_uuid,omitempty" tfsdk:"tenant_uuid"`
	Url                         *string                            `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackSecurityGroupHyperlinkRequest struct {
	Url *string `json:"url" tfsdk:"url"`
}

type OpenStackSecurityGroupRequest struct {
	Description *string                                   `json:"description,omitempty" tfsdk:"description"`
	Name        *string                                   `json:"name" tfsdk:"name"`
	Rules       []OpenStackSecurityGroupRuleCreateRequest `json:"rules" tfsdk:"rules"`
}

type OpenStackSecurityGroupRuleCreate struct {
	Cidr            *string `json:"cidr,omitempty" tfsdk:"cidr"`
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	Direction       *string `json:"direction,omitempty" tfsdk:"direction"`
	Ethertype       *string `json:"ethertype,omitempty" tfsdk:"ethertype"`
	FromPort        *int64  `json:"from_port,omitempty" tfsdk:"from_port"`
	Id              *int64  `json:"id,omitempty" tfsdk:"id"`
	Protocol        *string `json:"protocol,omitempty" tfsdk:"protocol"`
	RemoteGroup     *string `json:"remote_group,omitempty" tfsdk:"remote_group"`
	RemoteGroupName *string `json:"remote_group_name,omitempty" tfsdk:"remote_group_name"`
	RemoteGroupUuid *string `json:"remote_group_uuid,omitempty" tfsdk:"remote_group_uuid"`
	ToPort          *int64  `json:"to_port,omitempty" tfsdk:"to_port"`
}

type OpenStackSecurityGroupRuleCreateRequest struct {
	Cidr        *string `json:"cidr,omitempty" tfsdk:"cidr"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Direction   *string `json:"direction,omitempty" tfsdk:"direction"`
	Ethertype   *string `json:"ethertype,omitempty" tfsdk:"ethertype"`
	FromPort    *int64  `json:"from_port,omitempty" tfsdk:"from_port"`
	Protocol    *string `json:"protocol,omitempty" tfsdk:"protocol"`
	RemoteGroup *string `json:"remote_group,omitempty" tfsdk:"remote_group"`
	ToPort      *int64  `json:"to_port,omitempty" tfsdk:"to_port"`
}

type OpenStackSecurityGroupRuleUpdateByNameRequest struct {
	Cidr            *string `json:"cidr,omitempty" tfsdk:"cidr"`
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	Direction       *string `json:"direction,omitempty" tfsdk:"direction"`
	Ethertype       *string `json:"ethertype,omitempty" tfsdk:"ethertype"`
	FromPort        *int64  `json:"from_port,omitempty" tfsdk:"from_port"`
	Protocol        *string `json:"protocol,omitempty" tfsdk:"protocol"`
	RemoteGroup     *string `json:"remote_group,omitempty" tfsdk:"remote_group"`
	RemoteGroupName *string `json:"remote_group_name,omitempty" tfsdk:"remote_group_name"`
	ToPort          *int64  `json:"to_port,omitempty" tfsdk:"to_port"`
}

type OpenStackSecurityGroupRuleUpdateRequest struct {
	Cidr        *string `json:"cidr,omitempty" tfsdk:"cidr"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Direction   *string `json:"direction,omitempty" tfsdk:"direction"`
	Ethertype   *string `json:"ethertype,omitempty" tfsdk:"ethertype"`
	FromPort    *int64  `json:"from_port,omitempty" tfsdk:"from_port"`
	Protocol    *string `json:"protocol,omitempty" tfsdk:"protocol"`
	RemoteGroup *string `json:"remote_group,omitempty" tfsdk:"remote_group"`
	ToPort      *int64  `json:"to_port,omitempty" tfsdk:"to_port"`
}

type OpenStackSecurityGroupUpdate struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
}

type OpenStackSecurityGroupUpdateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
}

type OpenStackServerGroup struct {
	AccessUrl                   *string                   `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string                   `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created                     *string                   `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string                   `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string                   `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string                   `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string                   `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string                   `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string                   `json:"description,omitempty" tfsdk:"description"`
	DisplayName                 *string                   `json:"display_name,omitempty" tfsdk:"display_name"`
	ErrorMessage                *string                   `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string                   `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	Instances                   []OpenStackNestedInstance `json:"instances,omitempty" tfsdk:"instances"`
	IsLimitBased                *bool                     `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                     `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string                   `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                   `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                   `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                   `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                   `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                   `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                   `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                   `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string                   `json:"name,omitempty" tfsdk:"name"`
	Policy                      *string                   `json:"policy,omitempty" tfsdk:"policy"`
	Project                     *string                   `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string                   `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string                   `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ResourceType                *string                   `json:"resource_type,omitempty" tfsdk:"resource_type"`
	ServiceName                 *string                   `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string                   `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                   `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                   `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                   `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string                   `json:"state,omitempty" tfsdk:"state"`
	Tenant                      *string                   `json:"tenant,omitempty" tfsdk:"tenant"`
	TenantName                  *string                   `json:"tenant_name,omitempty" tfsdk:"tenant_name"`
	TenantUuid                  *string                   `json:"tenant_uuid,omitempty" tfsdk:"tenant_uuid"`
	Url                         *string                   `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackServerGroupHyperlinkRequest struct {
	Url *string `json:"url" tfsdk:"url"`
}

type OpenStackServerGroupRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
	Policy      *string `json:"policy,omitempty" tfsdk:"policy"`
}

type OpenStackSnapshot struct {
	AccessUrl                   *string                        `json:"access_url,omitempty" tfsdk:"access_url"`
	Action                      *string                        `json:"action,omitempty" tfsdk:"action"`
	BackendId                   *string                        `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created                     *string                        `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string                        `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string                        `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string                        `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string                        `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string                        `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string                        `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string                        `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string                        `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	IsLimitBased                *bool                          `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                          `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	KeptUntil                   *string                        `json:"kept_until,omitempty" tfsdk:"kept_until"`
	MarketplaceCategoryName     *string                        `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                        `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                        `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                        `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                        `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                        `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                        `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                        `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string                        `json:"name,omitempty" tfsdk:"name"`
	Project                     *string                        `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string                        `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string                        `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ResourceType                *string                        `json:"resource_type,omitempty" tfsdk:"resource_type"`
	Restorations                []OpenStackSnapshotRestoration `json:"restorations,omitempty" tfsdk:"restorations"`
	RuntimeState                *string                        `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceName                 *string                        `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string                        `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                        `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                        `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                        `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	Size                        *int64                         `json:"size,omitempty" tfsdk:"size"`
	SourceVolume                *string                        `json:"source_volume,omitempty" tfsdk:"source_volume"`
	SourceVolumeMarketplaceUuid *string                        `json:"source_volume_marketplace_uuid,omitempty" tfsdk:"source_volume_marketplace_uuid"`
	SourceVolumeName            *string                        `json:"source_volume_name,omitempty" tfsdk:"source_volume_name"`
	State                       *string                        `json:"state,omitempty" tfsdk:"state"`
	Url                         *string                        `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackSnapshotRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	KeptUntil   *string `json:"kept_until,omitempty" tfsdk:"kept_until"`
	Name        *string `json:"name" tfsdk:"name"`
}

type OpenStackSnapshotRestoration struct {
	Created            *string `json:"created,omitempty" tfsdk:"created"`
	Description        *string `json:"description,omitempty" tfsdk:"description"`
	Volume             *string `json:"volume,omitempty" tfsdk:"volume"`
	VolumeDevice       *string `json:"volume_device,omitempty" tfsdk:"volume_device"`
	VolumeName         *string `json:"volume_name,omitempty" tfsdk:"volume_name"`
	VolumeRuntimeState *string `json:"volume_runtime_state,omitempty" tfsdk:"volume_runtime_state"`
	VolumeSize         *int64  `json:"volume_size,omitempty" tfsdk:"volume_size"`
	VolumeState        *string `json:"volume_state,omitempty" tfsdk:"volume_state"`
}

type OpenStackSnapshotRestorationRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
}

type OpenStackStaticRoute struct {
	Destination *string `json:"destination,omitempty" tfsdk:"destination"`
	Nexthop     *string `json:"nexthop,omitempty" tfsdk:"nexthop"`
}

type OpenStackStaticRouteRequest struct {
	Destination *string `json:"destination" tfsdk:"destination"`
	Nexthop     *string `json:"nexthop" tfsdk:"nexthop"`
}

type OpenStackSubNet struct {
	AccessUrl                   *string                         `json:"access_url,omitempty" tfsdk:"access_url"`
	AllocationPools             []OpenStackSubNetAllocationPool `json:"allocation_pools,omitempty" tfsdk:"allocation_pools"`
	BackendId                   *string                         `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Cidr                        *string                         `json:"cidr,omitempty" tfsdk:"cidr"`
	Created                     *string                         `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string                         `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string                         `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string                         `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string                         `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string                         `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string                         `json:"description,omitempty" tfsdk:"description"`
	DisableGateway              *bool                           `json:"disable_gateway,omitempty" tfsdk:"disable_gateway"`
	DnsNameservers              []string                        `json:"dns_nameservers,omitempty" tfsdk:"dns_nameservers"`
	EnableDhcp                  *bool                           `json:"enable_dhcp,omitempty" tfsdk:"enable_dhcp"`
	ErrorMessage                *string                         `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string                         `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	GatewayIp                   *string                         `json:"gateway_ip,omitempty" tfsdk:"gateway_ip"`
	HostRoutes                  []OpenStackStaticRoute          `json:"host_routes,omitempty" tfsdk:"host_routes"`
	IpVersion                   *int64                          `json:"ip_version,omitempty" tfsdk:"ip_version"`
	IsConnected                 *bool                           `json:"is_connected,omitempty" tfsdk:"is_connected"`
	IsLimitBased                *bool                           `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                           `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string                         `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                         `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                         `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                         `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                         `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                         `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                         `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                         `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string                         `json:"name,omitempty" tfsdk:"name"`
	Network                     *string                         `json:"network,omitempty" tfsdk:"network"`
	NetworkName                 *string                         `json:"network_name,omitempty" tfsdk:"network_name"`
	Project                     *string                         `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string                         `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string                         `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ResourceType                *string                         `json:"resource_type,omitempty" tfsdk:"resource_type"`
	ServiceName                 *string                         `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string                         `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                         `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                         `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                         `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string                         `json:"state,omitempty" tfsdk:"state"`
	Tenant                      *string                         `json:"tenant,omitempty" tfsdk:"tenant"`
	TenantName                  *string                         `json:"tenant_name,omitempty" tfsdk:"tenant_name"`
	Url                         *string                         `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackSubNetAllocationPool struct {
	End   *string `json:"end,omitempty" tfsdk:"end"`
	Start *string `json:"start,omitempty" tfsdk:"start"`
}

type OpenStackSubNetAllocationPoolRequest struct {
	End   *string `json:"end" tfsdk:"end"`
	Start *string `json:"start" tfsdk:"start"`
}

type OpenStackSubNetRequest struct {
	AllocationPools []OpenStackSubNetAllocationPoolRequest `json:"allocation_pools,omitempty" tfsdk:"allocation_pools"`
	Cidr            *string                                `json:"cidr,omitempty" tfsdk:"cidr"`
	Description     *string                                `json:"description,omitempty" tfsdk:"description"`
	DisableGateway  *bool                                  `json:"disable_gateway,omitempty" tfsdk:"disable_gateway"`
	DnsNameservers  []string                               `json:"dns_nameservers,omitempty" tfsdk:"dns_nameservers"`
	GatewayIp       *string                                `json:"gateway_ip,omitempty" tfsdk:"gateway_ip"`
	HostRoutes      []OpenStackStaticRouteRequest          `json:"host_routes,omitempty" tfsdk:"host_routes"`
	Name            *string                                `json:"name" tfsdk:"name"`
}

type OpenStackTenant struct {
	AccessUrl                   *string `json:"access_url,omitempty" tfsdk:"access_url"`
	AvailabilityZone            *string `json:"availability_zone,omitempty" tfsdk:"availability_zone"`
	BackendId                   *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created                     *string `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	DefaultVolumeTypeName       *string `json:"default_volume_type_name,omitempty" tfsdk:"default_volume_type_name"`
	Description                 *string `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	ExternalNetworkId           *string `json:"external_network_id,omitempty" tfsdk:"external_network_id"`
	InternalNetworkId           *string `json:"internal_network_id,omitempty" tfsdk:"internal_network_id"`
	IsLimitBased                *bool   `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool   `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string `json:"name,omitempty" tfsdk:"name"`
	Project                     *string `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	Quotas                      []Quota `json:"quotas,omitempty" tfsdk:"quotas"`
	ResourceType                *string `json:"resource_type,omitempty" tfsdk:"resource_type"`
	ServiceName                 *string `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	SkipCreationOfDefaultRouter *bool   `json:"skip_creation_of_default_router,omitempty" tfsdk:"skip_creation_of_default_router"`
	State                       *string `json:"state,omitempty" tfsdk:"state"`
	Url                         *string `json:"url,omitempty" tfsdk:"url"`
	UserPassword                *string `json:"user_password,omitempty" tfsdk:"user_password"`
	UserUsername                *string `json:"user_username,omitempty" tfsdk:"user_username"`
}

type OpenStackTenantChangePasswordRequest struct {
	UserPassword *string `json:"user_password" tfsdk:"user_password"`
}

type OpenStackTenantCreateOrderAttributes struct {
	AvailabilityZone            *string                               `json:"availability_zone,omitempty" tfsdk:"availability_zone"`
	Description                 *string                               `json:"description,omitempty" tfsdk:"description"`
	Name                        *string                               `json:"name" tfsdk:"name"`
	SecurityGroups              []OpenStackTenantSecurityGroupRequest `json:"security_groups,omitempty" tfsdk:"security_groups"`
	SkipConnectionExtnet        *bool                                 `json:"skip_connection_extnet,omitempty" tfsdk:"skip_connection_extnet"`
	SkipCreationOfDefaultRouter *bool                                 `json:"skip_creation_of_default_router,omitempty" tfsdk:"skip_creation_of_default_router"`
	SkipCreationOfDefaultSubnet *bool                                 `json:"skip_creation_of_default_subnet,omitempty" tfsdk:"skip_creation_of_default_subnet"`
	SubnetCidr                  *string                               `json:"subnet_cidr,omitempty" tfsdk:"subnet_cidr"`
}

type OpenStackTenantQuota struct {
	Instances              *int64 `json:"instances,omitempty" tfsdk:"instances"`
	Ram                    *int64 `json:"ram,omitempty" tfsdk:"ram"`
	SecurityGroupCount     *int64 `json:"security_group_count,omitempty" tfsdk:"security_group_count"`
	SecurityGroupRuleCount *int64 `json:"security_group_rule_count,omitempty" tfsdk:"security_group_rule_count"`
	Snapshots              *int64 `json:"snapshots,omitempty" tfsdk:"snapshots"`
	Storage                *int64 `json:"storage,omitempty" tfsdk:"storage"`
	Vcpu                   *int64 `json:"vcpu,omitempty" tfsdk:"vcpu"`
	Volumes                *int64 `json:"volumes,omitempty" tfsdk:"volumes"`
}

type OpenStackTenantQuotaRequest struct {
	Instances              *int64 `json:"instances,omitempty" tfsdk:"instances"`
	Ram                    *int64 `json:"ram,omitempty" tfsdk:"ram"`
	SecurityGroupCount     *int64 `json:"security_group_count,omitempty" tfsdk:"security_group_count"`
	SecurityGroupRuleCount *int64 `json:"security_group_rule_count,omitempty" tfsdk:"security_group_rule_count"`
	Snapshots              *int64 `json:"snapshots,omitempty" tfsdk:"snapshots"`
	Storage                *int64 `json:"storage,omitempty" tfsdk:"storage"`
	Vcpu                   *int64 `json:"vcpu,omitempty" tfsdk:"vcpu"`
	Volumes                *int64 `json:"volumes,omitempty" tfsdk:"volumes"`
}

type OpenStackTenantRequest struct {
	AvailabilityZone            *string                               `json:"availability_zone,omitempty" tfsdk:"availability_zone"`
	DefaultVolumeTypeName       *string                               `json:"default_volume_type_name,omitempty" tfsdk:"default_volume_type_name"`
	Description                 *string                               `json:"description,omitempty" tfsdk:"description"`
	Name                        *string                               `json:"name" tfsdk:"name"`
	SecurityGroups              []OpenStackTenantSecurityGroupRequest `json:"security_groups,omitempty" tfsdk:"security_groups"`
	SkipCreationOfDefaultRouter *bool                                 `json:"skip_creation_of_default_router,omitempty" tfsdk:"skip_creation_of_default_router"`
	SkipCreationOfDefaultSubnet *bool                                 `json:"skip_creation_of_default_subnet,omitempty" tfsdk:"skip_creation_of_default_subnet"`
}

type OpenStackTenantSecurityGroup struct {
	Description *string                            `json:"description,omitempty" tfsdk:"description"`
	Name        *string                            `json:"name" tfsdk:"name"`
	Rules       []OpenStackSecurityGroupRuleCreate `json:"rules,omitempty" tfsdk:"rules"`
}

type OpenStackTenantSecurityGroupRequest struct {
	Description *string                                   `json:"description,omitempty" tfsdk:"description"`
	Name        *string                                   `json:"name" tfsdk:"name"`
	Rules       []OpenStackSecurityGroupRuleCreateRequest `json:"rules,omitempty" tfsdk:"rules"`
}

type OpenStackVolume struct {
	AccessUrl                   *string `json:"access_url,omitempty" tfsdk:"access_url"`
	Action                      *string `json:"action,omitempty" tfsdk:"action"`
	AvailabilityZone            *string `json:"availability_zone,omitempty" tfsdk:"availability_zone"`
	AvailabilityZoneName        *string `json:"availability_zone_name,omitempty" tfsdk:"availability_zone_name"`
	BackendId                   *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Bootable                    *bool   `json:"bootable,omitempty" tfsdk:"bootable"`
	Created                     *string `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string `json:"description,omitempty" tfsdk:"description"`
	Device                      *string `json:"device,omitempty" tfsdk:"device"`
	ErrorMessage                *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	ExtendEnabled               *bool   `json:"extend_enabled,omitempty" tfsdk:"extend_enabled"`
	Image                       *string `json:"image,omitempty" tfsdk:"image"`
	ImageMetadata               *string `json:"image_metadata,omitempty" tfsdk:"image_metadata"`
	ImageName                   *string `json:"image_name,omitempty" tfsdk:"image_name"`
	Instance                    *string `json:"instance,omitempty" tfsdk:"instance"`
	InstanceMarketplaceUuid     *string `json:"instance_marketplace_uuid,omitempty" tfsdk:"instance_marketplace_uuid"`
	InstanceName                *string `json:"instance_name,omitempty" tfsdk:"instance_name"`
	IsLimitBased                *bool   `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool   `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string `json:"name,omitempty" tfsdk:"name"`
	Project                     *string `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ResourceType                *string `json:"resource_type,omitempty" tfsdk:"resource_type"`
	RuntimeState                *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceName                 *string `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	Size                        *int64  `json:"size,omitempty" tfsdk:"size"`
	SourceSnapshot              *string `json:"source_snapshot,omitempty" tfsdk:"source_snapshot"`
	State                       *string `json:"state,omitempty" tfsdk:"state"`
	Tenant                      *string `json:"tenant,omitempty" tfsdk:"tenant"`
	TenantUuid                  *string `json:"tenant_uuid,omitempty" tfsdk:"tenant_uuid"`
	Type                        *string `json:"type,omitempty" tfsdk:"type"`
	TypeName                    *string `json:"type_name,omitempty" tfsdk:"type_name"`
	Url                         *string `json:"url,omitempty" tfsdk:"url"`
}

type OpenStackVolumeAvailabilityZone struct {
	Available *bool   `json:"available,omitempty" tfsdk:"available"`
	Name      *string `json:"name" tfsdk:"name"`
	Settings  *string `json:"settings,omitempty" tfsdk:"settings"`
	Url       *string `json:"url" tfsdk:"url"`
}

type OpenStackVolumeCreateOrderAttributes struct {
	AvailabilityZone *string `json:"availability_zone,omitempty" tfsdk:"availability_zone"`
	Description      *string `json:"description,omitempty" tfsdk:"description"`
	Image            *string `json:"image,omitempty" tfsdk:"image"`
	Name             *string `json:"name" tfsdk:"name"`
	Size             *int64  `json:"size,omitempty" tfsdk:"size"`
	Type             *string `json:"type,omitempty" tfsdk:"type"`
}

type OpenStackVolumeExtendRequest struct {
	DiskSize *int64 `json:"disk_size" tfsdk:"disk_size"`
}

type OpenStackVolumeRequest struct {
	Bootable    *bool   `json:"bootable,omitempty" tfsdk:"bootable"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
}

type OpenStackVolumeRetypeRequest struct {
	Type *string `json:"type" tfsdk:"type"`
}

type OpenStackVolumeType struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
	Settings    *string `json:"settings" tfsdk:"settings"`
	Url         *string `json:"url" tfsdk:"url"`
}

type OptionField struct {
	CascadeConfig             *CascadeConfig             `json:"cascade_config,omitempty" tfsdk:"cascade_config"`
	Choices                   []string                   `json:"choices,omitempty" tfsdk:"choices"`
	ComponentMultiplierConfig *ComponentMultiplierConfig `json:"component_multiplier_config,omitempty" tfsdk:"component_multiplier_config"`
	Default                   *string                    `json:"default,omitempty" tfsdk:"default"`
	DefaultConfigs            *K8sDefaultConfiguration   `json:"default_configs,omitempty" tfsdk:"default_configs"`
	HelpText                  *string                    `json:"help_text,omitempty" tfsdk:"help_text"`
	Label                     *string                    `json:"label" tfsdk:"label"`
	Max                       *int64                     `json:"max,omitempty" tfsdk:"max"`
	Min                       *int64                     `json:"min,omitempty" tfsdk:"min"`
	Required                  *bool                      `json:"required,omitempty" tfsdk:"required"`
	Type                      *string                    `json:"type" tfsdk:"type"`
}

type OptionFieldRequest struct {
	CascadeConfig             *CascadeConfigRequest             `json:"cascade_config,omitempty" tfsdk:"cascade_config"`
	Choices                   []string                          `json:"choices,omitempty" tfsdk:"choices"`
	ComponentMultiplierConfig *ComponentMultiplierConfigRequest `json:"component_multiplier_config,omitempty" tfsdk:"component_multiplier_config"`
	Default                   *string                           `json:"default,omitempty" tfsdk:"default"`
	DefaultConfigs            *K8sDefaultConfigurationRequest   `json:"default_configs,omitempty" tfsdk:"default_configs"`
	HelpText                  *string                           `json:"help_text,omitempty" tfsdk:"help_text"`
	Label                     *string                           `json:"label" tfsdk:"label"`
	Max                       *int64                            `json:"max,omitempty" tfsdk:"max"`
	Min                       *int64                            `json:"min,omitempty" tfsdk:"min"`
	Required                  *bool                             `json:"required,omitempty" tfsdk:"required"`
	Type                      *string                           `json:"type" tfsdk:"type"`
}

type OptionFieldTypeEnum struct {
}

type OrderAttachment struct {
	Attachment *string `json:"attachment,omitempty" tfsdk:"attachment"`
}

type OrderAttachmentRequest struct {
	Attachment *string `json:"attachment,omitempty" tfsdk:"attachment"`
}

type OrderAttachmentRequestForm struct {
	Attachment *string `json:"attachment,omitempty" tfsdk:"attachment"`
}

type OrderAttachmentRequestMultipart struct {
	Attachment *string `json:"attachment,omitempty" tfsdk:"attachment"`
}

type OrderBackendIDRequest struct {
	BackendId *string `json:"backend_id" tfsdk:"backend_id"`
}

type OrderCreateRequest struct {
	AcceptingTermsOfService *bool   `json:"accepting_terms_of_service,omitempty" tfsdk:"accepting_terms_of_service"`
	CallbackUrl             *string `json:"callback_url,omitempty" tfsdk:"callback_url"`
	Offering                *string `json:"offering" tfsdk:"offering"`
	Plan                    *string `json:"plan,omitempty" tfsdk:"plan"`
	Project                 *string `json:"project" tfsdk:"project"`
	RequestComment          *string `json:"request_comment,omitempty" tfsdk:"request_comment"`
	Slug                    *string `json:"slug,omitempty" tfsdk:"slug"`
	StartDate               *string `json:"start_date,omitempty" tfsdk:"start_date"`
	Type                    *string `json:"type,omitempty" tfsdk:"type"`
}

type OrderDetails struct {
	ActivationPrice            *float64 `json:"activation_price,omitempty" tfsdk:"activation_price"`
	Attachment                 *string  `json:"attachment,omitempty" tfsdk:"attachment"`
	BackendId                  *string  `json:"backend_id,omitempty" tfsdk:"backend_id"`
	CallbackUrl                *string  `json:"callback_url,omitempty" tfsdk:"callback_url"`
	CanTerminate               *bool    `json:"can_terminate,omitempty" tfsdk:"can_terminate"`
	CategoryIcon               *string  `json:"category_icon,omitempty" tfsdk:"category_icon"`
	CategoryTitle              *string  `json:"category_title,omitempty" tfsdk:"category_title"`
	CategoryUuid               *string  `json:"category_uuid,omitempty" tfsdk:"category_uuid"`
	CompletedAt                *string  `json:"completed_at,omitempty" tfsdk:"completed_at"`
	ConsumerReviewedAt         *string  `json:"consumer_reviewed_at,omitempty" tfsdk:"consumer_reviewed_at"`
	ConsumerReviewedBy         *string  `json:"consumer_reviewed_by,omitempty" tfsdk:"consumer_reviewed_by"`
	ConsumerReviewedByFullName *string  `json:"consumer_reviewed_by_full_name,omitempty" tfsdk:"consumer_reviewed_by_full_name"`
	ConsumerReviewedByUsername *string  `json:"consumer_reviewed_by_username,omitempty" tfsdk:"consumer_reviewed_by_username"`
	Cost                       *string  `json:"cost,omitempty" tfsdk:"cost"`
	Created                    *string  `json:"created,omitempty" tfsdk:"created"`
	CreatedByCivilNumber       *string  `json:"created_by_civil_number,omitempty" tfsdk:"created_by_civil_number"`
	CreatedByFullName          *string  `json:"created_by_full_name,omitempty" tfsdk:"created_by_full_name"`
	CreatedByUsername          *string  `json:"created_by_username,omitempty" tfsdk:"created_by_username"`
	CustomerName               *string  `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerSlug               *string  `json:"customer_slug,omitempty" tfsdk:"customer_slug"`
	CustomerUuid               *string  `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	ErrorMessage               *string  `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback             *string  `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	FixedPrice                 *float64 `json:"fixed_price,omitempty" tfsdk:"fixed_price"`
	MarketplaceResourceUuid    *string  `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                   *string  `json:"modified,omitempty" tfsdk:"modified"`
	NewCostEstimate            *string  `json:"new_cost_estimate,omitempty" tfsdk:"new_cost_estimate"`
	NewPlanName                *string  `json:"new_plan_name,omitempty" tfsdk:"new_plan_name"`
	NewPlanUuid                *string  `json:"new_plan_uuid,omitempty" tfsdk:"new_plan_uuid"`
	Offering                   *string  `json:"offering,omitempty" tfsdk:"offering"`
	OfferingBillable           *bool    `json:"offering_billable,omitempty" tfsdk:"offering_billable"`
	OfferingDescription        *string  `json:"offering_description,omitempty" tfsdk:"offering_description"`
	OfferingImage              *string  `json:"offering_image,omitempty" tfsdk:"offering_image"`
	OfferingName               *string  `json:"offering_name,omitempty" tfsdk:"offering_name"`
	OfferingShared             *bool    `json:"offering_shared,omitempty" tfsdk:"offering_shared"`
	OfferingThumbnail          *string  `json:"offering_thumbnail,omitempty" tfsdk:"offering_thumbnail"`
	OfferingType               *string  `json:"offering_type,omitempty" tfsdk:"offering_type"`
	OfferingUuid               *string  `json:"offering_uuid,omitempty" tfsdk:"offering_uuid"`
	OldCostEstimate            *float64 `json:"old_cost_estimate,omitempty" tfsdk:"old_cost_estimate"`
	OldPlanName                *string  `json:"old_plan_name,omitempty" tfsdk:"old_plan_name"`
	OldPlanUuid                *string  `json:"old_plan_uuid,omitempty" tfsdk:"old_plan_uuid"`
	OrderSubtype               *string  `json:"order_subtype,omitempty" tfsdk:"order_subtype"`
	Output                     *string  `json:"output,omitempty" tfsdk:"output"`
	Plan                       *string  `json:"plan,omitempty" tfsdk:"plan"`
	PlanDescription            *string  `json:"plan_description,omitempty" tfsdk:"plan_description"`
	PlanName                   *string  `json:"plan_name,omitempty" tfsdk:"plan_name"`
	PlanUnit                   *string  `json:"plan_unit,omitempty" tfsdk:"plan_unit"`
	PlanUuid                   *string  `json:"plan_uuid,omitempty" tfsdk:"plan_uuid"`
	ProjectDescription         *string  `json:"project_description,omitempty" tfsdk:"project_description"`
	ProjectName                *string  `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectSlug                *string  `json:"project_slug,omitempty" tfsdk:"project_slug"`
	ProjectUuid                *string  `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ProviderName               *string  `json:"provider_name,omitempty" tfsdk:"provider_name"`
	ProviderReviewedAt         *string  `json:"provider_reviewed_at,omitempty" tfsdk:"provider_reviewed_at"`
	ProviderReviewedBy         *string  `json:"provider_reviewed_by,omitempty" tfsdk:"provider_reviewed_by"`
	ProviderReviewedByFullName *string  `json:"provider_reviewed_by_full_name,omitempty" tfsdk:"provider_reviewed_by_full_name"`
	ProviderReviewedByUsername *string  `json:"provider_reviewed_by_username,omitempty" tfsdk:"provider_reviewed_by_username"`
	ProviderSlug               *string  `json:"provider_slug,omitempty" tfsdk:"provider_slug"`
	ProviderUuid               *string  `json:"provider_uuid,omitempty" tfsdk:"provider_uuid"`
	RequestComment             *string  `json:"request_comment,omitempty" tfsdk:"request_comment"`
	ResourceName               *string  `json:"resource_name,omitempty" tfsdk:"resource_name"`
	ResourceType               *string  `json:"resource_type,omitempty" tfsdk:"resource_type"`
	ResourceUuid               *string  `json:"resource_uuid,omitempty" tfsdk:"resource_uuid"`
	Slug                       *string  `json:"slug,omitempty" tfsdk:"slug"`
	StartDate                  *string  `json:"start_date,omitempty" tfsdk:"start_date"`
	State                      *string  `json:"state,omitempty" tfsdk:"state"`
	TerminationComment         *string  `json:"termination_comment,omitempty" tfsdk:"termination_comment"`
	Type                       *string  `json:"type,omitempty" tfsdk:"type"`
	Url                        *string  `json:"url,omitempty" tfsdk:"url"`

	Uuid *string `json:"uuid" tfsdk:"uuid"`
}

type OrderErrorDetailsRequest struct {
	ErrorMessage   *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
}

type OrderState struct {
}

type OrderUUID struct {
	OrderUuid *string `json:"order_uuid" tfsdk:"order_uuid"`
}

type OrganizationGroup struct {
	CustomersCount *int64  `json:"customers_count,omitempty" tfsdk:"customers_count"`
	Name           *string `json:"name,omitempty" tfsdk:"name"`
	Parent         *string `json:"parent,omitempty" tfsdk:"parent"`
	ParentName     *string `json:"parent_name,omitempty" tfsdk:"parent_name"`
	ParentUuid     *string `json:"parent_uuid,omitempty" tfsdk:"parent_uuid"`
	Url            *string `json:"url,omitempty" tfsdk:"url"`
}

type OrganizationGroupRequest struct {
	Name   *string `json:"name" tfsdk:"name"`
	Parent *string `json:"parent,omitempty" tfsdk:"parent"`
}

type OrganizationGroupsRequest struct {
	OrganizationGroups []string `json:"organization_groups,omitempty" tfsdk:"organization_groups"`
}

type PaidRequest struct {
	Date  *string `json:"date" tfsdk:"date"`
	Proof *string `json:"proof,omitempty" tfsdk:"proof"`
}

type PaidRequestForm struct {
	Date  *string `json:"date" tfsdk:"date"`
	Proof *string `json:"proof,omitempty" tfsdk:"proof"`
}

type PaidRequestMultipart struct {
	Date  *string `json:"date" tfsdk:"date"`
	Proof *string `json:"proof,omitempty" tfsdk:"proof"`
}

type PasswordChangeRequest struct {
	NewPassword *string `json:"new_password" tfsdk:"new_password"`
}

type PatchedAccessSubnetRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Inet        *string `json:"inet,omitempty" tfsdk:"inet"`
}

type PatchedAdminAnnouncementRequest struct {
	ActiveFrom  *string `json:"active_from,omitempty" tfsdk:"active_from"`
	ActiveTo    *string `json:"active_to,omitempty" tfsdk:"active_to"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Type        *string `json:"type,omitempty" tfsdk:"type"`
}

type PatchedAllocationRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Groupname   *string `json:"groupname,omitempty" tfsdk:"groupname"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
	NodeLimit   *int64  `json:"node_limit,omitempty" tfsdk:"node_limit"`
}

type PatchedAwsInstanceRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type PatchedAzurePublicIPRequest struct {
	Description   *string `json:"description,omitempty" tfsdk:"description"`
	Location      *string `json:"location,omitempty" tfsdk:"location"`
	ResourceGroup *string `json:"resource_group,omitempty" tfsdk:"resource_group"`
}

type PatchedAzureSqlDatabaseRequest struct {
	Charset     *string `json:"charset,omitempty" tfsdk:"charset"`
	Collation   *string `json:"collation,omitempty" tfsdk:"collation"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Server      *string `json:"server,omitempty" tfsdk:"server"`
}

type PatchedAzureSqlServerRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Location    *string `json:"location,omitempty" tfsdk:"location"`
	StorageMb   *int64  `json:"storage_mb,omitempty" tfsdk:"storage_mb"`
}

type PatchedAzureVirtualMachineRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Location    *string `json:"location,omitempty" tfsdk:"location"`
}

type PatchedBroadcastMessageRequest struct {
	Body    *string `json:"body,omitempty" tfsdk:"body"`
	SendAt  *string `json:"send_at,omitempty" tfsdk:"send_at"`
	Subject *string `json:"subject,omitempty" tfsdk:"subject"`
}

type PatchedCallManagingOrganisationRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Image       *string `json:"image,omitempty" tfsdk:"image"`
}

type PatchedCallManagingOrganisationRequestForm struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Image       *string `json:"image,omitempty" tfsdk:"image"`
}

type PatchedCallManagingOrganisationRequestMultipart struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Image       *string `json:"image,omitempty" tfsdk:"image"`
}

type PatchedCallResourceTemplateRequest struct {
	Description       *string `json:"description,omitempty" tfsdk:"description"`
	IsRequired        *bool   `json:"is_required,omitempty" tfsdk:"is_required"`
	Name              *string `json:"name,omitempty" tfsdk:"name"`
	RequestedOffering *string `json:"requested_offering,omitempty" tfsdk:"requested_offering"`
}

type PatchedCategoryColumnRequest struct {
	Attribute *string `json:"attribute,omitempty" tfsdk:"attribute"`
	Category  *string `json:"category,omitempty" tfsdk:"category"`
	Index     *int64  `json:"index,omitempty" tfsdk:"index"`
	Title     *string `json:"title,omitempty" tfsdk:"title"`
	Widget    *string `json:"widget,omitempty" tfsdk:"widget"`
}

type PatchedCategoryComponentsRequest struct {
	Category     *CategorySerializerForForNestedFieldsRequest `json:"category,omitempty" tfsdk:"category"`
	Description  *string                                      `json:"description,omitempty" tfsdk:"description"`
	MeasuredUnit *string                                      `json:"measured_unit,omitempty" tfsdk:"measured_unit"`
	Name         *string                                      `json:"name,omitempty" tfsdk:"name"`
	Type         *string                                      `json:"type,omitempty" tfsdk:"type"`
}

type PatchedCategoryGroupRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Icon        *string `json:"icon,omitempty" tfsdk:"icon"`
	Title       *string `json:"title,omitempty" tfsdk:"title"`
}

type PatchedCategoryGroupRequestForm struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Icon        *string `json:"icon,omitempty" tfsdk:"icon"`
	Title       *string `json:"title,omitempty" tfsdk:"title"`
}

type PatchedCategoryGroupRequestMultipart struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Icon        *string `json:"icon,omitempty" tfsdk:"icon"`
	Title       *string `json:"title,omitempty" tfsdk:"title"`
}

type PatchedCategoryHelpArticlesRequest struct {
	Categories []CategorySerializerForForNestedFieldsRequest `json:"categories,omitempty" tfsdk:"categories"`
	Title      *string                                       `json:"title,omitempty" tfsdk:"title"`
	Url        *string                                       `json:"url,omitempty" tfsdk:"url"`
}

type PatchedChecklistCategoryRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Icon        *string `json:"icon,omitempty" tfsdk:"icon"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type PatchedChecklistCategoryRequestForm struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Icon        *string `json:"icon,omitempty" tfsdk:"icon"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type PatchedChecklistCategoryRequestMultipart struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Icon        *string `json:"icon,omitempty" tfsdk:"icon"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type PatchedChecklistRequest struct {
	Category      *string `json:"category,omitempty" tfsdk:"category"`
	ChecklistType *string `json:"checklist_type,omitempty" tfsdk:"checklist_type"`
	Description   *string `json:"description,omitempty" tfsdk:"description"`
	Name          *string `json:"name,omitempty" tfsdk:"name"`
}

type PatchedClusterSecurityGroupRequest struct {
	Rules []RancherClusterSecurityGroupRuleRequest `json:"rules,omitempty" tfsdk:"rules"`
}

type PatchedCommentRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	IsPublic    *bool   `json:"is_public,omitempty" tfsdk:"is_public"`
}

type PatchedComponentUserUsageLimitRequest struct {
	Component *string `json:"component,omitempty" tfsdk:"component"`
	Limit     *string `json:"limit,omitempty" tfsdk:"limit"`
	Resource  *string `json:"resource,omitempty" tfsdk:"resource"`
	User      *string `json:"user,omitempty" tfsdk:"user"`
}

type PatchedCreateCustomerCreditRequest struct {
	ApplyAsMinimalConsumption *bool    `json:"apply_as_minimal_consumption,omitempty" tfsdk:"apply_as_minimal_consumption"`
	Customer                  *string  `json:"customer,omitempty" tfsdk:"customer"`
	EndDate                   *string  `json:"end_date,omitempty" tfsdk:"end_date"`
	ExpectedConsumption       *string  `json:"expected_consumption,omitempty" tfsdk:"expected_consumption"`
	GraceCoefficient          *string  `json:"grace_coefficient,omitempty" tfsdk:"grace_coefficient"`
	MinimalConsumptionLogic   *string  `json:"minimal_consumption_logic,omitempty" tfsdk:"minimal_consumption_logic"`
	Offerings                 []string `json:"offerings,omitempty" tfsdk:"offerings"`
	Value                     *string  `json:"value,omitempty" tfsdk:"value"`
}

type PatchedCustomerComponentUsagePolicyRequest struct {
	Actions            *string                                     `json:"actions,omitempty" tfsdk:"actions"`
	ComponentLimitsSet []NestedCustomerUsagePolicyComponentRequest `json:"component_limits_set,omitempty" tfsdk:"component_limits_set"`
	Scope              *string                                     `json:"scope,omitempty" tfsdk:"scope"`
}

type PatchedCustomerEstimatedCostPolicyRequest struct {
	Actions   *string `json:"actions,omitempty" tfsdk:"actions"`
	LimitCost *int64  `json:"limit_cost,omitempty" tfsdk:"limit_cost"`
	Period    *int64  `json:"period,omitempty" tfsdk:"period"`
	Scope     *string `json:"scope,omitempty" tfsdk:"scope"`
}

type PatchedCustomerRequest struct {
	Abbreviation                 *string  `json:"abbreviation,omitempty" tfsdk:"abbreviation"`
	AccessSubnets                *string  `json:"access_subnets,omitempty" tfsdk:"access_subnets"`
	AccountingStartDate          *string  `json:"accounting_start_date,omitempty" tfsdk:"accounting_start_date"`
	Address                      *string  `json:"address,omitempty" tfsdk:"address"`
	AgreementNumber              *string  `json:"agreement_number,omitempty" tfsdk:"agreement_number"`
	Archived                     *bool    `json:"archived,omitempty" tfsdk:"archived"`
	BackendId                    *string  `json:"backend_id,omitempty" tfsdk:"backend_id"`
	BankAccount                  *string  `json:"bank_account,omitempty" tfsdk:"bank_account"`
	BankName                     *string  `json:"bank_name,omitempty" tfsdk:"bank_name"`
	Blocked                      *bool    `json:"blocked,omitempty" tfsdk:"blocked"`
	ContactDetails               *string  `json:"contact_details,omitempty" tfsdk:"contact_details"`
	Country                      *string  `json:"country,omitempty" tfsdk:"country"`
	DefaultTaxPercent            *string  `json:"default_tax_percent,omitempty" tfsdk:"default_tax_percent"`
	Description                  *string  `json:"description,omitempty" tfsdk:"description"`
	DisplayBillingInfoInProjects *bool    `json:"display_billing_info_in_projects,omitempty" tfsdk:"display_billing_info_in_projects"`
	Domain                       *string  `json:"domain,omitempty" tfsdk:"domain"`
	Email                        *string  `json:"email,omitempty" tfsdk:"email"`
	GracePeriodDays              *int64   `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`
	Homepage                     *string  `json:"homepage,omitempty" tfsdk:"homepage"`
	Image                        *string  `json:"image,omitempty" tfsdk:"image"`
	Latitude                     *float64 `json:"latitude,omitempty" tfsdk:"latitude"`
	Longitude                    *float64 `json:"longitude,omitempty" tfsdk:"longitude"`
	MaxServiceAccounts           *int64   `json:"max_service_accounts,omitempty" tfsdk:"max_service_accounts"`
	Name                         *string  `json:"name,omitempty" tfsdk:"name"`
	NativeName                   *string  `json:"native_name,omitempty" tfsdk:"native_name"`
	NotificationEmails           *string  `json:"notification_emails,omitempty" tfsdk:"notification_emails"`
	PhoneNumber                  *string  `json:"phone_number,omitempty" tfsdk:"phone_number"`
	Postal                       *string  `json:"postal,omitempty" tfsdk:"postal"`
	ProjectMetadataChecklist     *string  `json:"project_metadata_checklist,omitempty" tfsdk:"project_metadata_checklist"`
	RegistrationCode             *string  `json:"registration_code,omitempty" tfsdk:"registration_code"`
	Slug                         *string  `json:"slug,omitempty" tfsdk:"slug"`
	SponsorNumber                *int64   `json:"sponsor_number,omitempty" tfsdk:"sponsor_number"`
	VatCode                      *string  `json:"vat_code,omitempty" tfsdk:"vat_code"`
}

type PatchedCustomerRequestForm struct {
	Abbreviation                 *string  `json:"abbreviation,omitempty" tfsdk:"abbreviation"`
	AccessSubnets                *string  `json:"access_subnets,omitempty" tfsdk:"access_subnets"`
	AccountingStartDate          *string  `json:"accounting_start_date,omitempty" tfsdk:"accounting_start_date"`
	Address                      *string  `json:"address,omitempty" tfsdk:"address"`
	AgreementNumber              *string  `json:"agreement_number,omitempty" tfsdk:"agreement_number"`
	Archived                     *bool    `json:"archived,omitempty" tfsdk:"archived"`
	BackendId                    *string  `json:"backend_id,omitempty" tfsdk:"backend_id"`
	BankAccount                  *string  `json:"bank_account,omitempty" tfsdk:"bank_account"`
	BankName                     *string  `json:"bank_name,omitempty" tfsdk:"bank_name"`
	Blocked                      *bool    `json:"blocked,omitempty" tfsdk:"blocked"`
	ContactDetails               *string  `json:"contact_details,omitempty" tfsdk:"contact_details"`
	Country                      *string  `json:"country,omitempty" tfsdk:"country"`
	DefaultTaxPercent            *string  `json:"default_tax_percent,omitempty" tfsdk:"default_tax_percent"`
	Description                  *string  `json:"description,omitempty" tfsdk:"description"`
	DisplayBillingInfoInProjects *bool    `json:"display_billing_info_in_projects,omitempty" tfsdk:"display_billing_info_in_projects"`
	Domain                       *string  `json:"domain,omitempty" tfsdk:"domain"`
	Email                        *string  `json:"email,omitempty" tfsdk:"email"`
	GracePeriodDays              *int64   `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`
	Homepage                     *string  `json:"homepage,omitempty" tfsdk:"homepage"`
	Image                        *string  `json:"image,omitempty" tfsdk:"image"`
	Latitude                     *float64 `json:"latitude,omitempty" tfsdk:"latitude"`
	Longitude                    *float64 `json:"longitude,omitempty" tfsdk:"longitude"`
	MaxServiceAccounts           *int64   `json:"max_service_accounts,omitempty" tfsdk:"max_service_accounts"`
	Name                         *string  `json:"name,omitempty" tfsdk:"name"`
	NativeName                   *string  `json:"native_name,omitempty" tfsdk:"native_name"`
	NotificationEmails           *string  `json:"notification_emails,omitempty" tfsdk:"notification_emails"`
	PhoneNumber                  *string  `json:"phone_number,omitempty" tfsdk:"phone_number"`
	Postal                       *string  `json:"postal,omitempty" tfsdk:"postal"`
	ProjectMetadataChecklist     *string  `json:"project_metadata_checklist,omitempty" tfsdk:"project_metadata_checklist"`
	RegistrationCode             *string  `json:"registration_code,omitempty" tfsdk:"registration_code"`
	Slug                         *string  `json:"slug,omitempty" tfsdk:"slug"`
	SponsorNumber                *int64   `json:"sponsor_number,omitempty" tfsdk:"sponsor_number"`
	VatCode                      *string  `json:"vat_code,omitempty" tfsdk:"vat_code"`
}

type PatchedCustomerRequestMultipart struct {
	Abbreviation                 *string  `json:"abbreviation,omitempty" tfsdk:"abbreviation"`
	AccessSubnets                *string  `json:"access_subnets,omitempty" tfsdk:"access_subnets"`
	AccountingStartDate          *string  `json:"accounting_start_date,omitempty" tfsdk:"accounting_start_date"`
	Address                      *string  `json:"address,omitempty" tfsdk:"address"`
	AgreementNumber              *string  `json:"agreement_number,omitempty" tfsdk:"agreement_number"`
	Archived                     *bool    `json:"archived,omitempty" tfsdk:"archived"`
	BackendId                    *string  `json:"backend_id,omitempty" tfsdk:"backend_id"`
	BankAccount                  *string  `json:"bank_account,omitempty" tfsdk:"bank_account"`
	BankName                     *string  `json:"bank_name,omitempty" tfsdk:"bank_name"`
	Blocked                      *bool    `json:"blocked,omitempty" tfsdk:"blocked"`
	ContactDetails               *string  `json:"contact_details,omitempty" tfsdk:"contact_details"`
	Country                      *string  `json:"country,omitempty" tfsdk:"country"`
	DefaultTaxPercent            *string  `json:"default_tax_percent,omitempty" tfsdk:"default_tax_percent"`
	Description                  *string  `json:"description,omitempty" tfsdk:"description"`
	DisplayBillingInfoInProjects *bool    `json:"display_billing_info_in_projects,omitempty" tfsdk:"display_billing_info_in_projects"`
	Domain                       *string  `json:"domain,omitempty" tfsdk:"domain"`
	Email                        *string  `json:"email,omitempty" tfsdk:"email"`
	GracePeriodDays              *int64   `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`
	Homepage                     *string  `json:"homepage,omitempty" tfsdk:"homepage"`
	Image                        *string  `json:"image,omitempty" tfsdk:"image"`
	Latitude                     *float64 `json:"latitude,omitempty" tfsdk:"latitude"`
	Longitude                    *float64 `json:"longitude,omitempty" tfsdk:"longitude"`
	MaxServiceAccounts           *int64   `json:"max_service_accounts,omitempty" tfsdk:"max_service_accounts"`
	Name                         *string  `json:"name,omitempty" tfsdk:"name"`
	NativeName                   *string  `json:"native_name,omitempty" tfsdk:"native_name"`
	NotificationEmails           *string  `json:"notification_emails,omitempty" tfsdk:"notification_emails"`
	PhoneNumber                  *string  `json:"phone_number,omitempty" tfsdk:"phone_number"`
	Postal                       *string  `json:"postal,omitempty" tfsdk:"postal"`
	ProjectMetadataChecklist     *string  `json:"project_metadata_checklist,omitempty" tfsdk:"project_metadata_checklist"`
	RegistrationCode             *string  `json:"registration_code,omitempty" tfsdk:"registration_code"`
	Slug                         *string  `json:"slug,omitempty" tfsdk:"slug"`
	SponsorNumber                *int64   `json:"sponsor_number,omitempty" tfsdk:"sponsor_number"`
	VatCode                      *string  `json:"vat_code,omitempty" tfsdk:"vat_code"`
}

type PatchedCustomerServiceAccountRequest struct {
	Customer            *string `json:"customer,omitempty" tfsdk:"customer"`
	Description         *string `json:"description,omitempty" tfsdk:"description"`
	Email               *string `json:"email,omitempty" tfsdk:"email"`
	ErrorTraceback      *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	PreferredIdentifier *string `json:"preferred_identifier,omitempty" tfsdk:"preferred_identifier"`
	Username            *string `json:"username,omitempty" tfsdk:"username"`
}

type PatchedDigitalOceanDropletRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type PatchedEmailHookRequest struct {
	Email       *string  `json:"email,omitempty" tfsdk:"email"`
	EventGroups []string `json:"event_groups,omitempty" tfsdk:"event_groups"`
	EventTypes  []string `json:"event_types,omitempty" tfsdk:"event_types"`
	IsActive    *bool    `json:"is_active,omitempty" tfsdk:"is_active"`
}

type PatchedExternalLinkRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Image       *string `json:"image,omitempty" tfsdk:"image"`
	Link        *string `json:"link,omitempty" tfsdk:"link"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type PatchedExternalLinkRequestForm struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Image       *string `json:"image,omitempty" tfsdk:"image"`
	Link        *string `json:"link,omitempty" tfsdk:"link"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type PatchedExternalLinkRequestMultipart struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Image       *string `json:"image,omitempty" tfsdk:"image"`
	Link        *string `json:"link,omitempty" tfsdk:"link"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type PatchedFirecrestJobRequest struct {
	Description  *string `json:"description,omitempty" tfsdk:"description"`
	Name         *string `json:"name,omitempty" tfsdk:"name"`
	RuntimeState *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
}

type PatchedIdentityProviderRequest struct {
	ClientId                 *string `json:"client_id,omitempty" tfsdk:"client_id"`
	ClientSecret             *string `json:"client_secret,omitempty" tfsdk:"client_secret"`
	DiscoveryUrl             *string `json:"discovery_url,omitempty" tfsdk:"discovery_url"`
	EnablePkce               *bool   `json:"enable_pkce,omitempty" tfsdk:"enable_pkce"`
	EnablePostLogoutRedirect *bool   `json:"enable_post_logout_redirect,omitempty" tfsdk:"enable_post_logout_redirect"`
	ExtraFields              *string `json:"extra_fields,omitempty" tfsdk:"extra_fields"`
	ExtraScope               *string `json:"extra_scope,omitempty" tfsdk:"extra_scope"`
	IsActive                 *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	Label                    *string `json:"label,omitempty" tfsdk:"label"`
	ManagementUrl            *string `json:"management_url,omitempty" tfsdk:"management_url"`
	Provider                 *string `json:"provider,omitempty" tfsdk:"provider"`
	UserClaim                *string `json:"user_claim,omitempty" tfsdk:"user_claim"`
	UserField                *string `json:"user_field,omitempty" tfsdk:"user_field"`
	VerifySsl                *bool   `json:"verify_ssl,omitempty" tfsdk:"verify_ssl"`
}

type PatchedInvitationUpdateRequest struct {
	Email *string `json:"email,omitempty" tfsdk:"email"`
	Role  *string `json:"role,omitempty" tfsdk:"role"`
}

type PatchedInvoiceItemUpdateRequest struct {
	ArticleCode *string `json:"article_code,omitempty" tfsdk:"article_code"`
	End         *string `json:"end,omitempty" tfsdk:"end"`
	Quantity    *string `json:"quantity,omitempty" tfsdk:"quantity"`
	Start       *string `json:"start,omitempty" tfsdk:"start"`
	UnitPrice   *string `json:"unit_price,omitempty" tfsdk:"unit_price"`
}

type PatchedIssueRequest struct {
	Assignee           *string `json:"assignee,omitempty" tfsdk:"assignee"`
	Description        *string `json:"description,omitempty" tfsdk:"description"`
	IsReportedManually *bool   `json:"is_reported_manually,omitempty" tfsdk:"is_reported_manually"`
	Summary            *string `json:"summary,omitempty" tfsdk:"summary"`
}

type PatchedIssueStatusRequest struct {
	Name *string `json:"name,omitempty" tfsdk:"name"`
	Type *int64  `json:"type,omitempty" tfsdk:"type"`
}

type PatchedKeycloakUserGroupMembershipRequest struct {
	Email     *string `json:"email,omitempty" tfsdk:"email"`
	Role      *string `json:"role,omitempty" tfsdk:"role"`
	ScopeUuid *string `json:"scope_uuid,omitempty" tfsdk:"scope_uuid"`
	Username  *string `json:"username,omitempty" tfsdk:"username"`
}

type PatchedLexisLinkRequest struct {
	HeappeProjectId *int64 `json:"heappe_project_id,omitempty" tfsdk:"heappe_project_id"`
}

type PatchedMaintenanceAnnouncementOfferingRequest struct {
	ImpactDescription *string `json:"impact_description,omitempty" tfsdk:"impact_description"`
	ImpactLevel       *int64  `json:"impact_level,omitempty" tfsdk:"impact_level"`
	Maintenance       *string `json:"maintenance,omitempty" tfsdk:"maintenance"`
	Offering          *string `json:"offering,omitempty" tfsdk:"offering"`
}

type PatchedMaintenanceAnnouncementOfferingTemplateRequest struct {
	ImpactDescription   *string `json:"impact_description,omitempty" tfsdk:"impact_description"`
	ImpactLevel         *int64  `json:"impact_level,omitempty" tfsdk:"impact_level"`
	MaintenanceTemplate *string `json:"maintenance_template,omitempty" tfsdk:"maintenance_template"`
	Offering            *string `json:"offering,omitempty" tfsdk:"offering"`
}

type PatchedMaintenanceAnnouncementRequest struct {
	ExternalReferenceUrl *string `json:"external_reference_url,omitempty" tfsdk:"external_reference_url"`
	InternalNotes        *string `json:"internal_notes,omitempty" tfsdk:"internal_notes"`
	MaintenanceType      *int64  `json:"maintenance_type,omitempty" tfsdk:"maintenance_type"`
	Message              *string `json:"message,omitempty" tfsdk:"message"`
	Name                 *string `json:"name,omitempty" tfsdk:"name"`
	ScheduledEnd         *string `json:"scheduled_end,omitempty" tfsdk:"scheduled_end"`
	ScheduledStart       *string `json:"scheduled_start,omitempty" tfsdk:"scheduled_start"`
	ServiceProvider      *string `json:"service_provider,omitempty" tfsdk:"service_provider"`
}

type PatchedMaintenanceAnnouncementTemplateRequest struct {
	MaintenanceType *int64  `json:"maintenance_type,omitempty" tfsdk:"maintenance_type"`
	Message         *string `json:"message,omitempty" tfsdk:"message"`
	Name            *string `json:"name,omitempty" tfsdk:"name"`
	ServiceProvider *string `json:"service_provider,omitempty" tfsdk:"service_provider"`
}

type PatchedMarketplaceCategoryRequest struct {
	DefaultTenantCategory *bool   `json:"default_tenant_category,omitempty" tfsdk:"default_tenant_category"`
	DefaultVmCategory     *bool   `json:"default_vm_category,omitempty" tfsdk:"default_vm_category"`
	DefaultVolumeCategory *bool   `json:"default_volume_category,omitempty" tfsdk:"default_volume_category"`
	Description           *string `json:"description,omitempty" tfsdk:"description"`
	Group                 *string `json:"group,omitempty" tfsdk:"group"`
	Icon                  *string `json:"icon,omitempty" tfsdk:"icon"`
	Title                 *string `json:"title,omitempty" tfsdk:"title"`
}

type PatchedMarketplaceCategoryRequestForm struct {
	DefaultTenantCategory *bool   `json:"default_tenant_category,omitempty" tfsdk:"default_tenant_category"`
	DefaultVmCategory     *bool   `json:"default_vm_category,omitempty" tfsdk:"default_vm_category"`
	DefaultVolumeCategory *bool   `json:"default_volume_category,omitempty" tfsdk:"default_volume_category"`
	Description           *string `json:"description,omitempty" tfsdk:"description"`
	Group                 *string `json:"group,omitempty" tfsdk:"group"`
	Icon                  *string `json:"icon,omitempty" tfsdk:"icon"`
	Title                 *string `json:"title,omitempty" tfsdk:"title"`
}

type PatchedMarketplaceCategoryRequestMultipart struct {
	DefaultTenantCategory *bool   `json:"default_tenant_category,omitempty" tfsdk:"default_tenant_category"`
	DefaultVmCategory     *bool   `json:"default_vm_category,omitempty" tfsdk:"default_vm_category"`
	DefaultVolumeCategory *bool   `json:"default_volume_category,omitempty" tfsdk:"default_volume_category"`
	Description           *string `json:"description,omitempty" tfsdk:"description"`
	Group                 *string `json:"group,omitempty" tfsdk:"group"`
	Icon                  *string `json:"icon,omitempty" tfsdk:"icon"`
	Title                 *string `json:"title,omitempty" tfsdk:"title"`
}

type PatchedMessageTemplateRequest struct {
	Body    *string `json:"body,omitempty" tfsdk:"body"`
	Name    *string `json:"name,omitempty" tfsdk:"name"`
	Subject *string `json:"subject,omitempty" tfsdk:"subject"`
}

type PatchedMigrationDetailsRequest struct {
	ErrorMessage   *string         `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback *string         `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	Mappings       *MappingRequest `json:"mappings,omitempty" tfsdk:"mappings"`
}

type PatchedNetworkRBACPolicyRequest struct {
	Network      *string `json:"network,omitempty" tfsdk:"network"`
	PolicyType   *string `json:"policy_type,omitempty" tfsdk:"policy_type"`
	TargetTenant *string `json:"target_tenant,omitempty" tfsdk:"target_tenant"`
}

type PatchedNotificationRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Key         *string `json:"key,omitempty" tfsdk:"key"`
}

type PatchedNotificationTemplateDetailSerializersRequest struct {
	Name *string `json:"name,omitempty" tfsdk:"name"`
	Path *string `json:"path,omitempty" tfsdk:"path"`
}

type PatchedOfferingEstimatedCostPolicyRequest struct {
	Actions            *string  `json:"actions,omitempty" tfsdk:"actions"`
	ApplyToAll         *bool    `json:"apply_to_all,omitempty" tfsdk:"apply_to_all"`
	LimitCost          *int64   `json:"limit_cost,omitempty" tfsdk:"limit_cost"`
	OrganizationGroups []string `json:"organization_groups,omitempty" tfsdk:"organization_groups"`
	Period             *int64   `json:"period,omitempty" tfsdk:"period"`
	Scope              *string  `json:"scope,omitempty" tfsdk:"scope"`
}

type PatchedOfferingPartitionUpdateRequest struct {
	CpuBind          *int64  `json:"cpu_bind,omitempty" tfsdk:"cpu_bind"`
	DefCpuPerGpu     *int64  `json:"def_cpu_per_gpu,omitempty" tfsdk:"def_cpu_per_gpu"`
	DefMemPerCpu     *int64  `json:"def_mem_per_cpu,omitempty" tfsdk:"def_mem_per_cpu"`
	DefMemPerGpu     *int64  `json:"def_mem_per_gpu,omitempty" tfsdk:"def_mem_per_gpu"`
	DefMemPerNode    *int64  `json:"def_mem_per_node,omitempty" tfsdk:"def_mem_per_node"`
	DefaultTime      *int64  `json:"default_time,omitempty" tfsdk:"default_time"`
	ExclusiveTopo    *bool   `json:"exclusive_topo,omitempty" tfsdk:"exclusive_topo"`
	ExclusiveUser    *bool   `json:"exclusive_user,omitempty" tfsdk:"exclusive_user"`
	GraceTime        *int64  `json:"grace_time,omitempty" tfsdk:"grace_time"`
	MaxCpusPerNode   *int64  `json:"max_cpus_per_node,omitempty" tfsdk:"max_cpus_per_node"`
	MaxCpusPerSocket *int64  `json:"max_cpus_per_socket,omitempty" tfsdk:"max_cpus_per_socket"`
	MaxMemPerCpu     *int64  `json:"max_mem_per_cpu,omitempty" tfsdk:"max_mem_per_cpu"`
	MaxMemPerNode    *int64  `json:"max_mem_per_node,omitempty" tfsdk:"max_mem_per_node"`
	MaxNodes         *int64  `json:"max_nodes,omitempty" tfsdk:"max_nodes"`
	MaxTime          *int64  `json:"max_time,omitempty" tfsdk:"max_time"`
	MinNodes         *int64  `json:"min_nodes,omitempty" tfsdk:"min_nodes"`
	PartitionName    *string `json:"partition_name,omitempty" tfsdk:"partition_name"`
	PartitionUuid    *string `json:"partition_uuid,omitempty" tfsdk:"partition_uuid"`
	PriorityTier     *int64  `json:"priority_tier,omitempty" tfsdk:"priority_tier"`
	Qos              *string `json:"qos,omitempty" tfsdk:"qos"`
	ReqResv          *bool   `json:"req_resv,omitempty" tfsdk:"req_resv"`
}

type PatchedOfferingSoftwareCatalogUpdateRequest struct {
	Catalog             *string `json:"catalog,omitempty" tfsdk:"catalog"`
	OfferingCatalogUuid *string `json:"offering_catalog_uuid,omitempty" tfsdk:"offering_catalog_uuid"`
	Partition           *string `json:"partition,omitempty" tfsdk:"partition"`
}

type PatchedOfferingTermsOfServiceRequest struct {
	GracePeriodDays    *int64  `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`
	IsActive           *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	TermsOfService     *string `json:"terms_of_service,omitempty" tfsdk:"terms_of_service"`
	TermsOfServiceLink *string `json:"terms_of_service_link,omitempty" tfsdk:"terms_of_service_link"`
}

type PatchedOfferingUsagePolicyRequest struct {
	Actions            *string                               `json:"actions,omitempty" tfsdk:"actions"`
	ApplyToAll         *bool                                 `json:"apply_to_all,omitempty" tfsdk:"apply_to_all"`
	ComponentLimitsSet []NestedOfferingComponentLimitRequest `json:"component_limits_set,omitempty" tfsdk:"component_limits_set"`
	OrganizationGroups []string                              `json:"organization_groups,omitempty" tfsdk:"organization_groups"`
	Period             *int64                                `json:"period,omitempty" tfsdk:"period"`
	Scope              *string                               `json:"scope,omitempty" tfsdk:"scope"`
}

type PatchedOfferingUserRequest struct {
	Offering     *string `json:"offering,omitempty" tfsdk:"offering"`
	OfferingUuid *string `json:"offering_uuid,omitempty" tfsdk:"offering_uuid"`
	User         *string `json:"user,omitempty" tfsdk:"user"`
	UserUuid     *string `json:"user_uuid,omitempty" tfsdk:"user_uuid"`
	Username     *string `json:"username,omitempty" tfsdk:"username"`
}

type PatchedOfferingUserRoleRequest struct {
	Name     *string `json:"name,omitempty" tfsdk:"name"`
	Offering *string `json:"offering,omitempty" tfsdk:"offering"`
}

type PatchedOfferingUserServiceProviderCommentRequest struct {
	ServiceProviderComment    *string `json:"service_provider_comment,omitempty" tfsdk:"service_provider_comment"`
	ServiceProviderCommentUrl *string `json:"service_provider_comment_url,omitempty" tfsdk:"service_provider_comment_url"`
}

type PatchedOnboardingCountryChecklistConfigurationRequest struct {
	Checklist *string `json:"checklist,omitempty" tfsdk:"checklist"`
	Country   *string `json:"country,omitempty" tfsdk:"country"`
	IsActive  *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
}

type PatchedOnboardingJustificationRequest struct {
	UserJustification *string `json:"user_justification,omitempty" tfsdk:"user_justification"`
	Verification      *string `json:"verification,omitempty" tfsdk:"verification"`
}

type PatchedOnboardingQuestionMetadataRequest struct {
	IntentField         *string `json:"intent_field,omitempty" tfsdk:"intent_field"`
	MapsToCustomerField *string `json:"maps_to_customer_field,omitempty" tfsdk:"maps_to_customer_field"`
	Question            *string `json:"question,omitempty" tfsdk:"question"`
}

type PatchedOnboardingVerificationRequest struct {
	Country               *string `json:"country,omitempty" tfsdk:"country"`
	ExpiresAt             *string `json:"expires_at,omitempty" tfsdk:"expires_at"`
	LegalName             *string `json:"legal_name,omitempty" tfsdk:"legal_name"`
	LegalPersonIdentifier *string `json:"legal_person_identifier,omitempty" tfsdk:"legal_person_identifier"`
}

type PatchedOpenStackBackupRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	KeptUntil   *string `json:"kept_until,omitempty" tfsdk:"kept_until"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type PatchedOpenStackInstanceRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type PatchedOpenStackNetworkRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type PatchedOpenStackPortRequest struct {
	Description    *string                                   `json:"description,omitempty" tfsdk:"description"`
	Name           *string                                   `json:"name,omitempty" tfsdk:"name"`
	SecurityGroups []OpenStackPortNestedSecurityGroupRequest `json:"security_groups,omitempty" tfsdk:"security_groups"`
	TargetTenant   *string                                   `json:"target_tenant,omitempty" tfsdk:"target_tenant"`
}

type PatchedOpenStackSecurityGroupUpdateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type PatchedOpenStackServerGroupRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
	Policy      *string `json:"policy,omitempty" tfsdk:"policy"`
}

type PatchedOpenStackSnapshotRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	KeptUntil   *string `json:"kept_until,omitempty" tfsdk:"kept_until"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type PatchedOpenStackSubNetRequest struct {
	AllocationPools []OpenStackSubNetAllocationPoolRequest `json:"allocation_pools,omitempty" tfsdk:"allocation_pools"`
	Cidr            *string                                `json:"cidr,omitempty" tfsdk:"cidr"`
	Description     *string                                `json:"description,omitempty" tfsdk:"description"`
	DisableGateway  *bool                                  `json:"disable_gateway,omitempty" tfsdk:"disable_gateway"`
	DnsNameservers  []string                               `json:"dns_nameservers,omitempty" tfsdk:"dns_nameservers"`
	GatewayIp       *string                                `json:"gateway_ip,omitempty" tfsdk:"gateway_ip"`
	HostRoutes      []OpenStackStaticRouteRequest          `json:"host_routes,omitempty" tfsdk:"host_routes"`
	Name            *string                                `json:"name,omitempty" tfsdk:"name"`
}

type PatchedOpenStackTenantRequest struct {
	AvailabilityZone            *string                               `json:"availability_zone,omitempty" tfsdk:"availability_zone"`
	DefaultVolumeTypeName       *string                               `json:"default_volume_type_name,omitempty" tfsdk:"default_volume_type_name"`
	Description                 *string                               `json:"description,omitempty" tfsdk:"description"`
	Name                        *string                               `json:"name,omitempty" tfsdk:"name"`
	SecurityGroups              []OpenStackTenantSecurityGroupRequest `json:"security_groups,omitempty" tfsdk:"security_groups"`
	SkipCreationOfDefaultRouter *bool                                 `json:"skip_creation_of_default_router,omitempty" tfsdk:"skip_creation_of_default_router"`
	SkipCreationOfDefaultSubnet *bool                                 `json:"skip_creation_of_default_subnet,omitempty" tfsdk:"skip_creation_of_default_subnet"`
}

type PatchedOpenStackVolumeRequest struct {
	Bootable    *bool   `json:"bootable,omitempty" tfsdk:"bootable"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type PatchedOrganizationGroupRequest struct {
	Name   *string `json:"name,omitempty" tfsdk:"name"`
	Parent *string `json:"parent,omitempty" tfsdk:"parent"`
}

type PatchedPaymentProfileRequest struct {
	Attributes   *PaymentProfileAttributesRequest `json:"attributes,omitempty" tfsdk:"attributes"`
	IsActive     *bool                            `json:"is_active,omitempty" tfsdk:"is_active"`
	Name         *string                          `json:"name,omitempty" tfsdk:"name"`
	Organization *string                          `json:"organization,omitempty" tfsdk:"organization"`
	PaymentType  *string                          `json:"payment_type,omitempty" tfsdk:"payment_type"`
}

type PatchedPaymentRequest struct {
	DateOfPayment *string `json:"date_of_payment,omitempty" tfsdk:"date_of_payment"`
	Profile       *string `json:"profile,omitempty" tfsdk:"profile"`
	Proof         *string `json:"proof,omitempty" tfsdk:"proof"`
	Sum           *string `json:"sum,omitempty" tfsdk:"sum"`
}

type PatchedPaymentRequestForm struct {
	DateOfPayment *string `json:"date_of_payment,omitempty" tfsdk:"date_of_payment"`
	Profile       *string `json:"profile,omitempty" tfsdk:"profile"`
	Proof         *string `json:"proof,omitempty" tfsdk:"proof"`
	Sum           *string `json:"sum,omitempty" tfsdk:"sum"`
}

type PatchedPaymentRequestMultipart struct {
	DateOfPayment *string `json:"date_of_payment,omitempty" tfsdk:"date_of_payment"`
	Profile       *string `json:"profile,omitempty" tfsdk:"profile"`
	Proof         *string `json:"proof,omitempty" tfsdk:"proof"`
	Sum           *string `json:"sum,omitempty" tfsdk:"sum"`
}

type PatchedProjectCreditRequest struct {
	ApplyAsMinimalConsumption                   *bool   `json:"apply_as_minimal_consumption,omitempty" tfsdk:"apply_as_minimal_consumption"`
	EndDate                                     *string `json:"end_date,omitempty" tfsdk:"end_date"`
	ExpectedConsumption                         *string `json:"expected_consumption,omitempty" tfsdk:"expected_consumption"`
	GraceCoefficient                            *string `json:"grace_coefficient,omitempty" tfsdk:"grace_coefficient"`
	MarkUnusedCreditAsSpentOnProjectTermination *bool   `json:"mark_unused_credit_as_spent_on_project_termination,omitempty" tfsdk:"mark_unused_credit_as_spent_on_project_termination"`
	MinimalConsumptionLogic                     *string `json:"minimal_consumption_logic,omitempty" tfsdk:"minimal_consumption_logic"`
	Project                                     *string `json:"project,omitempty" tfsdk:"project"`
	Value                                       *string `json:"value,omitempty" tfsdk:"value"`
}

type PatchedProjectEstimatedCostPolicyRequest struct {
	Actions   *string `json:"actions,omitempty" tfsdk:"actions"`
	LimitCost *int64  `json:"limit_cost,omitempty" tfsdk:"limit_cost"`
	Period    *int64  `json:"period,omitempty" tfsdk:"period"`
	Scope     *string `json:"scope,omitempty" tfsdk:"scope"`
}

type PatchedProjectInfoRequest struct {
	AllowedDestinations *string `json:"allowed_destinations,omitempty" tfsdk:"allowed_destinations"`
	Project             *string `json:"project,omitempty" tfsdk:"project"`
	Shortname           *string `json:"shortname,omitempty" tfsdk:"shortname"`
}

type PatchedProjectRequest struct {
	BackendId       *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Customer        *string `json:"customer,omitempty" tfsdk:"customer"`
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	EndDate         *string `json:"end_date,omitempty" tfsdk:"end_date"`
	GracePeriodDays *int64  `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`
	Image           *string `json:"image,omitempty" tfsdk:"image"`
	IsIndustry      *bool   `json:"is_industry,omitempty" tfsdk:"is_industry"`
	Kind            *string `json:"kind,omitempty" tfsdk:"kind"`
	Name            *string `json:"name,omitempty" tfsdk:"name"`
	OecdFos2007Code *string `json:"oecd_fos_2007_code,omitempty" tfsdk:"oecd_fos_2007_code"`
	Slug            *string `json:"slug,omitempty" tfsdk:"slug"`
	StaffNotes      *string `json:"staff_notes,omitempty" tfsdk:"staff_notes"`
	StartDate       *string `json:"start_date,omitempty" tfsdk:"start_date"`
	Type            *string `json:"type,omitempty" tfsdk:"type"`
}

type PatchedProjectRequestForm struct {
	BackendId       *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Customer        *string `json:"customer,omitempty" tfsdk:"customer"`
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	EndDate         *string `json:"end_date,omitempty" tfsdk:"end_date"`
	GracePeriodDays *int64  `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`
	Image           *string `json:"image,omitempty" tfsdk:"image"`
	IsIndustry      *bool   `json:"is_industry,omitempty" tfsdk:"is_industry"`
	Kind            *string `json:"kind,omitempty" tfsdk:"kind"`
	Name            *string `json:"name,omitempty" tfsdk:"name"`
	OecdFos2007Code *string `json:"oecd_fos_2007_code,omitempty" tfsdk:"oecd_fos_2007_code"`
	Slug            *string `json:"slug,omitempty" tfsdk:"slug"`
	StaffNotes      *string `json:"staff_notes,omitempty" tfsdk:"staff_notes"`
	StartDate       *string `json:"start_date,omitempty" tfsdk:"start_date"`
	Type            *string `json:"type,omitempty" tfsdk:"type"`
}

type PatchedProjectRequestMultipart struct {
	BackendId       *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Customer        *string `json:"customer,omitempty" tfsdk:"customer"`
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	EndDate         *string `json:"end_date,omitempty" tfsdk:"end_date"`
	GracePeriodDays *int64  `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`
	Image           *string `json:"image,omitempty" tfsdk:"image"`
	IsIndustry      *bool   `json:"is_industry,omitempty" tfsdk:"is_industry"`
	Kind            *string `json:"kind,omitempty" tfsdk:"kind"`
	Name            *string `json:"name,omitempty" tfsdk:"name"`
	OecdFos2007Code *string `json:"oecd_fos_2007_code,omitempty" tfsdk:"oecd_fos_2007_code"`
	Slug            *string `json:"slug,omitempty" tfsdk:"slug"`
	StaffNotes      *string `json:"staff_notes,omitempty" tfsdk:"staff_notes"`
	StartDate       *string `json:"start_date,omitempty" tfsdk:"start_date"`
	Type            *string `json:"type,omitempty" tfsdk:"type"`
}

type PatchedProjectServiceAccountRequest struct {
	Description         *string `json:"description,omitempty" tfsdk:"description"`
	Email               *string `json:"email,omitempty" tfsdk:"email"`
	ErrorTraceback      *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	PreferredIdentifier *string `json:"preferred_identifier,omitempty" tfsdk:"preferred_identifier"`
	Project             *string `json:"project,omitempty" tfsdk:"project"`
	Username            *string `json:"username,omitempty" tfsdk:"username"`
}

type PatchedProjectTemplateRequest struct {
	ApprovalLimit  *string  `json:"approval_limit,omitempty" tfsdk:"approval_limit"`
	Customer       *string  `json:"customer,omitempty" tfsdk:"customer"`
	Key            *string  `json:"key,omitempty" tfsdk:"key"`
	MaxCreditLimit *string  `json:"max_credit_limit,omitempty" tfsdk:"max_credit_limit"`
	Name           *string  `json:"name,omitempty" tfsdk:"name"`
	Offering       *string  `json:"offering,omitempty" tfsdk:"offering"`
	Offerings      []string `json:"offerings,omitempty" tfsdk:"offerings"`
	Portal         *string  `json:"portal,omitempty" tfsdk:"portal"`
	Provider       *string  `json:"provider,omitempty" tfsdk:"provider"`
	Shortname      *string  `json:"shortname,omitempty" tfsdk:"shortname"`
}

type PatchedProposalProjectRoleMappingRequest struct {
	Call         *string `json:"call,omitempty" tfsdk:"call"`
	ProjectRole  *string `json:"project_role,omitempty" tfsdk:"project_role"`
	ProposalRole *string `json:"proposal_role,omitempty" tfsdk:"proposal_role"`
}

type PatchedProposalReviewRequest struct {
	CommentProjectDescription             *string `json:"comment_project_description,omitempty" tfsdk:"comment_project_description"`
	CommentProjectDuration                *string `json:"comment_project_duration,omitempty" tfsdk:"comment_project_duration"`
	CommentProjectHasCivilianPurpose      *string `json:"comment_project_has_civilian_purpose,omitempty" tfsdk:"comment_project_has_civilian_purpose"`
	CommentProjectIsConfidential          *string `json:"comment_project_is_confidential,omitempty" tfsdk:"comment_project_is_confidential"`
	CommentProjectSummary                 *string `json:"comment_project_summary,omitempty" tfsdk:"comment_project_summary"`
	CommentProjectSupportingDocumentation *string `json:"comment_project_supporting_documentation,omitempty" tfsdk:"comment_project_supporting_documentation"`
	CommentProjectTitle                   *string `json:"comment_project_title,omitempty" tfsdk:"comment_project_title"`
	CommentResourceRequests               *string `json:"comment_resource_requests,omitempty" tfsdk:"comment_resource_requests"`
	CommentTeam                           *string `json:"comment_team,omitempty" tfsdk:"comment_team"`
	SummaryPrivateComment                 *string `json:"summary_private_comment,omitempty" tfsdk:"summary_private_comment"`
	SummaryPublicComment                  *string `json:"summary_public_comment,omitempty" tfsdk:"summary_public_comment"`
	SummaryScore                          *int64  `json:"summary_score,omitempty" tfsdk:"summary_score"`
}

type PatchedProtectedCallRequest struct {
	BackendId                           *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	ComplianceChecklist                 *string `json:"compliance_checklist,omitempty" tfsdk:"compliance_checklist"`
	CreatedBy                           *string `json:"created_by,omitempty" tfsdk:"created_by"`
	Description                         *string `json:"description,omitempty" tfsdk:"description"`
	ExternalUrl                         *string `json:"external_url,omitempty" tfsdk:"external_url"`
	FixedDurationInDays                 *int64  `json:"fixed_duration_in_days,omitempty" tfsdk:"fixed_duration_in_days"`
	Name                                *string `json:"name,omitempty" tfsdk:"name"`
	ProposalSlugTemplate                *string `json:"proposal_slug_template,omitempty" tfsdk:"proposal_slug_template"`
	ReferenceCode                       *string `json:"reference_code,omitempty" tfsdk:"reference_code"`
	ReviewerIdentityVisibleToSubmitters *bool   `json:"reviewer_identity_visible_to_submitters,omitempty" tfsdk:"reviewer_identity_visible_to_submitters"`
	ReviewsVisibleToSubmitters          *bool   `json:"reviews_visible_to_submitters,omitempty" tfsdk:"reviews_visible_to_submitters"`
	Slug                                *string `json:"slug,omitempty" tfsdk:"slug"`
}

type PatchedProtectedRoundRequest struct {
	AllocationDate           *string `json:"allocation_date,omitempty" tfsdk:"allocation_date"`
	AllocationTime           *string `json:"allocation_time,omitempty" tfsdk:"allocation_time"`
	CutoffTime               *string `json:"cutoff_time,omitempty" tfsdk:"cutoff_time"`
	DecidingEntity           *string `json:"deciding_entity,omitempty" tfsdk:"deciding_entity"`
	MinimalAverageScoring    *string `json:"minimal_average_scoring,omitempty" tfsdk:"minimal_average_scoring"`
	MinimumNumberOfReviewers *int64  `json:"minimum_number_of_reviewers,omitempty" tfsdk:"minimum_number_of_reviewers"`
	ReviewDurationInDays     *int64  `json:"review_duration_in_days,omitempty" tfsdk:"review_duration_in_days"`
	ReviewStrategy           *string `json:"review_strategy,omitempty" tfsdk:"review_strategy"`
	StartTime                *string `json:"start_time,omitempty" tfsdk:"start_time"`
}

type PatchedProviderPlanDetailsRequest struct {
	Archived    *bool   `json:"archived,omitempty" tfsdk:"archived"`
	ArticleCode *string `json:"article_code,omitempty" tfsdk:"article_code"`
	BackendId   *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	MaxAmount   *int64  `json:"max_amount,omitempty" tfsdk:"max_amount"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
	Unit        *string `json:"unit,omitempty" tfsdk:"unit"`
	UnitPrice   *string `json:"unit_price,omitempty" tfsdk:"unit_price"`
}

type PatchedQuestionAdminRequest struct {
	AlwaysRequiresReview    *bool   `json:"always_requires_review,omitempty" tfsdk:"always_requires_review"`
	AlwaysShowGuidance      *bool   `json:"always_show_guidance,omitempty" tfsdk:"always_show_guidance"`
	Checklist               *string `json:"checklist,omitempty" tfsdk:"checklist"`
	DependencyLogicOperator *string `json:"dependency_logic_operator,omitempty" tfsdk:"dependency_logic_operator"`
	Description             *string `json:"description,omitempty" tfsdk:"description"`
	GuidanceOperator        *string `json:"guidance_operator,omitempty" tfsdk:"guidance_operator"`
	MaxFileSizeMb           *int64  `json:"max_file_size_mb,omitempty" tfsdk:"max_file_size_mb"`
	MaxFilesCount           *int64  `json:"max_files_count,omitempty" tfsdk:"max_files_count"`
	MaxValue                *string `json:"max_value,omitempty" tfsdk:"max_value"`
	MinValue                *string `json:"min_value,omitempty" tfsdk:"min_value"`
	Operator                *string `json:"operator,omitempty" tfsdk:"operator"`
	Order                   *int64  `json:"order,omitempty" tfsdk:"order"`
	QuestionType            *string `json:"question_type,omitempty" tfsdk:"question_type"`
	Required                *bool   `json:"required,omitempty" tfsdk:"required"`
	UserGuidance            *string `json:"user_guidance,omitempty" tfsdk:"user_guidance"`
}

type PatchedQuestionDependencyRequest struct {
	DependsOnQuestion *string `json:"depends_on_question,omitempty" tfsdk:"depends_on_question"`
	Operator          *string `json:"operator,omitempty" tfsdk:"operator"`
	Question          *string `json:"question,omitempty" tfsdk:"question"`
}

type PatchedQuestionOptionsAdminRequest struct {
	Label *string `json:"label,omitempty" tfsdk:"label"`
	Order *int64  `json:"order,omitempty" tfsdk:"order"`
}

type PatchedRancherApplicationRequest struct {
	BackendId       *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage    *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback  *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	Name            *string `json:"name,omitempty" tfsdk:"name"`
	Namespace       *string `json:"namespace,omitempty" tfsdk:"namespace"`
	NamespaceName   *string `json:"namespace_name,omitempty" tfsdk:"namespace_name"`
	Project         *string `json:"project,omitempty" tfsdk:"project"`
	RancherProject  *string `json:"rancher_project,omitempty" tfsdk:"rancher_project"`
	RuntimeState    *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceSettings *string `json:"service_settings,omitempty" tfsdk:"service_settings"`
	Template        *string `json:"template,omitempty" tfsdk:"template"`
	Version         *string `json:"version,omitempty" tfsdk:"version"`
}

type PatchedRancherCatalogRequest struct {
	Branch      *string `json:"branch,omitempty" tfsdk:"branch"`
	CatalogUrl  *string `json:"catalog_url,omitempty" tfsdk:"catalog_url"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
	Scope       *string `json:"scope,omitempty" tfsdk:"scope"`
}

type PatchedRancherClusterRequest struct {
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	InstallLonghorn *bool   `json:"install_longhorn,omitempty" tfsdk:"install_longhorn"`
	Name            *string `json:"name,omitempty" tfsdk:"name"`
	SshPublicKey    *string `json:"ssh_public_key,omitempty" tfsdk:"ssh_public_key"`
	VmProject       *string `json:"vm_project,omitempty" tfsdk:"vm_project"`
}

type PatchedRancherHPARequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	MaxReplicas *int64  `json:"max_replicas,omitempty" tfsdk:"max_replicas"`
	MinReplicas *int64  `json:"min_replicas,omitempty" tfsdk:"min_replicas"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
	Workload    *string `json:"workload,omitempty" tfsdk:"workload"`
}

type PatchedRancherIngressRequest struct {
	BackendId       *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage    *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback  *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	Name            *string `json:"name,omitempty" tfsdk:"name"`
	Namespace       *string `json:"namespace,omitempty" tfsdk:"namespace"`
	Project         *string `json:"project,omitempty" tfsdk:"project"`
	RancherProject  *string `json:"rancher_project,omitempty" tfsdk:"rancher_project"`
	RuntimeState    *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceSettings *string `json:"service_settings,omitempty" tfsdk:"service_settings"`
}

type PatchedRancherServiceRequest struct {
	BackendId       *string                        `json:"backend_id,omitempty" tfsdk:"backend_id"`
	ClusterIp       *string                        `json:"cluster_ip,omitempty" tfsdk:"cluster_ip"`
	Description     *string                        `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage    *string                        `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback  *string                        `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	Name            *string                        `json:"name,omitempty" tfsdk:"name"`
	Namespace       *string                        `json:"namespace,omitempty" tfsdk:"namespace"`
	Project         *string                        `json:"project,omitempty" tfsdk:"project"`
	RuntimeState    *string                        `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceSettings *string                        `json:"service_settings,omitempty" tfsdk:"service_settings"`
	TargetWorkloads []RancherNestedWorkloadRequest `json:"target_workloads,omitempty" tfsdk:"target_workloads"`
}

type PatchedRancherWorkloadRequest struct {
	Cluster      *string `json:"cluster,omitempty" tfsdk:"cluster"`
	Name         *string `json:"name,omitempty" tfsdk:"name"`
	Namespace    *string `json:"namespace,omitempty" tfsdk:"namespace"`
	Project      *string `json:"project,omitempty" tfsdk:"project"`
	RuntimeState *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	Scale        *int64  `json:"scale,omitempty" tfsdk:"scale"`
}

type PatchedRemoteAllocationRequest struct {
	Description             *string `json:"description,omitempty" tfsdk:"description"`
	Name                    *string `json:"name,omitempty" tfsdk:"name"`
	NodeLimit               *int64  `json:"node_limit,omitempty" tfsdk:"node_limit"`
	RemoteProjectIdentifier *string `json:"remote_project_identifier,omitempty" tfsdk:"remote_project_identifier"`
}

type PatchedRemoteSynchronisationRequest struct {
	ApiUrl                 *string                            `json:"api_url,omitempty" tfsdk:"api_url"`
	IsActive               *bool                              `json:"is_active,omitempty" tfsdk:"is_active"`
	LocalServiceProvider   *string                            `json:"local_service_provider,omitempty" tfsdk:"local_service_provider"`
	RemoteOrganizationName *string                            `json:"remote_organization_name,omitempty" tfsdk:"remote_organization_name"`
	RemotelocalcategorySet []NestedRemoteLocalCategoryRequest `json:"remotelocalcategory_set,omitempty" tfsdk:"remotelocalcategory_set"`
	Token                  *string                            `json:"token,omitempty" tfsdk:"token"`
}

type PatchedRequestTypeAdminRequest struct {
	IsActive      *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	IssueTypeName *string `json:"issue_type_name,omitempty" tfsdk:"issue_type_name"`
	Name          *string `json:"name,omitempty" tfsdk:"name"`
	Order         *int64  `json:"order,omitempty" tfsdk:"order"`
}

type PatchedRequestedOfferingRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Plan        *string `json:"plan,omitempty" tfsdk:"plan"`
}

type PatchedRequestedResourceRequest struct {
	CallResourceTemplateUuid *string `json:"call_resource_template_uuid,omitempty" tfsdk:"call_resource_template_uuid"`
	Description              *string `json:"description,omitempty" tfsdk:"description"`
	RequestedOfferingUuid    *string `json:"requested_offering_uuid,omitempty" tfsdk:"requested_offering_uuid"`
}

type PatchedResourceUpdateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	EndDate     *string `json:"end_date,omitempty" tfsdk:"end_date"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type PatchedRobotAccountRequest struct {
	Description     *string  `json:"description,omitempty" tfsdk:"description"`
	Resource        *string  `json:"resource,omitempty" tfsdk:"resource"`
	ResponsibleUser *string  `json:"responsible_user,omitempty" tfsdk:"responsible_user"`
	Type            *string  `json:"type,omitempty" tfsdk:"type"`
	Username        *string  `json:"username,omitempty" tfsdk:"username"`
	Users           []string `json:"users,omitempty" tfsdk:"users"`
}

type PatchedRoleDetailsRequest struct {
	Description   *string `json:"description,omitempty" tfsdk:"description"`
	DescriptionAr *string `json:"description_ar,omitempty" tfsdk:"description_ar"`
	DescriptionCs *string `json:"description_cs,omitempty" tfsdk:"description_cs"`
	DescriptionDa *string `json:"description_da,omitempty" tfsdk:"description_da"`
	DescriptionDe *string `json:"description_de,omitempty" tfsdk:"description_de"`
	DescriptionEn *string `json:"description_en,omitempty" tfsdk:"description_en"`
	DescriptionEs *string `json:"description_es,omitempty" tfsdk:"description_es"`
	DescriptionEt *string `json:"description_et,omitempty" tfsdk:"description_et"`
	DescriptionFr *string `json:"description_fr,omitempty" tfsdk:"description_fr"`
	DescriptionIt *string `json:"description_it,omitempty" tfsdk:"description_it"`
	DescriptionLt *string `json:"description_lt,omitempty" tfsdk:"description_lt"`
	DescriptionLv *string `json:"description_lv,omitempty" tfsdk:"description_lv"`
	DescriptionNb *string `json:"description_nb,omitempty" tfsdk:"description_nb"`
	DescriptionRu *string `json:"description_ru,omitempty" tfsdk:"description_ru"`
	DescriptionSv *string `json:"description_sv,omitempty" tfsdk:"description_sv"`
	IsActive      *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	Name          *string `json:"name,omitempty" tfsdk:"name"`
}

type PatchedRuleRequest struct {
	Customer                          *string  `json:"customer,omitempty" tfsdk:"customer"`
	Name                              *string  `json:"name,omitempty" tfsdk:"name"`
	Plan                              *string  `json:"plan,omitempty" tfsdk:"plan"`
	ProjectRole                       *string  `json:"project_role,omitempty" tfsdk:"project_role"`
	ProjectRoleName                   *string  `json:"project_role_name,omitempty" tfsdk:"project_role_name"`
	UseUserOrganizationAsCustomerName *bool    `json:"use_user_organization_as_customer_name,omitempty" tfsdk:"use_user_organization_as_customer_name"`
	UserAffiliations                  []string `json:"user_affiliations,omitempty" tfsdk:"user_affiliations"`
	UserEmailPatterns                 []string `json:"user_email_patterns,omitempty" tfsdk:"user_email_patterns"`
}

type PatchedScreenshotRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type PatchedSectionRequest struct {
	Category     *string `json:"category,omitempty" tfsdk:"category"`
	IsStandalone *bool   `json:"is_standalone,omitempty" tfsdk:"is_standalone"`
	Key          *string `json:"key,omitempty" tfsdk:"key"`
	Title        *string `json:"title,omitempty" tfsdk:"title"`
}

type PatchedServiceProviderRequest struct {
	Description         *string `json:"description,omitempty" tfsdk:"description"`
	EnableNotifications *bool   `json:"enable_notifications,omitempty" tfsdk:"enable_notifications"`
	Image               *string `json:"image,omitempty" tfsdk:"image"`
}

type PatchedServiceProviderRequestForm struct {
	Description         *string `json:"description,omitempty" tfsdk:"description"`
	EnableNotifications *bool   `json:"enable_notifications,omitempty" tfsdk:"enable_notifications"`
	Image               *string `json:"image,omitempty" tfsdk:"image"`
}

type PatchedServiceProviderRequestMultipart struct {
	Description         *string `json:"description,omitempty" tfsdk:"description"`
	EnableNotifications *bool   `json:"enable_notifications,omitempty" tfsdk:"enable_notifications"`
	Image               *string `json:"image,omitempty" tfsdk:"image"`
}

type PatchedSlurmAllocationRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type PatchedSlurmPeriodicUsagePolicyRequest struct {
	Actions                *string                               `json:"actions,omitempty" tfsdk:"actions"`
	ApplyToAll             *bool                                 `json:"apply_to_all,omitempty" tfsdk:"apply_to_all"`
	CarryoverEnabled       *bool                                 `json:"carryover_enabled,omitempty" tfsdk:"carryover_enabled"`
	ComponentLimitsSet     []NestedOfferingComponentLimitRequest `json:"component_limits_set,omitempty" tfsdk:"component_limits_set"`
	FairshareDecayHalfLife *int64                                `json:"fairshare_decay_half_life,omitempty" tfsdk:"fairshare_decay_half_life"`
	GraceRatio             *float64                              `json:"grace_ratio,omitempty" tfsdk:"grace_ratio"`
	LimitType              *string                               `json:"limit_type,omitempty" tfsdk:"limit_type"`
	OrganizationGroups     []string                              `json:"organization_groups,omitempty" tfsdk:"organization_groups"`
	Period                 *int64                                `json:"period,omitempty" tfsdk:"period"`
	QosStrategy            *string                               `json:"qos_strategy,omitempty" tfsdk:"qos_strategy"`
	RawUsageReset          *bool                                 `json:"raw_usage_reset,omitempty" tfsdk:"raw_usage_reset"`
	Scope                  *string                               `json:"scope,omitempty" tfsdk:"scope"`
	TresBillingEnabled     *bool                                 `json:"tres_billing_enabled,omitempty" tfsdk:"tres_billing_enabled"`
}

type PatchedSoftwareCatalogRequest struct {
	AutoUpdateEnabled *bool   `json:"auto_update_enabled,omitempty" tfsdk:"auto_update_enabled"`
	CatalogType       *string `json:"catalog_type,omitempty" tfsdk:"catalog_type"`
	Description       *string `json:"description,omitempty" tfsdk:"description"`
	Name              *string `json:"name,omitempty" tfsdk:"name"`
	SourceUrl         *string `json:"source_url,omitempty" tfsdk:"source_url"`
	UpdateErrors      *string `json:"update_errors,omitempty" tfsdk:"update_errors"`
	Version           *string `json:"version,omitempty" tfsdk:"version"`
}

type PatchedSoftwarePackageRequest struct {
	Catalog        *string `json:"catalog,omitempty" tfsdk:"catalog"`
	Description    *string `json:"description,omitempty" tfsdk:"description"`
	Homepage       *string `json:"homepage,omitempty" tfsdk:"homepage"`
	IsExtension    *bool   `json:"is_extension,omitempty" tfsdk:"is_extension"`
	Name           *string `json:"name,omitempty" tfsdk:"name"`
	ParentSoftware *string `json:"parent_software,omitempty" tfsdk:"parent_software"`
}

type PatchedTemplateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	IssueType   *string `json:"issue_type,omitempty" tfsdk:"issue_type"`
	Name        *string `json:"name,omitempty" tfsdk:"name"`
}

type PatchedUserAgreementRequest struct {
	AgreementType *string `json:"agreement_type,omitempty" tfsdk:"agreement_type"`
	Content       *string `json:"content,omitempty" tfsdk:"content"`
	Language      *string `json:"language,omitempty" tfsdk:"language"`
}

type PatchedUserInfoRequest struct {
	Shortname *string `json:"shortname,omitempty" tfsdk:"shortname"`
	User      *string `json:"user,omitempty" tfsdk:"user"`
}

type PatchedUserOfferingConsentRequest struct {
	Version *string `json:"version,omitempty" tfsdk:"version"`
}

type PatchedUserRequest struct {
	AgreeWithPolicy      *bool   `json:"agree_with_policy,omitempty" tfsdk:"agree_with_policy"`
	BirthDate            *string `json:"birth_date,omitempty" tfsdk:"birth_date"`
	Description          *string `json:"description,omitempty" tfsdk:"description"`
	FirstName            *string `json:"first_name,omitempty" tfsdk:"first_name"`
	Image                *string `json:"image,omitempty" tfsdk:"image"`
	IsActive             *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	IsStaff              *bool   `json:"is_staff,omitempty" tfsdk:"is_staff"`
	IsSupport            *bool   `json:"is_support,omitempty" tfsdk:"is_support"`
	JobTitle             *string `json:"job_title,omitempty" tfsdk:"job_title"`
	LastName             *string `json:"last_name,omitempty" tfsdk:"last_name"`
	NativeName           *string `json:"native_name,omitempty" tfsdk:"native_name"`
	NotificationsEnabled *bool   `json:"notifications_enabled,omitempty" tfsdk:"notifications_enabled"`
	Organization         *string `json:"organization,omitempty" tfsdk:"organization"`
	PhoneNumber          *string `json:"phone_number,omitempty" tfsdk:"phone_number"`
	PreferredLanguage    *string `json:"preferred_language,omitempty" tfsdk:"preferred_language"`
	Slug                 *string `json:"slug,omitempty" tfsdk:"slug"`
	TokenLifetime        *int64  `json:"token_lifetime,omitempty" tfsdk:"token_lifetime"`
	Username             *string `json:"username,omitempty" tfsdk:"username"`
}

type PatchedUserRequestForm struct {
	AgreeWithPolicy      *bool   `json:"agree_with_policy,omitempty" tfsdk:"agree_with_policy"`
	BirthDate            *string `json:"birth_date,omitempty" tfsdk:"birth_date"`
	Description          *string `json:"description,omitempty" tfsdk:"description"`
	FirstName            *string `json:"first_name,omitempty" tfsdk:"first_name"`
	Image                *string `json:"image,omitempty" tfsdk:"image"`
	IsActive             *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	IsStaff              *bool   `json:"is_staff,omitempty" tfsdk:"is_staff"`
	IsSupport            *bool   `json:"is_support,omitempty" tfsdk:"is_support"`
	JobTitle             *string `json:"job_title,omitempty" tfsdk:"job_title"`
	LastName             *string `json:"last_name,omitempty" tfsdk:"last_name"`
	NativeName           *string `json:"native_name,omitempty" tfsdk:"native_name"`
	NotificationsEnabled *bool   `json:"notifications_enabled,omitempty" tfsdk:"notifications_enabled"`
	Organization         *string `json:"organization,omitempty" tfsdk:"organization"`
	PhoneNumber          *string `json:"phone_number,omitempty" tfsdk:"phone_number"`
	PreferredLanguage    *string `json:"preferred_language,omitempty" tfsdk:"preferred_language"`
	Slug                 *string `json:"slug,omitempty" tfsdk:"slug"`
	TokenLifetime        *int64  `json:"token_lifetime,omitempty" tfsdk:"token_lifetime"`
	Username             *string `json:"username,omitempty" tfsdk:"username"`
}

type PatchedUserRequestMultipart struct {
	AgreeWithPolicy      *bool   `json:"agree_with_policy,omitempty" tfsdk:"agree_with_policy"`
	BirthDate            *string `json:"birth_date,omitempty" tfsdk:"birth_date"`
	Description          *string `json:"description,omitempty" tfsdk:"description"`
	FirstName            *string `json:"first_name,omitempty" tfsdk:"first_name"`
	Image                *string `json:"image,omitempty" tfsdk:"image"`
	IsActive             *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	IsStaff              *bool   `json:"is_staff,omitempty" tfsdk:"is_staff"`
	IsSupport            *bool   `json:"is_support,omitempty" tfsdk:"is_support"`
	JobTitle             *string `json:"job_title,omitempty" tfsdk:"job_title"`
	LastName             *string `json:"last_name,omitempty" tfsdk:"last_name"`
	NativeName           *string `json:"native_name,omitempty" tfsdk:"native_name"`
	NotificationsEnabled *bool   `json:"notifications_enabled,omitempty" tfsdk:"notifications_enabled"`
	Organization         *string `json:"organization,omitempty" tfsdk:"organization"`
	PhoneNumber          *string `json:"phone_number,omitempty" tfsdk:"phone_number"`
	PreferredLanguage    *string `json:"preferred_language,omitempty" tfsdk:"preferred_language"`
	Slug                 *string `json:"slug,omitempty" tfsdk:"slug"`
	TokenLifetime        *int64  `json:"token_lifetime,omitempty" tfsdk:"token_lifetime"`
	Username             *string `json:"username,omitempty" tfsdk:"username"`
}

type PatchedVmwareVirtualMachineRequest struct {
	Cores          *int64  `json:"cores,omitempty" tfsdk:"cores"`
	CoresPerSocket *int64  `json:"cores_per_socket,omitempty" tfsdk:"cores_per_socket"`
	Description    *string `json:"description,omitempty" tfsdk:"description"`
	Ram            *int64  `json:"ram,omitempty" tfsdk:"ram"`
}

type PatchedWebHookRequest struct {
	ContentType    *string  `json:"content_type,omitempty" tfsdk:"content_type"`
	DestinationUrl *string  `json:"destination_url,omitempty" tfsdk:"destination_url"`
	EventGroups    []string `json:"event_groups,omitempty" tfsdk:"event_groups"`
	EventTypes     []string `json:"event_types,omitempty" tfsdk:"event_types"`
	IsActive       *bool    `json:"is_active,omitempty" tfsdk:"is_active"`
}

type Payment struct {
	CustomerUuid  *string `json:"customer_uuid" tfsdk:"customer_uuid"`
	DateOfPayment *string `json:"date_of_payment" tfsdk:"date_of_payment"`
	Invoice       *string `json:"invoice" tfsdk:"invoice"`
	InvoicePeriod *string `json:"invoice_period" tfsdk:"invoice_period"`
	InvoiceUuid   *string `json:"invoice_uuid" tfsdk:"invoice_uuid"`
	Profile       *string `json:"profile" tfsdk:"profile"`
	Proof         *string `json:"proof,omitempty" tfsdk:"proof"`
	Sum           *string `json:"sum,omitempty" tfsdk:"sum"`
	Url           *string `json:"url" tfsdk:"url"`
}

type PaymentProfile struct {
	Attributes         *PaymentProfileAttributes `json:"attributes,omitempty" tfsdk:"attributes"`
	IsActive           *bool                     `json:"is_active,omitempty" tfsdk:"is_active"`
	Name               *string                   `json:"name,omitempty" tfsdk:"name"`
	Organization       *string                   `json:"organization,omitempty" tfsdk:"organization"`
	OrganizationUuid   *string                   `json:"organization_uuid,omitempty" tfsdk:"organization_uuid"`
	PaymentType        *string                   `json:"payment_type,omitempty" tfsdk:"payment_type"`
	PaymentTypeDisplay *string                   `json:"payment_type_display,omitempty" tfsdk:"payment_type_display"`
	Url                *string                   `json:"url,omitempty" tfsdk:"url"`
}

type PaymentProfileAttributes struct {
	AgreementNumber *string `json:"agreement_number,omitempty" tfsdk:"agreement_number"`
	ContractSum     *int64  `json:"contract_sum,omitempty" tfsdk:"contract_sum"`
	EndDate         *string `json:"end_date,omitempty" tfsdk:"end_date"`
}

type PaymentProfileAttributesRequest struct {
	AgreementNumber *string `json:"agreement_number,omitempty" tfsdk:"agreement_number"`
	ContractSum     *int64  `json:"contract_sum,omitempty" tfsdk:"contract_sum"`
	EndDate         *string `json:"end_date,omitempty" tfsdk:"end_date"`
}

type PaymentProfileRequest struct {
	Attributes   *PaymentProfileAttributesRequest `json:"attributes,omitempty" tfsdk:"attributes"`
	IsActive     *bool                            `json:"is_active,omitempty" tfsdk:"is_active"`
	Name         *string                          `json:"name" tfsdk:"name"`
	Organization *string                          `json:"organization" tfsdk:"organization"`
	PaymentType  *string                          `json:"payment_type" tfsdk:"payment_type"`
}

type PaymentRequest struct {
	DateOfPayment *string `json:"date_of_payment" tfsdk:"date_of_payment"`
	Profile       *string `json:"profile" tfsdk:"profile"`
	Proof         *string `json:"proof,omitempty" tfsdk:"proof"`
	Sum           *string `json:"sum,omitempty" tfsdk:"sum"`
}

type PaymentRequestForm struct {
	DateOfPayment *string `json:"date_of_payment" tfsdk:"date_of_payment"`
	Profile       *string `json:"profile" tfsdk:"profile"`
	Proof         *string `json:"proof,omitempty" tfsdk:"proof"`
	Sum           *string `json:"sum,omitempty" tfsdk:"sum"`
}

type PaymentRequestMultipart struct {
	DateOfPayment *string `json:"date_of_payment" tfsdk:"date_of_payment"`
	Profile       *string `json:"profile" tfsdk:"profile"`
	Proof         *string `json:"proof,omitempty" tfsdk:"proof"`
	Sum           *string `json:"sum,omitempty" tfsdk:"sum"`
}

type PaymentTypeEnum struct {
}

type PaymentURLRequest struct {
	PaymentUrl *string `json:"payment_url,omitempty" tfsdk:"payment_url"`
}

type PeriodEnum struct {
}

type Permission struct {
	Created           *string `json:"created,omitempty" tfsdk:"created"`
	CreatedByFullName *string `json:"created_by_full_name,omitempty" tfsdk:"created_by_full_name"`
	CreatedByUsername *string `json:"created_by_username,omitempty" tfsdk:"created_by_username"`
	CustomerName      *string `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerUuid      *string `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	ExpirationTime    *string `json:"expiration_time,omitempty" tfsdk:"expiration_time"`
	RoleDescription   *string `json:"role_description,omitempty" tfsdk:"role_description"`
	RoleName          *string `json:"role_name,omitempty" tfsdk:"role_name"`
	RoleUuid          *string `json:"role_uuid,omitempty" tfsdk:"role_uuid"`
	ScopeName         *string `json:"scope_name,omitempty" tfsdk:"scope_name"`
	ScopeType         *string `json:"scope_type,omitempty" tfsdk:"scope_type"`
	ScopeUuid         *string `json:"scope_uuid,omitempty" tfsdk:"scope_uuid"`
	UserName          *string `json:"user_name,omitempty" tfsdk:"user_name"`
	UserSlug          *string `json:"user_slug,omitempty" tfsdk:"user_slug"`
	UserUuid          *string `json:"user_uuid,omitempty" tfsdk:"user_uuid"`
}

type PermissionMetadataResponse struct {
}

type PermissionRequest struct {
	Created             *string `json:"created" tfsdk:"created"`
	CreatedByEmail      *string `json:"created_by_email" tfsdk:"created_by_email"`
	CreatedByFullName   *string `json:"created_by_full_name" tfsdk:"created_by_full_name"`
	CreatedByUsername   *string `json:"created_by_username" tfsdk:"created_by_username"`
	CustomerName        *string `json:"customer_name" tfsdk:"customer_name"`
	CustomerUuid        *string `json:"customer_uuid" tfsdk:"customer_uuid"`
	Invitation          *string `json:"invitation" tfsdk:"invitation"`
	ProjectNameTemplate *string `json:"project_name_template" tfsdk:"project_name_template"`
	ReviewComment       *string `json:"review_comment,omitempty" tfsdk:"review_comment"`
	ReviewedAt          *string `json:"reviewed_at" tfsdk:"reviewed_at"`
	ReviewedByFullName  *string `json:"reviewed_by_full_name" tfsdk:"reviewed_by_full_name"`
	ReviewedByUsername  *string `json:"reviewed_by_username" tfsdk:"reviewed_by_username"`
	RoleDescription     *string `json:"role_description" tfsdk:"role_description"`
	RoleName            *string `json:"role_name" tfsdk:"role_name"`
	ScopeName           *string `json:"scope_name" tfsdk:"scope_name"`
	ScopeUuid           *string `json:"scope_uuid" tfsdk:"scope_uuid"`
	State               *string `json:"state" tfsdk:"state"`
	Url                 *string `json:"url" tfsdk:"url"`
}

type PlanComponent struct {
	Amount            *int64  `json:"amount,omitempty" tfsdk:"amount"`
	BillingType       *string `json:"billing_type" tfsdk:"billing_type"`
	ComponentName     *string `json:"component_name" tfsdk:"component_name"`
	DiscountRate      *int64  `json:"discount_rate,omitempty" tfsdk:"discount_rate"`
	DiscountThreshold *int64  `json:"discount_threshold,omitempty" tfsdk:"discount_threshold"`
	FuturePrice       *string `json:"future_price,omitempty" tfsdk:"future_price"`
	MeasuredUnit      *string `json:"measured_unit" tfsdk:"measured_unit"`
	OfferingName      *string `json:"offering_name" tfsdk:"offering_name"`
	PlanName          *string `json:"plan_name" tfsdk:"plan_name"`
	PlanUnit          *string `json:"plan_unit" tfsdk:"plan_unit"`
	Price             *string `json:"price,omitempty" tfsdk:"price"`
}

type PlanUsageResponse struct {
	CustomerProviderName *string `json:"customer_provider_name" tfsdk:"customer_provider_name"`
	CustomerProviderUuid *string `json:"customer_provider_uuid" tfsdk:"customer_provider_uuid"`
	Limit                *int64  `json:"limit" tfsdk:"limit"`
	OfferingName         *string `json:"offering_name" tfsdk:"offering_name"`
	OfferingUuid         *string `json:"offering_uuid" tfsdk:"offering_uuid"`
	PlanName             *string `json:"plan_name" tfsdk:"plan_name"`
	PlanUuid             *string `json:"plan_uuid" tfsdk:"plan_uuid"`
	Remaining            *int64  `json:"remaining" tfsdk:"remaining"`
	Usage                *int64  `json:"usage" tfsdk:"usage"`
}

type PluginComponent struct {
	BillingType  *string `json:"billing_type" tfsdk:"billing_type"`
	MeasuredUnit *string `json:"measured_unit" tfsdk:"measured_unit"`
	Name         *string `json:"name" tfsdk:"name"`
	Type         *string `json:"type" tfsdk:"type"`
}

type PluginOfferingType struct {
	AvailableLimits []string          `json:"available_limits" tfsdk:"available_limits"`
	Components      []PluginComponent `json:"components" tfsdk:"components"`
	OfferingType    *string           `json:"offering_type" tfsdk:"offering_type"`
}

type PolicyEnum struct {
}

type PolicyTypeEnum struct {
}

type PricesUpdateRequest struct {
}

type Priority struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	IconUrl     *string `json:"icon_url,omitempty" tfsdk:"icon_url"`
	Name        *string `json:"name" tfsdk:"name"`
	Url         *string `json:"url" tfsdk:"url"`
}

type Project struct {
	BackendId                            *string  `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created                              *string  `json:"created,omitempty" tfsdk:"created"`
	Customer                             *string  `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation                 *string  `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerDisplayBillingInfoInProjects *bool    `json:"customer_display_billing_info_in_projects,omitempty" tfsdk:"customer_display_billing_info_in_projects"`
	CustomerName                         *string  `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName                   *string  `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerSlug                         *string  `json:"customer_slug,omitempty" tfsdk:"customer_slug"`
	CustomerUuid                         *string  `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                          *string  `json:"description,omitempty" tfsdk:"description"`
	EndDate                              *string  `json:"end_date,omitempty" tfsdk:"end_date"`
	EndDateRequestedBy                   *string  `json:"end_date_requested_by,omitempty" tfsdk:"end_date_requested_by"`
	GracePeriodDays                      *int64   `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`
	Image                                *string  `json:"image,omitempty" tfsdk:"image"`
	IsIndustry                           *bool    `json:"is_industry,omitempty" tfsdk:"is_industry"`
	IsRemoved                            *bool    `json:"is_removed,omitempty" tfsdk:"is_removed"`
	Kind                                 *string  `json:"kind,omitempty" tfsdk:"kind"`
	MaxServiceAccounts                   *int64   `json:"max_service_accounts,omitempty" tfsdk:"max_service_accounts"`
	Name                                 *string  `json:"name,omitempty" tfsdk:"name"`
	OecdFos2007Code                      *string  `json:"oecd_fos_2007_code,omitempty" tfsdk:"oecd_fos_2007_code"`
	OecdFos2007Label                     *string  `json:"oecd_fos_2007_label,omitempty" tfsdk:"oecd_fos_2007_label"`
	ProjectCredit                        *float64 `json:"project_credit,omitempty" tfsdk:"project_credit"`
	ResourcesCount                       *int64   `json:"resources_count,omitempty" tfsdk:"resources_count"`
	Slug                                 *string  `json:"slug,omitempty" tfsdk:"slug"`
	StaffNotes                           *string  `json:"staff_notes,omitempty" tfsdk:"staff_notes"`
	StartDate                            *string  `json:"start_date,omitempty" tfsdk:"start_date"`
	Type                                 *string  `json:"type,omitempty" tfsdk:"type"`
	TypeName                             *string  `json:"type_name,omitempty" tfsdk:"type_name"`
	TypeUuid                             *string  `json:"type_uuid,omitempty" tfsdk:"type_uuid"`
	Url                                  *string  `json:"url,omitempty" tfsdk:"url"`
}

type ProjectAnswer struct {
	AnswersCount            *int64   `json:"answers_count" tfsdk:"answers_count"`
	CompletionPercentage    *float64 `json:"completion_percentage" tfsdk:"completion_percentage"`
	CompletionUuid          *string  `json:"completion_uuid" tfsdk:"completion_uuid"`
	IsCompleted             *bool    `json:"is_completed" tfsdk:"is_completed"`
	ProjectName             *string  `json:"project_name" tfsdk:"project_name"`
	ProjectUuid             *string  `json:"project_uuid" tfsdk:"project_uuid"`
	RequiresReview          *bool    `json:"requires_review" tfsdk:"requires_review"`
	UnansweredRequiredCount *int64   `json:"unanswered_required_count" tfsdk:"unanswered_required_count"`
}

type ProjectAttachRequest struct {
	ProjectUuid *string `json:"project_uuid" tfsdk:"project_uuid"`
}

type ProjectCredit struct {
	AllocatedCustomerCredit                     *float64               `json:"allocated_customer_credit" tfsdk:"allocated_customer_credit"`
	ApplyAsMinimalConsumption                   *bool                  `json:"apply_as_minimal_consumption,omitempty" tfsdk:"apply_as_minimal_consumption"`
	ConsumptionLastMonth                        *float64               `json:"consumption_last_month" tfsdk:"consumption_last_month"`
	CustomerCredit                              *string                `json:"customer_credit" tfsdk:"customer_credit"`
	CustomerName                                *string                `json:"customer_name" tfsdk:"customer_name"`
	CustomerSlug                                *string                `json:"customer_slug" tfsdk:"customer_slug"`
	CustomerUuid                                *string                `json:"customer_uuid" tfsdk:"customer_uuid"`
	EndDate                                     *string                `json:"end_date,omitempty" tfsdk:"end_date"`
	ExpectedConsumption                         *string                `json:"expected_consumption,omitempty" tfsdk:"expected_consumption"`
	GraceCoefficient                            *string                `json:"grace_coefficient,omitempty" tfsdk:"grace_coefficient"`
	MarkUnusedCreditAsSpentOnProjectTermination *bool                  `json:"mark_unused_credit_as_spent_on_project_termination,omitempty" tfsdk:"mark_unused_credit_as_spent_on_project_termination"`
	MinimalConsumption                          *float64               `json:"minimal_consumption" tfsdk:"minimal_consumption"`
	MinimalConsumptionLogic                     *string                `json:"minimal_consumption_logic,omitempty" tfsdk:"minimal_consumption_logic"`
	Offerings                                   []NestedPublicOffering `json:"offerings" tfsdk:"offerings"`
	Project                                     *string                `json:"project" tfsdk:"project"`
	ProjectName                                 *string                `json:"project_name" tfsdk:"project_name"`
	ProjectSlug                                 *string                `json:"project_slug" tfsdk:"project_slug"`
	ProjectUuid                                 *string                `json:"project_uuid" tfsdk:"project_uuid"`
	Url                                         *string                `json:"url" tfsdk:"url"`
	Value                                       *string                `json:"value,omitempty" tfsdk:"value"`
}

type ProjectCreditRequest struct {
	ApplyAsMinimalConsumption                   *bool   `json:"apply_as_minimal_consumption,omitempty" tfsdk:"apply_as_minimal_consumption"`
	EndDate                                     *string `json:"end_date,omitempty" tfsdk:"end_date"`
	ExpectedConsumption                         *string `json:"expected_consumption,omitempty" tfsdk:"expected_consumption"`
	GraceCoefficient                            *string `json:"grace_coefficient,omitempty" tfsdk:"grace_coefficient"`
	MarkUnusedCreditAsSpentOnProjectTermination *bool   `json:"mark_unused_credit_as_spent_on_project_termination,omitempty" tfsdk:"mark_unused_credit_as_spent_on_project_termination"`
	MinimalConsumptionLogic                     *string `json:"minimal_consumption_logic,omitempty" tfsdk:"minimal_consumption_logic"`
	Project                                     *string `json:"project" tfsdk:"project"`
	Value                                       *string `json:"value,omitempty" tfsdk:"value"`
}

type ProjectDetail struct {
	CompletionPercentage *float64 `json:"completion_percentage" tfsdk:"completion_percentage"`
	CompletionUuid       *string  `json:"completion_uuid" tfsdk:"completion_uuid"`
	IsCompleted          *bool    `json:"is_completed" tfsdk:"is_completed"`
	ProjectName          *string  `json:"project_name" tfsdk:"project_name"`
	ProjectUuid          *string  `json:"project_uuid" tfsdk:"project_uuid"`
	RequiresReview       *bool    `json:"requires_review" tfsdk:"requires_review"`
}

type ProjectDetailsResponse struct {
	FullyCompletedProjects  *int64          `json:"fully_completed_projects" tfsdk:"fully_completed_projects"`
	ProjectDetails          []ProjectDetail `json:"project_details" tfsdk:"project_details"`
	ProjectsRequiringReview *int64          `json:"projects_requiring_review" tfsdk:"projects_requiring_review"`
	ProjectsWithCompletions *int64          `json:"projects_with_completions" tfsdk:"projects_with_completions"`
	TotalProjects           *int64          `json:"total_projects" tfsdk:"total_projects"`
}

type ProjectEstimatedCostPolicy struct {
	Actions           *string  `json:"actions" tfsdk:"actions"`
	Created           *string  `json:"created" tfsdk:"created"`
	CreatedByFullName *string  `json:"created_by_full_name" tfsdk:"created_by_full_name"`
	CreatedByUsername *string  `json:"created_by_username" tfsdk:"created_by_username"`
	CustomerCredit    *float64 `json:"customer_credit" tfsdk:"customer_credit"`
	FiredDatetime     *string  `json:"fired_datetime" tfsdk:"fired_datetime"`
	HasFired          *bool    `json:"has_fired" tfsdk:"has_fired"`
	LimitCost         *int64   `json:"limit_cost" tfsdk:"limit_cost"`
	Period            *int64   `json:"period,omitempty" tfsdk:"period"`
	PeriodName        *string  `json:"period_name" tfsdk:"period_name"`
	ProjectCredit     *float64 `json:"project_credit" tfsdk:"project_credit"`
	Scope             *string  `json:"scope" tfsdk:"scope"`
	ScopeName         *string  `json:"scope_name" tfsdk:"scope_name"`
	ScopeUuid         *string  `json:"scope_uuid" tfsdk:"scope_uuid"`
	Url               *string  `json:"url" tfsdk:"url"`
}

type ProjectEstimatedCostPolicyRequest struct {
	Actions   *string `json:"actions" tfsdk:"actions"`
	LimitCost *int64  `json:"limit_cost" tfsdk:"limit_cost"`
	Period    *int64  `json:"period,omitempty" tfsdk:"period"`
	Scope     *string `json:"scope" tfsdk:"scope"`
}

type ProjectHyperlinkRequest struct {
	Url *string `json:"url" tfsdk:"url"`
}

type ProjectInfo struct {
	AllowedDestinations *string `json:"allowed_destinations,omitempty" tfsdk:"allowed_destinations"`
	Project             *string `json:"project" tfsdk:"project"`
	Shortname           *string `json:"shortname,omitempty" tfsdk:"shortname"`
}

type ProjectInfoRequest struct {
	AllowedDestinations *string `json:"allowed_destinations,omitempty" tfsdk:"allowed_destinations"`
	Project             *string `json:"project" tfsdk:"project"`
	Shortname           *string `json:"shortname,omitempty" tfsdk:"shortname"`
}

type ProjectPermissionLog struct {
	Created           *string `json:"created,omitempty" tfsdk:"created"`
	CreatedBy         *string `json:"created_by,omitempty" tfsdk:"created_by"`
	CreatedByFullName *string `json:"created_by_full_name,omitempty" tfsdk:"created_by_full_name"`
	CreatedByUsername *string `json:"created_by_username,omitempty" tfsdk:"created_by_username"`
	CustomerName      *string `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerUuid      *string `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	ExpirationTime    *string `json:"expiration_time,omitempty" tfsdk:"expiration_time"`
	Project           *string `json:"project,omitempty" tfsdk:"project"`
	ProjectCreated    *string `json:"project_created,omitempty" tfsdk:"project_created"`
	ProjectEndDate    *string `json:"project_end_date,omitempty" tfsdk:"project_end_date"`
	ProjectName       *string `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid       *string `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	Role              *string `json:"role,omitempty" tfsdk:"role"`
	RoleName          *string `json:"role_name,omitempty" tfsdk:"role_name"`
	User              *string `json:"user,omitempty" tfsdk:"user"`
	UserEmail         *string `json:"user_email,omitempty" tfsdk:"user_email"`
	UserFullName      *string `json:"user_full_name,omitempty" tfsdk:"user_full_name"`
	UserNativeName    *string `json:"user_native_name,omitempty" tfsdk:"user_native_name"`
	UserUsername      *string `json:"user_username,omitempty" tfsdk:"user_username"`
	UserUuid          *string `json:"user_uuid,omitempty" tfsdk:"user_uuid"`
}

type ProjectPermissionReview struct {
	Closed           *string `json:"closed" tfsdk:"closed"`
	Created          *string `json:"created" tfsdk:"created"`
	IsPending        *bool   `json:"is_pending" tfsdk:"is_pending"`
	ProjectName      *string `json:"project_name" tfsdk:"project_name"`
	ProjectUuid      *string `json:"project_uuid" tfsdk:"project_uuid"`
	ReviewerFullName *string `json:"reviewer_full_name" tfsdk:"reviewer_full_name"`
	ReviewerUuid     *string `json:"reviewer_uuid" tfsdk:"reviewer_uuid"`
	Url              *string `json:"url" tfsdk:"url"`
}

type ProjectQuotas struct {
	CustomerAbbreviation *string `json:"customer_abbreviation" tfsdk:"customer_abbreviation"`
	CustomerName         *string `json:"customer_name" tfsdk:"customer_name"`
	ProjectName          *string `json:"project_name" tfsdk:"project_name"`
	Value                *int64  `json:"value" tfsdk:"value"`
}

type ProjectRecoveryRequest struct {
	EndDate                          *string `json:"end_date,omitempty" tfsdk:"end_date"`
	RestoreTeamMembers               *bool   `json:"restore_team_members,omitempty" tfsdk:"restore_team_members"`
	SendInvitationsToPreviousMembers *bool   `json:"send_invitations_to_previous_members,omitempty" tfsdk:"send_invitations_to_previous_members"`
}

type ProjectRequest struct {
	BackendId       *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Customer        *string `json:"customer" tfsdk:"customer"`
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	EndDate         *string `json:"end_date,omitempty" tfsdk:"end_date"`
	GracePeriodDays *int64  `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`
	Image           *string `json:"image,omitempty" tfsdk:"image"`
	IsIndustry      *bool   `json:"is_industry,omitempty" tfsdk:"is_industry"`
	Kind            *string `json:"kind,omitempty" tfsdk:"kind"`
	Name            *string `json:"name" tfsdk:"name"`
	OecdFos2007Code *string `json:"oecd_fos_2007_code,omitempty" tfsdk:"oecd_fos_2007_code"`
	Slug            *string `json:"slug,omitempty" tfsdk:"slug"`
	StaffNotes      *string `json:"staff_notes,omitempty" tfsdk:"staff_notes"`
	StartDate       *string `json:"start_date,omitempty" tfsdk:"start_date"`
	Type            *string `json:"type,omitempty" tfsdk:"type"`
}

type ProjectRequestForm struct {
	BackendId       *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Customer        *string `json:"customer" tfsdk:"customer"`
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	EndDate         *string `json:"end_date,omitempty" tfsdk:"end_date"`
	GracePeriodDays *int64  `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`
	Image           *string `json:"image,omitempty" tfsdk:"image"`
	IsIndustry      *bool   `json:"is_industry,omitempty" tfsdk:"is_industry"`
	Kind            *string `json:"kind,omitempty" tfsdk:"kind"`
	Name            *string `json:"name" tfsdk:"name"`
	OecdFos2007Code *string `json:"oecd_fos_2007_code,omitempty" tfsdk:"oecd_fos_2007_code"`
	Slug            *string `json:"slug,omitempty" tfsdk:"slug"`
	StaffNotes      *string `json:"staff_notes,omitempty" tfsdk:"staff_notes"`
	StartDate       *string `json:"start_date,omitempty" tfsdk:"start_date"`
	Type            *string `json:"type,omitempty" tfsdk:"type"`
}

type ProjectRequestMultipart struct {
	BackendId       *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Customer        *string `json:"customer" tfsdk:"customer"`
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	EndDate         *string `json:"end_date,omitempty" tfsdk:"end_date"`
	GracePeriodDays *int64  `json:"grace_period_days,omitempty" tfsdk:"grace_period_days"`
	Image           *string `json:"image,omitempty" tfsdk:"image"`
	IsIndustry      *bool   `json:"is_industry,omitempty" tfsdk:"is_industry"`
	Kind            *string `json:"kind,omitempty" tfsdk:"kind"`
	Name            *string `json:"name" tfsdk:"name"`
	OecdFos2007Code *string `json:"oecd_fos_2007_code,omitempty" tfsdk:"oecd_fos_2007_code"`
	Slug            *string `json:"slug,omitempty" tfsdk:"slug"`
	StaffNotes      *string `json:"staff_notes,omitempty" tfsdk:"staff_notes"`
	StartDate       *string `json:"start_date,omitempty" tfsdk:"start_date"`
	Type            *string `json:"type,omitempty" tfsdk:"type"`
}

type ProjectServiceAccount struct {
	Created              *string `json:"created" tfsdk:"created"`
	CustomerAbbreviation *string `json:"customer_abbreviation" tfsdk:"customer_abbreviation"`
	CustomerName         *string `json:"customer_name" tfsdk:"customer_name"`
	CustomerUuid         *string `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description          *string `json:"description,omitempty" tfsdk:"description"`
	Email                *string `json:"email,omitempty" tfsdk:"email"`
	ErrorMessage         *string `json:"error_message" tfsdk:"error_message"`
	ErrorTraceback       *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	ExpiresAt            *string `json:"expires_at" tfsdk:"expires_at"`
	Modified             *string `json:"modified" tfsdk:"modified"`
	PreferredIdentifier  *string `json:"preferred_identifier,omitempty" tfsdk:"preferred_identifier"`
	Project              *string `json:"project" tfsdk:"project"`
	ProjectName          *string `json:"project_name" tfsdk:"project_name"`
	ProjectUuid          *string `json:"project_uuid" tfsdk:"project_uuid"`
	Token                *string `json:"token" tfsdk:"token"`
	Url                  *string `json:"url" tfsdk:"url"`
	Username             *string `json:"username,omitempty" tfsdk:"username"`
}

type ProjectServiceAccountRequest struct {
	Description         *string `json:"description,omitempty" tfsdk:"description"`
	Email               *string `json:"email,omitempty" tfsdk:"email"`
	ErrorTraceback      *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	PreferredIdentifier *string `json:"preferred_identifier,omitempty" tfsdk:"preferred_identifier"`
	Project             *string `json:"project" tfsdk:"project"`
	Username            *string `json:"username,omitempty" tfsdk:"username"`
}

type ProjectTemplate struct {
	ApprovalLimit  *string                   `json:"approval_limit,omitempty" tfsdk:"approval_limit"`
	Customer       *string                   `json:"customer" tfsdk:"customer"`
	Key            *string                   `json:"key,omitempty" tfsdk:"key"`
	MaxCreditLimit *string                   `json:"max_credit_limit,omitempty" tfsdk:"max_credit_limit"`
	Name           *string                   `json:"name" tfsdk:"name"`
	Offering       *string                   `json:"offering" tfsdk:"offering"`
	Offerings      []string                  `json:"offerings" tfsdk:"offerings"`
	OfferingsData  []ProviderOfferingDetails `json:"offerings_data" tfsdk:"offerings_data"`
	Portal         *string                   `json:"portal" tfsdk:"portal"`
	Provider       *string                   `json:"provider" tfsdk:"provider"`
	Shortname      *string                   `json:"shortname,omitempty" tfsdk:"shortname"`
}

type ProjectTemplateRequest struct {
	ApprovalLimit  *string  `json:"approval_limit,omitempty" tfsdk:"approval_limit"`
	Customer       *string  `json:"customer" tfsdk:"customer"`
	Key            *string  `json:"key,omitempty" tfsdk:"key"`
	MaxCreditLimit *string  `json:"max_credit_limit,omitempty" tfsdk:"max_credit_limit"`
	Name           *string  `json:"name" tfsdk:"name"`
	Offering       *string  `json:"offering" tfsdk:"offering"`
	Offerings      []string `json:"offerings" tfsdk:"offerings"`
	Portal         *string  `json:"portal" tfsdk:"portal"`
	Provider       *string  `json:"provider" tfsdk:"provider"`
	Shortname      *string  `json:"shortname,omitempty" tfsdk:"shortname"`
}

type ProjectType struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
	Url         *string `json:"url" tfsdk:"url"`
}

type ProjectUser struct {
	Email                *string `json:"email,omitempty" tfsdk:"email"`
	ExpirationTime       *string `json:"expiration_time" tfsdk:"expiration_time"`
	FullName             *string `json:"full_name" tfsdk:"full_name"`
	OfferingUserState    *string `json:"offering_user_state" tfsdk:"offering_user_state"`
	OfferingUserUsername *string `json:"offering_user_username" tfsdk:"offering_user_username"`
	Role                 *string `json:"role" tfsdk:"role"`
	Url                  *string `json:"url" tfsdk:"url"`
	Username             *string `json:"username" tfsdk:"username"`
}

type ProjectsLimitsGroupedByIndustryFlag struct {
}

type ProjectsLimitsGroupedByOecd struct {
}

type ProjectsUsagesGroupedByIndustryFlag struct {
}

type ProjectsUsagesGroupedByOecd struct {
}

type Proposal struct {
	AllocationComment            *string                 `json:"allocation_comment" tfsdk:"allocation_comment"`
	ApprovedBy                   *string                 `json:"approved_by" tfsdk:"approved_by"`
	CallManagingOrganisationUuid *string                 `json:"call_managing_organisation_uuid" tfsdk:"call_managing_organisation_uuid"`
	CallName                     *string                 `json:"call_name" tfsdk:"call_name"`
	CallUuid                     *string                 `json:"call_uuid" tfsdk:"call_uuid"`
	Created                      *string                 `json:"created" tfsdk:"created"`
	CreatedBy                    *string                 `json:"created_by" tfsdk:"created_by"`
	CreatedByName                *string                 `json:"created_by_name" tfsdk:"created_by_name"`
	CreatedByUuid                *string                 `json:"created_by_uuid" tfsdk:"created_by_uuid"`
	Description                  *string                 `json:"description,omitempty" tfsdk:"description"`
	DurationInDays               *int64                  `json:"duration_in_days,omitempty" tfsdk:"duration_in_days"`
	Name                         *string                 `json:"name" tfsdk:"name"`
	OecdFos2007Code              *string                 `json:"oecd_fos_2007_code,omitempty" tfsdk:"oecd_fos_2007_code"`
	OecdFos2007Label             *string                 `json:"oecd_fos_2007_label" tfsdk:"oecd_fos_2007_label"`
	Project                      *string                 `json:"project" tfsdk:"project"`
	ProjectHasCivilianPurpose    *bool                   `json:"project_has_civilian_purpose,omitempty" tfsdk:"project_has_civilian_purpose"`
	ProjectIsConfidential        *bool                   `json:"project_is_confidential,omitempty" tfsdk:"project_is_confidential"`
	ProjectName                  *string                 `json:"project_name" tfsdk:"project_name"`
	ProjectSummary               *string                 `json:"project_summary,omitempty" tfsdk:"project_summary"`
	Slug                         *string                 `json:"slug" tfsdk:"slug"`
	State                        *string                 `json:"state" tfsdk:"state"`
	SupportingDocumentation      []ProposalDocumentation `json:"supporting_documentation" tfsdk:"supporting_documentation"`
	Url                          *string                 `json:"url" tfsdk:"url"`
}

type ProposalApproveRequest struct {
	AllocationComment *string `json:"allocation_comment,omitempty" tfsdk:"allocation_comment"`
}

type ProposalChecklistAnswerSubmitResponse struct {
	Completion *ChecklistCompletionReviewer `json:"completion" tfsdk:"completion"`
	Detail     *string                      `json:"detail" tfsdk:"detail"`
}

type ProposalDetachDocumentsRequest struct {
	Documents []string `json:"documents" tfsdk:"documents"`
}

type ProposalDocumentation struct {
	Created  *string `json:"created" tfsdk:"created"`
	File     *string `json:"file,omitempty" tfsdk:"file"`
	FileName *string `json:"file_name" tfsdk:"file_name"`
	FileSize *int64  `json:"file_size" tfsdk:"file_size"`
}

type ProposalDocumentationRequest struct {
	File *string `json:"file,omitempty" tfsdk:"file"`
}

type ProposalDocumentationRequestForm struct {
	File *string `json:"file,omitempty" tfsdk:"file"`
}

type ProposalDocumentationRequestMultipart struct {
	File *string `json:"file,omitempty" tfsdk:"file"`
}

type ProposalProjectRoleMapping struct {
	Call         *string `json:"call" tfsdk:"call"`
	CallName     *string `json:"call_name" tfsdk:"call_name"`
	CallUuid     *string `json:"call_uuid" tfsdk:"call_uuid"`
	ProjectRole  *string `json:"project_role,omitempty" tfsdk:"project_role"`
	ProposalRole *string `json:"proposal_role" tfsdk:"proposal_role"`
	Url          *string `json:"url" tfsdk:"url"`
}

type ProposalProjectRoleMappingRequest struct {
	Call         *string `json:"call" tfsdk:"call"`
	ProjectRole  *string `json:"project_role,omitempty" tfsdk:"project_role"`
	ProposalRole *string `json:"proposal_role" tfsdk:"proposal_role"`
}

type ProposalRequest struct {
	Description               *string `json:"description,omitempty" tfsdk:"description"`
	DurationInDays            *int64  `json:"duration_in_days,omitempty" tfsdk:"duration_in_days"`
	Name                      *string `json:"name" tfsdk:"name"`
	OecdFos2007Code           *string `json:"oecd_fos_2007_code,omitempty" tfsdk:"oecd_fos_2007_code"`
	ProjectHasCivilianPurpose *bool   `json:"project_has_civilian_purpose,omitempty" tfsdk:"project_has_civilian_purpose"`
	ProjectIsConfidential     *bool   `json:"project_is_confidential,omitempty" tfsdk:"project_is_confidential"`
	ProjectSummary            *string `json:"project_summary,omitempty" tfsdk:"project_summary"`
	RoundUuid                 *string `json:"round_uuid" tfsdk:"round_uuid"`
}

type ProposalReview struct {
	AnonymousReviewerName                 *string `json:"anonymous_reviewer_name,omitempty" tfsdk:"anonymous_reviewer_name"`
	CallManagingOrganisationUuid          *string `json:"call_managing_organisation_uuid" tfsdk:"call_managing_organisation_uuid"`
	CallName                              *string `json:"call_name" tfsdk:"call_name"`
	CallSlug                              *string `json:"call_slug" tfsdk:"call_slug"`
	CallUuid                              *string `json:"call_uuid" tfsdk:"call_uuid"`
	CommentProjectDescription             *string `json:"comment_project_description,omitempty" tfsdk:"comment_project_description"`
	CommentProjectDuration                *string `json:"comment_project_duration,omitempty" tfsdk:"comment_project_duration"`
	CommentProjectHasCivilianPurpose      *string `json:"comment_project_has_civilian_purpose,omitempty" tfsdk:"comment_project_has_civilian_purpose"`
	CommentProjectIsConfidential          *string `json:"comment_project_is_confidential,omitempty" tfsdk:"comment_project_is_confidential"`
	CommentProjectSummary                 *string `json:"comment_project_summary,omitempty" tfsdk:"comment_project_summary"`
	CommentProjectSupportingDocumentation *string `json:"comment_project_supporting_documentation,omitempty" tfsdk:"comment_project_supporting_documentation"`
	CommentProjectTitle                   *string `json:"comment_project_title,omitempty" tfsdk:"comment_project_title"`
	CommentResourceRequests               *string `json:"comment_resource_requests,omitempty" tfsdk:"comment_resource_requests"`
	CommentTeam                           *string `json:"comment_team,omitempty" tfsdk:"comment_team"`
	Created                               *string `json:"created" tfsdk:"created"`
	Modified                              *string `json:"modified" tfsdk:"modified"`
	Proposal                              *string `json:"proposal" tfsdk:"proposal"`
	ProposalName                          *string `json:"proposal_name" tfsdk:"proposal_name"`
	ProposalSlug                          *string `json:"proposal_slug" tfsdk:"proposal_slug"`
	ProposalUuid                          *string `json:"proposal_uuid" tfsdk:"proposal_uuid"`
	ReviewEndDate                         *string `json:"review_end_date" tfsdk:"review_end_date"`
	Reviewer                              *string `json:"reviewer,omitempty" tfsdk:"reviewer"`
	ReviewerFullName                      *string `json:"reviewer_full_name,omitempty" tfsdk:"reviewer_full_name"`
	ReviewerUuid                          *string `json:"reviewer_uuid,omitempty" tfsdk:"reviewer_uuid"`
	RoundCutoffTime                       *string `json:"round_cutoff_time" tfsdk:"round_cutoff_time"`
	RoundName                             *string `json:"round_name" tfsdk:"round_name"`
	RoundSlug                             *string `json:"round_slug" tfsdk:"round_slug"`
	RoundStartTime                        *string `json:"round_start_time" tfsdk:"round_start_time"`
	RoundUuid                             *string `json:"round_uuid" tfsdk:"round_uuid"`
	State                                 *string `json:"state" tfsdk:"state"`
	SummaryPrivateComment                 *string `json:"summary_private_comment,omitempty" tfsdk:"summary_private_comment"`
	SummaryPublicComment                  *string `json:"summary_public_comment,omitempty" tfsdk:"summary_public_comment"`
	SummaryScore                          *int64  `json:"summary_score,omitempty" tfsdk:"summary_score"`
	Url                                   *string `json:"url" tfsdk:"url"`
}

type ProposalReviewRequest struct {
	CommentProjectDescription             *string `json:"comment_project_description,omitempty" tfsdk:"comment_project_description"`
	CommentProjectDuration                *string `json:"comment_project_duration,omitempty" tfsdk:"comment_project_duration"`
	CommentProjectHasCivilianPurpose      *string `json:"comment_project_has_civilian_purpose,omitempty" tfsdk:"comment_project_has_civilian_purpose"`
	CommentProjectIsConfidential          *string `json:"comment_project_is_confidential,omitempty" tfsdk:"comment_project_is_confidential"`
	CommentProjectSummary                 *string `json:"comment_project_summary,omitempty" tfsdk:"comment_project_summary"`
	CommentProjectSupportingDocumentation *string `json:"comment_project_supporting_documentation,omitempty" tfsdk:"comment_project_supporting_documentation"`
	CommentProjectTitle                   *string `json:"comment_project_title,omitempty" tfsdk:"comment_project_title"`
	CommentResourceRequests               *string `json:"comment_resource_requests,omitempty" tfsdk:"comment_resource_requests"`
	CommentTeam                           *string `json:"comment_team,omitempty" tfsdk:"comment_team"`
	Proposal                              *string `json:"proposal" tfsdk:"proposal"`
	Reviewer                              *string `json:"reviewer,omitempty" tfsdk:"reviewer"`
	SummaryPrivateComment                 *string `json:"summary_private_comment,omitempty" tfsdk:"summary_private_comment"`
	SummaryPublicComment                  *string `json:"summary_public_comment,omitempty" tfsdk:"summary_public_comment"`
	SummaryScore                          *int64  `json:"summary_score,omitempty" tfsdk:"summary_score"`
}

type ProposalReviewStateEnum struct {
}

type ProposalStates struct {
}

type ProposalUpdateProjectDetailsRequest struct {
	Description               *string `json:"description,omitempty" tfsdk:"description"`
	DurationInDays            *int64  `json:"duration_in_days,omitempty" tfsdk:"duration_in_days"`
	Name                      *string `json:"name" tfsdk:"name"`
	OecdFos2007Code           *string `json:"oecd_fos_2007_code,omitempty" tfsdk:"oecd_fos_2007_code"`
	ProjectHasCivilianPurpose *bool   `json:"project_has_civilian_purpose,omitempty" tfsdk:"project_has_civilian_purpose"`
	ProjectIsConfidential     *bool   `json:"project_is_confidential,omitempty" tfsdk:"project_is_confidential"`
	ProjectSummary            *string `json:"project_summary,omitempty" tfsdk:"project_summary"`
}

type ProtectedCall struct {
	BackendId                           *string                   `json:"backend_id,omitempty" tfsdk:"backend_id"`
	ComplianceChecklist                 *string                   `json:"compliance_checklist,omitempty" tfsdk:"compliance_checklist"`
	ComplianceChecklistName             *string                   `json:"compliance_checklist_name,omitempty" tfsdk:"compliance_checklist_name"`
	Created                             *string                   `json:"created,omitempty" tfsdk:"created"`
	CreatedBy                           *string                   `json:"created_by,omitempty" tfsdk:"created_by"`
	CustomerName                        *string                   `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerUuid                        *string                   `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                         *string                   `json:"description,omitempty" tfsdk:"description"`
	Documents                           []CallDocument            `json:"documents,omitempty" tfsdk:"documents"`
	EndDate                             *string                   `json:"end_date,omitempty" tfsdk:"end_date"`
	ExternalUrl                         *string                   `json:"external_url,omitempty" tfsdk:"external_url"`
	FixedDurationInDays                 *int64                    `json:"fixed_duration_in_days,omitempty" tfsdk:"fixed_duration_in_days"`
	Manager                             *string                   `json:"manager,omitempty" tfsdk:"manager"`
	ManagerUuid                         *string                   `json:"manager_uuid,omitempty" tfsdk:"manager_uuid"`
	Name                                *string                   `json:"name,omitempty" tfsdk:"name"`
	Offerings                           []NestedRequestedOffering `json:"offerings,omitempty" tfsdk:"offerings"`
	ProposalSlugTemplate                *string                   `json:"proposal_slug_template,omitempty" tfsdk:"proposal_slug_template"`
	ReferenceCode                       *string                   `json:"reference_code,omitempty" tfsdk:"reference_code"`
	ResourceTemplates                   []CallResourceTemplate    `json:"resource_templates,omitempty" tfsdk:"resource_templates"`
	ReviewerIdentityVisibleToSubmitters *bool                     `json:"reviewer_identity_visible_to_submitters,omitempty" tfsdk:"reviewer_identity_visible_to_submitters"`
	ReviewsVisibleToSubmitters          *bool                     `json:"reviews_visible_to_submitters,omitempty" tfsdk:"reviews_visible_to_submitters"`
	Rounds                              []NestedRound             `json:"rounds,omitempty" tfsdk:"rounds"`
	Slug                                *string                   `json:"slug,omitempty" tfsdk:"slug"`
	StartDate                           *string                   `json:"start_date,omitempty" tfsdk:"start_date"`
	State                               *string                   `json:"state,omitempty" tfsdk:"state"`
	Url                                 *string                   `json:"url,omitempty" tfsdk:"url"`
}

type ProtectedCallRequest struct {
	BackendId                           *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	ComplianceChecklist                 *string `json:"compliance_checklist,omitempty" tfsdk:"compliance_checklist"`
	CreatedBy                           *string `json:"created_by,omitempty" tfsdk:"created_by"`
	Description                         *string `json:"description,omitempty" tfsdk:"description"`
	ExternalUrl                         *string `json:"external_url,omitempty" tfsdk:"external_url"`
	FixedDurationInDays                 *int64  `json:"fixed_duration_in_days,omitempty" tfsdk:"fixed_duration_in_days"`
	Manager                             *string `json:"manager" tfsdk:"manager"`
	Name                                *string `json:"name" tfsdk:"name"`
	ProposalSlugTemplate                *string `json:"proposal_slug_template,omitempty" tfsdk:"proposal_slug_template"`
	ReferenceCode                       *string `json:"reference_code,omitempty" tfsdk:"reference_code"`
	ReviewerIdentityVisibleToSubmitters *bool   `json:"reviewer_identity_visible_to_submitters,omitempty" tfsdk:"reviewer_identity_visible_to_submitters"`
	ReviewsVisibleToSubmitters          *bool   `json:"reviews_visible_to_submitters,omitempty" tfsdk:"reviews_visible_to_submitters"`
	Slug                                *string `json:"slug,omitempty" tfsdk:"slug"`
}

type ProtectedProposalList struct {
	ApprovedByName *string `json:"approved_by_name" tfsdk:"approved_by_name"`
	Created        *string `json:"created" tfsdk:"created"`
	CreatedByName  *string `json:"created_by_name" tfsdk:"created_by_name"`
	Name           *string `json:"name" tfsdk:"name"`
	Slug           *string `json:"slug" tfsdk:"slug"`
	State          *string `json:"state" tfsdk:"state"`
}

type ProtectedProposalListRequest struct {
	Name *string `json:"name" tfsdk:"name"`
	Slug *string `json:"slug" tfsdk:"slug"`
}

type ProtectedRound struct {
	AllocationDate           *string                 `json:"allocation_date,omitempty" tfsdk:"allocation_date"`
	AllocationTime           *string                 `json:"allocation_time,omitempty" tfsdk:"allocation_time"`
	CutoffTime               *string                 `json:"cutoff_time" tfsdk:"cutoff_time"`
	DecidingEntity           *string                 `json:"deciding_entity,omitempty" tfsdk:"deciding_entity"`
	MinimalAverageScoring    *string                 `json:"minimal_average_scoring,omitempty" tfsdk:"minimal_average_scoring"`
	MinimumNumberOfReviewers *int64                  `json:"minimum_number_of_reviewers,omitempty" tfsdk:"minimum_number_of_reviewers"`
	Name                     *string                 `json:"name" tfsdk:"name"`
	Proposals                []ProtectedProposalList `json:"proposals" tfsdk:"proposals"`
	ReviewDurationInDays     *int64                  `json:"review_duration_in_days,omitempty" tfsdk:"review_duration_in_days"`
	ReviewStrategy           *string                 `json:"review_strategy,omitempty" tfsdk:"review_strategy"`
	Slug                     *string                 `json:"slug" tfsdk:"slug"`
	StartTime                *string                 `json:"start_time" tfsdk:"start_time"`
	Status                   *string                 `json:"status" tfsdk:"status"`
	Url                      *string                 `json:"url" tfsdk:"url"`
}

type ProtectedRoundRequest struct {
	AllocationDate           *string `json:"allocation_date,omitempty" tfsdk:"allocation_date"`
	AllocationTime           *string `json:"allocation_time,omitempty" tfsdk:"allocation_time"`
	CutoffTime               *string `json:"cutoff_time" tfsdk:"cutoff_time"`
	DecidingEntity           *string `json:"deciding_entity,omitempty" tfsdk:"deciding_entity"`
	MinimalAverageScoring    *string `json:"minimal_average_scoring,omitempty" tfsdk:"minimal_average_scoring"`
	MinimumNumberOfReviewers *int64  `json:"minimum_number_of_reviewers,omitempty" tfsdk:"minimum_number_of_reviewers"`
	ReviewDurationInDays     *int64  `json:"review_duration_in_days,omitempty" tfsdk:"review_duration_in_days"`
	ReviewStrategy           *string `json:"review_strategy,omitempty" tfsdk:"review_strategy"`
	StartTime                *string `json:"start_time" tfsdk:"start_time"`
}

type ProtocolEnum struct {
}

type ProviderOffering struct {
	CategoryTitle  *string             `json:"category_title,omitempty" tfsdk:"category_title"`
	Components     []OfferingComponent `json:"components,omitempty" tfsdk:"components"`
	CustomerUuid   *string             `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Name           *string             `json:"name,omitempty" tfsdk:"name"`
	Plans          []BaseProviderPlan  `json:"plans,omitempty" tfsdk:"plans"`
	ResourcesCount *int64              `json:"resources_count,omitempty" tfsdk:"resources_count"`
	Slug           *string             `json:"slug,omitempty" tfsdk:"slug"`
	State          *string             `json:"state,omitempty" tfsdk:"state"`
	Thumbnail      *string             `json:"thumbnail,omitempty" tfsdk:"thumbnail"`
	Type           *string             `json:"type,omitempty" tfsdk:"type"`
}

type ProviderOfferingCosts struct {
	Period *string  `json:"period" tfsdk:"period"`
	Price  *float64 `json:"price" tfsdk:"price"`
	Tax    *float64 `json:"tax" tfsdk:"tax"`
	Total  *float64 `json:"total" tfsdk:"total"`
}

type ProviderOfferingCustomer struct {
	Abbreviation *string `json:"abbreviation,omitempty" tfsdk:"abbreviation"`
	Email        *string `json:"email,omitempty" tfsdk:"email"`
	Name         *string `json:"name,omitempty" tfsdk:"name"`
	PhoneNumber  *string `json:"phone_number,omitempty" tfsdk:"phone_number"`
	Slug         *string `json:"slug,omitempty" tfsdk:"slug"`
}

type ProviderOfferingDetails struct {
	AccessUrl                 *string                 `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                 *string                 `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Billable                  *bool                   `json:"billable,omitempty" tfsdk:"billable"`
	BillingTypeClassification *string                 `json:"billing_type_classification,omitempty" tfsdk:"billing_type_classification"`
	Category                  *string                 `json:"category,omitempty" tfsdk:"category"`
	CategoryTitle             *string                 `json:"category_title,omitempty" tfsdk:"category_title"`
	CategoryUuid              *string                 `json:"category_uuid,omitempty" tfsdk:"category_uuid"`
	CitationCount             *int64                  `json:"citation_count,omitempty" tfsdk:"citation_count"`
	ComplianceChecklist       *string                 `json:"compliance_checklist,omitempty" tfsdk:"compliance_checklist"`
	Components                []OfferingComponent     `json:"components,omitempty" tfsdk:"components"`
	Country                   *string                 `json:"country,omitempty" tfsdk:"country"`
	Created                   *string                 `json:"created,omitempty" tfsdk:"created"`
	Customer                  *string                 `json:"customer,omitempty" tfsdk:"customer"`
	CustomerName              *string                 `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerUuid              *string                 `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	DataciteDoi               *string                 `json:"datacite_doi,omitempty" tfsdk:"datacite_doi"`
	Description               *string                 `json:"description,omitempty" tfsdk:"description"`
	Endpoints                 []NestedEndpoint        `json:"endpoints,omitempty" tfsdk:"endpoints"`
	Files                     []NestedOfferingFile    `json:"files,omitempty" tfsdk:"files"`
	FullDescription           *string                 `json:"full_description,omitempty" tfsdk:"full_description"`
	GettingStarted            *string                 `json:"getting_started,omitempty" tfsdk:"getting_started"`
	GoogleCalendarIsPublic    *bool                   `json:"google_calendar_is_public,omitempty" tfsdk:"google_calendar_is_public"`
	GoogleCalendarLink        *string                 `json:"google_calendar_link,omitempty" tfsdk:"google_calendar_link"`
	HasComplianceRequirements *bool                   `json:"has_compliance_requirements,omitempty" tfsdk:"has_compliance_requirements"`
	Image                     *string                 `json:"image,omitempty" tfsdk:"image"`
	IntegrationGuide          *string                 `json:"integration_guide,omitempty" tfsdk:"integration_guide"`
	IntegrationStatus         []IntegrationStatus     `json:"integration_status,omitempty" tfsdk:"integration_status"`
	Latitude                  *float64                `json:"latitude,omitempty" tfsdk:"latitude"`
	Longitude                 *float64                `json:"longitude,omitempty" tfsdk:"longitude"`
	Name                      *string                 `json:"name,omitempty" tfsdk:"name"`
	OrderCount                *int64                  `json:"order_count,omitempty" tfsdk:"order_count"`
	OrganizationGroups        []OrganizationGroup     `json:"organization_groups,omitempty" tfsdk:"organization_groups"`
	ParentDescription         *string                 `json:"parent_description,omitempty" tfsdk:"parent_description"`
	ParentName                *string                 `json:"parent_name,omitempty" tfsdk:"parent_name"`
	ParentUuid                *string                 `json:"parent_uuid,omitempty" tfsdk:"parent_uuid"`
	Partitions                []NestedPartition       `json:"partitions,omitempty" tfsdk:"partitions"`
	PausedReason              *string                 `json:"paused_reason,omitempty" tfsdk:"paused_reason"`
	Plans                     []BaseProviderPlan      `json:"plans,omitempty" tfsdk:"plans"`
	PrivacyPolicyLink         *string                 `json:"privacy_policy_link,omitempty" tfsdk:"privacy_policy_link"`
	Project                   *string                 `json:"project,omitempty" tfsdk:"project"`
	ProjectName               *string                 `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid               *string                 `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	Quotas                    []Quota                 `json:"quotas,omitempty" tfsdk:"quotas"`
	Roles                     []NestedRole            `json:"roles,omitempty" tfsdk:"roles"`
	Scope                     *string                 `json:"scope,omitempty" tfsdk:"scope"`
	ScopeErrorMessage         *string                 `json:"scope_error_message,omitempty" tfsdk:"scope_error_message"`
	ScopeName                 *string                 `json:"scope_name,omitempty" tfsdk:"scope_name"`
	ScopeState                *string                 `json:"scope_state,omitempty" tfsdk:"scope_state"`
	ScopeUuid                 *string                 `json:"scope_uuid,omitempty" tfsdk:"scope_uuid"`
	Screenshots               []NestedScreenshot      `json:"screenshots,omitempty" tfsdk:"screenshots"`
	Shared                    *bool                   `json:"shared,omitempty" tfsdk:"shared"`
	Slug                      *string                 `json:"slug,omitempty" tfsdk:"slug"`
	SoftwareCatalogs          []NestedSoftwareCatalog `json:"software_catalogs,omitempty" tfsdk:"software_catalogs"`
	State                     *string                 `json:"state,omitempty" tfsdk:"state"`
	Thumbnail                 *string                 `json:"thumbnail,omitempty" tfsdk:"thumbnail"`
	TotalCost                 *int64                  `json:"total_cost,omitempty" tfsdk:"total_cost"`
	TotalCostEstimated        *int64                  `json:"total_cost_estimated,omitempty" tfsdk:"total_cost_estimated"`
	TotalCustomers            *int64                  `json:"total_customers,omitempty" tfsdk:"total_customers"`
	Type                      *string                 `json:"type,omitempty" tfsdk:"type"`
	Url                       *string                 `json:"url,omitempty" tfsdk:"url"`
	VendorDetails             *string                 `json:"vendor_details,omitempty" tfsdk:"vendor_details"`
}

type ProviderOfferingDetailsRequest struct {
	AccessUrl           *string                    `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId           *string                    `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Billable            *bool                      `json:"billable,omitempty" tfsdk:"billable"`
	Category            *string                    `json:"category" tfsdk:"category"`
	ComplianceChecklist *string                    `json:"compliance_checklist,omitempty" tfsdk:"compliance_checklist"`
	Components          []OfferingComponentRequest `json:"components,omitempty" tfsdk:"components"`
	Country             *string                    `json:"country,omitempty" tfsdk:"country"`
	Customer            *string                    `json:"customer,omitempty" tfsdk:"customer"`
	DataciteDoi         *string                    `json:"datacite_doi,omitempty" tfsdk:"datacite_doi"`
	Description         *string                    `json:"description,omitempty" tfsdk:"description"`
	FullDescription     *string                    `json:"full_description,omitempty" tfsdk:"full_description"`
	GettingStarted      *string                    `json:"getting_started,omitempty" tfsdk:"getting_started"`
	Image               *string                    `json:"image,omitempty" tfsdk:"image"`
	IntegrationGuide    *string                    `json:"integration_guide,omitempty" tfsdk:"integration_guide"`
	Latitude            *float64                   `json:"latitude,omitempty" tfsdk:"latitude"`
	Longitude           *float64                   `json:"longitude,omitempty" tfsdk:"longitude"`
	Name                *string                    `json:"name" tfsdk:"name"`
	Plans               []BaseProviderPlanRequest  `json:"plans,omitempty" tfsdk:"plans"`
	PrivacyPolicyLink   *string                    `json:"privacy_policy_link,omitempty" tfsdk:"privacy_policy_link"`
	Shared              *bool                      `json:"shared,omitempty" tfsdk:"shared"`
	Slug                *string                    `json:"slug,omitempty" tfsdk:"slug"`
	Thumbnail           *string                    `json:"thumbnail,omitempty" tfsdk:"thumbnail"`
	Type                *string                    `json:"type" tfsdk:"type"`
	VendorDetails       *string                    `json:"vendor_details,omitempty" tfsdk:"vendor_details"`
}

type ProviderPlanDetails struct {
	Archived           *bool                 `json:"archived,omitempty" tfsdk:"archived"`
	ArticleCode        *string               `json:"article_code,omitempty" tfsdk:"article_code"`
	BackendId          *string               `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Components         []NestedPlanComponent `json:"components" tfsdk:"components"`
	Description        *string               `json:"description,omitempty" tfsdk:"description"`
	InitPrice          *float64              `json:"init_price" tfsdk:"init_price"`
	IsActive           *bool                 `json:"is_active" tfsdk:"is_active"`
	MaxAmount          *int64                `json:"max_amount,omitempty" tfsdk:"max_amount"`
	MinimalPrice       *float64              `json:"minimal_price" tfsdk:"minimal_price"`
	Name               *string               `json:"name" tfsdk:"name"`
	Offering           *string               `json:"offering" tfsdk:"offering"`
	OrganizationGroups []OrganizationGroup   `json:"organization_groups" tfsdk:"organization_groups"`
	PlanType           *string               `json:"plan_type" tfsdk:"plan_type"`
	ResourcesCount     *int64                `json:"resources_count" tfsdk:"resources_count"`
	SwitchPrice        *float64              `json:"switch_price" tfsdk:"switch_price"`
	Unit               *string               `json:"unit,omitempty" tfsdk:"unit"`
	UnitPrice          *string               `json:"unit_price,omitempty" tfsdk:"unit_price"`
	Url                *string               `json:"url" tfsdk:"url"`
}

type ProviderPlanDetailsRequest struct {
	Archived    *bool   `json:"archived,omitempty" tfsdk:"archived"`
	ArticleCode *string `json:"article_code,omitempty" tfsdk:"article_code"`
	BackendId   *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	MaxAmount   *int64  `json:"max_amount,omitempty" tfsdk:"max_amount"`
	Name        *string `json:"name" tfsdk:"name"`
	Offering    *string `json:"offering" tfsdk:"offering"`
	Unit        *string `json:"unit,omitempty" tfsdk:"unit"`
	UnitPrice   *string `json:"unit_price,omitempty" tfsdk:"unit_price"`
}

type ProviderProject struct {
	Image *string `json:"image,omitempty" tfsdk:"image"`
	Name  *string `json:"name,omitempty" tfsdk:"name"`
}

type ProviderRequestedOffering struct {
	Call                     *string             `json:"call" tfsdk:"call"`
	CallManagingOrganisation *string             `json:"call_managing_organisation" tfsdk:"call_managing_organisation"`
	CallName                 *string             `json:"call_name" tfsdk:"call_name"`
	CategoryName             *string             `json:"category_name" tfsdk:"category_name"`
	CategoryUuid             *string             `json:"category_uuid" tfsdk:"category_uuid"`
	Components               []OfferingComponent `json:"components" tfsdk:"components"`
	Created                  *string             `json:"created" tfsdk:"created"`
	CreatedByEmail           *string             `json:"created_by_email" tfsdk:"created_by_email"`
	CreatedByName            *string             `json:"created_by_name" tfsdk:"created_by_name"`
	Description              *string             `json:"description" tfsdk:"description"`
	Offering                 *string             `json:"offering" tfsdk:"offering"`
	OfferingName             *string             `json:"offering_name" tfsdk:"offering_name"`
	OfferingUuid             *string             `json:"offering_uuid" tfsdk:"offering_uuid"`
	Plan                     *string             `json:"plan" tfsdk:"plan"`
	ProviderName             *string             `json:"provider_name" tfsdk:"provider_name"`
	State                    *string             `json:"state" tfsdk:"state"`
	Url                      *string             `json:"url" tfsdk:"url"`
}

type ProviderRequestedResource struct {
	CallResourceTemplate     *string `json:"call_resource_template" tfsdk:"call_resource_template"`
	CallResourceTemplateName *string `json:"call_resource_template_name" tfsdk:"call_resource_template_name"`
	CreatedBy                *string `json:"created_by,omitempty" tfsdk:"created_by"`
	CreatedByName            *string `json:"created_by_name" tfsdk:"created_by_name"`
	Description              *string `json:"description,omitempty" tfsdk:"description"`
	Proposal                 *string `json:"proposal" tfsdk:"proposal"`
	ProposalName             *string `json:"proposal_name" tfsdk:"proposal_name"`
	Resource                 *string `json:"resource,omitempty" tfsdk:"resource"`
	ResourceName             *string `json:"resource_name" tfsdk:"resource_name"`
	Url                      *string `json:"url" tfsdk:"url"`
}

type ProviderUser struct {
	Email    *string `json:"email,omitempty" tfsdk:"email"`
	FullName *string `json:"full_name,omitempty" tfsdk:"full_name"`
	Image    *string `json:"image,omitempty" tfsdk:"image"`
}

type PublicCall struct {
	BackendId                           *string                   `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created                             *string                   `json:"created,omitempty" tfsdk:"created"`
	CustomerName                        *string                   `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerUuid                        *string                   `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                         *string                   `json:"description,omitempty" tfsdk:"description"`
	Documents                           []CallDocument            `json:"documents,omitempty" tfsdk:"documents"`
	EndDate                             *string                   `json:"end_date,omitempty" tfsdk:"end_date"`
	ExternalUrl                         *string                   `json:"external_url,omitempty" tfsdk:"external_url"`
	FixedDurationInDays                 *int64                    `json:"fixed_duration_in_days,omitempty" tfsdk:"fixed_duration_in_days"`
	Manager                             *string                   `json:"manager,omitempty" tfsdk:"manager"`
	ManagerUuid                         *string                   `json:"manager_uuid,omitempty" tfsdk:"manager_uuid"`
	Name                                *string                   `json:"name,omitempty" tfsdk:"name"`
	Offerings                           []NestedRequestedOffering `json:"offerings,omitempty" tfsdk:"offerings"`
	ResourceTemplates                   []CallResourceTemplate    `json:"resource_templates,omitempty" tfsdk:"resource_templates"`
	ReviewerIdentityVisibleToSubmitters *bool                     `json:"reviewer_identity_visible_to_submitters,omitempty" tfsdk:"reviewer_identity_visible_to_submitters"`
	ReviewsVisibleToSubmitters          *bool                     `json:"reviews_visible_to_submitters,omitempty" tfsdk:"reviews_visible_to_submitters"`
	Rounds                              []NestedRound             `json:"rounds,omitempty" tfsdk:"rounds"`
	Slug                                *string                   `json:"slug,omitempty" tfsdk:"slug"`
	StartDate                           *string                   `json:"start_date,omitempty" tfsdk:"start_date"`
	State                               *string                   `json:"state,omitempty" tfsdk:"state"`
	Url                                 *string                   `json:"url,omitempty" tfsdk:"url"`
}

type PublicMaintenanceAnnouncement struct {
	ActualEnd              *string                           `json:"actual_end" tfsdk:"actual_end"`
	ActualStart            *string                           `json:"actual_start" tfsdk:"actual_start"`
	AffectedOfferings      []MaintenanceAnnouncementOffering `json:"affected_offerings" tfsdk:"affected_offerings"`
	ExternalReferenceUrl   *string                           `json:"external_reference_url" tfsdk:"external_reference_url"`
	MaintenanceType        *int64                            `json:"maintenance_type" tfsdk:"maintenance_type"`
	MaintenanceTypeDisplay *string                           `json:"maintenance_type_display" tfsdk:"maintenance_type_display"`
	Message                *string                           `json:"message" tfsdk:"message"`
	Name                   *string                           `json:"name" tfsdk:"name"`
	ScheduledEnd           *string                           `json:"scheduled_end" tfsdk:"scheduled_end"`
	ScheduledStart         *string                           `json:"scheduled_start" tfsdk:"scheduled_start"`
	ServiceProviderName    *string                           `json:"service_provider_name" tfsdk:"service_provider_name"`
	State                  *string                           `json:"state" tfsdk:"state"`
	Url                    *string                           `json:"url" tfsdk:"url"`
}

type PublicMaintenanceAnnouncementStateEnum struct {
}

type PublicOfferingDetails struct {
	AccessUrl                 *string                 `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                 *string                 `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Billable                  *bool                   `json:"billable,omitempty" tfsdk:"billable"`
	BillingTypeClassification *string                 `json:"billing_type_classification,omitempty" tfsdk:"billing_type_classification"`
	Category                  *string                 `json:"category,omitempty" tfsdk:"category"`
	CategoryTitle             *string                 `json:"category_title,omitempty" tfsdk:"category_title"`
	CategoryUuid              *string                 `json:"category_uuid,omitempty" tfsdk:"category_uuid"`
	CitationCount             *int64                  `json:"citation_count,omitempty" tfsdk:"citation_count"`
	ComplianceChecklist       *string                 `json:"compliance_checklist,omitempty" tfsdk:"compliance_checklist"`
	Components                []OfferingComponent     `json:"components,omitempty" tfsdk:"components"`
	Country                   *string                 `json:"country,omitempty" tfsdk:"country"`
	Created                   *string                 `json:"created,omitempty" tfsdk:"created"`
	Customer                  *string                 `json:"customer,omitempty" tfsdk:"customer"`
	CustomerName              *string                 `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerUuid              *string                 `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	DataciteDoi               *string                 `json:"datacite_doi,omitempty" tfsdk:"datacite_doi"`
	Description               *string                 `json:"description,omitempty" tfsdk:"description"`
	Endpoints                 []NestedEndpoint        `json:"endpoints,omitempty" tfsdk:"endpoints"`
	Files                     []NestedOfferingFile    `json:"files,omitempty" tfsdk:"files"`
	FullDescription           *string                 `json:"full_description,omitempty" tfsdk:"full_description"`
	GettingStarted            *string                 `json:"getting_started,omitempty" tfsdk:"getting_started"`
	GoogleCalendarIsPublic    *bool                   `json:"google_calendar_is_public,omitempty" tfsdk:"google_calendar_is_public"`
	GoogleCalendarLink        *string                 `json:"google_calendar_link,omitempty" tfsdk:"google_calendar_link"`
	HasComplianceRequirements *bool                   `json:"has_compliance_requirements,omitempty" tfsdk:"has_compliance_requirements"`
	Image                     *string                 `json:"image,omitempty" tfsdk:"image"`
	IntegrationGuide          *string                 `json:"integration_guide,omitempty" tfsdk:"integration_guide"`
	Latitude                  *float64                `json:"latitude,omitempty" tfsdk:"latitude"`
	Longitude                 *float64                `json:"longitude,omitempty" tfsdk:"longitude"`
	Name                      *string                 `json:"name,omitempty" tfsdk:"name"`
	OrderCount                *int64                  `json:"order_count,omitempty" tfsdk:"order_count"`
	OrganizationGroups        []OrganizationGroup     `json:"organization_groups,omitempty" tfsdk:"organization_groups"`
	ParentDescription         *string                 `json:"parent_description,omitempty" tfsdk:"parent_description"`
	ParentName                *string                 `json:"parent_name,omitempty" tfsdk:"parent_name"`
	ParentUuid                *string                 `json:"parent_uuid,omitempty" tfsdk:"parent_uuid"`
	Partitions                []NestedPartition       `json:"partitions,omitempty" tfsdk:"partitions"`
	PausedReason              *string                 `json:"paused_reason,omitempty" tfsdk:"paused_reason"`
	Plans                     []BasePublicPlan        `json:"plans,omitempty" tfsdk:"plans"`
	PrivacyPolicyLink         *string                 `json:"privacy_policy_link,omitempty" tfsdk:"privacy_policy_link"`
	Project                   *string                 `json:"project,omitempty" tfsdk:"project"`
	ProjectName               *string                 `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid               *string                 `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	PromotionCampaigns        []NestedCampaign        `json:"promotion_campaigns,omitempty" tfsdk:"promotion_campaigns"`
	Quotas                    []Quota                 `json:"quotas,omitempty" tfsdk:"quotas"`
	Roles                     []NestedRole            `json:"roles,omitempty" tfsdk:"roles"`
	Scope                     *string                 `json:"scope,omitempty" tfsdk:"scope"`
	ScopeErrorMessage         *string                 `json:"scope_error_message,omitempty" tfsdk:"scope_error_message"`
	ScopeName                 *string                 `json:"scope_name,omitempty" tfsdk:"scope_name"`
	ScopeState                *string                 `json:"scope_state,omitempty" tfsdk:"scope_state"`
	ScopeUuid                 *string                 `json:"scope_uuid,omitempty" tfsdk:"scope_uuid"`
	Screenshots               []NestedScreenshot      `json:"screenshots,omitempty" tfsdk:"screenshots"`
	Shared                    *bool                   `json:"shared,omitempty" tfsdk:"shared"`
	Slug                      *string                 `json:"slug,omitempty" tfsdk:"slug"`
	SoftwareCatalogs          []NestedSoftwareCatalog `json:"software_catalogs,omitempty" tfsdk:"software_catalogs"`
	State                     *string                 `json:"state,omitempty" tfsdk:"state"`
	Thumbnail                 *string                 `json:"thumbnail,omitempty" tfsdk:"thumbnail"`
	TotalCost                 *int64                  `json:"total_cost,omitempty" tfsdk:"total_cost"`
	TotalCostEstimated        *int64                  `json:"total_cost_estimated,omitempty" tfsdk:"total_cost_estimated"`
	TotalCustomers            *int64                  `json:"total_customers,omitempty" tfsdk:"total_customers"`
	Type                      *string                 `json:"type,omitempty" tfsdk:"type"`
	Url                       *string                 `json:"url,omitempty" tfsdk:"url"`
	UserHasConsent            *bool                   `json:"user_has_consent,omitempty" tfsdk:"user_has_consent"`
	VendorDetails             *string                 `json:"vendor_details,omitempty" tfsdk:"vendor_details"`
}

type PullMarketplaceScriptResourceRequest struct {
	ResourceUuid *string `json:"resource_uuid" tfsdk:"resource_uuid"`
}

type QosStrategyEnum struct {
}

type QueryRequest struct {
	Query *string `json:"query" tfsdk:"query"`
}

type Question struct {
	AlwaysRequiresReview    *bool             `json:"always_requires_review,omitempty" tfsdk:"always_requires_review"`
	AlwaysShowGuidance      *bool             `json:"always_show_guidance,omitempty" tfsdk:"always_show_guidance"`
	DependencyLogicOperator *string           `json:"dependency_logic_operator,omitempty" tfsdk:"dependency_logic_operator"`
	Description             *string           `json:"description,omitempty" tfsdk:"description"`
	GuidanceOperator        *string           `json:"guidance_operator,omitempty" tfsdk:"guidance_operator"`
	MaxFileSizeMb           *int64            `json:"max_file_size_mb,omitempty" tfsdk:"max_file_size_mb"`
	MaxFilesCount           *int64            `json:"max_files_count,omitempty" tfsdk:"max_files_count"`
	MaxValue                *string           `json:"max_value,omitempty" tfsdk:"max_value"`
	MinValue                *string           `json:"min_value,omitempty" tfsdk:"min_value"`
	Operator                *string           `json:"operator,omitempty" tfsdk:"operator"`
	Order                   *int64            `json:"order,omitempty" tfsdk:"order"`
	QuestionOptions         []QuestionOptions `json:"question_options" tfsdk:"question_options"`
	QuestionType            *string           `json:"question_type,omitempty" tfsdk:"question_type"`
	Required                *bool             `json:"required,omitempty" tfsdk:"required"`
	UserGuidance            *string           `json:"user_guidance,omitempty" tfsdk:"user_guidance"`
}

type QuestionAdmin struct {
	AlwaysRequiresReview    *bool                  `json:"always_requires_review,omitempty" tfsdk:"always_requires_review"`
	AlwaysShowGuidance      *bool                  `json:"always_show_guidance,omitempty" tfsdk:"always_show_guidance"`
	Checklist               *string                `json:"checklist" tfsdk:"checklist"`
	ChecklistName           *string                `json:"checklist_name" tfsdk:"checklist_name"`
	ChecklistUuid           *string                `json:"checklist_uuid" tfsdk:"checklist_uuid"`
	DependencyLogicOperator *string                `json:"dependency_logic_operator,omitempty" tfsdk:"dependency_logic_operator"`
	Description             *string                `json:"description,omitempty" tfsdk:"description"`
	GuidanceOperator        *string                `json:"guidance_operator,omitempty" tfsdk:"guidance_operator"`
	MaxFileSizeMb           *int64                 `json:"max_file_size_mb,omitempty" tfsdk:"max_file_size_mb"`
	MaxFilesCount           *int64                 `json:"max_files_count,omitempty" tfsdk:"max_files_count"`
	MaxValue                *string                `json:"max_value,omitempty" tfsdk:"max_value"`
	MinValue                *string                `json:"min_value,omitempty" tfsdk:"min_value"`
	Operator                *string                `json:"operator,omitempty" tfsdk:"operator"`
	Order                   *int64                 `json:"order,omitempty" tfsdk:"order"`
	QuestionOptions         []QuestionOptionsAdmin `json:"question_options" tfsdk:"question_options"`
	QuestionType            *string                `json:"question_type,omitempty" tfsdk:"question_type"`
	Required                *bool                  `json:"required,omitempty" tfsdk:"required"`
	Url                     *string                `json:"url" tfsdk:"url"`
	UserGuidance            *string                `json:"user_guidance,omitempty" tfsdk:"user_guidance"`
}

type QuestionAdminRequest struct {
	AlwaysRequiresReview    *bool   `json:"always_requires_review,omitempty" tfsdk:"always_requires_review"`
	AlwaysShowGuidance      *bool   `json:"always_show_guidance,omitempty" tfsdk:"always_show_guidance"`
	Checklist               *string `json:"checklist" tfsdk:"checklist"`
	DependencyLogicOperator *string `json:"dependency_logic_operator,omitempty" tfsdk:"dependency_logic_operator"`
	Description             *string `json:"description,omitempty" tfsdk:"description"`
	GuidanceOperator        *string `json:"guidance_operator,omitempty" tfsdk:"guidance_operator"`
	MaxFileSizeMb           *int64  `json:"max_file_size_mb,omitempty" tfsdk:"max_file_size_mb"`
	MaxFilesCount           *int64  `json:"max_files_count,omitempty" tfsdk:"max_files_count"`
	MaxValue                *string `json:"max_value,omitempty" tfsdk:"max_value"`
	MinValue                *string `json:"min_value,omitempty" tfsdk:"min_value"`
	Operator                *string `json:"operator,omitempty" tfsdk:"operator"`
	Order                   *int64  `json:"order,omitempty" tfsdk:"order"`
	QuestionType            *string `json:"question_type,omitempty" tfsdk:"question_type"`
	Required                *bool   `json:"required,omitempty" tfsdk:"required"`
	UserGuidance            *string `json:"user_guidance,omitempty" tfsdk:"user_guidance"`
}

type QuestionAnswer struct {
	AnsweredProjectsCount *int64  `json:"answered_projects_count" tfsdk:"answered_projects_count"`
	MaxValue              *string `json:"max_value" tfsdk:"max_value"`
	MinValue              *string `json:"min_value" tfsdk:"min_value"`
	Order                 *int64  `json:"order" tfsdk:"order"`
	QuestionDescription   *string `json:"question_description" tfsdk:"question_description"`
	QuestionType          *string `json:"question_type" tfsdk:"question_type"`
	QuestionUuid          *string `json:"question_uuid" tfsdk:"question_uuid"`
	Required              *bool   `json:"required" tfsdk:"required"`
	TotalProjects         *int64  `json:"total_projects" tfsdk:"total_projects"`
}

type QuestionDependency struct {
	DependsOnQuestion     *string `json:"depends_on_question" tfsdk:"depends_on_question"`
	DependsOnQuestionName *string `json:"depends_on_question_name" tfsdk:"depends_on_question_name"`
	Operator              *string `json:"operator,omitempty" tfsdk:"operator"`
	Question              *string `json:"question" tfsdk:"question"`
	QuestionName          *string `json:"question_name" tfsdk:"question_name"`
	Url                   *string `json:"url" tfsdk:"url"`
}

type QuestionDependencyRequest struct {
	DependsOnQuestion *string `json:"depends_on_question" tfsdk:"depends_on_question"`
	Operator          *string `json:"operator,omitempty" tfsdk:"operator"`
	Question          *string `json:"question" tfsdk:"question"`
}

type QuestionOptions struct {
	Label *string `json:"label" tfsdk:"label"`
	Order *int64  `json:"order,omitempty" tfsdk:"order"`
}

type QuestionOptionsAdmin struct {
	Label        *string `json:"label" tfsdk:"label"`
	Order        *int64  `json:"order,omitempty" tfsdk:"order"`
	Question     *string `json:"question" tfsdk:"question"`
	QuestionUuid *string `json:"question_uuid" tfsdk:"question_uuid"`
	Url          *string `json:"url" tfsdk:"url"`
}

type QuestionOptionsAdminRequest struct {
	Label    *string `json:"label" tfsdk:"label"`
	Order    *int64  `json:"order,omitempty" tfsdk:"order"`
	Question *string `json:"question" tfsdk:"question"`
}

type QuestionTypeEnum struct {
}

type QuestionWithAnswer struct {
	Description   *string `json:"description" tfsdk:"description"`
	MaxFileSizeMb *int64  `json:"max_file_size_mb" tfsdk:"max_file_size_mb"`
	MaxFilesCount *int64  `json:"max_files_count" tfsdk:"max_files_count"`
	MaxValue      *string `json:"max_value" tfsdk:"max_value"`
	MinValue      *string `json:"min_value" tfsdk:"min_value"`
	Order         *int64  `json:"order" tfsdk:"order"`
	QuestionType  *string `json:"question_type" tfsdk:"question_type"`
	Required      *bool   `json:"required" tfsdk:"required"`
	UserGuidance  *string `json:"user_guidance" tfsdk:"user_guidance"`
}

type QuestionWithAnswerReviewer struct {
	AlwaysRequiresReview *bool   `json:"always_requires_review,omitempty" tfsdk:"always_requires_review"`
	Description          *string `json:"description" tfsdk:"description"`
	MaxFileSizeMb        *int64  `json:"max_file_size_mb" tfsdk:"max_file_size_mb"`
	MaxFilesCount        *int64  `json:"max_files_count" tfsdk:"max_files_count"`
	MaxValue             *string `json:"max_value" tfsdk:"max_value"`
	MinValue             *string `json:"min_value" tfsdk:"min_value"`
	Operator             *string `json:"operator,omitempty" tfsdk:"operator"`
	Order                *int64  `json:"order" tfsdk:"order"`
	QuestionType         *string `json:"question_type" tfsdk:"question_type"`
	Required             *bool   `json:"required" tfsdk:"required"`
	UserGuidance         *string `json:"user_guidance" tfsdk:"user_guidance"`
}

type Quota struct {
	Limit *int64  `json:"limit,omitempty" tfsdk:"limit"`
	Name  *string `json:"name,omitempty" tfsdk:"name"`
	Usage *int64  `json:"usage,omitempty" tfsdk:"usage"`
}

type QuotaRequest struct {
	Limit *int64  `json:"limit" tfsdk:"limit"`
	Name  *string `json:"name" tfsdk:"name"`
	Usage *int64  `json:"usage" tfsdk:"usage"`
}

type QuotasUpdateRequest struct {
}

type RancherApplication struct {
	AccessUrl                   *string `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	CatalogName                 *string `json:"catalog_name,omitempty" tfsdk:"catalog_name"`
	Created                     *string `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	ExternalUrl                 *string `json:"external_url,omitempty" tfsdk:"external_url"`
	IsLimitBased                *bool   `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool   `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string `json:"name,omitempty" tfsdk:"name"`
	Namespace                   *string `json:"namespace,omitempty" tfsdk:"namespace"`
	Project                     *string `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	RancherProject              *string `json:"rancher_project,omitempty" tfsdk:"rancher_project"`
	RancherProjectName          *string `json:"rancher_project_name,omitempty" tfsdk:"rancher_project_name"`
	ResourceType                *string `json:"resource_type,omitempty" tfsdk:"resource_type"`
	RuntimeState                *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceName                 *string `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string `json:"state,omitempty" tfsdk:"state"`
	Template                    *string `json:"template,omitempty" tfsdk:"template"`
	TemplateName                *string `json:"template_name,omitempty" tfsdk:"template_name"`
	Url                         *string `json:"url,omitempty" tfsdk:"url"`
	Version                     *string `json:"version,omitempty" tfsdk:"version"`
}

type RancherApplicationRequest struct {
	BackendId       *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage    *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback  *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	Name            *string `json:"name" tfsdk:"name"`
	Namespace       *string `json:"namespace,omitempty" tfsdk:"namespace"`
	NamespaceName   *string `json:"namespace_name,omitempty" tfsdk:"namespace_name"`
	Project         *string `json:"project" tfsdk:"project"`
	RancherProject  *string `json:"rancher_project" tfsdk:"rancher_project"`
	RuntimeState    *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceSettings *string `json:"service_settings" tfsdk:"service_settings"`
	Template        *string `json:"template" tfsdk:"template"`
	Version         *string `json:"version" tfsdk:"version"`
}

type RancherCatalog struct {
	Branch       *string `json:"branch" tfsdk:"branch"`
	CatalogUrl   *string `json:"catalog_url" tfsdk:"catalog_url"`
	Commit       *string `json:"commit" tfsdk:"commit"`
	Created      *string `json:"created" tfsdk:"created"`
	Description  *string `json:"description,omitempty" tfsdk:"description"`
	Modified     *string `json:"modified" tfsdk:"modified"`
	Name         *string `json:"name" tfsdk:"name"`
	RuntimeState *string `json:"runtime_state" tfsdk:"runtime_state"`
	Scope        *string `json:"scope" tfsdk:"scope"`
	ScopeType    *string `json:"scope_type" tfsdk:"scope_type"`
	Url          *string `json:"url" tfsdk:"url"`
}

type RancherCatalogCreate struct {
	Branch       *string `json:"branch" tfsdk:"branch"`
	CatalogUrl   *string `json:"catalog_url" tfsdk:"catalog_url"`
	Commit       *string `json:"commit" tfsdk:"commit"`
	Created      *string `json:"created" tfsdk:"created"`
	Description  *string `json:"description,omitempty" tfsdk:"description"`
	Modified     *string `json:"modified" tfsdk:"modified"`
	Name         *string `json:"name" tfsdk:"name"`
	Password     *string `json:"password,omitempty" tfsdk:"password"`
	RuntimeState *string `json:"runtime_state" tfsdk:"runtime_state"`
	Scope        *string `json:"scope" tfsdk:"scope"`
	ScopeType    *string `json:"scope_type" tfsdk:"scope_type"`
	Url          *string `json:"url" tfsdk:"url"`
	Username     *string `json:"username,omitempty" tfsdk:"username"`
}

type RancherCatalogCreateRequest struct {
	Branch      *string `json:"branch" tfsdk:"branch"`
	CatalogUrl  *string `json:"catalog_url" tfsdk:"catalog_url"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
	Password    *string `json:"password,omitempty" tfsdk:"password"`
	Scope       *string `json:"scope" tfsdk:"scope"`
	Username    *string `json:"username,omitempty" tfsdk:"username"`
}

type RancherCatalogRequest struct {
	Branch      *string `json:"branch" tfsdk:"branch"`
	CatalogUrl  *string `json:"catalog_url" tfsdk:"catalog_url"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
	Scope       *string `json:"scope" tfsdk:"scope"`
}

type RancherCatalogScopeType struct {
}

type RancherCatalogUpdate struct {
	Branch       *string `json:"branch" tfsdk:"branch"`
	CatalogUrl   *string `json:"catalog_url" tfsdk:"catalog_url"`
	Commit       *string `json:"commit" tfsdk:"commit"`
	Created      *string `json:"created" tfsdk:"created"`
	Description  *string `json:"description,omitempty" tfsdk:"description"`
	Modified     *string `json:"modified" tfsdk:"modified"`
	Name         *string `json:"name" tfsdk:"name"`
	RuntimeState *string `json:"runtime_state" tfsdk:"runtime_state"`
	Scope        *string `json:"scope" tfsdk:"scope"`
	ScopeType    *string `json:"scope_type" tfsdk:"scope_type"`
	Url          *string `json:"url" tfsdk:"url"`
}

type RancherCatalogUpdateRequest struct {
	Branch      *string `json:"branch" tfsdk:"branch"`
	CatalogUrl  *string `json:"catalog_url" tfsdk:"catalog_url"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
	Scope       *string `json:"scope" tfsdk:"scope"`
}

type RancherCluster struct {
	AccessUrl                   *string                 `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string                 `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created                     *string                 `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string                 `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string                 `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string                 `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string                 `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string                 `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string                 `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string                 `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string                 `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	InstallLonghorn             *bool                   `json:"install_longhorn,omitempty" tfsdk:"install_longhorn"`
	IsLimitBased                *bool                   `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                   `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	KubernetesVersion           *string                 `json:"kubernetes_version,omitempty" tfsdk:"kubernetes_version"`
	ManagementSecurityGroup     *string                 `json:"management_security_group,omitempty" tfsdk:"management_security_group"`
	MarketplaceCategoryName     *string                 `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                 `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                 `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                 `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                 `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                 `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                 `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                 `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string                 `json:"name,omitempty" tfsdk:"name"`
	Nodes                       []RancherNestedNode     `json:"nodes,omitempty" tfsdk:"nodes"`
	Project                     *string                 `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string                 `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string                 `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	PublicIps                   []RancherNestedPublicIP `json:"public_ips,omitempty" tfsdk:"public_ips"`
	ResourceType                *string                 `json:"resource_type,omitempty" tfsdk:"resource_type"`
	RuntimeState                *string                 `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceName                 *string                 `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string                 `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                 `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                 `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                 `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string                 `json:"state,omitempty" tfsdk:"state"`
	Tenant                      *string                 `json:"tenant,omitempty" tfsdk:"tenant"`
	TenantUuid                  *string                 `json:"tenant_uuid,omitempty" tfsdk:"tenant_uuid"`
	Url                         *string                 `json:"url,omitempty" tfsdk:"url"`
	VmProject                   *string                 `json:"vm_project,omitempty" tfsdk:"vm_project"`
}

type RancherClusterReference struct {
	MarketplaceUuid *string `json:"marketplace_uuid,omitempty" tfsdk:"marketplace_uuid"`
	Name            *string `json:"name,omitempty" tfsdk:"name"`
}

type RancherClusterRequest struct {
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	InstallLonghorn *bool   `json:"install_longhorn,omitempty" tfsdk:"install_longhorn"`
	Name            *string `json:"name" tfsdk:"name"`
	SshPublicKey    *string `json:"ssh_public_key,omitempty" tfsdk:"ssh_public_key"`
	VmProject       *string `json:"vm_project,omitempty" tfsdk:"vm_project"`
}

type RancherClusterSecurityGroupRule struct {
	Cidr        *string `json:"cidr,omitempty" tfsdk:"cidr"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Direction   *string `json:"direction,omitempty" tfsdk:"direction"`
	Ethertype   *string `json:"ethertype,omitempty" tfsdk:"ethertype"`
	FromPort    *int64  `json:"from_port,omitempty" tfsdk:"from_port"`
	Protocol    *string `json:"protocol,omitempty" tfsdk:"protocol"`
	ToPort      *int64  `json:"to_port,omitempty" tfsdk:"to_port"`
}

type RancherClusterSecurityGroupRuleRequest struct {
	Cidr        *string `json:"cidr,omitempty" tfsdk:"cidr"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Direction   *string `json:"direction,omitempty" tfsdk:"direction"`
	Ethertype   *string `json:"ethertype,omitempty" tfsdk:"ethertype"`
	FromPort    *int64  `json:"from_port,omitempty" tfsdk:"from_port"`
	Protocol    *string `json:"protocol,omitempty" tfsdk:"protocol"`
	ToPort      *int64  `json:"to_port,omitempty" tfsdk:"to_port"`
}

type RancherClusterTemplate struct {
	Created     *string                      `json:"created" tfsdk:"created"`
	Description *string                      `json:"description,omitempty" tfsdk:"description"`
	Modified    *string                      `json:"modified" tfsdk:"modified"`
	Name        *string                      `json:"name" tfsdk:"name"`
	Nodes       []RancherClusterTemplateNode `json:"nodes" tfsdk:"nodes"`
}

type RancherClusterTemplateNode struct {
	MinRam              *int64  `json:"min_ram" tfsdk:"min_ram"`
	MinVcpu             *int64  `json:"min_vcpu" tfsdk:"min_vcpu"`
	PreferredVolumeType *string `json:"preferred_volume_type,omitempty" tfsdk:"preferred_volume_type"`
	Role                *string `json:"role" tfsdk:"role"`
	SystemVolumeSize    *int64  `json:"system_volume_size" tfsdk:"system_volume_size"`
}

type RancherCreateNode struct {
	Cluster *string `json:"cluster" tfsdk:"cluster"`
	Role    *string `json:"role" tfsdk:"role"`
}

type RancherCreateNodeRequest struct {
	Cluster          *string             `json:"cluster" tfsdk:"cluster"`
	Cpu              *int64              `json:"cpu,omitempty" tfsdk:"cpu"`
	DataVolumes      []DataVolumeRequest `json:"data_volumes,omitempty" tfsdk:"data_volumes"`
	Flavor           *string             `json:"flavor,omitempty" tfsdk:"flavor"`
	Memory           *int64              `json:"memory,omitempty" tfsdk:"memory"`
	Role             *string             `json:"role" tfsdk:"role"`
	SshPublicKey     *string             `json:"ssh_public_key,omitempty" tfsdk:"ssh_public_key"`
	Subnet           *string             `json:"subnet" tfsdk:"subnet"`
	SystemVolumeSize *int64              `json:"system_volume_size,omitempty" tfsdk:"system_volume_size"`
	SystemVolumeType *string             `json:"system_volume_type,omitempty" tfsdk:"system_volume_type"`
	Tenant           *string             `json:"tenant,omitempty" tfsdk:"tenant"`
}

type RancherHPA struct {
	Cluster         *string `json:"cluster" tfsdk:"cluster"`
	ClusterName     *string `json:"cluster_name" tfsdk:"cluster_name"`
	ClusterUuid     *string `json:"cluster_uuid" tfsdk:"cluster_uuid"`
	Created         *string `json:"created" tfsdk:"created"`
	CurrentReplicas *int64  `json:"current_replicas" tfsdk:"current_replicas"`
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	DesiredReplicas *int64  `json:"desired_replicas" tfsdk:"desired_replicas"`
	MaxReplicas     *int64  `json:"max_replicas,omitempty" tfsdk:"max_replicas"`
	MinReplicas     *int64  `json:"min_replicas,omitempty" tfsdk:"min_replicas"`
	Modified        *string `json:"modified" tfsdk:"modified"`
	Name            *string `json:"name" tfsdk:"name"`
	Namespace       *string `json:"namespace" tfsdk:"namespace"`
	NamespaceName   *string `json:"namespace_name" tfsdk:"namespace_name"`
	NamespaceUuid   *string `json:"namespace_uuid" tfsdk:"namespace_uuid"`
	Project         *string `json:"project" tfsdk:"project"`
	ProjectName     *string `json:"project_name" tfsdk:"project_name"`
	ProjectUuid     *string `json:"project_uuid" tfsdk:"project_uuid"`
	RuntimeState    *string `json:"runtime_state" tfsdk:"runtime_state"`
	Url             *string `json:"url" tfsdk:"url"`
	Workload        *string `json:"workload,omitempty" tfsdk:"workload"`
	WorkloadName    *string `json:"workload_name" tfsdk:"workload_name"`
	WorkloadUuid    *string `json:"workload_uuid" tfsdk:"workload_uuid"`
}

type RancherHPARequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	MaxReplicas *int64  `json:"max_replicas,omitempty" tfsdk:"max_replicas"`
	MinReplicas *int64  `json:"min_replicas,omitempty" tfsdk:"min_replicas"`
	Name        *string `json:"name" tfsdk:"name"`
	Workload    *string `json:"workload,omitempty" tfsdk:"workload"`
}

type RancherImportYaml struct {
	DefaultNamespace *string `json:"default_namespace,omitempty" tfsdk:"default_namespace"`
	Namespace        *string `json:"namespace,omitempty" tfsdk:"namespace"`
	Yaml             *string `json:"yaml" tfsdk:"yaml"`
}

type RancherImportYamlRequest struct {
	DefaultNamespace *string `json:"default_namespace,omitempty" tfsdk:"default_namespace"`
	Namespace        *string `json:"namespace,omitempty" tfsdk:"namespace"`
	Yaml             *string `json:"yaml" tfsdk:"yaml"`
}

type RancherIngress struct {
	AccessUrl                   *string `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created                     *string `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	IsLimitBased                *bool   `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool   `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string `json:"name,omitempty" tfsdk:"name"`
	Namespace                   *string `json:"namespace,omitempty" tfsdk:"namespace"`
	NamespaceName               *string `json:"namespace_name,omitempty" tfsdk:"namespace_name"`
	Project                     *string `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	RancherProject              *string `json:"rancher_project,omitempty" tfsdk:"rancher_project"`
	RancherProjectName          *string `json:"rancher_project_name,omitempty" tfsdk:"rancher_project_name"`
	ResourceType                *string `json:"resource_type,omitempty" tfsdk:"resource_type"`
	RuntimeState                *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceName                 *string `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string `json:"state,omitempty" tfsdk:"state"`
	Url                         *string `json:"url,omitempty" tfsdk:"url"`
}

type RancherIngressRequest struct {
	BackendId       *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage    *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback  *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	Name            *string `json:"name" tfsdk:"name"`
	Namespace       *string `json:"namespace,omitempty" tfsdk:"namespace"`
	Project         *string `json:"project" tfsdk:"project"`
	RancherProject  *string `json:"rancher_project" tfsdk:"rancher_project"`
	RuntimeState    *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceSettings *string `json:"service_settings" tfsdk:"service_settings"`
}

type RancherNamespace struct {
	Created      *string `json:"created" tfsdk:"created"`
	Modified     *string `json:"modified" tfsdk:"modified"`
	Name         *string `json:"name" tfsdk:"name"`
	Project      *string `json:"project,omitempty" tfsdk:"project"`
	RuntimeState *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	Url          *string `json:"url" tfsdk:"url"`
}

type RancherNestedNamespace struct {
	Name *string `json:"name" tfsdk:"name"`
	Url  *string `json:"url" tfsdk:"url"`
}

type RancherNestedNode struct {
	BackendId      *string  `json:"backend_id,omitempty" tfsdk:"backend_id"`
	CpuAllocated   *float64 `json:"cpu_allocated,omitempty" tfsdk:"cpu_allocated"`
	CpuTotal       *int64   `json:"cpu_total,omitempty" tfsdk:"cpu_total"`
	Created        *string  `json:"created,omitempty" tfsdk:"created"`
	DockerVersion  *string  `json:"docker_version,omitempty" tfsdk:"docker_version"`
	ErrorMessage   *string  `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback *string  `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	Instance       *string  `json:"instance,omitempty" tfsdk:"instance"`
	K8sVersion     *string  `json:"k8s_version,omitempty" tfsdk:"k8s_version"`
	Modified       *string  `json:"modified,omitempty" tfsdk:"modified"`
	PodsAllocated  *int64   `json:"pods_allocated,omitempty" tfsdk:"pods_allocated"`
	PodsTotal      *int64   `json:"pods_total,omitempty" tfsdk:"pods_total"`
	RamAllocated   *int64   `json:"ram_allocated,omitempty" tfsdk:"ram_allocated"`
	RamTotal       *int64   `json:"ram_total,omitempty" tfsdk:"ram_total"`
	Role           *string  `json:"role,omitempty" tfsdk:"role"`
	RuntimeState   *string  `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	Url            *string  `json:"url,omitempty" tfsdk:"url"`
}

type RancherNestedNodeRequest struct {
	BackendId        *string             `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Cpu              *int64              `json:"cpu,omitempty" tfsdk:"cpu"`
	DataVolumes      []DataVolumeRequest `json:"data_volumes,omitempty" tfsdk:"data_volumes"`
	ErrorTraceback   *string             `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	Flavor           *string             `json:"flavor,omitempty" tfsdk:"flavor"`
	Memory           *int64              `json:"memory,omitempty" tfsdk:"memory"`
	Role             *string             `json:"role" tfsdk:"role"`
	Subnet           *string             `json:"subnet" tfsdk:"subnet"`
	SystemVolumeSize *int64              `json:"system_volume_size,omitempty" tfsdk:"system_volume_size"`
	SystemVolumeType *string             `json:"system_volume_type,omitempty" tfsdk:"system_volume_type"`
	Tenant           *string             `json:"tenant,omitempty" tfsdk:"tenant"`
}

type RancherNestedPublicIP struct {
	ExternalIpAddress *string `json:"external_ip_address,omitempty" tfsdk:"external_ip_address"`
	FloatingIp        *string `json:"floating_ip,omitempty" tfsdk:"floating_ip"`
	FloatingIpUuid    *string `json:"floating_ip_uuid,omitempty" tfsdk:"floating_ip_uuid"`
	IpAddress         *string `json:"ip_address,omitempty" tfsdk:"ip_address"`
}

type RancherNestedWorkload struct {
	Name *string `json:"name,omitempty" tfsdk:"name"`
	Url  *string `json:"url,omitempty" tfsdk:"url"`
}

type RancherNestedWorkloadRequest struct {
	Name *string `json:"name" tfsdk:"name"`
}

type RancherNode struct {
	BackendId               *string  `json:"backend_id" tfsdk:"backend_id"`
	Cluster                 *string  `json:"cluster" tfsdk:"cluster"`
	ClusterName             *string  `json:"cluster_name" tfsdk:"cluster_name"`
	ClusterUuid             *string  `json:"cluster_uuid" tfsdk:"cluster_uuid"`
	CpuAllocated            *float64 `json:"cpu_allocated" tfsdk:"cpu_allocated"`
	CpuTotal                *int64   `json:"cpu_total" tfsdk:"cpu_total"`
	Created                 *string  `json:"created" tfsdk:"created"`
	DockerVersion           *string  `json:"docker_version" tfsdk:"docker_version"`
	Instance                *string  `json:"instance" tfsdk:"instance"`
	InstanceMarketplaceUuid *string  `json:"instance_marketplace_uuid" tfsdk:"instance_marketplace_uuid"`
	InstanceName            *string  `json:"instance_name" tfsdk:"instance_name"`
	InstanceUuid            *string  `json:"instance_uuid" tfsdk:"instance_uuid"`
	K8sVersion              *string  `json:"k8s_version" tfsdk:"k8s_version"`
	Modified                *string  `json:"modified" tfsdk:"modified"`
	Name                    *string  `json:"name" tfsdk:"name"`
	PodsAllocated           *int64   `json:"pods_allocated" tfsdk:"pods_allocated"`
	PodsTotal               *int64   `json:"pods_total" tfsdk:"pods_total"`
	ProjectUuid             *string  `json:"project_uuid" tfsdk:"project_uuid"`
	RamAllocated            *int64   `json:"ram_allocated" tfsdk:"ram_allocated"`
	RamTotal                *int64   `json:"ram_total" tfsdk:"ram_total"`
	ResourceType            *string  `json:"resource_type" tfsdk:"resource_type"`
	Role                    *string  `json:"role" tfsdk:"role"`
	RuntimeState            *string  `json:"runtime_state" tfsdk:"runtime_state"`
	ServiceSettingsName     *string  `json:"service_settings_name" tfsdk:"service_settings_name"`
	ServiceSettingsUuid     *string  `json:"service_settings_uuid" tfsdk:"service_settings_uuid"`
	State                   *string  `json:"state" tfsdk:"state"`
	Url                     *string  `json:"url" tfsdk:"url"`
}

type RancherNodeRoleEnum struct {
}

type RancherProject struct {
	Cluster      *string                  `json:"cluster,omitempty" tfsdk:"cluster"`
	Created      *string                  `json:"created" tfsdk:"created"`
	Description  *string                  `json:"description,omitempty" tfsdk:"description"`
	Modified     *string                  `json:"modified" tfsdk:"modified"`
	Name         *string                  `json:"name" tfsdk:"name"`
	Namespaces   []RancherNestedNamespace `json:"namespaces" tfsdk:"namespaces"`
	RuntimeState *string                  `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	Url          *string                  `json:"url" tfsdk:"url"`
}

type RancherRoleScopeType struct {
}

type RancherService struct {
	AccessUrl                   *string                 `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string                 `json:"backend_id,omitempty" tfsdk:"backend_id"`
	ClusterIp                   *string                 `json:"cluster_ip,omitempty" tfsdk:"cluster_ip"`
	Created                     *string                 `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string                 `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string                 `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string                 `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string                 `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string                 `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string                 `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string                 `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string                 `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	IsLimitBased                *bool                   `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                   `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string                 `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                 `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                 `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                 `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                 `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                 `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                 `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                 `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string                 `json:"name,omitempty" tfsdk:"name"`
	Namespace                   *string                 `json:"namespace,omitempty" tfsdk:"namespace"`
	NamespaceName               *string                 `json:"namespace_name,omitempty" tfsdk:"namespace_name"`
	Project                     *string                 `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string                 `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string                 `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ResourceType                *string                 `json:"resource_type,omitempty" tfsdk:"resource_type"`
	RuntimeState                *string                 `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceName                 *string                 `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string                 `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                 `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                 `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                 `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string                 `json:"state,omitempty" tfsdk:"state"`
	TargetWorkloads             []RancherNestedWorkload `json:"target_workloads,omitempty" tfsdk:"target_workloads"`
	Url                         *string                 `json:"url,omitempty" tfsdk:"url"`
}

type RancherServiceCreate struct {
	AccessUrl                   *string                 `json:"access_url" tfsdk:"access_url"`
	BackendId                   *string                 `json:"backend_id,omitempty" tfsdk:"backend_id"`
	ClusterIp                   *string                 `json:"cluster_ip,omitempty" tfsdk:"cluster_ip"`
	Created                     *string                 `json:"created" tfsdk:"created"`
	Customer                    *string                 `json:"customer" tfsdk:"customer"`
	CustomerAbbreviation        *string                 `json:"customer_abbreviation" tfsdk:"customer_abbreviation"`
	CustomerName                *string                 `json:"customer_name" tfsdk:"customer_name"`
	CustomerNativeName          *string                 `json:"customer_native_name" tfsdk:"customer_native_name"`
	CustomerUuid                *string                 `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description                 *string                 `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string                 `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string                 `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	IsLimitBased                *bool                   `json:"is_limit_based" tfsdk:"is_limit_based"`
	IsUsageBased                *bool                   `json:"is_usage_based" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string                 `json:"marketplace_category_name" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string                 `json:"marketplace_category_uuid" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string                 `json:"marketplace_offering_name" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string                 `json:"marketplace_offering_uuid" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string                 `json:"marketplace_plan_uuid" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string                 `json:"marketplace_resource_state" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string                 `json:"marketplace_resource_uuid" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string                 `json:"modified" tfsdk:"modified"`
	Name                        *string                 `json:"name" tfsdk:"name"`
	Namespace                   *string                 `json:"namespace,omitempty" tfsdk:"namespace"`
	NamespaceName               *string                 `json:"namespace_name" tfsdk:"namespace_name"`
	Project                     *string                 `json:"project" tfsdk:"project"`
	ProjectName                 *string                 `json:"project_name" tfsdk:"project_name"`
	ProjectUuid                 *string                 `json:"project_uuid" tfsdk:"project_uuid"`
	ResourceType                *string                 `json:"resource_type" tfsdk:"resource_type"`
	RuntimeState                *string                 `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceName                 *string                 `json:"service_name" tfsdk:"service_name"`
	ServiceSettings             *string                 `json:"service_settings" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string                 `json:"service_settings_error_message" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string                 `json:"service_settings_state" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string                 `json:"service_settings_uuid" tfsdk:"service_settings_uuid"`
	State                       *string                 `json:"state" tfsdk:"state"`
	TargetWorkloads             []RancherWorkloadCreate `json:"target_workloads,omitempty" tfsdk:"target_workloads"`
	Url                         *string                 `json:"url" tfsdk:"url"`
}

type RancherServiceCreateRequest struct {
	BackendId       *string                        `json:"backend_id,omitempty" tfsdk:"backend_id"`
	ClusterIp       *string                        `json:"cluster_ip,omitempty" tfsdk:"cluster_ip"`
	Description     *string                        `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage    *string                        `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback  *string                        `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	Name            *string                        `json:"name" tfsdk:"name"`
	Namespace       *string                        `json:"namespace,omitempty" tfsdk:"namespace"`
	Project         *string                        `json:"project" tfsdk:"project"`
	RuntimeState    *string                        `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceSettings *string                        `json:"service_settings" tfsdk:"service_settings"`
	TargetWorkloads []RancherWorkloadCreateRequest `json:"target_workloads,omitempty" tfsdk:"target_workloads"`
}

type RancherServiceRequest struct {
	BackendId       *string                        `json:"backend_id,omitempty" tfsdk:"backend_id"`
	ClusterIp       *string                        `json:"cluster_ip,omitempty" tfsdk:"cluster_ip"`
	Description     *string                        `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage    *string                        `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback  *string                        `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	Name            *string                        `json:"name" tfsdk:"name"`
	Namespace       *string                        `json:"namespace,omitempty" tfsdk:"namespace"`
	Project         *string                        `json:"project" tfsdk:"project"`
	RuntimeState    *string                        `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceSettings *string                        `json:"service_settings" tfsdk:"service_settings"`
	TargetWorkloads []RancherNestedWorkloadRequest `json:"target_workloads" tfsdk:"target_workloads"`
}

type RancherTemplate struct {
	Catalog        *string  `json:"catalog,omitempty" tfsdk:"catalog"`
	CatalogName    *string  `json:"catalog_name" tfsdk:"catalog_name"`
	Cluster        *string  `json:"cluster,omitempty" tfsdk:"cluster"`
	Created        *string  `json:"created" tfsdk:"created"`
	DefaultVersion *string  `json:"default_version" tfsdk:"default_version"`
	Description    *string  `json:"description,omitempty" tfsdk:"description"`
	Icon           *string  `json:"icon,omitempty" tfsdk:"icon"`
	Modified       *string  `json:"modified" tfsdk:"modified"`
	Name           *string  `json:"name" tfsdk:"name"`
	Project        *string  `json:"project,omitempty" tfsdk:"project"`
	ProjectUrl     *string  `json:"project_url,omitempty" tfsdk:"project_url"`
	RuntimeState   *string  `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	Url            *string  `json:"url" tfsdk:"url"`
	Versions       []string `json:"versions" tfsdk:"versions"`
}

type RancherTemplateBaseQuestion struct {
	Default     *string `json:"default,omitempty" tfsdk:"default"`
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Group       *string `json:"group,omitempty" tfsdk:"group"`
	Label       *string `json:"label" tfsdk:"label"`
	Required    *bool   `json:"required,omitempty" tfsdk:"required"`
	ShowIf      *string `json:"showIf,omitempty" tfsdk:"showIf"`
	Type        *string `json:"type" tfsdk:"type"`
	Variable    *string `json:"variable" tfsdk:"variable"`
}

type RancherTemplateQuestion struct {
	Default           *string                       `json:"default,omitempty" tfsdk:"default"`
	Description       *string                       `json:"description,omitempty" tfsdk:"description"`
	Group             *string                       `json:"group,omitempty" tfsdk:"group"`
	Label             *string                       `json:"label" tfsdk:"label"`
	Required          *bool                         `json:"required,omitempty" tfsdk:"required"`
	ShowIf            *string                       `json:"showIf,omitempty" tfsdk:"showIf"`
	ShowSubquestionIf *string                       `json:"showSubquestionIf,omitempty" tfsdk:"showSubquestionIf"`
	Subquestions      []RancherTemplateBaseQuestion `json:"subquestions,omitempty" tfsdk:"subquestions"`
	Type              *string                       `json:"type" tfsdk:"type"`
	Variable          *string                       `json:"variable" tfsdk:"variable"`
}

type RancherTemplateQuestionType struct {
}

type RancherUser struct {
	ClusterRoles []RancherUserClusterLink `json:"cluster_roles" tfsdk:"cluster_roles"`
	FullName     *string                  `json:"full_name" tfsdk:"full_name"`
	IsActive     *bool                    `json:"is_active,omitempty" tfsdk:"is_active"`
	ProjectRoles []RancherUserProjectLink `json:"project_roles" tfsdk:"project_roles"`
	Settings     *string                  `json:"settings" tfsdk:"settings"`
	Url          *string                  `json:"url" tfsdk:"url"`
	User         *string                  `json:"user" tfsdk:"user"`
	UserName     *string                  `json:"user_name" tfsdk:"user_name"`
}

type RancherUserClusterLink struct {
	Cluster     *string `json:"cluster" tfsdk:"cluster"`
	ClusterName *string `json:"cluster_name" tfsdk:"cluster_name"`
	ClusterUuid *string `json:"cluster_uuid" tfsdk:"cluster_uuid"`
	Role        *string `json:"role" tfsdk:"role"`
}

type RancherUserProjectLink struct {
	Project     *string `json:"project" tfsdk:"project"`
	ProjectName *string `json:"project_name" tfsdk:"project_name"`
	ProjectUuid *string `json:"project_uuid" tfsdk:"project_uuid"`
	Role        *string `json:"role" tfsdk:"role"`
}

type RancherWorkload struct {
	Cluster       *string `json:"cluster,omitempty" tfsdk:"cluster"`
	ClusterName   *string `json:"cluster_name" tfsdk:"cluster_name"`
	ClusterUuid   *string `json:"cluster_uuid" tfsdk:"cluster_uuid"`
	Created       *string `json:"created" tfsdk:"created"`
	Modified      *string `json:"modified" tfsdk:"modified"`
	Name          *string `json:"name" tfsdk:"name"`
	Namespace     *string `json:"namespace,omitempty" tfsdk:"namespace"`
	NamespaceName *string `json:"namespace_name" tfsdk:"namespace_name"`
	NamespaceUuid *string `json:"namespace_uuid" tfsdk:"namespace_uuid"`
	Project       *string `json:"project,omitempty" tfsdk:"project"`
	ProjectName   *string `json:"project_name" tfsdk:"project_name"`
	ProjectUuid   *string `json:"project_uuid" tfsdk:"project_uuid"`
	RuntimeState  *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	Scale         *int64  `json:"scale" tfsdk:"scale"`
	Url           *string `json:"url" tfsdk:"url"`
}

type RancherWorkloadCreate struct {
	Url *string `json:"url" tfsdk:"url"`
}

type RancherWorkloadCreateRequest struct {
	Url *string `json:"url" tfsdk:"url"`
}

type RancherWorkloadRequest struct {
	Cluster      *string `json:"cluster,omitempty" tfsdk:"cluster"`
	Name         *string `json:"name" tfsdk:"name"`
	Namespace    *string `json:"namespace,omitempty" tfsdk:"namespace"`
	Project      *string `json:"project,omitempty" tfsdk:"project"`
	RuntimeState *string `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	Scale        *int64  `json:"scale" tfsdk:"scale"`
}

type ReferenceNumberRequest struct {
	ReferenceNumber *string `json:"reference_number,omitempty" tfsdk:"reference_number"`
}

type RemoteAllocation struct {
	AccessUrl                   *string `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created                     *string `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	IsActive                    *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	IsLimitBased                *bool   `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool   `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string `json:"name,omitempty" tfsdk:"name"`
	NodeLimit                   *int64  `json:"node_limit,omitempty" tfsdk:"node_limit"`
	NodeUsage                   *string `json:"node_usage,omitempty" tfsdk:"node_usage"`
	Project                     *string `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	RemoteProjectIdentifier     *string `json:"remote_project_identifier,omitempty" tfsdk:"remote_project_identifier"`
	ResourceType                *string `json:"resource_type,omitempty" tfsdk:"resource_type"`
	ServiceName                 *string `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string `json:"state,omitempty" tfsdk:"state"`
	Url                         *string `json:"url,omitempty" tfsdk:"url"`
}

type RemoteAllocationRequest struct {
	Description             *string `json:"description,omitempty" tfsdk:"description"`
	Name                    *string `json:"name" tfsdk:"name"`
	NodeLimit               *int64  `json:"node_limit,omitempty" tfsdk:"node_limit"`
	Project                 *string `json:"project" tfsdk:"project"`
	RemoteProjectIdentifier *string `json:"remote_project_identifier,omitempty" tfsdk:"remote_project_identifier"`
	ServiceSettings         *string `json:"service_settings" tfsdk:"service_settings"`
}

type RemoteAllocationSetLimits struct {
	NodeLimit *int64 `json:"node_limit" tfsdk:"node_limit"`
}

type RemoteAllocationSetLimitsRequest struct {
	NodeLimit *int64 `json:"node_limit" tfsdk:"node_limit"`
}

type RemoteAssociation struct {
	Allocation *string `json:"allocation" tfsdk:"allocation"`
}

type RemoteCredentialsRequest struct {
	ApiUrl *string `json:"api_url" tfsdk:"api_url"`
	Token  *string `json:"token" tfsdk:"token"`
}

type RemoteCustomer struct {
	Abbreviation *string `json:"abbreviation" tfsdk:"abbreviation"`
	Email        *string `json:"email" tfsdk:"email"`
	Name         *string `json:"name" tfsdk:"name"`
	PhoneNumber  *string `json:"phone_number" tfsdk:"phone_number"`
}

type RemoteEduteamsRequestRequest struct {
	Cuid *string `json:"cuid" tfsdk:"cuid"`
}

type RemoteEduteamsUUID struct {
}

type RemoteOffering struct {
	CategoryTitle *string `json:"category_title" tfsdk:"category_title"`
	Name          *string `json:"name" tfsdk:"name"`
	State         *string `json:"state" tfsdk:"state"`
	Type          *string `json:"type" tfsdk:"type"`
}

type RemoteOfferingCreateRequest struct {
	ApiUrl             *string `json:"api_url" tfsdk:"api_url"`
	LocalCategoryUuid  *string `json:"local_category_uuid" tfsdk:"local_category_uuid"`
	LocalCustomerUuid  *string `json:"local_customer_uuid" tfsdk:"local_customer_uuid"`
	RemoteCustomerUuid *string `json:"remote_customer_uuid" tfsdk:"remote_customer_uuid"`
	RemoteOfferingUuid *string `json:"remote_offering_uuid" tfsdk:"remote_offering_uuid"`
	Token              *string `json:"token" tfsdk:"token"`
}

type RemoteOfferingCreateResponse struct {
}

type RemoteProjectUpdateRequest struct {
	Created             *string `json:"created" tfsdk:"created"`
	CreatedBy           *int64  `json:"created_by,omitempty" tfsdk:"created_by"`
	CustomerName        *string `json:"customer_name" tfsdk:"customer_name"`
	CustomerUuid        *string `json:"customer_uuid" tfsdk:"customer_uuid"`
	NewDescription      *string `json:"new_description,omitempty" tfsdk:"new_description"`
	NewEndDate          *string `json:"new_end_date,omitempty" tfsdk:"new_end_date"`
	NewIsIndustry       *bool   `json:"new_is_industry,omitempty" tfsdk:"new_is_industry"`
	NewName             *string `json:"new_name,omitempty" tfsdk:"new_name"`
	NewOecdFos2007Code  *string `json:"new_oecd_fos_2007_code,omitempty" tfsdk:"new_oecd_fos_2007_code"`
	NewOecdFos2007Label *string `json:"new_oecd_fos_2007_label" tfsdk:"new_oecd_fos_2007_label"`
	OfferingName        *string `json:"offering_name" tfsdk:"offering_name"`
	OfferingUuid        *string `json:"offering_uuid" tfsdk:"offering_uuid"`
	OldDescription      *string `json:"old_description,omitempty" tfsdk:"old_description"`
	OldEndDate          *string `json:"old_end_date,omitempty" tfsdk:"old_end_date"`
	OldIsIndustry       *bool   `json:"old_is_industry,omitempty" tfsdk:"old_is_industry"`
	OldName             *string `json:"old_name,omitempty" tfsdk:"old_name"`
	OldOecdFos2007Code  *string `json:"old_oecd_fos_2007_code,omitempty" tfsdk:"old_oecd_fos_2007_code"`
	OldOecdFos2007Label *string `json:"old_oecd_fos_2007_label" tfsdk:"old_oecd_fos_2007_label"`
	ReviewComment       *string `json:"review_comment,omitempty" tfsdk:"review_comment"`
	ReviewedAt          *string `json:"reviewed_at" tfsdk:"reviewed_at"`
	ReviewedByFullName  *string `json:"reviewed_by_full_name" tfsdk:"reviewed_by_full_name"`
	ReviewedByUuid      *string `json:"reviewed_by_uuid" tfsdk:"reviewed_by_uuid"`
	State               *string `json:"state" tfsdk:"state"`
}

type RemoteResourceOrder struct {
	LocalState  *string `json:"local_state" tfsdk:"local_state"`
	OrderUuid   *string `json:"order_uuid" tfsdk:"order_uuid"`
	RemoteState *int64  `json:"remote_state" tfsdk:"remote_state"`
	SyncStatus  *string `json:"sync_status" tfsdk:"sync_status"`
}

type RemoteResourceOrderRemoteStateEnum struct {
}

type RemoteResourceSyncStatus struct {
	LastSync    *string `json:"last_sync" tfsdk:"last_sync"`
	LocalState  *string `json:"local_state" tfsdk:"local_state"`
	RemoteState *int64  `json:"remote_state" tfsdk:"remote_state"`
	SyncStatus  *string `json:"sync_status" tfsdk:"sync_status"`
}

type RemoteResourceSyncStatusRemoteStateEnum struct {
}

type RemoteResourceTeamMember struct {
	FullName   *string `json:"full_name" tfsdk:"full_name"`
	LocalRole  *string `json:"local_role" tfsdk:"local_role"`
	RemoteRole *string `json:"remote_role" tfsdk:"remote_role"`
	SyncStatus *string `json:"sync_status" tfsdk:"sync_status"`
}

type RemoteSynchronisation struct {
	ApiUrl                   *string                     `json:"api_url" tfsdk:"api_url"`
	Created                  *string                     `json:"created" tfsdk:"created"`
	ErrorMessage             *string                     `json:"error_message" tfsdk:"error_message"`
	GetStateDisplay          *string                     `json:"get_state_display" tfsdk:"get_state_display"`
	IsActive                 *bool                       `json:"is_active,omitempty" tfsdk:"is_active"`
	LastExecution            *string                     `json:"last_execution" tfsdk:"last_execution"`
	LastOutput               *string                     `json:"last_output" tfsdk:"last_output"`
	LocalServiceProvider     *string                     `json:"local_service_provider" tfsdk:"local_service_provider"`
	LocalServiceProviderName *string                     `json:"local_service_provider_name" tfsdk:"local_service_provider_name"`
	Modified                 *string                     `json:"modified" tfsdk:"modified"`
	RemoteOrganizationName   *string                     `json:"remote_organization_name" tfsdk:"remote_organization_name"`
	RemoteOrganizationUuid   *string                     `json:"remote_organization_uuid" tfsdk:"remote_organization_uuid"`
	RemotelocalcategorySet   []NestedRemoteLocalCategory `json:"remotelocalcategory_set" tfsdk:"remotelocalcategory_set"`
	Token                    *string                     `json:"token" tfsdk:"token"`
	Url                      *string                     `json:"url" tfsdk:"url"`
}

type RemoteSynchronisationRequest struct {
	ApiUrl                 *string                            `json:"api_url" tfsdk:"api_url"`
	IsActive               *bool                              `json:"is_active,omitempty" tfsdk:"is_active"`
	LocalServiceProvider   *string                            `json:"local_service_provider" tfsdk:"local_service_provider"`
	RemoteOrganizationName *string                            `json:"remote_organization_name" tfsdk:"remote_organization_name"`
	RemoteOrganizationUuid *string                            `json:"remote_organization_uuid" tfsdk:"remote_organization_uuid"`
	RemotelocalcategorySet []NestedRemoteLocalCategoryRequest `json:"remotelocalcategory_set" tfsdk:"remotelocalcategory_set"`
	Token                  *string                            `json:"token" tfsdk:"token"`
}

type RemoveOfferingComponentRequest struct {
}

type RemovePartitionRequest struct {
	PartitionUuid *string `json:"partition_uuid" tfsdk:"partition_uuid"`
}

type RemoveSoftwareCatalogRequest struct {
	OfferingCatalogUuid *string `json:"offering_catalog_uuid" tfsdk:"offering_catalog_uuid"`
}

type ReportSection struct {
	Body   *string `json:"body,omitempty" tfsdk:"body"`
	Header *string `json:"header,omitempty" tfsdk:"header"`
}

type ReportSectionRequest struct {
	Body   *string `json:"body" tfsdk:"body"`
	Header *string `json:"header" tfsdk:"header"`
}

type RequestType struct {
	IssueTypeName *string `json:"issue_type_name" tfsdk:"issue_type_name"`
	Name          *string `json:"name" tfsdk:"name"`
	Order         *int64  `json:"order,omitempty" tfsdk:"order"`
	Url           *string `json:"url" tfsdk:"url"`
}

type RequestTypeAdmin struct {
	BackendId     *int64  `json:"backend_id" tfsdk:"backend_id"`
	BackendName   *string `json:"backend_name" tfsdk:"backend_name"`
	IsActive      *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	IsSynced      *bool   `json:"is_synced" tfsdk:"is_synced"`
	IssueTypeName *string `json:"issue_type_name" tfsdk:"issue_type_name"`
	Name          *string `json:"name" tfsdk:"name"`
	Order         *int64  `json:"order,omitempty" tfsdk:"order"`
	Url           *string `json:"url" tfsdk:"url"`
}

type RequestTypeAdminRequest struct {
	IsActive      *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	IssueTypeName *string `json:"issue_type_name" tfsdk:"issue_type_name"`
	Name          *string `json:"name" tfsdk:"name"`
	Order         *int64  `json:"order,omitempty" tfsdk:"order"`
}

type RequestTypes struct {
}

type RequestedOffering struct {
	ApprovedBy               *string             `json:"approved_by" tfsdk:"approved_by"`
	ApprovedByName           *string             `json:"approved_by_name" tfsdk:"approved_by_name"`
	CallManagingOrganisation *string             `json:"call_managing_organisation" tfsdk:"call_managing_organisation"`
	CategoryName             *string             `json:"category_name" tfsdk:"category_name"`
	CategoryUuid             *string             `json:"category_uuid" tfsdk:"category_uuid"`
	Components               []OfferingComponent `json:"components" tfsdk:"components"`
	Created                  *string             `json:"created" tfsdk:"created"`
	CreatedBy                *string             `json:"created_by" tfsdk:"created_by"`
	CreatedByName            *string             `json:"created_by_name" tfsdk:"created_by_name"`
	Description              *string             `json:"description,omitempty" tfsdk:"description"`
	Offering                 *string             `json:"offering" tfsdk:"offering"`
	OfferingName             *string             `json:"offering_name" tfsdk:"offering_name"`
	OfferingUuid             *string             `json:"offering_uuid" tfsdk:"offering_uuid"`
	Plan                     *string             `json:"plan,omitempty" tfsdk:"plan"`
	ProviderName             *string             `json:"provider_name" tfsdk:"provider_name"`
	State                    *string             `json:"state" tfsdk:"state"`
	Url                      *string             `json:"url" tfsdk:"url"`
}

type RequestedOfferingRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Offering    *string `json:"offering" tfsdk:"offering"`
	Plan        *string `json:"plan,omitempty" tfsdk:"plan"`
}

type RequestedOfferingStates struct {
}

type RequestedResource struct {
	CallResourceTemplate     *string `json:"call_resource_template" tfsdk:"call_resource_template"`
	CallResourceTemplateName *string `json:"call_resource_template_name" tfsdk:"call_resource_template_name"`
	CreatedBy                *string `json:"created_by" tfsdk:"created_by"`
	CreatedByName            *string `json:"created_by_name" tfsdk:"created_by_name"`
	Description              *string `json:"description,omitempty" tfsdk:"description"`
	Resource                 *string `json:"resource" tfsdk:"resource"`
	ResourceName             *string `json:"resource_name" tfsdk:"resource_name"`
	Url                      *string `json:"url" tfsdk:"url"`
}

type RequestedResourceRequest struct {
	CallResourceTemplateUuid *string `json:"call_resource_template_uuid,omitempty" tfsdk:"call_resource_template_uuid"`
	Description              *string `json:"description,omitempty" tfsdk:"description"`
	RequestedOfferingUuid    *string `json:"requested_offering_uuid,omitempty" tfsdk:"requested_offering_uuid"`
}

type Resource struct {
	AvailableActions          []string            `json:"available_actions,omitempty" tfsdk:"available_actions"`
	BackendId                 *string             `json:"backend_id,omitempty" tfsdk:"backend_id"`
	CanTerminate              *bool               `json:"can_terminate,omitempty" tfsdk:"can_terminate"`
	CategoryIcon              *string             `json:"category_icon,omitempty" tfsdk:"category_icon"`
	CategoryTitle             *string             `json:"category_title,omitempty" tfsdk:"category_title"`
	CategoryUuid              *string             `json:"category_uuid,omitempty" tfsdk:"category_uuid"`
	Created                   *string             `json:"created,omitempty" tfsdk:"created"`
	CustomerName              *string             `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerSlug              *string             `json:"customer_slug,omitempty" tfsdk:"customer_slug"`
	CustomerUuid              *string             `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description               *string             `json:"description,omitempty" tfsdk:"description"`
	Downscaled                *bool               `json:"downscaled,omitempty" tfsdk:"downscaled"`
	EffectiveId               *string             `json:"effective_id,omitempty" tfsdk:"effective_id"`
	EndDate                   *string             `json:"end_date,omitempty" tfsdk:"end_date"`
	EndDateRequestedBy        *string             `json:"end_date_requested_by,omitempty" tfsdk:"end_date_requested_by"`
	Endpoints                 []NestedEndpoint    `json:"endpoints,omitempty" tfsdk:"endpoints"`
	ErrorMessage              *string             `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback            *string             `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	IsLimitBased              *bool               `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased              *bool               `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	LastSync                  *string             `json:"last_sync,omitempty" tfsdk:"last_sync"`
	Modified                  *string             `json:"modified,omitempty" tfsdk:"modified"`
	Name                      *string             `json:"name,omitempty" tfsdk:"name"`
	Offering                  *string             `json:"offering,omitempty" tfsdk:"offering"`
	OfferingBillable          *bool               `json:"offering_billable,omitempty" tfsdk:"offering_billable"`
	OfferingComponents        []OfferingComponent `json:"offering_components,omitempty" tfsdk:"offering_components"`
	OfferingDescription       *string             `json:"offering_description,omitempty" tfsdk:"offering_description"`
	OfferingImage             *string             `json:"offering_image,omitempty" tfsdk:"offering_image"`
	OfferingName              *string             `json:"offering_name,omitempty" tfsdk:"offering_name"`
	OfferingShared            *bool               `json:"offering_shared,omitempty" tfsdk:"offering_shared"`
	OfferingSlug              *string             `json:"offering_slug,omitempty" tfsdk:"offering_slug"`
	OfferingState             *string             `json:"offering_state,omitempty" tfsdk:"offering_state"`
	OfferingThumbnail         *string             `json:"offering_thumbnail,omitempty" tfsdk:"offering_thumbnail"`
	OfferingType              *string             `json:"offering_type,omitempty" tfsdk:"offering_type"`
	OfferingUuid              *string             `json:"offering_uuid,omitempty" tfsdk:"offering_uuid"`
	ParentName                *string             `json:"parent_name,omitempty" tfsdk:"parent_name"`
	ParentOfferingName        *string             `json:"parent_offering_name,omitempty" tfsdk:"parent_offering_name"`
	ParentOfferingSlug        *string             `json:"parent_offering_slug,omitempty" tfsdk:"parent_offering_slug"`
	ParentOfferingUuid        *string             `json:"parent_offering_uuid,omitempty" tfsdk:"parent_offering_uuid"`
	ParentUuid                *string             `json:"parent_uuid,omitempty" tfsdk:"parent_uuid"`
	Paused                    *bool               `json:"paused,omitempty" tfsdk:"paused"`
	Plan                      *string             `json:"plan,omitempty" tfsdk:"plan"`
	PlanDescription           *string             `json:"plan_description,omitempty" tfsdk:"plan_description"`
	PlanName                  *string             `json:"plan_name,omitempty" tfsdk:"plan_name"`
	PlanUnit                  *string             `json:"plan_unit,omitempty" tfsdk:"plan_unit"`
	PlanUuid                  *string             `json:"plan_uuid,omitempty" tfsdk:"plan_uuid"`
	Project                   *string             `json:"project,omitempty" tfsdk:"project"`
	ProjectDescription        *string             `json:"project_description,omitempty" tfsdk:"project_description"`
	ProjectEndDate            *string             `json:"project_end_date,omitempty" tfsdk:"project_end_date"`
	ProjectEndDateRequestedBy *string             `json:"project_end_date_requested_by,omitempty" tfsdk:"project_end_date_requested_by"`
	ProjectName               *string             `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectSlug               *string             `json:"project_slug,omitempty" tfsdk:"project_slug"`
	ProjectUuid               *string             `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ProviderName              *string             `json:"provider_name,omitempty" tfsdk:"provider_name"`
	ProviderSlug              *string             `json:"provider_slug,omitempty" tfsdk:"provider_slug"`
	ProviderUuid              *string             `json:"provider_uuid,omitempty" tfsdk:"provider_uuid"`
	Report                    []ReportSection     `json:"report,omitempty" tfsdk:"report"`
	ResourceType              *string             `json:"resource_type,omitempty" tfsdk:"resource_type"`
	ResourceUuid              *string             `json:"resource_uuid,omitempty" tfsdk:"resource_uuid"`
	RestrictMemberAccess      *bool               `json:"restrict_member_access,omitempty" tfsdk:"restrict_member_access"`
	Scope                     *string             `json:"scope,omitempty" tfsdk:"scope"`
	ServiceSettingsUuid       *string             `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	Slug                      *string             `json:"slug,omitempty" tfsdk:"slug"`
	State                     *string             `json:"state,omitempty" tfsdk:"state"`
	Url                       *string             `json:"url,omitempty" tfsdk:"url"`
	UserRequiresReconsent     *bool               `json:"user_requires_reconsent,omitempty" tfsdk:"user_requires_reconsent"`
	Username                  *string             `json:"username,omitempty" tfsdk:"username"`
}

type ResourceBackendIDRequest struct {
	BackendId *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
}

type ResourceBackendMetadataRequest struct {
}

type ResourceDownscaledRequest struct {
	Downscaled *bool `json:"downscaled,omitempty" tfsdk:"downscaled"`
}

type ResourceEndDateByProviderRequest struct {
	EndDate *string `json:"end_date,omitempty" tfsdk:"end_date"`
}

type ResourceLimitPeriod struct {
	BillingPeriods *int64  `json:"billing_periods,omitempty" tfsdk:"billing_periods"`
	End            *string `json:"end,omitempty" tfsdk:"end"`
	Quantity       *int64  `json:"quantity,omitempty" tfsdk:"quantity"`
	Start          *string `json:"start,omitempty" tfsdk:"start"`
	Total          *string `json:"total,omitempty" tfsdk:"total"`
}

type ResourceOffering struct {
	Name *string `json:"name" tfsdk:"name"`
}

type ResourceOptionsRequest struct {
}

type ResourcePausedRequest struct {
	Paused *bool `json:"paused,omitempty" tfsdk:"paused"`
}

type ResourcePlanPeriod struct {
	Components []BaseComponentUsage `json:"components" tfsdk:"components"`
	End        *string              `json:"end,omitempty" tfsdk:"end"`
	PlanName   *string              `json:"plan_name" tfsdk:"plan_name"`
	PlanUuid   *string              `json:"plan_uuid" tfsdk:"plan_uuid"`
	Start      *string              `json:"start,omitempty" tfsdk:"start"`
}

type ResourceProvisioningStats struct {
	AvgPendingDuration          *float64 `json:"avg_pending_duration" tfsdk:"avg_pending_duration"`
	AvgProvisioningDuration     *float64 `json:"avg_provisioning_duration" tfsdk:"avg_provisioning_duration"`
	OfferingName                *string  `json:"offering_name" tfsdk:"offering_name"`
	OfferingUuid                *string  `json:"offering_uuid" tfsdk:"offering_uuid"`
	ProvisioningCount           *int64   `json:"provisioning_count" tfsdk:"provisioning_count"`
	ProvisioningErrorCount      *int64   `json:"provisioning_error_count" tfsdk:"provisioning_error_count"`
	ProvisioningInProgressCount *int64   `json:"provisioning_in_progress_count" tfsdk:"provisioning_in_progress_count"`
	ProvisioningSuccessCount    *int64   `json:"provisioning_success_count" tfsdk:"provisioning_success_count"`
	ProvisioningSuccessRate     *float64 `json:"provisioning_success_rate" tfsdk:"provisioning_success_rate"`
	ServiceProviderName         *string  `json:"service_provider_name" tfsdk:"service_provider_name"`
	ServiceProviderUuid         *string  `json:"service_provider_uuid" tfsdk:"service_provider_uuid"`
}

type ResourceReallocateLimitsRequest struct {
	Targets []ResourceReallocateTargetRequest `json:"targets" tfsdk:"targets"`
}

type ResourceReallocateLimitsResponse struct {
	SourceOrderUuid  *string  `json:"source_order_uuid" tfsdk:"source_order_uuid"`
	TargetOrderUuids []string `json:"target_order_uuids" tfsdk:"target_order_uuids"`
}

type ResourceReallocateTargetRequest struct {
	ResourceUuid *string `json:"resource_uuid" tfsdk:"resource_uuid"`
}

type ResourceRenewRequest struct {
	Attachment      *string `json:"attachment,omitempty" tfsdk:"attachment"`
	ExtensionMonths *int64  `json:"extension_months" tfsdk:"extension_months"`
	RequestComment  *string `json:"request_comment,omitempty" tfsdk:"request_comment"`
}

type ResourceRenewRequestForm struct {
	Attachment      *string `json:"attachment,omitempty" tfsdk:"attachment"`
	ExtensionMonths *int64  `json:"extension_months" tfsdk:"extension_months"`
	RequestComment  *string `json:"request_comment,omitempty" tfsdk:"request_comment"`
}

type ResourceRenewRequestMultipart struct {
	Attachment      *string `json:"attachment,omitempty" tfsdk:"attachment"`
	ExtensionMonths *int64  `json:"extension_months" tfsdk:"extension_months"`
	RequestComment  *string `json:"request_comment,omitempty" tfsdk:"request_comment"`
}

type ResourceReportRequest struct {
	Report []ReportSectionRequest `json:"report" tfsdk:"report"`
}

type ResourceRequest struct {
	Downscaled *bool   `json:"downscaled,omitempty" tfsdk:"downscaled"`
	EndDate    *string `json:"end_date,omitempty" tfsdk:"end_date"`
	Name       *string `json:"name" tfsdk:"name"`
	Offering   *string `json:"offering" tfsdk:"offering"`
	Paused     *bool   `json:"paused,omitempty" tfsdk:"paused"`
	Plan       *string `json:"plan,omitempty" tfsdk:"plan"`
	Slug       *string `json:"slug,omitempty" tfsdk:"slug"`
}

type ResourceResponseStatus struct {
	Status *string `json:"status" tfsdk:"status"`
}

type ResourceRestrictMemberAccessRequest struct {
	RestrictMemberAccess *bool `json:"restrict_member_access,omitempty" tfsdk:"restrict_member_access"`
}

type ResourceSetLimitsRequest struct {
}

type ResourceSetStateErredRequest struct {
	ErrorMessage   *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
}

type ResourceSlugRequest struct {
	Slug *string `json:"slug" tfsdk:"slug"`
}

type ResourceState struct {
}

type ResourceSuggestNameRequest struct {
	Offering *string `json:"offering" tfsdk:"offering"`
	Project  *string `json:"project" tfsdk:"project"`
}

type ResourceSwitchPlanRequest struct {
	Plan *string `json:"plan" tfsdk:"plan"`
}

type ResourceTerminateRequest struct {
}

type ResourceUpdate struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	EndDate     *string `json:"end_date,omitempty" tfsdk:"end_date"`
	Name        *string `json:"name" tfsdk:"name"`
}

type ResourceUpdateLimitsRequest struct {
	RequestComment *string `json:"request_comment,omitempty" tfsdk:"request_comment"`
}

type ResourceUpdateRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	EndDate     *string `json:"end_date,omitempty" tfsdk:"end_date"`
	Name        *string `json:"name" tfsdk:"name"`
}

type ResourceUser struct {
	Resource     *string `json:"resource" tfsdk:"resource"`
	ResourceName *string `json:"resource_name" tfsdk:"resource_name"`
	ResourceUuid *string `json:"resource_uuid" tfsdk:"resource_uuid"`
	Role         *string `json:"role" tfsdk:"role"`
	RoleName     *string `json:"role_name" tfsdk:"role_name"`
	RoleUuid     *string `json:"role_uuid" tfsdk:"role_uuid"`
	User         *string `json:"user" tfsdk:"user"`
	UserFullName *string `json:"user_full_name" tfsdk:"user_full_name"`
	UserUsername *string `json:"user_username" tfsdk:"user_username"`
	UserUuid     *string `json:"user_uuid" tfsdk:"user_uuid"`
}

type ResourceUserRequest struct {
	Resource *string `json:"resource" tfsdk:"resource"`
	Role     *string `json:"role" tfsdk:"role"`
	User     *string `json:"user" tfsdk:"user"`
}

type ResourcesLimits struct {
	Name                  *string `json:"name" tfsdk:"name"`
	OfferingCountry       *string `json:"offering_country" tfsdk:"offering_country"`
	OfferingUuid          *string `json:"offering_uuid" tfsdk:"offering_uuid"`
	OrganizationGroupName *string `json:"organization_group_name" tfsdk:"organization_group_name"`
	OrganizationGroupUuid *string `json:"organization_group_uuid" tfsdk:"organization_group_uuid"`
	Value                 *int64  `json:"value" tfsdk:"value"`
}

type ReviewCommentRequest struct {
	Comment *string `json:"comment,omitempty" tfsdk:"comment"`
}

type ReviewStrategyEnum struct {
}

type ReviewSubmitRequest struct {
	SummaryPrivateComment *string `json:"summary_private_comment,omitempty" tfsdk:"summary_private_comment"`
	SummaryPublicComment  *string `json:"summary_public_comment,omitempty" tfsdk:"summary_public_comment"`
	SummaryScore          *int64  `json:"summary_score,omitempty" tfsdk:"summary_score"`
}

type RmqConnection struct {
	SourceIp *string `json:"source_ip" tfsdk:"source_ip"`
	Vhost    *string `json:"vhost" tfsdk:"vhost"`
}

type RmqSubscription struct {
	Created  *string `json:"created" tfsdk:"created"`
	SourceIp *string `json:"source_ip" tfsdk:"source_ip"`
}

type RmqUserStatsItem struct {
	Connections []RmqConnection `json:"connections" tfsdk:"connections"`
	Username    *string         `json:"username" tfsdk:"username"`
}

type RmqVHostStatsItem struct {
	Name          *string           `json:"name" tfsdk:"name"`
	Subscriptions []RmqSubscription `json:"subscriptions" tfsdk:"subscriptions"`
}

type RmqWaldurUser struct {
	Email    *string `json:"email" tfsdk:"email"`
	FullName *string `json:"full_name" tfsdk:"full_name"`
	Username *string `json:"username" tfsdk:"username"`
}

type RobotAccount struct {
	BackendId       *string       `json:"backend_id" tfsdk:"backend_id"`
	Created         *string       `json:"created" tfsdk:"created"`
	Description     *string       `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage    *string       `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback  *string       `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	Fingerprints    []Fingerprint `json:"fingerprints" tfsdk:"fingerprints"`
	Modified        *string       `json:"modified" tfsdk:"modified"`
	Resource        *string       `json:"resource" tfsdk:"resource"`
	ResponsibleUser *string       `json:"responsible_user,omitempty" tfsdk:"responsible_user"`
	State           *string       `json:"state,omitempty" tfsdk:"state"`
	Type            *string       `json:"type" tfsdk:"type"`
	Url             *string       `json:"url" tfsdk:"url"`
	Username        *string       `json:"username,omitempty" tfsdk:"username"`
	Users           []string      `json:"users,omitempty" tfsdk:"users"`
}

type RobotAccountDetails struct {
	BackendId      *string       `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created        *string       `json:"created,omitempty" tfsdk:"created"`
	CustomerName   *string       `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerUuid   *string       `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description    *string       `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage   *string       `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback *string       `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	Fingerprints   []Fingerprint `json:"fingerprints,omitempty" tfsdk:"fingerprints"`
	Modified       *string       `json:"modified,omitempty" tfsdk:"modified"`
	ProjectName    *string       `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid    *string       `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ProviderName   *string       `json:"provider_name,omitempty" tfsdk:"provider_name"`
	ProviderUuid   *string       `json:"provider_uuid,omitempty" tfsdk:"provider_uuid"`
	Resource       *string       `json:"resource,omitempty" tfsdk:"resource"`
	ResourceName   *string       `json:"resource_name,omitempty" tfsdk:"resource_name"`
	ResourceUuid   *string       `json:"resource_uuid,omitempty" tfsdk:"resource_uuid"`
	State          *string       `json:"state,omitempty" tfsdk:"state"`
	Type           *string       `json:"type,omitempty" tfsdk:"type"`
	Url            *string       `json:"url,omitempty" tfsdk:"url"`
	UserKeys       []SshKey      `json:"user_keys,omitempty" tfsdk:"user_keys"`
	Username       *string       `json:"username,omitempty" tfsdk:"username"`
	Users          []BasicUser   `json:"users,omitempty" tfsdk:"users"`
}

type RobotAccountErrorRequest struct {
	ErrorMessage *string `json:"error_message,omitempty" tfsdk:"error_message"`
}

type RobotAccountRequest struct {
	Description     *string  `json:"description,omitempty" tfsdk:"description"`
	Resource        *string  `json:"resource" tfsdk:"resource"`
	ResponsibleUser *string  `json:"responsible_user,omitempty" tfsdk:"responsible_user"`
	Type            *string  `json:"type" tfsdk:"type"`
	Username        *string  `json:"username,omitempty" tfsdk:"username"`
	Users           []string `json:"users,omitempty" tfsdk:"users"`
}

type RobotAccountStates struct {
}

type RoleDescription struct {
	Description   *string `json:"description,omitempty" tfsdk:"description"`
	DescriptionAr *string `json:"description_ar,omitempty" tfsdk:"description_ar"`
	DescriptionCs *string `json:"description_cs,omitempty" tfsdk:"description_cs"`
	DescriptionDa *string `json:"description_da,omitempty" tfsdk:"description_da"`
	DescriptionDe *string `json:"description_de,omitempty" tfsdk:"description_de"`
	DescriptionEn *string `json:"description_en,omitempty" tfsdk:"description_en"`
	DescriptionEs *string `json:"description_es,omitempty" tfsdk:"description_es"`
	DescriptionEt *string `json:"description_et,omitempty" tfsdk:"description_et"`
	DescriptionFr *string `json:"description_fr,omitempty" tfsdk:"description_fr"`
	DescriptionIt *string `json:"description_it,omitempty" tfsdk:"description_it"`
	DescriptionLt *string `json:"description_lt,omitempty" tfsdk:"description_lt"`
	DescriptionLv *string `json:"description_lv,omitempty" tfsdk:"description_lv"`
	DescriptionNb *string `json:"description_nb,omitempty" tfsdk:"description_nb"`
	DescriptionRu *string `json:"description_ru,omitempty" tfsdk:"description_ru"`
	DescriptionSv *string `json:"description_sv,omitempty" tfsdk:"description_sv"`
}

type RoleDescriptionRequest struct {
	Description   *string `json:"description,omitempty" tfsdk:"description"`
	DescriptionAr *string `json:"description_ar,omitempty" tfsdk:"description_ar"`
	DescriptionCs *string `json:"description_cs,omitempty" tfsdk:"description_cs"`
	DescriptionDa *string `json:"description_da,omitempty" tfsdk:"description_da"`
	DescriptionDe *string `json:"description_de,omitempty" tfsdk:"description_de"`
	DescriptionEn *string `json:"description_en,omitempty" tfsdk:"description_en"`
	DescriptionEs *string `json:"description_es,omitempty" tfsdk:"description_es"`
	DescriptionEt *string `json:"description_et,omitempty" tfsdk:"description_et"`
	DescriptionFr *string `json:"description_fr,omitempty" tfsdk:"description_fr"`
	DescriptionIt *string `json:"description_it,omitempty" tfsdk:"description_it"`
	DescriptionLt *string `json:"description_lt,omitempty" tfsdk:"description_lt"`
	DescriptionLv *string `json:"description_lv,omitempty" tfsdk:"description_lv"`
	DescriptionNb *string `json:"description_nb,omitempty" tfsdk:"description_nb"`
	DescriptionRu *string `json:"description_ru,omitempty" tfsdk:"description_ru"`
	DescriptionSv *string `json:"description_sv,omitempty" tfsdk:"description_sv"`
}

type RoleDetails struct {
	ContentType   *string  `json:"content_type,omitempty" tfsdk:"content_type"`
	Description   *string  `json:"description,omitempty" tfsdk:"description"`
	DescriptionAr *string  `json:"description_ar,omitempty" tfsdk:"description_ar"`
	DescriptionCs *string  `json:"description_cs,omitempty" tfsdk:"description_cs"`
	DescriptionDa *string  `json:"description_da,omitempty" tfsdk:"description_da"`
	DescriptionDe *string  `json:"description_de,omitempty" tfsdk:"description_de"`
	DescriptionEn *string  `json:"description_en,omitempty" tfsdk:"description_en"`
	DescriptionEs *string  `json:"description_es,omitempty" tfsdk:"description_es"`
	DescriptionEt *string  `json:"description_et,omitempty" tfsdk:"description_et"`
	DescriptionFr *string  `json:"description_fr,omitempty" tfsdk:"description_fr"`
	DescriptionIt *string  `json:"description_it,omitempty" tfsdk:"description_it"`
	DescriptionLt *string  `json:"description_lt,omitempty" tfsdk:"description_lt"`
	DescriptionLv *string  `json:"description_lv,omitempty" tfsdk:"description_lv"`
	DescriptionNb *string  `json:"description_nb,omitempty" tfsdk:"description_nb"`
	DescriptionRu *string  `json:"description_ru,omitempty" tfsdk:"description_ru"`
	DescriptionSv *string  `json:"description_sv,omitempty" tfsdk:"description_sv"`
	IsActive      *bool    `json:"is_active,omitempty" tfsdk:"is_active"`
	IsSystemRole  *bool    `json:"is_system_role,omitempty" tfsdk:"is_system_role"`
	Name          *string  `json:"name,omitempty" tfsdk:"name"`
	Permissions   []string `json:"permissions,omitempty" tfsdk:"permissions"`
	UsersCount    *int64   `json:"users_count,omitempty" tfsdk:"users_count"`
}

type RoleModifyRequest struct {
	ContentType   *string `json:"content_type" tfsdk:"content_type"`
	Description   *string `json:"description,omitempty" tfsdk:"description"`
	DescriptionAr *string `json:"description_ar,omitempty" tfsdk:"description_ar"`
	DescriptionCs *string `json:"description_cs,omitempty" tfsdk:"description_cs"`
	DescriptionDa *string `json:"description_da,omitempty" tfsdk:"description_da"`
	DescriptionDe *string `json:"description_de,omitempty" tfsdk:"description_de"`
	DescriptionEn *string `json:"description_en,omitempty" tfsdk:"description_en"`
	DescriptionEs *string `json:"description_es,omitempty" tfsdk:"description_es"`
	DescriptionEt *string `json:"description_et,omitempty" tfsdk:"description_et"`
	DescriptionFr *string `json:"description_fr,omitempty" tfsdk:"description_fr"`
	DescriptionIt *string `json:"description_it,omitempty" tfsdk:"description_it"`
	DescriptionLt *string `json:"description_lt,omitempty" tfsdk:"description_lt"`
	DescriptionLv *string `json:"description_lv,omitempty" tfsdk:"description_lv"`
	DescriptionNb *string `json:"description_nb,omitempty" tfsdk:"description_nb"`
	DescriptionRu *string `json:"description_ru,omitempty" tfsdk:"description_ru"`
	DescriptionSv *string `json:"description_sv,omitempty" tfsdk:"description_sv"`
	IsActive      *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	Name          *string `json:"name" tfsdk:"name"`
}

type RoleTemplate struct {
	DisplayName *string `json:"display_name" tfsdk:"display_name"`
	Name        *string `json:"name" tfsdk:"name"`
	ScopeType   *string `json:"scope_type" tfsdk:"scope_type"`
	Settings    *string `json:"settings" tfsdk:"settings"`
	Url         *string `json:"url" tfsdk:"url"`
}

type RoleType struct {
}

type RoundReviewer struct {
	AcceptedProposals *int64  `json:"accepted_proposals" tfsdk:"accepted_proposals"`
	Email             *string `json:"email" tfsdk:"email"`
	FullName          *string `json:"full_name" tfsdk:"full_name"`
	InReviewProposals *int64  `json:"in_review_proposals" tfsdk:"in_review_proposals"`
	RejectedProposals *int64  `json:"rejected_proposals" tfsdk:"rejected_proposals"`
}

type RoundStatus struct {
}

type Rule struct {
	CategoryTitle                     *string  `json:"category_title" tfsdk:"category_title"`
	CategoryUrl                       *string  `json:"category_url" tfsdk:"category_url"`
	Customer                          *string  `json:"customer,omitempty" tfsdk:"customer"`
	CustomerName                      *string  `json:"customer_name" tfsdk:"customer_name"`
	CustomerUuid                      *string  `json:"customer_uuid" tfsdk:"customer_uuid"`
	Name                              *string  `json:"name" tfsdk:"name"`
	OfferingName                      *string  `json:"offering_name" tfsdk:"offering_name"`
	OfferingUuid                      *string  `json:"offering_uuid" tfsdk:"offering_uuid"`
	Plan                              *string  `json:"plan,omitempty" tfsdk:"plan"`
	PlanName                          *string  `json:"plan_name" tfsdk:"plan_name"`
	ProjectRole                       *string  `json:"project_role,omitempty" tfsdk:"project_role"`
	ProjectRoleDescription            *string  `json:"project_role_description" tfsdk:"project_role_description"`
	ProjectRoleDisplayName            *string  `json:"project_role_display_name" tfsdk:"project_role_display_name"`
	Url                               *string  `json:"url" tfsdk:"url"`
	UseUserOrganizationAsCustomerName *bool    `json:"use_user_organization_as_customer_name,omitempty" tfsdk:"use_user_organization_as_customer_name"`
	UserAffiliations                  []string `json:"user_affiliations,omitempty" tfsdk:"user_affiliations"`
	UserEmailPatterns                 []string `json:"user_email_patterns,omitempty" tfsdk:"user_email_patterns"`
}

type RuleRequest struct {
	Customer                          *string  `json:"customer,omitempty" tfsdk:"customer"`
	Name                              *string  `json:"name" tfsdk:"name"`
	Plan                              *string  `json:"plan,omitempty" tfsdk:"plan"`
	ProjectRole                       *string  `json:"project_role,omitempty" tfsdk:"project_role"`
	ProjectRoleName                   *string  `json:"project_role_name,omitempty" tfsdk:"project_role_name"`
	UseUserOrganizationAsCustomerName *bool    `json:"use_user_organization_as_customer_name,omitempty" tfsdk:"use_user_organization_as_customer_name"`
	UserAffiliations                  []string `json:"user_affiliations,omitempty" tfsdk:"user_affiliations"`
	UserEmailPatterns                 []string `json:"user_email_patterns,omitempty" tfsdk:"user_email_patterns"`
}

type RuntimeStates struct {
	Label *string `json:"label" tfsdk:"label"`
	Value *string `json:"value" tfsdk:"value"`
}

type Saml2Login struct {
	Idp *string `json:"idp" tfsdk:"idp"`
}

type Saml2LoginComplete struct {
	SAMLResponse *string `json:"SAMLResponse" tfsdk:"SAMLResponse"`
}

type Saml2LoginCompleteRequest struct {
	SAMLResponse *string `json:"SAMLResponse" tfsdk:"SAMLResponse"`
}

type Saml2LoginRequest struct {
	Idp *string `json:"idp" tfsdk:"idp"`
}

type Saml2LogoutComplete struct {
	SAMLRequest  *string `json:"SAMLRequest,omitempty" tfsdk:"SAMLRequest"`
	SAMLResponse *string `json:"SAMLResponse,omitempty" tfsdk:"SAMLResponse"`
}

type Saml2LogoutCompleteRequest struct {
	SAMLRequest  *string `json:"SAMLRequest,omitempty" tfsdk:"SAMLRequest"`
	SAMLResponse *string `json:"SAMLResponse,omitempty" tfsdk:"SAMLResponse"`
}

type Saml2Provider struct {
	Name *string `json:"name" tfsdk:"name"`
	Url  *string `json:"url" tfsdk:"url"`
}

type Screenshot struct {
	Created      *string `json:"created" tfsdk:"created"`
	CustomerUuid *string `json:"customer_uuid" tfsdk:"customer_uuid"`
	Description  *string `json:"description,omitempty" tfsdk:"description"`
	Image        *string `json:"image" tfsdk:"image"`
	Name         *string `json:"name" tfsdk:"name"`
	Offering     *string `json:"offering" tfsdk:"offering"`
	Thumbnail    *string `json:"thumbnail" tfsdk:"thumbnail"`
	Url          *string `json:"url" tfsdk:"url"`
}

type ScreenshotRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Image       *string `json:"image" tfsdk:"image"`
	Name        *string `json:"name" tfsdk:"name"`
	Offering    *string `json:"offering" tfsdk:"offering"`
}

type ScreenshotRequestForm struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Image       *string `json:"image" tfsdk:"image"`
	Name        *string `json:"name" tfsdk:"name"`
	Offering    *string `json:"offering" tfsdk:"offering"`
}

type ScreenshotRequestMultipart struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Image       *string `json:"image" tfsdk:"image"`
	Name        *string `json:"name" tfsdk:"name"`
	Offering    *string `json:"offering" tfsdk:"offering"`
}

type Section struct {
	Category      *string `json:"category" tfsdk:"category"`
	CategoryTitle *string `json:"category_title" tfsdk:"category_title"`
	Created       *string `json:"created" tfsdk:"created"`
	IsStandalone  *bool   `json:"is_standalone,omitempty" tfsdk:"is_standalone"`
	Key           *string `json:"key" tfsdk:"key"`
	Title         *string `json:"title" tfsdk:"title"`
	Url           *string `json:"url" tfsdk:"url"`
}

type SectionRequest struct {
	Category     *string `json:"category" tfsdk:"category"`
	IsStandalone *bool   `json:"is_standalone,omitempty" tfsdk:"is_standalone"`
	Key          *string `json:"key" tfsdk:"key"`
	Title        *string `json:"title" tfsdk:"title"`
}

type ServiceAccountState struct {
}

type ServiceProvider struct {
	Created              *string             `json:"created,omitempty" tfsdk:"created"`
	Customer             *string             `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation *string             `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerCountry      *string             `json:"customer_country,omitempty" tfsdk:"customer_country"`
	CustomerImage        *string             `json:"customer_image,omitempty" tfsdk:"customer_image"`
	CustomerName         *string             `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName   *string             `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerSlug         *string             `json:"customer_slug,omitempty" tfsdk:"customer_slug"`
	CustomerUuid         *string             `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description          *string             `json:"description,omitempty" tfsdk:"description"`
	EnableNotifications  *bool               `json:"enable_notifications,omitempty" tfsdk:"enable_notifications"`
	Image                *string             `json:"image,omitempty" tfsdk:"image"`
	OfferingCount        *int64              `json:"offering_count,omitempty" tfsdk:"offering_count"`
	OrganizationGroups   []OrganizationGroup `json:"organization_groups,omitempty" tfsdk:"organization_groups"`
	Url                  *string             `json:"url,omitempty" tfsdk:"url"`
}

type ServiceProviderApiSecretCode struct {
	ApiSecretCode *string `json:"api_secret_code" tfsdk:"api_secret_code"`
}

type ServiceProviderChecklistSummary struct {
	CategoryName   *string `json:"category_name" tfsdk:"category_name"`
	ChecklistName  *string `json:"checklist_name" tfsdk:"checklist_name"`
	ChecklistUuid  *string `json:"checklist_uuid" tfsdk:"checklist_uuid"`
	OfferingsCount *int64  `json:"offerings_count" tfsdk:"offerings_count"`
	QuestionsCount *int64  `json:"questions_count" tfsdk:"questions_count"`
}

type ServiceProviderComplianceOverview struct {
	ChecklistName        *string  `json:"checklist_name" tfsdk:"checklist_name"`
	CompletedUsers       *int64   `json:"completed_users" tfsdk:"completed_users"`
	ComplianceRate       *float64 `json:"compliance_rate" tfsdk:"compliance_rate"`
	OfferingName         *string  `json:"offering_name" tfsdk:"offering_name"`
	OfferingUuid         *string  `json:"offering_uuid" tfsdk:"offering_uuid"`
	PendingUsers         *int64   `json:"pending_users" tfsdk:"pending_users"`
	TotalUsers           *int64   `json:"total_users" tfsdk:"total_users"`
	UsersWithCompletions *int64   `json:"users_with_completions" tfsdk:"users_with_completions"`
}

type ServiceProviderOfferingUserCompliance struct {
	ChecklistName        *string `json:"checklist_name" tfsdk:"checklist_name"`
	CompletionPercentage *int64  `json:"completion_percentage" tfsdk:"completion_percentage"`
	ComplianceStatus     *string `json:"compliance_status" tfsdk:"compliance_status"`
	Created              *string `json:"created" tfsdk:"created"`
	LastUpdated          *string `json:"last_updated" tfsdk:"last_updated"`
	OfferingName         *string `json:"offering_name" tfsdk:"offering_name"`
	State                *int64  `json:"state" tfsdk:"state"`
	UserEmail            *string `json:"user_email" tfsdk:"user_email"`
	UserFullName         *string `json:"user_full_name" tfsdk:"user_full_name"`
	Username             *string `json:"username" tfsdk:"username"`
}

type ServiceProviderOfferingUserComplianceStateEnum struct {
}

type ServiceProviderRequest struct {
	Customer            *string `json:"customer" tfsdk:"customer"`
	Description         *string `json:"description,omitempty" tfsdk:"description"`
	EnableNotifications *bool   `json:"enable_notifications,omitempty" tfsdk:"enable_notifications"`
	Image               *string `json:"image,omitempty" tfsdk:"image"`
}

type ServiceProviderRequestForm struct {
	Customer            *string `json:"customer" tfsdk:"customer"`
	Description         *string `json:"description,omitempty" tfsdk:"description"`
	EnableNotifications *bool   `json:"enable_notifications,omitempty" tfsdk:"enable_notifications"`
	Image               *string `json:"image,omitempty" tfsdk:"image"`
}

type ServiceProviderRequestMultipart struct {
	Customer            *string `json:"customer" tfsdk:"customer"`
	Description         *string `json:"description,omitempty" tfsdk:"description"`
	EnableNotifications *bool   `json:"enable_notifications,omitempty" tfsdk:"enable_notifications"`
	Image               *string `json:"image,omitempty" tfsdk:"image"`
}

type ServiceProviderRevenues struct {
	Month *int64 `json:"month" tfsdk:"month"`
	Total *int64 `json:"total" tfsdk:"total"`
	Year  *int64 `json:"year" tfsdk:"year"`
}

type ServiceProviderSignatureRequest struct {
	Customer *string `json:"customer" tfsdk:"customer"`
	Data     *string `json:"data" tfsdk:"data"`
	DryRun   *bool   `json:"dry_run,omitempty" tfsdk:"dry_run"`
}

type ServiceProviderStatistics struct {
	ActiveAndPausedOfferings *int64 `json:"active_and_paused_offerings" tfsdk:"active_and_paused_offerings"`
	ActiveCampaigns          *int64 `json:"active_campaigns" tfsdk:"active_campaigns"`
	ActiveResources          *int64 `json:"active_resources" tfsdk:"active_resources"`
	CurrentCustomers         *int64 `json:"current_customers" tfsdk:"current_customers"`
	CustomersNumberChange    *int64 `json:"customers_number_change" tfsdk:"customers_number_change"`
	ErredResources           *int64 `json:"erred_resources" tfsdk:"erred_resources"`
	PendingOrders            *int64 `json:"pending_orders" tfsdk:"pending_orders"`
	ResourcesNumberChange    *int64 `json:"resources_number_change" tfsdk:"resources_number_change"`
	UnresolvedTickets        *int64 `json:"unresolved_tickets" tfsdk:"unresolved_tickets"`
}

type ServiceSettings struct {
	Customer           *string `json:"customer,omitempty" tfsdk:"customer"`
	CustomerName       *string `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName *string `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	ErrorMessage       *string `json:"error_message,omitempty" tfsdk:"error_message"`
	Name               *string `json:"name,omitempty" tfsdk:"name"`
	Scope              *string `json:"scope,omitempty" tfsdk:"scope"`
	ScopeUuid          *string `json:"scope_uuid,omitempty" tfsdk:"scope_uuid"`
	Shared             *bool   `json:"shared,omitempty" tfsdk:"shared"`
	State              *string `json:"state,omitempty" tfsdk:"state"`
	TermsOfServices    *string `json:"terms_of_services,omitempty" tfsdk:"terms_of_services"`
	Type               *string `json:"type,omitempty" tfsdk:"type"`
	Url                *string `json:"url,omitempty" tfsdk:"url"`
}

type ServiceSettingsStateEnum struct {
}

type SetMtu struct {
	Mtu *int64 `json:"mtu" tfsdk:"mtu"`
}

type SetMtuRequest struct {
	Mtu *int64 `json:"mtu" tfsdk:"mtu"`
}

type SetOfferingsUsernameRequest struct {
	UserUuid *string `json:"user_uuid" tfsdk:"user_uuid"`
	Username *string `json:"username" tfsdk:"username"`
}

type SettingsMetadataResponse struct {
}

type SeverityEnum struct {
}

type SilenceActionRequest struct {
	DurationDays *int64 `json:"duration_days,omitempty" tfsdk:"duration_days"`
}

type SilenceActionResponse struct {
	DurationDays *int64  `json:"duration_days,omitempty" tfsdk:"duration_days"`
	Status       *string `json:"status" tfsdk:"status"`
}

type SlurmAllocation struct {
	AccessUrl                   *string `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	CpuLimit                    *int64  `json:"cpu_limit,omitempty" tfsdk:"cpu_limit"`
	CpuUsage                    *int64  `json:"cpu_usage,omitempty" tfsdk:"cpu_usage"`
	Created                     *string `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	Gateway                     *string `json:"gateway,omitempty" tfsdk:"gateway"`
	GpuLimit                    *int64  `json:"gpu_limit,omitempty" tfsdk:"gpu_limit"`
	GpuUsage                    *int64  `json:"gpu_usage,omitempty" tfsdk:"gpu_usage"`
	IsActive                    *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	IsLimitBased                *bool   `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool   `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string `json:"name,omitempty" tfsdk:"name"`
	Project                     *string `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	RamLimit                    *int64  `json:"ram_limit,omitempty" tfsdk:"ram_limit"`
	RamUsage                    *int64  `json:"ram_usage,omitempty" tfsdk:"ram_usage"`
	ResourceType                *string `json:"resource_type,omitempty" tfsdk:"resource_type"`
	ServiceName                 *string `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string `json:"state,omitempty" tfsdk:"state"`
	Url                         *string `json:"url,omitempty" tfsdk:"url"`
	Username                    *string `json:"username,omitempty" tfsdk:"username"`
}

type SlurmAllocationRequest struct {
	Description     *string `json:"description,omitempty" tfsdk:"description"`
	Name            *string `json:"name" tfsdk:"name"`
	Project         *string `json:"project" tfsdk:"project"`
	ServiceSettings *string `json:"service_settings" tfsdk:"service_settings"`
}

type SlurmAllocationSetLimits struct {
	CpuLimit *int64 `json:"cpu_limit" tfsdk:"cpu_limit"`
	GpuLimit *int64 `json:"gpu_limit" tfsdk:"gpu_limit"`
	RamLimit *int64 `json:"ram_limit" tfsdk:"ram_limit"`
}

type SlurmAllocationSetLimitsRequest struct {
	CpuLimit *int64 `json:"cpu_limit" tfsdk:"cpu_limit"`
	GpuLimit *int64 `json:"gpu_limit" tfsdk:"gpu_limit"`
	RamLimit *int64 `json:"ram_limit" tfsdk:"ram_limit"`
}

type SlurmAllocationUserUsage struct {
	Allocation *string `json:"allocation" tfsdk:"allocation"`
	CpuUsage   *int64  `json:"cpu_usage,omitempty" tfsdk:"cpu_usage"`
	FullName   *string `json:"full_name" tfsdk:"full_name"`
	GpuUsage   *int64  `json:"gpu_usage,omitempty" tfsdk:"gpu_usage"`
	Month      *int64  `json:"month" tfsdk:"month"`
	RamUsage   *int64  `json:"ram_usage,omitempty" tfsdk:"ram_usage"`
	User       *string `json:"user,omitempty" tfsdk:"user"`
	Username   *string `json:"username" tfsdk:"username"`
	Year       *int64  `json:"year" tfsdk:"year"`
}

type SlurmAssociation struct {
	Allocation *string `json:"allocation" tfsdk:"allocation"`
	Username   *string `json:"username" tfsdk:"username"`
}

type SlurmInvoicesSlurmPackageCreateOrderAttributes struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Name        *string `json:"name" tfsdk:"name"`
}

type SlurmPeriodicUsagePolicy struct {
	Actions                *string                        `json:"actions" tfsdk:"actions"`
	ApplyToAll             *bool                          `json:"apply_to_all,omitempty" tfsdk:"apply_to_all"`
	CarryoverEnabled       *bool                          `json:"carryover_enabled,omitempty" tfsdk:"carryover_enabled"`
	ComponentLimitsSet     []NestedOfferingComponentLimit `json:"component_limits_set" tfsdk:"component_limits_set"`
	Created                *string                        `json:"created" tfsdk:"created"`
	CreatedByFullName      *string                        `json:"created_by_full_name" tfsdk:"created_by_full_name"`
	CreatedByUsername      *string                        `json:"created_by_username" tfsdk:"created_by_username"`
	FairshareDecayHalfLife *int64                         `json:"fairshare_decay_half_life,omitempty" tfsdk:"fairshare_decay_half_life"`
	FiredDatetime          *string                        `json:"fired_datetime" tfsdk:"fired_datetime"`
	GraceRatio             *float64                       `json:"grace_ratio,omitempty" tfsdk:"grace_ratio"`
	HasFired               *bool                          `json:"has_fired" tfsdk:"has_fired"`
	LimitType              *string                        `json:"limit_type,omitempty" tfsdk:"limit_type"`
	OrganizationGroups     []string                       `json:"organization_groups,omitempty" tfsdk:"organization_groups"`
	Period                 *int64                         `json:"period,omitempty" tfsdk:"period"`
	PeriodName             *string                        `json:"period_name" tfsdk:"period_name"`
	QosStrategy            *string                        `json:"qos_strategy,omitempty" tfsdk:"qos_strategy"`
	RawUsageReset          *bool                          `json:"raw_usage_reset,omitempty" tfsdk:"raw_usage_reset"`
	Scope                  *string                        `json:"scope" tfsdk:"scope"`
	ScopeName              *string                        `json:"scope_name" tfsdk:"scope_name"`
	ScopeUuid              *string                        `json:"scope_uuid" tfsdk:"scope_uuid"`
	TresBillingEnabled     *bool                          `json:"tres_billing_enabled,omitempty" tfsdk:"tres_billing_enabled"`
	Url                    *string                        `json:"url" tfsdk:"url"`
}

type SlurmPeriodicUsagePolicyRequest struct {
	Actions                *string                               `json:"actions" tfsdk:"actions"`
	ApplyToAll             *bool                                 `json:"apply_to_all,omitempty" tfsdk:"apply_to_all"`
	CarryoverEnabled       *bool                                 `json:"carryover_enabled,omitempty" tfsdk:"carryover_enabled"`
	ComponentLimitsSet     []NestedOfferingComponentLimitRequest `json:"component_limits_set" tfsdk:"component_limits_set"`
	FairshareDecayHalfLife *int64                                `json:"fairshare_decay_half_life,omitempty" tfsdk:"fairshare_decay_half_life"`
	GraceRatio             *float64                              `json:"grace_ratio,omitempty" tfsdk:"grace_ratio"`
	LimitType              *string                               `json:"limit_type,omitempty" tfsdk:"limit_type"`
	OrganizationGroups     []string                              `json:"organization_groups,omitempty" tfsdk:"organization_groups"`
	Period                 *int64                                `json:"period,omitempty" tfsdk:"period"`
	QosStrategy            *string                               `json:"qos_strategy,omitempty" tfsdk:"qos_strategy"`
	RawUsageReset          *bool                                 `json:"raw_usage_reset,omitempty" tfsdk:"raw_usage_reset"`
	Scope                  *string                               `json:"scope" tfsdk:"scope"`
	TresBillingEnabled     *bool                                 `json:"tres_billing_enabled,omitempty" tfsdk:"tres_billing_enabled"`
}

type SmaxWebHookReceiver struct {
	Id *string `json:"id" tfsdk:"id"`
}

type SmaxWebHookReceiverRequest struct {
	Id *string `json:"id" tfsdk:"id"`
}

type SoftwareCatalog struct {
	AutoUpdateEnabled    *bool   `json:"auto_update_enabled,omitempty" tfsdk:"auto_update_enabled"`
	CatalogType          *string `json:"catalog_type,omitempty" tfsdk:"catalog_type"`
	CatalogTypeDisplay   *string `json:"catalog_type_display" tfsdk:"catalog_type_display"`
	Created              *string `json:"created" tfsdk:"created"`
	Description          *string `json:"description,omitempty" tfsdk:"description"`
	LastSuccessfulUpdate *string `json:"last_successful_update" tfsdk:"last_successful_update"`
	LastUpdateAttempt    *string `json:"last_update_attempt" tfsdk:"last_update_attempt"`
	Modified             *string `json:"modified" tfsdk:"modified"`
	Name                 *string `json:"name" tfsdk:"name"`
	PackageCount         *int64  `json:"package_count" tfsdk:"package_count"`
	SourceUrl            *string `json:"source_url,omitempty" tfsdk:"source_url"`
	UpdateErrors         *string `json:"update_errors,omitempty" tfsdk:"update_errors"`
	Url                  *string `json:"url" tfsdk:"url"`
	Version              *string `json:"version" tfsdk:"version"`
}

type SoftwareCatalogRequest struct {
	AutoUpdateEnabled *bool   `json:"auto_update_enabled,omitempty" tfsdk:"auto_update_enabled"`
	CatalogType       *string `json:"catalog_type,omitempty" tfsdk:"catalog_type"`
	Description       *string `json:"description,omitempty" tfsdk:"description"`
	Name              *string `json:"name" tfsdk:"name"`
	SourceUrl         *string `json:"source_url,omitempty" tfsdk:"source_url"`
	UpdateErrors      *string `json:"update_errors,omitempty" tfsdk:"update_errors"`
	Version           *string `json:"version" tfsdk:"version"`
}

type SoftwareCatalogUUID struct {
}

type SoftwarePackage struct {
	Catalog            *string                 `json:"catalog" tfsdk:"catalog"`
	CatalogName        *string                 `json:"catalog_name" tfsdk:"catalog_name"`
	CatalogType        *string                 `json:"catalog_type" tfsdk:"catalog_type"`
	CatalogTypeDisplay *string                 `json:"catalog_type_display" tfsdk:"catalog_type_display"`
	CatalogVersion     *string                 `json:"catalog_version" tfsdk:"catalog_version"`
	Created            *string                 `json:"created" tfsdk:"created"`
	Description        *string                 `json:"description,omitempty" tfsdk:"description"`
	ExtensionCount     *int64                  `json:"extension_count" tfsdk:"extension_count"`
	Homepage           *string                 `json:"homepage,omitempty" tfsdk:"homepage"`
	IsExtension        *bool                   `json:"is_extension,omitempty" tfsdk:"is_extension"`
	Modified           *string                 `json:"modified" tfsdk:"modified"`
	Name               *string                 `json:"name" tfsdk:"name"`
	ParentSoftware     *string                 `json:"parent_software,omitempty" tfsdk:"parent_software"`
	Url                *string                 `json:"url" tfsdk:"url"`
	VersionCount       *int64                  `json:"version_count" tfsdk:"version_count"`
	Versions           []NestedSoftwareVersion `json:"versions" tfsdk:"versions"`
}

type SoftwarePackageRequest struct {
	Catalog        *string `json:"catalog" tfsdk:"catalog"`
	Description    *string `json:"description,omitempty" tfsdk:"description"`
	Homepage       *string `json:"homepage,omitempty" tfsdk:"homepage"`
	IsExtension    *bool   `json:"is_extension,omitempty" tfsdk:"is_extension"`
	Name           *string `json:"name" tfsdk:"name"`
	ParentSoftware *string `json:"parent_software,omitempty" tfsdk:"parent_software"`
}

type SoftwareTarget struct {
	Created       *string `json:"created" tfsdk:"created"`
	Location      *string `json:"location" tfsdk:"location"`
	Modified      *string `json:"modified" tfsdk:"modified"`
	TargetName    *string `json:"target_name" tfsdk:"target_name"`
	TargetSubtype *string `json:"target_subtype" tfsdk:"target_subtype"`
	TargetType    *string `json:"target_type" tfsdk:"target_type"`
	Url           *string `json:"url" tfsdk:"url"`
}

type SoftwareVersion struct {
	CatalogType *string `json:"catalog_type" tfsdk:"catalog_type"`
	Created     *string `json:"created" tfsdk:"created"`
	Modified    *string `json:"modified" tfsdk:"modified"`
	PackageName *string `json:"package_name" tfsdk:"package_name"`
	ReleaseDate *string `json:"release_date" tfsdk:"release_date"`
	TargetCount *int64  `json:"target_count" tfsdk:"target_count"`
	Url         *string `json:"url" tfsdk:"url"`
	Version     *string `json:"version" tfsdk:"version"`
}

type SshKey struct {
	FingerprintMd5    *string `json:"fingerprint_md5,omitempty" tfsdk:"fingerprint_md5"`
	FingerprintSha256 *string `json:"fingerprint_sha256,omitempty" tfsdk:"fingerprint_sha256"`
	FingerprintSha512 *string `json:"fingerprint_sha512,omitempty" tfsdk:"fingerprint_sha512"`
	IsShared          *bool   `json:"is_shared,omitempty" tfsdk:"is_shared"`
	Name              *string `json:"name,omitempty" tfsdk:"name"`
	PublicKey         *string `json:"public_key,omitempty" tfsdk:"public_key"`
	Type              *string `json:"type,omitempty" tfsdk:"type"`
	Url               *string `json:"url,omitempty" tfsdk:"url"`
	UserUuid          *string `json:"user_uuid,omitempty" tfsdk:"user_uuid"`
}

type SshKeyRequest struct {
	Name      *string `json:"name,omitempty" tfsdk:"name"`
	PublicKey *string `json:"public_key" tfsdk:"public_key"`
}

type StateTransitionError struct {
	Detail *string `json:"detail" tfsdk:"detail"`
}

type StorageModeEnum struct {
}

type SubNetMapping struct {
	DstCidr *string `json:"dst_cidr" tfsdk:"dst_cidr"`
	SrcCidr *string `json:"src_cidr" tfsdk:"src_cidr"`
}

type SubNetMappingRequest struct {
	DstCidr *string `json:"dst_cidr" tfsdk:"dst_cidr"`
	SrcCidr *string `json:"src_cidr" tfsdk:"src_cidr"`
}

type SubmitRequestResponse struct {
	AutoApproved *bool   `json:"auto_approved" tfsdk:"auto_approved"`
	ScopeName    *string `json:"scope_name" tfsdk:"scope_name"`
	ScopeUuid    *string `json:"scope_uuid" tfsdk:"scope_uuid"`
}

type SubresourceOffering struct {
	Type *string `json:"type" tfsdk:"type"`
}

type SupportStats struct {
	ClosedThisMonthCount  *int64 `json:"closed_this_month_count" tfsdk:"closed_this_month_count"`
	OpenIssuesCount       *int64 `json:"open_issues_count" tfsdk:"open_issues_count"`
	RecentBroadcastsCount *int64 `json:"recent_broadcasts_count" tfsdk:"recent_broadcasts_count"`
}

type SupportUser struct {
	BackendId   *string `json:"backend_id" tfsdk:"backend_id"`
	BackendName *string `json:"backend_name,omitempty" tfsdk:"backend_name"`
	Name        *string `json:"name" tfsdk:"name"`
	Url         *string `json:"url" tfsdk:"url"`
	User        *string `json:"user" tfsdk:"user"`
}

type SupportedCountriesResponse struct {
	SupportedCountries []string `json:"supported_countries" tfsdk:"supported_countries"`
}

type SyncStatusEnum struct {
}

type TableSize struct {
	DataSize     *int64  `json:"data_size" tfsdk:"data_size"`
	ExternalSize *int64  `json:"external_size" tfsdk:"external_size"`
	TableName    *string `json:"table_name" tfsdk:"table_name"`
	TotalSize    *int64  `json:"total_size" tfsdk:"total_size"`
}

type Template struct {
	Attachments []TemplateAttachment `json:"attachments" tfsdk:"attachments"`
	Description *string              `json:"description" tfsdk:"description"`
	IssueType   *string              `json:"issue_type,omitempty" tfsdk:"issue_type"`
	Name        *string              `json:"name" tfsdk:"name"`
	Url         *string              `json:"url" tfsdk:"url"`
}

type TemplateAttachment struct {
	Created  *string `json:"created" tfsdk:"created"`
	File     *string `json:"file" tfsdk:"file"`
	FileSize *int64  `json:"file_size" tfsdk:"file_size"`
	MimeType *string `json:"mime_type" tfsdk:"mime_type"`
	Name     *string `json:"name" tfsdk:"name"`
}

type TemplateAttachmentRequest struct {
	File *string `json:"file" tfsdk:"file"`
	Name *string `json:"name" tfsdk:"name"`
}

type TemplateRequest struct {
	Description *string `json:"description" tfsdk:"description"`
	IssueType   *string `json:"issue_type,omitempty" tfsdk:"issue_type"`
	Name        *string `json:"name" tfsdk:"name"`
}

type TemplateVersion struct {
	AppReadme *string                   `json:"app_readme" tfsdk:"app_readme"`
	Questions []RancherTemplateQuestion `json:"questions" tfsdk:"questions"`
	Readme    *string                   `json:"readme" tfsdk:"readme"`
}

type Tenant struct {
	Name *string `json:"name" tfsdk:"name"`
	Url  *string `json:"url" tfsdk:"url"`
}

type TenantSecurityGroupUpdateRequest struct {
	Description *string                                         `json:"description,omitempty" tfsdk:"description"`
	Name        *string                                         `json:"name" tfsdk:"name"`
	Rules       []OpenStackSecurityGroupRuleUpdateByNameRequest `json:"rules,omitempty" tfsdk:"rules"`
}

type TimeSeriesToSData struct {
	Count *int64  `json:"count" tfsdk:"count"`
	Date  *string `json:"date" tfsdk:"date"`
}

type ToSConsentDashboard struct {
	AcceptedConsentsCount   *int64              `json:"accepted_consents_count" tfsdk:"accepted_consents_count"`
	ActiveUsersCount        *int64              `json:"active_users_count" tfsdk:"active_users_count"`
	ActiveUsersOverTime     []TimeSeriesToSData `json:"active_users_over_time" tfsdk:"active_users_over_time"`
	ActiveUsersPercentage   *float64            `json:"active_users_percentage" tfsdk:"active_users_percentage"`
	RevokedConsentsCount    *int64              `json:"revoked_consents_count" tfsdk:"revoked_consents_count"`
	RevokedConsentsOverTime []TimeSeriesToSData `json:"revoked_consents_over_time" tfsdk:"revoked_consents_over_time"`
	TosVersionAdoption      []VersionAdoption   `json:"tos_version_adoption" tfsdk:"tos_version_adoption"`
	TotalConsentsCount      *int64              `json:"total_consents_count" tfsdk:"total_consents_count"`
	TotalUsersCount         *int64              `json:"total_users_count" tfsdk:"total_users_count"`
}

type TokenRequest struct {
	Token *string `json:"token" tfsdk:"token"`
}

type TotalCustomerCost struct {
	Price *float64 `json:"price" tfsdk:"price"`
	Total *float64 `json:"total" tfsdk:"total"`
}

type UnsilenceActionResponse struct {
	Status *string `json:"status" tfsdk:"status"`
}

type UpdateActionsRequest struct {
	ProviderActionType *string `json:"provider_action_type,omitempty" tfsdk:"provider_action_type"`
}

type UpdateActionsResponse struct {
	Message            *string `json:"message" tfsdk:"message"`
	ProviderActionType *string `json:"provider_action_type,omitempty" tfsdk:"provider_action_type"`
	Status             *string `json:"status" tfsdk:"status"`
}

type UpdateOfferingComponentRequest struct {
	ArticleCode        *string `json:"article_code,omitempty" tfsdk:"article_code"`
	BillingType        *string `json:"billing_type" tfsdk:"billing_type"`
	DefaultLimit       *int64  `json:"default_limit,omitempty" tfsdk:"default_limit"`
	Description        *string `json:"description,omitempty" tfsdk:"description"`
	IsBoolean          *bool   `json:"is_boolean,omitempty" tfsdk:"is_boolean"`
	IsPrepaid          *bool   `json:"is_prepaid,omitempty" tfsdk:"is_prepaid"`
	LimitAmount        *int64  `json:"limit_amount,omitempty" tfsdk:"limit_amount"`
	LimitPeriod        *string `json:"limit_period,omitempty" tfsdk:"limit_period"`
	MaxAvailableLimit  *int64  `json:"max_available_limit,omitempty" tfsdk:"max_available_limit"`
	MaxPrepaidDuration *int64  `json:"max_prepaid_duration,omitempty" tfsdk:"max_prepaid_duration"`
	MaxValue           *int64  `json:"max_value,omitempty" tfsdk:"max_value"`
	MeasuredUnit       *string `json:"measured_unit,omitempty" tfsdk:"measured_unit"`
	MinPrepaidDuration *int64  `json:"min_prepaid_duration,omitempty" tfsdk:"min_prepaid_duration"`
	MinValue           *int64  `json:"min_value,omitempty" tfsdk:"min_value"`
	Name               *string `json:"name" tfsdk:"name"`
	OverageComponent   *string `json:"overage_component,omitempty" tfsdk:"overage_component"`
	Type               *string `json:"type" tfsdk:"type"`
	UnitFactor         *int64  `json:"unit_factor,omitempty" tfsdk:"unit_factor"`
}

type UrgencyEnum struct {
}

type User struct {
	AgreementDate                 *string      `json:"agreement_date,omitempty" tfsdk:"agreement_date"`
	BirthDate                     *string      `json:"birth_date,omitempty" tfsdk:"birth_date"`
	CivilNumber                   *string      `json:"civil_number,omitempty" tfsdk:"civil_number"`
	DateJoined                    *string      `json:"date_joined,omitempty" tfsdk:"date_joined"`
	Description                   *string      `json:"description,omitempty" tfsdk:"description"`
	Email                         *string      `json:"email,omitempty" tfsdk:"email"`
	FirstName                     *string      `json:"first_name,omitempty" tfsdk:"first_name"`
	FullName                      *string      `json:"full_name,omitempty" tfsdk:"full_name"`
	HasActiveSession              *bool        `json:"has_active_session,omitempty" tfsdk:"has_active_session"`
	IdentityProviderFields        []string     `json:"identity_provider_fields,omitempty" tfsdk:"identity_provider_fields"`
	IdentityProviderLabel         *string      `json:"identity_provider_label,omitempty" tfsdk:"identity_provider_label"`
	IdentityProviderManagementUrl *string      `json:"identity_provider_management_url,omitempty" tfsdk:"identity_provider_management_url"`
	IdentityProviderName          *string      `json:"identity_provider_name,omitempty" tfsdk:"identity_provider_name"`
	IdentitySource                *string      `json:"identity_source,omitempty" tfsdk:"identity_source"`
	Image                         *string      `json:"image,omitempty" tfsdk:"image"`
	IpAddress                     *string      `json:"ip_address,omitempty" tfsdk:"ip_address"`
	IsActive                      *bool        `json:"is_active,omitempty" tfsdk:"is_active"`
	IsStaff                       *bool        `json:"is_staff,omitempty" tfsdk:"is_staff"`
	IsSupport                     *bool        `json:"is_support,omitempty" tfsdk:"is_support"`
	JobTitle                      *string      `json:"job_title,omitempty" tfsdk:"job_title"`
	LastName                      *string      `json:"last_name,omitempty" tfsdk:"last_name"`
	NativeName                    *string      `json:"native_name,omitempty" tfsdk:"native_name"`
	NotificationsEnabled          *bool        `json:"notifications_enabled,omitempty" tfsdk:"notifications_enabled"`
	Organization                  *string      `json:"organization,omitempty" tfsdk:"organization"`
	Permissions                   []Permission `json:"permissions,omitempty" tfsdk:"permissions"`
	PhoneNumber                   *string      `json:"phone_number,omitempty" tfsdk:"phone_number"`
	PreferredLanguage             *string      `json:"preferred_language,omitempty" tfsdk:"preferred_language"`
	RegistrationMethod            *string      `json:"registration_method,omitempty" tfsdk:"registration_method"`
	RequestedEmail                *string      `json:"requested_email,omitempty" tfsdk:"requested_email"`
	Slug                          *string      `json:"slug,omitempty" tfsdk:"slug"`
	Token                         *string      `json:"token,omitempty" tfsdk:"token"`
	TokenExpiresAt                *string      `json:"token_expires_at,omitempty" tfsdk:"token_expires_at"`
	TokenLifetime                 *int64       `json:"token_lifetime,omitempty" tfsdk:"token_lifetime"`
	Url                           *string      `json:"url,omitempty" tfsdk:"url"`
	Username                      *string      `json:"username,omitempty" tfsdk:"username"`
}

type UserAction struct {
	ActionType            *string            `json:"action_type" tfsdk:"action_type"`
	CorrectiveActions     []CorrectiveAction `json:"corrective_actions" tfsdk:"corrective_actions"`
	Created               *string            `json:"created" tfsdk:"created"`
	DaysUntilDue          *int64             `json:"days_until_due" tfsdk:"days_until_due"`
	Description           *string            `json:"description" tfsdk:"description"`
	DueDate               *string            `json:"due_date,omitempty" tfsdk:"due_date"`
	IsEffectivelySilenced *bool              `json:"is_effectively_silenced" tfsdk:"is_effectively_silenced"`
	IsSilenced            *bool              `json:"is_silenced,omitempty" tfsdk:"is_silenced"`
	IsTemporarilySilenced *bool              `json:"is_temporarily_silenced" tfsdk:"is_temporarily_silenced"`
	Modified              *string            `json:"modified" tfsdk:"modified"`
	OfferingName          *string            `json:"offering_name,omitempty" tfsdk:"offering_name"`
	OfferingType          *string            `json:"offering_type,omitempty" tfsdk:"offering_type"`
	OrganizationName      *string            `json:"organization_name,omitempty" tfsdk:"organization_name"`
	OrganizationUuid      *string            `json:"organization_uuid,omitempty" tfsdk:"organization_uuid"`
	ProjectName           *string            `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid           *string            `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	RelatedObjectName     *string            `json:"related_object_name" tfsdk:"related_object_name"`
	RelatedObjectType     *string            `json:"related_object_type" tfsdk:"related_object_type"`
	RouteName             *string            `json:"route_name,omitempty" tfsdk:"route_name"`
	RouteParams           *string            `json:"route_params,omitempty" tfsdk:"route_params"`
	SilencedUntil         *string            `json:"silenced_until,omitempty" tfsdk:"silenced_until"`
	Title                 *string            `json:"title" tfsdk:"title"`
	Urgency               *string            `json:"urgency" tfsdk:"urgency"`
}

type UserActionExecution struct {
	CorrectiveActionLabel *string `json:"corrective_action_label" tfsdk:"corrective_action_label"`
	ErrorMessage          *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ExecutedAt            *string `json:"executed_at" tfsdk:"executed_at"`
	ExecutionMetadata     *string `json:"execution_metadata,omitempty" tfsdk:"execution_metadata"`
	Id                    *int64  `json:"id" tfsdk:"id"`
	Success               *bool   `json:"success,omitempty" tfsdk:"success"`
}

type UserActionProvider struct {
	ActionType          *string `json:"action_type" tfsdk:"action_type"`
	AppName             *string `json:"app_name" tfsdk:"app_name"`
	Id                  *int64  `json:"id" tfsdk:"id"`
	IsEnabled           *bool   `json:"is_enabled,omitempty" tfsdk:"is_enabled"`
	LastExecution       *string `json:"last_execution" tfsdk:"last_execution"`
	LastExecutionStatus *string `json:"last_execution_status" tfsdk:"last_execution_status"`
	ProviderClass       *string `json:"provider_class" tfsdk:"provider_class"`
	Schedule            *string `json:"schedule,omitempty" tfsdk:"schedule"`
}

type UserActionSummary struct {
	Overdue *int64 `json:"overdue" tfsdk:"overdue"`
	Total   *int64 `json:"total" tfsdk:"total"`
}

type UserAffiliationCount struct {
	Affiliation *string `json:"affiliation" tfsdk:"affiliation"`
	Count       *int64  `json:"count" tfsdk:"count"`
}

type UserAgreement struct {
	AgreementType *string `json:"agreement_type" tfsdk:"agreement_type"`
	Content       *string `json:"content" tfsdk:"content"`
	Created       *string `json:"created" tfsdk:"created"`
	Language      *string `json:"language" tfsdk:"language"`
	Modified      *string `json:"modified" tfsdk:"modified"`
	Url           *string `json:"url" tfsdk:"url"`
}

type UserAgreementRequest struct {
	AgreementType *string `json:"agreement_type" tfsdk:"agreement_type"`
	Content       *string `json:"content" tfsdk:"content"`
	Language      *string `json:"language" tfsdk:"language"`
}

type UserAuthMethodCount struct {
	Count  *int64  `json:"count" tfsdk:"count"`
	Method *string `json:"method" tfsdk:"method"`
}

type UserAuthToken struct {
	Created           *string `json:"created" tfsdk:"created"`
	Token             *string `json:"token" tfsdk:"token"`
	UserFirstName     *string `json:"user_first_name" tfsdk:"user_first_name"`
	UserIsActive      *bool   `json:"user_is_active" tfsdk:"user_is_active"`
	UserLastName      *string `json:"user_last_name" tfsdk:"user_last_name"`
	UserTokenLifetime *int64  `json:"user_token_lifetime" tfsdk:"user_token_lifetime"`
	UserUsername      *string `json:"user_username" tfsdk:"user_username"`
}

type UserChecklistCompletion struct {
	ChecklistDescription        *string  `json:"checklist_description" tfsdk:"checklist_description"`
	ChecklistName               *string  `json:"checklist_name" tfsdk:"checklist_name"`
	ChecklistUuid               *string  `json:"checklist_uuid" tfsdk:"checklist_uuid"`
	CompletionPercentage        *float64 `json:"completion_percentage" tfsdk:"completion_percentage"`
	Created                     *string  `json:"created" tfsdk:"created"`
	CustomerProviderName        *string  `json:"customer_provider_name" tfsdk:"customer_provider_name"`
	CustomerProviderUuid        *string  `json:"customer_provider_uuid" tfsdk:"customer_provider_uuid"`
	IsCompleted                 *bool    `json:"is_completed" tfsdk:"is_completed"`
	Modified                    *string  `json:"modified" tfsdk:"modified"`
	OfferingName                *string  `json:"offering_name" tfsdk:"offering_name"`
	OfferingUserUuid            *string  `json:"offering_user_uuid" tfsdk:"offering_user_uuid"`
	OfferingUuid                *string  `json:"offering_uuid" tfsdk:"offering_uuid"`
	RequiresReview              *bool    `json:"requires_review" tfsdk:"requires_review"`
	ReviewNotes                 *string  `json:"review_notes" tfsdk:"review_notes"`
	ReviewedAt                  *string  `json:"reviewed_at" tfsdk:"reviewed_at"`
	ReviewedBy                  *int64   `json:"reviewed_by" tfsdk:"reviewed_by"`
	UnansweredRequiredQuestions *int64   `json:"unanswered_required_questions" tfsdk:"unanswered_required_questions"`
}

type UserConsentInfo struct {
	AgreementDate *string `json:"agreement_date" tfsdk:"agreement_date"`
	IsRevoked     *bool   `json:"is_revoked" tfsdk:"is_revoked"`
	Version       *string `json:"version" tfsdk:"version"`
}

type UserEmailChangeRequest struct {
	Email *string `json:"email" tfsdk:"email"`
}

type UserIdentitySourceCount struct {
	Count          *int64  `json:"count" tfsdk:"count"`
	IdentitySource *string `json:"identity_source" tfsdk:"identity_source"`
}

type UserInfo struct {
	Shortname *string `json:"shortname,omitempty" tfsdk:"shortname"`
	User      *string `json:"user" tfsdk:"user"`
}

type UserInfoRequest struct {
	Shortname *string `json:"shortname,omitempty" tfsdk:"shortname"`
	User      *string `json:"user" tfsdk:"user"`
}

type UserOfferingConsent struct {
	AgreementDate     *string `json:"agreement_date" tfsdk:"agreement_date"`
	Created           *string `json:"created" tfsdk:"created"`
	HasConsent        *bool   `json:"has_consent" tfsdk:"has_consent"`
	Modified          *string `json:"modified" tfsdk:"modified"`
	OfferingName      *string `json:"offering_name" tfsdk:"offering_name"`
	OfferingSlug      *string `json:"offering_slug" tfsdk:"offering_slug"`
	OfferingUrl       *string `json:"offering_url" tfsdk:"offering_url"`
	OfferingUuid      *string `json:"offering_uuid" tfsdk:"offering_uuid"`
	RequiresReconsent *bool   `json:"requires_reconsent" tfsdk:"requires_reconsent"`
	RevocationDate    *string `json:"revocation_date" tfsdk:"revocation_date"`
	UserEmail         *string `json:"user_email" tfsdk:"user_email"`
	UserFullName      *string `json:"user_full_name" tfsdk:"user_full_name"`
	UserUsername      *string `json:"user_username" tfsdk:"user_username"`
	UserUuid          *string `json:"user_uuid" tfsdk:"user_uuid"`
	Version           *string `json:"version,omitempty" tfsdk:"version"`
}

type UserOfferingConsentCreate struct {
	Offering *string `json:"offering" tfsdk:"offering"`
}

type UserOfferingConsentCreateRequest struct {
	Offering *string `json:"offering" tfsdk:"offering"`
}

type UserOfferingConsentRequest struct {
	Version *string `json:"version,omitempty" tfsdk:"version"`
}

type UserOrganizationCount struct {
	Count        *int64  `json:"count" tfsdk:"count"`
	Organization *string `json:"organization" tfsdk:"organization"`
}

type UserRequest struct {
	AgreeWithPolicy      *bool   `json:"agree_with_policy,omitempty" tfsdk:"agree_with_policy"`
	BirthDate            *string `json:"birth_date,omitempty" tfsdk:"birth_date"`
	Description          *string `json:"description,omitempty" tfsdk:"description"`
	Email                *string `json:"email" tfsdk:"email"`
	FirstName            *string `json:"first_name,omitempty" tfsdk:"first_name"`
	Image                *string `json:"image,omitempty" tfsdk:"image"`
	IsActive             *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	IsStaff              *bool   `json:"is_staff,omitempty" tfsdk:"is_staff"`
	IsSupport            *bool   `json:"is_support,omitempty" tfsdk:"is_support"`
	JobTitle             *string `json:"job_title,omitempty" tfsdk:"job_title"`
	LastName             *string `json:"last_name,omitempty" tfsdk:"last_name"`
	NativeName           *string `json:"native_name,omitempty" tfsdk:"native_name"`
	NotificationsEnabled *bool   `json:"notifications_enabled,omitempty" tfsdk:"notifications_enabled"`
	Organization         *string `json:"organization,omitempty" tfsdk:"organization"`
	PhoneNumber          *string `json:"phone_number,omitempty" tfsdk:"phone_number"`
	PreferredLanguage    *string `json:"preferred_language,omitempty" tfsdk:"preferred_language"`
	Slug                 *string `json:"slug,omitempty" tfsdk:"slug"`
	TokenLifetime        *int64  `json:"token_lifetime,omitempty" tfsdk:"token_lifetime"`
	Username             *string `json:"username" tfsdk:"username"`
}

type UserRequestForm struct {
	AgreeWithPolicy      *bool   `json:"agree_with_policy,omitempty" tfsdk:"agree_with_policy"`
	BirthDate            *string `json:"birth_date,omitempty" tfsdk:"birth_date"`
	Description          *string `json:"description,omitempty" tfsdk:"description"`
	Email                *string `json:"email" tfsdk:"email"`
	FirstName            *string `json:"first_name,omitempty" tfsdk:"first_name"`
	Image                *string `json:"image,omitempty" tfsdk:"image"`
	IsActive             *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	IsStaff              *bool   `json:"is_staff,omitempty" tfsdk:"is_staff"`
	IsSupport            *bool   `json:"is_support,omitempty" tfsdk:"is_support"`
	JobTitle             *string `json:"job_title,omitempty" tfsdk:"job_title"`
	LastName             *string `json:"last_name,omitempty" tfsdk:"last_name"`
	NativeName           *string `json:"native_name,omitempty" tfsdk:"native_name"`
	NotificationsEnabled *bool   `json:"notifications_enabled,omitempty" tfsdk:"notifications_enabled"`
	Organization         *string `json:"organization,omitempty" tfsdk:"organization"`
	PhoneNumber          *string `json:"phone_number,omitempty" tfsdk:"phone_number"`
	PreferredLanguage    *string `json:"preferred_language,omitempty" tfsdk:"preferred_language"`
	Slug                 *string `json:"slug,omitempty" tfsdk:"slug"`
	TokenLifetime        *int64  `json:"token_lifetime,omitempty" tfsdk:"token_lifetime"`
	Username             *string `json:"username" tfsdk:"username"`
}

type UserRequestMultipart struct {
	AgreeWithPolicy      *bool   `json:"agree_with_policy,omitempty" tfsdk:"agree_with_policy"`
	BirthDate            *string `json:"birth_date,omitempty" tfsdk:"birth_date"`
	Description          *string `json:"description,omitempty" tfsdk:"description"`
	Email                *string `json:"email" tfsdk:"email"`
	FirstName            *string `json:"first_name,omitempty" tfsdk:"first_name"`
	Image                *string `json:"image,omitempty" tfsdk:"image"`
	IsActive             *bool   `json:"is_active,omitempty" tfsdk:"is_active"`
	IsStaff              *bool   `json:"is_staff,omitempty" tfsdk:"is_staff"`
	IsSupport            *bool   `json:"is_support,omitempty" tfsdk:"is_support"`
	JobTitle             *string `json:"job_title,omitempty" tfsdk:"job_title"`
	LastName             *string `json:"last_name,omitempty" tfsdk:"last_name"`
	NativeName           *string `json:"native_name,omitempty" tfsdk:"native_name"`
	NotificationsEnabled *bool   `json:"notifications_enabled,omitempty" tfsdk:"notifications_enabled"`
	Organization         *string `json:"organization,omitempty" tfsdk:"organization"`
	PhoneNumber          *string `json:"phone_number,omitempty" tfsdk:"phone_number"`
	PreferredLanguage    *string `json:"preferred_language,omitempty" tfsdk:"preferred_language"`
	Slug                 *string `json:"slug,omitempty" tfsdk:"slug"`
	TokenLifetime        *int64  `json:"token_lifetime,omitempty" tfsdk:"token_lifetime"`
	Username             *string `json:"username" tfsdk:"username"`
}

type UserRoleCreateRequest struct {
	ExpirationTime *string `json:"expiration_time,omitempty" tfsdk:"expiration_time"`
	Role           *string `json:"role" tfsdk:"role"`
	User           *string `json:"user" tfsdk:"user"`
}

type UserRoleDeleteRequest struct {
	ExpirationTime *string `json:"expiration_time,omitempty" tfsdk:"expiration_time"`
	Role           *string `json:"role" tfsdk:"role"`
	User           *string `json:"user" tfsdk:"user"`
}

type UserRoleDetails struct {
	Created           *string `json:"created,omitempty" tfsdk:"created"`
	CreatedByFullName *string `json:"created_by_full_name,omitempty" tfsdk:"created_by_full_name"`
	CreatedByUuid     *string `json:"created_by_uuid,omitempty" tfsdk:"created_by_uuid"`
	ExpirationTime    *string `json:"expiration_time,omitempty" tfsdk:"expiration_time"`
	RoleName          *string `json:"role_name,omitempty" tfsdk:"role_name"`
	RoleUuid          *string `json:"role_uuid,omitempty" tfsdk:"role_uuid"`
	UserEmail         *string `json:"user_email,omitempty" tfsdk:"user_email"`
	UserFullName      *string `json:"user_full_name,omitempty" tfsdk:"user_full_name"`
	UserImage         *string `json:"user_image,omitempty" tfsdk:"user_image"`
	UserUsername      *string `json:"user_username,omitempty" tfsdk:"user_username"`
	UserUuid          *string `json:"user_uuid,omitempty" tfsdk:"user_uuid"`
}

type UserRoleExpirationTime struct {
	ExpirationTime *string `json:"expiration_time" tfsdk:"expiration_time"`
}

type UserRoleUpdateRequest struct {
	ExpirationTime *string `json:"expiration_time,omitempty" tfsdk:"expiration_time"`
	Role           *string `json:"role" tfsdk:"role"`
	User           *string `json:"user" tfsdk:"user"`
}

type UsernameGenerationPolicyEnum struct {
}

type VMwareVirtualMachineCreateOrderAttributes struct {
	Cluster        *string `json:"cluster,omitempty" tfsdk:"cluster"`
	CoresPerSocket *int64  `json:"cores_per_socket,omitempty" tfsdk:"cores_per_socket"`
	Datastore      *string `json:"datastore,omitempty" tfsdk:"datastore"`
	Description    *string `json:"description,omitempty" tfsdk:"description"`
	GuestOs        *string `json:"guest_os,omitempty" tfsdk:"guest_os"`
	Name           *string `json:"name" tfsdk:"name"`
	Template       *string `json:"template,omitempty" tfsdk:"template"`
}

type ValidationDecisionEnum struct {
}

type ValidationMethodEnum struct {
}

type Version struct {
	LatestVersion *string `json:"latest_version,omitempty" tfsdk:"latest_version"`
	Version       *string `json:"version" tfsdk:"version"`
}

type VersionAdoption struct {
	UsersCount *int64  `json:"users_count" tfsdk:"users_count"`
	Version    *string `json:"version" tfsdk:"version"`
}

type VisibilityEnum struct {
}

type VisibleInvitationDetails struct {
	CreatedByFullName *string `json:"created_by_full_name" tfsdk:"created_by_full_name"`
	CreatedByImage    *string `json:"created_by_image" tfsdk:"created_by_image"`
	CreatedByUsername *string `json:"created_by_username" tfsdk:"created_by_username"`
	CustomerName      *string `json:"customer_name" tfsdk:"customer_name"`
	CustomerUuid      *string `json:"customer_uuid" tfsdk:"customer_uuid"`
	Email             *string `json:"email" tfsdk:"email"`
	ErrorMessage      *string `json:"error_message" tfsdk:"error_message"`
	ExecutionState    *string `json:"execution_state" tfsdk:"execution_state"`
	RoleDescription   *string `json:"role_description" tfsdk:"role_description"`
	RoleName          *string `json:"role_name" tfsdk:"role_name"`
	ScopeDescription  *string `json:"scope_description" tfsdk:"scope_description"`
	ScopeName         *string `json:"scope_name" tfsdk:"scope_name"`
	ScopeType         *string `json:"scope_type" tfsdk:"scope_type"`
	ScopeUuid         *string `json:"scope_uuid" tfsdk:"scope_uuid"`
	State             *string `json:"state" tfsdk:"state"`
}

type VmwareCluster struct {
	Name *string `json:"name" tfsdk:"name"`
	Url  *string `json:"url" tfsdk:"url"`
}

type VmwareDatastore struct {
	Capacity  *int64  `json:"capacity,omitempty" tfsdk:"capacity"`
	FreeSpace *int64  `json:"free_space,omitempty" tfsdk:"free_space"`
	Name      *string `json:"name" tfsdk:"name"`
	Type      *string `json:"type" tfsdk:"type"`
	Url       *string `json:"url" tfsdk:"url"`
}

type VmwareDisk struct {
	AccessUrl                   *string `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created                     *string `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	IsLimitBased                *bool   `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool   `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string `json:"name,omitempty" tfsdk:"name"`
	Project                     *string `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ResourceType                *string `json:"resource_type,omitempty" tfsdk:"resource_type"`
	ServiceName                 *string `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	Size                        *int64  `json:"size,omitempty" tfsdk:"size"`
	State                       *string `json:"state,omitempty" tfsdk:"state"`
	Url                         *string `json:"url,omitempty" tfsdk:"url"`
	Vm                          *string `json:"vm,omitempty" tfsdk:"vm"`
	VmName                      *string `json:"vm_name,omitempty" tfsdk:"vm_name"`
	VmUuid                      *string `json:"vm_uuid,omitempty" tfsdk:"vm_uuid"`
}

type VmwareDiskExtend struct {
	Size *int64 `json:"size" tfsdk:"size"`
}

type VmwareDiskExtendRequest struct {
	Size *int64 `json:"size" tfsdk:"size"`
}

type VmwareDiskRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Size        *int64  `json:"size" tfsdk:"size"`
}

type VmwareFolder struct {
	Name *string `json:"name" tfsdk:"name"`
	Url  *string `json:"url" tfsdk:"url"`
}

type VmwareLimit struct {
	MaxCoresPerSocket *int64 `json:"max_cores_per_socket" tfsdk:"max_cores_per_socket"`
	MaxCpu            *int64 `json:"max_cpu" tfsdk:"max_cpu"`
	MaxDisk           *int64 `json:"max_disk" tfsdk:"max_disk"`
	MaxDiskTotal      *int64 `json:"max_disk_total" tfsdk:"max_disk_total"`
	MaxRam            *int64 `json:"max_ram" tfsdk:"max_ram"`
}

type VmwareNestedDisk struct {
	Size *int64  `json:"size,omitempty" tfsdk:"size"`
	Url  *string `json:"url,omitempty" tfsdk:"url"`
}

type VmwareNestedDiskRequest struct {
	Size *int64 `json:"size" tfsdk:"size"`
}

type VmwareNestedNetwork struct {
	Url *string `json:"url" tfsdk:"url"`
}

type VmwareNestedNetworkRequest struct {
	Url *string `json:"url" tfsdk:"url"`
}

type VmwareNestedPort struct {
	MacAddress *string `json:"mac_address,omitempty" tfsdk:"mac_address"`
	Name       *string `json:"name,omitempty" tfsdk:"name"`
	Network    *string `json:"network,omitempty" tfsdk:"network"`
	Url        *string `json:"url,omitempty" tfsdk:"url"`
}

type VmwareNestedPortRequest struct {
	Name    *string `json:"name" tfsdk:"name"`
	Network *string `json:"network" tfsdk:"network"`
}

type VmwareNetwork struct {
	Name *string `json:"name" tfsdk:"name"`
	Type *string `json:"type" tfsdk:"type"`
	Url  *string `json:"url" tfsdk:"url"`
}

type VmwarePort struct {
	AccessUrl                   *string `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Created                     *string `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Description                 *string `json:"description,omitempty" tfsdk:"description"`
	ErrorMessage                *string `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	IsLimitBased                *bool   `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool   `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	MacAddress                  *string `json:"mac_address,omitempty" tfsdk:"mac_address"`
	MarketplaceCategoryName     *string `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string `json:"name,omitempty" tfsdk:"name"`
	Network                     *string `json:"network,omitempty" tfsdk:"network"`
	NetworkName                 *string `json:"network_name,omitempty" tfsdk:"network_name"`
	Project                     *string `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	ResourceType                *string `json:"resource_type,omitempty" tfsdk:"resource_type"`
	ServiceName                 *string `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string `json:"state,omitempty" tfsdk:"state"`
	Url                         *string `json:"url,omitempty" tfsdk:"url"`
	Vm                          *string `json:"vm,omitempty" tfsdk:"vm"`
	VmName                      *string `json:"vm_name,omitempty" tfsdk:"vm_name"`
	VmUuid                      *string `json:"vm_uuid,omitempty" tfsdk:"vm_uuid"`
}

type VmwarePortRequest struct {
	Description *string `json:"description,omitempty" tfsdk:"description"`
	Network     *string `json:"network" tfsdk:"network"`
}

type VmwareTemplate struct {
	Cores          *int64  `json:"cores,omitempty" tfsdk:"cores"`
	CoresPerSocket *int64  `json:"cores_per_socket,omitempty" tfsdk:"cores_per_socket"`
	Created        *string `json:"created" tfsdk:"created"`
	Description    *string `json:"description,omitempty" tfsdk:"description"`
	Disk           *int64  `json:"disk,omitempty" tfsdk:"disk"`
	GuestOs        *string `json:"guest_os" tfsdk:"guest_os"`
	GuestOsName    *string `json:"guest_os_name" tfsdk:"guest_os_name"`
	Modified       *string `json:"modified" tfsdk:"modified"`
	Name           *string `json:"name" tfsdk:"name"`
	Ram            *int64  `json:"ram,omitempty" tfsdk:"ram"`
	Url            *string `json:"url" tfsdk:"url"`
}

type VmwareVirtualMachine struct {
	AccessUrl                   *string            `json:"access_url,omitempty" tfsdk:"access_url"`
	BackendId                   *string            `json:"backend_id,omitempty" tfsdk:"backend_id"`
	Cluster                     *string            `json:"cluster,omitempty" tfsdk:"cluster"`
	ClusterName                 *string            `json:"cluster_name,omitempty" tfsdk:"cluster_name"`
	Cores                       *int64             `json:"cores,omitempty" tfsdk:"cores"`
	CoresPerSocket              *int64             `json:"cores_per_socket,omitempty" tfsdk:"cores_per_socket"`
	Created                     *string            `json:"created,omitempty" tfsdk:"created"`
	Customer                    *string            `json:"customer,omitempty" tfsdk:"customer"`
	CustomerAbbreviation        *string            `json:"customer_abbreviation,omitempty" tfsdk:"customer_abbreviation"`
	CustomerName                *string            `json:"customer_name,omitempty" tfsdk:"customer_name"`
	CustomerNativeName          *string            `json:"customer_native_name,omitempty" tfsdk:"customer_native_name"`
	CustomerUuid                *string            `json:"customer_uuid,omitempty" tfsdk:"customer_uuid"`
	Datastore                   *string            `json:"datastore,omitempty" tfsdk:"datastore"`
	DatastoreName               *string            `json:"datastore_name,omitempty" tfsdk:"datastore_name"`
	Description                 *string            `json:"description,omitempty" tfsdk:"description"`
	Disk                        *int64             `json:"disk,omitempty" tfsdk:"disk"`
	Disks                       []VmwareNestedDisk `json:"disks,omitempty" tfsdk:"disks"`
	ErrorMessage                *string            `json:"error_message,omitempty" tfsdk:"error_message"`
	ErrorTraceback              *string            `json:"error_traceback,omitempty" tfsdk:"error_traceback"`
	Folder                      *string            `json:"folder,omitempty" tfsdk:"folder"`
	FolderName                  *string            `json:"folder_name,omitempty" tfsdk:"folder_name"`
	GuestOs                     *string            `json:"guest_os,omitempty" tfsdk:"guest_os"`
	GuestOsName                 *string            `json:"guest_os_name,omitempty" tfsdk:"guest_os_name"`
	GuestPowerState             *string            `json:"guest_power_state,omitempty" tfsdk:"guest_power_state"`
	IsLimitBased                *bool              `json:"is_limit_based,omitempty" tfsdk:"is_limit_based"`
	IsUsageBased                *bool              `json:"is_usage_based,omitempty" tfsdk:"is_usage_based"`
	MarketplaceCategoryName     *string            `json:"marketplace_category_name,omitempty" tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid     *string            `json:"marketplace_category_uuid,omitempty" tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName     *string            `json:"marketplace_offering_name,omitempty" tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid     *string            `json:"marketplace_offering_uuid,omitempty" tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid         *string            `json:"marketplace_plan_uuid,omitempty" tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState    *string            `json:"marketplace_resource_state,omitempty" tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid     *string            `json:"marketplace_resource_uuid,omitempty" tfsdk:"marketplace_resource_uuid"`
	Modified                    *string            `json:"modified,omitempty" tfsdk:"modified"`
	Name                        *string            `json:"name,omitempty" tfsdk:"name"`
	Ports                       []VmwareNestedPort `json:"ports,omitempty" tfsdk:"ports"`
	Project                     *string            `json:"project,omitempty" tfsdk:"project"`
	ProjectName                 *string            `json:"project_name,omitempty" tfsdk:"project_name"`
	ProjectUuid                 *string            `json:"project_uuid,omitempty" tfsdk:"project_uuid"`
	Ram                         *int64             `json:"ram,omitempty" tfsdk:"ram"`
	ResourceType                *string            `json:"resource_type,omitempty" tfsdk:"resource_type"`
	RuntimeState                *string            `json:"runtime_state,omitempty" tfsdk:"runtime_state"`
	ServiceName                 *string            `json:"service_name,omitempty" tfsdk:"service_name"`
	ServiceSettings             *string            `json:"service_settings,omitempty" tfsdk:"service_settings"`
	ServiceSettingsErrorMessage *string            `json:"service_settings_error_message,omitempty" tfsdk:"service_settings_error_message"`
	ServiceSettingsState        *string            `json:"service_settings_state,omitempty" tfsdk:"service_settings_state"`
	ServiceSettingsUuid         *string            `json:"service_settings_uuid,omitempty" tfsdk:"service_settings_uuid"`
	State                       *string            `json:"state,omitempty" tfsdk:"state"`
	TemplateName                *string            `json:"template_name,omitempty" tfsdk:"template_name"`
	ToolsInstalled              *bool              `json:"tools_installed,omitempty" tfsdk:"tools_installed"`
	ToolsState                  *string            `json:"tools_state,omitempty" tfsdk:"tools_state"`
	Url                         *string            `json:"url,omitempty" tfsdk:"url"`
}

type VmwareVirtualMachineRequest struct {
	Cluster         *string                      `json:"cluster,omitempty" tfsdk:"cluster"`
	Cores           *int64                       `json:"cores,omitempty" tfsdk:"cores"`
	CoresPerSocket  *int64                       `json:"cores_per_socket,omitempty" tfsdk:"cores_per_socket"`
	Datastore       *string                      `json:"datastore,omitempty" tfsdk:"datastore"`
	Description     *string                      `json:"description,omitempty" tfsdk:"description"`
	Folder          *string                      `json:"folder,omitempty" tfsdk:"folder"`
	GuestOs         *string                      `json:"guest_os,omitempty" tfsdk:"guest_os"`
	Name            *string                      `json:"name" tfsdk:"name"`
	Networks        []VmwareNestedNetworkRequest `json:"networks,omitempty" tfsdk:"networks"`
	Project         *string                      `json:"project" tfsdk:"project"`
	Ram             *int64                       `json:"ram,omitempty" tfsdk:"ram"`
	ServiceSettings *string                      `json:"service_settings" tfsdk:"service_settings"`
	Template        *string                      `json:"template,omitempty" tfsdk:"template"`
}

type VolumeAttachRequest struct {
	Instance *string `json:"instance" tfsdk:"instance"`
}

type VolumeTypeEnum struct {
}

type VolumeTypeMapping struct {
	DstTypeUuid *string `json:"dst_type_uuid" tfsdk:"dst_type_uuid"`
	SrcTypeUuid *string `json:"src_type_uuid" tfsdk:"src_type_uuid"`
}

type VolumeTypeMappingRequest struct {
	DstTypeUuid *string `json:"dst_type_uuid" tfsdk:"dst_type_uuid"`
	SrcTypeUuid *string `json:"src_type_uuid" tfsdk:"src_type_uuid"`
}

type WebHook struct {
	AuthorEmail    *string  `json:"author_email" tfsdk:"author_email"`
	AuthorFullname *string  `json:"author_fullname" tfsdk:"author_fullname"`
	AuthorUsername *string  `json:"author_username" tfsdk:"author_username"`
	AuthorUuid     *string  `json:"author_uuid" tfsdk:"author_uuid"`
	ContentType    *string  `json:"content_type,omitempty" tfsdk:"content_type"`
	Created        *string  `json:"created" tfsdk:"created"`
	DestinationUrl *string  `json:"destination_url" tfsdk:"destination_url"`
	EventGroups    []string `json:"event_groups,omitempty" tfsdk:"event_groups"`
	EventTypes     []string `json:"event_types,omitempty" tfsdk:"event_types"`
	HookType       *string  `json:"hook_type" tfsdk:"hook_type"`
	IsActive       *bool    `json:"is_active,omitempty" tfsdk:"is_active"`
	Modified       *string  `json:"modified" tfsdk:"modified"`
	Url            *string  `json:"url" tfsdk:"url"`
}

type WebHookContentTypeEnum struct {
}

type WebHookReceiver struct {
	Comment            *JiraComment `json:"comment,omitempty" tfsdk:"comment"`
	Issue              *JiraIssue   `json:"issue" tfsdk:"issue"`
	IssueEventTypeName *string      `json:"issue_event_type_name,omitempty" tfsdk:"issue_event_type_name"`
	WebhookEvent       *string      `json:"webhookEvent" tfsdk:"webhookEvent"`
}

type WebHookReceiverRequest struct {
	Comment            *JiraCommentRequest `json:"comment,omitempty" tfsdk:"comment"`
	Issue              *JiraIssueRequest   `json:"issue" tfsdk:"issue"`
	IssueEventTypeName *string             `json:"issue_event_type_name,omitempty" tfsdk:"issue_event_type_name"`
	WebhookEvent       *string             `json:"webhookEvent" tfsdk:"webhookEvent"`
}

type WebHookRequest struct {
	ContentType    *string  `json:"content_type,omitempty" tfsdk:"content_type"`
	DestinationUrl *string  `json:"destination_url" tfsdk:"destination_url"`
	EventGroups    []string `json:"event_groups,omitempty" tfsdk:"event_groups"`
	EventTypes     []string `json:"event_types,omitempty" tfsdk:"event_types"`
	IsActive       *bool    `json:"is_active,omitempty" tfsdk:"is_active"`
}

type WidgetEnum struct {
}
