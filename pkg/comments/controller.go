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
	routes := router.Group("/v1/comment")
	routes.POST("", h.AddComment)
	routes.GET("", h.GetComments)
}
