/**
@author: yeebing
@date: 2022/10/3
**/

package util

import (
	"crypto/md5"
	"fmt"
)

func PasswordEncrypt(password string) string {
	data := md5.Sum([]byte(password))
	md5str := fmt.Sprintf("%x", data) //将[]byte转成16进制
	return md5str
}
