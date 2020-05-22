package board

import "fmt"

type Board struct {
	CurBoard     [][]string
	CurBoardSize int
}

func (b *Board) CleanBoard(size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			b.CurBoard[i][j] = ""
		}
	}
	b.CurBoardSize = size
}

func NewBoard(curBoardSize int) Board {
	new := Board{}
	new.CurBoard = make([][]string, curBoardSize)
	for i := range new.CurBoard {
		new.CurBoard[i] = make([]string, curBoardSize)
	}
	new.CurBoardSize = curBoardSize
	return new
}

func (b *Board) Display() {
	count := 0
	for i := 0; i < b.CurBoardSize; i++ {
		for j := 0; j < b.CurBoardSize; j++ {
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
	for i := 0; i < b.CurBoardSize; i++ {
		if b.CurBoard[i][i] != curSym {
			win = false
		}
	}
	if win == true {
		return true
	}

	//for second diagonal
	win = true
	for i := 0; i < b.CurBoardSize; i++ {
		j = b.CurBoardSize - 1 - i
		if b.CurBoard[i][j] != curSym {
			win = false
		}
	}
	if win == true {
		return true
	}

	//for row check
	for i := 0; i < b.CurBoardSize; i++ {
		win := true
		for j := 0; j < b.CurBoardSize; j++ {
			if b.CurBoard[i][j] != curSym {
				win = false
			}
		}
		if win == true {
			return true
		}
	}

	//for column check
	for i := 0; i < b.CurBoardSize; i++ {
		win := true
		for j := 0; j < b.CurBoardSize; j++ {
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
