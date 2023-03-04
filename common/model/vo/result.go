package vo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SUCCESS = "00000"
	ERROR   = "10000"
)

func BuildResult(code string, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Result{
		code,
		data,
		msg,
	})
}

func Success(c *gin.Context) {
	BuildResult(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func SuccessMsg(message string, c *gin.Context) {
	BuildResult(SUCCESS, map[string]interface{}{}, message, c)
}

func SuccessData(data interface{}, c *gin.Context) {
	BuildResult(SUCCESS, data, "操作成功", c)
}

func SuccessDetail(data interface{}, message string, c *gin.Context) {
	BuildResult(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	BuildResult(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailMsg(message string, c *gin.Context) {
	BuildResult(ERROR, map[string]interface{}{}, message, c)
}

func FailDetail(data interface{}, message string, c *gin.Context) {
	BuildResult(ERROR, data, message, c)
}
