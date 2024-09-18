package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesGetSettings_ldap_profile struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The key property
    key *string
    // The mail property
    mail *string
    // The name property
    name *string
    // The uid property
    uid *string
}
// NewGhesGetSettings_ldap_profile instantiates a new GhesGetSettings_ldap_profile and sets the default values.
func NewGhesGetSettings_ldap_profile()(*GhesGetSettings_ldap_profile) {
    m := &GhesGetSettings_ldap_profile{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesGetSettings_ldap_profileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesGetSettings_ldap_profileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesGetSettings_ldap_profile(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesGetSettings_ldap_profile) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesGetSettings_ldap_profile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    res["mail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMail(val)
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
    return res
}
// GetKey gets the key property value. The key property
// returns a *string when successful
func (m *GhesGetSettings_ldap_profile) GetKey()(*string) {
    return m.key
}
// GetMail gets the mail property value. The mail property
// returns a *string when successful
func (m *GhesGetSettings_ldap_profile) GetMail()(*string) {
    return m.mail
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *GhesGetSettings_ldap_profile) GetName()(*string) {
    return m.name
}
// GetUid gets the uid property value. The uid property
// returns a *string when successful
func (m *GhesGetSettings_ldap_profile) GetUid()(*string) {
    return m.uid
}
// Serialize serializes information the current object
func (m *GhesGetSettings_ldap_profile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mail", m.GetMail())
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
        err := writer.WriteStringValue("uid", m.GetUid())
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
func (m *GhesGetSettings_ldap_profile) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKey sets the key property value. The key property
func (m *GhesGetSettings_ldap_profile) SetKey(value *string)() {
    m.key = value
}
// SetMail sets the mail property value. The mail property
func (m *GhesGetSettings_ldap_profile) SetMail(value *string)() {
    m.mail = value
}
// SetName sets the name property value. The name property
func (m *GhesGetSettings_ldap_profile) SetName(value *string)() {
    m.name = value
}
// SetUid sets the uid property value. The uid property
func (m *GhesGetSettings_ldap_profile) SetUid(value *string)() {
    m.uid = value
}
type GhesGetSettings_ldap_profileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKey()(*string)
    GetMail()(*string)
    GetName()(*string)
    GetUid()(*string)
    SetKey(value *string)()
    SetMail(value *string)()
    SetName(value *string)()
    SetUid(value *string)()
}
