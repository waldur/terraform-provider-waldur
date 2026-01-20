package provider

import (
	"context"
	"net/http"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/waldur/terraform-provider-waldur/internal/actions"
	"github.com/waldur/terraform-provider-waldur/internal/client"
	"github.com/waldur/terraform-provider-waldur/internal/datasources"
	"github.com/waldur/terraform-provider-waldur/internal/resources"
)

// Ensure waldurProvider satisfies various provider interfaces.
var _ provider.Provider = &waldurProvider{}
var _ provider.ProviderWithActions = &waldurProvider{}
var _ provider.ProviderWithListResources = &waldurProvider{}

// waldurProvider defines the provider implementation.
type waldurProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance testing.
	version string
	// httpClient is an optional HTTP client for testing (e.g., with VCR)
	httpClient *http.Client
}

// waldurProviderModel describes the provider data model.
type waldurProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
	Token    types.String `tfsdk:"token"`
}

func (p *waldurProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "waldur"
	resp.Version = p.version
}

func (p *waldurProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "Waldur API endpoint URL. Can also be set via the `WALDUR_API_URL` environment variable.",
				Optional:            true,
			},
			"token": schema.StringAttribute{
				MarkdownDescription: "API authentication token. Can also be set via the `WALDUR_ACCESS_TOKEN` environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
		},
	}
}

func (p *waldurProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data waldurProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	endpoint := data.Endpoint.ValueString()
	if endpoint == "" {
		endpoint = os.Getenv("WALDUR_API_URL")
	}

	token := data.Token.ValueString()
	if token == "" {
		token = os.Getenv("WALDUR_ACCESS_TOKEN")
	}

	if endpoint == "" {
		resp.Diagnostics.AddError(
			"Missing API Endpoint",
			"The provider cannot be configured as there is no API endpoint URL. "+
				"Set the 'endpoint' value in the configuration or use the WALDUR_API_URL environment variable.",
		)
	}

	if token == "" {
		resp.Diagnostics.AddError(
			"Missing API Token",
			"The provider cannot be configured as there is no API token. "+
				"Set the 'token' value in the configuration or use the WALDUR_ACCESS_TOKEN environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Create API client
	apiClient, err := client.NewClient(&client.Config{
		Endpoint:   endpoint,
		Token:      token,
		HTTPClient: p.httpClient, // Pass through custom HTTP client for testing
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create API Client",
			"An error occurred when creating the API client: "+err.Error(),
		)
		return
	}

	// Make client available to resources and data sources
	resp.DataSourceData = apiClient
	resp.ResourceData = apiClient
}

func (p *waldurProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		resources.NewStructureProjectResource,
		resources.NewStructureCustomerResource,
		resources.NewMarketplaceOrderResource,
		resources.NewMarketplaceOfferingResource,
		resources.NewMarketplaceResourceResource,
		resources.NewOpenstackTenantResource,
		resources.NewOpenstackVolumeResource,
		resources.NewOpenstackVolumeAttachmentResource,
		resources.NewOpenstackInstanceResource,
		resources.NewOpenstackSecurityGroupResource,
		resources.NewOpenstackNetworkResource,
		resources.NewOpenstackServerGroupResource,
		resources.NewOpenstackSubnetResource,
		resources.NewOpenstackNetworkRbacPolicyResource,
		resources.NewOpenstackFloatingIpResource,
		resources.NewOpenstackPortResource,
	}
}

func (p *waldurProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		datasources.NewStructureProjectDataSource,
		datasources.NewStructureCustomerDataSource,
		datasources.NewCoreSshPublicKeyDataSource,
		datasources.NewMarketplaceOfferingDataSource,
		datasources.NewMarketplaceOrderDataSource,
		datasources.NewMarketplaceResourceDataSource,
		datasources.NewOpenstackTenantDataSource,
		datasources.NewOpenstackVolumeDataSource,
		datasources.NewOpenstackInstanceDataSource,
		datasources.NewOpenstackSecurityGroupDataSource,
		datasources.NewOpenstackSubnetDataSource,
		datasources.NewOpenstackNetworkDataSource,
		datasources.NewOpenstackFlavorDataSource,
		datasources.NewOpenstackImageDataSource,
		datasources.NewOpenstackVolumeTypeDataSource,
		datasources.NewOpenstackPortDataSource,
	}
}

func (p *waldurProvider) Actions(ctx context.Context) []func() action.Action {
	return []func() action.Action{
		actions.NewMarketplaceResourcePullAction,
		actions.NewMarketplaceResourceTerminateAction,
		actions.NewMarketplaceResourceUnlinkAction,
		actions.NewOpenstackTenantPullAction,
		actions.NewOpenstackTenantUnlinkAction,
		actions.NewOpenstackVolumePullAction,
		actions.NewOpenstackVolumeUnlinkAction,
		actions.NewOpenstackInstanceStartAction,
		actions.NewOpenstackInstanceStopAction,
		actions.NewOpenstackInstanceRestartAction,
		actions.NewOpenstackInstancePullAction,
		actions.NewOpenstackInstanceUnlinkAction,
		actions.NewOpenstackNetworkPullAction,
		actions.NewOpenstackNetworkUnlinkAction,
		actions.NewOpenstackSubnetConnectAction,
		actions.NewOpenstackSubnetDisconnectAction,
		actions.NewOpenstackSubnetPullAction,
		actions.NewOpenstackSubnetUnlinkAction,
		actions.NewOpenstackPortEnablePortAction,
		actions.NewOpenstackPortDisablePortAction,
		actions.NewOpenstackPortEnablePortSecurityAction,
		actions.NewOpenstackPortDisablePortSecurityAction,
		actions.NewOpenstackPortPullAction,
		actions.NewOpenstackPortUnlinkAction,
	}
}

func (p *waldurProvider) ListResources(ctx context.Context) []func() list.ListResource {
	return []func() list.ListResource{
		resources.NewStructureProjectList,
		resources.NewStructureCustomerList,
		resources.NewMarketplaceOrderList,
		resources.NewMarketplaceOfferingList,
		resources.NewMarketplaceResourceList,
		resources.NewOpenstackTenantList,
		resources.NewOpenstackVolumeList,
		resources.NewOpenstackVolumeAttachmentList,
		resources.NewOpenstackInstanceList,
		resources.NewOpenstackSecurityGroupList,
		resources.NewOpenstackNetworkList,
		resources.NewOpenstackServerGroupList,
		resources.NewOpenstackSubnetList,
		resources.NewOpenstackNetworkRbacPolicyList,
		resources.NewOpenstackFloatingIpList,
		resources.NewOpenstackPortList,
	}
}

func New(version string) func() provider.Provider {
	return NewWithHTTPClient(version, nil)
}

// NewWithHTTPClient creates a provider factory with a custom HTTP client.
// This is primarily used for testing with VCR or other HTTP mocking libraries.
func NewWithHTTPClient(version string, httpClient *http.Client) func() provider.Provider {
	return func() provider.Provider {
		return &waldurProvider{
			version:    version,
			httpClient: httpClient,
		}
	}
}
