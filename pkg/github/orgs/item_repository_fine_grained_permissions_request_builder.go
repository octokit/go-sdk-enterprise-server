package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ItemRepositoryFineGrainedPermissionsRequestBuilder builds and executes requests for operations under \orgs\{org}\repository-fine-grained-permissions
type ItemRepositoryFineGrainedPermissionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemRepositoryFineGrainedPermissionsRequestBuilderInternal instantiates a new ItemRepositoryFineGrainedPermissionsRequestBuilder and sets the default values.
func NewItemRepositoryFineGrainedPermissionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRepositoryFineGrainedPermissionsRequestBuilder) {
    m := &ItemRepositoryFineGrainedPermissionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/repository-fine-grained-permissions", pathParameters),
    }
    return m
}
// NewItemRepositoryFineGrainedPermissionsRequestBuilder instantiates a new ItemRepositoryFineGrainedPermissionsRequestBuilder and sets the default values.
func NewItemRepositoryFineGrainedPermissionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRepositoryFineGrainedPermissionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRepositoryFineGrainedPermissionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the fine-grained permissions that can be used in custom repository roles for an organization. For more information, see "[About custom repository roles](https://docs.github.com/enterprise-server@3.14/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-repository-roles)."The authenticated user must be an administrator of the organization or of a repository of the organization to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `admin:org` or `repo` scope to use this endpoint.
// returns a []RepositoryFineGrainedPermissionable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.14/rest/orgs/custom-roles#list-repository-fine-grained-permissions-for-an-organization
func (m *ItemRepositoryFineGrainedPermissionsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryFineGrainedPermissionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateRepositoryFineGrainedPermissionFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryFineGrainedPermissionable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryFineGrainedPermissionable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists the fine-grained permissions that can be used in custom repository roles for an organization. For more information, see "[About custom repository roles](https://docs.github.com/enterprise-server@3.14/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-repository-roles)."The authenticated user must be an administrator of the organization or of a repository of the organization to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `admin:org` or `repo` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemRepositoryFineGrainedPermissionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRepositoryFineGrainedPermissionsRequestBuilder when successful
func (m *ItemRepositoryFineGrainedPermissionsRequestBuilder) WithUrl(rawUrl string)(*ItemRepositoryFineGrainedPermissionsRequestBuilder) {
    return NewItemRepositoryFineGrainedPermissionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
