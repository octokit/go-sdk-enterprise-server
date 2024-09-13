package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesClusterStatus_nodes_services struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The details property
    details *string
    // The name property
    name *string
    // The status property
    status *GhesClusterStatusIndicator
}
// NewGhesClusterStatus_nodes_services instantiates a new GhesClusterStatus_nodes_services and sets the default values.
func NewGhesClusterStatus_nodes_services()(*GhesClusterStatus_nodes_services) {
    m := &GhesClusterStatus_nodes_services{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesClusterStatus_nodes_servicesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesClusterStatus_nodes_servicesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesClusterStatus_nodes_services(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesClusterStatus_nodes_services) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDetails gets the details property value. The details property
// returns a *string when successful
func (m *GhesClusterStatus_nodes_services) GetDetails()(*string) {
    return m.details
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesClusterStatus_nodes_services) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *GhesClusterStatus_nodes_services) GetName()(*string) {
    return m.name
}
// GetStatus gets the status property value. The status property
// returns a *GhesClusterStatusIndicator when successful
func (m *GhesClusterStatus_nodes_services) GetStatus()(*GhesClusterStatusIndicator) {
    return m.status
}
// Serialize serializes information the current object
func (m *GhesClusterStatus_nodes_services) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("details", m.GetDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *GhesClusterStatus_nodes_services) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDetails sets the details property value. The details property
func (m *GhesClusterStatus_nodes_services) SetDetails(value *string)() {
    m.details = value
}
// SetName sets the name property value. The name property
func (m *GhesClusterStatus_nodes_services) SetName(value *string)() {
    m.name = value
}
// SetStatus sets the status property value. The status property
func (m *GhesClusterStatus_nodes_services) SetStatus(value *GhesClusterStatusIndicator)() {
    m.status = value
}
type GhesClusterStatus_nodes_servicesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDetails()(*string)
    GetName()(*string)
    GetStatus()(*GhesClusterStatusIndicator)
    SetDetails(value *string)()
    SetName(value *string)()
    SetStatus(value *GhesClusterStatusIndicator)()
}
