package item
import (
    "errors"
)
// The state of enforcement for the hook on this repository.
type WithPre_receive_hook_PatchRequestBody_enforcement int

const (
    ENABLED_WITHPRE_RECEIVE_HOOK_PATCHREQUESTBODY_ENFORCEMENT WithPre_receive_hook_PatchRequestBody_enforcement = iota
    DISABLED_WITHPRE_RECEIVE_HOOK_PATCHREQUESTBODY_ENFORCEMENT
    TESTING_WITHPRE_RECEIVE_HOOK_PATCHREQUESTBODY_ENFORCEMENT
)

func (i WithPre_receive_hook_PatchRequestBody_enforcement) String() string {
    return []string{"enabled", "disabled", "testing"}[i]
}
func ParseWithPre_receive_hook_PatchRequestBody_enforcement(v string) (any, error) {
    result := ENABLED_WITHPRE_RECEIVE_HOOK_PATCHREQUESTBODY_ENFORCEMENT
    switch v {
        case "enabled":
            result = ENABLED_WITHPRE_RECEIVE_HOOK_PATCHREQUESTBODY_ENFORCEMENT
        case "disabled":
            result = DISABLED_WITHPRE_RECEIVE_HOOK_PATCHREQUESTBODY_ENFORCEMENT
        case "testing":
            result = TESTING_WITHPRE_RECEIVE_HOOK_PATCHREQUESTBODY_ENFORCEMENT
        default:
            return 0, errors.New("Unknown WithPre_receive_hook_PatchRequestBody_enforcement value: " + v)
    }
    return &result, nil
}
func SerializeWithPre_receive_hook_PatchRequestBody_enforcement(values []WithPre_receive_hook_PatchRequestBody_enforcement) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithPre_receive_hook_PatchRequestBody_enforcement) isMultiValue() bool {
    return false
}
