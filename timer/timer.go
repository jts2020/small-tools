package timer

import (
	"small-tools/conf"

	"github.com/robfig/cron"
)

type funcHandle func()

func Handle(handle funcHandle) {
	c := cron.New()
	c.AddFunc(conf.Ymlconf.App.Cron, handle)
	c.Start()
	defer c.Stop()
	select {}
}
