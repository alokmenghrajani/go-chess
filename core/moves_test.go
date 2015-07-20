package core

import (
	"fmt"
	"testing"
)

func TestNormalMovePerformMove(t *testing.T) {
	b, _ := Parse(
		"........\n"+
			".....k..\n"+
			"........\n"+
			".p......\n"+
			"........\n"+
			"...K....\n"+
			"........\n"+
			"........\n", White)
	b.en_passant = 1
	move := normal_move{abstract_move{xy{3, 2}, xy{4, 2}}}
	b = move.perform_move(b)
	_, ok := b.board[4][2].piece.(*king)
	if !ok {
		t.Error(fmt.Sprintf("expecting king, got %s", b.board[4][2]))
	}
	if b.en_passant != -2 {
		t.Error(fmt.Sprintf("expecting en-passant to be cleared, got %d",
			b.en_passant))
	}
}

func TestPawnPushTwoPerformMove(t *testing.T) {
	b, _ := Parse(
		"........\n"+
			".....k..\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"...K....\n"+
			".P......\n"+
			"........\n", White)
	move := pawn_push_two{abstract_move{xy{1, 1}, xy{1, 3}}}
	b = move.perform_move(b)
	_, ok := b.board[1][3].piece.(*pawn)
	if !ok {
		t.Error(fmt.Sprintf("expecting king, got %s", b.board[4][2]))
	}
	if b.en_passant != 1 {
		t.Error(fmt.Sprintf("expecting en-passant to be set, got %d",
			b.en_passant))
	}
}

func TestPawnPromotePerformMove(t *testing.T) {
	b, _ := Parse(
		"........\n"+
			".P...k..\n"+
			"........\n"+
			"........\n"+
			"........\n"+
			"...K....\n"+
			"........\n"+
			"........\n", White)
	move := pawn_promote{abstract_move{xy{1, 6}, xy{1, 7}}, new(queen)}
	b = move.perform_move(b)
	_, ok := b.board[1][7].piece.(*queen)
	if !ok {
		t.Error(fmt.Sprintf("expecting queen, got %s", b.board[1][7]))
	}
	if b.en_passant != -2 {
		t.Error(fmt.Sprintf("expecting en-passant to be cleared, got %d",
			b.en_passant))
	}
}

func TestPawnEnPassantPerformMove(t *testing.T) {
	b, _ := Parse(
		"........\n"+
			".....k..\n"+
			"........\n"+
			".Pp.....\n"+
			"........\n"+
			"...K....\n"+
			"........\n"+
			"........\n", White)
	move := pawn_en_passant{abstract_move{xy{1, 4}, xy{2, 5}}}
	b = move.perform_move(b)
	_, ok := b.board[2][5].piece.(*pawn)
	if !ok {
		t.Error(fmt.Sprintf("expecting pawn, got %s", b.board[2][5]))
	}
	if b.board[2][4].piece != nil {
		t.Error(fmt.Sprintf("expecting empty, got %s", b.board[2][4]))
	}
	if b.en_passant != -2 {
		t.Error(fmt.Sprintf("expecting en-passant to be cleared, got %d",
			b.en_passant))
	}
}
