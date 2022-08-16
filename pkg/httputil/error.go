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
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type HTTPError400 struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type HTTPError401 struct {
	Code    int    `json:"code" example:"401"`
	Message string `json:"message" example:"status unauthorized"`
}

type HTTPError404 struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status not found"`
}

type HTTPError405 struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"status method not allowed"`
}

type HTTPError408 struct {
	Code    int    `json:"code" example:"408"`
	Message string `json:"message" example:"status request timeout"`
}

type HTTPError500 struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"status internal server error"`
}

type HTTPError503 struct {
	Code    int    `json:"code" example:"503"`
	Message string `json:"message" example:"status service unavailable"`
}

type HTTPError505 struct {
	Code    int    `json:"code" example:"505"`
	Message string `json:"message" example:"status HTTP version not supported"`
}
