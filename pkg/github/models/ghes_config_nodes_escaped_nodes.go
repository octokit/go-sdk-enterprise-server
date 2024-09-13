package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesConfigNodes_nodes struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The cluster_roles property
    cluster_roles []GhesClusterRoles
    // The hostname property
    hostname *string
    // The replica property
    replica *bool
    // The uuid property
    uuid *string
}
// NewGhesConfigNodes_nodes instantiates a new GhesConfigNodes_nodes and sets the default values.
func NewGhesConfigNodes_nodes()(*GhesConfigNodes_nodes) {
    m := &GhesConfigNodes_nodes{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesConfigNodes_nodesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesConfigNodes_nodesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesConfigNodes_nodes(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesConfigNodes_nodes) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClusterRoles gets the cluster_roles property value. The cluster_roles property
// returns a []GhesClusterRoles when successful
func (m *GhesConfigNodes_nodes) GetClusterRoles()([]GhesClusterRoles) {
    return m.cluster_roles
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesConfigNodes_nodes) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cluster_roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseGhesClusterRoles)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GhesClusterRoles, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*GhesClusterRoles))
                }
            }
            m.SetClusterRoles(res)
        }
        return nil
    }
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
    res["replica"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReplica(val)
        }
        return nil
    }
    res["uuid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUuid(val)
        }
        return nil
    }
    return res
}
// GetHostname gets the hostname property value. The hostname property
// returns a *string when successful
func (m *GhesConfigNodes_nodes) GetHostname()(*string) {
    return m.hostname
}
// GetReplica gets the replica property value. The replica property
// returns a *bool when successful
func (m *GhesConfigNodes_nodes) GetReplica()(*bool) {
    return m.replica
}
// GetUuid gets the uuid property value. The uuid property
// returns a *string when successful
func (m *GhesConfigNodes_nodes) GetUuid()(*string) {
    return m.uuid
}
// Serialize serializes information the current object
func (m *GhesConfigNodes_nodes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetClusterRoles() != nil {
        err := writer.WriteCollectionOfStringValues("cluster_roles", SerializeGhesClusterRoles(m.GetClusterRoles()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hostname", m.GetHostname())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("replica", m.GetReplica())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("uuid", m.GetUuid())
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
func (m *GhesConfigNodes_nodes) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClusterRoles sets the cluster_roles property value. The cluster_roles property
func (m *GhesConfigNodes_nodes) SetClusterRoles(value []GhesClusterRoles)() {
    m.cluster_roles = value
}
// SetHostname sets the hostname property value. The hostname property
func (m *GhesConfigNodes_nodes) SetHostname(value *string)() {
    m.hostname = value
}
// SetReplica sets the replica property value. The replica property
func (m *GhesConfigNodes_nodes) SetReplica(value *bool)() {
    m.replica = value
}
// SetUuid sets the uuid property value. The uuid property
func (m *GhesConfigNodes_nodes) SetUuid(value *string)() {
    m.uuid = value
}
type GhesConfigNodes_nodesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClusterRoles()([]GhesClusterRoles)
    GetHostname()(*string)
    GetReplica()(*bool)
    GetUuid()(*string)
    SetClusterRoles(value []GhesClusterRoles)()
    SetHostname(value *string)()
    SetReplica(value *bool)()
    SetUuid(value *string)()
}
