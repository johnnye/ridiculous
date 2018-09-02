package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
	"path/filepath"
	"html/template"
	"fmt"
)


type PageData struct {
	PageTitle string
	Content   template.HTML
}

func homeHandler(rw http.ResponseWriter, r *http.Request) {
	input, err := ioutil.ReadFile("test.md")
	check(err)
	unsafe := blackfriday.Run(input)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	page := PageData{PageTitle:"Nice", Content:template.HTML(html)}

	layout := filepath.Join("templates", "index.html")

	tmpl := template.New("index")
	tmpl, _ = tmpl.ParseFiles(layout)

	err = tmpl.ExecuteTemplate(rw, "layout", page)
	check(err)

}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


	http.HandleFunc("/", homeHandler)
	fmt.Print(os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
