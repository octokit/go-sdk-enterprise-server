package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EnterpriseSettings_enterprise_ntp struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The primary_server property
    primary_server *string
    // The secondary_server property
    secondary_server *string
}
// NewEnterpriseSettings_enterprise_ntp instantiates a new EnterpriseSettings_enterprise_ntp and sets the default values.
func NewEnterpriseSettings_enterprise_ntp()(*EnterpriseSettings_enterprise_ntp) {
    m := &EnterpriseSettings_enterprise_ntp{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnterpriseSettings_enterprise_ntpFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnterpriseSettings_enterprise_ntpFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnterpriseSettings_enterprise_ntp(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EnterpriseSettings_enterprise_ntp) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnterpriseSettings_enterprise_ntp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["primary_server"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryServer(val)
        }
        return nil
    }
    res["secondary_server"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecondaryServer(val)
        }
        return nil
    }
    return res
}
// GetPrimaryServer gets the primary_server property value. The primary_server property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_ntp) GetPrimaryServer()(*string) {
    return m.primary_server
}
// GetSecondaryServer gets the secondary_server property value. The secondary_server property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_ntp) GetSecondaryServer()(*string) {
    return m.secondary_server
}
// Serialize serializes information the current object
func (m *EnterpriseSettings_enterprise_ntp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("primary_server", m.GetPrimaryServer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secondary_server", m.GetSecondaryServer())
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
func (m *EnterpriseSettings_enterprise_ntp) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPrimaryServer sets the primary_server property value. The primary_server property
func (m *EnterpriseSettings_enterprise_ntp) SetPrimaryServer(value *string)() {
    m.primary_server = value
}
// SetSecondaryServer sets the secondary_server property value. The secondary_server property
func (m *EnterpriseSettings_enterprise_ntp) SetSecondaryServer(value *string)() {
    m.secondary_server = value
}
type EnterpriseSettings_enterprise_ntpable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPrimaryServer()(*string)
    GetSecondaryServer()(*string)
    SetPrimaryServer(value *string)()
    SetSecondaryServer(value *string)()
}
