package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConfigurationStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The progress property
    progress []ConfigurationStatus_progressable
    // The status property
    status *string
}
// NewConfigurationStatus instantiates a new ConfigurationStatus and sets the default values.
func NewConfigurationStatus()(*ConfigurationStatus) {
    m := &ConfigurationStatus{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConfigurationStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConfigurationStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConfigurationStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ConfigurationStatus) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConfigurationStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["progress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConfigurationStatus_progressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConfigurationStatus_progressable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ConfigurationStatus_progressable)
                }
            }
            m.SetProgress(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
// GetProgress gets the progress property value. The progress property
// returns a []ConfigurationStatus_progressable when successful
func (m *ConfigurationStatus) GetProgress()([]ConfigurationStatus_progressable) {
    return m.progress
}
// GetStatus gets the status property value. The status property
// returns a *string when successful
func (m *ConfigurationStatus) GetStatus()(*string) {
    return m.status
}
// Serialize serializes information the current object
func (m *ConfigurationStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetProgress() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProgress()))
        for i, v := range m.GetProgress() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("progress", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("status", m.GetStatus())
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
func (m *ConfigurationStatus) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetProgress sets the progress property value. The progress property
func (m *ConfigurationStatus) SetProgress(value []ConfigurationStatus_progressable)() {
    m.progress = value
}
// SetStatus sets the status property value. The status property
func (m *ConfigurationStatus) SetStatus(value *string)() {
    m.status = value
}
type ConfigurationStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetProgress()([]ConfigurationStatus_progressable)
    GetStatus()(*string)
    SetProgress(value []ConfigurationStatus_progressable)()
    SetStatus(value *string)()
}
