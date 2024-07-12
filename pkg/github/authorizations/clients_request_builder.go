package authorizations

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ClientsRequestBuilder builds and executes requests for operations under \authorizations\clients
type ClientsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByClient_id gets an item from the github.com/octokit/go-sdk-enterprise-server/pkg/github.authorizations.clients.item collection
// Deprecated: 
// returns a *ClientsWithClient_ItemRequestBuilder when successful
func (m *ClientsRequestBuilder) ByClient_id(client_id string)(*ClientsWithClient_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if client_id != "" {
        urlTplParams["client_id"] = client_id
    }
    return NewClientsWithClient_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewClientsRequestBuilderInternal instantiates a new ClientsRequestBuilder and sets the default values.
func NewClientsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClientsRequestBuilder) {
    m := &ClientsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/authorizations/clients", pathParameters),
    }
    return m
}
// NewClientsRequestBuilder instantiates a new ClientsRequestBuilder and sets the default values.
func NewClientsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClientsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClientsRequestBuilderInternal(urlParams, requestAdapter)
}
