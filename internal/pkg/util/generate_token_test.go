/**
@author: yeebing
@date: 2022/10/3
**/

package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	key := "key"
	token1 := GenerateToken(key)
	t.Log("token1:", token1)
	time.Sleep(time.Millisecond)
	token2 := GenerateToken(key)
	t.Log("token2:", token2)
	assert.NotEqual(t, token1, token2)

}
