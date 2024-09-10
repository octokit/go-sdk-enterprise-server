package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PreReceiveHooksItemWithPre_receive_hook_PatchRequestBody_environment the pre-receive environment where the script is executed.
type PreReceiveHooksItemWithPre_receive_hook_PatchRequestBody_environment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewPreReceiveHooksItemWithPre_receive_hook_PatchRequestBody_environment instantiates a new PreReceiveHooksItemWithPre_receive_hook_PatchRequestBody_environment and sets the default values.
func NewPreReceiveHooksItemWithPre_receive_hook_PatchRequestBody_environment()(*PreReceiveHooksItemWithPre_receive_hook_PatchRequestBody_environment) {
    m := &PreReceiveHooksItemWithPre_receive_hook_PatchRequestBody_environment{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePreReceiveHooksItemWithPre_receive_hook_PatchRequestBody_environmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePreReceiveHooksItemWithPre_receive_hook_PatchRequestBody_environmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPreReceiveHooksItemWithPre_receive_hook_PatchRequestBody_environment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PreReceiveHooksItemWithPre_receive_hook_PatchRequestBody_environment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PreReceiveHooksItemWithPre_receive_hook_PatchRequestBody_environment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *PreReceiveHooksItemWithPre_receive_hook_PatchRequestBody_environment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PreReceiveHooksItemWithPre_receive_hook_PatchRequestBody_environment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type PreReceiveHooksItemWithPre_receive_hook_PatchRequestBody_environmentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
