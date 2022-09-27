/**
@author: yeebing
@date: 2022/9/27
**/

package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const (
// XRequestIDKey defines X-Request-ID key string.
//XRequestIDKey = "X-Request-ID"
)

// RequestID set request id to gin.context.
// XRequestIDKey is show which header key will store the request ID.
func RequestID(XRequestIDKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		rid := c.GetHeader(XRequestIDKey)

		// not set request_id, then create it.
		if rid == "" {
			rid = uuid.NewV4().String()
			c.Request.Header.Set(XRequestIDKey, rid)
			c.Set(XRequestIDKey, rid)
		}

		// Set XRequestIDKey header
		c.Writer.Header().Set(XRequestIDKey, rid)
		c.Next()
	}
}
