/**
@author: yeebing
@date: 2022/10/3
**/

package util

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

func GenerateToken(key string) string {
	salt := Salt()

	encoder := md5.New()
	encoder.Write([]byte(key))
	encoder.Write([]byte(salt))
	en := encoder.Sum(nil)
	str := fmt.Sprintf("%x", en)
	return str
}

func Salt() string {
	t := time.Now().UnixNano()
	return strconv.Itoa(int(t))
}
