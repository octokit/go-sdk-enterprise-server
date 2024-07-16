package manage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// V1ConfigSettingsRequestBuilder builds and executes requests for operations under \manage\v1\config\settings
type V1ConfigSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewV1ConfigSettingsRequestBuilderInternal instantiates a new V1ConfigSettingsRequestBuilder and sets the default values.
func NewV1ConfigSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ConfigSettingsRequestBuilder) {
    m := &V1ConfigSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/manage/v1/config/settings", pathParameters),
    }
    return m
}
// NewV1ConfigSettingsRequestBuilder instantiates a new V1ConfigSettingsRequestBuilder and sets the default values.
func NewV1ConfigSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ConfigSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1ConfigSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a list of settings for a GitHub Enterprise Server instance.
// returns a GhesGetSettingsable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.12/rest/enterprise-admin/manage-ghes#get-settings
func (m *V1ConfigSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.GhesGetSettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateGhesGetSettingsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.GhesGetSettingsable), nil
}
// Put updates the settings on your instance. For a list of the available settings, see the [Get settings endpoint](https://docs.github.com/enterprise-server@3.12/rest/enterprise-admin/manage-ghes#get-settings).**Notes:**- The request body only requires the settings parameters that should be updated to be specified, all other parameters will be unmodified or populated from the default values.- You cannot set the Management Console root site administrator password with this API endpoint. Use the `ghe-set-password` utility to change the management console password. For more information, see "[Command-line utilities](https://docs.github.com/enterprise-server@3.12/admin/configuration/configuring-your-enterprise/command-line-utilities#ghe-set-password)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.12/rest/enterprise-admin/manage-ghes#set-settings
func (m *V1ConfigSettingsRequestBuilder) Put(ctx context.Context, body V1ConfigSettingsPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToGetRequestInformation gets a list of settings for a GitHub Enterprise Server instance.
// returns a *RequestInformation when successful
func (m *V1ConfigSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPutRequestInformation updates the settings on your instance. For a list of the available settings, see the [Get settings endpoint](https://docs.github.com/enterprise-server@3.12/rest/enterprise-admin/manage-ghes#get-settings).**Notes:**- The request body only requires the settings parameters that should be updated to be specified, all other parameters will be unmodified or populated from the default values.- You cannot set the Management Console root site administrator password with this API endpoint. Use the `ghe-set-password` utility to change the management console password. For more information, see "[Command-line utilities](https://docs.github.com/enterprise-server@3.12/admin/configuration/configuring-your-enterprise/command-line-utilities#ghe-set-password)."
// returns a *RequestInformation when successful
func (m *V1ConfigSettingsRequestBuilder) ToPutRequestInformation(ctx context.Context, body V1ConfigSettingsPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1ConfigSettingsRequestBuilder when successful
func (m *V1ConfigSettingsRequestBuilder) WithUrl(rawUrl string)(*V1ConfigSettingsRequestBuilder) {
    return NewV1ConfigSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
