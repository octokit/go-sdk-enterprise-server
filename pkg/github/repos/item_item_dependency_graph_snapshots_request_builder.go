package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ItemItemDependencyGraphSnapshotsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\dependency-graph\snapshots
type ItemItemDependencyGraphSnapshotsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemDependencyGraphSnapshotsRequestBuilderInternal instantiates a new ItemItemDependencyGraphSnapshotsRequestBuilder and sets the default values.
func NewItemItemDependencyGraphSnapshotsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependencyGraphSnapshotsRequestBuilder) {
    m := &ItemItemDependencyGraphSnapshotsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/dependency-graph/snapshots", pathParameters),
    }
    return m
}
// NewItemItemDependencyGraphSnapshotsRequestBuilder instantiates a new ItemItemDependencyGraphSnapshotsRequestBuilder and sets the default values.
func NewItemItemDependencyGraphSnapshotsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependencyGraphSnapshotsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemDependencyGraphSnapshotsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a new snapshot of a repository's dependencies.The authenticated user must have access to the repository.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a ItemItemDependencyGraphSnapshotsPostResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.13/rest/dependency-graph/dependency-submission#create-a-snapshot-of-dependencies-for-a-repository
func (m *ItemItemDependencyGraphSnapshotsRequestBuilder) Post(ctx context.Context, body ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Snapshotable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ItemItemDependencyGraphSnapshotsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemDependencyGraphSnapshotsPostResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemDependencyGraphSnapshotsPostResponseable), nil
}
// ToPostRequestInformation create a new snapshot of a repository's dependencies.The authenticated user must have access to the repository.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemDependencyGraphSnapshotsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Snapshotable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemDependencyGraphSnapshotsRequestBuilder when successful
func (m *ItemItemDependencyGraphSnapshotsRequestBuilder) WithUrl(rawUrl string)(*ItemItemDependencyGraphSnapshotsRequestBuilder) {
    return NewItemItemDependencyGraphSnapshotsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
