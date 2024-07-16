package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LicenseInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The days_until_expiration property
    days_until_expiration *int32
    // The expire_at property
    expire_at *string
    // The kind property
    kind *string
    // The seats property
    seats LicenseInfo_LicenseInfo_seatsable
    // The seats_available property
    seats_available LicenseInfo_LicenseInfo_seats_availableable
    // The seats_used property
    seats_used *int32
}
// LicenseInfo_LicenseInfo_seats composed type wrapper for classes int32, string
type LicenseInfo_LicenseInfo_seats struct {
    // Composed type representation for type int32
    integer *int32
    // Composed type representation for type string
    string *string
}
// NewLicenseInfo_LicenseInfo_seats instantiates a new LicenseInfo_LicenseInfo_seats and sets the default values.
func NewLicenseInfo_LicenseInfo_seats()(*LicenseInfo_LicenseInfo_seats) {
    m := &LicenseInfo_LicenseInfo_seats{
    }
    return m
}
// CreateLicenseInfo_LicenseInfo_seatsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLicenseInfo_LicenseInfo_seatsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewLicenseInfo_LicenseInfo_seats()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetInt32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetInteger(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LicenseInfo_LicenseInfo_seats) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetInteger gets the integer property value. Composed type representation for type int32
// returns a *int32 when successful
func (m *LicenseInfo_LicenseInfo_seats) GetInteger()(*int32) {
    return m.integer
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *LicenseInfo_LicenseInfo_seats) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *LicenseInfo_LicenseInfo_seats) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *LicenseInfo_LicenseInfo_seats) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInteger() != nil {
        err := writer.WriteInt32Value("", m.GetInteger())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInteger sets the integer property value. Composed type representation for type int32
func (m *LicenseInfo_LicenseInfo_seats) SetInteger(value *int32)() {
    m.integer = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *LicenseInfo_LicenseInfo_seats) SetString(value *string)() {
    m.string = value
}
// LicenseInfo_LicenseInfo_seats_available composed type wrapper for classes int32, string
type LicenseInfo_LicenseInfo_seats_available struct {
    // Composed type representation for type int32
    integer *int32
    // Composed type representation for type string
    string *string
}
// NewLicenseInfo_LicenseInfo_seats_available instantiates a new LicenseInfo_LicenseInfo_seats_available and sets the default values.
func NewLicenseInfo_LicenseInfo_seats_available()(*LicenseInfo_LicenseInfo_seats_available) {
    m := &LicenseInfo_LicenseInfo_seats_available{
    }
    return m
}
// CreateLicenseInfo_LicenseInfo_seats_availableFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLicenseInfo_LicenseInfo_seats_availableFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewLicenseInfo_LicenseInfo_seats_available()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetInt32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetInteger(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LicenseInfo_LicenseInfo_seats_available) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetInteger gets the integer property value. Composed type representation for type int32
// returns a *int32 when successful
func (m *LicenseInfo_LicenseInfo_seats_available) GetInteger()(*int32) {
    return m.integer
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *LicenseInfo_LicenseInfo_seats_available) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *LicenseInfo_LicenseInfo_seats_available) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *LicenseInfo_LicenseInfo_seats_available) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInteger() != nil {
        err := writer.WriteInt32Value("", m.GetInteger())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInteger sets the integer property value. Composed type representation for type int32
func (m *LicenseInfo_LicenseInfo_seats_available) SetInteger(value *int32)() {
    m.integer = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *LicenseInfo_LicenseInfo_seats_available) SetString(value *string)() {
    m.string = value
}
type LicenseInfo_LicenseInfo_seats_availableable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteger()(*int32)
    GetString()(*string)
    SetInteger(value *int32)()
    SetString(value *string)()
}
type LicenseInfo_LicenseInfo_seatsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteger()(*int32)
    GetString()(*string)
    SetInteger(value *int32)()
    SetString(value *string)()
}
// NewLicenseInfo instantiates a new LicenseInfo and sets the default values.
func NewLicenseInfo()(*LicenseInfo) {
    m := &LicenseInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLicenseInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLicenseInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLicenseInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LicenseInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDaysUntilExpiration gets the days_until_expiration property value. The days_until_expiration property
// returns a *int32 when successful
func (m *LicenseInfo) GetDaysUntilExpiration()(*int32) {
    return m.days_until_expiration
}
// GetExpireAt gets the expire_at property value. The expire_at property
// returns a *string when successful
func (m *LicenseInfo) GetExpireAt()(*string) {
    return m.expire_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LicenseInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["days_until_expiration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDaysUntilExpiration(val)
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
    res["kind"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKind(val)
        }
        return nil
    }
    res["seats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLicenseInfo_LicenseInfo_seatsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeats(val.(LicenseInfo_LicenseInfo_seatsable))
        }
        return nil
    }
    res["seats_available"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLicenseInfo_LicenseInfo_seats_availableFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeatsAvailable(val.(LicenseInfo_LicenseInfo_seats_availableable))
        }
        return nil
    }
    res["seats_used"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeatsUsed(val)
        }
        return nil
    }
    return res
}
// GetKind gets the kind property value. The kind property
// returns a *string when successful
func (m *LicenseInfo) GetKind()(*string) {
    return m.kind
}
// GetSeats gets the seats property value. The seats property
// returns a LicenseInfo_LicenseInfo_seatsable when successful
func (m *LicenseInfo) GetSeats()(LicenseInfo_LicenseInfo_seatsable) {
    return m.seats
}
// GetSeatsAvailable gets the seats_available property value. The seats_available property
// returns a LicenseInfo_LicenseInfo_seats_availableable when successful
func (m *LicenseInfo) GetSeatsAvailable()(LicenseInfo_LicenseInfo_seats_availableable) {
    return m.seats_available
}
// GetSeatsUsed gets the seats_used property value. The seats_used property
// returns a *int32 when successful
func (m *LicenseInfo) GetSeatsUsed()(*int32) {
    return m.seats_used
}
// Serialize serializes information the current object
func (m *LicenseInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("days_until_expiration", m.GetDaysUntilExpiration())
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
        err := writer.WriteStringValue("kind", m.GetKind())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("seats", m.GetSeats())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("seats_available", m.GetSeatsAvailable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("seats_used", m.GetSeatsUsed())
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
func (m *LicenseInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDaysUntilExpiration sets the days_until_expiration property value. The days_until_expiration property
func (m *LicenseInfo) SetDaysUntilExpiration(value *int32)() {
    m.days_until_expiration = value
}
// SetExpireAt sets the expire_at property value. The expire_at property
func (m *LicenseInfo) SetExpireAt(value *string)() {
    m.expire_at = value
}
// SetKind sets the kind property value. The kind property
func (m *LicenseInfo) SetKind(value *string)() {
    m.kind = value
}
// SetSeats sets the seats property value. The seats property
func (m *LicenseInfo) SetSeats(value LicenseInfo_LicenseInfo_seatsable)() {
    m.seats = value
}
// SetSeatsAvailable sets the seats_available property value. The seats_available property
func (m *LicenseInfo) SetSeatsAvailable(value LicenseInfo_LicenseInfo_seats_availableable)() {
    m.seats_available = value
}
// SetSeatsUsed sets the seats_used property value. The seats_used property
func (m *LicenseInfo) SetSeatsUsed(value *int32)() {
    m.seats_used = value
}
type LicenseInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDaysUntilExpiration()(*int32)
    GetExpireAt()(*string)
    GetKind()(*string)
    GetSeats()(LicenseInfo_LicenseInfo_seatsable)
    GetSeatsAvailable()(LicenseInfo_LicenseInfo_seats_availableable)
    GetSeatsUsed()(*int32)
    SetDaysUntilExpiration(value *int32)()
    SetExpireAt(value *string)()
    SetKind(value *string)()
    SetSeats(value LicenseInfo_LicenseInfo_seatsable)()
    SetSeatsAvailable(value LicenseInfo_LicenseInfo_seats_availableable)()
    SetSeatsUsed(value *int32)()
}
