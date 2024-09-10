package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type HooksItemWithHook_PatchRequestBody struct {
    // Determines if notifications are sent when the webhook is triggered. Set to `true` to send notifications.
    active *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Key/value pairs to provide settings for this webhook.
    config HooksItemWithHook_PatchRequestBody_configable
    // The [events](https://docs.github.com/enterprise-server@3.11/webhooks/event-payloads) that trigger this webhook. A global webhook can be triggered by `user` and `organization` events. Default: `user` and `organization`.
    events []string
}
// NewHooksItemWithHook_PatchRequestBody instantiates a new HooksItemWithHook_PatchRequestBody and sets the default values.
func NewHooksItemWithHook_PatchRequestBody()(*HooksItemWithHook_PatchRequestBody) {
    m := &HooksItemWithHook_PatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateHooksItemWithHook_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHooksItemWithHook_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHooksItemWithHook_PatchRequestBody(), nil
}
// GetActive gets the active property value. Determines if notifications are sent when the webhook is triggered. Set to `true` to send notifications.
// returns a *bool when successful
func (m *HooksItemWithHook_PatchRequestBody) GetActive()(*bool) {
    return m.active
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *HooksItemWithHook_PatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfig gets the config property value. Key/value pairs to provide settings for this webhook.
// returns a HooksItemWithHook_PatchRequestBody_configable when successful
func (m *HooksItemWithHook_PatchRequestBody) GetConfig()(HooksItemWithHook_PatchRequestBody_configable) {
    return m.config
}
// GetEvents gets the events property value. The [events](https://docs.github.com/enterprise-server@3.11/webhooks/event-payloads) that trigger this webhook. A global webhook can be triggered by `user` and `organization` events. Default: `user` and `organization`.
// returns a []string when successful
func (m *HooksItemWithHook_PatchRequestBody) GetEvents()([]string) {
    return m.events
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HooksItemWithHook_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["active"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActive(val)
        }
        return nil
    }
    res["config"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateHooksItemWithHook_PatchRequestBody_configFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfig(val.(HooksItemWithHook_PatchRequestBody_configable))
        }
        return nil
    }
    res["events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetEvents(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *HooksItemWithHook_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("active", m.GetActive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("config", m.GetConfig())
        if err != nil {
            return err
        }
    }
    if m.GetEvents() != nil {
        err := writer.WriteCollectionOfStringValues("events", m.GetEvents())
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
// SetActive sets the active property value. Determines if notifications are sent when the webhook is triggered. Set to `true` to send notifications.
func (m *HooksItemWithHook_PatchRequestBody) SetActive(value *bool)() {
    m.active = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HooksItemWithHook_PatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfig sets the config property value. Key/value pairs to provide settings for this webhook.
func (m *HooksItemWithHook_PatchRequestBody) SetConfig(value HooksItemWithHook_PatchRequestBody_configable)() {
    m.config = value
}
// SetEvents sets the events property value. The [events](https://docs.github.com/enterprise-server@3.11/webhooks/event-payloads) that trigger this webhook. A global webhook can be triggered by `user` and `organization` events. Default: `user` and `organization`.
func (m *HooksItemWithHook_PatchRequestBody) SetEvents(value []string)() {
    m.events = value
}
type HooksItemWithHook_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActive()(*bool)
    GetConfig()(HooksItemWithHook_PatchRequestBody_configable)
    GetEvents()([]string)
    SetActive(value *bool)()
    SetConfig(value HooksItemWithHook_PatchRequestBody_configable)()
    SetEvents(value []string)()
}
