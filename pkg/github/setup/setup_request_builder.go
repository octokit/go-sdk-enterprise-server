package setup

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SetupRequestBuilder builds and executes requests for operations under \setup
type SetupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Api the api property
// returns a *ApiRequestBuilder when successful
func (m *SetupRequestBuilder) Api()(*ApiRequestBuilder) {
    return NewApiRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewSetupRequestBuilderInternal instantiates a new SetupRequestBuilder and sets the default values.
func NewSetupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetupRequestBuilder) {
    m := &SetupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/setup", pathParameters),
    }
    return m
}
// NewSetupRequestBuilder instantiates a new SetupRequestBuilder and sets the default values.
func NewSetupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetupRequestBuilderInternal(urlParams, requestAdapter)
}
