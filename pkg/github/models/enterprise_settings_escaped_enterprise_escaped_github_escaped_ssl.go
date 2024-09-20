package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EnterpriseSettings_enterprise_github_ssl struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The cert property
    cert *string
    // The enabled property
    enabled *bool
    // The key property
    key *string
}
// NewEnterpriseSettings_enterprise_github_ssl instantiates a new EnterpriseSettings_enterprise_github_ssl and sets the default values.
func NewEnterpriseSettings_enterprise_github_ssl()(*EnterpriseSettings_enterprise_github_ssl) {
    m := &EnterpriseSettings_enterprise_github_ssl{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnterpriseSettings_enterprise_github_sslFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnterpriseSettings_enterprise_github_sslFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnterpriseSettings_enterprise_github_ssl(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EnterpriseSettings_enterprise_github_ssl) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCert gets the cert property value. The cert property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_github_ssl) GetCert()(*string) {
    return m.cert
}
// GetEnabled gets the enabled property value. The enabled property
// returns a *bool when successful
func (m *EnterpriseSettings_enterprise_github_ssl) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnterpriseSettings_enterprise_github_ssl) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cert"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCert(val)
        }
        return nil
    }
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
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. The key property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_github_ssl) GetKey()(*string) {
    return m.key
}
// Serialize serializes information the current object
func (m *EnterpriseSettings_enterprise_github_ssl) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("cert", m.GetCert())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("key", m.GetKey())
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
func (m *EnterpriseSettings_enterprise_github_ssl) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCert sets the cert property value. The cert property
func (m *EnterpriseSettings_enterprise_github_ssl) SetCert(value *string)() {
    m.cert = value
}
// SetEnabled sets the enabled property value. The enabled property
func (m *EnterpriseSettings_enterprise_github_ssl) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetKey sets the key property value. The key property
func (m *EnterpriseSettings_enterprise_github_ssl) SetKey(value *string)() {
    m.key = value
}
type EnterpriseSettings_enterprise_github_sslable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCert()(*string)
    GetEnabled()(*bool)
    GetKey()(*string)
    SetCert(value *string)()
    SetEnabled(value *bool)()
    SetKey(value *string)()
}
