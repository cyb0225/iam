/**
@author: yeebing
@date: 2022/9/25
**/

package core

import (
	"github.com/cyb0225/iam/pkg/errno"
	"github.com/cyb0225/iam/pkg/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
		if coder.HTTPStatus() >= 200 && coder.HTTPStatus() < 400 {
			log.WithRequestID(c).Info("logical error",
				zap.Int("http", coder.HTTPStatus()),
				zap.Int("code", coder.Code()),
				zap.Error(err),
			)
		} else {
			// TODO: log stack trace.
			log.WithRequestID(c).Error("unexpected error",
				zap.Int("http", coder.HTTPStatus()),
				zap.Int("code", coder.Code()),
				zap.Any("error", err),
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
