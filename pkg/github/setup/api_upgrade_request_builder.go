package setup

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ApiUpgradeRequestBuilder builds and executes requests for operations under \setup\api\upgrade
type ApiUpgradeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewApiUpgradeRequestBuilderInternal instantiates a new ApiUpgradeRequestBuilder and sets the default values.
func NewApiUpgradeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiUpgradeRequestBuilder) {
    m := &ApiUpgradeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/setup/api/upgrade", pathParameters),
    }
    return m
}
// NewApiUpgradeRequestBuilder instantiates a new ApiUpgradeRequestBuilder and sets the default values.
func NewApiUpgradeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiUpgradeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApiUpgradeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post this API upgrades your license and also triggers the configuration process.**Note:** The request body for this operation must be submitted as `multipart/form-data` data. You can can reference the license file by prefixing the filename with the `@` symbol using `curl`. For more information, see the [`curl` documentation](https://curl.se/docs/manpage.html#-F).
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.12/rest/enterprise-admin/management-console#upgrade-a-license
func (m *ApiUpgradeRequestBuilder) Post(ctx context.Context, body i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.MultipartBody, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation this API upgrades your license and also triggers the configuration process.**Note:** The request body for this operation must be submitted as `multipart/form-data` data. You can can reference the license file by prefixing the filename with the `@` symbol using `curl`. For more information, see the [`curl` documentation](https://curl.se/docs/manpage.html#-F).
// returns a *RequestInformation when successful
func (m *ApiUpgradeRequestBuilder) ToPostRequestInformation(ctx context.Context, body i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.MultipartBody, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "multipart/form-data", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ApiUpgradeRequestBuilder when successful
func (m *ApiUpgradeRequestBuilder) WithUrl(rawUrl string)(*ApiUpgradeRequestBuilder) {
    return NewApiUpgradeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
