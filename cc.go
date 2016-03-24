package cc

// Solve populates solutions with valid boards of the dimensions columns*rows
// with valid configurations for pieces.
func Solve(columns, rows uint8, pieces []Piece, solutions *map[string]bool) {
	np := len(pieces)

	// No possible solutions
	if np != 1 && np >= int(columns*rows) {
		return
	}

	b := NewBoard(columns, rows)
	place(b, pieces, solutions)
}

// place tries to place the next piece on the board and recurse down
// the search tree. Appends the board and returns when a valid configuration
// is found.
func place(board Board, pieces []Piece, solutions *map[string]bool) {
	if len(pieces) == 0 {
		(*solutions)[board.Notation()] = true
		return
	}

	// Shift the pieces to get the first
	p, pieces := pieces[0], pieces[1:]

	for j := uint8(0); j < board.rows; j++ {
		for i := uint8(0); i < board.columns; i++ {
			c := board.cells[j][i]
			if c == Cell(Dead) || c != Cell(Blank) {
				continue
			}

			// Check so we don't threaten a placed piece
			canPlace := true
			tr := p.Threatening(&board, i, j)
			for _, t := range tr {
				tc := board.cells[t.y][t.x]
				if tc > Cell(Dead) {
					canPlace = false
				}
			}

			if canPlace {

				// Create a copy of the current board to use when reqursing down
				b2 := Board{
					columns: board.columns,
					rows:    board.rows,
					cells:   make([][]Cell, len(board.cells)),
				}
				for r := uint8(0); r < board.rows; r++ {
					b2.cells[r] = make([]Cell, len(board.cells[r]))
					copy(b2.cells[r], board.cells[r])
				}

				cc := Cell(0)
				switch p {
				case King:
					cc = Cell(King)
				case Rook:
					cc = Cell(Rook)
				case Queen:
					cc = Cell(Queen)
				case Bishop:
					cc = Cell(Bishop)
				case Knight:
					cc = Cell(Knight)
				}
				b2.cells[j][i] = cc // Place the piece

				// Mark all dead cells
				for _, t := range tr {
					b2.cells[t.y][t.x] = Cell(Dead)
				}

				// Recurse down with new board
				place(b2, pieces, solutions)
			}
		}
	}
}
