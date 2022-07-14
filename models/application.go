package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Application 
type Application struct {
    DirectoryObject
    // Specifies settings for an application that implements a web API.
    api IdentitySecurityDefaultsEnforcementPolicyable
    // The unique identifier for the application that is assigned by Azure AD. Not nullable. Read-only.
    appId *string
    // The appManagementPolicy applied to this application.
    appManagementPolicies []AppManagementPolicyable
    // The collection of roles assigned to the application. With app role assignments, these roles can be assigned to users, groups, or service principals associated with other applications. Not nullable.
    appRoles []AppRoleable
    // Specifies the certification status of the application.
    certification IdentitySecurityDefaultsEnforcementPolicyable
    // The connectorGroup the application is using with Azure AD Application Proxy. Nullable.
    connectorGroup IdentitySecurityDefaultsEnforcementPolicyable
    // The date and time the application was registered. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.  Supports $filter (eq, ne, not, ge, le, in, and eq on null values) and $orderBy.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The createdOnBehalfOf property
    createdOnBehalfOf IdentitySecurityDefaultsEnforcementPolicyable
    // The default redirect URI. If specified and there is no explicit redirect URI in the sign-in request for SAML and OIDC flows, Azure AD sends the token to this redirect URI. Azure AD also sends the token to this default URI in SAML IdP-initiated single sign-on. The value must match one of the configured redirect URIs for the application.
    defaultRedirectUri *string
    // Free text field to provide a description of the application object to end users. The maximum allowed size is 1024 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
    description *string
    // Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, not).
    disabledByMicrosoftStatus *string
    // The display name for the application. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
    displayName *string
    // Read-only. Nullable. Supports $expand and $filter (eq when counting empty collections).
    extensionProperties []ExtensionPropertyable
    // Federated identities for applications. Supports $expand and $filter (eq when counting empty collections).
    federatedIdentityCredentials []FederatedIdentityCredentialable
    // Configures the groups claim issued in a user or OAuth 2.0 access token that the application expects. To set this attribute, use one of the following string values: None, SecurityGroup (for security groups and Azure AD roles), All (this gets all security groups, distribution groups, and Azure AD directory roles that the signed-in user is a member of).
    groupMembershipClaims *string
    // The homeRealmDiscoveryPolicies property
    homeRealmDiscoveryPolicies []HomeRealmDiscoveryPolicyable
    // Also known as App ID URI, this value is set when an application is used as a resource app. The identifierUris acts as the prefix for the scopes you'll reference in your API's code, and it must be globally unique. You can use the default value provided, which is in the form api://<application-client-id>, or specify a more readable URI like https://contoso.com/api. For more information on valid identifierUris patterns and best practices, see Azure AD application registration security best practices. Not nullable. Supports $filter (eq, ne, ge, le, startsWith).
    identifierUris []string
    // Basic profile information of the application, such as it's marketing, support, terms of service, and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more information, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, not, ge, le, and eq on null values).
    info IdentitySecurityDefaultsEnforcementPolicyable
    // Specifies whether this application supports device authentication without a user. The default is false.
    isDeviceOnlyAuthSupported *bool
    // Specifies the fallback application type as public client, such as an installed application running on a mobile device. The default value is false which means the fallback application type is confidential client such as a web app. There are certain scenarios where Azure AD cannot determine the client application type. For example, the ROPC flow where the application is configured without specifying a redirect URI. In those cases Azure AD interprets the application type based on the value of this property.
    isFallbackPublicClient *bool
    // The collection of key credentials associated with the application. Not nullable. Supports $filter (eq, not, ge, le).
    keyCredentials []KeyCredentialable
    // The main logo for the application. Not nullable.
    logo []byte
    // Notes relevant for the management of the application.
    notes *string
    // Represents the set of properties required for configuring Application Proxy for this application. Configuring these properties allows you to publish your on-premises application for secure remote access.
    onPremisesPublishing IdentitySecurityDefaultsEnforcementPolicyable
    // Application developers can configure optional claims in their Azure AD applications to specify the claims that are sent to their application by the Microsoft security token service. For more information, see How to: Provide optional claims to your app.
    optionalClaims IdentitySecurityDefaultsEnforcementPolicyable
    // Directory objects that are owners of the application. Read-only. Nullable. Supports $expand.
    owners []DirectoryObjectable
    // Specifies parental control settings for an application.
    parentalControlSettings IdentitySecurityDefaultsEnforcementPolicyable
    // The collection of password credentials associated with the application. Not nullable.
    passwordCredentials []PasswordCredentialable
    // Specifies settings for installed clients such as desktop or mobile devices.
    publicClient IdentitySecurityDefaultsEnforcementPolicyable
    // The verified publisher domain for the application. Read-only. Supports $filter (eq, ne, ge, le, startsWith).
    publisherDomain *string
    // Specifies the resources that the application needs to access. This property also specifies the set of delegated permissions and application roles that it needs for each of those resources. This configuration of access to the required resources drives the consent experience. No more than 50 resource services (APIs) can be configured. Beginning mid-October 2021, the total number of required permissions must not exceed 400. Not nullable. Supports $filter (eq, not, ge, le).
    requiredResourceAccess []RequiredResourceAccessable
    // The URL where the service exposes SAML metadata for federation. This property is valid only for single-tenant applications. Nullable.
    samlMetadataUrl *string
    // References application or service contact information from a Service or Asset Management database. Nullable.
    serviceManagementReference *string
    // Specifies the Microsoft accounts that are supported for the current application. The possible values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount (default), and PersonalMicrosoftAccount. See more in the table below. Supports $filter (eq, ne, not).
    signInAudience *string
    // Specifies settings for a single-page application, including sign out URLs and redirect URIs for authorization codes and access tokens.
    spa IdentitySecurityDefaultsEnforcementPolicyable
    // The synchronization property
    synchronization IdentitySecurityDefaultsEnforcementPolicyable
    // Custom strings that can be used to categorize and identify the application. Not nullable.Supports $filter (eq, not, ge, le, startsWith).
    tags []string
    // Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD encrypts all the tokens it emits by using the key this property points to. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
    tokenEncryptionKeyId *string
    // The tokenIssuancePolicies property
    tokenIssuancePolicies []TokenIssuancePolicyable
    // The tokenLifetimePolicies assigned to this application. Supports $expand.
    tokenLifetimePolicies []TokenLifetimePolicyable
    // The unique identifier that can be assigned to an application as an alternative identifier. Immutable. Read-only.
    uniqueName *string
    // Specifies the verified publisher of the application. For more information about how publisher verification helps support application security, trustworthiness, and compliance, see Publisher verification.
    verifiedPublisher IdentitySecurityDefaultsEnforcementPolicyable
    // Specifies settings for a web application.
    web IdentitySecurityDefaultsEnforcementPolicyable
    // Specifies settings for apps running Microsoft Windows and published in the Microsoft Store or Xbox games store.
    windows IdentitySecurityDefaultsEnforcementPolicyable
}
// IdentitySecurityDefaultsEnforcementPolicy union type wrapper for classes certification, identitySecurityDefaultsEnforcementPolicyMember1
type IdentitySecurityDefaultsEnforcementPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type apiApplication
    apiApplication *ApiApplication
    // Union type representation for type certification
    certification *Certification
    // Union type representation for type connectorGroup
    connectorGroup *ConnectorGroup
    // Union type representation for type directoryObject
    directoryObject *DirectoryObject
    // Union type representation for type identitySecurityDefaultsEnforcementPolicyMember1
    identitySecurityDefaultsEnforcementPolicyMember1 *IdentitySecurityDefaultsEnforcementPolicyMember1
    // Union type representation for type informationalUrl
    informationalUrl *InformationalUrl
    // Union type representation for type onPremisesPublishing
    onPremisesPublishing *OnPremisesPublishing
    // Union type representation for type optionalClaims
    optionalClaims *OptionalClaims
    // Union type representation for type parentalControlSettings
    parentalControlSettings *ParentalControlSettings
    // Union type representation for type publicClientApplication
    publicClientApplication *PublicClientApplication
    // Union type representation for type spaApplication
    spaApplication *SpaApplication
    // Union type representation for type synchronization
    synchronization *Synchronization
    // Union type representation for type verifiedPublisher
    verifiedPublisher *VerifiedPublisher
    // Union type representation for type webApplication
    webApplication *WebApplication
    // Union type representation for type windowsApplication
    windowsApplication *WindowsApplication
}
// NewIdentitySecurityDefaultsEnforcementPolicy instantiates a new identitySecurityDefaultsEnforcementPolicy and sets the default values.
func NewIdentitySecurityDefaultsEnforcementPolicy()(*IdentitySecurityDefaultsEnforcementPolicy) {
    m := &IdentitySecurityDefaultsEnforcementPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentitySecurityDefaultsEnforcementPolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApiApplication gets the apiApplication property value. Union type representation for type apiApplication
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetApiApplication()(*ApiApplication) {
    if m == nil {
        return nil
    } else {
        return m.apiApplication
    }
}
// GetCertification gets the certification property value. Union type representation for type certification
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetCertification()(*Certification) {
    if m == nil {
        return nil
    } else {
        return m.certification
    }
}
// GetConnectorGroup gets the connectorGroup property value. Union type representation for type connectorGroup
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetConnectorGroup()(*ConnectorGroup) {
    if m == nil {
        return nil
    } else {
        return m.connectorGroup
    }
}
// GetDirectoryObject gets the directoryObject property value. Union type representation for type directoryObject
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetDirectoryObject()(*DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.directoryObject
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["apiApplication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApiApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApiApplication(val.(*ApiApplication))
        }
        return nil
    }
    res["certification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCertificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertification(val.(*Certification))
        }
        return nil
    }
    res["connectorGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConnectorGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorGroup(val.(*ConnectorGroup))
        }
        return nil
    }
    res["directoryObject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryObject(val.(*DirectoryObject))
        }
        return nil
    }
    res["identitySecurityDefaultsEnforcementPolicyMember1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySecurityDefaultsEnforcementPolicyMember1FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentitySecurityDefaultsEnforcementPolicyMember1(val.(*IdentitySecurityDefaultsEnforcementPolicyMember1))
        }
        return nil
    }
    res["informationalUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInformationalUrlFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInformationalUrl(val.(*InformationalUrl))
        }
        return nil
    }
    res["onPremisesPublishing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnPremisesPublishingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesPublishing(val.(*OnPremisesPublishing))
        }
        return nil
    }
    res["optionalClaims"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOptionalClaimsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptionalClaims(val.(*OptionalClaims))
        }
        return nil
    }
    res["parentalControlSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateParentalControlSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentalControlSettings(val.(*ParentalControlSettings))
        }
        return nil
    }
    res["publicClientApplication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePublicClientApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicClientApplication(val.(*PublicClientApplication))
        }
        return nil
    }
    res["spaApplication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSpaApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpaApplication(val.(*SpaApplication))
        }
        return nil
    }
    res["synchronization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSynchronization(val.(*Synchronization))
        }
        return nil
    }
    res["verifiedPublisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVerifiedPublisherFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedPublisher(val.(*VerifiedPublisher))
        }
        return nil
    }
    res["webApplication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebApplication(val.(*WebApplication))
        }
        return nil
    }
    res["windowsApplication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsApplication(val.(*WindowsApplication))
        }
        return nil
    }
    return res
}
// GetIdentitySecurityDefaultsEnforcementPolicyMember1 gets the identitySecurityDefaultsEnforcementPolicyMember1 property value. Union type representation for type identitySecurityDefaultsEnforcementPolicyMember1
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetIdentitySecurityDefaultsEnforcementPolicyMember1()(*IdentitySecurityDefaultsEnforcementPolicyMember1) {
    if m == nil {
        return nil
    } else {
        return m.identitySecurityDefaultsEnforcementPolicyMember1
    }
}
// GetInformationalUrl gets the informationalUrl property value. Union type representation for type informationalUrl
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetInformationalUrl()(*InformationalUrl) {
    if m == nil {
        return nil
    } else {
        return m.informationalUrl
    }
}
// GetOnPremisesPublishing gets the onPremisesPublishing property value. Union type representation for type onPremisesPublishing
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetOnPremisesPublishing()(*OnPremisesPublishing) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesPublishing
    }
}
// GetOptionalClaims gets the optionalClaims property value. Union type representation for type optionalClaims
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetOptionalClaims()(*OptionalClaims) {
    if m == nil {
        return nil
    } else {
        return m.optionalClaims
    }
}
// GetParentalControlSettings gets the parentalControlSettings property value. Union type representation for type parentalControlSettings
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetParentalControlSettings()(*ParentalControlSettings) {
    if m == nil {
        return nil
    } else {
        return m.parentalControlSettings
    }
}
// GetPublicClientApplication gets the publicClientApplication property value. Union type representation for type publicClientApplication
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetPublicClientApplication()(*PublicClientApplication) {
    if m == nil {
        return nil
    } else {
        return m.publicClientApplication
    }
}
// GetSpaApplication gets the spaApplication property value. Union type representation for type spaApplication
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetSpaApplication()(*SpaApplication) {
    if m == nil {
        return nil
    } else {
        return m.spaApplication
    }
}
// GetSynchronization gets the synchronization property value. Union type representation for type synchronization
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetSynchronization()(*Synchronization) {
    if m == nil {
        return nil
    } else {
        return m.synchronization
    }
}
// GetVerifiedPublisher gets the verifiedPublisher property value. Union type representation for type verifiedPublisher
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetVerifiedPublisher()(*VerifiedPublisher) {
    if m == nil {
        return nil
    } else {
        return m.verifiedPublisher
    }
}
// GetWebApplication gets the webApplication property value. Union type representation for type webApplication
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetWebApplication()(*WebApplication) {
    if m == nil {
        return nil
    } else {
        return m.webApplication
    }
}
// GetWindowsApplication gets the windowsApplication property value. Union type representation for type windowsApplication
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetWindowsApplication()(*WindowsApplication) {
    if m == nil {
        return nil
    } else {
        return m.windowsApplication
    }
}
// Serialize serializes information the current object
func (m *IdentitySecurityDefaultsEnforcementPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("apiApplication", m.GetApiApplication())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("certification", m.GetCertification())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("connectorGroup", m.GetConnectorGroup())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("directoryObject", m.GetDirectoryObject())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("identitySecurityDefaultsEnforcementPolicyMember1", m.GetIdentitySecurityDefaultsEnforcementPolicyMember1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("informationalUrl", m.GetInformationalUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("onPremisesPublishing", m.GetOnPremisesPublishing())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("optionalClaims", m.GetOptionalClaims())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("parentalControlSettings", m.GetParentalControlSettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("publicClientApplication", m.GetPublicClientApplication())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("spaApplication", m.GetSpaApplication())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("synchronization", m.GetSynchronization())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("verifiedPublisher", m.GetVerifiedPublisher())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("webApplication", m.GetWebApplication())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("windowsApplication", m.GetWindowsApplication())
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
func (m *IdentitySecurityDefaultsEnforcementPolicy) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApiApplication sets the apiApplication property value. Union type representation for type apiApplication
func (m *IdentitySecurityDefaultsEnforcementPolicy) SetApiApplication(value *ApiApplication)() {
    if m != nil {
        m.apiApplication = value
    }
}
// SetCertification sets the certification property value. Union type representation for type certification
func (m *IdentitySecurityDefaultsEnforcementPolicy) SetCertification(value *Certification)() {
    if m != nil {
        m.certification = value
    }
}
// SetConnectorGroup sets the connectorGroup property value. Union type representation for type connectorGroup
func (m *IdentitySecurityDefaultsEnforcementPolicy) SetConnectorGroup(value *ConnectorGroup)() {
    if m != nil {
        m.connectorGroup = value
    }
}
// SetDirectoryObject sets the directoryObject property value. Union type representation for type directoryObject
func (m *IdentitySecurityDefaultsEnforcementPolicy) SetDirectoryObject(value *DirectoryObject)() {
    if m != nil {
        m.directoryObject = value
    }
}
// SetIdentitySecurityDefaultsEnforcementPolicyMember1 sets the identitySecurityDefaultsEnforcementPolicyMember1 property value. Union type representation for type identitySecurityDefaultsEnforcementPolicyMember1
func (m *IdentitySecurityDefaultsEnforcementPolicy) SetIdentitySecurityDefaultsEnforcementPolicyMember1(value *IdentitySecurityDefaultsEnforcementPolicyMember1)() {
    if m != nil {
        m.identitySecurityDefaultsEnforcementPolicyMember1 = value
    }
}
// SetInformationalUrl sets the informationalUrl property value. Union type representation for type informationalUrl
func (m *IdentitySecurityDefaultsEnforcementPolicy) SetInformationalUrl(value *InformationalUrl)() {
    if m != nil {
        m.informationalUrl = value
    }
}
// SetOnPremisesPublishing sets the onPremisesPublishing property value. Union type representation for type onPremisesPublishing
func (m *IdentitySecurityDefaultsEnforcementPolicy) SetOnPremisesPublishing(value *OnPremisesPublishing)() {
    if m != nil {
        m.onPremisesPublishing = value
    }
}
// SetOptionalClaims sets the optionalClaims property value. Union type representation for type optionalClaims
func (m *IdentitySecurityDefaultsEnforcementPolicy) SetOptionalClaims(value *OptionalClaims)() {
    if m != nil {
        m.optionalClaims = value
    }
}
// SetParentalControlSettings sets the parentalControlSettings property value. Union type representation for type parentalControlSettings
func (m *IdentitySecurityDefaultsEnforcementPolicy) SetParentalControlSettings(value *ParentalControlSettings)() {
    if m != nil {
        m.parentalControlSettings = value
    }
}
// SetPublicClientApplication sets the publicClientApplication property value. Union type representation for type publicClientApplication
func (m *IdentitySecurityDefaultsEnforcementPolicy) SetPublicClientApplication(value *PublicClientApplication)() {
    if m != nil {
        m.publicClientApplication = value
    }
}
// SetSpaApplication sets the spaApplication property value. Union type representation for type spaApplication
func (m *IdentitySecurityDefaultsEnforcementPolicy) SetSpaApplication(value *SpaApplication)() {
    if m != nil {
        m.spaApplication = value
    }
}
// SetSynchronization sets the synchronization property value. Union type representation for type synchronization
func (m *IdentitySecurityDefaultsEnforcementPolicy) SetSynchronization(value *Synchronization)() {
    if m != nil {
        m.synchronization = value
    }
}
// SetVerifiedPublisher sets the verifiedPublisher property value. Union type representation for type verifiedPublisher
func (m *IdentitySecurityDefaultsEnforcementPolicy) SetVerifiedPublisher(value *VerifiedPublisher)() {
    if m != nil {
        m.verifiedPublisher = value
    }
}
// SetWebApplication sets the webApplication property value. Union type representation for type webApplication
func (m *IdentitySecurityDefaultsEnforcementPolicy) SetWebApplication(value *WebApplication)() {
    if m != nil {
        m.webApplication = value
    }
}
// SetWindowsApplication sets the windowsApplication property value. Union type representation for type windowsApplication
func (m *IdentitySecurityDefaultsEnforcementPolicy) SetWindowsApplication(value *WindowsApplication)() {
    if m != nil {
        m.windowsApplication = value
    }
}
// NewApplication instantiates a new Application and sets the default values.
func NewApplication()(*Application) {
    m := &Application{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// CreateApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplication(), nil
}
// GetApi gets the api property value. Specifies settings for an application that implements a web API.
func (m *Application) GetApi()(IdentitySecurityDefaultsEnforcementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.api
    }
}
// GetAppId gets the appId property value. The unique identifier for the application that is assigned by Azure AD. Not nullable. Read-only.
func (m *Application) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
// GetAppManagementPolicies gets the appManagementPolicies property value. The appManagementPolicy applied to this application.
func (m *Application) GetAppManagementPolicies()([]AppManagementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.appManagementPolicies
    }
}
// GetAppRoles gets the appRoles property value. The collection of roles assigned to the application. With app role assignments, these roles can be assigned to users, groups, or service principals associated with other applications. Not nullable.
func (m *Application) GetAppRoles()([]AppRoleable) {
    if m == nil {
        return nil
    } else {
        return m.appRoles
    }
}
// GetCertification gets the certification property value. Specifies the certification status of the application.
func (m *Application) GetCertification()(IdentitySecurityDefaultsEnforcementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.certification
    }
}
// GetConnectorGroup gets the connectorGroup property value. The connectorGroup the application is using with Azure AD Application Proxy. Nullable.
func (m *Application) GetConnectorGroup()(IdentitySecurityDefaultsEnforcementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.connectorGroup
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the application was registered. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.  Supports $filter (eq, ne, not, ge, le, in, and eq on null values) and $orderBy.
func (m *Application) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetCreatedOnBehalfOf gets the createdOnBehalfOf property value. The createdOnBehalfOf property
func (m *Application) GetCreatedOnBehalfOf()(IdentitySecurityDefaultsEnforcementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.createdOnBehalfOf
    }
}
// GetDefaultRedirectUri gets the defaultRedirectUri property value. The default redirect URI. If specified and there is no explicit redirect URI in the sign-in request for SAML and OIDC flows, Azure AD sends the token to this redirect URI. Azure AD also sends the token to this default URI in SAML IdP-initiated single sign-on. The value must match one of the configured redirect URIs for the application.
func (m *Application) GetDefaultRedirectUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultRedirectUri
    }
}
// GetDescription gets the description property value. Free text field to provide a description of the application object to end users. The maximum allowed size is 1024 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
func (m *Application) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisabledByMicrosoftStatus gets the disabledByMicrosoftStatus property value. Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, not).
func (m *Application) GetDisabledByMicrosoftStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.disabledByMicrosoftStatus
    }
}
// GetDisplayName gets the displayName property value. The display name for the application. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *Application) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetExtensionProperties gets the extensionProperties property value. Read-only. Nullable. Supports $expand and $filter (eq when counting empty collections).
func (m *Application) GetExtensionProperties()([]ExtensionPropertyable) {
    if m == nil {
        return nil
    } else {
        return m.extensionProperties
    }
}
// GetFederatedIdentityCredentials gets the federatedIdentityCredentials property value. Federated identities for applications. Supports $expand and $filter (eq when counting empty collections).
func (m *Application) GetFederatedIdentityCredentials()([]FederatedIdentityCredentialable) {
    if m == nil {
        return nil
    } else {
        return m.federatedIdentityCredentials
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Application) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["api"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApi(val.(IdentitySecurityDefaultsEnforcementPolicyable))
        }
        return nil
    }
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["appManagementPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppManagementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppManagementPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(AppManagementPolicyable)
            }
            m.SetAppManagementPolicies(res)
        }
        return nil
    }
    res["appRoles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppRoleable, len(val))
            for i, v := range val {
                res[i] = v.(AppRoleable)
            }
            m.SetAppRoles(res)
        }
        return nil
    }
    res["certification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertification(val.(IdentitySecurityDefaultsEnforcementPolicyable))
        }
        return nil
    }
    res["connectorGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorGroup(val.(IdentitySecurityDefaultsEnforcementPolicyable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["createdOnBehalfOf"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedOnBehalfOf(val.(IdentitySecurityDefaultsEnforcementPolicyable))
        }
        return nil
    }
    res["defaultRedirectUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultRedirectUri(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["disabledByMicrosoftStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisabledByMicrosoftStatus(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["extensionProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExtensionPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExtensionPropertyable, len(val))
            for i, v := range val {
                res[i] = v.(ExtensionPropertyable)
            }
            m.SetExtensionProperties(res)
        }
        return nil
    }
    res["federatedIdentityCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFederatedIdentityCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FederatedIdentityCredentialable, len(val))
            for i, v := range val {
                res[i] = v.(FederatedIdentityCredentialable)
            }
            m.SetFederatedIdentityCredentials(res)
        }
        return nil
    }
    res["groupMembershipClaims"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupMembershipClaims(val)
        }
        return nil
    }
    res["homeRealmDiscoveryPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHomeRealmDiscoveryPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HomeRealmDiscoveryPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(HomeRealmDiscoveryPolicyable)
            }
            m.SetHomeRealmDiscoveryPolicies(res)
        }
        return nil
    }
    res["identifierUris"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIdentifierUris(res)
        }
        return nil
    }
    res["info"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInfo(val.(IdentitySecurityDefaultsEnforcementPolicyable))
        }
        return nil
    }
    res["isDeviceOnlyAuthSupported"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeviceOnlyAuthSupported(val)
        }
        return nil
    }
    res["isFallbackPublicClient"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFallbackPublicClient(val)
        }
        return nil
    }
    res["keyCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyCredentialable, len(val))
            for i, v := range val {
                res[i] = v.(KeyCredentialable)
            }
            m.SetKeyCredentials(res)
        }
        return nil
    }
    res["logo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogo(val)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    res["onPremisesPublishing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesPublishing(val.(IdentitySecurityDefaultsEnforcementPolicyable))
        }
        return nil
    }
    res["optionalClaims"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptionalClaims(val.(IdentitySecurityDefaultsEnforcementPolicyable))
        }
        return nil
    }
    res["owners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetOwners(res)
        }
        return nil
    }
    res["parentalControlSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentalControlSettings(val.(IdentitySecurityDefaultsEnforcementPolicyable))
        }
        return nil
    }
    res["passwordCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePasswordCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PasswordCredentialable, len(val))
            for i, v := range val {
                res[i] = v.(PasswordCredentialable)
            }
            m.SetPasswordCredentials(res)
        }
        return nil
    }
    res["publicClient"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicClient(val.(IdentitySecurityDefaultsEnforcementPolicyable))
        }
        return nil
    }
    res["publisherDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisherDomain(val)
        }
        return nil
    }
    res["requiredResourceAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRequiredResourceAccessFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RequiredResourceAccessable, len(val))
            for i, v := range val {
                res[i] = v.(RequiredResourceAccessable)
            }
            m.SetRequiredResourceAccess(res)
        }
        return nil
    }
    res["samlMetadataUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSamlMetadataUrl(val)
        }
        return nil
    }
    res["serviceManagementReference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceManagementReference(val)
        }
        return nil
    }
    res["signInAudience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInAudience(val)
        }
        return nil
    }
    res["spa"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpa(val.(IdentitySecurityDefaultsEnforcementPolicyable))
        }
        return nil
    }
    res["synchronization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSynchronization(val.(IdentitySecurityDefaultsEnforcementPolicyable))
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTags(res)
        }
        return nil
    }
    res["tokenEncryptionKeyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenEncryptionKeyId(val)
        }
        return nil
    }
    res["tokenIssuancePolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTokenIssuancePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TokenIssuancePolicyable, len(val))
            for i, v := range val {
                res[i] = v.(TokenIssuancePolicyable)
            }
            m.SetTokenIssuancePolicies(res)
        }
        return nil
    }
    res["tokenLifetimePolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTokenLifetimePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TokenLifetimePolicyable, len(val))
            for i, v := range val {
                res[i] = v.(TokenLifetimePolicyable)
            }
            m.SetTokenLifetimePolicies(res)
        }
        return nil
    }
    res["uniqueName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueName(val)
        }
        return nil
    }
    res["verifiedPublisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedPublisher(val.(IdentitySecurityDefaultsEnforcementPolicyable))
        }
        return nil
    }
    res["web"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWeb(val.(IdentitySecurityDefaultsEnforcementPolicyable))
        }
        return nil
    }
    res["windows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindows(val.(IdentitySecurityDefaultsEnforcementPolicyable))
        }
        return nil
    }
    return res
}
// GetGroupMembershipClaims gets the groupMembershipClaims property value. Configures the groups claim issued in a user or OAuth 2.0 access token that the application expects. To set this attribute, use one of the following string values: None, SecurityGroup (for security groups and Azure AD roles), All (this gets all security groups, distribution groups, and Azure AD directory roles that the signed-in user is a member of).
func (m *Application) GetGroupMembershipClaims()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupMembershipClaims
    }
}
// GetHomeRealmDiscoveryPolicies gets the homeRealmDiscoveryPolicies property value. The homeRealmDiscoveryPolicies property
func (m *Application) GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.homeRealmDiscoveryPolicies
    }
}
// GetIdentifierUris gets the identifierUris property value. Also known as App ID URI, this value is set when an application is used as a resource app. The identifierUris acts as the prefix for the scopes you'll reference in your API's code, and it must be globally unique. You can use the default value provided, which is in the form api://<application-client-id>, or specify a more readable URI like https://contoso.com/api. For more information on valid identifierUris patterns and best practices, see Azure AD application registration security best practices. Not nullable. Supports $filter (eq, ne, ge, le, startsWith).
func (m *Application) GetIdentifierUris()([]string) {
    if m == nil {
        return nil
    } else {
        return m.identifierUris
    }
}
// GetInfo gets the info property value. Basic profile information of the application, such as it's marketing, support, terms of service, and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more information, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, not, ge, le, and eq on null values).
func (m *Application) GetInfo()(IdentitySecurityDefaultsEnforcementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.info
    }
}
// GetIsDeviceOnlyAuthSupported gets the isDeviceOnlyAuthSupported property value. Specifies whether this application supports device authentication without a user. The default is false.
func (m *Application) GetIsDeviceOnlyAuthSupported()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeviceOnlyAuthSupported
    }
}
// GetIsFallbackPublicClient gets the isFallbackPublicClient property value. Specifies the fallback application type as public client, such as an installed application running on a mobile device. The default value is false which means the fallback application type is confidential client such as a web app. There are certain scenarios where Azure AD cannot determine the client application type. For example, the ROPC flow where the application is configured without specifying a redirect URI. In those cases Azure AD interprets the application type based on the value of this property.
func (m *Application) GetIsFallbackPublicClient()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isFallbackPublicClient
    }
}
// GetKeyCredentials gets the keyCredentials property value. The collection of key credentials associated with the application. Not nullable. Supports $filter (eq, not, ge, le).
func (m *Application) GetKeyCredentials()([]KeyCredentialable) {
    if m == nil {
        return nil
    } else {
        return m.keyCredentials
    }
}
// GetLogo gets the logo property value. The main logo for the application. Not nullable.
func (m *Application) GetLogo()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.logo
    }
}
// GetNotes gets the notes property value. Notes relevant for the management of the application.
func (m *Application) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// GetOnPremisesPublishing gets the onPremisesPublishing property value. Represents the set of properties required for configuring Application Proxy for this application. Configuring these properties allows you to publish your on-premises application for secure remote access.
func (m *Application) GetOnPremisesPublishing()(IdentitySecurityDefaultsEnforcementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesPublishing
    }
}
// GetOptionalClaims gets the optionalClaims property value. Application developers can configure optional claims in their Azure AD applications to specify the claims that are sent to their application by the Microsoft security token service. For more information, see How to: Provide optional claims to your app.
func (m *Application) GetOptionalClaims()(IdentitySecurityDefaultsEnforcementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.optionalClaims
    }
}
// GetOwners gets the owners property value. Directory objects that are owners of the application. Read-only. Nullable. Supports $expand.
func (m *Application) GetOwners()([]DirectoryObjectable) {
    if m == nil {
        return nil
    } else {
        return m.owners
    }
}
// GetParentalControlSettings gets the parentalControlSettings property value. Specifies parental control settings for an application.
func (m *Application) GetParentalControlSettings()(IdentitySecurityDefaultsEnforcementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.parentalControlSettings
    }
}
// GetPasswordCredentials gets the passwordCredentials property value. The collection of password credentials associated with the application. Not nullable.
func (m *Application) GetPasswordCredentials()([]PasswordCredentialable) {
    if m == nil {
        return nil
    } else {
        return m.passwordCredentials
    }
}
// GetPublicClient gets the publicClient property value. Specifies settings for installed clients such as desktop or mobile devices.
func (m *Application) GetPublicClient()(IdentitySecurityDefaultsEnforcementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.publicClient
    }
}
// GetPublisherDomain gets the publisherDomain property value. The verified publisher domain for the application. Read-only. Supports $filter (eq, ne, ge, le, startsWith).
func (m *Application) GetPublisherDomain()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisherDomain
    }
}
// GetRequiredResourceAccess gets the requiredResourceAccess property value. Specifies the resources that the application needs to access. This property also specifies the set of delegated permissions and application roles that it needs for each of those resources. This configuration of access to the required resources drives the consent experience. No more than 50 resource services (APIs) can be configured. Beginning mid-October 2021, the total number of required permissions must not exceed 400. Not nullable. Supports $filter (eq, not, ge, le).
func (m *Application) GetRequiredResourceAccess()([]RequiredResourceAccessable) {
    if m == nil {
        return nil
    } else {
        return m.requiredResourceAccess
    }
}
// GetSamlMetadataUrl gets the samlMetadataUrl property value. The URL where the service exposes SAML metadata for federation. This property is valid only for single-tenant applications. Nullable.
func (m *Application) GetSamlMetadataUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.samlMetadataUrl
    }
}
// GetServiceManagementReference gets the serviceManagementReference property value. References application or service contact information from a Service or Asset Management database. Nullable.
func (m *Application) GetServiceManagementReference()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceManagementReference
    }
}
// GetSignInAudience gets the signInAudience property value. Specifies the Microsoft accounts that are supported for the current application. The possible values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount (default), and PersonalMicrosoftAccount. See more in the table below. Supports $filter (eq, ne, not).
func (m *Application) GetSignInAudience()(*string) {
    if m == nil {
        return nil
    } else {
        return m.signInAudience
    }
}
// GetSpa gets the spa property value. Specifies settings for a single-page application, including sign out URLs and redirect URIs for authorization codes and access tokens.
func (m *Application) GetSpa()(IdentitySecurityDefaultsEnforcementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.spa
    }
}
// GetSynchronization gets the synchronization property value. The synchronization property
func (m *Application) GetSynchronization()(IdentitySecurityDefaultsEnforcementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.synchronization
    }
}
// GetTags gets the tags property value. Custom strings that can be used to categorize and identify the application. Not nullable.Supports $filter (eq, not, ge, le, startsWith).
func (m *Application) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// GetTokenEncryptionKeyId gets the tokenEncryptionKeyId property value. Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD encrypts all the tokens it emits by using the key this property points to. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
func (m *Application) GetTokenEncryptionKeyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tokenEncryptionKeyId
    }
}
// GetTokenIssuancePolicies gets the tokenIssuancePolicies property value. The tokenIssuancePolicies property
func (m *Application) GetTokenIssuancePolicies()([]TokenIssuancePolicyable) {
    if m == nil {
        return nil
    } else {
        return m.tokenIssuancePolicies
    }
}
// GetTokenLifetimePolicies gets the tokenLifetimePolicies property value. The tokenLifetimePolicies assigned to this application. Supports $expand.
func (m *Application) GetTokenLifetimePolicies()([]TokenLifetimePolicyable) {
    if m == nil {
        return nil
    } else {
        return m.tokenLifetimePolicies
    }
}
// GetUniqueName gets the uniqueName property value. The unique identifier that can be assigned to an application as an alternative identifier. Immutable. Read-only.
func (m *Application) GetUniqueName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.uniqueName
    }
}
// GetVerifiedPublisher gets the verifiedPublisher property value. Specifies the verified publisher of the application. For more information about how publisher verification helps support application security, trustworthiness, and compliance, see Publisher verification.
func (m *Application) GetVerifiedPublisher()(IdentitySecurityDefaultsEnforcementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.verifiedPublisher
    }
}
// GetWeb gets the web property value. Specifies settings for a web application.
func (m *Application) GetWeb()(IdentitySecurityDefaultsEnforcementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.web
    }
}
// GetWindows gets the windows property value. Specifies settings for apps running Microsoft Windows and published in the Microsoft Store or Xbox games store.
func (m *Application) GetWindows()(IdentitySecurityDefaultsEnforcementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.windows
    }
}
// Serialize serializes information the current object
func (m *Application) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("api", m.GetApi())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    if m.GetAppManagementPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppManagementPolicies()))
        for i, v := range m.GetAppManagementPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appManagementPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppRoles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppRoles()))
        for i, v := range m.GetAppRoles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appRoles", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("certification", m.GetCertification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("connectorGroup", m.GetConnectorGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdOnBehalfOf", m.GetCreatedOnBehalfOf())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultRedirectUri", m.GetDefaultRedirectUri())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("disabledByMicrosoftStatus", m.GetDisabledByMicrosoftStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetExtensionProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtensionProperties()))
        for i, v := range m.GetExtensionProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("extensionProperties", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFederatedIdentityCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFederatedIdentityCredentials()))
        for i, v := range m.GetFederatedIdentityCredentials() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("federatedIdentityCredentials", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupMembershipClaims", m.GetGroupMembershipClaims())
        if err != nil {
            return err
        }
    }
    if m.GetHomeRealmDiscoveryPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHomeRealmDiscoveryPolicies()))
        for i, v := range m.GetHomeRealmDiscoveryPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("homeRealmDiscoveryPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIdentifierUris() != nil {
        err = writer.WriteCollectionOfStringValues("identifierUris", m.GetIdentifierUris())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("info", m.GetInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeviceOnlyAuthSupported", m.GetIsDeviceOnlyAuthSupported())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFallbackPublicClient", m.GetIsFallbackPublicClient())
        if err != nil {
            return err
        }
    }
    if m.GetKeyCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetKeyCredentials()))
        for i, v := range m.GetKeyCredentials() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("keyCredentials", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("logo", m.GetLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onPremisesPublishing", m.GetOnPremisesPublishing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("optionalClaims", m.GetOptionalClaims())
        if err != nil {
            return err
        }
    }
    if m.GetOwners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOwners()))
        for i, v := range m.GetOwners() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("owners", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parentalControlSettings", m.GetParentalControlSettings())
        if err != nil {
            return err
        }
    }
    if m.GetPasswordCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPasswordCredentials()))
        for i, v := range m.GetPasswordCredentials() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("passwordCredentials", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("publicClient", m.GetPublicClient())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisherDomain", m.GetPublisherDomain())
        if err != nil {
            return err
        }
    }
    if m.GetRequiredResourceAccess() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRequiredResourceAccess()))
        for i, v := range m.GetRequiredResourceAccess() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("requiredResourceAccess", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("samlMetadataUrl", m.GetSamlMetadataUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceManagementReference", m.GetServiceManagementReference())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signInAudience", m.GetSignInAudience())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("spa", m.GetSpa())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("synchronization", m.GetSynchronization())
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tokenEncryptionKeyId", m.GetTokenEncryptionKeyId())
        if err != nil {
            return err
        }
    }
    if m.GetTokenIssuancePolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTokenIssuancePolicies()))
        for i, v := range m.GetTokenIssuancePolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tokenIssuancePolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTokenLifetimePolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTokenLifetimePolicies()))
        for i, v := range m.GetTokenLifetimePolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tokenLifetimePolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("uniqueName", m.GetUniqueName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("verifiedPublisher", m.GetVerifiedPublisher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("web", m.GetWeb())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windows", m.GetWindows())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApi sets the api property value. Specifies settings for an application that implements a web API.
func (m *Application) SetApi(value IdentitySecurityDefaultsEnforcementPolicyable)() {
    if m != nil {
        m.api = value
    }
}
// SetAppId sets the appId property value. The unique identifier for the application that is assigned by Azure AD. Not nullable. Read-only.
func (m *Application) SetAppId(value *string)() {
    if m != nil {
        m.appId = value
    }
}
// SetAppManagementPolicies sets the appManagementPolicies property value. The appManagementPolicy applied to this application.
func (m *Application) SetAppManagementPolicies(value []AppManagementPolicyable)() {
    if m != nil {
        m.appManagementPolicies = value
    }
}
// SetAppRoles sets the appRoles property value. The collection of roles assigned to the application. With app role assignments, these roles can be assigned to users, groups, or service principals associated with other applications. Not nullable.
func (m *Application) SetAppRoles(value []AppRoleable)() {
    if m != nil {
        m.appRoles = value
    }
}
// SetCertification sets the certification property value. Specifies the certification status of the application.
func (m *Application) SetCertification(value IdentitySecurityDefaultsEnforcementPolicyable)() {
    if m != nil {
        m.certification = value
    }
}
// SetConnectorGroup sets the connectorGroup property value. The connectorGroup the application is using with Azure AD Application Proxy. Nullable.
func (m *Application) SetConnectorGroup(value IdentitySecurityDefaultsEnforcementPolicyable)() {
    if m != nil {
        m.connectorGroup = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the application was registered. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.  Supports $filter (eq, ne, not, ge, le, in, and eq on null values) and $orderBy.
func (m *Application) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetCreatedOnBehalfOf sets the createdOnBehalfOf property value. The createdOnBehalfOf property
func (m *Application) SetCreatedOnBehalfOf(value IdentitySecurityDefaultsEnforcementPolicyable)() {
    if m != nil {
        m.createdOnBehalfOf = value
    }
}
// SetDefaultRedirectUri sets the defaultRedirectUri property value. The default redirect URI. If specified and there is no explicit redirect URI in the sign-in request for SAML and OIDC flows, Azure AD sends the token to this redirect URI. Azure AD also sends the token to this default URI in SAML IdP-initiated single sign-on. The value must match one of the configured redirect URIs for the application.
func (m *Application) SetDefaultRedirectUri(value *string)() {
    if m != nil {
        m.defaultRedirectUri = value
    }
}
// SetDescription sets the description property value. Free text field to provide a description of the application object to end users. The maximum allowed size is 1024 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
func (m *Application) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisabledByMicrosoftStatus sets the disabledByMicrosoftStatus property value. Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, not).
func (m *Application) SetDisabledByMicrosoftStatus(value *string)() {
    if m != nil {
        m.disabledByMicrosoftStatus = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the application. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *Application) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetExtensionProperties sets the extensionProperties property value. Read-only. Nullable. Supports $expand and $filter (eq when counting empty collections).
func (m *Application) SetExtensionProperties(value []ExtensionPropertyable)() {
    if m != nil {
        m.extensionProperties = value
    }
}
// SetFederatedIdentityCredentials sets the federatedIdentityCredentials property value. Federated identities for applications. Supports $expand and $filter (eq when counting empty collections).
func (m *Application) SetFederatedIdentityCredentials(value []FederatedIdentityCredentialable)() {
    if m != nil {
        m.federatedIdentityCredentials = value
    }
}
// SetGroupMembershipClaims sets the groupMembershipClaims property value. Configures the groups claim issued in a user or OAuth 2.0 access token that the application expects. To set this attribute, use one of the following string values: None, SecurityGroup (for security groups and Azure AD roles), All (this gets all security groups, distribution groups, and Azure AD directory roles that the signed-in user is a member of).
func (m *Application) SetGroupMembershipClaims(value *string)() {
    if m != nil {
        m.groupMembershipClaims = value
    }
}
// SetHomeRealmDiscoveryPolicies sets the homeRealmDiscoveryPolicies property value. The homeRealmDiscoveryPolicies property
func (m *Application) SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicyable)() {
    if m != nil {
        m.homeRealmDiscoveryPolicies = value
    }
}
// SetIdentifierUris sets the identifierUris property value. Also known as App ID URI, this value is set when an application is used as a resource app. The identifierUris acts as the prefix for the scopes you'll reference in your API's code, and it must be globally unique. You can use the default value provided, which is in the form api://<application-client-id>, or specify a more readable URI like https://contoso.com/api. For more information on valid identifierUris patterns and best practices, see Azure AD application registration security best practices. Not nullable. Supports $filter (eq, ne, ge, le, startsWith).
func (m *Application) SetIdentifierUris(value []string)() {
    if m != nil {
        m.identifierUris = value
    }
}
// SetInfo sets the info property value. Basic profile information of the application, such as it's marketing, support, terms of service, and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more information, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, not, ge, le, and eq on null values).
func (m *Application) SetInfo(value IdentitySecurityDefaultsEnforcementPolicyable)() {
    if m != nil {
        m.info = value
    }
}
// SetIsDeviceOnlyAuthSupported sets the isDeviceOnlyAuthSupported property value. Specifies whether this application supports device authentication without a user. The default is false.
func (m *Application) SetIsDeviceOnlyAuthSupported(value *bool)() {
    if m != nil {
        m.isDeviceOnlyAuthSupported = value
    }
}
// SetIsFallbackPublicClient sets the isFallbackPublicClient property value. Specifies the fallback application type as public client, such as an installed application running on a mobile device. The default value is false which means the fallback application type is confidential client such as a web app. There are certain scenarios where Azure AD cannot determine the client application type. For example, the ROPC flow where the application is configured without specifying a redirect URI. In those cases Azure AD interprets the application type based on the value of this property.
func (m *Application) SetIsFallbackPublicClient(value *bool)() {
    if m != nil {
        m.isFallbackPublicClient = value
    }
}
// SetKeyCredentials sets the keyCredentials property value. The collection of key credentials associated with the application. Not nullable. Supports $filter (eq, not, ge, le).
func (m *Application) SetKeyCredentials(value []KeyCredentialable)() {
    if m != nil {
        m.keyCredentials = value
    }
}
// SetLogo sets the logo property value. The main logo for the application. Not nullable.
func (m *Application) SetLogo(value []byte)() {
    if m != nil {
        m.logo = value
    }
}
// SetNotes sets the notes property value. Notes relevant for the management of the application.
func (m *Application) SetNotes(value *string)() {
    if m != nil {
        m.notes = value
    }
}
// SetOnPremisesPublishing sets the onPremisesPublishing property value. Represents the set of properties required for configuring Application Proxy for this application. Configuring these properties allows you to publish your on-premises application for secure remote access.
func (m *Application) SetOnPremisesPublishing(value IdentitySecurityDefaultsEnforcementPolicyable)() {
    if m != nil {
        m.onPremisesPublishing = value
    }
}
// SetOptionalClaims sets the optionalClaims property value. Application developers can configure optional claims in their Azure AD applications to specify the claims that are sent to their application by the Microsoft security token service. For more information, see How to: Provide optional claims to your app.
func (m *Application) SetOptionalClaims(value IdentitySecurityDefaultsEnforcementPolicyable)() {
    if m != nil {
        m.optionalClaims = value
    }
}
// SetOwners sets the owners property value. Directory objects that are owners of the application. Read-only. Nullable. Supports $expand.
func (m *Application) SetOwners(value []DirectoryObjectable)() {
    if m != nil {
        m.owners = value
    }
}
// SetParentalControlSettings sets the parentalControlSettings property value. Specifies parental control settings for an application.
func (m *Application) SetParentalControlSettings(value IdentitySecurityDefaultsEnforcementPolicyable)() {
    if m != nil {
        m.parentalControlSettings = value
    }
}
// SetPasswordCredentials sets the passwordCredentials property value. The collection of password credentials associated with the application. Not nullable.
func (m *Application) SetPasswordCredentials(value []PasswordCredentialable)() {
    if m != nil {
        m.passwordCredentials = value
    }
}
// SetPublicClient sets the publicClient property value. Specifies settings for installed clients such as desktop or mobile devices.
func (m *Application) SetPublicClient(value IdentitySecurityDefaultsEnforcementPolicyable)() {
    if m != nil {
        m.publicClient = value
    }
}
// SetPublisherDomain sets the publisherDomain property value. The verified publisher domain for the application. Read-only. Supports $filter (eq, ne, ge, le, startsWith).
func (m *Application) SetPublisherDomain(value *string)() {
    if m != nil {
        m.publisherDomain = value
    }
}
// SetRequiredResourceAccess sets the requiredResourceAccess property value. Specifies the resources that the application needs to access. This property also specifies the set of delegated permissions and application roles that it needs for each of those resources. This configuration of access to the required resources drives the consent experience. No more than 50 resource services (APIs) can be configured. Beginning mid-October 2021, the total number of required permissions must not exceed 400. Not nullable. Supports $filter (eq, not, ge, le).
func (m *Application) SetRequiredResourceAccess(value []RequiredResourceAccessable)() {
    if m != nil {
        m.requiredResourceAccess = value
    }
}
// SetSamlMetadataUrl sets the samlMetadataUrl property value. The URL where the service exposes SAML metadata for federation. This property is valid only for single-tenant applications. Nullable.
func (m *Application) SetSamlMetadataUrl(value *string)() {
    if m != nil {
        m.samlMetadataUrl = value
    }
}
// SetServiceManagementReference sets the serviceManagementReference property value. References application or service contact information from a Service or Asset Management database. Nullable.
func (m *Application) SetServiceManagementReference(value *string)() {
    if m != nil {
        m.serviceManagementReference = value
    }
}
// SetSignInAudience sets the signInAudience property value. Specifies the Microsoft accounts that are supported for the current application. The possible values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount (default), and PersonalMicrosoftAccount. See more in the table below. Supports $filter (eq, ne, not).
func (m *Application) SetSignInAudience(value *string)() {
    if m != nil {
        m.signInAudience = value
    }
}
// SetSpa sets the spa property value. Specifies settings for a single-page application, including sign out URLs and redirect URIs for authorization codes and access tokens.
func (m *Application) SetSpa(value IdentitySecurityDefaultsEnforcementPolicyable)() {
    if m != nil {
        m.spa = value
    }
}
// SetSynchronization sets the synchronization property value. The synchronization property
func (m *Application) SetSynchronization(value IdentitySecurityDefaultsEnforcementPolicyable)() {
    if m != nil {
        m.synchronization = value
    }
}
// SetTags sets the tags property value. Custom strings that can be used to categorize and identify the application. Not nullable.Supports $filter (eq, not, ge, le, startsWith).
func (m *Application) SetTags(value []string)() {
    if m != nil {
        m.tags = value
    }
}
// SetTokenEncryptionKeyId sets the tokenEncryptionKeyId property value. Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD encrypts all the tokens it emits by using the key this property points to. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
func (m *Application) SetTokenEncryptionKeyId(value *string)() {
    if m != nil {
        m.tokenEncryptionKeyId = value
    }
}
// SetTokenIssuancePolicies sets the tokenIssuancePolicies property value. The tokenIssuancePolicies property
func (m *Application) SetTokenIssuancePolicies(value []TokenIssuancePolicyable)() {
    if m != nil {
        m.tokenIssuancePolicies = value
    }
}
// SetTokenLifetimePolicies sets the tokenLifetimePolicies property value. The tokenLifetimePolicies assigned to this application. Supports $expand.
func (m *Application) SetTokenLifetimePolicies(value []TokenLifetimePolicyable)() {
    if m != nil {
        m.tokenLifetimePolicies = value
    }
}
// SetUniqueName sets the uniqueName property value. The unique identifier that can be assigned to an application as an alternative identifier. Immutable. Read-only.
func (m *Application) SetUniqueName(value *string)() {
    if m != nil {
        m.uniqueName = value
    }
}
// SetVerifiedPublisher sets the verifiedPublisher property value. Specifies the verified publisher of the application. For more information about how publisher verification helps support application security, trustworthiness, and compliance, see Publisher verification.
func (m *Application) SetVerifiedPublisher(value IdentitySecurityDefaultsEnforcementPolicyable)() {
    if m != nil {
        m.verifiedPublisher = value
    }
}
// SetWeb sets the web property value. Specifies settings for a web application.
func (m *Application) SetWeb(value IdentitySecurityDefaultsEnforcementPolicyable)() {
    if m != nil {
        m.web = value
    }
}
// SetWindows sets the windows property value. Specifies settings for apps running Microsoft Windows and published in the Microsoft Store or Xbox games store.
func (m *Application) SetWindows(value IdentitySecurityDefaultsEnforcementPolicyable)() {
    if m != nil {
        m.windows = value
    }
}
