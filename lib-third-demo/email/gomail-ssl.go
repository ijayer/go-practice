package main

import (
	"crypto/tls"
	"fmt"
	"log"

	"gopkg.in/gomail.v2"
)

// Email
type Email struct {
	Username  string   // 服务器账号
	Password  string   // 服务器授权码
	Host      string   // 服务器地址
	Port      int      // 服务器端口
	Subject   string   // 邮件主题
	Body      string   // 邮件内容
	Receivers []string // 收信人
}

func GoMailWithSSL() {
	e := NewEmail()
	if err := e.SendEmail(); err != nil {
		log.Fatal(err)
	}
}

func NewEmail() *Email {
	return &Email{
		Username:  "xxx xxx",
		Password:  "xxx xxx",
		Host:      "xxx xxx",
		Port:      465,
		Body:      "Go email test with ssl",
		Subject:   fmt.Sprintf("%s", "xxx"),
		Receivers: []string{"one@qq.com"},
	}
}

func (e *Email) dialer() *gomail.Dialer {
	d := gomail.NewDialer(e.Host, e.Port, e.Username, e.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	d.SSL = true
	return d
}

func (e *Email) SendEmail() error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.Username)
	m.SetHeader("To", e.Receivers...)
	m.SetHeader("Subject", e.Subject)
	m.SetBody("text/plain", e.Body)
	return e.dialer().DialAndSend(m)
}
