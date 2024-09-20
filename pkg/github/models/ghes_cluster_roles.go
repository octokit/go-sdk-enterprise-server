package models
type GhesClusterRoles int

const (
    BLANK_GHESCLUSTERROLES GhesClusterRoles = iota
    ACTIONSSERVER_GHESCLUSTERROLES
    CONSULSERVER_GHESCLUSTERROLES
    ELASTICSEARCHSERVER_GHESCLUSTERROLES
    GITSERVER_GHESCLUSTERROLES
    JOBSERVER_GHESCLUSTERROLES
    LAUNCHSERVER_GHESCLUSTERROLES
    MEMCACHESERVER_GHESCLUSTERROLES
    METRICSSERVER_GHESCLUSTERROLES
    MSSQLSERVER_GHESCLUSTERROLES
    MYSQLSERVER_GHESCLUSTERROLES
    PAGESSERVER_GHESCLUSTERROLES
    REDISSERVER_GHESCLUSTERROLES
    STORAGESERVER_GHESCLUSTERROLES
    WEBSERVER_GHESCLUSTERROLES
)

func (i GhesClusterRoles) String() string {
    return []string{"Blank", "ActionsServer", "ConsulServer", "ElasticsearchServer", "GitServer", "JobServer", "LaunchServer", "MemcacheServer", "MetricsServer", "MssqlServer", "MysqlServer", "PagesServer", "RedisServer", "StorageServer", "WebServer"}[i]
}
func ParseGhesClusterRoles(v string) (any, error) {
    result := BLANK_GHESCLUSTERROLES
    switch v {
        case "Blank":
            result = BLANK_GHESCLUSTERROLES
        case "ActionsServer":
            result = ACTIONSSERVER_GHESCLUSTERROLES
        case "ConsulServer":
            result = CONSULSERVER_GHESCLUSTERROLES
        case "ElasticsearchServer":
            result = ELASTICSEARCHSERVER_GHESCLUSTERROLES
        case "GitServer":
            result = GITSERVER_GHESCLUSTERROLES
        case "JobServer":
            result = JOBSERVER_GHESCLUSTERROLES
        case "LaunchServer":
            result = LAUNCHSERVER_GHESCLUSTERROLES
        case "MemcacheServer":
            result = MEMCACHESERVER_GHESCLUSTERROLES
        case "MetricsServer":
            result = METRICSSERVER_GHESCLUSTERROLES
        case "MssqlServer":
            result = MSSQLSERVER_GHESCLUSTERROLES
        case "MysqlServer":
            result = MYSQLSERVER_GHESCLUSTERROLES
        case "PagesServer":
            result = PAGESSERVER_GHESCLUSTERROLES
        case "RedisServer":
            result = REDISSERVER_GHESCLUSTERROLES
        case "StorageServer":
            result = STORAGESERVER_GHESCLUSTERROLES
        case "WebServer":
            result = WEBSERVER_GHESCLUSTERROLES
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGhesClusterRoles(values []GhesClusterRoles) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GhesClusterRoles) isMultiValue() bool {
    return false
}
