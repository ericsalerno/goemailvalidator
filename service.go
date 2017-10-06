package goemailvalidator

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Service is the service listener for email validation
type Service struct {
	config *Configuration
}

// Listen for connections and respond
func (service *Service) Listen(config *Configuration) {
	service.config = config

	http.Handle("/validate", service)

	serverInfo := fmt.Sprintf(":%d", config.Port)
	log.Fatal(http.ListenAndServe(serverInfo, nil))
}

func (service *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := request{}

	if r.PostFormValue("email") == "" {
		log.Output(0, "Failed to process request without email post value.")

		response := service.getResponseError(&request, "You must post an email address with the variable name 'email'.")
		service.printOutput(w, response)
		return
	}

	request.inputEmail = r.PostFormValue("email")

	atPos := strings.Index(request.inputEmail, "@")

	if atPos == -1 {
		response := service.getResponseError(&request, "Invalid email, no @ found.")
		service.printOutput(w, response)
		return
	}

	request.inputUser = request.inputEmail[0:atPos]
	request.inputHost = request.inputEmail[atPos+1:]

	response := service.getResponseOutput(&request, true)
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

func (service *Service) getResponseError(req *request, errorString string) *Response {
	r := Response{}
	r.Status = 500
	r.Message = errorString
	r.Email = req.inputEmail
	r.Host = req.inputHost
	r.User = req.inputUser

	return &r
}

func (service *Service) getResponseOutput(req *request, isValid bool) *Response {
	r := Response{}
	r.Status = 200
	r.Message = "OK"
	r.Email = req.inputEmail
	r.Valid = isValid
	r.Host = req.inputHost
	r.User = req.inputUser

	return &r
}
