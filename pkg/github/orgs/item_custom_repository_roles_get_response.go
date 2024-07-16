package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

type ItemCustomRepositoryRolesGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The custom_roles property
    custom_roles []ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.OrganizationCustomRepositoryRoleable
    // The number of custom roles in this organization
    total_count *int32
}
// NewItemCustomRepositoryRolesGetResponse instantiates a new ItemCustomRepositoryRolesGetResponse and sets the default values.
func NewItemCustomRepositoryRolesGetResponse()(*ItemCustomRepositoryRolesGetResponse) {
    m := &ItemCustomRepositoryRolesGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCustomRepositoryRolesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCustomRepositoryRolesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCustomRepositoryRolesGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemCustomRepositoryRolesGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustomRoles gets the custom_roles property value. The custom_roles property
// returns a []OrganizationCustomRepositoryRoleable when successful
func (m *ItemCustomRepositoryRolesGetResponse) GetCustomRoles()([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.OrganizationCustomRepositoryRoleable) {
    return m.custom_roles
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemCustomRepositoryRolesGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["custom_roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateOrganizationCustomRepositoryRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.OrganizationCustomRepositoryRoleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.OrganizationCustomRepositoryRoleable)
                }
            }
            m.SetCustomRoles(res)
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
// GetTotalCount gets the total_count property value. The number of custom roles in this organization
// returns a *int32 when successful
func (m *ItemCustomRepositoryRolesGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemCustomRepositoryRolesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCustomRoles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomRoles()))
        for i, v := range m.GetCustomRoles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("custom_roles", cast)
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
func (m *ItemCustomRepositoryRolesGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustomRoles sets the custom_roles property value. The custom_roles property
func (m *ItemCustomRepositoryRolesGetResponse) SetCustomRoles(value []ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.OrganizationCustomRepositoryRoleable)() {
    m.custom_roles = value
}
// SetTotalCount sets the total_count property value. The number of custom roles in this organization
func (m *ItemCustomRepositoryRolesGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
type ItemCustomRepositoryRolesGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomRoles()([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.OrganizationCustomRepositoryRoleable)
    GetTotalCount()(*int32)
    SetCustomRoles(value []ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.OrganizationCustomRepositoryRoleable)()
    SetTotalCount(value *int32)()
}
