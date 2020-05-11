package checkGuess

import (
	"math"
	"numType"
)

func CheckGuess(numToBeGuessed, curGuess numType.Num) (string, bool) {

	hint := ""
	if numToBeGuessed == curGuess {
		return "PERFECT GUESS! CONGRATULATIONS!", true
	} else if numToBeGuessed > curGuess {
		hint += "TRY HIGHER"
	} else if numToBeGuessed < curGuess {
		hint += "TRY LOWER"
	}
	if math.Abs(float64(numToBeGuessed-curGuess)) <= 5 {
		hint += " ,VERY CLOSE"
	}
	return hint, false
}
