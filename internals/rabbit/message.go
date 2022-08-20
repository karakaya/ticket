package rabbit

type Message struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Body  string `json:"body"`
}
