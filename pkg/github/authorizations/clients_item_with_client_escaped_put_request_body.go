package authorizations

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ClientsItemWithClient_PutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
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
// NewClientsItemWithClient_PutRequestBody instantiates a new ClientsItemWithClient_PutRequestBody and sets the default values.
func NewClientsItemWithClient_PutRequestBody()(*ClientsItemWithClient_PutRequestBody) {
    m := &ClientsItemWithClient_PutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateClientsItemWithClient_PutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateClientsItemWithClient_PutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClientsItemWithClient_PutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ClientsItemWithClient_PutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClientSecret gets the client_secret property value. The OAuth app client secret for which to create the token.
// returns a *string when successful
func (m *ClientsItemWithClient_PutRequestBody) GetClientSecret()(*string) {
    return m.client_secret
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ClientsItemWithClient_PutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
func (m *ClientsItemWithClient_PutRequestBody) GetFingerprint()(*string) {
    return m.fingerprint
}
// GetNote gets the note property value. A note to remind you what the OAuth token is for.
// returns a *string when successful
func (m *ClientsItemWithClient_PutRequestBody) GetNote()(*string) {
    return m.note
}
// GetNoteUrl gets the note_url property value. A URL to remind you what app the OAuth token is for.
// returns a *string when successful
func (m *ClientsItemWithClient_PutRequestBody) GetNoteUrl()(*string) {
    return m.note_url
}
// GetScopes gets the scopes property value. A list of scopes that this authorization is in.
// returns a []string when successful
func (m *ClientsItemWithClient_PutRequestBody) GetScopes()([]string) {
    return m.scopes
}
// Serialize serializes information the current object
func (m *ClientsItemWithClient_PutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ClientsItemWithClient_PutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClientSecret sets the client_secret property value. The OAuth app client secret for which to create the token.
func (m *ClientsItemWithClient_PutRequestBody) SetClientSecret(value *string)() {
    m.client_secret = value
}
// SetFingerprint sets the fingerprint property value. A unique string to distinguish an authorization from others created for the same client ID and user.
func (m *ClientsItemWithClient_PutRequestBody) SetFingerprint(value *string)() {
    m.fingerprint = value
}
// SetNote sets the note property value. A note to remind you what the OAuth token is for.
func (m *ClientsItemWithClient_PutRequestBody) SetNote(value *string)() {
    m.note = value
}
// SetNoteUrl sets the note_url property value. A URL to remind you what app the OAuth token is for.
func (m *ClientsItemWithClient_PutRequestBody) SetNoteUrl(value *string)() {
    m.note_url = value
}
// SetScopes sets the scopes property value. A list of scopes that this authorization is in.
func (m *ClientsItemWithClient_PutRequestBody) SetScopes(value []string)() {
    m.scopes = value
}
type ClientsItemWithClient_PutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientSecret()(*string)
    GetFingerprint()(*string)
    GetNote()(*string)
    GetNoteUrl()(*string)
    GetScopes()([]string)
    SetClientSecret(value *string)()
    SetFingerprint(value *string)()
    SetNote(value *string)()
    SetNoteUrl(value *string)()
    SetScopes(value []string)()
}
