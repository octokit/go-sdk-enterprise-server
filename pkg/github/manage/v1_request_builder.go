package manage

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1RequestBuilder builds and executes requests for operations under \manage\v1
type V1RequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Access the access property
// returns a *V1AccessRequestBuilder when successful
func (m *V1RequestBuilder) Access()(*V1AccessRequestBuilder) {
    return NewV1AccessRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Checks the checks property
// returns a *V1ChecksRequestBuilder when successful
func (m *V1RequestBuilder) Checks()(*V1ChecksRequestBuilder) {
    return NewV1ChecksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cluster the cluster property
// returns a *V1ClusterRequestBuilder when successful
func (m *V1RequestBuilder) Cluster()(*V1ClusterRequestBuilder) {
    return NewV1ClusterRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Config the config property
// returns a *V1ConfigRequestBuilder when successful
func (m *V1RequestBuilder) Config()(*V1ConfigRequestBuilder) {
    return NewV1ConfigRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1RequestBuilderInternal instantiates a new V1RequestBuilder and sets the default values.
func NewV1RequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1RequestBuilder) {
    m := &V1RequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/manage/v1", pathParameters),
    }
    return m
}
// NewV1RequestBuilder instantiates a new V1RequestBuilder and sets the default values.
func NewV1RequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1RequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1RequestBuilderInternal(urlParams, requestAdapter)
}
// Maintenance the maintenance property
// returns a *V1MaintenanceRequestBuilder when successful
func (m *V1RequestBuilder) Maintenance()(*V1MaintenanceRequestBuilder) {
    return NewV1MaintenanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Replication the replication property
// returns a *V1ReplicationRequestBuilder when successful
func (m *V1RequestBuilder) Replication()(*V1ReplicationRequestBuilder) {
    return NewV1ReplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Version the version property
// returns a *V1VersionRequestBuilder when successful
func (m *V1RequestBuilder) Version()(*V1VersionRequestBuilder) {
    return NewV1VersionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
