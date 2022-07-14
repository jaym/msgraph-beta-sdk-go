package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleAssignmentable 
type UnifiedRoleAssignmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppScope()(RoleDefinitionsable)
    GetAppScopeId()(*string)
    GetCondition()(*string)
    GetDirectoryScope()(RoleDefinitionsable)
    GetDirectoryScopeId()(*string)
    GetPrincipal()(RoleDefinitionsable)
    GetPrincipalId()(*string)
    GetPrincipalOrganizationId()(*string)
    GetResourceScope()(*string)
    GetRoleDefinition()(RoleDefinitionsable)
    GetRoleDefinitionId()(*string)
    SetAppScope(value RoleDefinitionsable)()
    SetAppScopeId(value *string)()
    SetCondition(value *string)()
    SetDirectoryScope(value RoleDefinitionsable)()
    SetDirectoryScopeId(value *string)()
    SetPrincipal(value RoleDefinitionsable)()
    SetPrincipalId(value *string)()
    SetPrincipalOrganizationId(value *string)()
    SetResourceScope(value *string)()
    SetRoleDefinition(value RoleDefinitionsable)()
    SetRoleDefinitionId(value *string)()
}
