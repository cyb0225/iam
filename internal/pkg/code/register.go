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
	errno.RegisterWithArgs(ErrPasswordRequired, 200, "Password does not meet requirements")
	errno.RegisterWithArgs(ErrEmailRequired, 200, "Email does not meet requirements")
	errno.RegisterWithArgs(ErrUserAlreadyExisted, 200, "User already existed")
	errno.RegisterWithArgs(ErrSendEmail, 500, "Failed to send email")
	errno.RegisterWithArgs(ErrPassword, 200, "Error password")
	errno.RegisterWithArgs(ErrGetToken, 400, "Get Token from head failed")
}
