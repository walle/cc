package cc

import (
	"fmt"
	"math"
)

// Piece is the interface used to be able to use different implementation details
// for the different kinds of pieces.
// Since it only have one method it should be called Threatener according to Go
// best practices, but Piece carries more description of what really is.
type Piece interface {
	Threatening(b *Board, x, y int) []Cell
}

type king string
type queen string
type bishop string
type rook string
type knight string

// The pieces available
// Since they don't carry state use one instace
const (
	King   king   = "K" //"♚"
	Queen  queen  = "Q" //"♛"
	Bishop bishop = "B" //"♝"
	Rook   rook   = "R" //"♜"
	Knight knight = "N" //"♞"
)

func (p rook) Threatening(b *Board, x, y int) []Cell {
	if x < 0 || y < 0 || x >= b.columns || y >= b.rows {
		return []Cell{}
	}
	ret := make([]Cell, 0)
	for i := 0; i < b.columns; i++ {
		for j := 0; j < b.rows; j++ {
			if i == x && j != y {
				ret = append(ret, Cell{x: i, y: j})
			}
			if i != x && j == y {
				ret = append(ret, Cell{x: i, y: j})
			}
		}
	}
	return ret
}

func (p king) Threatening(b *Board, x, y int) []Cell {
	if x < 0 || y < 0 || x >= b.columns || y >= b.rows {
		return []Cell{}
	}
	t := []Cell{
		Cell{x: x - 1, y: y},
		Cell{x: x - 1, y: y - 1},
		Cell{x: x, y: y - 1},
		Cell{x: x + 1, y: y - 1},
		Cell{x: x + 1, y: y},
		Cell{x: x + 1, y: y + 1},
		Cell{x: x, y: y + 1},
		Cell{x: x - 1, y: y + 1},
	}
	ret := make([]Cell, 0)
	for _, c := range t {
		if c.x < 0 || c.y < 0 || c.x >= b.columns || c.y >= b.rows {
			continue
		}
		ret = append(ret, c)
	}
	return ret
}

func (p knight) Threatening(b *Board, x, y int) []Cell {
	if x < 0 || y < 0 || x >= b.columns || y >= b.rows {
		return []Cell{} // TODO: Should be error
	}
	t := []Cell{
		Cell{x: x - 1, y: y - 2},
		Cell{x: x + 1, y: y - 2},
		Cell{x: x - 2, y: y - 1},
		Cell{x: x + 2, y: y - 1},
		Cell{x: x - 2, y: y + 1},
		Cell{x: x + 2, y: y + 1},
		Cell{x: x - 1, y: y + 2},
		Cell{x: x + 1, y: y + 2},
	}
	ret := make([]Cell, 0)
	for _, c := range t {
		if c.x < 0 || c.y < 0 || c.x >= b.columns || c.y >= b.rows {
			continue
		}
		ret = append(ret, c)
	}
	return ret
}

func (p queen) Threatening(b *Board, x, y int) []Cell {
	if x < 0 || y < 0 || x >= b.columns || y >= b.rows {
		return []Cell{}
	}
	ret := make([]Cell, 0)
	for i := 0; i < b.columns; i++ {
		for j := 0; j < b.rows; j++ {
			if i == x || j == y || math.Abs(float64(j-y)) == math.Abs(float64(i-x)) {
				if i == x && j == y {
					continue
				}
				ret = append(ret, Cell{x: i, y: j})
			}
		}
	}
	return ret
}

func (p bishop) Threatening(b *Board, x, y int) []Cell {
	if x < 0 || y < 0 || x >= b.columns || y >= b.rows {
		return []Cell{}
	}
	ret := make([]Cell, 0)
	i := x
	j := y
	for {
		if i < 0 || i >= b.columns {
			break
		}
		if j < 0 || j >= b.rows {
			break
		}
		i++
		j++
		ret = append(ret, Cell{x: i, y: j})
	}
	i = x
	j = y
	for {
		if i < 0 || i >= b.columns {
			break
		}
		if j < 0 || j >= b.rows {
			break
		}
		i++
		j--
		ret = append(ret, Cell{x: i, y: j})
	}
	i = x
	j = y
	for {
		if i < 0 || i >= b.columns {
			break
		}
		if j < 0 || j >= b.rows {
			break
		}
		i--
		j--
		ret = append(ret, Cell{x: i, y: j})
	}
	i = x
	j = y
	for {
		if i < 0 || i >= b.columns {
			break
		}
		if j < 0 || j >= b.rows {
			break
		}
		i--
		j++
		ret = append(ret, Cell{x: i, y: j})
	}

	r := make([]Cell, 0)
	for _, c := range ret {
		if c.x < 0 || c.y < 0 || c.x >= b.columns || c.y >= b.rows {
			continue
		}
		r = append(r, c)
	}

	ret2 := make([]Cell, 0)
	m := make(map[string]Cell)
	for _, c := range r {
		m[fmt.Sprintf("%d,%d", c.x, c.y)] = c
	}

	for _, v := range m {
		ret2 = append(ret2, v)
	}

	return ret2
}
