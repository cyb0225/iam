/**
@author: yeebing
@date: 2022/10/3
**/

package user

import (
	"github.com/cyb0225/iam/internal/pkg/core"
	"github.com/gin-gonic/gin"
)

func (u *User) GetCode(c *gin.Context) {
	email := c.Query("email")
	res, err := u.srv.User().GetCode(c, email)
	core.WriteResponse(c, err, res)
}
