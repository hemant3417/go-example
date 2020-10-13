package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	name := "Hemanth Varma"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`

	nf, err := os.Create("index.html")

	defer nf.Close()

	if err != nil {
		log.Println("Error Occured creating index file ")
		return
	}

	io.Copy(nf, strings.NewReader(tpl))

}
