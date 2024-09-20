package admin

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
    ia19977d2cea6ae9d082220e4d6bf123ad448b9e828531d1ac9afdbee0875c77b "github.com/octokit/go-sdk-enterprise-server/pkg/github/admin/prereceiveenvironments"
)

// PreReceiveEnvironmentsRequestBuilder builds and executes requests for operations under \admin\pre-receive-environments
type PreReceiveEnvironmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PreReceiveEnvironmentsRequestBuilderGetQueryParameters list pre-receive environments
type PreReceiveEnvironmentsRequestBuilderGetQueryParameters struct {
    // The direction to sort the results by.
    Direction *ia19977d2cea6ae9d082220e4d6bf123ad448b9e828531d1ac9afdbee0875c77b.GetDirectionQueryParameterType `uriparametername:"direction"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    Sort *ia19977d2cea6ae9d082220e4d6bf123ad448b9e828531d1ac9afdbee0875c77b.GetSortQueryParameterType `uriparametername:"sort"`
}
// ByPre_receive_environment_id gets an item from the github.com/octokit/go-sdk-enterprise-server/pkg/github.admin.preReceiveEnvironments.item collection
// returns a *PreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder when successful
func (m *PreReceiveEnvironmentsRequestBuilder) ByPre_receive_environment_id(pre_receive_environment_id int32)(*PreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["pre_receive_environment_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(pre_receive_environment_id), 10)
    return NewPreReceiveEnvironmentsWithPre_receive_environment_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPreReceiveEnvironmentsRequestBuilderInternal instantiates a new PreReceiveEnvironmentsRequestBuilder and sets the default values.
func NewPreReceiveEnvironmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PreReceiveEnvironmentsRequestBuilder) {
    m := &PreReceiveEnvironmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/pre-receive-environments{?direction*,page*,per_page*,sort*}", pathParameters),
    }
    return m
}
// NewPreReceiveEnvironmentsRequestBuilder instantiates a new PreReceiveEnvironmentsRequestBuilder and sets the default values.
func NewPreReceiveEnvironmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PreReceiveEnvironmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPreReceiveEnvironmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list pre-receive environments
// returns a []PreReceiveEnvironmentable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.14/rest/enterprise-admin/pre-receive-environments#list-pre-receive-environments
func (m *PreReceiveEnvironmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[PreReceiveEnvironmentsRequestBuilderGetQueryParameters])([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.PreReceiveEnvironmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreatePreReceiveEnvironmentFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.PreReceiveEnvironmentable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.PreReceiveEnvironmentable)
        }
    }
    return val, nil
}
// Post create a pre-receive environment
// returns a PreReceiveEnvironmentable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.14/rest/enterprise-admin/pre-receive-environments#create-a-pre-receive-environment
func (m *PreReceiveEnvironmentsRequestBuilder) Post(ctx context.Context, body PreReceiveEnvironmentsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.PreReceiveEnvironmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreatePreReceiveEnvironmentFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.PreReceiveEnvironmentable), nil
}
// returns a *RequestInformation when successful
func (m *PreReceiveEnvironmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[PreReceiveEnvironmentsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// returns a *RequestInformation when successful
func (m *PreReceiveEnvironmentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body PreReceiveEnvironmentsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *PreReceiveEnvironmentsRequestBuilder when successful
func (m *PreReceiveEnvironmentsRequestBuilder) WithUrl(rawUrl string)(*PreReceiveEnvironmentsRequestBuilder) {
    return NewPreReceiveEnvironmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
