package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

type ItemItemCheckSuitesItemCheckRunsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The check_runs property
    check_runs []ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CheckRunable
    // The total_count property
    total_count *int32
}
// NewItemItemCheckSuitesItemCheckRunsGetResponse instantiates a new ItemItemCheckSuitesItemCheckRunsGetResponse and sets the default values.
func NewItemItemCheckSuitesItemCheckRunsGetResponse()(*ItemItemCheckSuitesItemCheckRunsGetResponse) {
    m := &ItemItemCheckSuitesItemCheckRunsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemCheckSuitesItemCheckRunsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemCheckSuitesItemCheckRunsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCheckSuitesItemCheckRunsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemCheckSuitesItemCheckRunsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCheckRuns gets the check_runs property value. The check_runs property
// returns a []CheckRunable when successful
func (m *ItemItemCheckSuitesItemCheckRunsGetResponse) GetCheckRuns()([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CheckRunable) {
    return m.check_runs
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemCheckSuitesItemCheckRunsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["check_runs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateCheckRunFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CheckRunable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CheckRunable)
                }
            }
            m.SetCheckRuns(res)
        }
        return nil
    }
    res["total_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCount(val)
        }
        return nil
    }
    return res
}
// GetTotalCount gets the total_count property value. The total_count property
// returns a *int32 when successful
func (m *ItemItemCheckSuitesItemCheckRunsGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemItemCheckSuitesItemCheckRunsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCheckRuns() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCheckRuns()))
        for i, v := range m.GetCheckRuns() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("check_runs", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_count", m.GetTotalCount())
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
func (m *ItemItemCheckSuitesItemCheckRunsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCheckRuns sets the check_runs property value. The check_runs property
func (m *ItemItemCheckSuitesItemCheckRunsGetResponse) SetCheckRuns(value []ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CheckRunable)() {
    m.check_runs = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemItemCheckSuitesItemCheckRunsGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
type ItemItemCheckSuitesItemCheckRunsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCheckRuns()([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CheckRunable)
    GetTotalCount()(*int32)
    SetCheckRuns(value []ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CheckRunable)()
    SetTotalCount(value *int32)()
}
