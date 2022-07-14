package models

import (
    i4dc2369546b1e74cf47f54e42619f3c9416a331be3c4b3d6d426ebe5b9e7aabd "defaultuserroleoverride"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DefaultUserRoleOverride 
type DefaultUserRoleOverride struct {
    Entity
    // The isDefault property
    isDefault *bool
    // The rolePermissions property
    rolePermissions []RoleDefinitionsable
}
// RoleDefinitions union type wrapper for classes unifiedRolePermission, roleDefinitionsMember1
type RoleDefinitions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type roleDefinitionsMember1
    roleDefinitionsMember1 RoleDefinitionsMember1able
    // Union type representation for type unifiedRolePermission
    unifiedRolePermission UnifiedRolePermissionable
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
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleDefinitions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["unifiedRolePermission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUnifiedRolePermissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnifiedRolePermission(val.(UnifiedRolePermissionable))
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
// GetUnifiedRolePermission gets the unifiedRolePermission property value. Union type representation for type unifiedRolePermission
func (m *RoleDefinitions) GetUnifiedRolePermission()(UnifiedRolePermissionable) {
    if m == nil {
        return nil
    } else {
        return m.unifiedRolePermission
    }
}
// Serialize serializes information the current object
func (m *RoleDefinitions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("roleDefinitionsMember1", m.GetRoleDefinitionsMember1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("unifiedRolePermission", m.GetUnifiedRolePermission())
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
// SetRoleDefinitionsMember1 sets the roleDefinitionsMember1 property value. Union type representation for type roleDefinitionsMember1
func (m *RoleDefinitions) SetRoleDefinitionsMember1(value RoleDefinitionsMember1able)() {
    if m != nil {
        m.roleDefinitionsMember1 = value
    }
}
// SetUnifiedRolePermission sets the unifiedRolePermission property value. Union type representation for type unifiedRolePermission
func (m *RoleDefinitions) SetUnifiedRolePermission(value UnifiedRolePermissionable)() {
    if m != nil {
        m.unifiedRolePermission = value
    }
}
// RoleDefinitionsable 
type RoleDefinitionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRoleDefinitionsMember1()(RoleDefinitionsMember1able)
    GetUnifiedRolePermission()(UnifiedRolePermissionable)
    SetRoleDefinitionsMember1(value RoleDefinitionsMember1able)()
    SetUnifiedRolePermission(value UnifiedRolePermissionable)()
}
// NewDefaultUserRoleOverride instantiates a new DefaultUserRoleOverride and sets the default values.
func NewDefaultUserRoleOverride()(*DefaultUserRoleOverride) {
    m := &DefaultUserRoleOverride{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDefaultUserRoleOverrideFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDefaultUserRoleOverrideFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDefaultUserRoleOverride(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DefaultUserRoleOverride) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["rolePermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRoleDefinitionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleDefinitionsable, len(val))
            for i, v := range val {
                res[i] = v.(RoleDefinitionsable)
            }
            m.SetRolePermissions(res)
        }
        return nil
    }
    return res
}
// GetIsDefault gets the isDefault property value. The isDefault property
func (m *DefaultUserRoleOverride) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// GetRolePermissions gets the rolePermissions property value. The rolePermissions property
func (m *DefaultUserRoleOverride) GetRolePermissions()([]RoleDefinitionsable) {
    if m == nil {
        return nil
    } else {
        return m.rolePermissions
    }
}
// Serialize serializes information the current object
func (m *DefaultUserRoleOverride) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    if m.GetRolePermissions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRolePermissions()))
        for i, v := range m.GetRolePermissions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("rolePermissions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsDefault sets the isDefault property value. The isDefault property
func (m *DefaultUserRoleOverride) SetIsDefault(value *bool)() {
    if m != nil {
        m.isDefault = value
    }
}
// SetRolePermissions sets the rolePermissions property value. The rolePermissions property
func (m *DefaultUserRoleOverride) SetRolePermissions(value []RoleDefinitionsable)() {
    if m != nil {
        m.rolePermissions = value
    }
}
