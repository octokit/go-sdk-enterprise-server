package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ItemItemActionsRunsItemApprovalsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\actions\runs\{run_id}\approvals
type ItemItemActionsRunsItemApprovalsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemActionsRunsItemApprovalsRequestBuilderInternal instantiates a new ItemItemActionsRunsItemApprovalsRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemApprovalsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemApprovalsRequestBuilder) {
    m := &ItemItemActionsRunsItemApprovalsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/actions/runs/{run_id}/approvals", pathParameters),
    }
    return m
}
// NewItemItemActionsRunsItemApprovalsRequestBuilder instantiates a new ItemItemActionsRunsItemApprovalsRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemApprovalsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemApprovalsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunsItemApprovalsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get anyone with read access to the repository can use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
// returns a []EnvironmentApprovalsable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.13/rest/actions/workflow-runs#get-the-review-history-for-a-workflow-run
func (m *ItemItemActionsRunsItemApprovalsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.EnvironmentApprovalsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateEnvironmentApprovalsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.EnvironmentApprovalsable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.EnvironmentApprovalsable)
        }
    }
    return val, nil
}
// ToGetRequestInformation anyone with read access to the repository can use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
// returns a *RequestInformation when successful
func (m *ItemItemActionsRunsItemApprovalsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemActionsRunsItemApprovalsRequestBuilder when successful
func (m *ItemItemActionsRunsItemApprovalsRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunsItemApprovalsRequestBuilder) {
    return NewItemItemActionsRunsItemApprovalsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
