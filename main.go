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
	var mailFlag bool

	watch.HeadBody(&body)

	watch.WatchCPU(&body, &mailFlag)

	watch.WatchDisk(&body, &mailFlag)

	watch.WatchCmd(&body, &mailFlag)

	mail.SendMail(body, mailFlag)

}
