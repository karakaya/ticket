package sender

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	gomail "gopkg.in/gomail.v2"
)

type Message struct {
	ID        uuid.UUID `json:"_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func SendEmail(msg []byte) {
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

	var message Message
	json.Unmarshal(msg, &message)

	e := gomail.NewMessage()
	e.SetHeader("From", from)
	e.SetHeader("To", message.Email)
	e.SetHeader("Subject", "Your ticket has beeen created")

	body := fmt.Sprintf(`
	<b> Ticket ID </b> %s <br>
	<b> Title </b> %s </b> <br>
	<b> Created At </b> %s <br>
	<b> Body </b> %s <br>
	`, message.ID, message.Title, message.CreatedAt, message.Body)

	e.SetBody("text/html", body)

	n := gomail.NewDialer(host, 587, user, password)

	if err := n.DialAndSend(e); err != nil {
		log.Fatalf("err to send email %v \n", err)
	}

}
