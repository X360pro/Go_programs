package main

import ("fmt")

func add(num1,num2 float64)float64{
	result := num1 + num2
	return result
}
func sub(num1,num2 float64)float64{
	result := num1 - num2 
	return result
}
func mul(num1,num2 float64)float64{
	result := num1 * num2 
	return result
}
func div(num1,num2 float64)float64{
	result := num1 / num2 
	return result
}
func mathOperations(result float64){
	fmt.Println(result)
}
func main(){
	mathOperations(add(10,5))
	mathOperations(sub(10,5))
	mathOperations(mul(10,5))
	mathOperations(div(10,5))
}