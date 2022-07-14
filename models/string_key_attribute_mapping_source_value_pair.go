package models

import (
    ia6e79e8e0310d4a0e52f9e35c123f9e74a86f4fed8ad2b963611107dc7f76859 "stringkeyattributemappingsourcevaluepair"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StringKeyAttributeMappingSourceValuePair 
type StringKeyAttributeMappingSourceValuePair struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The name of the parameter.
    key *string
    // The value of the parameter.
    value RoleAssignmentsable
}
// RoleAssignments union type wrapper for classes attributeMappingSource, roleAssignmentsMember1
type RoleAssignments struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type attributeMappingSource
    attributeMappingSource AttributeMappingSourceable
    // Union type representation for type roleAssignmentsMember1
    roleAssignmentsMember1 RoleAssignmentsMember1able
}
// NewRoleAssignments instantiates a new roleAssignments and sets the default values.
func NewRoleAssignments()(*RoleAssignments) {
    m := &RoleAssignments{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRoleAssignmentsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleAssignmentsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleAssignments(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RoleAssignments) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAttributeMappingSource gets the attributeMappingSource property value. Union type representation for type attributeMappingSource
func (m *RoleAssignments) GetAttributeMappingSource()(AttributeMappingSourceable) {
    if m == nil {
        return nil
    } else {
        return m.attributeMappingSource
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleAssignments) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attributeMappingSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAttributeMappingSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttributeMappingSource(val.(AttributeMappingSourceable))
        }
        return nil
    }
    res["roleAssignmentsMember1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoleAssignmentsMember1FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleAssignmentsMember1(val.(RoleAssignmentsMember1able))
        }
        return nil
    }
    return res
}
// GetRoleAssignmentsMember1 gets the roleAssignmentsMember1 property value. Union type representation for type roleAssignmentsMember1
func (m *RoleAssignments) GetRoleAssignmentsMember1()(RoleAssignmentsMember1able) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentsMember1
    }
}
// Serialize serializes information the current object
func (m *RoleAssignments) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attributeMappingSource", m.GetAttributeMappingSource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("roleAssignmentsMember1", m.GetRoleAssignmentsMember1())
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
func (m *RoleAssignments) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAttributeMappingSource sets the attributeMappingSource property value. Union type representation for type attributeMappingSource
func (m *RoleAssignments) SetAttributeMappingSource(value AttributeMappingSourceable)() {
    if m != nil {
        m.attributeMappingSource = value
    }
}
// SetRoleAssignmentsMember1 sets the roleAssignmentsMember1 property value. Union type representation for type roleAssignmentsMember1
func (m *RoleAssignments) SetRoleAssignmentsMember1(value RoleAssignmentsMember1able)() {
    if m != nil {
        m.roleAssignmentsMember1 = value
    }
}
// RoleAssignmentsable 
type RoleAssignmentsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttributeMappingSource()(AttributeMappingSourceable)
    GetRoleAssignmentsMember1()(RoleAssignmentsMember1able)
    SetAttributeMappingSource(value AttributeMappingSourceable)()
    SetRoleAssignmentsMember1(value RoleAssignmentsMember1able)()
}
// NewStringKeyAttributeMappingSourceValuePair instantiates a new stringKeyAttributeMappingSourceValuePair and sets the default values.
func NewStringKeyAttributeMappingSourceValuePair()(*StringKeyAttributeMappingSourceValuePair) {
    m := &StringKeyAttributeMappingSourceValuePair{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateStringKeyAttributeMappingSourceValuePairFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStringKeyAttributeMappingSourceValuePairFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStringKeyAttributeMappingSourceValuePair(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StringKeyAttributeMappingSourceValuePair) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StringKeyAttributeMappingSourceValuePair) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoleAssignmentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(RoleAssignmentsable))
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. The name of the parameter.
func (m *StringKeyAttributeMappingSourceValuePair) GetKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
// GetValue gets the value property value. The value of the parameter.
func (m *StringKeyAttributeMappingSourceValuePair) GetValue()(RoleAssignmentsable) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// Serialize serializes information the current object
func (m *StringKeyAttributeMappingSourceValuePair) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("value", m.GetValue())
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
func (m *StringKeyAttributeMappingSourceValuePair) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetKey sets the key property value. The name of the parameter.
func (m *StringKeyAttributeMappingSourceValuePair) SetKey(value *string)() {
    if m != nil {
        m.key = value
    }
}
// SetValue sets the value property value. The value of the parameter.
func (m *StringKeyAttributeMappingSourceValuePair) SetValue(value RoleAssignmentsable)() {
    if m != nil {
        m.value = value
    }
}
