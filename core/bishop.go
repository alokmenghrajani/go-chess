package core

type bishop struct {
}

func (bishop) list_moves(board *Board, point xy) moves {
	return list_moves_common(board, point,
		[]xy{xy{-1, -1}, xy{-1, 1}, xy{1, -1}, xy{1, 1}}, true)
}

func (bishop) does_capture_king(*Board, xy) bool {
	return false
}

func (bishop) String() string {
	return "B"
}
