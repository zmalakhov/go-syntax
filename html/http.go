package main

/**
* Syntax Go. Homework 6
* Zaur Malakhov, dated May 15, 2019
 */

import (
	"fmt"
	"html/template"
	"net/http"
)

type ViewData struct {
	Name string
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		name := r.URL.Query().Get("name")
		data := ViewData{
			Name: name,
		}
		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, data)
	})

	fmt.Println("server is running...")
	http.ListenAndServe(":8080", nil)
}
