package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ItemItemCodeownersErrorsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\codeowners\errors
type ItemItemCodeownersErrorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCodeownersErrorsRequestBuilderGetQueryParameters list any syntax errors that are detected in the CODEOWNERSfile.For more information about the correct CODEOWNERS syntax,see "[About code owners](https://docs.github.com/enterprise-server@3.13/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/about-code-owners)."
type ItemItemCodeownersErrorsRequestBuilderGetQueryParameters struct {
    // A branch, tag or commit name used to determine which version of the CODEOWNERS file to use. Default: the repository's default branch (e.g. `main`)
    Ref *string `uriparametername:"ref"`
}
// NewItemItemCodeownersErrorsRequestBuilderInternal instantiates a new ItemItemCodeownersErrorsRequestBuilder and sets the default values.
func NewItemItemCodeownersErrorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeownersErrorsRequestBuilder) {
    m := &ItemItemCodeownersErrorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/codeowners/errors{?ref*}", pathParameters),
    }
    return m
}
// NewItemItemCodeownersErrorsRequestBuilder instantiates a new ItemItemCodeownersErrorsRequestBuilder and sets the default values.
func NewItemItemCodeownersErrorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeownersErrorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeownersErrorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list any syntax errors that are detected in the CODEOWNERSfile.For more information about the correct CODEOWNERS syntax,see "[About code owners](https://docs.github.com/enterprise-server@3.13/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/about-code-owners)."
// returns a CodeownersErrorsable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.13/rest/repos/repos#list-codeowners-errors
func (m *ItemItemCodeownersErrorsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemCodeownersErrorsRequestBuilderGetQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CodeownersErrorsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateCodeownersErrorsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CodeownersErrorsable), nil
}
// ToGetRequestInformation list any syntax errors that are detected in the CODEOWNERSfile.For more information about the correct CODEOWNERS syntax,see "[About code owners](https://docs.github.com/enterprise-server@3.13/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/about-code-owners)."
// returns a *RequestInformation when successful
func (m *ItemItemCodeownersErrorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemCodeownersErrorsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemCodeownersErrorsRequestBuilder when successful
func (m *ItemItemCodeownersErrorsRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodeownersErrorsRequestBuilder) {
    return NewItemItemCodeownersErrorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
