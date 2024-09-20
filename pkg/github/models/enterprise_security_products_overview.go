package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EnterpriseSecurityProductsOverview struct {
    // The active_committers property
    active_committers *int32
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The advanced_security_enabled_repos property
    advanced_security_enabled_repos *int32
    // The code_scanning_default_setup_eligible_repos property
    code_scanning_default_setup_eligible_repos *int32
    // The code_scanning_default_setup_enabled_repos property
    code_scanning_default_setup_enabled_repos *int32
    // The code_scanning_enabled_repos property
    code_scanning_enabled_repos *int32
    // The code_scanning_pr_reviews_enabled_repos property
    code_scanning_pr_reviews_enabled_repos *int32
    // The dependabot_alerts_enabled_repos property
    dependabot_alerts_enabled_repos *int32
    // The dependabot_security_updates_enabled_repos property
    dependabot_security_updates_enabled_repos *int32
    // The dependabot_version_updates_enabled_repos property
    dependabot_version_updates_enabled_repos *int32
    // The maximum_committers property
    maximum_committers *int32
    // The nonarchived_repos property
    nonarchived_repos *int32
    // The purchased_committers property
    purchased_committers *int32
    // The secret_scanning_enabled_repos property
    secret_scanning_enabled_repos *int32
    // The secret_scanning_push_protection_enabled_repos property
    secret_scanning_push_protection_enabled_repos *int32
    // The total_repos property
    total_repos *int32
}
// NewEnterpriseSecurityProductsOverview instantiates a new EnterpriseSecurityProductsOverview and sets the default values.
func NewEnterpriseSecurityProductsOverview()(*EnterpriseSecurityProductsOverview) {
    m := &EnterpriseSecurityProductsOverview{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnterpriseSecurityProductsOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnterpriseSecurityProductsOverviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnterpriseSecurityProductsOverview(), nil
}
// GetActiveCommitters gets the active_committers property value. The active_committers property
// returns a *int32 when successful
func (m *EnterpriseSecurityProductsOverview) GetActiveCommitters()(*int32) {
    return m.active_committers
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EnterpriseSecurityProductsOverview) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdvancedSecurityEnabledRepos gets the advanced_security_enabled_repos property value. The advanced_security_enabled_repos property
// returns a *int32 when successful
func (m *EnterpriseSecurityProductsOverview) GetAdvancedSecurityEnabledRepos()(*int32) {
    return m.advanced_security_enabled_repos
}
// GetCodeScanningDefaultSetupEligibleRepos gets the code_scanning_default_setup_eligible_repos property value. The code_scanning_default_setup_eligible_repos property
// returns a *int32 when successful
func (m *EnterpriseSecurityProductsOverview) GetCodeScanningDefaultSetupEligibleRepos()(*int32) {
    return m.code_scanning_default_setup_eligible_repos
}
// GetCodeScanningDefaultSetupEnabledRepos gets the code_scanning_default_setup_enabled_repos property value. The code_scanning_default_setup_enabled_repos property
// returns a *int32 when successful
func (m *EnterpriseSecurityProductsOverview) GetCodeScanningDefaultSetupEnabledRepos()(*int32) {
    return m.code_scanning_default_setup_enabled_repos
}
// GetCodeScanningEnabledRepos gets the code_scanning_enabled_repos property value. The code_scanning_enabled_repos property
// returns a *int32 when successful
func (m *EnterpriseSecurityProductsOverview) GetCodeScanningEnabledRepos()(*int32) {
    return m.code_scanning_enabled_repos
}
// GetCodeScanningPrReviewsEnabledRepos gets the code_scanning_pr_reviews_enabled_repos property value. The code_scanning_pr_reviews_enabled_repos property
// returns a *int32 when successful
func (m *EnterpriseSecurityProductsOverview) GetCodeScanningPrReviewsEnabledRepos()(*int32) {
    return m.code_scanning_pr_reviews_enabled_repos
}
// GetDependabotAlertsEnabledRepos gets the dependabot_alerts_enabled_repos property value. The dependabot_alerts_enabled_repos property
// returns a *int32 when successful
func (m *EnterpriseSecurityProductsOverview) GetDependabotAlertsEnabledRepos()(*int32) {
    return m.dependabot_alerts_enabled_repos
}
// GetDependabotSecurityUpdatesEnabledRepos gets the dependabot_security_updates_enabled_repos property value. The dependabot_security_updates_enabled_repos property
// returns a *int32 when successful
func (m *EnterpriseSecurityProductsOverview) GetDependabotSecurityUpdatesEnabledRepos()(*int32) {
    return m.dependabot_security_updates_enabled_repos
}
// GetDependabotVersionUpdatesEnabledRepos gets the dependabot_version_updates_enabled_repos property value. The dependabot_version_updates_enabled_repos property
// returns a *int32 when successful
func (m *EnterpriseSecurityProductsOverview) GetDependabotVersionUpdatesEnabledRepos()(*int32) {
    return m.dependabot_version_updates_enabled_repos
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnterpriseSecurityProductsOverview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["active_committers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveCommitters(val)
        }
        return nil
    }
    res["advanced_security_enabled_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedSecurityEnabledRepos(val)
        }
        return nil
    }
    res["code_scanning_default_setup_eligible_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCodeScanningDefaultSetupEligibleRepos(val)
        }
        return nil
    }
    res["code_scanning_default_setup_enabled_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCodeScanningDefaultSetupEnabledRepos(val)
        }
        return nil
    }
    res["code_scanning_enabled_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCodeScanningEnabledRepos(val)
        }
        return nil
    }
    res["code_scanning_pr_reviews_enabled_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCodeScanningPrReviewsEnabledRepos(val)
        }
        return nil
    }
    res["dependabot_alerts_enabled_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependabotAlertsEnabledRepos(val)
        }
        return nil
    }
    res["dependabot_security_updates_enabled_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependabotSecurityUpdatesEnabledRepos(val)
        }
        return nil
    }
    res["dependabot_version_updates_enabled_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependabotVersionUpdatesEnabledRepos(val)
        }
        return nil
    }
    res["maximum_committers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumCommitters(val)
        }
        return nil
    }
    res["nonarchived_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNonarchivedRepos(val)
        }
        return nil
    }
    res["purchased_committers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPurchasedCommitters(val)
        }
        return nil
    }
    res["secret_scanning_enabled_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningEnabledRepos(val)
        }
        return nil
    }
    res["secret_scanning_push_protection_enabled_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningPushProtectionEnabledRepos(val)
        }
        return nil
    }
    res["total_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalRepos(val)
        }
        return nil
    }
    return res
}
// GetMaximumCommitters gets the maximum_committers property value. The maximum_committers property
// returns a *int32 when successful
func (m *EnterpriseSecurityProductsOverview) GetMaximumCommitters()(*int32) {
    return m.maximum_committers
}
// GetNonarchivedRepos gets the nonarchived_repos property value. The nonarchived_repos property
// returns a *int32 when successful
func (m *EnterpriseSecurityProductsOverview) GetNonarchivedRepos()(*int32) {
    return m.nonarchived_repos
}
// GetPurchasedCommitters gets the purchased_committers property value. The purchased_committers property
// returns a *int32 when successful
func (m *EnterpriseSecurityProductsOverview) GetPurchasedCommitters()(*int32) {
    return m.purchased_committers
}
// GetSecretScanningEnabledRepos gets the secret_scanning_enabled_repos property value. The secret_scanning_enabled_repos property
// returns a *int32 when successful
func (m *EnterpriseSecurityProductsOverview) GetSecretScanningEnabledRepos()(*int32) {
    return m.secret_scanning_enabled_repos
}
// GetSecretScanningPushProtectionEnabledRepos gets the secret_scanning_push_protection_enabled_repos property value. The secret_scanning_push_protection_enabled_repos property
// returns a *int32 when successful
func (m *EnterpriseSecurityProductsOverview) GetSecretScanningPushProtectionEnabledRepos()(*int32) {
    return m.secret_scanning_push_protection_enabled_repos
}
// GetTotalRepos gets the total_repos property value. The total_repos property
// returns a *int32 when successful
func (m *EnterpriseSecurityProductsOverview) GetTotalRepos()(*int32) {
    return m.total_repos
}
// Serialize serializes information the current object
func (m *EnterpriseSecurityProductsOverview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("active_committers", m.GetActiveCommitters())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("advanced_security_enabled_repos", m.GetAdvancedSecurityEnabledRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("code_scanning_default_setup_eligible_repos", m.GetCodeScanningDefaultSetupEligibleRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("code_scanning_default_setup_enabled_repos", m.GetCodeScanningDefaultSetupEnabledRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("code_scanning_enabled_repos", m.GetCodeScanningEnabledRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("code_scanning_pr_reviews_enabled_repos", m.GetCodeScanningPrReviewsEnabledRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("dependabot_alerts_enabled_repos", m.GetDependabotAlertsEnabledRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("dependabot_security_updates_enabled_repos", m.GetDependabotSecurityUpdatesEnabledRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("dependabot_version_updates_enabled_repos", m.GetDependabotVersionUpdatesEnabledRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maximum_committers", m.GetMaximumCommitters())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("nonarchived_repos", m.GetNonarchivedRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("purchased_committers", m.GetPurchasedCommitters())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("secret_scanning_enabled_repos", m.GetSecretScanningEnabledRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("secret_scanning_push_protection_enabled_repos", m.GetSecretScanningPushProtectionEnabledRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_repos", m.GetTotalRepos())
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
// SetActiveCommitters sets the active_committers property value. The active_committers property
func (m *EnterpriseSecurityProductsOverview) SetActiveCommitters(value *int32)() {
    m.active_committers = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EnterpriseSecurityProductsOverview) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdvancedSecurityEnabledRepos sets the advanced_security_enabled_repos property value. The advanced_security_enabled_repos property
func (m *EnterpriseSecurityProductsOverview) SetAdvancedSecurityEnabledRepos(value *int32)() {
    m.advanced_security_enabled_repos = value
}
// SetCodeScanningDefaultSetupEligibleRepos sets the code_scanning_default_setup_eligible_repos property value. The code_scanning_default_setup_eligible_repos property
func (m *EnterpriseSecurityProductsOverview) SetCodeScanningDefaultSetupEligibleRepos(value *int32)() {
    m.code_scanning_default_setup_eligible_repos = value
}
// SetCodeScanningDefaultSetupEnabledRepos sets the code_scanning_default_setup_enabled_repos property value. The code_scanning_default_setup_enabled_repos property
func (m *EnterpriseSecurityProductsOverview) SetCodeScanningDefaultSetupEnabledRepos(value *int32)() {
    m.code_scanning_default_setup_enabled_repos = value
}
// SetCodeScanningEnabledRepos sets the code_scanning_enabled_repos property value. The code_scanning_enabled_repos property
func (m *EnterpriseSecurityProductsOverview) SetCodeScanningEnabledRepos(value *int32)() {
    m.code_scanning_enabled_repos = value
}
// SetCodeScanningPrReviewsEnabledRepos sets the code_scanning_pr_reviews_enabled_repos property value. The code_scanning_pr_reviews_enabled_repos property
func (m *EnterpriseSecurityProductsOverview) SetCodeScanningPrReviewsEnabledRepos(value *int32)() {
    m.code_scanning_pr_reviews_enabled_repos = value
}
// SetDependabotAlertsEnabledRepos sets the dependabot_alerts_enabled_repos property value. The dependabot_alerts_enabled_repos property
func (m *EnterpriseSecurityProductsOverview) SetDependabotAlertsEnabledRepos(value *int32)() {
    m.dependabot_alerts_enabled_repos = value
}
// SetDependabotSecurityUpdatesEnabledRepos sets the dependabot_security_updates_enabled_repos property value. The dependabot_security_updates_enabled_repos property
func (m *EnterpriseSecurityProductsOverview) SetDependabotSecurityUpdatesEnabledRepos(value *int32)() {
    m.dependabot_security_updates_enabled_repos = value
}
// SetDependabotVersionUpdatesEnabledRepos sets the dependabot_version_updates_enabled_repos property value. The dependabot_version_updates_enabled_repos property
func (m *EnterpriseSecurityProductsOverview) SetDependabotVersionUpdatesEnabledRepos(value *int32)() {
    m.dependabot_version_updates_enabled_repos = value
}
// SetMaximumCommitters sets the maximum_committers property value. The maximum_committers property
func (m *EnterpriseSecurityProductsOverview) SetMaximumCommitters(value *int32)() {
    m.maximum_committers = value
}
// SetNonarchivedRepos sets the nonarchived_repos property value. The nonarchived_repos property
func (m *EnterpriseSecurityProductsOverview) SetNonarchivedRepos(value *int32)() {
    m.nonarchived_repos = value
}
// SetPurchasedCommitters sets the purchased_committers property value. The purchased_committers property
func (m *EnterpriseSecurityProductsOverview) SetPurchasedCommitters(value *int32)() {
    m.purchased_committers = value
}
// SetSecretScanningEnabledRepos sets the secret_scanning_enabled_repos property value. The secret_scanning_enabled_repos property
func (m *EnterpriseSecurityProductsOverview) SetSecretScanningEnabledRepos(value *int32)() {
    m.secret_scanning_enabled_repos = value
}
// SetSecretScanningPushProtectionEnabledRepos sets the secret_scanning_push_protection_enabled_repos property value. The secret_scanning_push_protection_enabled_repos property
func (m *EnterpriseSecurityProductsOverview) SetSecretScanningPushProtectionEnabledRepos(value *int32)() {
    m.secret_scanning_push_protection_enabled_repos = value
}
// SetTotalRepos sets the total_repos property value. The total_repos property
func (m *EnterpriseSecurityProductsOverview) SetTotalRepos(value *int32)() {
    m.total_repos = value
}
type EnterpriseSecurityProductsOverviewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveCommitters()(*int32)
    GetAdvancedSecurityEnabledRepos()(*int32)
    GetCodeScanningDefaultSetupEligibleRepos()(*int32)
    GetCodeScanningDefaultSetupEnabledRepos()(*int32)
    GetCodeScanningEnabledRepos()(*int32)
    GetCodeScanningPrReviewsEnabledRepos()(*int32)
    GetDependabotAlertsEnabledRepos()(*int32)
    GetDependabotSecurityUpdatesEnabledRepos()(*int32)
    GetDependabotVersionUpdatesEnabledRepos()(*int32)
    GetMaximumCommitters()(*int32)
    GetNonarchivedRepos()(*int32)
    GetPurchasedCommitters()(*int32)
    GetSecretScanningEnabledRepos()(*int32)
    GetSecretScanningPushProtectionEnabledRepos()(*int32)
    GetTotalRepos()(*int32)
    SetActiveCommitters(value *int32)()
    SetAdvancedSecurityEnabledRepos(value *int32)()
    SetCodeScanningDefaultSetupEligibleRepos(value *int32)()
    SetCodeScanningDefaultSetupEnabledRepos(value *int32)()
    SetCodeScanningEnabledRepos(value *int32)()
    SetCodeScanningPrReviewsEnabledRepos(value *int32)()
    SetDependabotAlertsEnabledRepos(value *int32)()
    SetDependabotSecurityUpdatesEnabledRepos(value *int32)()
    SetDependabotVersionUpdatesEnabledRepos(value *int32)()
    SetMaximumCommitters(value *int32)()
    SetNonarchivedRepos(value *int32)()
    SetPurchasedCommitters(value *int32)()
    SetSecretScanningEnabledRepos(value *int32)()
    SetSecretScanningPushProtectionEnabledRepos(value *int32)()
    SetTotalRepos(value *int32)()
}
