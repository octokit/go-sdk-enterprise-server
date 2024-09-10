package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MigrationsWithMigration_ItemRequestBuilder builds and executes requests for operations under \user\migrations\{migration_id}
type MigrationsWithMigration_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Archive the archive property
// returns a *MigrationsItemArchiveRequestBuilder when successful
func (m *MigrationsWithMigration_ItemRequestBuilder) Archive()(*MigrationsItemArchiveRequestBuilder) {
    return NewMigrationsItemArchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMigrationsWithMigration_ItemRequestBuilderInternal instantiates a new MigrationsWithMigration_ItemRequestBuilder and sets the default values.
func NewMigrationsWithMigration_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MigrationsWithMigration_ItemRequestBuilder) {
    m := &MigrationsWithMigration_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/migrations/{migration_id}", pathParameters),
    }
    return m
}
// NewMigrationsWithMigration_ItemRequestBuilder instantiates a new MigrationsWithMigration_ItemRequestBuilder and sets the default values.
func NewMigrationsWithMigration_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MigrationsWithMigration_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMigrationsWithMigration_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Repositories the repositories property
// returns a *MigrationsItemRepositoriesRequestBuilder when successful
func (m *MigrationsWithMigration_ItemRequestBuilder) Repositories()(*MigrationsItemRepositoriesRequestBuilder) {
    return NewMigrationsItemRepositoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
