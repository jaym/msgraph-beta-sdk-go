package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationStatusable 
type SynchronizationStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*SynchronizationStatusCode)
    GetCountSuccessiveCompleteFailures()(*int64)
    GetEscrowsPruned()(*bool)
    GetLastExecution()(RoleAssignmentsable)
    GetLastSuccessfulExecution()(RoleAssignmentsable)
    GetLastSuccessfulExecutionWithExports()(RoleAssignmentsable)
    GetProgress()([]RoleAssignmentsable)
    GetQuarantine()(RoleAssignmentsable)
    GetSteadyStateFirstAchievedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSteadyStateLastAchievedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSynchronizedEntryCountByType()([]RoleAssignmentsable)
    GetTroubleshootingUrl()(*string)
    SetCode(value *SynchronizationStatusCode)()
    SetCountSuccessiveCompleteFailures(value *int64)()
    SetEscrowsPruned(value *bool)()
    SetLastExecution(value RoleAssignmentsable)()
    SetLastSuccessfulExecution(value RoleAssignmentsable)()
    SetLastSuccessfulExecutionWithExports(value RoleAssignmentsable)()
    SetProgress(value []RoleAssignmentsable)()
    SetQuarantine(value RoleAssignmentsable)()
    SetSteadyStateFirstAchievedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSteadyStateLastAchievedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSynchronizedEntryCountByType(value []RoleAssignmentsable)()
    SetTroubleshootingUrl(value *string)()
}
