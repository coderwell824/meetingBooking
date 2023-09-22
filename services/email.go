package services

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
)

//MailBoxConfig 邮箱配置
//TODO: 抽离配置项
type MailBoxConfig struct {
	//邮箱标题
	Title string
	//邮箱内容
	Body string
	//收件人列表
	RecipientList []string
	//发件人账号
	Sender string
	//发件人密码
	SPassword string
	//SMTP服务器地址
	SMTPAddr string
	//SMTP端口
	SMTPPort int
}

func SendEmailService(emailCode string, recipientList []string) error {
	var mailConf MailBoxConfig
	mailConf.Title = "邮箱验证"
	
	//发送邮箱内容
	mailConf.Body = "content"
	mailConf.RecipientList = recipientList
	mailConf.Sender = `coderwell@163.com`
	mailConf.SPassword = "IBOVNYGHSUOLDEOJ"
	mailConf.SMTPAddr = `smtp.163.com`
	mailConf.SMTPPort = 25
	
	html := fmt.Sprintf(`<div>
        <div>
            尊敬的用户，您好！
        </div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>你本次的验证码为%s,为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复</p>
        </div>
    </div>`, emailCode)
	
	m := gomail.NewMessage()
	
	// 第三个参数是我们发送者的名称，但是如果对方有发送者的好友，优先显示对方好友备注名
	m.SetHeader(`From`, mailConf.Sender)
	m.SetHeader(`To`, mailConf.RecipientList...)
	m.SetHeader(`Subject`, mailConf.Title)
	m.SetBody(`text/html`, html)
	// m.Attach("./Dockerfile") //添加附件
	err := gomail.NewDialer(mailConf.SMTPAddr, mailConf.SMTPPort, mailConf.Sender, mailConf.SPassword).DialAndSend(m)
	if err != nil {
		log.Fatalf("Send Email Fail, %s", err.Error())
		return err
	}
	log.Printf("Send Email Success")
	return nil
}
