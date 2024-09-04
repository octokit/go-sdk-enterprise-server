package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

// ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\actions\runs\{run_id}\deployment_protection_rule
type ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Deployment_protection_rulePostRequestBody composed type wrapper for classes ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesCommentRequiredable, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesStateRequiredable
type Deployment_protection_rulePostRequestBody struct {
    // Composed type representation for type ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesCommentRequiredable
    deployment_protection_rulePostRequestBodyReviewCustomGatesCommentRequired ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesCommentRequiredable
    // Composed type representation for type ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesStateRequiredable
    deployment_protection_rulePostRequestBodyReviewCustomGatesStateRequired ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesStateRequiredable
    // Composed type representation for type ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesCommentRequiredable
    reviewCustomGatesCommentRequired ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesCommentRequiredable
    // Composed type representation for type ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesStateRequiredable
    reviewCustomGatesStateRequired ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesStateRequiredable
}
// NewDeployment_protection_rulePostRequestBody instantiates a new Deployment_protection_rulePostRequestBody and sets the default values.
func NewDeployment_protection_rulePostRequestBody()(*Deployment_protection_rulePostRequestBody) {
    m := &Deployment_protection_rulePostRequestBody{
    }
    return m
}
// CreateDeployment_protection_rulePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeployment_protection_rulePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewDeployment_protection_rulePostRequestBody()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateReviewCustomGatesCommentRequiredFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesCommentRequiredable); ok {
                result.SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired(cast)
            }
        } else if val, err := parseNode.GetObjectValue(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateReviewCustomGatesStateRequiredFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesStateRequiredable); ok {
                result.SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired(cast)
            }
        } else if val, err := parseNode.GetObjectValue(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateReviewCustomGatesCommentRequiredFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesCommentRequiredable); ok {
                result.SetReviewCustomGatesCommentRequired(cast)
            }
        } else if val, err := parseNode.GetObjectValue(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateReviewCustomGatesStateRequiredFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesStateRequiredable); ok {
                result.SetReviewCustomGatesStateRequired(cast)
            }
        }
    }
    return result, nil
}
// GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired gets the reviewCustomGatesCommentRequired property value. Composed type representation for type ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesCommentRequiredable
// returns a ReviewCustomGatesCommentRequiredable when successful
func (m *Deployment_protection_rulePostRequestBody) GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired()(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesCommentRequiredable) {
    return m.deployment_protection_rulePostRequestBodyReviewCustomGatesCommentRequired
}
// GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired gets the reviewCustomGatesStateRequired property value. Composed type representation for type ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesStateRequiredable
// returns a ReviewCustomGatesStateRequiredable when successful
func (m *Deployment_protection_rulePostRequestBody) GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired()(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesStateRequiredable) {
    return m.deployment_protection_rulePostRequestBodyReviewCustomGatesStateRequired
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Deployment_protection_rulePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *Deployment_protection_rulePostRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetReviewCustomGatesCommentRequired gets the reviewCustomGatesCommentRequired property value. Composed type representation for type ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesCommentRequiredable
// returns a ReviewCustomGatesCommentRequiredable when successful
func (m *Deployment_protection_rulePostRequestBody) GetReviewCustomGatesCommentRequired()(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesCommentRequiredable) {
    return m.reviewCustomGatesCommentRequired
}
// GetReviewCustomGatesStateRequired gets the reviewCustomGatesStateRequired property value. Composed type representation for type ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesStateRequiredable
// returns a ReviewCustomGatesStateRequiredable when successful
func (m *Deployment_protection_rulePostRequestBody) GetReviewCustomGatesStateRequired()(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesStateRequiredable) {
    return m.reviewCustomGatesStateRequired
}
// Serialize serializes information the current object
func (m *Deployment_protection_rulePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired() != nil {
        err := writer.WriteObjectValue("", m.GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired())
        if err != nil {
            return err
        }
    } else if m.GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired() != nil {
        err := writer.WriteObjectValue("", m.GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired())
        if err != nil {
            return err
        }
    } else if m.GetReviewCustomGatesCommentRequired() != nil {
        err := writer.WriteObjectValue("", m.GetReviewCustomGatesCommentRequired())
        if err != nil {
            return err
        }
    } else if m.GetReviewCustomGatesStateRequired() != nil {
        err := writer.WriteObjectValue("", m.GetReviewCustomGatesStateRequired())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired sets the reviewCustomGatesCommentRequired property value. Composed type representation for type ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesCommentRequiredable
func (m *Deployment_protection_rulePostRequestBody) SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired(value ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesCommentRequiredable)() {
    m.deployment_protection_rulePostRequestBodyReviewCustomGatesCommentRequired = value
}
// SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired sets the reviewCustomGatesStateRequired property value. Composed type representation for type ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesStateRequiredable
func (m *Deployment_protection_rulePostRequestBody) SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired(value ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesStateRequiredable)() {
    m.deployment_protection_rulePostRequestBodyReviewCustomGatesStateRequired = value
}
// SetReviewCustomGatesCommentRequired sets the reviewCustomGatesCommentRequired property value. Composed type representation for type ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesCommentRequiredable
func (m *Deployment_protection_rulePostRequestBody) SetReviewCustomGatesCommentRequired(value ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesCommentRequiredable)() {
    m.reviewCustomGatesCommentRequired = value
}
// SetReviewCustomGatesStateRequired sets the reviewCustomGatesStateRequired property value. Composed type representation for type ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesStateRequiredable
func (m *Deployment_protection_rulePostRequestBody) SetReviewCustomGatesStateRequired(value ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesStateRequiredable)() {
    m.reviewCustomGatesStateRequired = value
}
type Deployment_protection_rulePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired()(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesCommentRequiredable)
    GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired()(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesStateRequiredable)
    GetReviewCustomGatesCommentRequired()(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesCommentRequiredable)
    GetReviewCustomGatesStateRequired()(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesStateRequiredable)
    SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired(value ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesCommentRequiredable)()
    SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired(value ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesStateRequiredable)()
    SetReviewCustomGatesCommentRequired(value ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesCommentRequiredable)()
    SetReviewCustomGatesStateRequired(value ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ReviewCustomGatesStateRequiredable)()
}
// NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilderInternal instantiates a new ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) {
    m := &ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/actions/runs/{run_id}/deployment_protection_rule", pathParameters),
    }
    return m
}
// NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder instantiates a new ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilderInternal(urlParams, requestAdapter)
}
// Post approve or reject custom deployment protection rules provided by a GitHub App for a workflow run. For more information, see "[Using environments for deployment](https://docs.github.com/enterprise-server@3.14/actions/deployment/targeting-different-environments/using-environments-for-deployment)."> [!NOTE]> GitHub Apps can only review their own custom deployment protection rules. To approve or reject pending deployments that are waiting for review from a specific person or team, see [`POST /repos/{owner}/{repo}/actions/runs/{run_id}/pending_deployments`](/rest/actions/workflow-runs#review-pending-deployments-for-a-workflow-run).OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.14/rest/actions/workflow-runs#review-custom-deployment-protection-rules-for-a-workflow-run
func (m *ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) Post(ctx context.Context, body Deployment_protection_rulePostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation approve or reject custom deployment protection rules provided by a GitHub App for a workflow run. For more information, see "[Using environments for deployment](https://docs.github.com/enterprise-server@3.14/actions/deployment/targeting-different-environments/using-environments-for-deployment)."> [!NOTE]> GitHub Apps can only review their own custom deployment protection rules. To approve or reject pending deployments that are waiting for review from a specific person or team, see [`POST /repos/{owner}/{repo}/actions/runs/{run_id}/pending_deployments`](/rest/actions/workflow-runs#review-pending-deployments-for-a-workflow-run).OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with a private repository.
// returns a *RequestInformation when successful
func (m *ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) ToPostRequestInformation(ctx context.Context, body Deployment_protection_rulePostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder when successful
func (m *ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) {
    return NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
