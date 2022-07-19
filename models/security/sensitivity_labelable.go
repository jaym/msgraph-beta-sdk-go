package security

import (
    ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837 "msgraphbetasdkgo/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SensitivityLabelable 
type SensitivityLabelable interface {
    ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColor()(*string)
    GetContentFormats()([]string)
    GetDescription()(*string)
    GetHasProtection()(*bool)
    GetIsActive()(*bool)
    GetIsAppliable()(*bool)
    GetName()(*string)
    GetParent()(SensitivityLabelable)
    GetSensitivity()(*int32)
    GetTooltip()(*string)
    SetColor(value *string)()
    SetContentFormats(value []string)()
    SetDescription(value *string)()
    SetHasProtection(value *bool)()
    SetIsActive(value *bool)()
    SetIsAppliable(value *bool)()
    SetName(value *string)()
    SetParent(value SensitivityLabelable)()
    SetSensitivity(value *int32)()
    SetTooltip(value *string)()
}
