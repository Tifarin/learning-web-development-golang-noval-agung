package main

import (
	"fmt"
	"net/http"
)

func main() {
	//membuat handler (handler itu seperti controller)
	handlerIndex := func (w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("hello latifah"))
	}
	handlerData := func (w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("hello again"))
	}
	//route atau pengalamatan
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/data", handlerData)
	//menambah server
	
	fmt.Println("Server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}