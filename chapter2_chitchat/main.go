package main

import (
	"html/template"
	"net/http"
	"chapter2_chitchat/data"
)

func main() {
	mux:=http.NewServeMux()
	files:=http.FileServer(http.Dir("/public"))
	mux.Handle("/static/",http.StripPrefix("/static/",files))
	mux.HandleFunc("/",index)
	server:=&http.Server{
		Addr:              "0.0.0.0:8080",
		Handler:           mux,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	server.ListenAndServe()

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