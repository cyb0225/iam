/**
@author: yeebing
@date: 2022/10/3
**/

package util

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateCode create a code
func GenerateCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vcode
}
