/**
@author: yeebing
@date: 2022/10/3
**/

package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJudgePassword(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err := JudgePassword("1234abcd?")
		assert.Equal(t, nil, err)
	})

	t.Run("too long", func(t *testing.T) {
		err := JudgePassword("a2?")
		t.Log(err.Error())
		assert.NotEqual(t, nil, err)
	})

	t.Run("too short", func(t *testing.T) {
		err := JudgePassword("12345678900987654321abcdefghi")
		t.Log(err.Error())
		assert.NotEqual(t, nil, err)
	})

	t.Run("too simple", func(t *testing.T) {
		err := JudgePassword("123456789")
		t.Log(err.Error())
		assert.NotEqual(t, nil, err)
	})
}
