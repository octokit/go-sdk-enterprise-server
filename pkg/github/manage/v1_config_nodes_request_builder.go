package manage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// V1ConfigNodesRequestBuilder builds and executes requests for operations under \manage\v1\config\nodes
type V1ConfigNodesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1ConfigNodesRequestBuilderGetQueryParameters get node metadata for all configured nodes in the current cluster. For more information, see "[About clustering](https://docs.github.com/enterprise-server@3.12/admin/enterprise-management/configuring-clustering/about-clustering)."
type V1ConfigNodesRequestBuilderGetQueryParameters struct {
    // The cluster roles from the cluster configuration file.
    Cluster_roles *string `uriparametername:"cluster_roles"`
    // The UUID which identifies a node.
    Uuid *string `uriparametername:"uuid"`
}
// NewV1ConfigNodesRequestBuilderInternal instantiates a new V1ConfigNodesRequestBuilder and sets the default values.
func NewV1ConfigNodesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ConfigNodesRequestBuilder) {
    m := &V1ConfigNodesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/manage/v1/config/nodes{?cluster_roles*,uuid*}", pathParameters),
    }
    return m
}
// NewV1ConfigNodesRequestBuilder instantiates a new V1ConfigNodesRequestBuilder and sets the default values.
func NewV1ConfigNodesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ConfigNodesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1ConfigNodesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get node metadata for all configured nodes in the current cluster. For more information, see "[About clustering](https://docs.github.com/enterprise-server@3.12/admin/enterprise-management/configuring-clustering/about-clustering)."
// returns a GhesConfigNodesable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.12/rest/enterprise-admin/manage-ghes#get-ghes-node-metadata-for-all-nodes
func (m *V1ConfigNodesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V1ConfigNodesRequestBuilderGetQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.GhesConfigNodesable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateGhesConfigNodesFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.GhesConfigNodesable), nil
}
// ToGetRequestInformation get node metadata for all configured nodes in the current cluster. For more information, see "[About clustering](https://docs.github.com/enterprise-server@3.12/admin/enterprise-management/configuring-clustering/about-clustering)."
// returns a *RequestInformation when successful
func (m *V1ConfigNodesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V1ConfigNodesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1ConfigNodesRequestBuilder when successful
func (m *V1ConfigNodesRequestBuilder) WithUrl(rawUrl string)(*V1ConfigNodesRequestBuilder) {
    return NewV1ConfigNodesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
