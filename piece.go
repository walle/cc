package cc

// Piece is the interface used to be able to use different implementation details
// for the different kinds of pieces.
type Piece interface {
	String() string
	Threatening(b *Board, x, y uint8) []position
}

type blank cell
type dead cell
type king cell
type queen cell
type bishop cell
type rook cell
type knight cell

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

func (p rook) Threatening(b *Board, x, y uint8) []position {
	if x < 0 || y < 0 || x >= b.columns || y >= b.rows {
		return []position{}
	}
	var ret []position
	for i := uint8(0); i < b.columns; i++ {
		for j := uint8(0); j < b.rows; j++ {
			if i == x && j != y {
				ret = append(ret, position{x: i, y: j})
			}
			if i != x && j == y {
				ret = append(ret, position{x: i, y: j})
			}
		}
	}
	return ret
}

func (p rook) String() string {
	return "R"
}

func (p king) Threatening(b *Board, x, y uint8) []position {
	if x < 0 || y < 0 || x >= b.columns || y >= b.rows {
		return []position{}
	}
	t := []position{
		{x: x - 1, y: y},
		{x: x - 1, y: y - 1},
		{x: x, y: y - 1},
		{x: x + 1, y: y - 1},
		{x: x + 1, y: y},
		{x: x + 1, y: y + 1},
		{x: x, y: y + 1},
		{x: x - 1, y: y + 1},
	}
	var ret []position
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

func (p knight) Threatening(b *Board, x, y uint8) []position {
	if x < 0 || y < 0 || x >= b.columns || y >= b.rows {
		return []position{} // TODO: Should be error
	}
	t := []position{
		{x: x - 1, y: y - 2},
		{x: x + 1, y: y - 2},
		{x: x - 2, y: y - 1},
		{x: x + 2, y: y - 1},
		{x: x - 2, y: y + 1},
		{x: x + 2, y: y + 1},
		{x: x - 1, y: y + 2},
		{x: x + 1, y: y + 2},
	}
	var ret []position
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

func (p queen) Threatening(b *Board, x, y uint8) []position {
	if x < 0 || y < 0 || x >= b.columns || y >= b.rows {
		return []position{}
	}

	ret := make([]position, 0, b.columns*b.rows)
	ret = append(ret, Rook.Threatening(b, x, y)...)
	ret = append(ret, Bishop.Threatening(b, x, y)...)
	return ret
}

func (p queen) String() string {
	return "Q"
}

func (p bishop) Threatening(b *Board, x, y uint8) []position {
	if x < 0 || y < 0 || x >= b.columns || y >= b.rows {
		return []position{}
	}
	n := b.columns
	if b.rows > n {
		n = b.rows
	}
	var ret []position
	for step := 1; step < int(n); step++ {
		x1 := uint8(int(x) + (step * -1))
		y1 := uint8(int(y) + (step * -1))
		x2 := uint8(int(x) + (step * 1))
		y2 := uint8(int(y) + (step * 1))
		if x1 < b.columns && y1 < b.rows {
			ret = append(ret, position{x: x1, y: y1})
		}
		if x2 < b.columns && y2 < b.rows {
			ret = append(ret, position{x: x2, y: y2})
		}
		if x1 < b.columns && y2 < b.rows {
			ret = append(ret, position{x: x1, y: y2})
		}
		if x2 < b.columns && y1 < b.rows {
			ret = append(ret, position{x: x2, y: y1})
		}
	}
	return ret
}

func (p bishop) String() string {
	return "B"
}

func (p blank) String() string {
	return "."
}

func (p blank) Threatening(b *Board, x, y uint8) []position {
	return []position{}
}

func (p dead) String() string {
	return "X"
}

func (p dead) Threatening(b *Board, x, y uint8) []position {
	return []position{}
}
