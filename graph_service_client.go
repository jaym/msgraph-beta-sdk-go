package msgraphbetasdkgo

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i09893664b20e7c846b2bc7aaaf1cd7f554ed3d2c00ac11336bea4c3c3d859e09 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement"
    i310cda3e9f244aa61f9c9c78de433773f341b91c4a2310b8991671fe773be16e "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement"
    i5f794bd004f1cc95da776bcb1947ffabf97b71aae1f5c9f15255b24451e2929b "github.com/microsoftgraph/msgraph-beta-sdk-go/policies"
    i714cbeb65962cb4d3e58007792fa4832d175c04614ba3aa7efb22871aea885bf "github.com/microsoftgraph/msgraph-beta-sdk-go/settings"
    i761e9f0dec20dbf36c7fd626d107fb81ef94cafa7369422d2b2af143ffa16184 "github.com/microsoftgraph/msgraph-beta-sdk-go/security"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i9e22e53c888822daa9264a72be4d11f335e3170e9198aae4bba758214e319857 "github.com/microsoftgraph/msgraph-beta-sdk-go/domains"
    icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a "github.com/microsoftgraph/msgraph-beta-sdk-go/users"
    ie1b2fd35e4b1f7cbc7bd808e462c966c4ec16a274923b50216bdd8a2ae0a3129 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications"
    ie58948149bb028757a64f16376df00cc5a99ad93e0d57affa0ac24ff6d096aaf "github.com/microsoftgraph/msgraph-beta-sdk-go/organization"
    ia82ac058caa6740adb23fd487affca202963b9786c4890c80eaec27b6d7ced24 "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item"
    id9f276010196e2d79e634ec622333b5c53dc0fbc407a6c9aa27ca92d4f388ed3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item"
)

// GraphServiceClient the main entry point of the SDK, exposes the configuration and the fluent API.
type GraphServiceClient struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Applications the applications property
func (m *GraphServiceClient) Applications()(*ie1b2fd35e4b1f7cbc7bd808e462c966c4ec16a274923b50216bdd8a2ae0a3129.ApplicationsRequestBuilder) {
    return ie1b2fd35e4b1f7cbc7bd808e462c966c4ec16a274923b50216bdd8a2ae0a3129.NewApplicationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGraphServiceClient instantiates a new GraphServiceClient and sets the default values.
func NewGraphServiceClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GraphServiceClient) {
    m := &GraphServiceClient{
    }
    m.pathParameters = make(map[string]string);
    m.urlTemplate = "{+baseurl}";
    m.requestAdapter = requestAdapter;
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    if m.requestAdapter.GetBaseUrl() == "" {
        m.requestAdapter.SetBaseUrl("https://graph.microsoft.com/beta")
    }
    return m
}
// DeviceManagement the deviceManagement property
func (m *GraphServiceClient) DeviceManagement()(*i09893664b20e7c846b2bc7aaaf1cd7f554ed3d2c00ac11336bea4c3c3d859e09.DeviceManagementRequestBuilder) {
    return i09893664b20e7c846b2bc7aaaf1cd7f554ed3d2c00ac11336bea4c3c3d859e09.NewDeviceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Domains the domains property
func (m *GraphServiceClient) Domains()(*i9e22e53c888822daa9264a72be4d11f335e3170e9198aae4bba758214e319857.DomainsRequestBuilder) {
    return i9e22e53c888822daa9264a72be4d11f335e3170e9198aae4bba758214e319857.NewDomainsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DomainsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.domains.item collection
func (m *GraphServiceClient) DomainsById(id string)(*ia82ac058caa6740adb23fd487affca202963b9786c4890c80eaec27b6d7ced24.DomainItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["domain%2Did"] = id
    }
    return ia82ac058caa6740adb23fd487affca202963b9786c4890c80eaec27b6d7ced24.NewDomainItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Organization the organization property
func (m *GraphServiceClient) Organization()(*ie58948149bb028757a64f16376df00cc5a99ad93e0d57affa0ac24ff6d096aaf.OrganizationRequestBuilder) {
    return ie58948149bb028757a64f16376df00cc5a99ad93e0d57affa0ac24ff6d096aaf.NewOrganizationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Policies the policies property
func (m *GraphServiceClient) Policies()(*i5f794bd004f1cc95da776bcb1947ffabf97b71aae1f5c9f15255b24451e2929b.PoliciesRequestBuilder) {
    return i5f794bd004f1cc95da776bcb1947ffabf97b71aae1f5c9f15255b24451e2929b.NewPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleManagement the roleManagement property
func (m *GraphServiceClient) RoleManagement()(*i310cda3e9f244aa61f9c9c78de433773f341b91c4a2310b8991671fe773be16e.RoleManagementRequestBuilder) {
    return i310cda3e9f244aa61f9c9c78de433773f341b91c4a2310b8991671fe773be16e.NewRoleManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Security the security property
func (m *GraphServiceClient) Security()(*i761e9f0dec20dbf36c7fd626d107fb81ef94cafa7369422d2b2af143ffa16184.SecurityRequestBuilder) {
    return i761e9f0dec20dbf36c7fd626d107fb81ef94cafa7369422d2b2af143ffa16184.NewSecurityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Settings the settings property
func (m *GraphServiceClient) Settings()(*i714cbeb65962cb4d3e58007792fa4832d175c04614ba3aa7efb22871aea885bf.SettingsRequestBuilder) {
    return i714cbeb65962cb4d3e58007792fa4832d175c04614ba3aa7efb22871aea885bf.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Users the users property
func (m *GraphServiceClient) Users()(*icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a.UsersRequestBuilder) {
    return icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.users.item collection
func (m *GraphServiceClient) UsersById(id string)(*id9f276010196e2d79e634ec622333b5c53dc0fbc407a6c9aa27ca92d4f388ed3.UserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return id9f276010196e2d79e634ec622333b5c53dc0fbc407a6c9aa27ca92d4f388ed3.NewUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
