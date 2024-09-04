package manage

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1ClusterRequestBuilder builds and executes requests for operations under \manage\v1\cluster
type V1ClusterRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewV1ClusterRequestBuilderInternal instantiates a new V1ClusterRequestBuilder and sets the default values.
func NewV1ClusterRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ClusterRequestBuilder) {
    m := &V1ClusterRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/manage/v1/cluster", pathParameters),
    }
    return m
}
// NewV1ClusterRequestBuilder instantiates a new V1ClusterRequestBuilder and sets the default values.
func NewV1ClusterRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ClusterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1ClusterRequestBuilderInternal(urlParams, requestAdapter)
}
// Status the status property
// returns a *V1ClusterStatusRequestBuilder when successful
func (m *V1ClusterRequestBuilder) Status()(*V1ClusterStatusRequestBuilder) {
    return NewV1ClusterStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
