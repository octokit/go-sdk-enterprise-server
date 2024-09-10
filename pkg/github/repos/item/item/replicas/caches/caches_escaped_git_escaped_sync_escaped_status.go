package caches
type Caches_git_sync_status int

const (
    OFFLINE_CACHES_GIT_SYNC_STATUS Caches_git_sync_status = iota
    INACTIVE_CACHES_GIT_SYNC_STATUS
    IN_SYNC_CACHES_GIT_SYNC_STATUS
    NOT_IN_SYNC_CACHES_GIT_SYNC_STATUS
)

func (i Caches_git_sync_status) String() string {
    return []string{"offline", "inactive", "in_sync", "not_in_sync"}[i]
}
func ParseCaches_git_sync_status(v string) (any, error) {
    result := OFFLINE_CACHES_GIT_SYNC_STATUS
    switch v {
        case "offline":
            result = OFFLINE_CACHES_GIT_SYNC_STATUS
        case "inactive":
            result = INACTIVE_CACHES_GIT_SYNC_STATUS
        case "in_sync":
            result = IN_SYNC_CACHES_GIT_SYNC_STATUS
        case "not_in_sync":
            result = NOT_IN_SYNC_CACHES_GIT_SYNC_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCaches_git_sync_status(values []Caches_git_sync_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Caches_git_sync_status) isMultiValue() bool {
    return false
}
