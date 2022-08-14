package comment

import (
	"net/http"

	"github.com/CloudNua/go-api-2/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetComments(ctx *gin.Context) {
	var comments []models.Comment

	if result := h.DB.Find(&comments); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &comments)
}
