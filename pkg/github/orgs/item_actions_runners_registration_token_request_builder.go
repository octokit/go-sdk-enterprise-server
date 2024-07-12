package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ItemActionsRunnersRegistrationTokenRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\runners\registration-token
type ItemActionsRunnersRegistrationTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemActionsRunnersRegistrationTokenRequestBuilderInternal instantiates a new ItemActionsRunnersRegistrationTokenRequestBuilder and sets the default values.
func NewItemActionsRunnersRegistrationTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnersRegistrationTokenRequestBuilder) {
    m := &ItemActionsRunnersRegistrationTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/runners/registration-token", pathParameters),
    }
    return m
}
// NewItemActionsRunnersRegistrationTokenRequestBuilder instantiates a new ItemActionsRunnersRegistrationTokenRequestBuilder and sets the default values.
func NewItemActionsRunnersRegistrationTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnersRegistrationTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsRunnersRegistrationTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post returns a token that you can pass to the `config` script. The token expires after one hour.For example, you can replace `TOKEN` in the following example with the registration token provided by this endpoint to configure your self-hosted runner:```./config.sh --url https://github.com/octo-org --token TOKEN```Authenticated users must have admin access to the organization to use this endpoint.OAuth tokens and personal access tokens (classic) need the`admin:org` scope to use this endpoint. If the repository is private, OAuth tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a AuthenticationTokenable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.12/rest/actions/self-hosted-runners#create-a-registration-token-for-an-organization
func (m *ItemActionsRunnersRegistrationTokenRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.AuthenticationTokenable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateAuthenticationTokenFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.AuthenticationTokenable), nil
}
// ToPostRequestInformation returns a token that you can pass to the `config` script. The token expires after one hour.For example, you can replace `TOKEN` in the following example with the registration token provided by this endpoint to configure your self-hosted runner:```./config.sh --url https://github.com/octo-org --token TOKEN```Authenticated users must have admin access to the organization to use this endpoint.OAuth tokens and personal access tokens (classic) need the`admin:org` scope to use this endpoint. If the repository is private, OAuth tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemActionsRunnersRegistrationTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemActionsRunnersRegistrationTokenRequestBuilder when successful
func (m *ItemActionsRunnersRegistrationTokenRequestBuilder) WithUrl(rawUrl string)(*ItemActionsRunnersRegistrationTokenRequestBuilder) {
    return NewItemActionsRunnersRegistrationTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
