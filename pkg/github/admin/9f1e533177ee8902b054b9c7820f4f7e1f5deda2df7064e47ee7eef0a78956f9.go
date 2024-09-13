package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errors struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The code property
    code *string
    // The message property
    message *string
    // The resource property
    resource *string
}
// NewPreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errors instantiates a new PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errors and sets the default values.
func NewPreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errors()(*PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errors) {
    m := &PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errors{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errorsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errorsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errors(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errors) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCode gets the code property value. The code property
// returns a *string when successful
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errors) GetCode()(*string) {
    return m.code
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errors) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
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
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errors) GetMessage()(*string) {
    return m.message
}
// GetResource gets the resource property value. The resource property
// returns a *string when successful
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errors) GetResource()(*string) {
    return m.resource
}
// Serialize serializes information the current object
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errors) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
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
        err := writer.WriteStringValue("resource", m.GetResource())
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
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errors) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCode sets the code property value. The code property
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errors) SetCode(value *string)() {
    m.code = value
}
// SetMessage sets the message property value. The message property
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errors) SetMessage(value *string)() {
    m.message = value
}
// SetResource sets the resource property value. The resource property
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errors) SetResource(value *string)() {
    m.resource = value
}
type PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errorsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*string)
    GetMessage()(*string)
    GetResource()(*string)
    SetCode(value *string)()
    SetMessage(value *string)()
    SetResource(value *string)()
}
