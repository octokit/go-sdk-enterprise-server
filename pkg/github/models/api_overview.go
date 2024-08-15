package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApiOverview api Overview
type ApiOverview struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The domains property
    domains ApiOverview_domainsable
    // The installed_version property
    installed_version *string
    // The packages property
    packages []string
    // The verifiable_password_authentication property
    verifiable_password_authentication *bool
}
// NewApiOverview instantiates a new ApiOverview and sets the default values.
func NewApiOverview()(*ApiOverview) {
    m := &ApiOverview{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateApiOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApiOverviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApiOverview(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ApiOverview) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDomains gets the domains property value. The domains property
// returns a ApiOverview_domainsable when successful
func (m *ApiOverview) GetDomains()(ApiOverview_domainsable) {
    return m.domains
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApiOverview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["domains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApiOverview_domainsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomains(val.(ApiOverview_domainsable))
        }
        return nil
    }
    res["installed_version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstalledVersion(val)
        }
        return nil
    }
    res["packages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetPackages(res)
        }
        return nil
    }
    res["verifiable_password_authentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiablePasswordAuthentication(val)
        }
        return nil
    }
    return res
}
// GetInstalledVersion gets the installed_version property value. The installed_version property
// returns a *string when successful
func (m *ApiOverview) GetInstalledVersion()(*string) {
    return m.installed_version
}
// GetPackages gets the packages property value. The packages property
// returns a []string when successful
func (m *ApiOverview) GetPackages()([]string) {
    return m.packages
}
// GetVerifiablePasswordAuthentication gets the verifiable_password_authentication property value. The verifiable_password_authentication property
// returns a *bool when successful
func (m *ApiOverview) GetVerifiablePasswordAuthentication()(*bool) {
    return m.verifiable_password_authentication
}
// Serialize serializes information the current object
func (m *ApiOverview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("domains", m.GetDomains())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("installed_version", m.GetInstalledVersion())
        if err != nil {
            return err
        }
    }
    if m.GetPackages() != nil {
        err := writer.WriteCollectionOfStringValues("packages", m.GetPackages())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("verifiable_password_authentication", m.GetVerifiablePasswordAuthentication())
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
func (m *ApiOverview) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDomains sets the domains property value. The domains property
func (m *ApiOverview) SetDomains(value ApiOverview_domainsable)() {
    m.domains = value
}
// SetInstalledVersion sets the installed_version property value. The installed_version property
func (m *ApiOverview) SetInstalledVersion(value *string)() {
    m.installed_version = value
}
// SetPackages sets the packages property value. The packages property
func (m *ApiOverview) SetPackages(value []string)() {
    m.packages = value
}
// SetVerifiablePasswordAuthentication sets the verifiable_password_authentication property value. The verifiable_password_authentication property
func (m *ApiOverview) SetVerifiablePasswordAuthentication(value *bool)() {
    m.verifiable_password_authentication = value
}
type ApiOverviewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDomains()(ApiOverview_domainsable)
    GetInstalledVersion()(*string)
    GetPackages()([]string)
    GetVerifiablePasswordAuthentication()(*bool)
    SetDomains(value ApiOverview_domainsable)()
    SetInstalledVersion(value *string)()
    SetPackages(value []string)()
    SetVerifiablePasswordAuthentication(value *bool)()
}
