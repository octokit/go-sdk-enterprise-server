package setup

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ApiConfigcheckRequestBuilder builds and executes requests for operations under \setup\api\configcheck
type ApiConfigcheckRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewApiConfigcheckRequestBuilderInternal instantiates a new ApiConfigcheckRequestBuilder and sets the default values.
func NewApiConfigcheckRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiConfigcheckRequestBuilder) {
    m := &ApiConfigcheckRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/setup/api/configcheck", pathParameters),
    }
    return m
}
// NewApiConfigcheckRequestBuilder instantiates a new ApiConfigcheckRequestBuilder and sets the default values.
func NewApiConfigcheckRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiConfigcheckRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApiConfigcheckRequestBuilderInternal(urlParams, requestAdapter)
}
// Get this endpoint allows you to check the status of the most recent configuration process:Note that you may need to wait several seconds after you start a process before you can check its status.The different statuses are:| Status        | Description                       || ------------- | --------------------------------- || `PENDING`     | The job has not started yet       || `CONFIGURING` | The job is running                || `DONE`        | The job has finished correctly    || `FAILED`      | The job has finished unexpectedly |
// returns a ConfigurationStatusable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.12/rest/enterprise-admin/management-console#get-the-configuration-status
func (m *ApiConfigcheckRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ConfigurationStatusable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateConfigurationStatusFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ConfigurationStatusable), nil
}
// ToGetRequestInformation this endpoint allows you to check the status of the most recent configuration process:Note that you may need to wait several seconds after you start a process before you can check its status.The different statuses are:| Status        | Description                       || ------------- | --------------------------------- || `PENDING`     | The job has not started yet       || `CONFIGURING` | The job is running                || `DONE`        | The job has finished correctly    || `FAILED`      | The job has finished unexpectedly |
// returns a *RequestInformation when successful
func (m *ApiConfigcheckRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ApiConfigcheckRequestBuilder when successful
func (m *ApiConfigcheckRequestBuilder) WithUrl(rawUrl string)(*ApiConfigcheckRequestBuilder) {
    return NewApiConfigcheckRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
