package goemailvalidator

// Configuration for a service
type Configuration struct {
	Port int

	ValidateDisposable bool
	ValidateRegex      bool
	ValidateTest       bool
}
