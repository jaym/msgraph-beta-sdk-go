package getbyids

import (
    ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837 "msgraphbetasdkgo/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetByIdsResponse 
type GetByIdsResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The value property
    value []ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837.DirectoryObjectable
}
// NewGetByIdsResponse instantiates a new getByIdsResponse and sets the default values.
func NewGetByIdsResponse()(*GetByIdsResponse) {
    m := &GetByIdsResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetByIdsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetByIdsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetByIdsResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetByIdsResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetByIdsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837.CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837.DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837.DirectoryObjectable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *GetByIdsResponse) GetValue()([]ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837.DirectoryObjectable) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// Serialize serializes information the current object
func (m *GetByIdsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("value", cast)
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
func (m *GetByIdsResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetValue sets the value property value. The value property
func (m *GetByIdsResponse) SetValue(value []ic1cf40dc5fb70b4d4b54570999720cc0be6b4fc13aeb4599f8cf07d641e36837.DirectoryObjectable)() {
    if m != nil {
        m.value = value
    }
}