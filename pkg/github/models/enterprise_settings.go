package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EnterpriseSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The enterprise property
    enterprise EnterpriseSettings_enterpriseable
    // The run_list property
    run_list []string
}
// NewEnterpriseSettings instantiates a new EnterpriseSettings and sets the default values.
func NewEnterpriseSettings()(*EnterpriseSettings) {
    m := &EnterpriseSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnterpriseSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnterpriseSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnterpriseSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EnterpriseSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnterprise gets the enterprise property value. The enterprise property
// returns a EnterpriseSettings_enterpriseable when successful
func (m *EnterpriseSettings) GetEnterprise()(EnterpriseSettings_enterpriseable) {
    return m.enterprise
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnterpriseSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enterprise"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnterpriseSettings_enterpriseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterprise(val.(EnterpriseSettings_enterpriseable))
        }
        return nil
    }
    res["run_list"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetRunList(res)
        }
        return nil
    }
    return res
}
// GetRunList gets the run_list property value. The run_list property
// returns a []string when successful
func (m *EnterpriseSettings) GetRunList()([]string) {
    return m.run_list
}
// Serialize serializes information the current object
func (m *EnterpriseSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("enterprise", m.GetEnterprise())
        if err != nil {
            return err
        }
    }
    if m.GetRunList() != nil {
        err := writer.WriteCollectionOfStringValues("run_list", m.GetRunList())
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
func (m *EnterpriseSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnterprise sets the enterprise property value. The enterprise property
func (m *EnterpriseSettings) SetEnterprise(value EnterpriseSettings_enterpriseable)() {
    m.enterprise = value
}
// SetRunList sets the run_list property value. The run_list property
func (m *EnterpriseSettings) SetRunList(value []string)() {
    m.run_list = value
}
type EnterpriseSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnterprise()(EnterpriseSettings_enterpriseable)
    GetRunList()([]string)
    SetEnterprise(value EnterpriseSettings_enterpriseable)()
    SetRunList(value []string)()
}
