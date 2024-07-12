package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesReplicationStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The nodes property
    nodes []GhesReplicationStatus_nodesable
    // The status property
    status *GhesReplicationStatusIndicator
}
// NewGhesReplicationStatus instantiates a new GhesReplicationStatus and sets the default values.
func NewGhesReplicationStatus()(*GhesReplicationStatus) {
    m := &GhesReplicationStatus{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesReplicationStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesReplicationStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesReplicationStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesReplicationStatus) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesReplicationStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["nodes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGhesReplicationStatus_nodesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GhesReplicationStatus_nodesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GhesReplicationStatus_nodesable)
                }
            }
            m.SetNodes(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGhesReplicationStatusIndicator)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*GhesReplicationStatusIndicator))
        }
        return nil
    }
    return res
}
// GetNodes gets the nodes property value. The nodes property
// returns a []GhesReplicationStatus_nodesable when successful
func (m *GhesReplicationStatus) GetNodes()([]GhesReplicationStatus_nodesable) {
    return m.nodes
}
// GetStatus gets the status property value. The status property
// returns a *GhesReplicationStatusIndicator when successful
func (m *GhesReplicationStatus) GetStatus()(*GhesReplicationStatusIndicator) {
    return m.status
}
// Serialize serializes information the current object
func (m *GhesReplicationStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *GhesReplicationStatus) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNodes sets the nodes property value. The nodes property
func (m *GhesReplicationStatus) SetNodes(value []GhesReplicationStatus_nodesable)() {
    m.nodes = value
}
// SetStatus sets the status property value. The status property
func (m *GhesReplicationStatus) SetStatus(value *GhesReplicationStatusIndicator)() {
    m.status = value
}
type GhesReplicationStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNodes()([]GhesReplicationStatus_nodesable)
    GetStatus()(*GhesReplicationStatusIndicator)
    SetNodes(value []GhesReplicationStatus_nodesable)()
    SetStatus(value *GhesReplicationStatusIndicator)()
}
