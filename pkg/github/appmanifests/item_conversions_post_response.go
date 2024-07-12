package appmanifests

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

type ItemConversionsPostResponse struct {
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Integration
}
// NewItemConversionsPostResponse instantiates a new ItemConversionsPostResponse and sets the default values.
func NewItemConversionsPostResponse()(*ItemConversionsPostResponse) {
    m := &ItemConversionsPostResponse{
        Integration: *ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.NewIntegration(),
    }
    return m
}
// CreateItemConversionsPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemConversionsPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemConversionsPostResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemConversionsPostResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Integration.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ItemConversionsPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Integration.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type ItemConversionsPostResponseable interface {
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Integrationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
