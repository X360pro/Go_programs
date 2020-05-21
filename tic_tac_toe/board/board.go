package board

import "fmt"

type Board struct {
	CurBoard [][]string
}

func (b *Board) CleanBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b.CurBoard[i][j] = ""
		}
	}
}

func NewBoard() Board {
	new := Board{}
	new.CurBoard = make([][]string, 3)
	for i := range new.CurBoard {
		new.CurBoard[i] = make([]string, 3)
	}
	return new
}

func (b *Board) Display() {
	count := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			count++
			if b.CurBoard[i][j] == "" {
				fmt.Print(" -- (", count-1, ") ")
				continue
			}
			fmt.Print("  ", b.CurBoard[i][j], " (", count-1, ") ")
		}
		fmt.Printf("\n")
	}
}

//checks if player with current symbol has won
func (b *Board) CheckStatOfBoard(curSym string) bool {
	//for first diagonal
	win := true
	j := 0
	for i := 0; i < 3; i++ {
		if b.CurBoard[i][i] != curSym {
			win = false
		}
	}
	if win == true {
		return true
	}

	//for second diagonal
	win = true
	for i := 0; i < 3; i++ {
		j = 2 - i
		if b.CurBoard[i][j] != curSym {
			win = false
		}
	}
	if win == true {
		return true
	}

	//for row check
	for i := 0; i < 3; i++ {
		win := true
		for j := 0; j < 3; j++ {
			if b.CurBoard[i][j] != curSym {
				win = false
			}
		}
		if win == true {
			return true
		}
	}

	//for column check
	for i := 0; i < 3; i++ {
		win := true
		for j := 0; j < 3; j++ {
			if b.CurBoard[j][i] != curSym {
				win = false
			}
		}
		if win == true {
			return true
		}
	}
	return false
}
