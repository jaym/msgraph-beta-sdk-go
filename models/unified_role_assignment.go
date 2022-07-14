package models

import (
    i4372703afa69ccb6fa04b11e6dc12cc0dbe7cf1dee8f8a02ccd885769c86b4d2 "unifiedroleassignment"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleAssignment 
type UnifiedRoleAssignment struct {
    Entity
    // Details of the app specific scope when the assignment scope is app specific. Containment entity.
    appScope RoleDefinitionsable
    // Identifier of the app specific scope when the assignment scope is app specific. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. App scopes are scopes that are defined and understood by this application only.  For the entitlement management provider, use app scopes to specify a catalog, for example /AccessPackageCatalog/beedadfe-01d5-4025-910b-84abb9369997.
    appScopeId *string
    // The condition property
    condition *string
    // The directory object that is the scope of the assignment. Provided so that callers can get the directory object using $expand at the same time as getting the role assignment. Read-only. Supports $expand.
    directoryScope RoleDefinitionsable
    // Identifier of the directory object representing the scope of the assignment. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. App scopes are scopes that are defined and understood by this application only.
    directoryScopeId *string
    // The assigned principal. Provided so that callers can get the principal using $expand at the same time as getting the role assignment. Read-only. Supports $expand.
    principal RoleDefinitionsable
    // Identifier of the principal to which the assignment is granted. Supports $filter (eq operator only).
    principalId *string
    // The principalOrganizationId property
    principalOrganizationId *string
    // The scope at which the unifiedRoleAssignment applies. This is / for service-wide. DO NOT USE. This property will be deprecated soon.
    resourceScope *string
    // The roleDefinition the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment. roleDefinition.id will be auto expanded. Supports $expand.
    roleDefinition RoleDefinitionsable
    // Identifier of the unifiedRoleDefinition the assignment is for. Read-only. Supports $filter (eq operator only).
    roleDefinitionId *string
}
// RoleDefinitions union type wrapper for classes unifiedRoleDefinition, roleDefinitionsMember1
type RoleDefinitions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type appScope
    appScope AppScopeable
    // Union type representation for type directoryObject
    directoryObject DirectoryObjectable
    // Union type representation for type roleDefinitionsMember1
    roleDefinitionsMember1 RoleDefinitionsMember1able
    // Union type representation for type unifiedRoleDefinition
    unifiedRoleDefinition UnifiedRoleDefinitionable
}
// NewRoleDefinitions instantiates a new roleDefinitions and sets the default values.
func NewRoleDefinitions()(*RoleDefinitions) {
    m := &RoleDefinitions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRoleDefinitionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleDefinitionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleDefinitions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RoleDefinitions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppScope gets the appScope property value. Union type representation for type appScope
func (m *RoleDefinitions) GetAppScope()(AppScopeable) {
    if m == nil {
        return nil
    } else {
        return m.appScope
    }
}
// GetDirectoryObject gets the directoryObject property value. Union type representation for type directoryObject
func (m *RoleDefinitions) GetDirectoryObject()(DirectoryObjectable) {
    if m == nil {
        return nil
    } else {
        return m.directoryObject
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleDefinitions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppScope(val.(AppScopeable))
        }
        return nil
    }
    res["directoryObject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryObject(val.(DirectoryObjectable))
        }
        return nil
    }
    res["roleDefinitionsMember1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoleDefinitionsMember1FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinitionsMember1(val.(RoleDefinitionsMember1able))
        }
        return nil
    }
    res["unifiedRoleDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUnifiedRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnifiedRoleDefinition(val.(UnifiedRoleDefinitionable))
        }
        return nil
    }
    return res
}
// GetRoleDefinitionsMember1 gets the roleDefinitionsMember1 property value. Union type representation for type roleDefinitionsMember1
func (m *RoleDefinitions) GetRoleDefinitionsMember1()(RoleDefinitionsMember1able) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitionsMember1
    }
}
// GetUnifiedRoleDefinition gets the unifiedRoleDefinition property value. Union type representation for type unifiedRoleDefinition
func (m *RoleDefinitions) GetUnifiedRoleDefinition()(UnifiedRoleDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.unifiedRoleDefinition
    }
}
// Serialize serializes information the current object
func (m *RoleDefinitions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("appScope", m.GetAppScope())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("directoryObject", m.GetDirectoryObject())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("roleDefinitionsMember1", m.GetRoleDefinitionsMember1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("unifiedRoleDefinition", m.GetUnifiedRoleDefinition())
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
func (m *RoleDefinitions) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppScope sets the appScope property value. Union type representation for type appScope
func (m *RoleDefinitions) SetAppScope(value AppScopeable)() {
    if m != nil {
        m.appScope = value
    }
}
// SetDirectoryObject sets the directoryObject property value. Union type representation for type directoryObject
func (m *RoleDefinitions) SetDirectoryObject(value DirectoryObjectable)() {
    if m != nil {
        m.directoryObject = value
    }
}
// SetRoleDefinitionsMember1 sets the roleDefinitionsMember1 property value. Union type representation for type roleDefinitionsMember1
func (m *RoleDefinitions) SetRoleDefinitionsMember1(value RoleDefinitionsMember1able)() {
    if m != nil {
        m.roleDefinitionsMember1 = value
    }
}
// SetUnifiedRoleDefinition sets the unifiedRoleDefinition property value. Union type representation for type unifiedRoleDefinition
func (m *RoleDefinitions) SetUnifiedRoleDefinition(value UnifiedRoleDefinitionable)() {
    if m != nil {
        m.unifiedRoleDefinition = value
    }
}
// RoleDefinitionsable 
type RoleDefinitionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppScope()(AppScopeable)
    GetDirectoryObject()(DirectoryObjectable)
    GetRoleDefinitionsMember1()(RoleDefinitionsMember1able)
    GetUnifiedRoleDefinition()(UnifiedRoleDefinitionable)
    SetAppScope(value AppScopeable)()
    SetDirectoryObject(value DirectoryObjectable)()
    SetRoleDefinitionsMember1(value RoleDefinitionsMember1able)()
    SetUnifiedRoleDefinition(value UnifiedRoleDefinitionable)()
}
// NewUnifiedRoleAssignment instantiates a new UnifiedRoleAssignment and sets the default values.
func NewUnifiedRoleAssignment()(*UnifiedRoleAssignment) {
    m := &UnifiedRoleAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedRoleAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleAssignment(), nil
}
// GetAppScope gets the appScope property value. Details of the app specific scope when the assignment scope is app specific. Containment entity.
func (m *UnifiedRoleAssignment) GetAppScope()(RoleDefinitionsable) {
    if m == nil {
        return nil
    } else {
        return m.appScope
    }
}
// GetAppScopeId gets the appScopeId property value. Identifier of the app specific scope when the assignment scope is app specific. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. App scopes are scopes that are defined and understood by this application only.  For the entitlement management provider, use app scopes to specify a catalog, for example /AccessPackageCatalog/beedadfe-01d5-4025-910b-84abb9369997.
func (m *UnifiedRoleAssignment) GetAppScopeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appScopeId
    }
}
// GetCondition gets the condition property value. The condition property
func (m *UnifiedRoleAssignment) GetCondition()(*string) {
    if m == nil {
        return nil
    } else {
        return m.condition
    }
}
// GetDirectoryScope gets the directoryScope property value. The directory object that is the scope of the assignment. Provided so that callers can get the directory object using $expand at the same time as getting the role assignment. Read-only. Supports $expand.
func (m *UnifiedRoleAssignment) GetDirectoryScope()(RoleDefinitionsable) {
    if m == nil {
        return nil
    } else {
        return m.directoryScope
    }
}
// GetDirectoryScopeId gets the directoryScopeId property value. Identifier of the directory object representing the scope of the assignment. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. App scopes are scopes that are defined and understood by this application only.
func (m *UnifiedRoleAssignment) GetDirectoryScopeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.directoryScopeId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoleDefinitionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppScope(val.(RoleDefinitionsable))
        }
        return nil
    }
    res["appScopeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppScopeId(val)
        }
        return nil
    }
    res["condition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCondition(val)
        }
        return nil
    }
    res["directoryScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoleDefinitionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryScope(val.(RoleDefinitionsable))
        }
        return nil
    }
    res["directoryScopeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryScopeId(val)
        }
        return nil
    }
    res["principal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoleDefinitionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipal(val.(RoleDefinitionsable))
        }
        return nil
    }
    res["principalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalId(val)
        }
        return nil
    }
    res["principalOrganizationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalOrganizationId(val)
        }
        return nil
    }
    res["resourceScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceScope(val)
        }
        return nil
    }
    res["roleDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoleDefinitionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinition(val.(RoleDefinitionsable))
        }
        return nil
    }
    res["roleDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinitionId(val)
        }
        return nil
    }
    return res
}
// GetPrincipal gets the principal property value. The assigned principal. Provided so that callers can get the principal using $expand at the same time as getting the role assignment. Read-only. Supports $expand.
func (m *UnifiedRoleAssignment) GetPrincipal()(RoleDefinitionsable) {
    if m == nil {
        return nil
    } else {
        return m.principal
    }
}
// GetPrincipalId gets the principalId property value. Identifier of the principal to which the assignment is granted. Supports $filter (eq operator only).
func (m *UnifiedRoleAssignment) GetPrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalId
    }
}
// GetPrincipalOrganizationId gets the principalOrganizationId property value. The principalOrganizationId property
func (m *UnifiedRoleAssignment) GetPrincipalOrganizationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalOrganizationId
    }
}
// GetResourceScope gets the resourceScope property value. The scope at which the unifiedRoleAssignment applies. This is / for service-wide. DO NOT USE. This property will be deprecated soon.
func (m *UnifiedRoleAssignment) GetResourceScope()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceScope
    }
}
// GetRoleDefinition gets the roleDefinition property value. The roleDefinition the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment. roleDefinition.id will be auto expanded. Supports $expand.
func (m *UnifiedRoleAssignment) GetRoleDefinition()(RoleDefinitionsable) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinition
    }
}
// GetRoleDefinitionId gets the roleDefinitionId property value. Identifier of the unifiedRoleDefinition the assignment is for. Read-only. Supports $filter (eq operator only).
func (m *UnifiedRoleAssignment) GetRoleDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitionId
    }
}
// Serialize serializes information the current object
func (m *UnifiedRoleAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("appScope", m.GetAppScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appScopeId", m.GetAppScopeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("condition", m.GetCondition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("directoryScope", m.GetDirectoryScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("directoryScopeId", m.GetDirectoryScopeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("principal", m.GetPrincipal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalId", m.GetPrincipalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalOrganizationId", m.GetPrincipalOrganizationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceScope", m.GetResourceScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("roleDefinition", m.GetRoleDefinition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleDefinitionId", m.GetRoleDefinitionId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppScope sets the appScope property value. Details of the app specific scope when the assignment scope is app specific. Containment entity.
func (m *UnifiedRoleAssignment) SetAppScope(value RoleDefinitionsable)() {
    if m != nil {
        m.appScope = value
    }
}
// SetAppScopeId sets the appScopeId property value. Identifier of the app specific scope when the assignment scope is app specific. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. App scopes are scopes that are defined and understood by this application only.  For the entitlement management provider, use app scopes to specify a catalog, for example /AccessPackageCatalog/beedadfe-01d5-4025-910b-84abb9369997.
func (m *UnifiedRoleAssignment) SetAppScopeId(value *string)() {
    if m != nil {
        m.appScopeId = value
    }
}
// SetCondition sets the condition property value. The condition property
func (m *UnifiedRoleAssignment) SetCondition(value *string)() {
    if m != nil {
        m.condition = value
    }
}
// SetDirectoryScope sets the directoryScope property value. The directory object that is the scope of the assignment. Provided so that callers can get the directory object using $expand at the same time as getting the role assignment. Read-only. Supports $expand.
func (m *UnifiedRoleAssignment) SetDirectoryScope(value RoleDefinitionsable)() {
    if m != nil {
        m.directoryScope = value
    }
}
// SetDirectoryScopeId sets the directoryScopeId property value. Identifier of the directory object representing the scope of the assignment. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. App scopes are scopes that are defined and understood by this application only.
func (m *UnifiedRoleAssignment) SetDirectoryScopeId(value *string)() {
    if m != nil {
        m.directoryScopeId = value
    }
}
// SetPrincipal sets the principal property value. The assigned principal. Provided so that callers can get the principal using $expand at the same time as getting the role assignment. Read-only. Supports $expand.
func (m *UnifiedRoleAssignment) SetPrincipal(value RoleDefinitionsable)() {
    if m != nil {
        m.principal = value
    }
}
// SetPrincipalId sets the principalId property value. Identifier of the principal to which the assignment is granted. Supports $filter (eq operator only).
func (m *UnifiedRoleAssignment) SetPrincipalId(value *string)() {
    if m != nil {
        m.principalId = value
    }
}
// SetPrincipalOrganizationId sets the principalOrganizationId property value. The principalOrganizationId property
func (m *UnifiedRoleAssignment) SetPrincipalOrganizationId(value *string)() {
    if m != nil {
        m.principalOrganizationId = value
    }
}
// SetResourceScope sets the resourceScope property value. The scope at which the unifiedRoleAssignment applies. This is / for service-wide. DO NOT USE. This property will be deprecated soon.
func (m *UnifiedRoleAssignment) SetResourceScope(value *string)() {
    if m != nil {
        m.resourceScope = value
    }
}
// SetRoleDefinition sets the roleDefinition property value. The roleDefinition the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment. roleDefinition.id will be auto expanded. Supports $expand.
func (m *UnifiedRoleAssignment) SetRoleDefinition(value RoleDefinitionsable)() {
    if m != nil {
        m.roleDefinition = value
    }
}
// SetRoleDefinitionId sets the roleDefinitionId property value. Identifier of the unifiedRoleDefinition the assignment is for. Read-only. Supports $filter (eq operator only).
func (m *UnifiedRoleAssignment) SetRoleDefinitionId(value *string)() {
    if m != nil {
        m.roleDefinitionId = value
    }
}
