/**
@author: yeebing
@date: 2022/10/1
**/

package user

import (
	"github.com/cyb0225/iam/internal/apiserver/service/v1/model"
	"github.com/cyb0225/iam/internal/pkg/code"
	"github.com/cyb0225/iam/internal/pkg/core"
	"github.com/cyb0225/iam/pkg/errno"
	"github.com/gin-gonic/gin"
)

func (u *User) Update(c *gin.Context) {
	req := &model.UserUpdateRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		core.WriteResponse(c, errno.WithCode(code.ErrBind, err), nil)
		return
	}
	err := u.srv.User().Update(c, req)
	core.WriteResponse(c, err, nil)
}
