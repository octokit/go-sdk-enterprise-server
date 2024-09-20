package teams

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
    i74905effaddfad2641db8adcacd6698497990d022175c92bee611240571486b8 "github.com/octokit/go-sdk-enterprise-server/pkg/github/teams/item/discussions"
)

// ItemDiscussionsRequestBuilder builds and executes requests for operations under \teams\{team_id}\discussions
type ItemDiscussionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDiscussionsRequestBuilderGetQueryParameters > [!WARNING]> **Deprecation notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List discussions`](https://docs.github.com/enterprise-server@3.14/rest/teams/discussions#list-discussions) endpoint.List all discussions on a team's page.OAuth app tokens and personal access tokens (classic) need the `read:discussion` scope to use this endpoint.
type ItemDiscussionsRequestBuilderGetQueryParameters struct {
    // The direction to sort the results by.
    Direction *i74905effaddfad2641db8adcacd6698497990d022175c92bee611240571486b8.GetDirectionQueryParameterType `uriparametername:"direction"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// ByDiscussion_number gets an item from the github.com/octokit/go-sdk-enterprise-server/pkg/github.teams.item.discussions.item collection
// Deprecated: 
// returns a *ItemDiscussionsWithDiscussion_numberItemRequestBuilder when successful
func (m *ItemDiscussionsRequestBuilder) ByDiscussion_number(discussion_number int32)(*ItemDiscussionsWithDiscussion_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["discussion_number"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(discussion_number), 10)
    return NewItemDiscussionsWithDiscussion_numberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemDiscussionsRequestBuilderInternal instantiates a new ItemDiscussionsRequestBuilder and sets the default values.
func NewItemDiscussionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDiscussionsRequestBuilder) {
    m := &ItemDiscussionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team_id}/discussions{?direction*,page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemDiscussionsRequestBuilder instantiates a new ItemDiscussionsRequestBuilder and sets the default values.
func NewItemDiscussionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDiscussionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDiscussionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get > [!WARNING]> **Deprecation notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List discussions`](https://docs.github.com/enterprise-server@3.14/rest/teams/discussions#list-discussions) endpoint.List all discussions on a team's page.OAuth app tokens and personal access tokens (classic) need the `read:discussion` scope to use this endpoint.
// Deprecated: 
// returns a []TeamDiscussionable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.14/rest/teams/discussions#list-discussions-legacy
func (m *ItemDiscussionsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemDiscussionsRequestBuilderGetQueryParameters])([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.TeamDiscussionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateTeamDiscussionFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.TeamDiscussionable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.TeamDiscussionable)
        }
    }
    return val, nil
}
// Post > [!WARNING]> **Deprecation notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`Create a discussion`](https://docs.github.com/enterprise-server@3.14/rest/teams/discussions#create-a-discussion) endpoint.Creates a new discussion post on a team's page.This endpoint triggers [notifications](https://docs.github.com/enterprise-server@3.14/github/managing-subscriptions-and-notifications-on-github/about-notifications). Creating content too quickly using this endpoint may result in secondary rate limiting. For more information, see "[Rate limits for the API](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/rate-limits-for-the-rest-api#about-secondary-rate-limits)" and "[Best practices for using the REST API](https://docs.github.com/enterprise-server@3.14/rest/guides/best-practices-for-using-the-rest-api)."OAuth app tokens and personal access tokens (classic) need the `write:discussion` scope to use this endpoint.
// Deprecated: 
// returns a TeamDiscussionable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.14/rest/teams/discussions#create-a-discussion-legacy
func (m *ItemDiscussionsRequestBuilder) Post(ctx context.Context, body ItemDiscussionsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.TeamDiscussionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateTeamDiscussionFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.TeamDiscussionable), nil
}
// ToGetRequestInformation > [!WARNING]> **Deprecation notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`List discussions`](https://docs.github.com/enterprise-server@3.14/rest/teams/discussions#list-discussions) endpoint.List all discussions on a team's page.OAuth app tokens and personal access tokens (classic) need the `read:discussion` scope to use this endpoint.
// Deprecated: 
// returns a *RequestInformation when successful
func (m *ItemDiscussionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemDiscussionsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation > [!WARNING]> **Deprecation notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [`Create a discussion`](https://docs.github.com/enterprise-server@3.14/rest/teams/discussions#create-a-discussion) endpoint.Creates a new discussion post on a team's page.This endpoint triggers [notifications](https://docs.github.com/enterprise-server@3.14/github/managing-subscriptions-and-notifications-on-github/about-notifications). Creating content too quickly using this endpoint may result in secondary rate limiting. For more information, see "[Rate limits for the API](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/rate-limits-for-the-rest-api#about-secondary-rate-limits)" and "[Best practices for using the REST API](https://docs.github.com/enterprise-server@3.14/rest/guides/best-practices-for-using-the-rest-api)."OAuth app tokens and personal access tokens (classic) need the `write:discussion` scope to use this endpoint.
// Deprecated: 
// returns a *RequestInformation when successful
func (m *ItemDiscussionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemDiscussionsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemDiscussionsRequestBuilder when successful
func (m *ItemDiscussionsRequestBuilder) WithUrl(rawUrl string)(*ItemDiscussionsRequestBuilder) {
    return NewItemDiscussionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
