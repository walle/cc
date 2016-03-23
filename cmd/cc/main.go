package main

import (
	"fmt"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/walle/cc"
)

var args struct {
	Columns int `arg:"-c,required,help:the number of columns to use on the board"`
	Rows    int `arg:"-r,required,help:the number of rows to use on the board"`
	Kings   int `arg:"-k,help:the number of kings to use on the board"`
	Queens  int `arg:"-q,help:the number of queens to use on the board"`
	Bishops int `arg:"-b,help:the number of bishops to use on the board"`
	Rooks   int `arg:"-t,help:the number of rooks to use on the board"`
	Knights int `arg:"-n,help:the number of knights to use on the board"`
}

func main() {
	arg.MustParse(&args)
	start := time.Now()

	n := args.Kings + args.Queens + args.Bishops + args.Rooks + args.Knights
	pieces := make([]cc.Piece, 0, n)
	for i := 0; i < args.Queens; i++ {
		pieces = append(pieces, cc.Queen)
	}
	for i := 0; i < args.Rooks; i++ {
		pieces = append(pieces, cc.Rook)
	}
	for i := 0; i < args.Bishops; i++ {
		pieces = append(pieces, cc.Bishop)
	}
	for i := 0; i < args.Knights; i++ {
		pieces = append(pieces, cc.Knight)
	}
	for i := 0; i < args.Kings; i++ {
		pieces = append(pieces, cc.King)
	}

	solutions := make([]cc.Board, 0)
	cc.Solve(args.Columns, args.Rows, pieces, &solutions)

	combs := ""
	for _, s := range solutions {
		combs += s.Notation() + "; "
	}

	fmt.Printf("[%d] %s(%s)\n", len(solutions), combs, time.Now().Sub(start))
}
