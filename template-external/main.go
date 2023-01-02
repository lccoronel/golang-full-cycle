package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Name      string
	ClassTime int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Jave", 90},
		{"JS", 30},
	})
	if err != nil {
		panic(err)
	}
}
