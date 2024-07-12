package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesGetSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The admin_password property
    admin_password *string
    // The assets property
    assets *string
    // The auth_mode property
    auth_mode *string
    // The avatar property
    avatar GhesGetSettings_avatarable
    // The cas property
    cas GhesGetSettings_casable
    // The collectd property
    collectd GhesGetSettings_collectdable
    // The configuration_id property
    configuration_id *int32
    // The configuration_run_count property
    configuration_run_count *int32
    // The customer property
    customer GhesGetSettings_customerable
    // The expire_sessions property
    expire_sessions *bool
    // The github_hostname property
    github_hostname *string
    // The github_oauth property
    github_oauth GhesGetSettings_github_oauthable
    // The github_ssl property
    github_ssl GhesGetSettings_github_sslable
    // The http_proxy property
    http_proxy *string
    // The identicons_host property
    identicons_host *string
    // The ldap property
    ldap GhesGetSettings_ldapable
    // The license property
    license GhesGetSettings_licenseable
    // The load_balancer property
    load_balancer *string
    // The mapping property
    mapping GhesGetSettings_mappingable
    // The ntp property
    ntp GhesGetSettings_ntpable
    // The pages property
    pages GhesGetSettings_pagesable
    // The private_mode property
    private_mode *bool
    // The public_pages property
    public_pages *bool
    // The saml property
    saml GhesGetSettings_samlable
    // The signup_enabled property
    signup_enabled *bool
    // The smtp property
    smtp GhesGetSettings_smtpable
    // The snmp property
    snmp GhesGetSettings_snmpable
    // The subdomain_isolation property
    subdomain_isolation *bool
    // The syslog property
    syslog GhesGetSettings_syslogable
    // The timezone property
    timezone *string
}
// NewGhesGetSettings instantiates a new GhesGetSettings and sets the default values.
func NewGhesGetSettings()(*GhesGetSettings) {
    m := &GhesGetSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesGetSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesGetSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesGetSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesGetSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdminPassword gets the admin_password property value. The admin_password property
// returns a *string when successful
func (m *GhesGetSettings) GetAdminPassword()(*string) {
    return m.admin_password
}
// GetAssets gets the assets property value. The assets property
// returns a *string when successful
func (m *GhesGetSettings) GetAssets()(*string) {
    return m.assets
}
// GetAuthMode gets the auth_mode property value. The auth_mode property
// returns a *string when successful
func (m *GhesGetSettings) GetAuthMode()(*string) {
    return m.auth_mode
}
// GetAvatar gets the avatar property value. The avatar property
// returns a GhesGetSettings_avatarable when successful
func (m *GhesGetSettings) GetAvatar()(GhesGetSettings_avatarable) {
    return m.avatar
}
// GetCas gets the cas property value. The cas property
// returns a GhesGetSettings_casable when successful
func (m *GhesGetSettings) GetCas()(GhesGetSettings_casable) {
    return m.cas
}
// GetCollectd gets the collectd property value. The collectd property
// returns a GhesGetSettings_collectdable when successful
func (m *GhesGetSettings) GetCollectd()(GhesGetSettings_collectdable) {
    return m.collectd
}
// GetConfigurationId gets the configuration_id property value. The configuration_id property
// returns a *int32 when successful
func (m *GhesGetSettings) GetConfigurationId()(*int32) {
    return m.configuration_id
}
// GetConfigurationRunCount gets the configuration_run_count property value. The configuration_run_count property
// returns a *int32 when successful
func (m *GhesGetSettings) GetConfigurationRunCount()(*int32) {
    return m.configuration_run_count
}
// GetCustomer gets the customer property value. The customer property
// returns a GhesGetSettings_customerable when successful
func (m *GhesGetSettings) GetCustomer()(GhesGetSettings_customerable) {
    return m.customer
}
// GetExpireSessions gets the expire_sessions property value. The expire_sessions property
// returns a *bool when successful
func (m *GhesGetSettings) GetExpireSessions()(*bool) {
    return m.expire_sessions
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesGetSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["admin_password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminPassword(val)
        }
        return nil
    }
    res["assets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssets(val)
        }
        return nil
    }
    res["auth_mode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthMode(val)
        }
        return nil
    }
    res["avatar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGhesGetSettings_avatarFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvatar(val.(GhesGetSettings_avatarable))
        }
        return nil
    }
    res["cas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGhesGetSettings_casFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCas(val.(GhesGetSettings_casable))
        }
        return nil
    }
    res["collectd"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGhesGetSettings_collectdFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCollectd(val.(GhesGetSettings_collectdable))
        }
        return nil
    }
    res["configuration_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationId(val)
        }
        return nil
    }
    res["configuration_run_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationRunCount(val)
        }
        return nil
    }
    res["customer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGhesGetSettings_customerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomer(val.(GhesGetSettings_customerable))
        }
        return nil
    }
    res["expire_sessions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpireSessions(val)
        }
        return nil
    }
    res["github_hostname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGithubHostname(val)
        }
        return nil
    }
    res["github_oauth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGhesGetSettings_github_oauthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGithubOauth(val.(GhesGetSettings_github_oauthable))
        }
        return nil
    }
    res["github_ssl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGhesGetSettings_github_sslFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGithubSsl(val.(GhesGetSettings_github_sslable))
        }
        return nil
    }
    res["http_proxy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHttpProxy(val)
        }
        return nil
    }
    res["identicons_host"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdenticonsHost(val)
        }
        return nil
    }
    res["ldap"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGhesGetSettings_ldapFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLdap(val.(GhesGetSettings_ldapable))
        }
        return nil
    }
    res["license"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGhesGetSettings_licenseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicense(val.(GhesGetSettings_licenseable))
        }
        return nil
    }
    res["load_balancer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoadBalancer(val)
        }
        return nil
    }
    res["mapping"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGhesGetSettings_mappingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMapping(val.(GhesGetSettings_mappingable))
        }
        return nil
    }
    res["ntp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGhesGetSettings_ntpFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNtp(val.(GhesGetSettings_ntpable))
        }
        return nil
    }
    res["pages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGhesGetSettings_pagesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPages(val.(GhesGetSettings_pagesable))
        }
        return nil
    }
    res["private_mode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateMode(val)
        }
        return nil
    }
    res["public_pages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicPages(val)
        }
        return nil
    }
    res["saml"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGhesGetSettings_samlFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSaml(val.(GhesGetSettings_samlable))
        }
        return nil
    }
    res["signup_enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignupEnabled(val)
        }
        return nil
    }
    res["smtp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGhesGetSettings_smtpFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmtp(val.(GhesGetSettings_smtpable))
        }
        return nil
    }
    res["snmp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGhesGetSettings_snmpFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSnmp(val.(GhesGetSettings_snmpable))
        }
        return nil
    }
    res["subdomain_isolation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubdomainIsolation(val)
        }
        return nil
    }
    res["syslog"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGhesGetSettings_syslogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyslog(val.(GhesGetSettings_syslogable))
        }
        return nil
    }
    res["timezone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimezone(val)
        }
        return nil
    }
    return res
}
// GetGithubHostname gets the github_hostname property value. The github_hostname property
// returns a *string when successful
func (m *GhesGetSettings) GetGithubHostname()(*string) {
    return m.github_hostname
}
// GetGithubOauth gets the github_oauth property value. The github_oauth property
// returns a GhesGetSettings_github_oauthable when successful
func (m *GhesGetSettings) GetGithubOauth()(GhesGetSettings_github_oauthable) {
    return m.github_oauth
}
// GetGithubSsl gets the github_ssl property value. The github_ssl property
// returns a GhesGetSettings_github_sslable when successful
func (m *GhesGetSettings) GetGithubSsl()(GhesGetSettings_github_sslable) {
    return m.github_ssl
}
// GetHttpProxy gets the http_proxy property value. The http_proxy property
// returns a *string when successful
func (m *GhesGetSettings) GetHttpProxy()(*string) {
    return m.http_proxy
}
// GetIdenticonsHost gets the identicons_host property value. The identicons_host property
// returns a *string when successful
func (m *GhesGetSettings) GetIdenticonsHost()(*string) {
    return m.identicons_host
}
// GetLdap gets the ldap property value. The ldap property
// returns a GhesGetSettings_ldapable when successful
func (m *GhesGetSettings) GetLdap()(GhesGetSettings_ldapable) {
    return m.ldap
}
// GetLicense gets the license property value. The license property
// returns a GhesGetSettings_licenseable when successful
func (m *GhesGetSettings) GetLicense()(GhesGetSettings_licenseable) {
    return m.license
}
// GetLoadBalancer gets the load_balancer property value. The load_balancer property
// returns a *string when successful
func (m *GhesGetSettings) GetLoadBalancer()(*string) {
    return m.load_balancer
}
// GetMapping gets the mapping property value. The mapping property
// returns a GhesGetSettings_mappingable when successful
func (m *GhesGetSettings) GetMapping()(GhesGetSettings_mappingable) {
    return m.mapping
}
// GetNtp gets the ntp property value. The ntp property
// returns a GhesGetSettings_ntpable when successful
func (m *GhesGetSettings) GetNtp()(GhesGetSettings_ntpable) {
    return m.ntp
}
// GetPages gets the pages property value. The pages property
// returns a GhesGetSettings_pagesable when successful
func (m *GhesGetSettings) GetPages()(GhesGetSettings_pagesable) {
    return m.pages
}
// GetPrivateMode gets the private_mode property value. The private_mode property
// returns a *bool when successful
func (m *GhesGetSettings) GetPrivateMode()(*bool) {
    return m.private_mode
}
// GetPublicPages gets the public_pages property value. The public_pages property
// returns a *bool when successful
func (m *GhesGetSettings) GetPublicPages()(*bool) {
    return m.public_pages
}
// GetSaml gets the saml property value. The saml property
// returns a GhesGetSettings_samlable when successful
func (m *GhesGetSettings) GetSaml()(GhesGetSettings_samlable) {
    return m.saml
}
// GetSignupEnabled gets the signup_enabled property value. The signup_enabled property
// returns a *bool when successful
func (m *GhesGetSettings) GetSignupEnabled()(*bool) {
    return m.signup_enabled
}
// GetSmtp gets the smtp property value. The smtp property
// returns a GhesGetSettings_smtpable when successful
func (m *GhesGetSettings) GetSmtp()(GhesGetSettings_smtpable) {
    return m.smtp
}
// GetSnmp gets the snmp property value. The snmp property
// returns a GhesGetSettings_snmpable when successful
func (m *GhesGetSettings) GetSnmp()(GhesGetSettings_snmpable) {
    return m.snmp
}
// GetSubdomainIsolation gets the subdomain_isolation property value. The subdomain_isolation property
// returns a *bool when successful
func (m *GhesGetSettings) GetSubdomainIsolation()(*bool) {
    return m.subdomain_isolation
}
// GetSyslog gets the syslog property value. The syslog property
// returns a GhesGetSettings_syslogable when successful
func (m *GhesGetSettings) GetSyslog()(GhesGetSettings_syslogable) {
    return m.syslog
}
// GetTimezone gets the timezone property value. The timezone property
// returns a *string when successful
func (m *GhesGetSettings) GetTimezone()(*string) {
    return m.timezone
}
// Serialize serializes information the current object
func (m *GhesGetSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("admin_password", m.GetAdminPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("assets", m.GetAssets())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("auth_mode", m.GetAuthMode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("avatar", m.GetAvatar())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("cas", m.GetCas())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("collectd", m.GetCollectd())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("configuration_id", m.GetConfigurationId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("configuration_run_count", m.GetConfigurationRunCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("customer", m.GetCustomer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("expire_sessions", m.GetExpireSessions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("github_hostname", m.GetGithubHostname())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("github_oauth", m.GetGithubOauth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("github_ssl", m.GetGithubSsl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("http_proxy", m.GetHttpProxy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("identicons_host", m.GetIdenticonsHost())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("ldap", m.GetLdap())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("license", m.GetLicense())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("load_balancer", m.GetLoadBalancer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("mapping", m.GetMapping())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("ntp", m.GetNtp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("pages", m.GetPages())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("private_mode", m.GetPrivateMode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("public_pages", m.GetPublicPages())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("saml", m.GetSaml())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("signup_enabled", m.GetSignupEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("smtp", m.GetSmtp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("snmp", m.GetSnmp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("subdomain_isolation", m.GetSubdomainIsolation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("syslog", m.GetSyslog())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timezone", m.GetTimezone())
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
func (m *GhesGetSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdminPassword sets the admin_password property value. The admin_password property
func (m *GhesGetSettings) SetAdminPassword(value *string)() {
    m.admin_password = value
}
// SetAssets sets the assets property value. The assets property
func (m *GhesGetSettings) SetAssets(value *string)() {
    m.assets = value
}
// SetAuthMode sets the auth_mode property value. The auth_mode property
func (m *GhesGetSettings) SetAuthMode(value *string)() {
    m.auth_mode = value
}
// SetAvatar sets the avatar property value. The avatar property
func (m *GhesGetSettings) SetAvatar(value GhesGetSettings_avatarable)() {
    m.avatar = value
}
// SetCas sets the cas property value. The cas property
func (m *GhesGetSettings) SetCas(value GhesGetSettings_casable)() {
    m.cas = value
}
// SetCollectd sets the collectd property value. The collectd property
func (m *GhesGetSettings) SetCollectd(value GhesGetSettings_collectdable)() {
    m.collectd = value
}
// SetConfigurationId sets the configuration_id property value. The configuration_id property
func (m *GhesGetSettings) SetConfigurationId(value *int32)() {
    m.configuration_id = value
}
// SetConfigurationRunCount sets the configuration_run_count property value. The configuration_run_count property
func (m *GhesGetSettings) SetConfigurationRunCount(value *int32)() {
    m.configuration_run_count = value
}
// SetCustomer sets the customer property value. The customer property
func (m *GhesGetSettings) SetCustomer(value GhesGetSettings_customerable)() {
    m.customer = value
}
// SetExpireSessions sets the expire_sessions property value. The expire_sessions property
func (m *GhesGetSettings) SetExpireSessions(value *bool)() {
    m.expire_sessions = value
}
// SetGithubHostname sets the github_hostname property value. The github_hostname property
func (m *GhesGetSettings) SetGithubHostname(value *string)() {
    m.github_hostname = value
}
// SetGithubOauth sets the github_oauth property value. The github_oauth property
func (m *GhesGetSettings) SetGithubOauth(value GhesGetSettings_github_oauthable)() {
    m.github_oauth = value
}
// SetGithubSsl sets the github_ssl property value. The github_ssl property
func (m *GhesGetSettings) SetGithubSsl(value GhesGetSettings_github_sslable)() {
    m.github_ssl = value
}
// SetHttpProxy sets the http_proxy property value. The http_proxy property
func (m *GhesGetSettings) SetHttpProxy(value *string)() {
    m.http_proxy = value
}
// SetIdenticonsHost sets the identicons_host property value. The identicons_host property
func (m *GhesGetSettings) SetIdenticonsHost(value *string)() {
    m.identicons_host = value
}
// SetLdap sets the ldap property value. The ldap property
func (m *GhesGetSettings) SetLdap(value GhesGetSettings_ldapable)() {
    m.ldap = value
}
// SetLicense sets the license property value. The license property
func (m *GhesGetSettings) SetLicense(value GhesGetSettings_licenseable)() {
    m.license = value
}
// SetLoadBalancer sets the load_balancer property value. The load_balancer property
func (m *GhesGetSettings) SetLoadBalancer(value *string)() {
    m.load_balancer = value
}
// SetMapping sets the mapping property value. The mapping property
func (m *GhesGetSettings) SetMapping(value GhesGetSettings_mappingable)() {
    m.mapping = value
}
// SetNtp sets the ntp property value. The ntp property
func (m *GhesGetSettings) SetNtp(value GhesGetSettings_ntpable)() {
    m.ntp = value
}
// SetPages sets the pages property value. The pages property
func (m *GhesGetSettings) SetPages(value GhesGetSettings_pagesable)() {
    m.pages = value
}
// SetPrivateMode sets the private_mode property value. The private_mode property
func (m *GhesGetSettings) SetPrivateMode(value *bool)() {
    m.private_mode = value
}
// SetPublicPages sets the public_pages property value. The public_pages property
func (m *GhesGetSettings) SetPublicPages(value *bool)() {
    m.public_pages = value
}
// SetSaml sets the saml property value. The saml property
func (m *GhesGetSettings) SetSaml(value GhesGetSettings_samlable)() {
    m.saml = value
}
// SetSignupEnabled sets the signup_enabled property value. The signup_enabled property
func (m *GhesGetSettings) SetSignupEnabled(value *bool)() {
    m.signup_enabled = value
}
// SetSmtp sets the smtp property value. The smtp property
func (m *GhesGetSettings) SetSmtp(value GhesGetSettings_smtpable)() {
    m.smtp = value
}
// SetSnmp sets the snmp property value. The snmp property
func (m *GhesGetSettings) SetSnmp(value GhesGetSettings_snmpable)() {
    m.snmp = value
}
// SetSubdomainIsolation sets the subdomain_isolation property value. The subdomain_isolation property
func (m *GhesGetSettings) SetSubdomainIsolation(value *bool)() {
    m.subdomain_isolation = value
}
// SetSyslog sets the syslog property value. The syslog property
func (m *GhesGetSettings) SetSyslog(value GhesGetSettings_syslogable)() {
    m.syslog = value
}
// SetTimezone sets the timezone property value. The timezone property
func (m *GhesGetSettings) SetTimezone(value *string)() {
    m.timezone = value
}
type GhesGetSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdminPassword()(*string)
    GetAssets()(*string)
    GetAuthMode()(*string)
    GetAvatar()(GhesGetSettings_avatarable)
    GetCas()(GhesGetSettings_casable)
    GetCollectd()(GhesGetSettings_collectdable)
    GetConfigurationId()(*int32)
    GetConfigurationRunCount()(*int32)
    GetCustomer()(GhesGetSettings_customerable)
    GetExpireSessions()(*bool)
    GetGithubHostname()(*string)
    GetGithubOauth()(GhesGetSettings_github_oauthable)
    GetGithubSsl()(GhesGetSettings_github_sslable)
    GetHttpProxy()(*string)
    GetIdenticonsHost()(*string)
    GetLdap()(GhesGetSettings_ldapable)
    GetLicense()(GhesGetSettings_licenseable)
    GetLoadBalancer()(*string)
    GetMapping()(GhesGetSettings_mappingable)
    GetNtp()(GhesGetSettings_ntpable)
    GetPages()(GhesGetSettings_pagesable)
    GetPrivateMode()(*bool)
    GetPublicPages()(*bool)
    GetSaml()(GhesGetSettings_samlable)
    GetSignupEnabled()(*bool)
    GetSmtp()(GhesGetSettings_smtpable)
    GetSnmp()(GhesGetSettings_snmpable)
    GetSubdomainIsolation()(*bool)
    GetSyslog()(GhesGetSettings_syslogable)
    GetTimezone()(*string)
    SetAdminPassword(value *string)()
    SetAssets(value *string)()
    SetAuthMode(value *string)()
    SetAvatar(value GhesGetSettings_avatarable)()
    SetCas(value GhesGetSettings_casable)()
    SetCollectd(value GhesGetSettings_collectdable)()
    SetConfigurationId(value *int32)()
    SetConfigurationRunCount(value *int32)()
    SetCustomer(value GhesGetSettings_customerable)()
    SetExpireSessions(value *bool)()
    SetGithubHostname(value *string)()
    SetGithubOauth(value GhesGetSettings_github_oauthable)()
    SetGithubSsl(value GhesGetSettings_github_sslable)()
    SetHttpProxy(value *string)()
    SetIdenticonsHost(value *string)()
    SetLdap(value GhesGetSettings_ldapable)()
    SetLicense(value GhesGetSettings_licenseable)()
    SetLoadBalancer(value *string)()
    SetMapping(value GhesGetSettings_mappingable)()
    SetNtp(value GhesGetSettings_ntpable)()
    SetPages(value GhesGetSettings_pagesable)()
    SetPrivateMode(value *bool)()
    SetPublicPages(value *bool)()
    SetSaml(value GhesGetSettings_samlable)()
    SetSignupEnabled(value *bool)()
    SetSmtp(value GhesGetSettings_smtpable)()
    SetSnmp(value GhesGetSettings_snmpable)()
    SetSubdomainIsolation(value *bool)()
    SetSyslog(value GhesGetSettings_syslogable)()
    SetTimezone(value *string)()
}
