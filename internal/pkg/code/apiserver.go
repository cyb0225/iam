/**
@author: yeebing
@date: 2022/9/24
**/

package code

// User
const (
	// ErrUserNotFound - 404: User not found.
	ErrUserNotFound = iota + 200001

	// ErrTokenNotExisted - 404: Token not existed.
	ErrTokenNotExisted

	// ErrCodeNotExisted - 404: Code not existed.
	ErrCodeNotExisted

	// ErrPasswordRequired - 200: Password does not meet requirements
	ErrPasswordRequired

	// ErrEmailRequired - 200: Email does not meet requirements
	ErrEmailRequired

	// ErrUserAlreadyExisted - 200: User already existed
	ErrUserAlreadyExisted

	// ErrSendEmail -500: Send email failed
	ErrSendEmail

	// ErrPassword -200: Error password
	ErrPassword

	// ErrGetToken -400: Get Token from head failed.
	ErrGetToken

	// ErrGetTokenFromCtx -500: Get Token from context failed
	ErrGetTokenFromCtx

	// ErrGetUserIDFromCtx -500: Get User ID from context failed
	ErrGetUserIDFromCtx

	// ErrPasswordTooShort -200: Password is shorter than 8.
	ErrPasswordTooShort

	// ErrPasswordTooLong -200: Password is longer than 26.
	ErrPasswordTooLong

	// ErrPasswordTooSimple -200: The kind of character in password is less than 2.
	ErrPasswordTooSimple
)
