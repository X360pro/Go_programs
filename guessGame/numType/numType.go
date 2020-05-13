package numType

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Num int

func (n *Num) ReadInput() error {
	reader := bufio.NewReader(os.Stdin)
	guess, _ := reader.ReadString('\n')
	guess = strings.TrimSpace(guess)
	guessInt, err := strconv.Atoi(guess)
	if err != nil {
		return errors.New("Enter integer and not text")
	}
	*n = Num(guessInt)

	err = n.Checkinglimit()
	if err != nil {
		return err
	}
	return nil
}

func (n *Num) Checkinglimit() error {
	if *n > 50 || *n < 0 {
		return errors.New("NO. should lie between 1 and 50")
	}
	return nil
}
