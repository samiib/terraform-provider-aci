// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ciscoecosystem/aci-go-client/v2/client"

	// temporary unused until muxing is removed
	// "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	// temporary unused until muxing is removed
	// "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var globalAnnotation string

// Ensure AciProvider satisfies various provider interfaces.
var _ provider.Provider = &AciProvider{}

// AciProvider defines the provider implementation.
type AciProvider struct {
	version string
}

// AciProviderModel describes the provider data model.
type AciProviderModel struct {
	Username   types.String `tfsdk:"username"`
	Password   types.String `tfsdk:"password"`
	URL        types.String `tfsdk:"url"`
	PrivateKey types.String `tfsdk:"private_key"`
	Certname   types.String `tfsdk:"cert_name"`
	ProxyUrl   types.String `tfsdk:"proxy_url"`
	ProxyCreds types.String `tfsdk:"proxy_creds"`
	// IsInsecure         types.Bool   `tfsdk:"insecure"`
	// ValidateRelationDn types.Bool   `tfsdk:"validate_relation_dn"`
	// MaxRetries         types.Int64  `tfsdk:"retries"`
	IsInsecure         types.String `tfsdk:"insecure"`
	ValidateRelationDn types.String `tfsdk:"validate_relation_dn"`
	MaxRetries         types.String `tfsdk:"retries"`
	Annotation         types.String `tfsdk:"annotation"`
}

func (p *AciProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "aci"
	resp.Version = p.version
}

func (p *AciProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{

		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				Description: "Username for the APIC Account. This can also be set as the ACI_USERNAME environment variable.",
				Optional:    true,
			},
			"password": schema.StringAttribute{
				Description: "Password for the APIC Account. This can also be set as the ACI_PASSWORD environment variable.",
				Optional:    true,
			},
			"url": schema.StringAttribute{
				Description: "URL of the Cisco ACI web interface. This can also be set as the ACI_URL environment variable.",
				Optional:    true,
			},
			"private_key": schema.StringAttribute{
				Description: "Private key path for signature calculation. This can also be set as the ACI_PRIVATE_KEY environment variable.",
				Optional:    true,
			},
			"cert_name": schema.StringAttribute{
				Description: "Certificate name for the User in Cisco ACI. This can also be set as the ACI_CERT_NAME environment variable.",
				Optional:    true,
			},
			"proxy_url": schema.StringAttribute{
				Description: "Proxy Server URL with port number. This can also be set as the ACI_PROXY_URL environment variable.",
				Optional:    true,
			},
			"proxy_creds": schema.StringAttribute{
				Description: "Proxy server credentials in the form of username:password. This can also be set as the ACI_PROXY_CREDS environment variable.",
				Optional:    true,
			},
			// temporary used schema attributes until muxing is removed and the correct schema attributes can be used which are commented out below
			"insecure": schema.StringAttribute{
				Description: "Allow insecure HTTPS client. This can also be set as the ACI_INSECURE environment variable. Defaults to `true`.",
				Optional:    true,
			},
			"validate_relation_dn": schema.StringAttribute{
				Description: "Flag to validate if a object with entered relation Dn exists in the APIC. Defaults to `true`.",
				Optional:    true,
			},
			"retries": schema.StringAttribute{
				Description: "Number of retries for REST API calls. This can also be set as the ACI_RETRIES environment variable. Defaults to `2`.",
				Optional:    true,
			},
			// "insecure": schema.BoolAttribute{
			// 	Description: "Allow insecure HTTPS client. This can also be set as the ACI_INSECURE environment variable. Defaults to `true`.",
			// 	Optional:    true,
			// },
			// "validate_relation_dn": schema.BoolAttribute{
			// 	Description: "Flag to validate if a object with entered relation Dn exists in the APIC. Defaults to `true`.",
			// 	Optional:    true,
			// },
			// "retries": schema.Int64Attribute{
			// 	Description: "Number of retries for REST API calls. This can also be set as the ACI_RETRIES environment variable. Defaults to `2`.",
			// 	Optional:    true,
			// 	Validators: []validator.Int64{
			// 		int64validator.Between(0, 10),
			// 	},
			// },
			"annotation": schema.StringAttribute{
				Description: "Global annotation for the provider. This can also be set as the ACI_ANNOTATION environment variable.",
				Optional:    true,
			},
		},
	}
}

