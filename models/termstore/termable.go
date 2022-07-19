package termstore

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837 "msgraphbetasdkgo/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Termable 
type Termable interface {
    ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChildren()([]Termable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescriptions()([]LocalizedDescriptionable)
    GetLabels()([]LocalizedLabelable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetProperties()([]ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837.KeyValueable)
    GetRelations()([]Relationable)
    GetSet()(Setable)
    SetChildren(value []Termable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescriptions(value []LocalizedDescriptionable)()
    SetLabels(value []LocalizedLabelable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetProperties(value []ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837.KeyValueable)()
    SetRelations(value []Relationable)()
    SetSet(value Setable)()
}
