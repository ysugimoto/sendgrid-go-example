package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	apiKey          = os.Getenv("SENDGRID_API_KEY")
	testToAddress   = os.Getenv("SENDGRID_MAIL_TO_ADDRESS")
	testToName      = os.Getenv("SENDGRID_MAIL_TO_NAME")
	testFromName    = os.Getenv("SENDGRID_MAIL_FROM_NAME")
	testFromAddress = os.Getenv("SENDGRID_MAIL_FROM_ADDRESS")
)

func main() {
	client := sendgrid.NewSendClient(apiKey)
	// This mail only has HTML mail, not has text mail
	email := mail.NewV3MailInit(
		mail.NewEmail(testFromName, testFromAddress),
		fmt.Sprintf("This is test mail"),
		mail.NewEmail(testToName, testToAddress),
		mail.NewContent("text/html", "<h1>Confirm you got this mail</h1>"),
	)
	result, err := client.Send(email)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result)
}
