package models

import (
    i5a4fed406ba750f045425e4c41e27719543a35f9e43a3d81725a953319da3c5d "deviceconfigurationgroupassignment"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationGroupAssignment device configuration group assignment.
type DeviceConfigurationGroupAssignment struct {
    Entity
    // The navigation link to the Device Configuration being targeted.
    deviceConfiguration DeviceConfigurationsable
    // Indicates if this group is should be excluded. Defaults that the group should be included
    excludeGroup *bool
    // The Id of the AAD group we are targeting the device configuration to.
    targetGroupId *string
}
// DeviceConfigurations union type wrapper for classes deviceConfiguration, deviceConfigurationsMember1
type DeviceConfigurations struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type deviceConfiguration
    deviceConfiguration DeviceConfigurationable
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
// GetDeviceConfiguration gets the deviceConfiguration property value. Union type representation for type deviceConfiguration
func (m *DeviceConfigurations) GetDeviceConfiguration()(DeviceConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfiguration
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
    res["deviceConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceConfiguration(val.(DeviceConfigurationable))
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
        err := writer.WriteObjectValue("deviceConfiguration", m.GetDeviceConfiguration())
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
// SetDeviceConfiguration sets the deviceConfiguration property value. Union type representation for type deviceConfiguration
func (m *DeviceConfigurations) SetDeviceConfiguration(value DeviceConfigurationable)() {
    if m != nil {
        m.deviceConfiguration = value
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
    GetDeviceConfiguration()(DeviceConfigurationable)
    GetDeviceConfigurationsMember1()(DeviceConfigurationsMember1able)
    SetDeviceConfiguration(value DeviceConfigurationable)()
    SetDeviceConfigurationsMember1(value DeviceConfigurationsMember1able)()
}
// NewDeviceConfigurationGroupAssignment instantiates a new deviceConfigurationGroupAssignment and sets the default values.
func NewDeviceConfigurationGroupAssignment()(*DeviceConfigurationGroupAssignment) {
    m := &DeviceConfigurationGroupAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceConfigurationGroupAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationGroupAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationGroupAssignment(), nil
}
// GetDeviceConfiguration gets the deviceConfiguration property value. The navigation link to the Device Configuration being targeted.
func (m *DeviceConfigurationGroupAssignment) GetDeviceConfiguration()(DeviceConfigurationsable) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfiguration
    }
}
// GetExcludeGroup gets the excludeGroup property value. Indicates if this group is should be excluded. Defaults that the group should be included
func (m *DeviceConfigurationGroupAssignment) GetExcludeGroup()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.excludeGroup
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationGroupAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceConfigurationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceConfiguration(val.(DeviceConfigurationsable))
        }
        return nil
    }
    res["excludeGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcludeGroup(val)
        }
        return nil
    }
    res["targetGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetGroupId(val)
        }
        return nil
    }
    return res
}
// GetTargetGroupId gets the targetGroupId property value. The Id of the AAD group we are targeting the device configuration to.
func (m *DeviceConfigurationGroupAssignment) GetTargetGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetGroupId
    }
}
// Serialize serializes information the current object
func (m *DeviceConfigurationGroupAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("deviceConfiguration", m.GetDeviceConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("excludeGroup", m.GetExcludeGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetGroupId", m.GetTargetGroupId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceConfiguration sets the deviceConfiguration property value. The navigation link to the Device Configuration being targeted.
func (m *DeviceConfigurationGroupAssignment) SetDeviceConfiguration(value DeviceConfigurationsable)() {
    if m != nil {
        m.deviceConfiguration = value
    }
}
// SetExcludeGroup sets the excludeGroup property value. Indicates if this group is should be excluded. Defaults that the group should be included
func (m *DeviceConfigurationGroupAssignment) SetExcludeGroup(value *bool)() {
    if m != nil {
        m.excludeGroup = value
    }
}
// SetTargetGroupId sets the targetGroupId property value. The Id of the AAD group we are targeting the device configuration to.
func (m *DeviceConfigurationGroupAssignment) SetTargetGroupId(value *string)() {
    if m != nil {
        m.targetGroupId = value
    }
}
