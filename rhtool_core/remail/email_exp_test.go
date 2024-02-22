package remail

import (
	"fmt"
	"testing"
)

func TestEmail(t *testing.T) {
	form := "13********14@163.com"
	userName := "13*******14@163.com"
	password := "*****************"
	stmp := "smtp.163.com"
	port := "25"

	to := []string{"935549961@qq.com", "W13421100714@outlook.com"}
	cc := []string{}
	text := []byte("")
	html := []byte(`
  <ul>
<li><a "https://go-quiz.github.io/2020/01/10/godailylib/flag/">Go 每日一库之 flag</a></li>
<li><a "https://go-quiz.github.io/2020/01/10/godailylib/go-flags/">Go 每日一库之 go-flags</a></li>
<li><a "https://go-quiz.github.io/2020/01/14/godailylib/go-homedir/">Go 每日一库之 go-homedir</a></li>
<li><a "https://go-quiz.github.io/2020/01/15/godailylib/go-ini/">Go 每日一库之 go-ini</a></li>
<li><a "https://go-quiz.github.io/2020/01/17/godailylib/cobra/">Go 每日一库之 cobra</a></li>
<li><a "https://go-quiz.github.io/2020/01/18/godailylib/viper/">Go 每日一库之 viper</a></li>
<li><a "https://go-quiz.github.io/2020/01/19/godailylib/fsnotify/">Go 每日一库之 fsnotify</a></li>
<li><a "https://go-quiz.github.io/2020/01/20/godailylib/cast/">Go 每日一库之 cast</a></li>
</ul>
  `)
	email, err := NewREmail(form, userName, password, stmp, port)
	if err != nil {
		return
	}

	email.SetContent("靓仔", to, cc, text, html)
	err = email.Send()
	fmt.Println(err)
}
