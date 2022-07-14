package models

import (
    i4c85e2815974a1243ec30a23463567213d7a147ae794dbe7bf7e6eae36688582 "optionalclaims"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OptionalClaims 
type OptionalClaims struct {
    // The optional claims returned in the JWT access token.
    accessToken []PermissionGrantPoliciesable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The optional claims returned in the JWT ID token.
    idToken []PermissionGrantPoliciesable
    // The optional claims returned in the SAML token.
    saml2Token []PermissionGrantPoliciesable
}
// PermissionGrantPolicies union type wrapper for classes optionalClaim, permissionGrantPoliciesMember1
type PermissionGrantPolicies struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type optionalClaim
    optionalClaim OptionalClaimable
    // Union type representation for type permissionGrantPoliciesMember1
    permissionGrantPoliciesMember1 PermissionGrantPoliciesMember1able
}
// NewPermissionGrantPolicies instantiates a new permissionGrantPolicies and sets the default values.
func NewPermissionGrantPolicies()(*PermissionGrantPolicies) {
    m := &PermissionGrantPolicies{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePermissionGrantPoliciesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePermissionGrantPoliciesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPermissionGrantPolicies(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PermissionGrantPolicies) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PermissionGrantPolicies) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["optionalClaim"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOptionalClaimFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptionalClaim(val.(OptionalClaimable))
        }
        return nil
    }
    res["permissionGrantPoliciesMember1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePermissionGrantPoliciesMember1FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionGrantPoliciesMember1(val.(PermissionGrantPoliciesMember1able))
        }
        return nil
    }
    return res
}
// GetOptionalClaim gets the optionalClaim property value. Union type representation for type optionalClaim
func (m *PermissionGrantPolicies) GetOptionalClaim()(OptionalClaimable) {
    if m == nil {
        return nil
    } else {
        return m.optionalClaim
    }
}
// GetPermissionGrantPoliciesMember1 gets the permissionGrantPoliciesMember1 property value. Union type representation for type permissionGrantPoliciesMember1
func (m *PermissionGrantPolicies) GetPermissionGrantPoliciesMember1()(PermissionGrantPoliciesMember1able) {
    if m == nil {
        return nil
    } else {
        return m.permissionGrantPoliciesMember1
    }
}
// Serialize serializes information the current object
func (m *PermissionGrantPolicies) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("optionalClaim", m.GetOptionalClaim())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("permissionGrantPoliciesMember1", m.GetPermissionGrantPoliciesMember1())
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
func (m *PermissionGrantPolicies) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetOptionalClaim sets the optionalClaim property value. Union type representation for type optionalClaim
func (m *PermissionGrantPolicies) SetOptionalClaim(value OptionalClaimable)() {
    if m != nil {
        m.optionalClaim = value
    }
}
// SetPermissionGrantPoliciesMember1 sets the permissionGrantPoliciesMember1 property value. Union type representation for type permissionGrantPoliciesMember1
func (m *PermissionGrantPolicies) SetPermissionGrantPoliciesMember1(value PermissionGrantPoliciesMember1able)() {
    if m != nil {
        m.permissionGrantPoliciesMember1 = value
    }
}
// PermissionGrantPoliciesable 
type PermissionGrantPoliciesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOptionalClaim()(OptionalClaimable)
    GetPermissionGrantPoliciesMember1()(PermissionGrantPoliciesMember1able)
    SetOptionalClaim(value OptionalClaimable)()
    SetPermissionGrantPoliciesMember1(value PermissionGrantPoliciesMember1able)()
}
// NewOptionalClaims instantiates a new optionalClaims and sets the default values.
func NewOptionalClaims()(*OptionalClaims) {
    m := &OptionalClaims{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOptionalClaimsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOptionalClaimsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOptionalClaims(), nil
}
// GetAccessToken gets the accessToken property value. The optional claims returned in the JWT access token.
func (m *OptionalClaims) GetAccessToken()([]PermissionGrantPoliciesable) {
    if m == nil {
        return nil
    } else {
        return m.accessToken
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OptionalClaims) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OptionalClaims) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionGrantPoliciesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PermissionGrantPoliciesable, len(val))
            for i, v := range val {
                res[i] = v.(PermissionGrantPoliciesable)
            }
            m.SetAccessToken(res)
        }
        return nil
    }
    res["idToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionGrantPoliciesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PermissionGrantPoliciesable, len(val))
            for i, v := range val {
                res[i] = v.(PermissionGrantPoliciesable)
            }
            m.SetIdToken(res)
        }
        return nil
    }
    res["saml2Token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionGrantPoliciesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PermissionGrantPoliciesable, len(val))
            for i, v := range val {
                res[i] = v.(PermissionGrantPoliciesable)
            }
            m.SetSaml2Token(res)
        }
        return nil
    }
    return res
}
// GetIdToken gets the idToken property value. The optional claims returned in the JWT ID token.
func (m *OptionalClaims) GetIdToken()([]PermissionGrantPoliciesable) {
    if m == nil {
        return nil
    } else {
        return m.idToken
    }
}
// GetSaml2Token gets the saml2Token property value. The optional claims returned in the SAML token.
func (m *OptionalClaims) GetSaml2Token()([]PermissionGrantPoliciesable) {
    if m == nil {
        return nil
    } else {
        return m.saml2Token
    }
}
// Serialize serializes information the current object
func (m *OptionalClaims) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccessToken() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessToken()))
        for i, v := range m.GetAccessToken() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("accessToken", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIdToken() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIdToken()))
        for i, v := range m.GetIdToken() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("idToken", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSaml2Token() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSaml2Token()))
        for i, v := range m.GetSaml2Token() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("saml2Token", cast)
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
// SetAccessToken sets the accessToken property value. The optional claims returned in the JWT access token.
func (m *OptionalClaims) SetAccessToken(value []PermissionGrantPoliciesable)() {
    if m != nil {
        m.accessToken = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OptionalClaims) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIdToken sets the idToken property value. The optional claims returned in the JWT ID token.
func (m *OptionalClaims) SetIdToken(value []PermissionGrantPoliciesable)() {
    if m != nil {
        m.idToken = value
    }
}
// SetSaml2Token sets the saml2Token property value. The optional claims returned in the SAML token.
func (m *OptionalClaims) SetSaml2Token(value []PermissionGrantPoliciesable)() {
    if m != nil {
        m.saml2Token = value
    }
}
