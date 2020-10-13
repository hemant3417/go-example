package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	s1 := sage{
		"Gandhi",
		"Peace",
	}

	err := tpl.Execute(os.Stdout, s1)
	if err != nil {
		log.Fatalln(err)
	}

}
