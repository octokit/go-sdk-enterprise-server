package manage

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1ConfigRequestBuilder builds and executes requests for operations under \manage\v1\config
type V1ConfigRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Apply the apply property
// returns a *V1ConfigApplyRequestBuilder when successful
func (m *V1ConfigRequestBuilder) Apply()(*V1ConfigApplyRequestBuilder) {
    return NewV1ConfigApplyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewV1ConfigRequestBuilderInternal instantiates a new V1ConfigRequestBuilder and sets the default values.
func NewV1ConfigRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ConfigRequestBuilder) {
    m := &V1ConfigRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/manage/v1/config", pathParameters),
    }
    return m
}
// NewV1ConfigRequestBuilder instantiates a new V1ConfigRequestBuilder and sets the default values.
func NewV1ConfigRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ConfigRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1ConfigRequestBuilderInternal(urlParams, requestAdapter)
}
// Init the init property
// returns a *V1ConfigInitRequestBuilder when successful
func (m *V1ConfigRequestBuilder) Init()(*V1ConfigInitRequestBuilder) {
    return NewV1ConfigInitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// License the license property
// returns a *V1ConfigLicenseRequestBuilder when successful
func (m *V1ConfigRequestBuilder) License()(*V1ConfigLicenseRequestBuilder) {
    return NewV1ConfigLicenseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Nodes the nodes property
// returns a *V1ConfigNodesRequestBuilder when successful
func (m *V1ConfigRequestBuilder) Nodes()(*V1ConfigNodesRequestBuilder) {
    return NewV1ConfigNodesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Settings the settings property
// returns a *V1ConfigSettingsRequestBuilder when successful
func (m *V1ConfigRequestBuilder) Settings()(*V1ConfigSettingsRequestBuilder) {
    return NewV1ConfigSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
