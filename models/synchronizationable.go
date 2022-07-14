package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Synchronizationable 
type Synchronizationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetJobs()([]SynchronizationJobable)
    GetSecrets()([]RoleAssignmentsable)
    GetTemplates()([]SynchronizationTemplateable)
    SetJobs(value []SynchronizationJobable)()
    SetSecrets(value []RoleAssignmentsable)()
    SetTemplates(value []SynchronizationTemplateable)()
}
