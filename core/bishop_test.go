package core

import (
	"fmt"
	"testing"
)

func TestBishopListMoves(t *testing.T) {
	b, _ := Parse(
		"......k.\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			".....K..\n"+
			"........\n"+
			".B......\n"+
			"........\n", White)
	moves := b.board[1][1].piece.list_moves(b, xy{1, 1})
	expecting := "b2→a1,b2→a3,b2→c1,b2→c3,b2→d4,b2→e5,b2→f6,b2→g7,b2→h8"
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}

	b, _ = Parse(
		"........\n"+
			"......k.\n"+
			"........\n"+
			"....p...\n"+
			"......K.\n"+
			"P.......\n"+
			".B......\n"+
			"........\n", White)
	moves = b.board[1][1].piece.list_moves(b, xy{1, 1})
	expecting = "b2→a1,b2→c1,b2→c3,b2→d4,b2→e5"
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}
}
