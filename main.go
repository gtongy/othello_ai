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
	var board board.Board
	board.Initial()
	for {
		board.Print()
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("please input x,y : ex) 1,3 |>")
		text, _ := reader.ReadString('\n')
		splited := strings.Split(strings.TrimRight(text, "\n"), ",")
		x, _ := strconv.Atoi(splited[0])
		y, _ := strconv.Atoi(splited[1])
		if !board.Reverse(x, y) {
			fmt.Println("invalid refers to the board")
			continue
		}
		board.Update(x, y)
	}
}
