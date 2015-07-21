package core

import (
	"fmt"
	"testing"
)

func TestQueenListMoves(t *testing.T) {
	b, _ := Parse(
		"......k.\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"......K.\n"+
			"........\n"+
			".Q......\n"+
			"........\n", White)
	moves := b.board[1][1].piece.list_moves(b, xy{1, 1})
	expecting := "b2→a1,b2→a2,b2→a3,b2→b1,b2→b3,b2→b4,b2→b5,b2→b6,b2→b7,b2→b8,b2→c1,b2→c2,b2→d2,b2→e2,b2→f2,b2→g2,b2→h2,b2→c3,b2→d4,b2→e5,b2→f6,b2→g7,b2→h8"
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}

	b, _ = Parse(
		"........\n"+
			"......k.\n"+
			"........\n"+
			"........\n"+
			".P.p.K..\n"+
			"........\n"+
			".Q..p...\n"+
			"........\n", White)
	moves = b.board[1][1].piece.list_moves(b, xy{1, 1})
	expecting = "b2→a1,b2→a2,b2→a3,b2→b1,b2→b3,b2→c1,b2→c2,b2→d2,b2→e2,b2→c3,b2→d4"
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}
}
