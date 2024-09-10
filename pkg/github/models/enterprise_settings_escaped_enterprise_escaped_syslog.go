package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EnterpriseSettings_enterprise_syslog struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The enabled property
    enabled *bool
    // The protocol_name property
    protocol_name *string
    // The server property
    server *string
}
// NewEnterpriseSettings_enterprise_syslog instantiates a new EnterpriseSettings_enterprise_syslog and sets the default values.
func NewEnterpriseSettings_enterprise_syslog()(*EnterpriseSettings_enterprise_syslog) {
    m := &EnterpriseSettings_enterprise_syslog{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnterpriseSettings_enterprise_syslogFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnterpriseSettings_enterprise_syslogFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnterpriseSettings_enterprise_syslog(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EnterpriseSettings_enterprise_syslog) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnabled gets the enabled property value. The enabled property
// returns a *bool when successful
func (m *EnterpriseSettings_enterprise_syslog) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnterpriseSettings_enterprise_syslog) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["protocol_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtocolName(val)
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
    return res
}
// GetProtocolName gets the protocol_name property value. The protocol_name property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_syslog) GetProtocolName()(*string) {
    return m.protocol_name
}
// GetServer gets the server property value. The server property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_syslog) GetServer()(*string) {
    return m.server
}
// Serialize serializes information the current object
func (m *EnterpriseSettings_enterprise_syslog) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("protocol_name", m.GetProtocolName())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EnterpriseSettings_enterprise_syslog) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnabled sets the enabled property value. The enabled property
func (m *EnterpriseSettings_enterprise_syslog) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetProtocolName sets the protocol_name property value. The protocol_name property
func (m *EnterpriseSettings_enterprise_syslog) SetProtocolName(value *string)() {
    m.protocol_name = value
}
// SetServer sets the server property value. The server property
func (m *EnterpriseSettings_enterprise_syslog) SetServer(value *string)() {
    m.server = value
}
type EnterpriseSettings_enterprise_syslogable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnabled()(*bool)
    GetProtocolName()(*string)
    GetServer()(*string)
    SetEnabled(value *bool)()
    SetProtocolName(value *string)()
    SetServer(value *string)()
}
