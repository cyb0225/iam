/**
@author: yeebing
@date: 2022/9/27
**/

package middleware

import (
	"github.com/cyb0225/iam/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

var (
	fillInterval       = time.Second
	cap          int64 = 100
	quantum      int64 = 100
)

// RateLimitMiddleware set a rate limit bucket.
//  limit the number of requests in order to avoid excessive requests causing server crashes.
func RateLimitMiddleware() gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, cap, quantum)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			c.String(http.StatusForbidden, "rate limit")
			log.Logger.Warn("rate limit warn")
			c.Abort()
			return
		}
		c.Next()
	}
}
