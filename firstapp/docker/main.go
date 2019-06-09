package main

import (
	"net/http"
	"net"
	"log"
	"encoding/json"
	"fmt"
)

type ResponseStruct struct{
	Output string
	Status string
}

func printHello(w http.ResponseWriter, r *http.Request){

	msg := ResponseStruct{
		Output : "Hello World",
		Status: "success",
	}

	jsonData, err := json.Marshal(msg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
	w.Write(jsonData)
}

func getIP(w http.ResponseWriter, r *http.Request){

	conn, err := net.Dial("udp", "8.8.8.8:80")

	ip_response := ResponseStruct{
		Status: "",
		Output:"",
	}

	if err != nil{
		log.Fatal(err)
	}

	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	ipAddr, _, err := net.SplitHostPort(localAddr.String())

	if err != nil{
		ip_response.Status = "failure"
		log.Fatal(err)
	}else{

	ip_response.Output = ipAddr
	ip_response.Status = "success"

	jsonData, err := json.Marshal(ip_response)

	if err != nil {
		ip_response.Status = "failure"
		log.Fatal(err)
	}
		w.Write(jsonData)

	}


}

func main(){

	http.HandleFunc("/", printHello)
	http.HandleFunc("/ip", getIP)

	http.ListenAndServe(":8089", nil)
}
