package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EnterpriseSettings_enterprise_ldap_reconciliation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The org property
    org *string
    // The user property
    user *string
}
// NewEnterpriseSettings_enterprise_ldap_reconciliation instantiates a new EnterpriseSettings_enterprise_ldap_reconciliation and sets the default values.
func NewEnterpriseSettings_enterprise_ldap_reconciliation()(*EnterpriseSettings_enterprise_ldap_reconciliation) {
    m := &EnterpriseSettings_enterprise_ldap_reconciliation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnterpriseSettings_enterprise_ldap_reconciliationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnterpriseSettings_enterprise_ldap_reconciliationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnterpriseSettings_enterprise_ldap_reconciliation(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EnterpriseSettings_enterprise_ldap_reconciliation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnterpriseSettings_enterprise_ldap_reconciliation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["org"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrg(val)
        }
        return nil
    }
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val)
        }
        return nil
    }
    return res
}
// GetOrg gets the org property value. The org property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_ldap_reconciliation) GetOrg()(*string) {
    return m.org
}
// GetUser gets the user property value. The user property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_ldap_reconciliation) GetUser()(*string) {
    return m.user
}
// Serialize serializes information the current object
func (m *EnterpriseSettings_enterprise_ldap_reconciliation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("org", m.GetOrg())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("user", m.GetUser())
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
func (m *EnterpriseSettings_enterprise_ldap_reconciliation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOrg sets the org property value. The org property
func (m *EnterpriseSettings_enterprise_ldap_reconciliation) SetOrg(value *string)() {
    m.org = value
}
// SetUser sets the user property value. The user property
func (m *EnterpriseSettings_enterprise_ldap_reconciliation) SetUser(value *string)() {
    m.user = value
}
type EnterpriseSettings_enterprise_ldap_reconciliationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOrg()(*string)
    GetUser()(*string)
    SetOrg(value *string)()
    SetUser(value *string)()
}
