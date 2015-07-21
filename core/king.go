package core

type king struct {
}

func (king) list_moves(board *Board, point xy) moves {
	return list_moves_common(board, point,
		[]xy{xy{-1, -1}, xy{-1, 0}, xy{-1, 1}, xy{0, -1}, xy{0, 1}, xy{1, -1}, xy{1, 0}, xy{1, 1}}, false)
	// TODO: handle castling
}

func (king) can_capture_king(board *Board, point xy) bool {
	return can_capture_king_common(board, point,
		[]xy{xy{-1, -1}, xy{-1, 0}, xy{-1, 1}, xy{0, -1}, xy{0, 1}, xy{1, -1}, xy{1, 0}, xy{1, 1}}, false)
}

func (king) String() string {
	return "K"
}
