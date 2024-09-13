package models
type GhesVersion_version_platform int

const (
    AMI_GHESVERSION_VERSION_PLATFORM GhesVersion_version_platform = iota
    AZURE_GHESVERSION_VERSION_PLATFORM
    ESX_GHESVERSION_VERSION_PLATFORM
    GCE_GHESVERSION_VERSION_PLATFORM
    HYPERV_GHESVERSION_VERSION_PLATFORM
    KVM_GHESVERSION_VERSION_PLATFORM
)

func (i GhesVersion_version_platform) String() string {
    return []string{"ami", "azure", "esx", "gce", "hyperv", "kvm"}[i]
}
func ParseGhesVersion_version_platform(v string) (any, error) {
    result := AMI_GHESVERSION_VERSION_PLATFORM
    switch v {
        case "ami":
            result = AMI_GHESVERSION_VERSION_PLATFORM
        case "azure":
            result = AZURE_GHESVERSION_VERSION_PLATFORM
        case "esx":
            result = ESX_GHESVERSION_VERSION_PLATFORM
        case "gce":
            result = GCE_GHESVERSION_VERSION_PLATFORM
        case "hyperv":
            result = HYPERV_GHESVERSION_VERSION_PLATFORM
        case "kvm":
            result = KVM_GHESVERSION_VERSION_PLATFORM
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGhesVersion_version_platform(values []GhesVersion_version_platform) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GhesVersion_version_platform) isMultiValue() bool {
    return false
}
