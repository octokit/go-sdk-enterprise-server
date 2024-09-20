package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ItemSettingsBillingAdvancedSecurityRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\settings\billing\advanced-security
type ItemSettingsBillingAdvancedSecurityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSettingsBillingAdvancedSecurityRequestBuilderGetQueryParameters gets the GitHub Advanced Security active committers for an enterprise per repository.Each distinct user login across all repositories is counted as a single Advanced Security seat, so the `total_advanced_security_committers` is not the sum of active_users for each repository.The total number of repositories with committer information is tracked by the `total_count` field.
type ItemSettingsBillingAdvancedSecurityRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemSettingsBillingAdvancedSecurityRequestBuilderInternal instantiates a new ItemSettingsBillingAdvancedSecurityRequestBuilder and sets the default values.
func NewItemSettingsBillingAdvancedSecurityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingAdvancedSecurityRequestBuilder) {
    m := &ItemSettingsBillingAdvancedSecurityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/settings/billing/advanced-security{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemSettingsBillingAdvancedSecurityRequestBuilder instantiates a new ItemSettingsBillingAdvancedSecurityRequestBuilder and sets the default values.
func NewItemSettingsBillingAdvancedSecurityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingAdvancedSecurityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsBillingAdvancedSecurityRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the GitHub Advanced Security active committers for an enterprise per repository.Each distinct user login across all repositories is counted as a single Advanced Security seat, so the `total_advanced_security_committers` is not the sum of active_users for each repository.The total number of repositories with committer information is tracked by the `total_count` field.
// returns a AdvancedSecurityActiveCommittersable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.14/rest/enterprise-admin/billing#get-github-advanced-security-active-committers-for-an-enterprise
func (m *ItemSettingsBillingAdvancedSecurityRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemSettingsBillingAdvancedSecurityRequestBuilderGetQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.AdvancedSecurityActiveCommittersable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateAdvancedSecurityActiveCommittersFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.AdvancedSecurityActiveCommittersable), nil
}
// ToGetRequestInformation gets the GitHub Advanced Security active committers for an enterprise per repository.Each distinct user login across all repositories is counted as a single Advanced Security seat, so the `total_advanced_security_committers` is not the sum of active_users for each repository.The total number of repositories with committer information is tracked by the `total_count` field.
// returns a *RequestInformation when successful
func (m *ItemSettingsBillingAdvancedSecurityRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemSettingsBillingAdvancedSecurityRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSettingsBillingAdvancedSecurityRequestBuilder when successful
func (m *ItemSettingsBillingAdvancedSecurityRequestBuilder) WithUrl(rawUrl string)(*ItemSettingsBillingAdvancedSecurityRequestBuilder) {
    return NewItemSettingsBillingAdvancedSecurityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
