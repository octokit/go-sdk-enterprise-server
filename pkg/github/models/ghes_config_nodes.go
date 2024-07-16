package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesConfigNodes struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The nodes property
    nodes []GhesConfigNodes_nodesable
    // The topology property
    topology *GhesClusterTopology
}
// NewGhesConfigNodes instantiates a new GhesConfigNodes and sets the default values.
func NewGhesConfigNodes()(*GhesConfigNodes) {
    m := &GhesConfigNodes{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesConfigNodesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesConfigNodesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesConfigNodes(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesConfigNodes) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesConfigNodes) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["nodes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGhesConfigNodes_nodesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GhesConfigNodes_nodesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GhesConfigNodes_nodesable)
                }
            }
            m.SetNodes(res)
        }
        return nil
    }
    res["topology"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGhesClusterTopology)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopology(val.(*GhesClusterTopology))
        }
        return nil
    }
    return res
}
// GetNodes gets the nodes property value. The nodes property
// returns a []GhesConfigNodes_nodesable when successful
func (m *GhesConfigNodes) GetNodes()([]GhesConfigNodes_nodesable) {
    return m.nodes
}
// GetTopology gets the topology property value. The topology property
// returns a *GhesClusterTopology when successful
func (m *GhesConfigNodes) GetTopology()(*GhesClusterTopology) {
    return m.topology
}
// Serialize serializes information the current object
func (m *GhesConfigNodes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetTopology() != nil {
        cast := (*m.GetTopology()).String()
        err := writer.WriteStringValue("topology", &cast)
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
func (m *GhesConfigNodes) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNodes sets the nodes property value. The nodes property
func (m *GhesConfigNodes) SetNodes(value []GhesConfigNodes_nodesable)() {
    m.nodes = value
}
// SetTopology sets the topology property value. The topology property
func (m *GhesConfigNodes) SetTopology(value *GhesClusterTopology)() {
    m.topology = value
}
type GhesConfigNodesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNodes()([]GhesConfigNodes_nodesable)
    GetTopology()(*GhesClusterTopology)
    SetNodes(value []GhesConfigNodes_nodesable)()
    SetTopology(value *GhesClusterTopology)()
}
