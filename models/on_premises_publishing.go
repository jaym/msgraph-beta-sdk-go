package models

import (
    i7cc7b2988c0a6476329d81e10ddb2ab3a063f332f60e624716de9721058f8e7e "onpremisespublishing"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesPublishing 
type OnPremisesPublishing struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // If you are configuring a traffic manager in front of multiple App Proxy applications, the alternateUrl is the user-friendly URL that will point to the traffic manager.
    alternateUrl *string
    // The duration the connector will wait for a response from the backend application before closing the connection. Possible values are default, long. When set to default, the backend application timeout has a length of 85 seconds. When set to long, the backend timeout is increased to 180 seconds. Use long if your server takes more than 85 seconds to respond to requests or if you are unable to access the application and the error status is 'Backend Timeout'. Default value is default.
    applicationServerTimeout *string
    // Indicates if this application is an Application Proxy configured application. This is pre-set by the system. Read-only.
    applicationType *string
    // Details the pre-authentication setting for the application. Pre-authentication enforces that users must authenticate before accessing the app. Passthru does not require authentication. Possible values are: passthru, aadPreAuthentication.
    externalAuthenticationType Applicationsable
    // The published external url for the application. For example, https://intranet-contoso.msappproxy.net/.
    externalUrl *string
    // The internal url of the application. For example, https://intranet/.
    internalUrl *string
    // Indicates whether backend SSL certificate validation is enabled for the application. For all new Application Proxy apps, the property will be set to true by default. For all existing apps, the property will be set to false.
    isBackendCertificateValidationEnabled *bool
    // Indicates if the HTTPOnly cookie flag should be set in the HTTP response headers. Set this value to true to have Application Proxy cookies include the HTTPOnly flag in the HTTP response headers. If using Remote Desktop Services, set this value to False. Default value is false.
    isHttpOnlyCookieEnabled *bool
    // Indicates if the application is currently being published via Application Proxy or not. This is pre-set by the system. Read-only.
    isOnPremPublishingEnabled *bool
    // Indicates if the Persistent cookie flag should be set in the HTTP response headers. Keep this value set to false. Only use this setting for applications that can't share cookies between processes. For more information about cookie settings, see Cookie settings for accessing on-premises applications in Azure Active Directory. Default value is false.
    isPersistentCookieEnabled *bool
    // Indicates if the Secure cookie flag should be set in the HTTP response headers. Set this value to true to transmit cookies over a secure channel such as an encrypted HTTPS request. Default value is true.
    isSecureCookieEnabled *bool
    // Indicates whether validation of the state parameter when the client uses the OAuth 2.0 authorization code grant flow is enabled. This setting allows admins to specify whether they want to enable CSRF protection for their apps.
    isStateSessionEnabled *bool
    // Indicates if the application should translate urls in the reponse headers. Keep this value as true unless your application required the original host header in the authentication request. Default value is true.
    isTranslateHostHeaderEnabled *bool
    // Indicates if the application should translate urls in the application body. Keep this value as false unless you have hardcoded HTML links to other on-premises applications and don't use custom domains. For more information, see Link translation with Application Proxy. Default value is false.
    isTranslateLinksInBodyEnabled *bool
    // The onPremisesApplicationSegments property
    onPremisesApplicationSegments []Applicationsable
    // Represents the single sign-on configuration for the on-premises application.
    singleSignOnSettings Applicationsable
    // The useAlternateUrlForTranslationAndRedirect property
    useAlternateUrlForTranslationAndRedirect *bool
    // Details of the certificate associated with the application when a custom domain is in use. null when using the default domain. Read-only.
    verifiedCustomDomainCertificatesMetadata Applicationsable
    // The associated key credential for the custom domain used.
    verifiedCustomDomainKeyCredential Applicationsable
    // The associated password credential for the custom domain used.
    verifiedCustomDomainPasswordCredential Applicationsable
}
// Applications union type wrapper for classes verifiedCustomDomainCertificatesMetadata, applicationsMember1
type Applications struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type applicationsMember1
    applicationsMember1 ApplicationsMember1able
    // Union type representation for type externalAuthenticationType
    externalAuthenticationType *ExternalAuthenticationType
    // Union type representation for type keyCredential
    keyCredential KeyCredentialable
    // Union type representation for type onPremisesApplicationSegment
    onPremisesApplicationSegment OnPremisesApplicationSegmentable
    // Union type representation for type onPremisesPublishingSingleSignOn
    onPremisesPublishingSingleSignOn OnPremisesPublishingSingleSignOnable
    // Union type representation for type passwordCredential
    passwordCredential PasswordCredentialable
    // Union type representation for type verifiedCustomDomainCertificatesMetadata
    verifiedCustomDomainCertificatesMetadata VerifiedCustomDomainCertificatesMetadataable
}
// NewApplications instantiates a new applications and sets the default values.
func NewApplications()(*Applications) {
    m := &Applications{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateApplicationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplicationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplications(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Applications) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplicationsMember1 gets the applicationsMember1 property value. Union type representation for type applicationsMember1
func (m *Applications) GetApplicationsMember1()(ApplicationsMember1able) {
    if m == nil {
        return nil
    } else {
        return m.applicationsMember1
    }
}
// GetExternalAuthenticationType gets the externalAuthenticationType property value. Union type representation for type externalAuthenticationType
func (m *Applications) GetExternalAuthenticationType()(*ExternalAuthenticationType) {
    if m == nil {
        return nil
    } else {
        return m.externalAuthenticationType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Applications) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicationsMember1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApplicationsMember1FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationsMember1(val.(ApplicationsMember1able))
        }
        return nil
    }
    res["externalAuthenticationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExternalAuthenticationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalAuthenticationType(val.(*ExternalAuthenticationType))
        }
        return nil
    }
    res["keyCredential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateKeyCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyCredential(val.(KeyCredentialable))
        }
        return nil
    }
    res["onPremisesApplicationSegment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnPremisesApplicationSegmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesApplicationSegment(val.(OnPremisesApplicationSegmentable))
        }
        return nil
    }
    res["onPremisesPublishingSingleSignOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnPremisesPublishingSingleSignOnFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesPublishingSingleSignOn(val.(OnPremisesPublishingSingleSignOnable))
        }
        return nil
    }
    res["passwordCredential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePasswordCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordCredential(val.(PasswordCredentialable))
        }
        return nil
    }
    res["verifiedCustomDomainCertificatesMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVerifiedCustomDomainCertificatesMetadataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedCustomDomainCertificatesMetadata(val.(VerifiedCustomDomainCertificatesMetadataable))
        }
        return nil
    }
    return res
}
// GetKeyCredential gets the keyCredential property value. Union type representation for type keyCredential
func (m *Applications) GetKeyCredential()(KeyCredentialable) {
    if m == nil {
        return nil
    } else {
        return m.keyCredential
    }
}
// GetOnPremisesApplicationSegment gets the onPremisesApplicationSegment property value. Union type representation for type onPremisesApplicationSegment
func (m *Applications) GetOnPremisesApplicationSegment()(OnPremisesApplicationSegmentable) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesApplicationSegment
    }
}
// GetOnPremisesPublishingSingleSignOn gets the onPremisesPublishingSingleSignOn property value. Union type representation for type onPremisesPublishingSingleSignOn
func (m *Applications) GetOnPremisesPublishingSingleSignOn()(OnPremisesPublishingSingleSignOnable) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesPublishingSingleSignOn
    }
}
// GetPasswordCredential gets the passwordCredential property value. Union type representation for type passwordCredential
func (m *Applications) GetPasswordCredential()(PasswordCredentialable) {
    if m == nil {
        return nil
    } else {
        return m.passwordCredential
    }
}
// GetVerifiedCustomDomainCertificatesMetadata gets the verifiedCustomDomainCertificatesMetadata property value. Union type representation for type verifiedCustomDomainCertificatesMetadata
func (m *Applications) GetVerifiedCustomDomainCertificatesMetadata()(VerifiedCustomDomainCertificatesMetadataable) {
    if m == nil {
        return nil
    } else {
        return m.verifiedCustomDomainCertificatesMetadata
    }
}
// Serialize serializes information the current object
func (m *Applications) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("applicationsMember1", m.GetApplicationsMember1())
        if err != nil {
            return err
        }
    }
    if m.GetExternalAuthenticationType() != nil {
        cast := (*m.GetExternalAuthenticationType()).String()
        err := writer.WriteStringValue("externalAuthenticationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("keyCredential", m.GetKeyCredential())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("onPremisesApplicationSegment", m.GetOnPremisesApplicationSegment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("onPremisesPublishingSingleSignOn", m.GetOnPremisesPublishingSingleSignOn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("passwordCredential", m.GetPasswordCredential())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("verifiedCustomDomainCertificatesMetadata", m.GetVerifiedCustomDomainCertificatesMetadata())
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
func (m *Applications) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplicationsMember1 sets the applicationsMember1 property value. Union type representation for type applicationsMember1
func (m *Applications) SetApplicationsMember1(value ApplicationsMember1able)() {
    if m != nil {
        m.applicationsMember1 = value
    }
}
// SetExternalAuthenticationType sets the externalAuthenticationType property value. Union type representation for type externalAuthenticationType
func (m *Applications) SetExternalAuthenticationType(value *ExternalAuthenticationType)() {
    if m != nil {
        m.externalAuthenticationType = value
    }
}
// SetKeyCredential sets the keyCredential property value. Union type representation for type keyCredential
func (m *Applications) SetKeyCredential(value KeyCredentialable)() {
    if m != nil {
        m.keyCredential = value
    }
}
// SetOnPremisesApplicationSegment sets the onPremisesApplicationSegment property value. Union type representation for type onPremisesApplicationSegment
func (m *Applications) SetOnPremisesApplicationSegment(value OnPremisesApplicationSegmentable)() {
    if m != nil {
        m.onPremisesApplicationSegment = value
    }
}
// SetOnPremisesPublishingSingleSignOn sets the onPremisesPublishingSingleSignOn property value. Union type representation for type onPremisesPublishingSingleSignOn
func (m *Applications) SetOnPremisesPublishingSingleSignOn(value OnPremisesPublishingSingleSignOnable)() {
    if m != nil {
        m.onPremisesPublishingSingleSignOn = value
    }
}
// SetPasswordCredential sets the passwordCredential property value. Union type representation for type passwordCredential
func (m *Applications) SetPasswordCredential(value PasswordCredentialable)() {
    if m != nil {
        m.passwordCredential = value
    }
}
// SetVerifiedCustomDomainCertificatesMetadata sets the verifiedCustomDomainCertificatesMetadata property value. Union type representation for type verifiedCustomDomainCertificatesMetadata
func (m *Applications) SetVerifiedCustomDomainCertificatesMetadata(value VerifiedCustomDomainCertificatesMetadataable)() {
    if m != nil {
        m.verifiedCustomDomainCertificatesMetadata = value
    }
}
// Applicationsable 
type Applicationsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationsMember1()(ApplicationsMember1able)
    GetExternalAuthenticationType()(*ExternalAuthenticationType)
    GetKeyCredential()(KeyCredentialable)
    GetOnPremisesApplicationSegment()(OnPremisesApplicationSegmentable)
    GetOnPremisesPublishingSingleSignOn()(OnPremisesPublishingSingleSignOnable)
    GetPasswordCredential()(PasswordCredentialable)
    GetVerifiedCustomDomainCertificatesMetadata()(VerifiedCustomDomainCertificatesMetadataable)
    SetApplicationsMember1(value ApplicationsMember1able)()
    SetExternalAuthenticationType(value *ExternalAuthenticationType)()
    SetKeyCredential(value KeyCredentialable)()
    SetOnPremisesApplicationSegment(value OnPremisesApplicationSegmentable)()
    SetOnPremisesPublishingSingleSignOn(value OnPremisesPublishingSingleSignOnable)()
    SetPasswordCredential(value PasswordCredentialable)()
    SetVerifiedCustomDomainCertificatesMetadata(value VerifiedCustomDomainCertificatesMetadataable)()
}
// NewOnPremisesPublishing instantiates a new onPremisesPublishing and sets the default values.
func NewOnPremisesPublishing()(*OnPremisesPublishing) {
    m := &OnPremisesPublishing{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOnPremisesPublishingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesPublishingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesPublishing(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesPublishing) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAlternateUrl gets the alternateUrl property value. If you are configuring a traffic manager in front of multiple App Proxy applications, the alternateUrl is the user-friendly URL that will point to the traffic manager.
func (m *OnPremisesPublishing) GetAlternateUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alternateUrl
    }
}
// GetApplicationServerTimeout gets the applicationServerTimeout property value. The duration the connector will wait for a response from the backend application before closing the connection. Possible values are default, long. When set to default, the backend application timeout has a length of 85 seconds. When set to long, the backend timeout is increased to 180 seconds. Use long if your server takes more than 85 seconds to respond to requests or if you are unable to access the application and the error status is 'Backend Timeout'. Default value is default.
func (m *OnPremisesPublishing) GetApplicationServerTimeout()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationServerTimeout
    }
}
// GetApplicationType gets the applicationType property value. Indicates if this application is an Application Proxy configured application. This is pre-set by the system. Read-only.
func (m *OnPremisesPublishing) GetApplicationType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationType
    }
}
// GetExternalAuthenticationType gets the externalAuthenticationType property value. Details the pre-authentication setting for the application. Pre-authentication enforces that users must authenticate before accessing the app. Passthru does not require authentication. Possible values are: passthru, aadPreAuthentication.
func (m *OnPremisesPublishing) GetExternalAuthenticationType()(Applicationsable) {
    if m == nil {
        return nil
    } else {
        return m.externalAuthenticationType
    }
}
// GetExternalUrl gets the externalUrl property value. The published external url for the application. For example, https://intranet-contoso.msappproxy.net/.
func (m *OnPremisesPublishing) GetExternalUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesPublishing) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alternateUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlternateUrl(val)
        }
        return nil
    }
    res["applicationServerTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationServerTimeout(val)
        }
        return nil
    }
    res["applicationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationType(val)
        }
        return nil
    }
    res["externalAuthenticationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApplicationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalAuthenticationType(val.(Applicationsable))
        }
        return nil
    }
    res["externalUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalUrl(val)
        }
        return nil
    }
    res["internalUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalUrl(val)
        }
        return nil
    }
    res["isBackendCertificateValidationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBackendCertificateValidationEnabled(val)
        }
        return nil
    }
    res["isHttpOnlyCookieEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHttpOnlyCookieEnabled(val)
        }
        return nil
    }
    res["isOnPremPublishingEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOnPremPublishingEnabled(val)
        }
        return nil
    }
    res["isPersistentCookieEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPersistentCookieEnabled(val)
        }
        return nil
    }
    res["isSecureCookieEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSecureCookieEnabled(val)
        }
        return nil
    }
    res["isStateSessionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsStateSessionEnabled(val)
        }
        return nil
    }
    res["isTranslateHostHeaderEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTranslateHostHeaderEnabled(val)
        }
        return nil
    }
    res["isTranslateLinksInBodyEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTranslateLinksInBodyEnabled(val)
        }
        return nil
    }
    res["onPremisesApplicationSegments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApplicationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Applicationsable, len(val))
            for i, v := range val {
                res[i] = v.(Applicationsable)
            }
            m.SetOnPremisesApplicationSegments(res)
        }
        return nil
    }
    res["singleSignOnSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApplicationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSignOnSettings(val.(Applicationsable))
        }
        return nil
    }
    res["useAlternateUrlForTranslationAndRedirect"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseAlternateUrlForTranslationAndRedirect(val)
        }
        return nil
    }
    res["verifiedCustomDomainCertificatesMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApplicationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedCustomDomainCertificatesMetadata(val.(Applicationsable))
        }
        return nil
    }
    res["verifiedCustomDomainKeyCredential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApplicationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedCustomDomainKeyCredential(val.(Applicationsable))
        }
        return nil
    }
    res["verifiedCustomDomainPasswordCredential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApplicationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedCustomDomainPasswordCredential(val.(Applicationsable))
        }
        return nil
    }
    return res
}
// GetInternalUrl gets the internalUrl property value. The internal url of the application. For example, https://intranet/.
func (m *OnPremisesPublishing) GetInternalUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.internalUrl
    }
}
// GetIsBackendCertificateValidationEnabled gets the isBackendCertificateValidationEnabled property value. Indicates whether backend SSL certificate validation is enabled for the application. For all new Application Proxy apps, the property will be set to true by default. For all existing apps, the property will be set to false.
func (m *OnPremisesPublishing) GetIsBackendCertificateValidationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBackendCertificateValidationEnabled
    }
}
// GetIsHttpOnlyCookieEnabled gets the isHttpOnlyCookieEnabled property value. Indicates if the HTTPOnly cookie flag should be set in the HTTP response headers. Set this value to true to have Application Proxy cookies include the HTTPOnly flag in the HTTP response headers. If using Remote Desktop Services, set this value to False. Default value is false.
func (m *OnPremisesPublishing) GetIsHttpOnlyCookieEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isHttpOnlyCookieEnabled
    }
}
// GetIsOnPremPublishingEnabled gets the isOnPremPublishingEnabled property value. Indicates if the application is currently being published via Application Proxy or not. This is pre-set by the system. Read-only.
func (m *OnPremisesPublishing) GetIsOnPremPublishingEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOnPremPublishingEnabled
    }
}
// GetIsPersistentCookieEnabled gets the isPersistentCookieEnabled property value. Indicates if the Persistent cookie flag should be set in the HTTP response headers. Keep this value set to false. Only use this setting for applications that can't share cookies between processes. For more information about cookie settings, see Cookie settings for accessing on-premises applications in Azure Active Directory. Default value is false.
func (m *OnPremisesPublishing) GetIsPersistentCookieEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPersistentCookieEnabled
    }
}
// GetIsSecureCookieEnabled gets the isSecureCookieEnabled property value. Indicates if the Secure cookie flag should be set in the HTTP response headers. Set this value to true to transmit cookies over a secure channel such as an encrypted HTTPS request. Default value is true.
func (m *OnPremisesPublishing) GetIsSecureCookieEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSecureCookieEnabled
    }
}
// GetIsStateSessionEnabled gets the isStateSessionEnabled property value. Indicates whether validation of the state parameter when the client uses the OAuth 2.0 authorization code grant flow is enabled. This setting allows admins to specify whether they want to enable CSRF protection for their apps.
func (m *OnPremisesPublishing) GetIsStateSessionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isStateSessionEnabled
    }
}
// GetIsTranslateHostHeaderEnabled gets the isTranslateHostHeaderEnabled property value. Indicates if the application should translate urls in the reponse headers. Keep this value as true unless your application required the original host header in the authentication request. Default value is true.
func (m *OnPremisesPublishing) GetIsTranslateHostHeaderEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isTranslateHostHeaderEnabled
    }
}
// GetIsTranslateLinksInBodyEnabled gets the isTranslateLinksInBodyEnabled property value. Indicates if the application should translate urls in the application body. Keep this value as false unless you have hardcoded HTML links to other on-premises applications and don't use custom domains. For more information, see Link translation with Application Proxy. Default value is false.
func (m *OnPremisesPublishing) GetIsTranslateLinksInBodyEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isTranslateLinksInBodyEnabled
    }
}
// GetOnPremisesApplicationSegments gets the onPremisesApplicationSegments property value. The onPremisesApplicationSegments property
func (m *OnPremisesPublishing) GetOnPremisesApplicationSegments()([]Applicationsable) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesApplicationSegments
    }
}
// GetSingleSignOnSettings gets the singleSignOnSettings property value. Represents the single sign-on configuration for the on-premises application.
func (m *OnPremisesPublishing) GetSingleSignOnSettings()(Applicationsable) {
    if m == nil {
        return nil
    } else {
        return m.singleSignOnSettings
    }
}
// GetUseAlternateUrlForTranslationAndRedirect gets the useAlternateUrlForTranslationAndRedirect property value. The useAlternateUrlForTranslationAndRedirect property
func (m *OnPremisesPublishing) GetUseAlternateUrlForTranslationAndRedirect()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.useAlternateUrlForTranslationAndRedirect
    }
}
// GetVerifiedCustomDomainCertificatesMetadata gets the verifiedCustomDomainCertificatesMetadata property value. Details of the certificate associated with the application when a custom domain is in use. null when using the default domain. Read-only.
func (m *OnPremisesPublishing) GetVerifiedCustomDomainCertificatesMetadata()(Applicationsable) {
    if m == nil {
        return nil
    } else {
        return m.verifiedCustomDomainCertificatesMetadata
    }
}
// GetVerifiedCustomDomainKeyCredential gets the verifiedCustomDomainKeyCredential property value. The associated key credential for the custom domain used.
func (m *OnPremisesPublishing) GetVerifiedCustomDomainKeyCredential()(Applicationsable) {
    if m == nil {
        return nil
    } else {
        return m.verifiedCustomDomainKeyCredential
    }
}
// GetVerifiedCustomDomainPasswordCredential gets the verifiedCustomDomainPasswordCredential property value. The associated password credential for the custom domain used.
func (m *OnPremisesPublishing) GetVerifiedCustomDomainPasswordCredential()(Applicationsable) {
    if m == nil {
        return nil
    } else {
        return m.verifiedCustomDomainPasswordCredential
    }
}
// Serialize serializes information the current object
func (m *OnPremisesPublishing) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("alternateUrl", m.GetAlternateUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("applicationServerTimeout", m.GetApplicationServerTimeout())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("applicationType", m.GetApplicationType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("externalAuthenticationType", m.GetExternalAuthenticationType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalUrl", m.GetExternalUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("internalUrl", m.GetInternalUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isBackendCertificateValidationEnabled", m.GetIsBackendCertificateValidationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isHttpOnlyCookieEnabled", m.GetIsHttpOnlyCookieEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isOnPremPublishingEnabled", m.GetIsOnPremPublishingEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPersistentCookieEnabled", m.GetIsPersistentCookieEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSecureCookieEnabled", m.GetIsSecureCookieEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isStateSessionEnabled", m.GetIsStateSessionEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isTranslateHostHeaderEnabled", m.GetIsTranslateHostHeaderEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isTranslateLinksInBodyEnabled", m.GetIsTranslateLinksInBodyEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetOnPremisesApplicationSegments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOnPremisesApplicationSegments()))
        for i, v := range m.GetOnPremisesApplicationSegments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("onPremisesApplicationSegments", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("singleSignOnSettings", m.GetSingleSignOnSettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("useAlternateUrlForTranslationAndRedirect", m.GetUseAlternateUrlForTranslationAndRedirect())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("verifiedCustomDomainCertificatesMetadata", m.GetVerifiedCustomDomainCertificatesMetadata())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("verifiedCustomDomainKeyCredential", m.GetVerifiedCustomDomainKeyCredential())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("verifiedCustomDomainPasswordCredential", m.GetVerifiedCustomDomainPasswordCredential())
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
func (m *OnPremisesPublishing) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAlternateUrl sets the alternateUrl property value. If you are configuring a traffic manager in front of multiple App Proxy applications, the alternateUrl is the user-friendly URL that will point to the traffic manager.
func (m *OnPremisesPublishing) SetAlternateUrl(value *string)() {
    if m != nil {
        m.alternateUrl = value
    }
}
// SetApplicationServerTimeout sets the applicationServerTimeout property value. The duration the connector will wait for a response from the backend application before closing the connection. Possible values are default, long. When set to default, the backend application timeout has a length of 85 seconds. When set to long, the backend timeout is increased to 180 seconds. Use long if your server takes more than 85 seconds to respond to requests or if you are unable to access the application and the error status is 'Backend Timeout'. Default value is default.
func (m *OnPremisesPublishing) SetApplicationServerTimeout(value *string)() {
    if m != nil {
        m.applicationServerTimeout = value
    }
}
// SetApplicationType sets the applicationType property value. Indicates if this application is an Application Proxy configured application. This is pre-set by the system. Read-only.
func (m *OnPremisesPublishing) SetApplicationType(value *string)() {
    if m != nil {
        m.applicationType = value
    }
}
// SetExternalAuthenticationType sets the externalAuthenticationType property value. Details the pre-authentication setting for the application. Pre-authentication enforces that users must authenticate before accessing the app. Passthru does not require authentication. Possible values are: passthru, aadPreAuthentication.
func (m *OnPremisesPublishing) SetExternalAuthenticationType(value Applicationsable)() {
    if m != nil {
        m.externalAuthenticationType = value
    }
}
// SetExternalUrl sets the externalUrl property value. The published external url for the application. For example, https://intranet-contoso.msappproxy.net/.
func (m *OnPremisesPublishing) SetExternalUrl(value *string)() {
    if m != nil {
        m.externalUrl = value
    }
}
// SetInternalUrl sets the internalUrl property value. The internal url of the application. For example, https://intranet/.
func (m *OnPremisesPublishing) SetInternalUrl(value *string)() {
    if m != nil {
        m.internalUrl = value
    }
}
// SetIsBackendCertificateValidationEnabled sets the isBackendCertificateValidationEnabled property value. Indicates whether backend SSL certificate validation is enabled for the application. For all new Application Proxy apps, the property will be set to true by default. For all existing apps, the property will be set to false.
func (m *OnPremisesPublishing) SetIsBackendCertificateValidationEnabled(value *bool)() {
    if m != nil {
        m.isBackendCertificateValidationEnabled = value
    }
}
// SetIsHttpOnlyCookieEnabled sets the isHttpOnlyCookieEnabled property value. Indicates if the HTTPOnly cookie flag should be set in the HTTP response headers. Set this value to true to have Application Proxy cookies include the HTTPOnly flag in the HTTP response headers. If using Remote Desktop Services, set this value to False. Default value is false.
func (m *OnPremisesPublishing) SetIsHttpOnlyCookieEnabled(value *bool)() {
    if m != nil {
        m.isHttpOnlyCookieEnabled = value
    }
}
// SetIsOnPremPublishingEnabled sets the isOnPremPublishingEnabled property value. Indicates if the application is currently being published via Application Proxy or not. This is pre-set by the system. Read-only.
func (m *OnPremisesPublishing) SetIsOnPremPublishingEnabled(value *bool)() {
    if m != nil {
        m.isOnPremPublishingEnabled = value
    }
}
// SetIsPersistentCookieEnabled sets the isPersistentCookieEnabled property value. Indicates if the Persistent cookie flag should be set in the HTTP response headers. Keep this value set to false. Only use this setting for applications that can't share cookies between processes. For more information about cookie settings, see Cookie settings for accessing on-premises applications in Azure Active Directory. Default value is false.
func (m *OnPremisesPublishing) SetIsPersistentCookieEnabled(value *bool)() {
    if m != nil {
        m.isPersistentCookieEnabled = value
    }
}
// SetIsSecureCookieEnabled sets the isSecureCookieEnabled property value. Indicates if the Secure cookie flag should be set in the HTTP response headers. Set this value to true to transmit cookies over a secure channel such as an encrypted HTTPS request. Default value is true.
func (m *OnPremisesPublishing) SetIsSecureCookieEnabled(value *bool)() {
    if m != nil {
        m.isSecureCookieEnabled = value
    }
}
// SetIsStateSessionEnabled sets the isStateSessionEnabled property value. Indicates whether validation of the state parameter when the client uses the OAuth 2.0 authorization code grant flow is enabled. This setting allows admins to specify whether they want to enable CSRF protection for their apps.
func (m *OnPremisesPublishing) SetIsStateSessionEnabled(value *bool)() {
    if m != nil {
        m.isStateSessionEnabled = value
    }
}
// SetIsTranslateHostHeaderEnabled sets the isTranslateHostHeaderEnabled property value. Indicates if the application should translate urls in the reponse headers. Keep this value as true unless your application required the original host header in the authentication request. Default value is true.
func (m *OnPremisesPublishing) SetIsTranslateHostHeaderEnabled(value *bool)() {
    if m != nil {
        m.isTranslateHostHeaderEnabled = value
    }
}
// SetIsTranslateLinksInBodyEnabled sets the isTranslateLinksInBodyEnabled property value. Indicates if the application should translate urls in the application body. Keep this value as false unless you have hardcoded HTML links to other on-premises applications and don't use custom domains. For more information, see Link translation with Application Proxy. Default value is false.
func (m *OnPremisesPublishing) SetIsTranslateLinksInBodyEnabled(value *bool)() {
    if m != nil {
        m.isTranslateLinksInBodyEnabled = value
    }
}
// SetOnPremisesApplicationSegments sets the onPremisesApplicationSegments property value. The onPremisesApplicationSegments property
func (m *OnPremisesPublishing) SetOnPremisesApplicationSegments(value []Applicationsable)() {
    if m != nil {
        m.onPremisesApplicationSegments = value
    }
}
// SetSingleSignOnSettings sets the singleSignOnSettings property value. Represents the single sign-on configuration for the on-premises application.
func (m *OnPremisesPublishing) SetSingleSignOnSettings(value Applicationsable)() {
    if m != nil {
        m.singleSignOnSettings = value
    }
}
// SetUseAlternateUrlForTranslationAndRedirect sets the useAlternateUrlForTranslationAndRedirect property value. The useAlternateUrlForTranslationAndRedirect property
func (m *OnPremisesPublishing) SetUseAlternateUrlForTranslationAndRedirect(value *bool)() {
    if m != nil {
        m.useAlternateUrlForTranslationAndRedirect = value
    }
}
// SetVerifiedCustomDomainCertificatesMetadata sets the verifiedCustomDomainCertificatesMetadata property value. Details of the certificate associated with the application when a custom domain is in use. null when using the default domain. Read-only.
func (m *OnPremisesPublishing) SetVerifiedCustomDomainCertificatesMetadata(value Applicationsable)() {
    if m != nil {
        m.verifiedCustomDomainCertificatesMetadata = value
    }
}
// SetVerifiedCustomDomainKeyCredential sets the verifiedCustomDomainKeyCredential property value. The associated key credential for the custom domain used.
func (m *OnPremisesPublishing) SetVerifiedCustomDomainKeyCredential(value Applicationsable)() {
    if m != nil {
        m.verifiedCustomDomainKeyCredential = value
    }
}
// SetVerifiedCustomDomainPasswordCredential sets the verifiedCustomDomainPasswordCredential property value. The associated password credential for the custom domain used.
func (m *OnPremisesPublishing) SetVerifiedCustomDomainPasswordCredential(value Applicationsable)() {
    if m != nil {
        m.verifiedCustomDomainPasswordCredential = value
    }
}
