package manage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1ConfigInitRequestBuilder builds and executes requests for operations under \manage\v1\config\init
type V1ConfigInitRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewV1ConfigInitRequestBuilderInternal instantiates a new V1ConfigInitRequestBuilder and sets the default values.
func NewV1ConfigInitRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ConfigInitRequestBuilder) {
    m := &V1ConfigInitRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/manage/v1/config/init", pathParameters),
    }
    return m
}
// NewV1ConfigInitRequestBuilder instantiates a new V1ConfigInitRequestBuilder and sets the default values.
func NewV1ConfigInitRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ConfigInitRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1ConfigInitRequestBuilderInternal(urlParams, requestAdapter)
}
// Post when you boot a GitHub instance for the first time, you can use this endpoint to upload a license.Note that you afterwards need to `POST` to [`/manage/v1/config/apply`](https://docs.github.com/enterprise-server@3.12/rest/enterprise-admin/manage-ghes#start-configuration-apply-process) to start the actual configuration process.This endpoint also sets the root site administrator password which is used to authenticate with the GHES Manage API and the Management Console.**Note:** The request body for this operation must be submitted as `multipart/form-data` data. You can can reference the license file by prefixing the filename with the `@` symbol using `curl`. For more information, see the [`curl` documentation](https://curl.se/docs/manpage.html#-F).
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.12/rest/enterprise-admin/manage-ghes#initialize-instance-configuration-with-license-upload
func (m *V1ConfigInitRequestBuilder) Post(ctx context.Context, body i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.MultipartBody, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToPostRequestInformation when you boot a GitHub instance for the first time, you can use this endpoint to upload a license.Note that you afterwards need to `POST` to [`/manage/v1/config/apply`](https://docs.github.com/enterprise-server@3.12/rest/enterprise-admin/manage-ghes#start-configuration-apply-process) to start the actual configuration process.This endpoint also sets the root site administrator password which is used to authenticate with the GHES Manage API and the Management Console.**Note:** The request body for this operation must be submitted as `multipart/form-data` data. You can can reference the license file by prefixing the filename with the `@` symbol using `curl`. For more information, see the [`curl` documentation](https://curl.se/docs/manpage.html#-F).
// returns a *RequestInformation when successful
func (m *V1ConfigInitRequestBuilder) ToPostRequestInformation(ctx context.Context, body i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.MultipartBody, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "multipart/form-data", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1ConfigInitRequestBuilder when successful
func (m *V1ConfigInitRequestBuilder) WithUrl(rawUrl string)(*V1ConfigInitRequestBuilder) {
    return NewV1ConfigInitRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
