package security

import (
    ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837 "msgraphbetasdkgo/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InformationProtectionable 
type InformationProtectionable interface {
    ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLabelPolicySettings()(InformationProtectionPolicySettingable)
    GetSensitivityLabels()([]SensitivityLabelable)
    SetLabelPolicySettings(value InformationProtectionPolicySettingable)()
    SetSensitivityLabels(value []SensitivityLabelable)()
}
