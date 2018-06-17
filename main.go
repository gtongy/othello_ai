package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gtongy/othello_ai/board"
	"github.com/gtongy/othello_ai/evaluate"
)

func main() {
	var b board.Board
	var evaluation evaluate.Evaluation
	evaluation.Set()
	b.Initial(evaluation.Rows)
	for !b.EndGame() {
		b.Print()
		if !b.HasToPut(b.Turn) {
			fmt.Println("There is no place to put.")
			b.Turn = !b.Turn
			continue
		}
		if b.Turn == true {
			fmt.Println("my turn, can put cells are")
			fmt.Println(b.GetToPutNextCells(b.Turn))
			if !move(b, board.WhiteVal, board.BlackVal) {
				continue
			}
		}
		if b.Turn == false {
			fmt.Println("your turn, can put cells are")
			fmt.Println(b.GetToPutNextCells(b.Turn))
			if !move(b, board.BlackVal, board.WhiteVal) {
				continue
			}
		}
		b.Turn = !b.Turn
	}
	fmt.Println("End Game!!!")
	b.Print()
}

func move(b board.Board, myVal, yourVal int) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("please input x,y : ex) 1,3 |>")
	text, _ := reader.ReadString('\n')
	splited := strings.Split(strings.TrimRight(text, "\n"), ",")
	x, _ := strconv.Atoi(splited[0])
	y, _ := strconv.Atoi(splited[1])
	if !b.Reverse(x, y, myVal, yourVal) {
		fmt.Println("invalid refers to the board")
		return false
	}
	return true
}
