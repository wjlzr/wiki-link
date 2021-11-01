package util

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"wiki-link/common/lang"
	coreHttp "wiki-link/core/http"
	service "wiki-link/service/error"
)

//错误返回的json
func ResponseErrorJson(c *gin.Context, code int64, lang string, err error) {
	status := http.StatusOK
	if errors.Is(err, coreHttp.ErrVisitLimit) {
		status = http.StatusInternalServerError
		code = 2000000
	}
	c.JSON(status, gin.H{
		"code": code,
		"msg":  service.Errors(code, lang),
	})
	c.Abort()
}

//
func ResponseErrorJsonWithMsg(c *gin.Context, code int64, msg string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": code,
		"msg":  msg,
	})
	c.Abort()
}

//正确返回的json
func ResponseJson(c *gin.Context, msg, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
		"data": data,
	})
}

//
func ResponseErrorFormatJson(c *gin.Context, code int64, lang string, v ...interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": code,
		"msg":  fmt.Sprintf(service.Errors(code, lang), v...),
	})
}

//获取语言
func GetLang(c *gin.Context) string {
	language := c.Request.Header.Get("Language")
	if len(language) > 0 {
		return language
	}
	return lang.ZH_CN
}

//获取版本号
func GetVersion(c *gin.Context) string {
	version := c.Request.Header.Get("version")
	if len(version) > 0 {
		return version
	}
	return "1.1.0"
}

// 获取客户端
func GeClient(c *gin.Context) string {
	return c.Request.Header.Get("x-client")
}
