package models

import (
    ia60c72aa57578abddbf18c001bcb41236af765306672aca6d0b38c6646aab238 "synchronizationjob"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationJob 
type SynchronizationJob struct {
    Entity
    // Schedule used to run the job. Read-only.
    schedule RoleAssignmentsable
    // The synchronization schema configured for the job.
    schema RoleAssignmentsable
    // Status of the job, which includes when the job was last run, current job state, and errors.
    status RoleAssignmentsable
    // Settings associated with the job. Some settings are inherited from the template.
    synchronizationJobSettings []RoleAssignmentsable
    // Identifier of the synchronization template this job is based on.
    templateId *string
}
// RoleAssignments union type wrapper for classes synchronizationSchedule, roleAssignmentsMember1
type RoleAssignments struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type keyValuePair
    keyValuePair KeyValuePairable
    // Union type representation for type roleAssignmentsMember1
    roleAssignmentsMember1 RoleAssignmentsMember1able
    // Union type representation for type synchronizationSchedule
    synchronizationSchedule SynchronizationScheduleable
    // Union type representation for type synchronizationSchema
    synchronizationSchema SynchronizationSchemaable
    // Union type representation for type synchronizationStatus
    synchronizationStatus SynchronizationStatusable
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
    res["keyValuePair"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyValuePair(val.(KeyValuePairable))
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
    res["synchronizationSchedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSynchronizationSchedule(val.(SynchronizationScheduleable))
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
    res["synchronizationStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSynchronizationStatus(val.(SynchronizationStatusable))
        }
        return nil
    }
    return res
}
// GetKeyValuePair gets the keyValuePair property value. Union type representation for type keyValuePair
func (m *RoleAssignments) GetKeyValuePair()(KeyValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.keyValuePair
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
// GetSynchronizationSchedule gets the synchronizationSchedule property value. Union type representation for type synchronizationSchedule
func (m *RoleAssignments) GetSynchronizationSchedule()(SynchronizationScheduleable) {
    if m == nil {
        return nil
    } else {
        return m.synchronizationSchedule
    }
}
// GetSynchronizationSchema gets the synchronizationSchema property value. Union type representation for type synchronizationSchema
func (m *RoleAssignments) GetSynchronizationSchema()(SynchronizationSchemaable) {
    if m == nil {
        return nil
    } else {
        return m.synchronizationSchema
    }
}
// GetSynchronizationStatus gets the synchronizationStatus property value. Union type representation for type synchronizationStatus
func (m *RoleAssignments) GetSynchronizationStatus()(SynchronizationStatusable) {
    if m == nil {
        return nil
    } else {
        return m.synchronizationStatus
    }
}
// Serialize serializes information the current object
func (m *RoleAssignments) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("keyValuePair", m.GetKeyValuePair())
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
        err := writer.WriteObjectValue("synchronizationSchedule", m.GetSynchronizationSchedule())
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
        err := writer.WriteObjectValue("synchronizationStatus", m.GetSynchronizationStatus())
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
// SetKeyValuePair sets the keyValuePair property value. Union type representation for type keyValuePair
func (m *RoleAssignments) SetKeyValuePair(value KeyValuePairable)() {
    if m != nil {
        m.keyValuePair = value
    }
}
// SetRoleAssignmentsMember1 sets the roleAssignmentsMember1 property value. Union type representation for type roleAssignmentsMember1
func (m *RoleAssignments) SetRoleAssignmentsMember1(value RoleAssignmentsMember1able)() {
    if m != nil {
        m.roleAssignmentsMember1 = value
    }
}
// SetSynchronizationSchedule sets the synchronizationSchedule property value. Union type representation for type synchronizationSchedule
func (m *RoleAssignments) SetSynchronizationSchedule(value SynchronizationScheduleable)() {
    if m != nil {
        m.synchronizationSchedule = value
    }
}
// SetSynchronizationSchema sets the synchronizationSchema property value. Union type representation for type synchronizationSchema
func (m *RoleAssignments) SetSynchronizationSchema(value SynchronizationSchemaable)() {
    if m != nil {
        m.synchronizationSchema = value
    }
}
// SetSynchronizationStatus sets the synchronizationStatus property value. Union type representation for type synchronizationStatus
func (m *RoleAssignments) SetSynchronizationStatus(value SynchronizationStatusable)() {
    if m != nil {
        m.synchronizationStatus = value
    }
}
// RoleAssignmentsable 
type RoleAssignmentsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKeyValuePair()(KeyValuePairable)
    GetRoleAssignmentsMember1()(RoleAssignmentsMember1able)
    GetSynchronizationSchedule()(SynchronizationScheduleable)
    GetSynchronizationSchema()(SynchronizationSchemaable)
    GetSynchronizationStatus()(SynchronizationStatusable)
    SetKeyValuePair(value KeyValuePairable)()
    SetRoleAssignmentsMember1(value RoleAssignmentsMember1able)()
    SetSynchronizationSchedule(value SynchronizationScheduleable)()
    SetSynchronizationSchema(value SynchronizationSchemaable)()
    SetSynchronizationStatus(value SynchronizationStatusable)()
}
// NewSynchronizationJob instantiates a new synchronizationJob and sets the default values.
func NewSynchronizationJob()(*SynchronizationJob) {
    m := &SynchronizationJob{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSynchronizationJobFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationJobFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationJob(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationJob) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["schedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoleAssignmentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(RoleAssignmentsable))
        }
        return nil
    }
    res["schema"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoleAssignmentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchema(val.(RoleAssignmentsable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoleAssignmentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(RoleAssignmentsable))
        }
        return nil
    }
    res["synchronizationJobSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRoleAssignmentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleAssignmentsable, len(val))
            for i, v := range val {
                res[i] = v.(RoleAssignmentsable)
            }
            m.SetSynchronizationJobSettings(res)
        }
        return nil
    }
    res["templateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateId(val)
        }
        return nil
    }
    return res
}
// GetSchedule gets the schedule property value. Schedule used to run the job. Read-only.
func (m *SynchronizationJob) GetSchedule()(RoleAssignmentsable) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
// GetSchema gets the schema property value. The synchronization schema configured for the job.
func (m *SynchronizationJob) GetSchema()(RoleAssignmentsable) {
    if m == nil {
        return nil
    } else {
        return m.schema
    }
}
// GetStatus gets the status property value. Status of the job, which includes when the job was last run, current job state, and errors.
func (m *SynchronizationJob) GetStatus()(RoleAssignmentsable) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetSynchronizationJobSettings gets the synchronizationJobSettings property value. Settings associated with the job. Some settings are inherited from the template.
func (m *SynchronizationJob) GetSynchronizationJobSettings()([]RoleAssignmentsable) {
    if m == nil {
        return nil
    } else {
        return m.synchronizationJobSettings
    }
}
// GetTemplateId gets the templateId property value. Identifier of the synchronization template this job is based on.
func (m *SynchronizationJob) GetTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateId
    }
}
// Serialize serializes information the current object
func (m *SynchronizationJob) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("schedule", m.GetSchedule())
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
    {
        err = writer.WriteObjectValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    if m.GetSynchronizationJobSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSynchronizationJobSettings()))
        for i, v := range m.GetSynchronizationJobSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("synchronizationJobSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("templateId", m.GetTemplateId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSchedule sets the schedule property value. Schedule used to run the job. Read-only.
func (m *SynchronizationJob) SetSchedule(value RoleAssignmentsable)() {
    if m != nil {
        m.schedule = value
    }
}
// SetSchema sets the schema property value. The synchronization schema configured for the job.
func (m *SynchronizationJob) SetSchema(value RoleAssignmentsable)() {
    if m != nil {
        m.schema = value
    }
}
// SetStatus sets the status property value. Status of the job, which includes when the job was last run, current job state, and errors.
func (m *SynchronizationJob) SetStatus(value RoleAssignmentsable)() {
    if m != nil {
        m.status = value
    }
}
// SetSynchronizationJobSettings sets the synchronizationJobSettings property value. Settings associated with the job. Some settings are inherited from the template.
func (m *SynchronizationJob) SetSynchronizationJobSettings(value []RoleAssignmentsable)() {
    if m != nil {
        m.synchronizationJobSettings = value
    }
}
// SetTemplateId sets the templateId property value. Identifier of the synchronization template this job is based on.
func (m *SynchronizationJob) SetTemplateId(value *string)() {
    if m != nil {
        m.templateId = value
    }
}
