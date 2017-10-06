package goemailvalidator

// Response output for json marshalling
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Email   string `json:"email"`
	Valid   bool   `json:"valid"`

	Host string `json:"host"`
	User string `json:"user"`
}
