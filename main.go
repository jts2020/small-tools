package main

import (
	"fmt"
	"small-tools/conf"
	"small-tools/mail"
	"small-tools/timer"
	"small-tools/watch"
)

func main() {
	fmt.Printf("init Conf:%v\n", conf.Ymlconf)

	timer.Handle(handle)

}

func handle() {
	var body string

	watch.HeadBody(&body)

	watch.WatchCPU(&body)

	watch.WatchDisk(&body)

	watch.WatchCmd(&body)

	mail.SendMail(body, watch.MailFlag)

}
