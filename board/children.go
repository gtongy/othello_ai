package board

type Children struct {
	cells []Cell
}

func (c *Children) set(board *Board, x, y int) {
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i >= 0 && i <= 7 && j >= 0 && j <= 7 {
				c.cells = append(c.cells, board.rows[i].cells[j])
			}
		}
	}
}
