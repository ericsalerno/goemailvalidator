package main

import "testing"

func TestEmailParsing(t *testing.T) {
	r := request{}

	r.buildFromEmail("testing@email.com")

	if r.inputEmail != "testing@email.com" {
		t.Fatal("Invalid input email found.")
	}

	if r.inputHost != "email.com" {
		t.Fatal("Invalid host found")
	}

	if r.inputUser != "testing" {
		t.Fatal("Invalid user found.")
	}
}

func TestValidEmailUserValidation(t *testing.T) {
	individualValidEmailTest(t, "testing@email.com")
	individualValidEmailTest(t, "somet!$%*hing@localhost")
	individualValidEmailTest(t, "a@a")
	individualValidEmailTest(t, "helmet@123.224.123.255")
	individualValidEmailTest(t, "blargh@frufru.net.liver-disaster.com")
}

func individualValidEmailTest(t *testing.T, email string) {
	s := Service{}
	s.buildRegularExpressions()

	r := request{}
	r.buildFromEmail(email)

	complete := make(chan bool, 2)
	go r.validateUser(complete, s.validEmailUser)
	go r.validateHost(complete, s.validEmailHost, s.validEmailHostIP)
	<-complete
	<-complete

	if r.validUser == false {
		t.Fatal("Invalid user validation: " + email)
	}

	if r.validHost == false {
		t.Fatal("Invalid host validation: " + email)
	}
}

func TestInvalidEmailUserValidation(t *testing.T) {
	individualInvalidEmailTest(t, "test ing@email.com")
	individualInvalidEmailTest(t, "somet!$%*hing@123.1&556.734")
	individualInvalidEmailTest(t, "a@")
	individualInvalidEmailTest(t, "")
	individualInvalidEmailTest(t, "@123.224.123.255")
	individualInvalidEmailTest(t, "blargh@frufru.n et.liver-disaster.com")
}

func individualInvalidEmailTest(t *testing.T, email string) {
	s := Service{}
	s.buildRegularExpressions()

	r := request{}
	r.buildFromEmail(email)

	complete := make(chan bool, 2)
	go r.validateUser(complete, s.validEmailUser)
	go r.validateHost(complete, s.validEmailHost, s.validEmailHostIP)
	<-complete
	<-complete

	if r.validUser == true && r.validHost == true {
		t.Fatal("False positive validation: " + email)
	}
}
