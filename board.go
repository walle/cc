package cc

import (
	"fmt"
	"strconv"
	"strings"
)

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

// NewBoardFromString takes a notation and converts it to a board.
// Returns the board on success and an error if something goes wrong.
func NewBoardFromString(columns, rows uint8, notation string) (Board, error) {
	b := NewBoard(columns, rows)
	ps := strings.Split(notation, ",")
	for _, p := range ps {
		v := string(p[0])
		xs := p[1]
		ys := p[2]

		var cc Cell
		switch v {
		case "K":
			cc = Cell(King)
		case "R":
			cc = Cell(Rook)
		case "Q":
			cc = Cell(Queen)
		case "B":
			cc = Cell(Bishop)
		case "N":
			cc = Cell(Knight)
		}

		x := uint8(int(xs) - 97)
		yi, _ := strconv.Atoi(string(ys))
		y := uint8(yi - 1)
		if x >= columns || y >= rows {
			return b, fmt.Errorf(
				"cc: piece at %d, %d can't fit on board (%d,%d)",
				x, y, columns-1, rows-1,
			)
		}

		b.cells[y][x] = cc
	}
	return b, nil
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

	return ret[0 : len(ret)-1]
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

	return ret
}
