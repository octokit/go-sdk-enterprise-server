package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
    ic653fdac08035cb8bfc7dabb4472628f1539065e4c88c3d7cb1a363ab7fa780f "github.com/octokit/go-sdk-enterprise-server/pkg/github/repos/item/item/secretscanning/alerts"
)

// ItemItemSecretScanningAlertsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\secret-scanning\alerts
type ItemItemSecretScanningAlertsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemSecretScanningAlertsRequestBuilderGetQueryParameters lists secret scanning alerts for an eligible repository, from newest to oldest.The authenticated user must be an administrator for the repository or for the organization that owns the repository to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` or `security_events` scope to use this endpoint. If this endpoint is only used with public repositories, the token can use the `public_repo` scope instead.
type ItemItemSecretScanningAlertsRequestBuilderGetQueryParameters struct {
    // A cursor, as given in the [Link header](https://docs.github.com/enterprise-server@3.14/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for events after this cursor.  To receive an initial cursor on your first request, include an empty "after" query string.
    After *string `uriparametername:"after"`
    // A cursor, as given in the [Link header](https://docs.github.com/enterprise-server@3.14/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for events before this cursor. To receive an initial cursor on your first request, include an empty "before" query string.
    Before *string `uriparametername:"before"`
    // The direction to sort the results by.
    Direction *ic653fdac08035cb8bfc7dabb4472628f1539065e4c88c3d7cb1a363ab7fa780f.GetDirectionQueryParameterType `uriparametername:"direction"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // A comma-separated list of resolutions. Only secret scanning alerts with one of these resolutions are listed. Valid resolutions are `false_positive`, `wont_fix`, `revoked`, `pattern_edited`, `pattern_deleted` or `used_in_tests`.
    Resolution *string `uriparametername:"resolution"`
    // A comma-separated list of secret types to return. By default all secret types are returned.See "[Supported secret scanning patterns](https://docs.github.com/enterprise-server@3.14/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)"for a complete list of secret types.
    Secret_type *string `uriparametername:"secret_type"`
    // The property to sort the results by. `created` means when the alert was created. `updated` means when the alert was updated or resolved.
    Sort *ic653fdac08035cb8bfc7dabb4472628f1539065e4c88c3d7cb1a363ab7fa780f.GetSortQueryParameterType `uriparametername:"sort"`
    // Set to `open` or `resolved` to only list secret scanning alerts in a specific state.
    State *ic653fdac08035cb8bfc7dabb4472628f1539065e4c88c3d7cb1a363ab7fa780f.GetStateQueryParameterType `uriparametername:"state"`
    // A comma-separated list of validities that, when present, will return alerts that match the validities in this list. Valid options are `active`, `inactive`, and `unknown`.
    Validity *string `uriparametername:"validity"`
}
// ByAlert_number gets an item from the github.com/octokit/go-sdk-enterprise-server/pkg/github.repos.item.item.secretScanning.alerts.item collection
// returns a *ItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder when successful
func (m *ItemItemSecretScanningAlertsRequestBuilder) ByAlert_number(alert_number int32)(*ItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["alert_number"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(alert_number), 10)
    return NewItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemSecretScanningAlertsRequestBuilderInternal instantiates a new ItemItemSecretScanningAlertsRequestBuilder and sets the default values.
func NewItemItemSecretScanningAlertsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecretScanningAlertsRequestBuilder) {
    m := &ItemItemSecretScanningAlertsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/secret-scanning/alerts{?after*,before*,direction*,page*,per_page*,resolution*,secret_type*,sort*,state*,validity*}", pathParameters),
    }
    return m
}
// NewItemItemSecretScanningAlertsRequestBuilder instantiates a new ItemItemSecretScanningAlertsRequestBuilder and sets the default values.
func NewItemItemSecretScanningAlertsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecretScanningAlertsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemSecretScanningAlertsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists secret scanning alerts for an eligible repository, from newest to oldest.The authenticated user must be an administrator for the repository or for the organization that owns the repository to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` or `security_events` scope to use this endpoint. If this endpoint is only used with public repositories, the token can use the `public_repo` scope instead.
// returns a []SecretScanningAlertable when successful
// returns a Alerts503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.14/rest/secret-scanning/secret-scanning#list-secret-scanning-alerts-for-a-repository
func (m *ItemItemSecretScanningAlertsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemSecretScanningAlertsRequestBuilderGetQueryParameters])([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.SecretScanningAlertable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "503": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateAlerts503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateSecretScanningAlertFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.SecretScanningAlertable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.SecretScanningAlertable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists secret scanning alerts for an eligible repository, from newest to oldest.The authenticated user must be an administrator for the repository or for the organization that owns the repository to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` or `security_events` scope to use this endpoint. If this endpoint is only used with public repositories, the token can use the `public_repo` scope instead.
// returns a *RequestInformation when successful
func (m *ItemItemSecretScanningAlertsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemSecretScanningAlertsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemSecretScanningAlertsRequestBuilder when successful
func (m *ItemItemSecretScanningAlertsRequestBuilder) WithUrl(rawUrl string)(*ItemItemSecretScanningAlertsRequestBuilder) {
    return NewItemItemSecretScanningAlertsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
