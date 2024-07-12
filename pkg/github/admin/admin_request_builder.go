package admin

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AdminRequestBuilder builds and executes requests for operations under \admin
type AdminRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewAdminRequestBuilderInternal instantiates a new AdminRequestBuilder and sets the default values.
func NewAdminRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminRequestBuilder) {
    m := &AdminRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin", pathParameters),
    }
    return m
}
// NewAdminRequestBuilder instantiates a new AdminRequestBuilder and sets the default values.
func NewAdminRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdminRequestBuilderInternal(urlParams, requestAdapter)
}
// Hooks the hooks property
// returns a *HooksRequestBuilder when successful
func (m *AdminRequestBuilder) Hooks()(*HooksRequestBuilder) {
    return NewHooksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Keys the keys property
// returns a *KeysRequestBuilder when successful
func (m *AdminRequestBuilder) Keys()(*KeysRequestBuilder) {
    return NewKeysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ldap the ldap property
// returns a *LdapRequestBuilder when successful
func (m *AdminRequestBuilder) Ldap()(*LdapRequestBuilder) {
    return NewLdapRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Organizations the organizations property
// returns a *OrganizationsRequestBuilder when successful
func (m *AdminRequestBuilder) Organizations()(*OrganizationsRequestBuilder) {
    return NewOrganizationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PreReceiveEnvironments the preReceiveEnvironments property
// returns a *PreReceiveEnvironmentsRequestBuilder when successful
func (m *AdminRequestBuilder) PreReceiveEnvironments()(*PreReceiveEnvironmentsRequestBuilder) {
    return NewPreReceiveEnvironmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PreReceiveHooks the preReceiveHooks property
// returns a *PreReceiveHooksRequestBuilder when successful
func (m *AdminRequestBuilder) PreReceiveHooks()(*PreReceiveHooksRequestBuilder) {
    return NewPreReceiveHooksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tokens the tokens property
// returns a *TokensRequestBuilder when successful
func (m *AdminRequestBuilder) Tokens()(*TokensRequestBuilder) {
    return NewTokensRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users the users property
// returns a *UsersRequestBuilder when successful
func (m *AdminRequestBuilder) Users()(*UsersRequestBuilder) {
    return NewUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
