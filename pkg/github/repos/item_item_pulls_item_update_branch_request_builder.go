package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ItemItemPullsItemUpdateBranchRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\pulls\{pull_number}\update-branch
type ItemItemPullsItemUpdateBranchRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemPullsItemUpdateBranchRequestBuilderInternal instantiates a new ItemItemPullsItemUpdateBranchRequestBuilder and sets the default values.
func NewItemItemPullsItemUpdateBranchRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsItemUpdateBranchRequestBuilder) {
    m := &ItemItemPullsItemUpdateBranchRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/pulls/{pull_number}/update-branch", pathParameters),
    }
    return m
}
// NewItemItemPullsItemUpdateBranchRequestBuilder instantiates a new ItemItemPullsItemUpdateBranchRequestBuilder and sets the default values.
func NewItemItemPullsItemUpdateBranchRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsItemUpdateBranchRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemPullsItemUpdateBranchRequestBuilderInternal(urlParams, requestAdapter)
}
// Put updates the pull request branch with the latest upstream changes by merging HEAD from the base branch into the pull request branch.Note: If making a request on behalf of a GitHub App you must also have permissions to write the contents of the head repository.
// returns a ItemItemPullsItemUpdateBranchPutResponseable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.11/rest/pulls/pulls#update-a-pull-request-branch
func (m *ItemItemPullsItemUpdateBranchRequestBuilder) Put(ctx context.Context, body ItemItemPullsItemUpdateBranchPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ItemItemPullsItemUpdateBranchPutResponseable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
        "422": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemPullsItemUpdateBranchPutResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemPullsItemUpdateBranchPutResponseable), nil
}
// ToPutRequestInformation updates the pull request branch with the latest upstream changes by merging HEAD from the base branch into the pull request branch.Note: If making a request on behalf of a GitHub App you must also have permissions to write the contents of the head repository.
// returns a *RequestInformation when successful
func (m *ItemItemPullsItemUpdateBranchRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemItemPullsItemUpdateBranchPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemPullsItemUpdateBranchRequestBuilder when successful
func (m *ItemItemPullsItemUpdateBranchRequestBuilder) WithUrl(rawUrl string)(*ItemItemPullsItemUpdateBranchRequestBuilder) {
    return NewItemItemPullsItemUpdateBranchRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
