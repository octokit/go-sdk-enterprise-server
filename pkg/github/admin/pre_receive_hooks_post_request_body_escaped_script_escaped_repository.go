package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PreReceiveHooksPostRequestBody_script_repository the GitHub repository where the script is kept.
type PreReceiveHooksPostRequestBody_script_repository struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewPreReceiveHooksPostRequestBody_script_repository instantiates a new PreReceiveHooksPostRequestBody_script_repository and sets the default values.
func NewPreReceiveHooksPostRequestBody_script_repository()(*PreReceiveHooksPostRequestBody_script_repository) {
    m := &PreReceiveHooksPostRequestBody_script_repository{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePreReceiveHooksPostRequestBody_script_repositoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePreReceiveHooksPostRequestBody_script_repositoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPreReceiveHooksPostRequestBody_script_repository(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PreReceiveHooksPostRequestBody_script_repository) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PreReceiveHooksPostRequestBody_script_repository) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *PreReceiveHooksPostRequestBody_script_repository) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PreReceiveHooksPostRequestBody_script_repository) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type PreReceiveHooksPostRequestBody_script_repositoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
