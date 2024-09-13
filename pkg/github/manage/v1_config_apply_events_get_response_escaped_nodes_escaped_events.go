package manage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1ConfigApplyEventsGetResponse_nodes_events struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The body property
    body *string
    // The config_run_id property
    config_run_id *string
    // The event_name property
    event_name *string
    // The hostname property
    hostname *string
    // The severity_text property
    severity_text *string
    // The span_depth property
    span_depth *int32
    // The span_id property
    span_id *string
    // The span_parent_id property
    span_parent_id *string
    // The timestamp property
    timestamp *string
    // The topology property
    topology *string
    // The trace_id property
    trace_id *string
}
// NewV1ConfigApplyEventsGetResponse_nodes_events instantiates a new V1ConfigApplyEventsGetResponse_nodes_events and sets the default values.
func NewV1ConfigApplyEventsGetResponse_nodes_events()(*V1ConfigApplyEventsGetResponse_nodes_events) {
    m := &V1ConfigApplyEventsGetResponse_nodes_events{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1ConfigApplyEventsGetResponse_nodes_eventsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1ConfigApplyEventsGetResponse_nodes_eventsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1ConfigApplyEventsGetResponse_nodes_events(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1ConfigApplyEventsGetResponse_nodes_events) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBody gets the body property value. The body property
// returns a *string when successful
func (m *V1ConfigApplyEventsGetResponse_nodes_events) GetBody()(*string) {
    return m.body
}
// GetConfigRunId gets the config_run_id property value. The config_run_id property
// returns a *string when successful
func (m *V1ConfigApplyEventsGetResponse_nodes_events) GetConfigRunId()(*string) {
    return m.config_run_id
}
// GetEventName gets the event_name property value. The event_name property
// returns a *string when successful
func (m *V1ConfigApplyEventsGetResponse_nodes_events) GetEventName()(*string) {
    return m.event_name
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1ConfigApplyEventsGetResponse_nodes_events) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["body"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val)
        }
        return nil
    }
    res["config_run_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigRunId(val)
        }
        return nil
    }
    res["event_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventName(val)
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
    res["severity_text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverityText(val)
        }
        return nil
    }
    res["span_depth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpanDepth(val)
        }
        return nil
    }
    res["span_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpanId(val)
        }
        return nil
    }
    res["span_parent_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpanParentId(val)
        }
        return nil
    }
    res["timestamp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimestamp(val)
        }
        return nil
    }
    res["topology"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopology(val)
        }
        return nil
    }
    res["trace_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTraceId(val)
        }
        return nil
    }
    return res
}
// GetHostname gets the hostname property value. The hostname property
// returns a *string when successful
func (m *V1ConfigApplyEventsGetResponse_nodes_events) GetHostname()(*string) {
    return m.hostname
}
// GetSeverityText gets the severity_text property value. The severity_text property
// returns a *string when successful
func (m *V1ConfigApplyEventsGetResponse_nodes_events) GetSeverityText()(*string) {
    return m.severity_text
}
// GetSpanDepth gets the span_depth property value. The span_depth property
// returns a *int32 when successful
func (m *V1ConfigApplyEventsGetResponse_nodes_events) GetSpanDepth()(*int32) {
    return m.span_depth
}
// GetSpanId gets the span_id property value. The span_id property
// returns a *string when successful
func (m *V1ConfigApplyEventsGetResponse_nodes_events) GetSpanId()(*string) {
    return m.span_id
}
// GetSpanParentId gets the span_parent_id property value. The span_parent_id property
// returns a *string when successful
func (m *V1ConfigApplyEventsGetResponse_nodes_events) GetSpanParentId()(*string) {
    return m.span_parent_id
}
// GetTimestamp gets the timestamp property value. The timestamp property
// returns a *string when successful
func (m *V1ConfigApplyEventsGetResponse_nodes_events) GetTimestamp()(*string) {
    return m.timestamp
}
// GetTopology gets the topology property value. The topology property
// returns a *string when successful
func (m *V1ConfigApplyEventsGetResponse_nodes_events) GetTopology()(*string) {
    return m.topology
}
// GetTraceId gets the trace_id property value. The trace_id property
// returns a *string when successful
func (m *V1ConfigApplyEventsGetResponse_nodes_events) GetTraceId()(*string) {
    return m.trace_id
}
// Serialize serializes information the current object
func (m *V1ConfigApplyEventsGetResponse_nodes_events) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("config_run_id", m.GetConfigRunId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("event_name", m.GetEventName())
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
        err := writer.WriteStringValue("severity_text", m.GetSeverityText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("span_depth", m.GetSpanDepth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("span_id", m.GetSpanId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("span_parent_id", m.GetSpanParentId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timestamp", m.GetTimestamp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("topology", m.GetTopology())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("trace_id", m.GetTraceId())
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
func (m *V1ConfigApplyEventsGetResponse_nodes_events) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBody sets the body property value. The body property
func (m *V1ConfigApplyEventsGetResponse_nodes_events) SetBody(value *string)() {
    m.body = value
}
// SetConfigRunId sets the config_run_id property value. The config_run_id property
func (m *V1ConfigApplyEventsGetResponse_nodes_events) SetConfigRunId(value *string)() {
    m.config_run_id = value
}
// SetEventName sets the event_name property value. The event_name property
func (m *V1ConfigApplyEventsGetResponse_nodes_events) SetEventName(value *string)() {
    m.event_name = value
}
// SetHostname sets the hostname property value. The hostname property
func (m *V1ConfigApplyEventsGetResponse_nodes_events) SetHostname(value *string)() {
    m.hostname = value
}
// SetSeverityText sets the severity_text property value. The severity_text property
func (m *V1ConfigApplyEventsGetResponse_nodes_events) SetSeverityText(value *string)() {
    m.severity_text = value
}
// SetSpanDepth sets the span_depth property value. The span_depth property
func (m *V1ConfigApplyEventsGetResponse_nodes_events) SetSpanDepth(value *int32)() {
    m.span_depth = value
}
// SetSpanId sets the span_id property value. The span_id property
func (m *V1ConfigApplyEventsGetResponse_nodes_events) SetSpanId(value *string)() {
    m.span_id = value
}
// SetSpanParentId sets the span_parent_id property value. The span_parent_id property
func (m *V1ConfigApplyEventsGetResponse_nodes_events) SetSpanParentId(value *string)() {
    m.span_parent_id = value
}
// SetTimestamp sets the timestamp property value. The timestamp property
func (m *V1ConfigApplyEventsGetResponse_nodes_events) SetTimestamp(value *string)() {
    m.timestamp = value
}
// SetTopology sets the topology property value. The topology property
func (m *V1ConfigApplyEventsGetResponse_nodes_events) SetTopology(value *string)() {
    m.topology = value
}
// SetTraceId sets the trace_id property value. The trace_id property
func (m *V1ConfigApplyEventsGetResponse_nodes_events) SetTraceId(value *string)() {
    m.trace_id = value
}
type V1ConfigApplyEventsGetResponse_nodes_eventsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    GetConfigRunId()(*string)
    GetEventName()(*string)
    GetHostname()(*string)
    GetSeverityText()(*string)
    GetSpanDepth()(*int32)
    GetSpanId()(*string)
    GetSpanParentId()(*string)
    GetTimestamp()(*string)
    GetTopology()(*string)
    GetTraceId()(*string)
    SetBody(value *string)()
    SetConfigRunId(value *string)()
    SetEventName(value *string)()
    SetHostname(value *string)()
    SetSeverityText(value *string)()
    SetSpanDepth(value *int32)()
    SetSpanId(value *string)()
    SetSpanParentId(value *string)()
    SetTimestamp(value *string)()
    SetTopology(value *string)()
    SetTraceId(value *string)()
}
