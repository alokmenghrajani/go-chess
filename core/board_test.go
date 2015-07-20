package core

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	b, _ := Parse(
		"rnbqkbnr\n"+
			"pppppppp\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"PPPPPPPP\n"+
			"RNBQKBNR\n", White)
	// White pieces
	_, ok := b.board[0][1].piece.(*pawn)
	if !ok {
		t.Error("a2 does not contain a pawn:", b.board[0][1].piece)
	}
	_, ok = b.board[0][0].piece.(*rook)
	if !ok {
		t.Error("a1 does not contain a rook:", b.board[0][0].piece)
	}
	_, ok = b.board[1][0].piece.(*knight)
	if !ok {
		t.Error("b1 does not contain a knight:", b.board[1][0].piece)
	}
	_, ok = b.board[2][0].piece.(*bishop)
	if !ok {
		t.Error("c1 does not contain a bishop:", b.board[2][0].piece)
	}
	_, ok = b.board[3][0].piece.(*queen)
	if !ok {
		t.Error("d1 does not contain a queen:", b.board[3][0].piece)
	}
	_, ok = b.board[4][0].piece.(*king)
	if !ok {
		t.Error("e1 does not contain a king:", b.board[4][0].piece)
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 2; j++ {
			if b.board[i][j].color != White {
				t.Error("expecting White piece: ", i, j)
			}
		}
	}

	// Black pieces
	_, ok = b.board[0][6].piece.(*pawn)
	if !ok {
		t.Error("a7 does not contain a pawn:", b.board[0][6].piece)
	}
	_, ok = b.board[0][7].piece.(*rook)
	if !ok {
		t.Error("a8 does not contain a rook:", b.board[0][7].piece)
	}
	_, ok = b.board[1][7].piece.(*knight)
	if !ok {
		t.Error("b8 does not contain a knight:", b.board[1][7].piece)
	}
	_, ok = b.board[2][7].piece.(*bishop)
	if !ok {
		t.Error("c8 does not contain a bishop:", b.board[2][7].piece)
	}
	_, ok = b.board[3][7].piece.(*queen)
	if !ok {
		t.Error("d8 does not contain a queen:", b.board[3][7].piece)
	}
	_, ok = b.board[4][7].piece.(*king)
	if !ok {
		t.Error("e8 does not contain a king:", b.board[4][7].piece)
	}
	for i := 0; i < 8; i++ {
		for j := 6; j < 8; j++ {
			if b.board[i][j].color != Black {
				t.Error("expecting Black piece: ", i, j)
			}
		}
	}

	// Empty cells
	for i := 0; i < 8; i++ {
		for j := 2; j < 6; j++ {
			if b.board[i][j].piece != nil {
				t.Error("expecting empty cell:", i, j, b.board[i][j].piece)
			}
		}
	}
}

func TestString(t *testing.T) {
	s :=
		"8 r n b q k b n r \n" +
			"7 p p p p p p p p \n" +
			"6 . . . . . . . . \n" +
			"5 . . . . . . . . \n" +
			"4 . . . . . . . . \n" +
			"3 . . . . . . . . \n" +
			"2 P P P P P P P P \n" +
			"1 R N B Q K B N R \n"
	b, _ := Parse(s, White)
	s = s + "  a b c d e f g h "
	if b.String() != s {
		t.Error(fmt.Sprintf(">%s<\n\n>%s<\n", b, s))
	}
}
