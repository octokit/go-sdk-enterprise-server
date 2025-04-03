package advancedsecurity
type GetAdvanced_security_productQueryParameterType int

const (
    CODE_SECURITY_GETADVANCED_SECURITY_PRODUCTQUERYPARAMETERTYPE GetAdvanced_security_productQueryParameterType = iota
    SECRET_PROTECTION_GETADVANCED_SECURITY_PRODUCTQUERYPARAMETERTYPE
)

func (i GetAdvanced_security_productQueryParameterType) String() string {
    return []string{"code_security", "secret_protection"}[i]
}
func ParseGetAdvanced_security_productQueryParameterType(v string) (any, error) {
    result := CODE_SECURITY_GETADVANCED_SECURITY_PRODUCTQUERYPARAMETERTYPE
    switch v {
        case "code_security":
            result = CODE_SECURITY_GETADVANCED_SECURITY_PRODUCTQUERYPARAMETERTYPE
        case "secret_protection":
            result = SECRET_PROTECTION_GETADVANCED_SECURITY_PRODUCTQUERYPARAMETERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetAdvanced_security_productQueryParameterType(values []GetAdvanced_security_productQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetAdvanced_security_productQueryParameterType) isMultiValue() bool {
    return false
}
