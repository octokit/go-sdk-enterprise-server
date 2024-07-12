package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ItemItemMergeUpstreamRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\merge-upstream
type ItemItemMergeUpstreamRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemMergeUpstreamRequestBuilderInternal instantiates a new ItemItemMergeUpstreamRequestBuilder and sets the default values.
func NewItemItemMergeUpstreamRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemMergeUpstreamRequestBuilder) {
    m := &ItemItemMergeUpstreamRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/merge-upstream", pathParameters),
    }
    return m
}
// NewItemItemMergeUpstreamRequestBuilder instantiates a new ItemItemMergeUpstreamRequestBuilder and sets the default values.
func NewItemItemMergeUpstreamRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemMergeUpstreamRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemMergeUpstreamRequestBuilderInternal(urlParams, requestAdapter)
}
// Post sync a branch of a forked repository to keep it up-to-date with the upstream repository.
// returns a MergedUpstreamable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.11/rest/branches/branches#sync-a-fork-branch-with-the-upstream-repository
func (m *ItemItemMergeUpstreamRequestBuilder) Post(ctx context.Context, body ItemItemMergeUpstreamPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.MergedUpstreamable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateMergedUpstreamFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.MergedUpstreamable), nil
}
// ToPostRequestInformation sync a branch of a forked repository to keep it up-to-date with the upstream repository.
// returns a *RequestInformation when successful
func (m *ItemItemMergeUpstreamRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemMergeUpstreamPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemMergeUpstreamRequestBuilder when successful
func (m *ItemItemMergeUpstreamRequestBuilder) WithUrl(rawUrl string)(*ItemItemMergeUpstreamRequestBuilder) {
    return NewItemItemMergeUpstreamRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
