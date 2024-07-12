package authorizations

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemWithAuthorization_PatchRequestBody struct {
    // A list of scopes to add to this authorization.
    add_scopes []string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A unique string to distinguish an authorization from others created for the same client ID and user.
    fingerprint *string
    // A note to remind you what the OAuth token is for.
    note *string
    // A URL to remind you what app the OAuth token is for.
    note_url *string
    // A list of scopes to remove from this authorization.
    remove_scopes []string
    // A list of scopes that this authorization is in.
    scopes []string
}
// NewItemWithAuthorization_PatchRequestBody instantiates a new ItemWithAuthorization_PatchRequestBody and sets the default values.
func NewItemWithAuthorization_PatchRequestBody()(*ItemWithAuthorization_PatchRequestBody) {
    m := &ItemWithAuthorization_PatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemWithAuthorization_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemWithAuthorization_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemWithAuthorization_PatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemWithAuthorization_PatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddScopes gets the add_scopes property value. A list of scopes to add to this authorization.
// returns a []string when successful
func (m *ItemWithAuthorization_PatchRequestBody) GetAddScopes()([]string) {
    return m.add_scopes
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemWithAuthorization_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["add_scopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAddScopes(res)
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
    res["remove_scopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRemoveScopes(res)
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
func (m *ItemWithAuthorization_PatchRequestBody) GetFingerprint()(*string) {
    return m.fingerprint
}
// GetNote gets the note property value. A note to remind you what the OAuth token is for.
// returns a *string when successful
func (m *ItemWithAuthorization_PatchRequestBody) GetNote()(*string) {
    return m.note
}
// GetNoteUrl gets the note_url property value. A URL to remind you what app the OAuth token is for.
// returns a *string when successful
func (m *ItemWithAuthorization_PatchRequestBody) GetNoteUrl()(*string) {
    return m.note_url
}
// GetRemoveScopes gets the remove_scopes property value. A list of scopes to remove from this authorization.
// returns a []string when successful
func (m *ItemWithAuthorization_PatchRequestBody) GetRemoveScopes()([]string) {
    return m.remove_scopes
}
// GetScopes gets the scopes property value. A list of scopes that this authorization is in.
// returns a []string when successful
func (m *ItemWithAuthorization_PatchRequestBody) GetScopes()([]string) {
    return m.scopes
}
// Serialize serializes information the current object
func (m *ItemWithAuthorization_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAddScopes() != nil {
        err := writer.WriteCollectionOfStringValues("add_scopes", m.GetAddScopes())
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
    if m.GetRemoveScopes() != nil {
        err := writer.WriteCollectionOfStringValues("remove_scopes", m.GetRemoveScopes())
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
func (m *ItemWithAuthorization_PatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddScopes sets the add_scopes property value. A list of scopes to add to this authorization.
func (m *ItemWithAuthorization_PatchRequestBody) SetAddScopes(value []string)() {
    m.add_scopes = value
}
// SetFingerprint sets the fingerprint property value. A unique string to distinguish an authorization from others created for the same client ID and user.
func (m *ItemWithAuthorization_PatchRequestBody) SetFingerprint(value *string)() {
    m.fingerprint = value
}
// SetNote sets the note property value. A note to remind you what the OAuth token is for.
func (m *ItemWithAuthorization_PatchRequestBody) SetNote(value *string)() {
    m.note = value
}
// SetNoteUrl sets the note_url property value. A URL to remind you what app the OAuth token is for.
func (m *ItemWithAuthorization_PatchRequestBody) SetNoteUrl(value *string)() {
    m.note_url = value
}
// SetRemoveScopes sets the remove_scopes property value. A list of scopes to remove from this authorization.
func (m *ItemWithAuthorization_PatchRequestBody) SetRemoveScopes(value []string)() {
    m.remove_scopes = value
}
// SetScopes sets the scopes property value. A list of scopes that this authorization is in.
func (m *ItemWithAuthorization_PatchRequestBody) SetScopes(value []string)() {
    m.scopes = value
}
type ItemWithAuthorization_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddScopes()([]string)
    GetFingerprint()(*string)
    GetNote()(*string)
    GetNoteUrl()(*string)
    GetRemoveScopes()([]string)
    GetScopes()([]string)
    SetAddScopes(value []string)()
    SetFingerprint(value *string)()
    SetNote(value *string)()
    SetNoteUrl(value *string)()
    SetRemoveScopes(value []string)()
    SetScopes(value []string)()
}
