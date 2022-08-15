package comments

import (
	"net/http"

	"github.com/CloudNua/go-api-2/pkg/common/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type AddCommentRequestBody struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Slug   string `json:"slug"`
}

// @BasePath /v1/comment

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func (h handler) AddComment(ctx *gin.Context) {
	body := AddCommentRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var comment models.Comment

	comment.ID = uuid.NewV4().String()
	comment.Title = body.Title
	comment.Author = body.Author
	comment.Slug = body.Slug

	if result := h.DB.Create(&comment); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &comment)
}
