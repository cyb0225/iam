/**
@author: yeebing
@date: 2022/10/3
**/

package util

import (
	"errors"
	"github.com/cyb0225/iam/internal/pkg/code"
	"github.com/cyb0225/iam/pkg/errno"
	"regexp"
)

// JudgePassword Determine password strength
// rules:
// 1. length(8-26).
// 2. check the kind of acsii must be bigger than 2.
func JudgePassword(password string) error {
	if len(password) < 8 {
		return errno.WithCode(code.ErrPasswordTooShort, errors.New("password is too short"))
	}

	if len(password) > 26 {
		return errno.WithCode(code.ErrPasswordTooLong, errors.New("password is too long"))
	}

	level := 0
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[~!@#$%^&*?_-]+`}
	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, password)
		if match {
			level++
		}
	}

	// TODO: check the character that not allowed
	if level < 2 {
		return errno.WithCode(code.ErrPasswordTooSimple, errors.New("password is too simple"))
	}

	return nil
}
