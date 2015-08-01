package core

import (
	"fmt"
)

type king struct {
}

func (king) list_moves(board *Board, point xy) Moves {
	moves := list_moves_common(board, point,
		[]xy{xy{-1, -1}, xy{-1, 0}, xy{-1, 1}, xy{0, -1}, xy{0, 1}, xy{1, -1}, xy{1, 0}, xy{1, 1}}, false)
	// Castling rules:
	// - king is not currently in check
	// - king and rook have never moved
	// - spaces between king and rook are empty
	// - king doesn't cross over a square where it would be in check
	//
	// note: I don't explicitly check that the rook is owned by the same color. An
	// interesting edge case (pawn takes rook which hasn't moved and promotes to
	// rook) doesn't need to be handled because the king would be in check.
	dir, first_rank, _ := get_dir(board)
	first_rank = first_rank - dir
	if board.has_moved_king[board.to_play] {
		// king has moved
		return moves
	}
	if Is_in_check(board) {
		// king is in check
		return moves
	}
	assert(point.x == 4 && point.y == first_rank,
		fmt.Sprintf("king isn't in expected position: %s %d", point, first_rank))
	if king_can_castle_short(board, first_rank) {
		// Castling on the 'h' file
		moves = append(moves, king_castle{king: abstract_move{point, xy{6, first_rank}},
			rook: abstract_move{xy{7, first_rank}, xy{5, first_rank}}})
	}
	if king_can_castle_long(board, first_rank) {
		// Castling on the 'a' file
		moves = append(moves, king_castle{king: abstract_move{point, xy{2, first_rank}},
			rook: abstract_move{xy{0, first_rank}, xy{3, first_rank}}})
	}
	return moves
}

func king_can_castle_short(board *Board, first_rank int) bool {
	if board.has_moved_rook[board.to_play][1] {
		// rook has moved (short)
		return false
	}
	// check that every cell between king and rook is empty
	for _, offset := range []int{5, 6} {
		to := xy{offset, first_rank}
		if !is_empty(board, to) {
			// cell isn't empty (short)
			return false
		}
	}
	// check that king won't move into check
	for _, offset := range []int{5, 6} {
		from := xy{4, first_rank}
		to := xy{offset, first_rank}
		move := normal_move{abstract_move{from, to}}
		r := make(Moves, 0, 1)
		_, ok := append_if_not_in_check(board, r, move)
		if !ok {
			return false
		}
	}
	assert(is_rook(board, xy{7, first_rank}), "expecting rook")
	assert(board.board[7][first_rank].color == board.to_play,
		"expecting same color")
	return true
}

func king_can_castle_long(board *Board, first_rank int) bool {
	if board.has_moved_rook[board.to_play][0] {
		return false
	}
	// check that every cell between king and rook is empty
	for _, offset := range []int{3, 2, 1} {
		to := xy{offset, first_rank}
		if !is_empty(board, to) {
			fmt.Printf("cell %s isn't empty (long)\n", to)
			return false
		}
	}
	// check that king won't move into check
	for _, offset := range []int{3, 2} {
		from := xy{4, first_rank}
		to := xy{offset, first_rank}
		move := normal_move{abstract_move{from, to}}
		r := make(Moves, 0, 1)
		_, ok := append_if_not_in_check(board, r, move)
		if !ok {
			return false
		}
	}
	assert(is_rook(board, xy{0, first_rank}), "expecting rook")
	assert(board.board[0][first_rank].color == board.to_play,
		"expecting same color")
	return true
}

func (king) can_capture_king(board *Board, point xy) bool {
	return can_capture_king_common(board, point,
		[]xy{xy{-1, -1}, xy{-1, 0}, xy{-1, 1}, xy{0, -1}, xy{0, 1}, xy{1, -1}, xy{1, 0}, xy{1, 1}}, false)
}

func (king) String() string {
	return "K"
}
