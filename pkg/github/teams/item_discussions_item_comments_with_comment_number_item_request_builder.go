package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder builds and executes requests for operations under \teams\{team_id}\discussions\{discussion_number}\comments\{comment_number}
type ItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemDiscussionsItemCommentsWithComment_numberItemRequestBuilderInternal instantiates a new ItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder and sets the default values.
func NewItemDiscussionsItemCommentsWithComment_numberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) {
    m := &ItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team_id}/discussions/{discussion_number}/comments/{comment_number}", pathParameters),
    }
    return m
}
// NewItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder instantiates a new ItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder and sets the default values.
func NewItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDiscussionsItemCommentsWithComment_numberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Delete a discussion comment](https://docs.github.com/enterprise-server@3.12/rest/teams/discussion-comments#delete-a-discussion-comment) endpoint.Deletes a comment on a team discussion.OAuth app tokens and personal access tokens (classic) need the `write:discussion` scope to use this endpoint.
// Deprecated: 
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.12/rest/teams/discussion-comments#delete-a-discussion-comment-legacy
func (m *ItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Get a discussion comment](https://docs.github.com/enterprise-server@3.12/rest/teams/discussion-comments#get-a-discussion-comment) endpoint.Get a specific comment on a team discussion.OAuth app tokens and personal access tokens (classic) need the `read:discussion` scope to use this endpoint.
// Deprecated: 
// returns a TeamDiscussionCommentable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.12/rest/teams/discussion-comments#get-a-discussion-comment-legacy
func (m *ItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.TeamDiscussionCommentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateTeamDiscussionCommentFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.TeamDiscussionCommentable), nil
}
// Patch **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Update a discussion comment](https://docs.github.com/enterprise-server@3.12/rest/teams/discussion-comments#update-a-discussion-comment) endpoint.Edits the body text of a discussion comment.OAuth app tokens and personal access tokens (classic) need the `write:discussion` scope to use this endpoint.
// Deprecated: 
// returns a TeamDiscussionCommentable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.12/rest/teams/discussion-comments#update-a-discussion-comment-legacy
func (m *ItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) Patch(ctx context.Context, body ItemDiscussionsItemCommentsItemWithComment_numberPatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.TeamDiscussionCommentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateTeamDiscussionCommentFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.TeamDiscussionCommentable), nil
}
// Reactions the reactions property
// returns a *ItemDiscussionsItemCommentsItemReactionsRequestBuilder when successful
func (m *ItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) Reactions()(*ItemDiscussionsItemCommentsItemReactionsRequestBuilder) {
    return NewItemDiscussionsItemCommentsItemReactionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Delete a discussion comment](https://docs.github.com/enterprise-server@3.12/rest/teams/discussion-comments#delete-a-discussion-comment) endpoint.Deletes a comment on a team discussion.OAuth app tokens and personal access tokens (classic) need the `write:discussion` scope to use this endpoint.
// Deprecated: 
// returns a *RequestInformation when successful
func (m *ItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// ToGetRequestInformation **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Get a discussion comment](https://docs.github.com/enterprise-server@3.12/rest/teams/discussion-comments#get-a-discussion-comment) endpoint.Get a specific comment on a team discussion.OAuth app tokens and personal access tokens (classic) need the `read:discussion` scope to use this endpoint.
// Deprecated: 
// returns a *RequestInformation when successful
func (m *ItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Update a discussion comment](https://docs.github.com/enterprise-server@3.12/rest/teams/discussion-comments#update-a-discussion-comment) endpoint.Edits the body text of a discussion comment.OAuth app tokens and personal access tokens (classic) need the `write:discussion` scope to use this endpoint.
// Deprecated: 
// returns a *RequestInformation when successful
func (m *ItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemDiscussionsItemCommentsItemWithComment_numberPatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: 
// returns a *ItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder when successful
func (m *ItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) WithUrl(rawUrl string)(*ItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) {
    return NewItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