func (p *AciProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	var data AciProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	username := getStringAttribute(data.Username, "ACI_USERNAME")
	password := getStringAttribute(data.Password, "ACI_PASSWORD")
	url := getStringAttribute(data.URL, "ACI_URL")
	privateKey := getStringAttribute(data.PrivateKey, "ACI_PRIVATE_KEY")
	certName := getStringAttribute(data.Certname, "ACI_CERT_NAME")
	proxyUrl := getStringAttribute(data.ProxyUrl, "ACI_PROXY_URL")
	proxyCreds := getStringAttribute(data.ProxyCreds, "ACI_PROXY_CREDS")
	isInsecure := stringToBool(resp, "insecure", getStringAttribute(data.IsInsecure, "ACI_INSECURE"), true)
	validateRelationDn := stringToBool(resp, "insecure", getStringAttribute(data.ValidateRelationDn, "ACI_VAL_REL_DN"), true)
	maxRetries := stringToInt(resp, "retries", getStringAttribute(data.MaxRetries, "ACI_RETRIES"), 2)
	setGlobalAnnotation(data.Annotation, "ACI_ANNOTATION")

	if username == "" {
		resp.Diagnostics.AddError(
			"Username not provided",
			"Username must be provided for the ACI provider",
		)
	}

	if password == "" && (privateKey == "" || certName == "") {
		resp.Diagnostics.AddError(
			"Authentication details not provided",
			"Either 'password' OR 'private_key' and 'cert_name' must be provided for the ACI provider",
		)
	}

	if url == "" {
		resp.Diagnostics.AddError(
			"Url not provided",
			"Url must be provided for the ACI provider",
		)
	} else if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		resp.Diagnostics.AddError(
			"Incorrect url prefix",
			fmt.Sprintf("Url '%s' must start with 'http://' or 'https://'", url),
		)
	}

	// temporary conditional until muxing is removed and the correct retries schema attribute is used
	if maxRetries < 0 || maxRetries > 9 {
		resp.Diagnostics.AddError(
			"Incorrect retry amount",
			fmt.Sprintf("Retries must be between 0 and 9 inclusive, got: %d", maxRetries),
		)
	}

	var aciClient *client.Client
	if password != "" {
		aciClient = client.GetClient(url, username, client.Password(password), client.Insecure(isInsecure), client.ProxyUrl(proxyUrl), client.ProxyCreds(proxyCreds), client.ValidateRelationDn(validateRelationDn), client.MaxRetries(maxRetries))
	} else {
		aciClient = client.GetClient(url, username, client.PrivateKey(privateKey), client.AdminCert(certName), client.Insecure(isInsecure), client.ProxyUrl(proxyUrl), client.ProxyCreds(proxyCreds), client.ValidateRelationDn(validateRelationDn), client.MaxRetries(maxRetries))
	}

	resp.DataSourceData = aciClient
	resp.ResourceData = aciClient
}

func (p *AciProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewFvEpIpTagResource,
		NewFvEpMacTagResource,
		NewL3extConsLblResource,
		NewL3extRsRedistributePolResource,
		NewMgmtInstPResource,
		NewMgmtRsOoBConsResource,
		NewMgmtSubnetResource,
		NewPimRouteMapEntryResource,
		NewPimRouteMapPolResource,
		NewTagAnnotationResource,
		NewVzOOBBrCPResource,
		NewAciRestManagedResource,
	}
}

func (p *AciProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewFvEpIpTagDataSource,
		NewFvEpMacTagDataSource,
		NewL3extConsLblDataSource,
		NewL3extRsRedistributePolDataSource,
		NewMgmtInstPDataSource,
		NewMgmtRsOoBConsDataSource,
		NewMgmtSubnetDataSource,
		NewPimRouteMapEntryDataSource,
		NewPimRouteMapPolDataSource,
		NewTagAnnotationDataSource,
		NewVzOOBBrCPDataSource,
		NewAciRestManagedDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &AciProvider{
			version: version,
		}
	}
}

func setGlobalAnnotation(attribute basetypes.StringValue, envKey string) {

	if attribute.IsNull() {
		attributeValue, found := os.LookupEnv(envKey)
		if found {
			globalAnnotation = attributeValue
		} else {
			globalAnnotation = "orchestrator:terraform"
		}
	} else {
		globalAnnotation = attribute.ValueString()
	}
}

func getStringAttribute(attribute basetypes.StringValue, envKey string) string {

	if attribute.IsNull() {
		return os.Getenv(envKey)
	}
	return attribute.ValueString()

}

// Placeholder for future use when correct type can be used since muxing has been removed
func getBoolAttribute(resp *provider.ConfigureResponse, attribute basetypes.BoolValue, envKey string, defaultValue bool) bool {

	if attribute.IsNull() {
		envValue := os.Getenv(envKey)
		if envValue == "" {
			return defaultValue
		}
		boolValue, err := strconv.ParseBool(envValue)
		if err != nil {
			resp.Diagnostics.AddError(
				fmt.Sprintf("Invalid input '%s'", envValue),
				fmt.Sprintf("A boolean value must be provided for %s", envKey),
			)
		}
		return boolValue
	}
	return attribute.ValueBool()

}

// Placeholder for future use when correct type can be used since muxing has been removed
func getIntAttribute(resp *provider.ConfigureResponse, attribute basetypes.Int64Value, envKey string, defaultValue int) int {

	if attribute.IsNull() {
		envValue := os.Getenv(envKey)
		if envValue == "" {
			return defaultValue
		}
		intValue, err := strconv.Atoi(envValue)
		if err != nil {
			resp.Diagnostics.AddError(
				fmt.Sprintf("Invalid input '%s'", envValue),
				fmt.Sprintf("A integer value must be provided for %s", envKey),
			)
		}
		return intValue
	}
	return int(attribute.ValueInt64())
}

// Temporary function so correct type can be used untill since muxing has been removed
func stringToBool(resp *provider.ConfigureResponse, attributeName, stringValue string, boolValue bool) bool {
	var err error
	if stringValue != "" {
		boolValue, err = strconv.ParseBool(stringValue)
		if err != nil {
			resp.Diagnostics.AddError(
				fmt.Sprintf("Invalid input '%s'", stringValue),
				fmt.Sprintf("A boolean value must be provided for the %s attribute", attributeName),
			)
		}
	}
	return boolValue
}

// Temporary function so correct type can be used untill since muxing has been removed
func stringToInt(resp *provider.ConfigureResponse, attributeName, stringValue string, intValue int) int {
	var err error
	if stringValue != "" {
		intValue, err = strconv.Atoi(stringValue)
		if err != nil {
			resp.Diagnostics.AddError(
				fmt.Sprintf("Invalid input '%s'", stringValue),
				fmt.Sprintf("A integer must be provided for the %s attribute", attributeName),
			)
		}
	}
	return intValue
}
