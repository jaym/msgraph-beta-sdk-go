package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    ieec3b3cefaa49311ed258da81563b5b029fd831fffea9e08a2c1b69fe5888d73 "devicecompliancepolicy"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceCompliancePolicy this is the base class for Compliance policy. Compliance policies are platform specific and individual per-platform compliance policies inherit from here. 
type DeviceCompliancePolicy struct {
    Entity
    // The collection of assignments for this compliance policy.
    assignments []DeviceCompliancePolicyAssignmentable
    // DateTime the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Admin provided description of the Device Configuration.
    description *string
    // Compliance Setting State Device Summary
    deviceSettingStateSummaries []SettingStateDeviceSummaryable
    // List of DeviceComplianceDeviceStatus.
    deviceStatuses []DeviceComplianceDeviceStatusable
    // Device compliance devices status overview
    deviceStatusOverview DeviceCompliancePoliciesable
    // Admin provided name of the device configuration.
    displayName *string
    // DateTime the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string
    // The list of scheduled action for this rule
    scheduledActionsForRule []DeviceComplianceScheduledActionForRuleable
    // List of DeviceComplianceUserStatus.
    userStatuses []DeviceComplianceUserStatusable
    // Device compliance users status overview
    userStatusOverview DeviceCompliancePoliciesable
    // Version of the device configuration.
    version *int32
}
// DeviceCompliancePolicies union type wrapper for classes deviceComplianceUserOverview, deviceCompliancePoliciesMember1
type DeviceCompliancePolicies struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type deviceComplianceDeviceOverview
    deviceComplianceDeviceOverview DeviceComplianceDeviceOverviewable
    // Union type representation for type deviceCompliancePoliciesMember1
    deviceCompliancePoliciesMember1 DeviceCompliancePoliciesMember1able
    // Union type representation for type deviceComplianceUserOverview
    deviceComplianceUserOverview DeviceComplianceUserOverviewable
}
// NewDeviceCompliancePolicies instantiates a new deviceCompliancePolicies and sets the default values.
func NewDeviceCompliancePolicies()(*DeviceCompliancePolicies) {
    m := &DeviceCompliancePolicies{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceCompliancePoliciesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCompliancePoliciesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCompliancePolicies(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceCompliancePolicies) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceComplianceDeviceOverview gets the deviceComplianceDeviceOverview property value. Union type representation for type deviceComplianceDeviceOverview
func (m *DeviceCompliancePolicies) GetDeviceComplianceDeviceOverview()(DeviceComplianceDeviceOverviewable) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceDeviceOverview
    }
}
// GetDeviceCompliancePoliciesMember1 gets the deviceCompliancePoliciesMember1 property value. Union type representation for type deviceCompliancePoliciesMember1
func (m *DeviceCompliancePolicies) GetDeviceCompliancePoliciesMember1()(DeviceCompliancePoliciesMember1able) {
    if m == nil {
        return nil
    } else {
        return m.deviceCompliancePoliciesMember1
    }
}
// GetDeviceComplianceUserOverview gets the deviceComplianceUserOverview property value. Union type representation for type deviceComplianceUserOverview
func (m *DeviceCompliancePolicies) GetDeviceComplianceUserOverview()(DeviceComplianceUserOverviewable) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceUserOverview
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCompliancePolicies) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceComplianceDeviceOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceComplianceDeviceOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceComplianceDeviceOverview(val.(DeviceComplianceDeviceOverviewable))
        }
        return nil
    }
    res["deviceCompliancePoliciesMember1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceCompliancePoliciesMember1FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCompliancePoliciesMember1(val.(DeviceCompliancePoliciesMember1able))
        }
        return nil
    }
    res["deviceComplianceUserOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceComplianceUserOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceComplianceUserOverview(val.(DeviceComplianceUserOverviewable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceCompliancePolicies) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("deviceComplianceDeviceOverview", m.GetDeviceComplianceDeviceOverview())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("deviceCompliancePoliciesMember1", m.GetDeviceCompliancePoliciesMember1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("deviceComplianceUserOverview", m.GetDeviceComplianceUserOverview())
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
func (m *DeviceCompliancePolicies) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceComplianceDeviceOverview sets the deviceComplianceDeviceOverview property value. Union type representation for type deviceComplianceDeviceOverview
func (m *DeviceCompliancePolicies) SetDeviceComplianceDeviceOverview(value DeviceComplianceDeviceOverviewable)() {
    if m != nil {
        m.deviceComplianceDeviceOverview = value
    }
}
// SetDeviceCompliancePoliciesMember1 sets the deviceCompliancePoliciesMember1 property value. Union type representation for type deviceCompliancePoliciesMember1
func (m *DeviceCompliancePolicies) SetDeviceCompliancePoliciesMember1(value DeviceCompliancePoliciesMember1able)() {
    if m != nil {
        m.deviceCompliancePoliciesMember1 = value
    }
}
// SetDeviceComplianceUserOverview sets the deviceComplianceUserOverview property value. Union type representation for type deviceComplianceUserOverview
func (m *DeviceCompliancePolicies) SetDeviceComplianceUserOverview(value DeviceComplianceUserOverviewable)() {
    if m != nil {
        m.deviceComplianceUserOverview = value
    }
}
// DeviceCompliancePoliciesable 
type DeviceCompliancePoliciesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceComplianceDeviceOverview()(DeviceComplianceDeviceOverviewable)
    GetDeviceCompliancePoliciesMember1()(DeviceCompliancePoliciesMember1able)
    GetDeviceComplianceUserOverview()(DeviceComplianceUserOverviewable)
    SetDeviceComplianceDeviceOverview(value DeviceComplianceDeviceOverviewable)()
    SetDeviceCompliancePoliciesMember1(value DeviceCompliancePoliciesMember1able)()
    SetDeviceComplianceUserOverview(value DeviceComplianceUserOverviewable)()
}
// NewDeviceCompliancePolicy instantiates a new deviceCompliancePolicy and sets the default values.
func NewDeviceCompliancePolicy()(*DeviceCompliancePolicy) {
    m := &DeviceCompliancePolicy{
        Entity: *NewEntity(),
    }
    typeValue := "#microsoft.graph.deviceCompliancePolicy";
    m.SetType(&typeValue);
    return m
}
// CreateDeviceCompliancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCompliancePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCompliancePolicy(), nil
}
// GetAssignments gets the assignments property value. The collection of assignments for this compliance policy.
func (m *DeviceCompliancePolicy) GetAssignments()([]DeviceCompliancePolicyAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetCreatedDateTime gets the createdDateTime property value. DateTime the object was created.
func (m *DeviceCompliancePolicy) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. Admin provided description of the Device Configuration.
func (m *DeviceCompliancePolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDeviceSettingStateSummaries gets the deviceSettingStateSummaries property value. Compliance Setting State Device Summary
func (m *DeviceCompliancePolicy) GetDeviceSettingStateSummaries()([]SettingStateDeviceSummaryable) {
    if m == nil {
        return nil
    } else {
        return m.deviceSettingStateSummaries
    }
}
// GetDeviceStatuses gets the deviceStatuses property value. List of DeviceComplianceDeviceStatus.
func (m *DeviceCompliancePolicy) GetDeviceStatuses()([]DeviceComplianceDeviceStatusable) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatuses
    }
}
// GetDeviceStatusOverview gets the deviceStatusOverview property value. Device compliance devices status overview
func (m *DeviceCompliancePolicy) GetDeviceStatusOverview()(DeviceCompliancePoliciesable) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatusOverview
    }
}
// GetDisplayName gets the displayName property value. Admin provided name of the device configuration.
func (m *DeviceCompliancePolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCompliancePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceCompliancePolicyAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCompliancePolicyAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceCompliancePolicyAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["deviceSettingStateSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSettingStateDeviceSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SettingStateDeviceSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(SettingStateDeviceSummaryable)
            }
            m.SetDeviceSettingStateSummaries(res)
        }
        return nil
    }
    res["deviceStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceComplianceDeviceStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceComplianceDeviceStatusable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceComplianceDeviceStatusable)
            }
            m.SetDeviceStatuses(res)
        }
        return nil
    }
    res["deviceStatusOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceCompliancePoliciesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceStatusOverview(val.(DeviceCompliancePoliciesable))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["scheduledActionsForRule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceComplianceScheduledActionForRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceComplianceScheduledActionForRuleable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceComplianceScheduledActionForRuleable)
            }
            m.SetScheduledActionsForRule(res)
        }
        return nil
    }
    res["userStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceComplianceUserStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceComplianceUserStatusable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceComplianceUserStatusable)
            }
            m.SetUserStatuses(res)
        }
        return nil
    }
    res["userStatusOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceCompliancePoliciesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserStatusOverview(val.(DeviceCompliancePoliciesable))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *DeviceCompliancePolicy) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceCompliancePolicy) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// GetScheduledActionsForRule gets the scheduledActionsForRule property value. The list of scheduled action for this rule
func (m *DeviceCompliancePolicy) GetScheduledActionsForRule()([]DeviceComplianceScheduledActionForRuleable) {
    if m == nil {
        return nil
    } else {
        return m.scheduledActionsForRule
    }
}
// GetUserStatuses gets the userStatuses property value. List of DeviceComplianceUserStatus.
func (m *DeviceCompliancePolicy) GetUserStatuses()([]DeviceComplianceUserStatusable) {
    if m == nil {
        return nil
    } else {
        return m.userStatuses
    }
}
// GetUserStatusOverview gets the userStatusOverview property value. Device compliance users status overview
func (m *DeviceCompliancePolicy) GetUserStatusOverview()(DeviceCompliancePoliciesable) {
    if m == nil {
        return nil
    } else {
        return m.userStatusOverview
    }
}
// GetVersion gets the version property value. Version of the device configuration.
func (m *DeviceCompliancePolicy) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// Serialize serializes information the current object
func (m *DeviceCompliancePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
    if m.GetDeviceSettingStateSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceSettingStateSummaries()))
        for i, v := range m.GetDeviceSettingStateSummaries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceSettingStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceStatuses()))
        for i, v := range m.GetDeviceStatuses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceStatusOverview", m.GetDeviceStatusOverview())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    if m.GetScheduledActionsForRule() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetScheduledActionsForRule()))
        for i, v := range m.GetScheduledActionsForRule() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("scheduledActionsForRule", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserStatuses()))
        for i, v := range m.GetUserStatuses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userStatusOverview", m.GetUserStatusOverview())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The collection of assignments for this compliance policy.
