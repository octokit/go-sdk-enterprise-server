package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EnterpriseSettings_enterprise_github_oauth struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The client_id property
    client_id *string
    // The client_secret property
    client_secret *string
    // The organization_name property
    organization_name *string
    // The organization_team property
    organization_team *string
}
// NewEnterpriseSettings_enterprise_github_oauth instantiates a new EnterpriseSettings_enterprise_github_oauth and sets the default values.
func NewEnterpriseSettings_enterprise_github_oauth()(*EnterpriseSettings_enterprise_github_oauth) {
    m := &EnterpriseSettings_enterprise_github_oauth{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnterpriseSettings_enterprise_github_oauthFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnterpriseSettings_enterprise_github_oauthFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnterpriseSettings_enterprise_github_oauth(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EnterpriseSettings_enterprise_github_oauth) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClientId gets the client_id property value. The client_id property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_github_oauth) GetClientId()(*string) {
    return m.client_id
}
// GetClientSecret gets the client_secret property value. The client_secret property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_github_oauth) GetClientSecret()(*string) {
    return m.client_secret
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnterpriseSettings_enterprise_github_oauth) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["client_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientId(val)
        }
        return nil
    }
    res["client_secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientSecret(val)
        }
        return nil
    }
    res["organization_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationName(val)
        }
        return nil
    }
    res["organization_team"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationTeam(val)
        }
        return nil
    }
    return res
}
// GetOrganizationName gets the organization_name property value. The organization_name property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_github_oauth) GetOrganizationName()(*string) {
    return m.organization_name
}
// GetOrganizationTeam gets the organization_team property value. The organization_team property
// returns a *string when successful
func (m *EnterpriseSettings_enterprise_github_oauth) GetOrganizationTeam()(*string) {
    return m.organization_team
}
// Serialize serializes information the current object
func (m *EnterpriseSettings_enterprise_github_oauth) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("client_id", m.GetClientId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("client_secret", m.GetClientSecret())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("organization_name", m.GetOrganizationName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("organization_team", m.GetOrganizationTeam())
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
func (m *EnterpriseSettings_enterprise_github_oauth) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClientId sets the client_id property value. The client_id property
func (m *EnterpriseSettings_enterprise_github_oauth) SetClientId(value *string)() {
    m.client_id = value
}
// SetClientSecret sets the client_secret property value. The client_secret property
func (m *EnterpriseSettings_enterprise_github_oauth) SetClientSecret(value *string)() {
    m.client_secret = value
}
// SetOrganizationName sets the organization_name property value. The organization_name property
func (m *EnterpriseSettings_enterprise_github_oauth) SetOrganizationName(value *string)() {
    m.organization_name = value
}
// SetOrganizationTeam sets the organization_team property value. The organization_team property
func (m *EnterpriseSettings_enterprise_github_oauth) SetOrganizationTeam(value *string)() {
    m.organization_team = value
}
type EnterpriseSettings_enterprise_github_oauthable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientId()(*string)
    GetClientSecret()(*string)
    GetOrganizationName()(*string)
    GetOrganizationTeam()(*string)
    SetClientId(value *string)()
    SetClientSecret(value *string)()
    SetOrganizationName(value *string)()
    SetOrganizationTeam(value *string)()
}
