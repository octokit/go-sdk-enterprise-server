package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PreReceiveHook_script_repository struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The full_name property
    full_name *string
    // The html_url property
    html_url *string
    // The id property
    id *int32
    // The url property
    url *string
}
// NewPreReceiveHook_script_repository instantiates a new PreReceiveHook_script_repository and sets the default values.
func NewPreReceiveHook_script_repository()(*PreReceiveHook_script_repository) {
    m := &PreReceiveHook_script_repository{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePreReceiveHook_script_repositoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePreReceiveHook_script_repositoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPreReceiveHook_script_repository(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PreReceiveHook_script_repository) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PreReceiveHook_script_repository) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["full_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullName(val)
        }
        return nil
    }
    res["html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtmlUrl(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetFullName gets the full_name property value. The full_name property
// returns a *string when successful
func (m *PreReceiveHook_script_repository) GetFullName()(*string) {
    return m.full_name
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *PreReceiveHook_script_repository) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *PreReceiveHook_script_repository) GetId()(*int32) {
    return m.id
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *PreReceiveHook_script_repository) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *PreReceiveHook_script_repository) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("full_name", m.GetFullName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *PreReceiveHook_script_repository) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFullName sets the full_name property value. The full_name property
func (m *PreReceiveHook_script_repository) SetFullName(value *string)() {
    m.full_name = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *PreReceiveHook_script_repository) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *PreReceiveHook_script_repository) SetId(value *int32)() {
    m.id = value
}
// SetUrl sets the url property value. The url property
func (m *PreReceiveHook_script_repository) SetUrl(value *string)() {
    m.url = value
}
type PreReceiveHook_script_repositoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFullName()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetUrl()(*string)
    SetFullName(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetUrl(value *string)()
}
