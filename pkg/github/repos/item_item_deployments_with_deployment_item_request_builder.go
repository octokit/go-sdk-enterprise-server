package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ItemItemDeploymentsWithDeployment_ItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\deployments\{deployment_id}
type ItemItemDeploymentsWithDeployment_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemDeploymentsWithDeployment_ItemRequestBuilderInternal instantiates a new ItemItemDeploymentsWithDeployment_ItemRequestBuilder and sets the default values.
func NewItemItemDeploymentsWithDeployment_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDeploymentsWithDeployment_ItemRequestBuilder) {
    m := &ItemItemDeploymentsWithDeployment_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/deployments/{deployment_id}", pathParameters),
    }
    return m
}
// NewItemItemDeploymentsWithDeployment_ItemRequestBuilder instantiates a new ItemItemDeploymentsWithDeployment_ItemRequestBuilder and sets the default values.
func NewItemItemDeploymentsWithDeployment_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDeploymentsWithDeployment_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemDeploymentsWithDeployment_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete if the repository only has one deployment, you can delete the deployment regardless of its status. If the repository has more than one deployment, you can only delete inactive deployments. This ensures that repositories with multiple deployments will always have an active deployment.To set a deployment as inactive, you must:*   Create a new deployment that is active so that the system has a record of the current state, then delete the previously active deployment.*   Mark the active deployment as inactive by adding any non-successful deployment status.For more information, see "[Create a deployment](https://docs.github.com/enterprise-server@3.13/rest/deployments/deployments/#create-a-deployment)" and "[Create a deployment status](https://docs.github.com/enterprise-server@3.13/rest/deployments/statuses#create-a-deployment-status)."OAuth app tokens and personal access tokens (classic) need the `repo` or `repo_deployment` scope to use this endpoint.
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationErrorSimple error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.13/rest/deployments/deployments#delete-a-deployment
func (m *ItemItemDeploymentsWithDeployment_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
        "422": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get a deployment
// returns a Deploymentable when successful
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.13/rest/deployments/deployments#get-a-deployment
func (m *ItemItemDeploymentsWithDeployment_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Deploymentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateDeploymentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Deploymentable), nil
}
// Statuses the statuses property
// returns a *ItemItemDeploymentsItemStatusesRequestBuilder when successful
func (m *ItemItemDeploymentsWithDeployment_ItemRequestBuilder) Statuses()(*ItemItemDeploymentsItemStatusesRequestBuilder) {
    return NewItemItemDeploymentsItemStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation if the repository only has one deployment, you can delete the deployment regardless of its status. If the repository has more than one deployment, you can only delete inactive deployments. This ensures that repositories with multiple deployments will always have an active deployment.To set a deployment as inactive, you must:*   Create a new deployment that is active so that the system has a record of the current state, then delete the previously active deployment.*   Mark the active deployment as inactive by adding any non-successful deployment status.For more information, see "[Create a deployment](https://docs.github.com/enterprise-server@3.13/rest/deployments/deployments/#create-a-deployment)" and "[Create a deployment status](https://docs.github.com/enterprise-server@3.13/rest/deployments/statuses#create-a-deployment-status)."OAuth app tokens and personal access tokens (classic) need the `repo` or `repo_deployment` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemDeploymentsWithDeployment_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// returns a *RequestInformation when successful
func (m *ItemItemDeploymentsWithDeployment_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemDeploymentsWithDeployment_ItemRequestBuilder when successful
func (m *ItemItemDeploymentsWithDeployment_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemDeploymentsWithDeployment_ItemRequestBuilder) {
    return NewItemItemDeploymentsWithDeployment_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
