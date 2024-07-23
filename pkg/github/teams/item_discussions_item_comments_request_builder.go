package teams

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
    i7f894752788a21e319d42021818012cfae50b0fcbfd6a0ab4d54cf93b24873d7 "github.com/octokit/go-sdk-enterprise-server/pkg/github/teams/item/discussions/item/comments"
)

// ItemDiscussionsItemCommentsRequestBuilder builds and executes requests for operations under \teams\{team_id}\discussions\{discussion_number}\comments
type ItemDiscussionsItemCommentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDiscussionsItemCommentsRequestBuilderGetQueryParameters > [!WARNING]> **Deprecation notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [List discussion comments](https://docs.github.com/enterprise-server@3.13/rest/teams/discussion-comments#list-discussion-comments) endpoint.List all comments on a team discussion.OAuth app tokens and personal access tokens (classic) need the `read:discussion` scope to use this endpoint.
type ItemDiscussionsItemCommentsRequestBuilderGetQueryParameters struct {
    // The direction to sort the results by.
    Direction *i7f894752788a21e319d42021818012cfae50b0fcbfd6a0ab4d54cf93b24873d7.GetDirectionQueryParameterType `uriparametername:"direction"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.13/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.13/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// ByComment_number gets an item from the github.com/octokit/go-sdk-enterprise-server/pkg/github.teams.item.discussions.item.comments.item collection
// Deprecated: 
// returns a *ItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder when successful
func (m *ItemDiscussionsItemCommentsRequestBuilder) ByComment_number(comment_number int32)(*ItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["comment_number"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(comment_number), 10)
    return NewItemDiscussionsItemCommentsWithComment_numberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemDiscussionsItemCommentsRequestBuilderInternal instantiates a new ItemDiscussionsItemCommentsRequestBuilder and sets the default values.
func NewItemDiscussionsItemCommentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDiscussionsItemCommentsRequestBuilder) {
    m := &ItemDiscussionsItemCommentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team_id}/discussions/{discussion_number}/comments{?direction*,page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemDiscussionsItemCommentsRequestBuilder instantiates a new ItemDiscussionsItemCommentsRequestBuilder and sets the default values.
func NewItemDiscussionsItemCommentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDiscussionsItemCommentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDiscussionsItemCommentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get > [!WARNING]> **Deprecation notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [List discussion comments](https://docs.github.com/enterprise-server@3.13/rest/teams/discussion-comments#list-discussion-comments) endpoint.List all comments on a team discussion.OAuth app tokens and personal access tokens (classic) need the `read:discussion` scope to use this endpoint.
// Deprecated: 
// returns a []TeamDiscussionCommentable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.13/rest/teams/discussion-comments#list-discussion-comments-legacy
func (m *ItemDiscussionsItemCommentsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemDiscussionsItemCommentsRequestBuilderGetQueryParameters])([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.TeamDiscussionCommentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateTeamDiscussionCommentFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.TeamDiscussionCommentable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.TeamDiscussionCommentable)
        }
    }
    return val, nil
}
// Post > [!WARNING]> **Deprecation notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Create a discussion comment](https://docs.github.com/enterprise-server@3.13/rest/teams/discussion-comments#create-a-discussion-comment) endpoint.Creates a new comment on a team discussion.This endpoint triggers [notifications](https://docs.github.com/enterprise-server@3.13/github/managing-subscriptions-and-notifications-on-github/about-notifications). Creating content too quickly using this endpoint may result in secondary rate limiting. For more information, see "[Rate limits for the API](https://docs.github.com/enterprise-server@3.13/rest/using-the-rest-api/rate-limits-for-the-rest-api#about-secondary-rate-limits)" and "[Best practices for using the REST API](https://docs.github.com/enterprise-server@3.13/rest/guides/best-practices-for-using-the-rest-api)."OAuth app tokens and personal access tokens (classic) need the `write:discussion` scope to use this endpoint.
// Deprecated: 
// returns a TeamDiscussionCommentable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.13/rest/teams/discussion-comments#create-a-discussion-comment-legacy
func (m *ItemDiscussionsItemCommentsRequestBuilder) Post(ctx context.Context, body ItemDiscussionsItemCommentsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.TeamDiscussionCommentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation > [!WARNING]> **Deprecation notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [List discussion comments](https://docs.github.com/enterprise-server@3.13/rest/teams/discussion-comments#list-discussion-comments) endpoint.List all comments on a team discussion.OAuth app tokens and personal access tokens (classic) need the `read:discussion` scope to use this endpoint.
// Deprecated: 
// returns a *RequestInformation when successful
func (m *ItemDiscussionsItemCommentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemDiscussionsItemCommentsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation > [!WARNING]> **Deprecation notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Create a discussion comment](https://docs.github.com/enterprise-server@3.13/rest/teams/discussion-comments#create-a-discussion-comment) endpoint.Creates a new comment on a team discussion.This endpoint triggers [notifications](https://docs.github.com/enterprise-server@3.13/github/managing-subscriptions-and-notifications-on-github/about-notifications). Creating content too quickly using this endpoint may result in secondary rate limiting. For more information, see "[Rate limits for the API](https://docs.github.com/enterprise-server@3.13/rest/using-the-rest-api/rate-limits-for-the-rest-api#about-secondary-rate-limits)" and "[Best practices for using the REST API](https://docs.github.com/enterprise-server@3.13/rest/guides/best-practices-for-using-the-rest-api)."OAuth app tokens and personal access tokens (classic) need the `write:discussion` scope to use this endpoint.
// Deprecated: 
// returns a *RequestInformation when successful
func (m *ItemDiscussionsItemCommentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemDiscussionsItemCommentsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: 
// returns a *ItemDiscussionsItemCommentsRequestBuilder when successful
func (m *ItemDiscussionsItemCommentsRequestBuilder) WithUrl(rawUrl string)(*ItemDiscussionsItemCommentsRequestBuilder) {
    return NewItemDiscussionsItemCommentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
