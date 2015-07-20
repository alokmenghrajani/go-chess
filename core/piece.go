package core

/**
 * To build the list of valid moves, we need to proceed in two steps (because
 * the king is not allowed to move into check). We could:
 * 1. for each piece, build the list of potential moves
 * 2. from the opponents point of view, build the list of moves again. Check
 *    if any piece can capture the king.
 *
 * This can save on number of lines of code, but leads to hard to read to code,
 * in-optimal code or a combination of both.
 *
 * I therefore used an alternative approach which involves two methods:
 * 1. for each piece, build the list of moves.
 * 2. for each opponent piece, check if it can specifically capture the king.
 *    (this part can also be optimized with a reverse-check: what pieces can
 *    reach the king).
 */

type piece interface {
	list_moves(*Board, point) moves
	does_capture_king(*Board, point) bool
	String() string
}

type pawn struct {
}
type rook struct {
}
type knight struct {
}
type bishop struct {
}
type queen struct {
}
type king struct {
}

func (pawn) String() string {
	return "p"
}
func (rook) String() string {
	return "r"
}
func (knight) String() string {
	return "n"
}
func (bishop) String() string {
	return "b"
}
func (queen) String() string {
	return "q"
}
func (king) String() string {
	return "k"
}
