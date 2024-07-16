package scim

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// V2EnterprisesItemUsersRequestBuilder builds and executes requests for operations under \scim\v2\enterprises\{enterprise}\Users
type V2EnterprisesItemUsersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V2EnterprisesItemUsersRequestBuilderGetQueryParameters **Note:** The SCIM API endpoints for enterprise accounts are currently in *private* beta and are subject to change.Lists provisioned SCIM enterprise members.When a user with a SCIM-provisioned external identity is removed from an enterprise through a `patch` with `active` flag set to `false`, the account's metadata is preserved to allow the user to re-join the enterprise in the future. However, the user's account will be suspended and the user will not be able to sign-in. In order to permanently suspend the users account with no ability to re-join the enterprise in the future, use the `delete` request. Users that were not permanently deleted will be visible in the returned results.
type V2EnterprisesItemUsersRequestBuilderGetQueryParameters struct {
    // Used for pagination: the number of results to return per page.
    Count *int32 `uriparametername:"count"`
    // If specified, only results that match the specified filter will be returned. Multiple filters are not supported. Possible filters are `userName`, `externalId`, `id`, and `displayName`. For example, `?filter="externalId eq '9138790-10932-109120392-12321'"`.
    Filter *string `uriparametername:"filter"`
    // Used for pagination: the starting index of the first result to return when paginating through values.
    StartIndex *int32 `uriparametername:"startIndex"`
}
// ByScim_user_id gets an item from the github.com/octokit/go-sdk-enterprise-server/pkg/github.scim.v2.enterprises.item.Users.item collection
// returns a *V2EnterprisesItemUsersWithScim_user_ItemRequestBuilder when successful
func (m *V2EnterprisesItemUsersRequestBuilder) ByScim_user_id(scim_user_id string)(*V2EnterprisesItemUsersWithScim_user_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if scim_user_id != "" {
        urlTplParams["scim_user_id"] = scim_user_id
    }
    return NewV2EnterprisesItemUsersWithScim_user_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewV2EnterprisesItemUsersRequestBuilderInternal instantiates a new V2EnterprisesItemUsersRequestBuilder and sets the default values.
func NewV2EnterprisesItemUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V2EnterprisesItemUsersRequestBuilder) {
    m := &V2EnterprisesItemUsersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/scim/v2/enterprises/{enterprise}/Users{?count*,filter*,startIndex*}", pathParameters),
    }
    return m
}
// NewV2EnterprisesItemUsersRequestBuilder instantiates a new V2EnterprisesItemUsersRequestBuilder and sets the default values.
func NewV2EnterprisesItemUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V2EnterprisesItemUsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV2EnterprisesItemUsersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get **Note:** The SCIM API endpoints for enterprise accounts are currently in *private* beta and are subject to change.Lists provisioned SCIM enterprise members.When a user with a SCIM-provisioned external identity is removed from an enterprise through a `patch` with `active` flag set to `false`, the account's metadata is preserved to allow the user to re-join the enterprise in the future. However, the user's account will be suspended and the user will not be able to sign-in. In order to permanently suspend the users account with no ability to re-join the enterprise in the future, use the `delete` request. Users that were not permanently deleted will be visible in the returned results.
// returns a ScimEnterpriseUserListable when successful
// returns a ScimError error when the service returns a 400 status code
// returns a ScimError error when the service returns a 429 status code
// returns a ScimError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.10/rest/enterprise-admin/scim#list-scim-provisioned-identities-for-an-enterprise
func (m *V2EnterprisesItemUsersRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V2EnterprisesItemUsersRequestBuilderGetQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ScimEnterpriseUserListable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimErrorFromDiscriminatorValue,
        "429": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimErrorFromDiscriminatorValue,
        "500": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimEnterpriseUserListFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ScimEnterpriseUserListable), nil
}
// Post **Note:** The SCIM API endpoints for enterprise accounts are currently in *private* beta and are subject to change.Creates an external identity for a new SCIM enterprise user.SCIM does not authenticate users, it only provisions them. The authentication of users is done by SAML. However, when SCIM is enabled, all users need to be provisioned through SCIM before a user can sign in through SAML. The matching of a user to a SCIM provisioned user is done when the SAML assertion is consumed. The user will be matched on SAML response `NameID` to SCIM `userName`.When converting existing enterprise to use SCIM, the user handle (`userName`) from the SCIM payload will be used to match the provisioned user to an already existing user in the enterprise. Since the new identity record is created for newly provisioned users the matching for those records is done using a user's handle. Currently the matching will be performed to all of the users no matter if they were SAML JIT provisioned or created as local users.
// returns a ScimEnterpriseUserResponseable when successful
// returns a ScimError error when the service returns a 400 status code
// returns a ScimError error when the service returns a 429 status code
// returns a ScimError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.10/rest/enterprise-admin/scim#provision-a-scim-enterprise-user
func (m *V2EnterprisesItemUsersRequestBuilder) Post(ctx context.Context, body ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Userable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ScimEnterpriseUserResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimErrorFromDiscriminatorValue,
        "429": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimErrorFromDiscriminatorValue,
        "500": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateScimEnterpriseUserResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ScimEnterpriseUserResponseable), nil
}
// ToGetRequestInformation **Note:** The SCIM API endpoints for enterprise accounts are currently in *private* beta and are subject to change.Lists provisioned SCIM enterprise members.When a user with a SCIM-provisioned external identity is removed from an enterprise through a `patch` with `active` flag set to `false`, the account's metadata is preserved to allow the user to re-join the enterprise in the future. However, the user's account will be suspended and the user will not be able to sign-in. In order to permanently suspend the users account with no ability to re-join the enterprise in the future, use the `delete` request. Users that were not permanently deleted will be visible in the returned results.
// returns a *RequestInformation when successful
func (m *V2EnterprisesItemUsersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[V2EnterprisesItemUsersRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/scim+json")
    return requestInfo, nil
}
// ToPostRequestInformation **Note:** The SCIM API endpoints for enterprise accounts are currently in *private* beta and are subject to change.Creates an external identity for a new SCIM enterprise user.SCIM does not authenticate users, it only provisions them. The authentication of users is done by SAML. However, when SCIM is enabled, all users need to be provisioned through SCIM before a user can sign in through SAML. The matching of a user to a SCIM provisioned user is done when the SAML assertion is consumed. The user will be matched on SAML response `NameID` to SCIM `userName`.When converting existing enterprise to use SCIM, the user handle (`userName`) from the SCIM payload will be used to match the provisioned user to an already existing user in the enterprise. Since the new identity record is created for newly provisioned users the matching for those records is done using a user's handle. Currently the matching will be performed to all of the users no matter if they were SAML JIT provisioned or created as local users.
// returns a *RequestInformation when successful
func (m *V2EnterprisesItemUsersRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Userable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *V2EnterprisesItemUsersRequestBuilder when successful
func (m *V2EnterprisesItemUsersRequestBuilder) WithUrl(rawUrl string)(*V2EnterprisesItemUsersRequestBuilder) {
    return NewV2EnterprisesItemUsersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
