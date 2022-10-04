package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPasswordEncrypt(t *testing.T) {
	password := "123456"
	encrypt := PasswordEncrypt(password)
	assert.NotEqual(t, password, encrypt)
	t.Logf("password: %s, encrypt: %s\n", password, encrypt)
}
