package rolemanagement

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i60fcea844834d04268ab75ad2b465502d1530b6893a3b95822880356df4b76c9 "msgraphbetasdkgo/rolemanagement/directory"
)

// RoleManagementRequestBuilder builds and executes requests for operations under \roleManagement
type RoleManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewRoleManagementRequestBuilderInternal instantiates a new RoleManagementRequestBuilder and sets the default values.
func NewRoleManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementRequestBuilder) {
    m := &RoleManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleManagementRequestBuilder instantiates a new RoleManagementRequestBuilder and sets the default values.
func NewRoleManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// Directory the directory property
func (m *RoleManagementRequestBuilder) Directory()(*i60fcea844834d04268ab75ad2b465502d1530b6893a3b95822880356df4b76c9.DirectoryRequestBuilder) {
    return i60fcea844834d04268ab75ad2b465502d1530b6893a3b95822880356df4b76c9.NewDirectoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
