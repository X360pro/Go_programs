package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	content, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("content of file is : ", string(content))

	f, er := os.Open("data.txt")
	defer f.Close()
	checkError(er)
	bucket := make([]byte, 100)
	bytesRead, e := f.Read(bucket)
	checkError(e)
	fmt.Println("content of file is (limited): ", string(bucket[:bytesRead]))
	fmt.Println(bytesRead)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error occurred : ", err)
	}
}
