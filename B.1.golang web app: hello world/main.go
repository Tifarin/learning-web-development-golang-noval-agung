package main

import (
	"fmt"
	"net/http"
)
//http.responseWrite dan http.Request di perlukan oleh http.HandleFunc untuk penanganan request ke route yang  ditentukan, w untuk write, r untuk request
func handlerIndex(w http.ResponseWriter, r *http.Request) {
  var message = "Welcome"
  //Fungsi http.ResponseWriter.Write() pada dasarnya menerima parameter berupa slice of byte ([]byte). Oleh karena itu, sebelum mengirimkan pesan ke client, kita harus mengubah pesan yang ingin dikirim ke dalam bentuk []byte.
  w.Write([]byte(message))
}
func handlerHello(w http.ResponseWriter, r *http.Request) {
  var message = "Hello World!"
  w.Write([]byte(message))
}
func main() {
  //route
  //Fungsi http.HandleFunc() digunakan untuk routing. Parameter pertama adalah rute dan parameter ke-2 adalah handler-nya.
  http.HandleFunc("/", handlerIndex)
  http.HandleFunc("/index", handlerIndex)
  http.HandleFunc("/hello", handlerHello)
  //start server
  var address = "localhost:9000"
  fmt.Printf("server started at %s \n", address)
  //Fungsi http.ListenAndServe() digunakan membuat sekaligus start server baru, dengan parameter pertama adalah alamat web server yang diiginkan (bisa diisi host, host & port, atau port saja). Parameter kedua merupakan object mux atau multiplexer.
  err := http.ListenAndServe(address, nil)
  if err != nil {
    fmt.Println(err.Error())
  }
}
