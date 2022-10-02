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
)
