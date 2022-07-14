package models

import (
    ia37373c96814c8dcc7e756cf9aefbf2cf54e63d9eaa8f3c5e92e267ca8073c0d "onpremisesapplicationsegment"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesApplicationSegment 
type OnPremisesApplicationSegment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The alternateUrl property
    alternateUrl *string
    // The corsConfigurations property
    corsConfigurations []Applicationsable
    // The externalUrl property
    externalUrl *string
    // The internalUrl property
    internalUrl *string
}
// Applications union type wrapper for classes corsConfiguration, applicationsMember1
type Applications struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type applicationsMember1
    applicationsMember1 ApplicationsMember1able
    // Union type representation for type corsConfiguration
    corsConfiguration CorsConfigurationable
}
// NewApplications instantiates a new applications and sets the default values.
func NewApplications()(*Applications) {
    m := &Applications{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateApplicationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplicationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplications(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Applications) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplicationsMember1 gets the applicationsMember1 property value. Union type representation for type applicationsMember1
func (m *Applications) GetApplicationsMember1()(ApplicationsMember1able) {
    if m == nil {
        return nil
    } else {
        return m.applicationsMember1
    }
}
// GetCorsConfiguration gets the corsConfiguration property value. Union type representation for type corsConfiguration
func (m *Applications) GetCorsConfiguration()(CorsConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.corsConfiguration
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Applications) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicationsMember1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApplicationsMember1FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationsMember1(val.(ApplicationsMember1able))
        }
        return nil
    }
    res["corsConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCorsConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorsConfiguration(val.(CorsConfigurationable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *Applications) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("applicationsMember1", m.GetApplicationsMember1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("corsConfiguration", m.GetCorsConfiguration())
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
func (m *Applications) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplicationsMember1 sets the applicationsMember1 property value. Union type representation for type applicationsMember1
func (m *Applications) SetApplicationsMember1(value ApplicationsMember1able)() {
    if m != nil {
        m.applicationsMember1 = value
    }
}
// SetCorsConfiguration sets the corsConfiguration property value. Union type representation for type corsConfiguration
func (m *Applications) SetCorsConfiguration(value CorsConfigurationable)() {
    if m != nil {
        m.corsConfiguration = value
    }
}
// Applicationsable 
type Applicationsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationsMember1()(ApplicationsMember1able)
    GetCorsConfiguration()(CorsConfigurationable)
    SetApplicationsMember1(value ApplicationsMember1able)()
    SetCorsConfiguration(value CorsConfigurationable)()
}
// NewOnPremisesApplicationSegment instantiates a new onPremisesApplicationSegment and sets the default values.
func NewOnPremisesApplicationSegment()(*OnPremisesApplicationSegment) {
    m := &OnPremisesApplicationSegment{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOnPremisesApplicationSegmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesApplicationSegmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesApplicationSegment(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesApplicationSegment) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAlternateUrl gets the alternateUrl property value. The alternateUrl property
func (m *OnPremisesApplicationSegment) GetAlternateUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alternateUrl
    }
}
// GetCorsConfigurations gets the corsConfigurations property value. The corsConfigurations property
func (m *OnPremisesApplicationSegment) GetCorsConfigurations()([]Applicationsable) {
    if m == nil {
        return nil
    } else {
        return m.corsConfigurations
    }
}
// GetExternalUrl gets the externalUrl property value. The externalUrl property
func (m *OnPremisesApplicationSegment) GetExternalUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesApplicationSegment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alternateUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlternateUrl(val)
        }
        return nil
    }
    res["corsConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApplicationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Applicationsable, len(val))
            for i, v := range val {
                res[i] = v.(Applicationsable)
            }
            m.SetCorsConfigurations(res)
        }
        return nil
    }
    res["externalUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalUrl(val)
        }
        return nil
    }
    res["internalUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalUrl(val)
        }
        return nil
    }
    return res
}
// GetInternalUrl gets the internalUrl property value. The internalUrl property
func (m *OnPremisesApplicationSegment) GetInternalUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.internalUrl
    }
}
// Serialize serializes information the current object
func (m *OnPremisesApplicationSegment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("alternateUrl", m.GetAlternateUrl())
        if err != nil {
            return err
        }
    }
    if m.GetCorsConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCorsConfigurations()))
        for i, v := range m.GetCorsConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("corsConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalUrl", m.GetExternalUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("internalUrl", m.GetInternalUrl())
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
func (m *OnPremisesApplicationSegment) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAlternateUrl sets the alternateUrl property value. The alternateUrl property
func (m *OnPremisesApplicationSegment) SetAlternateUrl(value *string)() {
    if m != nil {
        m.alternateUrl = value
    }
}
// SetCorsConfigurations sets the corsConfigurations property value. The corsConfigurations property
func (m *OnPremisesApplicationSegment) SetCorsConfigurations(value []Applicationsable)() {
    if m != nil {
        m.corsConfigurations = value
    }
}
// SetExternalUrl sets the externalUrl property value. The externalUrl property
func (m *OnPremisesApplicationSegment) SetExternalUrl(value *string)() {
    if m != nil {
        m.externalUrl = value
    }
}
// SetInternalUrl sets the internalUrl property value. The internalUrl property
func (m *OnPremisesApplicationSegment) SetInternalUrl(value *string)() {
    if m != nil {
        m.internalUrl = value
    }
}
