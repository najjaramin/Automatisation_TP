package main

import (
    "net/smtp"
)

func main() {
    from := "expediteur@example.com"
    to := "aminnajjar3@gmail.com"
    subject := "Alerte système"
    body := "Message d'alerte personnalisé"

    msg := "Subject: " + subject + "\n\n" + body
    err := smtp.SendMail("localhost:25", nil, from, []string{to}, []byte(msg))
    if err != nil {
        panic(err)
    }
}