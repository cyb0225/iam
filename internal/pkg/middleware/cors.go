/**
@author: yeebing
@date: 2022/9/27
**/

package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	maxAge = 10 * time.Second
)

// Or you can use the default cors by "cors.Default()".

// Cors use to resolve the former domain name Vinita
// input maxAge means how long (with second-precision) the results of a preflight request can be cached.
func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "https://github.com"
		//},
		MaxAge: maxAge,
	})
}
