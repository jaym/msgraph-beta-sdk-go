package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    ib1755bf45aac675d3331dd21be8b6908c7df8eb8b5db77bcf4f8490b471220d4 "synchronizationstatus"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationStatus 
type SynchronizationStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The code property
    code *SynchronizationStatusCode
    // Number of consecutive times this job failed.
    countSuccessiveCompleteFailures *int64
    // true if the job's escrows (object-level errors) were pruned during initial synchronization. Escrows can be pruned if during the initial synchronization, you reach the threshold of errors that would normally put the job in quarantine. Instead of going into quarantine, the synchronization process clears the job's errors and continues until the initial synchronization is completed. When the initial synchronization is completed, the job will pause and wait for the customer to clean up the errors.
    escrowsPruned *bool
    // Details of the last execution of the job.
    lastExecution RoleAssignmentsable
    // Details of the last execution of this job, which didn't have any errors.
    lastSuccessfulExecution RoleAssignmentsable
    // Details of the last execution of the job, which exported objects into the target directory.
    lastSuccessfulExecutionWithExports RoleAssignmentsable
    // Details of the progress of a job toward completion.
    progress []RoleAssignmentsable
    // If job is in quarantine, quarantine details.
    quarantine RoleAssignmentsable
    // The time when steady state (no more changes to the process) was first achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    steadyStateFirstAchievedTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The time when steady state (no more changes to the process) was last achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    steadyStateLastAchievedTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Count of synchronized objects, listed by object type.
    synchronizedEntryCountByType []RoleAssignmentsable
    // In the event of an error, the URL with the troubleshooting steps for the issue.
    troubleshootingUrl *string
}
// RoleAssignments union type wrapper for classes synchronizationTaskExecution, roleAssignmentsMember1
type RoleAssignments struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type roleAssignmentsMember1
    roleAssignmentsMember1 RoleAssignmentsMember1able
    // Union type representation for type stringKeyLongValuePair
    stringKeyLongValuePair StringKeyLongValuePairable
    // Union type representation for type synchronizationProgress
    synchronizationProgress SynchronizationProgressable
    // Union type representation for type synchronizationQuarantine
    synchronizationQuarantine SynchronizationQuarantineable
    // Union type representation for type synchronizationTaskExecution
    synchronizationTaskExecution SynchronizationTaskExecutionable
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
    res["stringKeyLongValuePair"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStringKeyLongValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStringKeyLongValuePair(val.(StringKeyLongValuePairable))
        }
        return nil
    }
    res["synchronizationProgress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationProgressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSynchronizationProgress(val.(SynchronizationProgressable))
        }
        return nil
    }
    res["synchronizationQuarantine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationQuarantineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSynchronizationQuarantine(val.(SynchronizationQuarantineable))
        }
        return nil
    }
    res["synchronizationTaskExecution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationTaskExecutionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSynchronizationTaskExecution(val.(SynchronizationTaskExecutionable))
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
// GetStringKeyLongValuePair gets the stringKeyLongValuePair property value. Union type representation for type stringKeyLongValuePair
func (m *RoleAssignments) GetStringKeyLongValuePair()(StringKeyLongValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.stringKeyLongValuePair
    }
}
// GetSynchronizationProgress gets the synchronizationProgress property value. Union type representation for type synchronizationProgress
func (m *RoleAssignments) GetSynchronizationProgress()(SynchronizationProgressable) {
    if m == nil {
        return nil
    } else {
        return m.synchronizationProgress
    }
}
// GetSynchronizationQuarantine gets the synchronizationQuarantine property value. Union type representation for type synchronizationQuarantine
func (m *RoleAssignments) GetSynchronizationQuarantine()(SynchronizationQuarantineable) {
    if m == nil {
        return nil
    } else {
        return m.synchronizationQuarantine
    }
}
// GetSynchronizationTaskExecution gets the synchronizationTaskExecution property value. Union type representation for type synchronizationTaskExecution
func (m *RoleAssignments) GetSynchronizationTaskExecution()(SynchronizationTaskExecutionable) {
    if m == nil {
        return nil
    } else {
        return m.synchronizationTaskExecution
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
        err := writer.WriteObjectValue("stringKeyLongValuePair", m.GetStringKeyLongValuePair())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("synchronizationProgress", m.GetSynchronizationProgress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("synchronizationQuarantine", m.GetSynchronizationQuarantine())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("synchronizationTaskExecution", m.GetSynchronizationTaskExecution())
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
// SetStringKeyLongValuePair sets the stringKeyLongValuePair property value. Union type representation for type stringKeyLongValuePair
func (m *RoleAssignments) SetStringKeyLongValuePair(value StringKeyLongValuePairable)() {
    if m != nil {
        m.stringKeyLongValuePair = value
    }
}
// SetSynchronizationProgress sets the synchronizationProgress property value. Union type representation for type synchronizationProgress
func (m *RoleAssignments) SetSynchronizationProgress(value SynchronizationProgressable)() {
    if m != nil {
        m.synchronizationProgress = value
    }
}
// SetSynchronizationQuarantine sets the synchronizationQuarantine property value. Union type representation for type synchronizationQuarantine
func (m *RoleAssignments) SetSynchronizationQuarantine(value SynchronizationQuarantineable)() {
    if m != nil {
        m.synchronizationQuarantine = value
    }
}
// SetSynchronizationTaskExecution sets the synchronizationTaskExecution property value. Union type representation for type synchronizationTaskExecution
func (m *RoleAssignments) SetSynchronizationTaskExecution(value SynchronizationTaskExecutionable)() {
    if m != nil {
        m.synchronizationTaskExecution = value
    }
}
// RoleAssignmentsable 
type RoleAssignmentsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRoleAssignmentsMember1()(RoleAssignmentsMember1able)
    GetStringKeyLongValuePair()(StringKeyLongValuePairable)
    GetSynchronizationProgress()(SynchronizationProgressable)
    GetSynchronizationQuarantine()(SynchronizationQuarantineable)
    GetSynchronizationTaskExecution()(SynchronizationTaskExecutionable)
    SetRoleAssignmentsMember1(value RoleAssignmentsMember1able)()
    SetStringKeyLongValuePair(value StringKeyLongValuePairable)()
    SetSynchronizationProgress(value SynchronizationProgressable)()
    SetSynchronizationQuarantine(value SynchronizationQuarantineable)()
    SetSynchronizationTaskExecution(value SynchronizationTaskExecutionable)()
}
// NewSynchronizationStatus instantiates a new synchronizationStatus and sets the default values.
func NewSynchronizationStatus()(*SynchronizationStatus) {
    m := &SynchronizationStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSynchronizationStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCode gets the code property value. The code property
func (m *SynchronizationStatus) GetCode()(*SynchronizationStatusCode) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
// GetCountSuccessiveCompleteFailures gets the countSuccessiveCompleteFailures property value. Number of consecutive times this job failed.
func (m *SynchronizationStatus) GetCountSuccessiveCompleteFailures()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.countSuccessiveCompleteFailures
    }
}
// GetEscrowsPruned gets the escrowsPruned property value. true if the job's escrows (object-level errors) were pruned during initial synchronization. Escrows can be pruned if during the initial synchronization, you reach the threshold of errors that would normally put the job in quarantine. Instead of going into quarantine, the synchronization process clears the job's errors and continues until the initial synchronization is completed. When the initial synchronization is completed, the job will pause and wait for the customer to clean up the errors.
func (m *SynchronizationStatus) GetEscrowsPruned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.escrowsPruned
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSynchronizationStatusCode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val.(*SynchronizationStatusCode))
        }
        return nil
    }
    res["countSuccessiveCompleteFailures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountSuccessiveCompleteFailures(val)
        }
        return nil
    }
    res["escrowsPruned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEscrowsPruned(val)
        }
        return nil
    }
    res["lastExecution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoleAssignmentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastExecution(val.(RoleAssignmentsable))
        }
        return nil
    }
    res["lastSuccessfulExecution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoleAssignmentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSuccessfulExecution(val.(RoleAssignmentsable))
        }
        return nil
    }
    res["lastSuccessfulExecutionWithExports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoleAssignmentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSuccessfulExecutionWithExports(val.(RoleAssignmentsable))
        }
        return nil
    }
    res["progress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRoleAssignmentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleAssignmentsable, len(val))
            for i, v := range val {
                res[i] = v.(RoleAssignmentsable)
            }
            m.SetProgress(res)
        }
        return nil
    }
    res["quarantine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoleAssignmentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuarantine(val.(RoleAssignmentsable))
        }
        return nil
    }
    res["steadyStateFirstAchievedTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSteadyStateFirstAchievedTime(val)
        }
        return nil
    }
    res["steadyStateLastAchievedTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSteadyStateLastAchievedTime(val)
        }
        return nil
    }
    res["synchronizedEntryCountByType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRoleAssignmentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleAssignmentsable, len(val))
            for i, v := range val {
                res[i] = v.(RoleAssignmentsable)
            }
            m.SetSynchronizedEntryCountByType(res)
        }
        return nil
    }
    res["troubleshootingUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTroubleshootingUrl(val)
        }
        return nil
    }
    return res
}
// GetLastExecution gets the lastExecution property value. Details of the last execution of the job.
func (m *SynchronizationStatus) GetLastExecution()(RoleAssignmentsable) {
    if m == nil {
        return nil
    } else {
        return m.lastExecution
    }
}
// GetLastSuccessfulExecution gets the lastSuccessfulExecution property value. Details of the last execution of this job, which didn't have any errors.
func (m *SynchronizationStatus) GetLastSuccessfulExecution()(RoleAssignmentsable) {
    if m == nil {
        return nil
    } else {
        return m.lastSuccessfulExecution
    }
}
// GetLastSuccessfulExecutionWithExports gets the lastSuccessfulExecutionWithExports property value. Details of the last execution of the job, which exported objects into the target directory.
func (m *SynchronizationStatus) GetLastSuccessfulExecutionWithExports()(RoleAssignmentsable) {
    if m == nil {
        return nil
    } else {
        return m.lastSuccessfulExecutionWithExports
    }
}
// GetProgress gets the progress property value. Details of the progress of a job toward completion.
func (m *SynchronizationStatus) GetProgress()([]RoleAssignmentsable) {
    if m == nil {
        return nil
    } else {
        return m.progress
    }
}
// GetQuarantine gets the quarantine property value. If job is in quarantine, quarantine details.
func (m *SynchronizationStatus) GetQuarantine()(RoleAssignmentsable) {
    if m == nil {
        return nil
    } else {
        return m.quarantine
    }
}
// GetSteadyStateFirstAchievedTime gets the steadyStateFirstAchievedTime property value. The time when steady state (no more changes to the process) was first achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationStatus) GetSteadyStateFirstAchievedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.steadyStateFirstAchievedTime
    }
}
// GetSteadyStateLastAchievedTime gets the steadyStateLastAchievedTime property value. The time when steady state (no more changes to the process) was last achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationStatus) GetSteadyStateLastAchievedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.steadyStateLastAchievedTime
    }
}
// GetSynchronizedEntryCountByType gets the synchronizedEntryCountByType property value. Count of synchronized objects, listed by object type.
func (m *SynchronizationStatus) GetSynchronizedEntryCountByType()([]RoleAssignmentsable) {
    if m == nil {
        return nil
    } else {
        return m.synchronizedEntryCountByType
    }
}
// GetTroubleshootingUrl gets the troubleshootingUrl property value. In the event of an error, the URL with the troubleshooting steps for the issue.
func (m *SynchronizationStatus) GetTroubleshootingUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.troubleshootingUrl
    }
}
// Serialize serializes information the current object
func (m *SynchronizationStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCode() != nil {
        cast := (*m.GetCode()).String()
        err := writer.WriteStringValue("code", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("countSuccessiveCompleteFailures", m.GetCountSuccessiveCompleteFailures())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("escrowsPruned", m.GetEscrowsPruned())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lastExecution", m.GetLastExecution())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lastSuccessfulExecution", m.GetLastSuccessfulExecution())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lastSuccessfulExecutionWithExports", m.GetLastSuccessfulExecutionWithExports())
        if err != nil {
            return err
        }
    }
    if m.GetProgress() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProgress()))
        for i, v := range m.GetProgress() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("progress", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("quarantine", m.GetQuarantine())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("steadyStateFirstAchievedTime", m.GetSteadyStateFirstAchievedTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("steadyStateLastAchievedTime", m.GetSteadyStateLastAchievedTime())
        if err != nil {
            return err
        }
    }
    if m.GetSynchronizedEntryCountByType() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSynchronizedEntryCountByType()))
        for i, v := range m.GetSynchronizedEntryCountByType() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("synchronizedEntryCountByType", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("troubleshootingUrl", m.GetTroubleshootingUrl())
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
func (m *SynchronizationStatus) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCode sets the code property value. The code property
func (m *SynchronizationStatus) SetCode(value *SynchronizationStatusCode)() {
    if m != nil {
        m.code = value
    }
}
// SetCountSuccessiveCompleteFailures sets the countSuccessiveCompleteFailures property value. Number of consecutive times this job failed.
func (m *SynchronizationStatus) SetCountSuccessiveCompleteFailures(value *int64)() {
    if m != nil {
        m.countSuccessiveCompleteFailures = value
    }
}
// SetEscrowsPruned sets the escrowsPruned property value. true if the job's escrows (object-level errors) were pruned during initial synchronization. Escrows can be pruned if during the initial synchronization, you reach the threshold of errors that would normally put the job in quarantine. Instead of going into quarantine, the synchronization process clears the job's errors and continues until the initial synchronization is completed. When the initial synchronization is completed, the job will pause and wait for the customer to clean up the errors.
func (m *SynchronizationStatus) SetEscrowsPruned(value *bool)() {
    if m != nil {
        m.escrowsPruned = value
    }
}
// SetLastExecution sets the lastExecution property value. Details of the last execution of the job.
func (m *SynchronizationStatus) SetLastExecution(value RoleAssignmentsable)() {
    if m != nil {
        m.lastExecution = value
    }
}
// SetLastSuccessfulExecution sets the lastSuccessfulExecution property value. Details of the last execution of this job, which didn't have any errors.
func (m *SynchronizationStatus) SetLastSuccessfulExecution(value RoleAssignmentsable)() {
    if m != nil {
        m.lastSuccessfulExecution = value
    }
}
// SetLastSuccessfulExecutionWithExports sets the lastSuccessfulExecutionWithExports property value. Details of the last execution of the job, which exported objects into the target directory.
func (m *SynchronizationStatus) SetLastSuccessfulExecutionWithExports(value RoleAssignmentsable)() {
    if m != nil {
        m.lastSuccessfulExecutionWithExports = value
    }
}
// SetProgress sets the progress property value. Details of the progress of a job toward completion.
func (m *SynchronizationStatus) SetProgress(value []RoleAssignmentsable)() {
    if m != nil {
        m.progress = value
    }
}
// SetQuarantine sets the quarantine property value. If job is in quarantine, quarantine details.
func (m *SynchronizationStatus) SetQuarantine(value RoleAssignmentsable)() {
    if m != nil {
        m.quarantine = value
    }
}
// SetSteadyStateFirstAchievedTime sets the steadyStateFirstAchievedTime property value. The time when steady state (no more changes to the process) was first achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationStatus) SetSteadyStateFirstAchievedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.steadyStateFirstAchievedTime = value
    }
}
// SetSteadyStateLastAchievedTime sets the steadyStateLastAchievedTime property value. The time when steady state (no more changes to the process) was last achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationStatus) SetSteadyStateLastAchievedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.steadyStateLastAchievedTime = value
    }
}
// SetSynchronizedEntryCountByType sets the synchronizedEntryCountByType property value. Count of synchronized objects, listed by object type.
func (m *SynchronizationStatus) SetSynchronizedEntryCountByType(value []RoleAssignmentsable)() {
    if m != nil {
        m.synchronizedEntryCountByType = value
    }
}
// SetTroubleshootingUrl sets the troubleshootingUrl property value. In the event of an error, the URL with the troubleshooting steps for the issue.
func (m *SynchronizationStatus) SetTroubleshootingUrl(value *string)() {
    if m != nil {
        m.troubleshootingUrl = value
    }
}
