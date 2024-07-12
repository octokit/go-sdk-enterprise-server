package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ItemActionsRunnersRemoveTokenRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\actions\runners\remove-token
type ItemActionsRunnersRemoveTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemActionsRunnersRemoveTokenRequestBuilderInternal instantiates a new ItemActionsRunnersRemoveTokenRequestBuilder and sets the default values.
func NewItemActionsRunnersRemoveTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnersRemoveTokenRequestBuilder) {
    m := &ItemActionsRunnersRemoveTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/actions/runners/remove-token", pathParameters),
    }
    return m
}
// NewItemActionsRunnersRemoveTokenRequestBuilder instantiates a new ItemActionsRunnersRemoveTokenRequestBuilder and sets the default values.
func NewItemActionsRunnersRemoveTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnersRemoveTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsRunnersRemoveTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post returns a token that you can pass to the `config` script to remove a self-hosted runner from an enterprise. The token expires after one hour.Example using remove token:To remove your self-hosted runner from an enterprise, replace `TOKEN` with the remove token provided by thisendpoint.```./config.sh remove --token TOKEN```OAuth app tokens and personal access tokens (classic) need the `manage_runners:enterprise` scope to use this endpoint.
// returns a AuthenticationTokenable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.13/rest/actions/self-hosted-runners#create-a-remove-token-for-an-enterprise
func (m *ItemActionsRunnersRemoveTokenRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.AuthenticationTokenable, error) {
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
// ToPostRequestInformation returns a token that you can pass to the `config` script to remove a self-hosted runner from an enterprise. The token expires after one hour.Example using remove token:To remove your self-hosted runner from an enterprise, replace `TOKEN` with the remove token provided by thisendpoint.```./config.sh remove --token TOKEN```OAuth app tokens and personal access tokens (classic) need the `manage_runners:enterprise` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemActionsRunnersRemoveTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemActionsRunnersRemoveTokenRequestBuilder when successful
func (m *ItemActionsRunnersRemoveTokenRequestBuilder) WithUrl(rawUrl string)(*ItemActionsRunnersRemoveTokenRequestBuilder) {
    return NewItemActionsRunnersRemoveTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
