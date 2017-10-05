package goemailvalidator

// Response output for json marshalling
type Response struct {
	status  int
	message string
	email   string
	valid   bool
}
