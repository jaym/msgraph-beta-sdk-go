package models

import (
    i3b21d1e8301d1f95f11fc5930515630af3b73b5dc1681d57e13ee16285decd7b "deviceconfigurationassignment"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationAssignment the device configuration assignment entity assigns an AAD group to a specific device configuration.
type DeviceConfigurationAssignment struct {
    Entity
    // The admin intent to apply or remove the profile. Possible values are: apply, remove.
    intent DeviceConfigurationsable
    // Represents source of assignment.
    source *DeviceAndAppManagementAssignmentSource
    // The identifier of the source of the assignment. This property is read-only.
    sourceId *string
    // The assignment target for the device configuration.
    target DeviceConfigurationsable
}
// DeviceConfigurations union type wrapper for classes deviceAndAppManagementAssignmentTarget, deviceConfigurationsMember1
type DeviceConfigurations struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type deviceAndAppManagementAssignmentTarget
    deviceAndAppManagementAssignmentTarget DeviceAndAppManagementAssignmentTargetable
    // Union type representation for type deviceConfigAssignmentIntent
    deviceConfigAssignmentIntent *DeviceConfigAssignmentIntent
    // Union type representation for type deviceConfigurationsMember1
    deviceConfigurationsMember1 DeviceConfigurationsMember1able
}
// NewDeviceConfigurations instantiates a new deviceConfigurations and sets the default values.
func NewDeviceConfigurations()(*DeviceConfigurations) {
    m := &DeviceConfigurations{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceConfigurationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurations(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceConfigurations) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceAndAppManagementAssignmentTarget gets the deviceAndAppManagementAssignmentTarget property value. Union type representation for type deviceAndAppManagementAssignmentTarget
func (m *DeviceConfigurations) GetDeviceAndAppManagementAssignmentTarget()(DeviceAndAppManagementAssignmentTargetable) {
    if m == nil {
        return nil
    } else {
        return m.deviceAndAppManagementAssignmentTarget
    }
}
// GetDeviceConfigAssignmentIntent gets the deviceConfigAssignmentIntent property value. Union type representation for type deviceConfigAssignmentIntent
func (m *DeviceConfigurations) GetDeviceConfigAssignmentIntent()(*DeviceConfigAssignmentIntent) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigAssignmentIntent
    }
}
// GetDeviceConfigurationsMember1 gets the deviceConfigurationsMember1 property value. Union type representation for type deviceConfigurationsMember1
func (m *DeviceConfigurations) GetDeviceConfigurationsMember1()(DeviceConfigurationsMember1able) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationsMember1
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurations) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceAndAppManagementAssignmentTarget"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceAndAppManagementAssignmentTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAndAppManagementAssignmentTarget(val.(DeviceAndAppManagementAssignmentTargetable))
        }
        return nil
    }
    res["deviceConfigAssignmentIntent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceConfigAssignmentIntent)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceConfigAssignmentIntent(val.(*DeviceConfigAssignmentIntent))
        }
        return nil
    }
    res["deviceConfigurationsMember1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceConfigurationsMember1FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceConfigurationsMember1(val.(DeviceConfigurationsMember1able))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceConfigurations) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("deviceAndAppManagementAssignmentTarget", m.GetDeviceAndAppManagementAssignmentTarget())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceConfigAssignmentIntent() != nil {
        cast := (*m.GetDeviceConfigAssignmentIntent()).String()
        err := writer.WriteStringValue("deviceConfigAssignmentIntent", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("deviceConfigurationsMember1", m.GetDeviceConfigurationsMember1())
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
func (m *DeviceConfigurations) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceAndAppManagementAssignmentTarget sets the deviceAndAppManagementAssignmentTarget property value. Union type representation for type deviceAndAppManagementAssignmentTarget
func (m *DeviceConfigurations) SetDeviceAndAppManagementAssignmentTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    if m != nil {
        m.deviceAndAppManagementAssignmentTarget = value
    }
}
// SetDeviceConfigAssignmentIntent sets the deviceConfigAssignmentIntent property value. Union type representation for type deviceConfigAssignmentIntent
func (m *DeviceConfigurations) SetDeviceConfigAssignmentIntent(value *DeviceConfigAssignmentIntent)() {
    if m != nil {
        m.deviceConfigAssignmentIntent = value
    }
}
// SetDeviceConfigurationsMember1 sets the deviceConfigurationsMember1 property value. Union type representation for type deviceConfigurationsMember1
func (m *DeviceConfigurations) SetDeviceConfigurationsMember1(value DeviceConfigurationsMember1able)() {
    if m != nil {
        m.deviceConfigurationsMember1 = value
    }
}
// DeviceConfigurationsable 
type DeviceConfigurationsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceAndAppManagementAssignmentTarget()(DeviceAndAppManagementAssignmentTargetable)
    GetDeviceConfigAssignmentIntent()(*DeviceConfigAssignmentIntent)
    GetDeviceConfigurationsMember1()(DeviceConfigurationsMember1able)
    SetDeviceAndAppManagementAssignmentTarget(value DeviceAndAppManagementAssignmentTargetable)()
    SetDeviceConfigAssignmentIntent(value *DeviceConfigAssignmentIntent)()
    SetDeviceConfigurationsMember1(value DeviceConfigurationsMember1able)()
}
// NewDeviceConfigurationAssignment instantiates a new deviceConfigurationAssignment and sets the default values.
func NewDeviceConfigurationAssignment()(*DeviceConfigurationAssignment) {
    m := &DeviceConfigurationAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceConfigurationAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["intent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceConfigurationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntent(val.(DeviceConfigurationsable))
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAndAppManagementAssignmentSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*DeviceAndAppManagementAssignmentSource))
        }
        return nil
    }
    res["sourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceId(val)
        }
        return nil
    }
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceConfigurationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(DeviceConfigurationsable))
        }
        return nil
    }
    return res
}
// GetIntent gets the intent property value. The admin intent to apply or remove the profile. Possible values are: apply, remove.
func (m *DeviceConfigurationAssignment) GetIntent()(DeviceConfigurationsable) {
    if m == nil {
        return nil
    } else {
        return m.intent
    }
}
// GetSource gets the source property value. Represents source of assignment.
func (m *DeviceConfigurationAssignment) GetSource()(*DeviceAndAppManagementAssignmentSource) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
// GetSourceId gets the sourceId property value. The identifier of the source of the assignment. This property is read-only.
func (m *DeviceConfigurationAssignment) GetSourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceId
    }
}
// GetTarget gets the target property value. The assignment target for the device configuration.
func (m *DeviceConfigurationAssignment) GetTarget()(DeviceConfigurationsable) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// Serialize serializes information the current object
func (m *DeviceConfigurationAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("intent", m.GetIntent())
        if err != nil {
            return err
        }
    }
    if m.GetSource() != nil {
        cast := (*m.GetSource()).String()
        err = writer.WriteStringValue("source", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sourceId", m.GetSourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIntent sets the intent property value. The admin intent to apply or remove the profile. Possible values are: apply, remove.
func (m *DeviceConfigurationAssignment) SetIntent(value DeviceConfigurationsable)() {
    if m != nil {
        m.intent = value
    }
}
// SetSource sets the source property value. Represents source of assignment.
func (m *DeviceConfigurationAssignment) SetSource(value *DeviceAndAppManagementAssignmentSource)() {
    if m != nil {
        m.source = value
    }
}
// SetSourceId sets the sourceId property value. The identifier of the source of the assignment. This property is read-only.
func (m *DeviceConfigurationAssignment) SetSourceId(value *string)() {
    if m != nil {
        m.sourceId = value
    }
}
// SetTarget sets the target property value. The assignment target for the device configuration.
func (m *DeviceConfigurationAssignment) SetTarget(value DeviceConfigurationsable)() {
    if m != nil {
        m.target = value
    }
}
