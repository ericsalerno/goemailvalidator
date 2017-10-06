# goemailvalidator

Email validation service written in golang. It will load a huge blacklist file and run the regular expressions and hashmap look-ups for the blacklist in parallel.

This was built mostly for my own learning process with golang.

1. Run the service
2. HTTP POST an "email" variable to http://localhost/
3. read the json it returns

## Example Service Code

This is an example run of the the service using the blacklist.conf file from https://raw.githubusercontent.com/martenson/disposable-email-domains/master/disposable_email_blacklist.conf

    package main

    import "github.com/ericsalerno/goemailvalidator"

    func main() {
        c := goemailvalidator.Configuration{}
        c.Port = 8081
        c.LoadBlacklist("blacklist.conf")

        b := goemailvalidator.Service{}
        b.Listen(&c)
    }

## Example Test Client Code

    package main

    import (
        "encoding/json"
        "fmt"
        "io/ioutil"
        "net/http"
        "net/url"

        "github.com/ericsalerno/goemailvalidator"
    )

    func main() {
        testEmailValidate("", false)
        testEmailValidate("test@test.com", true)
        testEmailValidate("asdfasdf", false)
        testEmailValidate("@test.com", false)
        testEmailValidate("a@a.com", true)
        testEmailValidate("a@localhost", true)
        testEmailValidate("user!@123.456.331.531", true)
        testEmailValidate("blargh@host(*", false)
    }

    func testEmailValidate(email string, expected bool) {
        resp, err := http.PostForm("http://localhost:8081/", url.Values{"email": {email}})

        if err != nil {
            fmt.Printf("Sorry there was an error: %s\n", err)
            return
        }

        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)

        response := goemailvalidator.Response{}
        err = json.Unmarshal(body, &response)

        if err != nil {
            fmt.Printf("Sorry could not unmarshal json response: %s\n", err)
            return
        }

        pass := "FAIL!"
        if response.Valid == expected {
            pass = "pass"
        }

        fmt.Printf("%s - %s = %t (%d - %s)\n", pass, response.Email, response.Valid, resp.StatusCode, response.Message)
    }
