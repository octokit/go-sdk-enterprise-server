package gitignore

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// TemplatesWithNameItemRequestBuilder builds and executes requests for operations under \gitignore\templates\{name}
type TemplatesWithNameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewTemplatesWithNameItemRequestBuilderInternal instantiates a new TemplatesWithNameItemRequestBuilder and sets the default values.
func NewTemplatesWithNameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesWithNameItemRequestBuilder) {
    m := &TemplatesWithNameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/gitignore/templates/{name}", pathParameters),
    }
    return m
}
// NewTemplatesWithNameItemRequestBuilder instantiates a new TemplatesWithNameItemRequestBuilder and sets the default values.
func NewTemplatesWithNameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesWithNameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTemplatesWithNameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the content of a gitignore template.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.raw+json`**: Returns the raw .gitignore contents.
// returns a GitignoreTemplateable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.14/rest/gitignore/gitignore#get-a-gitignore-template
func (m *TemplatesWithNameItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.GitignoreTemplateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateGitignoreTemplateFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.GitignoreTemplateable), nil
}
// ToGetRequestInformation get the content of a gitignore template.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/enterprise-server@3.14/rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.raw+json`**: Returns the raw .gitignore contents.
// returns a *RequestInformation when successful
func (m *TemplatesWithNameItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TemplatesWithNameItemRequestBuilder when successful
func (m *TemplatesWithNameItemRequestBuilder) WithUrl(rawUrl string)(*TemplatesWithNameItemRequestBuilder) {
    return NewTemplatesWithNameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
