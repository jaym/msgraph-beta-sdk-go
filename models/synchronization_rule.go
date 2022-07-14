package models

import (
    iff24d2e6d08fb0746d44d7db432a679464a6202041af277015d47f4a7c86632d "synchronizationrule"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationRule 
type SynchronizationRule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // true if the synchronization rule can be customized; false if this rule is read-only and should not be changed.
    editable *bool
    // Synchronization rule identifier. Must be one of the identifiers recognized by the synchronization engine. Supported rule identifiers can be found in the synchronization template returned by the API.
    id *string
    // Additional extension properties. Unless instructed explicitly by the support team, metadata values should not be changed.
    metadata []RoleAssignmentsable
    // Human-readable name of the synchronization rule. Not nullable.
    name *string
    // Collection of object mappings supported by the rule. Tells the synchronization engine which objects should be synchronized.
    objectMappings []RoleAssignmentsable
    // Priority relative to other rules in the synchronizationSchema. Rules with the lowest priority number will be processed first.
    priority *int32
    // Name of the source directory. Must match one of the directory definitions in synchronizationSchema.
    sourceDirectoryName *string
    // Name of the target directory. Must match one of the directory definitions in synchronizationSchema.
    targetDirectoryName *string
}
// RoleAssignments union type wrapper for classes stringKeyStringValuePair, roleAssignmentsMember1
type RoleAssignments struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type objectMapping
    objectMapping ObjectMappingable
    // Union type representation for type roleAssignmentsMember1
    roleAssignmentsMember1 RoleAssignmentsMember1able
    // Union type representation for type stringKeyStringValuePair
    stringKeyStringValuePair StringKeyStringValuePairable
}
// NewRoleAssignments instantiates a new roleAssignments and sets the default values.
func NewRoleAssignments()(*RoleAssignments) {
    m := &RoleAssignments{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRoleAssignmentsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleAssignmentsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleAssignments(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RoleAssignments) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleAssignments) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["objectMapping"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateObjectMappingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectMapping(val.(ObjectMappingable))
        }
        return nil
    }
    res["roleAssignmentsMember1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoleAssignmentsMember1FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleAssignmentsMember1(val.(RoleAssignmentsMember1able))
        }
        return nil
    }
    res["stringKeyStringValuePair"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStringKeyStringValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStringKeyStringValuePair(val.(StringKeyStringValuePairable))
        }
        return nil
    }
    return res
}
// GetObjectMapping gets the objectMapping property value. Union type representation for type objectMapping
func (m *RoleAssignments) GetObjectMapping()(ObjectMappingable) {
    if m == nil {
        return nil
    } else {
        return m.objectMapping
    }
}
// GetRoleAssignmentsMember1 gets the roleAssignmentsMember1 property value. Union type representation for type roleAssignmentsMember1
func (m *RoleAssignments) GetRoleAssignmentsMember1()(RoleAssignmentsMember1able) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentsMember1
    }
}
// GetStringKeyStringValuePair gets the stringKeyStringValuePair property value. Union type representation for type stringKeyStringValuePair
func (m *RoleAssignments) GetStringKeyStringValuePair()(StringKeyStringValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.stringKeyStringValuePair
    }
}
// Serialize serializes information the current object
func (m *RoleAssignments) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("objectMapping", m.GetObjectMapping())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("roleAssignmentsMember1", m.GetRoleAssignmentsMember1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("stringKeyStringValuePair", m.GetStringKeyStringValuePair())
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
func (m *RoleAssignments) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetObjectMapping sets the objectMapping property value. Union type representation for type objectMapping
func (m *RoleAssignments) SetObjectMapping(value ObjectMappingable)() {
    if m != nil {
        m.objectMapping = value
    }
}
// SetRoleAssignmentsMember1 sets the roleAssignmentsMember1 property value. Union type representation for type roleAssignmentsMember1
func (m *RoleAssignments) SetRoleAssignmentsMember1(value RoleAssignmentsMember1able)() {
    if m != nil {
        m.roleAssignmentsMember1 = value
    }
}
// SetStringKeyStringValuePair sets the stringKeyStringValuePair property value. Union type representation for type stringKeyStringValuePair
func (m *RoleAssignments) SetStringKeyStringValuePair(value StringKeyStringValuePairable)() {
    if m != nil {
        m.stringKeyStringValuePair = value
    }
}
// RoleAssignmentsable 
type RoleAssignmentsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetObjectMapping()(ObjectMappingable)
    GetRoleAssignmentsMember1()(RoleAssignmentsMember1able)
    GetStringKeyStringValuePair()(StringKeyStringValuePairable)
    SetObjectMapping(value ObjectMappingable)()
    SetRoleAssignmentsMember1(value RoleAssignmentsMember1able)()
    SetStringKeyStringValuePair(value StringKeyStringValuePairable)()
}
// NewSynchronizationRule instantiates a new synchronizationRule and sets the default values.
func NewSynchronizationRule()(*SynchronizationRule) {
    m := &SynchronizationRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSynchronizationRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationRule(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationRule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEditable gets the editable property value. true if the synchronization rule can be customized; false if this rule is read-only and should not be changed.
func (m *SynchronizationRule) GetEditable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.editable
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["editable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEditable(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRoleAssignmentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleAssignmentsable, len(val))
            for i, v := range val {
                res[i] = v.(RoleAssignmentsable)
            }
            m.SetMetadata(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["objectMappings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRoleAssignmentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleAssignmentsable, len(val))
            for i, v := range val {
                res[i] = v.(RoleAssignmentsable)
            }
            m.SetObjectMappings(res)
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["sourceDirectoryName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceDirectoryName(val)
        }
        return nil
    }
    res["targetDirectoryName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetDirectoryName(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Synchronization rule identifier. Must be one of the identifiers recognized by the synchronization engine. Supported rule identifiers can be found in the synchronization template returned by the API.
func (m *SynchronizationRule) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetMetadata gets the metadata property value. Additional extension properties. Unless instructed explicitly by the support team, metadata values should not be changed.
func (m *SynchronizationRule) GetMetadata()([]RoleAssignmentsable) {
    if m == nil {
        return nil
    } else {
        return m.metadata
    }
}
// GetName gets the name property value. Human-readable name of the synchronization rule. Not nullable.
func (m *SynchronizationRule) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetObjectMappings gets the objectMappings property value. Collection of object mappings supported by the rule. Tells the synchronization engine which objects should be synchronized.
func (m *SynchronizationRule) GetObjectMappings()([]RoleAssignmentsable) {
    if m == nil {
        return nil
    } else {
        return m.objectMappings
    }
}
// GetPriority gets the priority property value. Priority relative to other rules in the synchronizationSchema. Rules with the lowest priority number will be processed first.
func (m *SynchronizationRule) GetPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
// GetSourceDirectoryName gets the sourceDirectoryName property value. Name of the source directory. Must match one of the directory definitions in synchronizationSchema.
func (m *SynchronizationRule) GetSourceDirectoryName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceDirectoryName
    }
}
// GetTargetDirectoryName gets the targetDirectoryName property value. Name of the target directory. Must match one of the directory definitions in synchronizationSchema.
func (m *SynchronizationRule) GetTargetDirectoryName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetDirectoryName
    }
}
// Serialize serializes information the current object
func (m *SynchronizationRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("editable", m.GetEditable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    if m.GetMetadata() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMetadata()))
        for i, v := range m.GetMetadata() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("metadata", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetObjectMappings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetObjectMappings()))
        for i, v := range m.GetObjectMappings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("objectMappings", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceDirectoryName", m.GetSourceDirectoryName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetDirectoryName", m.GetTargetDirectoryName())
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
func (m *SynchronizationRule) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEditable sets the editable property value. true if the synchronization rule can be customized; false if this rule is read-only and should not be changed.
func (m *SynchronizationRule) SetEditable(value *bool)() {
    if m != nil {
        m.editable = value
    }
}
// SetId sets the id property value. Synchronization rule identifier. Must be one of the identifiers recognized by the synchronization engine. Supported rule identifiers can be found in the synchronization template returned by the API.
func (m *SynchronizationRule) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetMetadata sets the metadata property value. Additional extension properties. Unless instructed explicitly by the support team, metadata values should not be changed.
func (m *SynchronizationRule) SetMetadata(value []RoleAssignmentsable)() {
    if m != nil {
        m.metadata = value
    }
}
// SetName sets the name property value. Human-readable name of the synchronization rule. Not nullable.
func (m *SynchronizationRule) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetObjectMappings sets the objectMappings property value. Collection of object mappings supported by the rule. Tells the synchronization engine which objects should be synchronized.
func (m *SynchronizationRule) SetObjectMappings(value []RoleAssignmentsable)() {
    if m != nil {
        m.objectMappings = value
    }
}
// SetPriority sets the priority property value. Priority relative to other rules in the synchronizationSchema. Rules with the lowest priority number will be processed first.
func (m *SynchronizationRule) SetPriority(value *int32)() {
    if m != nil {
        m.priority = value
    }
}
// SetSourceDirectoryName sets the sourceDirectoryName property value. Name of the source directory. Must match one of the directory definitions in synchronizationSchema.
func (m *SynchronizationRule) SetSourceDirectoryName(value *string)() {
    if m != nil {
        m.sourceDirectoryName = value
    }
}
// SetTargetDirectoryName sets the targetDirectoryName property value. Name of the target directory. Must match one of the directory definitions in synchronizationSchema.
func (m *SynchronizationRule) SetTargetDirectoryName(value *string)() {
    if m != nil {
        m.targetDirectoryName = value
    }
}
