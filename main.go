package main

import (
	"encoding/json"
	"log"
	"myapp/rps"
	"net/http"
	"strconv"
	"text/template"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// html := `<strong>Hello world</strong>`
	// w.Header().Set("Content-Type", "text/html")
	//fmt.Fprintf(w, html)
	renderPage(w, "index.html")

}

func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice , _:= strconv.Atoi(r.URL.Query().Get("c"))
	round := rps.PlayRound(playerChoice)
	out, err := json.MarshalIndent(round, "", "    ")
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/play", playRound)
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
