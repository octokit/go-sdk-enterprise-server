package authorizations

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthorizationsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OAuth app client key for which to create the token.
    client_id *string
    // The OAuth app client secret for which to create the token.
    client_secret *string
    // A unique string to distinguish an authorization from others created for the same client ID and user.
    fingerprint *string
    // A note to remind you what the OAuth token is for.
    note *string
    // A URL to remind you what app the OAuth token is for.
    note_url *string
    // A list of scopes that this authorization is in.
    scopes []string
}
// NewAuthorizationsPostRequestBody instantiates a new AuthorizationsPostRequestBody and sets the default values.
func NewAuthorizationsPostRequestBody()(*AuthorizationsPostRequestBody) {
    m := &AuthorizationsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthorizationsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthorizationsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthorizationsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AuthorizationsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClientId gets the client_id property value. The OAuth app client key for which to create the token.
// returns a *string when successful
func (m *AuthorizationsPostRequestBody) GetClientId()(*string) {
    return m.client_id
}
// GetClientSecret gets the client_secret property value. The OAuth app client secret for which to create the token.
// returns a *string when successful
func (m *AuthorizationsPostRequestBody) GetClientSecret()(*string) {
    return m.client_secret
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthorizationsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["client_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientId(val)
        }
        return nil
    }
    res["client_secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientSecret(val)
        }
        return nil
    }
    res["fingerprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFingerprint(val)
        }
        return nil
    }
    res["note"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNote(val)
        }
        return nil
    }
    res["note_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNoteUrl(val)
        }
        return nil
    }
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
// GetFingerprint gets the fingerprint property value. A unique string to distinguish an authorization from others created for the same client ID and user.
// returns a *string when successful
func (m *AuthorizationsPostRequestBody) GetFingerprint()(*string) {
    return m.fingerprint
}
// GetNote gets the note property value. A note to remind you what the OAuth token is for.
// returns a *string when successful
func (m *AuthorizationsPostRequestBody) GetNote()(*string) {
    return m.note
}
// GetNoteUrl gets the note_url property value. A URL to remind you what app the OAuth token is for.
// returns a *string when successful
func (m *AuthorizationsPostRequestBody) GetNoteUrl()(*string) {
    return m.note_url
}
// GetScopes gets the scopes property value. A list of scopes that this authorization is in.
// returns a []string when successful
func (m *AuthorizationsPostRequestBody) GetScopes()([]string) {
    return m.scopes
}
// Serialize serializes information the current object
func (m *AuthorizationsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("client_id", m.GetClientId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("client_secret", m.GetClientSecret())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fingerprint", m.GetFingerprint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("note", m.GetNote())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("note_url", m.GetNoteUrl())
        if err != nil {
            return err
        }
    }
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
func (m *AuthorizationsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClientId sets the client_id property value. The OAuth app client key for which to create the token.
func (m *AuthorizationsPostRequestBody) SetClientId(value *string)() {
    m.client_id = value
}
// SetClientSecret sets the client_secret property value. The OAuth app client secret for which to create the token.
func (m *AuthorizationsPostRequestBody) SetClientSecret(value *string)() {
    m.client_secret = value
}
// SetFingerprint sets the fingerprint property value. A unique string to distinguish an authorization from others created for the same client ID and user.
func (m *AuthorizationsPostRequestBody) SetFingerprint(value *string)() {
    m.fingerprint = value
}
// SetNote sets the note property value. A note to remind you what the OAuth token is for.
func (m *AuthorizationsPostRequestBody) SetNote(value *string)() {
    m.note = value
}
// SetNoteUrl sets the note_url property value. A URL to remind you what app the OAuth token is for.
func (m *AuthorizationsPostRequestBody) SetNoteUrl(value *string)() {
    m.note_url = value
}
// SetScopes sets the scopes property value. A list of scopes that this authorization is in.
func (m *AuthorizationsPostRequestBody) SetScopes(value []string)() {
    m.scopes = value
}
type AuthorizationsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientId()(*string)
    GetClientSecret()(*string)
    GetFingerprint()(*string)
    GetNote()(*string)
    GetNoteUrl()(*string)
    GetScopes()([]string)
    SetClientId(value *string)()
    SetClientSecret(value *string)()
    SetFingerprint(value *string)()
    SetNote(value *string)()
    SetNoteUrl(value *string)()
    SetScopes(value []string)()
}
