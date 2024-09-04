package github

import (
    "context"
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i019b368180824531ee10a9fabd22410b95ccd81f7e1b1f8416c99a317b24aff0 "github.com/octokit/go-sdk-enterprise-server/pkg/github/search"
    i18df8a9d03a66fdd1437c9a40ea0f9e0a7c520e50037e2c547caa46c8d57621c "github.com/octokit/go-sdk-enterprise-server/pkg/github/applications"
    i1920a84396ec87b264b65c3d4feffd636cd814025614ed8fc6837257346128db "github.com/octokit/go-sdk-enterprise-server/pkg/github/user"
    i19256e8e3069f58ac9e1484a0b995f1de29bf1316e48c80e893c8e451770516d "github.com/octokit/go-sdk-enterprise-server/pkg/github/admin"
    i1da9ff71b80748c2ffa3444d6d942401f510a5d59e7d0b85fba3f675a84cd93a "github.com/octokit/go-sdk-enterprise-server/pkg/github/setup"
    i204bf8331d9b4bbcba744334dc5275abce98054b6b9396febb5ec4627553393a "github.com/octokit/go-sdk-enterprise-server/pkg/github/installation"
    i22bad5e37e0fdcc5f6386ac06d16691f9e070a95085db42b968fccb55bb22a62 "github.com/octokit/go-sdk-enterprise-server/pkg/github/repos"
    i242b26146724e60c7c0c92e35852ec04f99105c6d9e5042280fcf7b36048b9de "github.com/octokit/go-sdk-enterprise-server/pkg/github/enterprise"
    i2553c08de368b7330dccbb084609f22b761d7062044d70c34052481b283f5b9a "github.com/octokit/go-sdk-enterprise-server/pkg/github/projects"
    i2ae21604fdbb7f04d11551ec8fc6db64b2d62842c0aafe38bc4bb2acbbcc1166 "github.com/octokit/go-sdk-enterprise-server/pkg/github/networks"
    i2c3692ea706dc939eb1fccbfa1f444338ef9d1dea009a854d1a90faa7e1f24f3 "github.com/octokit/go-sdk-enterprise-server/pkg/github/events"
    i390649da93a8a4e2b2f499d2309963d032a52c104175e7dcd881d14d0ca62db2 "github.com/octokit/go-sdk-enterprise-server/pkg/github/gitignore"
    i39384a1e70757d633c2ee950921fb9439950b78857111cb2dbd68b33f9d8830c "github.com/octokit/go-sdk-enterprise-server/pkg/github/advisories"
    i435ec22ab9e6920515f1fb3655f10415d71afd100ccd097f650723a910916999 "github.com/octokit/go-sdk-enterprise-server/pkg/github/meta"
    i4476ec40ff84502e96456ccdf4d8f7a9e12c4834106e8ad35b4847c7cb8ea9f4 "github.com/octokit/go-sdk-enterprise-server/pkg/github/gists"
    i46a0b04ed3ba56ba62fbc553ce7cfe242cedfc763368fc6fff06a19d3cbe08c2 "github.com/octokit/go-sdk-enterprise-server/pkg/github/orgs"
    i486611744201264ade6725e5c8b2be0c591c652d6fc7fde3bb8f77e6aa092114 "github.com/octokit/go-sdk-enterprise-server/pkg/github/markdown"
    i49aedbb6755a36d523c6f1f052735019f327e57cdd3408684ba45f2e885bd980 "github.com/octokit/go-sdk-enterprise-server/pkg/github/emojis"
    i516aafeb513add23cb0b85066636cf037269fe65da463faa134b8956732972ba "github.com/octokit/go-sdk-enterprise-server/pkg/github/octocat"
    i673f92a3a3e280895e22213d6efe43f261bc8c75f7ebd72e9afaba3d99f8a4df "github.com/octokit/go-sdk-enterprise-server/pkg/github/authorizations"
    i802736a68ae4a9f1242e056c805d91c3644e113ac0ed67708a7aafae694512f3 "github.com/octokit/go-sdk-enterprise-server/pkg/github/licenses"
    i88fc4f9ae1e68a7b479fa733a4c07fcea06906379a8f840f4d9d0356a0bfcb43 "github.com/octokit/go-sdk-enterprise-server/pkg/github/appmanifests"
    ia6bb22cd55a47079901290dbf3289a756984a02dbba6d3c8afe677eb2abcb352 "github.com/octokit/go-sdk-enterprise-server/pkg/github/apps"
    iab4fa88f4cfb41b10ba8be6db9ff835cabee88e1b3ebd8ca2a24889bd9a2bb66 "github.com/octokit/go-sdk-enterprise-server/pkg/github/teams"
    iac93f8f095011ae884edfdf015eb4bd77cd2c4073c9001899baf6e5394f67d26 "github.com/octokit/go-sdk-enterprise-server/pkg/github/issues"
    ib43f1d21bad1eaccb3220a6895cc4da70307341b670b60c2aa631d2995e47b3e "github.com/octokit/go-sdk-enterprise-server/pkg/github/rate_limit"
    ibd0b9c8a18deb92eb077cd9e715af68e5f543021a34f53d8c5ece9d4e750477c "github.com/octokit/go-sdk-enterprise-server/pkg/github/feeds"
    ic673ef3b555c335d84c8f292c04b9634308a76034507b576e75efb6aa8060af4 "github.com/octokit/go-sdk-enterprise-server/pkg/github/manage"
    icbde3b6cc6b306eaac2200351e57ca9a3b13479fa37b08baaeb4b17445835c36 "github.com/octokit/go-sdk-enterprise-server/pkg/github/scim"
    id285ec6c22949b7f657f9a7661f6e6849ef305c2114d4e0a5f1e0c0ce386d17a "github.com/octokit/go-sdk-enterprise-server/pkg/github/notifications"
    idaa4e03303c4f6a0851c16c06625aac826d0830582346e4b614c29f34eacbd22 "github.com/octokit/go-sdk-enterprise-server/pkg/github/enterprises"
    ideee9aaead60f0751b30729fcf84e07ef70a2f387455613efaccc0a48908813a "github.com/octokit/go-sdk-enterprise-server/pkg/github/zen"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
    ie2d2dc6169f03de2523572263aecda4690557f5da4d9eb8a3e4ad057f1e00609 "github.com/octokit/go-sdk-enterprise-server/pkg/github/app"
    ie80ab2c4f62921be85b8d011cd12092ac3cfcc2e38bbf97c34edad846cd82460 "github.com/octokit/go-sdk-enterprise-server/pkg/github/organizations"
    ieb3561664edbcd472a4089a7b601c05b4d7de53178c12aa3e14e0d6f1fcb91a7 "github.com/octokit/go-sdk-enterprise-server/pkg/github/repositories"
    if155ce7328d65c6829538f8121c72e07753ee85681edf8f539ba03b35a5b2a44 "github.com/octokit/go-sdk-enterprise-server/pkg/github/users"
    if280428a0c8dabdeeed4742ac65d1e1798c574dfe8b7088e61badb601f707afa "github.com/octokit/go-sdk-enterprise-server/pkg/github/codes_of_conduct"
)

