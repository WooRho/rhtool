package rcron

import (
	"context"
	"github.com/robfig/cron/v3"
	"strconv"
)

type Cron struct {
	c                 *cron.Cron
	ctx               context.Context
	isWithSecond      bool
	TaskListByEvery   []*CronTask
	TaskListByCronStr []*CronTask
	EntryIds          []cron.EntryID // 下标对应
}

type CronTask struct {
	EveryTime int
	CronStr   string
	F         func()
}

func NewCron(isWithSecond bool) *Cron {
	cn := &Cron{
		EntryIds: make([]cron.EntryID, 0),
	}
	if isWithSecond {
		cn.c = cron.New(cron.WithSeconds())
	}
	if !isWithSecond {
		cn.c = cron.New()
	}
	return cn
}

func (c *Cron) addFuncByEvery(everyTime int, f func()) (err error) {
	entryId, err := c.c.AddFunc("@every "+strconv.Itoa(everyTime)+"s", f)
	c.EntryIds = append(c.EntryIds, entryId)
	return
}

func (c *Cron) addFuncCronStr(cronStr string, f func()) (err error) {
	entryId, err := c.c.AddFunc(cronStr, f)
	c.EntryIds = append(c.EntryIds, entryId)
	return
}

func (c *Cron) Start() {
	c.c.Start()
}

func (c *Cron) Stop() {
	c.c.Stop()
}

func (c *Cron) Remove(entryId cron.EntryID) {
	c.c.Remove(entryId)
}

func (c *Cron) AddFuncByEvery() (err error) {
	if len(c.TaskListByEvery) == 0 {
		return
	}
	for _, f := range c.TaskListByEvery {
		c.addFuncByEvery(f.EveryTime, f.F)
	}
	return
}

// @yearly (or @annually)	Run once a year, 	midnight, Jan. 1st			0 0 1 1 *
// @monthly					Run once a month, 	midnight, first of month	0 0 1 * *
// @weekly					Run once a week, 	midnight between Sat/Sun	0 0 * * 0
// @daily (or @midnight)	Run once a day, 	midnight					0 0 * * *
// @hourly					Run once an hour, 	beginning of hour			0 * * * *
func (c *Cron) AddFuncByCronStr() (err error) {
	if len(c.TaskListByCronStr) == 0 {
		return
	}
	for _, f := range c.TaskListByCronStr {
		c.addFuncCronStr(f.CronStr, f.F)
	}
	return
}
