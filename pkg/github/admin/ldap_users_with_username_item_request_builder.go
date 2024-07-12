package admin

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LdapUsersWithUsernameItemRequestBuilder builds and executes requests for operations under \admin\ldap\users\{username}
type LdapUsersWithUsernameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewLdapUsersWithUsernameItemRequestBuilderInternal instantiates a new LdapUsersWithUsernameItemRequestBuilder and sets the default values.
func NewLdapUsersWithUsernameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LdapUsersWithUsernameItemRequestBuilder) {
    m := &LdapUsersWithUsernameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/ldap/users/{username}", pathParameters),
    }
    return m
}
// NewLdapUsersWithUsernameItemRequestBuilder instantiates a new LdapUsersWithUsernameItemRequestBuilder and sets the default values.
func NewLdapUsersWithUsernameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LdapUsersWithUsernameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLdapUsersWithUsernameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Mapping the mapping property
// returns a *LdapUsersItemMappingRequestBuilder when successful
func (m *LdapUsersWithUsernameItemRequestBuilder) Mapping()(*LdapUsersItemMappingRequestBuilder) {
    return NewLdapUsersItemMappingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sync the sync property
// returns a *LdapUsersItemSyncRequestBuilder when successful
func (m *LdapUsersWithUsernameItemRequestBuilder) Sync()(*LdapUsersItemSyncRequestBuilder) {
    return NewLdapUsersItemSyncRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
