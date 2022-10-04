/**
@author: yeebing
@date: 2022/9/25
**/

package user

import (
	"github.com/cyb0225/iam/internal/apiserver/service/v1/model"
	"github.com/cyb0225/iam/internal/pkg/code"
	"github.com/cyb0225/iam/internal/pkg/core"
	"github.com/cyb0225/iam/pkg/errno"
	"github.com/gin-gonic/gin"
)

func (u *User) Login(c *gin.Context) {
	req := &model.LoginRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		core.WriteResponse(c, errno.WithCode(code.ErrBind, err), nil)
		return
	}
	res, err := u.srv.User().Login(c, req)
	core.WriteResponse(c, err, res)
}
