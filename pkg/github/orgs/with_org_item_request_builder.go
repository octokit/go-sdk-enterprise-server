package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// WithOrgItemRequestBuilder builds and executes requests for operations under \orgs\{org}
type WithOrgItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Actions the actions property
// returns a *ItemActionsRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) Actions()(*ItemActionsRequestBuilder) {
    return NewItemActionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Announcement the announcement property
// returns a *ItemAnnouncementRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) Announcement()(*ItemAnnouncementRequestBuilder) {
    return NewItemAnnouncementRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AuditLog the auditLog property
// returns a *ItemAuditLogRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) AuditLog()(*ItemAuditLogRequestBuilder) {
    return NewItemAuditLogRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BySecurity_product gets an item from the github.com/octokit/go-sdk-enterprise-server/pkg/github.orgs.item.item collection
// returns a *ItemWithSecurity_productItemRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) BySecurity_product(security_product string)(*ItemWithSecurity_productItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if security_product != "" {
        urlTplParams["security_product"] = security_product
    }
    return NewItemWithSecurity_productItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// CodeScanning the codeScanning property
// returns a *ItemCodeScanningRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) CodeScanning()(*ItemCodeScanningRequestBuilder) {
    return NewItemCodeScanningRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWithOrgItemRequestBuilderInternal instantiates a new WithOrgItemRequestBuilder and sets the default values.
func NewWithOrgItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithOrgItemRequestBuilder) {
    m := &WithOrgItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}", pathParameters),
    }
    return m
}
// NewWithOrgItemRequestBuilder instantiates a new WithOrgItemRequestBuilder and sets the default values.
func NewWithOrgItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithOrgItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithOrgItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CustomRepositoryRoles the customRepositoryRoles property
// returns a *ItemCustomRepositoryRolesRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) CustomRepositoryRoles()(*ItemCustomRepositoryRolesRequestBuilder) {
    return NewItemCustomRepositoryRolesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete deletes an organization and all its repositories.The organization login will be unavailable for 90 days after deletion.Please review the Terms of Service regarding account deletion before using this endpoint:https://docs.github.com/enterprise-server@3.10/site-policy/github-terms/github-terms-of-service
// returns a ItemWithOrgDeleteResponseable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.10/rest/orgs/orgs#delete-an-organization
func (m *WithOrgItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ItemWithOrgDeleteResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
        "404": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemWithOrgDeleteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemWithOrgDeleteResponseable), nil
}
// Dependabot the dependabot property
// returns a *ItemDependabotRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) Dependabot()(*ItemDependabotRequestBuilder) {
    return NewItemDependabotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Docker the docker property
// returns a *ItemDockerRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) Docker()(*ItemDockerRequestBuilder) {
    return NewItemDockerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Events the events property
// returns a *ItemEventsRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) Events()(*ItemEventsRequestBuilder) {
    return NewItemEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExternalGroup the externalGroup property
// returns a *ItemExternalGroupRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) ExternalGroup()(*ItemExternalGroupRequestBuilder) {
    return NewItemExternalGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExternalGroups the externalGroups property
// returns a *ItemExternalGroupsRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) ExternalGroups()(*ItemExternalGroupsRequestBuilder) {
    return NewItemExternalGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get gets information about an organization.When the value of `two_factor_requirement_enabled` is `true`, the organization requires all members, billing managers, and outside collaborators to enable [two-factor authentication](https://docs.github.com/enterprise-server@3.10/articles/securing-your-account-with-two-factor-authentication-2fa/).To see the full details about an organization, the authenticated user must be an organization owner.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to see the full details about an organization.To see information about an organization's GitHub Enterprise Server plan, GitHub Apps need the `Organization plan` permission.
// returns a OrganizationFullable when successful
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.10/rest/orgs/orgs#get-an-organization
func (m *WithOrgItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.OrganizationFullable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateOrganizationFullFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.OrganizationFullable), nil
}
// Hooks the hooks property
// returns a *ItemHooksRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) Hooks()(*ItemHooksRequestBuilder) {
    return NewItemHooksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Installation the installation property
// returns a *ItemInstallationRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) Installation()(*ItemInstallationRequestBuilder) {
    return NewItemInstallationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Installations the installations property
// returns a *ItemInstallationsRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) Installations()(*ItemInstallationsRequestBuilder) {
    return NewItemInstallationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Issues the issues property
// returns a *ItemIssuesRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) Issues()(*ItemIssuesRequestBuilder) {
    return NewItemIssuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Members the members property
// returns a *ItemMembersRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) Members()(*ItemMembersRequestBuilder) {
    return NewItemMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Memberships the memberships property
// returns a *ItemMembershipsRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) Memberships()(*ItemMembershipsRequestBuilder) {
    return NewItemMembershipsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Migrations the migrations property
// returns a *ItemMigrationsRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) Migrations()(*ItemMigrationsRequestBuilder) {
    return NewItemMigrationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Outside_collaborators the outside_collaborators property
