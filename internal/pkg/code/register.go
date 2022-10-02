/**
@author: yeebing
@date: 2022/10/2
**/

package code

import "github.com/cyb0225/iam/pkg/errno"

func init() {
	errno.RegisterWithArgs(ErrDatabase, 500, "Database error")
	errno.RegisterWithArgs(ErrTypeAssertion, 500, "Type assertion error")

	errno.RegisterWithArgs(ErrUserNotFound, 404, "User not found")
	errno.RegisterWithArgs(ErrTokenNotExisted, 404, "Token not existed")
	errno.RegisterWithArgs(ErrCodeNotExisted, 404, "Code not existed")
}
