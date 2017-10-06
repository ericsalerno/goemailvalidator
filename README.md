# goemailvalidator

Email validation service written in golang.

This was built mostly for my own learning process with golang.

1. Run the service
2. HTTP POST an "email" variable to /validate
3. read the json it returns

## Example Client Code

    package main

    import (
        "fmt"
        "io/ioutil"
        "net/http"
        "net/url"
    )

    func main() {
        testEmailValidate("test@test.com")
        //Test Output (Code 200): {"status":200,"message":"OK","email":"test@test.com","valid":true}

        testEmailValidate("")
        //Test Output (Code 500): {"status":500,"message":"You must post an email address with the variable name 'email'.","email":"test@test.com","valid":false}
    }

    func testEmailValidate(email string) {
        resp, err := http.PostForm("http://localhost:8081/validate", url.Values{"email": {email}})

        if err != nil {
            fmt.Printf("Sorry there was an error: %s\n", err)
            return
        }

        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)

        fmt.Printf("Test Output (Code %d): %s\n", resp.StatusCode, body)
    }




