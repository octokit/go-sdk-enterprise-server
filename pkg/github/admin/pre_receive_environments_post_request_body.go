package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PreReceiveEnvironmentsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // URL from which to download a tarball of this environment.
    image_url *string
    // The new pre-receive environment's name.
    name *string
}
// NewPreReceiveEnvironmentsPostRequestBody instantiates a new PreReceiveEnvironmentsPostRequestBody and sets the default values.
func NewPreReceiveEnvironmentsPostRequestBody()(*PreReceiveEnvironmentsPostRequestBody) {
    m := &PreReceiveEnvironmentsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePreReceiveEnvironmentsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePreReceiveEnvironmentsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPreReceiveEnvironmentsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PreReceiveEnvironmentsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PreReceiveEnvironmentsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    return res
}
// GetImageUrl gets the image_url property value. URL from which to download a tarball of this environment.
// returns a *string when successful
func (m *PreReceiveEnvironmentsPostRequestBody) GetImageUrl()(*string) {
    return m.image_url
}
// GetName gets the name property value. The new pre-receive environment's name.
// returns a *string when successful
func (m *PreReceiveEnvironmentsPostRequestBody) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *PreReceiveEnvironmentsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PreReceiveEnvironmentsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetImageUrl sets the image_url property value. URL from which to download a tarball of this environment.
func (m *PreReceiveEnvironmentsPostRequestBody) SetImageUrl(value *string)() {
    m.image_url = value
}
// SetName sets the name property value. The new pre-receive environment's name.
func (m *PreReceiveEnvironmentsPostRequestBody) SetName(value *string)() {
    m.name = value
}
type PreReceiveEnvironmentsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetImageUrl()(*string)
    GetName()(*string)
    SetImageUrl(value *string)()
    SetName(value *string)()
}
