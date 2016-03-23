package cc

// Solve populates solutions with valid boards of the dimensions columns*rows
// with valid configurations for pieces.
func Solve(columns, rows int, pieces []Piece, solutions *[]Board) {
	np := len(pieces)

	b := NewBoard(columns, rows)
	place(b, pieces, solutions, 0)

	sm := make(map[string]Board)
	for _, s := range *solutions {
		sm[s.Notation()] = s
	}

	ret := make([]Board, 0)
	for _, v := range sm {
		ret = append(ret, v)
	}

	ret2 := make([]Board, 0)
	// Cull invalid
	for _, s := range ret {
		sum := 0
		for _, c := range s.cells {
			if c.piece != nil {
				sum++
			}
		}
		if sum == np {
			ret2 = append(ret2, s)
		}
	}

	//return ret
	*solutions = ret2
}

// place tries to place the next piece on the board and recurse down
// the search tree. Appends the board and returns when a valid configuration
// is found.
func place(board Board, pieces []Piece, solutions *[]Board, run int) {
	if len(pieces) == 0 {
		*solutions = append(*solutions, board)
		return
	}

	// Shift the pieces to get the first
	p, pieces := pieces[0], pieces[1:]

	for _, c := range board.cells {
		if c.dead || c.piece != nil {
			continue
		}
		canPlace := true
		tr := p.Threatening(&board, c.x, c.y)
		for _, t := range tr {
			tc := board.cells[board.columns*t.y+t.x]
			if tc.piece != nil {
				canPlace = false
			}
		}
		if canPlace {
			b2 := NewBoard(board.columns, board.rows)
			copy(b2.cells, board.cells)
			b2.cells[b2.columns*c.y+c.x].piece = p
			for _, t := range tr {
				b2.cells[b2.columns*t.y+t.x].dead = true
			}
			place(b2, pieces, solutions, run+1)
		}
	}
}
