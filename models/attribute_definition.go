package models

import (
    i62758439057444b51391f17c08afbeacb81186695ff1399f4e9d2eb3db26660b "attributedefinition"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttributeDefinition 
type AttributeDefinition struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // true if the attribute should be used as the anchor for the object. Anchor attributes must have a unique value identifying an object, and must be immutable. Default is false. One, and only one, of the object's attributes must be designated as the anchor to support synchronization.
    anchor *bool
    // The apiExpressions property
    apiExpressions []SecureScoresable
    // true if value of this attribute should be treated as case-sensitive. This setting affects how the synchronization engine detects changes for the attribute.
    caseExact *bool
    // The defaultValue property
    defaultValue *string
    // 'true' to allow null values for attributes.
    flowNullValues *bool
    // Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
    metadata []SecureScoresable
    // true if an attribute can have multiple values. Default is false.
    multivalued *bool
    // The mutability property
    mutability *Mutability
    // Name of the attribute. Must be unique within the object definition. Not nullable.
    name *string
    // For attributes with reference type, lists referenced objects (for example, the manager attribute would list User as the referenced object).
    referencedObjects []SecureScoresable
    // true if attribute is required. Object can not be created if any of the required attributes are missing. If during synchronization, the required attribute has no value, the default value will be used. If default the value was not set, synchronization will record an error.
    required *bool
    // The type property
    type_escaped *AttributeType
}
// SecureScores union type wrapper for classes referencedObject, secureScoresMember1
type SecureScores struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type metadataEntry
    metadataEntry MetadataEntryable
    // Union type representation for type referencedObject
    referencedObject ReferencedObjectable
    // Union type representation for type secureScoresMember1
    secureScoresMember1 SecureScoresMember1able
    // Union type representation for type stringKeyStringValuePair
    stringKeyStringValuePair StringKeyStringValuePairable
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
    res["referencedObject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateReferencedObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferencedObject(val.(ReferencedObjectable))
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
    res["stringKeyStringValuePair"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStringKeyStringValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStringKeyStringValuePair(val.(StringKeyStringValuePairable))
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
// GetReferencedObject gets the referencedObject property value. Union type representation for type referencedObject
func (m *SecureScores) GetReferencedObject()(ReferencedObjectable) {
    if m == nil {
        return nil
    } else {
        return m.referencedObject
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
// GetStringKeyStringValuePair gets the stringKeyStringValuePair property value. Union type representation for type stringKeyStringValuePair
func (m *SecureScores) GetStringKeyStringValuePair()(StringKeyStringValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.stringKeyStringValuePair
    }
}
// Serialize serializes information the current object
func (m *SecureScores) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("metadataEntry", m.GetMetadataEntry())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("referencedObject", m.GetReferencedObject())
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
        err := writer.WriteObjectValue("stringKeyStringValuePair", m.GetStringKeyStringValuePair())
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
// SetMetadataEntry sets the metadataEntry property value. Union type representation for type metadataEntry
func (m *SecureScores) SetMetadataEntry(value MetadataEntryable)() {
    if m != nil {
        m.metadataEntry = value
    }
}
// SetReferencedObject sets the referencedObject property value. Union type representation for type referencedObject
func (m *SecureScores) SetReferencedObject(value ReferencedObjectable)() {
    if m != nil {
        m.referencedObject = value
    }
}
// SetSecureScoresMember1 sets the secureScoresMember1 property value. Union type representation for type secureScoresMember1
func (m *SecureScores) SetSecureScoresMember1(value SecureScoresMember1able)() {
    if m != nil {
        m.secureScoresMember1 = value
    }
}
// SetStringKeyStringValuePair sets the stringKeyStringValuePair property value. Union type representation for type stringKeyStringValuePair
func (m *SecureScores) SetStringKeyStringValuePair(value StringKeyStringValuePairable)() {
    if m != nil {
        m.stringKeyStringValuePair = value
    }
}
// SecureScoresable 
type SecureScoresable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMetadataEntry()(MetadataEntryable)
    GetReferencedObject()(ReferencedObjectable)
    GetSecureScoresMember1()(SecureScoresMember1able)
    GetStringKeyStringValuePair()(StringKeyStringValuePairable)
    SetMetadataEntry(value MetadataEntryable)()
    SetReferencedObject(value ReferencedObjectable)()
    SetSecureScoresMember1(value SecureScoresMember1able)()
    SetStringKeyStringValuePair(value StringKeyStringValuePairable)()
}
// NewAttributeDefinition instantiates a new attributeDefinition and sets the default values.
func NewAttributeDefinition()(*AttributeDefinition) {
    m := &AttributeDefinition{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAttributeDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttributeDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttributeDefinition(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttributeDefinition) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAnchor gets the anchor property value. true if the attribute should be used as the anchor for the object. Anchor attributes must have a unique value identifying an object, and must be immutable. Default is false. One, and only one, of the object's attributes must be designated as the anchor to support synchronization.
func (m *AttributeDefinition) GetAnchor()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.anchor
    }
}
// GetApiExpressions gets the apiExpressions property value. The apiExpressions property
func (m *AttributeDefinition) GetApiExpressions()([]SecureScoresable) {
    if m == nil {
        return nil
    } else {
        return m.apiExpressions
    }
}
// GetCaseExact gets the caseExact property value. true if value of this attribute should be treated as case-sensitive. This setting affects how the synchronization engine detects changes for the attribute.
func (m *AttributeDefinition) GetCaseExact()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.caseExact
    }
}
// GetDefaultValue gets the defaultValue property value. The defaultValue property
func (m *AttributeDefinition) GetDefaultValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultValue
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttributeDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["anchor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnchor(val)
        }
        return nil
    }
    res["apiExpressions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecureScoresFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecureScoresable, len(val))
            for i, v := range val {
                res[i] = v.(SecureScoresable)
            }
            m.SetApiExpressions(res)
        }
        return nil
    }
    res["caseExact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaseExact(val)
        }
        return nil
    }
    res["defaultValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val)
        }
        return nil
    }
    res["flowNullValues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFlowNullValues(val)
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
    res["multivalued"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMultivalued(val)
        }
        return nil
    }
    res["mutability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMutability)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMutability(val.(*Mutability))
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
    res["referencedObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecureScoresFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecureScoresable, len(val))
            for i, v := range val {
                res[i] = v.(SecureScoresable)
            }
            m.SetReferencedObjects(res)
        }
        return nil
    }
    res["required"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequired(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttributeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*AttributeType))
        }
        return nil
    }
    return res
}
// GetFlowNullValues gets the flowNullValues property value. 'true' to allow null values for attributes.
func (m *AttributeDefinition) GetFlowNullValues()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.flowNullValues
    }
}
// GetMetadata gets the metadata property value. Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
func (m *AttributeDefinition) GetMetadata()([]SecureScoresable) {
    if m == nil {
        return nil
    } else {
        return m.metadata
    }
}
// GetMultivalued gets the multivalued property value. true if an attribute can have multiple values. Default is false.
func (m *AttributeDefinition) GetMultivalued()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.multivalued
    }
}
// GetMutability gets the mutability property value. The mutability property
func (m *AttributeDefinition) GetMutability()(*Mutability) {
    if m == nil {
        return nil
    } else {
        return m.mutability
    }
}
// GetName gets the name property value. Name of the attribute. Must be unique within the object definition. Not nullable.
func (m *AttributeDefinition) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetReferencedObjects gets the referencedObjects property value. For attributes with reference type, lists referenced objects (for example, the manager attribute would list User as the referenced object).
func (m *AttributeDefinition) GetReferencedObjects()([]SecureScoresable) {
    if m == nil {
        return nil
    } else {
        return m.referencedObjects
    }
}
// GetRequired gets the required property value. true if attribute is required. Object can not be created if any of the required attributes are missing. If during synchronization, the required attribute has no value, the default value will be used. If default the value was not set, synchronization will record an error.
func (m *AttributeDefinition) GetRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.required
    }
}
// GetType gets the type property value. The type property
func (m *AttributeDefinition) GetType()(*AttributeType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *AttributeDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("anchor", m.GetAnchor())
        if err != nil {
            return err
        }
    }
    if m.GetApiExpressions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApiExpressions()))
        for i, v := range m.GetApiExpressions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("apiExpressions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("caseExact", m.GetCaseExact())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("flowNullValues", m.GetFlowNullValues())
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
        err := writer.WriteBoolValue("multivalued", m.GetMultivalued())
        if err != nil {
            return err
        }
    }
    if m.GetMutability() != nil {
        cast := (*m.GetMutability()).String()
        err := writer.WriteStringValue("mutability", &cast)
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
    if m.GetReferencedObjects() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReferencedObjects()))
        for i, v := range m.GetReferencedObjects() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("referencedObjects", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("required", m.GetRequired())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *AttributeDefinition) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAnchor sets the anchor property value. true if the attribute should be used as the anchor for the object. Anchor attributes must have a unique value identifying an object, and must be immutable. Default is false. One, and only one, of the object's attributes must be designated as the anchor to support synchronization.
