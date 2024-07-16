package manage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// V1ChecksSystemRequirementsRequestBuilder builds and executes requests for operations under \manage\v1\checks\system-requirements
type V1ChecksSystemRequirementsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewV1ChecksSystemRequirementsRequestBuilderInternal instantiates a new V1ChecksSystemRequirementsRequestBuilder and sets the default values.
func NewV1ChecksSystemRequirementsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ChecksSystemRequirementsRequestBuilder) {
    m := &V1ChecksSystemRequirementsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/manage/v1/checks/system-requirements", pathParameters),
    }
    return m
}
// NewV1ChecksSystemRequirementsRequestBuilder instantiates a new V1ChecksSystemRequirementsRequestBuilder and sets the default values.
func NewV1ChecksSystemRequirementsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ChecksSystemRequirementsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1ChecksSystemRequirementsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get checks if the minimum requirements for system hardware resources are met on each configured cluster node.This endpoint may take several seconds to reply.
// returns a GhesChecksSystemRequirementsable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.12/rest/enterprise-admin/manage-ghes#get-system-requirements-check-results-on-configured-nodes
func (m *V1ChecksSystemRequirementsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.GhesChecksSystemRequirementsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateGhesChecksSystemRequirementsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.GhesChecksSystemRequirementsable), nil
}
// ToGetRequestInformation checks if the minimum requirements for system hardware resources are met on each configured cluster node.This endpoint may take several seconds to reply.
// returns a *RequestInformation when successful
func (m *V1ChecksSystemRequirementsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1ChecksSystemRequirementsRequestBuilder when successful
func (m *V1ChecksSystemRequirementsRequestBuilder) WithUrl(rawUrl string)(*V1ChecksSystemRequirementsRequestBuilder) {
    return NewV1ChecksSystemRequirementsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
