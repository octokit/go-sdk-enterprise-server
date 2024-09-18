package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
    ica6934b4a373fa1659e9f86c88392ef5c7c24513cb1e02c6bad9849147d46e61 "github.com/octokit/go-sdk-enterprise-server/pkg/github/repos/item/item/prereceivehooks"
)

// ItemItemPreReceiveHooksRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\pre-receive-hooks
type ItemItemPreReceiveHooksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemPreReceiveHooksRequestBuilderGetQueryParameters list all pre-receive hooks that are enabled or testing for this repository as well as any disabled hooks that are allowed to be enabled at the repository level. Pre-receive hooks that are disabled at a higher level and are not configurable will not be listed.
type ItemItemPreReceiveHooksRequestBuilderGetQueryParameters struct {
    // The direction to sort the results by.
    Direction *ica6934b4a373fa1659e9f86c88392ef5c7c24513cb1e02c6bad9849147d46e61.GetDirectionQueryParameterType `uriparametername:"direction"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    Sort *ica6934b4a373fa1659e9f86c88392ef5c7c24513cb1e02c6bad9849147d46e61.GetSortQueryParameterType `uriparametername:"sort"`
}
// ByPre_receive_hook_id gets an item from the github.com/octokit/go-sdk-enterprise-server/pkg/github.repos.item.item.preReceiveHooks.item collection
// returns a *ItemItemPreReceiveHooksWithPre_receive_hook_ItemRequestBuilder when successful
func (m *ItemItemPreReceiveHooksRequestBuilder) ByPre_receive_hook_id(pre_receive_hook_id int32)(*ItemItemPreReceiveHooksWithPre_receive_hook_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["pre_receive_hook_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(pre_receive_hook_id), 10)
    return NewItemItemPreReceiveHooksWithPre_receive_hook_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemPreReceiveHooksRequestBuilderInternal instantiates a new ItemItemPreReceiveHooksRequestBuilder and sets the default values.
func NewItemItemPreReceiveHooksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPreReceiveHooksRequestBuilder) {
    m := &ItemItemPreReceiveHooksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/pre-receive-hooks{?direction*,page*,per_page*,sort*}", pathParameters),
    }
    return m
}
// NewItemItemPreReceiveHooksRequestBuilder instantiates a new ItemItemPreReceiveHooksRequestBuilder and sets the default values.
func NewItemItemPreReceiveHooksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPreReceiveHooksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemPreReceiveHooksRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all pre-receive hooks that are enabled or testing for this repository as well as any disabled hooks that are allowed to be enabled at the repository level. Pre-receive hooks that are disabled at a higher level and are not configurable will not be listed.
// returns a []RepositoryPreReceiveHookable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.14/rest/enterprise-admin/repo-pre-receive-hooks#list-pre-receive-hooks-for-a-repository
func (m *ItemItemPreReceiveHooksRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemPreReceiveHooksRequestBuilderGetQueryParameters])([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryPreReceiveHookable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateRepositoryPreReceiveHookFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryPreReceiveHookable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryPreReceiveHookable)
        }
    }
    return val, nil
}
// ToGetRequestInformation list all pre-receive hooks that are enabled or testing for this repository as well as any disabled hooks that are allowed to be enabled at the repository level. Pre-receive hooks that are disabled at a higher level and are not configurable will not be listed.
// returns a *RequestInformation when successful
func (m *ItemItemPreReceiveHooksRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemPreReceiveHooksRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemPreReceiveHooksRequestBuilder when successful
func (m *ItemItemPreReceiveHooksRequestBuilder) WithUrl(rawUrl string)(*ItemItemPreReceiveHooksRequestBuilder) {
    return NewItemItemPreReceiveHooksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
