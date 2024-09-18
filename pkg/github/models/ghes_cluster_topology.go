package models
type GhesClusterTopology int

const (
    SINGLENODE_GHESCLUSTERTOPOLOGY GhesClusterTopology = iota
    HA_GHESCLUSTERTOPOLOGY
    CLUSTER_GHESCLUSTERTOPOLOGY
)

func (i GhesClusterTopology) String() string {
    return []string{"SingleNode", "Ha", "Cluster"}[i]
}
func ParseGhesClusterTopology(v string) (any, error) {
    result := SINGLENODE_GHESCLUSTERTOPOLOGY
    switch v {
        case "SingleNode":
            result = SINGLENODE_GHESCLUSTERTOPOLOGY
        case "Ha":
            result = HA_GHESCLUSTERTOPOLOGY
        case "Cluster":
            result = CLUSTER_GHESCLUSTERTOPOLOGY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGhesClusterTopology(values []GhesClusterTopology) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GhesClusterTopology) isMultiValue() bool {
    return false
}
