package board

import (
	"fmt"
)

type Cell struct {
	x   int
	y   int
	val int
}

type Row struct {
	num   int
	cells []Cell
}

type Board struct {
	Turn bool // true: my_turn, false: your_turn
	rows []Row
}

const (
	BOARD_SIZE_X = 8
	BOARD_SIZE_Y = 8
	SPACE_VAL    = 0
	WHITE_VAL    = 1
	BLACK_VAL    = 2
)

func (b *Board) Initial() {
	b.Turn = true
	for i := 0; i < BOARD_SIZE_X; i++ {
		var row Row
		for j := 0; j < BOARD_SIZE_Y; j++ {
			cell := Cell{x: i, y: j, val: initVal(i, j)}
			row.num = i
			row.cells = append(row.cells, cell)
		}
		b.rows = append(b.rows, row)
	}
}

func (b *Board) Reverse(x, y, my_val, your_val int) bool {
	reverceCells := b.ReverceCells(x, y, my_val, your_val)
	if reverceCells != nil {
		cellReverce(reverceCells, b, my_val)
		return true
	}
	return false
}

func (b *Board) ReverceCells(x, y, my_val, your_val int) []Cell {
	if b.rows[x].cells[y].val != SPACE_VAL {
		return nil
	}
	var children Children
	children.set(b, x, y)
	var reverceCells []Cell
	for _, cell := range children.cells {
		if cell.val != your_val {
			continue
		}
		var targetCells []Cell
		targetCells = append(targetCells, b.rows[x].cells[y])
		xIncrease := cell.x - x
		yIncrease := cell.y - y
		xTarget := x + xIncrease
		yTarget := y + yIncrease
		var exists bool
		exists = false
		for xTarget >= 0 && xTarget <= 7 && yTarget >= 0 && yTarget <= 7 {
			targetCells = append(targetCells, b.rows[xTarget].cells[yTarget])
			if b.rows[xTarget].cells[yTarget].val == my_val {
				exists = true
				break
			}
			xTarget += xIncrease
			yTarget += yIncrease
		}
		if exists {
			for _, targetCell := range targetCells {
				reverceCells = append(reverceCells, targetCell)
			}
		}
	}
	return reverceCells
}

func initVal(x, y int) int {
	if x == 3 && y == 3 || x == 4 && y == 4 {
		return WHITE_VAL
	}
	if x == 3 && y == 4 || x == 4 && y == 3 {
		return BLACK_VAL
	}
	return SPACE_VAL
}

func cellReverce(cells []Cell, b *Board, color int) {
	for _, cell := range cells {
		b.rows[cell.x].cells[cell.y].val = color
	}
}

func (b *Board) Print() {
	for _, row := range b.rows {
		for key, cell := range row.cells {
			if cell.val == SPACE_VAL {
				fmt.Print(" - ")
			}
			if cell.val == WHITE_VAL {
				fmt.Print(" ● ")
			}
			if cell.val == BLACK_VAL {
				fmt.Print(" ○ ")
			}
			if key == len(row.cells)-1 {
				fmt.Print("\n")
			}
		}
	}
}

func (b *Board) EndGame() bool {
	for _, row := range b.rows {
		for _, cell := range row.cells {
			if cell.val == SPACE_VAL {
				return false
			}
		}
	}
	return true
}

func (b *Board) HasToPut(turn bool) bool {
	var my_val int
	var your_val int
	if turn {
		my_val = WHITE_VAL
		your_val = BLACK_VAL
	} else {
		my_val = BLACK_VAL
		your_val = WHITE_VAL
	}
	for _, row := range b.rows {
		for _, cell := range row.cells {
			if b.ReverceCells(cell.x, cell.y, my_val, your_val) != nil {
				return true
			}
		}
	}
	return false
}
