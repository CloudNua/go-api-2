package comments

import (
	"net/http"

	"github.com/CloudNua/go-api-2/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// ListComments godoc
// @Summary      Get specific comment
// @Description  get string by ID
// @ID           get-string-by-int
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        id   	path      int  true  "Comment ID"
// @Success      200  	{array}   models.Comment
// @Failure      404    {object}  httputil.HTTPError404
// @Router       /comments/{id} [get]
func (h handler) GetComment(ctx *gin.Context) {
	id := ctx.Param("id")

	var comment models.Comment

	if result := h.DB.First(&comment, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &comment)
}
