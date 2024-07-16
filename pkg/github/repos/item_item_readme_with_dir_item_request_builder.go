package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ItemItemReadmeWithDirItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\readme\{dir}
type ItemItemReadmeWithDirItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemReadmeWithDirItemRequestBuilderGetQueryParameters gets the README from a repository directory.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/enterprise-server@3.10/rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.raw+json`**: Returns the raw file contents. This is the default if you do not specify a media type.- **`application/vnd.github.html+json`**: Returns the README in HTML. Markup languages are rendered to HTML using GitHub's open-source [Markup library](https://github.com/github/markup).
type ItemItemReadmeWithDirItemRequestBuilderGetQueryParameters struct {
    // The name of the commit/branch/tag. Default: the repository’s default branch.
    Ref *string `uriparametername:"ref"`
}
// NewItemItemReadmeWithDirItemRequestBuilderInternal instantiates a new ItemItemReadmeWithDirItemRequestBuilder and sets the default values.
func NewItemItemReadmeWithDirItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemReadmeWithDirItemRequestBuilder) {
    m := &ItemItemReadmeWithDirItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/readme/{dir}{?ref*}", pathParameters),
    }
    return m
}
// NewItemItemReadmeWithDirItemRequestBuilder instantiates a new ItemItemReadmeWithDirItemRequestBuilder and sets the default values.
func NewItemItemReadmeWithDirItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemReadmeWithDirItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemReadmeWithDirItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the README from a repository directory.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/enterprise-server@3.10/rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.raw+json`**: Returns the raw file contents. This is the default if you do not specify a media type.- **`application/vnd.github.html+json`**: Returns the README in HTML. Markup languages are rendered to HTML using GitHub's open-source [Markup library](https://github.com/github/markup).
// returns a ContentFileable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.10/rest/repos/contents#get-a-repository-readme-for-a-directory
func (m *ItemItemReadmeWithDirItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemReadmeWithDirItemRequestBuilderGetQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ContentFileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
        "422": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateContentFileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ContentFileable), nil
}
// ToGetRequestInformation gets the README from a repository directory.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/enterprise-server@3.10/rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.raw+json`**: Returns the raw file contents. This is the default if you do not specify a media type.- **`application/vnd.github.html+json`**: Returns the README in HTML. Markup languages are rendered to HTML using GitHub's open-source [Markup library](https://github.com/github/markup).
// returns a *RequestInformation when successful
func (m *ItemItemReadmeWithDirItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemReadmeWithDirItemRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemReadmeWithDirItemRequestBuilder when successful
func (m *ItemItemReadmeWithDirItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemReadmeWithDirItemRequestBuilder) {
    return NewItemItemReadmeWithDirItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
