package manage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1ConfigApplyEventsGetResponse_nodes struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The events property
    events []V1ConfigApplyEventsGetResponse_nodes_eventsable
    // Unique ID of the last response from a host used for pagination
    last_request_id *string
    // Hostname of the node
    node *string
}
// NewV1ConfigApplyEventsGetResponse_nodes instantiates a new V1ConfigApplyEventsGetResponse_nodes and sets the default values.
func NewV1ConfigApplyEventsGetResponse_nodes()(*V1ConfigApplyEventsGetResponse_nodes) {
    m := &V1ConfigApplyEventsGetResponse_nodes{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1ConfigApplyEventsGetResponse_nodesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1ConfigApplyEventsGetResponse_nodesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1ConfigApplyEventsGetResponse_nodes(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1ConfigApplyEventsGetResponse_nodes) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEvents gets the events property value. The events property
// returns a []V1ConfigApplyEventsGetResponse_nodes_eventsable when successful
func (m *V1ConfigApplyEventsGetResponse_nodes) GetEvents()([]V1ConfigApplyEventsGetResponse_nodes_eventsable) {
    return m.events
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1ConfigApplyEventsGetResponse_nodes) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1ConfigApplyEventsGetResponse_nodes_eventsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1ConfigApplyEventsGetResponse_nodes_eventsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1ConfigApplyEventsGetResponse_nodes_eventsable)
                }
            }
            m.SetEvents(res)
        }
        return nil
    }
    res["last_request_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRequestId(val)
        }
        return nil
    }
    res["node"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNode(val)
        }
        return nil
    }
    return res
}
// GetLastRequestId gets the last_request_id property value. Unique ID of the last response from a host used for pagination
// returns a *string when successful
func (m *V1ConfigApplyEventsGetResponse_nodes) GetLastRequestId()(*string) {
    return m.last_request_id
}
// GetNode gets the node property value. Hostname of the node
// returns a *string when successful
func (m *V1ConfigApplyEventsGetResponse_nodes) GetNode()(*string) {
    return m.node
}
// Serialize serializes information the current object
func (m *V1ConfigApplyEventsGetResponse_nodes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEvents()))
        for i, v := range m.GetEvents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("events", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("last_request_id", m.GetLastRequestId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("node", m.GetNode())
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
func (m *V1ConfigApplyEventsGetResponse_nodes) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEvents sets the events property value. The events property
func (m *V1ConfigApplyEventsGetResponse_nodes) SetEvents(value []V1ConfigApplyEventsGetResponse_nodes_eventsable)() {
    m.events = value
}
// SetLastRequestId sets the last_request_id property value. Unique ID of the last response from a host used for pagination
func (m *V1ConfigApplyEventsGetResponse_nodes) SetLastRequestId(value *string)() {
    m.last_request_id = value
}
// SetNode sets the node property value. Hostname of the node
func (m *V1ConfigApplyEventsGetResponse_nodes) SetNode(value *string)() {
    m.node = value
}
type V1ConfigApplyEventsGetResponse_nodesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEvents()([]V1ConfigApplyEventsGetResponse_nodes_eventsable)
    GetLastRequestId()(*string)
    GetNode()(*string)
    SetEvents(value []V1ConfigApplyEventsGetResponse_nodes_eventsable)()
    SetLastRequestId(value *string)()
    SetNode(value *string)()
}
