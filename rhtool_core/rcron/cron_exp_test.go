package rcron

import (
	"fmt"
	"testing"
	"time"
)

func TestCronByEvery(t *testing.T) {
	cron := NewCron(false)
	cron.TaskListByEvery = []*CronTask{
		{
			1, "", func() {
				fmt.Println("10s")
			},
		},
		{
			2, "", func() {
				fmt.Println("20s")
			},
		},
	}
	cron.AddFuncByEvery()

	cron.Start()
	fmt.Println("dfasfdf: ", cron.EntryIds)
	time.Sleep(3 * time.Second)
	cron.Remove(cron.EntryIds[1])
	time.Sleep(5 * time.Second)
	fmt.Println("dfasfdf: ", cron.EntryIds)
	cron.Stop()
}

func TestCronByCronStr(t *testing.T) {

	fmt.Println(time.Now())

	cron := NewCron(true)
	cron.TaskListByCronStr = []*CronTask{
		// 每一秒
		{
			0, "0/1 * * * * ?", func() {
				fmt.Println("10s")
			},
		},
		// 每两秒
		{
			0, "0/2 * * * * ?", func() {
				fmt.Println("20s")
			},
		},
		// 每天15点25分
		{
			0, "0 25 15 * * ?", func() {
				fmt.Println("57 11 * * * ?")
			},
		},
	}
	err := cron.AddFuncByCronStr()
	fmt.Println(err)

	cron.Start()
	time.Sleep(1000 * time.Second)
	cron.Stop()
}
