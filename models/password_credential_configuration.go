package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PasswordCredentialConfiguration 
type PasswordCredentialConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The maxLifetime property
    maxLifetime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Enforces the policy for an app created on or after the enforcement date. For existing applications, the enforcement date would be backdated. To apply to all applications, this date would be null.
    restrictForAppsCreatedAfterDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The type of restriction being applied. The possible values are: passwordAddition, passwordLifetime, symmetricKeyAddition, symmetricKeyLifetime,customPasswordAddition, unknownFutureValue. Each value of restrictionType can be used only once per policy.
    restrictionType AuthorizationPolicyable
}
// AuthorizationPolicy union type wrapper for classes appCredentialRestrictionType, authorizationPolicyMember1
type AuthorizationPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type appCredentialRestrictionType
    appCredentialRestrictionType *AppCredentialRestrictionType
    // Union type representation for type authorizationPolicyMember1
    authorizationPolicyMember1 *AuthorizationPolicyMember1
}
// NewAuthorizationPolicy instantiates a new authorizationPolicy and sets the default values.
func NewAuthorizationPolicy()(*AuthorizationPolicy) {
    m := &AuthorizationPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAuthorizationPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthorizationPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthorizationPolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthorizationPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppCredentialRestrictionType gets the appCredentialRestrictionType property value. Union type representation for type appCredentialRestrictionType
func (m *AuthorizationPolicy) GetAppCredentialRestrictionType()(*AppCredentialRestrictionType) {
    if m == nil {
        return nil
    } else {
        return m.appCredentialRestrictionType
    }
}
// GetAuthorizationPolicyMember1 gets the authorizationPolicyMember1 property value. Union type representation for type authorizationPolicyMember1
func (m *AuthorizationPolicy) GetAuthorizationPolicyMember1()(*AuthorizationPolicyMember1) {
    if m == nil {
        return nil
    } else {
        return m.authorizationPolicyMember1
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthorizationPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appCredentialRestrictionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppCredentialRestrictionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppCredentialRestrictionType(val.(*AppCredentialRestrictionType))
        }
        return nil
    }
    res["authorizationPolicyMember1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorizationPolicyMember1FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizationPolicyMember1(val.(*AuthorizationPolicyMember1))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AuthorizationPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppCredentialRestrictionType() != nil {
        cast := (*m.GetAppCredentialRestrictionType()).String()
        err := writer.WriteStringValue("appCredentialRestrictionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("authorizationPolicyMember1", m.GetAuthorizationPolicyMember1())
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
func (m *AuthorizationPolicy) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppCredentialRestrictionType sets the appCredentialRestrictionType property value. Union type representation for type appCredentialRestrictionType
func (m *AuthorizationPolicy) SetAppCredentialRestrictionType(value *AppCredentialRestrictionType)() {
    if m != nil {
        m.appCredentialRestrictionType = value
    }
}
// SetAuthorizationPolicyMember1 sets the authorizationPolicyMember1 property value. Union type representation for type authorizationPolicyMember1
func (m *AuthorizationPolicy) SetAuthorizationPolicyMember1(value *AuthorizationPolicyMember1)() {
    if m != nil {
        m.authorizationPolicyMember1 = value
    }
}
// NewPasswordCredentialConfiguration instantiates a new passwordCredentialConfiguration and sets the default values.
func NewPasswordCredentialConfiguration()(*PasswordCredentialConfiguration) {
    m := &PasswordCredentialConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePasswordCredentialConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePasswordCredentialConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPasswordCredentialConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PasswordCredentialConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PasswordCredentialConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["maxLifetime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxLifetime(val)
        }
        return nil
    }
    res["restrictForAppsCreatedAfterDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictForAppsCreatedAfterDateTime(val)
        }
        return nil
    }
    res["restrictionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorizationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictionType(val.(AuthorizationPolicyable))
        }
        return nil
    }
    return res
}
// GetMaxLifetime gets the maxLifetime property value. The maxLifetime property
func (m *PasswordCredentialConfiguration) GetMaxLifetime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.maxLifetime
    }
}
// GetRestrictForAppsCreatedAfterDateTime gets the restrictForAppsCreatedAfterDateTime property value. Enforces the policy for an app created on or after the enforcement date. For existing applications, the enforcement date would be backdated. To apply to all applications, this date would be null.
func (m *PasswordCredentialConfiguration) GetRestrictForAppsCreatedAfterDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.restrictForAppsCreatedAfterDateTime
    }
}
// GetRestrictionType gets the restrictionType property value. The type of restriction being applied. The possible values are: passwordAddition, passwordLifetime, symmetricKeyAddition, symmetricKeyLifetime,customPasswordAddition, unknownFutureValue. Each value of restrictionType can be used only once per policy.
func (m *PasswordCredentialConfiguration) GetRestrictionType()(AuthorizationPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.restrictionType
    }
}
// Serialize serializes information the current object
func (m *PasswordCredentialConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteISODurationValue("maxLifetime", m.GetMaxLifetime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("restrictForAppsCreatedAfterDateTime", m.GetRestrictForAppsCreatedAfterDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("restrictionType", m.GetRestrictionType())
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
func (m *PasswordCredentialConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMaxLifetime sets the maxLifetime property value. The maxLifetime property
func (m *PasswordCredentialConfiguration) SetMaxLifetime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    if m != nil {
        m.maxLifetime = value
    }
}
// SetRestrictForAppsCreatedAfterDateTime sets the restrictForAppsCreatedAfterDateTime property value. Enforces the policy for an app created on or after the enforcement date. For existing applications, the enforcement date would be backdated. To apply to all applications, this date would be null.
func (m *PasswordCredentialConfiguration) SetRestrictForAppsCreatedAfterDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.restrictForAppsCreatedAfterDateTime = value
    }
}
// SetRestrictionType sets the restrictionType property value. The type of restriction being applied. The possible values are: passwordAddition, passwordLifetime, symmetricKeyAddition, symmetricKeyLifetime,customPasswordAddition, unknownFutureValue. Each value of restrictionType can be used only once per policy.
func (m *PasswordCredentialConfiguration) SetRestrictionType(value AuthorizationPolicyable)() {
    if m != nil {
        m.restrictionType = value
    }
}
