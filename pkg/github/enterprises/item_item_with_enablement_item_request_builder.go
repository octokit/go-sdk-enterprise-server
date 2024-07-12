package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ItemItemWithEnablementItemRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\{security_product}\{enablement}
type ItemItemWithEnablementItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemWithEnablementItemRequestBuilderInternal instantiates a new ItemItemWithEnablementItemRequestBuilder and sets the default values.
func NewItemItemWithEnablementItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemWithEnablementItemRequestBuilder) {
    m := &ItemItemWithEnablementItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/{security_product}/{enablement}", pathParameters),
    }
    return m
}
// NewItemItemWithEnablementItemRequestBuilder instantiates a new ItemItemWithEnablementItemRequestBuilder and sets the default values.
func NewItemItemWithEnablementItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemWithEnablementItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemWithEnablementItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Post enables or disables the specified security feature for all repositories in an enterprise.The authenticated user must be an administrator of the enterprise to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `admin:enterprise` scope to use this endpoint.
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.11/rest/enterprise-admin/code-security-and-analysis#enable-or-disable-a-security-feature
func (m *ItemItemWithEnablementItemRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation enables or disables the specified security feature for all repositories in an enterprise.The authenticated user must be an administrator of the enterprise to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `admin:enterprise` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemWithEnablementItemRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemWithEnablementItemRequestBuilder when successful
func (m *ItemItemWithEnablementItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemWithEnablementItemRequestBuilder) {
    return NewItemItemWithEnablementItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
