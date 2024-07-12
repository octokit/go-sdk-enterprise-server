package manage

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1AccessRequestBuilder builds and executes requests for operations under \manage\v1\access
type V1AccessRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewV1AccessRequestBuilderInternal instantiates a new V1AccessRequestBuilder and sets the default values.
func NewV1AccessRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1AccessRequestBuilder) {
    m := &V1AccessRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/manage/v1/access", pathParameters),
    }
    return m
}
// NewV1AccessRequestBuilder instantiates a new V1AccessRequestBuilder and sets the default values.
func NewV1AccessRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1AccessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1AccessRequestBuilderInternal(urlParams, requestAdapter)
}
// Ssh the ssh property
// returns a *V1AccessSshRequestBuilder when successful
func (m *V1AccessRequestBuilder) Ssh()(*V1AccessSshRequestBuilder) {
    return NewV1AccessSshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
