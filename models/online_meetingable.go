package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnlineMeetingable 
type OnlineMeetingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowAttendeeToEnableCamera()(*bool)
    GetAllowAttendeeToEnableMic()(*bool)
    GetAllowedPresenters()(*OnlineMeetingPresenters)
    GetAllowMeetingChat()(*MeetingChatMode)
    GetAllowTeamworkReactions()(*bool)
    GetAlternativeRecording()([]byte)
    GetAnonymizeIdentityForRoles()([]string)
    GetAttendanceReports()([]MeetingAttendanceReportable)
    GetAttendeeReport()([]byte)
    GetAudioConferencing()(AudioConferencingable)
    GetBroadcastSettings()(BroadcastMeetingSettingsable)
    GetCapabilities()([]string)
    GetChatInfo()(ChatInfoable)
    GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExternalId()(*string)
    GetIsBroadcast()(*bool)
    GetIsEntryExitAnnounced()(*bool)
    GetJoinInformation()(ItemBodyable)
    GetJoinMeetingIdSettings()(JoinMeetingIdSettingsable)
    GetJoinUrl()(*string)
    GetJoinWebUrl()(*string)
    GetLobbyBypassSettings()(LobbyBypassSettingsable)
    GetMeetingAttendanceReport()(MeetingAttendanceReportable)
    GetParticipants()(MeetingParticipantsable)
    GetRecordAutomatically()(*bool)
    GetRecording()([]byte)
    GetRegistration()(MeetingRegistrationable)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSubject()(*string)
    GetVideoTeleconferenceId()(*string)
    SetAllowAttendeeToEnableCamera(value *bool)()
    SetAllowAttendeeToEnableMic(value *bool)()
    SetAllowedPresenters(value *OnlineMeetingPresenters)()
    SetAllowMeetingChat(value *MeetingChatMode)()
    SetAllowTeamworkReactions(value *bool)()
    SetAlternativeRecording(value []byte)()
    SetAnonymizeIdentityForRoles(value []string)()
    SetAttendanceReports(value []MeetingAttendanceReportable)()
    SetAttendeeReport(value []byte)()
    SetAudioConferencing(value AudioConferencingable)()
    SetBroadcastSettings(value BroadcastMeetingSettingsable)()
    SetCapabilities(value []string)()
    SetChatInfo(value ChatInfoable)()
    SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExternalId(value *string)()
    SetIsBroadcast(value *bool)()
    SetIsEntryExitAnnounced(value *bool)()
    SetJoinInformation(value ItemBodyable)()
    SetJoinMeetingIdSettings(value JoinMeetingIdSettingsable)()
    SetJoinUrl(value *string)()
    SetJoinWebUrl(value *string)()
    SetLobbyBypassSettings(value LobbyBypassSettingsable)()
    SetMeetingAttendanceReport(value MeetingAttendanceReportable)()
    SetParticipants(value MeetingParticipantsable)()
    SetRecordAutomatically(value *bool)()
    SetRecording(value []byte)()
    SetRegistration(value MeetingRegistrationable)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSubject(value *string)()
    SetVideoTeleconferenceId(value *string)()
}