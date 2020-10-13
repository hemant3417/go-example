package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	sa := []string{"Jamie", "ram", "abdul"}

	err := tpl.Execute(os.Stdout, sa)
	if err != nil {
		log.Fatalln(err)
	}
}
