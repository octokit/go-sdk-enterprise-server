package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
    i2a9b242c816cd3407964bab8573864203d494f3b0754df96c1843fd975098c4a "github.com/octokit/go-sdk-enterprise-server/pkg/github/orgs/item/dependabot/alerts"
)

// ItemDependabotAlertsRequestBuilder builds and executes requests for operations under \orgs\{org}\dependabot\alerts
type ItemDependabotAlertsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDependabotAlertsRequestBuilderGetQueryParameters lists Dependabot alerts for an organization.The authenticated user must be an owner or security manager for the organization to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint. If this endpoint is only used with public repositories, the token can use the `public_repo` scope instead.
type ItemDependabotAlertsRequestBuilderGetQueryParameters struct {
    // A cursor, as given in the [Link header](https://docs.github.com/enterprise-server@3.14/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results after this cursor. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    After *string `uriparametername:"after"`
    // A cursor, as given in the [Link header](https://docs.github.com/enterprise-server@3.14/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results before this cursor. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Before *string `uriparametername:"before"`
    // The direction to sort the results by.
    Direction *i2a9b242c816cd3407964bab8573864203d494f3b0754df96c1843fd975098c4a.GetDirectionQueryParameterType `uriparametername:"direction"`
    // A comma-separated list of ecosystems. If specified, only alerts for these ecosystems will be returned.Can be: `composer`, `go`, `maven`, `npm`, `nuget`, `pip`, `pub`, `rubygems`, `rust`
    Ecosystem *string `uriparametername:"ecosystem"`
    // **Deprecated**. The number of results per page (max 100), starting from the first matching result.This parameter must not be used in combination with `last`.Instead, use `per_page` in combination with `after` to fetch the first page of results.
    First *int32 `uriparametername:"first"`
    // **Deprecated**. The number of results per page (max 100), starting from the last matching result.This parameter must not be used in combination with `first`.Instead, use `per_page` in combination with `before` to fetch the last page of results.
    Last *int32 `uriparametername:"last"`
    // A comma-separated list of package names. If specified, only alerts for these packages will be returned.
    Package *string `uriparametername:"package"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The scope of the vulnerable dependency. If specified, only alerts with this scope will be returned.
    Scope *i2a9b242c816cd3407964bab8573864203d494f3b0754df96c1843fd975098c4a.GetScopeQueryParameterType `uriparametername:"scope"`
    // A comma-separated list of severities. If specified, only alerts with these severities will be returned.Can be: `low`, `medium`, `high`, `critical`
    Severity *string `uriparametername:"severity"`
    // The property by which to sort the results.`created` means when the alert was created.`updated` means when the alert's state last changed.
    Sort *i2a9b242c816cd3407964bab8573864203d494f3b0754df96c1843fd975098c4a.GetSortQueryParameterType `uriparametername:"sort"`
    // A comma-separated list of states. If specified, only alerts with these states will be returned.Can be: `auto_dismissed`, `dismissed`, `fixed`, `open`
    State *string `uriparametername:"state"`
}
// NewItemDependabotAlertsRequestBuilderInternal instantiates a new ItemDependabotAlertsRequestBuilder and sets the default values.
func NewItemDependabotAlertsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDependabotAlertsRequestBuilder) {
    m := &ItemDependabotAlertsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/dependabot/alerts{?after*,before*,direction*,ecosystem*,first*,last*,package*,per_page*,scope*,severity*,sort*,state*}", pathParameters),
    }
    return m
}
// NewItemDependabotAlertsRequestBuilder instantiates a new ItemDependabotAlertsRequestBuilder and sets the default values.
func NewItemDependabotAlertsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDependabotAlertsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDependabotAlertsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists Dependabot alerts for an organization.The authenticated user must be an owner or security manager for the organization to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint. If this endpoint is only used with public repositories, the token can use the `public_repo` scope instead.
// returns a []DependabotAlertWithRepositoryable when successful
// returns a BasicError error when the service returns a 400 status code
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationErrorSimple error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.14/rest/dependabot/alerts#list-dependabot-alerts-for-an-organization
func (m *ItemDependabotAlertsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemDependabotAlertsRequestBuilderGetQueryParameters])([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.DependabotAlertWithRepositoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
        "403": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
        "404": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
        "422": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateDependabotAlertWithRepositoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.DependabotAlertWithRepositoryable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.DependabotAlertWithRepositoryable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists Dependabot alerts for an organization.The authenticated user must be an owner or security manager for the organization to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint. If this endpoint is only used with public repositories, the token can use the `public_repo` scope instead.
// returns a *RequestInformation when successful
func (m *ItemDependabotAlertsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemDependabotAlertsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemDependabotAlertsRequestBuilder when successful
func (m *ItemDependabotAlertsRequestBuilder) WithUrl(rawUrl string)(*ItemDependabotAlertsRequestBuilder) {
    return NewItemDependabotAlertsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
