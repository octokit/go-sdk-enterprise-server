package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesGetSettings_saml struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The certificate property
    certificate *string
    // The certificate_path property
    certificate_path *string
    // The disable_admin_demote property
    disable_admin_demote *bool
    // The idp_initiated_sso property
    idp_initiated_sso *bool
    // The issuer property
    issuer *string
    // The sso_url property
    sso_url *string
}
// NewGhesGetSettings_saml instantiates a new GhesGetSettings_saml and sets the default values.
func NewGhesGetSettings_saml()(*GhesGetSettings_saml) {
    m := &GhesGetSettings_saml{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesGetSettings_samlFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesGetSettings_samlFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesGetSettings_saml(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesGetSettings_saml) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCertificate gets the certificate property value. The certificate property
// returns a *string when successful
func (m *GhesGetSettings_saml) GetCertificate()(*string) {
    return m.certificate
}
// GetCertificatePath gets the certificate_path property value. The certificate_path property
// returns a *string when successful
func (m *GhesGetSettings_saml) GetCertificatePath()(*string) {
    return m.certificate_path
}
// GetDisableAdminDemote gets the disable_admin_demote property value. The disable_admin_demote property
// returns a *bool when successful
func (m *GhesGetSettings_saml) GetDisableAdminDemote()(*bool) {
    return m.disable_admin_demote
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesGetSettings_saml) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["certificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificate(val)
        }
        return nil
    }
    res["certificate_path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificatePath(val)
        }
        return nil
    }
    res["disable_admin_demote"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableAdminDemote(val)
        }
        return nil
    }
    res["idp_initiated_sso"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdpInitiatedSso(val)
        }
        return nil
    }
    res["issuer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuer(val)
        }
        return nil
    }
    res["sso_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSsoUrl(val)
        }
        return nil
    }
    return res
}
// GetIdpInitiatedSso gets the idp_initiated_sso property value. The idp_initiated_sso property
// returns a *bool when successful
func (m *GhesGetSettings_saml) GetIdpInitiatedSso()(*bool) {
    return m.idp_initiated_sso
}
// GetIssuer gets the issuer property value. The issuer property
// returns a *string when successful
func (m *GhesGetSettings_saml) GetIssuer()(*string) {
    return m.issuer
}
// GetSsoUrl gets the sso_url property value. The sso_url property
// returns a *string when successful
func (m *GhesGetSettings_saml) GetSsoUrl()(*string) {
    return m.sso_url
}
// Serialize serializes information the current object
func (m *GhesGetSettings_saml) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("certificate", m.GetCertificate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("certificate_path", m.GetCertificatePath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("disable_admin_demote", m.GetDisableAdminDemote())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("idp_initiated_sso", m.GetIdpInitiatedSso())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issuer", m.GetIssuer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sso_url", m.GetSsoUrl())
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
func (m *GhesGetSettings_saml) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCertificate sets the certificate property value. The certificate property
func (m *GhesGetSettings_saml) SetCertificate(value *string)() {
    m.certificate = value
}
// SetCertificatePath sets the certificate_path property value. The certificate_path property
func (m *GhesGetSettings_saml) SetCertificatePath(value *string)() {
    m.certificate_path = value
}
// SetDisableAdminDemote sets the disable_admin_demote property value. The disable_admin_demote property
func (m *GhesGetSettings_saml) SetDisableAdminDemote(value *bool)() {
    m.disable_admin_demote = value
}
// SetIdpInitiatedSso sets the idp_initiated_sso property value. The idp_initiated_sso property
func (m *GhesGetSettings_saml) SetIdpInitiatedSso(value *bool)() {
    m.idp_initiated_sso = value
}
// SetIssuer sets the issuer property value. The issuer property
func (m *GhesGetSettings_saml) SetIssuer(value *string)() {
    m.issuer = value
}
// SetSsoUrl sets the sso_url property value. The sso_url property
func (m *GhesGetSettings_saml) SetSsoUrl(value *string)() {
    m.sso_url = value
}
type GhesGetSettings_samlable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificate()(*string)
    GetCertificatePath()(*string)
    GetDisableAdminDemote()(*bool)
    GetIdpInitiatedSso()(*bool)
    GetIssuer()(*string)
    GetSsoUrl()(*string)
    SetCertificate(value *string)()
    SetCertificatePath(value *string)()
    SetDisableAdminDemote(value *bool)()
    SetIdpInitiatedSso(value *bool)()
    SetIssuer(value *string)()
    SetSsoUrl(value *string)()
}
