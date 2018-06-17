package ai

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gtongy/othello_ai/board"
)

const CellMinSize = 0

func RandomChoiceCellIndex(cells []board.Cell) board.Cell {
	if len(cells) == 1 {
		return cells[0]
	}
	return cells[random(CellMinSize, len(cells)-1)]
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	fmt.Println(min)
	fmt.Println(max)
	return rand.Intn(max-min) + min
}
