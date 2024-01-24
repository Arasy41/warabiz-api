package smtp

// import (
// 	"encoding/base64"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"strings"
// )

// type MailData struct {
// 	Sender        string
// 	To            string
// 	Subject       string
// 	PlainTextBody string
// 	HtmlBody      string
// 	Files         []File
// }

// type File struct {
// 	Name string
// 	File io.Reader
// }

// func (data *MailData) HtmlBuilder(templatePath string, values map[string]interface{}) error {

// 	//* Read Template File
// 	template, err := ioutil.ReadFile(templatePath)
// 	if err != nil {
// 		return err
// 	}

// 	templateStr := Replacer(string(template), values)
// 	data.HtmlBody = templateStr
// 	return nil
// }

// func (data *MailData) PlainTextBodyBuilder(templatePath string, values map[string]interface{}) error {

// 	//* Read Template File
// 	template, err := ioutil.ReadFile(templatePath)
// 	if err != nil {
// 		return err
// 	}

// 	templateStr := Replacer(string(template), values)
// 	data.PlainTextBody = templateStr
// 	return nil
// }

// func Replacer(template string, values map[string]interface{}) string {
// 	for key, value := range values {
// 		template = strings.ReplaceAll(template, key, fmt.Sprintf("%v", value))
// 	}
// 	return template
// }

// func (data *MailData) getAttachmentMIMEs(files ...File) ([]string, error) {

// 	var attachmentMIMEs []string

// 	for i := 0; i < len(files); i++ {
// 		attachment, err := ioutil.ReadAll(files[i].File)
// 		if err != nil {
// 			return nil, err
// 		}

// 		attachmentContentType := "application/octet-stream"
// 		attachmentContentDisposition := fmt.Sprintf("attachment; filename=\"%s\"", files[i].Name)
// 		attachmentMIME := fmt.Sprintf("Content-Type: %s\r\nContent-Disposition: %s\r\nContent-Transfer-Encoding: base64\r\n\r\n%s\r\n",
// 			attachmentContentType, attachmentContentDisposition, encodeBase64(attachment))

// 		attachmentMIMEs = append(attachmentMIMEs, attachmentMIME)
// 	}
// 	return attachmentMIMEs, nil
// }

// func encodeBase64(data []byte) string {
// 	return strings.TrimRight(base64.StdEncoding.EncodeToString(data), "\r\n")
// }
