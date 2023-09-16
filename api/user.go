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

func bindUserApi(app core.App, ginEngine *gin.Engine) {
	api := &userApi{
		app: app,
	}

	subGroup := ginEngine.Group("/user", middleware.VerifySession(nil))
	subGroup.GET("/me", api.getCurrentUserInfo)
}

// getCurrentUserInfo Return current logged in user information
// @Summary Get current user information
// @Description Return current logged in user information
// @Tags profile
// @Accept json
// @Produce json
// @Success 200 {object} usecases.GetUserByAccountIDResponse
// @Router /user/me [get]
func (api *userApi) getCurrentUserInfo(c *gin.Context) {
	q := usecases.NewGetUserByAccountIDQuery(api.app)
	sessionContainer := session.GetSessionFromRequestContext(c.Request.Context())
	q.AccountID = sessionContainer.GetUserID()

	if user, err := q.Execute(); err != nil {
		response.NewBadRequestError("", err).WithGin(c)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
