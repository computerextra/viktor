package mail

import (
	"crypto/tls"
	"fmt"

	"viktor/db"

	"gopkg.in/gomail.v2"
)

type PaypalMail struct {
	Mitarbeiter  db.Mitarbeiter
	Benutzername string
	Betrag       string
}

func SendPaypalMail(props PaypalMail, server string, port int, user string, pass string, from string) error {
	body := "<!DOCTYPE html><html lang=\"de\"><head><meta charset=\"UTF-8\" />"
	body = fmt.Sprintf("%s<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\" />", body)
	body = fmt.Sprintf("%s<title>Paypal Abrechnung</title></head><body> <h1>Paypal Abrechnung</h1>", body)
	body = fmt.Sprintf("%s<p>Bitte Bezahle %s€ über Paypal.</p>", body, props.Betrag)
	body = fmt.Sprintf("%s<p>Link: <a href=\"http://paypal.me/%s/%s\">", body, props.Benutzername, props.Betrag)
	body = fmt.Sprintf("%shttp://paypal.me/%s/%s</a></p></body></html>", body, props.Benutzername, props.Betrag)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", *props.Mitarbeiter.Email)
	m.SetHeader("Subject", "PayPal Abrechnung")
	m.SetBody("text/html", body)

	d := gomail.NewDialer(server, port, user, pass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return d.DialAndSend(m)
}
