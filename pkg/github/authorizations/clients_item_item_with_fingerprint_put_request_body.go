package authorizations

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ClientsItemItemWithFingerprintPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OAuth app client secret for which to create the token.
    client_secret *string
    // A note to remind you what the OAuth token is for.
    note *string
    // A URL to remind you what app the OAuth token is for.
    note_url *string
    // A list of scopes that this authorization is in.
    scopes []string
}
// NewClientsItemItemWithFingerprintPutRequestBody instantiates a new ClientsItemItemWithFingerprintPutRequestBody and sets the default values.
func NewClientsItemItemWithFingerprintPutRequestBody()(*ClientsItemItemWithFingerprintPutRequestBody) {
    m := &ClientsItemItemWithFingerprintPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateClientsItemItemWithFingerprintPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateClientsItemItemWithFingerprintPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClientsItemItemWithFingerprintPutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ClientsItemItemWithFingerprintPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClientSecret gets the client_secret property value. The OAuth app client secret for which to create the token.
// returns a *string when successful
func (m *ClientsItemItemWithFingerprintPutRequestBody) GetClientSecret()(*string) {
    return m.client_secret
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ClientsItemItemWithFingerprintPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetNote gets the note property value. A note to remind you what the OAuth token is for.
// returns a *string when successful
func (m *ClientsItemItemWithFingerprintPutRequestBody) GetNote()(*string) {
    return m.note
}
// GetNoteUrl gets the note_url property value. A URL to remind you what app the OAuth token is for.
// returns a *string when successful
func (m *ClientsItemItemWithFingerprintPutRequestBody) GetNoteUrl()(*string) {
    return m.note_url
}
// GetScopes gets the scopes property value. A list of scopes that this authorization is in.
// returns a []string when successful
func (m *ClientsItemItemWithFingerprintPutRequestBody) GetScopes()([]string) {
    return m.scopes
}
// Serialize serializes information the current object
func (m *ClientsItemItemWithFingerprintPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("client_secret", m.GetClientSecret())
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
func (m *ClientsItemItemWithFingerprintPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClientSecret sets the client_secret property value. The OAuth app client secret for which to create the token.
func (m *ClientsItemItemWithFingerprintPutRequestBody) SetClientSecret(value *string)() {
    m.client_secret = value
}
// SetNote sets the note property value. A note to remind you what the OAuth token is for.
func (m *ClientsItemItemWithFingerprintPutRequestBody) SetNote(value *string)() {
    m.note = value
}
// SetNoteUrl sets the note_url property value. A URL to remind you what app the OAuth token is for.
func (m *ClientsItemItemWithFingerprintPutRequestBody) SetNoteUrl(value *string)() {
    m.note_url = value
}
// SetScopes sets the scopes property value. A list of scopes that this authorization is in.
func (m *ClientsItemItemWithFingerprintPutRequestBody) SetScopes(value []string)() {
    m.scopes = value
}
type ClientsItemItemWithFingerprintPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientSecret()(*string)
    GetNote()(*string)
    GetNoteUrl()(*string)
    GetScopes()([]string)
    SetClientSecret(value *string)()
    SetNote(value *string)()
    SetNoteUrl(value *string)()
    SetScopes(value []string)()
}
