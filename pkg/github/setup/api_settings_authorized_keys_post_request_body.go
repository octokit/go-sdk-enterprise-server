package setup

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ApiSettingsAuthorizedKeysPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The public SSH key.
    authorized_key *string
}
// NewApiSettingsAuthorizedKeysPostRequestBody instantiates a new ApiSettingsAuthorizedKeysPostRequestBody and sets the default values.
func NewApiSettingsAuthorizedKeysPostRequestBody()(*ApiSettingsAuthorizedKeysPostRequestBody) {
    m := &ApiSettingsAuthorizedKeysPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateApiSettingsAuthorizedKeysPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApiSettingsAuthorizedKeysPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApiSettingsAuthorizedKeysPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ApiSettingsAuthorizedKeysPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthorizedKey gets the authorized_key property value. The public SSH key.
// returns a *string when successful
func (m *ApiSettingsAuthorizedKeysPostRequestBody) GetAuthorizedKey()(*string) {
    return m.authorized_key
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApiSettingsAuthorizedKeysPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authorized_key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizedKey(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ApiSettingsAuthorizedKeysPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("authorized_key", m.GetAuthorizedKey())
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
func (m *ApiSettingsAuthorizedKeysPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthorizedKey sets the authorized_key property value. The public SSH key.
func (m *ApiSettingsAuthorizedKeysPostRequestBody) SetAuthorizedKey(value *string)() {
    m.authorized_key = value
}
type ApiSettingsAuthorizedKeysPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorizedKey()(*string)
    SetAuthorizedKey(value *string)()
}
