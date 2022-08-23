package emailsender

import (
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	gomail "gopkg.in/gomail.v2"
)

type Message struct {
	ID        uuid.UUID
	Title     string
	Body      string
	CreatedAt time.Time
}

func SendEmail(msg Message) {
	if os.Getenv("EMAIL_HOST") == "" {
		log.Fatal("define EMAIL_HOST env")
	}
	if os.Getenv("EMAIL_USER") == "" {
		log.Fatal("define EMAIL_USER env")
	}
	if os.Getenv("EMAIL_PASSWORD") == "" {
		log.Fatal("define EMAIL_PASSWORD env")
	}

	host := os.Getenv("EMAIL_HOST")
	user := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASSWORD")
	from := os.Getenv("EMAIL_FROM")

	e := gomail.NewMessage()
	e.SetHeader("From", from)
	e.SetHeader("To", "to@example.com")
	e.SetHeader("Subject", "Your ticket has beeen created")
	e.SetBody("text/html", "<b>This is the body of the mail</b>")

	n := gomail.NewDialer(host, 587, user, password)

	if err := n.DialAndSend(e); err != nil {
		log.Fatalf("err to send email %v \n", err)
	}

}
