package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ItemActionsCacheUsageRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\actions\cache\usage
type ItemActionsCacheUsageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemActionsCacheUsageRequestBuilderInternal instantiates a new ItemActionsCacheUsageRequestBuilder and sets the default values.
func NewItemActionsCacheUsageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsCacheUsageRequestBuilder) {
    m := &ItemActionsCacheUsageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/actions/cache/usage", pathParameters),
    }
    return m
}
// NewItemActionsCacheUsageRequestBuilder instantiates a new ItemActionsCacheUsageRequestBuilder and sets the default values.
func NewItemActionsCacheUsageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsCacheUsageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsCacheUsageRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the total GitHub Actions cache usage for an enterprise.The data fetched using this API is refreshed approximately every 5 minutes, so values returned from this endpoint may take at least 5 minutes to get updated.OAuth tokens and personal access tokens (classic) need the `admin:enterprise` scope to use this endpoint.
// returns a ActionsCacheUsageOrgEnterpriseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.12/rest/actions/cache#get-github-actions-cache-usage-for-an-enterprise
func (m *ItemActionsCacheUsageRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ActionsCacheUsageOrgEnterpriseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateActionsCacheUsageOrgEnterpriseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ActionsCacheUsageOrgEnterpriseable), nil
}
// ToGetRequestInformation gets the total GitHub Actions cache usage for an enterprise.The data fetched using this API is refreshed approximately every 5 minutes, so values returned from this endpoint may take at least 5 minutes to get updated.OAuth tokens and personal access tokens (classic) need the `admin:enterprise` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemActionsCacheUsageRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemActionsCacheUsageRequestBuilder when successful
func (m *ItemActionsCacheUsageRequestBuilder) WithUrl(rawUrl string)(*ItemActionsCacheUsageRequestBuilder) {
    return NewItemActionsCacheUsageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
