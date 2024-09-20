package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RepositoryPreReceiveHook struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The configuration_url property
    configuration_url *string
    // The enforcement property
    enforcement *string
    // The id property
    id *int32
    // The name property
    name *string
}
// NewRepositoryPreReceiveHook instantiates a new RepositoryPreReceiveHook and sets the default values.
func NewRepositoryPreReceiveHook()(*RepositoryPreReceiveHook) {
    m := &RepositoryPreReceiveHook{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryPreReceiveHookFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryPreReceiveHookFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryPreReceiveHook(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryPreReceiveHook) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfigurationUrl gets the configuration_url property value. The configuration_url property
// returns a *string when successful
func (m *RepositoryPreReceiveHook) GetConfigurationUrl()(*string) {
    return m.configuration_url
}
// GetEnforcement gets the enforcement property value. The enforcement property
// returns a *string when successful
func (m *RepositoryPreReceiveHook) GetEnforcement()(*string) {
    return m.enforcement
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryPreReceiveHook) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["configuration_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationUrl(val)
        }
        return nil
    }
    res["enforcement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforcement(val)
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
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *RepositoryPreReceiveHook) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *RepositoryPreReceiveHook) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *RepositoryPreReceiveHook) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("configuration_url", m.GetConfigurationUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("enforcement", m.GetEnforcement())
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
func (m *RepositoryPreReceiveHook) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfigurationUrl sets the configuration_url property value. The configuration_url property
func (m *RepositoryPreReceiveHook) SetConfigurationUrl(value *string)() {
    m.configuration_url = value
}
// SetEnforcement sets the enforcement property value. The enforcement property
func (m *RepositoryPreReceiveHook) SetEnforcement(value *string)() {
    m.enforcement = value
}
// SetId sets the id property value. The id property
func (m *RepositoryPreReceiveHook) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name property
func (m *RepositoryPreReceiveHook) SetName(value *string)() {
    m.name = value
}
type RepositoryPreReceiveHookable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfigurationUrl()(*string)
    GetEnforcement()(*string)
    GetId()(*int32)
    GetName()(*string)
    SetConfigurationUrl(value *string)()
    SetEnforcement(value *string)()
    SetId(value *int32)()
    SetName(value *string)()
}
