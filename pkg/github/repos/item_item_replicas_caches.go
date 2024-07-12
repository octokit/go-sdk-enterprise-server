package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemReplicasCaches status for a cache replica
type ItemItemReplicasCaches struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The git property
    git ItemItemReplicasCaches_gitable
    // The host property
    host *string
    // The location property
    location *string
}
// NewItemItemReplicasCaches instantiates a new ItemItemReplicasCaches and sets the default values.
func NewItemItemReplicasCaches()(*ItemItemReplicasCaches) {
    m := &ItemItemReplicasCaches{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemReplicasCachesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemReplicasCachesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemReplicasCaches(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemReplicasCaches) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemReplicasCaches) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["git"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemReplicasCaches_gitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGit(val.(ItemItemReplicasCaches_gitable))
        }
        return nil
    }
    res["host"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHost(val)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val)
        }
        return nil
    }
    return res
}
// GetGit gets the git property value. The git property
// returns a ItemItemReplicasCaches_gitable when successful
func (m *ItemItemReplicasCaches) GetGit()(ItemItemReplicasCaches_gitable) {
    return m.git
}
// GetHost gets the host property value. The host property
// returns a *string when successful
func (m *ItemItemReplicasCaches) GetHost()(*string) {
    return m.host
}
// GetLocation gets the location property value. The location property
// returns a *string when successful
func (m *ItemItemReplicasCaches) GetLocation()(*string) {
    return m.location
}
// Serialize serializes information the current object
func (m *ItemItemReplicasCaches) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("git", m.GetGit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("host", m.GetHost())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("location", m.GetLocation())
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
func (m *ItemItemReplicasCaches) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGit sets the git property value. The git property
func (m *ItemItemReplicasCaches) SetGit(value ItemItemReplicasCaches_gitable)() {
    m.git = value
}
// SetHost sets the host property value. The host property
func (m *ItemItemReplicasCaches) SetHost(value *string)() {
    m.host = value
}
// SetLocation sets the location property value. The location property
func (m *ItemItemReplicasCaches) SetLocation(value *string)() {
    m.location = value
}
type ItemItemReplicasCachesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGit()(ItemItemReplicasCaches_gitable)
    GetHost()(*string)
    GetLocation()(*string)
    SetGit(value ItemItemReplicasCaches_gitable)()
    SetHost(value *string)()
    SetLocation(value *string)()
}
