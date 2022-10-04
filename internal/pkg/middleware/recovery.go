/**
@author: yeebing
@date: 2022/9/27
**/

package middleware

import (
	"github.com/cyb0225/iam/pkg/email"
	"github.com/cyb0225/iam/pkg/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"runtime/debug"
)

var (
	toEmail = "yeebingchen@qq.com"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Logger.Error("panic", zap.Any("err", err), zap.String("stack", string(debug.Stack())))
				// send email
				subject := "iam 服务警报"
				text := []byte("出现panic错误，请及时处理")
				_ = email.Send([]string{toEmail}, subject, text)
			}
		}()
		c.Next()
	}
}
