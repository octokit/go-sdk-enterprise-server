package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EnterpriseSettings_enterprise_smtp struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The address property
    address *string
    // The authentication property
    authentication *string
    // The discardToNoreplyAddress property
    discardToNoreplyAddress *bool
    // The domain property
    domain *string
    // The enable_starttls_auto property
    enable_starttls_auto *bool
    // The enabled property
    enabled *bool
    // The noreply_address property
    noreply_address *string
    // The password property
    password *string
    // The port property
    port *string
    // The support_address property
    support_address *string
    // The support_address_type property
    support_address_type *string
    // The user_name property
    user_name *string
    // The username property
    username *string
}
// NewEnterpriseSettings_enterprise_smtp instantiates a new EnterpriseSettings_enterprise_smtp and sets the default values.
func NewEnterpriseSettings_enterprise_smtp()(*EnterpriseSettings_enterprise_smtp) {
    m := &EnterpriseSettings_enterprise_smtp{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnterpriseSettings_enterprise_smtpFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnterpriseSettings_enterprise_smtpFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnterpriseSettings_enterprise_smtp(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EnterpriseSettings_enterprise_smtp) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddress gets the address property value. The address property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_smtp) GetAddress()(*string) {
    return m.address
}
// GetAuthentication gets the authentication property value. The authentication property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_smtp) GetAuthentication()(*string) {
    return m.authentication
}
// GetDiscardToNoreplyAddress gets the discard-to-noreply-address property value. The discardToNoreplyAddress property
// returns a *bool when successful
func (m *EnterpriseSettings_enterprise_smtp) GetDiscardToNoreplyAddress()(*bool) {
    return m.discardToNoreplyAddress
}
// GetDomain gets the domain property value. The domain property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_smtp) GetDomain()(*string) {
    return m.domain
}
// GetEnabled gets the enabled property value. The enabled property
// returns a *bool when successful
func (m *EnterpriseSettings_enterprise_smtp) GetEnabled()(*bool) {
    return m.enabled
}
// GetEnableStarttlsAuto gets the enable_starttls_auto property value. The enable_starttls_auto property
// returns a *bool when successful
func (m *EnterpriseSettings_enterprise_smtp) GetEnableStarttlsAuto()(*bool) {
    return m.enable_starttls_auto
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnterpriseSettings_enterprise_smtp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
        }
        return nil
    }
    res["authentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthentication(val)
        }
        return nil
    }
    res["discard-to-noreply-address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscardToNoreplyAddress(val)
        }
        return nil
    }
    res["domain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomain(val)
        }
        return nil
    }
    res["enable_starttls_auto"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableStarttlsAuto(val)
        }
        return nil
    }
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
    res["noreply_address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNoreplyAddress(val)
        }
        return nil
    }
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["port"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPort(val)
        }
        return nil
    }
    res["support_address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportAddress(val)
        }
        return nil
    }
    res["support_address_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportAddressType(val)
        }
        return nil
    }
    res["user_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser_name(val)
        }
        return nil
    }
    res["username"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsername(val)
        }
        return nil
    }
    return res
}
// GetNoreplyAddress gets the noreply_address property value. The noreply_address property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_smtp) GetNoreplyAddress()(*string) {
    return m.noreply_address
}
// GetPassword gets the password property value. The password property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_smtp) GetPassword()(*string) {
    return m.password
}
// GetPort gets the port property value. The port property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_smtp) GetPort()(*string) {
    return m.port
}
// GetSupportAddress gets the support_address property value. The support_address property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_smtp) GetSupportAddress()(*string) {
    return m.support_address
}
// GetSupportAddressType gets the support_address_type property value. The support_address_type property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_smtp) GetSupportAddressType()(*string) {
    return m.support_address_type
}
// GetUser_name gets the user_name property value. The user_name property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_smtp) GetUser_name()(*string) {
    return m.user_name
}
// GetUsername gets the username property value. The username property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_smtp) GetUsername()(*string) {
    return m.username
}
// Serialize serializes information the current object
func (m *EnterpriseSettings_enterprise_smtp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("authentication", m.GetAuthentication())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("discard-to-noreply-address", m.GetDiscardToNoreplyAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("domain", m.GetDomain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enable_starttls_auto", m.GetEnableStarttlsAuto())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("noreply_address", m.GetNoreplyAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("port", m.GetPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("support_address", m.GetSupportAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("support_address_type", m.GetSupportAddressType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("username", m.GetUsername())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("user_name", m.GetUser_name())
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
func (m *EnterpriseSettings_enterprise_smtp) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddress sets the address property value. The address property
func (m *EnterpriseSettings_enterprise_smtp) SetAddress(value *string)() {
    m.address = value
}
// SetAuthentication sets the authentication property value. The authentication property
func (m *EnterpriseSettings_enterprise_smtp) SetAuthentication(value *string)() {
    m.authentication = value
}
// SetDiscardToNoreplyAddress sets the discard-to-noreply-address property value. The discardToNoreplyAddress property
func (m *EnterpriseSettings_enterprise_smtp) SetDiscardToNoreplyAddress(value *bool)() {
    m.discardToNoreplyAddress = value
}
// SetDomain sets the domain property value. The domain property
func (m *EnterpriseSettings_enterprise_smtp) SetDomain(value *string)() {
    m.domain = value
}
// SetEnabled sets the enabled property value. The enabled property
func (m *EnterpriseSettings_enterprise_smtp) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetEnableStarttlsAuto sets the enable_starttls_auto property value. The enable_starttls_auto property
func (m *EnterpriseSettings_enterprise_smtp) SetEnableStarttlsAuto(value *bool)() {
    m.enable_starttls_auto = value
}
// SetNoreplyAddress sets the noreply_address property value. The noreply_address property
func (m *EnterpriseSettings_enterprise_smtp) SetNoreplyAddress(value *string)() {
    m.noreply_address = value
}
// SetPassword sets the password property value. The password property
func (m *EnterpriseSettings_enterprise_smtp) SetPassword(value *string)() {
    m.password = value
}
// SetPort sets the port property value. The port property
func (m *EnterpriseSettings_enterprise_smtp) SetPort(value *string)() {
    m.port = value
}
// SetSupportAddress sets the support_address property value. The support_address property
func (m *EnterpriseSettings_enterprise_smtp) SetSupportAddress(value *string)() {
    m.support_address = value
}
// SetSupportAddressType sets the support_address_type property value. The support_address_type property
func (m *EnterpriseSettings_enterprise_smtp) SetSupportAddressType(value *string)() {
    m.support_address_type = value
}
// SetUser_name sets the user_name property value. The user_name property
func (m *EnterpriseSettings_enterprise_smtp) SetUser_name(value *string)() {
    m.user_name = value
}
// SetUsername sets the username property value. The username property
func (m *EnterpriseSettings_enterprise_smtp) SetUsername(value *string)() {
    m.username = value
}
type EnterpriseSettings_enterprise_smtpable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(*string)
    GetAuthentication()(*string)
    GetDiscardToNoreplyAddress()(*bool)
    GetDomain()(*string)
    GetEnabled()(*bool)
    GetEnableStarttlsAuto()(*bool)
    GetNoreplyAddress()(*string)
    GetPassword()(*string)
    GetPort()(*string)
    GetSupportAddress()(*string)
    GetSupportAddressType()(*string)
    GetUser_name()(*string)
    GetUsername()(*string)
    SetAddress(value *string)()
    SetAuthentication(value *string)()
    SetDiscardToNoreplyAddress(value *bool)()
    SetDomain(value *string)()
    SetEnabled(value *bool)()
    SetEnableStarttlsAuto(value *bool)()
    SetNoreplyAddress(value *string)()
    SetPassword(value *string)()
    SetPort(value *string)()
    SetSupportAddress(value *string)()
    SetSupportAddressType(value *string)()
    SetUser_name(value *string)()
    SetUsername(value *string)()
}
