package manage

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1ChecksRequestBuilder builds and executes requests for operations under \manage\v1\checks
type V1ChecksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewV1ChecksRequestBuilderInternal instantiates a new V1ChecksRequestBuilder and sets the default values.
func NewV1ChecksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ChecksRequestBuilder) {
    m := &V1ChecksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/manage/v1/checks", pathParameters),
    }
    return m
}
// NewV1ChecksRequestBuilder instantiates a new V1ChecksRequestBuilder and sets the default values.
func NewV1ChecksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ChecksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1ChecksRequestBuilderInternal(urlParams, requestAdapter)
}
// SystemRequirements the systemRequirements property
// returns a *V1ChecksSystemRequirementsRequestBuilder when successful
func (m *V1ChecksRequestBuilder) SystemRequirements()(*V1ChecksSystemRequirementsRequestBuilder) {
    return NewV1ChecksSystemRequirementsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
