package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EnterpriseSettings_enterprise struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The admin_password property
    admin_password *string
    // The assets property
    assets *string
    // The auth_mode property
    auth_mode *string
    // The avatar property
    avatar EnterpriseSettings_enterprise_avatarable
    // The cas property
    cas EnterpriseSettings_enterprise_casable
    // The collectd property
    collectd EnterpriseSettings_enterprise_collectdable
    // The configuration_id property
    configuration_id *int32
    // The configuration_run_count property
    configuration_run_count *int32
    // The customer property
    customer EnterpriseSettings_enterprise_customerable
    // The expire_sessions property
    expire_sessions *bool
    // The github_hostname property
    github_hostname *string
    // The github_oauth property
    github_oauth EnterpriseSettings_enterprise_github_oauthable
    // The github_ssl property
    github_ssl EnterpriseSettings_enterprise_github_sslable
    // The http_proxy property
    http_proxy *string
    // The identicons_host property
    identicons_host *string
    // The ldap property
    ldap EnterpriseSettings_enterprise_ldapable
    // The license property
    license EnterpriseSettings_enterprise_licenseable
    // The load_balancer property
    load_balancer *string
    // The mapping property
    mapping EnterpriseSettings_enterprise_mappingable
    // The ntp property
    ntp EnterpriseSettings_enterprise_ntpable
    // The pages property
    pages EnterpriseSettings_enterprise_pagesable
    // The private_mode property
    private_mode *bool
    // The public_pages property
    public_pages *bool
    // The saml property
    saml EnterpriseSettings_enterprise_samlable
    // The signup_enabled property
    signup_enabled *bool
    // The smtp property
    smtp EnterpriseSettings_enterprise_smtpable
    // The snmp property
    snmp EnterpriseSettings_enterprise_snmpable
    // The subdomain_isolation property
    subdomain_isolation *bool
    // The syslog property
    syslog EnterpriseSettings_enterprise_syslogable
    // The timezone property
    timezone *string
}
// NewEnterpriseSettings_enterprise instantiates a new EnterpriseSettings_enterprise and sets the default values.
func NewEnterpriseSettings_enterprise()(*EnterpriseSettings_enterprise) {
    m := &EnterpriseSettings_enterprise{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnterpriseSettings_enterpriseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnterpriseSettings_enterpriseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnterpriseSettings_enterprise(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EnterpriseSettings_enterprise) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdminPassword gets the admin_password property value. The admin_password property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise) GetAdminPassword()(*string) {
    return m.admin_password
}
// GetAssets gets the assets property value. The assets property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise) GetAssets()(*string) {
    return m.assets
}
// GetAuthMode gets the auth_mode property value. The auth_mode property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise) GetAuthMode()(*string) {
    return m.auth_mode
}
// GetAvatar gets the avatar property value. The avatar property
// returns a EnterpriseSettings_enterprise_avatarable when successful
func (m *EnterpriseSettings_enterprise) GetAvatar()(EnterpriseSettings_enterprise_avatarable) {
    return m.avatar
}
// GetCas gets the cas property value. The cas property
// returns a EnterpriseSettings_enterprise_casable when successful
func (m *EnterpriseSettings_enterprise) GetCas()(EnterpriseSettings_enterprise_casable) {
    return m.cas
}
// GetCollectd gets the collectd property value. The collectd property
// returns a EnterpriseSettings_enterprise_collectdable when successful
func (m *EnterpriseSettings_enterprise) GetCollectd()(EnterpriseSettings_enterprise_collectdable) {
    return m.collectd
}
// GetConfigurationId gets the configuration_id property value. The configuration_id property
// returns a *int32 when successful
func (m *EnterpriseSettings_enterprise) GetConfigurationId()(*int32) {
    return m.configuration_id
}
// GetConfigurationRunCount gets the configuration_run_count property value. The configuration_run_count property
// returns a *int32 when successful
func (m *EnterpriseSettings_enterprise) GetConfigurationRunCount()(*int32) {
    return m.configuration_run_count
}
// GetCustomer gets the customer property value. The customer property
// returns a EnterpriseSettings_enterprise_customerable when successful
func (m *EnterpriseSettings_enterprise) GetCustomer()(EnterpriseSettings_enterprise_customerable) {
    return m.customer
}
// GetExpireSessions gets the expire_sessions property value. The expire_sessions property
// returns a *bool when successful
func (m *EnterpriseSettings_enterprise) GetExpireSessions()(*bool) {
    return m.expire_sessions
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnterpriseSettings_enterprise) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(CreateEnterpriseSettings_enterprise_avatarFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvatar(val.(EnterpriseSettings_enterprise_avatarable))
        }
        return nil
    }
    res["cas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnterpriseSettings_enterprise_casFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCas(val.(EnterpriseSettings_enterprise_casable))
        }
        return nil
    }
    res["collectd"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnterpriseSettings_enterprise_collectdFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCollectd(val.(EnterpriseSettings_enterprise_collectdable))
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
        val, err := n.GetObjectValue(CreateEnterpriseSettings_enterprise_customerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomer(val.(EnterpriseSettings_enterprise_customerable))
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
        val, err := n.GetObjectValue(CreateEnterpriseSettings_enterprise_github_oauthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGithubOauth(val.(EnterpriseSettings_enterprise_github_oauthable))
        }
        return nil
    }
    res["github_ssl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnterpriseSettings_enterprise_github_sslFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGithubSsl(val.(EnterpriseSettings_enterprise_github_sslable))
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
        val, err := n.GetObjectValue(CreateEnterpriseSettings_enterprise_ldapFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLdap(val.(EnterpriseSettings_enterprise_ldapable))
        }
        return nil
    }
    res["license"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnterpriseSettings_enterprise_licenseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicense(val.(EnterpriseSettings_enterprise_licenseable))
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
        val, err := n.GetObjectValue(CreateEnterpriseSettings_enterprise_mappingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMapping(val.(EnterpriseSettings_enterprise_mappingable))
        }
        return nil
    }
    res["ntp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnterpriseSettings_enterprise_ntpFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNtp(val.(EnterpriseSettings_enterprise_ntpable))
        }
        return nil
    }
    res["pages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnterpriseSettings_enterprise_pagesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPages(val.(EnterpriseSettings_enterprise_pagesable))
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
        val, err := n.GetObjectValue(CreateEnterpriseSettings_enterprise_samlFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSaml(val.(EnterpriseSettings_enterprise_samlable))
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
        val, err := n.GetObjectValue(CreateEnterpriseSettings_enterprise_smtpFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmtp(val.(EnterpriseSettings_enterprise_smtpable))
        }
        return nil
    }
    res["snmp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnterpriseSettings_enterprise_snmpFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSnmp(val.(EnterpriseSettings_enterprise_snmpable))
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
        val, err := n.GetObjectValue(CreateEnterpriseSettings_enterprise_syslogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyslog(val.(EnterpriseSettings_enterprise_syslogable))
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
func (m *EnterpriseSettings_enterprise) GetGithubHostname()(*string) {
    return m.github_hostname
}
// GetGithubOauth gets the github_oauth property value. The github_oauth property
// returns a EnterpriseSettings_enterprise_github_oauthable when successful
func (m *EnterpriseSettings_enterprise) GetGithubOauth()(EnterpriseSettings_enterprise_github_oauthable) {
    return m.github_oauth
}
// GetGithubSsl gets the github_ssl property value. The github_ssl property
// returns a EnterpriseSettings_enterprise_github_sslable when successful
func (m *EnterpriseSettings_enterprise) GetGithubSsl()(EnterpriseSettings_enterprise_github_sslable) {
    return m.github_ssl
}
// GetHttpProxy gets the http_proxy property value. The http_proxy property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise) GetHttpProxy()(*string) {
    return m.http_proxy
}
// GetIdenticonsHost gets the identicons_host property value. The identicons_host property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise) GetIdenticonsHost()(*string) {
    return m.identicons_host
}
// GetLdap gets the ldap property value. The ldap property
// returns a EnterpriseSettings_enterprise_ldapable when successful
func (m *EnterpriseSettings_enterprise) GetLdap()(EnterpriseSettings_enterprise_ldapable) {
    return m.ldap
}
// GetLicense gets the license property value. The license property
// returns a EnterpriseSettings_enterprise_licenseable when successful
func (m *EnterpriseSettings_enterprise) GetLicense()(EnterpriseSettings_enterprise_licenseable) {
    return m.license
}
// GetLoadBalancer gets the load_balancer property value. The load_balancer property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise) GetLoadBalancer()(*string) {
    return m.load_balancer
}
// GetMapping gets the mapping property value. The mapping property
// returns a EnterpriseSettings_enterprise_mappingable when successful
func (m *EnterpriseSettings_enterprise) GetMapping()(EnterpriseSettings_enterprise_mappingable) {
    return m.mapping
}
// GetNtp gets the ntp property value. The ntp property
// returns a EnterpriseSettings_enterprise_ntpable when successful
func (m *EnterpriseSettings_enterprise) GetNtp()(EnterpriseSettings_enterprise_ntpable) {
    return m.ntp
}
// GetPages gets the pages property value. The pages property
// returns a EnterpriseSettings_enterprise_pagesable when successful
func (m *EnterpriseSettings_enterprise) GetPages()(EnterpriseSettings_enterprise_pagesable) {
    return m.pages
}
// GetPrivateMode gets the private_mode property value. The private_mode property
// returns a *bool when successful
func (m *EnterpriseSettings_enterprise) GetPrivateMode()(*bool) {
    return m.private_mode
}
// GetPublicPages gets the public_pages property value. The public_pages property
// returns a *bool when successful
func (m *EnterpriseSettings_enterprise) GetPublicPages()(*bool) {
    return m.public_pages
}
// GetSaml gets the saml property value. The saml property
// returns a EnterpriseSettings_enterprise_samlable when successful
func (m *EnterpriseSettings_enterprise) GetSaml()(EnterpriseSettings_enterprise_samlable) {
    return m.saml
}
// GetSignupEnabled gets the signup_enabled property value. The signup_enabled property
// returns a *bool when successful
func (m *EnterpriseSettings_enterprise) GetSignupEnabled()(*bool) {
    return m.signup_enabled
}
// GetSmtp gets the smtp property value. The smtp property
// returns a EnterpriseSettings_enterprise_smtpable when successful
func (m *EnterpriseSettings_enterprise) GetSmtp()(EnterpriseSettings_enterprise_smtpable) {
    return m.smtp
}
// GetSnmp gets the snmp property value. The snmp property
// returns a EnterpriseSettings_enterprise_snmpable when successful
func (m *EnterpriseSettings_enterprise) GetSnmp()(EnterpriseSettings_enterprise_snmpable) {
    return m.snmp
}
// GetSubdomainIsolation gets the subdomain_isolation property value. The subdomain_isolation property
// returns a *bool when successful
func (m *EnterpriseSettings_enterprise) GetSubdomainIsolation()(*bool) {
    return m.subdomain_isolation
}
// GetSyslog gets the syslog property value. The syslog property
// returns a EnterpriseSettings_enterprise_syslogable when successful
func (m *EnterpriseSettings_enterprise) GetSyslog()(EnterpriseSettings_enterprise_syslogable) {
    return m.syslog
}
// GetTimezone gets the timezone property value. The timezone property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise) GetTimezone()(*string) {
    return m.timezone
}
// Serialize serializes information the current object
func (m *EnterpriseSettings_enterprise) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *EnterpriseSettings_enterprise) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdminPassword sets the admin_password property value. The admin_password property
func (m *EnterpriseSettings_enterprise) SetAdminPassword(value *string)() {
    m.admin_password = value
}
// SetAssets sets the assets property value. The assets property
func (m *EnterpriseSettings_enterprise) SetAssets(value *string)() {
    m.assets = value
}
// SetAuthMode sets the auth_mode property value. The auth_mode property
func (m *EnterpriseSettings_enterprise) SetAuthMode(value *string)() {
    m.auth_mode = value
}
// SetAvatar sets the avatar property value. The avatar property
func (m *EnterpriseSettings_enterprise) SetAvatar(value EnterpriseSettings_enterprise_avatarable)() {
    m.avatar = value
}
// SetCas sets the cas property value. The cas property
func (m *EnterpriseSettings_enterprise) SetCas(value EnterpriseSettings_enterprise_casable)() {
    m.cas = value
}
// SetCollectd sets the collectd property value. The collectd property
func (m *EnterpriseSettings_enterprise) SetCollectd(value EnterpriseSettings_enterprise_collectdable)() {
    m.collectd = value
}
// SetConfigurationId sets the configuration_id property value. The configuration_id property
func (m *EnterpriseSettings_enterprise) SetConfigurationId(value *int32)() {
    m.configuration_id = value
}
// SetConfigurationRunCount sets the configuration_run_count property value. The configuration_run_count property
func (m *EnterpriseSettings_enterprise) SetConfigurationRunCount(value *int32)() {
    m.configuration_run_count = value
}
// SetCustomer sets the customer property value. The customer property
func (m *EnterpriseSettings_enterprise) SetCustomer(value EnterpriseSettings_enterprise_customerable)() {
    m.customer = value
}
// SetExpireSessions sets the expire_sessions property value. The expire_sessions property
func (m *EnterpriseSettings_enterprise) SetExpireSessions(value *bool)() {
    m.expire_sessions = value
}
// SetGithubHostname sets the github_hostname property value. The github_hostname property
func (m *EnterpriseSettings_enterprise) SetGithubHostname(value *string)() {
    m.github_hostname = value
}
// SetGithubOauth sets the github_oauth property value. The github_oauth property
func (m *EnterpriseSettings_enterprise) SetGithubOauth(value EnterpriseSettings_enterprise_github_oauthable)() {
    m.github_oauth = value
}
// SetGithubSsl sets the github_ssl property value. The github_ssl property
func (m *EnterpriseSettings_enterprise) SetGithubSsl(value EnterpriseSettings_enterprise_github_sslable)() {
    m.github_ssl = value
}
// SetHttpProxy sets the http_proxy property value. The http_proxy property
func (m *EnterpriseSettings_enterprise) SetHttpProxy(value *string)() {
    m.http_proxy = value
}
// SetIdenticonsHost sets the identicons_host property value. The identicons_host property
func (m *EnterpriseSettings_enterprise) SetIdenticonsHost(value *string)() {
    m.identicons_host = value
}
// SetLdap sets the ldap property value. The ldap property
func (m *EnterpriseSettings_enterprise) SetLdap(value EnterpriseSettings_enterprise_ldapable)() {
    m.ldap = value
}
// SetLicense sets the license property value. The license property
func (m *EnterpriseSettings_enterprise) SetLicense(value EnterpriseSettings_enterprise_licenseable)() {
    m.license = value
}
// SetLoadBalancer sets the load_balancer property value. The load_balancer property
func (m *EnterpriseSettings_enterprise) SetLoadBalancer(value *string)() {
    m.load_balancer = value
}
// SetMapping sets the mapping property value. The mapping property
func (m *EnterpriseSettings_enterprise) SetMapping(value EnterpriseSettings_enterprise_mappingable)() {
    m.mapping = value
}
// SetNtp sets the ntp property value. The ntp property
func (m *EnterpriseSettings_enterprise) SetNtp(value EnterpriseSettings_enterprise_ntpable)() {
    m.ntp = value
}
// SetPages sets the pages property value. The pages property
func (m *EnterpriseSettings_enterprise) SetPages(value EnterpriseSettings_enterprise_pagesable)() {
    m.pages = value
}
// SetPrivateMode sets the private_mode property value. The private_mode property
func (m *EnterpriseSettings_enterprise) SetPrivateMode(value *bool)() {
    m.private_mode = value
}
// SetPublicPages sets the public_pages property value. The public_pages property
func (m *EnterpriseSettings_enterprise) SetPublicPages(value *bool)() {
    m.public_pages = value
}
// SetSaml sets the saml property value. The saml property
func (m *EnterpriseSettings_enterprise) SetSaml(value EnterpriseSettings_enterprise_samlable)() {
    m.saml = value
}
// SetSignupEnabled sets the signup_enabled property value. The signup_enabled property
func (m *EnterpriseSettings_enterprise) SetSignupEnabled(value *bool)() {
    m.signup_enabled = value
}
// SetSmtp sets the smtp property value. The smtp property
func (m *EnterpriseSettings_enterprise) SetSmtp(value EnterpriseSettings_enterprise_smtpable)() {
    m.smtp = value
}
// SetSnmp sets the snmp property value. The snmp property
func (m *EnterpriseSettings_enterprise) SetSnmp(value EnterpriseSettings_enterprise_snmpable)() {
    m.snmp = value
}
// SetSubdomainIsolation sets the subdomain_isolation property value. The subdomain_isolation property
func (m *EnterpriseSettings_enterprise) SetSubdomainIsolation(value *bool)() {
    m.subdomain_isolation = value
}
// SetSyslog sets the syslog property value. The syslog property
func (m *EnterpriseSettings_enterprise) SetSyslog(value EnterpriseSettings_enterprise_syslogable)() {
    m.syslog = value
}
// SetTimezone sets the timezone property value. The timezone property
func (m *EnterpriseSettings_enterprise) SetTimezone(value *string)() {
    m.timezone = value
}
type EnterpriseSettings_enterpriseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdminPassword()(*string)
    GetAssets()(*string)
    GetAuthMode()(*string)
    GetAvatar()(EnterpriseSettings_enterprise_avatarable)
    GetCas()(EnterpriseSettings_enterprise_casable)
    GetCollectd()(EnterpriseSettings_enterprise_collectdable)
    GetConfigurationId()(*int32)
    GetConfigurationRunCount()(*int32)
    GetCustomer()(EnterpriseSettings_enterprise_customerable)
    GetExpireSessions()(*bool)
    GetGithubHostname()(*string)
    GetGithubOauth()(EnterpriseSettings_enterprise_github_oauthable)
    GetGithubSsl()(EnterpriseSettings_enterprise_github_sslable)
    GetHttpProxy()(*string)
    GetIdenticonsHost()(*string)
    GetLdap()(EnterpriseSettings_enterprise_ldapable)
    GetLicense()(EnterpriseSettings_enterprise_licenseable)
    GetLoadBalancer()(*string)
    GetMapping()(EnterpriseSettings_enterprise_mappingable)
    GetNtp()(EnterpriseSettings_enterprise_ntpable)
    GetPages()(EnterpriseSettings_enterprise_pagesable)
    GetPrivateMode()(*bool)
    GetPublicPages()(*bool)
    GetSaml()(EnterpriseSettings_enterprise_samlable)
    GetSignupEnabled()(*bool)
    GetSmtp()(EnterpriseSettings_enterprise_smtpable)
    GetSnmp()(EnterpriseSettings_enterprise_snmpable)
    GetSubdomainIsolation()(*bool)
    GetSyslog()(EnterpriseSettings_enterprise_syslogable)
    GetTimezone()(*string)
    SetAdminPassword(value *string)()
    SetAssets(value *string)()
    SetAuthMode(value *string)()
    SetAvatar(value EnterpriseSettings_enterprise_avatarable)()
    SetCas(value EnterpriseSettings_enterprise_casable)()
    SetCollectd(value EnterpriseSettings_enterprise_collectdable)()
    SetConfigurationId(value *int32)()
    SetConfigurationRunCount(value *int32)()
    SetCustomer(value EnterpriseSettings_enterprise_customerable)()
    SetExpireSessions(value *bool)()
    SetGithubHostname(value *string)()
    SetGithubOauth(value EnterpriseSettings_enterprise_github_oauthable)()
    SetGithubSsl(value EnterpriseSettings_enterprise_github_sslable)()
    SetHttpProxy(value *string)()
    SetIdenticonsHost(value *string)()
    SetLdap(value EnterpriseSettings_enterprise_ldapable)()
    SetLicense(value EnterpriseSettings_enterprise_licenseable)()
    SetLoadBalancer(value *string)()
    SetMapping(value EnterpriseSettings_enterprise_mappingable)()
    SetNtp(value EnterpriseSettings_enterprise_ntpable)()
    SetPages(value EnterpriseSettings_enterprise_pagesable)()
    SetPrivateMode(value *bool)()
    SetPublicPages(value *bool)()
    SetSaml(value EnterpriseSettings_enterprise_samlable)()
    SetSignupEnabled(value *bool)()
    SetSmtp(value EnterpriseSettings_enterprise_smtpable)()
    SetSnmp(value EnterpriseSettings_enterprise_snmpable)()
    SetSubdomainIsolation(value *bool)()
    SetSyslog(value EnterpriseSettings_enterprise_syslogable)()
    SetTimezone(value *string)()
}
