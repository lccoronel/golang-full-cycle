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
	curso := Curso{"Go", 40}

	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Name}}, Carge Horaria: {{.ClassTime}}"))

	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
