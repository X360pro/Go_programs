package main

import "fmt"
import "github.com/X360pro/Go/tree/master/greeting"

func negate(myBoolean * bool){
	*myBoolean = !(*myBoolean)
}
func main(){
	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)
	greeting.Hello()
}