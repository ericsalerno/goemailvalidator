package goemailvalidator

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

// Service is the service listener for email validation
type Service struct {
	config           *Configuration
	validEmailUser   *regexp.Regexp
	validEmailHost   *regexp.Regexp
	validEmailHostIP *regexp.Regexp
}

// Listen for connections and respond
func (service *Service) Listen(config *Configuration) {
	service.config = config
	service.buildRegularExpressions()

	http.Handle("/", service)

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

	request.buildFromEmail(r.PostFormValue("email"))

	if !request.validPreliminary {
		response := service.getResponseError(&request, "Invalid email: "+request.invalidReason)
		service.printOutput(w, response)
		return
	}

	complete := make(chan bool, 3)

	go request.validateHost(complete, service.validEmailHost, service.validEmailHostIP)
	go request.validateUser(complete, service.validEmailUser)
	go request.validateBlackList(complete, service.config)

	<-complete
	<-complete
	<-complete

	response := service.getResponseOutput(&request, request.validHost && request.validUser && request.validPreliminary)
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

func (service *Service) buildRegularExpressions() {
	service.validEmailUser = regexp.MustCompile(`^[a-zA-Z0-9!#$%&'*+/=\?^_\{\}|~\.-]+$`)
	service.validEmailHost = regexp.MustCompile(`^[a-zA-Z0-9!#$%&\.-]+$`)
	service.validEmailHostIP = regexp.MustCompile(`^\d{1-3}\.\d{1-3\.\d{1-3}\.\d{1-3}$`)
}
