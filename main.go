package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"path/filepath"
	"html/template"
	"fmt"
)


func homeHandler(rw http.ResponseWriter, r *http.Request) {

	page, pageerror := page(filepath.Clean(r.URL.Path))

	check(pageerror) //TODO: 404 and 500 here rather then panic

	layout := filepath.Join("templates", "index.html")

	tmpl := template.New("index")
	tmpl, _ = tmpl.ParseFiles(layout)

	err := tmpl.ExecuteTemplate(rw, "layout", page)
	check(err)

}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	
	http.HandleFunc("/", homeHandler)
	fmt.Print(os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
