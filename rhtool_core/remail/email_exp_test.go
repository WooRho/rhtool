package remail

import (
	"fmt"
	"testing"
)

func TestEmail(t *testing.T) {
	form := "13421100714@163.com"
	userName := "13421100714@163.com"
	password := "ZHACDWEHOMYGOVIN"
	email := NewREmail(form, userName, password)

	stmp := "smtp.163.com"
	to := []string{"935549961@qq.com", "W13421100714@outlook.com"}
	//cc := []string{"13421100714@163.com"}
	text := []byte("dsaffsafasdfasf")
	html := []byte("")

	email.SetContent("靓仔", stmp, to, nil, text, html)
	err := email.Send()
	fmt.Println(err)
}
