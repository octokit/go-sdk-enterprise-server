package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GlobalHook2 struct {
    // The active property
    active *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The config property
    config GlobalHook2_configable
    // The created_at property
    created_at *string
    // The events property
    events []string
    // The id property
    id *int32
    // The name property
    name *string
    // The ping_url property
    ping_url *string
    // The type property
    typeEscaped *string
    // The updated_at property
    updated_at *string
    // The url property
    url *string
}
// NewGlobalHook2 instantiates a new GlobalHook2 and sets the default values.
func NewGlobalHook2()(*GlobalHook2) {
    m := &GlobalHook2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGlobalHook2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGlobalHook2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGlobalHook2(), nil
}
// GetActive gets the active property value. The active property
// returns a *bool when successful
func (m *GlobalHook2) GetActive()(*bool) {
    return m.active
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GlobalHook2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfig gets the config property value. The config property
// returns a GlobalHook2_configable when successful
func (m *GlobalHook2) GetConfig()(GlobalHook2_configable) {
    return m.config
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *string when successful
func (m *GlobalHook2) GetCreatedAt()(*string) {
    return m.created_at
}
// GetEvents gets the events property value. The events property
// returns a []string when successful
func (m *GlobalHook2) GetEvents()([]string) {
    return m.events
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GlobalHook2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["active"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActive(val)
        }
        return nil
    }
    res["config"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGlobalHook2_configFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfig(val.(GlobalHook2_configable))
        }
        return nil
    }
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
    res["events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetEvents(res)
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
    res["ping_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPingUrl(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
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
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *GlobalHook2) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *GlobalHook2) GetName()(*string) {
    return m.name
}
// GetPingUrl gets the ping_url property value. The ping_url property
// returns a *string when successful
func (m *GlobalHook2) GetPingUrl()(*string) {
    return m.ping_url
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *GlobalHook2) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *string when successful
func (m *GlobalHook2) GetUpdatedAt()(*string) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *GlobalHook2) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *GlobalHook2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("active", m.GetActive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("config", m.GetConfig())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    if m.GetEvents() != nil {
        err := writer.WriteCollectionOfStringValues("events", m.GetEvents())
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
        err := writer.WriteStringValue("ping_url", m.GetPingUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("updated_at", m.GetUpdatedAt())
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
// SetActive sets the active property value. The active property
func (m *GlobalHook2) SetActive(value *bool)() {
    m.active = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GlobalHook2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfig sets the config property value. The config property
func (m *GlobalHook2) SetConfig(value GlobalHook2_configable)() {
    m.config = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *GlobalHook2) SetCreatedAt(value *string)() {
    m.created_at = value
}
// SetEvents sets the events property value. The events property
func (m *GlobalHook2) SetEvents(value []string)() {
    m.events = value
}
// SetId sets the id property value. The id property
func (m *GlobalHook2) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name property
func (m *GlobalHook2) SetName(value *string)() {
    m.name = value
}
// SetPingUrl sets the ping_url property value. The ping_url property
func (m *GlobalHook2) SetPingUrl(value *string)() {
    m.ping_url = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *GlobalHook2) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *GlobalHook2) SetUpdatedAt(value *string)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *GlobalHook2) SetUrl(value *string)() {
    m.url = value
}
type GlobalHook2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActive()(*bool)
    GetConfig()(GlobalHook2_configable)
    GetCreatedAt()(*string)
    GetEvents()([]string)
    GetId()(*int32)
    GetName()(*string)
    GetPingUrl()(*string)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*string)
    GetUrl()(*string)
    SetActive(value *bool)()
    SetConfig(value GlobalHook2_configable)()
    SetCreatedAt(value *string)()
    SetEvents(value []string)()
    SetId(value *int32)()
    SetName(value *string)()
    SetPingUrl(value *string)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *string)()
    SetUrl(value *string)()
}
