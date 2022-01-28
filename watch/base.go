package watch

import (
	"fmt"
	"net"
	"os"
	"small-tools/conf"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
)

func HeadBody(body *string) {
	*body += "<p><b>" + time.Now().Format("2006-01-02 15:04:05")
	info, _ := host.Info()
	*body += "[" + info.Hostname + "][" + GetIp() + "]异常情况</b></p><hr/>"
}

func WatchDisk(body *string, mailFlag *bool) {
	diskThreshold := conf.Ymlconf.Watch.DiskThreshold
	diskPaths := conf.Ymlconf.Watch.Disks
	for i := range diskPaths {
		diskPath := diskPaths[i]
		if len(diskPath) != 0 {
			diskInfo, _ := disk.Usage(diskPath)
			usage := int(diskInfo.UsedPercent)
			if usage > diskThreshold {
				*mailFlag = true
				usageStr := strconv.Itoa(usage)
				*body += "<p>[" + diskPath + "]" + " disk usage:" + usageStr + "%</p>"
			}
		}
	}

}

func WatchCPU(body *string, mailFlag *bool) {
	cpuThreshold := conf.Ymlconf.Watch.CpuThreshold
	percent, _ := cpu.Percent(time.Second, false)
	usage := int(percent[0])
	if usage > cpuThreshold {
		*mailFlag = true
		usageStr := strconv.Itoa(usage)
		*body += "<p>CPU usage:" + usageStr + "%</p>"
	}

}

func GetIp() string {
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
