package watch

import (
	"os/exec"
	"small-tools/conf"
	"syscall"
)

func WatchCmd(body *string) {
	name := conf.Ymlconf.Cmd.Name
	arg1 := conf.Ymlconf.Cmd.Arg1
	cmdPaths := conf.Ymlconf.Cmd.Paths
	for i := range cmdPaths {
		cmdPath := cmdPaths[i]
		if len(cmdPath) != 0 {
			command := exec.Command(name, arg1, cmdPath) //初始化Cmd
			buf, _ := command.Output()
			code := command.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
			if code != 0 {
				MailFlag = true
				*body += "<p>[" + cmdPath + "]:<br/>" + string(buf) + "</p>"
			}
		}
	}

}
