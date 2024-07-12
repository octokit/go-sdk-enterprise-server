package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ItemExternalGroupWithGroup_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\external-group\{group_id}
type ItemExternalGroupWithGroup_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemExternalGroupWithGroup_ItemRequestBuilderInternal instantiates a new ItemExternalGroupWithGroup_ItemRequestBuilder and sets the default values.
func NewItemExternalGroupWithGroup_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemExternalGroupWithGroup_ItemRequestBuilder) {
    m := &ItemExternalGroupWithGroup_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/external-group/{group_id}", pathParameters),
    }
    return m
}
// NewItemExternalGroupWithGroup_ItemRequestBuilder instantiates a new ItemExternalGroupWithGroup_ItemRequestBuilder and sets the default values.
func NewItemExternalGroupWithGroup_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemExternalGroupWithGroup_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemExternalGroupWithGroup_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get displays information about the specific group's usage.  Provides a list of the group's external members as well as a list of teams that this group is connected to.You can manage team membership with your identity provider using Enterprise Managed Users for GitHub Enterprise Cloud. For more information, see "[GitHub's products](https://docs.github.com/enterprise-server@3.12/github/getting-started-with-github/githubs-products)" in the GitHub Help documentation.
// returns a ExternalGroupable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.12/rest/teams/external-groups#get-an-external-group
func (m *ItemExternalGroupWithGroup_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ExternalGroupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateExternalGroupFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ExternalGroupable), nil
}
// ToGetRequestInformation displays information about the specific group's usage.  Provides a list of the group's external members as well as a list of teams that this group is connected to.You can manage team membership with your identity provider using Enterprise Managed Users for GitHub Enterprise Cloud. For more information, see "[GitHub's products](https://docs.github.com/enterprise-server@3.12/github/getting-started-with-github/githubs-products)" in the GitHub Help documentation.
// returns a *RequestInformation when successful
func (m *ItemExternalGroupWithGroup_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemExternalGroupWithGroup_ItemRequestBuilder when successful
func (m *ItemExternalGroupWithGroup_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemExternalGroupWithGroup_ItemRequestBuilder) {
    return NewItemExternalGroupWithGroup_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