// returns a *ItemOutside_collaboratorsRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) Outside_collaborators()(*ItemOutside_collaboratorsRequestBuilder) {
    return NewItemOutside_collaboratorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Packages the packages property
// returns a *ItemPackagesRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) Packages()(*ItemPackagesRequestBuilder) {
    return NewItemPackagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch **Parameter Deprecation Notice:** GitHub Enterprise Server will replace and discontinue `members_allowed_repository_creation_type` in favor of more granular permissions. The new input parameters are `members_can_create_public_repositories`, `members_can_create_private_repositories` for all organizations and `members_can_create_internal_repositories` for organizations associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+. For more information, see the [blog post](https://developer.github.com/changes/2019-12-03-internal-visibility-changes).Updates the organization's profile and member privileges.The authenticated user must be an organization owner to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `admin:org` or `repo` scope to use this endpoint.
// returns a OrganizationFullable when successful
// returns a BasicError error when the service returns a 409 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.10/rest/orgs/orgs#update-an-organization
func (m *WithOrgItemRequestBuilder) Patch(ctx context.Context, body ItemWithOrgPatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.OrganizationFullable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "409": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateOrganizationFullFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.OrganizationFullable), nil
}
// PersonalAccessTokenRequests the personalAccessTokenRequests property
// returns a *ItemPersonalAccessTokenRequestsRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) PersonalAccessTokenRequests()(*ItemPersonalAccessTokenRequestsRequestBuilder) {
    return NewItemPersonalAccessTokenRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PersonalAccessTokens the personalAccessTokens property
// returns a *ItemPersonalAccessTokensRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) PersonalAccessTokens()(*ItemPersonalAccessTokensRequestBuilder) {
    return NewItemPersonalAccessTokensRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PreReceiveHooks the preReceiveHooks property
// returns a *ItemPreReceiveHooksRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) PreReceiveHooks()(*ItemPreReceiveHooksRequestBuilder) {
    return NewItemPreReceiveHooksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Projects the projects property
// returns a *ItemProjectsRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) Projects()(*ItemProjectsRequestBuilder) {
    return NewItemProjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Public_members the public_members property
// returns a *ItemPublic_membersRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) Public_members()(*ItemPublic_membersRequestBuilder) {
    return NewItemPublic_membersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repos the repos property
// returns a *ItemReposRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) Repos()(*ItemReposRequestBuilder) {
    return NewItemReposRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RepositoryFineGrainedPermissions the repositoryFineGrainedPermissions property
// returns a *ItemRepositoryFineGrainedPermissionsRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) RepositoryFineGrainedPermissions()(*ItemRepositoryFineGrainedPermissionsRequestBuilder) {
    return NewItemRepositoryFineGrainedPermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecretScanning the secretScanning property
// returns a *ItemSecretScanningRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) SecretScanning()(*ItemSecretScanningRequestBuilder) {
    return NewItemSecretScanningRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecurityManagers the securityManagers property
// returns a *ItemSecurityManagersRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) SecurityManagers()(*ItemSecurityManagersRequestBuilder) {
    return NewItemSecurityManagersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Settings the settings property
// returns a *ItemSettingsRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) Settings()(*ItemSettingsRequestBuilder) {
    return NewItemSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Teams the teams property
// returns a *ItemTeamsRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) Teams()(*ItemTeamsRequestBuilder) {
    return NewItemTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deletes an organization and all its repositories.The organization login will be unavailable for 90 days after deletion.Please review the Terms of Service regarding account deletion before using this endpoint:https://docs.github.com/enterprise-server@3.10/site-policy/github-terms/github-terms-of-service
// returns a *RequestInformation when successful
func (m *WithOrgItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation gets information about an organization.When the value of `two_factor_requirement_enabled` is `true`, the organization requires all members, billing managers, and outside collaborators to enable [two-factor authentication](https://docs.github.com/enterprise-server@3.10/articles/securing-your-account-with-two-factor-authentication-2fa/).To see the full details about an organization, the authenticated user must be an organization owner.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to see the full details about an organization.To see information about an organization's GitHub Enterprise Server plan, GitHub Apps need the `Organization plan` permission.
// returns a *RequestInformation when successful
func (m *WithOrgItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation **Parameter Deprecation Notice:** GitHub Enterprise Server will replace and discontinue `members_allowed_repository_creation_type` in favor of more granular permissions. The new input parameters are `members_can_create_public_repositories`, `members_can_create_private_repositories` for all organizations and `members_can_create_internal_repositories` for organizations associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+. For more information, see the [blog post](https://developer.github.com/changes/2019-12-03-internal-visibility-changes).Updates the organization's profile and member privileges.The authenticated user must be an organization owner to use this endpoint.OAuth app tokens and personal access tokens (classic) need the `admin:org` or `repo` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *WithOrgItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemWithOrgPatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WithOrgItemRequestBuilder when successful
func (m *WithOrgItemRequestBuilder) WithUrl(rawUrl string)(*WithOrgItemRequestBuilder) {
    return NewWithOrgItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
