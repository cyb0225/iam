/**
@author: yeebing
@date: 2022/9/27
**/

package email

import (
	"errors"
	"github.com/jordan-wright/email"
	"net/smtp"
	"time"
)

type Option struct {
	SMTPKey   string `yaml:"smtp_key"`
	FromEmail string `yaml:"from_email"` // the which account to send the email
}

var (
	Pool *email.Pool

	timeout   = time.Second * 10
	fromEmail = ""
	smtpAddr  = "smtp.qq.com:25"
	smtpHost  = "smtp.qq.com"
)

// New create an email http request pool to send email.
func New(opts Option) (*email.Pool, error) {
	p, err := email.NewPool(
		smtpAddr,
		100, // thread pool cap
		smtp.PlainAuth("",
			opts.FromEmail,
			opts.SMTPKey,
			smtpHost,
		))

	if err != nil {
		return nil, err
	}

	Pool = p
	fromEmail = opts.FromEmail
	return Pool, nil
}

func Send(toEmail []string, subject string, text []byte) error {
	if Pool == nil {
		return errors.New("pool haven't init")
	}

	e := email.NewEmail()
	e.To = toEmail
	e.From = fromEmail
	e.Subject = subject
	e.Text = text

	return Pool.Send(e, timeout)
}