func (m *DeviceCompliancePolicy) SetAssignments(value []DeviceCompliancePolicyAssignmentable)() {
    if m != nil {
        m.assignments = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. DateTime the object was created.
func (m *DeviceCompliancePolicy) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. Admin provided description of the Device Configuration.
func (m *DeviceCompliancePolicy) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDeviceSettingStateSummaries sets the deviceSettingStateSummaries property value. Compliance Setting State Device Summary
func (m *DeviceCompliancePolicy) SetDeviceSettingStateSummaries(value []SettingStateDeviceSummaryable)() {
    if m != nil {
        m.deviceSettingStateSummaries = value
    }
}
// SetDeviceStatuses sets the deviceStatuses property value. List of DeviceComplianceDeviceStatus.
func (m *DeviceCompliancePolicy) SetDeviceStatuses(value []DeviceComplianceDeviceStatusable)() {
    if m != nil {
        m.deviceStatuses = value
    }
}
// SetDeviceStatusOverview sets the deviceStatusOverview property value. Device compliance devices status overview
func (m *DeviceCompliancePolicy) SetDeviceStatusOverview(value DeviceCompliancePoliciesable)() {
    if m != nil {
        m.deviceStatusOverview = value
    }
}
// SetDisplayName sets the displayName property value. Admin provided name of the device configuration.
func (m *DeviceCompliancePolicy) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *DeviceCompliancePolicy) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceCompliancePolicy) SetRoleScopeTagIds(value []string)() {
    if m != nil {
        m.roleScopeTagIds = value
    }
}
// SetScheduledActionsForRule sets the scheduledActionsForRule property value. The list of scheduled action for this rule
func (m *DeviceCompliancePolicy) SetScheduledActionsForRule(value []DeviceComplianceScheduledActionForRuleable)() {
    if m != nil {
        m.scheduledActionsForRule = value
    }
}
// SetUserStatuses sets the userStatuses property value. List of DeviceComplianceUserStatus.
func (m *DeviceCompliancePolicy) SetUserStatuses(value []DeviceComplianceUserStatusable)() {
    if m != nil {
        m.userStatuses = value
    }
}
// SetUserStatusOverview sets the userStatusOverview property value. Device compliance users status overview
func (m *DeviceCompliancePolicy) SetUserStatusOverview(value DeviceCompliancePoliciesable)() {
    if m != nil {
        m.userStatusOverview = value
    }
}
// SetVersion sets the version property value. Version of the device configuration.
func (m *DeviceCompliancePolicy) SetVersion(value *int32)() {
    if m != nil {
        m.version = value
    }
}
