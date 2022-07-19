package policies

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i154ced0e8b3656059650a4bd061d0aa6f15bc39d29746142cde157c94c8ce663 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/identitysecuritydefaultsenforcementpolicy"
    i3baef4935efa79b1ea53b0f8a4d12e8e33271d34f1b7fe4e5e3deefd698fb11d "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/adminconsentrequestpolicy"
    i4235b128855e4ef4e736e93e75306be2329042607217b14308bae073614f6b91 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/permissiongrantpolicies"
    idf9e8b9f9da041893e6b6cb24dd45425a8696e3da093aef925b748ba4b81e936 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/authorizationpolicy"
)

// PoliciesRequestBuilder builds and executes requests for operations under \policies
type PoliciesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AdminConsentRequestPolicy the adminConsentRequestPolicy property
func (m *PoliciesRequestBuilder) AdminConsentRequestPolicy()(*i3baef4935efa79b1ea53b0f8a4d12e8e33271d34f1b7fe4e5e3deefd698fb11d.AdminConsentRequestPolicyRequestBuilder) {
    return i3baef4935efa79b1ea53b0f8a4d12e8e33271d34f1b7fe4e5e3deefd698fb11d.NewAdminConsentRequestPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthorizationPolicy the authorizationPolicy property
func (m *PoliciesRequestBuilder) AuthorizationPolicy()(*idf9e8b9f9da041893e6b6cb24dd45425a8696e3da093aef925b748ba4b81e936.AuthorizationPolicyRequestBuilder) {
    return idf9e8b9f9da041893e6b6cb24dd45425a8696e3da093aef925b748ba4b81e936.NewAuthorizationPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPoliciesRequestBuilderInternal instantiates a new PoliciesRequestBuilder and sets the default values.
func NewPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PoliciesRequestBuilder) {
    m := &PoliciesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPoliciesRequestBuilder instantiates a new PoliciesRequestBuilder and sets the default values.
func NewPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// IdentitySecurityDefaultsEnforcementPolicy the identitySecurityDefaultsEnforcementPolicy property
func (m *PoliciesRequestBuilder) IdentitySecurityDefaultsEnforcementPolicy()(*i154ced0e8b3656059650a4bd061d0aa6f15bc39d29746142cde157c94c8ce663.IdentitySecurityDefaultsEnforcementPolicyRequestBuilder) {
    return i154ced0e8b3656059650a4bd061d0aa6f15bc39d29746142cde157c94c8ce663.NewIdentitySecurityDefaultsEnforcementPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantPolicies the permissionGrantPolicies property
func (m *PoliciesRequestBuilder) PermissionGrantPolicies()(*i4235b128855e4ef4e736e93e75306be2329042607217b14308bae073614f6b91.PermissionGrantPoliciesRequestBuilder) {
    return i4235b128855e4ef4e736e93e75306be2329042607217b14308bae073614f6b91.NewPermissionGrantPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
