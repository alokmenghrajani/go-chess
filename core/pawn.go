package core

//  "fmt"

type pawn struct {
}

func (pawn) list_moves(board *Board, point xy) moves {
	dir, start, end := pawn_get_dir(board)
	r := make(moves, 0, 4)
	to := xy{x: point.x, y: point.y + dir}
	if is_empty(board, to) {
		// try to move forward
		var ok bool
		r, ok = pawn_append_move(board, r, normal_move{abstract_move{point, to}}, end)
		if ok && point.y == start {
			// try to move two squares forward
			to = xy{point.x, point.y + dir + dir}
			m := abstract_move{point, to}
			r, _ = append_if_not_in_check(board, r, pawn_push_two{m})
		}
	}
	if point.x >= 1 {
		to = xy{point.x - 1, point.y + dir}
		if is_opponent(board, to) {
			// try to capture
			r, _ = pawn_append_move(board, r, normal_move{abstract_move{point, to}},
				end)
		}
	}
	if point.x < 7 {
		to = xy{point.x + 1, point.y + dir}
		if is_opponent(board, to) {
			// try to capture
			r, _ = pawn_append_move(board, r, normal_move{abstract_move{point, to}},
				end)
		}
	}
	if point.y == end-3*dir {
		if board.en_passant == point.x+1 {
			// en-passant
			_, ok := board.board[point.x+1][point.y].piece.(*pawn)
			assert(ok, "check opponent pawn")
			assert(is_opponent(board, xy{point.x + 1, point.y}),
				"check opponent pawn")
			assert(board.board[point.x+1][point.y+dir].piece == nil,
				"check empty square")
			to = xy{point.x + 1, point.y + dir}
			m := abstract_move{point, to}
			r, _ = append_if_not_in_check(board, r, pawn_en_passant{m})
		} else if board.en_passant == point.x-1 {
			// en-passant
			_, ok := board.board[point.x-1][point.y].piece.(*pawn)
			assert(ok, "check opponent pawn")
			assert(is_opponent(board, xy{point.x - 1, point.y}),
				"check opponent pawn")
			assert(board.board[point.x-1][point.y+dir].piece == nil,
				"check empty square")
			to = xy{point.x - 1, point.y + dir}
			m := abstract_move{point, to}
			r, _ = append_if_not_in_check(board, r, pawn_en_passant{m})
		}
	}
	return r
}

func (pawn) can_capture_king(board *Board, point xy) bool {
	dir, _, _ := pawn_get_dir(board)
	return is_opponent_king(board, point.x-1, point.y+dir) ||
		is_opponent_king(board, point.x+1, point.y+dir)
}

func pawn_get_dir(board *Board) (int, int, int) {
	if board.to_play == White {
		return 1, 1, 7
	} else {
		return -1, 6, 0
	}
}

/**
 * Validates if move causes check or not.
 * Then handles the case where pawn gets promoted.
 */
func pawn_append_move(board *Board, r moves, m normal_move, end int) (moves, bool) {
	if m.to.y == end {
		_, ok := append_if_not_in_check(board, r, m)
		if ok {
			// we are promoting
			abstract_move := abstract_move{m.from, m.to}
			r = append(r, pawn_promote{abstract_move: abstract_move, piece: new(queen)})
			r = append(r, pawn_promote{abstract_move: abstract_move, piece: new(rook)})
			r = append(r, pawn_promote{abstract_move: abstract_move, piece: new(bishop)})
			r = append(r, pawn_promote{abstract_move: abstract_move, piece: new(knight)})
		}
		return r, ok
	} else {
		return append_if_not_in_check(board, r, m)
	}
}

func (pawn) String() string {
	return "P"
}
