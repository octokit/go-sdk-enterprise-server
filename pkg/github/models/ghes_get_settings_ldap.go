package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesGetSettings_ldap struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The admin_group property
    admin_group *string
    // The base property
    base []string
    // The bind_dn property
    bind_dn *string
    // The host property
    host *string
    // The method property
    method *string
    // The password property
    password *string
    // The port property
    port *int32
    // The posix_support property
    posix_support *bool
    // The profile property
    profile GhesGetSettings_ldap_profileable
    // The reconciliation property
    reconciliation GhesGetSettings_ldap_reconciliationable
    // The recursive_group_search property
    recursive_group_search *bool
    // The search_strategy property
    search_strategy *string
    // The sync_enabled property
    sync_enabled *bool
    // The team_sync_interval property
    team_sync_interval *int32
    // The uid property
    uid *string
    // The user_groups property
    user_groups []string
    // The user_sync_emails property
    user_sync_emails *bool
    // The user_sync_interval property
    user_sync_interval *int32
    // The user_sync_keys property
    user_sync_keys *bool
    // The virtual_attribute_enabled property
    virtual_attribute_enabled *bool
}
// NewGhesGetSettings_ldap instantiates a new GhesGetSettings_ldap and sets the default values.
func NewGhesGetSettings_ldap()(*GhesGetSettings_ldap) {
    m := &GhesGetSettings_ldap{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesGetSettings_ldapFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesGetSettings_ldapFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesGetSettings_ldap(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesGetSettings_ldap) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdminGroup gets the admin_group property value. The admin_group property
// returns a *string when successful
func (m *GhesGetSettings_ldap) GetAdminGroup()(*string) {
    return m.admin_group
}
// GetBase gets the base property value. The base property
// returns a []string when successful
func (m *GhesGetSettings_ldap) GetBase()([]string) {
    return m.base
}
// GetBindDn gets the bind_dn property value. The bind_dn property
// returns a *string when successful
func (m *GhesGetSettings_ldap) GetBindDn()(*string) {
    return m.bind_dn
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesGetSettings_ldap) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["admin_group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminGroup(val)
        }
        return nil
    }
    res["base"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetBase(res)
        }
        return nil
    }
    res["bind_dn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBindDn(val)
        }
        return nil
    }
    res["host"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHost(val)
        }
        return nil
    }
    res["method"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMethod(val)
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
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPort(val)
        }
        return nil
    }
    res["posix_support"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPosixSupport(val)
        }
        return nil
    }
    res["profile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGhesGetSettings_ldap_profileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfile(val.(GhesGetSettings_ldap_profileable))
        }
        return nil
    }
    res["reconciliation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGhesGetSettings_ldap_reconciliationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReconciliation(val.(GhesGetSettings_ldap_reconciliationable))
        }
        return nil
    }
    res["recursive_group_search"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecursiveGroupSearch(val)
        }
        return nil
    }
    res["search_strategy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearchStrategy(val)
        }
        return nil
    }
    res["sync_enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyncEnabled(val)
        }
        return nil
    }
    res["team_sync_interval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamSyncInterval(val)
        }
        return nil
    }
    res["uid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUid(val)
        }
        return nil
    }
    res["user_groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetUserGroups(res)
        }
        return nil
    }
    res["user_sync_emails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserSyncEmails(val)
        }
        return nil
    }
    res["user_sync_interval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserSyncInterval(val)
        }
        return nil
    }
    res["user_sync_keys"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserSyncKeys(val)
        }
        return nil
    }
    res["virtual_attribute_enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualAttributeEnabled(val)
        }
        return nil
    }
    return res
}
// GetHost gets the host property value. The host property
// returns a *string when successful
func (m *GhesGetSettings_ldap) GetHost()(*string) {
    return m.host
}
// GetMethod gets the method property value. The method property
// returns a *string when successful
func (m *GhesGetSettings_ldap) GetMethod()(*string) {
    return m.method
}
// GetPassword gets the password property value. The password property
// returns a *string when successful
func (m *GhesGetSettings_ldap) GetPassword()(*string) {
    return m.password
}
// GetPort gets the port property value. The port property
// returns a *int32 when successful
func (m *GhesGetSettings_ldap) GetPort()(*int32) {
    return m.port
}
// GetPosixSupport gets the posix_support property value. The posix_support property
// returns a *bool when successful
func (m *GhesGetSettings_ldap) GetPosixSupport()(*bool) {
    return m.posix_support
}
// GetProfile gets the profile property value. The profile property
// returns a GhesGetSettings_ldap_profileable when successful
func (m *GhesGetSettings_ldap) GetProfile()(GhesGetSettings_ldap_profileable) {
    return m.profile
}
// GetReconciliation gets the reconciliation property value. The reconciliation property
// returns a GhesGetSettings_ldap_reconciliationable when successful
func (m *GhesGetSettings_ldap) GetReconciliation()(GhesGetSettings_ldap_reconciliationable) {
    return m.reconciliation
}
// GetRecursiveGroupSearch gets the recursive_group_search property value. The recursive_group_search property
// returns a *bool when successful
func (m *GhesGetSettings_ldap) GetRecursiveGroupSearch()(*bool) {
    return m.recursive_group_search
}
// GetSearchStrategy gets the search_strategy property value. The search_strategy property
// returns a *string when successful
func (m *GhesGetSettings_ldap) GetSearchStrategy()(*string) {
    return m.search_strategy
}
// GetSyncEnabled gets the sync_enabled property value. The sync_enabled property
// returns a *bool when successful
func (m *GhesGetSettings_ldap) GetSyncEnabled()(*bool) {
    return m.sync_enabled
}
// GetTeamSyncInterval gets the team_sync_interval property value. The team_sync_interval property
// returns a *int32 when successful
func (m *GhesGetSettings_ldap) GetTeamSyncInterval()(*int32) {
    return m.team_sync_interval
}
// GetUid gets the uid property value. The uid property
// returns a *string when successful
func (m *GhesGetSettings_ldap) GetUid()(*string) {
    return m.uid
}
// GetUserGroups gets the user_groups property value. The user_groups property
// returns a []string when successful
func (m *GhesGetSettings_ldap) GetUserGroups()([]string) {
    return m.user_groups
}
// GetUserSyncEmails gets the user_sync_emails property value. The user_sync_emails property
// returns a *bool when successful
func (m *GhesGetSettings_ldap) GetUserSyncEmails()(*bool) {
    return m.user_sync_emails
}
// GetUserSyncInterval gets the user_sync_interval property value. The user_sync_interval property
// returns a *int32 when successful
func (m *GhesGetSettings_ldap) GetUserSyncInterval()(*int32) {
    return m.user_sync_interval
}
// GetUserSyncKeys gets the user_sync_keys property value. The user_sync_keys property
// returns a *bool when successful
func (m *GhesGetSettings_ldap) GetUserSyncKeys()(*bool) {
    return m.user_sync_keys
}
// GetVirtualAttributeEnabled gets the virtual_attribute_enabled property value. The virtual_attribute_enabled property
// returns a *bool when successful
func (m *GhesGetSettings_ldap) GetVirtualAttributeEnabled()(*bool) {
    return m.virtual_attribute_enabled
}
// Serialize serializes information the current object
func (m *GhesGetSettings_ldap) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("admin_group", m.GetAdminGroup())
        if err != nil {
            return err
        }
    }
    if m.GetBase() != nil {
        err := writer.WriteCollectionOfStringValues("base", m.GetBase())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bind_dn", m.GetBindDn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("host", m.GetHost())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("method", m.GetMethod())
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
        err := writer.WriteInt32Value("port", m.GetPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("posix_support", m.GetPosixSupport())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("profile", m.GetProfile())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("reconciliation", m.GetReconciliation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("recursive_group_search", m.GetRecursiveGroupSearch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("search_strategy", m.GetSearchStrategy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sync_enabled", m.GetSyncEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("team_sync_interval", m.GetTeamSyncInterval())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("uid", m.GetUid())
        if err != nil {
            return err
        }
    }
    if m.GetUserGroups() != nil {
        err := writer.WriteCollectionOfStringValues("user_groups", m.GetUserGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("user_sync_emails", m.GetUserSyncEmails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("user_sync_interval", m.GetUserSyncInterval())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("user_sync_keys", m.GetUserSyncKeys())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("virtual_attribute_enabled", m.GetVirtualAttributeEnabled())
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
func (m *GhesGetSettings_ldap) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdminGroup sets the admin_group property value. The admin_group property
func (m *GhesGetSettings_ldap) SetAdminGroup(value *string)() {
    m.admin_group = value
}
// SetBase sets the base property value. The base property
func (m *GhesGetSettings_ldap) SetBase(value []string)() {
    m.base = value
}
// SetBindDn sets the bind_dn property value. The bind_dn property
func (m *GhesGetSettings_ldap) SetBindDn(value *string)() {
    m.bind_dn = value
}
// SetHost sets the host property value. The host property
func (m *GhesGetSettings_ldap) SetHost(value *string)() {
    m.host = value
}
// SetMethod sets the method property value. The method property
func (m *GhesGetSettings_ldap) SetMethod(value *string)() {
    m.method = value
}
// SetPassword sets the password property value. The password property
func (m *GhesGetSettings_ldap) SetPassword(value *string)() {
    m.password = value
}
// SetPort sets the port property value. The port property
func (m *GhesGetSettings_ldap) SetPort(value *int32)() {
    m.port = value
}
// SetPosixSupport sets the posix_support property value. The posix_support property
func (m *GhesGetSettings_ldap) SetPosixSupport(value *bool)() {
    m.posix_support = value
}
// SetProfile sets the profile property value. The profile property
func (m *GhesGetSettings_ldap) SetProfile(value GhesGetSettings_ldap_profileable)() {
    m.profile = value
}
// SetReconciliation sets the reconciliation property value. The reconciliation property
func (m *GhesGetSettings_ldap) SetReconciliation(value GhesGetSettings_ldap_reconciliationable)() {
    m.reconciliation = value
}
// SetRecursiveGroupSearch sets the recursive_group_search property value. The recursive_group_search property
func (m *GhesGetSettings_ldap) SetRecursiveGroupSearch(value *bool)() {
    m.recursive_group_search = value
}
// SetSearchStrategy sets the search_strategy property value. The search_strategy property
func (m *GhesGetSettings_ldap) SetSearchStrategy(value *string)() {
    m.search_strategy = value
}
// SetSyncEnabled sets the sync_enabled property value. The sync_enabled property
func (m *GhesGetSettings_ldap) SetSyncEnabled(value *bool)() {
    m.sync_enabled = value
}
// SetTeamSyncInterval sets the team_sync_interval property value. The team_sync_interval property
func (m *GhesGetSettings_ldap) SetTeamSyncInterval(value *int32)() {
    m.team_sync_interval = value
}
// SetUid sets the uid property value. The uid property
func (m *GhesGetSettings_ldap) SetUid(value *string)() {
    m.uid = value
}
// SetUserGroups sets the user_groups property value. The user_groups property
func (m *GhesGetSettings_ldap) SetUserGroups(value []string)() {
    m.user_groups = value
}
// SetUserSyncEmails sets the user_sync_emails property value. The user_sync_emails property
func (m *GhesGetSettings_ldap) SetUserSyncEmails(value *bool)() {
    m.user_sync_emails = value
}
// SetUserSyncInterval sets the user_sync_interval property value. The user_sync_interval property
func (m *GhesGetSettings_ldap) SetUserSyncInterval(value *int32)() {
    m.user_sync_interval = value
}
// SetUserSyncKeys sets the user_sync_keys property value. The user_sync_keys property
func (m *GhesGetSettings_ldap) SetUserSyncKeys(value *bool)() {
    m.user_sync_keys = value
}
// SetVirtualAttributeEnabled sets the virtual_attribute_enabled property value. The virtual_attribute_enabled property
func (m *GhesGetSettings_ldap) SetVirtualAttributeEnabled(value *bool)() {
    m.virtual_attribute_enabled = value
}
type GhesGetSettings_ldapable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdminGroup()(*string)
    GetBase()([]string)
    GetBindDn()(*string)
    GetHost()(*string)
    GetMethod()(*string)
    GetPassword()(*string)
    GetPort()(*int32)
    GetPosixSupport()(*bool)
    GetProfile()(GhesGetSettings_ldap_profileable)
    GetReconciliation()(GhesGetSettings_ldap_reconciliationable)
    GetRecursiveGroupSearch()(*bool)
    GetSearchStrategy()(*string)
    GetSyncEnabled()(*bool)
    GetTeamSyncInterval()(*int32)
    GetUid()(*string)
    GetUserGroups()([]string)
    GetUserSyncEmails()(*bool)
    GetUserSyncInterval()(*int32)
    GetUserSyncKeys()(*bool)
    GetVirtualAttributeEnabled()(*bool)
    SetAdminGroup(value *string)()
    SetBase(value []string)()
    SetBindDn(value *string)()
    SetHost(value *string)()
    SetMethod(value *string)()
    SetPassword(value *string)()
    SetPort(value *int32)()
    SetPosixSupport(value *bool)()
    SetProfile(value GhesGetSettings_ldap_profileable)()
    SetReconciliation(value GhesGetSettings_ldap_reconciliationable)()
    SetRecursiveGroupSearch(value *bool)()
    SetSearchStrategy(value *string)()
    SetSyncEnabled(value *bool)()
    SetTeamSyncInterval(value *int32)()
    SetUid(value *string)()
    SetUserGroups(value []string)()
    SetUserSyncEmails(value *bool)()
    SetUserSyncInterval(value *int32)()
    SetUserSyncKeys(value *bool)()
    SetVirtualAttributeEnabled(value *bool)()
}
