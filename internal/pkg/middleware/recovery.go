/**
@author: yeebing
@date: 2022/9/27
**/

package middleware

import (
	"github.com/cyb0225/iam/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
	"runtime/debug"
)

func Recovery(email *email.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Logger.Error("panic", zap.Any("err", err), zap.String("stack", string(debug.Stack())))
			}
		}()
		c.Next()
	}
}
