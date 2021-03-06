package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gtongy/othello_ai/ai"
	"github.com/gtongy/othello_ai/board"
	"github.com/gtongy/othello_ai/evaluate"
)

func main() {
	var b board.Board
	var evaluation evaluate.Evaluation
	evaluation.Set()
	b.Initial(evaluation.Rows)
	for !b.EndGame() {
		if !b.HasToPut(b.Turn) {
			fmt.Println("There is no place to put.")
			b.Turn = !b.Turn
			continue
		}
		if b.Turn == true {
			b.Print()
			fmt.Println(b.GetToPutNextCells(b.Turn))
			x, y := choiceMyTurnCell()
			if !move(b, x, y) {
				continue
			}
		}
		if b.Turn == false {
			cells := b.GetToPutNextCells(b.Turn)
			// random choice cell
			choiceCell := ai.RandomChoiceCellIndex(cells)
			x, y := board.GetPosition(choiceCell)
			if !move(b, x, y) {
				continue
			}
		}
		b.Turn = !b.Turn
	}
	fmt.Println("End Game!!!")
	myCount, yourCount := b.Calculate()
	fmt.Printf("white %v, black %v\n", myCount, yourCount)
	b.Print()
}

func move(b board.Board, x, y int) bool {
	myVal, yourVal := board.GetColor(b.Turn)
	if !b.Reverse(x, y, myVal, yourVal) {
		fmt.Println("invalid refers to the board")
		return false
	}
	return true
}

func choiceMyTurnCell() (int, int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("please input x,y : ex) 1,3 |>")
	text, _ := reader.ReadString('\n')
	splited := strings.Split(strings.TrimRight(text, "\n"), ",")
	x, _ := strconv.Atoi(splited[0])
	y, _ := strconv.Atoi(splited[1])
	return x, y
}
