package main

import (
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
	postRequest("Rock Attack")
	getRequest()

	//resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
}
