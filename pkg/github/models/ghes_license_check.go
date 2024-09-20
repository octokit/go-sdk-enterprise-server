package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesLicenseCheck struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The license status of the currently active enterprise license.
    status *GhesLicenseCheck_status
}
// NewGhesLicenseCheck instantiates a new GhesLicenseCheck and sets the default values.
func NewGhesLicenseCheck()(*GhesLicenseCheck) {
    m := &GhesLicenseCheck{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesLicenseCheckFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesLicenseCheckFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesLicenseCheck(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesLicenseCheck) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesLicenseCheck) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGhesLicenseCheck_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*GhesLicenseCheck_status))
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. The license status of the currently active enterprise license.
// returns a *GhesLicenseCheck_status when successful
func (m *GhesLicenseCheck) GetStatus()(*GhesLicenseCheck_status) {
    return m.status
}
// Serialize serializes information the current object
func (m *GhesLicenseCheck) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *GhesLicenseCheck) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetStatus sets the status property value. The license status of the currently active enterprise license.
func (m *GhesLicenseCheck) SetStatus(value *GhesLicenseCheck_status)() {
    m.status = value
}
type GhesLicenseCheckable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetStatus()(*GhesLicenseCheck_status)
    SetStatus(value *GhesLicenseCheck_status)()
}
