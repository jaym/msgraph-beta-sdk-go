package directory

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ic09a099122154ece987c7f85636fcc11aac7d11f4704eb7ea18e15d32cb118be "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roledefinitions"
    if3967cc171135f8ae89f635c0a197e13e64649a0043e1dcc96293c3c4f5ebbbe "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignments"
)

// DirectoryRequestBuilder builds and executes requests for operations under \roleManagement\directory
type DirectoryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewDirectoryRequestBuilderInternal instantiates a new DirectoryRequestBuilder and sets the default values.
func NewDirectoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRequestBuilder) {
    m := &DirectoryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/directory";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryRequestBuilder instantiates a new DirectoryRequestBuilder and sets the default values.
func NewDirectoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRequestBuilderInternal(urlParams, requestAdapter)
}
// RoleAssignments the roleAssignments property
func (m *DirectoryRequestBuilder) RoleAssignments()(*if3967cc171135f8ae89f635c0a197e13e64649a0043e1dcc96293c3c4f5ebbbe.RoleAssignmentsRequestBuilder) {
    return if3967cc171135f8ae89f635c0a197e13e64649a0043e1dcc96293c3c4f5ebbbe.NewRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinitions the roleDefinitions property
func (m *DirectoryRequestBuilder) RoleDefinitions()(*ic09a099122154ece987c7f85636fcc11aac7d11f4704eb7ea18e15d32cb118be.RoleDefinitionsRequestBuilder) {
    return ic09a099122154ece987c7f85636fcc11aac7d11f4704eb7ea18e15d32cb118be.NewRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
