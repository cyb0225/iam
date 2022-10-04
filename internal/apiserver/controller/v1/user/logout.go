/**
@author: yeebing
@date: 2022/9/25
**/

package user

import (
	"errors"
	"github.com/cyb0225/iam/internal/pkg/code"
	"github.com/cyb0225/iam/internal/pkg/core"
	"github.com/cyb0225/iam/pkg/errno"
	"github.com/gin-gonic/gin"
)

func (u *User) Logout(c *gin.Context) {
	token := c.GetString("token")
	if len(token) == 0 {
		core.WriteResponse(c, errno.WithCode(code.ErrGetTokenFromCtx, errors.New("get token from context failed")), nil)
		return
	}

	err := u.srv.User().Logout(c, token)
	core.WriteResponse(c, err, nil)
}
