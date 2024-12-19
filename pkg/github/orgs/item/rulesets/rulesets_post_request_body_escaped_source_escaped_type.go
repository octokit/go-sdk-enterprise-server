package rulesets
// The type of the source of the ruleset
type RulesetsPostRequestBody_source_type int

const (
    REPOSITORY_RULESETSPOSTREQUESTBODY_SOURCE_TYPE RulesetsPostRequestBody_source_type = iota
    ORGANIZATION_RULESETSPOSTREQUESTBODY_SOURCE_TYPE
)

func (i RulesetsPostRequestBody_source_type) String() string {
    return []string{"Repository", "Organization"}[i]
}
func ParseRulesetsPostRequestBody_source_type(v string) (any, error) {
    result := REPOSITORY_RULESETSPOSTREQUESTBODY_SOURCE_TYPE
    switch v {
        case "Repository":
            result = REPOSITORY_RULESETSPOSTREQUESTBODY_SOURCE_TYPE
        case "Organization":
            result = ORGANIZATION_RULESETSPOSTREQUESTBODY_SOURCE_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRulesetsPostRequestBody_source_type(values []RulesetsPostRequestBody_source_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RulesetsPostRequestBody_source_type) isMultiValue() bool {
    return false
}
