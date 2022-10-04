/**
@author: yeebing
@date: 2022/10/1
**/

package middleware

import (
	"errors"
	"github.com/cyb0225/iam/internal/apiserver/cache"
	"github.com/cyb0225/iam/internal/pkg/code"
	"github.com/cyb0225/iam/internal/pkg/core"
	"github.com/cyb0225/iam/pkg/errno"
	"github.com/gin-gonic/gin"
)

func TokenAuth(ca cache.TokenCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		// got token from head
		token := c.GetHeader("token")
		if len(token) == 0 {
			err := errno.WithCode(code.ErrGetToken, errors.New("get token from head failed"))
			core.WriteResponse(c, err, nil)
			c.Abort()
			return
		}

		// check the token with cache
		val, err := ca.Get(c.Request.Context(), token)
		if err != nil {
			err := errno.WithCode(code.ErrTokenNotExisted, errors.New("token not existed"))
			core.WriteResponse(c, err, nil)
			c.Abort()
			return
		}
		// set token and user id to context
		c.Set("token", token)
		c.Set("id", val.UserID)
	}
}
