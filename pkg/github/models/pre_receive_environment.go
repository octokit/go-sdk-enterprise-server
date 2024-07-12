package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PreReceiveEnvironment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created_at property
    created_at *string
    // The default_environment property
    default_environment *bool
    // The download property
    download PreReceiveEnvironment_downloadable
    // The hooks_count property
    hooks_count *int32
    // The html_url property
    html_url *string
    // The id property
    id *int64
    // The image_url property
    image_url *string
    // The name property
    name *string
    // The url property
    url *string
}
// NewPreReceiveEnvironment instantiates a new PreReceiveEnvironment and sets the default values.
func NewPreReceiveEnvironment()(*PreReceiveEnvironment) {
    m := &PreReceiveEnvironment{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePreReceiveEnvironmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePreReceiveEnvironmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPreReceiveEnvironment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PreReceiveEnvironment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *string when successful
func (m *PreReceiveEnvironment) GetCreatedAt()(*string) {
    return m.created_at
}
// GetDefaultEnvironment gets the default_environment property value. The default_environment property
// returns a *bool when successful
func (m *PreReceiveEnvironment) GetDefaultEnvironment()(*bool) {
    return m.default_environment
}
// GetDownload gets the download property value. The download property
// returns a PreReceiveEnvironment_downloadable when successful
func (m *PreReceiveEnvironment) GetDownload()(PreReceiveEnvironment_downloadable) {
    return m.download
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PreReceiveEnvironment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["default_environment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultEnvironment(val)
        }
        return nil
    }
    res["download"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePreReceiveEnvironment_downloadFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDownload(val.(PreReceiveEnvironment_downloadable))
        }
        return nil
    }
    res["hooks_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHooksCount(val)
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
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["image_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageUrl(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
// GetHooksCount gets the hooks_count property value. The hooks_count property
// returns a *int32 when successful
func (m *PreReceiveEnvironment) GetHooksCount()(*int32) {
    return m.hooks_count
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *PreReceiveEnvironment) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *PreReceiveEnvironment) GetId()(*int64) {
    return m.id
}
// GetImageUrl gets the image_url property value. The image_url property
// returns a *string when successful
func (m *PreReceiveEnvironment) GetImageUrl()(*string) {
    return m.image_url
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *PreReceiveEnvironment) GetName()(*string) {
    return m.name
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *PreReceiveEnvironment) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *PreReceiveEnvironment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("default_environment", m.GetDefaultEnvironment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("download", m.GetDownload())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("hooks_count", m.GetHooksCount())
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
        err := writer.WriteInt64Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("image_url", m.GetImageUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *PreReceiveEnvironment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *PreReceiveEnvironment) SetCreatedAt(value *string)() {
    m.created_at = value
}
// SetDefaultEnvironment sets the default_environment property value. The default_environment property
func (m *PreReceiveEnvironment) SetDefaultEnvironment(value *bool)() {
    m.default_environment = value
}
// SetDownload sets the download property value. The download property
func (m *PreReceiveEnvironment) SetDownload(value PreReceiveEnvironment_downloadable)() {
    m.download = value
}
// SetHooksCount sets the hooks_count property value. The hooks_count property
func (m *PreReceiveEnvironment) SetHooksCount(value *int32)() {
    m.hooks_count = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *PreReceiveEnvironment) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *PreReceiveEnvironment) SetId(value *int64)() {
    m.id = value
}
// SetImageUrl sets the image_url property value. The image_url property
func (m *PreReceiveEnvironment) SetImageUrl(value *string)() {
    m.image_url = value
}
// SetName sets the name property value. The name property
func (m *PreReceiveEnvironment) SetName(value *string)() {
    m.name = value
}
// SetUrl sets the url property value. The url property
func (m *PreReceiveEnvironment) SetUrl(value *string)() {
    m.url = value
}
type PreReceiveEnvironmentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*string)
    GetDefaultEnvironment()(*bool)
    GetDownload()(PreReceiveEnvironment_downloadable)
    GetHooksCount()(*int32)
    GetHtmlUrl()(*string)
    GetId()(*int64)
    GetImageUrl()(*string)
    GetName()(*string)
    GetUrl()(*string)
    SetCreatedAt(value *string)()
    SetDefaultEnvironment(value *bool)()
    SetDownload(value PreReceiveEnvironment_downloadable)()
    SetHooksCount(value *int32)()
    SetHtmlUrl(value *string)()
    SetId(value *int64)()
    SetImageUrl(value *string)()
    SetName(value *string)()
    SetUrl(value *string)()
}
