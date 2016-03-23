package cc

// Cell is one cell on a board.
type Cell struct {
	x     int
	y     int
	piece Piece
	dead  bool
	skip  bool
}