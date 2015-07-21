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
	expecting = "h4→g3,h4→g4,h4→h5"
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}

	b, _ = Parse(
		"r.......\n"+
			"......k.\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			".P......\n"+
			"K.......\n", White)
	moves = b.board[0][0].piece.list_moves(b, xy{0, 0})
	expecting = "a1→b1"
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}

	// long castling
	b, _ = Parse(
		"........\n"+
			"......k.\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"R...K..R\n", White)
	b.has_moved_king[0] = false
	b.has_moved_rook[0][0] = false
	moves = b.board[4][0].piece.list_moves(b, xy{4, 0})
	expecting = "e1→d1,e1→d2,e1→e2,e1→f1,e1→f2,0-0-0"
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}

	// short castling
	b, _ = Parse(
		"........\n"+
			"......k.\n"+
			".......r\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"R...K..R\n", White)
	b.has_moved_king[0] = false
	b.has_moved_rook[0][1] = false
	moves = b.board[4][0].piece.list_moves(b, xy{4, 0})
	expecting = "e1→d1,e1→d2,e1→e2,e1→f1,e1→f2,0-0"
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}

	// don't jump over check when castling
	b, _ = Parse(
		"........\n"+
			"......k.\n"+
			".....r..\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"....K..R\n", White)
	b.has_moved_king[0] = false
	b.has_moved_rook[0][1] = false
	moves = b.board[4][0].piece.list_moves(b, xy{4, 0})
	expecting = "e1→d1,e1→d2,e1→e2"
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}

	// don't castle into check
	b, _ = Parse(
		"........\n"+
			"......k.\n"+
			"......r.\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"....K..R\n", White)
	b.has_moved_king[0] = false
	b.has_moved_rook[0][1] = false
	moves = b.board[4][0].piece.list_moves(b, xy{4, 0})
	expecting = "e1→d1,e1→d2,e1→e2,e1→f1,e1→f2"
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}

	// don't castle when in check
	b, _ = Parse(
		"........\n"+
			"......k.\n"+
			"....r...\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"....K..R\n", White)
	b.has_moved_king[0] = false
	b.has_moved_rook[0][1] = false
	moves = b.board[4][0].piece.list_moves(b, xy{4, 0})
	expecting = "e1→d1,e1→d2,e1→f1,e1→f2"
	if moves.String() != expecting {
		t.Error(fmt.Sprintf("expecting %s, got %s", expecting, moves))
	}
}
