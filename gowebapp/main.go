package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080",nil)
}

func handler(writer http.ResponseWriter,request *http.Request){
	_, _ = fmt.Fprintf(writer, "hello world, %s!", request.URL.Path[1:])
}