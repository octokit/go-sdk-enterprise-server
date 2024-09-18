package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesChecksSystemRequirements struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The nodes property
    nodes []GhesChecksSystemRequirements_nodesable
    // The status property
    status *GhesChecksResultIndicator
}
// NewGhesChecksSystemRequirements instantiates a new GhesChecksSystemRequirements and sets the default values.
func NewGhesChecksSystemRequirements()(*GhesChecksSystemRequirements) {
    m := &GhesChecksSystemRequirements{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesChecksSystemRequirementsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesChecksSystemRequirementsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesChecksSystemRequirements(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesChecksSystemRequirements) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesChecksSystemRequirements) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["nodes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGhesChecksSystemRequirements_nodesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GhesChecksSystemRequirements_nodesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GhesChecksSystemRequirements_nodesable)
                }
            }
            m.SetNodes(res)
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
// GetNodes gets the nodes property value. The nodes property
// returns a []GhesChecksSystemRequirements_nodesable when successful
func (m *GhesChecksSystemRequirements) GetNodes()([]GhesChecksSystemRequirements_nodesable) {
    return m.nodes
}
// GetStatus gets the status property value. The status property
// returns a *GhesChecksResultIndicator when successful
func (m *GhesChecksSystemRequirements) GetStatus()(*GhesChecksResultIndicator) {
    return m.status
}
// Serialize serializes information the current object
func (m *GhesChecksSystemRequirements) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetNodes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNodes()))
        for i, v := range m.GetNodes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("nodes", cast)
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
func (m *GhesChecksSystemRequirements) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNodes sets the nodes property value. The nodes property
func (m *GhesChecksSystemRequirements) SetNodes(value []GhesChecksSystemRequirements_nodesable)() {
    m.nodes = value
}
// SetStatus sets the status property value. The status property
func (m *GhesChecksSystemRequirements) SetStatus(value *GhesChecksResultIndicator)() {
    m.status = value
}
type GhesChecksSystemRequirementsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNodes()([]GhesChecksSystemRequirements_nodesable)
    GetStatus()(*GhesChecksResultIndicator)
    SetNodes(value []GhesChecksSystemRequirements_nodesable)()
    SetStatus(value *GhesChecksResultIndicator)()
}
