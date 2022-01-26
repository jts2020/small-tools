package mail

import (
	"crypto/tls"
	"fmt"
	"small-tools/conf"

	"gopkg.in/gomail.v2"
)

func sendMail(mailTo []string, subject string, body string) error {
	mailInfo := conf.Ymlconf.Mail

	msg := gomail.NewMessage()

	msg.SetHeader("From", msg.FormatAddress(mailInfo.User, mailInfo.Name)) //这种方式可以添加别名，即“XX官方”
	msg.SetHeader("To", mailTo...)                                         //发送给多个用户
	msg.SetHeader("Subject", mailInfo.Subject)                             //设置邮件主题
	msg.SetBody("text/html", body)                                         //设置邮件正文

	d := gomail.NewDialer(mailInfo.Host, mailInfo.Port, mailInfo.User, mailInfo.Pass)

	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	err := d.DialAndSend(msg)
	return err

}

func SendMail(body string, mailFlag bool) {
	if !mailFlag {
		return
	}
	fmt.Println("send body:", body)
	enable := conf.Ymlconf.Mail.Enable
	if !enable {
		return
	}
	mailTo := conf.Ymlconf.Mail.MailTo
	fmt.Println("mailTo:", mailTo)
	subject := conf.Ymlconf.Mail.Subject

	err := sendMail(mailTo, subject, body)
	if err != nil {
		fmt.Println(err)
		fmt.Println("send fail")
		return
	}

	fmt.Println("send successfully")

}
