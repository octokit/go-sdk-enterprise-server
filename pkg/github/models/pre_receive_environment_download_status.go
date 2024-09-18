package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PreReceiveEnvironmentDownloadStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The downloaded_at property
    downloaded_at *string
    // The message property
    message *string
    // The state property
    state *string
    // The url property
    url *string
}
// NewPreReceiveEnvironmentDownloadStatus instantiates a new PreReceiveEnvironmentDownloadStatus and sets the default values.
func NewPreReceiveEnvironmentDownloadStatus()(*PreReceiveEnvironmentDownloadStatus) {
    m := &PreReceiveEnvironmentDownloadStatus{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePreReceiveEnvironmentDownloadStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePreReceiveEnvironmentDownloadStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPreReceiveEnvironmentDownloadStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PreReceiveEnvironmentDownloadStatus) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDownloadedAt gets the downloaded_at property value. The downloaded_at property
// returns a *string when successful
func (m *PreReceiveEnvironmentDownloadStatus) GetDownloadedAt()(*string) {
    return m.downloaded_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PreReceiveEnvironmentDownloadStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["downloaded_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDownloadedAt(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
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
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *PreReceiveEnvironmentDownloadStatus) GetMessage()(*string) {
    return m.message
}
// GetState gets the state property value. The state property
// returns a *string when successful
func (m *PreReceiveEnvironmentDownloadStatus) GetState()(*string) {
    return m.state
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *PreReceiveEnvironmentDownloadStatus) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *PreReceiveEnvironmentDownloadStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("downloaded_at", m.GetDownloadedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
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
func (m *PreReceiveEnvironmentDownloadStatus) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDownloadedAt sets the downloaded_at property value. The downloaded_at property
func (m *PreReceiveEnvironmentDownloadStatus) SetDownloadedAt(value *string)() {
    m.downloaded_at = value
}
// SetMessage sets the message property value. The message property
func (m *PreReceiveEnvironmentDownloadStatus) SetMessage(value *string)() {
    m.message = value
}
// SetState sets the state property value. The state property
func (m *PreReceiveEnvironmentDownloadStatus) SetState(value *string)() {
    m.state = value
}
// SetUrl sets the url property value. The url property
func (m *PreReceiveEnvironmentDownloadStatus) SetUrl(value *string)() {
    m.url = value
}
type PreReceiveEnvironmentDownloadStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDownloadedAt()(*string)
    GetMessage()(*string)
    GetState()(*string)
    GetUrl()(*string)
    SetDownloadedAt(value *string)()
    SetMessage(value *string)()
    SetState(value *string)()
    SetUrl(value *string)()
}
