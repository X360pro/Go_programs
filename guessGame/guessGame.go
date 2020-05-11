package main

import (
	"checkGuess"
	"fmt"
	genRandom "genRandom"
	"numType"
)

func rules() {
	fmt.Println("      Rules :")
	fmt.Println("      		1)The no. generated will be between 1-50")
	fmt.Println("     		2)Challenge is to guess the number")
	fmt.Println("      		3)You have 5 chances to Guess the no.")
	fmt.Println("      		4)if guessed no. is smaller than original no. then hint will be HIGHER!")
	fmt.Println("      		5)if guessed no. is greater than original no. then hint will be LOWER!")
	fmt.Println("      		6)if the difference between the guessed number and original no. is less than 5")
	fmt.Println("      			then hint given will be VERY CLOSE !")
}

func main() {
	numToBeGuessed := numType.Num(genRandom.RandomNumGenerator(1, 50))
	err := numToBeGuessed.CheckingType()

	if err != nil {
		fmt.Println(err)
	}

	rules()
	fmt.Println("Enter your guess")

	var guess numType.Num
	var hint string
	var flag bool
	for i := 1; i <= 5; i++ {
		fmt.Println(" ")
		fmt.Println("Attempt no. ", i)

		for {
			err := numType.Scan(&guess)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Try Again with integer input")
				continue
			}
			newerr := guess.CheckingType()
			if newerr != nil {
				fmt.Println(newerr)
				continue
			}
			break
		}
		if hint, flag = checkGuess.CheckGuess(numToBeGuessed, guess); flag == true {
			fmt.Println(hint)
			break
		} else {
			fmt.Println(hint)
		}
	}

	fmt.Println("                  Original no. was", numToBeGuessed)
}
