package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// PreReceiveEnvironmentsItemDownloadsRequestBuilder builds and executes requests for operations under \admin\pre-receive-environments\{pre_receive_environment_id}\downloads
type PreReceiveEnvironmentsItemDownloadsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewPreReceiveEnvironmentsItemDownloadsRequestBuilderInternal instantiates a new PreReceiveEnvironmentsItemDownloadsRequestBuilder and sets the default values.
func NewPreReceiveEnvironmentsItemDownloadsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PreReceiveEnvironmentsItemDownloadsRequestBuilder) {
    m := &PreReceiveEnvironmentsItemDownloadsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/pre-receive-environments/{pre_receive_environment_id}/downloads", pathParameters),
    }
    return m
}
// NewPreReceiveEnvironmentsItemDownloadsRequestBuilder instantiates a new PreReceiveEnvironmentsItemDownloadsRequestBuilder and sets the default values.
func NewPreReceiveEnvironmentsItemDownloadsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PreReceiveEnvironmentsItemDownloadsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPreReceiveEnvironmentsItemDownloadsRequestBuilderInternal(urlParams, requestAdapter)
}
// Latest the latest property
// returns a *PreReceiveEnvironmentsItemDownloadsLatestRequestBuilder when successful
func (m *PreReceiveEnvironmentsItemDownloadsRequestBuilder) Latest()(*PreReceiveEnvironmentsItemDownloadsLatestRequestBuilder) {
    return NewPreReceiveEnvironmentsItemDownloadsLatestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post triggers a new download of the environment tarball from the environment's `image_url`. When the download is finished, the newly downloaded tarball will overwrite the existing environment.If a download cannot be triggered, you will receive a `422 Unprocessable Entity` response.The possible error messages are:* _Cannot modify or delete the default environment_* _Can not start a new download when a download is in progress_
// returns a PreReceiveEnvironmentDownloadStatusable when successful
// returns a PreReceiveEnvironmentsItemDownloadsPreReceiveEnvironmentDownloadStatus422Error error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.14/rest/enterprise-admin/pre-receive-environments#start-a-pre-receive-environment-download
func (m *PreReceiveEnvironmentsItemDownloadsRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.PreReceiveEnvironmentDownloadStatusable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": CreatePreReceiveEnvironmentsItemDownloadsPreReceiveEnvironmentDownloadStatus422ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreatePreReceiveEnvironmentDownloadStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.PreReceiveEnvironmentDownloadStatusable), nil
}
// ToPostRequestInformation triggers a new download of the environment tarball from the environment's `image_url`. When the download is finished, the newly downloaded tarball will overwrite the existing environment.If a download cannot be triggered, you will receive a `422 Unprocessable Entity` response.The possible error messages are:* _Cannot modify or delete the default environment_* _Can not start a new download when a download is in progress_
// returns a *RequestInformation when successful
func (m *PreReceiveEnvironmentsItemDownloadsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *PreReceiveEnvironmentsItemDownloadsRequestBuilder when successful
func (m *PreReceiveEnvironmentsItemDownloadsRequestBuilder) WithUrl(rawUrl string)(*PreReceiveEnvironmentsItemDownloadsRequestBuilder) {
    return NewPreReceiveEnvironmentsItemDownloadsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
