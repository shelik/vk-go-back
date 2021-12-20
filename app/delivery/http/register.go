package http

import (
	"github.com/gin-gonic/gin"
	"github.com/shelik/mtranslate/app"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.Engine, uc app.Usecase) {
	h := NewHandler(uc)

	apiEndpoints := router.Group("/api")
	{
		apiEndpoints.GET("/", h.Translate)
	}
}
