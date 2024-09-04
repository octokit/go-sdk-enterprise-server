package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UsersPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // **Required for built-in authentication.** The user's emailaddress. This parameter can be omitted when using CAS, LDAP, or SAML.For more information, see "[About authentication for your enterprise](https://docs.github.com/enterprise-server@3.14/admin/identity-and-access-management/managing-iam-for-your-enterprise/about-authentication-for-your-enterprise)."
    email *string
    // The user's username.
    login *string
    // Whether to set the user as suspended when the user is created.
    suspended *bool
}
// NewUsersPostRequestBody instantiates a new UsersPostRequestBody and sets the default values.
func NewUsersPostRequestBody()(*UsersPostRequestBody) {
    m := &UsersPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUsersPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUsersPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *UsersPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmail gets the email property value. **Required for built-in authentication.** The user's emailaddress. This parameter can be omitted when using CAS, LDAP, or SAML.For more information, see "[About authentication for your enterprise](https://docs.github.com/enterprise-server@3.14/admin/identity-and-access-management/managing-iam-for-your-enterprise/about-authentication-for-your-enterprise)."
// returns a *string when successful
func (m *UsersPostRequestBody) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UsersPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["login"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogin(val)
        }
        return nil
    }
    res["suspended"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuspended(val)
        }
        return nil
    }
    return res
}
// GetLogin gets the login property value. The user's username.
// returns a *string when successful
func (m *UsersPostRequestBody) GetLogin()(*string) {
    return m.login
}
// GetSuspended gets the suspended property value. Whether to set the user as suspended when the user is created.
// returns a *bool when successful
func (m *UsersPostRequestBody) GetSuspended()(*bool) {
    return m.suspended
}
// Serialize serializes information the current object
func (m *UsersPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("login", m.GetLogin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("suspended", m.GetSuspended())
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
func (m *UsersPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmail sets the email property value. **Required for built-in authentication.** The user's emailaddress. This parameter can be omitted when using CAS, LDAP, or SAML.For more information, see "[About authentication for your enterprise](https://docs.github.com/enterprise-server@3.14/admin/identity-and-access-management/managing-iam-for-your-enterprise/about-authentication-for-your-enterprise)."
func (m *UsersPostRequestBody) SetEmail(value *string)() {
    m.email = value
}
// SetLogin sets the login property value. The user's username.
func (m *UsersPostRequestBody) SetLogin(value *string)() {
    m.login = value
}
// SetSuspended sets the suspended property value. Whether to set the user as suspended when the user is created.
func (m *UsersPostRequestBody) SetSuspended(value *bool)() {
    m.suspended = value
}
type UsersPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmail()(*string)
    GetLogin()(*string)
    GetSuspended()(*bool)
    SetEmail(value *string)()
    SetLogin(value *string)()
    SetSuspended(value *bool)()
}
