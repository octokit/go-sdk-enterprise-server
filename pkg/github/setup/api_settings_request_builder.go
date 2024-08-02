package setup

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ApiSettingsRequestBuilder builds and executes requests for operations under \setup\api\settings
type ApiSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthorizedKeys the authorizedKeys property
// returns a *ApiSettingsAuthorizedKeysRequestBuilder when successful
func (m *ApiSettingsRequestBuilder) AuthorizedKeys()(*ApiSettingsAuthorizedKeysRequestBuilder) {
    return NewApiSettingsAuthorizedKeysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewApiSettingsRequestBuilderInternal instantiates a new ApiSettingsRequestBuilder and sets the default values.
func NewApiSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiSettingsRequestBuilder) {
    m := &ApiSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/setup/api/settings", pathParameters),
    }
    return m
}
// NewApiSettingsRequestBuilder instantiates a new ApiSettingsRequestBuilder and sets the default values.
func NewApiSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApiSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the settings for your instance. To change settings, see the [Set settings endpoint](https://docs.github.com/enterprise-server@3.10/rest/enterprise-admin/management-console#set-settings).> [!NOTE]> You cannot retrieve the management console password with the Enterprise administration API.
// returns a EnterpriseSettingsable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.10/rest/enterprise-admin/management-console#get-settings
func (m *ApiSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.EnterpriseSettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateEnterpriseSettingsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.EnterpriseSettingsable), nil
}
// Put applies settings on your instance. For a list of the available settings, see the [Get settings endpoint](https://docs.github.com/enterprise-server@3.10/rest/enterprise-admin/management-console#get-settings).**Notes:**- The request body for this operation must be submitted as `application/x-www-form-urlencoded` data. You can submit a parameter value as a string, or you can use a tool such as `curl` to submit a parameter value as the contents of a text file. For more information, see the [`curl` documentation](https://curl.se/docs/manpage.html#--data-urlencode).- You cannot set the management console password with the Enterprise administration API. Use the `ghe-set-password` utility to change the management console password. For more information, see "[Command-line utilities](https://docs.github.com/enterprise-server@3.10/admin/configuration/configuring-your-enterprise/command-line-utilities#ghe-set-password)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.10/rest/enterprise-admin/management-console#set-settings
func (m *ApiSettingsRequestBuilder) Put(ctx context.Context, body ApiSettingsPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToGetRequestInformation gets the settings for your instance. To change settings, see the [Set settings endpoint](https://docs.github.com/enterprise-server@3.10/rest/enterprise-admin/management-console#set-settings).> [!NOTE]> You cannot retrieve the management console password with the Enterprise administration API.
// returns a *RequestInformation when successful
func (m *ApiSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPutRequestInformation applies settings on your instance. For a list of the available settings, see the [Get settings endpoint](https://docs.github.com/enterprise-server@3.10/rest/enterprise-admin/management-console#get-settings).**Notes:**- The request body for this operation must be submitted as `application/x-www-form-urlencoded` data. You can submit a parameter value as a string, or you can use a tool such as `curl` to submit a parameter value as the contents of a text file. For more information, see the [`curl` documentation](https://curl.se/docs/manpage.html#--data-urlencode).- You cannot set the management console password with the Enterprise administration API. Use the `ghe-set-password` utility to change the management console password. For more information, see "[Command-line utilities](https://docs.github.com/enterprise-server@3.10/admin/configuration/configuring-your-enterprise/command-line-utilities#ghe-set-password)."
// returns a *RequestInformation when successful
func (m *ApiSettingsRequestBuilder) ToPutRequestInformation(ctx context.Context, body ApiSettingsPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/x-www-form-urlencoded", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ApiSettingsRequestBuilder when successful
func (m *ApiSettingsRequestBuilder) WithUrl(rawUrl string)(*ApiSettingsRequestBuilder) {
    return NewApiSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
