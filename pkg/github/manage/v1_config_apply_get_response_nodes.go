package manage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1ConfigApplyGetResponse_nodes struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The hostname of the node
    hostname *string
    // The unique Run ID of the ghe-config-apply run on the host
    run_id *string
    // Whether the ghe-config-apply run is still running on the host
    running *bool
    // Whether the ghe-config-apply run was successful on the host
    successful *bool
}
// NewV1ConfigApplyGetResponse_nodes instantiates a new V1ConfigApplyGetResponse_nodes and sets the default values.
func NewV1ConfigApplyGetResponse_nodes()(*V1ConfigApplyGetResponse_nodes) {
    m := &V1ConfigApplyGetResponse_nodes{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1ConfigApplyGetResponse_nodesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1ConfigApplyGetResponse_nodesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1ConfigApplyGetResponse_nodes(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1ConfigApplyGetResponse_nodes) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1ConfigApplyGetResponse_nodes) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["run_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunId(val)
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
// GetHostname gets the hostname property value. The hostname of the node
// returns a *string when successful
func (m *V1ConfigApplyGetResponse_nodes) GetHostname()(*string) {
    return m.hostname
}
// GetRunId gets the run_id property value. The unique Run ID of the ghe-config-apply run on the host
// returns a *string when successful
func (m *V1ConfigApplyGetResponse_nodes) GetRunId()(*string) {
    return m.run_id
}
// GetRunning gets the running property value. Whether the ghe-config-apply run is still running on the host
// returns a *bool when successful
func (m *V1ConfigApplyGetResponse_nodes) GetRunning()(*bool) {
    return m.running
}
// GetSuccessful gets the successful property value. Whether the ghe-config-apply run was successful on the host
// returns a *bool when successful
func (m *V1ConfigApplyGetResponse_nodes) GetSuccessful()(*bool) {
    return m.successful
}
// Serialize serializes information the current object
func (m *V1ConfigApplyGetResponse_nodes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("hostname", m.GetHostname())
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
        err := writer.WriteStringValue("run_id", m.GetRunId())
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
func (m *V1ConfigApplyGetResponse_nodes) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHostname sets the hostname property value. The hostname of the node
func (m *V1ConfigApplyGetResponse_nodes) SetHostname(value *string)() {
    m.hostname = value
}
// SetRunId sets the run_id property value. The unique Run ID of the ghe-config-apply run on the host
func (m *V1ConfigApplyGetResponse_nodes) SetRunId(value *string)() {
    m.run_id = value
}
// SetRunning sets the running property value. Whether the ghe-config-apply run is still running on the host
func (m *V1ConfigApplyGetResponse_nodes) SetRunning(value *bool)() {
    m.running = value
}
// SetSuccessful sets the successful property value. Whether the ghe-config-apply run was successful on the host
func (m *V1ConfigApplyGetResponse_nodes) SetSuccessful(value *bool)() {
    m.successful = value
}
type V1ConfigApplyGetResponse_nodesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHostname()(*string)
    GetRunId()(*string)
    GetRunning()(*bool)
    GetSuccessful()(*bool)
    SetHostname(value *string)()
    SetRunId(value *string)()
    SetRunning(value *bool)()
    SetSuccessful(value *bool)()
}
