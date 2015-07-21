package core

type rook struct {
}

func (rook) list_moves(board *Board, point xy) moves {
	return list_moves_common(board, point,
		[]xy{xy{-1, 0}, xy{0, -1}, xy{0, 1}, xy{1, 0}}, true)
}

func (rook) does_capture_king(*Board, xy) bool {
	return false
}

func (rook) String() string {
	return "R"
}
