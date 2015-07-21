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
 *
 * Note: when building the list of moves for a piece, if a piece can't move in
 * a given direction (because it causes the king to be in check), an
 * optimization is to stop further searches in that direction.
 */

type piece interface {
	list_moves(*Board, xy) moves
	does_capture_king(*Board, xy) bool
	String() string
}
