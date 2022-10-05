/**
@author: yeebing
@date: 2022/9/25
**/

package user

import (
	"github.com/cyb0225/iam/internal/pkg/core"
	"github.com/gin-gonic/gin"
)

func (u *User) Logout(c *gin.Context) {
	err := u.srv.User().Logout(c)
	core.WriteResponse(c, err, nil)
}
