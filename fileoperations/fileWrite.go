package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Hello from go")
	err := ioutil.WriteFile("data1.txt", data, 0644)
	checkError(err)

	f, er := os.Create("myFile.txt")
	defer f.Close()
	checkError(er)
	bytesWritten, e := f.WriteString("Hello")
	checkError(e)
	fmt.Println("Bytes written is ", bytesWritten)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error occurred : ", err)
	}
}
