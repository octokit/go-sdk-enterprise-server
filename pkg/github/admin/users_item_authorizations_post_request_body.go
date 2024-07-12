package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UsersItemAuthorizationsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A list of [scopes](https://docs.github.com/enterprise-server@3.13/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).
    scopes []string
}
// NewUsersItemAuthorizationsPostRequestBody instantiates a new UsersItemAuthorizationsPostRequestBody and sets the default values.
func NewUsersItemAuthorizationsPostRequestBody()(*UsersItemAuthorizationsPostRequestBody) {
    m := &UsersItemAuthorizationsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUsersItemAuthorizationsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUsersItemAuthorizationsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemAuthorizationsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *UsersItemAuthorizationsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UsersItemAuthorizationsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["scopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetScopes(res)
        }
        return nil
    }
    return res
}
// GetScopes gets the scopes property value. A list of [scopes](https://docs.github.com/enterprise-server@3.13/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).
// returns a []string when successful
func (m *UsersItemAuthorizationsPostRequestBody) GetScopes()([]string) {
    return m.scopes
}
// Serialize serializes information the current object
func (m *UsersItemAuthorizationsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetScopes() != nil {
        err := writer.WriteCollectionOfStringValues("scopes", m.GetScopes())
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
func (m *UsersItemAuthorizationsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetScopes sets the scopes property value. A list of [scopes](https://docs.github.com/enterprise-server@3.13/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).
func (m *UsersItemAuthorizationsPostRequestBody) SetScopes(value []string)() {
    m.scopes = value
}
type UsersItemAuthorizationsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetScopes()([]string)
    SetScopes(value []string)()
}
