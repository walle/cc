package cc

import (
	"testing"
)

func TestAscii(t *testing.T) {
	b := NewBoard(3, 3)

	expected := " . . .\n . . .\n . . .\n"
	if b.ASCII() != expected {
		t.Errorf("Expected %s got %s", expected, b.ASCII())
	}

	expected = " R . .\n . R .\n . . .\n"
	b.cells[0][0] = cell(Rook)
	b.cells[1][1] = cell(Rook)
	if b.ASCII() != expected {
		t.Errorf("Expected %s got %s", expected, b.ASCII())
	}

	expected = " R X K\n X R Q\n N B .\n"
	b.cells[0][0] = cell(Rook)
	b.cells[0][1] = cell(Dead)
	b.cells[0][2] = cell(King)
	b.cells[1][0] = cell(Dead)
	b.cells[1][1] = cell(Rook)
	b.cells[1][2] = cell(Queen)
	b.cells[2][0] = cell(Knight)
	b.cells[2][1] = cell(Bishop)
	if b.ASCII() != expected {
		t.Errorf("Expected %s got %s", expected, b.ASCII())
	}
}

func TestNewBoardFromString(t *testing.T) {
	notation := "Kb2,Ke2,Qg3,Na6,Bb6,Bc6,Qf7"
	b, err := NewBoardFromString(7, 7, notation)
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	if b.Notation() != notation {
		t.Errorf("Expected %s got %s", notation, b.Notation())
	}
	b, err = NewBoardFromString(2, 2, notation)
	if err == nil {
		t.Errorf("Error did not occur")
	}
	notation = "Rb1,Ra2"
	b, err = NewBoardFromString(2, 2, notation)
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	if b.Notation() != notation {
		t.Errorf("Expected %s got %s", notation, b.Notation())
	}
	notation = "Rxyb1,Fa2"
	b, err = NewBoardFromString(2, 2, notation)
	if err == nil {
		t.Errorf("Error did not occur with misformed input")
	}
	notation = "Rb1,Fa2"
	b, err = NewBoardFromString(2, 2, notation)
	if err == nil {
		t.Errorf("Error did not occur with misformed input")
	}
}
