package models

import (
    ie861254c8c989772edbc44dcb4211b4d571fc16370454a2a18f6fc023d7ef8c4 "objectdefinition"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ObjectDefinition 
type ObjectDefinition struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The attributes property
    attributes []SecureScoresable
    // The metadata property
    metadata []SecureScoresable
    // The name property
    name *string
    // The supportedApis property
    supportedApis []string
}
// SecureScores union type wrapper for classes attributeDefinition, secureScoresMember1
type SecureScores struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type attributeDefinition
    attributeDefinition AttributeDefinitionable
    // Union type representation for type metadataEntry
    metadataEntry MetadataEntryable
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
// GetAttributeDefinition gets the attributeDefinition property value. Union type representation for type attributeDefinition
func (m *SecureScores) GetAttributeDefinition()(AttributeDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.attributeDefinition
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecureScores) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attributeDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAttributeDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttributeDefinition(val.(AttributeDefinitionable))
        }
        return nil
    }
    res["metadataEntry"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMetadataEntryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadataEntry(val.(MetadataEntryable))
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
// GetMetadataEntry gets the metadataEntry property value. Union type representation for type metadataEntry
func (m *SecureScores) GetMetadataEntry()(MetadataEntryable) {
    if m == nil {
        return nil
    } else {
        return m.metadataEntry
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
        err := writer.WriteObjectValue("attributeDefinition", m.GetAttributeDefinition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("metadataEntry", m.GetMetadataEntry())
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
// SetAttributeDefinition sets the attributeDefinition property value. Union type representation for type attributeDefinition
func (m *SecureScores) SetAttributeDefinition(value AttributeDefinitionable)() {
    if m != nil {
        m.attributeDefinition = value
    }
}
// SetMetadataEntry sets the metadataEntry property value. Union type representation for type metadataEntry
func (m *SecureScores) SetMetadataEntry(value MetadataEntryable)() {
    if m != nil {
        m.metadataEntry = value
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
    GetAttributeDefinition()(AttributeDefinitionable)
    GetMetadataEntry()(MetadataEntryable)
    GetSecureScoresMember1()(SecureScoresMember1able)
    SetAttributeDefinition(value AttributeDefinitionable)()
    SetMetadataEntry(value MetadataEntryable)()
    SetSecureScoresMember1(value SecureScoresMember1able)()
}
// NewObjectDefinition instantiates a new objectDefinition and sets the default values.
func NewObjectDefinition()(*ObjectDefinition) {
    m := &ObjectDefinition{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateObjectDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateObjectDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewObjectDefinition(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ObjectDefinition) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAttributes gets the attributes property value. The attributes property
func (m *ObjectDefinition) GetAttributes()([]SecureScoresable) {
    if m == nil {
        return nil
    } else {
        return m.attributes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ObjectDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attributes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecureScoresFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecureScoresable, len(val))
            for i, v := range val {
                res[i] = v.(SecureScoresable)
            }
            m.SetAttributes(res)
        }
        return nil
    }
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecureScoresFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecureScoresable, len(val))
            for i, v := range val {
                res[i] = v.(SecureScoresable)
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
    res["supportedApis"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSupportedApis(res)
        }
        return nil
    }
    return res
}
// GetMetadata gets the metadata property value. The metadata property
func (m *ObjectDefinition) GetMetadata()([]SecureScoresable) {
    if m == nil {
        return nil
    } else {
        return m.metadata
    }
}
// GetName gets the name property value. The name property
func (m *ObjectDefinition) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetSupportedApis gets the supportedApis property value. The supportedApis property
func (m *ObjectDefinition) GetSupportedApis()([]string) {
    if m == nil {
        return nil
    } else {
        return m.supportedApis
    }
}
// Serialize serializes information the current object
func (m *ObjectDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAttributes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttributes()))
        for i, v := range m.GetAttributes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("attributes", cast)
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
    if m.GetSupportedApis() != nil {
        err := writer.WriteCollectionOfStringValues("supportedApis", m.GetSupportedApis())
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
func (m *ObjectDefinition) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAttributes sets the attributes property value. The attributes property
func (m *ObjectDefinition) SetAttributes(value []SecureScoresable)() {
    if m != nil {
        m.attributes = value
    }
}
// SetMetadata sets the metadata property value. The metadata property
func (m *ObjectDefinition) SetMetadata(value []SecureScoresable)() {
    if m != nil {
        m.metadata = value
    }
}
// SetName sets the name property value. The name property
func (m *ObjectDefinition) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetSupportedApis sets the supportedApis property value. The supportedApis property
func (m *ObjectDefinition) SetSupportedApis(value []string)() {
    if m != nil {
        m.supportedApis = value
    }
}
