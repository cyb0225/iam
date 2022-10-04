/**
@author: yeebing
@date: 2022/10/1
**/

package user

import (
	"github.com/cyb0225/iam/internal/pkg/core"
	"github.com/gin-gonic/gin"
)

func (u *User) List(c *gin.Context) {
	res, err := u.srv.User().List(c)
	core.WriteResponse(c, err, res)
}
