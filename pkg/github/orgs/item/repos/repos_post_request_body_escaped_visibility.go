package repos
// The visibility of the repository. **Note**: For GitHub Enterprise Server, this endpoint will only list repositories available to all users on the enterprise. For more information, see "[Creating an internal repository](https://docs.github.com/enterprise-server@3.14/github/creating-cloning-and-archiving-repositories/about-repository-visibility#about-internal-repositories)" in the GitHub Help documentation.  The `visibility` parameter overrides the `private` parameter when you use both parameters with the `nebula-preview` preview header.
type ReposPostRequestBody_visibility int

const (
    PUBLIC_REPOSPOSTREQUESTBODY_VISIBILITY ReposPostRequestBody_visibility = iota
    PRIVATE_REPOSPOSTREQUESTBODY_VISIBILITY
    INTERNAL_REPOSPOSTREQUESTBODY_VISIBILITY
)

func (i ReposPostRequestBody_visibility) String() string {
    return []string{"public", "private", "internal"}[i]
}
func ParseReposPostRequestBody_visibility(v string) (any, error) {
    result := PUBLIC_REPOSPOSTREQUESTBODY_VISIBILITY
    switch v {
        case "public":
            result = PUBLIC_REPOSPOSTREQUESTBODY_VISIBILITY
        case "private":
            result = PRIVATE_REPOSPOSTREQUESTBODY_VISIBILITY
        case "internal":
            result = INTERNAL_REPOSPOSTREQUESTBODY_VISIBILITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeReposPostRequestBody_visibility(values []ReposPostRequestBody_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ReposPostRequestBody_visibility) isMultiValue() bool {
    return false
}
