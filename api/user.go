package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stewie1520/elasticpmapi/api/middleware"
	"github.com/stewie1520/elasticpmapi/api/response"
	"github.com/stewie1520/elasticpmapi/core"
	"github.com/stewie1520/elasticpmapi/usecases"
	"github.com/supertokens/supertokens-golang/recipe/session"
)

type userApi struct {
	app core.App
}

func bindUserApi(app core.App, ginEngine *gin.Engine) error {
	api := &userApi{
		app: app,
	}

	subGroup := ginEngine.Group("/user", middleware.VerifySession(nil))
	subGroup.GET("/me", api.getUser)

	return nil
}

func (api *userApi) getUser(c *gin.Context) {
	q := usecases.NewGetUserByAccountIDQuery(api.app)
	sessionContainer := session.GetSessionFromRequestContext(c.Request.Context())
	// TODO: this is wrong, fix later
	q.AccountID = sessionContainer.GetUserID()

	if user, err := q.Execute(); err != nil {
		response.NewBadRequestError("", err).WithGin(c)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
