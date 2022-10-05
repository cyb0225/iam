/**
@author: yeebing
@date: 2022/10/2
**/

package code

import "github.com/cyb0225/iam/pkg/errno"

func init() {
	errno.RegisterWithArgs(ErrDatabase, 500, "Database error")
	errno.RegisterWithArgs(ErrTypeAssertion, 500, "Type assertion error")
	errno.RegisterWithArgs(ErrTypeTransform, 500, "Type transform error")
	errno.RegisterWithArgs(ErrBind, 400, "Error occurred while binding the request body to the struct.")
	errno.RegisterWithArgs(ErrGetFormFile, 400, "Can not get post form file")
	errno.RegisterWithArgs(ErrSaveUploadFile, 500, "Can not save the upload file")

	errno.RegisterWithArgs(ErrUserNotFound, 404, "User not found")
	errno.RegisterWithArgs(ErrTokenNotExisted, 404, "Token not existed")
	errno.RegisterWithArgs(ErrCodeNotExisted, 404, "Code not existed")
	errno.RegisterWithArgs(ErrPasswordRequired, 200, "Password does not meet requirements")
	errno.RegisterWithArgs(ErrEmailRequired, 200, "Email does not meet requirements")
	errno.RegisterWithArgs(ErrUserAlreadyExisted, 200, "User already existed")
	errno.RegisterWithArgs(ErrSendEmail, 500, "Failed to send email")
	errno.RegisterWithArgs(ErrPassword, 200, "Error password")
	errno.RegisterWithArgs(ErrGetToken, 400, "Get Token from head failed")
	errno.RegisterWithArgs(ErrGetTokenFromCtx, 500, "Get Token from context failed")
	errno.RegisterWithArgs(ErrGetUserIDFromCtx, 500, "Get User ID from context failed")
	errno.RegisterWithArgs(ErrPasswordTooShort, 200, "Password is shorter than 8.")
	errno.RegisterWithArgs(ErrPasswordTooLong, 200, "Password is longer than 26.")
	errno.RegisterWithArgs(ErrPasswordTooSimple, 200, "The kind of character in password is less than 2.")
}
