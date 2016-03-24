package cc

import (
	"testing"
)

var threateningTests = []struct {
	c   uint8 // columns on board
	r   uint8 // rows on board
	p   Piece // piece to test
	x   uint8 // column to place piece
	y   uint8 // row to place piece
	out int   // Number of threatened cells
}{
	{1, 1, King, 0, 0, 0},
	{2, 2, King, 0, 0, 3},
	{100, 100, King, 0, 0, 3},
	{100, 100, King, 50, 50, 8},
	{100, 100, King, 99, 99, 3},
	{2, 2, King, 5, 3, 0}, // Out of bounds
	{1, 1, Rook, 0, 0, 0},
	{2, 2, Rook, 0, 0, 2},
	{10, 10, Rook, 0, 0, 18},
	{10, 10, Rook, 9, 9, 18},
	{10, 10, Rook, 4, 4, 18},
	{2, 2, Rook, 5, 3, 0}, // Out of bounds
	{1, 1, Knight, 0, 0, 0},
	{2, 2, Knight, 0, 0, 0},
	{5, 5, Knight, 2, 2, 8},
	{5, 5, Knight, 4, 4, 2},
	{5, 5, Knight, 0, 0, 2},
	{5, 5, Knight, 2, 0, 4},
	{2, 2, Knight, 5, 3, 0}, // Out of bounds
	{1, 1, Queen, 0, 0, 0},
	{2, 2, Queen, 0, 0, 3},
	{5, 5, Queen, 2, 2, 16},
	{5, 5, Queen, 4, 4, 12},
	{5, 5, Queen, 0, 0, 12},
	{2, 2, Queen, 5, 3, 0}, // Out of bounds
	{1, 1, Bishop, 0, 0, 0},
	{2, 2, Bishop, 0, 0, 1},
	{5, 5, Bishop, 2, 2, 8},
	{5, 5, Bishop, 4, 4, 4},
	{5, 5, Bishop, 0, 0, 4},
	{2, 2, Bishop, 5, 3, 0}, // Out of bounds
	{1, 1, Blank, 0, 0, 0},
	{1, 1, Dead, 0, 0, 0},
}

func TestThretening(t *testing.T) {
	for _, c := range threateningTests {
		b := NewBoard(c.c, c.r)
		tr := c.p.Threatening(&b, c.x, c.y)
		if len(tr) != c.out {
			t.Errorf("Expected %d got %d: %+v", c.out, len(tr), c)
		}
	}
}

// Benchmarks

var smallBoard = NewBoard(3, 3)
var largeBoard = NewBoard(30, 30)

func BenchmarkThreteningKingSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tr := King.Threatening(&smallBoard, 1, 1)
		if len(tr) != 8 {
			b.Errorf("Expected %d got %d", 8, len(tr))
		}
	}
}

func BenchmarkThreteningKingLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tr := King.Threatening(&largeBoard, 10, 10)
		if len(tr) != 8 {
			b.Errorf("Expected %d got %d", 8, len(tr))
		}
	}
}

func BenchmarkThreteningRookSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tr := Rook.Threatening(&smallBoard, 1, 1)
		if len(tr) != 4 {
			b.Errorf("Expected %d got %d", 8, len(tr))
		}
	}
}

func BenchmarkThreteningRookLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tr := Rook.Threatening(&largeBoard, 14, 14)
		if len(tr) != 58 {
			b.Errorf("Expected %d got %d", 8, len(tr))
		}
	}
}

func BenchmarkThreteningKnightSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tr := Knight.Threatening(&smallBoard, 0, 0)
		if len(tr) != 2 {
			b.Errorf("Expected %d got %d", 2, len(tr))
		}
	}
}

func BenchmarkThreteningKnightLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tr := Knight.Threatening(&largeBoard, 10, 10)
		if len(tr) != 8 {
			b.Errorf("Expected %d got %d", 8, len(tr))
		}
	}
}

func BenchmarkThreteningQueenSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tr := Queen.Threatening(&smallBoard, 0, 0)
		if len(tr) != 6 {
			b.Errorf("Expected %d got %d", 6, len(tr))
		}
	}
}

func BenchmarkThreteningQueenLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tr := Queen.Threatening(&largeBoard, 10, 10)
		if len(tr) != 107 {
			b.Errorf("Expected %d got %d", 107, len(tr))
		}
	}
}

func BenchmarkThreteningBishopSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tr := Bishop.Threatening(&smallBoard, 0, 0)
		if len(tr) != 2 {
			b.Errorf("Expected %d got %d", 2, len(tr))
		}
	}
}

func BenchmarkThreteningBishopLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tr := Bishop.Threatening(&largeBoard, 10, 10)
		if len(tr) != 49 {
			b.Errorf("Expected %d got %d", 49, len(tr))
		}
	}
}
