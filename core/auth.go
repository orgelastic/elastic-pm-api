package core

import (
	"fmt"
	"net/http"

	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword/tpepmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func initSuperToken(app App) error {
	config := app.Config()

	apiBasePath := "/auth"
	websiteBasePath := "/auth"
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: config.SuperTokens.ConnectionUrl,
			APIKey:        config.SuperTokens.ApiKey,
		},
		AppInfo: supertokens.AppInfo{
			AppName:         "elastic-pm",
			APIDomain:       config.ApiDomain,
			WebsiteDomain:   config.WebsiteDomain,
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			thirdpartyemailpassword.Init(&tpepmodels.TypeInput{
				Providers: []tpmodels.ProviderInput{
					{
						Config: tpmodels.ProviderConfig{
							ThirdPartyId: "google",
							Clients: []tpmodels.ProviderClientConfig{
								{
									ClientID:     config.Google.ClientID,
									ClientSecret: config.Google.ClientSecret,
								},
							},
						},
					},
				},

				Override: &tpepmodels.OverrideStruct{
					Functions: func(originalImplementation tpepmodels.RecipeInterface) tpepmodels.RecipeInterface {
						originalEmailPasswordSignUp := *originalImplementation.EmailPasswordSignUp
						originalThirdPartySignInUp := *originalImplementation.ThirdPartySignInUp

						// override the email password sign up function
						(*originalImplementation.EmailPasswordSignUp) = func(email, password string, tenantId string, userContext supertokens.UserContext) (tpepmodels.SignUpResponse, error) {
							resp, err := originalEmailPasswordSignUp(email, password, tenantId, userContext)
							if err != nil {
								return tpepmodels.SignUpResponse{}, err
							}

							if resp.OK != nil {
								app.OnAfterAccountCreated().Trigger(&AccountCreatedEvent{
									ID:         resp.OK.User.ID,
									Email:      resp.OK.User.Email,
									TimeJoined: resp.OK.User.TimeJoined,
									ThirdParty: resp.OK.User.ThirdParty,
									TenantIds:  resp.OK.User.TenantIds,
								})
							}

							return resp, err
						}

						// override the thirdparty sign in / up function
						(*originalImplementation.ThirdPartySignInUp) = func(thirdPartyID, thirdPartyUserID, email string, oAuthTokens tpmodels.TypeOAuthTokens, rawUserInfoFromProvider tpmodels.TypeRawUserInfoFromProvider, tenantId string, userContext supertokens.UserContext) (tpepmodels.SignInUpResponse, error) {
							resp, err := originalThirdPartySignInUp(thirdPartyID, thirdPartyUserID, email, oAuthTokens, rawUserInfoFromProvider, tenantId, userContext)
							if err != nil {
								return tpepmodels.SignInUpResponse{}, err
							}

							if resp.OK != nil {
								user := resp.OK.User
								fmt.Println(user)

								accessToken := resp.OK.OAuthTokens["access_token"].(string)
								firstName := resp.OK.RawUserInfoFromProvider.FromUserInfoAPI["first_name"].(string)

								fmt.Println(accessToken)
								fmt.Println(firstName)

								if resp.OK.CreatedNewUser {
									app.OnAfterAccountCreated().Trigger(&AccountCreatedEvent{
										ID:         resp.OK.User.ID,
										Email:      resp.OK.User.Email,
										TimeJoined: resp.OK.User.TimeJoined,
										ThirdParty: resp.OK.User.ThirdParty,
										TenantIds:  resp.OK.User.TenantIds,
									})
								}
							}

							return resp, err
						}

						return originalImplementation
					},
				},
			}),

			session.Init(&sessmodels.TypeInput{
				ErrorHandlers: &sessmodels.ErrorHandlers{
					OnUnauthorised: func(message string, req *http.Request, res http.ResponseWriter) error {
						return app.OnUnauthorizedAccess().Trigger(&UnauthorizedAccessEvent{message, req, res})
					},
				},
			}),
			dashboard.Init(nil),
		},
	})

	return err
}
