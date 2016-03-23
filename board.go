package cc

import "fmt"

// Board represents a chess board. It contains columns*rows cells that can be
// occupied by pieces.
type Board struct {
	columns int
	rows    int
	cells   []Cell
}

// NewBoard creates a new board of dimensions (columns*rows) with only
// empty cells.
func NewBoard(columns, rows int) Board {
	c := make([]Cell, columns*rows)
	for j := 0; j < rows; j++ {
		for i := 0; i < columns; i++ {
			c[columns*j+i] = Cell{x: i, y: j, piece: nil, dead: false}
		}
	}
	return Board{columns: columns, rows: rows, cells: c}
}

// Notation returns the board configuration using common notation.
func (b Board) Notation() string {
	ret := ""
	for _, c := range b.cells {
		if c.piece != nil {
			ret += fmt.Sprintf("%s%s%d,", c.piece, string(97+c.x), c.y+1)
		}
	}
	return ret
}

// Ascii returns the board configuration as an ascii drawing
// dot (.) marks empty cells, X marks dead cells, and pieces are represented
// by their individual symbol.
func (b Board) Ascii() string {
	ret := ""
	for j := 0; j < b.rows; j++ {
		for i := 0; i < b.columns; i++ {
			c := b.cells[b.columns*j+i]
			if c.piece == nil {
				if c.dead {
					ret += "X"
				} else {
					ret += "."
				}
			} else {
				ret += fmt.Sprintf("%s", c.piece)
			}
		}
		ret += "\n"
	}
	ret += "\n"
	return ret
}
