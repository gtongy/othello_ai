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

func (b *Board) Update(x, y int) {
	for row_key, row := range b.rows {
		for cell_key, cell := range row.cells {
			if cell.x == x && cell.y == y {
				b.rows[row_key].cells[cell_key].val = WHITE_VAL
			}
		}
	}
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
