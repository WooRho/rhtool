package remail

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

type rEmail struct {
	email    *email.Email
	userName string
	password string

	smtp string
}

func NewREmail(from, userName, password string) *rEmail {
	re := &rEmail{
		email:    email.NewEmail(),
		userName: userName,
		password: password,
	}
	re.email.From = from
	return re
}

func (r *rEmail) SetSubject(subject string) {
	r.email.Subject = subject
}

func (r *rEmail) SetSmtp(smtp string) {
	r.smtp = smtp
}

func (r *rEmail) SetTo(to []string) {
	r.email.To = to
}

func (r *rEmail) SetCC(cc []string) {
	r.email.Cc = cc
}

func (r *rEmail) SetText(text []byte) {
	r.email.Text = text
}

func (r *rEmail) SetHtml(html []byte) {
	r.email.HTML = html
}

func (r *rEmail) SetContent(subject, smtp string, to, cc []string, text, html []byte) {
	r.SetSubject(subject)
	r.SetSmtp(smtp)
	r.SetTo(to)
	r.SetCC(cc)
	r.SetText(text)
	r.SetHtml(html)
}

func (r *rEmail) Send() error {
	err := r.email.Send(r.smtp+":25", smtp.PlainAuth("", r.userName, r.password, r.smtp))
	if err != nil {
		return err
	}
	return nil
}
