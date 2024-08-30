package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesClusterStatus_nodes struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The hostname property
    hostname *string
    // The services property
    services []GhesClusterStatus_nodes_servicesable
    // The status property
    status *GhesClusterStatusIndicator
}
// NewGhesClusterStatus_nodes instantiates a new GhesClusterStatus_nodes and sets the default values.
func NewGhesClusterStatus_nodes()(*GhesClusterStatus_nodes) {
    m := &GhesClusterStatus_nodes{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesClusterStatus_nodesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesClusterStatus_nodesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesClusterStatus_nodes(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesClusterStatus_nodes) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesClusterStatus_nodes) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["services"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGhesClusterStatus_nodes_servicesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GhesClusterStatus_nodes_servicesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GhesClusterStatus_nodes_servicesable)
                }
            }
            m.SetServices(res)
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
// GetHostname gets the hostname property value. The hostname property
// returns a *string when successful
func (m *GhesClusterStatus_nodes) GetHostname()(*string) {
    return m.hostname
}
// GetServices gets the services property value. The services property
// returns a []GhesClusterStatus_nodes_servicesable when successful
func (m *GhesClusterStatus_nodes) GetServices()([]GhesClusterStatus_nodes_servicesable) {
    return m.services
}
// GetStatus gets the status property value. The status property
// returns a *GhesClusterStatusIndicator when successful
func (m *GhesClusterStatus_nodes) GetStatus()(*GhesClusterStatusIndicator) {
    return m.status
}
// Serialize serializes information the current object
func (m *GhesClusterStatus_nodes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("hostname", m.GetHostname())
        if err != nil {
            return err
        }
    }
    if m.GetServices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServices()))
        for i, v := range m.GetServices() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("services", cast)
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
func (m *GhesClusterStatus_nodes) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHostname sets the hostname property value. The hostname property
func (m *GhesClusterStatus_nodes) SetHostname(value *string)() {
    m.hostname = value
}
// SetServices sets the services property value. The services property
func (m *GhesClusterStatus_nodes) SetServices(value []GhesClusterStatus_nodes_servicesable)() {
    m.services = value
}
// SetStatus sets the status property value. The status property
func (m *GhesClusterStatus_nodes) SetStatus(value *GhesClusterStatusIndicator)() {
    m.status = value
}
type GhesClusterStatus_nodesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHostname()(*string)
    GetServices()([]GhesClusterStatus_nodes_servicesable)
    GetStatus()(*GhesClusterStatusIndicator)
    SetHostname(value *string)()
    SetServices(value []GhesClusterStatus_nodes_servicesable)()
    SetStatus(value *GhesClusterStatusIndicator)()
}
