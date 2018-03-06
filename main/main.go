package main

import (
	"fmt"
	"../model"
	"../data"
)

func main() {
	model.Authors = data.ImportAuthors("/home/bogdanboboc97/Projects/goworkshop2/data/authors.json")
	model.Books = data.ImportBooks("/home/bogdanboboc97/Projects/goworkshop2/data/books.json")

	fmt.Println(model.Books)
	fmt.Println(model.Authors)
}


