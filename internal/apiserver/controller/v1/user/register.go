/**
@author: yeebing
@date: 2022/9/27
**/

package user

import (
	"github.com/cyb0225/iam/internal/apiserver/service/v1/model"
	"github.com/cyb0225/iam/internal/pkg/code"
	"github.com/cyb0225/iam/internal/pkg/core"
	"github.com/cyb0225/iam/pkg/errno"
	"github.com/gin-gonic/gin"
)

// Register register a user account
func (u *User) Register(c *gin.Context) {
	req := &model.RegisterRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		core.WriteResponse(c, errno.WithCode(code.ErrBind, err), nil)
		return
	}
	res, err := u.srv.User().Register(c, req)
	core.WriteResponse(c, err, res)
}
