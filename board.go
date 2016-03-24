package cc

import "fmt"

// Board represents a chess board. It contains columns*rows cells that can be
// occupied by pieces.
type Board struct {
	columns uint8
	rows    uint8
	cells   [][]Cell
}

// NewBoard creates a new board of dimensions (columns*rows) with only
// empty cells.
func NewBoard(columns, rows uint8) Board {
	c := make([][]Cell, rows)
	for j := uint8(0); j < rows; j++ {
		c[j] = make([]Cell, columns)
		for i := uint8(0); i < columns; i++ {
			c[j][i] = Cell(Blank)
		}
	}
	return Board{columns: columns, rows: rows, cells: c}
}

// Notation returns the board configuration using common notation.
func (b Board) Notation() string {
	ret := ""
	for j := uint8(0); j < b.rows; j++ {
		for i := uint8(0); i < b.columns; i++ {
			var cc Piece
			c := b.cells[j][i]
			switch c {
			case Cell(Blank), Cell(Dead):
				continue
			case Cell(King):
				cc = King
			case Cell(Rook):
				cc = Rook
			case Cell(Queen):
				cc = Queen
			case Cell(Bishop):
				cc = Bishop
			case Cell(Knight):
				cc = Knight
			}
			ret += fmt.Sprintf("%s%s%d,", cc, string(97+i), j+1)
		}
	}
	return ret
}

// Ascii returns the board configuration as an ascii drawing
// dot (.) marks empty cells, X marks dead cells, and pieces are represented
// by their individual symbol.
func (b Board) Ascii() string {
	ret := ""
	for j := uint8(0); j < b.rows; j++ {
		for i := uint8(0); i < b.columns; i++ {
			c := b.cells[j][i]
			var cc Piece
			switch c {
			case Cell(Blank):
				cc = Blank
			case Cell(Dead):
				cc = Dead
			case Cell(King):
				cc = King
			case Cell(Rook):
				cc = Rook
			case Cell(Queen):
				cc = Queen
			case Cell(Bishop):
				cc = Bishop
			case Cell(Knight):
				cc = Knight
			}
			ret += fmt.Sprintf(" %s", cc)
		}
		ret += "\n"
	}
	ret += "\n"
	return ret
}
