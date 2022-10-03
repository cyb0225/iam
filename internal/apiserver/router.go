/**
@author: yeebing
@date: 2022/9/24
**/

package apiserver

import (
	"github.com/cyb0225/iam/internal/apiserver/controller/v1/user"
	"github.com/cyb0225/iam/internal/pkg/middleware"
	"github.com/cyb0225/iam/pkg/email"
	"github.com/gin-gonic/gin"
	"os"
)

func InitRouter(opts *Option) *gin.Engine {
	var e *gin.Engine
	if opts.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
		e = gin.New()
		e.Use(middleware.Recovery(email.Pool))
		e.Use(middleware.Logger(os.Stdout)) // TODO:
	} else { // debug is default mode
		e = gin.New()
		e.Use(gin.Recovery())
		e.Use(gin.Logger())
	}
	initMiddleware(e, opts)
	initController(e, opts)
	return e
}

func initMiddleware(e *gin.Engine, opts *Option) {
	e.MaxMultipartMemory = 8 << 20 // 8 MiB, maximum file upload size supported.
	e.Use(middleware.App())
	e.Use(middleware.RequestID())
	e.Use(middleware.RateLimitMiddleware())
	e.Use(middleware.Cors())
}

func initController(r *gin.Engine, opts *Option) {
	app := r.Group("/iam")

	v1 := app.Group("/v1")
	{
		userv1 := v1.Group("/user")
		{
			userController := user.New()
			userv1.GET(":id", userController.Get)
			userv1.GET("/list", userController.List)
			userv1.POST("/login", userController.Login)
			userv1.POST("/register", userController.Register)
			userv1.GET("/code", userController.GetCode)
			userv1.Use(middleware.TokenAuth())
			userv1.DELETE("/logout", userController.Logout)
			userv1.PUT("/update/password", userController.ChangePassword)
			userv1.PUT("/update/email", userController.ChangeEmail)
			userv1.PUT("/update/avatar", userController.UploadAvatar)
			userv1.PUT("/update", userController.Update)
		}
	}

}
