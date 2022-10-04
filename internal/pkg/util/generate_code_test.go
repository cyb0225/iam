/**
@author: yeebing
@date: 2022/10/3
**/

package util

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestGenerateCode(t *testing.T) {
	vcode := GenerateCode()
	_, err := strconv.Atoi(vcode)
	t.Log(vcode)
	assert.Equal(t, nil, err)
	assert.Equal(t, 6, len(vcode))
}
