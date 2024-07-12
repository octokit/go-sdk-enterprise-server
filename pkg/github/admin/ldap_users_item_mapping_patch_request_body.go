package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LdapUsersItemMappingPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The [distinguished name](https://www.ldap.com/ldap-dns-and-rdns) (DN) of the LDAP entry to map to a team.
    ldap_dn *string
}
// NewLdapUsersItemMappingPatchRequestBody instantiates a new LdapUsersItemMappingPatchRequestBody and sets the default values.
func NewLdapUsersItemMappingPatchRequestBody()(*LdapUsersItemMappingPatchRequestBody) {
    m := &LdapUsersItemMappingPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLdapUsersItemMappingPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLdapUsersItemMappingPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLdapUsersItemMappingPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LdapUsersItemMappingPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LdapUsersItemMappingPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    return res
}
// GetLdapDn gets the ldap_dn property value. The [distinguished name](https://www.ldap.com/ldap-dns-and-rdns) (DN) of the LDAP entry to map to a team.
// returns a *string when successful
func (m *LdapUsersItemMappingPatchRequestBody) GetLdapDn()(*string) {
    return m.ldap_dn
}
// Serialize serializes information the current object
func (m *LdapUsersItemMappingPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("ldap_dn", m.GetLdapDn())
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
func (m *LdapUsersItemMappingPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLdapDn sets the ldap_dn property value. The [distinguished name](https://www.ldap.com/ldap-dns-and-rdns) (DN) of the LDAP entry to map to a team.
func (m *LdapUsersItemMappingPatchRequestBody) SetLdapDn(value *string)() {
    m.ldap_dn = value
}
type LdapUsersItemMappingPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLdapDn()(*string)
    SetLdapDn(value *string)()
}
