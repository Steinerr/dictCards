package main

import (
	"io/ioutil"
	"net/http"
	"html/template"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	return ioutil.WriteFile(getWikiPageFullPath(p.Title), p.Body, 0600)
}

func getWikiPageFullPath(pageTitle string) string {
	return "pages/" + pageTitle + ".txt"
}

func loadPage(title string) (*Page, error) {
	body, err := ioutil.ReadFile(getWikiPageFullPath(title))
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func wikiPageView(w http.ResponseWriter, r *http.Request) {
	pageTitle := r.URL.Path[1:]
	body, err := ioutil.ReadFile(getWikiPageFullPath(pageTitle))
	p := &Page{Title: pageTitle}
	if err == nil {
		p.Body = body
	}

	t, _ := template.ParseFiles("templates/page.html")
	t.Execute(w, p)
}

func main(){
	http.HandleFunc("/", wikiPageView)
	http.ListenAndServe(":8080", nil)
}