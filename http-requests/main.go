package main

import (
	"io"
	"net/http"
)

func main() {
	request, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	defer request.Body.Close()

	response, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	println(string(response))
}

// defer faz com que a linha seja executada por ultimo
