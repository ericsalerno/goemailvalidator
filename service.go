package goemailvalidator

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Service is the service listener for email validation
type Service struct {
	config     *Configuration
	inputEmail string
}

// Listen for connections and respond
func (service *Service) Listen(config *Configuration) {
	service.config = config

	http.Handle("/validate", service)

	serverInfo := fmt.Sprintf(":%d", config.Port)
	log.Fatal(http.ListenAndServe(serverInfo, nil))
}

func (service *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.PostFormValue("email") == "" {
		log.Output(0, "Failed to process request without email post value.")

		response := service.getResponseError("You must post an email address with the variable name 'email'.")
		service.printOutput(w, response)
		return
	}

	service.inputEmail = r.PostFormValue("email")

	response := service.getResponseOutput(true)
	service.printOutput(w, response)
}

func (service *Service) printOutput(w http.ResponseWriter, r *Response) {
	output, err := json.Marshal(r)

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(r.Status)
	w.Write(output)
}

func (service *Service) getResponseError(errorString string) *Response {
	r := Response{}
	r.Status = 500
	r.Message = errorString
	r.Email = service.inputEmail

	return &r
}

func (service *Service) getResponseOutput(isValid bool) *Response {
	r := Response{}
	r.Status = 200
	r.Message = "OK"
	r.Email = service.inputEmail
	r.Valid = isValid

	return &r
}
