package models
type GhesGetMaintenance_status int

const (
    ON_GHESGETMAINTENANCE_STATUS GhesGetMaintenance_status = iota
    OFF_GHESGETMAINTENANCE_STATUS
    SCHEDULED_GHESGETMAINTENANCE_STATUS
)

func (i GhesGetMaintenance_status) String() string {
    return []string{"on", "off", "scheduled"}[i]
}
func ParseGhesGetMaintenance_status(v string) (any, error) {
    result := ON_GHESGETMAINTENANCE_STATUS
    switch v {
        case "on":
            result = ON_GHESGETMAINTENANCE_STATUS
        case "off":
            result = OFF_GHESGETMAINTENANCE_STATUS
        case "scheduled":
            result = SCHEDULED_GHESGETMAINTENANCE_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGhesGetMaintenance_status(values []GhesGetMaintenance_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GhesGetMaintenance_status) isMultiValue() bool {
    return false
}
