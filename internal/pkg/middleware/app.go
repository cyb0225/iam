/**
@author: yeebing
@date: 2022/9/27
**/

package middleware

import "github.com/gin-gonic/gin"

// App set server server's information(server name and version)
func App(name string, version string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", name)
		c.Set("app_version", version)
		c.Next()
	}
}
