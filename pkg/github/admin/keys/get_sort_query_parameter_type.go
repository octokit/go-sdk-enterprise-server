package keys
import (
    "errors"
)
type GetSortQueryParameterType int

const (
    CREATED_GETSORTQUERYPARAMETERTYPE GetSortQueryParameterType = iota
    UPDATED_GETSORTQUERYPARAMETERTYPE
    ACCESSED_GETSORTQUERYPARAMETERTYPE
)

func (i GetSortQueryParameterType) String() string {
    return []string{"created", "updated", "accessed"}[i]
}
func ParseGetSortQueryParameterType(v string) (any, error) {
    result := CREATED_GETSORTQUERYPARAMETERTYPE
    switch v {
        case "created":
            result = CREATED_GETSORTQUERYPARAMETERTYPE
        case "updated":
            result = UPDATED_GETSORTQUERYPARAMETERTYPE
        case "accessed":
            result = ACCESSED_GETSORTQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetSortQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetSortQueryParameterType(values []GetSortQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetSortQueryParameterType) isMultiValue() bool {
    return false
}
