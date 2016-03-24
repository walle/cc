package cc

import (
	"testing"
)

func TestAscii(t *testing.T) {
	b := NewBoard(3, 3)
	expected := " . . .\n . . .\n . . .\n\n"
	if b.Ascii() != expected {
		t.Errorf("Expected %s got %s", expected, b.Ascii())
	}
	expected = " R . .\n . R .\n . . .\n\n"
	b.cells[0][0] = Cell(Rook)
	b.cells[1][1] = Cell(Rook)
	if b.Ascii() != expected {
		t.Errorf("Expected %s got %s", expected, b.Ascii())
	}
	expected = " R X K\n X R Q\n N B .\n\n"
	b.cells[0][0] = Cell(Rook)
	b.cells[0][1] = Cell(Dead)
	b.cells[0][2] = Cell(King)
	b.cells[1][0] = Cell(Dead)
	b.cells[1][1] = Cell(Rook)
	b.cells[1][2] = Cell(Queen)
	b.cells[2][0] = Cell(Knight)
	b.cells[2][1] = Cell(Bishop)
	if b.Ascii() != expected {
		t.Errorf("Expected %s got %s", expected, b.Ascii())
	}
}
