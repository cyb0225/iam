/**
@author: yeebing
@date: 2022/10/2
**/

package code

import "github.com/cyb0225/iam/pkg/errno"

func init() {
	errno.RegisterWithArgs(ErrDatabase, 500, "Database error")

	errno.RegisterWithArgs(ErrUserNotFound, 404, "User not found")
}
