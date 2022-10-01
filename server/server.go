package server

import (
	// import local packages
	"butterfly/server/router"
)

func InitServer() {
	server := router.InitRouter()
	server.Run(":7887")
}
