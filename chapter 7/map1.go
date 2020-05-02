package main

import "fmt"

func main(){
	jewellery := make(map[string]float64)
	jewellery["necklace"] = 89.99
	jewellery["earrings"] = 79.99
	my,ok := jewellery["earrings"]
	fmt.Println(my,ok)
	clothing := map[string]float64{"pants":59.99,"shirt":39.99}
	fmt.Println(jewellery,clothing)
}