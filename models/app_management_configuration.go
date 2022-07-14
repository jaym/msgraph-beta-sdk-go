package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppManagementConfiguration 
type AppManagementConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Collection of keyCredential restrictions settings to be applied to an application or service principal.
    keyCredentials []AuthorizationPolicyable
    // Collection of password restrictions settings to be applied to an application or service principal.
    passwordCredentials []AuthorizationPolicyable
}
// AuthorizationPolicy union type wrapper for classes keyCredentialConfiguration, authorizationPolicyMember1
type AuthorizationPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type authorizationPolicyMember1
    authorizationPolicyMember1 *AuthorizationPolicyMember1
    // Union type representation for type keyCredentialConfiguration
    keyCredentialConfiguration *KeyCredentialConfiguration
    // Union type representation for type passwordCredentialConfiguration
    passwordCredentialConfiguration *PasswordCredentialConfiguration
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
    res["keyCredentialConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateKeyCredentialConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyCredentialConfiguration(val.(*KeyCredentialConfiguration))
        }
        return nil
    }
    res["passwordCredentialConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePasswordCredentialConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordCredentialConfiguration(val.(*PasswordCredentialConfiguration))
        }
        return nil
    }
    return res
}
// GetKeyCredentialConfiguration gets the keyCredentialConfiguration property value. Union type representation for type keyCredentialConfiguration
func (m *AuthorizationPolicy) GetKeyCredentialConfiguration()(*KeyCredentialConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.keyCredentialConfiguration
    }
}
// GetPasswordCredentialConfiguration gets the passwordCredentialConfiguration property value. Union type representation for type passwordCredentialConfiguration
func (m *AuthorizationPolicy) GetPasswordCredentialConfiguration()(*PasswordCredentialConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.passwordCredentialConfiguration
    }
}
// Serialize serializes information the current object
func (m *AuthorizationPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("authorizationPolicyMember1", m.GetAuthorizationPolicyMember1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("keyCredentialConfiguration", m.GetKeyCredentialConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("passwordCredentialConfiguration", m.GetPasswordCredentialConfiguration())
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
// SetAuthorizationPolicyMember1 sets the authorizationPolicyMember1 property value. Union type representation for type authorizationPolicyMember1
func (m *AuthorizationPolicy) SetAuthorizationPolicyMember1(value *AuthorizationPolicyMember1)() {
    if m != nil {
        m.authorizationPolicyMember1 = value
    }
}
// SetKeyCredentialConfiguration sets the keyCredentialConfiguration property value. Union type representation for type keyCredentialConfiguration
func (m *AuthorizationPolicy) SetKeyCredentialConfiguration(value *KeyCredentialConfiguration)() {
    if m != nil {
        m.keyCredentialConfiguration = value
    }
}
// SetPasswordCredentialConfiguration sets the passwordCredentialConfiguration property value. Union type representation for type passwordCredentialConfiguration
func (m *AuthorizationPolicy) SetPasswordCredentialConfiguration(value *PasswordCredentialConfiguration)() {
    if m != nil {
        m.passwordCredentialConfiguration = value
    }
}
// NewAppManagementConfiguration instantiates a new appManagementConfiguration and sets the default values.
func NewAppManagementConfiguration()(*AppManagementConfiguration) {
    m := &AppManagementConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAppManagementConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppManagementConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppManagementConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppManagementConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppManagementConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["keyCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthorizationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthorizationPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(AuthorizationPolicyable)
            }
            m.SetKeyCredentials(res)
        }
        return nil
    }
    res["passwordCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthorizationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthorizationPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(AuthorizationPolicyable)
            }
            m.SetPasswordCredentials(res)
        }
        return nil
    }
    return res
}
// GetKeyCredentials gets the keyCredentials property value. Collection of keyCredential restrictions settings to be applied to an application or service principal.
func (m *AppManagementConfiguration) GetKeyCredentials()([]AuthorizationPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.keyCredentials
    }
}
// GetPasswordCredentials gets the passwordCredentials property value. Collection of password restrictions settings to be applied to an application or service principal.
func (m *AppManagementConfiguration) GetPasswordCredentials()([]AuthorizationPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.passwordCredentials
    }
}
// Serialize serializes information the current object
func (m *AppManagementConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetKeyCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetKeyCredentials()))
        for i, v := range m.GetKeyCredentials() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("keyCredentials", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPasswordCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPasswordCredentials()))
        for i, v := range m.GetPasswordCredentials() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("passwordCredentials", cast)
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
func (m *AppManagementConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetKeyCredentials sets the keyCredentials property value. Collection of keyCredential restrictions settings to be applied to an application or service principal.
func (m *AppManagementConfiguration) SetKeyCredentials(value []AuthorizationPolicyable)() {
    if m != nil {
        m.keyCredentials = value
    }
}
// SetPasswordCredentials sets the passwordCredentials property value. Collection of password restrictions settings to be applied to an application or service principal.
func (m *AppManagementConfiguration) SetPasswordCredentials(value []AuthorizationPolicyable)() {
    if m != nil {
        m.passwordCredentials = value
    }
}
