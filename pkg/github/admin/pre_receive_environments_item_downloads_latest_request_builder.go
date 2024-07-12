package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// PreReceiveEnvironmentsItemDownloadsLatestRequestBuilder builds and executes requests for operations under \admin\pre-receive-environments\{pre_receive_environment_id}\downloads\latest
type PreReceiveEnvironmentsItemDownloadsLatestRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewPreReceiveEnvironmentsItemDownloadsLatestRequestBuilderInternal instantiates a new PreReceiveEnvironmentsItemDownloadsLatestRequestBuilder and sets the default values.
func NewPreReceiveEnvironmentsItemDownloadsLatestRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PreReceiveEnvironmentsItemDownloadsLatestRequestBuilder) {
    m := &PreReceiveEnvironmentsItemDownloadsLatestRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/pre-receive-environments/{pre_receive_environment_id}/downloads/latest", pathParameters),
    }
    return m
}
// NewPreReceiveEnvironmentsItemDownloadsLatestRequestBuilder instantiates a new PreReceiveEnvironmentsItemDownloadsLatestRequestBuilder and sets the default values.
func NewPreReceiveEnvironmentsItemDownloadsLatestRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PreReceiveEnvironmentsItemDownloadsLatestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPreReceiveEnvironmentsItemDownloadsLatestRequestBuilderInternal(urlParams, requestAdapter)
}
// Get in addition to seeing the download status at the "[Get a pre-receive environment](#get-a-pre-receive-environment)" endpoint, there is also this separate endpoint for just the download status.
// returns a PreReceiveEnvironmentDownloadStatusable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.13/rest/enterprise-admin/pre-receive-environments#get-the-download-status-for-a-pre-receive-environment
func (m *PreReceiveEnvironmentsItemDownloadsLatestRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.PreReceiveEnvironmentDownloadStatusable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreatePreReceiveEnvironmentDownloadStatusFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.PreReceiveEnvironmentDownloadStatusable), nil
}
// ToGetRequestInformation in addition to seeing the download status at the "[Get a pre-receive environment](#get-a-pre-receive-environment)" endpoint, there is also this separate endpoint for just the download status.
// returns a *RequestInformation when successful
func (m *PreReceiveEnvironmentsItemDownloadsLatestRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *PreReceiveEnvironmentsItemDownloadsLatestRequestBuilder when successful
func (m *PreReceiveEnvironmentsItemDownloadsLatestRequestBuilder) WithUrl(rawUrl string)(*PreReceiveEnvironmentsItemDownloadsLatestRequestBuilder) {
    return NewPreReceiveEnvironmentsItemDownloadsLatestRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
