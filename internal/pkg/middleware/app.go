/**
@author: yeebing
@date: 2022/9/27
**/

package middleware

import "github.com/gin-gonic/gin"

var (
	appName = "iam"
	version = "v1.0.0"
)

// App set server server's information(server name and version)
func App() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", appName)
		c.Set("app_version", version)
		c.Next()
	}
}
