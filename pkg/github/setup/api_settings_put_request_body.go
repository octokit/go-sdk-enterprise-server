package setup

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ApiSettingsPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A JSON string with the new settings. Note that you only need to pass the specific settings you want to modify. For a list of the available settings, see the [Get settings endpoint](https://docs.github.com/enterprise-server@3.11/rest/enterprise-admin/management-console#get-settings).
    settings *string
}
// NewApiSettingsPutRequestBody instantiates a new ApiSettingsPutRequestBody and sets the default values.
func NewApiSettingsPutRequestBody()(*ApiSettingsPutRequestBody) {
    m := &ApiSettingsPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateApiSettingsPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApiSettingsPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApiSettingsPutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ApiSettingsPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApiSettingsPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val)
        }
        return nil
    }
    return res
}
// GetSettings gets the settings property value. A JSON string with the new settings. Note that you only need to pass the specific settings you want to modify. For a list of the available settings, see the [Get settings endpoint](https://docs.github.com/enterprise-server@3.11/rest/enterprise-admin/management-console#get-settings).
// returns a *string when successful
func (m *ApiSettingsPutRequestBody) GetSettings()(*string) {
    return m.settings
}
// Serialize serializes information the current object
func (m *ApiSettingsPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("settings", m.GetSettings())
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
func (m *ApiSettingsPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSettings sets the settings property value. A JSON string with the new settings. Note that you only need to pass the specific settings you want to modify. For a list of the available settings, see the [Get settings endpoint](https://docs.github.com/enterprise-server@3.11/rest/enterprise-admin/management-console#get-settings).
func (m *ApiSettingsPutRequestBody) SetSettings(value *string)() {
    m.settings = value
}
type ApiSettingsPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSettings()(*string)
    SetSettings(value *string)()
}
