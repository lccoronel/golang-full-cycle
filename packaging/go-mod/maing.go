package main

import "github.com/google/uuid"

func main() {
	var id = uuid.New().String()
	print(id)
}
