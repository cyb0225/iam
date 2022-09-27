/**
@author: yeebing
@date: 2022/9/27
**/

package middleware

import (
	"github.com/gin-gonic/gin"
	"io"
)

// Logger set log file and log format.
func Logger(writer io.Writer) gin.HandlerFunc {
	return func(c *gin.Context) {
		gin.LoggerWithConfig(gin.LoggerConfig{
			Output: writer,
		})
	}
}
