package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", handleHello)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("connect success")
	fmt.Println("remoteAddr: ", r.RemoteAddr)
	fmt.Println("url: ", r.URL)
	fmt.Println("header: ", r.Header)
	fmt.Println("body: ", r.Body)

	w.Write([]byte("hello http"))

}
