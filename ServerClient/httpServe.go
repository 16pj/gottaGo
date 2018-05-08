package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func main() {

	var serverAddr = "localhost"
	var serverPort = "8080"

	http.HandleFunc("/", handler)
	fmt.Println("Starting server at addr:", serverAddr, "and port:", serverPort)
	http.ListenAndServe(serverAddr+":"+serverPort, nil)

}
