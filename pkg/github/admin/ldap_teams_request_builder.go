package admin

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LdapTeamsRequestBuilder builds and executes requests for operations under \admin\ldap\teams
type LdapTeamsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByTeam_id gets an item from the github.com/octokit/go-sdk-enterprise-server/pkg/github.admin.ldap.teams.item collection
// returns a *LdapTeamsWithTeam_ItemRequestBuilder when successful
func (m *LdapTeamsRequestBuilder) ByTeam_id(team_id int32)(*LdapTeamsWithTeam_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["team_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(team_id), 10)
    return NewLdapTeamsWithTeam_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewLdapTeamsRequestBuilderInternal instantiates a new LdapTeamsRequestBuilder and sets the default values.
func NewLdapTeamsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LdapTeamsRequestBuilder) {
    m := &LdapTeamsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/ldap/teams", pathParameters),
    }
    return m
}
// NewLdapTeamsRequestBuilder instantiates a new LdapTeamsRequestBuilder and sets the default values.
func NewLdapTeamsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LdapTeamsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLdapTeamsRequestBuilderInternal(urlParams, requestAdapter)
}
