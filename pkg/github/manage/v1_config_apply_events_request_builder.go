package manage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1ConfigApplyEventsRequestBuilder builds and executes requests for operations under \manage\v1\config\apply\events
type V1ConfigApplyEventsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1ConfigApplyEventsRequestBuilderGetQueryParameters lists events from an in-process `ghe-config-apply` run on your Github Enterprise Server instance.
type V1ConfigApplyEventsRequestBuilderGetQueryParameters struct {
    // The unique ID of the last response from a host, used for pagination.
    Last_request_id *string `uriparametername:"last_request_id"`
}
// NewV1ConfigApplyEventsRequestBuilderInternal instantiates a new V1ConfigApplyEventsRequestBuilder and sets the default values.
func NewV1ConfigApplyEventsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ConfigApplyEventsRequestBuilder) {
    m := &V1ConfigApplyEventsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/manage/v1/config/apply/events{?last_request_id*}", pathParameters),
    }
    return m
}
// NewV1ConfigApplyEventsRequestBuilder instantiates a new V1ConfigApplyEventsRequestBuilder and sets the default values.
func NewV1ConfigApplyEventsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ConfigApplyEventsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1ConfigApplyEventsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists events from an in-process `ghe-config-apply` run on your Github Enterprise Server instance.
// returns a V1ConfigApplyEventsGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.14/rest/enterprise-admin/manage-ghes#list-events-from-ghe-config-apply
func (m *V1ConfigApplyEventsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V1ConfigApplyEventsRequestBuilderGetQueryParameters])(V1ConfigApplyEventsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1ConfigApplyEventsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1ConfigApplyEventsGetResponseable), nil
}
// ToGetRequestInformation lists events from an in-process `ghe-config-apply` run on your Github Enterprise Server instance.
// returns a *RequestInformation when successful
func (m *V1ConfigApplyEventsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V1ConfigApplyEventsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1ConfigApplyEventsRequestBuilder when successful
func (m *V1ConfigApplyEventsRequestBuilder) WithUrl(rawUrl string)(*V1ConfigApplyEventsRequestBuilder) {
    return NewV1ConfigApplyEventsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
