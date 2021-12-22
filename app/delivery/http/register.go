package http

import (
	"github.com/gin-gonic/gin"
	"github.com/shelik/vk-go-back/app"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.Engine, uc app.Usecase) {
	h := NewHandler(uc)

	apiEndpoints := router.Group("/api")
	{
		apiEndpoints.GET("/get_galleries", h.GetGalleries)
		apiEndpoints.POST("/get_photos", h.GetPhotos)
	}
}
