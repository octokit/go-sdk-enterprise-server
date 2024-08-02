package setup

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ApiStartRequestBuilder builds and executes requests for operations under \setup\api\start
type ApiStartRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewApiStartRequestBuilderInternal instantiates a new ApiStartRequestBuilder and sets the default values.
func NewApiStartRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiStartRequestBuilder) {
    m := &ApiStartRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/setup/api/start", pathParameters),
    }
    return m
}
// NewApiStartRequestBuilder instantiates a new ApiStartRequestBuilder and sets the default values.
func NewApiStartRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiStartRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApiStartRequestBuilderInternal(urlParams, requestAdapter)
}
// Post when you boot a GitHub instance for the first time, you can use the following endpoint to upload a license.Note that you need to `POST` to [`/setup/api/configure`](https://docs.github.com/enterprise-server@3.10/rest/enterprise-admin/management-console#start-a-configuration-process) to start the actual configuration process.When using this endpoint, your GitHub instance must have a password set. This can be accomplished two ways:1.  If you're working directly with the API before accessing the web interface, you must pass in the password parameter to set your password.2.  If you set up your instance via the web interface before accessing the API, your calls to this endpoint do not need the password parameter.> [!NOTE]> The request body for this operation must be submitted as `multipart/form-data` data. You can can reference the license file by prefixing the filename with the `@` symbol using `curl`. For more information, see the [`curl` documentation](https://curl.se/docs/manpage.html#-F).
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.10/rest/enterprise-admin/management-console#create-a-github-license
func (m *ApiStartRequestBuilder) Post(ctx context.Context, body i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.MultipartBody, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToPostRequestInformation when you boot a GitHub instance for the first time, you can use the following endpoint to upload a license.Note that you need to `POST` to [`/setup/api/configure`](https://docs.github.com/enterprise-server@3.10/rest/enterprise-admin/management-console#start-a-configuration-process) to start the actual configuration process.When using this endpoint, your GitHub instance must have a password set. This can be accomplished two ways:1.  If you're working directly with the API before accessing the web interface, you must pass in the password parameter to set your password.2.  If you set up your instance via the web interface before accessing the API, your calls to this endpoint do not need the password parameter.> [!NOTE]> The request body for this operation must be submitted as `multipart/form-data` data. You can can reference the license file by prefixing the filename with the `@` symbol using `curl`. For more information, see the [`curl` documentation](https://curl.se/docs/manpage.html#-F).
// returns a *RequestInformation when successful
func (m *ApiStartRequestBuilder) ToPostRequestInformation(ctx context.Context, body i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.MultipartBody, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "multipart/form-data", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ApiStartRequestBuilder when successful
func (m *ApiStartRequestBuilder) WithUrl(rawUrl string)(*ApiStartRequestBuilder) {
    return NewApiStartRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
