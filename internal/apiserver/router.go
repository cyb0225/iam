/**
@author: yeebing
@date: 2022/9/24
**/

package apiserver

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.New()
	initMiddleware(router)
	initController(router)
	return router
}

func initMiddleware(r *gin.Engine) {
}

func initController(r *gin.Engine) {
}
