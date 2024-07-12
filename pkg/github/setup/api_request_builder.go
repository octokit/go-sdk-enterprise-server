package setup

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ApiRequestBuilder builds and executes requests for operations under \setup\api
type ApiRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Configcheck the configcheck property
// returns a *ApiConfigcheckRequestBuilder when successful
func (m *ApiRequestBuilder) Configcheck()(*ApiConfigcheckRequestBuilder) {
    return NewApiConfigcheckRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Configure the configure property
// returns a *ApiConfigureRequestBuilder when successful
func (m *ApiRequestBuilder) Configure()(*ApiConfigureRequestBuilder) {
    return NewApiConfigureRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewApiRequestBuilderInternal instantiates a new ApiRequestBuilder and sets the default values.
func NewApiRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiRequestBuilder) {
    m := &ApiRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/setup/api", pathParameters),
    }
    return m
}
// NewApiRequestBuilder instantiates a new ApiRequestBuilder and sets the default values.
func NewApiRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApiRequestBuilderInternal(urlParams, requestAdapter)
}
// Maintenance the maintenance property
// returns a *ApiMaintenanceRequestBuilder when successful
func (m *ApiRequestBuilder) Maintenance()(*ApiMaintenanceRequestBuilder) {
    return NewApiMaintenanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Settings the settings property
// returns a *ApiSettingsRequestBuilder when successful
func (m *ApiRequestBuilder) Settings()(*ApiSettingsRequestBuilder) {
    return NewApiSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Start the start property
// returns a *ApiStartRequestBuilder when successful
func (m *ApiRequestBuilder) Start()(*ApiStartRequestBuilder) {
    return NewApiStartRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Upgrade the upgrade property
// returns a *ApiUpgradeRequestBuilder when successful
func (m *ApiRequestBuilder) Upgrade()(*ApiUpgradeRequestBuilder) {
    return NewApiUpgradeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
