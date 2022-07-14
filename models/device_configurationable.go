package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationable 
type DeviceConfigurationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]DeviceConfigurationAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDeviceManagementApplicabilityRuleDeviceMode()(DeviceConfigurationsable)
    GetDeviceManagementApplicabilityRuleOsEdition()(DeviceConfigurationsable)
    GetDeviceManagementApplicabilityRuleOsVersion()(DeviceConfigurationsable)
    GetDeviceSettingStateSummaries()([]SettingStateDeviceSummaryable)
    GetDeviceStatuses()([]DeviceConfigurationDeviceStatusable)
    GetDeviceStatusOverview()(DeviceConfigurationsable)
    GetDisplayName()(*string)
    GetGroupAssignments()([]DeviceConfigurationGroupAssignmentable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRoleScopeTagIds()([]string)
    GetSupportsScopeTags()(*bool)
    GetUserStatuses()([]DeviceConfigurationUserStatusable)
    GetUserStatusOverview()(DeviceConfigurationsable)
    GetVersion()(*int32)
    SetAssignments(value []DeviceConfigurationAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDeviceManagementApplicabilityRuleDeviceMode(value DeviceConfigurationsable)()
    SetDeviceManagementApplicabilityRuleOsEdition(value DeviceConfigurationsable)()
    SetDeviceManagementApplicabilityRuleOsVersion(value DeviceConfigurationsable)()
    SetDeviceSettingStateSummaries(value []SettingStateDeviceSummaryable)()
    SetDeviceStatuses(value []DeviceConfigurationDeviceStatusable)()
    SetDeviceStatusOverview(value DeviceConfigurationsable)()
    SetDisplayName(value *string)()
    SetGroupAssignments(value []DeviceConfigurationGroupAssignmentable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRoleScopeTagIds(value []string)()
    SetSupportsScopeTags(value *bool)()
    SetUserStatuses(value []DeviceConfigurationUserStatusable)()
    SetUserStatusOverview(value DeviceConfigurationsable)()
    SetVersion(value *int32)()
}
