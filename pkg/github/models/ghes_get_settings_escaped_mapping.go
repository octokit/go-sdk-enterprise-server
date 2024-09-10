package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GhesGetSettings_mapping struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The basemap property
    basemap *string
    // The enabled property
    enabled *bool
    // The tileserver property
    tileserver *string
    // The token property
    token *string
}
// NewGhesGetSettings_mapping instantiates a new GhesGetSettings_mapping and sets the default values.
func NewGhesGetSettings_mapping()(*GhesGetSettings_mapping) {
    m := &GhesGetSettings_mapping{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGhesGetSettings_mappingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGhesGetSettings_mappingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGhesGetSettings_mapping(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GhesGetSettings_mapping) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBasemap gets the basemap property value. The basemap property
// returns a *string when successful
func (m *GhesGetSettings_mapping) GetBasemap()(*string) {
    return m.basemap
}
// GetEnabled gets the enabled property value. The enabled property
// returns a *bool when successful
func (m *GhesGetSettings_mapping) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GhesGetSettings_mapping) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["basemap"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBasemap(val)
        }
        return nil
    }
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["tileserver"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTileserver(val)
        }
        return nil
    }
    res["token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToken(val)
        }
        return nil
    }
    return res
}
// GetTileserver gets the tileserver property value. The tileserver property
// returns a *string when successful
func (m *GhesGetSettings_mapping) GetTileserver()(*string) {
    return m.tileserver
}
// GetToken gets the token property value. The token property
// returns a *string when successful
func (m *GhesGetSettings_mapping) GetToken()(*string) {
    return m.token
}
// Serialize serializes information the current object
func (m *GhesGetSettings_mapping) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("basemap", m.GetBasemap())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tileserver", m.GetTileserver())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("token", m.GetToken())
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
func (m *GhesGetSettings_mapping) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBasemap sets the basemap property value. The basemap property
func (m *GhesGetSettings_mapping) SetBasemap(value *string)() {
    m.basemap = value
}
// SetEnabled sets the enabled property value. The enabled property
func (m *GhesGetSettings_mapping) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetTileserver sets the tileserver property value. The tileserver property
func (m *GhesGetSettings_mapping) SetTileserver(value *string)() {
    m.tileserver = value
}
// SetToken sets the token property value. The token property
func (m *GhesGetSettings_mapping) SetToken(value *string)() {
    m.token = value
}
type GhesGetSettings_mappingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBasemap()(*string)
    GetEnabled()(*bool)
    GetTileserver()(*string)
    GetToken()(*string)
    SetBasemap(value *string)()
    SetEnabled(value *bool)()
    SetTileserver(value *string)()
    SetToken(value *string)()
}
