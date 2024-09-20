package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
    icd0d0eeb49e0011ccac02b4e4a0d91480a7f99532147b4eb2ededb1cd9194f0f "github.com/octokit/go-sdk-enterprise-server/pkg/github/user/starred"
)

// StarredRequestBuilder builds and executes requests for operations under \user\starred
type StarredRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// StarredRequestBuilderGetQueryParameters lists repositories the authenticated user has starred.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.star+json`**: Includes a timestamp of when the star was created.
type StarredRequestBuilderGetQueryParameters struct {
    // The direction to sort the results by.
    Direction *icd0d0eeb49e0011ccac02b4e4a0d91480a7f99532147b4eb2ededb1cd9194f0f.GetDirectionQueryParameterType `uriparametername:"direction"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The property to sort the results by. `created` means when the repository was starred. `updated` means when the repository was last pushed to.
    Sort *icd0d0eeb49e0011ccac02b4e4a0d91480a7f99532147b4eb2ededb1cd9194f0f.GetSortQueryParameterType `uriparametername:"sort"`
}
// ByOwner gets an item from the github.com/octokit/go-sdk-enterprise-server/pkg/github.user.starred.item collection
// returns a *StarredWithOwnerItemRequestBuilder when successful
func (m *StarredRequestBuilder) ByOwner(owner string)(*StarredWithOwnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if owner != "" {
        urlTplParams["owner"] = owner
    }
    return NewStarredWithOwnerItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewStarredRequestBuilderInternal instantiates a new StarredRequestBuilder and sets the default values.
func NewStarredRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StarredRequestBuilder) {
    m := &StarredRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/starred{?direction*,page*,per_page*,sort*}", pathParameters),
    }
    return m
}
// NewStarredRequestBuilder instantiates a new StarredRequestBuilder and sets the default values.
func NewStarredRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StarredRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStarredRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists repositories the authenticated user has starred.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.star+json`**: Includes a timestamp of when the star was created.
// returns a []Repositoryable when successful
// returns a BasicError error when the service returns a 401 status code
// returns a BasicError error when the service returns a 403 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.14/rest/activity/starring#list-repositories-starred-by-the-authenticated-user
func (m *StarredRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[StarredRequestBuilderGetQueryParameters])([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Repositoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
        "403": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateRepositoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Repositoryable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Repositoryable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists repositories the authenticated user has starred.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.star+json`**: Includes a timestamp of when the star was created.
// returns a *RequestInformation when successful
func (m *StarredRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[StarredRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *StarredRequestBuilder when successful
func (m *StarredRequestBuilder) WithUrl(rawUrl string)(*StarredRequestBuilder) {
    return NewStarredRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
