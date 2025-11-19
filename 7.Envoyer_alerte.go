package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	from := "tonmail@gmail.com"
	password := "tonpassword"
	to := []string{"destinataire@gmail.com"}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("Subject: Alerte\r\n\r\nCeci est un test d'alerte")

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println("Erreur:", err)
	} else {
		fmt.Println("Mail envoyé")
	}
}
