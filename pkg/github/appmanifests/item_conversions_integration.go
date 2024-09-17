package appmanifests

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

type ItemConversionsIntegration struct {
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Integration
    // The client_id property
    client_id *string
}
// NewItemConversionsIntegration instantiates a new ItemConversionsIntegration and sets the default values.
func NewItemConversionsIntegration()(*ItemConversionsIntegration) {
    m := &ItemConversionsIntegration{
        Integration: *ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.NewIntegration(),
    }
    return m
}
// CreateItemConversionsIntegrationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemConversionsIntegrationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemConversionsIntegration(), nil
}
// GetClientId gets the client_id property value. The client_id property
// returns a *string when successful
func (m *ItemConversionsIntegration) GetClientId()(*string) {
    return m.client_id
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemConversionsIntegration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Integration.GetFieldDeserializers()
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
    return res
}
// Serialize serializes information the current object
func (m *ItemConversionsIntegration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Integration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("client_id", m.GetClientId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClientId sets the client_id property value. The client_id property
func (m *ItemConversionsIntegration) SetClientId(value *string)() {
    m.client_id = value
}
type ItemConversionsIntegrationable interface {
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Integrationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientId()(*string)
    SetClientId(value *string)()
}
