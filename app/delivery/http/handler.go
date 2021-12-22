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

type PhotoAlbums struct {
	OwnerID string   `json:"owner_id"`
	Albums  []string `json:"gallery_ids"`
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

	galleries := h.uc.GetGalleries(c.Request.Context(), ownerID)

	fmt.Println(galleries)
	c.JSON(http.StatusOK, galleries)
}

func (h *Handler) GetPhotos(c *gin.Context) {
	photoAlbums := new(PhotoAlbums)

	if err := c.BindJSON(photoAlbums); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(err)

		return
	}

	fmt.Println("From front:")
	fmt.Println(photoAlbums.Albums)

	photos := h.uc.GetPhotos(c.Request.Context(), photoAlbums.OwnerID, photoAlbums.Albums)

	c.JSON(http.StatusOK, photos)
}
