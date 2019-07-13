package mailutil

import (
	"log"

	gomail "gopkg.in/gomail.v2"
)

//Configuration settings for sending mail
type configuration struct {
	smtpServer   string
	smtpPort     int
	smtpAccount  string
	smtpPassword string
	mailFrom     string
	mailTo       string
}

//MailConfig holds the settings when sending email via mailutil
var MailConfig configuration

// SetConfiguration for sending email
func SetConfiguration(smtpServer string, port int, account string, password string, mailFrom string, mailTo string) {
	config := configuration{}
	config.smtpServer = smtpServer
	config.smtpPort = port
	config.smtpAccount = account
	config.smtpPassword = password
	config.mailFrom = mailFrom
	config.mailTo = mailTo
	MailConfig = config
}

//SendMail provides a quick way to send a message to recipients
func SendMail(subject string, message string) {
	if &MailConfig == nil {
		log.Panicln("Need to call SetConfiguration first.")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", MailConfig.mailFrom)
	m.SetHeader("To", MailConfig.mailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", message)

	dial := gomail.NewDialer(MailConfig.smtpServer, MailConfig.smtpPort, MailConfig.smtpAccount, MailConfig.smtpPassword)
	err := dial.DialAndSend(m)
	if err != nil {
		log.Print("Error sending mail.")
		log.Println(err)
	}

}
