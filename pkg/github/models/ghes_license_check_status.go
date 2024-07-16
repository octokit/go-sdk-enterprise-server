package models
import (
    "errors"
)
// The license status of the currently active enterprise license.
type GhesLicenseCheck_status int

const (
    VALID_GHESLICENSECHECK_STATUS GhesLicenseCheck_status = iota
    INVALID_GHESLICENSECHECK_STATUS
    EXPIRED_GHESLICENSECHECK_STATUS
    CLUSTERMODENOTSUPPORTED_GHESLICENSECHECK_STATUS
)

func (i GhesLicenseCheck_status) String() string {
    return []string{"valid", "invalid", "expired", "cluster mode not supported"}[i]
}
func ParseGhesLicenseCheck_status(v string) (any, error) {
    result := VALID_GHESLICENSECHECK_STATUS
    switch v {
        case "valid":
            result = VALID_GHESLICENSECHECK_STATUS
        case "invalid":
            result = INVALID_GHESLICENSECHECK_STATUS
        case "expired":
            result = EXPIRED_GHESLICENSECHECK_STATUS
        case "cluster mode not supported":
            result = CLUSTERMODENOTSUPPORTED_GHESLICENSECHECK_STATUS
        default:
            return 0, errors.New("Unknown GhesLicenseCheck_status value: " + v)
    }
    return &result, nil
}
func SerializeGhesLicenseCheck_status(values []GhesLicenseCheck_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GhesLicenseCheck_status) isMultiValue() bool {
    return false
}
