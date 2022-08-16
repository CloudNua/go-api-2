package comments

import (
	"net/http"

	"github.com/CloudNua/go-api-2/pkg/common/models"
	"github.com/CloudNua/go-api-2/pkg/httputil"
	"github.com/gin-gonic/gin"
)

// UpdateComment godoc
// @Summary      Update a comment
// @Description  Update by json comment
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        id       path      int                  true  "Comment ID"
// @Param        comment  body      models.UpdateCommentRequestBody  true  "Update comment"
// @Success      200      {object}  models.Comment
// @Failure      400      {object}  httputil.HTTPError400
// @Failure      404      {object}  httputil.HTTPError404
// @Failure      500      {object}  httputil.HTTPError500
// @Router       /comments/{id} [put]
func (h handler) UpdateComment(ctx *gin.Context) {
	id := ctx.Param("id")
	body := models.UpdateCommentRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		// ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var comment models.Comment

	if result := h.DB.First(&comment, id); result.Error != nil {
		httputil.NewError(ctx, http.StatusNotFound, result.Error)
		// ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	comment.Title = body.Title
	comment.Author = body.Author
	comment.Slug = body.Slug

	h.DB.Save(&comment)

	ctx.JSON(http.StatusOK, &comment)
}
