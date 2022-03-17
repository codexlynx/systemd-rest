package routes

import (
	"github.com/codexlynx/systemd-rest/internal/http/controllers"
	"github.com/gin-gonic/gin"
)

// StartRouter to start Gin router.
func StartRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		api.GET("units", controllers.GetUnits)
		api.GET("units/:name", controllers.GetUnit)
		api.POST("units/:name/start", controllers.StartUnit)
		api.POST("units/:name/stop", controllers.StopUnit)
		api.GET("journal/:name", controllers.GetUnitJournal)
	}
	return router
}
