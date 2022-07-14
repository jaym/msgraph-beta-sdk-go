package policies

import (
    i21588b40caaddef52a034946c3674bc41513befe24f7c98ed96222418d4caa3f "msgraphbetasdkgo/policies/identitysecuritydefaultsenforcementpolicy"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i393f4dff1da5b4803faf521f41eb43a4d6f8df76b774b123243d28efb959e42d "msgraphbetasdkgo/policies/adminconsentrequestpolicy"
    i8e71945d343786d387efdbf97cb7746942f2578ebd3076e69a816c11c172bd7c "msgraphbetasdkgo/policies/permissiongrantpolicies"
    iaa7c639a2bbe5bec73cd8ac5783a8640d2a37f6ec32d95f58df8acafe74e64f6 "msgraphbetasdkgo/policies/authorizationpolicy"
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
func (m *PoliciesRequestBuilder) AdminConsentRequestPolicy()(*i393f4dff1da5b4803faf521f41eb43a4d6f8df76b774b123243d28efb959e42d.AdminConsentRequestPolicyRequestBuilder) {
    return i393f4dff1da5b4803faf521f41eb43a4d6f8df76b774b123243d28efb959e42d.NewAdminConsentRequestPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthorizationPolicy the authorizationPolicy property
func (m *PoliciesRequestBuilder) AuthorizationPolicy()(*iaa7c639a2bbe5bec73cd8ac5783a8640d2a37f6ec32d95f58df8acafe74e64f6.AuthorizationPolicyRequestBuilder) {
    return iaa7c639a2bbe5bec73cd8ac5783a8640d2a37f6ec32d95f58df8acafe74e64f6.NewAuthorizationPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *PoliciesRequestBuilder) IdentitySecurityDefaultsEnforcementPolicy()(*i21588b40caaddef52a034946c3674bc41513befe24f7c98ed96222418d4caa3f.IdentitySecurityDefaultsEnforcementPolicyRequestBuilder) {
    return i21588b40caaddef52a034946c3674bc41513befe24f7c98ed96222418d4caa3f.NewIdentitySecurityDefaultsEnforcementPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantPolicies the permissionGrantPolicies property
func (m *PoliciesRequestBuilder) PermissionGrantPolicies()(*i8e71945d343786d387efdbf97cb7746942f2578ebd3076e69a816c11c172bd7c.PermissionGrantPoliciesRequestBuilder) {
    return i8e71945d343786d387efdbf97cb7746942f2578ebd3076e69a816c11c172bd7c.NewPermissionGrantPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
