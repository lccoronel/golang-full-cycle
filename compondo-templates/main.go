package main

import (
	"html/template"
	"os"
	"strings"
)

type Curso struct {
	Name      string
	ClassTime int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Jave", 90},
		{"JS", 30},
	})
	if err != nil {
		panic(err)
	}
}
