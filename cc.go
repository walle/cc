package cc

import (
	"sync"
)

// Solve populates solutions with valid boards of the dimensions columns*rows
// with valid configurations for pieces.
func Solve(columns, rows uint8, pieces []Piece, solutions *map[string]bool) {
	np := len(pieces)

	// No possible solutions
	if np != 1 && np >= int(columns*rows) {
		return
	}

	wg := &sync.WaitGroup{}
	ch := make(chan string)

	// Start a goroutine for each possible starting position
	for j := uint8(0); j < rows; j++ {
		for i := uint8(0); i < columns; i++ {
			wg.Add(1)
			go func(i, j uint8) {
				b := NewBoard(columns, rows)
				// Shift the pieces to get the first
				p, pieces := pieces[0], pieces[1:]
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
				b.cells[j][i] = cc // Place the piece

				// Mark all dead cells
				tr := p.Threatening(&b, i, j)
				for _, t := range tr {
					b.cells[t.y][t.x] = Cell(Dead)
				}
				place(b, pieces, ch)
				wg.Done()
			}(i, j)
		}
	}

	// Syncronize the go routines by closing the channel when they are finished
	go func(wg *sync.WaitGroup, ch chan string) {
		wg.Wait()
		close(ch)
	}(wg, ch)

	// Syncronize the read from the channel so we dont exit to fast
	done := make(chan bool, 1)
	go func(ch <-chan string, done chan<- bool) {
		for s := range ch {
			(*solutions)[s] = true
		}
		done <- true
	}(ch, done)

	<-done
}

// place tries to place the next piece on the board and recurse down
// the search tree. Appends the board and returns when a valid configuration
// is found.
func place(board Board, pieces []Piece, ch chan<- string) {
	if len(pieces) == 0 {
		//(*solutions)[board.Notation()] = true
		ch <- board.Notation()
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
				place(b2, pieces, ch)
			}
		}
	}
}
