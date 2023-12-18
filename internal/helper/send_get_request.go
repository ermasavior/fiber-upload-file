package helper

import (
	"fmt"
	"net/http"
)

func SendGetRequest(requestURL string) bool {
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return false
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	return true
}
