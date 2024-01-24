package smtp

// import (
// 	"strings"
// 	"warabiz/api/config"
// 	"time"

// 	"gopkg.in/mail.v2"
// )

// //* Functions used to send mail with the configurations
// func (data *MailData) SendGoEmail(cfg *config.SMTPAccount) error {

// 	dialer := mail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)
// 	dialer.Timeout = cfg.Timeout * time.Minute

// 	msg := mail.NewMessage()
// 	msg.SetHeader("To", data.To)
// 	msg.SetHeader("Subject", data.Subject)
// 	msg.SetHeader("From", data.Sender)
// 	msg.SetBody("text/plain", data.PlainTextBody)
// 	data.HtmlBody = strings.ReplaceAll(data.HtmlBody, "{{.PlainText}}", data.PlainTextBody)
// 	msg.AddAlternative("text/html", data.HtmlBody)

// 	for _, file := range data.Files {
// 		msg.AttachReader(file.Name, file.File)
// 	}

// 	return dialer.DialAndSend(msg)
// }
