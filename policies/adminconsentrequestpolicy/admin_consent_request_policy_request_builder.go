package adminconsentrequestpolicy

import (
    ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837 "msgraphbetasdkgo/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i6da0336faa817434736692f8c8cc36680c0b59a923d4499bb6c797b96dff5879 "msgraphbetasdkgo/models/odataerrors"
)

// AdminConsentRequestPolicyRequestBuilder builds and executes requests for operations under \policies\adminConsentRequestPolicy
type AdminConsentRequestPolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AdminConsentRequestPolicyRequestBuilderGetQueryParameters the policy by which consent requests are created and managed for the entire tenant.
type AdminConsentRequestPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AdminConsentRequestPolicyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminConsentRequestPolicyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdminConsentRequestPolicyRequestBuilderGetQueryParameters
}
// NewAdminConsentRequestPolicyRequestBuilderInternal instantiates a new AdminConsentRequestPolicyRequestBuilder and sets the default values.
func NewAdminConsentRequestPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminConsentRequestPolicyRequestBuilder) {
    m := &AdminConsentRequestPolicyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/adminConsentRequestPolicy{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAdminConsentRequestPolicyRequestBuilder instantiates a new AdminConsentRequestPolicyRequestBuilder and sets the default values.
func NewAdminConsentRequestPolicyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminConsentRequestPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdminConsentRequestPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the policy by which consent requests are created and managed for the entire tenant.
func (m *AdminConsentRequestPolicyRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the policy by which consent requests are created and managed for the entire tenant.
func (m *AdminConsentRequestPolicyRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AdminConsentRequestPolicyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get the policy by which consent requests are created and managed for the entire tenant.
func (m *AdminConsentRequestPolicyRequestBuilder) Get()(ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837.AdminConsentRequestPolicyable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the policy by which consent requests are created and managed for the entire tenant.
func (m *AdminConsentRequestPolicyRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *AdminConsentRequestPolicyRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837.AdminConsentRequestPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i6da0336faa817434736692f8c8cc36680c0b59a923d4499bb6c797b96dff5879.CreateODataErrorFromDiscriminatorValue,
        "5XX": i6da0336faa817434736692f8c8cc36680c0b59a923d4499bb6c797b96dff5879.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837.CreateAdminConsentRequestPolicyFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837.AdminConsentRequestPolicyable), nil
}
