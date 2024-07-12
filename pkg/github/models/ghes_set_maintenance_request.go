package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesSetMaintenanceRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether to enable maintenance mode.
    enabled *bool
    // The list of IP addresses to exclude from maintenance mode. IPv4, IPv6, and CIDR addresses are supported.
    ip_exception_list []string
    // The message to display to users when maintenance mode is enabled.
    maintenance_mode_message *string
    // The UUID of the node to target. This parameter is incompatible with maintenance mode scheduling. Only use `uuid` if the value of `when` is empty or `now`.
    uuid *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The time to enable maintenance mode. If this parameter is empty or set to `now`, maintenance mode is enabled immediately. Otherwise, maintenance mode is enabled at the specified time. The format is ISO 8601.
    when *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewGhesSetMaintenanceRequest instantiates a new GhesSetMaintenanceRequest and sets the default values.
func NewGhesSetMaintenanceRequest()(*GhesSetMaintenanceRequest) {
    m := &GhesSetMaintenanceRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesSetMaintenanceRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesSetMaintenanceRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesSetMaintenanceRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesSetMaintenanceRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnabled gets the enabled property value. Whether to enable maintenance mode.
// returns a *bool when successful
func (m *GhesSetMaintenanceRequest) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesSetMaintenanceRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
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
    res["when"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWhen(val)
        }
        return nil
    }
    return res
}
// GetIpExceptionList gets the ip_exception_list property value. The list of IP addresses to exclude from maintenance mode. IPv4, IPv6, and CIDR addresses are supported.
// returns a []string when successful
func (m *GhesSetMaintenanceRequest) GetIpExceptionList()([]string) {
    return m.ip_exception_list
}
// GetMaintenanceModeMessage gets the maintenance_mode_message property value. The message to display to users when maintenance mode is enabled.
// returns a *string when successful
func (m *GhesSetMaintenanceRequest) GetMaintenanceModeMessage()(*string) {
    return m.maintenance_mode_message
}
// GetUuid gets the uuid property value. The UUID of the node to target. This parameter is incompatible with maintenance mode scheduling. Only use `uuid` if the value of `when` is empty or `now`.
// returns a *UUID when successful
func (m *GhesSetMaintenanceRequest) GetUuid()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.uuid
}
// GetWhen gets the when property value. The time to enable maintenance mode. If this parameter is empty or set to `now`, maintenance mode is enabled immediately. Otherwise, maintenance mode is enabled at the specified time. The format is ISO 8601.
// returns a *Time when successful
func (m *GhesSetMaintenanceRequest) GetWhen()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.when
}
// Serialize serializes information the current object
func (m *GhesSetMaintenanceRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
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
        err := writer.WriteUUIDValue("uuid", m.GetUuid())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("when", m.GetWhen())
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
func (m *GhesSetMaintenanceRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnabled sets the enabled property value. Whether to enable maintenance mode.
func (m *GhesSetMaintenanceRequest) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetIpExceptionList sets the ip_exception_list property value. The list of IP addresses to exclude from maintenance mode. IPv4, IPv6, and CIDR addresses are supported.
func (m *GhesSetMaintenanceRequest) SetIpExceptionList(value []string)() {
    m.ip_exception_list = value
}
// SetMaintenanceModeMessage sets the maintenance_mode_message property value. The message to display to users when maintenance mode is enabled.
func (m *GhesSetMaintenanceRequest) SetMaintenanceModeMessage(value *string)() {
    m.maintenance_mode_message = value
}
// SetUuid sets the uuid property value. The UUID of the node to target. This parameter is incompatible with maintenance mode scheduling. Only use `uuid` if the value of `when` is empty or `now`.
func (m *GhesSetMaintenanceRequest) SetUuid(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.uuid = value
}
// SetWhen sets the when property value. The time to enable maintenance mode. If this parameter is empty or set to `now`, maintenance mode is enabled immediately. Otherwise, maintenance mode is enabled at the specified time. The format is ISO 8601.
func (m *GhesSetMaintenanceRequest) SetWhen(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.when = value
}
type GhesSetMaintenanceRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnabled()(*bool)
    GetIpExceptionList()([]string)
    GetMaintenanceModeMessage()(*string)
    GetUuid()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetWhen()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetEnabled(value *bool)()
    SetIpExceptionList(value []string)()
    SetMaintenanceModeMessage(value *string)()
    SetUuid(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetWhen(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
