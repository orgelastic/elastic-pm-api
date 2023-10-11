package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stewie1520/elasticpmapi/core"
	"net/http"
)

type healthApi struct {
	app core.App
}

func bindHealthApi(app core.App, ginEngine *gin.Engine) {
	api := &healthApi{
		app: app,
	}

	subGroup := ginEngine.Group("/health")
	subGroup.GET("/ready", api.readiness)
	subGroup.GET("/live", api.liveness)
}

func (api *healthApi) liveness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"healthy": true,
	})
}

func (api *healthApi) readiness(c *gin.Context) {
	err := api.app.DB().Ping(c)
	unhealthyResponse := gin.H{
		"healthy": false,
	}

	if err != nil {
		if api.app.IsDebug() {
			unhealthyResponse["database"] = err.Error()
		} else {
			unhealthyResponse["database"] = "database error"
		}

		c.JSON(http.StatusOK, unhealthyResponse)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"healthy": true,
	})
}
