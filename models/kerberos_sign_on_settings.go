package models

import (
    iaf507e9fab54dd393f53d684f9aada4c4cc0f1d3f7bf666cebb7853967f69ee2 "kerberossignonsettings"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// KerberosSignOnSettings 
type KerberosSignOnSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Internal Application SPN of the application server. This SPN needs to be in the list of services to which the connector can present delegated credentials.
    kerberosServicePrincipalName *string
    // The Delegated Login Identity for the connector to use on behalf of your users. For more information, see Working with different on-premises and cloud identities . Possible values are: userPrincipalName, onPremisesUserPrincipalName, userPrincipalUsername, onPremisesUserPrincipalUsername, onPremisesSAMAccountName.
    kerberosSignOnMappingAttributeType Applicationsable
}
// Applications union type wrapper for classes kerberosSignOnMappingAttributeType, applicationsMember1
type Applications struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type applicationsMember1
    applicationsMember1 ApplicationsMember1able
    // Union type representation for type kerberosSignOnMappingAttributeType
    kerberosSignOnMappingAttributeType *KerberosSignOnMappingAttributeType
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
    res["kerberosSignOnMappingAttributeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKerberosSignOnMappingAttributeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKerberosSignOnMappingAttributeType(val.(*KerberosSignOnMappingAttributeType))
        }
        return nil
    }
    return res
}
// GetKerberosSignOnMappingAttributeType gets the kerberosSignOnMappingAttributeType property value. Union type representation for type kerberosSignOnMappingAttributeType
func (m *Applications) GetKerberosSignOnMappingAttributeType()(*KerberosSignOnMappingAttributeType) {
    if m == nil {
        return nil
    } else {
        return m.kerberosSignOnMappingAttributeType
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
    if m.GetKerberosSignOnMappingAttributeType() != nil {
        cast := (*m.GetKerberosSignOnMappingAttributeType()).String()
        err := writer.WriteStringValue("kerberosSignOnMappingAttributeType", &cast)
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
// SetKerberosSignOnMappingAttributeType sets the kerberosSignOnMappingAttributeType property value. Union type representation for type kerberosSignOnMappingAttributeType
func (m *Applications) SetKerberosSignOnMappingAttributeType(value *KerberosSignOnMappingAttributeType)() {
    if m != nil {
        m.kerberosSignOnMappingAttributeType = value
    }
}
// Applicationsable 
type Applicationsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationsMember1()(ApplicationsMember1able)
    GetKerberosSignOnMappingAttributeType()(*KerberosSignOnMappingAttributeType)
    SetApplicationsMember1(value ApplicationsMember1able)()
    SetKerberosSignOnMappingAttributeType(value *KerberosSignOnMappingAttributeType)()
}
// NewKerberosSignOnSettings instantiates a new kerberosSignOnSettings and sets the default values.
func NewKerberosSignOnSettings()(*KerberosSignOnSettings) {
    m := &KerberosSignOnSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateKerberosSignOnSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateKerberosSignOnSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKerberosSignOnSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *KerberosSignOnSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *KerberosSignOnSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["kerberosServicePrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKerberosServicePrincipalName(val)
        }
        return nil
    }
    res["kerberosSignOnMappingAttributeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApplicationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKerberosSignOnMappingAttributeType(val.(Applicationsable))
        }
        return nil
    }
    return res
}
// GetKerberosServicePrincipalName gets the kerberosServicePrincipalName property value. The Internal Application SPN of the application server. This SPN needs to be in the list of services to which the connector can present delegated credentials.
func (m *KerberosSignOnSettings) GetKerberosServicePrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.kerberosServicePrincipalName
    }
}
// GetKerberosSignOnMappingAttributeType gets the kerberosSignOnMappingAttributeType property value. The Delegated Login Identity for the connector to use on behalf of your users. For more information, see Working with different on-premises and cloud identities . Possible values are: userPrincipalName, onPremisesUserPrincipalName, userPrincipalUsername, onPremisesUserPrincipalUsername, onPremisesSAMAccountName.
func (m *KerberosSignOnSettings) GetKerberosSignOnMappingAttributeType()(Applicationsable) {
    if m == nil {
        return nil
    } else {
        return m.kerberosSignOnMappingAttributeType
    }
}
// Serialize serializes information the current object
func (m *KerberosSignOnSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("kerberosServicePrincipalName", m.GetKerberosServicePrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("kerberosSignOnMappingAttributeType", m.GetKerberosSignOnMappingAttributeType())
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
func (m *KerberosSignOnSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetKerberosServicePrincipalName sets the kerberosServicePrincipalName property value. The Internal Application SPN of the application server. This SPN needs to be in the list of services to which the connector can present delegated credentials.
func (m *KerberosSignOnSettings) SetKerberosServicePrincipalName(value *string)() {
    if m != nil {
        m.kerberosServicePrincipalName = value
    }
}
// SetKerberosSignOnMappingAttributeType sets the kerberosSignOnMappingAttributeType property value. The Delegated Login Identity for the connector to use on behalf of your users. For more information, see Working with different on-premises and cloud identities . Possible values are: userPrincipalName, onPremisesUserPrincipalName, userPrincipalUsername, onPremisesUserPrincipalUsername, onPremisesSAMAccountName.
func (m *KerberosSignOnSettings) SetKerberosSignOnMappingAttributeType(value Applicationsable)() {
    if m != nil {
        m.kerberosSignOnMappingAttributeType = value
    }
}
