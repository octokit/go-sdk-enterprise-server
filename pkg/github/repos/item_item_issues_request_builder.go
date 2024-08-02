package repos

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
    i7e4c62d8268966954a5c0ac9afd824b3b6979d9f8715928af4b65a3e87a07c54 "github.com/octokit/go-sdk-enterprise-server/pkg/github/repos/item/item/issues"
)

// ItemItemIssuesRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\issues
type ItemItemIssuesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemIssuesRequestBuilderGetQueryParameters list issues in a repository. Only open issues will be listed.> [!NOTE]> GitHub's REST API considers every pull request an issue, but not every issue is a pull request. For this reason, "Issues" endpoints may return both issues and pull requests in the response. You can identify pull requests by the `pull_request` key. Be aware that the `id` of a pull request returned from "Issues" endpoints will be an _issue id_. To find out the pull request id, use the "[List pull requests](https://docs.github.com/enterprise-server@3.10/rest/pulls/pulls#list-pull-requests)" endpoint.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/enterprise-server@3.10/rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.raw+json`**: Returns the raw markdown body. Response will include `body`. This is the default if you do not pass any specific media type.- **`application/vnd.github.text+json`**: Returns a text only representation of the markdown body. Response will include `body_text`.- **`application/vnd.github.html+json`**: Returns HTML rendered from the body's markdown. Response will include `body_html`.- **`application/vnd.github.full+json`**: Returns raw, text, and HTML representations. Response will include `body`, `body_text`, and `body_html`.
type ItemItemIssuesRequestBuilderGetQueryParameters struct {
    // Can be the name of a user. Pass in `none` for issues with no assigned user, and `*` for issues assigned to any user.
    Assignee *string `uriparametername:"assignee"`
    // The user that created the issue.
    Creator *string `uriparametername:"creator"`
    // The direction to sort the results by.
    Direction *i7e4c62d8268966954a5c0ac9afd824b3b6979d9f8715928af4b65a3e87a07c54.GetDirectionQueryParameterType `uriparametername:"direction"`
    // A list of comma separated label names. Example: `bug,ui,@high`
    Labels *string `uriparametername:"labels"`
    // A user that's mentioned in the issue.
    Mentioned *string `uriparametername:"mentioned"`
    // If an `integer` is passed, it should refer to a milestone by its `number` field. If the string `*` is passed, issues with any milestone are accepted. If the string `none` is passed, issues without milestones are returned.
    Milestone *string `uriparametername:"milestone"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.10/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.10/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // Only show results that were last updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
    Since *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"since"`
    // What to sort results by.
    Sort *i7e4c62d8268966954a5c0ac9afd824b3b6979d9f8715928af4b65a3e87a07c54.GetSortQueryParameterType `uriparametername:"sort"`
    // Indicates the state of the issues to return.
    State *i7e4c62d8268966954a5c0ac9afd824b3b6979d9f8715928af4b65a3e87a07c54.GetStateQueryParameterType `uriparametername:"state"`
}
// ByIssue_number gets an item from the github.com/octokit/go-sdk-enterprise-server/pkg/github.repos.item.item.issues.item collection
// returns a *ItemItemIssuesWithIssue_numberItemRequestBuilder when successful
func (m *ItemItemIssuesRequestBuilder) ByIssue_number(issue_number int32)(*ItemItemIssuesWithIssue_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["issue_number"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(issue_number), 10)
    return NewItemItemIssuesWithIssue_numberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Comments the comments property
// returns a *ItemItemIssuesCommentsRequestBuilder when successful
func (m *ItemItemIssuesRequestBuilder) Comments()(*ItemItemIssuesCommentsRequestBuilder) {
    return NewItemItemIssuesCommentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemIssuesRequestBuilderInternal instantiates a new ItemItemIssuesRequestBuilder and sets the default values.
func NewItemItemIssuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesRequestBuilder) {
    m := &ItemItemIssuesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/issues{?assignee*,creator*,direction*,labels*,mentioned*,milestone*,page*,per_page*,since*,sort*,state*}", pathParameters),
    }
    return m
}
// NewItemItemIssuesRequestBuilder instantiates a new ItemItemIssuesRequestBuilder and sets the default values.
func NewItemItemIssuesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemIssuesRequestBuilderInternal(urlParams, requestAdapter)
}
// Events the events property
// returns a *ItemItemIssuesEventsRequestBuilder when successful
func (m *ItemItemIssuesRequestBuilder) Events()(*ItemItemIssuesEventsRequestBuilder) {
    return NewItemItemIssuesEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list issues in a repository. Only open issues will be listed.> [!NOTE]> GitHub's REST API considers every pull request an issue, but not every issue is a pull request. For this reason, "Issues" endpoints may return both issues and pull requests in the response. You can identify pull requests by the `pull_request` key. Be aware that the `id` of a pull request returned from "Issues" endpoints will be an _issue id_. To find out the pull request id, use the "[List pull requests](https://docs.github.com/enterprise-server@3.10/rest/pulls/pulls#list-pull-requests)" endpoint.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/enterprise-server@3.10/rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.raw+json`**: Returns the raw markdown body. Response will include `body`. This is the default if you do not pass any specific media type.- **`application/vnd.github.text+json`**: Returns a text only representation of the markdown body. Response will include `body_text`.- **`application/vnd.github.html+json`**: Returns HTML rendered from the body's markdown. Response will include `body_html`.- **`application/vnd.github.full+json`**: Returns raw, text, and HTML representations. Response will include `body`, `body_text`, and `body_html`.
// returns a []Issueable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.10/rest/issues/issues#list-repository-issues
func (m *ItemItemIssuesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemIssuesRequestBuilderGetQueryParameters])([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Issueable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
        "422": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateIssueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Issueable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Issueable)
        }
    }
    return val, nil
}
// Post any user with pull access to a repository can create an issue. If [issues are disabled in the repository](https://docs.github.com/enterprise-server@3.10/articles/disabling-issues/), the API returns a `410 Gone` status.This endpoint triggers [notifications](https://docs.github.com/enterprise-server@3.10/github/managing-subscriptions-and-notifications-on-github/about-notifications). Creating content too quickly using this endpoint may result in secondary rate limiting. For more information, see "[Rate limits for the API](https://docs.github.com/enterprise-server@3.10/rest/using-the-rest-api/rate-limits-for-the-rest-api#about-secondary-rate-limits)"and "[Best practices for using the REST API](https://docs.github.com/enterprise-server@3.10/rest/guides/best-practices-for-using-the-rest-api)."This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/enterprise-server@3.10/rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.raw+json`**: Returns the raw markdown body. Response will include `body`. This is the default if you do not pass any specific media type.- **`application/vnd.github.text+json`**: Returns a text only representation of the markdown body. Response will include `body_text`.- **`application/vnd.github.html+json`**: Returns HTML rendered from the body's markdown. Response will include `body_html`.- **`application/vnd.github.full+json`**: Returns raw, text, and HTML representations. Response will include `body`, `body_text`, and `body_html`.
// returns a Issueable when successful
// returns a BasicError error when the service returns a 400 status code
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 410 status code
// returns a ValidationError error when the service returns a 422 status code
// returns a Issue503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.10/rest/issues/issues#create-an-issue
func (m *ItemItemIssuesRequestBuilder) Post(ctx context.Context, body ItemItemIssuesPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Issueable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
        "403": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
        "404": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
        "410": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
        "422": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateValidationErrorFromDiscriminatorValue,
        "503": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateIssue503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateIssueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Issueable), nil
}
// ToGetRequestInformation list issues in a repository. Only open issues will be listed.> [!NOTE]> GitHub's REST API considers every pull request an issue, but not every issue is a pull request. For this reason, "Issues" endpoints may return both issues and pull requests in the response. You can identify pull requests by the `pull_request` key. Be aware that the `id` of a pull request returned from "Issues" endpoints will be an _issue id_. To find out the pull request id, use the "[List pull requests](https://docs.github.com/enterprise-server@3.10/rest/pulls/pulls#list-pull-requests)" endpoint.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/enterprise-server@3.10/rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.raw+json`**: Returns the raw markdown body. Response will include `body`. This is the default if you do not pass any specific media type.- **`application/vnd.github.text+json`**: Returns a text only representation of the markdown body. Response will include `body_text`.- **`application/vnd.github.html+json`**: Returns HTML rendered from the body's markdown. Response will include `body_html`.- **`application/vnd.github.full+json`**: Returns raw, text, and HTML representations. Response will include `body`, `body_text`, and `body_html`.
// returns a *RequestInformation when successful
func (m *ItemItemIssuesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemIssuesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation any user with pull access to a repository can create an issue. If [issues are disabled in the repository](https://docs.github.com/enterprise-server@3.10/articles/disabling-issues/), the API returns a `410 Gone` status.This endpoint triggers [notifications](https://docs.github.com/enterprise-server@3.10/github/managing-subscriptions-and-notifications-on-github/about-notifications). Creating content too quickly using this endpoint may result in secondary rate limiting. For more information, see "[Rate limits for the API](https://docs.github.com/enterprise-server@3.10/rest/using-the-rest-api/rate-limits-for-the-rest-api#about-secondary-rate-limits)"and "[Best practices for using the REST API](https://docs.github.com/enterprise-server@3.10/rest/guides/best-practices-for-using-the-rest-api)."This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/enterprise-server@3.10/rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.raw+json`**: Returns the raw markdown body. Response will include `body`. This is the default if you do not pass any specific media type.- **`application/vnd.github.text+json`**: Returns a text only representation of the markdown body. Response will include `body_text`.- **`application/vnd.github.html+json`**: Returns HTML rendered from the body's markdown. Response will include `body_html`.- **`application/vnd.github.full+json`**: Returns raw, text, and HTML representations. Response will include `body`, `body_text`, and `body_html`.
// returns a *RequestInformation when successful
func (m *ItemItemIssuesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemIssuesPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemIssuesRequestBuilder when successful
func (m *ItemItemIssuesRequestBuilder) WithUrl(rawUrl string)(*ItemItemIssuesRequestBuilder) {
    return NewItemItemIssuesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
