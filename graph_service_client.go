package msgraphbetasdkgo

import (
    i4b2d55bb354e12d758e74ae9fe8ba9df2480de0e2c09162d79d9b8c5986c28b9 "msgraphbetasdkgo/policies"
    i6455bbcebecc3db18fad5f32313e7969578fe6c4bb85f7af86268614eac8461f "msgraphbetasdkgo/security"
    i782388dedc41a62acb7543de41ce1036e876276bd42b682e3032a7f3b53f4887 "msgraphbetasdkgo/organization"
    i79ef632ebe4800278ff66763d66d68550464a260065b1342136af919099a8478 "msgraphbetasdkgo/settings"
    i9c34a78adbfe3782d6e5ccd507fee3aebe9b22d9acd2cd87aa2581854bc3d6c3 "msgraphbetasdkgo/domains"
    ib1a75c74a5d481c8b8846d6b99ecb600db5e8063b65d888e4f551d66baa4d129 "msgraphbetasdkgo/devicemanagement"
    ibb12c00680dda05aab78922c0e76be04fe8206dc7cfb7e5d194c5f5b9798a8b9 "msgraphbetasdkgo/rolemanagement"
    ic6e6215f2555e2bb2c9363dbe950a8cbb1db0a5b4605ffc718f0dd47dc47f22e "msgraphbetasdkgo/users"
    id715ab657bcb7663bf224d754abbfe797891ada76b074972274e49079a3db48a "msgraphbetasdkgo/applications"
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i48cb45b3658128273922cc3b73ad57746451a113b41496c3315a00b079cca49b "msgraphbetasdkgo/domains/item"
    i6a64ca7801c71f2061d2a5373c6837be393e418538fb89a7a860556f56f27dd6 "msgraphbetasdkgo/users/item"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
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
func (m *GraphServiceClient) Applications()(*id715ab657bcb7663bf224d754abbfe797891ada76b074972274e49079a3db48a.ApplicationsRequestBuilder) {
    return id715ab657bcb7663bf224d754abbfe797891ada76b074972274e49079a3db48a.NewApplicationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *GraphServiceClient) DeviceManagement()(*ib1a75c74a5d481c8b8846d6b99ecb600db5e8063b65d888e4f551d66baa4d129.DeviceManagementRequestBuilder) {
    return ib1a75c74a5d481c8b8846d6b99ecb600db5e8063b65d888e4f551d66baa4d129.NewDeviceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Domains the domains property
func (m *GraphServiceClient) Domains()(*i9c34a78adbfe3782d6e5ccd507fee3aebe9b22d9acd2cd87aa2581854bc3d6c3.DomainsRequestBuilder) {
    return i9c34a78adbfe3782d6e5ccd507fee3aebe9b22d9acd2cd87aa2581854bc3d6c3.NewDomainsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DomainsById gets an item from the msgraphbetasdkgo.domains.item collection
func (m *GraphServiceClient) DomainsById(id string)(*i48cb45b3658128273922cc3b73ad57746451a113b41496c3315a00b079cca49b.DomainItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["domain%2Did"] = id
    }
    return i48cb45b3658128273922cc3b73ad57746451a113b41496c3315a00b079cca49b.NewDomainItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Organization the organization property
func (m *GraphServiceClient) Organization()(*i782388dedc41a62acb7543de41ce1036e876276bd42b682e3032a7f3b53f4887.OrganizationRequestBuilder) {
    return i782388dedc41a62acb7543de41ce1036e876276bd42b682e3032a7f3b53f4887.NewOrganizationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Policies the policies property
func (m *GraphServiceClient) Policies()(*i4b2d55bb354e12d758e74ae9fe8ba9df2480de0e2c09162d79d9b8c5986c28b9.PoliciesRequestBuilder) {
    return i4b2d55bb354e12d758e74ae9fe8ba9df2480de0e2c09162d79d9b8c5986c28b9.NewPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleManagement the roleManagement property
func (m *GraphServiceClient) RoleManagement()(*ibb12c00680dda05aab78922c0e76be04fe8206dc7cfb7e5d194c5f5b9798a8b9.RoleManagementRequestBuilder) {
    return ibb12c00680dda05aab78922c0e76be04fe8206dc7cfb7e5d194c5f5b9798a8b9.NewRoleManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Security the security property
func (m *GraphServiceClient) Security()(*i6455bbcebecc3db18fad5f32313e7969578fe6c4bb85f7af86268614eac8461f.SecurityRequestBuilder) {
    return i6455bbcebecc3db18fad5f32313e7969578fe6c4bb85f7af86268614eac8461f.NewSecurityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Settings the settings property
func (m *GraphServiceClient) Settings()(*i79ef632ebe4800278ff66763d66d68550464a260065b1342136af919099a8478.SettingsRequestBuilder) {
    return i79ef632ebe4800278ff66763d66d68550464a260065b1342136af919099a8478.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Users the users property
func (m *GraphServiceClient) Users()(*ic6e6215f2555e2bb2c9363dbe950a8cbb1db0a5b4605ffc718f0dd47dc47f22e.UsersRequestBuilder) {
    return ic6e6215f2555e2bb2c9363dbe950a8cbb1db0a5b4605ffc718f0dd47dc47f22e.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsersById gets an item from the msgraphbetasdkgo.users.item collection
func (m *GraphServiceClient) UsersById(id string)(*i6a64ca7801c71f2061d2a5373c6837be393e418538fb89a7a860556f56f27dd6.UserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return i6a64ca7801c71f2061d2a5373c6837be393e418538fb89a7a860556f56f27dd6.NewUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
