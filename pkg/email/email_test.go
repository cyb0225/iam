/**
@author: yeebing
@date: 2022/10/1
**/

package email

import (
	"github.com/jordan-wright/email"
	"testing"
	"time"
)

// TestSendEmail used to test
func TestSendEmail(t *testing.T) {
	opts := Option{
		SMTPKey:         "kyxrxqinzltbfagd",
		FromEmail:       "yeebing<yeebingchen@qq.com>",
		EmailServerAddr: "smtp.qq.com:25",
		EmailServerHost: "smtp.qq.com",
	}

	if _, err := New(opts); err != nil {
		t.Fatal("init email failed")
	}

	e := email.NewEmail()
	e.From = "yeebing<yeebingchen@qq.com>"
	e.To = []string{"yeebingchen@qq.com"}
	e.Subject = "测试邮件"
	e.Text = []byte("测试邮件内容")
	if err := Pool.Send(e, time.Second*10); err != nil {
		t.Fatalf("send email failed: %v", err)
	}
}
