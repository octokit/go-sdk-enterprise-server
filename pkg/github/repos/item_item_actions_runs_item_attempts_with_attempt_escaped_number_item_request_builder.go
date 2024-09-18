package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\actions\runs\{run_id}\attempts\{attempt_number}
type ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderGetQueryParameters gets a specific workflow run attempt.Anyone with read access to the repository can use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
type ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderGetQueryParameters struct {
    // If `true` pull requests are omitted from the response (empty array).
    Exclude_pull_requests *bool `uriparametername:"exclude_pull_requests"`
}
// NewItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderInternal instantiates a new ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) {
    m := &ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/actions/runs/{run_id}/attempts/{attempt_number}{?exclude_pull_requests*}", pathParameters),
    }
    return m
}
// NewItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder instantiates a new ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a specific workflow run attempt.Anyone with read access to the repository can use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
// returns a WorkflowRunable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.14/rest/actions/workflow-runs#get-a-workflow-run-attempt
func (m *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderGetQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.WorkflowRunable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateWorkflowRunFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.WorkflowRunable), nil
}
// Jobs the jobs property
// returns a *ItemItemActionsRunsItemAttemptsItemJobsRequestBuilder when successful
func (m *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) Jobs()(*ItemItemActionsRunsItemAttemptsItemJobsRequestBuilder) {
    return NewItemItemActionsRunsItemAttemptsItemJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Logs the logs property
// returns a *ItemItemActionsRunsItemAttemptsItemLogsRequestBuilder when successful
func (m *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) Logs()(*ItemItemActionsRunsItemAttemptsItemLogsRequestBuilder) {
    return NewItemItemActionsRunsItemAttemptsItemLogsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation gets a specific workflow run attempt.Anyone with read access to the repository can use this endpoint.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
// returns a *RequestInformation when successful
func (m *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder when successful
func (m *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) {
    return NewItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
