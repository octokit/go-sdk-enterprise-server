package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActionsCacheUsagePolicyForRepository gitHub Actions cache usage policy for repository.
type ActionsCacheUsagePolicyForRepository struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The size limit for the sum of all caches, in gigabytes.
    repo_cache_size_limit_in_gb *int32
}
// NewActionsCacheUsagePolicyForRepository instantiates a new ActionsCacheUsagePolicyForRepository and sets the default values.
func NewActionsCacheUsagePolicyForRepository()(*ActionsCacheUsagePolicyForRepository) {
    m := &ActionsCacheUsagePolicyForRepository{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActionsCacheUsagePolicyForRepositoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActionsCacheUsagePolicyForRepositoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActionsCacheUsagePolicyForRepository(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActionsCacheUsagePolicyForRepository) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActionsCacheUsagePolicyForRepository) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["repo_cache_size_limit_in_gb"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepoCacheSizeLimitInGb(val)
        }
        return nil
    }
    return res
}
// GetRepoCacheSizeLimitInGb gets the repo_cache_size_limit_in_gb property value. The size limit for the sum of all caches, in gigabytes.
// returns a *int32 when successful
func (m *ActionsCacheUsagePolicyForRepository) GetRepoCacheSizeLimitInGb()(*int32) {
    return m.repo_cache_size_limit_in_gb
}
// Serialize serializes information the current object
func (m *ActionsCacheUsagePolicyForRepository) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("repo_cache_size_limit_in_gb", m.GetRepoCacheSizeLimitInGb())
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
func (m *ActionsCacheUsagePolicyForRepository) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRepoCacheSizeLimitInGb sets the repo_cache_size_limit_in_gb property value. The size limit for the sum of all caches, in gigabytes.
func (m *ActionsCacheUsagePolicyForRepository) SetRepoCacheSizeLimitInGb(value *int32)() {
    m.repo_cache_size_limit_in_gb = value
}
type ActionsCacheUsagePolicyForRepositoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRepoCacheSizeLimitInGb()(*int32)
    SetRepoCacheSizeLimitInGb(value *int32)()
}
