package core

type rook struct {
}

func (rook) list_moves(board *Board, point xy) Moves {
	return list_moves_common(board, point,
		[]xy{xy{-1, 0}, xy{0, -1}, xy{0, 1}, xy{1, 0}}, true)
}

func (rook) can_capture_king(board *Board, point xy) bool {
	return can_capture_king_common(board, point,
		[]xy{xy{-1, 0}, xy{0, -1}, xy{0, 1}, xy{1, 0}}, true)
}

func (rook) String() string {
	return "R"
}
