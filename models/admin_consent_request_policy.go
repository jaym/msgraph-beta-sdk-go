package models

import (
    i97652f2af6ecfcdde39b09d5596e701a32c51a726e73110e5f71ac12a911d07b "adminconsentrequestpolicy"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdminConsentRequestPolicy 
type AdminConsentRequestPolicy struct {
    Entity
    // Specifies whether the admin consent request feature is enabled or disabled. Required.
    isEnabled *bool
    // Specifies whether reviewers will receive notifications. Required.
    notifyReviewers *bool
    // Specifies whether reviewers will receive reminder emails. Required.
    remindersEnabled *bool
    // Specifies the duration the request is active before it automatically expires if no decision is applied.
    requestDurationInDays *int32
    // Required.
    reviewers []AdminConsentRequestPolicyWrapperable
    // Specifies the version of this policy. When the policy is updated, this version is updated. Read-only.
    version *int32
}
// AdminConsentRequestPolicyWrapper union type wrapper for classes accessReviewReviewerScope, adminConsentRequestPolicyMember1
type AdminConsentRequestPolicyWrapper struct {
    // Union type representation for type accessReviewReviewerScope
    accessReviewReviewerScope AccessReviewReviewerScopeable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type adminConsentRequestPolicyMember1
    adminConsentRequestPolicyMember1 AdminConsentRequestPolicyMember1able
}
// NewAdminConsentRequestPolicyWrapper instantiates a new adminConsentRequestPolicyWrapper and sets the default values.
func NewAdminConsentRequestPolicyWrapper()(*AdminConsentRequestPolicyWrapper) {
    m := &AdminConsentRequestPolicyWrapper{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAdminConsentRequestPolicyWrapperFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdminConsentRequestPolicyWrapperFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdminConsentRequestPolicyWrapper(), nil
}
// GetAccessReviewReviewerScope gets the accessReviewReviewerScope property value. Union type representation for type accessReviewReviewerScope
func (m *AdminConsentRequestPolicyWrapper) GetAccessReviewReviewerScope()(AccessReviewReviewerScopeable) {
    if m == nil {
        return nil
    } else {
        return m.accessReviewReviewerScope
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AdminConsentRequestPolicyWrapper) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAdminConsentRequestPolicyMember1 gets the adminConsentRequestPolicyMember1 property value. Union type representation for type adminConsentRequestPolicyMember1
func (m *AdminConsentRequestPolicyWrapper) GetAdminConsentRequestPolicyMember1()(AdminConsentRequestPolicyMember1able) {
    if m == nil {
        return nil
    } else {
        return m.adminConsentRequestPolicyMember1
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdminConsentRequestPolicyWrapper) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessReviewReviewerScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessReviewReviewerScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessReviewReviewerScope(val.(AccessReviewReviewerScopeable))
        }
        return nil
    }
    res["adminConsentRequestPolicyMember1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAdminConsentRequestPolicyMember1FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminConsentRequestPolicyMember1(val.(AdminConsentRequestPolicyMember1able))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AdminConsentRequestPolicyWrapper) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("accessReviewReviewerScope", m.GetAccessReviewReviewerScope())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("adminConsentRequestPolicyMember1", m.GetAdminConsentRequestPolicyMember1())
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
// SetAccessReviewReviewerScope sets the accessReviewReviewerScope property value. Union type representation for type accessReviewReviewerScope
func (m *AdminConsentRequestPolicyWrapper) SetAccessReviewReviewerScope(value AccessReviewReviewerScopeable)() {
    if m != nil {
        m.accessReviewReviewerScope = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AdminConsentRequestPolicyWrapper) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAdminConsentRequestPolicyMember1 sets the adminConsentRequestPolicyMember1 property value. Union type representation for type adminConsentRequestPolicyMember1
func (m *AdminConsentRequestPolicyWrapper) SetAdminConsentRequestPolicyMember1(value AdminConsentRequestPolicyMember1able)() {
    if m != nil {
        m.adminConsentRequestPolicyMember1 = value
    }
}
// AdminConsentRequestPolicyWrapperable 
type AdminConsentRequestPolicyWrapperable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessReviewReviewerScope()(AccessReviewReviewerScopeable)
    GetAdminConsentRequestPolicyMember1()(AdminConsentRequestPolicyMember1able)
    SetAccessReviewReviewerScope(value AccessReviewReviewerScopeable)()
    SetAdminConsentRequestPolicyMember1(value AdminConsentRequestPolicyMember1able)()
}
// NewAdminConsentRequestPolicy instantiates a new adminConsentRequestPolicy and sets the default values.
func NewAdminConsentRequestPolicy()(*AdminConsentRequestPolicy) {
    m := &AdminConsentRequestPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAdminConsentRequestPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdminConsentRequestPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdminConsentRequestPolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdminConsentRequestPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["notifyReviewers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotifyReviewers(val)
        }
        return nil
    }
    res["remindersEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemindersEnabled(val)
        }
        return nil
    }
    res["requestDurationInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestDurationInDays(val)
        }
        return nil
    }
    res["reviewers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAdminConsentRequestPolicyWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AdminConsentRequestPolicyWrapperable, len(val))
            for i, v := range val {
                res[i] = v.(AdminConsentRequestPolicyWrapperable)
            }
            m.SetReviewers(res)
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
// GetIsEnabled gets the isEnabled property value. Specifies whether the admin consent request feature is enabled or disabled. Required.
func (m *AdminConsentRequestPolicy) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetNotifyReviewers gets the notifyReviewers property value. Specifies whether reviewers will receive notifications. Required.
func (m *AdminConsentRequestPolicy) GetNotifyReviewers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.notifyReviewers
    }
}
// GetRemindersEnabled gets the remindersEnabled property value. Specifies whether reviewers will receive reminder emails. Required.
func (m *AdminConsentRequestPolicy) GetRemindersEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.remindersEnabled
    }
}
// GetRequestDurationInDays gets the requestDurationInDays property value. Specifies the duration the request is active before it automatically expires if no decision is applied.
func (m *AdminConsentRequestPolicy) GetRequestDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.requestDurationInDays
    }
}
// GetReviewers gets the reviewers property value. Required.
func (m *AdminConsentRequestPolicy) GetReviewers()([]AdminConsentRequestPolicyWrapperable) {
    if m == nil {
        return nil
    } else {
        return m.reviewers
    }
}
// GetVersion gets the version property value. Specifies the version of this policy. When the policy is updated, this version is updated. Read-only.
func (m *AdminConsentRequestPolicy) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// Serialize serializes information the current object
func (m *AdminConsentRequestPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("notifyReviewers", m.GetNotifyReviewers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("remindersEnabled", m.GetRemindersEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("requestDurationInDays", m.GetRequestDurationInDays())
        if err != nil {
            return err
        }
    }
    if m.GetReviewers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReviewers()))
        for i, v := range m.GetReviewers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("reviewers", cast)
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
// SetIsEnabled sets the isEnabled property value. Specifies whether the admin consent request feature is enabled or disabled. Required.
func (m *AdminConsentRequestPolicy) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
// SetNotifyReviewers sets the notifyReviewers property value. Specifies whether reviewers will receive notifications. Required.
func (m *AdminConsentRequestPolicy) SetNotifyReviewers(value *bool)() {
    if m != nil {
        m.notifyReviewers = value
    }
}
// SetRemindersEnabled sets the remindersEnabled property value. Specifies whether reviewers will receive reminder emails. Required.
func (m *AdminConsentRequestPolicy) SetRemindersEnabled(value *bool)() {
    if m != nil {
        m.remindersEnabled = value
    }
}
// SetRequestDurationInDays sets the requestDurationInDays property value. Specifies the duration the request is active before it automatically expires if no decision is applied.
func (m *AdminConsentRequestPolicy) SetRequestDurationInDays(value *int32)() {
    if m != nil {
        m.requestDurationInDays = value
    }
}
// SetReviewers sets the reviewers property value. Required.
func (m *AdminConsentRequestPolicy) SetReviewers(value []AdminConsentRequestPolicyWrapperable)() {
    if m != nil {
        m.reviewers = value
    }
}
// SetVersion sets the version property value. Specifies the version of this policy. When the policy is updated, this version is updated. Read-only.
func (m *AdminConsentRequestPolicy) SetVersion(value *int32)() {
    if m != nil {
        m.version = value
    }
}
