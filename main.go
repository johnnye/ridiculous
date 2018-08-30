package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
)

func homeHandler(rw http.ResponseWriter, r *http.Request) {
	input, err := ioutil.ReadFile("test.md")
	check(err)
	fmt.Print(string(input))
	unsafe := blackfriday.Run(input)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	fmt.Fprintln(rw, string(html))
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


	http.HandleFunc("/", homeHandler)
	log.Print(os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
