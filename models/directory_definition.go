package models

import (
    i1cb342a555ce84307e361151fb7065a03742d88d55275c9e8e21d5fd02d105e2 "directorydefinition"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectoryDefinition 
type DirectoryDefinition struct {
    Entity
    // The discoverabilities property
    discoverabilities *DirectoryDefinitionDiscoverabilities
    // Represents the discovery date and time using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    discoveryDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Name of the directory. Must be unique within the synchronization schema. Not nullable.
    name *string
    // Collection of objects supported by the directory.
    objects []SecureScoresable
    // The readOnly property
    readOnly *bool
    // Read only value that indicates version discovered. null if discovery has not yet occurred.
    version *string
}
// SecureScores union type wrapper for classes objectDefinition, secureScoresMember1
type SecureScores struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type objectDefinition
    objectDefinition ObjectDefinitionable
    // Union type representation for type secureScoresMember1
    secureScoresMember1 SecureScoresMember1able
}
// NewSecureScores instantiates a new secureScores and sets the default values.
func NewSecureScores()(*SecureScores) {
    m := &SecureScores{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSecureScoresFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecureScoresFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecureScores(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecureScores) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecureScores) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["objectDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateObjectDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectDefinition(val.(ObjectDefinitionable))
        }
        return nil
    }
    res["secureScoresMember1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecureScoresMember1FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecureScoresMember1(val.(SecureScoresMember1able))
        }
        return nil
    }
    return res
}
// GetObjectDefinition gets the objectDefinition property value. Union type representation for type objectDefinition
func (m *SecureScores) GetObjectDefinition()(ObjectDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.objectDefinition
    }
}
// GetSecureScoresMember1 gets the secureScoresMember1 property value. Union type representation for type secureScoresMember1
func (m *SecureScores) GetSecureScoresMember1()(SecureScoresMember1able) {
    if m == nil {
        return nil
    } else {
        return m.secureScoresMember1
    }
}
// Serialize serializes information the current object
func (m *SecureScores) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("objectDefinition", m.GetObjectDefinition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("secureScoresMember1", m.GetSecureScoresMember1())
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
func (m *SecureScores) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetObjectDefinition sets the objectDefinition property value. Union type representation for type objectDefinition
func (m *SecureScores) SetObjectDefinition(value ObjectDefinitionable)() {
    if m != nil {
        m.objectDefinition = value
    }
}
// SetSecureScoresMember1 sets the secureScoresMember1 property value. Union type representation for type secureScoresMember1
func (m *SecureScores) SetSecureScoresMember1(value SecureScoresMember1able)() {
    if m != nil {
        m.secureScoresMember1 = value
    }
}
// SecureScoresable 
type SecureScoresable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetObjectDefinition()(ObjectDefinitionable)
    GetSecureScoresMember1()(SecureScoresMember1able)
    SetObjectDefinition(value ObjectDefinitionable)()
    SetSecureScoresMember1(value SecureScoresMember1able)()
}
// NewDirectoryDefinition instantiates a new DirectoryDefinition and sets the default values.
func NewDirectoryDefinition()(*DirectoryDefinition) {
    m := &DirectoryDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDirectoryDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectoryDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryDefinition(), nil
}
// GetDiscoverabilities gets the discoverabilities property value. The discoverabilities property
func (m *DirectoryDefinition) GetDiscoverabilities()(*DirectoryDefinitionDiscoverabilities) {
    if m == nil {
        return nil
    } else {
        return m.discoverabilities
    }
}
// GetDiscoveryDateTime gets the discoveryDateTime property value. Represents the discovery date and time using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *DirectoryDefinition) GetDiscoveryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.discoveryDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectoryDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["discoverabilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDirectoryDefinitionDiscoverabilities)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscoverabilities(val.(*DirectoryDefinitionDiscoverabilities))
        }
        return nil
    }
    res["discoveryDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscoveryDateTime(val)
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
    res["objects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecureScoresFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecureScoresable, len(val))
            for i, v := range val {
                res[i] = v.(SecureScoresable)
            }
            m.SetObjects(res)
        }
        return nil
    }
    res["readOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReadOnly(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
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
// GetName gets the name property value. Name of the directory. Must be unique within the synchronization schema. Not nullable.
func (m *DirectoryDefinition) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetObjects gets the objects property value. Collection of objects supported by the directory.
func (m *DirectoryDefinition) GetObjects()([]SecureScoresable) {
    if m == nil {
        return nil
    } else {
        return m.objects
    }
}
// GetReadOnly gets the readOnly property value. The readOnly property
func (m *DirectoryDefinition) GetReadOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.readOnly
    }
}
// GetVersion gets the version property value. Read only value that indicates version discovered. null if discovery has not yet occurred.
func (m *DirectoryDefinition) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// Serialize serializes information the current object
func (m *DirectoryDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDiscoverabilities() != nil {
        cast := (*m.GetDiscoverabilities()).String()
        err = writer.WriteStringValue("discoverabilities", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("discoveryDateTime", m.GetDiscoveryDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetObjects() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetObjects()))
        for i, v := range m.GetObjects() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("objects", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("readOnly", m.GetReadOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDiscoverabilities sets the discoverabilities property value. The discoverabilities property
func (m *DirectoryDefinition) SetDiscoverabilities(value *DirectoryDefinitionDiscoverabilities)() {
    if m != nil {
        m.discoverabilities = value
    }
}
// SetDiscoveryDateTime sets the discoveryDateTime property value. Represents the discovery date and time using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *DirectoryDefinition) SetDiscoveryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.discoveryDateTime = value
    }
}
// SetName sets the name property value. Name of the directory. Must be unique within the synchronization schema. Not nullable.
func (m *DirectoryDefinition) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetObjects sets the objects property value. Collection of objects supported by the directory.
func (m *DirectoryDefinition) SetObjects(value []SecureScoresable)() {
    if m != nil {
        m.objects = value
    }
}
// SetReadOnly sets the readOnly property value. The readOnly property
func (m *DirectoryDefinition) SetReadOnly(value *bool)() {
    if m != nil {
        m.readOnly = value
    }
}
// SetVersion sets the version property value. Read only value that indicates version discovered. null if discovery has not yet occurred.
func (m *DirectoryDefinition) SetVersion(value *string)() {
    if m != nil {
        m.version = value
    }
}
