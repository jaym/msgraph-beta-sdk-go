package devicemanagement

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i06897cd84592b50e6e76477755f51f0fc302e5664c3a9d4325a123deb6d935d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations"
    i358b7a34135f775242bdf4c2e565d2d6445c815cd929fbf919e431687e5d64e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies"
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
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicies()(*i358b7a34135f775242bdf4c2e565d2d6445c815cd929fbf919e431687e5d64e0.DeviceCompliancePoliciesRequestBuilder) {
    return i358b7a34135f775242bdf4c2e565d2d6445c815cd929fbf919e431687e5d64e0.NewDeviceCompliancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurations the deviceConfigurations property
func (m *DeviceManagementRequestBuilder) DeviceConfigurations()(*i06897cd84592b50e6e76477755f51f0fc302e5664c3a9d4325a123deb6d935d4.DeviceConfigurationsRequestBuilder) {
    return i06897cd84592b50e6e76477755f51f0fc302e5664c3a9d4325a123deb6d935d4.NewDeviceConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
