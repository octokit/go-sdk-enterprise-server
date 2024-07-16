package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesSetMaintenanceResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The error property
    error *string
    // The hostname property
    hostname *string
    // The message property
    message *string
    // The uuid property
    uuid *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
}
// NewGhesSetMaintenanceResponse instantiates a new GhesSetMaintenanceResponse and sets the default values.
func NewGhesSetMaintenanceResponse()(*GhesSetMaintenanceResponse) {
    m := &GhesSetMaintenanceResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesSetMaintenanceResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesSetMaintenanceResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesSetMaintenanceResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesSetMaintenanceResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetError gets the error property value. The error property
// returns a *string when successful
func (m *GhesSetMaintenanceResponse) GetError()(*string) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesSetMaintenanceResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetHostname gets the hostname property value. The hostname property
// returns a *string when successful
func (m *GhesSetMaintenanceResponse) GetHostname()(*string) {
    return m.hostname
}
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *GhesSetMaintenanceResponse) GetMessage()(*string) {
    return m.message
}
// GetUuid gets the uuid property value. The uuid property
// returns a *UUID when successful
func (m *GhesSetMaintenanceResponse) GetUuid()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.uuid
}
// Serialize serializes information the current object
func (m *GhesSetMaintenanceResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *GhesSetMaintenanceResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetError sets the error property value. The error property
func (m *GhesSetMaintenanceResponse) SetError(value *string)() {
    m.error = value
}
// SetHostname sets the hostname property value. The hostname property
func (m *GhesSetMaintenanceResponse) SetHostname(value *string)() {
    m.hostname = value
}
// SetMessage sets the message property value. The message property
func (m *GhesSetMaintenanceResponse) SetMessage(value *string)() {
    m.message = value
}
// SetUuid sets the uuid property value. The uuid property
func (m *GhesSetMaintenanceResponse) SetUuid(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.uuid = value
}
type GhesSetMaintenanceResponseable interface {
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
