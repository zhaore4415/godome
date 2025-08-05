package main

import (
	"html/template"
	"os"
)

type PageVariables struct {
	Title   string
	Message string
	Items   []string
}

func main() {
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		panic(err)
	}

	data := PageVariables{
		Title:   "My Website",
		Message: "Welcome to my website!",
		Items:   []string{"Item 1", "Item 2", "Item 3"},
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
