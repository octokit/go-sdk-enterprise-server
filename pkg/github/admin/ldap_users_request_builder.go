package admin

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LdapUsersRequestBuilder builds and executes requests for operations under \admin\ldap\users
type LdapUsersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByUsername gets an item from the github.com/octokit/go-sdk-enterprise-server/pkg/github.admin.ldap.users.item collection
// returns a *LdapUsersWithUsernameItemRequestBuilder when successful
func (m *LdapUsersRequestBuilder) ByUsername(username string)(*LdapUsersWithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if username != "" {
        urlTplParams["username"] = username
    }
    return NewLdapUsersWithUsernameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewLdapUsersRequestBuilderInternal instantiates a new LdapUsersRequestBuilder and sets the default values.
func NewLdapUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LdapUsersRequestBuilder) {
    m := &LdapUsersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/ldap/users", pathParameters),
    }
    return m
}
// NewLdapUsersRequestBuilder instantiates a new LdapUsersRequestBuilder and sets the default values.
func NewLdapUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LdapUsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLdapUsersRequestBuilderInternal(urlParams, requestAdapter)
}
