package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EnterpriseOverview struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The comments property
    comments EnterpriseCommentOverviewable
    // The gists property
    gists EnterpriseGistOverviewable
    // The hooks property
    hooks EnterpriseHookOverviewable
    // The issues property
    issues EnterpriseIssueOverviewable
    // The milestones property
    milestones EnterpriseMilestoneOverviewable
    // The orgs property
    orgs EnterpriseOrganizationOverviewable
    // The pages property
    pages EnterprisePageOverviewable
    // The pulls property
    pulls EnterprisePullRequestOverviewable
    // The repos property
    repos EnterpriseRepositoryOverviewable
    // The users property
    users EnterpriseUserOverviewable
}
// NewEnterpriseOverview instantiates a new EnterpriseOverview and sets the default values.
func NewEnterpriseOverview()(*EnterpriseOverview) {
    m := &EnterpriseOverview{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnterpriseOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnterpriseOverviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnterpriseOverview(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EnterpriseOverview) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComments gets the comments property value. The comments property
// returns a EnterpriseCommentOverviewable when successful
func (m *EnterpriseOverview) GetComments()(EnterpriseCommentOverviewable) {
    return m.comments
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnterpriseOverview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnterpriseCommentOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComments(val.(EnterpriseCommentOverviewable))
        }
        return nil
    }
    res["gists"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnterpriseGistOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGists(val.(EnterpriseGistOverviewable))
        }
        return nil
    }
    res["hooks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnterpriseHookOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHooks(val.(EnterpriseHookOverviewable))
        }
        return nil
    }
    res["issues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnterpriseIssueOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssues(val.(EnterpriseIssueOverviewable))
        }
        return nil
    }
    res["milestones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnterpriseMilestoneOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMilestones(val.(EnterpriseMilestoneOverviewable))
        }
        return nil
    }
    res["orgs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnterpriseOrganizationOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrgs(val.(EnterpriseOrganizationOverviewable))
        }
        return nil
    }
    res["pages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnterprisePageOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPages(val.(EnterprisePageOverviewable))
        }
        return nil
    }
    res["pulls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnterprisePullRequestOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPulls(val.(EnterprisePullRequestOverviewable))
        }
        return nil
    }
    res["repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnterpriseRepositoryOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepos(val.(EnterpriseRepositoryOverviewable))
        }
        return nil
    }
    res["users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnterpriseUserOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsers(val.(EnterpriseUserOverviewable))
        }
        return nil
    }
    return res
}
// GetGists gets the gists property value. The gists property
// returns a EnterpriseGistOverviewable when successful
func (m *EnterpriseOverview) GetGists()(EnterpriseGistOverviewable) {
    return m.gists
}
// GetHooks gets the hooks property value. The hooks property
// returns a EnterpriseHookOverviewable when successful
func (m *EnterpriseOverview) GetHooks()(EnterpriseHookOverviewable) {
    return m.hooks
}
// GetIssues gets the issues property value. The issues property
// returns a EnterpriseIssueOverviewable when successful
func (m *EnterpriseOverview) GetIssues()(EnterpriseIssueOverviewable) {
    return m.issues
}
// GetMilestones gets the milestones property value. The milestones property
// returns a EnterpriseMilestoneOverviewable when successful
func (m *EnterpriseOverview) GetMilestones()(EnterpriseMilestoneOverviewable) {
    return m.milestones
}
// GetOrgs gets the orgs property value. The orgs property
// returns a EnterpriseOrganizationOverviewable when successful
func (m *EnterpriseOverview) GetOrgs()(EnterpriseOrganizationOverviewable) {
    return m.orgs
}
// GetPages gets the pages property value. The pages property
// returns a EnterprisePageOverviewable when successful
func (m *EnterpriseOverview) GetPages()(EnterprisePageOverviewable) {
    return m.pages
}
// GetPulls gets the pulls property value. The pulls property
// returns a EnterprisePullRequestOverviewable when successful
func (m *EnterpriseOverview) GetPulls()(EnterprisePullRequestOverviewable) {
    return m.pulls
}
// GetRepos gets the repos property value. The repos property
// returns a EnterpriseRepositoryOverviewable when successful
func (m *EnterpriseOverview) GetRepos()(EnterpriseRepositoryOverviewable) {
    return m.repos
}
// GetUsers gets the users property value. The users property
// returns a EnterpriseUserOverviewable when successful
func (m *EnterpriseOverview) GetUsers()(EnterpriseUserOverviewable) {
    return m.users
}
// Serialize serializes information the current object
func (m *EnterpriseOverview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("comments", m.GetComments())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("gists", m.GetGists())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("hooks", m.GetHooks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("issues", m.GetIssues())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("milestones", m.GetMilestones())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("orgs", m.GetOrgs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("pages", m.GetPages())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("pulls", m.GetPulls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("repos", m.GetRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("users", m.GetUsers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EnterpriseOverview) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComments sets the comments property value. The comments property
func (m *EnterpriseOverview) SetComments(value EnterpriseCommentOverviewable)() {
    m.comments = value
}
// SetGists sets the gists property value. The gists property
func (m *EnterpriseOverview) SetGists(value EnterpriseGistOverviewable)() {
    m.gists = value
}
// SetHooks sets the hooks property value. The hooks property
func (m *EnterpriseOverview) SetHooks(value EnterpriseHookOverviewable)() {
    m.hooks = value
}
// SetIssues sets the issues property value. The issues property
func (m *EnterpriseOverview) SetIssues(value EnterpriseIssueOverviewable)() {
    m.issues = value
}
// SetMilestones sets the milestones property value. The milestones property
func (m *EnterpriseOverview) SetMilestones(value EnterpriseMilestoneOverviewable)() {
    m.milestones = value
}
// SetOrgs sets the orgs property value. The orgs property
func (m *EnterpriseOverview) SetOrgs(value EnterpriseOrganizationOverviewable)() {
    m.orgs = value
}
// SetPages sets the pages property value. The pages property
func (m *EnterpriseOverview) SetPages(value EnterprisePageOverviewable)() {
    m.pages = value
}
// SetPulls sets the pulls property value. The pulls property
func (m *EnterpriseOverview) SetPulls(value EnterprisePullRequestOverviewable)() {
    m.pulls = value
}
// SetRepos sets the repos property value. The repos property
func (m *EnterpriseOverview) SetRepos(value EnterpriseRepositoryOverviewable)() {
    m.repos = value
}
// SetUsers sets the users property value. The users property
func (m *EnterpriseOverview) SetUsers(value EnterpriseUserOverviewable)() {
    m.users = value
}
type EnterpriseOverviewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComments()(EnterpriseCommentOverviewable)
    GetGists()(EnterpriseGistOverviewable)
    GetHooks()(EnterpriseHookOverviewable)
    GetIssues()(EnterpriseIssueOverviewable)
    GetMilestones()(EnterpriseMilestoneOverviewable)
    GetOrgs()(EnterpriseOrganizationOverviewable)
    GetPages()(EnterprisePageOverviewable)
    GetPulls()(EnterprisePullRequestOverviewable)
    GetRepos()(EnterpriseRepositoryOverviewable)
    GetUsers()(EnterpriseUserOverviewable)
    SetComments(value EnterpriseCommentOverviewable)()
    SetGists(value EnterpriseGistOverviewable)()
    SetHooks(value EnterpriseHookOverviewable)()
    SetIssues(value EnterpriseIssueOverviewable)()
    SetMilestones(value EnterpriseMilestoneOverviewable)()
    SetOrgs(value EnterpriseOrganizationOverviewable)()
    SetPages(value EnterprisePageOverviewable)()
    SetPulls(value EnterprisePullRequestOverviewable)()
    SetRepos(value EnterpriseRepositoryOverviewable)()
    SetUsers(value EnterpriseUserOverviewable)()
}
