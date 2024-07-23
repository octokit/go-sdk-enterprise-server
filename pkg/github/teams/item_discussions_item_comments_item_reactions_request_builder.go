package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
    i0e502917b2f0701ab8a9a6e2189eca63fadab16c6e7ceb276965613570b6ac9a "github.com/octokit/go-sdk-enterprise-server/pkg/github/teams/item/discussions/item/comments/item/reactions"
)

// ItemDiscussionsItemCommentsItemReactionsRequestBuilder builds and executes requests for operations under \teams\{team_id}\discussions\{discussion_number}\comments\{comment_number}\reactions
type ItemDiscussionsItemCommentsItemReactionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDiscussionsItemCommentsItemReactionsRequestBuilderGetQueryParameters > [!WARNING]> **Deprecation notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List reactions for a team discussion comment`](https://docs.github.com/enterprise-server@3.12/rest/reactions/reactions#list-reactions-for-a-team-discussion-comment) endpoint.List the reactions to a [team discussion comment](https://docs.github.com/enterprise-server@3.12/rest/teams/discussion-comments#get-a-discussion-comment).OAuth app tokens and personal access tokens (classic) need the `read:discussion` scope to use this endpoint.
type ItemDiscussionsItemCommentsItemReactionsRequestBuilderGetQueryParameters struct {
    // Returns a single [reaction type](https://docs.github.com/enterprise-server@3.12/rest/reactions/reactions#about-reactions). Omit this parameter to list all reactions to a team discussion comment.
    Content *i0e502917b2f0701ab8a9a6e2189eca63fadab16c6e7ceb276965613570b6ac9a.GetContentQueryParameterType `uriparametername:"content"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.12/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.12/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemDiscussionsItemCommentsItemReactionsRequestBuilderInternal instantiates a new ItemDiscussionsItemCommentsItemReactionsRequestBuilder and sets the default values.
func NewItemDiscussionsItemCommentsItemReactionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDiscussionsItemCommentsItemReactionsRequestBuilder) {
    m := &ItemDiscussionsItemCommentsItemReactionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team_id}/discussions/{discussion_number}/comments/{comment_number}/reactions{?content*,page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemDiscussionsItemCommentsItemReactionsRequestBuilder instantiates a new ItemDiscussionsItemCommentsItemReactionsRequestBuilder and sets the default values.
func NewItemDiscussionsItemCommentsItemReactionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDiscussionsItemCommentsItemReactionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDiscussionsItemCommentsItemReactionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get > [!WARNING]> **Deprecation notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List reactions for a team discussion comment`](https://docs.github.com/enterprise-server@3.12/rest/reactions/reactions#list-reactions-for-a-team-discussion-comment) endpoint.List the reactions to a [team discussion comment](https://docs.github.com/enterprise-server@3.12/rest/teams/discussion-comments#get-a-discussion-comment).OAuth app tokens and personal access tokens (classic) need the `read:discussion` scope to use this endpoint.
// Deprecated: 
// returns a []Reactionable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.12/rest/reactions/reactions#list-reactions-for-a-team-discussion-comment-legacy
func (m *ItemDiscussionsItemCommentsItemReactionsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemDiscussionsItemCommentsItemReactionsRequestBuilderGetQueryParameters])([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Reactionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateReactionFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Reactionable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Reactionable)
        }
    }
    return val, nil
}
// Post > [!WARNING]> **Deprecation notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new "[Create reaction for a team discussion comment](https://docs.github.com/enterprise-server@3.12/rest/reactions/reactions#create-reaction-for-a-team-discussion-comment)" endpoint.Create a reaction to a [team discussion comment](https://docs.github.com/enterprise-server@3.12/rest/teams/discussion-comments#get-a-discussion-comment).A response with an HTTP `200` status means that you already added the reaction type to this team discussion comment.OAuth app tokens and personal access tokens (classic) need the `write:discussion` scope to use this endpoint.
// Deprecated: 
// returns a Reactionable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.12/rest/reactions/reactions#create-reaction-for-a-team-discussion-comment-legacy
func (m *ItemDiscussionsItemCommentsItemReactionsRequestBuilder) Post(ctx context.Context, body ItemDiscussionsItemCommentsItemReactionsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Reactionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateReactionFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Reactionable), nil
}
// ToGetRequestInformation > [!WARNING]> **Deprecation notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List reactions for a team discussion comment`](https://docs.github.com/enterprise-server@3.12/rest/reactions/reactions#list-reactions-for-a-team-discussion-comment) endpoint.List the reactions to a [team discussion comment](https://docs.github.com/enterprise-server@3.12/rest/teams/discussion-comments#get-a-discussion-comment).OAuth app tokens and personal access tokens (classic) need the `read:discussion` scope to use this endpoint.
// Deprecated: 
// returns a *RequestInformation when successful
func (m *ItemDiscussionsItemCommentsItemReactionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemDiscussionsItemCommentsItemReactionsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation > [!WARNING]> **Deprecation notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new "[Create reaction for a team discussion comment](https://docs.github.com/enterprise-server@3.12/rest/reactions/reactions#create-reaction-for-a-team-discussion-comment)" endpoint.Create a reaction to a [team discussion comment](https://docs.github.com/enterprise-server@3.12/rest/teams/discussion-comments#get-a-discussion-comment).A response with an HTTP `200` status means that you already added the reaction type to this team discussion comment.OAuth app tokens and personal access tokens (classic) need the `write:discussion` scope to use this endpoint.
// Deprecated: 
// returns a *RequestInformation when successful
func (m *ItemDiscussionsItemCommentsItemReactionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemDiscussionsItemCommentsItemReactionsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemDiscussionsItemCommentsItemReactionsRequestBuilder when successful
func (m *ItemDiscussionsItemCommentsItemReactionsRequestBuilder) WithUrl(rawUrl string)(*ItemDiscussionsItemCommentsItemReactionsRequestBuilder) {
    return NewItemDiscussionsItemCommentsItemReactionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
