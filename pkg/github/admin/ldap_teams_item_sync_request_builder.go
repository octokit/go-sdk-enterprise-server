package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LdapTeamsItemSyncRequestBuilder builds and executes requests for operations under \admin\ldap\teams\{team_id}\sync
type LdapTeamsItemSyncRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewLdapTeamsItemSyncRequestBuilderInternal instantiates a new LdapTeamsItemSyncRequestBuilder and sets the default values.
func NewLdapTeamsItemSyncRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LdapTeamsItemSyncRequestBuilder) {
    m := &LdapTeamsItemSyncRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/ldap/teams/{team_id}/sync", pathParameters),
    }
    return m
}
// NewLdapTeamsItemSyncRequestBuilder instantiates a new LdapTeamsItemSyncRequestBuilder and sets the default values.
func NewLdapTeamsItemSyncRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LdapTeamsItemSyncRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLdapTeamsItemSyncRequestBuilderInternal(urlParams, requestAdapter)
}
// Post note that this API call does not automatically initiate an LDAP sync. Rather, if a `201` is returned, the sync job is queued successfully, and is performed when the instance is ready.
// returns a LdapTeamsItemSyncPostResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.11/rest/enterprise-admin/ldap#sync-ldap-mapping-for-a-team
func (m *LdapTeamsItemSyncRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(LdapTeamsItemSyncPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateLdapTeamsItemSyncPostResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(LdapTeamsItemSyncPostResponseable), nil
}
// ToPostRequestInformation note that this API call does not automatically initiate an LDAP sync. Rather, if a `201` is returned, the sync job is queued successfully, and is performed when the instance is ready.
// returns a *RequestInformation when successful
func (m *LdapTeamsItemSyncRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *LdapTeamsItemSyncRequestBuilder when successful
func (m *LdapTeamsItemSyncRequestBuilder) WithUrl(rawUrl string)(*LdapTeamsItemSyncRequestBuilder) {
    return NewLdapTeamsItemSyncRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
