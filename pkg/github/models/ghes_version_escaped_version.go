package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesVersion_version struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The build_date property
    build_date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The build_id property
    build_id *string
    // The platform property
    platform *GhesVersion_version_platform
    // The version property
    version *string
}
// NewGhesVersion_version instantiates a new GhesVersion_version and sets the default values.
func NewGhesVersion_version()(*GhesVersion_version) {
    m := &GhesVersion_version{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesVersion_versionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesVersion_versionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesVersion_version(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesVersion_version) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBuildDate gets the build_date property value. The build_date property
// returns a *DateOnly when successful
func (m *GhesVersion_version) GetBuildDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.build_date
}
// GetBuildId gets the build_id property value. The build_id property
// returns a *string when successful
func (m *GhesVersion_version) GetBuildId()(*string) {
    return m.build_id
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesVersion_version) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["build_date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuildDate(val)
        }
        return nil
    }
    res["build_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuildId(val)
        }
        return nil
    }
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGhesVersion_version_platform)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*GhesVersion_version_platform))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetPlatform gets the platform property value. The platform property
// returns a *GhesVersion_version_platform when successful
func (m *GhesVersion_version) GetPlatform()(*GhesVersion_version_platform) {
    return m.platform
}
// GetVersion gets the version property value. The version property
// returns a *string when successful
func (m *GhesVersion_version) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *GhesVersion_version) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteDateOnlyValue("build_date", m.GetBuildDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("build_id", m.GetBuildId())
        if err != nil {
            return err
        }
    }
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err := writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("version", m.GetVersion())
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
func (m *GhesVersion_version) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBuildDate sets the build_date property value. The build_date property
func (m *GhesVersion_version) SetBuildDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.build_date = value
}
// SetBuildId sets the build_id property value. The build_id property
func (m *GhesVersion_version) SetBuildId(value *string)() {
    m.build_id = value
}
// SetPlatform sets the platform property value. The platform property
func (m *GhesVersion_version) SetPlatform(value *GhesVersion_version_platform)() {
    m.platform = value
}
// SetVersion sets the version property value. The version property
func (m *GhesVersion_version) SetVersion(value *string)() {
    m.version = value
}
type GhesVersion_versionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBuildDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetBuildId()(*string)
    GetPlatform()(*GhesVersion_version_platform)
    GetVersion()(*string)
    SetBuildDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetBuildId(value *string)()
    SetPlatform(value *GhesVersion_version_platform)()
    SetVersion(value *string)()
}
