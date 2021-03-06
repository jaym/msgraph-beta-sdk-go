package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceable 
type ManagedDeviceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAadRegistered()(*bool)
    GetActivationLockBypassCode()(*string)
    GetAndroidSecurityPatchLevel()(*string)
    GetAssignmentFilterEvaluationStatusDetails()([]AssignmentFilterEvaluationStatusDetailsable)
    GetAutopilotEnrolled()(*bool)
    GetAzureActiveDirectoryDeviceId()(*string)
    GetAzureADDeviceId()(*string)
    GetAzureADRegistered()(*bool)
    GetBootstrapTokenEscrowed()(*bool)
    GetChassisType()(*ChassisType)
    GetChromeOSDeviceInfo()([]ChromeOSDevicePropertyable)
    GetCloudPcRemoteActionResults()([]CloudPcRemoteActionResultable)
    GetComplianceGracePeriodExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetComplianceState()(*ComplianceState)
    GetConfigurationManagerClientEnabledFeatures()(ConfigurationManagerClientEnabledFeaturesable)
    GetConfigurationManagerClientHealthState()(ConfigurationManagerClientHealthStateable)
    GetConfigurationManagerClientInformation()(ConfigurationManagerClientInformationable)
    GetDetectedApps()([]DetectedAppable)
    GetDeviceActionResults()([]DeviceActionResultable)
    GetDeviceCategory()(DeviceCategoryable)
    GetDeviceCategoryDisplayName()(*string)
    GetDeviceCompliancePolicyStates()([]DeviceCompliancePolicyStateable)
    GetDeviceConfigurationStates()([]DeviceConfigurationStateable)
    GetDeviceEnrollmentType()(*DeviceEnrollmentType)
    GetDeviceFirmwareConfigurationInterfaceManaged()(*bool)
    GetDeviceHealthAttestationState()(DeviceHealthAttestationStateable)
    GetDeviceName()(*string)
    GetDeviceRegistrationState()(*DeviceRegistrationState)
    GetDeviceType()(*DeviceType)
    GetEasActivated()(*bool)
    GetEasActivationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEasDeviceId()(*string)
    GetEmailAddress()(*string)
    GetEnrolledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEnrollmentProfileName()(*string)
    GetEthernetMacAddress()(*string)
    GetExchangeAccessState()(*DeviceManagementExchangeAccessState)
    GetExchangeAccessStateReason()(*DeviceManagementExchangeAccessStateReason)
    GetExchangeLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFreeStorageSpaceInBytes()(*int64)
    GetHardwareInformation()(HardwareInformationable)
    GetIccid()(*string)
    GetImei()(*string)
    GetIsEncrypted()(*bool)
    GetIsSupervised()(*bool)
    GetJailBroken()(*string)
    GetJoinType()(*JoinType)
    GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLogCollectionRequests()([]DeviceLogCollectionResponseable)
    GetLostModeState()(*LostModeState)
    GetManagedDeviceMobileAppConfigurationStates()([]ManagedDeviceMobileAppConfigurationStateable)
    GetManagedDeviceName()(*string)
    GetManagedDeviceOwnerType()(*ManagedDeviceOwnerType)
    GetManagementAgent()(*ManagementAgentType)
    GetManagementCertificateExpirationDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagementFeatures()(*ManagedDeviceManagementFeatures)
    GetManagementState()(*ManagementState)
    GetManufacturer()(*string)
    GetMeid()(*string)
    GetModel()(*string)
    GetNotes()(*string)
    GetOperatingSystem()(*string)
    GetOsVersion()(*string)
    GetOwnerType()(*OwnerType)
    GetPartnerReportedThreatState()(*ManagedDevicePartnerReportedHealthState)
    GetPhoneNumber()(*string)
    GetPhysicalMemoryInBytes()(*int64)
    GetPreferMdmOverGroupPolicyAppliedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetProcessorArchitecture()(*ManagedDeviceArchitecture)
    GetRemoteAssistanceSessionErrorDetails()(*string)
    GetRemoteAssistanceSessionUrl()(*string)
    GetRequireUserEnrollmentApproval()(*bool)
    GetRetireAfterDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRoleScopeTagIds()([]string)
    GetSecurityBaselineStates()([]SecurityBaselineStateable)
    GetSerialNumber()(*string)
    GetSkuFamily()(*string)
    GetSkuNumber()(*int32)
    GetSpecificationVersion()(*string)
    GetSubscriberCarrier()(*string)
    GetTotalStorageSpaceInBytes()(*int64)
    GetUdid()(*string)
    GetUserDisplayName()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    GetUsers()([]Userable)
    GetUsersLoggedOn()([]LoggedOnUserable)
    GetWiFiMacAddress()(*string)
    GetWindowsActiveMalwareCount()(*int32)
    GetWindowsProtectionState()(WindowsProtectionStateable)
    GetWindowsRemediatedMalwareCount()(*int32)
    SetAadRegistered(value *bool)()
    SetActivationLockBypassCode(value *string)()
    SetAndroidSecurityPatchLevel(value *string)()
    SetAssignmentFilterEvaluationStatusDetails(value []AssignmentFilterEvaluationStatusDetailsable)()
    SetAutopilotEnrolled(value *bool)()
    SetAzureActiveDirectoryDeviceId(value *string)()
    SetAzureADDeviceId(value *string)()
    SetAzureADRegistered(value *bool)()
    SetBootstrapTokenEscrowed(value *bool)()
    SetChassisType(value *ChassisType)()
    SetChromeOSDeviceInfo(value []ChromeOSDevicePropertyable)()
    SetCloudPcRemoteActionResults(value []CloudPcRemoteActionResultable)()
    SetComplianceGracePeriodExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetComplianceState(value *ComplianceState)()
    SetConfigurationManagerClientEnabledFeatures(value ConfigurationManagerClientEnabledFeaturesable)()
    SetConfigurationManagerClientHealthState(value ConfigurationManagerClientHealthStateable)()
    SetConfigurationManagerClientInformation(value ConfigurationManagerClientInformationable)()
    SetDetectedApps(value []DetectedAppable)()
    SetDeviceActionResults(value []DeviceActionResultable)()
    SetDeviceCategory(value DeviceCategoryable)()
    SetDeviceCategoryDisplayName(value *string)()
    SetDeviceCompliancePolicyStates(value []DeviceCompliancePolicyStateable)()
    SetDeviceConfigurationStates(value []DeviceConfigurationStateable)()
    SetDeviceEnrollmentType(value *DeviceEnrollmentType)()
    SetDeviceFirmwareConfigurationInterfaceManaged(value *bool)()
    SetDeviceHealthAttestationState(value DeviceHealthAttestationStateable)()
    SetDeviceName(value *string)()
    SetDeviceRegistrationState(value *DeviceRegistrationState)()
    SetDeviceType(value *DeviceType)()
    SetEasActivated(value *bool)()
    SetEasActivationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEasDeviceId(value *string)()
    SetEmailAddress(value *string)()
    SetEnrolledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEnrollmentProfileName(value *string)()
    SetEthernetMacAddress(value *string)()
    SetExchangeAccessState(value *DeviceManagementExchangeAccessState)()
    SetExchangeAccessStateReason(value *DeviceManagementExchangeAccessStateReason)()
    SetExchangeLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFreeStorageSpaceInBytes(value *int64)()
    SetHardwareInformation(value HardwareInformationable)()
    SetIccid(value *string)()
    SetImei(value *string)()
    SetIsEncrypted(value *bool)()
    SetIsSupervised(value *bool)()
    SetJailBroken(value *string)()
    SetJoinType(value *JoinType)()
    SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLogCollectionRequests(value []DeviceLogCollectionResponseable)()
    SetLostModeState(value *LostModeState)()
    SetManagedDeviceMobileAppConfigurationStates(value []ManagedDeviceMobileAppConfigurationStateable)()
    SetManagedDeviceName(value *string)()
    SetManagedDeviceOwnerType(value *ManagedDeviceOwnerType)()
    SetManagementAgent(value *ManagementAgentType)()
    SetManagementCertificateExpirationDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagementFeatures(value *ManagedDeviceManagementFeatures)()
    SetManagementState(value *ManagementState)()
    SetManufacturer(value *string)()
    SetMeid(value *string)()
    SetModel(value *string)()
    SetNotes(value *string)()
    SetOperatingSystem(value *string)()
    SetOsVersion(value *string)()
    SetOwnerType(value *OwnerType)()
    SetPartnerReportedThreatState(value *ManagedDevicePartnerReportedHealthState)()
    SetPhoneNumber(value *string)()
    SetPhysicalMemoryInBytes(value *int64)()
    SetPreferMdmOverGroupPolicyAppliedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetProcessorArchitecture(value *ManagedDeviceArchitecture)()
    SetRemoteAssistanceSessionErrorDetails(value *string)()
    SetRemoteAssistanceSessionUrl(value *string)()
    SetRequireUserEnrollmentApproval(value *bool)()
    SetRetireAfterDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRoleScopeTagIds(value []string)()
    SetSecurityBaselineStates(value []SecurityBaselineStateable)()
    SetSerialNumber(value *string)()
    SetSkuFamily(value *string)()
    SetSkuNumber(value *int32)()
    SetSpecificationVersion(value *string)()
    SetSubscriberCarrier(value *string)()
    SetTotalStorageSpaceInBytes(value *int64)()
    SetUdid(value *string)()
    SetUserDisplayName(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
    SetUsers(value []Userable)()
    SetUsersLoggedOn(value []LoggedOnUserable)()
    SetWiFiMacAddress(value *string)()
    SetWindowsActiveMalwareCount(value *int32)()
    SetWindowsProtectionState(value WindowsProtectionStateable)()
    SetWindowsRemediatedMalwareCount(value *int32)()
}
