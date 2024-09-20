package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemPreReceiveHooksItemWithPre_receive_hook_PatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether repositories can override enforcement.
    allow_downstream_configuration *bool
    // The state of enforcement for the hook on this repository.
    enforcement *string
}
// NewItemPreReceiveHooksItemWithPre_receive_hook_PatchRequestBody instantiates a new ItemPreReceiveHooksItemWithPre_receive_hook_PatchRequestBody and sets the default values.
func NewItemPreReceiveHooksItemWithPre_receive_hook_PatchRequestBody()(*ItemPreReceiveHooksItemWithPre_receive_hook_PatchRequestBody) {
    m := &ItemPreReceiveHooksItemWithPre_receive_hook_PatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemPreReceiveHooksItemWithPre_receive_hook_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPreReceiveHooksItemWithPre_receive_hook_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPreReceiveHooksItemWithPre_receive_hook_PatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemPreReceiveHooksItemWithPre_receive_hook_PatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowDownstreamConfiguration gets the allow_downstream_configuration property value. Whether repositories can override enforcement.
// returns a *bool when successful
func (m *ItemPreReceiveHooksItemWithPre_receive_hook_PatchRequestBody) GetAllowDownstreamConfiguration()(*bool) {
    return m.allow_downstream_configuration
}
// GetEnforcement gets the enforcement property value. The state of enforcement for the hook on this repository.
// returns a *string when successful
func (m *ItemPreReceiveHooksItemWithPre_receive_hook_PatchRequestBody) GetEnforcement()(*string) {
    return m.enforcement
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemPreReceiveHooksItemWithPre_receive_hook_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allow_downstream_configuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDownstreamConfiguration(val)
        }
        return nil
    }
    res["enforcement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforcement(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemPreReceiveHooksItemWithPre_receive_hook_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allow_downstream_configuration", m.GetAllowDownstreamConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("enforcement", m.GetEnforcement())
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
func (m *ItemPreReceiveHooksItemWithPre_receive_hook_PatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowDownstreamConfiguration sets the allow_downstream_configuration property value. Whether repositories can override enforcement.
func (m *ItemPreReceiveHooksItemWithPre_receive_hook_PatchRequestBody) SetAllowDownstreamConfiguration(value *bool)() {
    m.allow_downstream_configuration = value
}
// SetEnforcement sets the enforcement property value. The state of enforcement for the hook on this repository.
func (m *ItemPreReceiveHooksItemWithPre_receive_hook_PatchRequestBody) SetEnforcement(value *string)() {
    m.enforcement = value
}
type ItemPreReceiveHooksItemWithPre_receive_hook_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowDownstreamConfiguration()(*bool)
    GetEnforcement()(*string)
    SetAllowDownstreamConfiguration(value *bool)()
    SetEnforcement(value *string)()
}
