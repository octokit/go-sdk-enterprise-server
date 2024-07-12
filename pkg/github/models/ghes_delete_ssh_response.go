package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesDeleteSshResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Error message indicating the reason for the SSH key removal failure. This field is only present if the SSH key removal failed.
    error *string
    // Hostname of the node where the SSH key was removed.
    hostname *string
    // Message indicating the result of the SSH key removal.
    message *string
    // UUID of the SSH key that was removed.
    uuid *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
}
// NewGhesDeleteSshResponse instantiates a new GhesDeleteSshResponse and sets the default values.
func NewGhesDeleteSshResponse()(*GhesDeleteSshResponse) {
    m := &GhesDeleteSshResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesDeleteSshResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesDeleteSshResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesDeleteSshResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesDeleteSshResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetError gets the error property value. Error message indicating the reason for the SSH key removal failure. This field is only present if the SSH key removal failed.
// returns a *string when successful
func (m *GhesDeleteSshResponse) GetError()(*string) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesDeleteSshResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val)
        }
        return nil
    }
    res["hostname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostname(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["uuid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUuid(val)
        }
        return nil
    }
    return res
}
// GetHostname gets the hostname property value. Hostname of the node where the SSH key was removed.
// returns a *string when successful
func (m *GhesDeleteSshResponse) GetHostname()(*string) {
    return m.hostname
}
// GetMessage gets the message property value. Message indicating the result of the SSH key removal.
// returns a *string when successful
func (m *GhesDeleteSshResponse) GetMessage()(*string) {
    return m.message
}
// GetUuid gets the uuid property value. UUID of the SSH key that was removed.
// returns a *UUID when successful
func (m *GhesDeleteSshResponse) GetUuid()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.uuid
}
// Serialize serializes information the current object
func (m *GhesDeleteSshResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hostname", m.GetHostname())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("uuid", m.GetUuid())
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
func (m *GhesDeleteSshResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetError sets the error property value. Error message indicating the reason for the SSH key removal failure. This field is only present if the SSH key removal failed.
func (m *GhesDeleteSshResponse) SetError(value *string)() {
    m.error = value
}
// SetHostname sets the hostname property value. Hostname of the node where the SSH key was removed.
func (m *GhesDeleteSshResponse) SetHostname(value *string)() {
    m.hostname = value
}
// SetMessage sets the message property value. Message indicating the result of the SSH key removal.
func (m *GhesDeleteSshResponse) SetMessage(value *string)() {
    m.message = value
}
// SetUuid sets the uuid property value. UUID of the SSH key that was removed.
func (m *GhesDeleteSshResponse) SetUuid(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.uuid = value
}
type GhesDeleteSshResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetError()(*string)
    GetHostname()(*string)
    GetMessage()(*string)
    GetUuid()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    SetError(value *string)()
    SetHostname(value *string)()
    SetMessage(value *string)()
    SetUuid(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
}
