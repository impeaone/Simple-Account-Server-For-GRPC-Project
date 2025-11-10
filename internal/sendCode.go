package internal

import (
	"log"
	"net/smtp"
)

func SendCodeToTheEmail(code, email, emailServ, emailServPass string) {

	smtpHost := "smtp.yandex.com"
	smtpPort := "587"
	from := emailServ
	password := emailServPass

	auth := smtp.PlainAuth("", from, password, smtpHost)

	htmlBody := `
				<!DOCTYPE html>
				<html>
				<head>
    				<meta charset="UTF-8">
    				<title>Код подтверждения</title>
				</head>
				<style>
					.code_form {
						width: 20%;
						display: block;
						border-radius: 30px;
						padding: 20px;
						box-shadow: 1px 2px 2px 0px rgba(0, 0, 0, 0.4);
						transition: all 0.3s ease;
					}
					.code_form:hover {
						cursor: pointer;
						box-shadow: 1px 2px 2px 3px rgba(0, 0, 0, 0.8);
						transform: translateY(-3px);
					}
					
				</style>
				<body>
					<div class="code_form">
						<h2>Ваш код подтверждения</h2>
						<p style="font-size: 24px; font-weight: bold; color: #2563eb;">` + code + `</p>
						<p>Никому не сообщайте этот код.</p>
					<div>
				</body>
				</html>
				`
	msg := "From: " + from + "\r\n" +
		"To: " + email + "\r\n" +
		"Subject: Ваш код подтверждения\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		htmlBody

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}
