package enterprise

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EnterpriseRequestBuilder builds and executes requests for operations under \enterprise
type EnterpriseRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Announcement the announcement property
// returns a *AnnouncementRequestBuilder when successful
func (m *EnterpriseRequestBuilder) Announcement()(*AnnouncementRequestBuilder) {
    return NewAnnouncementRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEnterpriseRequestBuilderInternal instantiates a new EnterpriseRequestBuilder and sets the default values.
func NewEnterpriseRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseRequestBuilder) {
    m := &EnterpriseRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprise", pathParameters),
    }
    return m
}
// NewEnterpriseRequestBuilder instantiates a new EnterpriseRequestBuilder and sets the default values.
func NewEnterpriseRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnterpriseRequestBuilderInternal(urlParams, requestAdapter)
}
// Settings the settings property
// returns a *SettingsRequestBuilder when successful
func (m *EnterpriseRequestBuilder) Settings()(*SettingsRequestBuilder) {
    return NewSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stats the stats property
// returns a *StatsRequestBuilder when successful
func (m *EnterpriseRequestBuilder) Stats()(*StatsRequestBuilder) {
    return NewStatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
