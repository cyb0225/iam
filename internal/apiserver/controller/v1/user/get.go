/**
@author: yeebing
@date: 2022/9/25
**/

package user

import (
	"github.com/cyb0225/iam/internal/pkg/code"
	"github.com/cyb0225/iam/internal/pkg/core"
	"github.com/cyb0225/iam/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func (u *User) Get(c *gin.Context) {
	id := c.Param("id")
	userID, err := cast.ToUint64E(id)
	if err != nil {
		core.WriteResponse(c, errno.WithCode(code.ErrTypeTransform, err), nil)
		return
	}

	res, err := u.srv.User().Get(c, userID)
	core.WriteResponse(c, err, res)
}
