package models
import (
    "errors"
)
type GhesChecksResultIndicator int

const (
    OK_GHESCHECKSRESULTINDICATOR GhesChecksResultIndicator = iota
    FAILED_GHESCHECKSRESULTINDICATOR
)

func (i GhesChecksResultIndicator) String() string {
    return []string{"OK", "FAILED"}[i]
}
func ParseGhesChecksResultIndicator(v string) (any, error) {
    result := OK_GHESCHECKSRESULTINDICATOR
    switch v {
        case "OK":
            result = OK_GHESCHECKSRESULTINDICATOR
        case "FAILED":
            result = FAILED_GHESCHECKSRESULTINDICATOR
        default:
            return 0, errors.New("Unknown GhesChecksResultIndicator value: " + v)
    }
    return &result, nil
}
func SerializeGhesChecksResultIndicator(values []GhesChecksResultIndicator) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GhesChecksResultIndicator) isMultiValue() bool {
    return false
}
