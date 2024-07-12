package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
    i6c770f215866341afbb0263b70f7d303a7c6d6d999489894513c17b798403478 "github.com/octokit/go-sdk-enterprise-server/pkg/github/repos/item/item/codescanning/alerts"
)

// ItemItemCodeScanningAlertsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\code-scanning\alerts
type ItemItemCodeScanningAlertsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCodeScanningAlertsRequestBuilderGetQueryParameters lists code scanning alerts.The response includes a `most_recent_instance` object.This provides details of the most recent instance of this alertfor the default branch (or for the specified Git reference if you used `ref` in the request).OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
type ItemItemCodeScanningAlertsRequestBuilderGetQueryParameters struct {
    // The direction to sort the results by.
    Direction *i6c770f215866341afbb0263b70f7d303a7c6d6d999489894513c17b798403478.GetDirectionQueryParameterType `uriparametername:"direction"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.10/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.10/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The Git reference for the results you want to list. The `ref` for a branch can be formatted either as `refs/heads/<branch name>` or simply `<branch name>`. To reference a pull request use `refs/pull/<number>/merge`.
    Ref *string `uriparametername:"ref"`
    // If specified, only code scanning alerts with this severity will be returned.
    Severity *ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CodeScanningAlertSeverity `uriparametername:"severity"`
    // The property by which to sort the results.
    Sort *i6c770f215866341afbb0263b70f7d303a7c6d6d999489894513c17b798403478.GetSortQueryParameterType `uriparametername:"sort"`
    // If specified, only code scanning alerts with this state will be returned.
    State *ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CodeScanningAlertStateQuery `uriparametername:"state"`
    // The GUID of a code scanning tool. Only results by this tool will be listed. Note that some code scanning tools may not include a GUID in their analysis data. You can specify the tool by using either `tool_guid` or `tool_name`, but not both.
    Tool_guid *string `uriparametername:"tool_guid"`
    // The name of a code scanning tool. Only results by this tool will be listed. You can specify the tool by using either `tool_name` or `tool_guid`, but not both.
    Tool_name *string `uriparametername:"tool_name"`
}
// ByAlert_number gets an item from the github.com/octokit/go-sdk-enterprise-server/pkg/github.repos.item.item.codeScanning.alerts.item collection
// returns a *ItemItemCodeScanningAlertsWithAlert_numberItemRequestBuilder when successful
func (m *ItemItemCodeScanningAlertsRequestBuilder) ByAlert_number(alert_number int32)(*ItemItemCodeScanningAlertsWithAlert_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["alert_number"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(alert_number), 10)
    return NewItemItemCodeScanningAlertsWithAlert_numberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCodeScanningAlertsRequestBuilderInternal instantiates a new ItemItemCodeScanningAlertsRequestBuilder and sets the default values.
func NewItemItemCodeScanningAlertsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningAlertsRequestBuilder) {
    m := &ItemItemCodeScanningAlertsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/code-scanning/alerts{?direction*,page*,per_page*,ref*,severity*,sort*,state*,tool_guid*,tool_name*}", pathParameters),
    }
    return m
}
// NewItemItemCodeScanningAlertsRequestBuilder instantiates a new ItemItemCodeScanningAlertsRequestBuilder and sets the default values.
func NewItemItemCodeScanningAlertsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningAlertsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningAlertsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists code scanning alerts.The response includes a `most_recent_instance` object.This provides details of the most recent instance of this alertfor the default branch (or for the specified Git reference if you used `ref` in the request).OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
// returns a []CodeScanningAlertItemsable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a Alerts503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.10/rest/code-scanning/code-scanning#list-code-scanning-alerts-for-a-repository
func (m *ItemItemCodeScanningAlertsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemCodeScanningAlertsRequestBuilderGetQueryParameters])([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CodeScanningAlertItemsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
        "404": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
        "503": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateAlerts503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateCodeScanningAlertItemsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CodeScanningAlertItemsable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CodeScanningAlertItemsable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists code scanning alerts.The response includes a `most_recent_instance` object.This provides details of the most recent instance of this alertfor the default branch (or for the specified Git reference if you used `ref` in the request).OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
// returns a *RequestInformation when successful
func (m *ItemItemCodeScanningAlertsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemCodeScanningAlertsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemCodeScanningAlertsRequestBuilder when successful
func (m *ItemItemCodeScanningAlertsRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodeScanningAlertsRequestBuilder) {
    return NewItemItemCodeScanningAlertsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
