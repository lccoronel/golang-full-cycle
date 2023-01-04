package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	api := http.Client{Timeout: time.Second}

	response, err := api.Get("http://google.com")
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
