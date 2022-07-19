package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAndAppManagementAssignmentTarget base type for assignment targets.
type DeviceAndAppManagementAssignmentTarget struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Id of the filter for the target assignment.
    deviceAndAppManagementAssignmentFilterId *string
    // Represents type of the assignment filter.
    deviceAndAppManagementAssignmentFilterType *DeviceAndAppManagementAssignmentFilterType
    // The OdataType property
    odataType *string
}
// NewDeviceAndAppManagementAssignmentTarget instantiates a new deviceAndAppManagementAssignmentTarget and sets the default values.
func NewDeviceAndAppManagementAssignmentTarget()(*DeviceAndAppManagementAssignmentTarget) {
    m := &DeviceAndAppManagementAssignmentTarget{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.deviceAndAppManagementAssignmentTarget";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceAndAppManagementAssignmentTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAndAppManagementAssignmentTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAndAppManagementAssignmentTarget(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAndAppManagementAssignmentTarget) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceAndAppManagementAssignmentFilterId gets the deviceAndAppManagementAssignmentFilterId property value. The Id of the filter for the target assignment.
func (m *DeviceAndAppManagementAssignmentTarget) GetDeviceAndAppManagementAssignmentFilterId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceAndAppManagementAssignmentFilterId
    }
}
// GetDeviceAndAppManagementAssignmentFilterType gets the deviceAndAppManagementAssignmentFilterType property value. Represents type of the assignment filter.
func (m *DeviceAndAppManagementAssignmentTarget) GetDeviceAndAppManagementAssignmentFilterType()(*DeviceAndAppManagementAssignmentFilterType) {
    if m == nil {
        return nil
    } else {
        return m.deviceAndAppManagementAssignmentFilterType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAndAppManagementAssignmentTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceAndAppManagementAssignmentFilterId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAndAppManagementAssignmentFilterId(val)
        }
        return nil
    }
    res["deviceAndAppManagementAssignmentFilterType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAndAppManagementAssignmentFilterType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAndAppManagementAssignmentFilterType(val.(*DeviceAndAppManagementAssignmentFilterType))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceAndAppManagementAssignmentTarget) GetOdataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.odataType
    }
}
// Serialize serializes information the current object
func (m *DeviceAndAppManagementAssignmentTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceAndAppManagementAssignmentFilterId", m.GetDeviceAndAppManagementAssignmentFilterId())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceAndAppManagementAssignmentFilterType() != nil {
        cast := (*m.GetDeviceAndAppManagementAssignmentFilterType()).String()
        err := writer.WriteStringValue("deviceAndAppManagementAssignmentFilterType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
func (m *DeviceAndAppManagementAssignmentTarget) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceAndAppManagementAssignmentFilterId sets the deviceAndAppManagementAssignmentFilterId property value. The Id of the filter for the target assignment.
func (m *DeviceAndAppManagementAssignmentTarget) SetDeviceAndAppManagementAssignmentFilterId(value *string)() {
    if m != nil {
        m.deviceAndAppManagementAssignmentFilterId = value
    }
}
// SetDeviceAndAppManagementAssignmentFilterType sets the deviceAndAppManagementAssignmentFilterType property value. Represents type of the assignment filter.
func (m *DeviceAndAppManagementAssignmentTarget) SetDeviceAndAppManagementAssignmentFilterType(value *DeviceAndAppManagementAssignmentFilterType)() {
    if m != nil {
        m.deviceAndAppManagementAssignmentFilterType = value
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceAndAppManagementAssignmentTarget) SetOdataType(value *string)() {
    if m != nil {
        m.odataType = value
    }
}
