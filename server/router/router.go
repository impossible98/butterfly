package router

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
	// import local packages
	"butterfly/server/api"
)

func InitRouter() *gin.Engine {
	app := gin.Default()

	app.POST("/api/version", api.Version)
	// return
	return app
}
