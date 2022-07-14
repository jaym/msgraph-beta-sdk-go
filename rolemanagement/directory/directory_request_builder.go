package directory

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib89059cc8ea0d6757a3540d223d0b309e8b0ea5cc5680375eb74ecadefd4f0f4 "msgraphbetasdkgo/rolemanagement/directory/roledefinitions"
    if4fb3a19bcac32bdd6dfdabc1c6da7811c8660566df2348f1708f4317bd5852d "msgraphbetasdkgo/rolemanagement/directory/roleassignments"
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
func (m *DirectoryRequestBuilder) RoleAssignments()(*if4fb3a19bcac32bdd6dfdabc1c6da7811c8660566df2348f1708f4317bd5852d.RoleAssignmentsRequestBuilder) {
    return if4fb3a19bcac32bdd6dfdabc1c6da7811c8660566df2348f1708f4317bd5852d.NewRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinitions the roleDefinitions property
func (m *DirectoryRequestBuilder) RoleDefinitions()(*ib89059cc8ea0d6757a3540d223d0b309e8b0ea5cc5680375eb74ecadefd4f0f4.RoleDefinitionsRequestBuilder) {
    return ib89059cc8ea0d6757a3540d223d0b309e8b0ea5cc5680375eb74ecadefd4f0f4.NewRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
