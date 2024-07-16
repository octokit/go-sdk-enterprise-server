package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MaintenanceStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The connection_services property
    connection_services []MaintenanceStatus_connection_servicesable
    // The scheduled_time property
    scheduled_time *string
    // The status property
    status *string
}
// NewMaintenanceStatus instantiates a new MaintenanceStatus and sets the default values.
func NewMaintenanceStatus()(*MaintenanceStatus) {
    m := &MaintenanceStatus{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMaintenanceStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMaintenanceStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMaintenanceStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MaintenanceStatus) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConnectionServices gets the connection_services property value. The connection_services property
// returns a []MaintenanceStatus_connection_servicesable when successful
func (m *MaintenanceStatus) GetConnectionServices()([]MaintenanceStatus_connection_servicesable) {
    return m.connection_services
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MaintenanceStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["connection_services"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMaintenanceStatus_connection_servicesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MaintenanceStatus_connection_servicesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MaintenanceStatus_connection_servicesable)
                }
            }
            m.SetConnectionServices(res)
        }
        return nil
    }
    res["scheduled_time"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduledTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
// GetScheduledTime gets the scheduled_time property value. The scheduled_time property
// returns a *string when successful
func (m *MaintenanceStatus) GetScheduledTime()(*string) {
    return m.scheduled_time
}
// GetStatus gets the status property value. The status property
// returns a *string when successful
func (m *MaintenanceStatus) GetStatus()(*string) {
    return m.status
}
// Serialize serializes information the current object
func (m *MaintenanceStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetConnectionServices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConnectionServices()))
        for i, v := range m.GetConnectionServices() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("connection_services", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("scheduled_time", m.GetScheduledTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("status", m.GetStatus())
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
func (m *MaintenanceStatus) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConnectionServices sets the connection_services property value. The connection_services property
func (m *MaintenanceStatus) SetConnectionServices(value []MaintenanceStatus_connection_servicesable)() {
    m.connection_services = value
}
// SetScheduledTime sets the scheduled_time property value. The scheduled_time property
func (m *MaintenanceStatus) SetScheduledTime(value *string)() {
    m.scheduled_time = value
}
// SetStatus sets the status property value. The status property
func (m *MaintenanceStatus) SetStatus(value *string)() {
    m.status = value
}
type MaintenanceStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectionServices()([]MaintenanceStatus_connection_servicesable)
    GetScheduledTime()(*string)
    GetStatus()(*string)
    SetConnectionServices(value []MaintenanceStatus_connection_servicesable)()
    SetScheduledTime(value *string)()
    SetStatus(value *string)()
}
