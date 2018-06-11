package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gtongy/othello_ai/board"
)

func main() {
	var b board.Board
	b.Initial()
	for !b.EndGame() {
		b.Print()
		if !b.HasToPut(b.Turn) {
			fmt.Println("There is no place to put.")
			b.Turn = !b.Turn
			continue
		}
		if b.Turn == true {
			fmt.Println("my turn")
			if !move(b, board.WHITE_VAL, board.BLACK_VAL) {
				continue
			}
		}
		if b.Turn == false {
			fmt.Println("your turn")
			if !move(b, board.BLACK_VAL, board.WHITE_VAL) {
				continue
			}
		}
		b.Turn = !b.Turn
	}
	fmt.Println("End Game!!!")
	b.Print()
}

func move(b board.Board, my_val, your_val int) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("please input x,y : ex) 1,3 |>")
	text, _ := reader.ReadString('\n')
	splited := strings.Split(strings.TrimRight(text, "\n"), ",")
	x, _ := strconv.Atoi(splited[0])
	y, _ := strconv.Atoi(splited[1])
	if !b.Reverse(x, y, my_val, your_val) {
		fmt.Println("invalid refers to the board")
		return false
	}
	return true
}
