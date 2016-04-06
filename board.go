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
	cells   [][]cell
}

// NewBoard creates a new board of dimensions (columns*rows) with only
// empty cells.
func NewBoard(columns, rows uint8) Board {
	c := make([][]cell, rows)

	for j := uint8(0); j < rows; j++ {
		c[j] = make([]cell, columns)
		for i := uint8(0); i < columns; i++ {
			c[j][i] = cell(Blank)
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
		if len(p) != 3 {
			return b, fmt.Errorf("cc: the notation %s is not valid", p)
		}
		v := string(p[0])
		xs := p[1]
		ys := p[2]

		var cc cell
		switch v {
		case "K":
			cc = cell(King)
		case "R":
			cc = cell(Rook)
		case "Q":
			cc = cell(Queen)
		case "B":
			cc = cell(Bishop)
		case "N":
			cc = cell(Knight)
		default:
			return b, fmt.Errorf("cc: piece %s is not recognized", v)
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
			case cell(Blank), cell(Dead):
				continue
			case cell(King):
				cc = King
			case cell(Rook):
				cc = Rook
			case cell(Queen):
				cc = Queen
			case cell(Bishop):
				cc = Bishop
			case cell(Knight):
				cc = Knight
			}

			ret += fmt.Sprintf("%s%s%d,", cc, string(97+i), j+1)
		}
	}

	return ret[0 : len(ret)-1]
}

// ASCII returns the board configuration as an ascii drawing
// dot (.) marks empty cells, X marks dead cells, and pieces are represented
// by their individual symbol.
func (b Board) ASCII() string {
	ret := ""

	for j := uint8(0); j < b.rows; j++ {
		for i := uint8(0); i < b.columns; i++ {
			c := b.cells[j][i]
			var cc Piece
			switch c {
			case cell(Blank):
				cc = Blank
			case cell(Dead):
				cc = Dead
			case cell(King):
				cc = King
			case cell(Rook):
				cc = Rook
			case cell(Queen):
				cc = Queen
			case cell(Bishop):
				cc = Bishop
			case cell(Knight):
				cc = Knight
			}

			ret += fmt.Sprintf(" %s", cc)
		}

		ret += "\n"
	}

	return ret
}
