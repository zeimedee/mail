package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	sendmail "github.com/zeimedee/mail/sendMail"
)

func main() {
	Email := os.Getenv("EMAIL")
	Pass := os.Getenv("PASSWORD")
	if Email == "" || Pass == "" {
		err := godotenv.Load()
		sendmail.Check(err)
	}

	em := os.Getenv("EMAIL")
	pw := os.Getenv("PASSWORD")

	Email = em
	Password := pw
	Recipient := em
	CC := []string{em}

	path, err := os.Getwd()
	sendmail.Check(err)

	sender := sendmail.NewSender(Email, Password)

	Msg, err := sender.WriteMessage("alex", path+"/mailTemplate/mail.html")
	sendmail.Check(err)

	body := sender.WriteEmail(Recipient, "alex", Msg, CC)

	mail, err := sender.Mail("alex", string(body), Recipient)
	if err != nil {
		panic(err)
	}
	fmt.Println(mail)
}
