package cc

import "fmt"

// Piece is the interface used to be able to use different implementation details
// for the different kinds of pieces.
type Piece interface {
	String() string
	Threatening(b *Board, x, y uint8) []Position
}

type blank Cell
type dead Cell
type king Cell
type queen Cell
type bishop Cell
type rook Cell
type knight Cell

// The pieces available
// Since they don't carry state use one instace
const (
	Blank  blank  = 0 //"."
	Dead   dead   = 1 //"X"
	King   king   = 5 //"K" //"♚"
	Queen  queen  = 6 //"Q" //"♛"
	Bishop bishop = 7 //"B" //"♝"
	Rook   rook   = 8 //"R" //"♜"
	Knight knight = 9 //"N" //"♞"
)

func (p rook) Threatening(b *Board, x, y uint8) []Position {
	if x < 0 || y < 0 || x >= b.columns || y >= b.rows {
		return []Position{}
	}
	ret := make([]Position, 0)
	for i := uint8(0); i < b.columns; i++ {
		for j := uint8(0); j < b.rows; j++ {
			if i == x && j != y {
				ret = append(ret, Position{x: i, y: j})
			}
			if i != x && j == y {
				ret = append(ret, Position{x: i, y: j})
			}
		}
	}
	return ret
}

func (p rook) String() string {
	return "R"
}

func (p king) Threatening(b *Board, x, y uint8) []Position {
	if x < 0 || y < 0 || x >= b.columns || y >= b.rows {
		return []Position{}
	}
	t := []Position{
		{x: x - 1, y: y},
		{x: x - 1, y: y - 1},
		{x: x, y: y - 1},
		{x: x + 1, y: y - 1},
		{x: x + 1, y: y},
		{x: x + 1, y: y + 1},
		{x: x, y: y + 1},
		{x: x - 1, y: y + 1},
	}
	ret := make([]Position, 0)
	for _, c := range t {
		if c.x < 0 || c.y < 0 || c.x >= b.columns || c.y >= b.rows {
			continue
		}
		ret = append(ret, c)
	}
	return ret
}

func (p king) String() string {
	return "K"
}

func (p knight) Threatening(b *Board, x, y uint8) []Position {
	if x < 0 || y < 0 || x >= b.columns || y >= b.rows {
		return []Position{} // TODO: Should be error
	}
	t := []Position{
		{x: x - 1, y: y - 2},
		{x: x + 1, y: y - 2},
		{x: x - 2, y: y - 1},
		{x: x + 2, y: y - 1},
		{x: x - 2, y: y + 1},
		{x: x + 2, y: y + 1},
		{x: x - 1, y: y + 2},
		{x: x + 1, y: y + 2},
	}
	ret := make([]Position, 0)
	for _, c := range t {
		if c.x < 0 || c.y < 0 || c.x >= b.columns || c.y >= b.rows {
			continue
		}
		ret = append(ret, c)
	}
	return ret
}

func (p knight) String() string {
	return "N"
}

func (p queen) Threatening(b *Board, x, y uint8) []Position {
	if x < 0 || y < 0 || x >= b.columns || y >= b.rows {
		return []Position{}
	}

	ret := make([]Position, 0, b.columns*b.rows)
	ret = append(ret, Rook.Threatening(b, x, y)...)
	ret = append(ret, Bishop.Threatening(b, x, y)...)
	return ret
}

func (p queen) String() string {
	return "Q"
}

func (p bishop) Threatening(b *Board, x, y uint8) []Position {
	if x < 0 || y < 0 || x >= b.columns || y >= b.rows {
		return []Position{}
	}
	m := make(map[string]Position)
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
		m[fmt.Sprintf("%d,%d", i, j)] = Position{x: i, y: j}
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
		m[fmt.Sprintf("%d,%d", i, j)] = Position{x: i, y: j}
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
		m[fmt.Sprintf("%d,%d", i, j)] = Position{x: i, y: j}
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
		m[fmt.Sprintf("%d,%d", i, j)] = Position{x: i, y: j}
	}

	ret := make([]Position, 0, b.columns*b.rows)
	for _, c := range m {
		if c.x < 0 || c.y < 0 || c.x >= b.columns || c.y >= b.rows {
			continue
		}
		ret = append(ret, c)
	}
	return ret
}

func (p bishop) String() string {
	return "B"
}

func (p blank) String() string {
	return "."
}

func (p blank) Threatening(b *Board, x, y uint8) []Position {
	return []Position{}
}

func (p dead) String() string {
	return "X"
}

func (p dead) Threatening(b *Board, x, y uint8) []Position {
	return []Position{}
}
