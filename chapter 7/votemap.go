package main

import ("fmt"
		"os"
		"bufio")

func main(){
	file,_ := os.Open("data_ch_7.txt")
	scanner := bufio.NewScanner(file)
	names := make(map[string]int)
	for scanner.Scan(){
		if _,ok := names[scanner.Text()] ; ok {
			names[scanner.Text()]++
		} else {
			names[scanner.Text()] = 0
		}
	}
	fmt.Println(names)
}