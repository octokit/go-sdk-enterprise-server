package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ItemItemReplicasCachesRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\replicas\caches
type ItemItemReplicasCachesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemReplicasCachesRequestBuilderGetQueryParameters lists the status of each repository cache replica.
type ItemItemReplicasCachesRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.11/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.11/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemItemReplicasCachesRequestBuilderInternal instantiates a new ItemItemReplicasCachesRequestBuilder and sets the default values.
func NewItemItemReplicasCachesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemReplicasCachesRequestBuilder) {
    m := &ItemItemReplicasCachesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/replicas/caches{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemItemReplicasCachesRequestBuilder instantiates a new ItemItemReplicasCachesRequestBuilder and sets the default values.
func NewItemItemReplicasCachesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemReplicasCachesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemReplicasCachesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the status of each repository cache replica.
// returns a []ItemItemReplicasCachesable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.11/rest/repos/repos#list-repository-cache-replication-status
func (m *ItemItemReplicasCachesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemReplicasCachesRequestBuilderGetQueryParameters])([]ItemItemReplicasCachesable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
        "404": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, CreateItemItemReplicasCachesFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ItemItemReplicasCachesable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ItemItemReplicasCachesable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists the status of each repository cache replica.
// returns a *RequestInformation when successful
func (m *ItemItemReplicasCachesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemReplicasCachesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemReplicasCachesRequestBuilder when successful
func (m *ItemItemReplicasCachesRequestBuilder) WithUrl(rawUrl string)(*ItemItemReplicasCachesRequestBuilder) {
    return NewItemItemReplicasCachesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
