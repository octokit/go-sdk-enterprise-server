package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesChecksSystemRequirements_nodes_roles_status struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The role property
    role *string
    // The status property
    status *GhesChecksResultIndicator
}
// NewGhesChecksSystemRequirements_nodes_roles_status instantiates a new GhesChecksSystemRequirements_nodes_roles_status and sets the default values.
func NewGhesChecksSystemRequirements_nodes_roles_status()(*GhesChecksSystemRequirements_nodes_roles_status) {
    m := &GhesChecksSystemRequirements_nodes_roles_status{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesChecksSystemRequirements_nodes_roles_statusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesChecksSystemRequirements_nodes_roles_statusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesChecksSystemRequirements_nodes_roles_status(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesChecksSystemRequirements_nodes_roles_status) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesChecksSystemRequirements_nodes_roles_status) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGhesChecksResultIndicator)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*GhesChecksResultIndicator))
        }
        return nil
    }
    return res
}
// GetRole gets the role property value. The role property
// returns a *string when successful
func (m *GhesChecksSystemRequirements_nodes_roles_status) GetRole()(*string) {
    return m.role
}
// GetStatus gets the status property value. The status property
// returns a *GhesChecksResultIndicator when successful
func (m *GhesChecksSystemRequirements_nodes_roles_status) GetStatus()(*GhesChecksResultIndicator) {
    return m.status
}
// Serialize serializes information the current object
func (m *GhesChecksSystemRequirements_nodes_roles_status) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("role", m.GetRole())
        if err != nil {
            return err
        }
    }
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
func (m *GhesChecksSystemRequirements_nodes_roles_status) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRole sets the role property value. The role property
func (m *GhesChecksSystemRequirements_nodes_roles_status) SetRole(value *string)() {
    m.role = value
}
// SetStatus sets the status property value. The status property
func (m *GhesChecksSystemRequirements_nodes_roles_status) SetStatus(value *GhesChecksResultIndicator)() {
    m.status = value
}
type GhesChecksSystemRequirements_nodes_roles_statusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRole()(*string)
    GetStatus()(*GhesChecksResultIndicator)
    SetRole(value *string)()
    SetStatus(value *GhesChecksResultIndicator)()
}
