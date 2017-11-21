# goemailvalidator

Email validation service package written in golang. It will load a huge blacklist file and run the regular expressions and hashmap look-ups for the blacklist in parallel.

This was built mostly for my own learning process with golang.

1. Run the service
2. HTTP POST an "email" variable to http://localhost/
3. read the json it returns

## Running via Docker

The docker container pulls a blacklist.conf file from https://raw.githubusercontent.com/martenson/disposable-email-domains/master/disposable_email_blacklist.conf

    go get github.com/ericsalerno/goemailvalidator
    docker build -t goemailvalidator ~/src/github.com/ericsalerno/goemailvalidator
    docker run --publish 8081:8081 --name emailvalidator --rm goemailvalidator

With the service now running 

## Example Test Client Code

    package main

    import (
        "encoding/json"
        "fmt"
        "io/ioutil"
        "net/http"
        "net/url"
        "os"
    )

    type response struct {
        Status  int    `json:"status"`
        Message string `json:"message"`
        Email   string `json:"email"`
        Valid   bool   `json:"valid"`

        Host string `json:"host"`
        User string `json:"user"`
    }

    func main() {
        if len(os.Args) < 2 {
            fmt.Println("Usage: testemailpost <emailaddress>")
            return
        }

        if validateEmail(os.Args[1]) == true {
            fmt.Println(os.Args[1] + " is valid!")
        } else {
            fmt.Println(os.Args[1] + " is invalid!")
        }
    }

    func validateEmail(email string) bool {
        resp, err := http.PostForm("http://localhost:8081/", url.Values{"email": {email}})

        if err != nil {
            fmt.Printf("Sorry there was an error: %s\n", err)
            return false
        }

        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)

        response := response{}
        err = json.Unmarshal(body, &response)

        if err != nil {
            fmt.Printf("Sorry could not unmarshal json response: %s\n", err)
            return false
        }

        return response.Valid
    }

## Service JSON Output

Example valid response

    {"status":200,"message":"OK","email":"a@a.com","valid":true,"host":"a.com","user":"a"}

Example invalid response

    {"status":200,"message":"OK","email":"blargh@host(*","valid":false,"host":"host(*","user":"blargh"}