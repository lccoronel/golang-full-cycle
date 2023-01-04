package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	api := http.Client{}

	jsonVar := bytes.NewBuffer([]byte(`{"name": "lucas"}`))

	response, err := api.Post("http://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	io.CopyBuffer(os.Stdout, response.Body, nil)
}
