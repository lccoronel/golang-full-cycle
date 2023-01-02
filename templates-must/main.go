package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Name      string
	ClassTime int
}

func main() {
	curso := Curso{Name: "Go", ClassTime: 40}
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso: {{.Name}}, Carge Horaria: {{.ClassTime}}")

	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
