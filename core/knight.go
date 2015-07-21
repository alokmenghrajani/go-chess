package core

type knight struct {
}

func (knight) list_moves(board *Board, point xy) moves {
	return list_moves_common(board, point,
		[]xy{xy{-2, -1}, xy{-2, 1}, xy{-1, -2}, xy{-1, 2}, xy{1, -2}, xy{1, 2}, xy{2, -1}, xy{2, 1}}, false)
}

func (knight) can_capture_king(board *Board, point xy) bool {
	return can_capture_king_common(board, point,
		[]xy{xy{-2, -1}, xy{-2, 1}, xy{-1, -2}, xy{-1, 2}, xy{1, -2}, xy{1, 2}, xy{2, -1}, xy{2, 1}}, false)
}

func (knight) String() string {
	return "N"
}
