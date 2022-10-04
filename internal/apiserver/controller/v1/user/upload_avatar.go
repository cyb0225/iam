/**
@author: yeebing
@date: 2022/10/1
**/

package user

import (
	"errors"
	"github.com/cyb0225/iam/internal/pkg/code"
	"github.com/cyb0225/iam/internal/pkg/core"
	"github.com/cyb0225/iam/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

var (
	dir = "./storage/avatar/"
)

func (u *User) UploadAvatar(c *gin.Context) {
	userID := c.GetUint64("id")
	if userID == 0 {
		core.WriteResponse(c, errno.WithCode(code.ErrGetUserIDFromCtx, errors.New("get user id from context failed")), nil)
		return
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		core.WriteResponse(c, errno.WithCode(code.ErrGetFormFile, err), nil)
		return
	}

	// save file
	strID := cast.ToString(userID)
	dst := dir + strID
	if err := c.SaveUploadedFile(file, dst); err != nil {
		core.WriteResponse(c, errno.WithCode(code.ErrSaveUploadFile, err), nil)
		return
	}

	// save avatar
	err = u.srv.User().UploadAvatar(c, dst)
	core.WriteResponse(c, err, nil)
}
