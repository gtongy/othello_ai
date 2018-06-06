package main

import (
	"github.com/gtongy/othello_ai/board"
)

func main() {
	var board board.Board
	board.Initial()
	board.Print()
}
