package comments

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	// Simple group: v1
	v1 := router.Group("/api/v1/")
	{
		comments := v1.Group("/comments")
		{
			comments.POST("", h.AddComment)
			comments.GET("", h.GetComments)
			comments.GET("/:id", h.GetComment)
			comments.PUT("/:id", h.UpdateComment)
			comments.DELETE("/:id", h.DeleteComment)
		}
	}
}
