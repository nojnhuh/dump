package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func getRequest() {
	resp, err := http.Get("http://localhost:8000/")
	if err != nil {
		fmt.Println("Oh no")
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Read body: %v", err)
		return
	}
	fmt.Println(string(data))
}

func postRequest(toRead string) {
	myReader := strings.NewReader(toRead)
	resp, err := http.Post("http://localhost:8000/", "Post Request Header", myReader)
	if err != nil {
		fmt.Println("Oh no")
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Read body: %v", err)
		return
	}
	fmt.Println(string(data))
}

func main() {

	requestPtr := flag.String("requestName", "get", "a request type")

	bodyPtr := flag.String("body", "Rock Soup", "the post request body") // as a note, you may need to add quotes around the text you pass in here

	flag.Parse()

	if *requestPtr == "get" {
		getRequest()
	}
	if *requestPtr == "post" {
		postRequest(*bodyPtr)
	}

	//resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
}
