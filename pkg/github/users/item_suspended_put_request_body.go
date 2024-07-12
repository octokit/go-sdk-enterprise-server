package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemSuspendedPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The reason the user is being suspended. This message will be logged in the [audit log](https://docs.github.com/enterprise-server@3.10/admin/monitoring-activity-in-your-enterprise/reviewing-audit-logs-for-your-enterprise/about-the-audit-log-for-your-enterprise). If you don't provide a `reason`, it will default to "Suspended via API by _SITE\_ADMINISTRATOR_", where _SITE\_ADMINISTRATOR_ is the person who performed the action.
    reason *string
}
// NewItemSuspendedPutRequestBody instantiates a new ItemSuspendedPutRequestBody and sets the default values.
func NewItemSuspendedPutRequestBody()(*ItemSuspendedPutRequestBody) {
    m := &ItemSuspendedPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemSuspendedPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSuspendedPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSuspendedPutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemSuspendedPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemSuspendedPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val)
        }
        return nil
    }
    return res
}
// GetReason gets the reason property value. The reason the user is being suspended. This message will be logged in the [audit log](https://docs.github.com/enterprise-server@3.10/admin/monitoring-activity-in-your-enterprise/reviewing-audit-logs-for-your-enterprise/about-the-audit-log-for-your-enterprise). If you don't provide a `reason`, it will default to "Suspended via API by _SITE\_ADMINISTRATOR_", where _SITE\_ADMINISTRATOR_ is the person who performed the action.
// returns a *string when successful
func (m *ItemSuspendedPutRequestBody) GetReason()(*string) {
    return m.reason
}
// Serialize serializes information the current object
func (m *ItemSuspendedPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("reason", m.GetReason())
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
func (m *ItemSuspendedPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetReason sets the reason property value. The reason the user is being suspended. This message will be logged in the [audit log](https://docs.github.com/enterprise-server@3.10/admin/monitoring-activity-in-your-enterprise/reviewing-audit-logs-for-your-enterprise/about-the-audit-log-for-your-enterprise). If you don't provide a `reason`, it will default to "Suspended via API by _SITE\_ADMINISTRATOR_", where _SITE\_ADMINISTRATOR_ is the person who performed the action.
func (m *ItemSuspendedPutRequestBody) SetReason(value *string)() {
    m.reason = value
}
type ItemSuspendedPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReason()(*string)
    SetReason(value *string)()
}
