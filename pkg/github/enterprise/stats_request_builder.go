package enterprise

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// StatsRequestBuilder builds and executes requests for operations under \enterprise\stats
type StatsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// All the all property
// returns a *StatsAllRequestBuilder when successful
func (m *StatsRequestBuilder) All()(*StatsAllRequestBuilder) {
    return NewStatsAllRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Comments the comments property
// returns a *StatsCommentsRequestBuilder when successful
func (m *StatsRequestBuilder) Comments()(*StatsCommentsRequestBuilder) {
    return NewStatsCommentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewStatsRequestBuilderInternal instantiates a new StatsRequestBuilder and sets the default values.
func NewStatsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StatsRequestBuilder) {
    m := &StatsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprise/stats", pathParameters),
    }
    return m
}
// NewStatsRequestBuilder instantiates a new StatsRequestBuilder and sets the default values.
func NewStatsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StatsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStatsRequestBuilderInternal(urlParams, requestAdapter)
}
// Gists the gists property
// returns a *StatsGistsRequestBuilder when successful
func (m *StatsRequestBuilder) Gists()(*StatsGistsRequestBuilder) {
    return NewStatsGistsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Hooks the hooks property
// returns a *StatsHooksRequestBuilder when successful
func (m *StatsRequestBuilder) Hooks()(*StatsHooksRequestBuilder) {
    return NewStatsHooksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Issues the issues property
// returns a *StatsIssuesRequestBuilder when successful
func (m *StatsRequestBuilder) Issues()(*StatsIssuesRequestBuilder) {
    return NewStatsIssuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Milestones the milestones property
// returns a *StatsMilestonesRequestBuilder when successful
func (m *StatsRequestBuilder) Milestones()(*StatsMilestonesRequestBuilder) {
    return NewStatsMilestonesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Orgs the orgs property
// returns a *StatsOrgsRequestBuilder when successful
func (m *StatsRequestBuilder) Orgs()(*StatsOrgsRequestBuilder) {
    return NewStatsOrgsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Pages the pages property
// returns a *StatsPagesRequestBuilder when successful
func (m *StatsRequestBuilder) Pages()(*StatsPagesRequestBuilder) {
    return NewStatsPagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Pulls the pulls property
// returns a *StatsPullsRequestBuilder when successful
func (m *StatsRequestBuilder) Pulls()(*StatsPullsRequestBuilder) {
    return NewStatsPullsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repos the repos property
// returns a *StatsReposRequestBuilder when successful
func (m *StatsRequestBuilder) Repos()(*StatsReposRequestBuilder) {
    return NewStatsReposRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecurityProducts the securityProducts property
// returns a *StatsSecurityProductsRequestBuilder when successful
func (m *StatsRequestBuilder) SecurityProducts()(*StatsSecurityProductsRequestBuilder) {
    return NewStatsSecurityProductsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users the users property
// returns a *StatsUsersRequestBuilder when successful
func (m *StatsRequestBuilder) Users()(*StatsUsersRequestBuilder) {
    return NewStatsUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
