package manage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// V1MaintenanceRequestBuilder builds and executes requests for operations under \manage\v1\maintenance
type V1MaintenanceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1MaintenanceRequestBuilderGetQueryParameters gets the status and details of maintenance mode on all available nodes. For more information, see "[Enabling and scheduling maintenance mode](https://docs.github.com/enterprise-server@3.14/admin/configuration/configuring-your-enterprise/enabling-and-scheduling-maintenance-mode)."
type V1MaintenanceRequestBuilderGetQueryParameters struct {
    // The cluster roles from the cluster configuration file.
    Cluster_roles *string `uriparametername:"cluster_roles"`
    // The UUID which identifies a node.
    Uuid *string `uriparametername:"uuid"`
}
// NewV1MaintenanceRequestBuilderInternal instantiates a new V1MaintenanceRequestBuilder and sets the default values.
func NewV1MaintenanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1MaintenanceRequestBuilder) {
    m := &V1MaintenanceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/manage/v1/maintenance{?cluster_roles*,uuid*}", pathParameters),
    }
    return m
}
// NewV1MaintenanceRequestBuilder instantiates a new V1MaintenanceRequestBuilder and sets the default values.
func NewV1MaintenanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1MaintenanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1MaintenanceRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the status and details of maintenance mode on all available nodes. For more information, see "[Enabling and scheduling maintenance mode](https://docs.github.com/enterprise-server@3.14/admin/configuration/configuring-your-enterprise/enabling-and-scheduling-maintenance-mode)."
// returns a []GhesGetMaintenanceable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.14/rest/enterprise-admin/manage-ghes#get-the-status-of-maintenance-mode
func (m *V1MaintenanceRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V1MaintenanceRequestBuilderGetQueryParameters])([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.GhesGetMaintenanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateGhesGetMaintenanceFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.GhesGetMaintenanceable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.GhesGetMaintenanceable)
        }
    }
    return val, nil
}
// Post sets or schedules the maintenance mode. For more information, see "[Enabling and scheduling maintenance mode](https://docs.github.com/enterprise-server@3.14/admin/configuration/configuring-your-enterprise/enabling-and-scheduling-maintenance-mode)."
// returns a []GhesSetMaintenanceResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.14/rest/enterprise-admin/manage-ghes#set-the-status-of-maintenance-mode
func (m *V1MaintenanceRequestBuilder) Post(ctx context.Context, body ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.GhesSetMaintenanceRequestable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.GhesSetMaintenanceResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateGhesSetMaintenanceResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.GhesSetMaintenanceResponseable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.GhesSetMaintenanceResponseable)
        }
    }
    return val, nil
}
// ToGetRequestInformation gets the status and details of maintenance mode on all available nodes. For more information, see "[Enabling and scheduling maintenance mode](https://docs.github.com/enterprise-server@3.14/admin/configuration/configuring-your-enterprise/enabling-and-scheduling-maintenance-mode)."
// returns a *RequestInformation when successful
func (m *V1MaintenanceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V1MaintenanceRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation sets or schedules the maintenance mode. For more information, see "[Enabling and scheduling maintenance mode](https://docs.github.com/enterprise-server@3.14/admin/configuration/configuring-your-enterprise/enabling-and-scheduling-maintenance-mode)."
// returns a *RequestInformation when successful
func (m *V1MaintenanceRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.GhesSetMaintenanceRequestable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1MaintenanceRequestBuilder when successful
func (m *V1MaintenanceRequestBuilder) WithUrl(rawUrl string)(*V1MaintenanceRequestBuilder) {
    return NewV1MaintenanceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
