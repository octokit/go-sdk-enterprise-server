package manage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1ConfigApplyGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The nodes property
    nodes []V1ConfigApplyGetResponse_nodesable
    // Whether the ghe-config-apply run is still running in the environment
    running *bool
    // Whether the ghe-config-apply run was successful in the environment
    successful *bool
}
// NewV1ConfigApplyGetResponse instantiates a new V1ConfigApplyGetResponse and sets the default values.
func NewV1ConfigApplyGetResponse()(*V1ConfigApplyGetResponse) {
    m := &V1ConfigApplyGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1ConfigApplyGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1ConfigApplyGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1ConfigApplyGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1ConfigApplyGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1ConfigApplyGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["nodes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1ConfigApplyGetResponse_nodesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1ConfigApplyGetResponse_nodesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1ConfigApplyGetResponse_nodesable)
                }
            }
            m.SetNodes(res)
        }
        return nil
    }
    res["running"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunning(val)
        }
        return nil
    }
    res["successful"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessful(val)
        }
        return nil
    }
    return res
}
// GetNodes gets the nodes property value. The nodes property
// returns a []V1ConfigApplyGetResponse_nodesable when successful
func (m *V1ConfigApplyGetResponse) GetNodes()([]V1ConfigApplyGetResponse_nodesable) {
    return m.nodes
}
// GetRunning gets the running property value. Whether the ghe-config-apply run is still running in the environment
// returns a *bool when successful
func (m *V1ConfigApplyGetResponse) GetRunning()(*bool) {
    return m.running
}
// GetSuccessful gets the successful property value. Whether the ghe-config-apply run was successful in the environment
// returns a *bool when successful
func (m *V1ConfigApplyGetResponse) GetSuccessful()(*bool) {
    return m.successful
}
// Serialize serializes information the current object
func (m *V1ConfigApplyGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteBoolValue("running", m.GetRunning())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("successful", m.GetSuccessful())
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
func (m *V1ConfigApplyGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNodes sets the nodes property value. The nodes property
func (m *V1ConfigApplyGetResponse) SetNodes(value []V1ConfigApplyGetResponse_nodesable)() {
    m.nodes = value
}
// SetRunning sets the running property value. Whether the ghe-config-apply run is still running in the environment
func (m *V1ConfigApplyGetResponse) SetRunning(value *bool)() {
    m.running = value
}
// SetSuccessful sets the successful property value. Whether the ghe-config-apply run was successful in the environment
func (m *V1ConfigApplyGetResponse) SetSuccessful(value *bool)() {
    m.successful = value
}
type V1ConfigApplyGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNodes()([]V1ConfigApplyGetResponse_nodesable)
    GetRunning()(*bool)
    GetSuccessful()(*bool)
    SetNodes(value []V1ConfigApplyGetResponse_nodesable)()
    SetRunning(value *bool)()
    SetSuccessful(value *bool)()
}
