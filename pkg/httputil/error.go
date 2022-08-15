package httputil

import "github.com/gin-gonic/gin"

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError example
type HTTPError400 struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type HTTPError404 struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"not found"`
}

type HTTPError500 struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"internal server error"`
}
