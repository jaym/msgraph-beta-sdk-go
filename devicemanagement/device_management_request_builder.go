package devicemanagement

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i3dbc32ed6e7f2382197bf0faf9babee9e91f243b2d9fbce1ea9d7d7df535349b "msgraphbetasdkgo/devicemanagement/devicecompliancepolicies"
    ida6191aa5e91e1cefce390636108fe8ab55e2bbb3d65f9dfc31c46fe420e017c "msgraphbetasdkgo/devicemanagement/deviceconfigurations"
)

// DeviceManagementRequestBuilder builds and executes requests for operations under \deviceManagement
type DeviceManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewDeviceManagementRequestBuilderInternal instantiates a new DeviceManagementRequestBuilder and sets the default values.
func NewDeviceManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementRequestBuilder) {
    m := &DeviceManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementRequestBuilder instantiates a new DeviceManagementRequestBuilder and sets the default values.
func NewDeviceManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// DeviceCompliancePolicies the deviceCompliancePolicies property
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicies()(*i3dbc32ed6e7f2382197bf0faf9babee9e91f243b2d9fbce1ea9d7d7df535349b.DeviceCompliancePoliciesRequestBuilder) {
    return i3dbc32ed6e7f2382197bf0faf9babee9e91f243b2d9fbce1ea9d7d7df535349b.NewDeviceCompliancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurations the deviceConfigurations property
func (m *DeviceManagementRequestBuilder) DeviceConfigurations()(*ida6191aa5e91e1cefce390636108fe8ab55e2bbb3d65f9dfc31c46fe420e017c.DeviceConfigurationsRequestBuilder) {
    return ida6191aa5e91e1cefce390636108fe8ab55e2bbb3d65f9dfc31c46fe420e017c.NewDeviceConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
