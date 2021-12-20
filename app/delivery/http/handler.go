package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelik/vk-go-back/app"
)

// Handler ...
type Handler struct {
	uc app.Usecase
}

// NewHandler ...
func NewHandler(uc app.Usecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

// HelloWorld ...
func (h *Handler) GetGalleries(c *gin.Context) {
	ownerID := c.Request.URL.Query().Get("owner_id")
	fmt.Println(ownerID)

	galleries := h.uc.GetGalleries(c.Request.Context(), ownerID)

	c.JSON(http.StatusOK, galleries)
}
