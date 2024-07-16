package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LdapMappingTeam struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The description property
    description *string
    // The html_url property
    html_url *string
    // The id property
    id *int32
    // The ldap_dn property
    ldap_dn *string
    // The members_url property
    members_url *string
    // The name property
    name *string
    // The node_id property
    node_id *string
    // The notification_setting property
    notification_setting *string
    // The parent property
    parent i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
    // The permission property
    permission *string
    // The privacy property
    privacy *string
    // The repositories_url property
    repositories_url *string
    // The slug property
    slug *string
    // The url property
    url *string
}
// NewLdapMappingTeam instantiates a new LdapMappingTeam and sets the default values.
func NewLdapMappingTeam()(*LdapMappingTeam) {
    m := &LdapMappingTeam{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLdapMappingTeamFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLdapMappingTeamFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLdapMappingTeam(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LdapMappingTeam) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *LdapMappingTeam) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LdapMappingTeam) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtmlUrl(val)
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
    res["ldap_dn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLdapDn(val)
        }
        return nil
    }
    res["members_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembersUrl(val)
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
    res["node_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeId(val)
        }
        return nil
    }
    res["notification_setting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationSetting(val)
        }
        return nil
    }
    res["parent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParent(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
        }
        return nil
    }
    res["permission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermission(val)
        }
        return nil
    }
    res["privacy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacy(val)
        }
        return nil
    }
    res["repositories_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoriesUrl(val)
        }
        return nil
    }
    res["slug"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSlug(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *LdapMappingTeam) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *LdapMappingTeam) GetId()(*int32) {
    return m.id
}
// GetLdapDn gets the ldap_dn property value. The ldap_dn property
// returns a *string when successful
func (m *LdapMappingTeam) GetLdapDn()(*string) {
    return m.ldap_dn
}
// GetMembersUrl gets the members_url property value. The members_url property
// returns a *string when successful
func (m *LdapMappingTeam) GetMembersUrl()(*string) {
    return m.members_url
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *LdapMappingTeam) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *LdapMappingTeam) GetNodeId()(*string) {
    return m.node_id
}
// GetNotificationSetting gets the notification_setting property value. The notification_setting property
// returns a *string when successful
func (m *LdapMappingTeam) GetNotificationSetting()(*string) {
    return m.notification_setting
}
// GetParent gets the parent property value. The parent property
// returns a UntypedNodeable when successful
func (m *LdapMappingTeam) GetParent()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.parent
}
// GetPermission gets the permission property value. The permission property
// returns a *string when successful
func (m *LdapMappingTeam) GetPermission()(*string) {
    return m.permission
}
// GetPrivacy gets the privacy property value. The privacy property
// returns a *string when successful
func (m *LdapMappingTeam) GetPrivacy()(*string) {
    return m.privacy
}
// GetRepositoriesUrl gets the repositories_url property value. The repositories_url property
// returns a *string when successful
func (m *LdapMappingTeam) GetRepositoriesUrl()(*string) {
    return m.repositories_url
}
// GetSlug gets the slug property value. The slug property
// returns a *string when successful
func (m *LdapMappingTeam) GetSlug()(*string) {
    return m.slug
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *LdapMappingTeam) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *LdapMappingTeam) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
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
        err := writer.WriteStringValue("ldap_dn", m.GetLdapDn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("members_url", m.GetMembersUrl())
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
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("notification_setting", m.GetNotificationSetting())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("parent", m.GetParent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("permission", m.GetPermission())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("privacy", m.GetPrivacy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repositories_url", m.GetRepositoriesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("slug", m.GetSlug())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *LdapMappingTeam) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description property
func (m *LdapMappingTeam) SetDescription(value *string)() {
    m.description = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *LdapMappingTeam) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *LdapMappingTeam) SetId(value *int32)() {
    m.id = value
}
// SetLdapDn sets the ldap_dn property value. The ldap_dn property
func (m *LdapMappingTeam) SetLdapDn(value *string)() {
    m.ldap_dn = value
}
// SetMembersUrl sets the members_url property value. The members_url property
func (m *LdapMappingTeam) SetMembersUrl(value *string)() {
    m.members_url = value
}
// SetName sets the name property value. The name property
func (m *LdapMappingTeam) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *LdapMappingTeam) SetNodeId(value *string)() {
    m.node_id = value
}
// SetNotificationSetting sets the notification_setting property value. The notification_setting property
func (m *LdapMappingTeam) SetNotificationSetting(value *string)() {
    m.notification_setting = value
}
// SetParent sets the parent property value. The parent property
func (m *LdapMappingTeam) SetParent(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.parent = value
}
// SetPermission sets the permission property value. The permission property
func (m *LdapMappingTeam) SetPermission(value *string)() {
    m.permission = value
}
// SetPrivacy sets the privacy property value. The privacy property
func (m *LdapMappingTeam) SetPrivacy(value *string)() {
    m.privacy = value
}
// SetRepositoriesUrl sets the repositories_url property value. The repositories_url property
func (m *LdapMappingTeam) SetRepositoriesUrl(value *string)() {
    m.repositories_url = value
}
// SetSlug sets the slug property value. The slug property
func (m *LdapMappingTeam) SetSlug(value *string)() {
    m.slug = value
}
// SetUrl sets the url property value. The url property
func (m *LdapMappingTeam) SetUrl(value *string)() {
    m.url = value
}
type LdapMappingTeamable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetLdapDn()(*string)
    GetMembersUrl()(*string)
    GetName()(*string)
    GetNodeId()(*string)
    GetNotificationSetting()(*string)
    GetParent()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    GetPermission()(*string)
    GetPrivacy()(*string)
    GetRepositoriesUrl()(*string)
    GetSlug()(*string)
    GetUrl()(*string)
    SetDescription(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetLdapDn(value *string)()
    SetMembersUrl(value *string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetNotificationSetting(value *string)()
    SetParent(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
    SetPermission(value *string)()
    SetPrivacy(value *string)()
    SetRepositoriesUrl(value *string)()
    SetSlug(value *string)()
    SetUrl(value *string)()
}
