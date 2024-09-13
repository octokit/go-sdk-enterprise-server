package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesGetSettings_collectd struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The enabled property
    enabled *bool
    // The encryption property
    encryption *string
    // The password property
    password *string
    // The port property
    port *int32
    // The server property
    server *string
    // The username property
    username *string
}
// NewGhesGetSettings_collectd instantiates a new GhesGetSettings_collectd and sets the default values.
func NewGhesGetSettings_collectd()(*GhesGetSettings_collectd) {
    m := &GhesGetSettings_collectd{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesGetSettings_collectdFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesGetSettings_collectdFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesGetSettings_collectd(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesGetSettings_collectd) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnabled gets the enabled property value. The enabled property
// returns a *bool when successful
func (m *GhesGetSettings_collectd) GetEnabled()(*bool) {
    return m.enabled
}
// GetEncryption gets the encryption property value. The encryption property
// returns a *string when successful
func (m *GhesGetSettings_collectd) GetEncryption()(*string) {
    return m.encryption
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesGetSettings_collectd) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["encryption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryption(val)
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
    res["port"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPort(val)
        }
        return nil
    }
    res["server"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServer(val)
        }
        return nil
    }
    res["username"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsername(val)
        }
        return nil
    }
    return res
}
// GetPassword gets the password property value. The password property
// returns a *string when successful
func (m *GhesGetSettings_collectd) GetPassword()(*string) {
    return m.password
}
// GetPort gets the port property value. The port property
// returns a *int32 when successful
func (m *GhesGetSettings_collectd) GetPort()(*int32) {
    return m.port
}
// GetServer gets the server property value. The server property
// returns a *string when successful
func (m *GhesGetSettings_collectd) GetServer()(*string) {
    return m.server
}
// GetUsername gets the username property value. The username property
// returns a *string when successful
func (m *GhesGetSettings_collectd) GetUsername()(*string) {
    return m.username
}
// Serialize serializes information the current object
func (m *GhesGetSettings_collectd) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("encryption", m.GetEncryption())
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
        err := writer.WriteInt32Value("port", m.GetPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("server", m.GetServer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("username", m.GetUsername())
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
func (m *GhesGetSettings_collectd) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnabled sets the enabled property value. The enabled property
func (m *GhesGetSettings_collectd) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetEncryption sets the encryption property value. The encryption property
func (m *GhesGetSettings_collectd) SetEncryption(value *string)() {
    m.encryption = value
}
// SetPassword sets the password property value. The password property
func (m *GhesGetSettings_collectd) SetPassword(value *string)() {
    m.password = value
}
// SetPort sets the port property value. The port property
func (m *GhesGetSettings_collectd) SetPort(value *int32)() {
    m.port = value
}
// SetServer sets the server property value. The server property
func (m *GhesGetSettings_collectd) SetServer(value *string)() {
    m.server = value
}
// SetUsername sets the username property value. The username property
func (m *GhesGetSettings_collectd) SetUsername(value *string)() {
    m.username = value
}
type GhesGetSettings_collectdable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnabled()(*bool)
    GetEncryption()(*string)
    GetPassword()(*string)
    GetPort()(*int32)
    GetServer()(*string)
    GetUsername()(*string)
    SetEnabled(value *bool)()
    SetEncryption(value *string)()
    SetPassword(value *string)()
    SetPort(value *int32)()
    SetServer(value *string)()
    SetUsername(value *string)()
}
