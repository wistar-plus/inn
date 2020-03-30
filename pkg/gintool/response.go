package gintool

import (
	"inn/pkg/e"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//ResJSON response json
func ResJSON(c *gin.Context, httpcode, code int, msg string, data interface{}) {
	c.JSON(httpcode, response{
		Code: code,
		Msg:  e.GetMsg(code) + msg,
		Data: data,
	})
}

//ResSuccess response success
func ResSuccess(c *gin.Context, data interface{}) {
	ResJSON(c, 200, e.SUCCESS, "", data)
}

//ResError response error
func ResError(c *gin.Context, code int, err error) {
	ResJSON(c, 200, code, "："+err.Error(), nil)
}

//ResCode response code
func ResCode(c *gin.Context, code int) {
	ResJSON(c, 200, code, "", nil)
}