// ApiClient the main entry point of the SDK, exposes the configuration and the fluent API.
type ApiClient struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Admin the admin property
// returns a *AdminRequestBuilder when successful
func (m *ApiClient) Admin()(*i19256e8e3069f58ac9e1484a0b995f1de29bf1316e48c80e893c8e451770516d.AdminRequestBuilder) {
    return i19256e8e3069f58ac9e1484a0b995f1de29bf1316e48c80e893c8e451770516d.NewAdminRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Advisories the advisories property
// returns a *AdvisoriesRequestBuilder when successful
func (m *ApiClient) Advisories()(*i39384a1e70757d633c2ee950921fb9439950b78857111cb2dbd68b33f9d8830c.AdvisoriesRequestBuilder) {
    return i39384a1e70757d633c2ee950921fb9439950b78857111cb2dbd68b33f9d8830c.NewAdvisoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// App the app property
// returns a *AppRequestBuilder when successful
func (m *ApiClient) App()(*ie2d2dc6169f03de2523572263aecda4690557f5da4d9eb8a3e4ad057f1e00609.AppRequestBuilder) {
    return ie2d2dc6169f03de2523572263aecda4690557f5da4d9eb8a3e4ad057f1e00609.NewAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Applications the applications property
// returns a *ApplicationsRequestBuilder when successful
func (m *ApiClient) Applications()(*i18df8a9d03a66fdd1437c9a40ea0f9e0a7c520e50037e2c547caa46c8d57621c.ApplicationsRequestBuilder) {
    return i18df8a9d03a66fdd1437c9a40ea0f9e0a7c520e50037e2c547caa46c8d57621c.NewApplicationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppManifests the appManifests property
// returns a *AppManifestsRequestBuilder when successful
func (m *ApiClient) AppManifests()(*i88fc4f9ae1e68a7b479fa733a4c07fcea06906379a8f840f4d9d0356a0bfcb43.AppManifestsRequestBuilder) {
    return i88fc4f9ae1e68a7b479fa733a4c07fcea06906379a8f840f4d9d0356a0bfcb43.NewAppManifestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Apps the apps property
// returns a *AppsRequestBuilder when successful
func (m *ApiClient) Apps()(*ia6bb22cd55a47079901290dbf3289a756984a02dbba6d3c8afe677eb2abcb352.AppsRequestBuilder) {
    return ia6bb22cd55a47079901290dbf3289a756984a02dbba6d3c8afe677eb2abcb352.NewAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Authorizations the authorizations property
// returns a *AuthorizationsRequestBuilder when successful
func (m *ApiClient) Authorizations()(*i673f92a3a3e280895e22213d6efe43f261bc8c75f7ebd72e9afaba3d99f8a4df.AuthorizationsRequestBuilder) {
    return i673f92a3a3e280895e22213d6efe43f261bc8c75f7ebd72e9afaba3d99f8a4df.NewAuthorizationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Codes_of_conduct the codes_of_conduct property
// returns a *Codes_of_conductRequestBuilder when successful
func (m *ApiClient) Codes_of_conduct()(*if280428a0c8dabdeeed4742ac65d1e1798c574dfe8b7088e61badb601f707afa.Codes_of_conductRequestBuilder) {
    return if280428a0c8dabdeeed4742ac65d1e1798c574dfe8b7088e61badb601f707afa.NewCodes_of_conductRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewApiClient instantiates a new ApiClient and sets the default values.
func NewApiClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiClient) {
    m := &ApiClient{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}", map[string]string{}),
    }
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426.NewMultipartSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.BaseRequestBuilder.RequestAdapter.GetBaseUrl() == "" {
        m.BaseRequestBuilder.RequestAdapter.SetBaseUrl("{protocol}://{hostname}/api/v3")
    }
    m.BaseRequestBuilder.PathParameters["baseurl"] = m.BaseRequestBuilder.RequestAdapter.GetBaseUrl()
    return m
}
// Emojis the emojis property
// returns a *EmojisRequestBuilder when successful
func (m *ApiClient) Emojis()(*i49aedbb6755a36d523c6f1f052735019f327e57cdd3408684ba45f2e885bd980.EmojisRequestBuilder) {
    return i49aedbb6755a36d523c6f1f052735019f327e57cdd3408684ba45f2e885bd980.NewEmojisRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Enterprise the enterprise property
// returns a *EnterpriseRequestBuilder when successful
func (m *ApiClient) Enterprise()(*i242b26146724e60c7c0c92e35852ec04f99105c6d9e5042280fcf7b36048b9de.EnterpriseRequestBuilder) {
    return i242b26146724e60c7c0c92e35852ec04f99105c6d9e5042280fcf7b36048b9de.NewEnterpriseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Enterprises the enterprises property
// returns a *EnterprisesRequestBuilder when successful
func (m *ApiClient) Enterprises()(*idaa4e03303c4f6a0851c16c06625aac826d0830582346e4b614c29f34eacbd22.EnterprisesRequestBuilder) {
    return idaa4e03303c4f6a0851c16c06625aac826d0830582346e4b614c29f34eacbd22.NewEnterprisesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Events the events property
// returns a *EventsRequestBuilder when successful
func (m *ApiClient) Events()(*i2c3692ea706dc939eb1fccbfa1f444338ef9d1dea009a854d1a90faa7e1f24f3.EventsRequestBuilder) {
    return i2c3692ea706dc939eb1fccbfa1f444338ef9d1dea009a854d1a90faa7e1f24f3.NewEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Feeds the feeds property
// returns a *FeedsRequestBuilder when successful
func (m *ApiClient) Feeds()(*ibd0b9c8a18deb92eb077cd9e715af68e5f543021a34f53d8c5ece9d4e750477c.FeedsRequestBuilder) {
    return ibd0b9c8a18deb92eb077cd9e715af68e5f543021a34f53d8c5ece9d4e750477c.NewFeedsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get Hypermedia links to resources accessible in GitHub's REST API
// returns a Rootable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-server@3.14/rest/meta/meta#github-api-root
func (m *ApiClient) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Rootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateRootFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.Rootable), nil
}
// Gists the gists property
// returns a *GistsRequestBuilder when successful
func (m *ApiClient) Gists()(*i4476ec40ff84502e96456ccdf4d8f7a9e12c4834106e8ad35b4847c7cb8ea9f4.GistsRequestBuilder) {
    return i4476ec40ff84502e96456ccdf4d8f7a9e12c4834106e8ad35b4847c7cb8ea9f4.NewGistsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Gitignore the gitignore property
// returns a *GitignoreRequestBuilder when successful
func (m *ApiClient) Gitignore()(*i390649da93a8a4e2b2f499d2309963d032a52c104175e7dcd881d14d0ca62db2.GitignoreRequestBuilder) {
    return i390649da93a8a4e2b2f499d2309963d032a52c104175e7dcd881d14d0ca62db2.NewGitignoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Installation the installation property
// returns a *InstallationRequestBuilder when successful
func (m *ApiClient) Installation()(*i204bf8331d9b4bbcba744334dc5275abce98054b6b9396febb5ec4627553393a.InstallationRequestBuilder) {
    return i204bf8331d9b4bbcba744334dc5275abce98054b6b9396febb5ec4627553393a.NewInstallationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Issues the issues property
// returns a *IssuesRequestBuilder when successful
func (m *ApiClient) Issues()(*iac93f8f095011ae884edfdf015eb4bd77cd2c4073c9001899baf6e5394f67d26.IssuesRequestBuilder) {
    return iac93f8f095011ae884edfdf015eb4bd77cd2c4073c9001899baf6e5394f67d26.NewIssuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Licenses the licenses property
// returns a *LicensesRequestBuilder when successful
func (m *ApiClient) Licenses()(*i802736a68ae4a9f1242e056c805d91c3644e113ac0ed67708a7aafae694512f3.LicensesRequestBuilder) {
    return i802736a68ae4a9f1242e056c805d91c3644e113ac0ed67708a7aafae694512f3.NewLicensesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Manage the manage property
// returns a *ManageRequestBuilder when successful
func (m *ApiClient) Manage()(*ic673ef3b555c335d84c8f292c04b9634308a76034507b576e75efb6aa8060af4.ManageRequestBuilder) {
    return ic673ef3b555c335d84c8f292c04b9634308a76034507b576e75efb6aa8060af4.NewManageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Markdown the markdown property
// returns a *MarkdownRequestBuilder when successful
func (m *ApiClient) Markdown()(*i486611744201264ade6725e5c8b2be0c591c652d6fc7fde3bb8f77e6aa092114.MarkdownRequestBuilder) {
    return i486611744201264ade6725e5c8b2be0c591c652d6fc7fde3bb8f77e6aa092114.NewMarkdownRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Meta the meta property
// returns a *MetaRequestBuilder when successful
func (m *ApiClient) Meta()(*i435ec22ab9e6920515f1fb3655f10415d71afd100ccd097f650723a910916999.MetaRequestBuilder) {
    return i435ec22ab9e6920515f1fb3655f10415d71afd100ccd097f650723a910916999.NewMetaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Networks the networks property
// returns a *NetworksRequestBuilder when successful
func (m *ApiClient) Networks()(*i2ae21604fdbb7f04d11551ec8fc6db64b2d62842c0aafe38bc4bb2acbbcc1166.NetworksRequestBuilder) {
    return i2ae21604fdbb7f04d11551ec8fc6db64b2d62842c0aafe38bc4bb2acbbcc1166.NewNetworksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Notifications the notifications property
// returns a *NotificationsRequestBuilder when successful
func (m *ApiClient) Notifications()(*id285ec6c22949b7f657f9a7661f6e6849ef305c2114d4e0a5f1e0c0ce386d17a.NotificationsRequestBuilder) {
    return id285ec6c22949b7f657f9a7661f6e6849ef305c2114d4e0a5f1e0c0ce386d17a.NewNotificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Octocat the octocat property
// returns a *OctocatRequestBuilder when successful
func (m *ApiClient) Octocat()(*i516aafeb513add23cb0b85066636cf037269fe65da463faa134b8956732972ba.OctocatRequestBuilder) {
    return i516aafeb513add23cb0b85066636cf037269fe65da463faa134b8956732972ba.NewOctocatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Organizations the organizations property
// returns a *OrganizationsRequestBuilder when successful
func (m *ApiClient) Organizations()(*ie80ab2c4f62921be85b8d011cd12092ac3cfcc2e38bbf97c34edad846cd82460.OrganizationsRequestBuilder) {
    return ie80ab2c4f62921be85b8d011cd12092ac3cfcc2e38bbf97c34edad846cd82460.NewOrganizationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Orgs the orgs property
// returns a *OrgsRequestBuilder when successful
func (m *ApiClient) Orgs()(*i46a0b04ed3ba56ba62fbc553ce7cfe242cedfc763368fc6fff06a19d3cbe08c2.OrgsRequestBuilder) {
    return i46a0b04ed3ba56ba62fbc553ce7cfe242cedfc763368fc6fff06a19d3cbe08c2.NewOrgsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Projects the projects property
// returns a *ProjectsRequestBuilder when successful
func (m *ApiClient) Projects()(*i2553c08de368b7330dccbb084609f22b761d7062044d70c34052481b283f5b9a.ProjectsRequestBuilder) {
    return i2553c08de368b7330dccbb084609f22b761d7062044d70c34052481b283f5b9a.NewProjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rate_limit the rate_limit property
// returns a *Rate_limitRequestBuilder when successful
func (m *ApiClient) Rate_limit()(*ib43f1d21bad1eaccb3220a6895cc4da70307341b670b60c2aa631d2995e47b3e.Rate_limitRequestBuilder) {
    return ib43f1d21bad1eaccb3220a6895cc4da70307341b670b60c2aa631d2995e47b3e.NewRate_limitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repos the repos property
// returns a *ReposRequestBuilder when successful
func (m *ApiClient) Repos()(*i22bad5e37e0fdcc5f6386ac06d16691f9e070a95085db42b968fccb55bb22a62.ReposRequestBuilder) {
    return i22bad5e37e0fdcc5f6386ac06d16691f9e070a95085db42b968fccb55bb22a62.NewReposRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repositories the repositories property
// returns a *RepositoriesRequestBuilder when successful
func (m *ApiClient) Repositories()(*ieb3561664edbcd472a4089a7b601c05b4d7de53178c12aa3e14e0d6f1fcb91a7.RepositoriesRequestBuilder) {
    return ieb3561664edbcd472a4089a7b601c05b4d7de53178c12aa3e14e0d6f1fcb91a7.NewRepositoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Scim the scim property
// returns a *ScimRequestBuilder when successful
func (m *ApiClient) Scim()(*icbde3b6cc6b306eaac2200351e57ca9a3b13479fa37b08baaeb4b17445835c36.ScimRequestBuilder) {
    return icbde3b6cc6b306eaac2200351e57ca9a3b13479fa37b08baaeb4b17445835c36.NewScimRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Search the search property
// returns a *SearchRequestBuilder when successful
func (m *ApiClient) Search()(*i019b368180824531ee10a9fabd22410b95ccd81f7e1b1f8416c99a317b24aff0.SearchRequestBuilder) {
    return i019b368180824531ee10a9fabd22410b95ccd81f7e1b1f8416c99a317b24aff0.NewSearchRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Setup the setup property
// returns a *SetupRequestBuilder when successful
func (m *ApiClient) Setup()(*i1da9ff71b80748c2ffa3444d6d942401f510a5d59e7d0b85fba3f675a84cd93a.SetupRequestBuilder) {
    return i1da9ff71b80748c2ffa3444d6d942401f510a5d59e7d0b85fba3f675a84cd93a.NewSetupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Teams the teams property
// returns a *TeamsRequestBuilder when successful
func (m *ApiClient) Teams()(*iab4fa88f4cfb41b10ba8be6db9ff835cabee88e1b3ebd8ca2a24889bd9a2bb66.TeamsRequestBuilder) {
    return iab4fa88f4cfb41b10ba8be6db9ff835cabee88e1b3ebd8ca2a24889bd9a2bb66.NewTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get Hypermedia links to resources accessible in GitHub's REST API
// returns a *RequestInformation when successful
func (m *ApiClient) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// User the user property
// returns a *UserRequestBuilder when successful
func (m *ApiClient) User()(*i1920a84396ec87b264b65c3d4feffd636cd814025614ed8fc6837257346128db.UserRequestBuilder) {
    return i1920a84396ec87b264b65c3d4feffd636cd814025614ed8fc6837257346128db.NewUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users the users property
// returns a *UsersRequestBuilder when successful
func (m *ApiClient) Users()(*if155ce7328d65c6829538f8121c72e07753ee85681edf8f539ba03b35a5b2a44.UsersRequestBuilder) {
    return if155ce7328d65c6829538f8121c72e07753ee85681edf8f539ba03b35a5b2a44.NewUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Zen the zen property
// returns a *ZenRequestBuilder when successful
func (m *ApiClient) Zen()(*ideee9aaead60f0751b30729fcf84e07ef70a2f387455613efaccc0a48908813a.ZenRequestBuilder) {
    return ideee9aaead60f0751b30729fcf84e07ef70a2f387455613efaccc0a48908813a.NewZenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
