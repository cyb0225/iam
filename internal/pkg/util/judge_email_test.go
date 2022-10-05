/**
@author: yeebing
@date: 2022/10/3
**/

package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJudgeEmail(t *testing.T) {
	t.Run("return true", func(t *testing.T) {
		b := JudgeEmail("2103561941@qq.com")
		assert.Equal(t, true, b)
	})

	t.Run("return false", func(t *testing.T) {
		b := JudgeEmail("helloworld.com")
		assert.Equal(t, false, b)
	})
}
