package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemReplicasRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\replicas
type ItemItemReplicasRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Caches the caches property
// returns a *ItemItemReplicasCachesRequestBuilder when successful
func (m *ItemItemReplicasRequestBuilder) Caches()(*ItemItemReplicasCachesRequestBuilder) {
    return NewItemItemReplicasCachesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemReplicasRequestBuilderInternal instantiates a new ItemItemReplicasRequestBuilder and sets the default values.
func NewItemItemReplicasRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemReplicasRequestBuilder) {
    m := &ItemItemReplicasRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/replicas", pathParameters),
    }
    return m
}
// NewItemItemReplicasRequestBuilder instantiates a new ItemItemReplicasRequestBuilder and sets the default values.
func NewItemItemReplicasRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemReplicasRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemReplicasRequestBuilderInternal(urlParams, requestAdapter)
}
