package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	api := http.Client{}

	request, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Accept", "application/json")

	response, err := api.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))
}
