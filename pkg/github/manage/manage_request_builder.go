package manage

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ManageRequestBuilder builds and executes requests for operations under \manage
type ManageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewManageRequestBuilderInternal instantiates a new ManageRequestBuilder and sets the default values.
func NewManageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageRequestBuilder) {
    m := &ManageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/manage", pathParameters),
    }
    return m
}
// NewManageRequestBuilder instantiates a new ManageRequestBuilder and sets the default values.
func NewManageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageRequestBuilderInternal(urlParams, requestAdapter)
}
// V1 the v1 property
// returns a *V1RequestBuilder when successful
func (m *ManageRequestBuilder) V1()(*V1RequestBuilder) {
    return NewV1RequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
