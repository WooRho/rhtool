package remail

import (
	"errors"
	"github.com/jordan-wright/email"
	"net/smtp"
	"sync"
)

type rEmail struct {
	email    *email.Email
	emailCh  chan *email.Email
	emails   []*email.Email
	wg       sync.WaitGroup
	userName string
	password string
	smtp     string
	port     string
	addr     string
	auth     smtp.Auth
	pool     *email.Pool
}

// _smtp   eg: "smtp.163.com"
func NewREmail(from, userName, password, _smtp, port string) (*rEmail, error) {
	re := &rEmail{
		email:    email.NewEmail(),
		userName: userName,
		password: password,
		smtp:     _smtp,
		port:     port,
		addr:     _smtp + ":" + port,
	}
	re.auth = smtp.PlainAuth("", userName, password, _smtp)
	re.email.From = from
	err := re.validate()
	if err != nil {
		return nil, err
	}
	return re, nil
}

func NewEmptyREmail() *rEmail {
	return &rEmail{
		email: email.NewEmail(),
	}
}

func (r *rEmail) Append(subject string, to, cc []string, text, html []byte) {
	e := &email.Email{}
	e.Subject = subject
	e.To = to
	e.Cc = cc
	e.Text = text
	e.HTML = html
	r.emails = append(r.emails, e)
}

func (r *rEmail) SetFrom(from string) {
	r.email.From = from
}

func (r *rEmail) SetPassword(password string) {
	r.password = password
}

func (r *rEmail) SetUserName(userName string) {
	r.userName = userName
}

func (r *rEmail) SetPort(port string) {
	r.port = port
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

func (r *rEmail) SetContent(subject string, to, cc []string, text, html []byte) {
	r.SetSubject(subject)
	r.SetTo(to)
	r.SetCC(cc)
	r.SetText(text)
	r.SetHtml(html)
}

func (r *rEmail) Send() error {
	err := r.email.Send(r.addr, r.auth)
	if err != nil {
		return err
	}
	return nil
}

func (r *rEmail) validate() error {

	if r.smtp == "" {
		return errors.New("need smtp")
	}
	if r.userName == "" {
		return errors.New("need userName")
	}
	if r.password == "" {
		return errors.New("need password")
	}
	if r.port == "" {
		return errors.New("need port")
	}
	return nil
}

// ------------------- Pool ----------------- //
// use NewEmptyREmail()
func (r *rEmail) NewREmailPool(count int) error {
	var err error
	r.pool, err = email.NewPool(r.addr, count, r.auth)
	return err
}

// gnum 并发数
func (r *rEmail) SendUsePool(gnum int) {

}
