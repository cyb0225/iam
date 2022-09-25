/**
@author: yeebing
@date: 2022/9/25
**/

package core

import (
	"github.com/cyb0225/iam/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// WriteResponse write response back to client.
func WriteResponse(c *gin.Context, err error, data interface{}) {
	// check the error
	if err != nil {
		coder := errno.ParseCoder(err)
		if coder.Code() >= 200 && coder.Code() < 300 {
			// TODO: do a info log
		} else {
			// TODO: do a error log.
		}
		c.JSON(coder.HTTPStatus(), response{
			Code: coder.Code(),
			Msg:  coder.String(),
		})
		return
	}

	// TODO: send success code back
	c.JSON(http.StatusOK, response{
		Code: 0,
	})
}
