package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesClusterStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The nodes property
    nodes []GhesClusterStatus_nodesable
    // The status property
    status *GhesClusterStatusIndicator
}
// NewGhesClusterStatus instantiates a new GhesClusterStatus and sets the default values.
func NewGhesClusterStatus()(*GhesClusterStatus) {
    m := &GhesClusterStatus{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesClusterStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesClusterStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesClusterStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesClusterStatus) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesClusterStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["nodes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGhesClusterStatus_nodesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GhesClusterStatus_nodesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GhesClusterStatus_nodesable)
                }
            }
            m.SetNodes(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGhesClusterStatusIndicator)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*GhesClusterStatusIndicator))
        }
        return nil
    }
    return res
}
// GetNodes gets the nodes property value. The nodes property
// returns a []GhesClusterStatus_nodesable when successful
func (m *GhesClusterStatus) GetNodes()([]GhesClusterStatus_nodesable) {
    return m.nodes
}
// GetStatus gets the status property value. The status property
// returns a *GhesClusterStatusIndicator when successful
func (m *GhesClusterStatus) GetStatus()(*GhesClusterStatusIndicator) {
    return m.status
}
// Serialize serializes information the current object
func (m *GhesClusterStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *GhesClusterStatus) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNodes sets the nodes property value. The nodes property
func (m *GhesClusterStatus) SetNodes(value []GhesClusterStatus_nodesable)() {
    m.nodes = value
}
// SetStatus sets the status property value. The status property
func (m *GhesClusterStatus) SetStatus(value *GhesClusterStatusIndicator)() {
    m.status = value
}
type GhesClusterStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNodes()([]GhesClusterStatus_nodesable)
    GetStatus()(*GhesClusterStatusIndicator)
    SetNodes(value []GhesClusterStatus_nodesable)()
    SetStatus(value *GhesClusterStatusIndicator)()
}
