package manage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1ConfigApplyRequestBuilder builds and executes requests for operations under \manage\v1\config\apply
type V1ConfigApplyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1ConfigApplyRequestBuilderGetQueryParameters displays the current status of `ghe-config-apply` in the environment or the status of a historical run by ID.
type V1ConfigApplyRequestBuilderGetQueryParameters struct {
    // The unique run ID of the `ghe-config-apply` run.
    Run_id *string `uriparametername:"run_id"`
}
// NewV1ConfigApplyRequestBuilderInternal instantiates a new V1ConfigApplyRequestBuilder and sets the default values.
func NewV1ConfigApplyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ConfigApplyRequestBuilder) {
    m := &V1ConfigApplyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/manage/v1/config/apply{?run_id*}", pathParameters),
    }
    return m
}
// NewV1ConfigApplyRequestBuilder instantiates a new V1ConfigApplyRequestBuilder and sets the default values.
func NewV1ConfigApplyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ConfigApplyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1ConfigApplyRequestBuilderInternal(urlParams, requestAdapter)
}
// Events the events property
// returns a *V1ConfigApplyEventsRequestBuilder when successful
func (m *V1ConfigApplyRequestBuilder) Events()(*V1ConfigApplyEventsRequestBuilder) {
    return NewV1ConfigApplyEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get displays the current status of `ghe-config-apply` in the environment or the status of a historical run by ID.
// returns a V1ConfigApplyGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.12/rest/enterprise-admin/manage-ghes#get-the-status-of-a-ghe-config-apply-run
func (m *V1ConfigApplyRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V1ConfigApplyRequestBuilderGetQueryParameters])(V1ConfigApplyGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1ConfigApplyGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1ConfigApplyGetResponseable), nil
}
// Post triggers a run of `ghe-config-apply` from the `ghes-manage` agent on your Nomad Delegate instance.You can provide a run ID or allow one to be generated randomly.
// returns a V1ConfigApplyPostResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.12/rest/enterprise-admin/manage-ghes#trigger-a-ghe-config-apply-run
func (m *V1ConfigApplyRequestBuilder) Post(ctx context.Context, body V1ConfigApplyPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(V1ConfigApplyPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1ConfigApplyPostResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1ConfigApplyPostResponseable), nil
}
// ToGetRequestInformation displays the current status of `ghe-config-apply` in the environment or the status of a historical run by ID.
// returns a *RequestInformation when successful
func (m *V1ConfigApplyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V1ConfigApplyRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation triggers a run of `ghe-config-apply` from the `ghes-manage` agent on your Nomad Delegate instance.You can provide a run ID or allow one to be generated randomly.
// returns a *RequestInformation when successful
func (m *V1ConfigApplyRequestBuilder) ToPostRequestInformation(ctx context.Context, body V1ConfigApplyPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1ConfigApplyRequestBuilder when successful
func (m *V1ConfigApplyRequestBuilder) WithUrl(rawUrl string)(*V1ConfigApplyRequestBuilder) {
    return NewV1ConfigApplyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
