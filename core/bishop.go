package core

type bishop struct {
}

func (bishop) list_moves(board *Board, point xy) Moves {
	return list_moves_common(board, point,
		[]xy{xy{-1, -1}, xy{-1, 1}, xy{1, -1}, xy{1, 1}}, true)
}

func (bishop) can_capture_king(board *Board, point xy) bool {
	return can_capture_king_common(board, point,
		[]xy{xy{-1, -1}, xy{-1, 1}, xy{1, -1}, xy{1, 1}}, true)
}

func (bishop) String() string {
	return "B"
}
