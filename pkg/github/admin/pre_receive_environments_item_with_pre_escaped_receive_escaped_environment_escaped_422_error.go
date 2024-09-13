package admin

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PreReceiveEnvironmentsItemWithPre_receive_environment_422Error struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ApiError
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The errors property
    errors []PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errorsable
    // The message property
    message *string
}
// NewPreReceiveEnvironmentsItemWithPre_receive_environment_422Error instantiates a new PreReceiveEnvironmentsItemWithPre_receive_environment_422Error and sets the default values.
func NewPreReceiveEnvironmentsItemWithPre_receive_environment_422Error()(*PreReceiveEnvironmentsItemWithPre_receive_environment_422Error) {
    m := &PreReceiveEnvironmentsItemWithPre_receive_environment_422Error{
        ApiError: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewApiError(),
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePreReceiveEnvironmentsItemWithPre_receive_environment_422ErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePreReceiveEnvironmentsItemWithPre_receive_environment_422ErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPreReceiveEnvironmentsItemWithPre_receive_environment_422Error(), nil
}
// Error the primary error message.
// returns a string when successful
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_422Error) Error()(string) {
    return m.ApiError.Error()
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_422Error) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetErrors gets the errors property value. The errors property
// returns a []PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errorsable when successful
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_422Error) GetErrors()([]PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errorsable) {
    return m.errors
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_422Error) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["errors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errorsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errorsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errorsable)
                }
            }
            m.SetErrors(res)
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
    return res
}
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_422Error) GetMessage()(*string) {
    return m.message
}
// Serialize serializes information the current object
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_422Error) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetErrors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetErrors()))
        for i, v := range m.GetErrors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("errors", cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_422Error) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetErrors sets the errors property value. The errors property
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_422Error) SetErrors(value []PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errorsable)() {
    m.errors = value
}
// SetMessage sets the message property value. The message property
func (m *PreReceiveEnvironmentsItemWithPre_receive_environment_422Error) SetMessage(value *string)() {
    m.message = value
}
type PreReceiveEnvironmentsItemWithPre_receive_environment_422Errorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetErrors()([]PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errorsable)
    GetMessage()(*string)
    SetErrors(value []PreReceiveEnvironmentsItemWithPre_receive_environment_422Error_errorsable)()
    SetMessage(value *string)()
}
