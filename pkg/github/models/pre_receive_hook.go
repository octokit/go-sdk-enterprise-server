package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PreReceiveHook struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The allow_downstream_configuration property
    allow_downstream_configuration *bool
    // The enforcement property
    enforcement *string
    // The environment property
    environment PreReceiveHook_environmentable
    // The id property
    id *int32
    // The name property
    name *string
    // The script property
    script *string
    // The script_repository property
    script_repository PreReceiveHook_script_repositoryable
}
// NewPreReceiveHook instantiates a new PreReceiveHook and sets the default values.
func NewPreReceiveHook()(*PreReceiveHook) {
    m := &PreReceiveHook{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePreReceiveHookFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePreReceiveHookFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPreReceiveHook(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PreReceiveHook) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowDownstreamConfiguration gets the allow_downstream_configuration property value. The allow_downstream_configuration property
// returns a *bool when successful
func (m *PreReceiveHook) GetAllowDownstreamConfiguration()(*bool) {
    return m.allow_downstream_configuration
}
// GetEnforcement gets the enforcement property value. The enforcement property
// returns a *string when successful
func (m *PreReceiveHook) GetEnforcement()(*string) {
    return m.enforcement
}
// GetEnvironment gets the environment property value. The environment property
// returns a PreReceiveHook_environmentable when successful
func (m *PreReceiveHook) GetEnvironment()(PreReceiveHook_environmentable) {
    return m.environment
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PreReceiveHook) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allow_downstream_configuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDownstreamConfiguration(val)
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
    res["environment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePreReceiveHook_environmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnvironment(val.(PreReceiveHook_environmentable))
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
    res["script"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScript(val)
        }
        return nil
    }
    res["script_repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePreReceiveHook_script_repositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScriptRepository(val.(PreReceiveHook_script_repositoryable))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *PreReceiveHook) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *PreReceiveHook) GetName()(*string) {
    return m.name
}
// GetScript gets the script property value. The script property
// returns a *string when successful
func (m *PreReceiveHook) GetScript()(*string) {
    return m.script
}
// GetScriptRepository gets the script_repository property value. The script_repository property
// returns a PreReceiveHook_script_repositoryable when successful
func (m *PreReceiveHook) GetScriptRepository()(PreReceiveHook_script_repositoryable) {
    return m.script_repository
}
// Serialize serializes information the current object
func (m *PreReceiveHook) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allow_downstream_configuration", m.GetAllowDownstreamConfiguration())
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
        err := writer.WriteObjectValue("environment", m.GetEnvironment())
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
        err := writer.WriteStringValue("script", m.GetScript())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("script_repository", m.GetScriptRepository())
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
func (m *PreReceiveHook) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowDownstreamConfiguration sets the allow_downstream_configuration property value. The allow_downstream_configuration property
func (m *PreReceiveHook) SetAllowDownstreamConfiguration(value *bool)() {
    m.allow_downstream_configuration = value
}
// SetEnforcement sets the enforcement property value. The enforcement property
func (m *PreReceiveHook) SetEnforcement(value *string)() {
    m.enforcement = value
}
// SetEnvironment sets the environment property value. The environment property
func (m *PreReceiveHook) SetEnvironment(value PreReceiveHook_environmentable)() {
    m.environment = value
}
// SetId sets the id property value. The id property
func (m *PreReceiveHook) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name property
func (m *PreReceiveHook) SetName(value *string)() {
    m.name = value
}
// SetScript sets the script property value. The script property
func (m *PreReceiveHook) SetScript(value *string)() {
    m.script = value
}
// SetScriptRepository sets the script_repository property value. The script_repository property
func (m *PreReceiveHook) SetScriptRepository(value PreReceiveHook_script_repositoryable)() {
    m.script_repository = value
}
type PreReceiveHookable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowDownstreamConfiguration()(*bool)
    GetEnforcement()(*string)
    GetEnvironment()(PreReceiveHook_environmentable)
    GetId()(*int32)
    GetName()(*string)
    GetScript()(*string)
    GetScriptRepository()(PreReceiveHook_script_repositoryable)
    SetAllowDownstreamConfiguration(value *bool)()
    SetEnforcement(value *string)()
    SetEnvironment(value PreReceiveHook_environmentable)()
    SetId(value *int32)()
    SetName(value *string)()
    SetScript(value *string)()
    SetScriptRepository(value PreReceiveHook_script_repositoryable)()
}
