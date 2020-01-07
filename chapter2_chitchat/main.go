package main

import (
	"html/template"
	"net/http"
)

func main() {
	mux:=http.NewServeMux()
	files:=http.FileServer(http.Dir("/public"))
	mux.HandleFunc("/",index)
	mux.Handle("/static/",http.StripPrefix("/static/",files))
}

func index(w http.ResponseWriter,r *http.Request){
	files:=[]string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",

	}
	templates:=template.Must(template.ParseFiles(files...))
	threads,err:=data.Threads()
	if err==nil{
		templates.ExecuteTemplate(w,"layout",threads)
	}
}