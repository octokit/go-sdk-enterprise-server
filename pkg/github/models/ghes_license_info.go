package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesLicenseInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether the GitHub Advanced Security feature is enabled.
    advancedSecurityEnabled *bool
    // If the GitHub Advanced Security feature is enabled, the seat count.
    advancedSecuritySeats *int32
    // Whether the cluster support feature is enabled.
    clusterSupport *bool
    // The company under which the license is issued.
    company *string
    // Whether the Github Connect feature is enabled.
    croquetSupport *bool
    // Whether this license is issued under custom terms.
    customTerms *bool
    // Wheter this license is issued as an evaluation license.
    evaluation *bool
    // The expiration date of the license.
    expireAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Whether the insights feature is enabled.
    insightsEnabled *bool
    // If the insights feature is enabled, the expiration date.
    insightsExpireAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // If the learning lab feature is enabled, the expiration date.
    learningLabEvaluationExpires *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // If the learning lab feature is enabled, the seat count.
    learningLabSeats *int32
    // Whether the perpetual feature is enabled.
    perpetual *bool
    // The unique reference number of the license.
    referenceNumber *string
    // If the license is issued with limited seating, the seat count.
    seats *int32
    // Whether the SSH feature is enabled.
    sshAllowed *bool
    // The support key of the license.
    supportKey *string
    // Whether the license is issued with unlimited seat count.
    unlimitedSeating *bool
}
// NewGhesLicenseInfo instantiates a new GhesLicenseInfo and sets the default values.
func NewGhesLicenseInfo()(*GhesLicenseInfo) {
    m := &GhesLicenseInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesLicenseInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesLicenseInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesLicenseInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesLicenseInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdvancedSecurityEnabled gets the advancedSecurityEnabled property value. Whether the GitHub Advanced Security feature is enabled.
// returns a *bool when successful
func (m *GhesLicenseInfo) GetAdvancedSecurityEnabled()(*bool) {
    return m.advancedSecurityEnabled
}
// GetAdvancedSecuritySeats gets the advancedSecuritySeats property value. If the GitHub Advanced Security feature is enabled, the seat count.
// returns a *int32 when successful
func (m *GhesLicenseInfo) GetAdvancedSecuritySeats()(*int32) {
    return m.advancedSecuritySeats
}
// GetClusterSupport gets the clusterSupport property value. Whether the cluster support feature is enabled.
// returns a *bool when successful
func (m *GhesLicenseInfo) GetClusterSupport()(*bool) {
    return m.clusterSupport
}
// GetCompany gets the company property value. The company under which the license is issued.
// returns a *string when successful
func (m *GhesLicenseInfo) GetCompany()(*string) {
    return m.company
}
// GetCroquetSupport gets the croquetSupport property value. Whether the Github Connect feature is enabled.
// returns a *bool when successful
func (m *GhesLicenseInfo) GetCroquetSupport()(*bool) {
    return m.croquetSupport
}
// GetCustomTerms gets the customTerms property value. Whether this license is issued under custom terms.
// returns a *bool when successful
func (m *GhesLicenseInfo) GetCustomTerms()(*bool) {
    return m.customTerms
}
// GetEvaluation gets the evaluation property value. Wheter this license is issued as an evaluation license.
// returns a *bool when successful
func (m *GhesLicenseInfo) GetEvaluation()(*bool) {
    return m.evaluation
}
// GetExpireAt gets the expireAt property value. The expiration date of the license.
// returns a *Time when successful
func (m *GhesLicenseInfo) GetExpireAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expireAt
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesLicenseInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["advancedSecurityEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedSecurityEnabled(val)
        }
        return nil
    }
    res["advancedSecuritySeats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedSecuritySeats(val)
        }
        return nil
    }
    res["clusterSupport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClusterSupport(val)
        }
        return nil
    }
    res["company"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompany(val)
        }
        return nil
    }
    res["croquetSupport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCroquetSupport(val)
        }
        return nil
    }
    res["customTerms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomTerms(val)
        }
        return nil
    }
    res["evaluation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvaluation(val)
        }
        return nil
    }
    res["expireAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpireAt(val)
        }
        return nil
    }
    res["insightsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInsightsEnabled(val)
        }
        return nil
    }
    res["insightsExpireAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInsightsExpireAt(val)
        }
        return nil
    }
    res["learningLabEvaluationExpires"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLearningLabEvaluationExpires(val)
        }
        return nil
    }
    res["learningLabSeats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLearningLabSeats(val)
        }
        return nil
    }
    res["perpetual"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPerpetual(val)
        }
        return nil
    }
    res["referenceNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferenceNumber(val)
        }
        return nil
    }
    res["seats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeats(val)
        }
        return nil
    }
    res["sshAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSshAllowed(val)
        }
        return nil
    }
    res["supportKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportKey(val)
        }
        return nil
    }
    res["unlimitedSeating"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnlimitedSeating(val)
        }
        return nil
    }
    return res
}
// GetInsightsEnabled gets the insightsEnabled property value. Whether the insights feature is enabled.
// returns a *bool when successful
func (m *GhesLicenseInfo) GetInsightsEnabled()(*bool) {
    return m.insightsEnabled
}
// GetInsightsExpireAt gets the insightsExpireAt property value. If the insights feature is enabled, the expiration date.
// returns a *Time when successful
func (m *GhesLicenseInfo) GetInsightsExpireAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.insightsExpireAt
}
// GetLearningLabEvaluationExpires gets the learningLabEvaluationExpires property value. If the learning lab feature is enabled, the expiration date.
// returns a *Time when successful
func (m *GhesLicenseInfo) GetLearningLabEvaluationExpires()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.learningLabEvaluationExpires
}
// GetLearningLabSeats gets the learningLabSeats property value. If the learning lab feature is enabled, the seat count.
// returns a *int32 when successful
func (m *GhesLicenseInfo) GetLearningLabSeats()(*int32) {
    return m.learningLabSeats
}
// GetPerpetual gets the perpetual property value. Whether the perpetual feature is enabled.
// returns a *bool when successful
func (m *GhesLicenseInfo) GetPerpetual()(*bool) {
    return m.perpetual
}
// GetReferenceNumber gets the referenceNumber property value. The unique reference number of the license.
// returns a *string when successful
func (m *GhesLicenseInfo) GetReferenceNumber()(*string) {
    return m.referenceNumber
}
// GetSeats gets the seats property value. If the license is issued with limited seating, the seat count.
// returns a *int32 when successful
func (m *GhesLicenseInfo) GetSeats()(*int32) {
    return m.seats
}
// GetSshAllowed gets the sshAllowed property value. Whether the SSH feature is enabled.
// returns a *bool when successful
func (m *GhesLicenseInfo) GetSshAllowed()(*bool) {
    return m.sshAllowed
}
// GetSupportKey gets the supportKey property value. The support key of the license.
// returns a *string when successful
func (m *GhesLicenseInfo) GetSupportKey()(*string) {
    return m.supportKey
}
// GetUnlimitedSeating gets the unlimitedSeating property value. Whether the license is issued with unlimited seat count.
// returns a *bool when successful
func (m *GhesLicenseInfo) GetUnlimitedSeating()(*bool) {
    return m.unlimitedSeating
}
// Serialize serializes information the current object
func (m *GhesLicenseInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("advancedSecurityEnabled", m.GetAdvancedSecurityEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("advancedSecuritySeats", m.GetAdvancedSecuritySeats())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("clusterSupport", m.GetClusterSupport())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("company", m.GetCompany())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("croquetSupport", m.GetCroquetSupport())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("customTerms", m.GetCustomTerms())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("evaluation", m.GetEvaluation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("expireAt", m.GetExpireAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("insightsEnabled", m.GetInsightsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("insightsExpireAt", m.GetInsightsExpireAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("learningLabEvaluationExpires", m.GetLearningLabEvaluationExpires())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("learningLabSeats", m.GetLearningLabSeats())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("perpetual", m.GetPerpetual())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("referenceNumber", m.GetReferenceNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("seats", m.GetSeats())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sshAllowed", m.GetSshAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("supportKey", m.GetSupportKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("unlimitedSeating", m.GetUnlimitedSeating())
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
func (m *GhesLicenseInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdvancedSecurityEnabled sets the advancedSecurityEnabled property value. Whether the GitHub Advanced Security feature is enabled.
func (m *GhesLicenseInfo) SetAdvancedSecurityEnabled(value *bool)() {
    m.advancedSecurityEnabled = value
}
// SetAdvancedSecuritySeats sets the advancedSecuritySeats property value. If the GitHub Advanced Security feature is enabled, the seat count.
func (m *GhesLicenseInfo) SetAdvancedSecuritySeats(value *int32)() {
    m.advancedSecuritySeats = value
}
// SetClusterSupport sets the clusterSupport property value. Whether the cluster support feature is enabled.
func (m *GhesLicenseInfo) SetClusterSupport(value *bool)() {
    m.clusterSupport = value
}
// SetCompany sets the company property value. The company under which the license is issued.
func (m *GhesLicenseInfo) SetCompany(value *string)() {
    m.company = value
}
// SetCroquetSupport sets the croquetSupport property value. Whether the Github Connect feature is enabled.
func (m *GhesLicenseInfo) SetCroquetSupport(value *bool)() {
    m.croquetSupport = value
}
// SetCustomTerms sets the customTerms property value. Whether this license is issued under custom terms.
func (m *GhesLicenseInfo) SetCustomTerms(value *bool)() {
    m.customTerms = value
}
// SetEvaluation sets the evaluation property value. Wheter this license is issued as an evaluation license.
func (m *GhesLicenseInfo) SetEvaluation(value *bool)() {
    m.evaluation = value
}
// SetExpireAt sets the expireAt property value. The expiration date of the license.
func (m *GhesLicenseInfo) SetExpireAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expireAt = value
}
// SetInsightsEnabled sets the insightsEnabled property value. Whether the insights feature is enabled.
func (m *GhesLicenseInfo) SetInsightsEnabled(value *bool)() {
    m.insightsEnabled = value
}
// SetInsightsExpireAt sets the insightsExpireAt property value. If the insights feature is enabled, the expiration date.
func (m *GhesLicenseInfo) SetInsightsExpireAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.insightsExpireAt = value
}
// SetLearningLabEvaluationExpires sets the learningLabEvaluationExpires property value. If the learning lab feature is enabled, the expiration date.
func (m *GhesLicenseInfo) SetLearningLabEvaluationExpires(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.learningLabEvaluationExpires = value
}
// SetLearningLabSeats sets the learningLabSeats property value. If the learning lab feature is enabled, the seat count.
func (m *GhesLicenseInfo) SetLearningLabSeats(value *int32)() {
    m.learningLabSeats = value
}
// SetPerpetual sets the perpetual property value. Whether the perpetual feature is enabled.
func (m *GhesLicenseInfo) SetPerpetual(value *bool)() {
    m.perpetual = value
}
// SetReferenceNumber sets the referenceNumber property value. The unique reference number of the license.
func (m *GhesLicenseInfo) SetReferenceNumber(value *string)() {
    m.referenceNumber = value
}
// SetSeats sets the seats property value. If the license is issued with limited seating, the seat count.
func (m *GhesLicenseInfo) SetSeats(value *int32)() {
    m.seats = value
}
// SetSshAllowed sets the sshAllowed property value. Whether the SSH feature is enabled.
func (m *GhesLicenseInfo) SetSshAllowed(value *bool)() {
    m.sshAllowed = value
}
// SetSupportKey sets the supportKey property value. The support key of the license.
func (m *GhesLicenseInfo) SetSupportKey(value *string)() {
    m.supportKey = value
}
// SetUnlimitedSeating sets the unlimitedSeating property value. Whether the license is issued with unlimited seat count.
func (m *GhesLicenseInfo) SetUnlimitedSeating(value *bool)() {
    m.unlimitedSeating = value
}
type GhesLicenseInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvancedSecurityEnabled()(*bool)
    GetAdvancedSecuritySeats()(*int32)
    GetClusterSupport()(*bool)
    GetCompany()(*string)
    GetCroquetSupport()(*bool)
    GetCustomTerms()(*bool)
    GetEvaluation()(*bool)
    GetExpireAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetInsightsEnabled()(*bool)
    GetInsightsExpireAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLearningLabEvaluationExpires()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLearningLabSeats()(*int32)
    GetPerpetual()(*bool)
    GetReferenceNumber()(*string)
    GetSeats()(*int32)
    GetSshAllowed()(*bool)
    GetSupportKey()(*string)
    GetUnlimitedSeating()(*bool)
    SetAdvancedSecurityEnabled(value *bool)()
    SetAdvancedSecuritySeats(value *int32)()
    SetClusterSupport(value *bool)()
    SetCompany(value *string)()
    SetCroquetSupport(value *bool)()
    SetCustomTerms(value *bool)()
    SetEvaluation(value *bool)()
    SetExpireAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetInsightsEnabled(value *bool)()
    SetInsightsExpireAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLearningLabEvaluationExpires(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLearningLabSeats(value *int32)()
    SetPerpetual(value *bool)()
    SetReferenceNumber(value *string)()
    SetSeats(value *int32)()
    SetSshAllowed(value *bool)()
    SetSupportKey(value *string)()
    SetUnlimitedSeating(value *bool)()
}
