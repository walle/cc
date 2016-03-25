package cc

import (
	"testing"
)

var solveTests = []struct {
	c   uint8   // columns on board
	r   uint8   // rows on board
	p   []Piece // The pieces to use
	out int     // The number of solutions
}{
	{2, 2, []Piece{Rook, Rook}, 2},
	{2, 2, []Piece{King, King}, 0},
	{2, 2, []Piece{Knight, Knight}, 6},
	{3, 2, []Piece{Bishop, Bishop}, 11},
	{3, 3, []Piece{Rook, King, King}, 4},
	{4, 4, []Piece{Rook, Rook, Knight, Knight, Knight, Knight}, 8},
	{3, 3, []Piece{Bishop, Bishop, Bishop, Bishop, Bishop, Bishop, Bishop, Bishop, Bishop, Bishop}, 0},
	{1, 1, []Piece{Queen}, 1},
	{2, 2, []Piece{Queen, Queen}, 0},
	{4, 4, []Piece{Queen, Queen, Queen, Queen}, 2},
	{5, 5, []Piece{Queen, Queen, Queen, Queen, Queen}, 10},
	{6, 6, []Piece{Queen, Queen, Queen, Queen, Queen, Queen}, 4},
	{7, 7, []Piece{Queen, Queen, Queen, Queen, Queen, Queen, Queen}, 40},
	{2, 2, []Piece{Rook, Rook, Rook}, 0},
	//{8, 8, []Piece{Queen, Queen, Queen, Queen, Queen, Queen, Queen, Queen}, 92},
	//{7, 7, []Piece{Queen, Queen, Bishop, Bishop, Knight, King, King}, 3062636},
}

func TestSolve(t *testing.T) {
	for _, tc := range solveTests {
		solutions := make(map[string]bool)
		Solve(tc.c, tc.r, tc.p, &solutions)
		if len(solutions) != tc.out {
			t.Errorf("Expected %d got %d: %+v", tc.out, len(solutions), tc)
		}
	}
}

// Benchmarks

func Benchmark2x2R2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions := make(map[string]bool)
		p := []Piece{Rook, Rook}
		Solve(2, 2, p, &solutions)
		if len(solutions) != 2 {
			b.Errorf("Expected %d got %d", 2, len(solutions))
		}
	}
}

func Benchmark3x3R1K2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions := make(map[string]bool)
		p := []Piece{Rook, King, King}
		Solve(3, 3, p, &solutions)
		if len(solutions) != 4 {
			b.Errorf("Expected %d got %d", 4, len(solutions))
		}
	}
}

func Benchmark4x4R2N4(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions := make(map[string]bool)
		p := []Piece{Rook, Rook, Knight, Knight, Knight, Knight}
		Solve(4, 4, p, &solutions)
		if len(solutions) != 8 {
			b.Errorf("Expected %d got %d", 8, len(solutions))
		}
	}
}

func Benchmark2Q(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions := make(map[string]bool)
		p := []Piece{Queen, Queen}
		Solve(2, 2, p, &solutions)
		if len(solutions) != 0 {
			b.Errorf("Expected %d got %d", 0, len(solutions))
		}
	}
}

func Benchmark4Q(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions := make(map[string]bool)
		p := []Piece{Queen, Queen, Queen, Queen}
		Solve(4, 4, p, &solutions)
		if len(solutions) != 2 {
			b.Errorf("Expected %d got %d", 2, len(solutions))
		}
	}
}

func Benchmark5Q(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions := make(map[string]bool)
		p := []Piece{Queen, Queen, Queen, Queen, Queen}
		Solve(5, 5, p, &solutions)
		if len(solutions) != 10 {
			b.Errorf("Expected %d got %d", 10, len(solutions))
		}
	}
}

func Benchmark6Q(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions := make(map[string]bool)
		p := []Piece{Queen, Queen, Queen, Queen, Queen, Queen}
		Solve(6, 6, p, &solutions)
		if len(solutions) != 4 {
			b.Errorf("Expected %d got %d", 4, len(solutions))
		}
	}
}

func Benchmark7Q(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions := make(map[string]bool)
		p := []Piece{Queen, Queen, Queen, Queen, Queen, Queen, Queen}
		Solve(7, 7, p, &solutions)
		if len(solutions) != 40 {
			b.Errorf("Expected %d got %d", 40, len(solutions))
		}
	}
}

func Benchmark8Q(b *testing.B) {
	b.Skip("Slow test")
	for n := 0; n < b.N; n++ {
		solutions := make(map[string]bool)
		p := []Piece{Queen, Queen, Queen, Queen, Queen, Queen, Queen, Queen}
		Solve(8, 8, p, &solutions)
		if len(solutions) != 92 {
			b.Errorf("Expected %d got %d", 92, len(solutions))
		}
	}
}

func Benchmark7x7Q2B2N1K2(b *testing.B) {
	b.Skip("Slow test")
	for n := 0; n < b.N; n++ {
		solutions := make(map[string]bool)
		p := []Piece{Queen, Queen, Bishop, Bishop, Knight, King, King}
		Solve(7, 7, p, &solutions)
		if len(solutions) != 3062636 {
			b.Errorf("Expected %d got %d", 3062636, len(solutions))
		}
	}
}

func Benchmark3x3B10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions := make(map[string]bool)
		p := []Piece{Bishop, Bishop, Bishop, Bishop, Bishop, Bishop, Bishop, Bishop, Bishop, Bishop}
		Solve(3, 3, p, &solutions)
		if len(solutions) != 0 {
			b.Errorf("Expected %d got %d", 0, len(solutions))
		}
	}
}
