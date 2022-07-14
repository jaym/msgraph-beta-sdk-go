package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// KeyCredentialConfiguration 
type KeyCredentialConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The maxLifetime property
    maxLifetime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Timestamp when the policy is enforced for all apps created on or after the specified date. For existing applications, the enforcement date would be back dated. To apply to all applications regardless of their creation date, this property would be null. Nullable.
    restrictForAppsCreatedAfterDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The type of restriction being applied. Possible values are asymmetricKeyLifetime, unknownFutureValue. Each value of restrictionType can be used only once per policy.
    restrictionType AuthorizationPolicyable
}
// AuthorizationPolicy union type wrapper for classes appKeyCredentialRestrictionType, authorizationPolicyMember1
type AuthorizationPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type appKeyCredentialRestrictionType
    appKeyCredentialRestrictionType *AppKeyCredentialRestrictionType
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
// GetAppKeyCredentialRestrictionType gets the appKeyCredentialRestrictionType property value. Union type representation for type appKeyCredentialRestrictionType
func (m *AuthorizationPolicy) GetAppKeyCredentialRestrictionType()(*AppKeyCredentialRestrictionType) {
    if m == nil {
        return nil
    } else {
        return m.appKeyCredentialRestrictionType
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
    res["appKeyCredentialRestrictionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppKeyCredentialRestrictionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppKeyCredentialRestrictionType(val.(*AppKeyCredentialRestrictionType))
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
    if m.GetAppKeyCredentialRestrictionType() != nil {
        cast := (*m.GetAppKeyCredentialRestrictionType()).String()
        err := writer.WriteStringValue("appKeyCredentialRestrictionType", &cast)
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
// SetAppKeyCredentialRestrictionType sets the appKeyCredentialRestrictionType property value. Union type representation for type appKeyCredentialRestrictionType
func (m *AuthorizationPolicy) SetAppKeyCredentialRestrictionType(value *AppKeyCredentialRestrictionType)() {
    if m != nil {
        m.appKeyCredentialRestrictionType = value
    }
}
// SetAuthorizationPolicyMember1 sets the authorizationPolicyMember1 property value. Union type representation for type authorizationPolicyMember1
func (m *AuthorizationPolicy) SetAuthorizationPolicyMember1(value *AuthorizationPolicyMember1)() {
    if m != nil {
        m.authorizationPolicyMember1 = value
    }
}
// NewKeyCredentialConfiguration instantiates a new keyCredentialConfiguration and sets the default values.
func NewKeyCredentialConfiguration()(*KeyCredentialConfiguration) {
    m := &KeyCredentialConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateKeyCredentialConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateKeyCredentialConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKeyCredentialConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *KeyCredentialConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *KeyCredentialConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *KeyCredentialConfiguration) GetMaxLifetime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.maxLifetime
    }
}
// GetRestrictForAppsCreatedAfterDateTime gets the restrictForAppsCreatedAfterDateTime property value. Timestamp when the policy is enforced for all apps created on or after the specified date. For existing applications, the enforcement date would be back dated. To apply to all applications regardless of their creation date, this property would be null. Nullable.
func (m *KeyCredentialConfiguration) GetRestrictForAppsCreatedAfterDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.restrictForAppsCreatedAfterDateTime
    }
}
// GetRestrictionType gets the restrictionType property value. The type of restriction being applied. Possible values are asymmetricKeyLifetime, unknownFutureValue. Each value of restrictionType can be used only once per policy.
func (m *KeyCredentialConfiguration) GetRestrictionType()(AuthorizationPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.restrictionType
    }
}
// Serialize serializes information the current object
func (m *KeyCredentialConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *KeyCredentialConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMaxLifetime sets the maxLifetime property value. The maxLifetime property
func (m *KeyCredentialConfiguration) SetMaxLifetime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    if m != nil {
        m.maxLifetime = value
    }
}
// SetRestrictForAppsCreatedAfterDateTime sets the restrictForAppsCreatedAfterDateTime property value. Timestamp when the policy is enforced for all apps created on or after the specified date. For existing applications, the enforcement date would be back dated. To apply to all applications regardless of their creation date, this property would be null. Nullable.
func (m *KeyCredentialConfiguration) SetRestrictForAppsCreatedAfterDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.restrictForAppsCreatedAfterDateTime = value
    }
}
// SetRestrictionType sets the restrictionType property value. The type of restriction being applied. Possible values are asymmetricKeyLifetime, unknownFutureValue. Each value of restrictionType can be used only once per policy.
func (m *KeyCredentialConfiguration) SetRestrictionType(value AuthorizationPolicyable)() {
    if m != nil {
        m.restrictionType = value
    }
}
