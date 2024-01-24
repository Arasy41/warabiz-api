package smtp

// func (data *MailData) SendEmail(cfg *config.SMTPAccount) error {

// 	smtpAddr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
// 	auth := smtp.PlainAuth("", data.Sender, cfg.Password, cfg.Host)

// 	subject := data.Subject
// 	htmlBody := data.HtmlBody
// 	plainTextBody := data.PlainTextBody

// 	// Create the plain text MIME part
// 	plainText := fmt.Sprintf("Content-Type: text/plain; charset=utf-8\r\nContent-Transfer-Encoding: quoted-printable\r\n\r\n%s\r\n", plainTextBody)

// 	// Create the HTML MIME part
// 	html := fmt.Sprintf("Content-Type: text/html; charset=utf-8\r\nContent-Transfer-Encoding: quoted-printable\r\n\r\n%s\r\n", htmlBody)

// 	// Create the full email message
// 	message := fmt.Sprintf("Subject: %s\r\nMIME-Version: 1.0\r\n\r\n%s%s\r\n", subject, plainText, html)

// 	if len(data.Files) != 0 {
// 		attachmentMIMEs, err := data.getAttachmentMIMEs(data.Files...)
// 		if err != nil {
// 			return err
// 		}
// 		message += "Content-Type: multipart/mixed; boundary=boundary\r\n\r\n" + "--boundary\r\n"
// 		for _, attachmentMIME := range attachmentMIMEs {
// 			message += fmt.Sprintf("--boundary\r\n%s\r\n", attachmentMIME)
// 		}
// 		message += "--boundary--"
// 	}

// 	if err := smtp.SendMail(smtpAddr, auth, data.Sender, []string{data.To}, []byte(message)); err != nil {
// 		return err
// 	}

// 	fmt.Println("hahahah")

// 	return nil
// }
