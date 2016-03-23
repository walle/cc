package cc

import (
	"testing"
)

func TestAscii(t *testing.T) {
	b := NewBoard(2, 2)
	expected := "..\n..\n\n"
	if b.Ascii() != expected {
		t.Errorf("Expected %s got %s", expected, b.Ascii())
	}
	expected = "R.\n.R\n\n"
	b.cells[0].piece = Rook
	b.cells[3].piece = Rook
	if b.Ascii() != expected {
		t.Errorf("Expected %s got %s", expected, b.Ascii())
	}
	expected = "RX\nXR\n\n"
	b.cells[0].piece = Rook
	b.cells[1].dead = true
	b.cells[2].dead = true
	b.cells[3].piece = Rook
	if b.Ascii() != expected {
		t.Errorf("Expected %s got %s", expected, b.Ascii())
	}
}
