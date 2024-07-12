package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\rulesets\rule-suites\{rule_suite_id}
type ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilderInternal instantiates a new ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder and sets the default values.
func NewItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) {
    m := &ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/rulesets/rule-suites/{rule_suite_id}", pathParameters),
    }
    return m
}
// NewItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder instantiates a new ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder and sets the default values.
func NewItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets information about a suite of rule evaluations from within a repository.For more information, see "[Managing rulesets for a repository](https://docs.github.com/enterprise-server@3.13/repositories/configuring-branches-and-merges-in-your-repository/managing-rulesets/managing-rulesets-for-a-repository#viewing-insights-for-rulesets)."
// returns a RuleSuiteable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.13/rest/repos/rule-suites#get-a-repository-rule-suite
func (m *ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RuleSuiteable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
        "500": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateRuleSuiteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RuleSuiteable), nil
}
// ToGetRequestInformation gets information about a suite of rule evaluations from within a repository.For more information, see "[Managing rulesets for a repository](https://docs.github.com/enterprise-server@3.13/repositories/configuring-branches-and-merges-in-your-repository/managing-rulesets/managing-rulesets-for-a-repository#viewing-insights-for-rulesets)."
// returns a *RequestInformation when successful
func (m *ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder when successful
func (m *ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder) {
    return NewItemItemRulesetsRuleSuitesWithRule_suite_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
