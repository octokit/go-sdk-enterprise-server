package scim

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder builds and executes requests for operations under \scim\v2\enterprises\{enterprise}\Groups\{scim_group_id}
type V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilderGetQueryParameters **Note:** The SCIM API endpoints for enterprise accounts are currently in *private* beta and are subject to change.Gets information about a SCIM group.
type V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilderGetQueryParameters struct {
    // Excludes the specified attribute from being returned in the results. Using this parameter can speed up response time.
    ExcludedAttributes *string `uriparametername:"excludedAttributes"`
}
// NewV2EnterprisesItemGroupsWithScim_group_ItemRequestBuilderInternal instantiates a new V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder and sets the default values.
func NewV2EnterprisesItemGroupsWithScim_group_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder) {
    m := &V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/scim/v2/enterprises/{enterprise}/Groups/{scim_group_id}{?excludedAttributes*}", pathParameters),
    }
    return m
}
// NewV2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder instantiates a new V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder and sets the default values.
func NewV2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV2EnterprisesItemGroupsWithScim_group_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete **Note:** The SCIM API endpoints for enterprise accounts are currently in *private* beta and are subject to change. Deletes a SCIM group from an enterprise.
// returns a ScimError error when the service returns a 400 status code
// returns a BasicError error when the service returns a 404 status code
// returns a ScimError error when the service returns a 429 status code
// returns a ScimError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.10/rest/enterprise-admin/scim#delete-a-scim-group-from-an-enterprise
func (m *V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimErrorFromDiscriminatorValue,
        "404": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
        "429": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimErrorFromDiscriminatorValue,
        "500": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get **Note:** The SCIM API endpoints for enterprise accounts are currently in *private* beta and are subject to change.Gets information about a SCIM group.
// returns a ScimEnterpriseGroupResponseable when successful
// returns a ScimError error when the service returns a 400 status code
// returns a BasicError error when the service returns a 404 status code
// returns a ScimError error when the service returns a 429 status code
// returns a ScimError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.10/rest/enterprise-admin/scim#get-scim-provisioning-information-for-an-enterprise-group
func (m *V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilderGetQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ScimEnterpriseGroupResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimErrorFromDiscriminatorValue,
        "404": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
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
// Patch **Note:** The SCIM API endpoints for enterprise accounts are currently in *private* beta and are subject to change.Update a provisioned group’s individual attributes.To change a group’s values, you must provide a specific Operations JSON format that contains at least one of the add, remove, or replace operations. For examples and more information on the SCIM operations format, see the [SCIM specification](https://tools.ietf.org/html/rfc7644#section-3.5.2).  Update can also be used to add group memberships.Group memberships can be sent one at a time or in batches for faster performance. **Note**: The memberships are referenced through a local user `id`, and the user will need to be created before they are referenced here.
// returns a ScimEnterpriseGroupResponseable when successful
// returns a ScimError error when the service returns a 400 status code
// returns a BasicError error when the service returns a 404 status code
// returns a ScimError error when the service returns a 429 status code
// returns a ScimError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.10/rest/enterprise-admin/scim#update-an-attribute-for-a-scim-enterprise-group
func (m *V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder) Patch(ctx context.Context, body ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.PatchSchemaable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ScimEnterpriseGroupResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimErrorFromDiscriminatorValue,
        "404": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
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
// Put **Note:** The SCIM API endpoints for enterprise accounts are currently in *private* beta and are subject to change.Replaces an existing provisioned group’s information.You must provide all the information required for the group as if you were provisioning it for the first time. Any existing group information that you don't provide will be removed, including group membership. If you want to only update a specific attribute, use the [Update an attribute for a SCIM enterprise group](#update-an-attribute-for-a-scim-enterprise-group) endpoint instead.
// returns a ScimEnterpriseGroupResponseable when successful
// returns a ScimError error when the service returns a 400 status code
// returns a BasicError error when the service returns a 404 status code
// returns a ScimError error when the service returns a 429 status code
// returns a ScimError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.10/rest/enterprise-admin/scim#set-scim-information-for-a-provisioned-enterprise-group
func (m *V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder) Put(ctx context.Context, body ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Groupable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ScimEnterpriseGroupResponseable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimErrorFromDiscriminatorValue,
        "404": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
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
// ToDeleteRequestInformation **Note:** The SCIM API endpoints for enterprise accounts are currently in *private* beta and are subject to change. Deletes a SCIM group from an enterprise.
// returns a *RequestInformation when successful
func (m *V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json, application/scim+json")
    return requestInfo, nil
}
// ToGetRequestInformation **Note:** The SCIM API endpoints for enterprise accounts are currently in *private* beta and are subject to change.Gets information about a SCIM group.
// returns a *RequestInformation when successful
func (m *V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/scim+json")
    return requestInfo, nil
}
// ToPatchRequestInformation **Note:** The SCIM API endpoints for enterprise accounts are currently in *private* beta and are subject to change.Update a provisioned group’s individual attributes.To change a group’s values, you must provide a specific Operations JSON format that contains at least one of the add, remove, or replace operations. For examples and more information on the SCIM operations format, see the [SCIM specification](https://tools.ietf.org/html/rfc7644#section-3.5.2).  Update can also be used to add group memberships.Group memberships can be sent one at a time or in batches for faster performance. **Note**: The memberships are referenced through a local user `id`, and the user will need to be created before they are referenced here.
// returns a *RequestInformation when successful
func (m *V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.PatchSchemaable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/scim+json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// ToPutRequestInformation **Note:** The SCIM API endpoints for enterprise accounts are currently in *private* beta and are subject to change.Replaces an existing provisioned group’s information.You must provide all the information required for the group as if you were provisioning it for the first time. Any existing group information that you don't provide will be removed, including group membership. If you want to only update a specific attribute, use the [Update an attribute for a SCIM enterprise group](#update-an-attribute-for-a-scim-enterprise-group) endpoint instead.
// returns a *RequestInformation when successful
func (m *V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Groupable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/scim+json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder when successful
func (m *V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder) WithUrl(rawUrl string)(*V2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder) {
    return NewV2EnterprisesItemGroupsWithScim_group_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
