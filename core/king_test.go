package core

import (
	"fmt"
	"testing"
)

func TestKingListMoves(t *testing.T) {
	b, _ := Parse(
		"......k.\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"..K.....\n"+
			"........\n"+
			"........\n", White)
	moves := b.board[2][2].piece.list_moves(b, xy{2, 2})
	expecting := "c3→b2,c3→b3,c3→b4,c3→c2,c3→c4,c3→d2,c3→d3,c3→d4"
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}

	b, _ = Parse(
		"........\n"+
			"......k.\n"+
			"........\n"+
			"......P.\n"+
			"......pK\n"+
			"........\n"+
			"........\n"+
			"........\n", White)
	moves = b.board[7][3].piece.list_moves(b, xy{7, 3})
	expecting = "h4→g3,h4→g4,h4→h3,h4→h5"
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}
}
