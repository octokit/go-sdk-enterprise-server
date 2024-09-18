package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActionsCacheUsagePolicyEnterprise gitHub Actions cache usage policy for an enterprise.
type ActionsCacheUsagePolicyEnterprise struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // For repositories in an enterprise, the maximum value that can be set as the limit for the sum of all caches in a repository, in gigabytes.
    max_repo_cache_size_limit_in_gb *int32
    // For repositories in an enterprise, the default size limit for the sum of all caches in a repository, in gigabytes.
    repo_cache_size_limit_in_gb *int32
}
// NewActionsCacheUsagePolicyEnterprise instantiates a new ActionsCacheUsagePolicyEnterprise and sets the default values.
func NewActionsCacheUsagePolicyEnterprise()(*ActionsCacheUsagePolicyEnterprise) {
    m := &ActionsCacheUsagePolicyEnterprise{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActionsCacheUsagePolicyEnterpriseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActionsCacheUsagePolicyEnterpriseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActionsCacheUsagePolicyEnterprise(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActionsCacheUsagePolicyEnterprise) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActionsCacheUsagePolicyEnterprise) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["max_repo_cache_size_limit_in_gb"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxRepoCacheSizeLimitInGb(val)
        }
        return nil
    }
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
// GetMaxRepoCacheSizeLimitInGb gets the max_repo_cache_size_limit_in_gb property value. For repositories in an enterprise, the maximum value that can be set as the limit for the sum of all caches in a repository, in gigabytes.
// returns a *int32 when successful
func (m *ActionsCacheUsagePolicyEnterprise) GetMaxRepoCacheSizeLimitInGb()(*int32) {
    return m.max_repo_cache_size_limit_in_gb
}
// GetRepoCacheSizeLimitInGb gets the repo_cache_size_limit_in_gb property value. For repositories in an enterprise, the default size limit for the sum of all caches in a repository, in gigabytes.
// returns a *int32 when successful
func (m *ActionsCacheUsagePolicyEnterprise) GetRepoCacheSizeLimitInGb()(*int32) {
    return m.repo_cache_size_limit_in_gb
}
// Serialize serializes information the current object
func (m *ActionsCacheUsagePolicyEnterprise) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("max_repo_cache_size_limit_in_gb", m.GetMaxRepoCacheSizeLimitInGb())
        if err != nil {
            return err
        }
    }
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
func (m *ActionsCacheUsagePolicyEnterprise) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMaxRepoCacheSizeLimitInGb sets the max_repo_cache_size_limit_in_gb property value. For repositories in an enterprise, the maximum value that can be set as the limit for the sum of all caches in a repository, in gigabytes.
func (m *ActionsCacheUsagePolicyEnterprise) SetMaxRepoCacheSizeLimitInGb(value *int32)() {
    m.max_repo_cache_size_limit_in_gb = value
}
// SetRepoCacheSizeLimitInGb sets the repo_cache_size_limit_in_gb property value. For repositories in an enterprise, the default size limit for the sum of all caches in a repository, in gigabytes.
func (m *ActionsCacheUsagePolicyEnterprise) SetRepoCacheSizeLimitInGb(value *int32)() {
    m.repo_cache_size_limit_in_gb = value
}
type ActionsCacheUsagePolicyEnterpriseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaxRepoCacheSizeLimitInGb()(*int32)
    GetRepoCacheSizeLimitInGb()(*int32)
    SetMaxRepoCacheSizeLimitInGb(value *int32)()
    SetRepoCacheSizeLimitInGb(value *int32)()
}
