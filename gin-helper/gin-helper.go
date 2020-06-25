package gin_helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


type ErrResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"error_code"`
	Message string `json:"error_message"`
}

type OkResponse struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
}

func WriteResponse(c *gin.Context, payload interface{}) {
	// always return http.StatusOK
	c.JSON(http.StatusOK, OkResponse{
		Success: true,
		Result:  payload,
	})
}

func WriteError(c *gin.Context, errCode int, errMsg string) {
	code := http.StatusOK
	c.JSON(code, ErrResponse{
		Success: false,
		Code:    errCode,
		Message: errMsg,
	})
}