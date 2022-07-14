package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppManagementPolicy 
type AppManagementPolicy struct {
    PolicyBase
    // The appliesTo property
    appliesTo []DirectoryObjectable
    // The isEnabled property
    isEnabled *bool
    // The restrictions property
    restrictions AuthorizationPolicyable
}
// AuthorizationPolicy union type wrapper for classes appManagementConfiguration, authorizationPolicyMember1
type AuthorizationPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type appManagementConfiguration
    appManagementConfiguration *AppManagementConfiguration
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
// GetAppManagementConfiguration gets the appManagementConfiguration property value. Union type representation for type appManagementConfiguration
func (m *AuthorizationPolicy) GetAppManagementConfiguration()(*AppManagementConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.appManagementConfiguration
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
    res["appManagementConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppManagementConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppManagementConfiguration(val.(*AppManagementConfiguration))
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
    {
        err := writer.WriteObjectValue("appManagementConfiguration", m.GetAppManagementConfiguration())
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
// SetAppManagementConfiguration sets the appManagementConfiguration property value. Union type representation for type appManagementConfiguration
func (m *AuthorizationPolicy) SetAppManagementConfiguration(value *AppManagementConfiguration)() {
    if m != nil {
        m.appManagementConfiguration = value
    }
}
// SetAuthorizationPolicyMember1 sets the authorizationPolicyMember1 property value. Union type representation for type authorizationPolicyMember1
func (m *AuthorizationPolicy) SetAuthorizationPolicyMember1(value *AuthorizationPolicyMember1)() {
    if m != nil {
        m.authorizationPolicyMember1 = value
    }
}
// NewAppManagementPolicy instantiates a new appManagementPolicy and sets the default values.
func NewAppManagementPolicy()(*AppManagementPolicy) {
    m := &AppManagementPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    return m
}
// CreateAppManagementPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppManagementPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppManagementPolicy(), nil
}
// GetAppliesTo gets the appliesTo property value. The appliesTo property
func (m *AppManagementPolicy) GetAppliesTo()([]DirectoryObjectable) {
    if m == nil {
        return nil
    } else {
        return m.appliesTo
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppManagementPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["appliesTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetAppliesTo(res)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["restrictions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorizationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictions(val.(AuthorizationPolicyable))
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. The isEnabled property
func (m *AppManagementPolicy) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetRestrictions gets the restrictions property value. The restrictions property
func (m *AppManagementPolicy) GetRestrictions()(AuthorizationPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.restrictions
    }
}
// Serialize serializes information the current object
func (m *AppManagementPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppliesTo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppliesTo()))
        for i, v := range m.GetAppliesTo() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appliesTo", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("restrictions", m.GetRestrictions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppliesTo sets the appliesTo property value. The appliesTo property
func (m *AppManagementPolicy) SetAppliesTo(value []DirectoryObjectable)() {
    if m != nil {
        m.appliesTo = value
    }
}
// SetIsEnabled sets the isEnabled property value. The isEnabled property
func (m *AppManagementPolicy) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
// SetRestrictions sets the restrictions property value. The restrictions property
func (m *AppManagementPolicy) SetRestrictions(value AuthorizationPolicyable)() {
    if m != nil {
        m.restrictions = value
    }
}
