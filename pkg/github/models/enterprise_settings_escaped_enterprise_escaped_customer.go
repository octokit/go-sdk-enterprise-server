package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EnterpriseSettings_enterprise_customer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The email property
    email *string
    // The name property
    name *string
    // The public_key_data property
    public_key_data *string
    // The secret_key_data property
    secret_key_data *string
    // The uuid property
    uuid *string
}
// NewEnterpriseSettings_enterprise_customer instantiates a new EnterpriseSettings_enterprise_customer and sets the default values.
func NewEnterpriseSettings_enterprise_customer()(*EnterpriseSettings_enterprise_customer) {
    m := &EnterpriseSettings_enterprise_customer{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnterpriseSettings_enterprise_customerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnterpriseSettings_enterprise_customerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnterpriseSettings_enterprise_customer(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EnterpriseSettings_enterprise_customer) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_customer) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnterpriseSettings_enterprise_customer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
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
    res["public_key_data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicKeyData(val)
        }
        return nil
    }
    res["secret_key_data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretKeyData(val)
        }
        return nil
    }
    res["uuid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUuid(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_customer) GetName()(*string) {
    return m.name
}
// GetPublicKeyData gets the public_key_data property value. The public_key_data property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_customer) GetPublicKeyData()(*string) {
    return m.public_key_data
}
// GetSecretKeyData gets the secret_key_data property value. The secret_key_data property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_customer) GetSecretKeyData()(*string) {
    return m.secret_key_data
}
// GetUuid gets the uuid property value. The uuid property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_customer) GetUuid()(*string) {
    return m.uuid
}
// Serialize serializes information the current object
func (m *EnterpriseSettings_enterprise_customer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("email", m.GetEmail())
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
        err := writer.WriteStringValue("public_key_data", m.GetPublicKeyData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secret_key_data", m.GetSecretKeyData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("uuid", m.GetUuid())
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
func (m *EnterpriseSettings_enterprise_customer) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmail sets the email property value. The email property
func (m *EnterpriseSettings_enterprise_customer) SetEmail(value *string)() {
    m.email = value
}
// SetName sets the name property value. The name property
func (m *EnterpriseSettings_enterprise_customer) SetName(value *string)() {
    m.name = value
}
// SetPublicKeyData sets the public_key_data property value. The public_key_data property
func (m *EnterpriseSettings_enterprise_customer) SetPublicKeyData(value *string)() {
    m.public_key_data = value
}
// SetSecretKeyData sets the secret_key_data property value. The secret_key_data property
func (m *EnterpriseSettings_enterprise_customer) SetSecretKeyData(value *string)() {
    m.secret_key_data = value
}
// SetUuid sets the uuid property value. The uuid property
func (m *EnterpriseSettings_enterprise_customer) SetUuid(value *string)() {
    m.uuid = value
}
type EnterpriseSettings_enterprise_customerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmail()(*string)
    GetName()(*string)
    GetPublicKeyData()(*string)
    GetSecretKeyData()(*string)
    GetUuid()(*string)
    SetEmail(value *string)()
    SetName(value *string)()
    SetPublicKeyData(value *string)()
    SetSecretKeyData(value *string)()
    SetUuid(value *string)()
}
