package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

	name := os.Args[1]

	tpl := fmt.Sprint(`
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
	`)

	nf, err := os.Create("index.html")

	defer nf.Close()

	if err != nil {
		log.Fatal("Error creating file ", err)
	}

	io.Copy(nf, strings.NewReader(tpl))

}
