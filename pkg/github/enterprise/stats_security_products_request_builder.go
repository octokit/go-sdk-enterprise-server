package enterprise

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// StatsSecurityProductsRequestBuilder builds and executes requests for operations under \enterprise\stats\security-products
type StatsSecurityProductsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewStatsSecurityProductsRequestBuilderInternal instantiates a new StatsSecurityProductsRequestBuilder and sets the default values.
func NewStatsSecurityProductsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StatsSecurityProductsRequestBuilder) {
    m := &StatsSecurityProductsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprise/stats/security-products", pathParameters),
    }
    return m
}
// NewStatsSecurityProductsRequestBuilder instantiates a new StatsSecurityProductsRequestBuilder and sets the default values.
func NewStatsSecurityProductsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StatsSecurityProductsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStatsSecurityProductsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the statistics about security products for a GitHub Enterprise Server instance.To use this endpoint, you must be a site admin.
// returns a EnterpriseSecurityProductsOverviewable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.11/rest/enterprise-admin/admin-stats#get-security-products-statistics
func (m *StatsSecurityProductsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.EnterpriseSecurityProductsOverviewable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateEnterpriseSecurityProductsOverviewFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.EnterpriseSecurityProductsOverviewable), nil
}
// ToGetRequestInformation gets the statistics about security products for a GitHub Enterprise Server instance.To use this endpoint, you must be a site admin.
// returns a *RequestInformation when successful
func (m *StatsSecurityProductsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *StatsSecurityProductsRequestBuilder when successful
func (m *StatsSecurityProductsRequestBuilder) WithUrl(rawUrl string)(*StatsSecurityProductsRequestBuilder) {
    return NewStatsSecurityProductsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
