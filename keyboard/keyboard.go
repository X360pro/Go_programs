package keyboard

import ("bufio"
		"os"
		"strconv"
		"strings")

func GetFloat()(float64,error){
	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString(\n) 
	input = strings.TrimSpace(input)
	number,_ := strconv.ParseFloat(input,64)
	return  number,nil
}