package comments

import (
	"net/http"

	"github.com/CloudNua/go-api-2/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// DeleteComment godoc
// @Summary      Delete a comment
// @Description  Delete by comment ID
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        id   	   path      int  true  "Comment ID"  Format(int64)
// @Success      200      {object}  models.Comment
// @Failure      400      {object}  httputil.HTTPError400
// @Failure      404      {object}  httputil.HTTPError404
// @Router       /comments/{id} [delete]
func (h handler) DeleteComment(ctx *gin.Context) {
	id := ctx.Param("id")

	var comment models.Comment

	if result := h.DB.First(&comment, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&comment)

	ctx.Status(http.StatusOK)
}
