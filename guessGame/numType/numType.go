package numType

import (
	"errors"
	"fmt"
)

type Num int

type CheckType interface {
	CheckingType() error
}

func Scan(a *Num) error {
	_, err := fmt.Scan(a)
	return err
}

func (n *Num) CheckingType() error {
	if *n > 50 || *n < 0 {
		return errors.New("NO. should lie between 1 and 50")
	}
	return nil
}
