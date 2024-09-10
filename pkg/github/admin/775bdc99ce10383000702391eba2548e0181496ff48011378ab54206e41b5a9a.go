package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PreReceiveEnvironmentsItemWithPre_receive_environment_PatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // URL from which to download a tarball of this environment.
    image_url *string
    // This pre-receive environment's new name.
    name *string
}
// NewPreReceiveEnvironmentsItemWithPre_receive_environment_PatchRequestBody instantiates a new PreReceiveEnvironmentsItemWithPre_receive_environment_PatchRequestBody and sets the default values.
func NewPreReceiveEnvironmentsItemWithPre_receive_environment_PatchRequestBody()(*PreReceiveEnvironmentsItemWithPre_receive_environment_PatchRequestBody) {
    m := &PreReceiveEnvironmentsItemWithPre_receive_environment_PatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePreReceiveEnvironmentsItemWithPre_receive_environment_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePreReceiveEnvironmentsItemWithPre_receive_environment_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPreReceiveEnvironmentsItemWithPre_receive_environment_PatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_PatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_PatchRequestBody) GetImageUrl()(*string) {
    return m.image_url
}
// GetName gets the name property value. This pre-receive environment's new name.
// returns a *string when successful
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_PatchRequestBody) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_PatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetImageUrl sets the image_url property value. URL from which to download a tarball of this environment.
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_PatchRequestBody) SetImageUrl(value *string)() {
    m.image_url = value
}
// SetName sets the name property value. This pre-receive environment's new name.
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_PatchRequestBody) SetName(value *string)() {
    m.name = value
}
type PreReceiveEnvironmentsItemWithPre_receive_environment_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetImageUrl()(*string)
    GetName()(*string)
    SetImageUrl(value *string)()
    SetName(value *string)()
}
