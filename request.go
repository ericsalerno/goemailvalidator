package goemailvalidator

import (
	"regexp"
	"strings"
)

type request struct {
	inputEmail string
	inputHost  string
	inputUser  string

	validPreliminary bool
	validUser        bool
	validHost        bool

	invalidReason string
}

func (r *request) buildFromEmail(email string) {

	r.inputEmail = email

	atPos := strings.Index(r.inputEmail, "@")

	if atPos == -1 {
		r.invalidReason = "Missing @"
		r.validPreliminary = false
		return
	}

	r.inputUser = r.inputEmail[0:atPos]
	r.inputHost = r.inputEmail[atPos+1:]

	if r.inputUser == "" {
		r.invalidReason = "Missing user"
		r.validPreliminary = false
		return
	}

	if r.inputHost == "" {
		r.invalidReason = "Missing host"
		r.validPreliminary = false
		return
	}

	r.validPreliminary = true
}

func (r *request) validateUser(validUserRegex *regexp.Regexp) {
	r.validUser = validUserRegex.MatchString(r.inputUser)

	if r.validUser == false {
		r.invalidReason = "User " + r.inputUser + " does not appear valid."
	}
}

func (r *request) validateHost(validHostRegex *regexp.Regexp, validHostIPRegex *regexp.Regexp) {
	r.validHost = validHostRegex.MatchString(r.inputHost) || validHostIPRegex.MatchString(r.inputHost)

	if r.validHost == false {
		r.invalidReason = "Host " + r.inputHost + " does not appear valid."
	}
}

func (r *request) validateBlackList(c *Configuration) {
	hostValue, ok := c.HostList[r.inputHost]

	if ok == false {
		return
	}

	if hostValue == 1 {
		r.invalidReason = "Host " + r.inputHost + " found in blacklist."
		r.validHost = false
		return
	}

	r.validHost = true
}
