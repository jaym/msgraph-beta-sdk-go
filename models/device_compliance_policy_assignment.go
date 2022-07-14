package models

import (
    i7471c8cea9dadde74743028149b45b8223bdbd35cf42901be88b32a6115181bb "devicecompliancepolicyassignment"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceCompliancePolicyAssignment device compliance policy assignment.
type DeviceCompliancePolicyAssignment struct {
    Entity
    // Represents source of assignment.
    source *DeviceAndAppManagementAssignmentSource
    // The identifier of the source of the assignment.
    sourceId *string
    // Target for the compliance policy assignment.
    target DeviceCompliancePoliciesable
}
// DeviceCompliancePolicies union type wrapper for classes deviceAndAppManagementAssignmentTarget, deviceCompliancePoliciesMember1
type DeviceCompliancePolicies struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type deviceAndAppManagementAssignmentTarget
    deviceAndAppManagementAssignmentTarget DeviceAndAppManagementAssignmentTargetable
    // Union type representation for type deviceCompliancePoliciesMember1
    deviceCompliancePoliciesMember1 DeviceCompliancePoliciesMember1able
}
// NewDeviceCompliancePolicies instantiates a new deviceCompliancePolicies and sets the default values.
func NewDeviceCompliancePolicies()(*DeviceCompliancePolicies) {
    m := &DeviceCompliancePolicies{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceCompliancePoliciesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCompliancePoliciesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCompliancePolicies(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceCompliancePolicies) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceAndAppManagementAssignmentTarget gets the deviceAndAppManagementAssignmentTarget property value. Union type representation for type deviceAndAppManagementAssignmentTarget
func (m *DeviceCompliancePolicies) GetDeviceAndAppManagementAssignmentTarget()(DeviceAndAppManagementAssignmentTargetable) {
    if m == nil {
        return nil
    } else {
        return m.deviceAndAppManagementAssignmentTarget
    }
}
// GetDeviceCompliancePoliciesMember1 gets the deviceCompliancePoliciesMember1 property value. Union type representation for type deviceCompliancePoliciesMember1
func (m *DeviceCompliancePolicies) GetDeviceCompliancePoliciesMember1()(DeviceCompliancePoliciesMember1able) {
    if m == nil {
        return nil
    } else {
        return m.deviceCompliancePoliciesMember1
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCompliancePolicies) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["deviceCompliancePoliciesMember1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceCompliancePoliciesMember1FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCompliancePoliciesMember1(val.(DeviceCompliancePoliciesMember1able))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceCompliancePolicies) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("deviceAndAppManagementAssignmentTarget", m.GetDeviceAndAppManagementAssignmentTarget())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("deviceCompliancePoliciesMember1", m.GetDeviceCompliancePoliciesMember1())
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
func (m *DeviceCompliancePolicies) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceAndAppManagementAssignmentTarget sets the deviceAndAppManagementAssignmentTarget property value. Union type representation for type deviceAndAppManagementAssignmentTarget
func (m *DeviceCompliancePolicies) SetDeviceAndAppManagementAssignmentTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    if m != nil {
        m.deviceAndAppManagementAssignmentTarget = value
    }
}
// SetDeviceCompliancePoliciesMember1 sets the deviceCompliancePoliciesMember1 property value. Union type representation for type deviceCompliancePoliciesMember1
func (m *DeviceCompliancePolicies) SetDeviceCompliancePoliciesMember1(value DeviceCompliancePoliciesMember1able)() {
    if m != nil {
        m.deviceCompliancePoliciesMember1 = value
    }
}
// DeviceCompliancePoliciesable 
type DeviceCompliancePoliciesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceAndAppManagementAssignmentTarget()(DeviceAndAppManagementAssignmentTargetable)
    GetDeviceCompliancePoliciesMember1()(DeviceCompliancePoliciesMember1able)
    SetDeviceAndAppManagementAssignmentTarget(value DeviceAndAppManagementAssignmentTargetable)()
    SetDeviceCompliancePoliciesMember1(value DeviceCompliancePoliciesMember1able)()
}
// NewDeviceCompliancePolicyAssignment instantiates a new deviceCompliancePolicyAssignment and sets the default values.
func NewDeviceCompliancePolicyAssignment()(*DeviceCompliancePolicyAssignment) {
    m := &DeviceCompliancePolicyAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceCompliancePolicyAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCompliancePolicyAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCompliancePolicyAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCompliancePolicyAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
        val, err := n.GetObjectValue(CreateDeviceCompliancePoliciesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(DeviceCompliancePoliciesable))
        }
        return nil
    }
    return res
}
// GetSource gets the source property value. Represents source of assignment.
func (m *DeviceCompliancePolicyAssignment) GetSource()(*DeviceAndAppManagementAssignmentSource) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
// GetSourceId gets the sourceId property value. The identifier of the source of the assignment.
func (m *DeviceCompliancePolicyAssignment) GetSourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceId
    }
}
// GetTarget gets the target property value. Target for the compliance policy assignment.
func (m *DeviceCompliancePolicyAssignment) GetTarget()(DeviceCompliancePoliciesable) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// Serialize serializes information the current object
func (m *DeviceCompliancePolicyAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
// SetSource sets the source property value. Represents source of assignment.
func (m *DeviceCompliancePolicyAssignment) SetSource(value *DeviceAndAppManagementAssignmentSource)() {
    if m != nil {
        m.source = value
    }
}
// SetSourceId sets the sourceId property value. The identifier of the source of the assignment.
func (m *DeviceCompliancePolicyAssignment) SetSourceId(value *string)() {
    if m != nil {
        m.sourceId = value
    }
}
// SetTarget sets the target property value. Target for the compliance policy assignment.
func (m *DeviceCompliancePolicyAssignment) SetTarget(value DeviceCompliancePoliciesable)() {
    if m != nil {
        m.target = value
    }
}
