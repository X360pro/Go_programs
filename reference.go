package main

import "fmt"

func modifyValueInt(num int) {
	num = 10
	fmt.Println("value inside for passing value ", num)
}

func modifyPointerInt(num *int) {
	*num = 10
	fmt.Println("value inside for passing pointer", *num)
}
func modifyValueFloat(num float64) {
	num = 10
	fmt.Println("value inside for passing value ", num)
}

func modifyPointerFloat(num *float64) {
	*num = 10
	fmt.Println("value inside for passing pointer", *num)
}

func modifyValueString(text string) {
	text = "newtext"
	fmt.Println("value inside for passing value ", text)
}

func modifyPointerString(text *string) {
	*text = "newtext"
	fmt.Println("value inside for passing pointer", *text)
}
func modifyValueBool(flag bool) {
	flag = true
	fmt.Println("value inside for passing value ", flag)
}

func modifyPointerBool(flag *bool) {
	*flag = true
	fmt.Println("value inside for passing pointer", *flag)
}

func main() {
	num := 5
	modifyValueInt(num)
	fmt.Println("value in main for passing vlaue", num)
	modifyPointerInt(&num)
	fmt.Println("value in main for passing pointer", num)
	text := "old"
	modifyValueString(text)
	fmt.Println("value in main for passing vlaue", text)
	modifyPointerString(&text)
	fmt.Println("value in main for passing pointer", text)
	floatnum := 5.0
	modifyValueFloat(floatnum)
	fmt.Println("value in main for passing vlaue", floatnum)
	modifyPointerFloat(&floatnum)
	fmt.Println("value in main for passing pointer", floatnum)
	flag := false
	modifyValueBool(flag)
	fmt.Println("value in main for passing vlaue", flag)
	modifyPointerBool(&flag)
	fmt.Println("value in main for passing pointer", flag)
}
