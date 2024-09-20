package admin

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LdapTeamsWithTeam_ItemRequestBuilder builds and executes requests for operations under \admin\ldap\teams\{team_id}
type LdapTeamsWithTeam_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewLdapTeamsWithTeam_ItemRequestBuilderInternal instantiates a new LdapTeamsWithTeam_ItemRequestBuilder and sets the default values.
func NewLdapTeamsWithTeam_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LdapTeamsWithTeam_ItemRequestBuilder) {
    m := &LdapTeamsWithTeam_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/ldap/teams/{team_id}", pathParameters),
    }
    return m
}
// NewLdapTeamsWithTeam_ItemRequestBuilder instantiates a new LdapTeamsWithTeam_ItemRequestBuilder and sets the default values.
func NewLdapTeamsWithTeam_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LdapTeamsWithTeam_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLdapTeamsWithTeam_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Mapping the mapping property
// returns a *LdapTeamsItemMappingRequestBuilder when successful
func (m *LdapTeamsWithTeam_ItemRequestBuilder) Mapping()(*LdapTeamsItemMappingRequestBuilder) {
    return NewLdapTeamsItemMappingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sync the sync property
// returns a *LdapTeamsItemSyncRequestBuilder when successful
func (m *LdapTeamsWithTeam_ItemRequestBuilder) Sync()(*LdapTeamsItemSyncRequestBuilder) {
    return NewLdapTeamsItemSyncRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
