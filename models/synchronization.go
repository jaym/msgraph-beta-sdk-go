package models

import (
    ib67cd4ca4e32bf16c29c68d60b17792bc72277eb8eee122da56fa9274e1c9b28 "synchronization"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Synchronization 
type Synchronization struct {
    Entity
    // Performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
    jobs []SynchronizationJobable
    // Represents a collection of credentials to access provisioned cloud applications.
    secrets []RoleAssignmentsable
    // Pre-configured synchronization settings for a particular application.
    templates []SynchronizationTemplateable
}
// RoleAssignments union type wrapper for classes synchronizationSecretKeyStringValuePair, roleAssignmentsMember1
type RoleAssignments struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type roleAssignmentsMember1
    roleAssignmentsMember1 RoleAssignmentsMember1able
    // Union type representation for type synchronizationSecretKeyStringValuePair
    synchronizationSecretKeyStringValuePair SynchronizationSecretKeyStringValuePairable
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
    res["synchronizationSecretKeyStringValuePair"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationSecretKeyStringValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSynchronizationSecretKeyStringValuePair(val.(SynchronizationSecretKeyStringValuePairable))
        }
        return nil
    }
    return res
}
// GetRoleAssignmentsMember1 gets the roleAssignmentsMember1 property value. Union type representation for type roleAssignmentsMember1
func (m *RoleAssignments) GetRoleAssignmentsMember1()(RoleAssignmentsMember1able) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentsMember1
    }
}
// GetSynchronizationSecretKeyStringValuePair gets the synchronizationSecretKeyStringValuePair property value. Union type representation for type synchronizationSecretKeyStringValuePair
func (m *RoleAssignments) GetSynchronizationSecretKeyStringValuePair()(SynchronizationSecretKeyStringValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.synchronizationSecretKeyStringValuePair
    }
}
// Serialize serializes information the current object
func (m *RoleAssignments) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("roleAssignmentsMember1", m.GetRoleAssignmentsMember1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("synchronizationSecretKeyStringValuePair", m.GetSynchronizationSecretKeyStringValuePair())
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
// SetRoleAssignmentsMember1 sets the roleAssignmentsMember1 property value. Union type representation for type roleAssignmentsMember1
func (m *RoleAssignments) SetRoleAssignmentsMember1(value RoleAssignmentsMember1able)() {
    if m != nil {
        m.roleAssignmentsMember1 = value
    }
}
// SetSynchronizationSecretKeyStringValuePair sets the synchronizationSecretKeyStringValuePair property value. Union type representation for type synchronizationSecretKeyStringValuePair
func (m *RoleAssignments) SetSynchronizationSecretKeyStringValuePair(value SynchronizationSecretKeyStringValuePairable)() {
    if m != nil {
        m.synchronizationSecretKeyStringValuePair = value
    }
}
// RoleAssignmentsable 
type RoleAssignmentsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRoleAssignmentsMember1()(RoleAssignmentsMember1able)
    GetSynchronizationSecretKeyStringValuePair()(SynchronizationSecretKeyStringValuePairable)
    SetRoleAssignmentsMember1(value RoleAssignmentsMember1able)()
    SetSynchronizationSecretKeyStringValuePair(value SynchronizationSecretKeyStringValuePairable)()
}
// NewSynchronization instantiates a new Synchronization and sets the default values.
func NewSynchronization()(*Synchronization) {
    m := &Synchronization{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSynchronizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronization(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Synchronization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["jobs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSynchronizationJobFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SynchronizationJobable, len(val))
            for i, v := range val {
                res[i] = v.(SynchronizationJobable)
            }
            m.SetJobs(res)
        }
        return nil
    }
    res["secrets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRoleAssignmentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleAssignmentsable, len(val))
            for i, v := range val {
                res[i] = v.(RoleAssignmentsable)
            }
            m.SetSecrets(res)
        }
        return nil
    }
    res["templates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSynchronizationTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SynchronizationTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(SynchronizationTemplateable)
            }
            m.SetTemplates(res)
        }
        return nil
    }
    return res
}
// GetJobs gets the jobs property value. Performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
func (m *Synchronization) GetJobs()([]SynchronizationJobable) {
    if m == nil {
        return nil
    } else {
        return m.jobs
    }
}
// GetSecrets gets the secrets property value. Represents a collection of credentials to access provisioned cloud applications.
func (m *Synchronization) GetSecrets()([]RoleAssignmentsable) {
    if m == nil {
        return nil
    } else {
        return m.secrets
    }
}
// GetTemplates gets the templates property value. Pre-configured synchronization settings for a particular application.
func (m *Synchronization) GetTemplates()([]SynchronizationTemplateable) {
    if m == nil {
        return nil
    } else {
        return m.templates
    }
}
// Serialize serializes information the current object
func (m *Synchronization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetJobs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetJobs()))
        for i, v := range m.GetJobs() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("jobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecrets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSecrets()))
        for i, v := range m.GetSecrets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("secrets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTemplates()))
        for i, v := range m.GetTemplates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("templates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetJobs sets the jobs property value. Performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
func (m *Synchronization) SetJobs(value []SynchronizationJobable)() {
    if m != nil {
        m.jobs = value
    }
}
// SetSecrets sets the secrets property value. Represents a collection of credentials to access provisioned cloud applications.
func (m *Synchronization) SetSecrets(value []RoleAssignmentsable)() {
    if m != nil {
        m.secrets = value
    }
}
// SetTemplates sets the templates property value. Pre-configured synchronization settings for a particular application.
func (m *Synchronization) SetTemplates(value []SynchronizationTemplateable)() {
    if m != nil {
        m.templates = value
    }
}
