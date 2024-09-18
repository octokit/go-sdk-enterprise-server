package repositories
type GetVisibilityQueryParameterType int

const (
    ALL_GETVISIBILITYQUERYPARAMETERTYPE GetVisibilityQueryParameterType = iota
    PUBLIC_GETVISIBILITYQUERYPARAMETERTYPE
)

func (i GetVisibilityQueryParameterType) String() string {
    return []string{"all", "public"}[i]
}
func ParseGetVisibilityQueryParameterType(v string) (any, error) {
    result := ALL_GETVISIBILITYQUERYPARAMETERTYPE
    switch v {
        case "all":
            result = ALL_GETVISIBILITYQUERYPARAMETERTYPE
        case "public":
            result = PUBLIC_GETVISIBILITYQUERYPARAMETERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetVisibilityQueryParameterType(values []GetVisibilityQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetVisibilityQueryParameterType) isMultiValue() bool {
    return false
}
