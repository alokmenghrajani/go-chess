package core

import (
	"fmt"
	"testing"
)

func TestPawnListMoves(t *testing.T) {
	// move b one or two
	b, _ := Parse(
		"........\n"+
			"......k.\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"...K....\n"+
			".P......\n"+
			"........\n", White)
	moves := b.board[1][1].piece.list_moves(b, xy{1, 1})
	expecting := "b2→b3,b2→b4"
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}

	// move by one only
	b, _ = Parse(
		"........\n"+
			"......k.\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			".P.K....\n"+
			"........\n"+
			"........\n", White)
	moves = b.board[1][2].piece.list_moves(b, xy{1, 2})
	expecting = "b3→b4"
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}

	// captures
	b, _ = Parse(
		"........\n"+
			"......k.\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"r.r.....\n"+
			".P...K..\n"+
			"........\n", White)
	moves = b.board[1][1].piece.list_moves(b, xy{1, 1})
	expecting = "b2→b3,b2→b4,b2→a3,b2→c3"
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}

	b, _ = Parse(
		"........\n"+
			"......k.\n"+
			"........\n"+
			"........\n"+
			"rr......\n"+
			"P.......\n"+
			".....K..\n"+
			"........\n", White)
	moves = b.board[0][2].piece.list_moves(b, xy{0, 2})
	expecting = "a3→b4"
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}

	b, _ = Parse(
		"........\n"+
			"......k.\n"+
			"......r.\n"+
			".......P\n"+
			"........\n"+
			"........\n"+
			".....K..\n"+
			"........\n", White)
	moves = b.board[7][4].piece.list_moves(b, xy{7, 4})
	expecting = "h5→h6,h5→g6"
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}

	// promotions
	b, _ = Parse(
		"........\n"+
			"..P...k.\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			".....K..\n"+
			"........\n", White)
	moves = b.board[2][6].piece.list_moves(b, xy{2, 6})
	expecting = "c7→c8Q,c7→c8R,c7→c8B,c7→c8N"
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}

	// en passant
	b, _ = Parse(
		"........\n"+
			"......k.\n"+
			"........\n"+
			".Pp.....\n"+
			"........\n"+
			"........\n"+
			".....K..\n"+
			"........\n", White)
	b.en_passant = 2
	moves = b.board[1][4].piece.list_moves(b, xy{1, 4})
	expecting = "b5→b6,b5→c6e.p."
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}
}
