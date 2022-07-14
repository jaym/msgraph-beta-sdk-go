package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ObjectMappingable 
type ObjectMappingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttributeMappings()([]RoleAssignmentsable)
    GetEnabled()(*bool)
    GetFlowTypes()(*ObjectFlowTypes)
    GetMetadata()([]RoleAssignmentsable)
    GetName()(*string)
    GetScope()(RoleAssignmentsable)
    GetSourceObjectName()(*string)
    GetTargetObjectName()(*string)
    SetAttributeMappings(value []RoleAssignmentsable)()
    SetEnabled(value *bool)()
    SetFlowTypes(value *ObjectFlowTypes)()
    SetMetadata(value []RoleAssignmentsable)()
    SetName(value *string)()
    SetScope(value RoleAssignmentsable)()
    SetSourceObjectName(value *string)()
    SetTargetObjectName(value *string)()
}
