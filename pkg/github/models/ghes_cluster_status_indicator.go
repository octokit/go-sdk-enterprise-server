package models
type GhesClusterStatusIndicator int

const (
    UNKNOWN_GHESCLUSTERSTATUSINDICATOR GhesClusterStatusIndicator = iota
    OK_GHESCLUSTERSTATUSINDICATOR
    WARNING_GHESCLUSTERSTATUSINDICATOR
    CRITICAL_GHESCLUSTERSTATUSINDICATOR
)

func (i GhesClusterStatusIndicator) String() string {
    return []string{"UNKNOWN", "OK", "WARNING", "CRITICAL"}[i]
}
func ParseGhesClusterStatusIndicator(v string) (any, error) {
    result := UNKNOWN_GHESCLUSTERSTATUSINDICATOR
    switch v {
        case "UNKNOWN":
            result = UNKNOWN_GHESCLUSTERSTATUSINDICATOR
        case "OK":
            result = OK_GHESCLUSTERSTATUSINDICATOR
        case "WARNING":
            result = WARNING_GHESCLUSTERSTATUSINDICATOR
        case "CRITICAL":
            result = CRITICAL_GHESCLUSTERSTATUSINDICATOR
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGhesClusterStatusIndicator(values []GhesClusterStatusIndicator) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GhesClusterStatusIndicator) isMultiValue() bool {
    return false
}
