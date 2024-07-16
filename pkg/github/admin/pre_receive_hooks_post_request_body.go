package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PreReceiveHooksPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether enforcement can be overridden at the org or repo level. default: `false`
    allow_downstream_configuration *bool
    // The state of enforcement for this hook. default: `disabled`
    enforcement *string
    // The pre-receive environment where the script is executed.
    environment PreReceiveHooksPostRequestBody_environmentable
    // The name of the hook.
    name *string
    // The script that the hook runs.
    script *string
    // The GitHub repository where the script is kept.
    script_repository PreReceiveHooksPostRequestBody_script_repositoryable
}
// NewPreReceiveHooksPostRequestBody instantiates a new PreReceiveHooksPostRequestBody and sets the default values.
func NewPreReceiveHooksPostRequestBody()(*PreReceiveHooksPostRequestBody) {
    m := &PreReceiveHooksPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePreReceiveHooksPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePreReceiveHooksPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPreReceiveHooksPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PreReceiveHooksPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowDownstreamConfiguration gets the allow_downstream_configuration property value. Whether enforcement can be overridden at the org or repo level. default: `false`
// returns a *bool when successful
func (m *PreReceiveHooksPostRequestBody) GetAllowDownstreamConfiguration()(*bool) {
    return m.allow_downstream_configuration
}
// GetEnforcement gets the enforcement property value. The state of enforcement for this hook. default: `disabled`
// returns a *string when successful
func (m *PreReceiveHooksPostRequestBody) GetEnforcement()(*string) {
    return m.enforcement
}
// GetEnvironment gets the environment property value. The pre-receive environment where the script is executed.
// returns a PreReceiveHooksPostRequestBody_environmentable when successful
func (m *PreReceiveHooksPostRequestBody) GetEnvironment()(PreReceiveHooksPostRequestBody_environmentable) {
    return m.environment
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PreReceiveHooksPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(CreatePreReceiveHooksPostRequestBody_environmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnvironment(val.(PreReceiveHooksPostRequestBody_environmentable))
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
        val, err := n.GetObjectValue(CreatePreReceiveHooksPostRequestBody_script_repositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScriptRepository(val.(PreReceiveHooksPostRequestBody_script_repositoryable))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the hook.
// returns a *string when successful
func (m *PreReceiveHooksPostRequestBody) GetName()(*string) {
    return m.name
}
// GetScript gets the script property value. The script that the hook runs.
// returns a *string when successful
func (m *PreReceiveHooksPostRequestBody) GetScript()(*string) {
    return m.script
}
// GetScriptRepository gets the script_repository property value. The GitHub repository where the script is kept.
// returns a PreReceiveHooksPostRequestBody_script_repositoryable when successful
func (m *PreReceiveHooksPostRequestBody) GetScriptRepository()(PreReceiveHooksPostRequestBody_script_repositoryable) {
    return m.script_repository
}
// Serialize serializes information the current object
func (m *PreReceiveHooksPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PreReceiveHooksPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowDownstreamConfiguration sets the allow_downstream_configuration property value. Whether enforcement can be overridden at the org or repo level. default: `false`
func (m *PreReceiveHooksPostRequestBody) SetAllowDownstreamConfiguration(value *bool)() {
    m.allow_downstream_configuration = value
}
// SetEnforcement sets the enforcement property value. The state of enforcement for this hook. default: `disabled`
func (m *PreReceiveHooksPostRequestBody) SetEnforcement(value *string)() {
    m.enforcement = value
}
// SetEnvironment sets the environment property value. The pre-receive environment where the script is executed.
func (m *PreReceiveHooksPostRequestBody) SetEnvironment(value PreReceiveHooksPostRequestBody_environmentable)() {
    m.environment = value
}
// SetName sets the name property value. The name of the hook.
func (m *PreReceiveHooksPostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetScript sets the script property value. The script that the hook runs.
func (m *PreReceiveHooksPostRequestBody) SetScript(value *string)() {
    m.script = value
}
// SetScriptRepository sets the script_repository property value. The GitHub repository where the script is kept.
func (m *PreReceiveHooksPostRequestBody) SetScriptRepository(value PreReceiveHooksPostRequestBody_script_repositoryable)() {
    m.script_repository = value
}
type PreReceiveHooksPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowDownstreamConfiguration()(*bool)
    GetEnforcement()(*string)
    GetEnvironment()(PreReceiveHooksPostRequestBody_environmentable)
    GetName()(*string)
    GetScript()(*string)
    GetScriptRepository()(PreReceiveHooksPostRequestBody_script_repositoryable)
    SetAllowDownstreamConfiguration(value *bool)()
    SetEnforcement(value *string)()
    SetEnvironment(value PreReceiveHooksPostRequestBody_environmentable)()
    SetName(value *string)()
    SetScript(value *string)()
    SetScriptRepository(value PreReceiveHooksPostRequestBody_script_repositoryable)()
}
