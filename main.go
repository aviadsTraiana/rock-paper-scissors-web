package main

import (
	"log"
	"net/http"
	"text/template"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// html := `<strong>Hello world</strong>`
	// w.Header().Set("Content-Type", "text/html")
	//fmt.Fprintf(w, html)
	renderPage(w, "index.html")

}

func main() {
	http.HandleFunc("/", homePage)
	log.Println("Starting web server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func renderPage(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
