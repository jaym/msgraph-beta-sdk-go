package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    if28b7c7fbe90206b93837a76712128e6b7972d3ae4bd2142ac85f7b7b892aff7 "securescore"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecureScore 
type SecureScore struct {
    Entity
    // Active user count of the given tenant.
    activeUserCount *int32
    // Average score by different scopes (for example, average by industry, average by seating) and control category (Identity, Data, Device, Apps, Infrastructure) within the scope.
    averageComparativeScores []RoleDefinitionsable
    // GUID string for tenant ID.
    azureTenantId *string
    // Contains tenant scores for a set of controls.
    controlScores []RoleDefinitionsable
    // The date when the entity is created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Tenant current attained score on specified date.
    currentScore *float64
    // Microsoft-provided services for the tenant (for example, Exchange online, Skype, SharePoint).
    enabledServices []string
    // Licensed user count of the given tenant.
    licensedUserCount *int32
    // Tenant maximum possible score on specified date.
    maxScore *float64
    // Complex type containing details about the security product/service vendor, provider, and subprovider (for example, vendor=Microsoft; provider=SecureScore). Required.
    vendorInformation RoleDefinitionsable
}
// RoleDefinitions union type wrapper for classes averageComparativeScore, roleDefinitionsMember1
type RoleDefinitions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type averageComparativeScore
    averageComparativeScore AverageComparativeScoreable
    // Union type representation for type controlScore
    controlScore ControlScoreable
    // Union type representation for type roleDefinitionsMember1
    roleDefinitionsMember1 RoleDefinitionsMember1able
    // Union type representation for type securityVendorInformation
    securityVendorInformation SecurityVendorInformationable
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
// GetAverageComparativeScore gets the averageComparativeScore property value. Union type representation for type averageComparativeScore
func (m *RoleDefinitions) GetAverageComparativeScore()(AverageComparativeScoreable) {
    if m == nil {
        return nil
    } else {
        return m.averageComparativeScore
    }
}
// GetControlScore gets the controlScore property value. Union type representation for type controlScore
func (m *RoleDefinitions) GetControlScore()(ControlScoreable) {
    if m == nil {
        return nil
    } else {
        return m.controlScore
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleDefinitions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["averageComparativeScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAverageComparativeScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageComparativeScore(val.(AverageComparativeScoreable))
        }
        return nil
    }
    res["controlScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateControlScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetControlScore(val.(ControlScoreable))
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
    res["securityVendorInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecurityVendorInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityVendorInformation(val.(SecurityVendorInformationable))
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
// GetSecurityVendorInformation gets the securityVendorInformation property value. Union type representation for type securityVendorInformation
func (m *RoleDefinitions) GetSecurityVendorInformation()(SecurityVendorInformationable) {
    if m == nil {
        return nil
    } else {
        return m.securityVendorInformation
    }
}
// Serialize serializes information the current object
func (m *RoleDefinitions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("averageComparativeScore", m.GetAverageComparativeScore())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("controlScore", m.GetControlScore())
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
        err := writer.WriteObjectValue("securityVendorInformation", m.GetSecurityVendorInformation())
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
// SetAverageComparativeScore sets the averageComparativeScore property value. Union type representation for type averageComparativeScore
func (m *RoleDefinitions) SetAverageComparativeScore(value AverageComparativeScoreable)() {
    if m != nil {
        m.averageComparativeScore = value
    }
}
// SetControlScore sets the controlScore property value. Union type representation for type controlScore
func (m *RoleDefinitions) SetControlScore(value ControlScoreable)() {
    if m != nil {
        m.controlScore = value
    }
}
// SetRoleDefinitionsMember1 sets the roleDefinitionsMember1 property value. Union type representation for type roleDefinitionsMember1
func (m *RoleDefinitions) SetRoleDefinitionsMember1(value RoleDefinitionsMember1able)() {
    if m != nil {
        m.roleDefinitionsMember1 = value
    }
}
// SetSecurityVendorInformation sets the securityVendorInformation property value. Union type representation for type securityVendorInformation
func (m *RoleDefinitions) SetSecurityVendorInformation(value SecurityVendorInformationable)() {
    if m != nil {
        m.securityVendorInformation = value
    }
}
// RoleDefinitionsable 
type RoleDefinitionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAverageComparativeScore()(AverageComparativeScoreable)
    GetControlScore()(ControlScoreable)
    GetRoleDefinitionsMember1()(RoleDefinitionsMember1able)
    GetSecurityVendorInformation()(SecurityVendorInformationable)
    SetAverageComparativeScore(value AverageComparativeScoreable)()
    SetControlScore(value ControlScoreable)()
    SetRoleDefinitionsMember1(value RoleDefinitionsMember1able)()
    SetSecurityVendorInformation(value SecurityVendorInformationable)()
}
// NewSecureScore instantiates a new SecureScore and sets the default values.
func NewSecureScore()(*SecureScore) {
    m := &SecureScore{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSecureScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecureScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecureScore(), nil
}
// GetActiveUserCount gets the activeUserCount property value. Active user count of the given tenant.
func (m *SecureScore) GetActiveUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.activeUserCount
    }
}
// GetAverageComparativeScores gets the averageComparativeScores property value. Average score by different scopes (for example, average by industry, average by seating) and control category (Identity, Data, Device, Apps, Infrastructure) within the scope.
func (m *SecureScore) GetAverageComparativeScores()([]RoleDefinitionsable) {
    if m == nil {
        return nil
    } else {
        return m.averageComparativeScores
    }
}
// GetAzureTenantId gets the azureTenantId property value. GUID string for tenant ID.
func (m *SecureScore) GetAzureTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureTenantId
    }
}
// GetControlScores gets the controlScores property value. Contains tenant scores for a set of controls.
func (m *SecureScore) GetControlScores()([]RoleDefinitionsable) {
    if m == nil {
        return nil
    } else {
        return m.controlScores
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date when the entity is created.
func (m *SecureScore) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetCurrentScore gets the currentScore property value. Tenant current attained score on specified date.
func (m *SecureScore) GetCurrentScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.currentScore
    }
}
// GetEnabledServices gets the enabledServices property value. Microsoft-provided services for the tenant (for example, Exchange online, Skype, SharePoint).
func (m *SecureScore) GetEnabledServices()([]string) {
    if m == nil {
        return nil
    } else {
        return m.enabledServices
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecureScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveUserCount(val)
        }
        return nil
    }
    res["averageComparativeScores"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRoleDefinitionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleDefinitionsable, len(val))
            for i, v := range val {
                res[i] = v.(RoleDefinitionsable)
            }
            m.SetAverageComparativeScores(res)
        }
        return nil
    }
    res["azureTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureTenantId(val)
        }
        return nil
    }
    res["controlScores"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRoleDefinitionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleDefinitionsable, len(val))
            for i, v := range val {
                res[i] = v.(RoleDefinitionsable)
            }
            m.SetControlScores(res)
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
    res["currentScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentScore(val)
        }
        return nil
    }
    res["enabledServices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEnabledServices(res)
        }
        return nil
    }
    res["licensedUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicensedUserCount(val)
        }
        return nil
    }
    res["maxScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxScore(val)
        }
        return nil
    }
    res["vendorInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoleDefinitionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorInformation(val.(RoleDefinitionsable))
        }
        return nil
    }
    return res
}
// GetLicensedUserCount gets the licensedUserCount property value. Licensed user count of the given tenant.
func (m *SecureScore) GetLicensedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.licensedUserCount
    }
}
// GetMaxScore gets the maxScore property value. Tenant maximum possible score on specified date.
func (m *SecureScore) GetMaxScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.maxScore
    }
}
// GetVendorInformation gets the vendorInformation property value. Complex type containing details about the security product/service vendor, provider, and subprovider (for example, vendor=Microsoft; provider=SecureScore). Required.
func (m *SecureScore) GetVendorInformation()(RoleDefinitionsable) {
    if m == nil {
        return nil
    } else {
        return m.vendorInformation
    }
}
// Serialize serializes information the current object
func (m *SecureScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("activeUserCount", m.GetActiveUserCount())
        if err != nil {
            return err
        }
    }
    if m.GetAverageComparativeScores() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAverageComparativeScores()))
        for i, v := range m.GetAverageComparativeScores() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("averageComparativeScores", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureTenantId", m.GetAzureTenantId())
        if err != nil {
            return err
        }
    }
    if m.GetControlScores() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetControlScores()))
        for i, v := range m.GetControlScores() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("controlScores", cast)
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
        err = writer.WriteFloat64Value("currentScore", m.GetCurrentScore())
        if err != nil {
            return err
        }
    }
    if m.GetEnabledServices() != nil {
        err = writer.WriteCollectionOfStringValues("enabledServices", m.GetEnabledServices())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("licensedUserCount", m.GetLicensedUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("maxScore", m.GetMaxScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("vendorInformation", m.GetVendorInformation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveUserCount sets the activeUserCount property value. Active user count of the given tenant.
func (m *SecureScore) SetActiveUserCount(value *int32)() {
    if m != nil {
        m.activeUserCount = value
    }
}
// SetAverageComparativeScores sets the averageComparativeScores property value. Average score by different scopes (for example, average by industry, average by seating) and control category (Identity, Data, Device, Apps, Infrastructure) within the scope.
func (m *SecureScore) SetAverageComparativeScores(value []RoleDefinitionsable)() {
    if m != nil {
        m.averageComparativeScores = value
    }
}
// SetAzureTenantId sets the azureTenantId property value. GUID string for tenant ID.
func (m *SecureScore) SetAzureTenantId(value *string)() {
    if m != nil {
        m.azureTenantId = value
    }
}
// SetControlScores sets the controlScores property value. Contains tenant scores for a set of controls.
func (m *SecureScore) SetControlScores(value []RoleDefinitionsable)() {
    if m != nil {
        m.controlScores = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date when the entity is created.
func (m *SecureScore) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetCurrentScore sets the currentScore property value. Tenant current attained score on specified date.
func (m *SecureScore) SetCurrentScore(value *float64)() {
    if m != nil {
        m.currentScore = value
    }
}
// SetEnabledServices sets the enabledServices property value. Microsoft-provided services for the tenant (for example, Exchange online, Skype, SharePoint).
func (m *SecureScore) SetEnabledServices(value []string)() {
    if m != nil {
        m.enabledServices = value
    }
}
// SetLicensedUserCount sets the licensedUserCount property value. Licensed user count of the given tenant.
func (m *SecureScore) SetLicensedUserCount(value *int32)() {
    if m != nil {
        m.licensedUserCount = value
    }
}
// SetMaxScore sets the maxScore property value. Tenant maximum possible score on specified date.
func (m *SecureScore) SetMaxScore(value *float64)() {
    if m != nil {
        m.maxScore = value
    }
}
// SetVendorInformation sets the vendorInformation property value. Complex type containing details about the security product/service vendor, provider, and subprovider (for example, vendor=Microsoft; provider=SecureScore). Required.
func (m *SecureScore) SetVendorInformation(value RoleDefinitionsable)() {
    if m != nil {
        m.vendorInformation = value
    }
}
