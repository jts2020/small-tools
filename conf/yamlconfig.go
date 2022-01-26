package conf

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type appconf struct {
	Version string `yaml:"version"`
	Cron    string `yaml:"cron"`
}

type watchconf struct {
	Disks         []string `yaml:"disks"`
	DiskThreshold int      `yaml:"diskThreshold"`
	CpuThreshold  int      `yaml:"cpuThreshold"`
}

type cmdconf struct {
	Name  string   `yaml:"name"`
	Arg1  string   `yaml:"arg1"`
	Paths []string `yaml:"paths"`
}

type mailconf struct {
	Enable  bool     `yaml:"enable"`
	Subject string   `yaml:"subject"`
	Name    string   `yaml:"name"`
	User    string   `yaml:"user"`
	Pass    string   `yaml:"pass"`
	Host    string   `yaml:"host"`
	Port    int      `yaml:"port"`
	MailTo  []string `yaml:"mailTo"`
}

type ymlconf struct {
	App   appconf   `yaml:"app"`
	Watch watchconf `yaml:"watch"`
	Cmd   cmdconf   `yaml:"cmd"`
	Mail  mailconf  `yaml:"mail"`
}

func (c *ymlconf) getConf(pathFile string) *ymlconf {
	ymlFile, err := ioutil.ReadFile(pathFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(ymlFile, c)
	return c
}

var Ymlconf ymlconf

func init() {
	pathFile := "config/conf.yml"
	if len(os.Args) > 1 && len(os.Args[1]) != 0 {
		pathFile = os.Args[1]
	}
	Ymlconf.getConf(pathFile)
}
