package orgs

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
    id2409d36c7b20f4f264b157e026cce9c1d36ffdde31f38521d5d3f2e2562a94c "github.com/octokit/go-sdk-enterprise-server/pkg/github/orgs/item/prereceivehooks"
)

// ItemPreReceiveHooksRequestBuilder builds and executes requests for operations under \orgs\{org}\pre-receive-hooks
type ItemPreReceiveHooksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPreReceiveHooksRequestBuilderGetQueryParameters list all pre-receive hooks that are enabled or testing for this organization as well as any disabled hooks that can be configured at the organization level. Globally disabled pre-receive hooks that do not allow downstream configuration are not listed.
type ItemPreReceiveHooksRequestBuilderGetQueryParameters struct {
    // The direction to sort the results by.
    Direction *id2409d36c7b20f4f264b157e026cce9c1d36ffdde31f38521d5d3f2e2562a94c.GetDirectionQueryParameterType `uriparametername:"direction"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.13/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.13/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The sort order for the response collection.
    Sort *id2409d36c7b20f4f264b157e026cce9c1d36ffdde31f38521d5d3f2e2562a94c.GetSortQueryParameterType `uriparametername:"sort"`
}
// ByPre_receive_hook_id gets an item from the github.com/octokit/go-sdk-enterprise-server/pkg/github.orgs.item.preReceiveHooks.item collection
// returns a *ItemPreReceiveHooksWithPre_receive_hook_ItemRequestBuilder when successful
func (m *ItemPreReceiveHooksRequestBuilder) ByPre_receive_hook_id(pre_receive_hook_id int32)(*ItemPreReceiveHooksWithPre_receive_hook_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["pre_receive_hook_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(pre_receive_hook_id), 10)
    return NewItemPreReceiveHooksWithPre_receive_hook_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPreReceiveHooksRequestBuilderInternal instantiates a new ItemPreReceiveHooksRequestBuilder and sets the default values.
func NewItemPreReceiveHooksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPreReceiveHooksRequestBuilder) {
    m := &ItemPreReceiveHooksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/pre-receive-hooks{?direction*,page*,per_page*,sort*}", pathParameters),
    }
    return m
}
// NewItemPreReceiveHooksRequestBuilder instantiates a new ItemPreReceiveHooksRequestBuilder and sets the default values.
func NewItemPreReceiveHooksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPreReceiveHooksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPreReceiveHooksRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all pre-receive hooks that are enabled or testing for this organization as well as any disabled hooks that can be configured at the organization level. Globally disabled pre-receive hooks that do not allow downstream configuration are not listed.
// returns a []OrgPreReceiveHookable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.13/rest/enterprise-admin/org-pre-receive-hooks#list-pre-receive-hooks-for-an-organization
func (m *ItemPreReceiveHooksRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemPreReceiveHooksRequestBuilderGetQueryParameters])([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.OrgPreReceiveHookable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateOrgPreReceiveHookFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.OrgPreReceiveHookable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.OrgPreReceiveHookable)
        }
    }
    return val, nil
}
// ToGetRequestInformation list all pre-receive hooks that are enabled or testing for this organization as well as any disabled hooks that can be configured at the organization level. Globally disabled pre-receive hooks that do not allow downstream configuration are not listed.
// returns a *RequestInformation when successful
func (m *ItemPreReceiveHooksRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemPreReceiveHooksRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemPreReceiveHooksRequestBuilder when successful
func (m *ItemPreReceiveHooksRequestBuilder) WithUrl(rawUrl string)(*ItemPreReceiveHooksRequestBuilder) {
    return NewItemPreReceiveHooksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
