package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EnterpriseSettings_enterprise_avatar struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The enabled property
    enabled *bool
    // The uri property
    uri *string
}
// NewEnterpriseSettings_enterprise_avatar instantiates a new EnterpriseSettings_enterprise_avatar and sets the default values.
func NewEnterpriseSettings_enterprise_avatar()(*EnterpriseSettings_enterprise_avatar) {
    m := &EnterpriseSettings_enterprise_avatar{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnterpriseSettings_enterprise_avatarFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnterpriseSettings_enterprise_avatarFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnterpriseSettings_enterprise_avatar(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EnterpriseSettings_enterprise_avatar) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnabled gets the enabled property value. The enabled property
// returns a *bool when successful
func (m *EnterpriseSettings_enterprise_avatar) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnterpriseSettings_enterprise_avatar) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["uri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUri(val)
        }
        return nil
    }
    return res
}
// GetUri gets the uri property value. The uri property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_avatar) GetUri()(*string) {
    return m.uri
}
// Serialize serializes information the current object
func (m *EnterpriseSettings_enterprise_avatar) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("uri", m.GetUri())
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
func (m *EnterpriseSettings_enterprise_avatar) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnabled sets the enabled property value. The enabled property
func (m *EnterpriseSettings_enterprise_avatar) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetUri sets the uri property value. The uri property
func (m *EnterpriseSettings_enterprise_avatar) SetUri(value *string)() {
    m.uri = value
}
type EnterpriseSettings_enterprise_avatarable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnabled()(*bool)
    GetUri()(*string)
    SetEnabled(value *bool)()
    SetUri(value *string)()
}
