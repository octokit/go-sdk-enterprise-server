package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ItemSecurityManagersRequestBuilder builds and executes requests for operations under \orgs\{org}\security-managers
type ItemSecurityManagersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemSecurityManagersRequestBuilderInternal instantiates a new ItemSecurityManagersRequestBuilder and sets the default values.
func NewItemSecurityManagersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityManagersRequestBuilder) {
    m := &ItemSecurityManagersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/security-managers", pathParameters),
    }
    return m
}
// NewItemSecurityManagersRequestBuilder instantiates a new ItemSecurityManagersRequestBuilder and sets the default values.
func NewItemSecurityManagersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityManagersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSecurityManagersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists teams that are security managers for an organization. For more information, see "[Managing security managers in your organization](https://docs.github.com/enterprise-server@3.11/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."The authenticated user must be an administrator or security manager for the organization to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `read:org` scope to use this endpoint.
// returns a []TeamSimpleable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.11/rest/orgs/security-managers#list-security-manager-teams
func (m *ItemSecurityManagersRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.TeamSimpleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateTeamSimpleFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.TeamSimpleable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.TeamSimpleable)
        }
    }
    return val, nil
}
// Teams the teams property
// returns a *ItemSecurityManagersTeamsRequestBuilder when successful
func (m *ItemSecurityManagersRequestBuilder) Teams()(*ItemSecurityManagersTeamsRequestBuilder) {
    return NewItemSecurityManagersTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation lists teams that are security managers for an organization. For more information, see "[Managing security managers in your organization](https://docs.github.com/enterprise-server@3.11/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."The authenticated user must be an administrator or security manager for the organization to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `read:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemSecurityManagersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSecurityManagersRequestBuilder when successful
func (m *ItemSecurityManagersRequestBuilder) WithUrl(rawUrl string)(*ItemSecurityManagersRequestBuilder) {
    return NewItemSecurityManagersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
