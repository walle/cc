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
	{1, 1, Rook, 0, 0, 0},
	{2, 2, Rook, 0, 0, 2},
	{10, 10, Rook, 0, 0, 18},
	{10, 10, Rook, 9, 9, 18},
	{10, 10, Rook, 4, 4, 18},
	{1, 1, Knight, 0, 0, 0},
	{2, 2, Knight, 0, 0, 0},
	{5, 5, Knight, 2, 2, 8},
	{5, 5, Knight, 4, 4, 2},
	{5, 5, Knight, 0, 0, 2},
	{5, 5, Knight, 2, 0, 4},
	{1, 1, Queen, 0, 0, 0},
	{2, 2, Queen, 0, 0, 3},
	{5, 5, Queen, 2, 2, 16},
	{5, 5, Queen, 4, 4, 12},
	{5, 5, Queen, 0, 0, 12},
	{1, 1, Bishop, 0, 0, 0},
	{2, 2, Bishop, 0, 0, 1},
	{5, 5, Bishop, 2, 2, 8},
	{5, 5, Bishop, 4, 4, 4},
	{5, 5, Bishop, 0, 0, 4},
	{1, 1, Blank, 0, 0, 0},
	{1, 1, Dead, 0, 0, 0},
}

func TestThreatening(t *testing.T) {
	for _, c := range threateningTests {
		b := NewBoard(c.c, c.r)
		tr, err := c.p.Threatening(&b, c.x, c.y)
		if err != nil {
			t.Errorf("Error occured: %s", err)
		}
		if len(tr) != c.out {
			t.Errorf("Expected %d got %d: %+v", c.out, len(tr), c)
		}
	}
}

func TestThreateningOutOfBounds(t *testing.T) {
	b := NewBoard(2, 2)
	_, err := King.Threatening(&b, 5, 3)
	if err == nil {
		t.Errorf("Expected out of bounds error")
	}
	_, err = Rook.Threatening(&b, 5, 3)
	if err == nil {
		t.Errorf("Expected out of bounds error")
	}
	_, err = Knight.Threatening(&b, 5, 3)
	if err == nil {
		t.Errorf("Expected out of bounds error")
	}
	_, err = Queen.Threatening(&b, 5, 3)
	if err == nil {
		t.Errorf("Expected out of bounds error")
	}
	_, err = Bishop.Threatening(&b, 5, 3)
	if err == nil {
		t.Errorf("Expected out of bounds error")
	}
}

// Benchmarks

var smallBoard = NewBoard(3, 3)
var largeBoard = NewBoard(30, 30)

func BenchmarkThreateningKingSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tr, _ := King.Threatening(&smallBoard, 1, 1)
		if len(tr) != 8 {
			b.Errorf("Expected %d got %d", 8, len(tr))
		}
	}
}

func BenchmarkThreateningKingLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tr, _ := King.Threatening(&largeBoard, 10, 10)
		if len(tr) != 8 {
			b.Errorf("Expected %d got %d", 8, len(tr))
		}
	}
}

func BenchmarkThreateningRookSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tr, _ := Rook.Threatening(&smallBoard, 1, 1)
		if len(tr) != 4 {
			b.Errorf("Expected %d got %d", 8, len(tr))
		}
	}
}

func BenchmarkThreateningRookLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tr, _ := Rook.Threatening(&largeBoard, 14, 14)
		if len(tr) != 58 {
			b.Errorf("Expected %d got %d", 8, len(tr))
		}
	}
}

func BenchmarkThreateningKnightSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tr, _ := Knight.Threatening(&smallBoard, 0, 0)
		if len(tr) != 2 {
			b.Errorf("Expected %d got %d", 2, len(tr))
		}
	}
}

func BenchmarkThreateningKnightLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tr, _ := Knight.Threatening(&largeBoard, 10, 10)
		if len(tr) != 8 {
			b.Errorf("Expected %d got %d", 8, len(tr))
		}
	}
}

func BenchmarkThreateningQueenSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tr, _ := Queen.Threatening(&smallBoard, 0, 0)
		if len(tr) != 6 {
			b.Errorf("Expected %d got %d", 6, len(tr))
		}
	}
}

func BenchmarkThreateningQueenLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tr, _ := Queen.Threatening(&largeBoard, 10, 10)
		if len(tr) != 107 {
			b.Errorf("Expected %d got %d", 107, len(tr))
		}
	}
}

func BenchmarkThreateningBishopSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tr, _ := Bishop.Threatening(&smallBoard, 0, 0)
		if len(tr) != 2 {
			b.Errorf("Expected %d got %d", 2, len(tr))
		}
	}
}

func BenchmarkThreateningBishopLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tr, _ := Bishop.Threatening(&largeBoard, 10, 10)
		if len(tr) != 49 {
			b.Errorf("Expected %d got %d", 49, len(tr))
		}
	}
}
