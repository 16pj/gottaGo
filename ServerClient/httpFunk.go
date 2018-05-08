package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func getBodyCore(stringChan chan string, fullUrl string) {

	resp, err := http.Get(fullUrl)

	if err != nil {
		log.Fatal("Unable to access page!")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)

		if err2 != nil {
			log.Fatal("Unable to extract the body from site!")
		}
		bodyString := string(bodyBytes)

		stringChan <- bodyString
	}

}

func getBody(addr string, other_params ...string) string {

	fullUrl := ""

	if len(other_params) > 2 {
		log.Fatal("Too many arguments. Please provide maximum of 3.")
	} else if len(other_params) > 0 {
		if len(other_params) > 1 {
			fullUrl = "http://" + addr + ":" + other_params[0] + "/" + other_params[1]
		} else {
			fullUrl = "http://" + addr + ":" + other_params[0]
		}
	} else {
		fullUrl = "http://" + addr
	}

	stringChan := make(chan string)

	go getBodyCore(stringChan, fullUrl)
	time.Sleep(time.Millisecond * 10)
	body := <-stringChan
	return body
}

func main() {
	body := getBody("localhost", "8080")
	//body := getBody("google.com")

	fmt.Println(body)
}
