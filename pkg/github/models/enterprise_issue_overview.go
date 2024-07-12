package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EnterpriseIssueOverview struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The closed_issues property
    closed_issues *int32
    // The open_issues property
    open_issues *int32
    // The total_issues property
    total_issues *int32
}
// NewEnterpriseIssueOverview instantiates a new EnterpriseIssueOverview and sets the default values.
func NewEnterpriseIssueOverview()(*EnterpriseIssueOverview) {
    m := &EnterpriseIssueOverview{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnterpriseIssueOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnterpriseIssueOverviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnterpriseIssueOverview(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EnterpriseIssueOverview) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClosedIssues gets the closed_issues property value. The closed_issues property
// returns a *int32 when successful
func (m *EnterpriseIssueOverview) GetClosedIssues()(*int32) {
    return m.closed_issues
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnterpriseIssueOverview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["closed_issues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClosedIssues(val)
        }
        return nil
    }
    res["open_issues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOpenIssues(val)
        }
        return nil
    }
    res["total_issues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalIssues(val)
        }
        return nil
    }
    return res
}
// GetOpenIssues gets the open_issues property value. The open_issues property
// returns a *int32 when successful
func (m *EnterpriseIssueOverview) GetOpenIssues()(*int32) {
    return m.open_issues
}
// GetTotalIssues gets the total_issues property value. The total_issues property
// returns a *int32 when successful
func (m *EnterpriseIssueOverview) GetTotalIssues()(*int32) {
    return m.total_issues
}
// Serialize serializes information the current object
func (m *EnterpriseIssueOverview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("closed_issues", m.GetClosedIssues())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("open_issues", m.GetOpenIssues())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_issues", m.GetTotalIssues())
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
func (m *EnterpriseIssueOverview) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClosedIssues sets the closed_issues property value. The closed_issues property
func (m *EnterpriseIssueOverview) SetClosedIssues(value *int32)() {
    m.closed_issues = value
}
// SetOpenIssues sets the open_issues property value. The open_issues property
func (m *EnterpriseIssueOverview) SetOpenIssues(value *int32)() {
    m.open_issues = value
}
// SetTotalIssues sets the total_issues property value. The total_issues property
func (m *EnterpriseIssueOverview) SetTotalIssues(value *int32)() {
    m.total_issues = value
}
type EnterpriseIssueOverviewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClosedIssues()(*int32)
    GetOpenIssues()(*int32)
    GetTotalIssues()(*int32)
    SetClosedIssues(value *int32)()
    SetOpenIssues(value *int32)()
    SetTotalIssues(value *int32)()
}
