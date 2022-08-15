package comments

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/CloudNua/go-api-2/pkg/common/models"
	"github.com/CloudNua/go-api-2/pkg/httputil"
)

// ListComments godoc
// @Summary      List Comments
// @Description  get comments
// @Tags         comments
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Comment
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /comments [get]
func (h handler) GetComments(ctx *gin.Context) {
	var comments []models.Comment

	if result := h.DB.Find(&comments); result.Error != nil {
		httputil.NewError(ctx, http.StatusNotFound, result.Error)
		// ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &comments)
}
