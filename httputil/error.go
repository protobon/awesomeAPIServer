package httputil

import "github.com/gin-gonic/gin"

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:  status,
		Error: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError example
type HTTPError struct {
	Code  int    `json:"code" example:"400"`
	Error string `json:"error" example:"Invalid request payload"`
}
