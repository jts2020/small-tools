package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"small-tools/conf"
	_ "small-tools/conf"
	"small-tools/mail"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
)

var confMap map[string]string
var flag bool

type funcHandle func()

func main() {
	confMap = conf.MyConfig.Mymap
	fmt.Printf("init Conf:%v\n", confMap)

	TimeTask(handle)

}

func handle() {
	body := "<p><b>" + time.Now().Format("2006-01-02 15:04:05")
	info, _ := host.Info()
	body = body + "[" + info.Hostname + "]"
	body = body + "[" + getIp() + "]异常情况</b></p><hr/>"

	watchCPU(&body)

	watchDisk(&body)

	watchCmd(&body)

	sendMail(body)

}

func watchCPU(body *string) {
	cpuThreshold, _ := strconv.Atoi(confMap["watch"+conf.Middle+"cpuThreshold"])
	percent, _ := cpu.Percent(time.Second, false)
	usage := int(percent[0])
	if usage > cpuThreshold {
		flag = true
		usageStr := strconv.Itoa(usage)
		*body = *body + "<p>CPU usage:" + usageStr + "%</p>"
	}

}

func watchDisk(body *string) {
	diskThreshold, _ := strconv.Atoi(confMap["watch"+conf.Middle+"diskThreshold"])
	diskPaths := strings.Split(confMap["watch"+conf.Middle+"disk"], ",")
	for i := range diskPaths {
		diskPath := diskPaths[i]
		if len(diskPath) != 0 {
			diskInfo, _ := disk.Usage(diskPath)
			usage := int(diskInfo.UsedPercent)
			if usage > diskThreshold {
				flag = true
				usageStr := strconv.Itoa(usage)
				*body = *body + "<p>[" + diskPath + "]" + " disk usage:" + usageStr + "%</p>"
			}
		}
	}

}

func sendMail(body string) {
	if !flag {
		return
	}
	fmt.Println("send body:", body)
	enable := confMap["mail"+conf.Middle+"enable"]
	if strings.EqualFold(enable, "false") {
		return
	}
	mailTo := strings.Split(confMap["mail"+conf.Middle+"mailTo"], ",")
	fmt.Println("mailTo:", mailTo)
	subject := confMap["mail"+conf.Middle+"subject"]

	err := mail.SendMail(mailTo, subject, body)
	if err != nil {
		fmt.Println(err)
		fmt.Println("send fail")
		return
	}

	fmt.Println("send successfully")

}

func TimeTask(handle funcHandle) {
	var ch chan int
	//定时任务
	ti, _ := strconv.Atoi(confMap["app"+conf.Middle+"time"])
	ticker := time.NewTicker(time.Duration(ti) * time.Second)
	go func() {
		for range ticker.C {
			handle()
		}
		ch <- 1
	}()
	<-ch
}

func watchCmd(body *string) {
	name := confMap["cmd"+conf.Middle+"name"]
	arg1 := confMap["cmd"+conf.Middle+"arg1"]
	cmdPaths := strings.Split(confMap["cmd"+conf.Middle+"path"], ",")
	for i := range cmdPaths {
		cmdPath := cmdPaths[i]
		if len(cmdPath) != 0 {
			command := exec.Command(name, arg1, cmdPath) //初始化Cmd
			buf, _ := command.Output()
			code := command.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
			if code != 0 {
				flag = true
				*body = *body + "<p>[" + cmdPath + "]:<br/>" + string(buf) + "</p>"
			}
		}
	}

}

func getIp() string {

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}
	return "127.0.0.1"

}
