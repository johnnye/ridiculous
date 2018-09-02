package main
import (
	"html/template"
	"io/ioutil"
	"gopkg.in/russross/blackfriday.v2"
	"github.com/microcosm-cc/bluemonday"
	"fmt"
	"path/filepath"
)


type PageData struct {
	PageTitle string
	Content   template.HTML
}

func allPages() {

}

func page(name string) PageData {
	fmt.Println(len(name))

	if (len(name) <= 1) {
		fmt.Println("No Path show the index")
		name = "index"
	}

	filetoread := filepath.Join("content", name)
	filetoread = filetoread + ".md"
	fmt.Println(filetoread)

	input, err := ioutil.ReadFile(filetoread)
	check(err)
	unsafe := blackfriday.Run(input)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	return PageData{PageTitle:"Nice", Content:template.HTML(html)}
}