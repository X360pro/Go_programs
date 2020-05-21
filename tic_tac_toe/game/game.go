package game

import (
	"board"
	"bufio"
	"fmt"
	"os"
	"player"
	"strconv"
	"strings"
)

type Game struct {
	board.Board
	player1 player.Player
	player2 player.Player
}

func NewGame() Game {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name of 1st player : ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	player1 := player.Player{input, "X", false, 0}
	fmt.Print("Enter name of 2nd player : ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	player2 := player.Player{input, "O", false, 0}
	new := Game{board.NewBoard(), player1, player2}
	return new
}

func (g *Game) Play() {
	count := 1
	win := false
	reader := bufio.NewReader(os.Stdin)
PlayAgain:
	fmt.Printf("%v  is X \n", g.player1.Name)
	fmt.Printf("%v  is O \n", g.player2.Name)
	g.player1.Turn = true
	g.Board.Display()
	for chances := 0; chances < 9; chances++ {
		if g.player1.Turn {
			fmt.Printf("%v enter your position : ", g.player1.Name)
			g.move(g.player1.Symbol)
			win = g.CheckStatOfBoard(g.player1.Symbol)
			if win {
				g.player1.Score++
				g.Display()
				fmt.Printf("Congratulations %v has won this round\n", g.player1.Name)
				break
			}
			g.Display()
			g.player1.Turn = false
			g.player2.Turn = true
		} else {
			fmt.Printf("%s enter your position : ", g.player2.Name)
			g.move(g.player2.Symbol)
			win = g.CheckStatOfBoard(g.player2.Symbol)
			if win {
				g.player2.Score++
				g.Display()
				fmt.Printf("Congratulations %v has won this round\n", g.player2.Name)
				break
			}
			g.Display()
			g.player2.Turn = false
			g.player1.Turn = true
		}
	}
	if !win {
		fmt.Println("Match is a Draw")
	}
	fmt.Println("ENTER 1  PLAY AGAIN AND 0 TO EXIT")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "1" {
		temp := g.player1
		g.player1 = g.player2
		g.player2 = temp
		g.player1.Symbol = "X"
		g.player2.Symbol = "O"
		count++
		g.CleanBoard()
		goto PlayAgain
	}
	fmt.Printf("Final Score after %d rounds\n", count)
	fmt.Printf(" %v : %d \n", g.player1.Name, g.player1.Score)
	fmt.Printf(" %v : %d \n", g.player2.Name, g.player2.Score)
	if g.player1.Score > g.player2.Score {
		fmt.Printf("%v has won !!!", g.player1.Name)
	} else if g.player1.Score < g.player2.Score {
		fmt.Printf("%v has won !!!", g.player2.Name)
	} else {
		fmt.Println("Both have scored equal points")
	}
	fmt.Println("Thankyou for playing")
}

func (g *Game) move(curSym string) {
	reader := bufio.NewReader(os.Stdin)
Repeat:
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	move, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err, " Try Again")
		goto Repeat
	}
	if move < 0 || move > 9 || g.CurBoard[move/3][move%3] != "" {
		fmt.Println("Positon is either occupied or out of bounds ")
		fmt.Println(" Try Again")
		goto Repeat
	}
	g.CurBoard[move/3][move%3] = curSym
}
