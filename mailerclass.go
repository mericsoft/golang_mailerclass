package mailerclass

import (
	"encoding/base64"
	"io/ioutil"
	"net/smtp"
)

type MailCon struct { // basic auth connetcion
	Host     string
	Port     string
	UserName string
	Pass     string
}
type Mailer struct {
	Conn          MailCon
	From          string
	To            []string
	Name          string
	Subject       string
	Body          string
	File          string // need ful path
	ContentType   string // for html : text/ftml
	ReturnMessage string
}

func (m Mailer) SendMail() string {
	conn := smtp.PlainAuth("", m.Conn.UserName, m.Conn.Pass, m.Conn.Host) // auth
	mesbody := "From : " + m.From + "\r\n"                                //message headers
	mesbody += "To: " + m.To[0] + "\r\n"
	mesbody += "Subject: " + m.Subject + "\r\n"
	mesbody += "MIME-Version: 1.0\r\n"
	mesbody += "Content-Type: " + m.ContentType + "\r\n"
	if m.File != "" {
		mesbody += "Content-Transfer-Encoding: base64\r\n" // for attachment need base64
		mesbody += "Content-Disposition: attachment;filename=\"" + m.File + "\"\r\n"
		fileread, err := ioutil.ReadFile(m.File) // read file
		if err != nil {
			return err.Error() + "\r\n"
		}
		mesbody += "\r\n" + base64.StdEncoding.EncodeToString(fileread) //atatch file into the message
	}
	mesbody += m.Body + "\r\n" + m.Name
	err := smtp.SendMail(m.Conn.Host+":"+m.Conn.Port, conn, m.From, m.To, []byte(mesbody))
	if err != nil {
		return err.Error() + "\r\n"
	} else {
		return m.ReturnMessage
	}
}
