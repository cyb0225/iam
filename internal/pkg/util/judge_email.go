/**
@author: yeebing
@date: 2022/10/3
**/

package util

import "regexp"

func JudgeEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
