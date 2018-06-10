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
	if b.rows[x].cells[y].val != SPACE_VAL {
		return false
	}
	var children Children
	children.set(b, x, y)
	for _, cell := range children.cells {
		if cell.val != your_val {
			continue
		}
		var target_cells []Cell
		target_cells = append(target_cells, b.rows[x].cells[y])
		x_increase := cell.x - x
		y_increase := cell.y - y
		x_target := x + x_increase
		y_target := y + y_increase
		for x_target >= 0 && x_target <= 7 && y_target >= 0 && y_target <= 7 {
			target_cells = append(target_cells, b.rows[x_target].cells[y_target])
			if b.rows[x_target].cells[y_target].val == my_val {
				fmt.Println(target_cells)
				cellReverce(target_cells, b, my_val)
				return true
			}
			x_target += x_increase
			y_target += y_increase
		}
	}
	// TODO: split this condition
	fmt.Println("reverse cell is not exists")
	return false
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
