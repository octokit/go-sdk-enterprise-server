package manage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// V1VersionRequestBuilder builds and executes requests for operations under \manage\v1\version
type V1VersionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1VersionRequestBuilderGetQueryParameters gets the GitHub Enterprise Server release versions that are currently installed on all available nodes. For more information, see "[GitHub Enterprise Server releases](https://docs.github.com/enterprise-server@3.11/admin/all-releases)."
type V1VersionRequestBuilderGetQueryParameters struct {
    // The cluster roles from the cluster configuration file.
    Cluster_roles *string `uriparametername:"cluster_roles"`
    // The UUID which identifies a node.
    Uuid *string `uriparametername:"uuid"`
}
// NewV1VersionRequestBuilderInternal instantiates a new V1VersionRequestBuilder and sets the default values.
func NewV1VersionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1VersionRequestBuilder) {
    m := &V1VersionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/manage/v1/version{?cluster_roles*,uuid*}", pathParameters),
    }
    return m
}
// NewV1VersionRequestBuilder instantiates a new V1VersionRequestBuilder and sets the default values.
func NewV1VersionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1VersionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1VersionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the GitHub Enterprise Server release versions that are currently installed on all available nodes. For more information, see "[GitHub Enterprise Server releases](https://docs.github.com/enterprise-server@3.11/admin/all-releases)."
// returns a []GhesVersionable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.11/rest/enterprise-admin/manage-ghes#get-all-ghes-release-versions-for-all-nodes
func (m *V1VersionRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V1VersionRequestBuilderGetQueryParameters])([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.GhesVersionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateGhesVersionFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.GhesVersionable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.GhesVersionable)
        }
    }
    return val, nil
}
// ToGetRequestInformation gets the GitHub Enterprise Server release versions that are currently installed on all available nodes. For more information, see "[GitHub Enterprise Server releases](https://docs.github.com/enterprise-server@3.11/admin/all-releases)."
// returns a *RequestInformation when successful
func (m *V1VersionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V1VersionRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1VersionRequestBuilder when successful
func (m *V1VersionRequestBuilder) WithUrl(rawUrl string)(*V1VersionRequestBuilder) {
    return NewV1VersionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