func (m *AttributeDefinition) SetAnchor(value *bool)() {
    if m != nil {
        m.anchor = value
    }
}
// SetApiExpressions sets the apiExpressions property value. The apiExpressions property
func (m *AttributeDefinition) SetApiExpressions(value []SecureScoresable)() {
    if m != nil {
        m.apiExpressions = value
    }
}
// SetCaseExact sets the caseExact property value. true if value of this attribute should be treated as case-sensitive. This setting affects how the synchronization engine detects changes for the attribute.
func (m *AttributeDefinition) SetCaseExact(value *bool)() {
    if m != nil {
        m.caseExact = value
    }
}
// SetDefaultValue sets the defaultValue property value. The defaultValue property
func (m *AttributeDefinition) SetDefaultValue(value *string)() {
    if m != nil {
        m.defaultValue = value
    }
}
// SetFlowNullValues sets the flowNullValues property value. 'true' to allow null values for attributes.
func (m *AttributeDefinition) SetFlowNullValues(value *bool)() {
    if m != nil {
        m.flowNullValues = value
    }
}
// SetMetadata sets the metadata property value. Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
func (m *AttributeDefinition) SetMetadata(value []SecureScoresable)() {
    if m != nil {
        m.metadata = value
    }
}
// SetMultivalued sets the multivalued property value. true if an attribute can have multiple values. Default is false.
func (m *AttributeDefinition) SetMultivalued(value *bool)() {
    if m != nil {
        m.multivalued = value
    }
}
// SetMutability sets the mutability property value. The mutability property
func (m *AttributeDefinition) SetMutability(value *Mutability)() {
    if m != nil {
        m.mutability = value
    }
}
// SetName sets the name property value. Name of the attribute. Must be unique within the object definition. Not nullable.
func (m *AttributeDefinition) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetReferencedObjects sets the referencedObjects property value. For attributes with reference type, lists referenced objects (for example, the manager attribute would list User as the referenced object).
func (m *AttributeDefinition) SetReferencedObjects(value []SecureScoresable)() {
    if m != nil {
        m.referencedObjects = value
    }
}
// SetRequired sets the required property value. true if attribute is required. Object can not be created if any of the required attributes are missing. If during synchronization, the required attribute has no value, the default value will be used. If default the value was not set, synchronization will record an error.
func (m *AttributeDefinition) SetRequired(value *bool)() {
    if m != nil {
        m.required = value
    }
}
// SetType sets the type property value. The type property
func (m *AttributeDefinition) SetType(value *AttributeType)() {
    if m != nil {
        m.type_escaped = value
    }
}
