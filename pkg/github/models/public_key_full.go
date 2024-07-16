package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PublicKeyFull struct {
    // The added_by property
    added_by *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The id property
    id *int64
    // The key property
    key *string
    // The last_used property
    last_used *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The read_only property
    read_only *bool
    // The repository_id property
    repository_id *int64
    // The title property
    title *string
    // The url property
    url *string
    // The user_id property
    user_id *int32
    // The verified property
    verified *bool
}
// NewPublicKeyFull instantiates a new PublicKeyFull and sets the default values.
func NewPublicKeyFull()(*PublicKeyFull) {
    m := &PublicKeyFull{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePublicKeyFullFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePublicKeyFullFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPublicKeyFull(), nil
}
// GetAddedBy gets the added_by property value. The added_by property
// returns a *string when successful
func (m *PublicKeyFull) GetAddedBy()(*string) {
    return m.added_by
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PublicKeyFull) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *PublicKeyFull) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PublicKeyFull) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["added_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddedBy(val)
        }
        return nil
    }
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
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
    res["last_used"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUsed(val)
        }
        return nil
    }
    res["read_only"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReadOnly(val)
        }
        return nil
    }
    res["repository_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryId(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    res["user_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["verified"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerified(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *PublicKeyFull) GetId()(*int64) {
    return m.id
}
// GetKey gets the key property value. The key property
// returns a *string when successful
func (m *PublicKeyFull) GetKey()(*string) {
    return m.key
}
// GetLastUsed gets the last_used property value. The last_used property
// returns a *Time when successful
func (m *PublicKeyFull) GetLastUsed()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.last_used
}
// GetReadOnly gets the read_only property value. The read_only property
// returns a *bool when successful
func (m *PublicKeyFull) GetReadOnly()(*bool) {
    return m.read_only
}
// GetRepositoryId gets the repository_id property value. The repository_id property
// returns a *int64 when successful
func (m *PublicKeyFull) GetRepositoryId()(*int64) {
    return m.repository_id
}
// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *PublicKeyFull) GetTitle()(*string) {
    return m.title
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *PublicKeyFull) GetUrl()(*string) {
    return m.url
}
// GetUserId gets the user_id property value. The user_id property
// returns a *int32 when successful
func (m *PublicKeyFull) GetUserId()(*int32) {
    return m.user_id
}
// GetVerified gets the verified property value. The verified property
// returns a *bool when successful
func (m *PublicKeyFull) GetVerified()(*bool) {
    return m.verified
}
// Serialize serializes information the current object
func (m *PublicKeyFull) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("added_by", m.GetAddedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("last_used", m.GetLastUsed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("read_only", m.GetReadOnly())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("repository_id", m.GetRepositoryId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("user_id", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("verified", m.GetVerified())
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
// SetAddedBy sets the added_by property value. The added_by property
func (m *PublicKeyFull) SetAddedBy(value *string)() {
    m.added_by = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PublicKeyFull) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *PublicKeyFull) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetId sets the id property value. The id property
func (m *PublicKeyFull) SetId(value *int64)() {
    m.id = value
}
// SetKey sets the key property value. The key property
func (m *PublicKeyFull) SetKey(value *string)() {
    m.key = value
}
// SetLastUsed sets the last_used property value. The last_used property
func (m *PublicKeyFull) SetLastUsed(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.last_used = value
}
// SetReadOnly sets the read_only property value. The read_only property
func (m *PublicKeyFull) SetReadOnly(value *bool)() {
    m.read_only = value
}
// SetRepositoryId sets the repository_id property value. The repository_id property
func (m *PublicKeyFull) SetRepositoryId(value *int64)() {
    m.repository_id = value
}
// SetTitle sets the title property value. The title property
func (m *PublicKeyFull) SetTitle(value *string)() {
    m.title = value
}
// SetUrl sets the url property value. The url property
func (m *PublicKeyFull) SetUrl(value *string)() {
    m.url = value
}
// SetUserId sets the user_id property value. The user_id property
func (m *PublicKeyFull) SetUserId(value *int32)() {
    m.user_id = value
}
// SetVerified sets the verified property value. The verified property
func (m *PublicKeyFull) SetVerified(value *bool)() {
    m.verified = value
}
type PublicKeyFullable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddedBy()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*int64)
    GetKey()(*string)
    GetLastUsed()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReadOnly()(*bool)
    GetRepositoryId()(*int64)
    GetTitle()(*string)
    GetUrl()(*string)
    GetUserId()(*int32)
    GetVerified()(*bool)
    SetAddedBy(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *int64)()
    SetKey(value *string)()
    SetLastUsed(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReadOnly(value *bool)()
    SetRepositoryId(value *int64)()
    SetTitle(value *string)()
    SetUrl(value *string)()
    SetUserId(value *int32)()
    SetVerified(value *bool)()
}
