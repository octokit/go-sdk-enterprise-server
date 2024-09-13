package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesChecksSystemRequirements_nodes struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The hostname property
    hostname *string
    // The roles_status property
    roles_status []GhesChecksSystemRequirements_nodes_roles_statusable
    // The status property
    status *GhesChecksResultIndicator
}
// NewGhesChecksSystemRequirements_nodes instantiates a new GhesChecksSystemRequirements_nodes and sets the default values.
func NewGhesChecksSystemRequirements_nodes()(*GhesChecksSystemRequirements_nodes) {
    m := &GhesChecksSystemRequirements_nodes{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesChecksSystemRequirements_nodesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesChecksSystemRequirements_nodesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesChecksSystemRequirements_nodes(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesChecksSystemRequirements_nodes) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesChecksSystemRequirements_nodes) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["roles_status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGhesChecksSystemRequirements_nodes_roles_statusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GhesChecksSystemRequirements_nodes_roles_statusable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GhesChecksSystemRequirements_nodes_roles_statusable)
                }
            }
            m.SetRolesStatus(res)
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
// GetHostname gets the hostname property value. The hostname property
// returns a *string when successful
func (m *GhesChecksSystemRequirements_nodes) GetHostname()(*string) {
    return m.hostname
}
// GetRolesStatus gets the roles_status property value. The roles_status property
// returns a []GhesChecksSystemRequirements_nodes_roles_statusable when successful
func (m *GhesChecksSystemRequirements_nodes) GetRolesStatus()([]GhesChecksSystemRequirements_nodes_roles_statusable) {
    return m.roles_status
}
// GetStatus gets the status property value. The status property
// returns a *GhesChecksResultIndicator when successful
func (m *GhesChecksSystemRequirements_nodes) GetStatus()(*GhesChecksResultIndicator) {
    return m.status
}
// Serialize serializes information the current object
func (m *GhesChecksSystemRequirements_nodes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("hostname", m.GetHostname())
        if err != nil {
            return err
        }
    }
    if m.GetRolesStatus() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRolesStatus()))
        for i, v := range m.GetRolesStatus() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("roles_status", cast)
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
func (m *GhesChecksSystemRequirements_nodes) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHostname sets the hostname property value. The hostname property
func (m *GhesChecksSystemRequirements_nodes) SetHostname(value *string)() {
    m.hostname = value
}
// SetRolesStatus sets the roles_status property value. The roles_status property
func (m *GhesChecksSystemRequirements_nodes) SetRolesStatus(value []GhesChecksSystemRequirements_nodes_roles_statusable)() {
    m.roles_status = value
}
// SetStatus sets the status property value. The status property
func (m *GhesChecksSystemRequirements_nodes) SetStatus(value *GhesChecksResultIndicator)() {
    m.status = value
}
type GhesChecksSystemRequirements_nodesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHostname()(*string)
    GetRolesStatus()([]GhesChecksSystemRequirements_nodes_roles_statusable)
    GetStatus()(*GhesChecksResultIndicator)
    SetHostname(value *string)()
    SetRolesStatus(value []GhesChecksSystemRequirements_nodes_roles_statusable)()
    SetStatus(value *GhesChecksResultIndicator)()
}
