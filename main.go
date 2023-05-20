package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	
	"github.com/julienschmidt/httprouter"
)

const maxUploadSize = 10 * 1024 * 1024 // 10 mb
var uploadPath = "./uploads"
var posterpath = "./poster"
var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/vid", Vid)
	router.GET("/hello/:name", Hello)

	router.ServeFiles("/v/*filepath", http.Dir("./uploads"))
	log.Fatal(http.ListenAndServe(":4000", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Vid(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl.ExecuteTemplate(w, "head.html", nil)
	tmpl.ExecuteTemplate(w, "viewpost.html", nil)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
