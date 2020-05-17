package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func responSize(url string, myChannel chan int) {

	fmt.Println("Getting : ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	myChannel <- len(body)
}

func main() {

	myChannel := make(chan int)

	go responSize("http://example.com", myChannel)
	go responSize("http://golang.org/", myChannel)
	go responSize("http://golang.org/doc", myChannel)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
