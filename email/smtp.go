// simple mail transfer protocol demo based on "net/smtp"
package main

import (
	"fmt"
	"net/smtp"
	"strconv"
	"strings"

	"instance.golang.com/utils"
)

func Smtp() {
	// authentication config
	ec := NewEmailConfig(
		"example@163.com", // email server host
		25,                // email server port
		"example@163.com", // your email account
		"******",          // authorization code not password
	)
	// plain authentication
	auth := smtp.PlainAuth("", ec.Username, ec.Password, ec.Host)

	// sender & receivers
	sender := "example@163.com"
	receivers := []string{
		"xxx@qq.com",
		"xxx@163.com",
		// ...
	}

	// email
	subject := utils.Now() + " [Hello]"
	body := "Verification code: " + utils.RandNumMath() + "\n\nThanks!"
	contentType := "Content-Type: text/plain; charset=UTF-8" // text/html
	message := []byte(
		"To: " + strings.Join(receivers, ",") + "\r\n" +
			"Form: " + "<" + ec.Username + ">" + "\r\n" +
			"Subject: " + subject + "\r\n" + contentType +
			"\r\n\r\n" +
			body,
	)

	// send email
	fmt.Println("start sending test mail...")
	err := smtp.SendMail(ec.Host+":"+strconv.Itoa(ec.Port), auth, sender, receivers, message)
	if err != nil {
		fmt.Printf("send failed: %v\n", err)
	} else {
		fmt.Println("send success")
	}
}
