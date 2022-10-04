/**
@author: yeebing
@date: 2022/9/25
**/

package core

import (
	"github.com/cyb0225/iam/pkg/log"
	"go.uber.org/zap"
	"net/http"

	"github.com/cyb0225/iam/pkg/errno"
	"github.com/gin-gonic/gin"
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
			log.Logger.Info("logical error",
				zap.Int("http", coder.HTTPStatus()),
				zap.Int("code", coder.Code()),
				zap.Error(err),
			)
		} else {
			// log stack trace.
			log.Logger.Error("unexpected error",
				zap.Int("http", coder.HTTPStatus()),
				zap.Int("code", coder.Code()),
				zap.Any("error", errno.StackError(err)),
			)
		}
		c.JSON(coder.HTTPStatus(), response{
			Code: coder.Code(),
			Msg:  coder.String(),
		})
		return
	}

	c.JSON(http.StatusOK, response{
		Code: 0,
		Data: data,
	})
}
