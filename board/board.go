package board

import (
	"fmt"
)

type Cell struct {
	x    int
	y    int
	val  int
	eval int
}

type Row struct {
	num   int
	cells []Cell
}

type Board struct {
	Turn bool
	rows []Row
}

// board constant
const (
	BoardSizeX = 8
	BoardSizeY = 8
	SpaceVal   = 0
	WhiteVal   = 1
	BlackVal   = 2
)

func (b *Board) Initial(evals [][]int) {
	b.Turn = true
	for i := 0; i < BoardSizeX; i++ {
		var row Row
		for j := 0; j < BoardSizeY; j++ {
			cell := Cell{x: i, y: j, val: initVal(i, j), eval: evals[i][j]}
			row.num = i
			row.cells = append(row.cells, cell)
		}
		b.rows = append(b.rows, row)
	}
}

func (b *Board) Reverse(x, y, myVal, yourVal int) bool {
	reverceCells := b.ReverceCells(x, y, myVal, yourVal)
	if reverceCells != nil {
		cellReverce(reverceCells, b, myVal)
		return true
	}
	return false
}

func (b *Board) ReverceCells(x, y, myVal, yourVal int) []Cell {
	if b.rows[x].cells[y].val != SpaceVal {
		return nil
	}
	var children Children
	children.set(b, x, y)
	var reverceCells []Cell
	for _, cell := range children.cells {
		if cell.val != yourVal {
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
			if b.rows[xTarget].cells[yTarget].val == myVal {
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
		return WhiteVal
	}
	if x == 3 && y == 4 || x == 4 && y == 3 {
		return BlackVal
	}
	return SpaceVal
}

func cellReverce(cells []Cell, b *Board, color int) {
	for _, cell := range cells {
		b.rows[cell.x].cells[cell.y].val = color
	}
}

func (b *Board) Print() {
	for _, row := range b.rows {
		for key, cell := range row.cells {
			if cell.val == SpaceVal {
				fmt.Print(" - ")
			}
			if cell.val == WhiteVal {
				fmt.Print(" ● ")
			}
			if cell.val == BlackVal {
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
			if cell.val == SpaceVal {
				return false
			}
		}
	}
	return true
}

func (b *Board) HasToPut(turn bool) bool {
	myVal, yourVal := getColor(turn)
	for _, row := range b.rows {
		for _, cell := range row.cells {
			if b.ReverceCells(cell.x, cell.y, myVal, yourVal) != nil {
				return true
			}
		}
	}
	return false
}

func (b *Board) GetToPutNextCells(turn bool) []Cell {
	var cells []Cell
	myVal, yourVal := getColor(turn)
	for _, row := range b.rows {
		for _, cell := range row.cells {
			if b.ReverceCells(cell.x, cell.y, myVal, yourVal) != nil {
				cells = append(cells, cell)
			}
		}
	}
	return cells
}

func getColor(turn bool) (myVal int, yourVal int) {
	if turn {
		return WhiteVal, BlackVal
	}
	return BlackVal, WhiteVal

}
