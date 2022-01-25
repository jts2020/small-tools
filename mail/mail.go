package mail

import (
	"crypto/tls"
	"small-tools/conf"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(mailTo []string, subject string, body string) error {
	mailInfo := conf.MyConfig.Mymap
	pix := "mail" + conf.Middle
	mailConn := map[string]string{
		"subject": mailInfo[pix+"subject"],
		"name":    mailInfo[pix+"name"],
		"user":    mailInfo[pix+"user"],
		"pass":    mailInfo[pix+"pass"],
		"host":    mailInfo[pix+"host"],
		"port":    mailInfo[pix+"port"],
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], mailConn["name"])) //这种方式可以添加别名，即“XX官方”
	m.SetHeader("To", mailTo...)                                             //发送给多个用户
	m.SetHeader("Subject", subject)                                          //设置邮件主题
	m.SetBody("text/html", body)                                             //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	err := d.DialAndSend(m)
	return err

}
