package admin

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LdapRequestBuilder builds and executes requests for operations under \admin\ldap
type LdapRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewLdapRequestBuilderInternal instantiates a new LdapRequestBuilder and sets the default values.
func NewLdapRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LdapRequestBuilder) {
    m := &LdapRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/ldap", pathParameters),
    }
    return m
}
// NewLdapRequestBuilder instantiates a new LdapRequestBuilder and sets the default values.
func NewLdapRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LdapRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLdapRequestBuilderInternal(urlParams, requestAdapter)
}
// Teams the teams property
// returns a *LdapTeamsRequestBuilder when successful
func (m *LdapRequestBuilder) Teams()(*LdapTeamsRequestBuilder) {
    return NewLdapTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users the users property
// returns a *LdapUsersRequestBuilder when successful
func (m *LdapRequestBuilder) Users()(*LdapUsersRequestBuilder) {
    return NewLdapUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
