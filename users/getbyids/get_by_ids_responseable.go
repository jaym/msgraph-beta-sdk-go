package getbyids

import (
    ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837 "msgraphbetasdkgo/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetByIdsResponseable 
type GetByIdsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837.DirectoryObjectable)
    SetValue(value []ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837.DirectoryObjectable)()
}
