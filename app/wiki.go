package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"net/http"
	"html/template"
)

type Page struct {
	Title string
	Body []byte
}

//func (p *Page) save() error {
//	return ioutil.WriteFile(getWikiPageFullPath(p.Title), p.Body, 0600)
//}

func wikiPageView(w http.ResponseWriter, r *http.Request) {
	// todo: remove hardcode in the future
	// todo: investigate ssl mode
	// todo: how can i print a detailed error?
	// todo: WTF??!! WHY????!
	db, err := sql.Open("postgres", "postgres://dictcard:dictcard@db/dictcard?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	var body string
	pageTitle := r.URL.Path[1:]
	p := &Page{Title: pageTitle}

	result := db.QueryRow("SELECT p.body FROM page as p WHERE p.title = $1", pageTitle).Scan(&body)

	if result == sql.ErrNoRows {
		p.Body = []byte("404: Page not found ¯\\_(ツ)_/¯")
	} else if result != nil {
		p.Body = []byte("500: Internal server error :(")
		//log.Fatal(err)
		log.Println(result)
	} else {
		p.Body = []byte(body)
	}

	t, _ := template.ParseFiles("templates/page.html")
	t.Execute(w, p)
}

func main(){
	http.HandleFunc("/", wikiPageView)
	http.ListenAndServe(":8080", nil)
}