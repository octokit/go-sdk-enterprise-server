package manage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// V1ReplicationStatusRequestBuilder builds and executes requests for operations under \manage\v1\replication\status
type V1ReplicationStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1ReplicationStatusRequestBuilderGetQueryParameters gets the status of all services running on each replica node.This endpoint may take several seconds to reply.
type V1ReplicationStatusRequestBuilderGetQueryParameters struct {
    // The cluster roles from the cluster configuration file.
    Cluster_roles *string `uriparametername:"cluster_roles"`
    // The UUID which identifies a node.
    Uuid *string `uriparametername:"uuid"`
}
// NewV1ReplicationStatusRequestBuilderInternal instantiates a new V1ReplicationStatusRequestBuilder and sets the default values.
func NewV1ReplicationStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ReplicationStatusRequestBuilder) {
    m := &V1ReplicationStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/manage/v1/replication/status{?cluster_roles*,uuid*}", pathParameters),
    }
    return m
}
// NewV1ReplicationStatusRequestBuilder instantiates a new V1ReplicationStatusRequestBuilder and sets the default values.
func NewV1ReplicationStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ReplicationStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1ReplicationStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the status of all services running on each replica node.This endpoint may take several seconds to reply.
// returns a GhesReplicationStatusable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.11/rest/enterprise-admin/manage-ghes#get-the-status-of-services-running-on-all-replica-nodes
func (m *V1ReplicationStatusRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V1ReplicationStatusRequestBuilderGetQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.GhesReplicationStatusable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateGhesReplicationStatusFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.GhesReplicationStatusable), nil
}
// ToGetRequestInformation gets the status of all services running on each replica node.This endpoint may take several seconds to reply.
// returns a *RequestInformation when successful
func (m *V1ReplicationStatusRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V1ReplicationStatusRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1ReplicationStatusRequestBuilder when successful
func (m *V1ReplicationStatusRequestBuilder) WithUrl(rawUrl string)(*V1ReplicationStatusRequestBuilder) {
    return NewV1ReplicationStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
