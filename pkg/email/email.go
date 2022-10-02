/**
@author: yeebing
@date: 2022/9/27
**/

package email

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

type Option struct {
	SMTPKey         string `yaml:"smtp_key"`
	FromEmail       string `yaml:"from_email"`        // the which account to send the email
	EmailServerAddr string `yaml:"email_server_addr"` // the email address (domain/ip + port)
	EmailServerHost string `yaml:"email_server_host"` // the email ip host (domain/ip)
	PoolCap         int    `yaml:"pool_cap"`          // the capacity of http pool
}

var (
	Pool *email.Pool
)

// New create an email http request pool to send email.
func New(opts Option) (*email.Pool, error) {
	p, err := email.NewPool(opts.EmailServerAddr, opts.PoolCap, smtp.PlainAuth("", opts.FromEmail, opts.SMTPKey, opts.EmailServerHost))
	if err != nil {
		return nil, err
	}

	Pool = p
	return Pool, nil
}
