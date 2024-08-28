package setup

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ApiStartPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The content of your _.ghl_ license file.
    license *string
    // You **must** provide a password _only if_ you are uploading your license for the first time. If you previously set a password through the web interface, you don't need this parameter.
    password *string
    // An optional JSON string containing the installation settings. For a list of the available settings, see the [Get settings endpoint](https://docs.github.com/enterprise-server@3.12/rest/enterprise-admin/management-console#get-settings).
    settings *string
}
// NewApiStartPostRequestBody instantiates a new ApiStartPostRequestBody and sets the default values.
func NewApiStartPostRequestBody()(*ApiStartPostRequestBody) {
    m := &ApiStartPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateApiStartPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApiStartPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApiStartPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ApiStartPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApiStartPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["license"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicense(val)
        }
        return nil
    }
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
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
// GetLicense gets the license property value. The content of your _.ghl_ license file.
// returns a *string when successful
func (m *ApiStartPostRequestBody) GetLicense()(*string) {
    return m.license
}
// GetPassword gets the password property value. You **must** provide a password _only if_ you are uploading your license for the first time. If you previously set a password through the web interface, you don't need this parameter.
// returns a *string when successful
func (m *ApiStartPostRequestBody) GetPassword()(*string) {
    return m.password
}
// GetSettings gets the settings property value. An optional JSON string containing the installation settings. For a list of the available settings, see the [Get settings endpoint](https://docs.github.com/enterprise-server@3.12/rest/enterprise-admin/management-console#get-settings).
// returns a *string when successful
func (m *ApiStartPostRequestBody) GetSettings()(*string) {
    return m.settings
}
// Serialize serializes information the current object
func (m *ApiStartPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("license", m.GetLicense())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
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
func (m *ApiStartPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLicense sets the license property value. The content of your _.ghl_ license file.
func (m *ApiStartPostRequestBody) SetLicense(value *string)() {
    m.license = value
}
// SetPassword sets the password property value. You **must** provide a password _only if_ you are uploading your license for the first time. If you previously set a password through the web interface, you don't need this parameter.
func (m *ApiStartPostRequestBody) SetPassword(value *string)() {
    m.password = value
}
// SetSettings sets the settings property value. An optional JSON string containing the installation settings. For a list of the available settings, see the [Get settings endpoint](https://docs.github.com/enterprise-server@3.12/rest/enterprise-admin/management-console#get-settings).
func (m *ApiStartPostRequestBody) SetSettings(value *string)() {
    m.settings = value
}
type ApiStartPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLicense()(*string)
    GetPassword()(*string)
    GetSettings()(*string)
    SetLicense(value *string)()
    SetPassword(value *string)()
    SetSettings(value *string)()
}
