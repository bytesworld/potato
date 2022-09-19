package response

import (
	"gitee.com/bytesworld/tomato/internal/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	ErrorCode int `json:"error_code"`
	Data interface{} `json:"data"`
	Message string `json:"message" yaml:"message"`
}

func Success(ctx *gin.Context,data interface{})  {
	ctx.JSON(
		http.StatusOK,
		Response{
			0,
			data,
			"ok",
		})
}

func Fail(c *gin.Context, errorCode int, msg string) {
	c.JSON(http.StatusOK, Response{
		errorCode,
		nil,
		msg,
	})
}

// FailByError 失败响应 返回自定义错误的错误码、错误信息
func FailByError(c *gin.Context, error controllers.CustomError) {
	Fail(c, error.ErrorCode, error.ErrorMsg)
}

// ValidateFail 请求参数验证失败
func ValidateFail(c *gin.Context, msg string)  {
	Fail(c, controllers.Errors.ValidateError.ErrorCode, msg)
}

// BusinessFail 业务逻辑失败
func BusinessFail(c *gin.Context, msg string) {
	Fail(c, controllers.Errors.BusinessError.ErrorCode, msg)
}

