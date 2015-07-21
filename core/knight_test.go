package core

import (
	"fmt"
	"testing"
)

func TestKnightListMoves(t *testing.T) {
	b, _ := Parse(
		"........\n"+
			"......k.\n"+
			"........\n"+
			"........\n"+
			"......K.\n"+
			"..N.....\n"+
			"........\n"+
			"........\n", White)
	moves := b.board[2][2].piece.list_moves(b, xy{2, 2})
	expecting := "c3→a2,c3→a4,c3→b1,c3→b5,c3→d1,c3→d5,c3→e2,c3→e4"
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}

	b, _ = Parse(
		"........\n"+
			"......k.\n"+
			"........\n"+
			"........\n"+
			"..P..K..\n"+
			"...p....\n"+
			".N......\n"+
			"........\n", White)
	moves = b.board[1][1].piece.list_moves(b, xy{1, 1})
	expecting = "b2→a4,b2→d1,b2→d3"
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}
}
