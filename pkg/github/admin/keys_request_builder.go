package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
    i7b5e92fdca75c4784d3f41cbe18fab270185eef539028487994d96e4d2495c56 "github.com/octokit/go-sdk-enterprise-server/pkg/github/admin/keys"
)

// KeysRequestBuilder builds and executes requests for operations under \admin\keys
type KeysRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// KeysRequestBuilderGetQueryParameters list public keys
type KeysRequestBuilderGetQueryParameters struct {
    // The direction to sort the results by.
    Direction *i7b5e92fdca75c4784d3f41cbe18fab270185eef539028487994d96e4d2495c56.GetDirectionQueryParameterType `uriparametername:"direction"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // Only show public keys accessed after the given time.
    Since *string `uriparametername:"since"`
    Sort *i7b5e92fdca75c4784d3f41cbe18fab270185eef539028487994d96e4d2495c56.GetSortQueryParameterType `uriparametername:"sort"`
}
// ByKey_ids gets an item from the github.com/octokit/go-sdk-enterprise-server/pkg/github.admin.keys.item collection
// returns a *KeysWithKey_idsItemRequestBuilder when successful
func (m *KeysRequestBuilder) ByKey_ids(key_ids string)(*KeysWithKey_idsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if key_ids != "" {
        urlTplParams["key_ids"] = key_ids
    }
    return NewKeysWithKey_idsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewKeysRequestBuilderInternal instantiates a new KeysRequestBuilder and sets the default values.
func NewKeysRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*KeysRequestBuilder) {
    m := &KeysRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/keys{?direction*,page*,per_page*,since*,sort*}", pathParameters),
    }
    return m
}
// NewKeysRequestBuilder instantiates a new KeysRequestBuilder and sets the default values.
func NewKeysRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*KeysRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewKeysRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list public keys
// returns a []PublicKeyFullable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.14/rest/enterprise-admin/users#list-public-keys
func (m *KeysRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[KeysRequestBuilderGetQueryParameters])([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.PublicKeyFullable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreatePublicKeyFullFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.PublicKeyFullable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.PublicKeyFullable)
        }
    }
    return val, nil
}
// returns a *RequestInformation when successful
func (m *KeysRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[KeysRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *KeysRequestBuilder when successful
func (m *KeysRequestBuilder) WithUrl(rawUrl string)(*KeysRequestBuilder) {
    return NewKeysRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
