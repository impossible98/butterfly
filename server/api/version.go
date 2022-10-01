package api

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
	// import local packages
	"butterfly/global"
	"butterfly/utils/ecode"
	"butterfly/utils/format"
)

func Version(ctx *gin.Context) {
	format.HTTP(ctx, ecode.Success, nil, gin.H{
		"version":     global.VERSION,
		"api_version": global.API_VERSION,
	})
}
