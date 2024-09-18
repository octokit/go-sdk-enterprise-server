package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesConfigInitRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The content of your _.ghl_ license file.
    license *string
    // The root site administrator password.
    password *string
}
// NewGhesConfigInitRequest instantiates a new GhesConfigInitRequest and sets the default values.
func NewGhesConfigInitRequest()(*GhesConfigInitRequest) {
    m := &GhesConfigInitRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesConfigInitRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesConfigInitRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesConfigInitRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesConfigInitRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesConfigInitRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["license"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicense(val)
        }
        return nil
    }
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    return res
}
// GetLicense gets the license property value. The content of your _.ghl_ license file.
// returns a *string when successful
func (m *GhesConfigInitRequest) GetLicense()(*string) {
    return m.license
}
// GetPassword gets the password property value. The root site administrator password.
// returns a *string when successful
func (m *GhesConfigInitRequest) GetPassword()(*string) {
    return m.password
}
// Serialize serializes information the current object
func (m *GhesConfigInitRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("license", m.GetLicense())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("password", m.GetPassword())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GhesConfigInitRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLicense sets the license property value. The content of your _.ghl_ license file.
func (m *GhesConfigInitRequest) SetLicense(value *string)() {
    m.license = value
}
// SetPassword sets the password property value. The root site administrator password.
func (m *GhesConfigInitRequest) SetPassword(value *string)() {
    m.password = value
}
type GhesConfigInitRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLicense()(*string)
    GetPassword()(*string)
    SetLicense(value *string)()
    SetPassword(value *string)()
}
