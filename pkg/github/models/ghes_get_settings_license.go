package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesGetSettings_license struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The cluster_support property
    cluster_support *bool
    // The evaluation property
    evaluation *bool
    // The expire_at property
    expire_at *string
    // The perpetual property
    perpetual *bool
    // The seats property
    seats *int32
    // The ssh_allowed property
    ssh_allowed *bool
    // The support_key property
    support_key *string
    // The unlimited_seating property
    unlimited_seating *bool
}
// NewGhesGetSettings_license instantiates a new GhesGetSettings_license and sets the default values.
func NewGhesGetSettings_license()(*GhesGetSettings_license) {
    m := &GhesGetSettings_license{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesGetSettings_licenseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesGetSettings_licenseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesGetSettings_license(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesGetSettings_license) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClusterSupport gets the cluster_support property value. The cluster_support property
// returns a *bool when successful
func (m *GhesGetSettings_license) GetClusterSupport()(*bool) {
    return m.cluster_support
}
// GetEvaluation gets the evaluation property value. The evaluation property
// returns a *bool when successful
func (m *GhesGetSettings_license) GetEvaluation()(*bool) {
    return m.evaluation
}
// GetExpireAt gets the expire_at property value. The expire_at property
// returns a *string when successful
func (m *GhesGetSettings_license) GetExpireAt()(*string) {
    return m.expire_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesGetSettings_license) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cluster_support"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClusterSupport(val)
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
    res["expire_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpireAt(val)
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
    res["ssh_allowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSshAllowed(val)
        }
        return nil
    }
    res["support_key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportKey(val)
        }
        return nil
    }
    res["unlimited_seating"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetPerpetual gets the perpetual property value. The perpetual property
// returns a *bool when successful
func (m *GhesGetSettings_license) GetPerpetual()(*bool) {
    return m.perpetual
}
// GetSeats gets the seats property value. The seats property
// returns a *int32 when successful
func (m *GhesGetSettings_license) GetSeats()(*int32) {
    return m.seats
}
// GetSshAllowed gets the ssh_allowed property value. The ssh_allowed property
// returns a *bool when successful
func (m *GhesGetSettings_license) GetSshAllowed()(*bool) {
    return m.ssh_allowed
}
// GetSupportKey gets the support_key property value. The support_key property
// returns a *string when successful
func (m *GhesGetSettings_license) GetSupportKey()(*string) {
    return m.support_key
}
// GetUnlimitedSeating gets the unlimited_seating property value. The unlimited_seating property
// returns a *bool when successful
func (m *GhesGetSettings_license) GetUnlimitedSeating()(*bool) {
    return m.unlimited_seating
}
// Serialize serializes information the current object
func (m *GhesGetSettings_license) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("cluster_support", m.GetClusterSupport())
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
        err := writer.WriteStringValue("expire_at", m.GetExpireAt())
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
        err := writer.WriteInt32Value("seats", m.GetSeats())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ssh_allowed", m.GetSshAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("support_key", m.GetSupportKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("unlimited_seating", m.GetUnlimitedSeating())
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
func (m *GhesGetSettings_license) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClusterSupport sets the cluster_support property value. The cluster_support property
func (m *GhesGetSettings_license) SetClusterSupport(value *bool)() {
    m.cluster_support = value
}
// SetEvaluation sets the evaluation property value. The evaluation property
func (m *GhesGetSettings_license) SetEvaluation(value *bool)() {
    m.evaluation = value
}
// SetExpireAt sets the expire_at property value. The expire_at property
func (m *GhesGetSettings_license) SetExpireAt(value *string)() {
    m.expire_at = value
}
// SetPerpetual sets the perpetual property value. The perpetual property
func (m *GhesGetSettings_license) SetPerpetual(value *bool)() {
    m.perpetual = value
}
// SetSeats sets the seats property value. The seats property
func (m *GhesGetSettings_license) SetSeats(value *int32)() {
    m.seats = value
}
// SetSshAllowed sets the ssh_allowed property value. The ssh_allowed property
func (m *GhesGetSettings_license) SetSshAllowed(value *bool)() {
    m.ssh_allowed = value
}
// SetSupportKey sets the support_key property value. The support_key property
func (m *GhesGetSettings_license) SetSupportKey(value *string)() {
    m.support_key = value
}
// SetUnlimitedSeating sets the unlimited_seating property value. The unlimited_seating property
func (m *GhesGetSettings_license) SetUnlimitedSeating(value *bool)() {
    m.unlimited_seating = value
}
type GhesGetSettings_licenseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClusterSupport()(*bool)
    GetEvaluation()(*bool)
    GetExpireAt()(*string)
    GetPerpetual()(*bool)
    GetSeats()(*int32)
    GetSshAllowed()(*bool)
    GetSupportKey()(*string)
    GetUnlimitedSeating()(*bool)
    SetClusterSupport(value *bool)()
    SetEvaluation(value *bool)()
    SetExpireAt(value *string)()
    SetPerpetual(value *bool)()
    SetSeats(value *int32)()
    SetSshAllowed(value *bool)()
    SetSupportKey(value *string)()
    SetUnlimitedSeating(value *bool)()
}
