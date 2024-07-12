package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
    ia5b56a324dabaabc88906ce951d195a926d514d542852d6ef1858a0d0dfc3641 "github.com/octokit/go-sdk-enterprise-server/pkg/github/enterprises/item/auditlog"
)

// ItemAuditLogRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\audit-log
type ItemAuditLogRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuditLogRequestBuilderGetQueryParameters gets the audit log for an enterprise.The authenticated user must be an enterprise admin to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `admin:enterprise` scope to use this endpoint.
type ItemAuditLogRequestBuilderGetQueryParameters struct {
    // A cursor, as given in the [Link header](https://docs.github.com/enterprise-server@3.11/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for events after this cursor.
    After *string `uriparametername:"after"`
    // A cursor, as given in the [Link header](https://docs.github.com/enterprise-server@3.11/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for events before this cursor.
    Before *string `uriparametername:"before"`
    // The event types to include:- `web` - returns web (non-Git) events.- `git` - returns Git events.- `all` - returns both web and Git events.The default is `web`.
    Include *ia5b56a324dabaabc88906ce951d195a926d514d542852d6ef1858a0d0dfc3641.GetIncludeQueryParameterType `uriparametername:"include"`
    // The order of audit log events. To list newest events first, specify `desc`. To list oldest events first, specify `asc`.The default is `desc`.
    Order *ia5b56a324dabaabc88906ce951d195a926d514d542852d6ef1858a0d0dfc3641.GetOrderQueryParameterType `uriparametername:"order"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.11/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.11/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // A search phrase. For more information, see [Searching the audit log](https://docs.github.com/enterprise-server@3.11/admin/monitoring-activity-in-your-enterprise/reviewing-audit-logs-for-your-enterprise/searching-the-audit-log-for-your-enterprise#searching-the-audit-log).
    Phrase *string `uriparametername:"phrase"`
}
// NewItemAuditLogRequestBuilderInternal instantiates a new ItemAuditLogRequestBuilder and sets the default values.
func NewItemAuditLogRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuditLogRequestBuilder) {
    m := &ItemAuditLogRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/audit-log{?after*,before*,include*,order*,page*,per_page*,phrase*}", pathParameters),
    }
    return m
}
// NewItemAuditLogRequestBuilder instantiates a new ItemAuditLogRequestBuilder and sets the default values.
func NewItemAuditLogRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuditLogRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuditLogRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the audit log for an enterprise.The authenticated user must be an enterprise admin to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `admin:enterprise` scope to use this endpoint.
// returns a []AuditLogEventable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.11/rest/enterprise-admin/audit-log#get-the-audit-log-for-an-enterprise
func (m *ItemAuditLogRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemAuditLogRequestBuilderGetQueryParameters])([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.AuditLogEventable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateAuditLogEventFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.AuditLogEventable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.AuditLogEventable)
        }
    }
    return val, nil
}
// ToGetRequestInformation gets the audit log for an enterprise.The authenticated user must be an enterprise admin to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `admin:enterprise` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemAuditLogRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemAuditLogRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAuditLogRequestBuilder when successful
func (m *ItemAuditLogRequestBuilder) WithUrl(rawUrl string)(*ItemAuditLogRequestBuilder) {
    return NewItemAuditLogRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
