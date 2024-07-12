package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesGetMaintenance struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The can_unset_maintenance property
    can_unset_maintenance *bool
    // The connection_services property
    connection_services []GhesGetMaintenance_connection_servicesable
    // The hostname property
    hostname *string
    // The ip_exception_list property
    ip_exception_list []string
    // The maintenance_mode_message property
    maintenance_mode_message *string
    // The scheduled_time property
    scheduled_time *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The status property
    status *GhesGetMaintenance_status
    // The uuid property
    uuid *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
}
// NewGhesGetMaintenance instantiates a new GhesGetMaintenance and sets the default values.
func NewGhesGetMaintenance()(*GhesGetMaintenance) {
    m := &GhesGetMaintenance{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesGetMaintenanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesGetMaintenanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesGetMaintenance(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesGetMaintenance) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCanUnsetMaintenance gets the can_unset_maintenance property value. The can_unset_maintenance property
// returns a *bool when successful
func (m *GhesGetMaintenance) GetCanUnsetMaintenance()(*bool) {
    return m.can_unset_maintenance
}
// GetConnectionServices gets the connection_services property value. The connection_services property
// returns a []GhesGetMaintenance_connection_servicesable when successful
func (m *GhesGetMaintenance) GetConnectionServices()([]GhesGetMaintenance_connection_servicesable) {
    return m.connection_services
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesGetMaintenance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["can_unset_maintenance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanUnsetMaintenance(val)
        }
        return nil
    }
    res["connection_services"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGhesGetMaintenance_connection_servicesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GhesGetMaintenance_connection_servicesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GhesGetMaintenance_connection_servicesable)
                }
            }
            m.SetConnectionServices(res)
        }
        return nil
    }
    res["hostname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostname(val)
        }
        return nil
    }
    res["ip_exception_list"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetIpExceptionList(res)
        }
        return nil
    }
    res["maintenance_mode_message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaintenanceModeMessage(val)
        }
        return nil
    }
    res["scheduled_time"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduledTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGhesGetMaintenance_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*GhesGetMaintenance_status))
        }
        return nil
    }
    res["uuid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUuid(val)
        }
        return nil
    }
    return res
}
// GetHostname gets the hostname property value. The hostname property
// returns a *string when successful
func (m *GhesGetMaintenance) GetHostname()(*string) {
    return m.hostname
}
// GetIpExceptionList gets the ip_exception_list property value. The ip_exception_list property
// returns a []string when successful
func (m *GhesGetMaintenance) GetIpExceptionList()([]string) {
    return m.ip_exception_list
}
// GetMaintenanceModeMessage gets the maintenance_mode_message property value. The maintenance_mode_message property
// returns a *string when successful
func (m *GhesGetMaintenance) GetMaintenanceModeMessage()(*string) {
    return m.maintenance_mode_message
}
// GetScheduledTime gets the scheduled_time property value. The scheduled_time property
// returns a *DateOnly when successful
func (m *GhesGetMaintenance) GetScheduledTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.scheduled_time
}
// GetStatus gets the status property value. The status property
// returns a *GhesGetMaintenance_status when successful
func (m *GhesGetMaintenance) GetStatus()(*GhesGetMaintenance_status) {
    return m.status
}
// GetUuid gets the uuid property value. The uuid property
// returns a *UUID when successful
func (m *GhesGetMaintenance) GetUuid()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.uuid
}
// Serialize serializes information the current object
func (m *GhesGetMaintenance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("can_unset_maintenance", m.GetCanUnsetMaintenance())
        if err != nil {
            return err
        }
    }
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
        err := writer.WriteStringValue("hostname", m.GetHostname())
        if err != nil {
            return err
        }
    }
    if m.GetIpExceptionList() != nil {
        err := writer.WriteCollectionOfStringValues("ip_exception_list", m.GetIpExceptionList())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("maintenance_mode_message", m.GetMaintenanceModeMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteDateOnlyValue("scheduled_time", m.GetScheduledTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("uuid", m.GetUuid())
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
func (m *GhesGetMaintenance) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCanUnsetMaintenance sets the can_unset_maintenance property value. The can_unset_maintenance property
func (m *GhesGetMaintenance) SetCanUnsetMaintenance(value *bool)() {
    m.can_unset_maintenance = value
}
// SetConnectionServices sets the connection_services property value. The connection_services property
func (m *GhesGetMaintenance) SetConnectionServices(value []GhesGetMaintenance_connection_servicesable)() {
    m.connection_services = value
}
// SetHostname sets the hostname property value. The hostname property
func (m *GhesGetMaintenance) SetHostname(value *string)() {
    m.hostname = value
}
// SetIpExceptionList sets the ip_exception_list property value. The ip_exception_list property
func (m *GhesGetMaintenance) SetIpExceptionList(value []string)() {
    m.ip_exception_list = value
}
// SetMaintenanceModeMessage sets the maintenance_mode_message property value. The maintenance_mode_message property
func (m *GhesGetMaintenance) SetMaintenanceModeMessage(value *string)() {
    m.maintenance_mode_message = value
}
// SetScheduledTime sets the scheduled_time property value. The scheduled_time property
func (m *GhesGetMaintenance) SetScheduledTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.scheduled_time = value
}
// SetStatus sets the status property value. The status property
func (m *GhesGetMaintenance) SetStatus(value *GhesGetMaintenance_status)() {
    m.status = value
}
// SetUuid sets the uuid property value. The uuid property
func (m *GhesGetMaintenance) SetUuid(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.uuid = value
}
type GhesGetMaintenanceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCanUnsetMaintenance()(*bool)
    GetConnectionServices()([]GhesGetMaintenance_connection_servicesable)
    GetHostname()(*string)
    GetIpExceptionList()([]string)
    GetMaintenanceModeMessage()(*string)
    GetScheduledTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetStatus()(*GhesGetMaintenance_status)
    GetUuid()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    SetCanUnsetMaintenance(value *bool)()
    SetConnectionServices(value []GhesGetMaintenance_connection_servicesable)()
    SetHostname(value *string)()
    SetIpExceptionList(value []string)()
    SetMaintenanceModeMessage(value *string)()
    SetScheduledTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetStatus(value *GhesGetMaintenance_status)()
    SetUuid(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
}
