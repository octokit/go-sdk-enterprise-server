package scim

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// V2EnterprisesItemGroupsRequestBuilder builds and executes requests for operations under \scim\v2\enterprises\{enterprise}\Groups
type V2EnterprisesItemGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V2EnterprisesItemGroupsRequestBuilderGetQueryParameters **Note:** The SCIM API endpoints for enterprise accounts are currently in *private* beta and are subject to change.Lists provisioned SCIM groups in an enterprise.You can improve query search time by using the `excludedAttributes` query parameter with a value of `members` to exclude members from the response.
type V2EnterprisesItemGroupsRequestBuilderGetQueryParameters struct {
    // Used for pagination: the number of results to return per page.
    Count *int32 `uriparametername:"count"`
    // Excludes the specified attribute from being returned in the results. Using this parameter can speed up response time.
    ExcludedAttributes *string `uriparametername:"excludedAttributes"`
    // If specified, only results that match the specified filter will be returned. Multiple filters are not supported. Possible filters are `externalId`, `id`, and `displayName`. For example, `?filter="externalId eq '9138790-10932-109120392-12321'"`.
    Filter *string `uriparametername:"filter"`
    // Used for pagination: the starting index of the first result to return when paginating through values.
    StartIndex *int32 `uriparametername:"startIndex"`
}
// ByScim_group_id gets an item from the github.com/octokit/go-sdk-enterprise-server/pkg/github.scim.v2.enterprises.item.Groups.item collection
// returns a *V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder when successful
func (m *V2EnterprisesItemGroupsRequestBuilder) ByScim_group_id(scim_group_id string)(*V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if scim_group_id != "" {
        urlTplParams["scim_group_id"] = scim_group_id
    }
    return NewV2EnterprisesItemGroupsWithScim_group_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV2EnterprisesItemGroupsRequestBuilderInternal instantiates a new V2EnterprisesItemGroupsRequestBuilder and sets the default values.
func NewV2EnterprisesItemGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V2EnterprisesItemGroupsRequestBuilder) {
    m := &V2EnterprisesItemGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/scim/v2/enterprises/{enterprise}/Groups{?count*,excludedAttributes*,filter*,startIndex*}", pathParameters),
    }
    return m
}
// NewV2EnterprisesItemGroupsRequestBuilder instantiates a new V2EnterprisesItemGroupsRequestBuilder and sets the default values.
func NewV2EnterprisesItemGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V2EnterprisesItemGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV2EnterprisesItemGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get **Note:** The SCIM API endpoints for enterprise accounts are currently in *private* beta and are subject to change.Lists provisioned SCIM groups in an enterprise.You can improve query search time by using the `excludedAttributes` query parameter with a value of `members` to exclude members from the response.
// returns a ScimEnterpriseGroupListable when successful
// returns a ScimError error when the service returns a 400 status code
// returns a ScimError error when the service returns a 429 status code
// returns a ScimError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.13/rest/enterprise-admin/scim#list-provisioned-scim-groups-for-an-enterprise
func (m *V2EnterprisesItemGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V2EnterprisesItemGroupsRequestBuilderGetQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ScimEnterpriseGroupListable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimErrorFromDiscriminatorValue,
        "429": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimErrorFromDiscriminatorValue,
        "500": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimEnterpriseGroupListFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ScimEnterpriseGroupListable), nil
}
// Post **Note:** The SCIM API endpoints for enterprise accounts are currently in *private* beta and are subject to change.Creates a SCIM group for an enterprise.If members are included as part of the group provisioning payload, they will be created as external group members. It is up to a provider to store a mapping between the `externalId` and `id` of each user.
// returns a ScimEnterpriseGroupResponseable when successful
// returns a ScimError error when the service returns a 400 status code
// returns a ScimError error when the service returns a 429 status code
// returns a ScimError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.13/rest/enterprise-admin/scim#provision-a-scim-enterprise-group
func (m *V2EnterprisesItemGroupsRequestBuilder) Post(ctx context.Context, body ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Groupable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ScimEnterpriseGroupResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimErrorFromDiscriminatorValue,
        "429": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimErrorFromDiscriminatorValue,
        "500": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimEnterpriseGroupResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ScimEnterpriseGroupResponseable), nil
}
// ToGetRequestInformation **Note:** The SCIM API endpoints for enterprise accounts are currently in *private* beta and are subject to change.Lists provisioned SCIM groups in an enterprise.You can improve query search time by using the `excludedAttributes` query parameter with a value of `members` to exclude members from the response.
// returns a *RequestInformation when successful
func (m *V2EnterprisesItemGroupsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V2EnterprisesItemGroupsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/scim+json")
    return requestInfo, nil
}
// ToPostRequestInformation **Note:** The SCIM API endpoints for enterprise accounts are currently in *private* beta and are subject to change.Creates a SCIM group for an enterprise.If members are included as part of the group provisioning payload, they will be created as external group members. It is up to a provider to store a mapping between the `externalId` and `id` of each user.
// returns a *RequestInformation when successful
func (m *V2EnterprisesItemGroupsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Groupable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/scim+json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V2EnterprisesItemGroupsRequestBuilder when successful
func (m *V2EnterprisesItemGroupsRequestBuilder) WithUrl(rawUrl string)(*V2EnterprisesItemGroupsRequestBuilder) {
    return NewV2EnterprisesItemGroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
