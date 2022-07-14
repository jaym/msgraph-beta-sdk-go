package models

import (
    ib98ad0ae79f6901bfacb150978e99efcbcf093f9da05b4d2106cc32c0a085b33 "apiapplication"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApiApplication 
type ApiApplication struct {
    // When true, allows an application to use claims mapping without specifying a custom signing key.
    acceptMappedClaims *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app. If you set the appID of the client app to this value, the user only consents once to the client app. Azure AD knows that consenting to the client means implicitly consenting to the web API and automatically provisions service principals for both APIs at the same time. Both the client and the web API app must be registered in the same tenant.
    knownClientApplications []string
    // The definition of the delegated permissions exposed by the web API represented by this application registration. These delegated permissions may be requested by a client application, and may be granted by users or administrators during consent. Delegated permissions are sometimes referred to as OAuth 2.0 scopes.
    oauth2PermissionScopes []PermissionScopeable
    // Lists the client applications that are pre-authorized with the specified delegated permissions to access this application's APIs. Users are not required to consent to any pre-authorized application (for the permissions specified). However, any additional permissions not listed in preAuthorizedApplications (requested through incremental consent for example) will require user consent.
    preAuthorizedApplications []Applicationsable
    // Specifies the access token version expected by this resource. This changes the version and format of the JWT produced independent of the endpoint or client used to request the access token.  The endpoint used, v1.0 or v2.0, is chosen by the client and only impacts the version of id_tokens. Resources need to explicitly configure requestedAccessTokenVersion to indicate the supported access token format.  Possible values for requestedAccessTokenVersion are 1, 2, or null. If the value is null, this defaults to 1, which corresponds to the v1.0 endpoint.  If signInAudience on the application is configured as AzureADandPersonalMicrosoftAccount, the value for this property must be 2
    requestedAccessTokenVersion *int32
}
// Applications union type wrapper for classes preAuthorizedApplication, applicationsMember1
type Applications struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type applicationsMember1
    applicationsMember1 ApplicationsMember1able
    // Union type representation for type preAuthorizedApplication
    preAuthorizedApplication PreAuthorizedApplicationable
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
    res["preAuthorizedApplication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePreAuthorizedApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreAuthorizedApplication(val.(PreAuthorizedApplicationable))
        }
        return nil
    }
    return res
}
// GetPreAuthorizedApplication gets the preAuthorizedApplication property value. Union type representation for type preAuthorizedApplication
func (m *Applications) GetPreAuthorizedApplication()(PreAuthorizedApplicationable) {
    if m == nil {
        return nil
    } else {
        return m.preAuthorizedApplication
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
    {
        err := writer.WriteObjectValue("preAuthorizedApplication", m.GetPreAuthorizedApplication())
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
// SetPreAuthorizedApplication sets the preAuthorizedApplication property value. Union type representation for type preAuthorizedApplication
func (m *Applications) SetPreAuthorizedApplication(value PreAuthorizedApplicationable)() {
    if m != nil {
        m.preAuthorizedApplication = value
    }
}
// Applicationsable 
type Applicationsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationsMember1()(ApplicationsMember1able)
    GetPreAuthorizedApplication()(PreAuthorizedApplicationable)
    SetApplicationsMember1(value ApplicationsMember1able)()
    SetPreAuthorizedApplication(value PreAuthorizedApplicationable)()
}
// NewApiApplication instantiates a new apiApplication and sets the default values.
func NewApiApplication()(*ApiApplication) {
    m := &ApiApplication{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateApiApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApiApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApiApplication(), nil
}
// GetAcceptMappedClaims gets the acceptMappedClaims property value. When true, allows an application to use claims mapping without specifying a custom signing key.
func (m *ApiApplication) GetAcceptMappedClaims()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.acceptMappedClaims
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApiApplication) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApiApplication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["acceptMappedClaims"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceptMappedClaims(val)
        }
        return nil
    }
    res["knownClientApplications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetKnownClientApplications(res)
        }
        return nil
    }
    res["oauth2PermissionScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PermissionScopeable, len(val))
            for i, v := range val {
                res[i] = v.(PermissionScopeable)
            }
            m.SetOauth2PermissionScopes(res)
        }
        return nil
    }
    res["preAuthorizedApplications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApplicationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Applicationsable, len(val))
            for i, v := range val {
                res[i] = v.(Applicationsable)
            }
            m.SetPreAuthorizedApplications(res)
        }
        return nil
    }
    res["requestedAccessTokenVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedAccessTokenVersion(val)
        }
        return nil
    }
    return res
}
// GetKnownClientApplications gets the knownClientApplications property value. Used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app. If you set the appID of the client app to this value, the user only consents once to the client app. Azure AD knows that consenting to the client means implicitly consenting to the web API and automatically provisions service principals for both APIs at the same time. Both the client and the web API app must be registered in the same tenant.
func (m *ApiApplication) GetKnownClientApplications()([]string) {
    if m == nil {
        return nil
    } else {
        return m.knownClientApplications
    }
}
// GetOauth2PermissionScopes gets the oauth2PermissionScopes property value. The definition of the delegated permissions exposed by the web API represented by this application registration. These delegated permissions may be requested by a client application, and may be granted by users or administrators during consent. Delegated permissions are sometimes referred to as OAuth 2.0 scopes.
func (m *ApiApplication) GetOauth2PermissionScopes()([]PermissionScopeable) {
    if m == nil {
        return nil
    } else {
        return m.oauth2PermissionScopes
    }
}
// GetPreAuthorizedApplications gets the preAuthorizedApplications property value. Lists the client applications that are pre-authorized with the specified delegated permissions to access this application's APIs. Users are not required to consent to any pre-authorized application (for the permissions specified). However, any additional permissions not listed in preAuthorizedApplications (requested through incremental consent for example) will require user consent.
func (m *ApiApplication) GetPreAuthorizedApplications()([]Applicationsable) {
    if m == nil {
        return nil
    } else {
        return m.preAuthorizedApplications
    }
}
// GetRequestedAccessTokenVersion gets the requestedAccessTokenVersion property value. Specifies the access token version expected by this resource. This changes the version and format of the JWT produced independent of the endpoint or client used to request the access token.  The endpoint used, v1.0 or v2.0, is chosen by the client and only impacts the version of id_tokens. Resources need to explicitly configure requestedAccessTokenVersion to indicate the supported access token format.  Possible values for requestedAccessTokenVersion are 1, 2, or null. If the value is null, this defaults to 1, which corresponds to the v1.0 endpoint.  If signInAudience on the application is configured as AzureADandPersonalMicrosoftAccount, the value for this property must be 2
func (m *ApiApplication) GetRequestedAccessTokenVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.requestedAccessTokenVersion
    }
}
// Serialize serializes information the current object
func (m *ApiApplication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("acceptMappedClaims", m.GetAcceptMappedClaims())
        if err != nil {
            return err
        }
    }
    if m.GetKnownClientApplications() != nil {
        err := writer.WriteCollectionOfStringValues("knownClientApplications", m.GetKnownClientApplications())
        if err != nil {
            return err
        }
    }
    if m.GetOauth2PermissionScopes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOauth2PermissionScopes()))
        for i, v := range m.GetOauth2PermissionScopes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("oauth2PermissionScopes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPreAuthorizedApplications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPreAuthorizedApplications()))
        for i, v := range m.GetPreAuthorizedApplications() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("preAuthorizedApplications", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("requestedAccessTokenVersion", m.GetRequestedAccessTokenVersion())
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
// SetAcceptMappedClaims sets the acceptMappedClaims property value. When true, allows an application to use claims mapping without specifying a custom signing key.
func (m *ApiApplication) SetAcceptMappedClaims(value *bool)() {
    if m != nil {
        m.acceptMappedClaims = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApiApplication) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetKnownClientApplications sets the knownClientApplications property value. Used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app. If you set the appID of the client app to this value, the user only consents once to the client app. Azure AD knows that consenting to the client means implicitly consenting to the web API and automatically provisions service principals for both APIs at the same time. Both the client and the web API app must be registered in the same tenant.
func (m *ApiApplication) SetKnownClientApplications(value []string)() {
    if m != nil {
        m.knownClientApplications = value
    }
}
// SetOauth2PermissionScopes sets the oauth2PermissionScopes property value. The definition of the delegated permissions exposed by the web API represented by this application registration. These delegated permissions may be requested by a client application, and may be granted by users or administrators during consent. Delegated permissions are sometimes referred to as OAuth 2.0 scopes.
func (m *ApiApplication) SetOauth2PermissionScopes(value []PermissionScopeable)() {
    if m != nil {
        m.oauth2PermissionScopes = value
    }
}
// SetPreAuthorizedApplications sets the preAuthorizedApplications property value. Lists the client applications that are pre-authorized with the specified delegated permissions to access this application's APIs. Users are not required to consent to any pre-authorized application (for the permissions specified). However, any additional permissions not listed in preAuthorizedApplications (requested through incremental consent for example) will require user consent.
func (m *ApiApplication) SetPreAuthorizedApplications(value []Applicationsable)() {
    if m != nil {
        m.preAuthorizedApplications = value
    }
}
// SetRequestedAccessTokenVersion sets the requestedAccessTokenVersion property value. Specifies the access token version expected by this resource. This changes the version and format of the JWT produced independent of the endpoint or client used to request the access token.  The endpoint used, v1.0 or v2.0, is chosen by the client and only impacts the version of id_tokens. Resources need to explicitly configure requestedAccessTokenVersion to indicate the supported access token format.  Possible values for requestedAccessTokenVersion are 1, 2, or null. If the value is null, this defaults to 1, which corresponds to the v1.0 endpoint.  If signInAudience on the application is configured as AzureADandPersonalMicrosoftAccount, the value for this property must be 2
func (m *ApiApplication) SetRequestedAccessTokenVersion(value *int32)() {
    if m != nil {
        m.requestedAccessTokenVersion = value
    }
}
