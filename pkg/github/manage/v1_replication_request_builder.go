package manage

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1ReplicationRequestBuilder builds and executes requests for operations under \manage\v1\replication
type V1ReplicationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewV1ReplicationRequestBuilderInternal instantiates a new V1ReplicationRequestBuilder and sets the default values.
func NewV1ReplicationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ReplicationRequestBuilder) {
    m := &V1ReplicationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/manage/v1/replication", pathParameters),
    }
    return m
}
// NewV1ReplicationRequestBuilder instantiates a new V1ReplicationRequestBuilder and sets the default values.
func NewV1ReplicationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ReplicationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1ReplicationRequestBuilderInternal(urlParams, requestAdapter)
}
// Status the status property
// returns a *V1ReplicationStatusRequestBuilder when successful
func (m *V1ReplicationRequestBuilder) Status()(*V1ReplicationStatusRequestBuilder) {
    return NewV1ReplicationStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
