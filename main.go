package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/", mainHandle)
	http.HandleFunc("/click", buttonClickedHandler)

	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":3000", nil)
}

func mainHandle(w http.ResponseWriter, r *http.Request) {
	users := map[string][]User{
		"Users": {
			{ID: "AA715721", Name: "Kenneth Sheldon"},
			{ID: "AA245917", Name: "William Henrion"},
			{ID: "AA739440", Name: "Kevin Wong"},
		},
	}
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	tmpl.Execute(w, users)
}

func buttonClickedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Button clicked")
	data := User{ID: "AA999999", Name: "New User Name"}
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	tmpl.ExecuteTemplate(w, "user-element", data)
}
