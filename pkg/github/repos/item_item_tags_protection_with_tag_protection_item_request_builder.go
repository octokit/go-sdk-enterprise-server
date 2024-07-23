package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\tags\protection\{tag_protection_id}
type ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemTagsProtectionWithTag_protection_ItemRequestBuilderInternal instantiates a new ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder and sets the default values.
func NewItemItemTagsProtectionWithTag_protection_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder) {
    m := &ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/tags/protection/{tag_protection_id}", pathParameters),
    }
    return m
}
// NewItemItemTagsProtectionWithTag_protection_ItemRequestBuilder instantiates a new ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder and sets the default values.
func NewItemItemTagsProtectionWithTag_protection_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemTagsProtectionWithTag_protection_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete > [!WARNING]> **Deprecation notice:** This operation is deprecated and will be removed after August 30, 2024. Use the "[Repository Rulesets](https://docs.github.com/enterprise-server@3.12/rest/repos/rules#delete-a-repository-ruleset)" endpoint instead.This deletes a tag protection state for a repository.This endpoint is only available to repository administrators.
// Deprecated: 
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.12/rest/repos/tags#deprecated---delete-a-tag-protection-state-for-a-repository
func (m *ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
        "404": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation > [!WARNING]> **Deprecation notice:** This operation is deprecated and will be removed after August 30, 2024. Use the "[Repository Rulesets](https://docs.github.com/enterprise-server@3.12/rest/repos/rules#delete-a-repository-ruleset)" endpoint instead.This deletes a tag protection state for a repository.This endpoint is only available to repository administrators.
// Deprecated: 
// returns a *RequestInformation when successful
func (m *ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: 
// returns a *ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder when successful
func (m *ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder) {
    return NewItemItemTagsProtectionWithTag_protection_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
