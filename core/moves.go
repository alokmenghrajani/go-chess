package core

import (
	"fmt"
	"strings"
)

// We could have one struct with enough information to represent the different
// types of moves, but it felt cleaner to use multiple types.
type move interface {
	perform_move(board *Board) *Board
	String() string
}

type xy struct {
	x, y int
}

type abstract_move struct {
	from, to xy
}

type normal_move struct {
	abstract_move
}

/* Used for keeping track of en-passant */
type pawn_push_two struct {
	abstract_move
}

type pawn_promote struct {
	abstract_move
	piece piece
}

type pawn_en_passant struct {
	abstract_move
}

type king_castle struct {
	king abstract_move
	rook abstract_move
}

type moves []move

func (m normal_move) perform_move(board *Board) *Board {
	b := *board // make a copy of the board
	b.board[m.to.x][m.to.y] = b.board[m.from.x][m.from.y]
	b.board[m.from.x][m.from.y] = square{nil, White}
	b.en_passant = -2
	b.to_play = flip_color(b.to_play)
	return &b
}

func (m pawn_push_two) perform_move(board *Board) *Board {
	b := *board // make a copy of the board
	b.board[m.to.x][m.to.y] = b.board[m.from.x][m.from.y]
	b.board[m.from.x][m.from.y] = square{nil, White}
	assert(m.to.x == m.from.x, "pawn push_push_to not in straight line")
	b.en_passant = m.to.x
	b.to_play = flip_color(b.to_play)
	return &b
}

func (m pawn_promote) perform_move(board *Board) *Board {
	b := *board // make a copy of the board
	b.board[m.to.x][m.to.y] = b.board[m.from.x][m.from.y]
	b.board[m.to.x][m.to.y].piece = m.piece
	b.board[m.from.x][m.from.y] = square{nil, White}
	b.en_passant = -2
	b.to_play = flip_color(b.to_play)
	return &b
}

func (m pawn_en_passant) perform_move(board *Board) *Board {
	b := *board // make a copy of the board
	b.board[m.to.x][m.to.y] = b.board[m.from.x][m.from.y]
	b.board[m.from.x][m.from.y] = square{nil, White}
	b.board[m.to.x][m.from.y] = square{nil, White}
	b.en_passant = -2
	b.to_play = flip_color(b.to_play)
	return &b
}

func is_empty(board *Board, point xy) bool {
	return board.board[point.x][point.y].piece == nil
}

func is_opponent(board *Board, point xy) bool {
	return !is_empty(board, point) &&
		board.board[point.x][point.y].color != board.to_play
}

func is_in_bounds(x int, y int) bool {
	return x >= 0 && x < 8 && y >= 0 && y < 8
}

func is_opponent_king(board *Board, x int, y int) bool {
	if !is_in_bounds(x, y) {
		return false
	}
	if !is_opponent(board, xy{x, y}) {
		return false
	}
	_, ok := board.board[x][y].piece.(*king)
	return ok
}

func append_if_not_in_check(board *Board, moves moves, move move) (moves, bool) {
	b := move.perform_move(board)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			point := xy{i, j}
			if !is_empty(b, point) && !is_opponent(b, point) {
				if b.board[i][j].piece.can_capture_king(b, point) {
					return moves, false
				}
			}
		}
	}
	return append(moves, move), true
}

/**
 * Handles piece movement for all pieces except pawns
 */
func list_moves_common(board *Board, point xy, offsets []xy, repeat bool) moves {
	r := make(moves, 0, 30)
	for _, offset := range offsets {
		r = list_moves_direction(r, board, point, offset, repeat)
	}
	return r
}

func list_moves_direction(r moves, board *Board, point xy, offset xy, repeat bool) moves {
	for i := 1; i < 8; i++ {
		to := xy{point.x + offset.x*i, point.y + offset.y*i}
		if !is_in_bounds(to.x, to.y) {
			return r
		}
		if is_empty(board, to) || is_opponent(board, to) {
			// we are moving into an empty square or capturing a piece
			var ok bool
			r, ok = append_if_not_in_check(board, r,
				normal_move{abstract_move{point, to}})
			if !ok {
				// moving this piece caused a check, so we are done.
				return r
			}
		}
		if !repeat || !is_empty(board, to) {
			// we are done
			return r
		}
	}
	return r
}

func can_capture_king_common(board *Board, point xy, offsets []xy, repeat bool) bool {
	for _, offset := range offsets {
		if can_capture_king_direction(board, point, offset, repeat) {
			return true
		}
	}
	return false
}

func can_capture_king_direction(board *Board, point xy, offset xy, repeat bool) bool {
	for i := 1; i < 8; i++ {
		to := xy{point.x + offset.x*i, point.y + offset.y*i}
		if !is_in_bounds(to.x, to.y) {
			return false
		}
		if is_opponent_king(board, to.x, to.y) {
			return true
		}
		if !repeat || !is_empty(board, to) {
			// we are done
			return false
		}
	}
	return false
}

func flip_color(c color) color {
	if c == White {
		return Black
	} else {
		return White
	}
}

func (m normal_move) String() string {
	return fmt.Sprintf("%c%d→%c%d", 'a'+m.from.x, m.from.y+1, 'a'+m.to.x,
		m.to.y+1)
}

func (m pawn_push_two) String() string {
	return fmt.Sprintf("%c%d→%c%d", 'a'+m.from.x, m.from.y+1, 'a'+m.to.x,
		m.to.y+1)
}

func (m pawn_promote) String() string {
	return fmt.Sprintf("%c%d→%c%d%s", 'a'+m.from.x, m.from.y+1, 'a'+m.to.x,
		m.to.y+1, m.piece.String())
}

func (m pawn_en_passant) String() string {
	return fmt.Sprintf("%c%d→%c%de.p.", 'a'+m.from.x, m.from.y+1, 'a'+m.to.x,
		m.to.y+1)
}

func (moves moves) String() string {
	r := make([]string, len(moves))
	for k, m := range moves {
		r[k] = m.String()
	}
	return strings.Join(r, ",")
}
