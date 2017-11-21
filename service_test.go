package main

import "testing"

func TestServerCreate(t *testing.T) {
	c := Configuration{
		Port: 8081,
	}

	b := Service{
		Config: &c,
	}

	if b.Config.Port != 8081 {
		t.Fatal("Missing config value (somehow)")
	}
}

func TestServerRegexes(t *testing.T) {
	b := Service{}

	b.buildRegularExpressions()

	// [a-zA-Z0-9\.-]+
	if b.validEmailHost.MatchString("test-HOSTval124id.net") == false {
		t.Fatal("Invalid email host regex (plain)!")
	}

	if b.validEmailHost.MatchString("localhost") == false {
		t.Fatal("Invalid email host regex (local)!")
	}

	if b.validEmailHost.MatchString("87dfa@$22") == true {
		t.Fatal("Invalid email host regex (invalid characters)!")
	}

	if b.validEmailHost.MatchString("87dfa@$22.lol") == true {
		t.Fatal("Invalid email host regex (invalid characters)!")
	}

	// ^\d{1-3}\.\d{1-3\.\d{1-3}\.\d{1-3}$
	if b.validEmailHostIP.MatchString("12.14.156.255") == false {
		t.Fatal("Invalid host IP regex (plain)!")
	}

	if b.validEmailHostIP.MatchString("12.14.156") == true {
		t.Fatal("Invalid host IP regex (missing octet)!")
	}

	if b.validEmailHostIP.MatchString("12.14.156.123.554") == true {
		t.Fatal("Invalid host IP regex (extra octet)!")
	}

	// [a-zA-Z0-9!#$%&'*+/=\?^_\{\}|~\.-]+
	if b.validEmailUser.MatchString("geezer123") == false {
		t.Fatal("Invalid user regex (plain)!")
	}

	if b.validEmailUser.MatchString("") == true {
		t.Fatal("Invalid user regex (empty)!")
	}

	if b.validEmailUser.MatchString("b&l*a4{#'1_+r=gh?12}3~-.|^as%df$!") == false {
		t.Fatal("Invalid user regex (special characters)!")
	}
}
