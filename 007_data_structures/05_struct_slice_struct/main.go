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

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	si := []sage{{
		"Gandhi",
		"Peace",
	},
		{
			"Rambo",
			"Warrior",
		},
		{
			"James",
			"Spy",
		},
	}

	ci := []car{
		{
			"vw",
			"polo",
			5,
		},
		{
			"tata",
			"tiago",
			5,
		},
	}

	di := items{
		Wisdom:    si,
		Transport: ci,
	}
	err := tpl.Execute(os.Stdout, di)
	if err != nil {
		log.Fatalln(err)
	}

}
