package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", `Never Back Down : else Go to Hell`)
	if err != nil {
		log.Fatalln(err)
	}

}
