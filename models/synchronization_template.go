package models

import (
    ie48d5ed2cc2fb4d9ec5d97eb5b0d1efa799b1a66e5556abb92853d79f7a4526c "synchronizationtemplate"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationTemplate 
type SynchronizationTemplate struct {
    Entity
    // Identifier of the application this template belongs to.
    applicationId *string
    // true if this template is recommended to be the default for the application.
    default_escaped *bool
    // Description of the template.
    description *string
    // true if this template should appear in the collection of templates available for the application instance (service principal).
    discoverable *bool
    // One of the well-known factory tags supported by the synchronization engine. The factoryTag tells the synchronization engine which implementation to use when processing jobs based on this template.
    factoryTag *string
    // Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
    metadata []RoleDefinitionsable
    // Default synchronization schema for the jobs based on this template.
    schema RoleDefinitionsable
}
// RoleDefinitions union type wrapper for classes synchronizationSchema, roleDefinitionsMember1
type RoleDefinitions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type metadataEntry
    metadataEntry MetadataEntryable
    // Union type representation for type roleDefinitionsMember1
    roleDefinitionsMember1 RoleDefinitionsMember1able
    // Union type representation for type synchronizationSchema
    synchronizationSchema SynchronizationSchemaable
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
    res["metadataEntry"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMetadataEntryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadataEntry(val.(MetadataEntryable))
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
    res["synchronizationSchema"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationSchemaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSynchronizationSchema(val.(SynchronizationSchemaable))
        }
        return nil
    }
    return res
}
// GetMetadataEntry gets the metadataEntry property value. Union type representation for type metadataEntry
func (m *RoleDefinitions) GetMetadataEntry()(MetadataEntryable) {
    if m == nil {
        return nil
    } else {
        return m.metadataEntry
    }
}
// GetRoleDefinitionsMember1 gets the roleDefinitionsMember1 property value. Union type representation for type roleDefinitionsMember1
func (m *RoleDefinitions) GetRoleDefinitionsMember1()(RoleDefinitionsMember1able) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitionsMember1
    }
}
// GetSynchronizationSchema gets the synchronizationSchema property value. Union type representation for type synchronizationSchema
func (m *RoleDefinitions) GetSynchronizationSchema()(SynchronizationSchemaable) {
    if m == nil {
        return nil
    } else {
        return m.synchronizationSchema
    }
}
// Serialize serializes information the current object
func (m *RoleDefinitions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("metadataEntry", m.GetMetadataEntry())
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
        err := writer.WriteObjectValue("synchronizationSchema", m.GetSynchronizationSchema())
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
// SetMetadataEntry sets the metadataEntry property value. Union type representation for type metadataEntry
func (m *RoleDefinitions) SetMetadataEntry(value MetadataEntryable)() {
    if m != nil {
        m.metadataEntry = value
    }
}
// SetRoleDefinitionsMember1 sets the roleDefinitionsMember1 property value. Union type representation for type roleDefinitionsMember1
func (m *RoleDefinitions) SetRoleDefinitionsMember1(value RoleDefinitionsMember1able)() {
    if m != nil {
        m.roleDefinitionsMember1 = value
    }
}
// SetSynchronizationSchema sets the synchronizationSchema property value. Union type representation for type synchronizationSchema
func (m *RoleDefinitions) SetSynchronizationSchema(value SynchronizationSchemaable)() {
    if m != nil {
        m.synchronizationSchema = value
    }
}
// RoleDefinitionsable 
type RoleDefinitionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMetadataEntry()(MetadataEntryable)
    GetRoleDefinitionsMember1()(RoleDefinitionsMember1able)
    GetSynchronizationSchema()(SynchronizationSchemaable)
    SetMetadataEntry(value MetadataEntryable)()
    SetRoleDefinitionsMember1(value RoleDefinitionsMember1able)()
    SetSynchronizationSchema(value SynchronizationSchemaable)()
}
// NewSynchronizationTemplate instantiates a new SynchronizationTemplate and sets the default values.
func NewSynchronizationTemplate()(*SynchronizationTemplate) {
    m := &SynchronizationTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSynchronizationTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationTemplate(), nil
}
// GetApplicationId gets the applicationId property value. Identifier of the application this template belongs to.
func (m *SynchronizationTemplate) GetApplicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationId
    }
}
// GetDefault gets the default property value. true if this template is recommended to be the default for the application.
func (m *SynchronizationTemplate) GetDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.default_escaped
    }
}
// GetDescription gets the description property value. Description of the template.
func (m *SynchronizationTemplate) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDiscoverable gets the discoverable property value. true if this template should appear in the collection of templates available for the application instance (service principal).
func (m *SynchronizationTemplate) GetDiscoverable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.discoverable
    }
}
// GetFactoryTag gets the factoryTag property value. One of the well-known factory tags supported by the synchronization engine. The factoryTag tells the synchronization engine which implementation to use when processing jobs based on this template.
func (m *SynchronizationTemplate) GetFactoryTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.factoryTag
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationId(val)
        }
        return nil
    }
    res["default"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefault(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["discoverable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscoverable(val)
        }
        return nil
    }
    res["factoryTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFactoryTag(val)
        }
        return nil
    }
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRoleDefinitionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleDefinitionsable, len(val))
            for i, v := range val {
                res[i] = v.(RoleDefinitionsable)
            }
            m.SetMetadata(res)
        }
        return nil
    }
    res["schema"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoleDefinitionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchema(val.(RoleDefinitionsable))
        }
        return nil
    }
    return res
}
// GetMetadata gets the metadata property value. Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
func (m *SynchronizationTemplate) GetMetadata()([]RoleDefinitionsable) {
    if m == nil {
        return nil
    } else {
        return m.metadata
    }
}
// GetSchema gets the schema property value. Default synchronization schema for the jobs based on this template.
func (m *SynchronizationTemplate) GetSchema()(RoleDefinitionsable) {
    if m == nil {
        return nil
    } else {
        return m.schema
    }
}
// Serialize serializes information the current object
func (m *SynchronizationTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("applicationId", m.GetApplicationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("default", m.GetDefault())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("discoverable", m.GetDiscoverable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("factoryTag", m.GetFactoryTag())
        if err != nil {
            return err
        }
    }
    if m.GetMetadata() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMetadata()))
        for i, v := range m.GetMetadata() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("metadata", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("schema", m.GetSchema())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationId sets the applicationId property value. Identifier of the application this template belongs to.
func (m *SynchronizationTemplate) SetApplicationId(value *string)() {
    if m != nil {
        m.applicationId = value
    }
}
// SetDefault sets the default property value. true if this template is recommended to be the default for the application.
func (m *SynchronizationTemplate) SetDefault(value *bool)() {
    if m != nil {
        m.default_escaped = value
    }
}
// SetDescription sets the description property value. Description of the template.
func (m *SynchronizationTemplate) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDiscoverable sets the discoverable property value. true if this template should appear in the collection of templates available for the application instance (service principal).
func (m *SynchronizationTemplate) SetDiscoverable(value *bool)() {
    if m != nil {
        m.discoverable = value
    }
}
// SetFactoryTag sets the factoryTag property value. One of the well-known factory tags supported by the synchronization engine. The factoryTag tells the synchronization engine which implementation to use when processing jobs based on this template.
func (m *SynchronizationTemplate) SetFactoryTag(value *string)() {
    if m != nil {
        m.factoryTag = value
    }
}
// SetMetadata sets the metadata property value. Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
func (m *SynchronizationTemplate) SetMetadata(value []RoleDefinitionsable)() {
    if m != nil {
        m.metadata = value
    }
}
// SetSchema sets the schema property value. Default synchronization schema for the jobs based on this template.
func (m *SynchronizationTemplate) SetSchema(value RoleDefinitionsable)() {
    if m != nil {
        m.schema = value
    }
}
