/**
@author: yeebing
@date: 2022/10/1
**/

package email

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestSendEmail used to test
func TestSendEmail(t *testing.T) {
	sends := []string{"yeebingchen@qq.com"}
	subject := "subject"
	text := []byte("text")

	t.Run("with out init email Pool", func(t *testing.T) {
		err := Send(sends, subject, text)
		assert.NotEqual(t, nil, err)
	})

	opts := Option{
		SMTPKey:   "123",
		FromEmail: "yeebingchen@qq.com",
	}

	if _, err := New(opts); err != nil {
		t.Fatal("init email failed")
	}

	t.Run("smtp_key error", func(t *testing.T) {
		err := Send(sends, subject, text)
		assert.NotEqual(t, nil, err)
	})

	opts.SMTPKey = "123"

	if _, err := New(opts); err != nil {
		t.Fatal("init email failed")
	}

	t.Run("success", func(t *testing.T) {
		err := Send(sends, subject, text)
		assert.Equal(t, nil, err)
	})

}
