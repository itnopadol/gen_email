package model

import (
	"bytes"
	"fmt"
	"net/smtp"
	"html/template"
)

type GenEmail struct {

}

type MailData struct {
	ArNameMail string `json:"ar_name_mail"`
	UrlLink string `json:"url_link"`
}

type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; image/png; charset=\"UTF-8\";\n\n"
)

func (p *GenEmail) GenEmailAuto(access_token string, ar_code string, ar_name string, doc_no string, email string) error {
	subject := "แจ้งใบวางบิล"
	receiver := email
	r := NewRequest([]string{receiver}, subject)

	data := &MailData{}
	data.ArNameMail = ar_name
	data.UrlLink = "http://app.nopadol.com:20000/email/html?ar_code=" + ar_code + "&doc_no=" + doc_no + "&access_token=" +access_token

	t, err := template.ParseFiles("./templates/letter.html")
	if err != nil {
		return err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return err
	}
	r.body = buffer.String()


	//r.body = "http://venus:8099/email/html?ar_code=" + ar_code + "&doc_no=" + doc_no + "&access_token=" +access_token

	body := "To: " + r.to[0] + "\r\nSubject: " + r.subject + "\r\n" + MIME + "\r\n" + r.body
	SMTP := fmt.Sprintf("%s:%d", "smtp.gmail.com", 587)
	if err := smtp.SendMail(SMTP, smtp.PlainAuth("", "nopadol_mailauto@nopadol.com", "[vdw,jwfh2012", "smtp.gmail.com"), "satit@nopadol.com", r.to, []byte(body)); err != nil {
		return err
	}
	return nil
}

func NewRequest(to []string, subject string) *Request {
	return &Request{
		to:      to,
		subject: subject,
	}
}