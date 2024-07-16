package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// HooksItemPingsRequestBuilder builds and executes requests for operations under \admin\hooks\{hook_id}\pings
type HooksItemPingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewHooksItemPingsRequestBuilderInternal instantiates a new HooksItemPingsRequestBuilder and sets the default values.
func NewHooksItemPingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HooksItemPingsRequestBuilder) {
    m := &HooksItemPingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/hooks/{hook_id}/pings", pathParameters),
    }
    return m
}
// NewHooksItemPingsRequestBuilder instantiates a new HooksItemPingsRequestBuilder and sets the default values.
func NewHooksItemPingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HooksItemPingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHooksItemPingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post this will trigger a [ping event](https://docs.github.com/enterprise-server@3.11/webhooks/#ping-event) to be sent to the webhook.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.11/rest/enterprise-admin/global-webhooks#ping-a-global-webhook
func (m *HooksItemPingsRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation this will trigger a [ping event](https://docs.github.com/enterprise-server@3.11/webhooks/#ping-event) to be sent to the webhook.
// returns a *RequestInformation when successful
func (m *HooksItemPingsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *HooksItemPingsRequestBuilder when successful
func (m *HooksItemPingsRequestBuilder) WithUrl(rawUrl string)(*HooksItemPingsRequestBuilder) {
    return NewHooksItemPingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
