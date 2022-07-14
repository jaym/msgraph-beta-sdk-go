package models

import (
    i3e99bea73dbf3dd0a7b65ada7638f529645a432ac74a606cf6a86142f6f24158 "onpremisespublishingsinglesignon"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesPublishingSingleSignOn 
type OnPremisesPublishingSingleSignOn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Kerberos Constrained Delegation settings for applications that use Integrated Window Authentication.
    kerberosSignOnSettings Applicationsable
    // The preferred single-sign on mode for the application. Possible values are: none, onPremisesKerberos, aadHeaderBased,pingHeaderBased.
    singleSignOnMode Applicationsable
}
// Applications union type wrapper for classes singleSignOnMode, applicationsMember1
type Applications struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type applicationsMember1
    applicationsMember1 ApplicationsMember1able
    // Union type representation for type kerberosSignOnSettings
    kerberosSignOnSettings KerberosSignOnSettingsable
    // Union type representation for type singleSignOnMode
    singleSignOnMode *SingleSignOnMode
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
    res["kerberosSignOnSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateKerberosSignOnSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKerberosSignOnSettings(val.(KerberosSignOnSettingsable))
        }
        return nil
    }
    res["singleSignOnMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSingleSignOnMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSignOnMode(val.(*SingleSignOnMode))
        }
        return nil
    }
    return res
}
// GetKerberosSignOnSettings gets the kerberosSignOnSettings property value. Union type representation for type kerberosSignOnSettings
func (m *Applications) GetKerberosSignOnSettings()(KerberosSignOnSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.kerberosSignOnSettings
    }
}
// GetSingleSignOnMode gets the singleSignOnMode property value. Union type representation for type singleSignOnMode
func (m *Applications) GetSingleSignOnMode()(*SingleSignOnMode) {
    if m == nil {
        return nil
    } else {
        return m.singleSignOnMode
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
        err := writer.WriteObjectValue("kerberosSignOnSettings", m.GetKerberosSignOnSettings())
        if err != nil {
            return err
        }
    }
    if m.GetSingleSignOnMode() != nil {
        cast := (*m.GetSingleSignOnMode()).String()
        err := writer.WriteStringValue("singleSignOnMode", &cast)
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
// SetKerberosSignOnSettings sets the kerberosSignOnSettings property value. Union type representation for type kerberosSignOnSettings
func (m *Applications) SetKerberosSignOnSettings(value KerberosSignOnSettingsable)() {
    if m != nil {
        m.kerberosSignOnSettings = value
    }
}
// SetSingleSignOnMode sets the singleSignOnMode property value. Union type representation for type singleSignOnMode
func (m *Applications) SetSingleSignOnMode(value *SingleSignOnMode)() {
    if m != nil {
        m.singleSignOnMode = value
    }
}
// Applicationsable 
type Applicationsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationsMember1()(ApplicationsMember1able)
    GetKerberosSignOnSettings()(KerberosSignOnSettingsable)
    GetSingleSignOnMode()(*SingleSignOnMode)
    SetApplicationsMember1(value ApplicationsMember1able)()
    SetKerberosSignOnSettings(value KerberosSignOnSettingsable)()
    SetSingleSignOnMode(value *SingleSignOnMode)()
}
// NewOnPremisesPublishingSingleSignOn instantiates a new onPremisesPublishingSingleSignOn and sets the default values.
func NewOnPremisesPublishingSingleSignOn()(*OnPremisesPublishingSingleSignOn) {
    m := &OnPremisesPublishingSingleSignOn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOnPremisesPublishingSingleSignOnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesPublishingSingleSignOnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesPublishingSingleSignOn(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesPublishingSingleSignOn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesPublishingSingleSignOn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["kerberosSignOnSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApplicationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKerberosSignOnSettings(val.(Applicationsable))
        }
        return nil
    }
    res["singleSignOnMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApplicationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSignOnMode(val.(Applicationsable))
        }
        return nil
    }
    return res
}
// GetKerberosSignOnSettings gets the kerberosSignOnSettings property value. The Kerberos Constrained Delegation settings for applications that use Integrated Window Authentication.
func (m *OnPremisesPublishingSingleSignOn) GetKerberosSignOnSettings()(Applicationsable) {
    if m == nil {
        return nil
    } else {
        return m.kerberosSignOnSettings
    }
}
// GetSingleSignOnMode gets the singleSignOnMode property value. The preferred single-sign on mode for the application. Possible values are: none, onPremisesKerberos, aadHeaderBased,pingHeaderBased.
func (m *OnPremisesPublishingSingleSignOn) GetSingleSignOnMode()(Applicationsable) {
    if m == nil {
        return nil
    } else {
        return m.singleSignOnMode
    }
}
// Serialize serializes information the current object
func (m *OnPremisesPublishingSingleSignOn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("kerberosSignOnSettings", m.GetKerberosSignOnSettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("singleSignOnMode", m.GetSingleSignOnMode())
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
func (m *OnPremisesPublishingSingleSignOn) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetKerberosSignOnSettings sets the kerberosSignOnSettings property value. The Kerberos Constrained Delegation settings for applications that use Integrated Window Authentication.
func (m *OnPremisesPublishingSingleSignOn) SetKerberosSignOnSettings(value Applicationsable)() {
    if m != nil {
        m.kerberosSignOnSettings = value
    }
}
// SetSingleSignOnMode sets the singleSignOnMode property value. The preferred single-sign on mode for the application. Possible values are: none, onPremisesKerberos, aadHeaderBased,pingHeaderBased.
func (m *OnPremisesPublishingSingleSignOn) SetSingleSignOnMode(value Applicationsable)() {
    if m != nil {
        m.singleSignOnMode = value
    }
}
