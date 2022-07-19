package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Entity 
type Entity struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The id property
    id *string
    // The OdataType property
    odataType *string
}
// NewEntity instantiates a new entity and sets the default values.
func NewEntity()(*Entity) {
    m := &Entity{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.entity";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEntityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEntityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.accessReviewInstance":
                        return NewAccessReviewInstance(), nil
                    case "#microsoft.graph.accessReviewInstanceDecisionItem":
                        return NewAccessReviewInstanceDecisionItem(), nil
                    case "#microsoft.graph.accessReviewReviewer":
                        return NewAccessReviewReviewer(), nil
                    case "#microsoft.graph.accessReviewScheduleDefinition":
                        return NewAccessReviewScheduleDefinition(), nil
                    case "#microsoft.graph.accessReviewStage":
                        return NewAccessReviewStage(), nil
                    case "#microsoft.graph.activityHistoryItem":
                        return NewActivityHistoryItem(), nil
                    case "#microsoft.graph.activityStatistics":
                        return NewActivityStatistics(), nil
                    case "#microsoft.graph.adminConsentRequestPolicy":
                        return NewAdminConsentRequestPolicy(), nil
                    case "#microsoft.graph.agreementAcceptance":
                        return NewAgreementAcceptance(), nil
                    case "#microsoft.graph.appConsentRequest":
                        return NewAppConsentRequest(), nil
                    case "#microsoft.graph.application":
                        return NewApplication(), nil
                    case "#microsoft.graph.appLogCollectionRequest":
                        return NewAppLogCollectionRequest(), nil
                    case "#microsoft.graph.appManagementPolicy":
                        return NewAppManagementPolicy(), nil
                    case "#microsoft.graph.appRoleAssignment":
                        return NewAppRoleAssignment(), nil
                    case "#microsoft.graph.approval":
                        return NewApproval(), nil
                    case "#microsoft.graph.approvalStep":
                        return NewApprovalStep(), nil
                    case "#microsoft.graph.appScope":
                        return NewAppScope(), nil
                    case "#microsoft.graph.assignmentFilterEvaluationStatusDetails":
                        return NewAssignmentFilterEvaluationStatusDetails(), nil
                    case "#microsoft.graph.associatedTeamInfo":
                        return NewAssociatedTeamInfo(), nil
                    case "#microsoft.graph.attachment":
                        return NewAttachment(), nil
                    case "#microsoft.graph.attachmentBase":
                        return NewAttachmentBase(), nil
                    case "#microsoft.graph.attachmentSession":
                        return NewAttachmentSession(), nil
                    case "#microsoft.graph.attendanceRecord":
                        return NewAttendanceRecord(), nil
                    case "#microsoft.graph.authentication":
                        return NewAuthentication(), nil
                    case "#microsoft.graph.authenticationMethod":
                        return NewAuthenticationMethod(), nil
                    case "#microsoft.graph.authorizationPolicy":
                        return NewAuthorizationPolicy(), nil
                    case "#microsoft.graph.baseItem":
                        return NewBaseItem(), nil
                    case "#microsoft.graph.baseItemVersion":
                        return NewBaseItemVersion(), nil
                    case "#microsoft.graph.baseTask":
                        return NewBaseTask(), nil
                    case "#microsoft.graph.baseTaskList":
                        return NewBaseTaskList(), nil
                    case "#microsoft.graph.bitlocker":
                        return NewBitlocker(), nil
                    case "#microsoft.graph.bitlockerRecoveryKey":
                        return NewBitlockerRecoveryKey(), nil
                    case "#microsoft.graph.calendar":
                        return NewCalendar(), nil
                    case "#microsoft.graph.calendarGroup":
                        return NewCalendarGroup(), nil
                    case "#microsoft.graph.calendarPermission":
                        return NewCalendarPermission(), nil
                    case "#microsoft.graph.certificateBasedAuthConfiguration":
                        return NewCertificateBasedAuthConfiguration(), nil
                    case "#microsoft.graph.changeTrackedEntity":
                        return NewChangeTrackedEntity(), nil
                    case "#microsoft.graph.channel":
                        return NewChannel(), nil
                    case "#microsoft.graph.chat":
                        return NewChat(), nil
                    case "#microsoft.graph.chatMessage":
                        return NewChatMessage(), nil
                    case "#microsoft.graph.chatMessageHostedContent":
                        return NewChatMessageHostedContent(), nil
                    case "#microsoft.graph.chatMessageInfo":
                        return NewChatMessageInfo(), nil
                    case "#microsoft.graph.checklistItem":
                        return NewChecklistItem(), nil
                    case "#microsoft.graph.cloudPC":
                        return NewCloudPC(), nil
                    case "#microsoft.graph.columnDefinition":
                        return NewColumnDefinition(), nil
                    case "#microsoft.graph.columnLink":
                        return NewColumnLink(), nil
                    case "#microsoft.graph.command":
                        return NewCommand(), nil
                    case "#microsoft.graph.connector":
                        return NewConnector(), nil
                    case "#microsoft.graph.connectorGroup":
                        return NewConnectorGroup(), nil
                    case "#microsoft.graph.contact":
                        return NewContact(), nil
                    case "#microsoft.graph.contactFolder":
                        return NewContactFolder(), nil
                    case "#microsoft.graph.contactMergeSuggestions":
                        return NewContactMergeSuggestions(), nil
                    case "#microsoft.graph.contentType":
                        return NewContentType(), nil
                    case "#microsoft.graph.conversation":
                        return NewConversation(), nil
                    case "#microsoft.graph.conversationMember":
                        return NewConversationMember(), nil
                    case "#microsoft.graph.conversationThread":
                        return NewConversationThread(), nil
                    case "#microsoft.graph.dataLossPreventionPolicy":
                        return NewDataLossPreventionPolicy(), nil
                    case "#microsoft.graph.defaultUserRoleOverride":
                        return NewDefaultUserRoleOverride(), nil
                    case "#microsoft.graph.detectedApp":
                        return NewDetectedApp(), nil
                    case "#microsoft.graph.device":
                        return NewDevice(), nil
                    case "#microsoft.graph.deviceCategory":
                        return NewDeviceCategory(), nil
                    case "#microsoft.graph.deviceComplianceActionItem":
                        return NewDeviceComplianceActionItem(), nil
                    case "#microsoft.graph.deviceComplianceDeviceOverview":
                        return NewDeviceComplianceDeviceOverview(), nil
                    case "#microsoft.graph.deviceComplianceDeviceStatus":
                        return NewDeviceComplianceDeviceStatus(), nil
                    case "#microsoft.graph.deviceCompliancePolicy":
                        return NewDeviceCompliancePolicy(), nil
                    case "#microsoft.graph.deviceCompliancePolicyAssignment":
                        return NewDeviceCompliancePolicyAssignment(), nil
                    case "#microsoft.graph.deviceCompliancePolicyState":
                        return NewDeviceCompliancePolicyState(), nil
                    case "#microsoft.graph.deviceComplianceScheduledActionForRule":
                        return NewDeviceComplianceScheduledActionForRule(), nil
                    case "#microsoft.graph.deviceComplianceUserOverview":
                        return NewDeviceComplianceUserOverview(), nil
                    case "#microsoft.graph.deviceComplianceUserStatus":
                        return NewDeviceComplianceUserStatus(), nil
                    case "#microsoft.graph.deviceConfiguration":
                        return NewDeviceConfiguration(), nil
                    case "#microsoft.graph.deviceConfigurationAssignment":
                        return NewDeviceConfigurationAssignment(), nil
                    case "#microsoft.graph.deviceConfigurationDeviceOverview":
                        return NewDeviceConfigurationDeviceOverview(), nil
                    case "#microsoft.graph.deviceConfigurationDeviceStatus":
                        return NewDeviceConfigurationDeviceStatus(), nil
                    case "#microsoft.graph.deviceConfigurationGroupAssignment":
                        return NewDeviceConfigurationGroupAssignment(), nil
                    case "#microsoft.graph.deviceConfigurationState":
                        return NewDeviceConfigurationState(), nil
                    case "#microsoft.graph.deviceConfigurationUserOverview":
                        return NewDeviceConfigurationUserOverview(), nil
                    case "#microsoft.graph.deviceConfigurationUserStatus":
                        return NewDeviceConfigurationUserStatus(), nil
                    case "#microsoft.graph.deviceEnrollmentConfiguration":
                        return NewDeviceEnrollmentConfiguration(), nil
                    case "#microsoft.graph.deviceLogCollectionResponse":
                        return NewDeviceLogCollectionResponse(), nil
                    case "#microsoft.graph.deviceManagementTroubleshootingEvent":
                        return NewDeviceManagementTroubleshootingEvent(), nil
                    case "#microsoft.graph.directoryDefinition":
                        return NewDirectoryDefinition(), nil
                    case "#microsoft.graph.directoryObject":
                        return NewDirectoryObject(), nil
                    case "#microsoft.graph.directorySetting":
                        return NewDirectorySetting(), nil
                    case "#microsoft.graph.documentSetVersion":
                        return NewDocumentSetVersion(), nil
                    case "#microsoft.graph.domain":
                        return NewDomain(), nil
                    case "#microsoft.graph.domainDnsRecord":
                        return NewDomainDnsRecord(), nil
                    case "#microsoft.graph.drive":
                        return NewDrive(), nil
                    case "#microsoft.graph.driveItem":
                        return NewDriveItem(), nil
                    case "#microsoft.graph.driveItemVersion":
                        return NewDriveItemVersion(), nil
                    case "#microsoft.graph.educationalActivity":
                        return NewEducationalActivity(), nil
                    case "#microsoft.graph.emailAuthenticationMethod":
                        return NewEmailAuthenticationMethod(), nil
                    case "#microsoft.graph.endpoint":
                        return NewEndpoint(), nil
                    case "#microsoft.graph.enrollmentConfigurationAssignment":
                        return NewEnrollmentConfigurationAssignment(), nil
                    case "#microsoft.graph.event":
                        return NewEvent(), nil
                    case "#microsoft.graph.extension":
                        return NewExtension(), nil
                    case "#microsoft.graph.extensionProperty":
                        return NewExtensionProperty(), nil
                    case "#microsoft.graph.federatedIdentityCredential":
                        return NewFederatedIdentityCredential(), nil
                    case "#microsoft.graph.fido2AuthenticationMethod":
                        return NewFido2AuthenticationMethod(), nil
                    case "#microsoft.graph.fieldValueSet":
                        return NewFieldValueSet(), nil
                    case "#microsoft.graph.governanceInsight":
                        return NewGovernanceInsight(), nil
                    case "#microsoft.graph.group":
                        return NewGroup(), nil
                    case "#microsoft.graph.groupLifecyclePolicy":
                        return NewGroupLifecyclePolicy(), nil
                    case "#microsoft.graph.homeRealmDiscoveryPolicy":
                        return NewHomeRealmDiscoveryPolicy(), nil
                    case "#microsoft.graph.identityProviderBase":
                        return NewIdentityProviderBase(), nil
                    case "#microsoft.graph.identitySecurityDefaultsEnforcementPolicy":
                        return NewIdentitySecurityDefaultsEnforcementPolicy(), nil
                    case "#microsoft.graph.inferenceClassification":
                        return NewInferenceClassification(), nil
                    case "#microsoft.graph.inferenceClassificationOverride":
                        return NewInferenceClassificationOverride(), nil
                    case "#microsoft.graph.informationProtection":
                        return NewInformationProtection(), nil
                    case "#microsoft.graph.informationProtectionLabel":
                        return NewInformationProtectionLabel(), nil
                    case "#microsoft.graph.informationProtectionPolicy":
                        return NewInformationProtectionPolicy(), nil
                    case "#microsoft.graph.insightsSettings":
                        return NewInsightsSettings(), nil
                    case "#microsoft.graph.internalDomainFederation":
                        return NewInternalDomainFederation(), nil
                    case "#microsoft.graph.itemActivity":
                        return NewItemActivity(), nil
                    case "#microsoft.graph.itemActivityOLD":
                        return NewItemActivityOLD(), nil
                    case "#microsoft.graph.itemActivityStat":
                        return NewItemActivityStat(), nil
                    case "#microsoft.graph.itemAddress":
                        return NewItemAddress(), nil
                    case "#microsoft.graph.itemAnalytics":
                        return NewItemAnalytics(), nil
                    case "#microsoft.graph.itemEmail":
                        return NewItemEmail(), nil
                    case "#microsoft.graph.itemFacet":
                        return NewItemFacet(), nil
                    case "#microsoft.graph.itemInsights":
                        return NewItemInsights(), nil
                    case "#microsoft.graph.itemPatent":
                        return NewItemPatent(), nil
                    case "#microsoft.graph.itemPhone":
                        return NewItemPhone(), nil
                    case "#microsoft.graph.itemPublication":
                        return NewItemPublication(), nil
                    case "#microsoft.graph.languageProficiency":
                        return NewLanguageProficiency(), nil
                    case "#microsoft.graph.licenseDetails":
                        return NewLicenseDetails(), nil
                    case "#microsoft.graph.linkedResource":
                        return NewLinkedResource(), nil
                    case "#microsoft.graph.linkedResource_v2":
                        return NewLinkedResource_v2(), nil
                    case "#microsoft.graph.list":
                        return NewList(), nil
                    case "#microsoft.graph.listItem":
                        return NewListItem(), nil
                    case "#microsoft.graph.listItemVersion":
                        return NewListItemVersion(), nil
                    case "#microsoft.graph.longRunningOperation":
                        return NewLongRunningOperation(), nil
                    case "#microsoft.graph.mailFolder":
                        return NewMailFolder(), nil
                    case "#microsoft.graph.managedAppOperation":
                        return NewManagedAppOperation(), nil
                    case "#microsoft.graph.managedAppPolicy":
                        return NewManagedAppPolicy(), nil
                    case "#microsoft.graph.managedAppRegistration":
                        return NewManagedAppRegistration(), nil
                    case "#microsoft.graph.managedDevice":
                        return NewManagedDevice(), nil
                    case "#microsoft.graph.managedDeviceMobileAppConfigurationState":
                        return NewManagedDeviceMobileAppConfigurationState(), nil
                    case "#microsoft.graph.meetingAttendanceReport":
                        return NewMeetingAttendanceReport(), nil
                    case "#microsoft.graph.meetingRegistrantBase":
                        return NewMeetingRegistrantBase(), nil
                    case "#microsoft.graph.meetingRegistration":
                        return NewMeetingRegistration(), nil
                    case "#microsoft.graph.meetingRegistrationBase":
                        return NewMeetingRegistrationBase(), nil
                    case "#microsoft.graph.meetingRegistrationQuestion":
                        return NewMeetingRegistrationQuestion(), nil
                    case "#microsoft.graph.mention":
                        return NewMention(), nil
                    case "#microsoft.graph.message":
                        return NewMessage(), nil
                    case "#microsoft.graph.messageRule":
                        return NewMessageRule(), nil
                    case "#microsoft.graph.microsoftApplicationDataAccessSettings":
                        return NewMicrosoftApplicationDataAccessSettings(), nil
                    case "#microsoft.graph.microsoftAuthenticatorAuthenticationMethod":
                        return NewMicrosoftAuthenticatorAuthenticationMethod(), nil
                    case "#microsoft.graph.mobileAppIntentAndState":
                        return NewMobileAppIntentAndState(), nil
                    case "#microsoft.graph.mobileAppTroubleshootingEvent":
                        return NewMobileAppTroubleshootingEvent(), nil
                    case "#microsoft.graph.multiValueLegacyExtendedProperty":
                        return NewMultiValueLegacyExtendedProperty(), nil
                    case "#microsoft.graph.notebook":
                        return NewNotebook(), nil
                    case "#microsoft.graph.notification":
                        return NewNotification(), nil
                    case "#microsoft.graph.oAuth2PermissionGrant":
                        return NewOAuth2PermissionGrant(), nil
                    case "#microsoft.graph.offerShiftRequest":
                        return NewOfferShiftRequest(), nil
                    case "#microsoft.graph.officeGraphInsights":
                        return NewOfficeGraphInsights(), nil
                    case "#microsoft.graph.onenote":
                        return NewOnenote(), nil
                    case "#microsoft.graph.onenoteEntityBaseModel":
                        return NewOnenoteEntityBaseModel(), nil
                    case "#microsoft.graph.onenoteEntityHierarchyModel":
                        return NewOnenoteEntityHierarchyModel(), nil
                    case "#microsoft.graph.onenoteEntitySchemaObjectModel":
                        return NewOnenoteEntitySchemaObjectModel(), nil
                    case "#microsoft.graph.onenoteOperation":
                        return NewOnenoteOperation(), nil
                    case "#microsoft.graph.onenotePage":
                        return NewOnenotePage(), nil
                    case "#microsoft.graph.onenoteResource":
                        return NewOnenoteResource(), nil
                    case "#microsoft.graph.onenoteSection":
                        return NewOnenoteSection(), nil
                    case "#microsoft.graph.onlineMeeting":
                        return NewOnlineMeeting(), nil
                    case "#microsoft.graph.openShift":
                        return NewOpenShift(), nil
                    case "#microsoft.graph.openShiftChangeRequest":
                        return NewOpenShiftChangeRequest(), nil
                    case "#microsoft.graph.operation":
                        return NewOperation(), nil
                    case "#microsoft.graph.organization":
                        return NewOrganization(), nil
                    case "#microsoft.graph.organizationalBranding":
                        return NewOrganizationalBranding(), nil
                    case "#microsoft.graph.organizationalBrandingLocalization":
                        return NewOrganizationalBrandingLocalization(), nil
                    case "#microsoft.graph.organizationalBrandingProperties":
                        return NewOrganizationalBrandingProperties(), nil
                    case "#microsoft.graph.organizationSettings":
                        return NewOrganizationSettings(), nil
                    case "#microsoft.graph.outlookCategory":
                        return NewOutlookCategory(), nil
                    case "#microsoft.graph.outlookItem":
                        return NewOutlookItem(), nil
                    case "#microsoft.graph.outlookTask":
                        return NewOutlookTask(), nil
                    case "#microsoft.graph.outlookTaskFolder":
                        return NewOutlookTaskFolder(), nil
                    case "#microsoft.graph.outlookTaskGroup":
                        return NewOutlookTaskGroup(), nil
                    case "#microsoft.graph.outlookUser":
                        return NewOutlookUser(), nil
                    case "#microsoft.graph.passwordAuthenticationMethod":
                        return NewPasswordAuthenticationMethod(), nil
                    case "#microsoft.graph.passwordlessMicrosoftAuthenticatorAuthenticationMethod":
                        return NewPasswordlessMicrosoftAuthenticatorAuthenticationMethod(), nil
                    case "#microsoft.graph.payloadResponse":
                        return NewPayloadResponse(), nil
                    case "#microsoft.graph.permission":
                        return NewPermission(), nil
                    case "#microsoft.graph.permissionGrantConditionSet":
                        return NewPermissionGrantConditionSet(), nil
                    case "#microsoft.graph.permissionGrantPolicy":
                        return NewPermissionGrantPolicy(), nil
                    case "#microsoft.graph.person":
                        return NewPerson(), nil
                    case "#microsoft.graph.personAnnotation":
                        return NewPersonAnnotation(), nil
                    case "#microsoft.graph.personAnnualEvent":
                        return NewPersonAnnualEvent(), nil
                    case "#microsoft.graph.personAward":
                        return NewPersonAward(), nil
                    case "#microsoft.graph.personCertification":
                        return NewPersonCertification(), nil
                    case "#microsoft.graph.personInterest":
                        return NewPersonInterest(), nil
                    case "#microsoft.graph.personName":
                        return NewPersonName(), nil
                    case "#microsoft.graph.personWebsite":
                        return NewPersonWebsite(), nil
                    case "#microsoft.graph.phoneAuthenticationMethod":
                        return NewPhoneAuthenticationMethod(), nil
                    case "#microsoft.graph.pinnedChatMessageInfo":
                        return NewPinnedChatMessageInfo(), nil
                    case "#microsoft.graph.plannerAssignedToTaskBoardTaskFormat":
                        return NewPlannerAssignedToTaskBoardTaskFormat(), nil
                    case "#microsoft.graph.plannerBucket":
                        return NewPlannerBucket(), nil
                    case "#microsoft.graph.plannerBucketTaskBoardTaskFormat":
                        return NewPlannerBucketTaskBoardTaskFormat(), nil
                    case "#microsoft.graph.plannerDelta":
                        return NewPlannerDelta(), nil
                    case "#microsoft.graph.plannerGroup":
                        return NewPlannerGroup(), nil
                    case "#microsoft.graph.plannerPlan":
                        return NewPlannerPlan(), nil
                    case "#microsoft.graph.plannerPlanDetails":
                        return NewPlannerPlanDetails(), nil
                    case "#microsoft.graph.plannerProgressTaskBoardTaskFormat":
                        return NewPlannerProgressTaskBoardTaskFormat(), nil
                    case "#microsoft.graph.plannerTask":
                        return NewPlannerTask(), nil
                    case "#microsoft.graph.plannerTaskDetails":
                        return NewPlannerTaskDetails(), nil
                    case "#microsoft.graph.plannerUser":
                        return NewPlannerUser(), nil
                    case "#microsoft.graph.policyBase":
                        return NewPolicyBase(), nil
                    case "#microsoft.graph.post":
                        return NewPost(), nil
                    case "#microsoft.graph.presence":
                        return NewPresence(), nil
                    case "#microsoft.graph.printConnector":
                        return NewPrintConnector(), nil
                    case "#microsoft.graph.printDocument":
                        return NewPrintDocument(), nil
                    case "#microsoft.graph.printer":
                        return NewPrinter(), nil
                    case "#microsoft.graph.printerBase":
                        return NewPrinterBase(), nil
                    case "#microsoft.graph.printerShare":
                        return NewPrinterShare(), nil
                    case "#microsoft.graph.printJob":
                        return NewPrintJob(), nil
                    case "#microsoft.graph.printTask":
                        return NewPrintTask(), nil
                    case "#microsoft.graph.printTaskDefinition":
                        return NewPrintTaskDefinition(), nil
                    case "#microsoft.graph.printTaskTrigger":
                        return NewPrintTaskTrigger(), nil
                    case "#microsoft.graph.profile":
                        return NewProfile(), nil
                    case "#microsoft.graph.profileCardProperty":
                        return NewProfileCardProperty(), nil
                    case "#microsoft.graph.profilePhoto":
                        return NewProfilePhoto(), nil
                    case "#microsoft.graph.projectParticipation":
                        return NewProjectParticipation(), nil
                    case "#microsoft.graph.regionalAndLanguageSettings":
                        return NewRegionalAndLanguageSettings(), nil
                    case "#microsoft.graph.request":
                        return NewRequest(), nil
                    case "#microsoft.graph.resourceSpecificPermissionGrant":
                        return NewResourceSpecificPermissionGrant(), nil
                    case "#microsoft.graph.richLongRunningOperation":
                        return NewRichLongRunningOperation(), nil
                    case "#microsoft.graph.samlOrWsFedProvider":
                        return NewSamlOrWsFedProvider(), nil
                    case "#microsoft.graph.schedule":
                        return NewSchedule(), nil
                    case "#microsoft.graph.scheduleChangeRequest":
                        return NewScheduleChangeRequest(), nil
                    case "#microsoft.graph.schedulingGroup":
                        return NewSchedulingGroup(), nil
                    case "#microsoft.graph.scopedRoleMembership":
                        return NewScopedRoleMembership(), nil
                    case "#microsoft.graph.sectionGroup":
                        return NewSectionGroup(), nil
                    case "#microsoft.graph.secureScore":
                        return NewSecureScore(), nil
                    case "#microsoft.graph.securityBaselineSettingState":
                        return NewSecurityBaselineSettingState(), nil
                    case "#microsoft.graph.securityBaselineState":
                        return NewSecurityBaselineState(), nil
                    case "#microsoft.graph.sensitivityLabel":
                        return NewSensitivityLabel(), nil
                    case "#microsoft.graph.sensitivityPolicySettings":
                        return NewSensitivityPolicySettings(), nil
                    case "#microsoft.graph.settingStateDeviceSummary":
                        return NewSettingStateDeviceSummary(), nil
                    case "#microsoft.graph.sharedEmailDomainInvitation":
                        return NewSharedEmailDomainInvitation(), nil
                    case "#microsoft.graph.sharedInsight":
                        return NewSharedInsight(), nil
                    case "#microsoft.graph.sharedWithChannelTeamInfo":
                        return NewSharedWithChannelTeamInfo(), nil
                    case "#microsoft.graph.shift":
                        return NewShift(), nil
                    case "#microsoft.graph.shiftPreferences":
                        return NewShiftPreferences(), nil
                    case "#microsoft.graph.singleValueLegacyExtendedProperty":
                        return NewSingleValueLegacyExtendedProperty(), nil
                    case "#microsoft.graph.site":
                        return NewSite(), nil
                    case "#microsoft.graph.sitePage":
                        return NewSitePage(), nil
                    case "#microsoft.graph.skillProficiency":
                        return NewSkillProficiency(), nil
                    case "#microsoft.graph.softwareOathAuthenticationMethod":
                        return NewSoftwareOathAuthenticationMethod(), nil
                    case "#microsoft.graph.stsPolicy":
                        return NewStsPolicy(), nil
                    case "#microsoft.graph.subscription":
                        return NewSubscription(), nil
                    case "#microsoft.graph.swapShiftsChangeRequest":
                        return NewSwapShiftsChangeRequest(), nil
                    case "#microsoft.graph.synchronization":
                        return NewSynchronization(), nil
                    case "#microsoft.graph.synchronizationJob":
                        return NewSynchronizationJob(), nil
                    case "#microsoft.graph.synchronizationSchema":
                        return NewSynchronizationSchema(), nil
                    case "#microsoft.graph.synchronizationTemplate":
                        return NewSynchronizationTemplate(), nil
                    case "#microsoft.graph.tasks":
                        return NewTasks(), nil
                    case "#microsoft.graph.team":
                        return NewTeam(), nil
                    case "#microsoft.graph.teamInfo":
                        return NewTeamInfo(), nil
                    case "#microsoft.graph.teamsApp":
                        return NewTeamsApp(), nil
                    case "#microsoft.graph.teamsAppDefinition":
                        return NewTeamsAppDefinition(), nil
                    case "#microsoft.graph.teamsAppIcon":
                        return NewTeamsAppIcon(), nil
                    case "#microsoft.graph.teamsAppInstallation":
                        return NewTeamsAppInstallation(), nil
                    case "#microsoft.graph.teamsAsyncOperation":
                        return NewTeamsAsyncOperation(), nil
                    case "#microsoft.graph.teamsTab":
                        return NewTeamsTab(), nil
                    case "#microsoft.graph.teamsTemplate":
                        return NewTeamsTemplate(), nil
                    case "#microsoft.graph.teamworkBot":
                        return NewTeamworkBot(), nil
                    case "#microsoft.graph.teamworkHostedContent":
                        return NewTeamworkHostedContent(), nil
                    case "#microsoft.graph.teamworkTag":
                        return NewTeamworkTag(), nil
                    case "#microsoft.graph.teamworkTagMember":
                        return NewTeamworkTagMember(), nil
                    case "#microsoft.graph.temporaryAccessPassAuthenticationMethod":
                        return NewTemporaryAccessPassAuthenticationMethod(), nil
                    case "#microsoft.graph.threatAssessmentRequest":
                        return NewThreatAssessmentRequest(), nil
                    case "#microsoft.graph.threatAssessmentResult":
                        return NewThreatAssessmentResult(), nil
                    case "#microsoft.graph.thumbnailSet":
                        return NewThumbnailSet(), nil
                    case "#microsoft.graph.timeCard":
                        return NewTimeCard(), nil
                    case "#microsoft.graph.timeOff":
                        return NewTimeOff(), nil
                    case "#microsoft.graph.timeOffReason":
                        return NewTimeOffReason(), nil
                    case "#microsoft.graph.timeOffRequest":
                        return NewTimeOffRequest(), nil
                    case "#microsoft.graph.todo":
                        return NewTodo(), nil
                    case "#microsoft.graph.todoTask":
                        return NewTodoTask(), nil
                    case "#microsoft.graph.todoTaskList":
                        return NewTodoTaskList(), nil
                    case "#microsoft.graph.tokenIssuancePolicy":
                        return NewTokenIssuancePolicy(), nil
                    case "#microsoft.graph.tokenLifetimePolicy":
                        return NewTokenLifetimePolicy(), nil
                    case "#microsoft.graph.trending":
                        return NewTrending(), nil
                    case "#microsoft.graph.unifiedRoleAssignment":
                        return NewUnifiedRoleAssignment(), nil
                    case "#microsoft.graph.unifiedRoleDefinition":
                        return NewUnifiedRoleDefinition(), nil
                    case "#microsoft.graph.usageRight":
                        return NewUsageRight(), nil
                    case "#microsoft.graph.usedInsight":
                        return NewUsedInsight(), nil
                    case "#microsoft.graph.user":
                        return NewUser(), nil
                    case "#microsoft.graph.userAccountInformation":
                        return NewUserAccountInformation(), nil
                    case "#microsoft.graph.userActivity":
                        return NewUserActivity(), nil
                    case "#microsoft.graph.userAnalytics":
                        return NewUserAnalytics(), nil
                    case "#microsoft.graph.userConfiguration":
                        return NewUserConfiguration(), nil
                    case "#microsoft.graph.userConsentRequest":
                        return NewUserConsentRequest(), nil
                    case "#microsoft.graph.userInsightsSettings":
                        return NewUserInsightsSettings(), nil
                    case "#microsoft.graph.userScopeTeamsAppInstallation":
                        return NewUserScopeTeamsAppInstallation(), nil
                    case "#microsoft.graph.userSettings":
                        return NewUserSettings(), nil
                    case "#microsoft.graph.userTeamwork":
                        return NewUserTeamwork(), nil
                    case "#microsoft.graph.webAccount":
                        return NewWebAccount(), nil
                    case "#microsoft.graph.windowsDeviceMalwareState":
                        return NewWindowsDeviceMalwareState(), nil
                    case "#microsoft.graph.windowsHelloForBusinessAuthenticationMethod":
                        return NewWindowsHelloForBusinessAuthenticationMethod(), nil
                    case "#microsoft.graph.windowsInformationProtectionDeviceRegistration":
                        return NewWindowsInformationProtectionDeviceRegistration(), nil
                    case "#microsoft.graph.windowsProtectionState":
                        return NewWindowsProtectionState(), nil
                    case "#microsoft.graph.workbook":
                        return NewWorkbook(), nil
                    case "#microsoft.graph.workbookApplication":
                        return NewWorkbookApplication(), nil
                    case "#microsoft.graph.workbookChart":
                        return NewWorkbookChart(), nil
                    case "#microsoft.graph.workbookChartAreaFormat":
                        return NewWorkbookChartAreaFormat(), nil
                    case "#microsoft.graph.workbookChartAxes":
                        return NewWorkbookChartAxes(), nil
                    case "#microsoft.graph.workbookChartAxis":
                        return NewWorkbookChartAxis(), nil
                    case "#microsoft.graph.workbookChartAxisFormat":
                        return NewWorkbookChartAxisFormat(), nil
                    case "#microsoft.graph.workbookChartAxisTitle":
                        return NewWorkbookChartAxisTitle(), nil
                    case "#microsoft.graph.workbookChartAxisTitleFormat":
                        return NewWorkbookChartAxisTitleFormat(), nil
                    case "#microsoft.graph.workbookChartDataLabelFormat":
                        return NewWorkbookChartDataLabelFormat(), nil
                    case "#microsoft.graph.workbookChartDataLabels":
                        return NewWorkbookChartDataLabels(), nil
                    case "#microsoft.graph.workbookChartFill":
                        return NewWorkbookChartFill(), nil
                    case "#microsoft.graph.workbookChartFont":
                        return NewWorkbookChartFont(), nil
                    case "#microsoft.graph.workbookChartGridlines":
                        return NewWorkbookChartGridlines(), nil
                    case "#microsoft.graph.workbookChartGridlinesFormat":
                        return NewWorkbookChartGridlinesFormat(), nil
                    case "#microsoft.graph.workbookChartLegend":
                        return NewWorkbookChartLegend(), nil
                    case "#microsoft.graph.workbookChartLegendFormat":
                        return NewWorkbookChartLegendFormat(), nil
                    case "#microsoft.graph.workbookChartLineFormat":
                        return NewWorkbookChartLineFormat(), nil
                    case "#microsoft.graph.workbookChartPoint":
                        return NewWorkbookChartPoint(), nil
                    case "#microsoft.graph.workbookChartPointFormat":
                        return NewWorkbookChartPointFormat(), nil
                    case "#microsoft.graph.workbookChartSeries":
                        return NewWorkbookChartSeries(), nil
                    case "#microsoft.graph.workbookChartSeriesFormat":
                        return NewWorkbookChartSeriesFormat(), nil
                    case "#microsoft.graph.workbookChartTitle":
                        return NewWorkbookChartTitle(), nil
                    case "#microsoft.graph.workbookChartTitleFormat":
                        return NewWorkbookChartTitleFormat(), nil
                    case "#microsoft.graph.workbookComment":
                        return NewWorkbookComment(), nil
                    case "#microsoft.graph.workbookCommentReply":
                        return NewWorkbookCommentReply(), nil
                    case "#microsoft.graph.workbookFilter":
                        return NewWorkbookFilter(), nil
                    case "#microsoft.graph.workbookFunctions":
                        return NewWorkbookFunctions(), nil
                    case "#microsoft.graph.workbookNamedItem":
                        return NewWorkbookNamedItem(), nil
                    case "#microsoft.graph.workbookOperation":
                        return NewWorkbookOperation(), nil
                    case "#microsoft.graph.workbookPivotTable":
                        return NewWorkbookPivotTable(), nil
                    case "#microsoft.graph.workbookTable":
                        return NewWorkbookTable(), nil
                    case "#microsoft.graph.workbookTableColumn":
                        return NewWorkbookTableColumn(), nil
                    case "#microsoft.graph.workbookTableRow":
                        return NewWorkbookTableRow(), nil
                    case "#microsoft.graph.workbookTableSort":
                        return NewWorkbookTableSort(), nil
                    case "#microsoft.graph.workbookWorksheet":
                        return NewWorkbookWorksheet(), nil
                    case "#microsoft.graph.workbookWorksheetProtection":
                        return NewWorkbookWorksheetProtection(), nil
                    case "#microsoft.graph.workPosition":
                        return NewWorkPosition(), nil
                }
            }
        }
    }
    return NewEntity(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Entity) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Entity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
func (m *Entity) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Entity) GetOdataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.odataType
    }
}
// Serialize serializes information the current object
func (m *Entity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Entity) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetId sets the id property value. The id property
func (m *Entity) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Entity) SetOdataType(value *string)() {
    if m != nil {
        m.odataType = value
    }
}
