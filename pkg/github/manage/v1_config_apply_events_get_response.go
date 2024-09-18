package manage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1ConfigApplyEventsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The nodes property
    nodes []V1ConfigApplyEventsGetResponse_nodesable
}
// NewV1ConfigApplyEventsGetResponse instantiates a new V1ConfigApplyEventsGetResponse and sets the default values.
func NewV1ConfigApplyEventsGetResponse()(*V1ConfigApplyEventsGetResponse) {
    m := &V1ConfigApplyEventsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1ConfigApplyEventsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1ConfigApplyEventsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1ConfigApplyEventsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1ConfigApplyEventsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1ConfigApplyEventsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["nodes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1ConfigApplyEventsGetResponse_nodesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1ConfigApplyEventsGetResponse_nodesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1ConfigApplyEventsGetResponse_nodesable)
                }
            }
            m.SetNodes(res)
        }
        return nil
    }
    return res
}
// GetNodes gets the nodes property value. The nodes property
// returns a []V1ConfigApplyEventsGetResponse_nodesable when successful
func (m *V1ConfigApplyEventsGetResponse) GetNodes()([]V1ConfigApplyEventsGetResponse_nodesable) {
    return m.nodes
}
// Serialize serializes information the current object
func (m *V1ConfigApplyEventsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetNodes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNodes()))
        for i, v := range m.GetNodes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("nodes", cast)
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
func (m *V1ConfigApplyEventsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNodes sets the nodes property value. The nodes property
func (m *V1ConfigApplyEventsGetResponse) SetNodes(value []V1ConfigApplyEventsGetResponse_nodesable)() {
    m.nodes = value
}
type V1ConfigApplyEventsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNodes()([]V1ConfigApplyEventsGetResponse_nodesable)
    SetNodes(value []V1ConfigApplyEventsGetResponse_nodesable)()
}
