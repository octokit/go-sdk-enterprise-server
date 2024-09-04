package models
type GhesReplicationStatusIndicator int

const (
    UNKNOWN_GHESREPLICATIONSTATUSINDICATOR GhesReplicationStatusIndicator = iota
    OK_GHESREPLICATIONSTATUSINDICATOR
    WARNING_GHESREPLICATIONSTATUSINDICATOR
    CRITICAL_GHESREPLICATIONSTATUSINDICATOR
)

func (i GhesReplicationStatusIndicator) String() string {
    return []string{"UNKNOWN", "OK", "WARNING", "CRITICAL"}[i]
}
func ParseGhesReplicationStatusIndicator(v string) (any, error) {
    result := UNKNOWN_GHESREPLICATIONSTATUSINDICATOR
    switch v {
        case "UNKNOWN":
            result = UNKNOWN_GHESREPLICATIONSTATUSINDICATOR
        case "OK":
            result = OK_GHESREPLICATIONSTATUSINDICATOR
        case "WARNING":
            result = WARNING_GHESREPLICATIONSTATUSINDICATOR
        case "CRITICAL":
            result = CRITICAL_GHESREPLICATIONSTATUSINDICATOR
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGhesReplicationStatusIndicator(values []GhesReplicationStatusIndicator) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GhesReplicationStatusIndicator) isMultiValue() bool {
    return false
}
