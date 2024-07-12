package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// PreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder builds and executes requests for operations under \admin\pre-receive-environments\{pre_receive_environment_id}
type PreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewPreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilderInternal instantiates a new PreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder and sets the default values.
func NewPreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder) {
    m := &PreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/pre-receive-environments/{pre_receive_environment_id}", pathParameters),
    }
    return m
}
// NewPreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder instantiates a new PreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder and sets the default values.
func NewPreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete if you attempt to delete an environment that cannot be deleted, you will receive a `422 Unprocessable Entity` response.The possible error messages are:*   _Cannot modify or delete the default environment_*   _Cannot delete environment that has hooks_*   _Cannot delete environment when download is in progress_
// returns a PreReceiveEnvironmentsItemWithPre_receive_environment_422Error error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.13/rest/enterprise-admin/pre-receive-environments#delete-a-pre-receive-environment
func (m *PreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": CreatePreReceiveEnvironmentsItemWithPre_receive_environment_422ErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Downloads the downloads property
// returns a *PreReceiveEnvironmentsItemDownloadsRequestBuilder when successful
func (m *PreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder) Downloads()(*PreReceiveEnvironmentsItemDownloadsRequestBuilder) {
    return NewPreReceiveEnvironmentsItemDownloadsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a pre-receive environment
// returns a PreReceiveEnvironmentable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.13/rest/enterprise-admin/pre-receive-environments#get-a-pre-receive-environment
func (m *PreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.PreReceiveEnvironmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreatePreReceiveEnvironmentFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.PreReceiveEnvironmentable), nil
}
// Patch you cannot modify the default environment. If you attempt to modify the default environment, you will receive a `422 Unprocessable Entity` response.
// returns a PreReceiveEnvironmentable when successful
// returns a PreReceiveEnvironmentsItemPreReceiveEnvironment422Error error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.13/rest/enterprise-admin/pre-receive-environments#update-a-pre-receive-environment
func (m *PreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder) Patch(ctx context.Context, body PreReceiveEnvironmentsItemWithPre_receive_environment_PatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.PreReceiveEnvironmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": CreatePreReceiveEnvironmentsItemPreReceiveEnvironment422ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreatePreReceiveEnvironmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.PreReceiveEnvironmentable), nil
}
// ToDeleteRequestInformation if you attempt to delete an environment that cannot be deleted, you will receive a `422 Unprocessable Entity` response.The possible error messages are:*   _Cannot modify or delete the default environment_*   _Cannot delete environment that has hooks_*   _Cannot delete environment when download is in progress_
// returns a *RequestInformation when successful
func (m *PreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// returns a *RequestInformation when successful
func (m *PreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation you cannot modify the default environment. If you attempt to modify the default environment, you will receive a `422 Unprocessable Entity` response.
// returns a *RequestInformation when successful
func (m *PreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body PreReceiveEnvironmentsItemWithPre_receive_environment_PatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *PreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder when successful
func (m *PreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder) WithUrl(rawUrl string)(*PreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder) {
    return NewPreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
