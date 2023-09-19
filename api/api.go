package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stewie1520/elasticpmapi/api/middleware"
	"github.com/stewie1520/elasticpmapi/api/response"
	"github.com/stewie1520/elasticpmapi/core"
)

func InitApi(app core.App) (*gin.Engine, error) {
	if app.IsDebug() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(middleware.Cors(app))
	engine.Use(middleware.SuperToken)
	handleUnAuthorizedError(app)

	bindUserApi(app, engine)

	// TODO: add more api group here

	return engine, nil
}

func handleUnAuthorizedError(app core.App) {
	app.OnUnauthorizedAccess().Add(func(event *core.UnauthorizedAccessEvent) error {
		response.
			NewUnauthorizedError(event.Message, nil).
			WithResponseWriter(event.Res)

		return nil
	})
}
